package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
)

func NewRoleMembership() resource.Resource {
	return core.NewGenericResource[RoleMembershipTFModel, httpclient.RoleMembershipRequest, httpclient.RoleMembership](
		RoleMembershipOps{},
		core.RoleMembershipImportParser[RoleMembershipTFModel](),
	)
}

type RoleMembershipTFModel struct {
	core.BaseStringIdentifiable
	core.BaseOrganizationID
	core.BaseRoleID
	UserUUID  types.String `tfsdk:"user_uuid"`
	UserEmail types.String `tfsdk:"user_email"`
	JoinedAt  types.String `tfsdk:"joined_at"`
}

type RoleMembershipOps struct{}

func (o RoleMembershipOps) ResourceName() string {
	return "role_membership"
}

func (o RoleMembershipOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a user's membership in a role within a PostHog organization." + core.EnterpriseRBACNote,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The UUID of the role membership.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": core.OrganizationIDSchemaAttribute(),
			"role_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The UUID of the role.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"user_uuid": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The UUID of the user to add to the role.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"user_email": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The email address of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"joined_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the user joined the role.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o RoleMembershipOps) BuildCreateRequest(_ context.Context, model RoleMembershipTFModel) (httpclient.RoleMembershipRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.RoleMembershipRequest{
		UserUUID: model.UserUUID.ValueString(),
	}

	return req, diags
}

// BuildUpdateRequest returns an empty request because Update sends nothing to the API.
func (o RoleMembershipOps) BuildUpdateRequest(_ context.Context, _, _ RoleMembershipTFModel) (httpclient.RoleMembershipRequest, diag.Diagnostics) {
	return httpclient.RoleMembershipRequest{}, nil
}

func (o RoleMembershipOps) MapResponseToModel(_ context.Context, resp httpclient.RoleMembership, model *RoleMembershipTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)

	if resp.RoleID != "" {
		model.RoleID = types.StringValue(resp.RoleID)
	}

	if resp.User != nil {
		model.UserUUID = types.StringValue(resp.User.UUID)
		model.UserEmail = core.PtrToStringNullIfEmptyTrimmed(resp.User.Email)
	}

	model.JoinedAt = core.PtrToStringNullIfEmptyTrimmed(resp.JoinedAt)

	return diags
}

func (o RoleMembershipOps) Create(ctx context.Context, client httpclient.PosthogClient, model RoleMembershipTFModel, req httpclient.RoleMembershipRequest) (httpclient.RoleMembership, error) {
	return client.CreateRoleMembership(ctx, model.GetEffectiveOrganizationID(), model.GetRoleID(), req)
}

func (o RoleMembershipOps) Read(ctx context.Context, client httpclient.PosthogClient, model RoleMembershipTFModel) (httpclient.RoleMembership, httpclient.HTTPStatusCode, error) {
	return client.GetRoleMembership(ctx, model.GetEffectiveOrganizationID(), model.GetRoleID(), model.GetID())
}

// Update re-reads the membership. PostHog cannot update a role membership, and every other
// attribute forces replacement, so the only in-place change is organization_id moving off the
// retired "@current" alias, which needs nothing from the API.
func (o RoleMembershipOps) Update(ctx context.Context, client httpclient.PosthogClient, model RoleMembershipTFModel, _ httpclient.RoleMembershipRequest) (httpclient.RoleMembership, httpclient.HTTPStatusCode, error) {
	return o.Read(ctx, client, model)
}

func (o RoleMembershipOps) Delete(ctx context.Context, client httpclient.PosthogClient, model RoleMembershipTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteRoleMembership(ctx, model.GetEffectiveOrganizationID(), model.GetRoleID(), model.GetID())
}
