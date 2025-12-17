package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewAlert() resource.Resource {
	return core.NewGenericResource[AlertResourceTFModel, httpclient.AlertRequest, httpclient.Alert](AlertOps{})
}

type AlertResourceTFModel struct {
	core.BaseStringIdentifiable
	Name                   types.String  `tfsdk:"name"`
	Insight                types.Int64   `tfsdk:"insight"`
	Enabled                types.Bool    `tfsdk:"enabled"`
	SubscribedUsers        types.Set     `tfsdk:"subscribed_users"`
	ThresholdType          types.String  `tfsdk:"threshold_type"`
	ThresholdLower         types.Float64 `tfsdk:"threshold_lower"`
	ThresholdUpper         types.Float64 `tfsdk:"threshold_upper"`
	ConditionType          types.String  `tfsdk:"condition_type"`
	SeriesIndex            types.Int64   `tfsdk:"series_index"`
	CheckOngoingInterval   types.Bool    `tfsdk:"check_ongoing_interval"`
	CalculationInterval    types.String  `tfsdk:"calculation_interval"`
	SkipWeekend            types.Bool    `tfsdk:"skip_weekend"`
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
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Condition type: `absolute_value`, `relative_increase`, or `relative_decrease`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"series_index": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Index of the trend series to monitor (0-based). Used for trends alerts.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
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

	if !model.ConditionType.IsNull() && !model.ConditionType.IsUnknown() {
		req.Condition = &httpclient.AlertCondition{
			Type: model.ConditionType.ValueString(),
		}
	}

	if !model.SeriesIndex.IsNull() && !model.SeriesIndex.IsUnknown() || !model.CheckOngoingInterval.IsNull() && !model.CheckOngoingInterval.IsUnknown() {
		req.Config = &httpclient.TrendsAlertConfig{
			Type: "TrendsAlertConfig",
		}
		if !model.SeriesIndex.IsNull() && !model.SeriesIndex.IsUnknown() {
			seriesIndex := int(model.SeriesIndex.ValueInt64())
			req.Config.SeriesIndex = &seriesIndex
		}
		if !model.CheckOngoingInterval.IsNull() && !model.CheckOngoingInterval.IsUnknown() {
			checkOngoing := model.CheckOngoingInterval.ValueBool()
			req.Config.CheckOngoingInterval = &checkOngoing
		}
	}

	if !model.CalculationInterval.IsNull() && !model.CalculationInterval.IsUnknown() {
		interval := model.CalculationInterval.ValueString()
		req.CalculationInterval = &interval
	}

	if !model.SkipWeekend.IsNull() && !model.SkipWeekend.IsUnknown() {
		skip := model.SkipWeekend.ValueBool()
		req.SkipWeekend = &skip
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

	return diags
}

func (o AlertOps) Create(ctx context.Context, client httpclient.PosthogClient, req httpclient.AlertRequest) (httpclient.Alert, error) {
	return client.CreateAlert(ctx, req)
}

func (o AlertOps) Read(ctx context.Context, client httpclient.PosthogClient, id string) (httpclient.Alert, httpclient.HTTPStatusCode, error) {
	return client.GetAlert(ctx, id)
}

func (o AlertOps) Update(ctx context.Context, client httpclient.PosthogClient, id string, req httpclient.AlertRequest) (httpclient.Alert, httpclient.HTTPStatusCode, error) {
	return client.UpdateAlert(ctx, id, req)
}

func (o AlertOps) Delete(ctx context.Context, client httpclient.PosthogClient, id string) (httpclient.HTTPStatusCode, error) {
	return client.DeleteAlert(ctx, id)
}
