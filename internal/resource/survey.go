package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
	"github.com/posthog/terraform-provider/internal/util"
)

func NewSurvey() resource.Resource {
	return core.NewGenericResource[SurveyTFModel, httpclient.SurveyRequest, httpclient.Survey](
		SurveyOps{},
		core.ProjectScopedImportParser[SurveyTFModel](),
	)
}

type SurveyTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name                         types.String `tfsdk:"name"`
	Description                  types.String `tfsdk:"description"`
	Type                         types.String `tfsdk:"type"`
	Schedule                     types.String `tfsdk:"schedule"`
	LinkedFlagID                 types.Int64  `tfsdk:"linked_flag_id"`
	LinkedInsightID              types.Int64  `tfsdk:"linked_insight_id"`
	TargetingFlagID              types.Int64  `tfsdk:"targeting_flag_id"`
	TargetingFlagFiltersJSON     types.String `tfsdk:"targeting_flag_filters_json"`
	QuestionsJSON                types.String `tfsdk:"questions_json"`
	ConditionsJSON               types.String `tfsdk:"conditions_json"`
	AppearanceJSON               types.String `tfsdk:"appearance_json"`
	StartDate                    types.String `tfsdk:"start_date"`
	EndDate                      types.String `tfsdk:"end_date"`
	Archived                     types.Bool   `tfsdk:"archived"`
	ResponsesLimit               types.Int64  `tfsdk:"responses_limit"`
	IterationCount               types.Int64  `tfsdk:"iteration_count"`
	IterationFrequencyDays       types.Int64  `tfsdk:"iteration_frequency_days"`
	ResponseSamplingStartDate    types.String `tfsdk:"response_sampling_start_date"`
	ResponseSamplingIntervalType types.String `tfsdk:"response_sampling_interval_type"`
	ResponseSamplingInterval     types.Int64  `tfsdk:"response_sampling_interval"`
	ResponseSamplingLimit        types.Int64  `tfsdk:"response_sampling_limit"`
	ResponseSamplingDailyLimits  types.String `tfsdk:"response_sampling_daily_limits_json"`
	EnablePartialResponses       types.Bool   `tfsdk:"enable_partial_responses"`
	EnableIframeEmbedding        types.Bool   `tfsdk:"enable_iframe_embedding"`
	TranslationsJSON             types.String `tfsdk:"translations_json"`
	CreateInFolder               types.String `tfsdk:"create_in_folder"`
	FormContentJSON              types.String `tfsdk:"form_content_json"`
	CreatedAt                    types.String `tfsdk:"created_at"`
	CreatedByJSON                types.String `tfsdk:"created_by_json"`
	LinkedFlagJSON               types.String `tfsdk:"linked_flag_json"`
	TargetingFlagJSON            types.String `tfsdk:"targeting_flag_json"`
	InternalTargetingFlagJSON    types.String `tfsdk:"internal_targeting_flag_json"`
}

type SurveyOps struct{}

func (o SurveyOps) ResourceName() string {
	return "Survey"
}

func (o SurveyOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog surveys via the project-scoped surveys API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the survey.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Survey name.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Survey description.",
			},
			"type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Survey type. Supported values are `popover`, `widget`, `external_survey`, and `api`.",
			},
			"schedule": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Survey schedule. Supported values are `once`, `recurring`, and `always`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"linked_flag_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Feature flag ID linked to the survey. Preserved across refreshes when explicitly configured.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"linked_insight_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Insight ID linked to the survey. This field is write-only in the PostHog API and is preserved from Terraform state when configured.",
			},
			"targeting_flag_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Existing targeting feature flag ID to use for this survey.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"targeting_flag_filters_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object describing targeting flag filters. This input is write-only in the PostHog API and is preserved from Terraform state when configured.",
			},
			"questions_json": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "JSON array describing the survey questions.",
			},
			"conditions_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object describing survey display and targeting conditions.",
			},
			"appearance_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object describing survey appearance settings.",
			},
			"start_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "RFC3339 start date for the survey.",
			},
			"end_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "RFC3339 end date for the survey.",
			},
			"archived": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the survey is archived.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"responses_limit": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Maximum number of responses allowed for the survey.",
			},
			"iteration_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Number of survey recurrences when `schedule` is `recurring`.",
			},
			"iteration_frequency_days": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Number of days between recurrences when `schedule` is `recurring`.",
			},
			"response_sampling_start_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "RFC3339 date when response sampling starts.",
			},
			"response_sampling_interval_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Response sampling interval type. Supported values are `day`, `week`, and `month`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"response_sampling_interval": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Response sampling interval value.",
			},
			"response_sampling_limit": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Maximum responses allowed during each sampling interval.",
			},
			"response_sampling_daily_limits_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object describing daily response sampling limits.",
			},
			"enable_partial_responses": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether partial responses are stored when a respondent exits early.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enable_iframe_embedding": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the survey is embeddable in an iframe.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"translations_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object describing translated survey content.",
			},
			"create_in_folder": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Folder identifier used only during survey creation.",
			},
			"form_content_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object describing custom form content.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "RFC3339 creation timestamp of the survey.",
			},
			"created_by_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "JSON object describing the survey creator.",
			},
			"linked_flag_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "JSON object describing the linked feature flag returned by the API.",
			},
			"targeting_flag_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "JSON object describing the targeting feature flag returned by the API.",
			},
			"internal_targeting_flag_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "JSON object describing the internal targeting feature flag returned by the API.",
			},
		},
	}
}

func (o SurveyOps) BuildCreateRequest(ctx context.Context, model SurveyTFModel) (httpclient.SurveyRequest, diag.Diagnostics) {
	return o.buildRequest(ctx, model, SurveyTFModel{})
}

func (o SurveyOps) BuildUpdateRequest(ctx context.Context, plan, state SurveyTFModel) (httpclient.SurveyRequest, diag.Diagnostics) {
	return o.buildRequest(ctx, plan, state)
}

func (o SurveyOps) buildRequest(ctx context.Context, plan, state SurveyTFModel) (httpclient.SurveyRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.SurveyRequest{
		Name: plan.Name.ValueString(),
		Type: plan.Type.ValueString(),
	}

	applySurveyDescription(plan, state, &req)
	applySurveyRelationships(plan, &req)

	req.Questions, diags = parseSurveyQuestions(plan.QuestionsJSON)
	if diags.HasError() {
		return req, diags
	}

	diags = appendSurveyJSONFields(plan, state, &req)
	if diags.HasError() {
		return req, diags
	}

	applySurveyTimingFields(plan, &req)
	applySurveySamplingFields(plan, &req)
	applySurveyBooleanFields(plan, &req)
	applySurveyTargetingClear(plan, state, &req)

	return req, diags
}

func applySurveyDescription(plan, state SurveyTFModel, req *httpclient.SurveyRequest) {
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		description := plan.Description.ValueString()
		req.Description = &description
		return
	}

	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = util.StringPtr("")
	}
}

func applySurveyRelationships(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.LinkedFlagID = int64PtrFromValue(plan.LinkedFlagID)
	req.LinkedInsightID = int64PtrFromValue(plan.LinkedInsightID)
	req.TargetingFlagID = int64PtrFromValue(plan.TargetingFlagID)
	req.Schedule = stringPtrFromValue(plan.Schedule)
}

func parseSurveyQuestions(attr types.String) ([]interface{}, diag.Diagnostics) {
	if attr.IsNull() || attr.IsUnknown() {
		return nil, nil
	}

	return parseJSONArrayAttribute(attr, "questions_json")
}

func appendSurveyJSONFields(plan, state SurveyTFModel, req *httpclient.SurveyRequest) diag.Diagnostics {
	var diags diag.Diagnostics

	diags.Append(assignJSONAttribute(plan.ConditionsJSON, "conditions_json", &req.Conditions)...)
	diags.Append(assignJSONAttribute(plan.AppearanceJSON, "appearance_json", &req.Appearance)...)
	diags.Append(assignJSONAttribute(plan.ResponseSamplingDailyLimits, "response_sampling_daily_limits_json", &req.ResponseSamplingDailyLimits)...)
	diags.Append(assignJSONAttribute(plan.TranslationsJSON, "translations_json", &req.Translations)...)
	diags.Append(assignJSONAttribute(plan.FormContentJSON, "form_content_json", &req.FormContent)...)
	diags.Append(assignSurveyTargetingFilters(plan, state, req)...)

	return diags
}

func assignJSONAttribute(attr types.String, name string, target *interface{}) diag.Diagnostics {
	value, diags := parseJSONAttribute(attr, name)
	if diags.HasError() {
		return diags
	}

	*target = value
	return nil
}

func assignSurveyTargetingFilters(plan, state SurveyTFModel, req *httpclient.SurveyRequest) diag.Diagnostics {
	if !plan.TargetingFlagFiltersJSON.IsNull() && !plan.TargetingFlagFiltersJSON.IsUnknown() {
		value, diags := parseJSONAttribute(plan.TargetingFlagFiltersJSON, "targeting_flag_filters_json")
		if diags.HasError() {
			return diags
		}
		req.TargetingFlagFilters = value
		return nil
	}

	if !state.TargetingFlagFiltersJSON.IsNull() && !state.TargetingFlagFiltersJSON.IsUnknown() {
		// Preserve the write-only value across normal updates when PostHog does not
		// expose the original filters back to the provider.
		value, diags := parseJSONAttribute(state.TargetingFlagFiltersJSON, "targeting_flag_filters_json")
		if diags.HasError() {
			return diags
		}
		req.TargetingFlagFilters = value
	}

	return nil
}

func applySurveyTimingFields(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.StartDate = stringPtrFromValue(plan.StartDate)
	req.EndDate = stringPtrFromValue(plan.EndDate)
	req.Archived = boolPtrFromValue(plan.Archived)
	req.ResponsesLimit = int64PtrFromValue(plan.ResponsesLimit)
	req.IterationCount = int64PtrFromValue(plan.IterationCount)
	req.IterationFrequencyDays = int64PtrFromValue(plan.IterationFrequencyDays)
}

func applySurveySamplingFields(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.ResponseSamplingStartDate = stringPtrFromValue(plan.ResponseSamplingStartDate)
	req.ResponseSamplingIntervalType = stringPtrFromValue(plan.ResponseSamplingIntervalType)
	req.ResponseSamplingInterval = int64PtrFromValue(plan.ResponseSamplingInterval)
	req.ResponseSamplingLimit = int64PtrFromValue(plan.ResponseSamplingLimit)
}

func applySurveyBooleanFields(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.EnablePartialResponses = boolPtrFromValue(plan.EnablePartialResponses)
	req.EnableIframeEmbedding = boolPtrFromValue(plan.EnableIframeEmbedding)
	req.CreateInFolder = stringPtrFromValue(plan.CreateInFolder)
}

func applySurveyTargetingClear(plan, state SurveyTFModel, req *httpclient.SurveyRequest) {
	if !plan.TargetingFlagID.IsNull() || state.TargetingFlagID.IsNull() || state.TargetingFlagID.IsUnknown() {
		return
	}

	// Explicitly clear an existing external targeting flag when the config removes it.
	removeTargetingFlag := true
	req.RemoveTargetingFlag = &removeTargetingFlag
	req.TargetingFlagID = nil
	req.TargetingFlagFilters = nil
}

func (o SurveyOps) MapResponseToModel(_ context.Context, resp httpclient.Survey, model *SurveyTFModel) diag.Diagnostics {
	model.ID = types.StringValue(resp.ID)
	mapSurveyCoreFields(resp, model)
	mapSurveyRelationshipFields(resp, model)
	mapSurveyQuestionsField(resp, model)
	mapSurveyStructuredJSONFields(resp, model)
	preserveWriteOnlySurveyFields(model)
	mapSurveyScheduleFields(resp, model)
	mapSurveyComputedFields(resp, model)

	return nil
}

func mapSurveyCoreFields(resp httpclient.Survey, model *SurveyTFModel) {
	model.Name = core.PtrToStringTrimmed(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.Type = core.PtrToStringTrimmed(resp.Type)
	model.Schedule = core.PtrToStringNullIfEmptyTrimmed(resp.Schedule)
}

func mapSurveyRelationshipFields(resp httpclient.Survey, model *SurveyTFModel) {
	model.LinkedFlagID = int64ValueFromMapOrNull(resp.LinkedFlag)
	model.TargetingFlagID = int64ValueFromMapOrNull(resp.TargetingFlag)

	if model.LinkedInsightID.IsUnknown() {
		model.LinkedInsightID = types.Int64Null()
	}
}

func mapSurveyQuestionsField(resp httpclient.Survey, model *SurveyTFModel) {
	if len(resp.Questions) == 0 {
		return
	}

	normalized, err := normalizeSurveyQuestionsForState(resp.Questions, valueStringOrEmpty(model.QuestionsJSON))
	if err == nil {
		model.QuestionsJSON = types.StringValue(normalized)
	}
}

func mapSurveyStructuredJSONFields(resp httpclient.Survey, model *SurveyTFModel) {
	applyNormalizedJSONString(resp.Conditions, valueStringOrEmpty(model.ConditionsJSON), &model.ConditionsJSON)
	applyNormalizedJSONString(resp.Appearance, valueStringOrEmpty(model.AppearanceJSON), &model.AppearanceJSON)
	applyNormalizedJSONString(resp.ResponseSamplingDailyLimits, valueStringOrEmpty(model.ResponseSamplingDailyLimits), &model.ResponseSamplingDailyLimits)
	applyNormalizedJSONString(resp.Translations, valueStringOrEmpty(model.TranslationsJSON), &model.TranslationsJSON)
	applyNormalizedJSONString(resp.FormContent, valueStringOrEmpty(model.FormContentJSON), &model.FormContentJSON)
}

func preserveWriteOnlySurveyFields(model *SurveyTFModel) {
	// targeting_flag_filters is write-only. Preserve the configured value during
	// ordinary refreshes instead of trying to reconstruct it from augmented API
	// state. Imported resources keep this attribute null.
	if model.TargetingFlagFiltersJSON.IsUnknown() {
		model.TargetingFlagFiltersJSON = types.StringNull()
	}
}

func mapSurveyScheduleFields(resp httpclient.Survey, model *SurveyTFModel) {
	model.StartDate = core.PtrToStringNullIfEmptyTrimmed(resp.StartDate)
	model.EndDate = core.PtrToStringNullIfEmptyTrimmed(resp.EndDate)
	model.Archived = core.PtrToBool(resp.Archived)
	model.ResponsesLimit = int64ValueOrNull(resp.ResponsesLimit)
	model.IterationCount = int64ValueOrNull(resp.IterationCount)
	model.IterationFrequencyDays = int64ValueOrNull(resp.IterationFrequencyDays)
	model.ResponseSamplingStartDate = core.PtrToStringNullIfEmptyTrimmed(resp.ResponseSamplingStartDate)
	model.ResponseSamplingIntervalType = core.PtrToStringNullIfEmptyTrimmed(resp.ResponseSamplingIntervalType)
	model.ResponseSamplingInterval = int64ValueOrNull(resp.ResponseSamplingInterval)
	model.ResponseSamplingLimit = int64ValueOrNull(resp.ResponseSamplingLimit)
	model.EnablePartialResponses = core.PtrToBool(resp.EnablePartialResponses)
	model.EnableIframeEmbedding = core.PtrToBool(resp.EnableIframeEmbedding)
}

func mapSurveyComputedFields(resp httpclient.Survey, model *SurveyTFModel) {
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.CreatedByJSON = jsonStringValue(resp.CreatedBy)
	model.LinkedFlagJSON = jsonStringValue(resp.LinkedFlag)
	model.TargetingFlagJSON = jsonStringValue(resp.TargetingFlag)
	model.InternalTargetingFlagJSON = jsonStringValue(resp.InternalTargetingFlag)
}

func applyNormalizedJSONString(apiData interface{}, current string, target *types.String) {
	if normalized, ok := normalizedJSONString(apiData, current); ok {
		*target = normalized
	}
}

func (o SurveyOps) Create(ctx context.Context, client httpclient.PosthogClient, model SurveyTFModel, req httpclient.SurveyRequest) (httpclient.Survey, error) {
	return client.CreateSurvey(ctx, model.GetEffectiveProjectID(), req)
}

func (o SurveyOps) Read(ctx context.Context, client httpclient.PosthogClient, model SurveyTFModel) (httpclient.Survey, httpclient.HTTPStatusCode, error) {
	return client.GetSurvey(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o SurveyOps) Update(ctx context.Context, client httpclient.PosthogClient, model SurveyTFModel, req httpclient.SurveyRequest) (httpclient.Survey, httpclient.HTTPStatusCode, error) {
	return client.UpdateSurvey(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o SurveyOps) Delete(ctx context.Context, client httpclient.PosthogClient, model SurveyTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteSurvey(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func parseJSONAttribute(attr types.String, name string) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	if attr.IsNull() || attr.IsUnknown() {
		return nil, diags
	}

	var parsed interface{}
	if err := json.Unmarshal([]byte(strings.TrimSpace(attr.ValueString())), &parsed); err != nil {
		diags.AddError("Invalid "+name, fmt.Sprintf("%s must contain valid JSON: %s", name, err.Error()))
	}
	return parsed, diags
}

func parseJSONArrayAttribute(attr types.String, name string) ([]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	if attr.IsNull() || attr.IsUnknown() {
		diags.AddError("Missing "+name, fmt.Sprintf("%s is required", name))
		return nil, diags
	}

	var parsed []interface{}
	if err := json.Unmarshal([]byte(strings.TrimSpace(attr.ValueString())), &parsed); err != nil {
		diags.AddError("Invalid "+name, fmt.Sprintf("%s must contain a valid JSON array: %s", name, err.Error()))
	}
	return parsed, diags
}

func normalizeSurveyQuestionsForState(apiData interface{}, userJSON string) (string, error) {
	return normalizeJSONForState(stripSurveyQuestionServerFields(apiData), userJSON)
}

func stripSurveyQuestionServerFields(v interface{}) interface{} {
	switch value := v.(type) {
	case map[string]interface{}:
		cleaned := make(map[string]interface{}, len(value))
		for k, inner := range value {
			if k == "id" {
				continue
			}
			cleaned[k] = stripSurveyQuestionServerFields(inner)
		}
		return cleaned
	case []interface{}:
		cleaned := make([]interface{}, len(value))
		for i, item := range value {
			cleaned[i] = stripSurveyQuestionServerFields(item)
		}
		return cleaned
	default:
		return v
	}
}

func normalizedJSONString(apiData interface{}, current string) (types.String, bool) {
	if apiData == nil {
		return types.StringNull(), true
	}

	normalized, err := normalizeJSONForState(apiData, current)
	if err != nil {
		return types.StringNull(), false
	}
	if normalized == "" {
		return types.StringNull(), true
	}
	return types.StringValue(normalized), true
}

func jsonStringValue(value interface{}) types.String {
	if value == nil {
		return types.StringNull()
	}

	data, err := json.Marshal(value)
	if err != nil {
		return types.StringNull()
	}
	return types.StringValue(string(data))
}

func int64PtrFromValue(value types.Int64) *int64 {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	v := value.ValueInt64()
	return &v
}

func stringPtrFromValue(value types.String) *string {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	v := value.ValueString()
	return &v
}

func boolPtrFromValue(value types.Bool) *bool {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	v := value.ValueBool()
	return &v
}

func int64ValueOrNull(value *int64) types.Int64 {
	if value == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*value)
}

func int64ValueFromMapOrNull(value map[string]interface{}) types.Int64 {
	if id, ok := int64FromMap(value); ok {
		return types.Int64Value(id)
	}
	return types.Int64Null()
}

func int64FromMap(value map[string]interface{}) (int64, bool) {
	if len(value) == 0 {
		return 0, false
	}

	raw, ok := value["id"]
	if !ok {
		return 0, false
	}
	switch typed := raw.(type) {
	case float64:
		return int64(typed), true
	case int64:
		return typed, true
	case int:
		return int64(typed), true
	default:
		return 0, false
	}
}

func valueStringOrEmpty(value types.String) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}
	return value.ValueString()
}
