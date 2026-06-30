package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

func NewProjectSettings() resource.Resource {
	return core.NewGenericResource[ProjectSettingsModel, httpclient.EnvironmentSettingsRequest, httpclient.Environment](
		ProjectSettingsOps{},
		core.ProjectSingletonImportParser[ProjectSettingsModel](),
	)
}

// ProjectSettingsModel manages environment-level settings for a PostHog project.
// This is a singleton per project - there is only one set of settings per project.
type ProjectSettingsModel struct {
	core.BaseProjectID
	// ID is the project/environment ID since there is only one settings object per project.
	ID                         types.String `tfsdk:"id"`
	HeatmapsOptIn              types.Bool   `tfsdk:"heatmaps_opt_in"`
	AutocaptureExceptionsOptIn types.Bool   `tfsdk:"autocapture_exceptions_opt_in"`
	SessionRecordingOptIn      types.Bool   `tfsdk:"session_recording_opt_in"`
	SurveysOptIn               types.Bool   `tfsdk:"surveys_opt_in"`
	CookielessServerHashMode   types.Int64  `tfsdk:"cookieless_server_hash_mode"`
	AutocaptureWebVitalsOptIn  types.Bool   `tfsdk:"autocapture_web_vitals_opt_in"`
}

func (m ProjectSettingsModel) GetID() string {
	return m.ID.ValueString()
}

func (m ProjectSettingsModel) HasValidID() bool {
	return m.GetEffectiveProjectID() != ""
}

type ProjectSettingsOps struct{}

func (o ProjectSettingsOps) ResourceName() string {
	return "project_settings"
}

func (o ProjectSettingsOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: `Manages environment-level settings for a PostHog project (environment).

These settings live on the PostHog environment object (` + "`/api/environments/{id}/`" + `), which is distinct from the organization-scoped ` + "`posthog_project`" + ` resource. Each setting is optional - omitted attributes are left at whatever value PostHog currently has, and PostHog's server-side default is read back into state.

~> **Singleton:** There is only one settings object per project. This resource is a singleton - creating multiple instances for the same project will cause conflicts.

~> **Destroy Behavior:** Destroying this resource is a no-op. It stops Terraform from managing the settings but does **not** reset any values on PostHog; the last-applied settings remain in effect.

~> **Plan-gated settings:** If PostHog accepts the update but silently ignores a setting (for example a feature that is not enabled for your plan), Terraform reports a generic "Provider produced inconsistent result after apply" error for that attribute. This usually means the setting cannot be toggled for your project rather than a provider bug.

~> **Out of scope:** Managing the web-analytics capture-domain allowlist ("add domains for web analytics") is intentionally not supported here. That allowlist is not writable via the Personal API Key, so there is no clean API field to manage. Follow up if the API exposes it.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The project (environment) ID, used as the resource identifier.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"heatmaps_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether heatmaps are enabled for web (the PostHog toolbar heatmap feature).",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"autocapture_exceptions_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether exception autocapture is enabled.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"session_recording_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether session recording (recording user sessions) is enabled.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"surveys_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether surveys are enabled.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"cookieless_server_hash_mode": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The cookieless server hash mode: `0` (disabled), `1` (stateless), or `2` (stateful). Matches PostHog's `CookielessServerHashMode` enum.",
				Validators: []validator.Int64{
					int64validator.OneOf(0, 1, 2),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"autocapture_web_vitals_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether web vitals autocapture is enabled.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o ProjectSettingsOps) BuildCreateRequest(_ context.Context, model ProjectSettingsModel) (httpclient.EnvironmentSettingsRequest, diag.Diagnostics) {
	req := httpclient.EnvironmentSettingsRequest{
		HeatmapsOptIn:              util.BoolPtrFromValue(model.HeatmapsOptIn),
		AutocaptureExceptionsOptIn: util.BoolPtrFromValue(model.AutocaptureExceptionsOptIn),
		SessionRecordingOptIn:      util.BoolPtrFromValue(model.SessionRecordingOptIn),
		SurveysOptIn:               util.BoolPtrFromValue(model.SurveysOptIn),
		AutocaptureWebVitalsOptIn:  util.BoolPtrFromValue(model.AutocaptureWebVitalsOptIn),
		CookielessServerHashMode:   util.Int64PtrFromValue(model.CookielessServerHashMode),
	}
	return req, nil
}

func (o ProjectSettingsOps) BuildUpdateRequest(ctx context.Context, plan, _ ProjectSettingsModel) (httpclient.EnvironmentSettingsRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o ProjectSettingsOps) MapResponseToModel(_ context.Context, resp httpclient.Environment, model *ProjectSettingsModel) diag.Diagnostics {
	model.ID = types.StringValue(model.GetEffectiveProjectID())
	model.HeatmapsOptIn = core.PtrToBool(resp.HeatmapsOptIn)
	model.AutocaptureExceptionsOptIn = core.PtrToBool(resp.AutocaptureExceptionsOptIn)
	model.SessionRecordingOptIn = core.PtrToBool(resp.SessionRecordingOptIn)
	model.SurveysOptIn = core.PtrToBool(resp.SurveysOptIn)
	model.AutocaptureWebVitalsOptIn = core.PtrToBool(resp.AutocaptureWebVitalsOptIn)
	model.CookielessServerHashMode = util.PtrToInt64(resp.CookielessServerHashMode)
	return nil
}

func (o ProjectSettingsOps) Create(ctx context.Context, client httpclient.PosthogClient, model ProjectSettingsModel, req httpclient.EnvironmentSettingsRequest) (httpclient.Environment, error) {
	result, _, err := client.UpdateEnvironment(ctx, model.GetEffectiveProjectID(), req)
	return result, err
}

func (o ProjectSettingsOps) Read(ctx context.Context, client httpclient.PosthogClient, model ProjectSettingsModel) (httpclient.Environment, httpclient.HTTPStatusCode, error) {
	return client.GetEnvironment(ctx, model.GetEffectiveProjectID())
}

func (o ProjectSettingsOps) Update(ctx context.Context, client httpclient.PosthogClient, model ProjectSettingsModel, req httpclient.EnvironmentSettingsRequest) (httpclient.Environment, httpclient.HTTPStatusCode, error) {
	return client.UpdateEnvironment(ctx, model.GetEffectiveProjectID(), req)
}

func (o ProjectSettingsOps) Delete(_ context.Context, _ httpclient.PosthogClient, _ ProjectSettingsModel) (httpclient.HTTPStatusCode, error) {
	// Settings live on the environment, which is never destroyed by this resource.
	// Deleting simply stops Terraform from managing the settings; PostHog keeps the
	// last-applied values. This is intentionally a no-op.
	return 0, nil
}
