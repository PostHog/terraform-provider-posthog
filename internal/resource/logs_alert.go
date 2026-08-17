package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
)

// hhmmPattern matches a 24-hour HH:MM local time, as used by quiet-hours windows.
var hhmmPattern = regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)

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
			"PostHog UI. An alert with no destination still evaluates, but notifies nobody.",
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
				Optional:            true,
				MarkdownDescription: "Human-readable name for the alert. PostHog defaults this to `Untitled alert`.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
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
					"[logs alerts API](https://posthog.com/docs/api/logs). Compared semantically, so key ordering and " +
					"whitespace differences from the PostHog API do not produce a diff. Use this for anything beyond " +
					"severity and service, such as filtering on a log attribute.",
			},
			"threshold_count": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of matching log entries that breaches the threshold within the window. Defaults to 100. Use `0` with the `above` operator to fire on any matching log.",
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
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Total number of check periods in the sliding evaluation window (the `M` in N-of-M). Defaults to 1.",
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"datapoints_to_alarm": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "How many periods within the evaluation window must breach the threshold to fire (the `N` in N-of-M). Defaults to 1.",
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
				MarkdownDescription: "Quiet hours: local time windows during which the alert must not run. Times use the " +
					"project timezone. Each window must span at least 30 minutes. Omit the attribute to disable quiet " +
					"hours.\n\n" +
					"~> PostHog merges overlapping or identical windows when saving. Declare non-overlapping windows, " +
					"otherwise the merged result read back from the API will differ from your configuration on every plan.",
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
				Computed:            true,
				MarkdownDescription: "Current evaluation state of the alert, such as `not_firing`, `firing`, or `snoozed`.",
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

	if filterGroup != nil {
		encoded, err := json.Marshal(filterGroup)
		if err != nil {
			diags.AddError("Invalid filter group in response", fmt.Sprintf("Could not encode filterGroup returned by PostHog: %s", err.Error()))
			return diags
		}
		model.FilterGroupJSON = jsontypes.NewNormalizedValue(string(encoded))
	} else {
		model.FilterGroupJSON = jsontypes.NewNormalizedNull()
	}

	model.ThresholdCount = int64PtrToValue(resp.ThresholdCount)
	model.ThresholdOperator = core.PtrToStringNullIfEmptyTrimmed(resp.ThresholdOperator)
	model.WindowMinutes = int64PtrToValue(resp.WindowMinutes)
	model.EvaluationPeriods = int64PtrToValue(resp.EvaluationPeriods)
	model.DatapointsToAlarm = int64PtrToValue(resp.DatapointsToAlarm)
	model.CooldownMinutes = int64PtrToValue(resp.CooldownMinutes)

	blockedWindows, d := blockedWindowsToSet(ctx, resp.ScheduleRestriction)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.BlockedWindows = blockedWindows

	return diags
}

func int64PtrToValue(v *int64) types.Int64 {
	if v == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*v)
}

// blockedWindowsToSet converts the API's schedule_restriction into the flattened
// blocked_windows set. A null or empty restriction means quiet hours are off.
func blockedWindowsToSet(ctx context.Context, schedule *httpclient.LogsAlertSchedule) (types.Set, diag.Diagnostics) {
	objectType := types.ObjectType{AttrTypes: blockedWindowAttrTypes}
	if schedule == nil || len(schedule.BlockedWindows) == 0 {
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
