package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewProject() resource.Resource {
	return core.NewGenericResource[ProjectTFModel, httpclient.ProjectRequest, httpclient.Project](
		ProjectOps{},
		core.OrganizationScopedImportParser[ProjectTFModel](),
	)
}

type ProjectTFModel struct {
	core.BaseInt64Identifiable
	core.BaseOrganizationID
	Name     types.String `tfsdk:"name"`
	APIToken types.String `tfsdk:"api_token"`
	Timezone types.String `tfsdk:"timezone"`
}

type ProjectOps struct{}

func (o ProjectOps) ResourceName() string {
	return "project"
}

func (o ProjectOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a PostHog project within an organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The numeric identifier of the project.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the project.",
				Required:            true,
			},
			"organization_id": core.OrganizationIDSchemaAttribute(),
			"api_token": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The API token for this project. This is used to send events to PostHog.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"timezone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The timezone for this project (e.g., 'UTC', 'America/New_York', 'Europe/London'). Defaults to 'UTC'.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o ProjectOps) BuildCreateRequest(ctx context.Context, model ProjectTFModel) (httpclient.ProjectRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := model.Name.ValueString()
	req := httpclient.ProjectRequest{
		Name: &name,
	}

	if !model.Timezone.IsNull() && !model.Timezone.IsUnknown() {
		tz := model.Timezone.ValueString()
		req.Timezone = &tz
	}

	return req, diags
}

func (o ProjectOps) BuildUpdateRequest(ctx context.Context, plan, state ProjectTFModel) (httpclient.ProjectRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o ProjectOps) MapResponseToModel(ctx context.Context, resp httpclient.Project, model *ProjectTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.APIToken = core.PtrToStringNullIfEmptyTrimmed(resp.APIToken)
	model.Timezone = core.PtrToStringNullIfEmptyTrimmed(resp.Timezone)

	return diags
}

func (o ProjectOps) Create(ctx context.Context, client httpclient.PosthogClient, model ProjectTFModel, req httpclient.ProjectRequest) (httpclient.Project, error) {
	return client.CreateProject(ctx, model.GetEffectiveOrganizationID(), req)
}

func (o ProjectOps) Read(ctx context.Context, client httpclient.PosthogClient, model ProjectTFModel) (httpclient.Project, httpclient.HTTPStatusCode, error) {
	return client.GetProject(ctx, model.GetEffectiveOrganizationID(), model.GetID())
}

func (o ProjectOps) Update(ctx context.Context, client httpclient.PosthogClient, model ProjectTFModel, req httpclient.ProjectRequest) (httpclient.Project, httpclient.HTTPStatusCode, error) {
	return client.UpdateProject(ctx, model.GetEffectiveOrganizationID(), model.GetID(), req)
}

func (o ProjectOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ProjectTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteProject(ctx, model.GetEffectiveOrganizationID(), model.GetID())
}
