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

func NewRole() datasource.DataSource {
	return &RoleDataSource{}
}

type RoleDataSource struct {
	client                httpclient.PosthogClient
	defaultOrganizationID string
}

type RoleDataSourceModel struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	Name           types.String `tfsdk:"name"`
	ID             types.String `tfsdk:"id"`
	CreatedAt      types.String `tfsdk:"created_at"`
	CreatedBy      types.String `tfsdk:"created_by"`
}

var (
	_ datasource.DataSource              = (*RoleDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*RoleDataSource)(nil)
)

func (d *RoleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_role", req.ProviderTypeName)
}

func (d *RoleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Look up a role by name within an organization." + core.EnterpriseRBACNote,
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The organization ID to search for the role in. Defaults to the provider-level organization_id.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the role to look up. Matching is case-insensitive.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The UUID of the role.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the role was created.",
			},
			"created_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the user who created the role.",
			},
		},
	}
}

func (d *RoleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *RoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RoleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

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

	name := strings.TrimSpace(config.Name.ValueString())

	tflog.Debug(ctx, "Looking up role by name", map[string]any{
		"organization_id": orgID,
		"name":            name,
	})

	roles, err := d.client.ListRoles(ctx, orgID)
	if err != nil {
		resp.Diagnostics.AddError("Error listing roles", err.Error())
		return
	}

	var found *httpclient.Role
	for i := range roles {
		if strings.EqualFold(strings.TrimSpace(roles[i].Name), name) {
			found = &roles[i]
			break
		}
	}

	if found == nil {
		resp.Diagnostics.AddError(
			"Role not found",
			fmt.Sprintf("No role with name %q found in organization %s", name, orgID),
		)
		return
	}

	config.ID = types.StringValue(found.ID)
	config.Name = types.StringValue(found.Name)
	config.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(found.CreatedAt)
	config.CreatedBy = core.PtrToStringNullIfEmptyTrimmed(found.CreatedBy)

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
