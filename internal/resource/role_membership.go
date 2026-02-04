package resource

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
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

func (o RoleMembershipOps) BuildUpdateRequest(_ context.Context, _, _ RoleMembershipTFModel) (httpclient.RoleMembershipRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	diags.AddError(
		"Update not supported",
		"Role memberships do not support updates. This is a bug if you see this error.",
	)
	return httpclient.RoleMembershipRequest{}, diags
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

func (o RoleMembershipOps) Update(ctx context.Context, client httpclient.PosthogClient, _ RoleMembershipTFModel, _ httpclient.RoleMembershipRequest) (httpclient.RoleMembership, httpclient.HTTPStatusCode, error) {
	return httpclient.RoleMembership{}, 0, fmt.Errorf("role memberships do not support updates; this is a bug if you see this error")
}

func (o RoleMembershipOps) Delete(ctx context.Context, client httpclient.PosthogClient, model RoleMembershipTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteRoleMembership(ctx, model.GetEffectiveOrganizationID(), model.GetRoleID(), model.GetID())
}
