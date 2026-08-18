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

// nonBlankPattern requires the name to start and end with a non-whitespace character.
// PostHog trims names server-side (verified against a live instance), so a name with
// surrounding whitespace reads back different from the config and fails every apply
// without ever self-healing, since the config keeps the whitespace.
// rfc3339Pattern is deliberately loose about the offset: the API echoes timestamps back
// normalised to UTC, and the response mapper only writes them into state when the
// practitioner declared one, so exact-format round-tripping is not required here.
var rfc3339Pattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[+-]\d{2}:\d{2})$`)

var nonBlankPattern = regexp.MustCompile(`^\S(.*\S)?$`)

const minutesPerDay = 24 * 60

// Defaults PostHog applies when an attribute is omitted. Plan-time validation resolves
// omitted attributes to these so a config that trips an invariant only by relying on a
// default is still caught.
const (
	defaultEnabled           = true
	defaultThresholdCount    = 100
	defaultThresholdOperator = thresholdOperatorAbove
	defaultEvaluationPeriods = 1
	defaultDatapointsToAlarm = 1
)

// The threshold operators PostHog accepts. Named so the schema validator, the default and
// the plan-time check cannot drift apart.
const (
	thresholdOperatorAbove = "above"
	thresholdOperatorBelow = "below"
)

// The alert states that mean nobody is being notified. Named for the same reason as the
// operators: the warning switch and the schema description both refer to them.
const (
	stateBroken  = "broken"
	stateSnoozed = "snoozed"
)

// notNotifyingWarnings reports the states in which an alert exists and plans clean but
// notifies nobody. Terraform manages neither, so without a warning on refresh the first
// signal is a page that never arrived. It stays quiet in every healthy state.
func notNotifyingWarnings(resp httpclient.LogsAlert) diag.Diagnostics {
	var diags diag.Diagnostics

	// Name the alert: during import there is no config for Terraform to attach the
	// diagnostic to, and a stack managing many alerts would otherwise emit warnings that
	// read identically.
	who := fmt.Sprintf("%q (%s)", strings.TrimSpace(util.PtrToString(resp.Name)), resp.ID)

	switch strings.TrimSpace(util.PtrToString(resp.State)) {
	case stateBroken:
		diags.AddWarning(
			"Log alert is not notifying",
			fmt.Sprintf("PostHog marked alert %s broken after repeated failed checks, so it has stopped "+
				"evaluating and notifies nobody. Reset it from the PostHog UI. Terraform does not manage this "+
				"state, so plans stay clean until you do.", who),
		)
	case stateSnoozed:
		diags.AddWarning(
			"Log alert is snoozed",
			fmt.Sprintf("Alert %s is snoozed and notifies nobody until the snooze expires or is cleared from "+
				"the PostHog UI. Terraform does not manage snoozing, so plans stay clean meanwhile.", who),
		)
	}

	return diags
}

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
	SnoozeUntil       types.String         `tfsdk:"snooze_until"`
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
			"~> **An alert with no destination still evaluates, but notifies nobody.** Delivery is not a field " +
			"on this resource. When the alert fires PostHog emits the internal event `$logs_alert_firing`, and " +
			"routing that to Slack, a webhook or Microsoft Teams is a separate `posthog_hog_function` of " +
			"`type = \"internal_destination\"` filtered to this alert's id. That is the same pattern insight " +
			"alerts use, so notifications are manageable as code today: see " +
			"`examples/alert-notifications/logs-alert.tf`. Destinations attached from the PostHog UI are hog " +
			"functions too, so Terraform does not adopt them unless you declare them.\n\n" +
			"Removing `severity_levels`, `service_names`, `filter_group_json`, or `blocked_windows` from your " +
			"configuration clears them server-side. The remaining optional attributes are computed, so removing one " +
			"retains its last applied value rather than restoring the documented default. Set it explicitly to " +
			"change it back. Drift works the same way. Terraform corrects a PostHog UI edit to an attribute you " +
			"declared, but silently adopts one to a computed attribute you left out.",
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
					// Not just LengthAtLeast(1): PostHog trims names, and the response mapper
					// nulls out any name that is blank after trimming, so both a
					// whitespace-only name and one with surrounding whitespace read back
					// different from the config and fail the apply.
					stringvalidator.RegexMatches(nonBlankPattern, "must not be blank or start or end with whitespace"),
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
					"tracked either. If someone edits it in the PostHog UI Terraform will not detect the drift, so " +
					"declare every field you care about. An imported alert adopts PostHog's stored filter group " +
					"verbatim, so the first plan after an import may show one diff that clears on apply.",
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
					stringvalidator.OneOf(thresholdOperatorAbove, thresholdOperatorBelow),
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
				MarkdownDescription: "Quiet hours: time windows during which the alert is not evaluated. Times " +
					"use the project timezone. Windows must not overlap or touch each other. A window may cross " +
					"midnight (for example `22:00` to `06:00`), but only as the sole window, because PostHog " +
					"stores a crossing window as two windows when anything else is configured. PostHog enforces " +
					"its own limits on window length and count, and reports them on apply. Omit the attribute, " +
					"or set it to an empty list, to disable quiet hours.",
				Validators: []validator.Set{},
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
			"snooze_until": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "RFC 3339 timestamp until which the alert is silenced, for example " +
					"`2026-01-31T09:00:00Z`. Useful for planned work when you want a window of silence in " +
					"version control rather than in someone's memory.\n\n" +
					"Managed only when you set it. Leave it out and Terraform never sends it, so a snooze an " +
					"operator sets in the PostHog UI is left alone rather than being reverted on the next apply.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(rfc3339Pattern, "must be an RFC 3339 timestamp, for example 2026-01-31T09:00:00Z"),
				},
			},
			"state": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Current evaluation state of the alert, as PostHog reports it: `not_firing`, " +
					"`firing`, `pending_resolve`, `errored`, `snoozed` or `broken`. A `broken` alert has stopped " +
					"evaluating after repeated failed checks and notifies nobody; clearing that is a PostHog UI " +
					"action, so the provider warns about it on refresh rather than managing it.",
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

	req.SnoozeUntil = util.StringPtrFromValue(model.SnoozeUntil)
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

	// Only reflected when the configuration declares it. Adopting a snooze set in the
	// PostHog UI would make the next apply revert it, which is the operator's call to make,
	// not Terraform's.
	if !model.SnoozeUntil.IsNull() && !model.SnoozeUntil.IsUnknown() {
		model.SnoozeUntil = core.PtrToStringNullIfEmptyTrimmed(resp.SnoozeUntil)
	}

	diags.Append(notNotifyingWarnings(resp)...)

	// The API omits filter keys that are not set, so an absent key means null rather than
	// an empty collection.
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

// blockedWindowsToSet keeps an explicitly empty set empty. Flipping it to null would show
// as drift on every plan, the same reason core.TagsToSetPreserveEmpty exists.
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

// nonEmptyJSONObjectValidator rejects JSON that parses but is not a filter group.
// jsontypes.Normalized only checks that the string is well-formed, so without this an
// empty object, a bare null or an array reaches the API and fails mid-apply.
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

// ModifyResourcePlan rejects configurations PostHog would reshape or refuse, so they fail
// at plan time with a message naming the problem.
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

	// On update Terraform carries the refreshed state value into the plan for a computed
	// attribute, but PostHog owns this one and can move it between refresh and apply:
	// disabling an alert resets it, and every alert is re-evaluated on a 5-minute cycle.
	// Planning it as unknown lets the applied value differ without failing as an
	// inconsistent result.
	//
	// Only when something else is actually changing. Marking it unknown on an otherwise
	// no-op plan would make state a permanent diff and never converge.
	if !req.State.Raw.IsNull() && !req.Plan.Raw.Equal(req.State.Raw) {
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("state"), types.StringUnknown())...)
	}
}

// validateLogsAlertPlan is separate from ModifyResourcePlan so it can be unit tested
// against a model. It needs both plan and config because neither alone gives the effective
// value of an Optional+Computed attribute; see util.ResolveInt64.
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

	datapoints, datapointsKnown := util.ResolveInt64(plan.DatapointsToAlarm, config.DatapointsToAlarm, defaultDatapointsToAlarm)
	periods, periodsKnown := util.ResolveInt64(plan.EvaluationPeriods, config.EvaluationPeriods, defaultEvaluationPeriods)
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
	operator, operatorKnown := util.ResolveString(plan.ThresholdOperator, config.ThresholdOperator, defaultThresholdOperator)
	count, countKnown := util.ResolveInt64(plan.ThresholdCount, config.ThresholdCount, defaultThresholdCount)
	if operatorKnown && countKnown && operator == thresholdOperatorBelow && count == 0 {
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

	enabled, enabledKnown := util.ResolveBool(plan.Enabled, config.Enabled, defaultEnabled)
	if !enabledKnown || !enabled {
		return diags
	}
	// An unresolved filter might turn out to be non-empty, so it counts as one and the API
	// gets the final say.
	if util.IsEmptySet(plan.SeverityLevels) && util.IsEmptySet(plan.ServiceNames) && plan.FilterGroupJSON.IsNull() {
		diags.AddError(
			"Log alert has no filters",
			"An enabled log alert needs at least one of severity_levels, service_names, or filter_group_json. "+
				"Set enabled = false to save it without filters.",
		)
	}
	return diags
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

		// A window with the same start and end blocks no time. Skip it. Treating it as a
		// span would read 00:00-00:00 as a whole day. PostHog rejects it on apply.
		if start == end {
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
