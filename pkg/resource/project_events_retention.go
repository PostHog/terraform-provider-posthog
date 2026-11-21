package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
)

func NewProjectEventsRetention() resource.Resource {
	return core.NewGenericResource[ProjectEventsRetentionModel, httpclient.ProjectEventsRetentionRequest, httpclient.ProjectEventsRetention](ProjectEventsRetentionOps{}, core.ProjectSingletonImportParser[ProjectEventsRetentionModel]())
}

type ProjectEventsRetentionModel struct {
	core.BaseProjectID
	ID              types.String `tfsdk:"id"`
	RetentionMonths types.Int64  `tfsdk:"retention_months"`
}

func (m ProjectEventsRetentionModel) GetID() string    { return m.ID.ValueString() }
func (m ProjectEventsRetentionModel) HasValidID() bool { return m.GetEffectiveProjectID() != "" }

type ProjectEventsRetentionOps struct{}

func (o ProjectEventsRetentionOps) ResourceName() string { return "project_events_retention" }
func (o ProjectEventsRetentionOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages the event retention period for a PostHog project. Only organization admins can manage this setting. There is one retention setting per project. Omitting `retention_months` makes the project inherit the organization's period, which may permanently delete older events if that period is shorter. Destroying the resource stops Terraform from managing the setting and leaves the override in place.",
		Attributes: map[string]schema.Attribute{
			"id":               schema.StringAttribute{Computed: true, MarkdownDescription: "The project ID (used as the resource identifier).", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"project_id":       core.ProjectIDSchemaAttribute(),
			"retention_months": schema.Int64Attribute{Optional: true, MarkdownDescription: "The project's event retention period in months. Omit to inherit the organization's period. Available values depend on the PostHog plan; see `retention_months_choices` in the [Events retention API reference](https://posthog.com/docs/api/events-retention)."},
		},
	}
}
func (o ProjectEventsRetentionOps) BuildCreateRequest(_ context.Context, model ProjectEventsRetentionModel) (httpclient.ProjectEventsRetentionRequest, diag.Diagnostics) {
	return httpclient.ProjectEventsRetentionRequest{RetentionMonths: util.Int64PtrFromValue(model.RetentionMonths)}, nil
}
func (o ProjectEventsRetentionOps) BuildUpdateRequest(ctx context.Context, plan, _ ProjectEventsRetentionModel) (httpclient.ProjectEventsRetentionRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}
func (o ProjectEventsRetentionOps) MapResponseToModel(_ context.Context, resp httpclient.ProjectEventsRetention, model *ProjectEventsRetentionModel) diag.Diagnostics {
	model.ID = types.StringValue(model.GetEffectiveProjectID())
	model.RetentionMonths = util.PtrToInt64(resp.ProjectRetentionMonths)
	return nil
}
func (o ProjectEventsRetentionOps) Create(ctx context.Context, client httpclient.PosthogClient, model ProjectEventsRetentionModel, req httpclient.ProjectEventsRetentionRequest) (httpclient.ProjectEventsRetention, error) {
	result, _, err := client.SetProjectEventsRetention(ctx, model.GetEffectiveProjectID(), req)
	return result, err
}
func (o ProjectEventsRetentionOps) Read(ctx context.Context, client httpclient.PosthogClient, model ProjectEventsRetentionModel) (httpclient.ProjectEventsRetention, httpclient.HTTPStatusCode, error) {
	return client.GetProjectEventsRetention(ctx, model.GetEffectiveProjectID())
}
func (o ProjectEventsRetentionOps) Update(ctx context.Context, client httpclient.PosthogClient, model ProjectEventsRetentionModel, req httpclient.ProjectEventsRetentionRequest) (httpclient.ProjectEventsRetention, httpclient.HTTPStatusCode, error) {
	return client.SetProjectEventsRetention(ctx, model.GetEffectiveProjectID(), req)
}
func (o ProjectEventsRetentionOps) Delete(ctx context.Context, _ httpclient.PosthogClient, model ProjectEventsRetentionModel) (httpclient.HTTPStatusCode, error) {
	tflog.Warn(ctx, "posthog_project_events_retention destroy stops managing the setting; the override remains on PostHog", map[string]any{
		"project_id": model.GetEffectiveProjectID(),
	})
	return 0, nil
}
