package resource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

// MembershipLevel represents an organization membership level.
// See: https://app.posthog.com/api/schema/ (OrganizationMembershipLevel enum)
type MembershipLevel struct {
	Name  string
	Value int
}

var (
	MembershipLevelMember = MembershipLevel{Name: "member", Value: 1}
	MembershipLevelAdmin  = MembershipLevel{Name: "admin", Value: 8}
	MembershipLevelOwner  = MembershipLevel{Name: "owner", Value: 15}

	// membershipLevels maps level names to their definitions
	membershipLevelsByName = map[string]MembershipLevel{
		MembershipLevelMember.Name: MembershipLevelMember,
		MembershipLevelAdmin.Name:  MembershipLevelAdmin,
		MembershipLevelOwner.Name:  MembershipLevelOwner,
	}

	// membershipLevelsByValue maps level values to their definitions
	membershipLevelsByValue = map[int]MembershipLevel{
		MembershipLevelMember.Value: MembershipLevelMember,
		MembershipLevelAdmin.Value:  MembershipLevelAdmin,
		MembershipLevelOwner.Value:  MembershipLevelOwner,
	}

	// membershipLevelNames is used for schema validation
	membershipLevelNames = []string{
		MembershipLevelMember.Name,
		MembershipLevelAdmin.Name,
		MembershipLevelOwner.Name,
	}
)

func levelFromString(name string) *int {
	if level, ok := membershipLevelsByName[name]; ok {
		return &level.Value
	}
	return &MembershipLevelMember.Value
}

func levelToString(value *int) string {
	if value == nil {
		return MembershipLevelMember.Name
	}
	if level, ok := membershipLevelsByValue[*value]; ok {
		return level.Name
	}
	return MembershipLevelMember.Name
}

func NewOrganizationMember() resource.Resource {
	return core.NewGenericResource[OrganizationMemberTFModel, httpclient.OrganizationMemberRequest, httpclient.OrganizationMember](
		OrganizationMemberOps{},
		core.OrganizationScopedImportParser[OrganizationMemberTFModel](),
	)
}

type OrganizationMemberTFModel struct {
	core.BaseStringIdentifiable
	core.BaseOrganizationID
	UserUUID     types.String `tfsdk:"user_uuid"`
	Level        types.String `tfsdk:"level"`
	Email        types.String `tfsdk:"email"`
	FirstName    types.String `tfsdk:"first_name"`
	LastName     types.String `tfsdk:"last_name"`
	JoinedAt     types.String `tfsdk:"joined_at"`
	Is2FAEnabled types.Bool   `tfsdk:"is_2fa_enabled"`
}

type OrganizationMemberOps struct{}

func (o OrganizationMemberOps) ResourceName() string {
	return "organization_member"
}

func (o OrganizationMemberOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: `Manages a user's membership in a PostHog organization.

~> **Warning:** Destroying this resource will **remove the user from the organization entirely**. Users must already be members of the organization (e.g., via invite) before this resource can manage them.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The member ID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": core.OrganizationIDSchemaAttribute(),
			"user_uuid": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The UUID of the user. Use the posthog_user data source to look up users by email.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"level": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The access level for the member. Valid values are 'member', 'admin', or 'owner'. Defaults to 'member'.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(membershipLevelNames...),
				},
			},
			"email": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The email address of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"first_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The first name of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The last name of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"joined_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the user joined the organization.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_2fa_enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether two-factor authentication is enabled for the user.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o OrganizationMemberOps) BuildCreateRequest(_ context.Context, model OrganizationMemberTFModel) (httpclient.OrganizationMemberRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.OrganizationMemberRequest{}
	if !model.Level.IsNull() && !model.Level.IsUnknown() {
		req.Level = levelFromString(model.Level.ValueString())
	}

	return req, diags
}

func (o OrganizationMemberOps) BuildUpdateRequest(ctx context.Context, plan, _ OrganizationMemberTFModel) (httpclient.OrganizationMemberRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o OrganizationMemberOps) MapResponseToModel(_ context.Context, resp httpclient.OrganizationMember, model *OrganizationMemberTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Level = types.StringValue(levelToString(resp.Level))

	if resp.User != nil {
		model.UserUUID = types.StringValue(resp.User.UUID)
		model.Email = core.PtrToStringNullIfEmptyTrimmed(resp.User.Email)
		model.FirstName = core.PtrToStringNullIfEmptyTrimmed(resp.User.FirstName)
		model.LastName = core.PtrToStringNullIfEmptyTrimmed(resp.User.LastName)
	}

	model.JoinedAt = core.PtrToStringNullIfEmptyTrimmed(resp.JoinedAt)
	model.Is2FAEnabled = core.PtrToBool(resp.Is2FAEnabled)

	return diags
}

// Create "adopts" an existing organization member and optionally updates their level.
// Users must already be members of the organization (e.g., via invite) before this resource can manage them.
func (o OrganizationMemberOps) Create(ctx context.Context, client httpclient.PosthogClient, model OrganizationMemberTFModel, req httpclient.OrganizationMemberRequest) (httpclient.OrganizationMember, error) {
	orgID := model.GetEffectiveOrganizationID()
	userUUID := model.UserUUID.ValueString()

	// First, verify the user is already a member
	member, status, err := client.GetOrganizationMember(ctx, orgID, userUUID)
	if err != nil {
		if status == http.StatusNotFound {
			return httpclient.OrganizationMember{}, fmt.Errorf("user %s is not a member of organization %s; users must be invited to the organization first", userUUID, orgID)
		}
		return httpclient.OrganizationMember{}, err
	}

	// If a level is specified, update it
	if req.Level != nil {
		member, _, err = client.UpdateOrganizationMember(ctx, orgID, userUUID, req)
		if err != nil {
			return httpclient.OrganizationMember{}, fmt.Errorf("error updating organization member level: %w", err)
		}
	}

	return member, nil
}

func (o OrganizationMemberOps) Read(ctx context.Context, client httpclient.PosthogClient, model OrganizationMemberTFModel) (httpclient.OrganizationMember, httpclient.HTTPStatusCode, error) {
	return client.GetOrganizationMember(ctx, model.GetEffectiveOrganizationID(), model.UserUUID.ValueString())
}

func (o OrganizationMemberOps) Update(ctx context.Context, client httpclient.PosthogClient, model OrganizationMemberTFModel, req httpclient.OrganizationMemberRequest) (httpclient.OrganizationMember, httpclient.HTTPStatusCode, error) {
	return client.UpdateOrganizationMember(ctx, model.GetEffectiveOrganizationID(), model.UserUUID.ValueString(), req)
}

func (o OrganizationMemberOps) Delete(ctx context.Context, client httpclient.PosthogClient, model OrganizationMemberTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteOrganizationMember(ctx, model.GetEffectiveOrganizationID(), model.UserUUID.ValueString())
}
