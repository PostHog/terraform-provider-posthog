package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
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

// hhmmPattern matches a 24-hour HH:MM local time, as used by quiet-hours windows.
var hhmmPattern = regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)

var nonBlankPattern = regexp.MustCompile(`\S`)

const (
	minutesPerDay = 24 * 60
	// minBlockedWindowMinutes is the shortest quiet-hours window PostHog accepts.
	minBlockedWindowMinutes = 30
)

// Defaults PostHog applies when an attribute is omitted. Plan-time validation resolves
// omitted attributes to these so a config that trips an invariant only by relying on a
// default is still caught.
const (
	defaultThresholdCount    = 100
	defaultThresholdOperator = "above"
	defaultEvaluationPeriods = 1
	defaultDatapointsToAlarm = 1
)

// blockedWindowAttrTypes mirrors BlockedWindowTFModel for set conversions.
var blockedWindowAttrTypes = map[string]attr.Type{
	"start": types.StringType,
	"end":   types.StringType,
}

func NewLogsAlert() resource.Resource {
	return core.NewGenericResource[LogsAlertTFModel, httpclient.LogsAlertRequest, httpclient.LogsAlert](
		LogsAlertOps{},
		core.ProjectScopedImportParser[LogsAlertTFModel](),
	)
}

type LogsAlertTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name              types.String         `tfsdk:"name"`
	Enabled           types.Bool           `tfsdk:"enabled"`
	SeverityLevels    types.Set            `tfsdk:"severity_levels"`
	ServiceNames      types.Set            `tfsdk:"service_names"`
	FilterGroupJSON   jsontypes.Normalized `tfsdk:"filter_group_json"`
	ThresholdCount    types.Int64          `tfsdk:"threshold_count"`
	ThresholdOperator types.String         `tfsdk:"threshold_operator"`
	WindowMinutes     types.Int64          `tfsdk:"window_minutes"`
	EvaluationPeriods types.Int64          `tfsdk:"evaluation_periods"`
	DatapointsToAlarm types.Int64          `tfsdk:"datapoints_to_alarm"`
	CooldownMinutes   types.Int64          `tfsdk:"cooldown_minutes"`
	BlockedWindows    types.Set            `tfsdk:"blocked_windows"`
	State             types.String         `tfsdk:"state"`
}

type BlockedWindowTFModel struct {
	Start types.String `tfsdk:"start"`
	End   types.String `tfsdk:"end"`
}

type LogsAlertOps struct{}

func (o LogsAlertOps) ResourceName() string {
	return "Logs Alert"
}

func (o LogsAlertOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog [log alerts](https://posthog.com/docs/logs/alerts). A log alert " +
			"periodically counts the log entries matching its filters over a rolling window and fires when that " +
			"count crosses the threshold.\n\n" +
			"At least one of `severity_levels`, `service_names`, or `filter_group_json` is required unless the " +
			"alert is disabled (`enabled = false`). A project may hold at most 20 log alerts.\n\n" +
			"~> **Log alerts are gated behind a feature flag.** The API returns `403` with " +
			"`This action requires feature flag 'logs-alerting' to be enabled for your organization` until " +
			"PostHog enables it for you. A `403` on the first apply means the flag is off, not that the " +
			"credentials or configuration are wrong.\n\n" +
			"~> **Notification destinations are not managed by this resource.** PostHog attaches Slack, webhook, " +
			"and Microsoft Teams destinations through a separate endpoint that the alert API does not read back, " +
			"so Terraform cannot track them without reporting permanent drift. Attach destinations from the " +
			"PostHog UI. An alert with no destination still evaluates, but notifies nobody. This also means any " +
			"change that replaces the resource — notably changing `project_id` — creates a new alert with no " +
			"destinations attached, which you must re-attach manually.\n\n" +
			"~> An alert whose `state` is `broken` or `snoozed` has stopped notifying. Terraform does not manage " +
			"either condition, so `terraform plan` stays clean while the alert is silent; reset or unsnooze it " +
			"from the PostHog UI.\n\n" +
			"Removing `severity_levels`, `service_names`, `filter_group_json`, or `blocked_windows` from your " +
			"configuration clears them server-side. The remaining optional attributes are computed, so removing one " +
			"retains its last applied value rather than restoring the documented default — set it explicitly to " +
			"change it back.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the log alert.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Human-readable name for the alert. PostHog defaults this to `Untitled alert` " +
					"when omitted, so this attribute is computed: leaving it out adopts the server's default rather " +
					"than producing a diff.",
				Validators: []validator.String{
					// Not just LengthAtLeast(1): the response mapper nulls out any name that
					// is blank after trimming, so a whitespace-only name would come back as
					// null against a non-null config and fail the apply.
					stringvalidator.RegexMatches(nonBlankPattern, "must contain at least one non-whitespace character"),
					stringvalidator.LengthAtMost(255),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the alert is actively evaluated. Defaults to true. Disabling resets the alert state to `not_firing`.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"severity_levels": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Log severity levels to count: `trace`, `debug`, `info`, `warn`, `error`, or `fatal`.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf(
						"trace",
						"debug",
						"info",
						"warn",
						"error",
						"fatal",
					)),
				},
			},
			"service_names": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Service names to scope the alert to.",
			},
			"filter_group_json": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
				MarkdownDescription: "Attribute-level filters as JSON, matching the `filters.filterGroup` object of the " +
					"[logs alerts API](https://posthog.com/docs/api/logs). Use this for anything beyond severity and " +
					"service, such as filtering on a log attribute. Must be a non-empty JSON object. Only the fields " +
					"you declare are tracked: PostHog annotates saved filters with defaults (such as `label`) that " +
					"would otherwise surface as permanent drift. The flip side is that a field you omit is not " +
					"tracked either, so if someone edits it in the PostHog UI Terraform will not detect the drift — " +
					"declare every field you care about.",
				Validators: []validator.String{
					nonEmptyJSONObjectValidator{},
				},
			},
			"threshold_count": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Log entry count to compare against. The alert fires when the number of matching " +
					"entries in the window is `above` (or `below`) this value. Defaults to 100. Use `0` with `above` to " +
					"fire on any matching log.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"threshold_operator": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the alert fires when the count is `above` or `below` the threshold. Defaults to `above`.",
				Validators: []validator.String{
					stringvalidator.OneOf("above", "below"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"window_minutes": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Time window in minutes over which log entries are counted: `5`, `10`, `15`, `30`, or `60`. Defaults to 5.",
				Validators: []validator.Int64{
					int64validator.OneOf(5, 10, 15, 30, 60),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"evaluation_periods": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "How many of the most recent check periods to consider. PostHog checks a log alert " +
					"every 5 minutes, so 3 periods covers the last 15 minutes. Must be between 1 and 10. Defaults to 1.",
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"datapoints_to_alarm": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "How many of the `evaluation_periods` most recent check periods must breach the " +
					"threshold before the alert fires. Must be between 1 and 10 and must not exceed " +
					"`evaluation_periods`. Defaults to 1.",
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"cooldown_minutes": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Minimum minutes between repeated notifications after the alert fires. Defaults to 0, meaning no cooldown.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"blocked_windows": schema.SetNestedAttribute{
				Optional: true,
				MarkdownDescription: "Quiet hours: up to 5 time windows during which the alert is not evaluated. Times " +
					"use the project timezone. Each window must span at least 30 minutes, and windows must not " +
					"overlap or touch each other. A window may cross midnight (for example `22:00` to `06:00`), but " +
					"only as the sole window: PostHog stores blocked windows on a single merged 24-hour timeline, " +
					"and a crossing window alongside another one is stored as two windows rather than one. Omit the " +
					"attribute, or set it to an empty list, to disable quiet hours.",
				Validators: []validator.Set{
					setvalidator.SizeAtMost(5),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"start": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Start time as `HH:MM` (24-hour, project timezone). Inclusive.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(hhmmPattern, "must be a 24-hour time in HH:MM format"),
							},
						},
						"end": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "End time as `HH:MM` (24-hour, project timezone). Exclusive.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(hhmmPattern, "must be a 24-hour time in HH:MM format"),
							},
						},
					},
				},
			},
			"state": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Current evaluation state of the alert: `not_firing`, `firing`, `snoozed`, or " +
					"`broken`. `broken` means PostHog stopped evaluating the alert after repeated failed checks — it " +
					"notifies nobody until it is reset from the PostHog UI.",
			},
		},
	}
}

func (o LogsAlertOps) BuildCreateRequest(ctx context.Context, model LogsAlertTFModel) (httpclient.LogsAlertRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	// filters and schedule_restriction are whole-object replacements, so filters is
	// always sent, even when empty (which is how a draft alert clears its filters).
	req := httpclient.LogsAlertRequest{
		Filters: &httpclient.LogsAlertFilters{},
	}

	req.Name = util.StringPtrFromValue(model.Name)
	req.Enabled = util.BoolPtrFromValue(model.Enabled)

	severityLevels, d := core.ExtractTags(ctx, model.SeverityLevels)
	diags.Append(d...)
	if diags.HasError() {
		return req, diags
	}
	req.Filters.SeverityLevels = severityLevels

	serviceNames, d := core.ExtractTags(ctx, model.ServiceNames)
	diags.Append(d...)
	if diags.HasError() {
		return req, diags
	}
	req.Filters.ServiceNames = serviceNames

	filterGroup, d := util.ParseJSONStringMap("filter_group_json", model.FilterGroupJSON.StringValue)
	diags.Append(d...)
	if diags.HasError() {
		return req, diags
	}
	req.Filters.FilterGroup = filterGroup

	req.ThresholdCount = util.Int64PtrFromValue(model.ThresholdCount)
	req.ThresholdOperator = util.StringPtrFromValue(model.ThresholdOperator)
	req.WindowMinutes = util.Int64PtrFromValue(model.WindowMinutes)
	req.EvaluationPeriods = util.Int64PtrFromValue(model.EvaluationPeriods)
	req.DatapointsToAlarm = util.Int64PtrFromValue(model.DatapointsToAlarm)
	req.CooldownMinutes = util.Int64PtrFromValue(model.CooldownMinutes)

	if !model.BlockedWindows.IsNull() && !model.BlockedWindows.IsUnknown() {
		var windows []BlockedWindowTFModel
		diags.Append(model.BlockedWindows.ElementsAs(ctx, &windows, false)...)
		if diags.HasError() {
			return req, diags
		}
		blockedWindows := make([]httpclient.LogsAlertBlockedWindow, len(windows))
		for i, w := range windows {
			blockedWindows[i] = httpclient.LogsAlertBlockedWindow{
				Start: w.Start.ValueString(),
				End:   w.End.ValueString(),
			}
		}
		req.ScheduleRestriction = &httpclient.LogsAlertSchedule{BlockedWindows: blockedWindows}
	}

	return req, diags
}

func (o LogsAlertOps) BuildUpdateRequest(ctx context.Context, plan, state LogsAlertTFModel) (httpclient.LogsAlertRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o LogsAlertOps) MapResponseToModel(ctx context.Context, resp httpclient.LogsAlert, model *LogsAlertTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.Enabled = core.PtrToBool(resp.Enabled)
	model.State = core.PtrToStringNullIfEmptyTrimmed(resp.State)

	// Terraform does not manage state, so a broken or snoozed alert leaves plans clean
	// while notifying nobody. Warn on refresh rather than leaving the docs as the only
	// signal. It fires only in the degraded state, so it does not nag in steady state.
	switch strings.TrimSpace(util.PtrToString(resp.State)) {
	case "broken":
		diags.AddWarning(
			"Log alert is not notifying",
			"PostHog marked this alert broken after repeated failed checks, so it has stopped evaluating and "+
				"notifies nobody. Reset it from the PostHog UI. Terraform does not manage this state, so plans "+
				"stay clean until you do.",
		)
	case "snoozed":
		diags.AddWarning(
			"Log alert is snoozed",
			"This alert is snoozed and notifies nobody until the snooze expires or is cleared from the PostHog "+
				"UI. Terraform does not manage snoozing, so plans stay clean meanwhile.",
		)
	}

	// The API echoes back only the filter keys that are set, so an absent key maps
	// to null rather than an empty collection.
	var severityLevels, serviceNames []string
	var filterGroup map[string]any
	if resp.Filters != nil {
		severityLevels = resp.Filters.SeverityLevels
		serviceNames = resp.Filters.ServiceNames
		filterGroup = resp.Filters.FilterGroup
	}

	severitySet, d := core.TagsToSetPreserveEmpty(ctx, severityLevels, model.SeverityLevels)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.SeverityLevels = severitySet

	serviceSet, d := core.TagsToSetPreserveEmpty(ctx, serviceNames, model.ServiceNames)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.ServiceNames = serviceSet

	// PostHog annotates saved filters with its own defaults (a leaf sent as
	// {key,type,operator,value} reads back with "label":null), so project the response
	// onto the fields the user actually declared. Without this the state written after
	// apply differs from the config and Terraform reports an inconsistent result.
	// Length, not nil: an empty object means no filter group. Storing "{}" would put a
	// value in state that no config can express, since the validator rejects an empty
	// object, leaving drift that cannot be resolved from configuration.
	if len(filterGroup) > 0 {
		encoded, err := normalizeJSONForState(filterGroup, model.FilterGroupJSON.ValueString())
		if err != nil {
			diags.AddError("Invalid filter group in response", fmt.Sprintf("Could not encode filterGroup returned by PostHog: %s", err.Error()))
			return diags
		}
		model.FilterGroupJSON = jsontypes.NewNormalizedValue(encoded)
	} else {
		model.FilterGroupJSON = jsontypes.NewNormalizedNull()
	}

	model.ThresholdCount = util.PtrToInt64(resp.ThresholdCount)
	model.ThresholdOperator = core.PtrToStringNullIfEmptyTrimmed(resp.ThresholdOperator)
	model.WindowMinutes = util.PtrToInt64(resp.WindowMinutes)
	model.EvaluationPeriods = util.PtrToInt64(resp.EvaluationPeriods)
	model.DatapointsToAlarm = util.PtrToInt64(resp.DatapointsToAlarm)
	model.CooldownMinutes = util.PtrToInt64(resp.CooldownMinutes)

	blockedWindows, d := blockedWindowsToSet(ctx, resp.ScheduleRestriction, model.BlockedWindows)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.BlockedWindows = blockedWindows

	return diags
}

// blockedWindowsToSet converts the API's schedule_restriction into the flattened
// blocked_windows set. A null or empty restriction means quiet hours are off, but a
// set the user explicitly configured as empty stays empty — mirroring
// core.TagsToSetPreserveEmpty, so an empty config does not read back as null.
func blockedWindowsToSet(ctx context.Context, schedule *httpclient.LogsAlertSchedule, current types.Set) (types.Set, diag.Diagnostics) {
	objectType := types.ObjectType{AttrTypes: blockedWindowAttrTypes}
	if schedule == nil || len(schedule.BlockedWindows) == 0 {
		if !current.IsNull() && !current.IsUnknown() {
			return types.SetValueFrom(ctx, objectType, []BlockedWindowTFModel{})
		}
		return types.SetNull(objectType), nil
	}

	windows := make([]BlockedWindowTFModel, len(schedule.BlockedWindows))
	for i, w := range schedule.BlockedWindows {
		windows[i] = BlockedWindowTFModel{
			Start: types.StringValue(w.Start),
			End:   types.StringValue(w.End),
		}
	}
	return types.SetValueFrom(ctx, objectType, windows)
}

// nonEmptyJSONObjectValidator rejects filter_group_json values that are well-formed JSON
// but not a usable filter group. jsontypes.Normalized only checks well-formedness, so
// without this `jsonencode({})`, `"null"`, and JSON arrays all fail at apply time — the
// last with a raw Go unmarshal error.
type nonEmptyJSONObjectValidator struct{}

func (v nonEmptyJSONObjectValidator) Description(_ context.Context) string {
	return "must be a non-empty JSON object"
}

func (v nonEmptyJSONObjectValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v nonEmptyJSONObjectValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	var decoded map[string]any
	if err := json.Unmarshal([]byte(req.ConfigValue.ValueString()), &decoded); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid filter_group_json",
			`filter_group_json must be a JSON object, such as `+
				`jsonencode({ type = "AND", values = [...] }).`,
		)
		return
	}
	if len(decoded) == 0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid filter_group_json",
			"filter_group_json must be a non-empty JSON object. Omit the attribute entirely to apply no "+
				"attribute-level filter.",
		)
	}
}

// ModifyResourcePlan enforces the invariants PostHog documents but only checks server-side,
// so a bad config fails at plan time with an actionable message instead of mid-apply.
func (o LogsAlertOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan, config LogsAlertTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(validateLogsAlertPlan(ctx, plan, config)...)
}

// validateLogsAlertPlan is split out from ModifyResourcePlan so the invariants can be unit
// tested directly against a model, including the unknown-value cases. It takes both plan
// and config because neither alone identifies the effective value; see resolveInt64.
func validateLogsAlertPlan(ctx context.Context, plan, config LogsAlertTFModel) diag.Diagnostics {
	var diags diag.Diagnostics
	diags.Append(validateCanEverFire(plan, config)...)
	diags.Append(validateHasFilters(plan, config)...)
	diags.Append(validateBlockedWindows(ctx, plan.BlockedWindows)...)
	return diags
}

// validateCanEverFire rejects threshold/period combinations that no log volume can satisfy.
func validateCanEverFire(plan, config LogsAlertTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	datapoints, datapointsKnown := resolveInt64(plan.DatapointsToAlarm, config.DatapointsToAlarm, defaultDatapointsToAlarm)
	periods, periodsKnown := resolveInt64(plan.EvaluationPeriods, config.EvaluationPeriods, defaultEvaluationPeriods)
	if datapointsKnown && periodsKnown && datapoints > periods {
		diags.AddAttributeError(
			path.Root("datapoints_to_alarm"),
			"Alert can never fire",
			fmt.Sprintf(
				"datapoints_to_alarm (%d) must not exceed evaluation_periods (%d): the alert requires more breaching "+
					"periods than it ever evaluates, so it would never fire.",
				datapoints, periods,
			),
		)
	}

	// A log count is never negative, so "below 0" is unsatisfiable.
	operator, operatorKnown := resolveString(plan.ThresholdOperator, config.ThresholdOperator, defaultThresholdOperator)
	count, countKnown := resolveInt64(plan.ThresholdCount, config.ThresholdCount, defaultThresholdCount)
	if operatorKnown && countKnown && operator == "below" && count == 0 {
		diags.AddAttributeError(
			path.Root("threshold_count"),
			"Alert can never fire",
			"threshold_operator = \"below\" with threshold_count = 0 can never be satisfied, because a log count is "+
				"never negative. Use threshold_count = 1 to fire when no matching logs arrive in the window, or "+
				"threshold_operator = \"above\" with threshold_count = 0 to fire on any matching log.",
		)
	}

	return diags
}

// validateHasFilters enforces PostHog's rule that an enabled alert must filter on something.
func validateHasFilters(plan, config LogsAlertTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	enabled, enabledKnown := resolveBool(plan.Enabled, config.Enabled, true)
	if !enabledKnown || !enabled {
		return diags
	}
	// An unresolved filter could turn out to be non-empty, so it counts as a filter and
	// the rule is left to the API: isEmptySet is false for an unknown set, and an unknown
	// filter_group_json is not null.
	if isEmptySet(plan.SeverityLevels) && isEmptySet(plan.ServiceNames) && plan.FilterGroupJSON.IsNull() {
		diags.AddError(
			"Log alert has no filters",
			"An enabled log alert needs at least one of severity_levels, service_names, or filter_group_json. "+
				"Set enabled = false to save it without filters.",
		)
	}
	return diags
}

// resolveInt64 reports the value PostHog will end up with, and whether it is knowable yet.
// The plan already carries the effective value whenever it is known — including on update,
// where Terraform copies the prior state into the plan for an Optional+Computed attribute
// the config omits, so an omitted attribute keeps its last applied value. The plan is
// unknown only on create, or when the config references something not yet resolvable; a
// null config means the attribute was omitted on create, so the server default applies.
func resolveInt64(plan, config types.Int64, def int64) (int64, bool) {
	if !plan.IsUnknown() && !plan.IsNull() {
		return plan.ValueInt64(), true
	}
	if config.IsNull() {
		return def, true
	}
	return 0, false
}

func resolveString(plan, config types.String, def string) (string, bool) {
	if !plan.IsUnknown() && !plan.IsNull() {
		return plan.ValueString(), true
	}
	if config.IsNull() {
		return def, true
	}
	return "", false
}

func resolveBool(plan, config types.Bool, def bool) (bool, bool) {
	if !plan.IsUnknown() && !plan.IsNull() {
		return plan.ValueBool(), true
	}
	if config.IsNull() {
		return def, true
	}
	return false, false
}

// isEmptySet reports whether a set contributes no values. An unknown set is not empty — its
// elements are simply not resolvable yet — and Elements() returns nothing for unknown, so
// the IsUnknown check has to come first.
func isEmptySet(v types.Set) bool {
	if v.IsUnknown() {
		return false
	}
	return v.IsNull() || len(v.Elements()) == 0
}

// validateBlockedWindows enforces the ≥30-minute span rule, plus the two config shapes
// PostHog would silently reshape. PostHog does not store the windows it is given: it lays
// them on a single 24-hour timeline, merges, and re-derives windows from the result. A
// reshaped config reads back differently from the plan, which fails the apply with an
// inconsistent result rather than showing up as recoverable drift.
func validateBlockedWindows(ctx context.Context, windows types.Set) diag.Diagnostics {
	var diags diag.Diagnostics
	if windows.IsNull() || windows.IsUnknown() {
		return diags
	}

	var parsed []BlockedWindowTFModel
	diags.Append(windows.ElementsAs(ctx, &parsed, false)...)
	if diags.HasError() {
		return diags
	}

	// source is the index of the configured window an interval came from. Two intervals of
	// the same window are the two halves of one midnight crossing, not a conflict. The
	// label is only message text, so it must not be used to decide identity.
	type interval struct {
		start, end int
		source     int
		label      string
	}
	var intervals []interval
	windowCount := 0
	wrappingWindows := 0

	for i, w := range parsed {
		if w.Start.IsUnknown() || w.End.IsUnknown() {
			return diags
		}
		start, okStart := parseHHMM(w.Start.ValueString())
		end, okEnd := parseHHMM(w.End.ValueString())
		if !okStart || !okEnd {
			// The HH:MM regex validator already reports malformed values.
			return diags
		}

		label := fmt.Sprintf("%s-%s", w.Start.ValueString(), w.End.ValueString())

		// An identical start and end covers no time at all rather than a whole day, so it
		// gets its own message — the wrap correction below would otherwise report 1440
		// minutes while complaining the window is too short.
		if start == end {
			diags.AddAttributeError(
				path.Root("blocked_windows"),
				"Quiet-hours window covers no time",
				fmt.Sprintf(
					"Window %s has the same start and end time, so it blocks nothing. Set an end at least %d minutes "+
						"after the start.", label, minBlockedWindowMinutes,
				),
			)
			continue
		}

		span := end - start
		if span < 0 {
			span += minutesPerDay
		}
		if span < minBlockedWindowMinutes {
			diags.AddAttributeError(
				path.Root("blocked_windows"),
				"Quiet-hours window is too short",
				fmt.Sprintf("Window %s spans %d minutes; PostHog requires at least %d.", label, span, minBlockedWindowMinutes),
			)
			continue
		}

		windowCount++

		// Split a window that runs past midnight so overlap is a plain interval comparison.
		// One ending exactly at midnight stops at the end of the day and needs no second
		// half; emitting an empty [0, 0) half would collide with every window starting at
		// midnight in the check below.
		switch {
		case end > start:
			intervals = append(intervals, interval{start: start, end: end, source: i, label: label})
		case end == 0:
			intervals = append(intervals, interval{start: start, end: minutesPerDay, source: i, label: label})
		default:
			wrappingWindows++
			intervals = append(intervals,
				interval{start: start, end: minutesPerDay, source: i, label: label},
				interval{start: 0, end: end, source: i, label: label},
			)
		}
	}

	// Windows that merely touch count as overlapping. PostHog merges on `next.start <=
	// prev.end`, so 01:00-02:00 and 02:00-03:00 are saved as a single 01:00-03:00 window.
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i].source == intervals[j].source {
				continue
			}
			if intervals[i].start <= intervals[j].end && intervals[j].start <= intervals[i].end {
				diags.AddAttributeError(
					path.Root("blocked_windows"),
					"Quiet-hours windows overlap",
					fmt.Sprintf(
						"Windows %s and %s overlap or run straight into each other. PostHog merges them into a "+
							"single window when saving, so the alert would read back with different windows than "+
							"configured and every apply would fail. Combine them into one window.",
						intervals[i].label, intervals[j].label,
					),
				)
				return diags
			}
		}
	}

	// A window written as crossing midnight always expands to two intervals, and PostHog
	// only rejoins them when nothing else is on the timeline. Alongside any other window it
	// stays split, so the alert reads back with one more window than was configured.
	if wrappingWindows > 0 && windowCount > 1 {
		diags.AddAttributeError(
			path.Root("blocked_windows"),
			"Quiet hours crossing midnight must be the only window",
			"A window whose end is before its start crosses midnight, and PostHog stores it as one window per "+
				"side of midnight unless it is the only window configured. Either make it the only window, or "+
				"split it yourself into a window ending at 00:00 and one starting at 00:00.",
		)
		return diags
	}

	// Two windows that between them block both sides of midnight get rejoined into a single
	// crossing window, but only while they are the whole timeline. A third window anywhere
	// in the day leaves all of them stored as written.
	if len(intervals) == 2 && intervals[0].source != intervals[1].source &&
		((intervals[0].start == 0 && intervals[1].end == minutesPerDay) ||
			(intervals[1].start == 0 && intervals[0].end == minutesPerDay)) {
		diags.AddAttributeError(
			path.Root("blocked_windows"),
			"Quiet-hours windows meeting at midnight are stored as one",
			fmt.Sprintf(
				"Windows %s and %s together block midnight from both sides, and PostHog rejoins them into a "+
					"single crossing window when they are the only two, so the alert would read back differently "+
					"than configured. Write them as one window crossing midnight, add a third window elsewhere "+
					"in the day, or end the evening window at 23:59.",
				intervals[0].label, intervals[1].label,
			),
		)
	}

	return diags
}

// parseHHMM converts a validated HH:MM string to minutes past local midnight.
func parseHHMM(v string) (int, bool) {
	hh, mm, found := strings.Cut(v, ":")
	if !found {
		return 0, false
	}
	hours, err := strconv.Atoi(hh)
	if err != nil {
		return 0, false
	}
	minutes, err := strconv.Atoi(mm)
	if err != nil {
		return 0, false
	}
	return hours*60 + minutes, true
}

func (o LogsAlertOps) Create(ctx context.Context, client httpclient.PosthogClient, model LogsAlertTFModel, req httpclient.LogsAlertRequest) (httpclient.LogsAlert, error) {
	return client.CreateLogsAlert(ctx, model.GetEffectiveProjectID(), req)
}

func (o LogsAlertOps) Read(ctx context.Context, client httpclient.PosthogClient, model LogsAlertTFModel) (httpclient.LogsAlert, httpclient.HTTPStatusCode, error) {
	return client.GetLogsAlert(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o LogsAlertOps) Update(ctx context.Context, client httpclient.PosthogClient, model LogsAlertTFModel, req httpclient.LogsAlertRequest) (httpclient.LogsAlert, httpclient.HTTPStatusCode, error) {
	return client.UpdateLogsAlert(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o LogsAlertOps) Delete(ctx context.Context, client httpclient.PosthogClient, model LogsAlertTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteLogsAlert(ctx, model.GetEffectiveProjectID(), model.GetID())
}
