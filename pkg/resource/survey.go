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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
)

// surveyTargetingFlagKeyPrefix starts the key of every feature flag PostHog creates from
// targeting_flag_filters.
const surveyTargetingFlagKeyPrefix = "survey-targeting-"

// surveyQuestionServerFields are fields the API generates per-question
// that should be stripped during import, when there is no configured JSON.
// Configured IDs must be preserved during normal reads and writes.
var surveyQuestionServerFields = map[string]struct{}{
	"id": {},
}

func NewSurvey() resource.Resource {
	return core.NewGenericResource[SurveyTFModel, httpclient.SurveyRequest, httpclient.Survey](
		SurveyOps{},
		core.ProjectScopedImportParser[SurveyTFModel](),
	)
}

type SurveyTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name                         types.String        `tfsdk:"name"`
	Description                  types.String        `tfsdk:"description"`
	Type                         types.String        `tfsdk:"type"`
	Schedule                     types.String        `tfsdk:"schedule"`
	LinkedFlagID                 types.Int64         `tfsdk:"linked_flag_id"`
	LinkedInsightID              types.Int64         `tfsdk:"linked_insight_id"`
	TargetingFlagID              types.Int64         `tfsdk:"targeting_flag_id"`
	TargetingFlagFiltersJSON     types.String        `tfsdk:"targeting_flag_filters_json"`
	QuestionsJSON                core.NormalizedJSON `tfsdk:"questions_json"`
	ConditionsJSON               core.NormalizedJSON `tfsdk:"conditions_json"`
	AppearanceJSON               core.NormalizedJSON `tfsdk:"appearance_json"`
	StartDate                    types.String        `tfsdk:"start_date"`
	EndDate                      types.String        `tfsdk:"end_date"`
	Archived                     types.Bool          `tfsdk:"archived"`
	ResponsesLimit               types.Int64         `tfsdk:"responses_limit"`
	IterationCount               types.Int64         `tfsdk:"iteration_count"`
	IterationFrequencyDays       types.Int64         `tfsdk:"iteration_frequency_days"`
	ResponseSamplingStartDate    types.String        `tfsdk:"response_sampling_start_date"`
	ResponseSamplingIntervalType types.String        `tfsdk:"response_sampling_interval_type"`
	ResponseSamplingInterval     types.Int64         `tfsdk:"response_sampling_interval"`
	ResponseSamplingLimit        types.Int64         `tfsdk:"response_sampling_limit"`
	ResponseSamplingDailyLimits  core.NormalizedJSON `tfsdk:"response_sampling_daily_limits_json"`
	EnablePartialResponses       types.Bool          `tfsdk:"enable_partial_responses"`
	EnableIframeEmbedding        types.Bool          `tfsdk:"enable_iframe_embedding"`
	TranslationsJSON             core.NormalizedJSON `tfsdk:"translations_json"`
	CreateInFolder               types.String        `tfsdk:"create_in_folder"`
	FormContentJSON              core.NormalizedJSON `tfsdk:"form_content_json"`
	CreatedAt                    types.String        `tfsdk:"created_at"`
	CreatedByJSON                types.String        `tfsdk:"created_by_json"`
	LinkedFlagJSON               types.String        `tfsdk:"linked_flag_json"`
	TargetingFlagJSON            types.String        `tfsdk:"targeting_flag_json"`
	InternalTargetingFlagJSON    types.String        `tfsdk:"internal_targeting_flag_json"`
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
			// linked_flag_id, linked_insight_id, and targeting_flag_id are intentionally
			// Optional-only (no Computed + UseStateForUnknown). The server never populates
			// a default for these, and Computed + UseStateForUnknown would carry the prior
			// value into plan when config drops the attribute — making the unlink path
			// (linked_flag_id: null) and the remove_targeting_flag clear path unreachable.
			// targeting_flag_id stays null while targeting_flag_filters_json is set, although
			// PostHog then creates a targeting flag: see mapSurveyRelationshipFields.
			"linked_flag_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Feature flag ID linked to the survey. Remove this attribute to unlink the flag from the survey.",
			},
			"linked_insight_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Insight ID linked to the survey. This field is write-only in the PostHog API and is preserved from Terraform state when configured.",
			},
			"targeting_flag_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Existing targeting feature flag ID to use for this survey. Remove this attribute to detach the targeting flag.",
			},
			"targeting_flag_filters_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object describing targeting flag filters. PostHog creates a feature flag from them; its ID is in `targeting_flag_json`, and `targeting_flag_id` stays unset. Removing this attribute deletes that flag. This input is write-only in the PostHog API and is preserved from Terraform state when configured.",
			},
			"questions_json": schema.StringAttribute{
				CustomType:          core.NormalizedJSONType{},
				Required:            true,
				MarkdownDescription: "JSON array describing the survey questions.",
			},
			"conditions_json": schema.StringAttribute{
				CustomType:          core.NormalizedJSONType{},
				Optional:            true,
				MarkdownDescription: "JSON object describing survey display and targeting conditions.",
			},
			// No UseStateForUnknown: an unset appearance must be unknown on updates so it is not
			// sent. State may hold only some of its fields, so sending it would drop the others.
			"appearance_json": schema.StringAttribute{
				CustomType:          core.NormalizedJSONType{},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "JSON object describing survey appearance settings. If unset, the survey keeps the appearance it has in PostHog, which adds a default thank-you message to new surveys. Removing this attribute keeps the current appearance; set it to `jsonencode({})` to clear it.",
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
				CustomType:          core.NormalizedJSONType{},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Daily response sampling limits, which PostHog calculates from the other `response_sampling_*` attributes as a JSON string. Leave this unset: PostHog ignores any other value, so the apply fails.",
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
				CustomType:          core.NormalizedJSONType{},
				Optional:            true,
				MarkdownDescription: "JSON object describing translated survey content.",
			},
			"create_in_folder": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Folder identifier used only during survey creation.",
			},
			"form_content_json": schema.StringAttribute{
				CustomType:          core.NormalizedJSONType{},
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
	diags.Append(applySurveyTargeting(plan, state, &req)...)

	return req, diags
}

func applySurveyDescription(plan, state SurveyTFModel, req *httpclient.SurveyRequest) {
	if core.IsKnown(plan.Description) {
		description := plan.Description.ValueString()
		req.Description = &description
		return
	}

	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = util.StringPtr("")
	}
}

func applySurveyRelationships(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.LinkedFlagID = util.Int64PtrFromValue(plan.LinkedFlagID)
	req.LinkedInsightID = util.Int64PtrFromValue(plan.LinkedInsightID)
	req.TargetingFlagID = util.Int64PtrFromValue(plan.TargetingFlagID)
	req.Schedule = util.StringPtrFromValue(plan.Schedule)
}

func parseSurveyQuestions(attr core.NormalizedJSON) ([]interface{}, diag.Diagnostics) {
	if attr.IsNull() || attr.IsUnknown() {
		return nil, nil
	}

	return parseJSONArrayAttribute(attr, "questions_json")
}

func appendSurveyJSONFields(plan, state SurveyTFModel, req *httpclient.SurveyRequest) diag.Diagnostics {
	var diags diag.Diagnostics

	diags.Append(assignClearableJSONAttribute(plan.ConditionsJSON, state.ConditionsJSON, "conditions_json", &req.Conditions)...)
	diags.Append(assignJSONAttribute(plan.AppearanceJSON, "appearance_json", &req.Appearance)...)
	diags.Append(assignJSONAttribute(plan.ResponseSamplingDailyLimits, "response_sampling_daily_limits_json", &req.ResponseSamplingDailyLimits)...)
	diags.Append(assignClearableJSONAttribute(plan.TranslationsJSON, state.TranslationsJSON, "translations_json", &req.Translations)...)
	diags.Append(assignClearableJSONAttribute(plan.FormContentJSON, state.FormContentJSON, "form_content_json", &req.FormContent)...)

	return diags
}

func assignJSONAttribute(attr core.NormalizedJSON, name string, target *interface{}) diag.Diagnostics {
	value, diags := parseJSONAttribute(attr.StringValue, name)
	if diags.HasError() {
		return diags
	}

	*target = value
	return nil
}

// assignClearableJSONAttribute sends null when the attribute was removed from the config.
// Leaving it out of the request would keep the old value in PostHog.
func assignClearableJSONAttribute(plan, state core.NormalizedJSON, name string, target *interface{}) diag.Diagnostics {
	if plan.IsNull() && core.IsKnown(state) {
		*target = util.JSONNull()
		return nil
	}
	return assignJSONAttribute(plan, name, target)
}

func applySurveyTimingFields(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.StartDate = util.StringPtrFromValue(plan.StartDate)
	req.EndDate = util.StringPtrFromValue(plan.EndDate)
	req.Archived = util.BoolPtrFromValue(plan.Archived)
	req.ResponsesLimit = util.Int64PtrFromValue(plan.ResponsesLimit)
	req.IterationCount = util.Int64PtrFromValue(plan.IterationCount)
	req.IterationFrequencyDays = util.Int64PtrFromValue(plan.IterationFrequencyDays)
}

func applySurveySamplingFields(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.ResponseSamplingStartDate = util.StringPtrFromValue(plan.ResponseSamplingStartDate)
	req.ResponseSamplingIntervalType = util.StringPtrFromValue(plan.ResponseSamplingIntervalType)
	req.ResponseSamplingInterval = util.Int64PtrFromValue(plan.ResponseSamplingInterval)
	req.ResponseSamplingLimit = util.Int64PtrFromValue(plan.ResponseSamplingLimit)
}

func applySurveyBooleanFields(plan SurveyTFModel, req *httpclient.SurveyRequest) {
	req.EnablePartialResponses = util.BoolPtrFromValue(plan.EnablePartialResponses)
	req.EnableIframeEmbedding = util.BoolPtrFromValue(plan.EnableIframeEmbedding)
	req.CreateInFolder = util.StringPtrFromValue(plan.CreateInFolder)
}

// applySurveyTargeting sends the configured targeting filters, or removes the targeting
// flag when the config no longer sets targeting_flag_id or targeting_flag_filters_json.
func applySurveyTargeting(plan, state SurveyTFModel, req *httpclient.SurveyRequest) diag.Diagnostics {
	var diags diag.Diagnostics

	if core.IsKnown(plan.TargetingFlagFiltersJSON) {
		// PostHog applies filters to the flag the survey already targets, so dropping
		// targeting_flag_id in favor of filters would rewrite a flag managed elsewhere.
		if plan.TargetingFlagID.IsNull() && core.IsKnown(state.TargetingFlagID) && !isSurveyOwnedFlag(state.TargetingFlagJSON) {
			diags.AddError("Cannot replace targeting_flag_id with targeting_flag_filters_json",
				fmt.Sprintf("PostHog would overwrite the filters of feature flag %d, which this survey targets but does not own. "+
					"Keep targeting_flag_id and change the filters on that feature flag instead.", state.TargetingFlagID.ValueInt64()))
			return diags
		}
		req.TargetingFlagFilters, diags = parseJSONAttribute(plan.TargetingFlagFiltersJSON, "targeting_flag_filters_json")
		return diags
	}

	if plan.TargetingFlagID.IsNull() && (core.IsKnown(state.TargetingFlagID) || core.IsKnown(state.TargetingFlagFiltersJSON)) {
		removeTargetingFlag := true
		req.RemoveTargetingFlag = &removeTargetingFlag
	}
	return diags
}

// isSurveyOwnedFlag reports whether PostHog created the targeting flag for a survey.
func isSurveyOwnedFlag(targetingFlagJSON types.String) bool {
	var flag struct {
		Key string `json:"key"`
	}
	if json.Unmarshal([]byte(targetingFlagJSON.ValueString()), &flag) != nil {
		return false
	}
	return strings.HasPrefix(flag.Key, surveyTargetingFlagKeyPrefix)
}

func (o SurveyOps) MapResponseToModel(_ context.Context, resp httpclient.Survey, model *SurveyTFModel) diag.Diagnostics {
	model.ID = types.StringValue(resp.ID)
	mapSurveyCoreFields(resp, model)
	mapSurveyRelationshipFields(resp, model)
	mapSurveyQuestionsField(resp, model)
	mapSurveyStructuredJSONFields(resp, model)
	mapSurveyTargetingFilters(resp, model)
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
	// Filters make PostHog create the targeting flag, so the config has no ID for it and
	// state must not either. The flag is still in targeting_flag_json.
	if !core.IsKnown(model.TargetingFlagFiltersJSON) || !model.TargetingFlagID.IsNull() {
		model.TargetingFlagID = int64ValueFromMapOrNull(resp.TargetingFlag)
	}

	if model.LinkedInsightID.IsUnknown() {
		model.LinkedInsightID = types.Int64Null()
	}
}

func mapSurveyQuestionsField(resp httpclient.Survey, model *SurveyTFModel) {
	if len(resp.Questions) == 0 {
		return
	}

	normalized, err := normalizeSurveyQuestionsForState(resp.Questions, model.QuestionsJSON.ValueString())
	if err == nil {
		model.QuestionsJSON = core.NewNormalizedJSONValue(normalized)
	}
}

func mapSurveyStructuredJSONFields(resp httpclient.Survey, model *SurveyTFModel) {
	applyNormalizedJSONString(resp.Conditions, model.ConditionsJSON.ValueString(), &model.ConditionsJSON)
	applyNormalizedJSONString(resp.Appearance, model.AppearanceJSON.ValueString(), &model.AppearanceJSON)
	applyNormalizedJSONString(resp.ResponseSamplingDailyLimits, model.ResponseSamplingDailyLimits.ValueString(), &model.ResponseSamplingDailyLimits)
	applyNormalizedJSONString(resp.Translations, model.TranslationsJSON.ValueString(), &model.TranslationsJSON)
	applyNormalizedJSONString(resp.FormContent, model.FormContentJSON.ValueString(), &model.FormContentJSON)
}

func mapSurveyTargetingFilters(resp httpclient.Survey, model *SurveyTFModel) {
	// targeting_flag_filters is write-only. Preserve the configured value during
	// ordinary refreshes instead of trying to reconstruct it from augmented API
	// state. Imported resources keep this attribute null.
	if model.TargetingFlagFiltersJSON.IsUnknown() {
		model.TargetingFlagFiltersJSON = types.StringNull()
		return
	}

	// Filters always leave a targeting flag in PostHog, except empty ones, which PostHog
	// ignores. Without a flag the survey targets everyone, so let the plan restore them.
	filters := json.RawMessage(model.TargetingFlagFiltersJSON.ValueString())
	if len(resp.TargetingFlag) == 0 && !isEmptyJSONObjectOrNull(filters) {
		model.TargetingFlagFiltersJSON = types.StringNull()
	}
}

func mapSurveyScheduleFields(resp httpclient.Survey, model *SurveyTFModel) {
	model.StartDate = core.PtrToStringNullIfEmptyTrimmed(resp.StartDate)
	model.EndDate = core.PtrToStringNullIfEmptyTrimmed(resp.EndDate)
	model.Archived = core.PtrToBool(resp.Archived)
	model.ResponsesLimit = util.PtrToInt64(resp.ResponsesLimit)
	model.IterationCount = util.PtrToInt64(resp.IterationCount)
	model.IterationFrequencyDays = util.PtrToInt64(resp.IterationFrequencyDays)
	model.ResponseSamplingStartDate = core.PtrToStringNullIfEmptyTrimmed(resp.ResponseSamplingStartDate)
	model.ResponseSamplingIntervalType = core.PtrToStringNullIfEmptyTrimmed(resp.ResponseSamplingIntervalType)
	model.ResponseSamplingInterval = util.PtrToInt64(resp.ResponseSamplingInterval)
	model.ResponseSamplingLimit = util.PtrToInt64(resp.ResponseSamplingLimit)
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

func applyNormalizedJSONString(apiData interface{}, current string, target *core.NormalizedJSON) {
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

func parseJSONArrayAttribute(attr core.NormalizedJSON, name string) ([]interface{}, diag.Diagnostics) {
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
	if strings.TrimSpace(userJSON) == "" {
		apiData = util.StripFields(apiData, surveyQuestionServerFields)
	}
	// The configured field shape keeps explicit IDs and omits generated IDs
	// independently for each question, while still surfacing remote ID changes.
	return normalizeJSONForState(apiData, userJSON)
}

func normalizedJSONString(apiData interface{}, current string) (core.NormalizedJSON, bool) {
	if apiData == nil {
		return core.NewNormalizedJSONNull(), true
	}

	normalized, err := normalizeJSONForState(apiData, current)
	if err != nil {
		return core.NewNormalizedJSONNull(), false
	}
	if normalized == "" {
		return core.NewNormalizedJSONNull(), true
	}
	return core.NewNormalizedJSONValue(normalized), true
}

func jsonStringValue(value interface{}) types.String {
	if value == nil {
		return types.StringNull()
	}

	data, err := json.Marshal(value)
	if err != nil {
		return types.StringNull()
	}
	// A typed-nil map/slice marshals to the JSON literal "null"; surface that as
	// a Terraform null instead of the four-character string "null".
	if string(data) == "null" {
		return types.StringNull()
	}
	return types.StringValue(string(data))
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
