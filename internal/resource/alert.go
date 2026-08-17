package resource

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewAlert() resource.Resource {
	return core.NewGenericResource[AlertResourceTFModel, httpclient.AlertRequest, httpclient.Alert](
		AlertOps{},
		core.ProjectScopedImportParser[AlertResourceTFModel](),
	)
}

type AlertResourceTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name                 types.String  `tfsdk:"name"`
	Insight              types.Int64   `tfsdk:"insight"`
	Enabled              types.Bool    `tfsdk:"enabled"`
	SubscribedUsers      types.Set     `tfsdk:"subscribed_users"`
	ThresholdType        types.String  `tfsdk:"threshold_type"`
	ThresholdLower       types.Float64 `tfsdk:"threshold_lower"`
	ThresholdUpper       types.Float64 `tfsdk:"threshold_upper"`
	ConditionType        types.String  `tfsdk:"condition_type"`
	SeriesIndex          types.Int64   `tfsdk:"series_index"`
	CheckOngoingInterval types.Bool    `tfsdk:"check_ongoing_interval"`
	CalculationInterval  types.String  `tfsdk:"calculation_interval"`
	SkipWeekend          types.Bool    `tfsdk:"skip_weekend"`
	ScheduleRestriction  types.Object  `tfsdk:"schedule_restriction"`
}

type AlertBlockedWindowModel struct {
	Start types.String `tfsdk:"start"`
	End   types.String `tfsdk:"end"`
}

type AlertScheduleRestrictionModel struct {
	BlockedWindows types.Set `tfsdk:"blocked_windows"`
}

var alertBlockedWindowObjectType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"start": types.StringType,
		"end":   types.StringType,
	},
}

var alertScheduleRestrictionAttrTypes = map[string]attr.Type{
	"blocked_windows": types.SetType{ElemType: alertBlockedWindowObjectType},
}

var alertTimeOfDayValidator = stringvalidator.RegexMatches(
	regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`),
	"must be a 24-hour time in HH:MM format",
)

const (
	alertMinutesPerDay = 24 * 60
	// The quiet-hour limits PostHog enforces. Referenced by the validator, its error
	// text and the schema description, so they live in one place.
	alertMinBlockedWindowMinutes = 30
	alertMaxBlockedWindows       = 5
)

// blockedWindowsValidator enforces the quiet-hour rules PostHog applies server-side, so a
// bad config fails during plan. PostHog does not store the windows it is given: it lays
// them on a single 24-hour timeline, merges, and re-derives windows from the result. A
// config it would reshape comes back different from the planned one, so the apply fails
// with an inconsistent-result error instead of showing as drift.
type blockedWindowsValidator struct{}

func (v blockedWindowsValidator) Description(context.Context) string {
	return fmt.Sprintf(
		"blocked windows must not overlap or touch, must each span at least %d minutes, "+
			"and a window crossing midnight must be the only one",
		alertMinBlockedWindowMinutes,
	)
}

func (v blockedWindowsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v blockedWindowsValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var windows []AlertBlockedWindowModel
	resp.Diagnostics.Append(req.ConfigValue.ElementsAs(ctx, &windows, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// window is the index of the configured window a span came from. Two spans of the same
	// window are the two halves of one midnight crossing, not a conflict. The label is only
	// message text, so it must not be used to decide identity.
	type interval struct {
		window     int
		label      string
		start, end int
	}

	var spans []interval
	windowCount := 0
	wrappingWindows := 0
	for i, window := range windows {
		start, startOK := alertMinutesSinceMidnight(window.Start)
		end, endOK := alertMinutesSinceMidnight(window.End)
		// Malformed times are the format validator's problem, equal bounds are the API's.
		if !startOK || !endOK || start == end {
			continue
		}
		label := window.Start.ValueString() + "-" + window.End.ValueString()
		length := end - start
		if length < 0 {
			length += alertMinutesPerDay
		}
		if length < alertMinBlockedWindowMinutes {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Blocked window is too short",
				fmt.Sprintf(
					"Window %s spans %d minutes. PostHog requires each blocked window to span at least %d minutes.",
					label, length, alertMinBlockedWindowMinutes,
				),
			)
			return
		}
		windowCount++
		switch {
		case end > start:
			spans = append(spans, interval{i, label, start, end})
		case end == 0:
			// Runs to the end of the day without wrapping into the next morning, so it
			// needs no second half. An empty [0, 0) half would collide with every window
			// starting at midnight in the check below.
			spans = append(spans, interval{i, label, start, alertMinutesPerDay})
		default:
			// Wraps midnight, as the overnight 22:00-07:00 preset does.
			wrappingWindows++
			spans = append(spans, interval{i, label, start, alertMinutesPerDay}, interval{i, label, 0, end})
		}
	}

	// Windows that merely touch count as overlapping. PostHog merges on `next.start <=
	// prev.end`, so 00:00-06:00 and 06:00-09:00 are saved as a single 00:00-09:00 window.
	for i := range spans {
		for j := i + 1; j < len(spans); j++ {
			// The two halves of one wrapped window are not in conflict with each other.
			if spans[i].window == spans[j].window {
				continue
			}
			if spans[i].start <= spans[j].end && spans[j].start <= spans[i].end {
				resp.Diagnostics.AddAttributeError(
					req.Path,
					"Overlapping blocked windows",
					fmt.Sprintf(
						"Windows %s and %s overlap or run straight into each other. PostHog merges them into a single window when saving, so the alert would not match this configuration. Combine them into a single window.",
						spans[i].label, spans[j].label,
					),
				)
				return
			}
		}
	}

	// A window written as crossing midnight always expands to two spans, and PostHog only
	// rejoins them when nothing else is on the timeline. Alongside any other window it
	// stays split, so the alert reads back with one more window than was configured.
	if wrappingWindows > 0 && windowCount > 1 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Blocked window crossing midnight must be the only window",
			"A window whose end is before its start crosses midnight, and PostHog stores it as one window per "+
				"side of midnight unless it is the only window configured. Either make it the only window, or "+
				"split it yourself into a window ending at 00:00 and one starting at 00:00.",
		)
		return
	}

	// Two windows that between them block both sides of midnight get rejoined into a single
	// crossing window, but only while they are the whole timeline. A third window anywhere
	// in the day leaves all of them stored as written.
	if len(spans) == 2 && spans[0].window != spans[1].window &&
		((spans[0].start == 0 && spans[1].end == alertMinutesPerDay) ||
			(spans[1].start == 0 && spans[0].end == alertMinutesPerDay)) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Blocked windows meeting at midnight are stored as one",
			fmt.Sprintf(
				"Windows %s and %s together block midnight from both sides, and PostHog rejoins them into a "+
					"single crossing window when they are the only two, so the alert would not match this "+
					"configuration. Write them as one window crossing midnight, add a third window elsewhere in "+
					"the day, or end the evening window at 23:59.",
				spans[0].label, spans[1].label,
			),
		)
	}
}

func alertMinutesSinceMidnight(value types.String) (int, bool) {
	if value.IsNull() || value.IsUnknown() {
		return 0, false
	}
	var hours, minutes int
	if _, err := fmt.Sscanf(value.ValueString(), "%d:%d", &hours, &minutes); err != nil {
		return 0, false
	}
	return hours*60 + minutes, true
}

type AlertOps struct{}

func (o AlertOps) ResourceName() string {
	return "Alert"
}

func (o AlertOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog alerts. Alerts notify you when an insight's value crosses a threshold.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the alert.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the alert.",
			},
			"insight": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "ID of the insight this alert monitors.",
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the alert is enabled. Defaults to true.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"subscribed_users": schema.SetAttribute{
				Required:            true,
				ElementType:         types.Int64Type,
				MarkdownDescription: "List of user IDs to notify when the alert fires.",
			},
			"threshold_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Type of threshold: `absolute` for fixed values, `percentage` for relative changes.",
			},
			"threshold_lower": schema.Float64Attribute{
				Optional:            true,
				MarkdownDescription: "Lower bound of the threshold. Alert fires when value goes below this.",
			},
			"threshold_upper": schema.Float64Attribute{
				Optional:            true,
				MarkdownDescription: "Upper bound of the threshold. Alert fires when value goes above this.",
			},
			"condition_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Condition type: `absolute_value`, `relative_increase`, or `relative_decrease`.",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"absolute_value",
						"relative_increase",
						"relative_decrease",
					),
				},
			},
			"series_index": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Index of the trend series to monitor (0-based). Used for trends alerts.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"check_ongoing_interval": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to check the ongoing (incomplete) interval. When false, only completed intervals are checked.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"calculation_interval": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "How often to check the alert: `hourly`, `daily`, `weekly`, or `monthly`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"skip_weekend": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to skip checking the alert on weekends.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"schedule_restriction": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Quiet hours: local time windows during which the alert is not evaluated. Times use the project timezone.",
				Attributes: map[string]schema.Attribute{
					"blocked_windows": schema.SetNestedAttribute{
						Required:            true,
						MarkdownDescription: "Blocked time windows, half-open `[start, end)`, each spanning at least 30 minutes. Windows must not overlap or touch. A window may wrap midnight (`end` before `start`), but only as the sole window. Between 1 and 5 windows; remove `schedule_restriction` to disable quiet hours.",
						Validators: []validator.Set{
							// An empty set is not the same as no quiet hours: PostHog
							// normalizes it to null, which would not match the configured
							// (non-null) block and would fail the apply.
							setvalidator.SizeAtLeast(1),
							setvalidator.SizeAtMost(alertMaxBlockedWindows),
							blockedWindowsValidator{},
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"start": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "Start time `HH:MM` (24-hour, project timezone). Inclusive.",
									Validators: []validator.String{
										alertTimeOfDayValidator,
									},
								},
								"end": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "End time `HH:MM` (24-hour, project timezone). Exclusive.",
									Validators: []validator.String{
										alertTimeOfDayValidator,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (o AlertOps) BuildCreateRequest(ctx context.Context, model AlertResourceTFModel) (httpclient.AlertRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.AlertRequest{
		Insight: model.Insight.ValueInt64(),
		Threshold: &httpclient.AlertThreshold{
			Configuration: httpclient.ThresholdConfiguration{
				Type: model.ThresholdType.ValueString(),
			},
		},
	}

	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		name := model.Name.ValueString()
		req.Name = &name
	}

	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		enabled := model.Enabled.ValueBool()
		req.Enabled = &enabled
	}

	if !model.ThresholdLower.IsNull() || !model.ThresholdUpper.IsNull() {
		req.Threshold.Configuration.Bounds = &httpclient.ThresholdBounds{}
		if !model.ThresholdLower.IsNull() && !model.ThresholdLower.IsUnknown() {
			lower := model.ThresholdLower.ValueFloat64()
			req.Threshold.Configuration.Bounds.Lower = &lower
		}
		if !model.ThresholdUpper.IsNull() && !model.ThresholdUpper.IsUnknown() {
			upper := model.ThresholdUpper.ValueFloat64()
			req.Threshold.Configuration.Bounds.Upper = &upper
		}
	}

	req.Condition = &httpclient.AlertCondition{
		Type: model.ConditionType.ValueString(),
	}

	seriesIndex := int(model.SeriesIndex.ValueInt64())
	req.Config = &httpclient.TrendsAlertConfig{
		Type:        "TrendsAlertConfig",
		SeriesIndex: &seriesIndex,
	}
	if !model.CheckOngoingInterval.IsNull() && !model.CheckOngoingInterval.IsUnknown() {
		checkOngoing := model.CheckOngoingInterval.ValueBool()
		req.Config.CheckOngoingInterval = &checkOngoing
	}

	if !model.CalculationInterval.IsNull() && !model.CalculationInterval.IsUnknown() {
		interval := model.CalculationInterval.ValueString()
		req.CalculationInterval = &interval
	}

	if !model.SkipWeekend.IsNull() && !model.SkipWeekend.IsUnknown() {
		skip := model.SkipWeekend.ValueBool()
		req.SkipWeekend = &skip
	}

	// Left nil when the block is absent, which sends an explicit null and clears quiet hours.
	if !model.ScheduleRestriction.IsNull() && !model.ScheduleRestriction.IsUnknown() {
		var restriction AlertScheduleRestrictionModel
		diags.Append(model.ScheduleRestriction.As(ctx, &restriction, basetypes.ObjectAsOptions{})...)
		if diags.HasError() {
			return req, diags
		}

		var windows []AlertBlockedWindowModel
		diags.Append(restriction.BlockedWindows.ElementsAs(ctx, &windows, false)...)
		if diags.HasError() {
			return req, diags
		}

		blockedWindows := make([]httpclient.AlertBlockedWindow, len(windows))
		for i, window := range windows {
			blockedWindows[i] = httpclient.AlertBlockedWindow{
				Start: window.Start.ValueString(),
				End:   window.End.ValueString(),
			}
		}
		req.ScheduleRestriction = &httpclient.AlertScheduleRestriction{BlockedWindows: blockedWindows}
	}

	if !model.SubscribedUsers.IsNull() && !model.SubscribedUsers.IsUnknown() {
		var userIDs []int64
		diags.Append(model.SubscribedUsers.ElementsAs(ctx, &userIDs, false)...)
		if diags.HasError() {
			return req, diags
		}
		req.SubscribedUsers = userIDs
	}

	return req, diags
}

func (o AlertOps) BuildUpdateRequest(ctx context.Context, plan, state AlertResourceTFModel) (httpclient.AlertRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o AlertOps) MapResponseToModel(ctx context.Context, resp httpclient.Alert, model *AlertResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Insight = types.Int64Value(resp.Insight.ID)

	if resp.Name != nil {
		model.Name = types.StringValue(*resp.Name)
	} else {
		model.Name = types.StringNull()
	}

	if resp.Enabled != nil {
		model.Enabled = types.BoolValue(*resp.Enabled)
	} else {
		model.Enabled = types.BoolNull()
	}

	userIDs := make([]int64, len(resp.SubscribedUsers))
	for i, user := range resp.SubscribedUsers {
		userIDs[i] = user.ID
	}
	subscribedUsersSet, d := types.SetValueFrom(ctx, types.Int64Type, userIDs)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.SubscribedUsers = subscribedUsersSet

	if resp.Threshold != nil {
		model.ThresholdType = types.StringValue(resp.Threshold.Configuration.Type)
		if resp.Threshold.Configuration.Bounds != nil {
			if resp.Threshold.Configuration.Bounds.Lower != nil {
				model.ThresholdLower = types.Float64Value(*resp.Threshold.Configuration.Bounds.Lower)
			} else {
				model.ThresholdLower = types.Float64Null()
			}
			if resp.Threshold.Configuration.Bounds.Upper != nil {
				model.ThresholdUpper = types.Float64Value(*resp.Threshold.Configuration.Bounds.Upper)
			} else {
				model.ThresholdUpper = types.Float64Null()
			}
		} else {
			model.ThresholdLower = types.Float64Null()
			model.ThresholdUpper = types.Float64Null()
		}
	}

	if resp.Condition != nil && resp.Condition.Type != "" {
		model.ConditionType = types.StringValue(resp.Condition.Type)
	} else {
		model.ConditionType = types.StringNull()
	}

	if resp.Config != nil && resp.Config.SeriesIndex != nil {
		model.SeriesIndex = types.Int64Value(int64(*resp.Config.SeriesIndex))
	} else {
		model.SeriesIndex = types.Int64Null()
	}

	if resp.Config != nil && resp.Config.CheckOngoingInterval != nil {
		model.CheckOngoingInterval = types.BoolValue(*resp.Config.CheckOngoingInterval)
	} else {
		model.CheckOngoingInterval = types.BoolNull()
	}

	if resp.CalculationInterval != nil {
		model.CalculationInterval = types.StringValue(*resp.CalculationInterval)
	} else {
		model.CalculationInterval = types.StringNull()
	}

	if resp.SkipWeekend != nil {
		model.SkipWeekend = types.BoolValue(*resp.SkipWeekend)
	} else {
		model.SkipWeekend = types.BoolNull()
	}

	// An empty window list means the same thing as no restriction at all. Treating it as a
	// populated object would leave a non-null object against a null config and fail the
	// apply with an inconsistent result.
	if resp.ScheduleRestriction == nil || len(resp.ScheduleRestriction.BlockedWindows) == 0 {
		model.ScheduleRestriction = types.ObjectNull(alertScheduleRestrictionAttrTypes)
	} else {
		windows := make([]AlertBlockedWindowModel, len(resp.ScheduleRestriction.BlockedWindows))
		for i, window := range resp.ScheduleRestriction.BlockedWindows {
			windows[i] = AlertBlockedWindowModel{
				Start: types.StringValue(window.Start),
				End:   types.StringValue(window.End),
			}
		}
		windowSet, d := types.SetValueFrom(ctx, alertBlockedWindowObjectType, windows)
		diags.Append(d...)
		if diags.HasError() {
			return diags
		}
		restriction, d := types.ObjectValueFrom(ctx, alertScheduleRestrictionAttrTypes, AlertScheduleRestrictionModel{
			BlockedWindows: windowSet,
		})
		diags.Append(d...)
		model.ScheduleRestriction = restriction
	}

	return diags
}

func (o AlertOps) Create(ctx context.Context, client httpclient.PosthogClient, model AlertResourceTFModel, req httpclient.AlertRequest) (httpclient.Alert, error) {
	return client.CreateAlert(ctx, model.GetEffectiveProjectID(), req)
}

func (o AlertOps) Read(ctx context.Context, client httpclient.PosthogClient, model AlertResourceTFModel) (httpclient.Alert, httpclient.HTTPStatusCode, error) {
	return client.GetAlert(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o AlertOps) Update(ctx context.Context, client httpclient.PosthogClient, model AlertResourceTFModel, req httpclient.AlertRequest) (httpclient.Alert, httpclient.HTTPStatusCode, error) {
	return client.UpdateAlert(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o AlertOps) Delete(ctx context.Context, client httpclient.PosthogClient, model AlertResourceTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteAlert(ctx, model.GetEffectiveProjectID(), model.GetID())
}
