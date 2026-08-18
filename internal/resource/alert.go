package resource

import (
	"context"

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
	core.QuietHoursTimePattern,
	"must be a 24-hour time in HH:MM format",
)

// blockedWindowsValidator adapts this resource's nested window shape onto the shared
// quiet-hours rules. The rules themselves live in core because posthog_logs_alert exposes
// the same windows under a different attribute and must reject exactly the same shapes.
type blockedWindowsValidator struct{}

func (v blockedWindowsValidator) Description(context.Context) string {
	return "blocked windows must not be reshaped by PostHog: they may not overlap or touch, a window " +
		"crossing midnight has to be the only one, and two windows meeting at midnight are only allowed " +
		"alongside a third"
}

func (v blockedWindowsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v blockedWindowsValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	windows, ok := alertQuietHoursWindows(ctx, req.ConfigValue, resp)
	if !ok {
		return
	}
	resp.Diagnostics.Append(core.ValidateQuietHoursWindows(windows, req.Path)...)
}

// alertQuietHoursWindows converts the configured set into the shared window type. Elements
// that are null or unknown are skipped rather than converted: ElementsAs reflects into a
// plain struct, which cannot represent either, and would fail the plan with the framework's
// "report this to the provider developer" error for what is a configuration shape.
func alertQuietHoursWindows(ctx context.Context, set types.Set, resp *validator.SetResponse) ([]core.QuietHoursWindow, bool) {
	var windows []core.QuietHoursWindow
	for _, element := range set.Elements() {
		if element.IsNull() || element.IsUnknown() {
			continue
		}
		object, isObject := element.(types.Object)
		if !isObject {
			continue
		}
		var window AlertBlockedWindowModel
		diags := object.As(ctx, &window, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return nil, false
		}
		if window.Start.IsNull() || window.Start.IsUnknown() || window.End.IsNull() || window.End.IsUnknown() {
			continue
		}
		windows = append(windows, core.QuietHoursWindow{
			Start: window.Start.ValueString(),
			End:   window.End.ValueString(),
		})
	}
	return windows, true
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
						Required: true,
						MarkdownDescription: "Blocked time windows, half-open `[start, end)`. Windows must not " +
							"overlap or touch, except that one may end at `00:00` where another starts. A window " +
							"may wrap midnight (`end` before `start`), but only as the sole window, and two windows " +
							"meeting at midnight are only allowed alongside a third. PostHog enforces its own " +
							"limits on window length and count and reports them on apply. Remove " +
							"`schedule_restriction` to disable quiet hours.",
						Validators: []validator.Set{
							// An empty set is not the same as no quiet hours: PostHog
							// normalizes it to null, which would not match the configured
							// (non-null) block and would fail the apply.
							setvalidator.SizeAtLeast(1),
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
