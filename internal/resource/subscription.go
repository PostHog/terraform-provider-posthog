package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

func NewSubscription() resource.Resource {
	return core.NewGenericResource[SubscriptionResourceTFModel, httpclient.SubscriptionRequest, httpclient.Subscription](
		SubscriptionOps{},
		core.ProjectScopedImportParser[SubscriptionResourceTFModel](),
	)
}

type SubscriptionResourceTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	TargetType              types.String         `tfsdk:"target_type"`
	TargetValue             types.String         `tfsdk:"target_value"`
	IntegrationID           types.Int64          `tfsdk:"integration_id"`
	DashboardID             types.Int64          `tfsdk:"dashboard_id"`
	InsightID               types.Int64          `tfsdk:"insight_id"`
	AIPrompt                types.String         `tfsdk:"ai_prompt"`
	AIPromptConfig          jsontypes.Normalized `tfsdk:"ai_prompt_config"`
	SummaryEnabled          types.Bool           `tfsdk:"summary_enabled"`
	SummaryPromptGuide      types.String         `tfsdk:"summary_prompt_guide"`
	DashboardExportInsights types.Set            `tfsdk:"dashboard_export_insights"`
	Frequency               types.String         `tfsdk:"frequency"`
	Interval                types.Int64          `tfsdk:"interval"`
	StartDate               types.String         `tfsdk:"start_date"`
	ByWeekday               types.Set            `tfsdk:"byweekday"`
	BySetPos                types.Int64          `tfsdk:"bysetpos"`
	Enabled                 types.Bool           `tfsdk:"enabled"`
	Title                   types.String         `tfsdk:"title"`
	ResourceType            types.String         `tfsdk:"resource_type"`
	Summary                 types.String         `tfsdk:"summary"`
	NextDeliveryDate        types.String         `tfsdk:"next_delivery_date"`
	CreatedAt               types.String         `tfsdk:"created_at"`
}

type SubscriptionOps struct{}

func (o SubscriptionOps) ResourceName() string {
	return "Subscription"
}

func (o SubscriptionOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog subscriptions. A subscription delivers a recurring dashboard or insight " +
			"digest to email or Slack on an rrule schedule.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Numeric ID of the subscription.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"target_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Delivery channel type: `email` or `slack`.",
				Validators: []validator.String{
					stringvalidator.OneOf("email", "slack"),
				},
			},
			"target_value": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Delivery target. For `email` a plain address (or comma-separated addresses). " +
					"For `slack` the channel in the form `<channelId>|#<channel-name>` " +
					"(e.g. `C0B9A53J8RF|#reports`). Stored verbatim; the API does not parse the pipe/#name suffix.",
			},
			"integration_id": schema.Int64Attribute{
				Optional: true,
				MarkdownDescription: "ID of the Slack integration to deliver through. Required when `target_type` is " +
					"`slack`, and must be omitted for `email`. To find it, connect Slack under " +
					"**Settings → Integrations** in the PostHog UI, then list this project's integrations via " +
					"`GET /api/projects/<project_id>/integrations/` and use the `id` of the entry whose " +
					"`kind` is `slack`.",
			},
			"dashboard_id": schema.Int64Attribute{
				Optional: true,
				MarkdownDescription: "ID of the dashboard to send a digest for. Exactly one of `dashboard_id`, " +
					"`insight_id`, or `ai_prompt` must be set. Changing which subject is targeted forces replacement.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(path.MatchRoot("insight_id"), path.MatchRoot("ai_prompt")),
				},
			},
			"insight_id": schema.Int64Attribute{
				Optional: true,
				MarkdownDescription: "ID of a single insight to subscribe to. Exactly one of `dashboard_id`, " +
					"`insight_id`, or `ai_prompt` must be set. Changing which subject is targeted forces replacement.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"ai_prompt": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "Natural-language prompt for an AI-summary subscription. PostHog runs the prompt " +
					"over your project data and delivers the result on the schedule. Exactly one of `dashboard_id`, " +
					"`insight_id`, or `ai_prompt` must be set. The prompt text itself can be edited in place; " +
					"switching to or from a dashboard/insight subject forces replacement.",
			},
			"ai_prompt_config": schema.StringAttribute{
				Optional:   true,
				CustomType: jsontypes.NormalizedType{},
				MarkdownDescription: "Optional JSON config for an `ai_prompt` subscription, e.g. the analysis window " +
					"(`{\"window\":{\"mode\":\"last_n_days\",\"start_days_ago\":7,\"end_days_ago\":null}}`). Compared " +
					"semantically, so key ordering and whitespace do not produce a diff. Only meaningful when `ai_prompt` is set.",
			},
			"summary_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to attach an AI-generated summary to the delivery. Defaults to false.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"summary_prompt_guide": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Optional free-text guidance steering the AI summary (only used when `summary_enabled` is true).",
			},
			"dashboard_export_insights": schema.SetAttribute{
				Optional:    true,
				ElementType: types.Int64Type,
				MarkdownDescription: "Insight IDs from the dashboard to include in the digest. Required and non-empty " +
					"when `dashboard_id` is set (the API rejects an empty selection). Ignored for insight subscriptions.",
			},
			"frequency": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Recurrence unit: `daily`, `weekly`, `monthly`, or `yearly`.",
				Validators: []validator.String{
					stringvalidator.OneOf("daily", "weekly", "monthly", "yearly"),
				},
			},
			"interval": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Every N of the frequency unit (e.g. `interval = 2` with `frequency = \"weekly\"` is every two weeks).",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"start_date": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Required RFC3339 datetime that anchors the delivery schedule " +
					"(the recurrence's `dtstart`). It sets *when in the cycle* deliveries land — the " +
					"time of day for `daily`, and the weekday/day-of-month plus time for " +
					"`weekly`/`monthly`. PostHog requires it and has no server default, so it must " +
					"always be set. Example: `2026-08-17T09:00:00Z` delivers at 09:00 UTC.",
			},
			"byweekday": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Lowercase weekday names (`monday`..`sunday`) used with weekly/monthly frequency.",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf(
						"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday",
					)),
				},
			},
			"bysetpos": schema.Int64Attribute{
				Optional: true,
				MarkdownDescription: "Selects *which* occurrence within the period to use (rrule `BYSETPOS`), " +
					"used together with `byweekday` on a `monthly` frequency. For example " +
					"`byweekday = [\"monday\"]` with `bysetpos = 1` means the **first Monday of each " +
					"month**, and `bysetpos = -1` means the **last Monday**. Positive values count from " +
					"the start of the period, negative from the end. Leave unset for simple " +
					"daily/weekly schedules.",
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the subscription is active. Defaults to true. Set to false to pause without deleting.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"title": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Optional human-readable label for the subscription.",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Server-inferred subject type: `dashboard`, `insight`, or `ai_prompt`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"summary": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Server-generated human-readable schedule summary (e.g. `sent every week on Monday`). " +
					"Recomputed whenever the schedule changes.",
				// No UseStateForUnknown: the summary is derived from the schedule fields, so it must
				// become unknown (and be recomputed) when any of them change on update. Freezing it to
				// the prior state would produce an "inconsistent result after apply" when the schedule
				// is edited (e.g. clearing byweekday/bysetpos).
			},
			"next_delivery_date": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Computed timestamp of the next scheduled delivery.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the subscription was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o SubscriptionOps) BuildCreateRequest(ctx context.Context, model SubscriptionResourceTFModel) (httpclient.SubscriptionRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.SubscriptionRequest{
		TargetType:  model.TargetType.ValueString(),
		TargetValue: model.TargetValue.ValueString(),
		Frequency:   model.Frequency.ValueString(),
		// interval is required with no server default; send unconditionally.
		Interval:  model.Interval.ValueInt64(),
		StartDate: model.StartDate.ValueString(),
	}

	req.IntegrationID = util.Int64PtrFromValue(model.IntegrationID)
	req.Dashboard = util.Int64PtrFromValue(model.DashboardID)
	req.Insight = util.Int64PtrFromValue(model.InsightID)
	req.Prompt = util.StringPtrFromValue(model.AIPrompt)
	req.AIPromptConfig = rawFromNormalized(model.AIPromptConfig)
	req.SummaryEnabled = util.BoolPtrFromValue(model.SummaryEnabled)
	req.SummaryPromptGuide = util.StringPtrFromValue(model.SummaryPromptGuide)

	// Default the slices to a non-nil empty slice: the request struct drops ,omitempty on
	// these fields so an emptied set clears on update, but the API rejects a null slice on
	// create. [] both creates-as-unset and clears-on-update. (An empty selection on a
	// dashboard sub is still rejected server-side, which is the intended validation error.)
	ids := []int64{}
	if !model.DashboardExportInsights.IsNull() && !model.DashboardExportInsights.IsUnknown() {
		diags.Append(model.DashboardExportInsights.ElementsAs(ctx, &ids, false)...)
		if diags.HasError() {
			return req, diags
		}
	}
	req.DashboardExportInsights = ids

	days := []string{}
	if !model.ByWeekday.IsNull() && !model.ByWeekday.IsUnknown() {
		diags.Append(model.ByWeekday.ElementsAs(ctx, &days, false)...)
		if diags.HasError() {
			return req, diags
		}
	}
	req.ByWeekday = days

	req.BySetPos = util.Int64PtrFromValue(model.BySetPos)
	req.Enabled = util.BoolPtrFromValue(model.Enabled)
	req.Title = util.StringPtrFromValue(model.Title)

	return req, diags
}

func (o SubscriptionOps) BuildUpdateRequest(ctx context.Context, plan, state SubscriptionResourceTFModel) (httpclient.SubscriptionRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	// integration_id, bysetpos, byweekday and dashboard_export_insights clear on removal
	// without extra work here: their request-struct fields drop ,omitempty, so BuildCreateRequest
	// already emits an explicit null (int64s) or [] (slices) whenever the plan value is null.
	// Only the string fields below need the pointer-to-"" trick, since ,omitempty is kept on them.

	// Clear the optional title if it was removed from config.
	if core.ShouldClearString(plan.Title, state.Title) {
		req.Title = util.StringPtr("")
	}

	// Clear summary_prompt_guide if it was removed from config.
	if core.ShouldClearString(plan.SummaryPromptGuide, state.SummaryPromptGuide) {
		req.SummaryPromptGuide = util.StringPtr("")
	}

	// Clear ai_prompt_config if it was removed from config. Send an explicit empty object;
	// a nil RawMessage would be dropped by omitempty and leave the old config in place.
	if plan.AIPromptConfig.IsNull() && !state.AIPromptConfig.IsNull() {
		req.AIPromptConfig = json.RawMessage("{}")
	}

	return req, diags
}

func (o SubscriptionOps) MapResponseToModel(ctx context.Context, resp httpclient.Subscription, model *SubscriptionResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.TargetType = types.StringValue(resp.TargetType)
	// target_value stored verbatim to avoid drift.
	model.TargetValue = types.StringValue(resp.TargetValue)
	model.Frequency = types.StringValue(resp.Frequency)
	model.Interval = types.Int64Value(resp.Interval)

	model.IntegrationID = util.PtrToInt64(resp.IntegrationID)
	model.DashboardID = util.PtrToInt64(resp.Dashboard)
	model.InsightID = util.PtrToInt64(resp.Insight)
	model.BySetPos = util.PtrToInt64(resp.BySetPos)

	model.Enabled = core.PtrToBool(resp.Enabled)
	model.Title = core.PtrToStringNullIfEmptyTrimmed(resp.Title)

	model.AIPrompt = core.PtrToStringNullIfEmptyTrimmed(resp.Prompt)
	model.SummaryEnabled = core.PtrToBool(resp.SummaryEnabled)
	model.SummaryPromptGuide = core.PtrToStringNullIfEmptyTrimmed(resp.SummaryPromptGuide)
	model.AIPromptConfig = aiPromptConfigToModel(resp.AIPromptConfig, model.AIPromptConfig)

	model.ResourceType = core.PtrToStringNullIfEmptyTrimmed(resp.ResourceType)
	model.Summary = core.PtrToStringNullIfEmptyTrimmed(resp.Summary)
	model.NextDeliveryDate = core.PtrToStringNullIfEmptyTrimmed(resp.NextDeliveryDate)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)

	// Reconcile start_date: the API re-emits the datetime in a possibly-different
	// ISO-8601 format, so we can't store it verbatim (perpetual diff) nor blindly
	// normalize it (start_date is Required, so a non-UTC config value would then
	// differ from plan and trip an inconsistent-result error).
	model.StartDate = reconcileStartDate(model.StartDate, resp.StartDate, &diags)

	exportSet, d := core.Int64SetPreserveEmpty(ctx, resp.DashboardExportInsights, model.DashboardExportInsights)
	diags.Append(d...)
	model.DashboardExportInsights = exportSet

	weekdaySet, d := core.TagsToSetPreserveEmpty(ctx, resp.ByWeekday, model.ByWeekday)
	diags.Append(d...)
	model.ByWeekday = weekdaySet

	return diags
}

// aiPromptConfigToModel maps the API's ai_prompt_config blob to state. It maps
// {}/empty/null to a null value so an unset config does not perpetually diff,
// then defers to normalizeRawForState (the same helper experiment.go uses) so
// the value is filtered to the user's configured keys — this keeps drift
// protection consistent if PostHog ever echoes server-added keys.
func aiPromptConfigToModel(raw json.RawMessage, current jsontypes.Normalized) jsontypes.Normalized {
	trimmed := strings.TrimSpace(string(raw))
	if trimmed == "" || trimmed == "{}" || trimmed == "null" {
		return jsontypes.NewNormalizedNull()
	}
	return normalizeRawForState(json.RawMessage(trimmed), current)
}

// reconcileStartDate keeps the user's configured start_date when it denotes the
// same instant as the API's value, so a non-UTC (but valid RFC3339) config value
// neither perpetually diffs nor trips an inconsistent-result error on this
// Required attribute. When they represent different instants (a genuine drift),
// it adopts the API's normalized value.
func reconcileStartDate(config types.String, apiValue string, diags *diag.Diagnostics) types.String {
	apiState := normalizeDateTimeToModel(apiValue, diags)
	if !config.IsNull() && !config.IsUnknown() && !apiState.IsNull() {
		if cfgNorm, err := core.NormalizeRFC3339(config.ValueString()); err == nil && cfgNorm == apiState.ValueString() {
			return config
		}
	}
	return apiState
}

// normalizeDateTimeToModel canonicalizes an RFC3339 datetime for state. On a parse
// failure it falls back to the raw API value and records a warning rather than failing
// the apply.
func normalizeDateTimeToModel(apiValue string, diags *diag.Diagnostics) types.String {
	normalized, err := core.NormalizeRFC3339(apiValue)
	if err != nil {
		diags.AddWarning(
			"Unparseable datetime from API",
			"Could not normalize datetime "+apiValue+" as RFC3339; storing it verbatim: "+err.Error(),
		)
		return types.StringValue(apiValue)
	}
	return types.StringValue(normalized)
}

func (o SubscriptionOps) Create(ctx context.Context, client httpclient.PosthogClient, model SubscriptionResourceTFModel, req httpclient.SubscriptionRequest) (httpclient.Subscription, error) {
	return client.CreateSubscription(ctx, model.GetEffectiveProjectID(), req)
}

func (o SubscriptionOps) Read(ctx context.Context, client httpclient.PosthogClient, model SubscriptionResourceTFModel) (httpclient.Subscription, httpclient.HTTPStatusCode, error) {
	sub, code, err := client.GetSubscription(ctx, model.GetEffectiveProjectID(), model.GetID())
	if err != nil {
		return sub, code, err
	}
	// A soft-deleted subscription is still returned by GET (200) with deleted=true; only the
	// list endpoint filters it out. Surface it as not-found so the generic Read drops it from
	// state and a subsequent plan recreates it.
	if sub.Deleted != nil && *sub.Deleted {
		return sub, http.StatusNotFound, fmt.Errorf("subscription %d is soft-deleted", sub.ID)
	}
	return sub, code, nil
}

func (o SubscriptionOps) Update(ctx context.Context, client httpclient.PosthogClient, model SubscriptionResourceTFModel, req httpclient.SubscriptionRequest) (httpclient.Subscription, httpclient.HTTPStatusCode, error) {
	return client.UpdateSubscription(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o SubscriptionOps) Delete(ctx context.Context, client httpclient.PosthogClient, model SubscriptionResourceTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteSubscription(ctx, model.GetEffectiveProjectID(), model.GetID())
}
