package datasource

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewUser() datasource.DataSource {
	return &UserDataSource{}
}

type UserDataSource struct {
	client                httpclient.PosthogClient
	defaultOrganizationID string
}

type UserDataSourceModel struct {
	OrganizationID       types.String `tfsdk:"organization_id"`
	Email                types.String `tfsdk:"email"`
	UUID                 types.String `tfsdk:"uuid"`
	OrganizationMemberID types.String `tfsdk:"organization_member_id"`
	FirstName            types.String `tfsdk:"first_name"`
	LastName             types.String `tfsdk:"last_name"`
	IsEmailVerified      types.Bool   `tfsdk:"is_email_verified"`
}

var (
	_ datasource.DataSource              = (*UserDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*UserDataSource)(nil)
)

func (d *UserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_user", req.ProviderTypeName)
}

func (d *UserDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Look up a user by email within an organization. The user must be a member of the organization.",
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The organization ID to search for the user in. Defaults to the provider-level organization_id.",
			},
			"email": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The email address of the user to look up.",
			},
			"uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The UUID of the user.",
			},
			"organization_member_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The organization membership ID. Use this for `organization_member` in access controls.",
			},
			"first_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The first name of the user.",
			},
			"last_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The last name of the user.",
			},
			"is_email_verified": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the user's email address has been verified.",
			},
		},
	}
}

func (d *UserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(data.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = providerData.Client
	d.defaultOrganizationID = providerData.DefaultOrganizationID
}

func (d *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config UserDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use provider default if organization_id not explicitly set
	orgID := config.OrganizationID.ValueString()
	if config.OrganizationID.IsNull() || config.OrganizationID.IsUnknown() || orgID == "" {
		orgID = d.defaultOrganizationID
		config.OrganizationID = types.StringValue(orgID)
	}

	if orgID == "" {
		resp.Diagnostics.AddError(
			"Missing organization_id",
			"organization_id must be set either in the data source or at the provider level",
		)
		return
	}

	email := strings.ToLower(strings.TrimSpace(config.Email.ValueString()))

	tflog.Debug(ctx, "Looking up user by email", map[string]any{
		"organization_id": orgID,
		"email":           email,
	})

	members, err := d.client.ListOrganizationMembers(ctx, orgID)
	if err != nil {
		resp.Diagnostics.AddError("Error listing organization members", err.Error())
		return
	}

	var found *httpclient.OrganizationMember
	for i := range members {
		member := &members[i]
		if member.User != nil && member.User.Email != nil {
			memberEmail := strings.ToLower(strings.TrimSpace(*member.User.Email))
			if memberEmail == email {
				found = member
				break
			}
		}
	}

	if found == nil {
		resp.Diagnostics.AddError(
			"User not found",
			fmt.Sprintf("No user with email %q found in organization %s", email, orgID),
		)
		return
	}

	config.UUID = types.StringValue(found.User.UUID)
	config.OrganizationMemberID = types.StringValue(found.ID)
	config.FirstName = core.PtrToStringNullIfEmptyTrimmed(found.User.FirstName)
	config.LastName = core.PtrToStringNullIfEmptyTrimmed(found.User.LastName)
	config.IsEmailVerified = core.PtrToBool(found.User.IsEmailVerified)

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
