package resource

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	AppURLs                    types.List   `tfsdk:"app_urls"`
	RecordingDomains           types.List   `tfsdk:"recording_domains"`
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

~> **Plan-gated settings:** If PostHog accepts the update but silently ignores a setting (for example a feature that is not enabled for your plan), Terraform reports a generic "Provider produced inconsistent result after apply" error. The provider also emits a warning naming the specific attribute(s) PostHog did not apply, so you can tell which setting is unavailable for your project rather than a provider bug.

~> **Domains:** ` + "`app_urls`" + ` is the project's authorized-domains list — shown in PostHog settings as **Web analytics domains** and reused as the toolbar's Authorized URLs (web analytics has no separate allowlist of its own). ` + "`recording_domains`" + ` is the session-replay domains list.`,
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
				MarkdownDescription: "The cookieless server hash mode. Known values: `0` (disabled), `1` (stateless), `2` (stateful) — matching PostHog's `CookielessServerHashMode` enum. PostHog may add modes over time, so any non-negative value is accepted rather than a fixed set; consult the PostHog docs for the current options.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
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
			"app_urls": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The project's authorized domains — shown in PostHog settings as **Web analytics domains** (and used as the toolbar's Authorized URLs). These are the domains tracked in web analytics and where the toolbar is enabled. Maps to the team `app_urls` field. Wildcards are not allowed; order is preserved.",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"recording_domains": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Authorized domains for session replay. Maps to the team `recording_domains` field. Order is preserved.",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o ProjectSettingsOps) BuildCreateRequest(ctx context.Context, model ProjectSettingsModel) (httpclient.EnvironmentSettingsRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	req := httpclient.EnvironmentSettingsRequest{
		HeatmapsOptIn:              util.BoolPtrFromValue(model.HeatmapsOptIn),
		AutocaptureExceptionsOptIn: util.BoolPtrFromValue(model.AutocaptureExceptionsOptIn),
		SessionRecordingOptIn:      util.BoolPtrFromValue(model.SessionRecordingOptIn),
		SurveysOptIn:               util.BoolPtrFromValue(model.SurveysOptIn),
		AutocaptureWebVitalsOptIn:  util.BoolPtrFromValue(model.AutocaptureWebVitalsOptIn),
		CookielessServerHashMode:   util.Int64PtrFromValue(model.CookielessServerHashMode),
	}

	appURLs, d := util.StringListToSlicePtr(ctx, model.AppURLs)
	diags.Append(d...)
	req.AppURLs = appURLs

	recordingDomains, d := util.StringListToSlicePtr(ctx, model.RecordingDomains)
	diags.Append(d...)
	req.RecordingDomains = recordingDomains

	return req, diags
}

func (o ProjectSettingsOps) BuildUpdateRequest(ctx context.Context, plan, _ ProjectSettingsModel) (httpclient.EnvironmentSettingsRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o ProjectSettingsOps) MapResponseToModel(ctx context.Context, resp httpclient.Environment, model *ProjectSettingsModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Before overwriting state, detect any setting the user explicitly configured
	// that PostHog returned with a different value (e.g. a feature not enabled for
	// their plan). Naming those attributes turns the framework's otherwise generic
	// "inconsistent result after apply" error into an actionable message instead of
	// leaving the user to trial-and-error each field. Unconfigured (null/unknown)
	// attributes are skipped, so this only fires on a genuine divergence.
	var ignored []string
	if boolDiverged(model.HeatmapsOptIn, resp.HeatmapsOptIn) {
		ignored = append(ignored, "heatmaps_opt_in")
	}
	if boolDiverged(model.AutocaptureExceptionsOptIn, resp.AutocaptureExceptionsOptIn) {
		ignored = append(ignored, "autocapture_exceptions_opt_in")
	}
	if boolDiverged(model.SessionRecordingOptIn, resp.SessionRecordingOptIn) {
		ignored = append(ignored, "session_recording_opt_in")
	}
	if boolDiverged(model.SurveysOptIn, resp.SurveysOptIn) {
		ignored = append(ignored, "surveys_opt_in")
	}
	if boolDiverged(model.AutocaptureWebVitalsOptIn, resp.AutocaptureWebVitalsOptIn) {
		ignored = append(ignored, "autocapture_web_vitals_opt_in")
	}
	if int64Diverged(model.CookielessServerHashMode, resp.CookielessServerHashMode) {
		ignored = append(ignored, "cookieless_server_hash_mode")
	}
	if len(ignored) > 0 {
		diags.AddWarning(
			"PostHog did not apply some configured settings",
			fmt.Sprintf(
				"PostHog returned a different value for: %s. These settings are likely not available for your project or plan, "+
					"so PostHog silently ignored them. Remove them from your configuration or enable the feature in PostHog.",
				strings.Join(ignored, ", "),
			),
		)
	}

	model.ID = types.StringValue(model.GetEffectiveProjectID())
	model.HeatmapsOptIn = core.PtrToBool(resp.HeatmapsOptIn)
	model.AutocaptureExceptionsOptIn = core.PtrToBool(resp.AutocaptureExceptionsOptIn)
	model.SessionRecordingOptIn = core.PtrToBool(resp.SessionRecordingOptIn)
	model.SurveysOptIn = core.PtrToBool(resp.SurveysOptIn)
	model.AutocaptureWebVitalsOptIn = core.PtrToBool(resp.AutocaptureWebVitalsOptIn)
	model.CookielessServerHashMode = util.PtrToInt64(resp.CookielessServerHashMode)

	appURLs, d := util.StringSlicePtrToList(ctx, resp.AppURLs)
	diags.Append(d...)
	model.AppURLs = appURLs

	recordingDomains, d := util.StringSlicePtrToList(ctx, resp.RecordingDomains)
	diags.Append(d...)
	model.RecordingDomains = recordingDomains

	return diags
}

// boolDiverged reports whether a configured (non-null, non-unknown) bool attribute
// differs from the value PostHog returned (nil server value counts as divergent).
func boolDiverged(configured types.Bool, server *bool) bool {
	if configured.IsNull() || configured.IsUnknown() {
		return false
	}
	return server == nil || *server != configured.ValueBool()
}

// int64Diverged is the int64 counterpart of boolDiverged.
func int64Diverged(configured types.Int64, server *int64) bool {
	if configured.IsNull() || configured.IsUnknown() {
		return false
	}
	return server == nil || *server != configured.ValueInt64()
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

func (o ProjectSettingsOps) Delete(ctx context.Context, _ httpclient.PosthogClient, model ProjectSettingsModel) (httpclient.HTTPStatusCode, error) {
	// Settings live on the environment, which is never destroyed by this resource.
	// Deleting simply stops Terraform from managing the settings; PostHog keeps the
	// last-applied values. This is intentionally a no-op, so log it explicitly:
	// otherwise the generic framework logs "Deleted ..." and an operator could
	// wrongly conclude the settings were reset server-side.
	tflog.Warn(ctx, "posthog_project_settings destroy is a no-op: settings remain on PostHog and are NOT reset; only Terraform stops managing them", map[string]any{
		"project_id": model.GetEffectiveProjectID(),
	})
	return 0, nil
}
