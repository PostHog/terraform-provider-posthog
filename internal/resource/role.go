package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewRole() resource.Resource {
	return core.NewGenericResource[RoleTFModel, httpclient.RoleRequest, httpclient.Role](
		RoleOps{},
		core.OrganizationScopedImportParser[RoleTFModel](),
	)
}

type RoleTFModel struct {
	core.BaseStringIdentifiable
	core.BaseOrganizationID
	Name      types.String `tfsdk:"name"`
	CreatedAt types.String `tfsdk:"created_at"`
	CreatedBy types.String `tfsdk:"created_by"`
}

type RoleOps struct{}

func (o RoleOps) ResourceName() string {
	return "role"
}

func (o RoleOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a role within a PostHog organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The UUID of the role.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": core.OrganizationIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the role.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the role was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_by": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the user who created the role.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o RoleOps) BuildCreateRequest(ctx context.Context, model RoleTFModel) (httpclient.RoleRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.RoleRequest{
		Name: model.Name.ValueString(),
	}

	return req, diags
}

func (o RoleOps) BuildUpdateRequest(ctx context.Context, plan, state RoleTFModel) (httpclient.RoleRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o RoleOps) MapResponseToModel(ctx context.Context, resp httpclient.Role, model *RoleTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.CreatedBy = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedBy)

	return diags
}

func (o RoleOps) Create(ctx context.Context, client httpclient.PosthogClient, model RoleTFModel, req httpclient.RoleRequest) (httpclient.Role, error) {
	return client.CreateRole(ctx, model.GetEffectiveOrganizationID(), req)
}

func (o RoleOps) Read(ctx context.Context, client httpclient.PosthogClient, model RoleTFModel) (httpclient.Role, httpclient.HTTPStatusCode, error) {
	return client.GetRole(ctx, model.GetEffectiveOrganizationID(), model.GetID())
}

func (o RoleOps) Update(ctx context.Context, client httpclient.PosthogClient, model RoleTFModel, req httpclient.RoleRequest) (httpclient.Role, httpclient.HTTPStatusCode, error) {
	return client.UpdateRole(ctx, model.GetEffectiveOrganizationID(), model.GetID(), req)
}

func (o RoleOps) Delete(ctx context.Context, client httpclient.PosthogClient, model RoleTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteRole(ctx, model.GetEffectiveOrganizationID(), model.GetID())
}
