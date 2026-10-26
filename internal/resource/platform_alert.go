package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

const sourceKindLogs = "logs"

// The defaults PostHog stores when a create leaves a setting out. The provider declares them in
// the schema, so they show in the plan and a PostHog UI edit to one reads as drift.
const (
	defaultPlatformAlertName = "Untitled alert"
	defaultCooldownMinutes   = 0
	defaultMaxInstances      = 100
	maxInstancesCeiling      = 100
	groupingModeSingle       = "single"
	groupingModeByLabels     = "by_result_labels"
	onOverflowDegradeVisible = "degrade_visible"
	defaultWindowMinutes     = 5
)

// The grouping PostHog stores when a create leaves it out: one group for the whole alert.
var defaultPlatformAlertGrouping = httpclient.PlatformAlertGrouping{
	Mode:         groupingModeSingle,
	Keys:         []string{},
	MaxInstances: defaultMaxInstances,
	OnOverflow:   onOverflowDegradeVisible,
}

// PostHog requires the logs condition on this API. These are the logs alerts API's defaults, so an
// alert moved from posthog_logs_alert that never set them plans no change.
var defaultPlatformLogsCondition = httpclient.PlatformAlertLogsCondition{
	ThresholdCount:    defaultThresholdCount,
	ThresholdOperator: defaultThresholdOperator,
	WindowMinutes:     defaultWindowMinutes,
}

// The result labels a logs alert can group by.
var logsGroupingKeys = []string{"service_name", "severity_text"}

// legacyIDPrefix marks a state id that names the source product's alert rather than the platform
// alert. A move writes it, because Terraform does not configure the provider before a move, so the
// mover cannot look up the platform id. The next read resolves it.
const legacyIDPrefix = "legacy:"

var platformAlertGroupingAttrTypes = map[string]attr.Type{
	"mode":          types.StringType,
	"keys":          types.ListType{ElemType: types.StringType},
	"max_instances": types.Int64Type,
	"on_overflow":   types.StringType,
}

var platformLogsConditionAttrTypes = map[string]attr.Type{
	"threshold_count":    types.Int64Type,
	"threshold_operator": types.StringType,
	"window_minutes":     types.Int64Type,
}

func NewPlatformAlert() resource.Resource {
	return core.NewGenericResource[PlatformAlertTFModel, httpclient.PlatformAlertRequest, httpclient.PlatformAlert](
		PlatformAlertOps{},
		core.ProjectScopedImportParser[PlatformAlertTFModel](),
	)
}

type PlatformAlertTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	SourceKind            types.String                  `tfsdk:"source_kind"`
	Name                  types.String                  `tfsdk:"name"`
	Enabled               types.Bool                    `tfsdk:"enabled"`
	EvaluationPeriods     types.Int64                   `tfsdk:"evaluation_periods"`
	DatapointsToAlarm     types.Int64                   `tfsdk:"datapoints_to_alarm"`
	CooldownMinutes       types.Int64                   `tfsdk:"cooldown_minutes"`
	BlockedWindows        types.Set                     `tfsdk:"blocked_windows"`
	Grouping              *PlatformAlertGroupingTFModel `tfsdk:"grouping"`
	Logs                  *PlatformAlertLogsTFModel     `tfsdk:"logs"`
	LegacyConfigurationID types.String                  `tfsdk:"legacy_configuration_id"`
}

type PlatformAlertGroupingTFModel struct {
	Mode         types.String `tfsdk:"mode"`
	Keys         types.List   `tfsdk:"keys"`
	MaxInstances types.Int64  `tfsdk:"max_instances"`
	OnOverflow   types.String `tfsdk:"on_overflow"`
}

type PlatformAlertLogsTFModel struct {
	Filters   *PlatformAlertLogsFiltersTFModel   `tfsdk:"filters"`
	Condition *PlatformAlertLogsConditionTFModel `tfsdk:"condition"`
}

type PlatformAlertLogsFiltersTFModel struct {
	SeverityLevels  types.Set            `tfsdk:"severity_levels"`
	ServiceNames    types.Set            `tfsdk:"service_names"`
	FilterGroupJSON jsontypes.Normalized `tfsdk:"filter_group_json"`
}

type PlatformAlertLogsConditionTFModel struct {
	ThresholdCount    types.Int64  `tfsdk:"threshold_count"`
	ThresholdOperator types.String `tfsdk:"threshold_operator"`
	WindowMinutes     types.Int64  `tfsdk:"window_minutes"`
}

type PlatformAlertOps struct{}

func (o PlatformAlertOps) ResourceName() string {
	return "Platform Alert"
}

func (o PlatformAlertOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an alert on PostHog's shared alerts platform. One resource covers every " +
			"kind of alert, chosen by `source_kind`, and the settings that kind needs go in the block of the same " +
			"name. `logs` is the only kind so far: it counts the log entries matching its filters over a rolling " +
			"window and fires when that count crosses the threshold.\n\n" +
			"~> **One alert, one resource.** Do not manage an alert with this resource and with " +
			"`posthog_logs_alert` at the same time. Both write the same alert, so each apply undoes the other's " +
			"changes. The provider cannot detect this. To move an alert from `posthog_logs_alert`, use a `moved` " +
			"block (see below) and remove the old resource in the same change.\n\n" +
			"The platform alerts API is behind the `platform-alerts` feature flag. Until it is on for your " +
			"organization, every request fails with 403. The API key needs the Alert scope at write, plus the " +
			"scope of each source kind it manages at write (Logs for `logs`). Write includes read. With the " +
			"source kind's scope missing PostHog answers 404, and Terraform then plans to create the alert again.\n\n" +
			"Every optional setting has the default PostHog stores, declared here so it shows in the plan. " +
			"Removing a setting from your configuration returns it to that default, and Terraform corrects an " +
			"edit made in the PostHog UI.\n\n" +
			"An alert with no destination notifies nobody. A `logs` alert keeps a logs alert row beside it, whose " +
			"id is `legacy_configuration_id`, so attach destinations with `posthog_logs_alert_destination` and " +
			"`alert_id = posthog_platform_alert.<name>.legacy_configuration_id`.\n\n" +
			"**Moving from `posthog_logs_alert`.** With Terraform 1.8 and later, a `moved` block from the " +
			"`posthog_logs_alert` address to this one adopts the alert with no destroy (see the example). " +
			"The flat attributes move into the `logs` block: `severity_levels`, `service_names` and " +
			"`filter_group_json` under `logs.filters`, `threshold_count`, `threshold_operator` and " +
			"`window_minutes` under `logs.condition`. The rest keep their names. The alert must be one the " +
			"platform already evaluates, which is true once its project has begun moving logs alerts to the " +
			"platform. Otherwise the plan fails and names the alert. `snooze_until` is not managed by this " +
			"resource, so a snooze stays as it is.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the platform alert. Not the id of the `posthog_logs_alert` it may have moved from; that is `legacy_configuration_id`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"source_kind": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Product whose data the alert evaluates: `logs`. Changing it replaces the alert, " +
					"because PostHog cannot change the kind of an existing alert. Set the block of the same name.",
				Validators: []validator.String{
					stringvalidator.OneOf(sourceKindLogs),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(defaultPlatformAlertName),
				MarkdownDescription: fmt.Sprintf("Human-readable name for the alert. Defaults to `%s`.", defaultPlatformAlertName),
				Validators: []validator.String{
					stringvalidator.RegexMatches(nonBlankPattern, "must not be blank or start or end with whitespace"),
					stringvalidator.LengthAtMost(255),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(defaultEnabled),
				MarkdownDescription: "Whether the alert is evaluated on schedule. Defaults to true. Disabling it ends any firing.",
			},
			"evaluation_periods": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(defaultEvaluationPeriods),
				MarkdownDescription: "How many of the most recent checks to consider. A `logs` alert is checked " +
					"every 5 minutes, so 3 periods cover the last 15 minutes. Must be between 1 and 10. Defaults to 1.",
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
			},
			"datapoints_to_alarm": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(defaultDatapointsToAlarm),
				MarkdownDescription: "How many of the `evaluation_periods` most recent checks must breach before " +
					"the alert fires. Must be between 1 and 10 and must not exceed `evaluation_periods`. Defaults to 1.",
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
			},
			"cooldown_minutes": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(defaultCooldownMinutes),
				MarkdownDescription: "Minimum minutes between notifications for the same alert. Defaults to 0, meaning no cooldown.",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"blocked_windows": schema.SetNestedAttribute{
				Optional: true,
				MarkdownDescription: "Quiet hours: time windows during which the alert does not notify. Times " +
					"use the project timezone. Windows must not overlap or touch each other. A window may cross " +
					"midnight (for example `22:00` to `06:00`), but only as the sole window, because PostHog " +
					"stores a crossing window as two windows when anything else is configured. PostHog enforces " +
					"its own limits on window length and count, and reports them on apply. Omit the attribute, " +
					"or set it to an empty list, to disable quiet hours.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"start": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Start time as `HH:MM` (24-hour, project timezone). Inclusive.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(core.QuietHoursTimePattern, "must be a 24-hour time in HH:MM format"),
							},
						},
						"end": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "End time as `HH:MM` (24-hour, project timezone). Exclusive.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(core.QuietHoursTimePattern, "must be a 24-hour time in HH:MM format"),
							},
						},
					},
				},
			},
			"grouping": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Default:  objectdefault.StaticValue(groupingObject(defaultPlatformAlertGrouping)),
				MarkdownDescription: "How the alert splits its results into groups that fire and resolve on their " +
					"own. Defaults to one group for the whole alert.\n\n" +
					"~> **Changing `keys` drops every existing group silently.** PostHog deletes the groups and " +
					"closes their notification threads without a message, so a group that was firing sends no " +
					"resolve. The PostHog UI warns before this change. Terraform cannot, so check the plan.",
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString(groupingModeSingle),
						MarkdownDescription: "`single` keeps one group for the whole alert. `by_result_labels` keeps " +
							"one group per distinct value of the labels in `keys`. Defaults to `single`.",
						Validators: []validator.String{
							stringvalidator.OneOf(groupingModeSingle, groupingModeByLabels),
						},
					},
					"keys": schema.ListAttribute{
						Optional:    true,
						Computed:    true,
						ElementType: types.StringType,
						Default:     listdefault.StaticValue(stringList(defaultPlatformAlertGrouping.Keys)),
						MarkdownDescription: "Result labels whose values make a group. Empty for `single`, required " +
							"for `by_result_labels`. A `logs` alert can group by `service_name` and `severity_text`, " +
							"and only with `threshold_operator = \"above\"`.",
						Validators: []validator.List{
							listvalidator.UniqueValues(),
						},
					},
					"max_instances": schema.Int64Attribute{
						Optional: true,
						Computed: true,
						Default:  int64default.StaticInt64(defaultMaxInstances),
						MarkdownDescription: fmt.Sprintf("Most groups the alert keeps at once, from 1 to %d. Groups "+
							"past it are not tracked. Defaults to %d.", maxInstancesCeiling, defaultMaxInstances),
						Validators: []validator.Int64{
							int64validator.Between(1, maxInstancesCeiling),
						},
					},
					"on_overflow": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString(onOverflowDegradeVisible),
						MarkdownDescription: "What happens to groups past `max_instances`. `degrade_visible`, the " +
							"only value, says in the last message how many were left out.",
						Validators: []validator.String{
							stringvalidator.OneOf(onOverflowDegradeVisible),
						},
					},
				},
			},
			"logs": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Settings for a `logs` alert. Required when `source_kind = \"logs\"`.",
				Attributes: map[string]schema.Attribute{
					"filters": schema.SingleNestedAttribute{
						Optional: true,
						MarkdownDescription: "The logs the alert counts. An enabled alert needs at least one of " +
							"`severity_levels`, `service_names` or `filter_group_json`. Omit the attribute to save a " +
							"disabled draft with no filters.",
						Attributes: map[string]schema.Attribute{
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
								MarkdownDescription: "Attribute-level filters as JSON, the same value `filter_group_json` takes " +
									"on `posthog_logs_alert`. It stays JSON because a filter group nests property " +
									"filters of many kinds. Must be a non-empty JSON object. Only the " +
									"fields you declare are tracked, because PostHog annotates saved filters with " +
									"defaults (such as `label`). A field you omit is therefore not tracked either.",
								Validators: []validator.String{
									nonEmptyJSONObjectValidator{},
								},
							},
						},
					},
					"condition": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,
						Default:  objectdefault.StaticValue(logsConditionObject(defaultPlatformLogsCondition)),
						MarkdownDescription: "The bound the count is checked against. Defaults to above 100 in 5 " +
							"minutes, the defaults of `posthog_logs_alert`.",
						Attributes: map[string]schema.Attribute{
							"threshold_count": schema.Int64Attribute{
								Optional: true,
								Computed: true,
								Default:  int64default.StaticInt64(defaultPlatformLogsCondition.ThresholdCount),
								MarkdownDescription: "Log entry count to compare against. Defaults to 100. Use `0` " +
									"with `above` to fire on any matching log.",
								Validators: []validator.Int64{
									int64validator.Between(0, 2147483647),
								},
							},
							"threshold_operator": schema.StringAttribute{
								Optional:            true,
								Computed:            true,
								Default:             stringdefault.StaticString(defaultPlatformLogsCondition.ThresholdOperator),
								MarkdownDescription: "Whether the alert fires when the count is `above` or `below` the threshold. Defaults to `above`.",
								Validators: []validator.String{
									stringvalidator.OneOf(thresholdOperatorAbove, thresholdOperatorBelow),
								},
							},
							"window_minutes": schema.Int64Attribute{
								Optional:            true,
								Computed:            true,
								Default:             int64default.StaticInt64(defaultPlatformLogsCondition.WindowMinutes),
								MarkdownDescription: "Time window in minutes over which log entries are counted: `5`, `10`, `15`, `30`, or `60`. Defaults to 5.",
								Validators: []validator.Int64{
									int64validator.OneOf(5, 10, 15, 30, 60),
								},
							},
						},
					},
				},
			},
			"legacy_configuration_id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Id of the source product's own alert behind this one. For `logs`, the id " +
					"`posthog_logs_alert_destination` takes as `alert_id`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o PlatformAlertOps) BuildCreateRequest(ctx context.Context, model PlatformAlertTFModel) (httpclient.PlatformAlertRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.PlatformAlertRequest{
		SourceKind:        model.SourceKind.ValueString(),
		Name:              model.Name.ValueString(),
		Enabled:           model.Enabled.ValueBool(),
		EvaluationPeriods: model.EvaluationPeriods.ValueInt64(),
		DatapointsToAlarm: model.DatapointsToAlarm.ValueInt64(),
		CooldownMinutes:   model.CooldownMinutes.ValueInt64(),
		Grouping:          defaultPlatformAlertGrouping,
	}

	// Defaults fill both blocks in every plan, so the nil checks only guard a hand-built model.
	if model.Grouping != nil {
		req.Grouping.Mode = model.Grouping.Mode.ValueString()
		req.Grouping.MaxInstances = model.Grouping.MaxInstances.ValueInt64()
		req.Grouping.OnOverflow = model.Grouping.OnOverflow.ValueString()
		keys, d := util.StringListToSlice(ctx, model.Grouping.Keys)
		diags.Append(d...)
		if keys != nil {
			req.Grouping.Keys = keys
		}
	}

	schedule, d := blockedWindowsFromSet(ctx, model.BlockedWindows)
	diags.Append(d...)
	req.ScheduleRestriction = schedule

	if model.SourceKind.ValueString() == sourceKindLogs {
		sourceConfig, d := buildLogsSourceConfig(ctx, model.Logs)
		diags.Append(d...)
		req.SourceConfig = sourceConfig
	}

	return req, diags
}

func buildLogsSourceConfig(ctx context.Context, logs *PlatformAlertLogsTFModel) (httpclient.PlatformAlertLogsSourceConfig, diag.Diagnostics) {
	var diags diag.Diagnostics
	config := httpclient.PlatformAlertLogsSourceConfig{Condition: defaultPlatformLogsCondition}
	if logs == nil {
		return config, diags
	}
	if logs.Condition != nil {
		config.Condition = httpclient.PlatformAlertLogsCondition{
			ThresholdCount:    logs.Condition.ThresholdCount.ValueInt64(),
			ThresholdOperator: logs.Condition.ThresholdOperator.ValueString(),
			WindowMinutes:     logs.Condition.WindowMinutes.ValueInt64(),
		}
	}
	if logs.Filters == nil {
		return config, diags
	}

	severityLevels, d := core.ExtractTags(ctx, logs.Filters.SeverityLevels)
	diags.Append(d...)
	serviceNames, d := core.ExtractTags(ctx, logs.Filters.ServiceNames)
	diags.Append(d...)
	filterGroup, d := util.ParseJSONStringMap("logs.filters.filter_group_json", logs.Filters.FilterGroupJSON.StringValue)
	diags.Append(d...)
	config.Filters = httpclient.PlatformAlertLogsFilters{
		SeverityLevels: severityLevels,
		ServiceNames:   serviceNames,
		FilterGroup:    filterGroup,
	}
	return config, diags
}

func (o PlatformAlertOps) BuildUpdateRequest(ctx context.Context, plan, state PlatformAlertTFModel) (httpclient.PlatformAlertRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o PlatformAlertOps) MapResponseToModel(ctx context.Context, resp httpclient.PlatformAlert, model *PlatformAlertTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.SourceKind = core.PtrToStringNullIfEmptyTrimmed(resp.SourceKind)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.Enabled = core.PtrToBool(resp.Enabled)
	model.EvaluationPeriods = util.PtrToInt64(resp.EvaluationPeriods)
	model.DatapointsToAlarm = util.PtrToInt64(resp.DatapointsToAlarm)
	model.CooldownMinutes = util.PtrToInt64(resp.CooldownMinutes)
	model.LegacyConfigurationID = core.PtrToStringNullIfEmptyTrimmed(resp.LegacyConfigurationID)

	diags.Append(platformAlertWarnings(resp)...)

	blockedWindows, d := blockedWindowsToSet(ctx, resp.ScheduleRestriction, model.BlockedWindows)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.BlockedWindows = blockedWindows

	model.Grouping = groupingToModel(resp.Grouping)

	if model.SourceKind.ValueString() == sourceKindLogs {
		logs, d := logsSourceConfigToModel(ctx, resp.SourceConfig, model.Logs)
		diags.Append(d...)
		if diags.HasError() {
			return diags
		}
		model.Logs = logs
	}

	return diags
}

func groupingToModel(grouping *httpclient.PlatformAlertGrouping) *PlatformAlertGroupingTFModel {
	stored := defaultPlatformAlertGrouping
	if grouping != nil {
		stored = *grouping
	}
	return &PlatformAlertGroupingTFModel{
		Mode:         types.StringValue(stored.Mode),
		Keys:         stringList(stored.Keys),
		MaxInstances: types.Int64Value(stored.MaxInstances),
		OnOverflow:   types.StringValue(stored.OnOverflow),
	}
}

// stringList builds a list of strings, empty rather than null for no strings.
func stringList(values []string) types.List {
	elements := make([]attr.Value, len(values))
	for i, value := range values {
		elements[i] = types.StringValue(value)
	}
	return types.ListValueMust(types.StringType, elements)
}

func groupingObject(grouping httpclient.PlatformAlertGrouping) types.Object {
	return types.ObjectValueMust(platformAlertGroupingAttrTypes, map[string]attr.Value{
		"mode":          types.StringValue(grouping.Mode),
		"keys":          stringList(grouping.Keys),
		"max_instances": types.Int64Value(grouping.MaxInstances),
		"on_overflow":   types.StringValue(grouping.OnOverflow),
	})
}

func logsConditionObject(condition httpclient.PlatformAlertLogsCondition) types.Object {
	return types.ObjectValueMust(platformLogsConditionAttrTypes, map[string]attr.Value{
		"threshold_count":    types.Int64Value(condition.ThresholdCount),
		"threshold_operator": types.StringValue(condition.ThresholdOperator),
		"window_minutes":     types.Int64Value(condition.WindowMinutes),
	})
}

// logsSourceConfigToModel reads source_config for a logs alert. current is the model the
// response is mapped onto, which decides whether an absent filter reads as null or empty.
func logsSourceConfigToModel(ctx context.Context, raw json.RawMessage, current *PlatformAlertLogsTFModel) (*PlatformAlertLogsTFModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	var sourceConfig httpclient.PlatformAlertLogsSourceConfig
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &sourceConfig); err != nil {
			diags.AddError("Invalid source_config in response", fmt.Sprintf("Could not decode the logs source_config PostHog returned: %s", err.Error()))
			return current, diags
		}
	}

	logs := &PlatformAlertLogsTFModel{
		Condition: &PlatformAlertLogsConditionTFModel{
			ThresholdCount:    types.Int64Value(sourceConfig.Condition.ThresholdCount),
			ThresholdOperator: types.StringValue(sourceConfig.Condition.ThresholdOperator),
			WindowMinutes:     types.Int64Value(sourceConfig.Condition.WindowMinutes),
		},
	}

	var currentFilters *PlatformAlertLogsFiltersTFModel
	if current != nil {
		currentFilters = current.Filters
	}
	filters := sourceConfig.Filters
	if currentFilters == nil && len(filters.SeverityLevels) == 0 && len(filters.ServiceNames) == 0 && len(filters.FilterGroup) == 0 {
		return logs, diags
	}
	if currentFilters == nil {
		currentFilters = &PlatformAlertLogsFiltersTFModel{
			SeverityLevels:  types.SetNull(types.StringType),
			ServiceNames:    types.SetNull(types.StringType),
			FilterGroupJSON: jsontypes.NewNormalizedNull(),
		}
	}

	severitySet, d := core.TagsToSetPreserveEmpty(ctx, filters.SeverityLevels, currentFilters.SeverityLevels)
	diags.Append(d...)
	serviceSet, d := core.TagsToSetPreserveEmpty(ctx, filters.ServiceNames, currentFilters.ServiceNames)
	diags.Append(d...)
	if diags.HasError() {
		return current, diags
	}

	filterGroupJSON := jsontypes.NewNormalizedNull()
	// Projected onto the declared fields for the reason given in posthog_logs_alert's mapper:
	// PostHog adds its own defaults to a saved filter group.
	if len(filters.FilterGroup) > 0 {
		encoded, err := normalizeJSONForState(filters.FilterGroup, currentFilters.FilterGroupJSON.ValueString())
		if err != nil {
			diags.AddError("Invalid filter group in response", fmt.Sprintf("Could not encode filterGroup returned by PostHog: %s", err.Error()))
			return current, diags
		}
		filterGroupJSON = jsontypes.NewNormalizedValue(encoded)
	}

	logs.Filters = &PlatformAlertLogsFiltersTFModel{
		SeverityLevels:  severitySet,
		ServiceNames:    serviceSet,
		FilterGroupJSON: filterGroupJSON,
	}
	return logs, diags
}

// platformAlertWarnings reports states in which the alert plans clean but either notifies nobody
// or cannot be changed through this resource.
func platformAlertWarnings(resp httpclient.PlatformAlert) diag.Diagnostics {
	var diags diag.Diagnostics
	who := fmt.Sprintf("%q (%s)", strings.TrimSpace(util.PtrToString(resp.Name)), resp.ID)

	if strings.TrimSpace(util.PtrToString(resp.CheckStatus)) == stateBroken {
		diags.AddWarning(
			"Platform alert is not notifying",
			fmt.Sprintf("PostHog marked alert %s broken after repeated failed checks, so it has stopped "+
				"evaluating and notifies nobody. Reset it from the PostHog UI. Terraform does not manage this "+
				"state, so plans stay clean until you do.", who),
		)
	}
	if snoozedUntil(resp.SnoozeUntil, time.Now()) {
		diags.AddWarning(
			"Platform alert is snoozed",
			fmt.Sprintf("Alert %s is snoozed until %s and notifies nobody until then. Terraform does not "+
				"manage snoozing, so plans stay clean meanwhile.", who, util.PtrToString(resp.SnoozeUntil)),
		)
	}
	if owner := strings.TrimSpace(util.PtrToString(resp.Owner)); owner != "" && owner != "platform" {
		diags.AddWarning(
			"Platform alert is managed by its source product",
			fmt.Sprintf("Alert %s is evaluated by the %s product rather than the alerts platform, so PostHog "+
				"refuses changes to it through this resource. Apply will fail until it is moved back to the "+
				"platform.", who, util.PtrToString(resp.SourceKind)),
		)
	}
	return diags
}

// snoozedUntil reports whether a snooze is still running at now. PostHog keeps the timestamp after
// it passes. An unparseable value counts as running, so the warning errs toward showing.
func snoozedUntil(snoozeUntil *string, now time.Time) bool {
	value := strings.TrimSpace(util.PtrToString(snoozeUntil))
	if value == "" {
		return false
	}
	until, err := time.Parse(time.RFC3339, value)
	return err != nil || until.After(now)
}

// ModifyResourcePlan rejects configurations PostHog would refuse, so they fail at plan time with
// a message naming the problem. Every optional attribute has a schema default, so the plan holds
// the value PostHog will receive.
func (o PlatformAlertOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	// Under -refresh=false a moved alert still holds the placeholder id, and the update that resolves
	// it returns the platform id. Planning the id as known would fail that apply as inconsistent.
	if !req.State.Raw.IsNull() && !req.Plan.Raw.Equal(req.State.Raw) {
		var stateID types.String
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &stateID)...)
		if strings.HasPrefix(stateID.ValueString(), legacyIDPrefix) {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("id"), types.StringUnknown())...)
		}
	}
	var plan PlatformAlertTFModel
	// An unknown nested object cannot be read into the model. The API checks it on apply instead.
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		return
	}
	resp.Diagnostics.Append(validatePlatformAlertPlan(ctx, plan)...)
}

func validatePlatformAlertPlan(ctx context.Context, plan PlatformAlertTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if !plan.SourceKind.IsUnknown() && plan.SourceKind.ValueString() == sourceKindLogs && plan.Logs == nil {
		diags.AddAttributeError(path.Root("logs"), "Missing logs settings",
			"source_kind = \"logs\" needs a logs block with the alert's filters and condition.")
	}

	if !plan.DatapointsToAlarm.IsUnknown() && !plan.EvaluationPeriods.IsUnknown() &&
		plan.DatapointsToAlarm.ValueInt64() > plan.EvaluationPeriods.ValueInt64() {
		diags.AddAttributeError(path.Root("datapoints_to_alarm"), "Alert can never fire",
			fmt.Sprintf("datapoints_to_alarm (%d) must not exceed evaluation_periods (%d): the alert requires more "+
				"breaching checks than it ever evaluates, so it would never fire.",
				plan.DatapointsToAlarm.ValueInt64(), plan.EvaluationPeriods.ValueInt64()))
	}

	diags.Append(validateBlockedWindows(ctx, plan.BlockedWindows)...)

	operator := ""
	if plan.Logs != nil && plan.Logs.Condition != nil {
		condition := plan.Logs.Condition
		operator = condition.ThresholdOperator.ValueString()
		if !condition.ThresholdOperator.IsUnknown() && !condition.ThresholdCount.IsUnknown() &&
			operator == thresholdOperatorBelow && condition.ThresholdCount.ValueInt64() == 0 {
			diags.AddAttributeError(path.Root("logs").AtName("condition").AtName("threshold_count"), "Alert can never fire",
				"threshold_operator = \"below\" with threshold_count = 0 can never be satisfied, because a log count is "+
					"never negative. Use threshold_count = 1 to fire when no matching logs arrive in the window, or "+
					"threshold_operator = \"above\" with threshold_count = 0 to fire on any matching log.")
		}
	}

	if plan.Logs != nil && !plan.Enabled.IsUnknown() && plan.Enabled.ValueBool() {
		filters := plan.Logs.Filters
		// An unresolved filter might turn out to be non-empty, so it counts as one and the API gets
		// the final say.
		if filters == nil || (util.IsEmptySet(filters.SeverityLevels) && util.IsEmptySet(filters.ServiceNames) && filters.FilterGroupJSON.IsNull()) {
			diags.AddAttributeError(path.Root("logs").AtName("filters"), "Log alert has no filters",
				"An enabled logs alert needs at least one of severity_levels, service_names, or filter_group_json "+
					"in logs.filters. Set enabled = false to save it without filters.")
		}
	}

	diags.Append(validateGrouping(plan.Grouping, plan.Logs != nil, operator)...)
	return diags
}

func validateGrouping(grouping *PlatformAlertGroupingTFModel, isLogs bool, operator string) diag.Diagnostics {
	var diags diag.Diagnostics
	if grouping == nil || grouping.Mode.IsUnknown() || grouping.Keys.IsUnknown() {
		return diags
	}
	keysPath := path.Root("grouping").AtName("keys")
	elements := grouping.Keys.Elements()
	var keys []string
	for _, element := range elements {
		if value, ok := element.(types.String); ok && !value.IsUnknown() && !value.IsNull() {
			keys = append(keys, value.ValueString())
		}
	}

	switch grouping.Mode.ValueString() {
	case groupingModeSingle:
		if len(elements) > 0 {
			diags.AddAttributeError(keysPath, "Grouping keys need by_result_labels",
				"grouping.mode = \"single\" keeps one group for the whole alert and takes no keys. Set mode = "+
					"\"by_result_labels\" to group by them, or remove keys.")
		}
		return diags
	case groupingModeByLabels:
		if len(elements) == 0 {
			diags.AddAttributeError(keysPath, "Grouping needs keys",
				"grouping.mode = \"by_result_labels\" needs at least one key to group by.")
			return diags
		}
	}

	if !isLogs {
		return diags
	}
	for _, key := range keys {
		if !slices.Contains(logsGroupingKeys, key) {
			diags.AddAttributeError(keysPath, "Unsupported grouping key",
				fmt.Sprintf("A logs alert can group by %s, not %q.", strings.Join(logsGroupingKeys, " and "), key))
		}
	}
	if operator != "" && operator != thresholdOperatorAbove {
		diags.AddAttributeError(keysPath, "Grouped alert must fire above",
			"A grouped logs alert can only fire above a threshold. A group with no matching log returns no "+
				"result, so a service that goes silent is never seen.")
	}
	return diags
}

// logsAlertSourceProviders are the provider addresses a moved posthog_logs_alert may come from. A
// config without required_providers names this provider hashicorp/posthog, as the acceptance tests do.
var logsAlertSourceProviders = map[string]bool{
	"registry.terraform.io/posthog/posthog":   true,
	"registry.terraform.io/hashicorp/posthog": true,
}

// StateMovers accepts a `moved` block from posthog_logs_alert. The mover cannot call the API, so it
// stores the logs alert's id behind legacyIDPrefix, and the read that follows the move resolves it.
func (o PlatformAlertOps) StateMovers() []resource.StateMover {
	logsSchema := LogsAlertOps{}.Schema()
	return []resource.StateMover{
		{
			SourceSchema: &logsSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "posthog_logs_alert" || !logsAlertSourceProviders[req.SourceProviderAddress] {
					return
				}
				var source LogsAlertTFModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &source)...)
				if resp.Diagnostics.HasError() {
					return
				}
				target := platformAlertFromLogsAlert(source)
				if !source.SnoozeUntil.IsNull() && source.SnoozeUntil.ValueString() != "" {
					resp.Diagnostics.AddWarning("snooze_until is not moved",
						fmt.Sprintf("posthog_platform_alert does not manage snoozing, so the snooze on logs alert %s "+
							"stays as it is and Terraform stops changing it.", source.ID.ValueString()))
				}
				resp.Diagnostics.Append(resp.TargetState.Set(ctx, target)...)
			},
		},
	}
}

func platformAlertFromLogsAlert(source LogsAlertTFModel) PlatformAlertTFModel {
	target := PlatformAlertTFModel{
		SourceKind:            types.StringValue(sourceKindLogs),
		Name:                  source.Name,
		Enabled:               source.Enabled,
		EvaluationPeriods:     source.EvaluationPeriods,
		DatapointsToAlarm:     source.DatapointsToAlarm,
		CooldownMinutes:       source.CooldownMinutes,
		BlockedWindows:        source.BlockedWindows,
		LegacyConfigurationID: source.ID,
		Grouping:              groupingToModel(nil),
		Logs: &PlatformAlertLogsTFModel{
			Condition: &PlatformAlertLogsConditionTFModel{
				ThresholdCount:    source.ThresholdCount,
				ThresholdOperator: source.ThresholdOperator,
				WindowMinutes:     source.WindowMinutes,
			},
		},
	}
	target.ID = types.StringValue(legacyIDPrefix + source.ID.ValueString())
	target.ProjectID = source.ProjectID
	if !source.SeverityLevels.IsNull() || !source.ServiceNames.IsNull() || !source.FilterGroupJSON.IsNull() {
		target.Logs.Filters = &PlatformAlertLogsFiltersTFModel{
			SeverityLevels:  source.SeverityLevels,
			ServiceNames:    source.ServiceNames,
			FilterGroupJSON: source.FilterGroupJSON,
		}
	}
	return target
}

// resolvePlatformAlertID returns the platform alert id for a model. When a move left the source
// alert's id in state, it looks the alert up and returns it too. A miss is an error rather than a
// 404: a 404 would drop the alert from state and plan a second alert beside the logs alert.
func resolvePlatformAlertID(ctx context.Context, client httpclient.PosthogClient, model PlatformAlertTFModel) (string, *httpclient.PlatformAlert, error) {
	legacyID, moved := strings.CutPrefix(model.GetID(), legacyIDPrefix)
	if !moved {
		return model.GetID(), nil, nil
	}
	projectID := model.GetEffectiveProjectID()
	alert, found, err := client.FindPlatformAlertByLegacyID(ctx, projectID, legacyID)
	if err != nil {
		return "", nil, fmt.Errorf("looking up the platform alert for logs alert %s: %w", legacyID, err)
	}
	if !found {
		return "", nil, fmt.Errorf("logs alert %s has no platform alert in project %s, so it cannot move to "+
			"posthog_platform_alert yet. Its project has not begun moving logs alerts to the alerts platform. "+
			"Remove the moved block and keep managing it with posthog_logs_alert", legacyID, projectID)
	}
	if owner := util.PtrToString(alert.Owner); owner != "platform" {
		return "", nil, fmt.Errorf("logs alert %s is still evaluated by the logs product (owner %q), so "+
			"PostHog refuses changes to it through posthog_platform_alert. Remove the moved block and keep "+
			"managing it with posthog_logs_alert until its project moves to the alerts platform", legacyID, owner)
	}
	return alert.ID, &alert, nil
}

func (o PlatformAlertOps) Create(ctx context.Context, client httpclient.PosthogClient, model PlatformAlertTFModel, req httpclient.PlatformAlertRequest) (httpclient.PlatformAlert, error) {
	return client.CreatePlatformAlert(ctx, model.GetEffectiveProjectID(), req)
}

func (o PlatformAlertOps) Read(ctx context.Context, client httpclient.PosthogClient, model PlatformAlertTFModel) (httpclient.PlatformAlert, httpclient.HTTPStatusCode, error) {
	id, resolved, err := resolvePlatformAlertID(ctx, client, model)
	if err != nil {
		return httpclient.PlatformAlert{}, 0, err
	}
	if resolved != nil {
		return *resolved, 0, nil
	}
	return client.GetPlatformAlert(ctx, model.GetEffectiveProjectID(), id)
}

func (o PlatformAlertOps) Update(ctx context.Context, client httpclient.PosthogClient, model PlatformAlertTFModel, req httpclient.PlatformAlertRequest) (httpclient.PlatformAlert, httpclient.HTTPStatusCode, error) {
	id, _, err := resolvePlatformAlertID(ctx, client, model)
	if err != nil {
		return httpclient.PlatformAlert{}, 0, err
	}
	return client.UpdatePlatformAlert(ctx, model.GetEffectiveProjectID(), id, req)
}

func (o PlatformAlertOps) Delete(ctx context.Context, client httpclient.PosthogClient, model PlatformAlertTFModel) (httpclient.HTTPStatusCode, error) {
	id, _, err := resolvePlatformAlertID(ctx, client, model)
	if err != nil {
		return 0, err
	}
	return client.DeletePlatformAlert(ctx, model.GetEffectiveProjectID(), id)
}
