package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
)

func NewProjectSettings() resource.Resource {
	return core.NewGenericResource[ProjectSettingsModel, projectSettingsRequest, projectSettingsResponse](
		ProjectSettingsOps{},
		core.ProjectSingletonImportParser[ProjectSettingsModel](),
	)
}

// projectSettingsRequest is the environment PATCH body plus the events retention
// period, which PostHog serves from its own endpoint rather than the environment.
type projectSettingsRequest struct {
	httpclient.EnvironmentSettingsRequest
	// EventsRetentionMonths is the plan's period, null or unknown when the config
	// leaves PostHog's value alone.
	EventsRetentionMonths types.Int64
	// PlannedFromEventsRetentionMonths is the period the plan was made against:
	// state's value on update (null for the organization's period), unknown on
	// create, which has no state.
	PlannedFromEventsRetentionMonths types.Int64
}

// projectSettingsResponse is the environment plus the project's events retention period.
type projectSettingsResponse struct {
	httpclient.Environment
	// EventsRetentionMonths is the project's own period, or null while the project
	// uses its organization's period.
	EventsRetentionMonths types.Int64
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
	CapturePerformanceOptIn    types.Bool   `tfsdk:"capture_performance_opt_in"`
	AnonymizeIps               types.Bool   `tfsdk:"anonymize_ips"`
	// NetworkPayloadCapture holds a NetworkPayloadCaptureModel object.
	NetworkPayloadCapture types.Object `tfsdk:"session_recording_network_payload_capture_config"`
	AppURLs               types.List   `tfsdk:"app_urls"`
	RecordingDomains      types.List   `tfsdk:"recording_domains"`
	// TestAccountFilters is the "internal and test users" filter list as a normalized
	// JSON array; server-injected cohort_name is stripped in MapResponseToModel.
	TestAccountFilters               core.NormalizedJSON `tfsdk:"test_account_filters"`
	TestAccountFiltersDefaultChecked types.Bool          `tfsdk:"test_account_filters_default_checked"`
	EventsRetentionMonths            types.Int64         `tfsdk:"events_retention_months"`
}

// NetworkPayloadCaptureModel is the Terraform shape of the
// session_recording_network_payload_capture_config JSON blob.
type NetworkPayloadCaptureModel struct {
	RecordHeaders types.Bool `tfsdk:"record_headers"`
	RecordBody    types.Bool `tfsdk:"record_body"`
}

var networkPayloadCaptureAttrTypes = map[string]attr.Type{
	"record_headers": types.BoolType,
	"record_body":    types.BoolType,
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

These settings live on the PostHog environment object (` + "`/api/environments/{id}/`" + `), which is distinct from the organization-scoped ` + "`posthog_project`" + ` resource. Each setting is optional - omitted attributes are left at whatever value PostHog currently has, and PostHog's server-side default is read back into state. Updates send only configured settings. When a plan changes settings, omitted attributes become unknown, show as ` + "`(known after apply)`" + `, and refresh from PostHog after the update. References to omitted attributes can also be unknown until apply.

~> **Singleton:** There is only one settings object per project. This resource is a singleton - creating multiple instances for the same project will cause conflicts.

~> **Destroy Behavior:** Destroying this resource is a no-op. It stops Terraform from managing the settings but does **not** reset any values on PostHog; the last-applied settings remain in effect.

~> **Plan-gated settings:** If PostHog accepts the update but silently ignores a setting (for example a feature that is not enabled for your plan), Terraform reports a generic "Provider produced inconsistent result after apply" error. The provider also emits a warning naming the specific attribute(s) PostHog did not apply, so you can tell which setting is unavailable for your project rather than a provider bug.

~> **Domains:** ` + "`app_urls`" + ` is the project's authorized-domains list — shown in PostHog settings as **Web analytics domains** and reused as the toolbar's Authorized URLs (web analytics has no separate allowlist of its own). ` + "`recording_domains`" + ` is the session-replay domains list.

~> **Events retention:** ` + "`events_retention_months`" + ` is read and set through PostHog's events retention endpoint (` + "`/api/projects/{id}/events_retention/`" + `), not the environment, so it has its own access rules and is read on every refresh.`,
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
			},
			"autocapture_exceptions_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether exception autocapture is enabled.",
			},
			"session_recording_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether session recording (recording user sessions) is enabled.",
			},
			"surveys_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether surveys are enabled.",
			},
			"cookieless_server_hash_mode": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The cookieless server hash mode. Known values: `0` (disabled), `1` (stateless), `2` (stateful) — matching PostHog's `CookielessServerHashMode` enum. PostHog may add modes over time, so any non-negative value is accepted rather than a fixed set; consult the PostHog docs for the current options.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"autocapture_web_vitals_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether web vitals autocapture is enabled.",
			},
			"capture_performance_opt_in": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether network performance capture is enabled for session replay. This must be `true` for `session_recording_network_payload_capture_config` to have any effect — without it, session replay does not capture network requests at all.",
			},
			"anonymize_ips": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to discard the client IP address at ingestion (shown in PostHog as **Discard client IP data**). Enabling this degrades IP-based event matching for destinations that rely on it (for example Meta Conversions API), since `$ip` is never stored.",
			},
			"session_recording_network_payload_capture_config": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Controls whether session replay records the headers and bodies of captured network requests. Only takes effect when `capture_performance_opt_in` is `true`. " +
					"**Security:** recorded headers and bodies can contain sensitive data such as `Authorization` tokens, cookies, and PII from request/response payloads — enable these deliberately and review PostHog's network-capture redaction options in your SDK configuration. " +
					"PostHog replaces this JSON object wholesale on update, so both `record_headers` and `record_body` must be set when the block is configured; omitting the whole block leaves the current PostHog value untouched.",
				Attributes: map[string]schema.Attribute{
					"record_headers": schema.BoolAttribute{
						Required:            true,
						MarkdownDescription: "Whether to record the headers of network requests in session replay.",
					},
					"record_body": schema.BoolAttribute{
						Required:            true,
						MarkdownDescription: "Whether to record the bodies of network requests in session replay.",
					},
				},
			},
			"app_urls": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The project's authorized domains — shown in PostHog settings as **Web analytics domains** (and used as the toolbar's Authorized URLs). These are the domains tracked in web analytics and where the toolbar is enabled. Maps to the team `app_urls` field. Wildcards are not allowed; order is preserved.",
			},
			"recording_domains": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Authorized domains for session replay. Maps to the team `recording_domains` field. Order is preserved.",
			},
			"test_account_filters": schema.StringAttribute{
				CustomType: core.NormalizedJSONType{},
				Optional:   true,
				Computed:   true,
				MarkdownDescription: "The **internal and test users** filter list as a JSON array of filter objects — the definition every managed insight and hog function references through `filter_test_accounts`. " +
					"Managing it here makes that shared definition owned by Terraform, so a UI edit surfaces as drift instead of silently changing what all filtered insights measure. " +
					"Compared semantically, so key ordering and whitespace differences from the PostHog API do not produce a diff. " +
					"A filter object may reference a cohort (`{\"key\":\"id\",\"type\":\"cohort\",\"value\":<cohort_id>,\"operator\":\"in\"}`); PostHog injects a read-only `cohort_name` into such objects, which this provider strips from state so the filter round-trips without a perpetual diff.",
			},
			"test_account_filters_default_checked": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the **internal and test users** filters (`test_account_filters`) are applied by default across the project. PostHog may report this as null until it is first set.",
			},
			"events_retention_months": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "How many months of events this project keeps, overriding the organization's period. Null while the project uses the organization's period. " +
					"PostHog currently accepts `1`, `3`, `12`, `24`, `36`, `60` and `84`, up to the longest period the organization's plan allows. PostHog checks the value when it's set, so a period it rejects fails the apply. " +
					"Only organization admins can change it, and Terraform only sends it when it differs from PostHog's value. " +
					"**A shorter period permanently deletes older events** at PostHog's next monthly cleanup, and they can't be restored. Queries may stop returning those events as soon as the period is set. The plan warns when it may shorten it, and an apply stops if the period changed in PostHog after its plan was made. " +
					"Removing the attribute leaves the period as it is. To return the project to the organization's period, clear it in PostHog's project settings.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
		},
	}
}

func (o ProjectSettingsOps) BuildCreateRequest(ctx context.Context, model ProjectSettingsModel) (projectSettingsRequest, diag.Diagnostics) {
	req, diags := o.buildRequest(ctx, model)
	req.PlannedFromEventsRetentionMonths = types.Int64Unknown()
	diags.Append(networkCaptureWarning(model.NetworkPayloadCapture, model.CapturePerformanceOptIn)...)
	return req, diags
}

// ModifyResourcePlan warns when the plan may shorten the events retention
// period, then plans no change when every configured setting already matches
// state. Otherwise the plan is left as the framework built it, with the
// settings the config leaves out unknown until the update reads them back.
func (o ProjectSettingsOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, state, plan types.Object
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := state.Attributes()

	// Without state for this project, the plan is a create or a replacement the
	// framework has already planned, and the project's current events retention
	// period is unknown. project_id is compared plan to state because the config
	// leaves it null when the provider-level project applies.
	hasPrior := plan.Attributes()["project_id"].Equal(prior["project_id"])
	have := types.Int64Unknown()
	if hasPrior {
		have = prior["events_retention_months"].(types.Int64)
	}
	resp.Diagnostics.Append(shorterRetentionWarning(config.Attributes()["events_retention_months"].(types.Int64), have)...)
	if !hasPrior {
		return
	}
	for name, configured := range config.Attributes() {
		if name == "project_id" || configured.IsNull() {
			continue
		}
		if !configured.Equal(prior[name]) {
			return
		}
	}

	// Restore the whole prior state, not only the omitted settings: required
	// children of an omitted computed object can make the framework propose
	// changes even when every configured setting still matches state.
	resp.Plan.Raw = req.State.Raw
}

// shorterRetentionWarning warns when the configured events retention period may
// be shorter than the project's current one: PostHog then deletes the older
// events for good. configured is unknown when it comes from a value only known
// at apply. have is null while the project uses the organization's period, and
// unknown when Terraform hasn't read the project's period.
func shorterRetentionWarning(configured, have types.Int64) diag.Diagnostics {
	var change string
	switch {
	case configured.IsNull():
		return nil
	case configured.IsUnknown():
		change = "sets the project's events retention to a period only known at apply, which may be shorter than its current one"
	case have.IsUnknown():
		change = fmt.Sprintf("sets the project's events retention to %d months, replacing its current period in PostHog, which may be longer", configured.ValueInt64())
	case have.IsNull():
		change = fmt.Sprintf("sets the project's events retention to %d months, replacing the organization's period, which may be longer", configured.ValueInt64())
	case configured.ValueInt64() >= have.ValueInt64():
		return nil
	default:
		change = fmt.Sprintf("shortens the project's events retention from %d to %d months", have.ValueInt64(), configured.ValueInt64())
	}
	var diags diag.Diagnostics
	diags.AddWarning(
		"events_retention_months may delete events",
		"This plan "+change+". Queries may stop returning events older than the new period as soon as it's set, and PostHog permanently deletes them at its next monthly cleanup. Deleted events can't be restored.",
	)
	return diags
}

func (o ProjectSettingsOps) buildRequest(ctx context.Context, model ProjectSettingsModel) (projectSettingsRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	req := projectSettingsRequest{
		EnvironmentSettingsRequest: httpclient.EnvironmentSettingsRequest{
			HeatmapsOptIn:              util.BoolPtrFromValue(model.HeatmapsOptIn),
			AutocaptureExceptionsOptIn: util.BoolPtrFromValue(model.AutocaptureExceptionsOptIn),
			SessionRecordingOptIn:      util.BoolPtrFromValue(model.SessionRecordingOptIn),
			SurveysOptIn:               util.BoolPtrFromValue(model.SurveysOptIn),
			AutocaptureWebVitalsOptIn:  util.BoolPtrFromValue(model.AutocaptureWebVitalsOptIn),
			CapturePerformanceOptIn:    util.BoolPtrFromValue(model.CapturePerformanceOptIn),
			AnonymizeIps:               util.BoolPtrFromValue(model.AnonymizeIps),
			CookielessServerHashMode:   util.Int64PtrFromValue(model.CookielessServerHashMode),
		},
		EventsRetentionMonths: model.EventsRetentionMonths,
	}

	// An unset block stays nil so the PATCH omits the field entirely (leave
	// untouched). When set, both children are sent: PostHog replaces the whole
	// JSON blob, so a partial object would wipe the missing key server-side.
	if core.IsKnown(model.NetworkPayloadCapture) {
		var cfg NetworkPayloadCaptureModel
		diags.Append(model.NetworkPayloadCapture.As(ctx, &cfg, basetypes.ObjectAsOptions{})...)
		if !diags.HasError() {
			req.SessionRecordingNetworkPayloadCaptureConfig = &httpclient.NetworkPayloadCaptureConfig{
				RecordHeaders: util.BoolPtrFromValue(cfg.RecordHeaders),
				RecordBody:    util.BoolPtrFromValue(cfg.RecordBody),
			}
		}
	}

	appURLs, d := util.StringListToSlicePtr(ctx, model.AppURLs)
	diags.Append(d...)
	req.AppURLs = appURLs

	recordingDomains, d := util.StringListToSlicePtr(ctx, model.RecordingDomains)
	diags.Append(d...)
	req.RecordingDomains = recordingDomains

	// Only send test_account_filters when the user configured it, so an unset
	// attribute leaves PostHog's current value untouched. A pointer to the parsed
	// slice is used (rather than a bare slice) so an explicit empty array "[]" is
	// sent to clear the filters instead of being dropped by omitempty.
	if core.IsKnown(model.TestAccountFilters) {
		var filters []interface{}
		if err := json.Unmarshal([]byte(model.TestAccountFilters.ValueString()), &filters); err != nil {
			diags.AddError("Invalid test_account_filters JSON", fmt.Sprintf("Could not parse test_account_filters: %s", err.Error()))
		} else {
			req.TestAccountFilters = &filters
		}
	}

	req.TestAccountFiltersDefaultChecked = util.BoolPtrFromValue(model.TestAccountFiltersDefaultChecked)

	return req, diags
}

func (o ProjectSettingsOps) BuildUpdateRequest(ctx context.Context, plan, state ProjectSettingsModel) (projectSettingsRequest, diag.Diagnostics) {
	req, diags := o.buildRequest(ctx, plan)
	req.PlannedFromEventsRetentionMonths = state.EventsRetentionMonths
	// An omitted capture_performance_opt_in is unknown in an update plan, so the
	// warning goes by the value PostHog last reported. The request still omits it.
	capturePerformance := plan.CapturePerformanceOptIn
	if capturePerformance.IsUnknown() {
		capturePerformance = state.CapturePerformanceOptIn
	}
	diags.Append(networkCaptureWarning(plan.NetworkPayloadCapture, capturePerformance)...)
	return req, diags
}

// networkCaptureWarning warns when network payload capture is configured but
// capture_performance_opt_in, which it depends on, is not known to be on.
func networkCaptureWarning(payload types.Object, capturePerformance types.Bool) diag.Diagnostics {
	var diags diag.Diagnostics
	// ValueBool is false for null and unknown too: warn unless performance
	// capture is known to be enabled.
	if payload.IsNull() || payload.IsUnknown() || capturePerformance.ValueBool() {
		return diags
	}
	diags.AddWarning(
		"session_recording_network_payload_capture_config has no effect without capture_performance_opt_in",
		"Session replay only captures network requests when network performance capture is enabled, "+
			"so the configured payload capture settings will do nothing until capture_performance_opt_in is true. "+
			"Set capture_performance_opt_in = true in this resource, or ignore this warning if network capture "+
			"is already enabled for the project outside Terraform.",
	)
	return diags
}

func (o ProjectSettingsOps) MapResponseToModel(ctx context.Context, resp projectSettingsResponse, model *ProjectSettingsModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Before overwriting state, detect any setting the user explicitly configured
	// that PostHog returned with a different value (e.g. a feature not enabled for
	// their plan). Naming those attributes turns the framework's otherwise generic
	// "inconsistent result after apply" error into an actionable message instead of
	// leaving the user to trial-and-error each field. Unconfigured (null/unknown)
	// attributes are skipped, so this only fires on a genuine divergence.
	var ignored []string
	if util.BoolDiverged(model.HeatmapsOptIn, resp.HeatmapsOptIn) {
		ignored = append(ignored, "heatmaps_opt_in")
	}
	if util.BoolDiverged(model.AutocaptureExceptionsOptIn, resp.AutocaptureExceptionsOptIn) {
		ignored = append(ignored, "autocapture_exceptions_opt_in")
	}
	if util.BoolDiverged(model.SessionRecordingOptIn, resp.SessionRecordingOptIn) {
		ignored = append(ignored, "session_recording_opt_in")
	}
	if util.BoolDiverged(model.SurveysOptIn, resp.SurveysOptIn) {
		ignored = append(ignored, "surveys_opt_in")
	}
	if util.BoolDiverged(model.AutocaptureWebVitalsOptIn, resp.AutocaptureWebVitalsOptIn) {
		ignored = append(ignored, "autocapture_web_vitals_opt_in")
	}
	if util.Int64Diverged(model.CookielessServerHashMode, resp.CookielessServerHashMode) {
		ignored = append(ignored, "cookieless_server_hash_mode")
	}
	if util.BoolDiverged(model.CapturePerformanceOptIn, resp.CapturePerformanceOptIn) {
		ignored = append(ignored, "capture_performance_opt_in")
	}
	if util.BoolDiverged(model.AnonymizeIps, resp.AnonymizeIps) {
		ignored = append(ignored, "anonymize_ips")
	}
	if networkPayloadCaptureDiverged(ctx, model.NetworkPayloadCapture, resp.SessionRecordingNetworkPayloadCaptureConfig) {
		ignored = append(ignored, "session_recording_network_payload_capture_config")
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
	model.CapturePerformanceOptIn = core.PtrToBool(resp.CapturePerformanceOptIn)
	model.AnonymizeIps = core.PtrToBool(resp.AnonymizeIps)
	model.CookielessServerHashMode = util.PtrToInt64(resp.CookielessServerHashMode)

	if resp.SessionRecordingNetworkPayloadCaptureConfig == nil {
		// PostHog stores null until the setting is first written; map that to a
		// null object rather than an object of null bools.
		model.NetworkPayloadCapture = types.ObjectNull(networkPayloadCaptureAttrTypes)
	} else {
		obj, d := types.ObjectValueFrom(ctx, networkPayloadCaptureAttrTypes, NetworkPayloadCaptureModel{
			RecordHeaders: core.PtrToBool(resp.SessionRecordingNetworkPayloadCaptureConfig.RecordHeaders),
			RecordBody:    core.PtrToBool(resp.SessionRecordingNetworkPayloadCaptureConfig.RecordBody),
		})
		diags.Append(d...)
		model.NetworkPayloadCapture = obj
	}

	appURLs, d := util.StringSlicePtrToList(ctx, resp.AppURLs)
	diags.Append(d...)
	model.AppURLs = appURLs

	recordingDomains, d := util.StringSlicePtrToList(ctx, resp.RecordingDomains)
	diags.Append(d...)
	model.RecordingDomains = recordingDomains

	if resp.TestAccountFilters == nil {
		model.TestAccountFilters = core.NewNormalizedJSONNull()
	} else {
		// PostHog injects a server-computed cohort_name into cohort-referencing filter
		// objects; strip it so the filter round-trips without a perpetual diff.
		stripped := util.StripFields(*resp.TestAccountFilters, serverComputedFilterKeys)
		jsonStr, err := marshalJSON(stripped)
		if err != nil {
			diags.AddError("Failed to normalize test_account_filters", err.Error())
			return diags
		}
		model.TestAccountFilters = core.NewNormalizedJSONValue(jsonStr)
	}

	model.TestAccountFiltersDefaultChecked = core.PtrToBool(resp.TestAccountFiltersDefaultChecked)
	model.EventsRetentionMonths = resp.EventsRetentionMonths

	return diags
}

// networkPayloadCaptureDiverged reports whether a configured (non-null,
// non-unknown) payload-capture object differs from what the server returned.
// It is the object counterpart of util.BoolDiverged: unconfigured objects
// never diverge; a nil server value or an unparseable configured value counts
// as divergent.
func networkPayloadCaptureDiverged(ctx context.Context, configured types.Object, server *httpclient.NetworkPayloadCaptureConfig) bool {
	if configured.IsNull() || configured.IsUnknown() {
		return false
	}
	var cfg NetworkPayloadCaptureModel
	if d := configured.As(ctx, &cfg, basetypes.ObjectAsOptions{}); d.HasError() {
		// Unreachable for a schema-derived object, but if it ever fires the
		// attribute must be named in the divergence warning rather than
		// vanishing into the generic "inconsistent result after apply" error.
		tflog.Error(ctx, "failed to parse configured session_recording_network_payload_capture_config; treating as diverged", map[string]any{
			"diagnostics": fmt.Sprintf("%v", d),
		})
		return true
	}
	if server == nil {
		return true
	}
	return util.BoolDiverged(cfg.RecordHeaders, server.RecordHeaders) ||
		util.BoolDiverged(cfg.RecordBody, server.RecordBody)
}

func (o ProjectSettingsOps) Create(ctx context.Context, client httpclient.PosthogClient, model ProjectSettingsModel, req projectSettingsRequest) (projectSettingsResponse, error) {
	result, _, err := o.Update(ctx, client, model, req)
	return result, err
}

func (o ProjectSettingsOps) Read(ctx context.Context, client httpclient.PosthogClient, model ProjectSettingsModel) (projectSettingsResponse, httpclient.HTTPStatusCode, error) {
	env, status, err := client.GetEnvironment(ctx, model.GetEffectiveProjectID())
	if err != nil {
		return projectSettingsResponse{}, status, err
	}
	months, _, err := readEventsRetention(ctx, client, model.GetEffectiveProjectID())
	return projectSettingsResponse{Environment: env, EventsRetentionMonths: months}, status, err
}

// Update PATCHes the environment, then sets the events retention period when
// it differs from PostHog's. If the set fails, the environment change has
// already applied; the next apply sends both again.
//
// A set needs an organization admin and can delete events, so the period is
// never sent unchanged, and the apply stops before any write when PostHog's
// period is neither the one the plan was made against nor the one it sets:
// someone changed it after the plan, which never showed or warned about it.
func (o ProjectSettingsOps) Update(ctx context.Context, client httpclient.PosthogClient, model ProjectSettingsModel, req projectSettingsRequest) (projectSettingsResponse, httpclient.HTTPStatusCode, error) {
	projectID := model.GetEffectiveProjectID()
	months, found, err := readEventsRetention(ctx, client, projectID)
	if err != nil {
		return projectSettingsResponse{}, 0, err
	}
	want := req.EventsRetentionMonths
	if core.IsKnown(want) && !found {
		return projectSettingsResponse{}, 0, fmt.Errorf("PostHog returned 404 for /api/projects/%s/events_retention/, which events_retention_months needs: "+
			"the project doesn't exist or this API key can't see it, or this PostHog version predates the events retention endpoint", projectID)
	}
	setsPeriod := core.IsKnown(want) && !want.Equal(months)
	if from := req.PlannedFromEventsRetentionMonths; setsPeriod && !from.IsUnknown() && !from.Equal(months) {
		return projectSettingsResponse{}, 0, fmt.Errorf(
			"events_retention_months is %s in PostHog, but this plan was made when it was %s. Run terraform plan again, with refresh enabled, to review the change",
			describeRetention(months), describeRetention(from))
	}
	env, status, err := client.UpdateEnvironment(ctx, projectID, req.EnvironmentSettingsRequest)
	if err != nil {
		return projectSettingsResponse{}, status, err
	}
	if setsPeriod {
		tflog.Info(ctx, "setting events retention", map[string]any{
			"project_id": projectID,
			"from":       describeRetention(months),
			"to":         describeRetention(want),
		})
		updated, _, err := client.SetEventsRetention(ctx, projectID, httpclient.EventsRetentionRequest{RetentionMonths: want.ValueInt64()})
		if err != nil {
			// Not the set's own status: the generic update takes a 404 to mean
			// the settings are gone and drops them from state.
			return projectSettingsResponse{}, status, fmt.Errorf("the other settings were saved, but setting events_retention_months to %d failed: %w", want.ValueInt64(), err)
		}
		months = util.PtrToInt64(updated.ProjectRetentionMonths)
	}
	return projectSettingsResponse{Environment: env, EventsRetentionMonths: months}, status, nil
}

// describeRetention names a project's events retention period for messages.
func describeRetention(months types.Int64) string {
	if months.IsNull() {
		return "the organization's period"
	}
	return fmt.Sprintf("%d months", months.ValueInt64())
}

// readEventsRetention returns the project's own events retention period, or null
// while it uses the organization's, and whether PostHog has the events
// retention endpoint. Older self-hosted versions answer 404; that reads as no
// period so existing configurations keep refreshing.
func readEventsRetention(ctx context.Context, client httpclient.PosthogClient, projectID string) (types.Int64, bool, error) {
	current, status, err := client.GetEventsRetention(ctx, projectID)
	if status == http.StatusNotFound {
		tflog.Warn(ctx, "events retention endpoint not found; reading events_retention_months as null", map[string]any{
			"project_id": projectID,
		})
		return types.Int64Null(), false, nil
	}
	if err != nil {
		return types.Int64Null(), false, fmt.Errorf("failed to read events retention: %w", err)
	}
	return util.PtrToInt64(current.ProjectRetentionMonths), true, nil
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
