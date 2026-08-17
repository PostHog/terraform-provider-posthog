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

const (
	minutesPerDay = 24 * 60
	// minBlockedWindowMinutes is the shortest quiet-hours window PostHog accepts.
	minBlockedWindowMinutes = 30
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
			"alert is a draft (`enabled = false`). A project may hold at most 20 log alerts.\n\n" +
			"~> **Notification destinations are not managed by this resource.** PostHog attaches Slack, webhook, " +
			"and Microsoft Teams destinations through a separate endpoint that the alert API does not read back, " +
			"so Terraform cannot track them without reporting permanent drift. Attach destinations from the " +
			"PostHog UI. An alert with no destination still evaluates, but notifies nobody. This also means any " +
			"change that replaces the resource — notably changing `project_id` — creates a new alert with no " +
			"destinations attached, which you must re-attach manually.\n\n" +
			"~> An alert whose `state` is `broken` or `snoozed` has stopped notifying. Terraform does not manage " +
			"either condition, so `terraform plan` stays clean while the alert is silent; reset or unsnooze it " +
			"from the PostHog UI.",
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
					"would otherwise surface as permanent drift.",
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
					"every 5 minutes, so 3 periods covers the last 15 minutes. Defaults to 1.",
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
				MarkdownDescription: "How many of those periods must breach the threshold before the alert fires. " +
					"Must not exceed `evaluation_periods`. Defaults to 1.",
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
					"use the project timezone. A window may cross midnight (for example `22:00` to `06:00`) and must " +
					"span at least 30 minutes. Windows must not overlap each other. Omit the attribute to disable quiet hours.",
				Validators: []validator.Set{
					setvalidator.SizeAtMost(5),
					setvalidator.SizeAtLeast(1),
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

	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		name := model.Name.ValueString()
		req.Name = &name
	}

	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		enabled := model.Enabled.ValueBool()
		req.Enabled = &enabled
	}

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

	if !model.FilterGroupJSON.IsNull() && !model.FilterGroupJSON.IsUnknown() {
		var filterGroup map[string]any
		if err := json.Unmarshal([]byte(model.FilterGroupJSON.ValueString()), &filterGroup); err != nil {
			diags.AddError("Invalid filter_group_json", fmt.Sprintf("Could not parse filter_group_json: %s", err.Error()))
			return req, diags
		}
		req.Filters.FilterGroup = filterGroup
	}

	if !model.ThresholdCount.IsNull() && !model.ThresholdCount.IsUnknown() {
		count := model.ThresholdCount.ValueInt64()
		req.ThresholdCount = &count
	}

	if !model.ThresholdOperator.IsNull() && !model.ThresholdOperator.IsUnknown() {
		operator := model.ThresholdOperator.ValueString()
		req.ThresholdOperator = &operator
	}

	if !model.WindowMinutes.IsNull() && !model.WindowMinutes.IsUnknown() {
		window := model.WindowMinutes.ValueInt64()
		req.WindowMinutes = &window
	}

	if !model.EvaluationPeriods.IsNull() && !model.EvaluationPeriods.IsUnknown() {
		periods := model.EvaluationPeriods.ValueInt64()
		req.EvaluationPeriods = &periods
	}

	if !model.DatapointsToAlarm.IsNull() && !model.DatapointsToAlarm.IsUnknown() {
		datapoints := model.DatapointsToAlarm.ValueInt64()
		req.DatapointsToAlarm = &datapoints
	}

	if !model.CooldownMinutes.IsNull() && !model.CooldownMinutes.IsUnknown() {
		cooldown := model.CooldownMinutes.ValueInt64()
		req.CooldownMinutes = &cooldown
	}

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
	if filterGroup != nil {
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
			"Invalid filter group",
			fmt.Sprintf("Must be a JSON object, got: %s", err.Error()),
		)
		return
	}
	if len(decoded) == 0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Empty filter group",
			"Must be a non-empty JSON object. Omit the attribute entirely to apply no attribute-level filter.",
		)
	}
}

// ModifyResourcePlan enforces the invariants PostHog documents but only checks server-side,
// so a bad config fails at plan time with an actionable message instead of mid-apply.
func (o LogsAlertOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan LogsAlertTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An alert firing on N of M periods can never fire when N > M.
	if isKnown(plan.DatapointsToAlarm) && isKnown(plan.EvaluationPeriods) &&
		plan.DatapointsToAlarm.ValueInt64() > plan.EvaluationPeriods.ValueInt64() {
		resp.Diagnostics.AddAttributeError(
			path.Root("datapoints_to_alarm"),
			"Alert can never fire",
			fmt.Sprintf(
				"datapoints_to_alarm (%d) must not exceed evaluation_periods (%d): the alert requires more breaching "+
					"periods than it ever evaluates, so it would never fire.",
				plan.DatapointsToAlarm.ValueInt64(), plan.EvaluationPeriods.ValueInt64(),
			),
		)
	}

	// PostHog rejects an enabled alert with no filters; catch it before the API does.
	// enabled is Optional+Computed, so an unknown value means the server default (true).
	enabled := plan.Enabled.IsUnknown() || plan.Enabled.IsNull() || plan.Enabled.ValueBool()
	noFilters := isEmptySet(plan.SeverityLevels) && isEmptySet(plan.ServiceNames) &&
		plan.FilterGroupJSON.IsNull()
	if enabled && noFilters && !plan.FilterGroupJSON.IsUnknown() {
		resp.Diagnostics.AddError(
			"Log alert has no filters",
			"An enabled log alert needs at least one of severity_levels, service_names, or filter_group_json. "+
				"Set enabled = false to keep it as a draft.",
		)
	}

	resp.Diagnostics.Append(validateBlockedWindows(ctx, plan.BlockedWindows)...)
}

func isKnown(v types.Int64) bool {
	return !v.IsNull() && !v.IsUnknown()
}

func isEmptySet(v types.Set) bool {
	return v.IsNull() || len(v.Elements()) == 0
}

// validateBlockedWindows enforces the ≥30-minute span and non-overlap rules. PostHog
// silently merges overlapping windows on save, which would come back as a set that
// differs from the config — an apply failure rather than recoverable drift.
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

	type interval struct{ start, end int }
	var intervals []interval
	var labels []string

	for _, w := range parsed {
		if w.Start.IsUnknown() || w.End.IsUnknown() {
			return diags
		}
		start, okStart := parseHHMM(w.Start.ValueString())
		end, okEnd := parseHHMM(w.End.ValueString())
		if !okStart || !okEnd {
			// The HH:MM regex validator already reports malformed values.
			return diags
		}

		span := end - start
		if span <= 0 {
			span += minutesPerDay
		}
		label := fmt.Sprintf("%s-%s", w.Start.ValueString(), w.End.ValueString())
		if start == end || span < minBlockedWindowMinutes {
			diags.AddAttributeError(
				path.Root("blocked_windows"),
				"Quiet-hours window is too short",
				fmt.Sprintf("Window %s spans %d minutes; PostHog requires at least %d.", label, span, minBlockedWindowMinutes),
			)
			continue
		}

		// Split a window that crosses midnight so overlap is a plain interval comparison.
		if end > start {
			intervals = append(intervals, interval{start, end})
			labels = append(labels, label)
		} else {
			intervals = append(intervals, interval{start, minutesPerDay}, interval{0, end})
			labels = append(labels, label, label)
		}
	}

	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if labels[i] == labels[j] {
				continue
			}
			if intervals[i].start < intervals[j].end && intervals[j].start < intervals[i].end {
				diags.AddAttributeError(
					path.Root("blocked_windows"),
					"Quiet-hours windows overlap",
					fmt.Sprintf(
						"Windows %s and %s overlap. PostHog merges overlapping windows when saving, so the alert "+
							"would read back with different windows than configured and every apply would fail.",
						labels[i], labels[j],
					),
				)
				return diags
			}
		}
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
