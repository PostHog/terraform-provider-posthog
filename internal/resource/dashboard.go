package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewDashboard() resource.Resource {
	return core.NewGenericResource[DashboardResourceTFModel, httpclient.DashboardRequest, httpclient.Dashboard](DashboardOps{})
}

type DashboardResourceTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Pinned      types.Bool   `tfsdk:"pinned"`
	Tags        types.Set    `tfsdk:"tags"`
	Deleted     types.Bool   `tfsdk:"deleted"`
}

type DashboardOps struct{}

func (o DashboardOps) ResourceName() string {
	return "Dashboard"
}

func (o DashboardOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "PostHog Dashboard resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Dashboard ID",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Dashboard name",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Dashboard description",
			},
			"pinned": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the dashboard is pinned",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Set of tags for the dashboard",
			},
			"deleted": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the dashboard is deleted (soft delete)",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o DashboardOps) BuildCreateRequest(ctx context.Context, model DashboardResourceTFModel) (httpclient.DashboardRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := model.Name.ValueString()
	req := httpclient.DashboardRequest{
		Name: &name,
	}

	if !model.Description.IsNull() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}

	if !model.Pinned.IsNull() {
		pinned := model.Pinned.ValueBool()
		req.Pinned = &pinned
	}

	if !model.Tags.IsNull() {
		tags, d := core.ExtractTags(ctx, model.Tags)
		diags.Append(d...)
		req.Tags = tags
	}

	return req, diags
}

func (o DashboardOps) BuildUpdateRequest(ctx context.Context, plan, state DashboardResourceTFModel) (httpclient.DashboardRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	// Clear description if removed from config
	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = core.StringPtr("")
	}

	if !plan.Deleted.IsNull() {
		deleted := plan.Deleted.ValueBool()
		req.Deleted = &deleted
	}

	return req, diags
}

func (o DashboardOps) MapResponseToModel(ctx context.Context, resp httpclient.Dashboard, model *DashboardResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.Pinned = core.PtrToBool(resp.Pinned)
	model.Deleted = core.PtrToBool(resp.Deleted)

	tagsSet, d := core.TagsToSet(ctx, resp.Tags)
	diags.Append(d...)
	model.Tags = tagsSet

	return diags
}

func (o DashboardOps) Create(ctx context.Context, client httpclient.PosthogClient, model DashboardResourceTFModel, req httpclient.DashboardRequest) (httpclient.Dashboard, error) {
	return client.CreateDashboard(ctx, model.GetEffectiveProjectID(), req)
}

func (o DashboardOps) Read(ctx context.Context, client httpclient.PosthogClient, model DashboardResourceTFModel) (httpclient.Dashboard, httpclient.HTTPStatusCode, error) {
	return client.GetDashboard(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o DashboardOps) Update(ctx context.Context, client httpclient.PosthogClient, model DashboardResourceTFModel, req httpclient.DashboardRequest) (httpclient.Dashboard, httpclient.HTTPStatusCode, error) {
	return client.UpdateDashboard(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o DashboardOps) Delete(ctx context.Context, client httpclient.PosthogClient, model DashboardResourceTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteDashboard(ctx, model.GetEffectiveProjectID(), model.GetID())
}
