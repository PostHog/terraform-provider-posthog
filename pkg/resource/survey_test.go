package resource

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSurveyName               = "survey"
	testSurveyDescription        = "desc"
	testSurveyType               = "popover"
	testSurveyQuestionJSON       = `[{"type":"open","question":"How are you?"}]`
	testSurveyQuestionNormalized = `[{"question":"How are you?","type":"open"}]`
	testSurveyStartDate          = "2026-04-20T00:00:00Z"
	testSurveyEndDate            = "2026-04-21T00:00:00Z"
	testSurveySamplingStartDate  = "2026-04-22T00:00:00Z"
	testSurveyProjectID          = "project-123"
	testSurveyID                 = "survey-123"
	testSurveyCollectionPath     = "/api/projects/project-123/surveys/"
	testSurveyResourcePath       = "/api/projects/project-123/surveys/survey-123/"
)

func TestSurveyResourceMetadataAndSchema(t *testing.T) {
	resourceInstance := NewSurvey()
	require.NotNil(t, resourceInstance)

	ops := SurveyOps{}
	assert.Equal(t, "Survey", ops.ResourceName())

	resourceSchema := ops.Schema()
	assert.Equal(t, "Manage PostHog surveys via the project-scoped surveys API.", resourceSchema.MarkdownDescription)

	nameAttr, ok := resourceSchema.Attributes["name"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, nameAttr.Required)

	scheduleAttr, ok := resourceSchema.Attributes["schedule"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, scheduleAttr.Optional)
	assert.True(t, scheduleAttr.Computed)

	linkedInsightAttr, ok := resourceSchema.Attributes["linked_insight_id"].(schema.Int64Attribute)
	require.True(t, ok)
	assert.Contains(t, linkedInsightAttr.MarkdownDescription, "write-only")
}

func TestSurveyJSON_ComparedByContent(t *testing.T) {
	attributes := SurveyOps{}.Schema().Attributes
	for _, name := range []string{
		"questions_json",
		"conditions_json",
		"appearance_json",
		"response_sampling_daily_limits_json",
		"translations_json",
		"form_content_json",
	} {
		assert.Equal(t, core.NormalizedJSONType{}, attributes[name].GetType(), name)
	}
}

// conditions_json, appearance_json, response_sampling_daily_limits_json,
// translations_json and form_content_json are all mapped the same way, so
// conditions_json stands in for the five.
func TestSurveyMapResponseToModel_OptionalJSON(t *testing.T) {
	configured := core.NewNormalizedJSONValue("{\n  \"url\": \"https://example.com\"\n}\n")
	tests := map[string]struct {
		before core.NormalizedJSON
		api    any
		after  core.NormalizedJSON
	}{
		"unset":                      {core.NewNormalizedJSONNull(), nil, core.NewNormalizedJSONNull()},
		"unset, value in PostHog":    {core.NewNormalizedJSONNull(), map[string]any{"url": "a"}, core.NewNormalizedJSONValue(`{"url":"a"}`)},
		"gone in PostHog":            {configured, nil, core.NewNormalizedJSONNull()},
		"changed in PostHog":         {configured, map[string]any{"url": "a"}, core.NewNormalizedJSONValue(`{"url":"a"}`)},
		"field added by PostHog":     {configured, map[string]any{"url": "https://example.com", "actions": nil}, core.NewNormalizedJSONValue(`{"url":"https://example.com"}`)},
		"field removed in PostHog":   {configured, map[string]any{}, core.NewNormalizedJSONValue(`{}`)},
		"JSON string from PostHog":   {core.NewNormalizedJSONValue("\"[1]\"\n"), "[1]", core.NewNormalizedJSONValue(`"[1]"`)},
		"JSON string on import":      {core.NewNormalizedJSONNull(), "[1]", core.NewNormalizedJSONValue(`"[1]"`)},
		"different type from config": {configured, "[1]", core.NewNormalizedJSONValue(`"[1]"`)},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := SurveyTFModel{ConditionsJSON: tt.before}
			diags := SurveyOps{}.MapResponseToModel(context.Background(), httpclient.Survey{Conditions: tt.api}, &model)
			require.False(t, diags.HasError(), diags.Errors())
			assert.Equal(t, tt.after, model.ConditionsJSON)
		})
	}
}

// State keeps the configured text only when it equals what PostHog returns, so every
// attribute must map back to the same content however its JSON is formatted.
func TestSurveyJSON_FormattedConfigRoundTrips(t *testing.T) {
	formatted := func(compact string) core.NormalizedJSON {
		var out bytes.Buffer
		require.NoError(t, json.Indent(&out, []byte(compact), "", "  "))
		return core.NewNormalizedJSONValue(out.String() + "\n")
	}
	plan := SurveyTFModel{
		Name:                        types.StringValue(testSurveyName),
		Type:                        types.StringValue(testSurveyType),
		QuestionsJSON:               formatted(`[{"type":"open","question":"Q&A","id":"set-in-config"}]`),
		ConditionsJSON:              formatted(`{"url":"https://example.com/?a=1&b=2","selector":"div > p"}`),
		AppearanceJSON:              formatted(`{"thankYouMessageHeader":"<b>Thanks</b>","backgroundColor":"#eeeded"}`),
		ResponseSamplingDailyLimits: formatted(`"[{\"daily_response_limit\": 10}]"`),
		TranslationsJSON:            formatted(`{"fr":{"thankYouMessageHeader":"\u00e0 bient\u00f4t","name":"Bonjour"}}`),
		FormContentJSON:             formatted(`{"type":"doc","content":[{"type":"paragraph","attrs":null}]}`),
	}

	req, diags := SurveyOps{}.BuildCreateRequest(context.Background(), plan)
	require.False(t, diags.HasError(), diags.Errors())
	body, err := json.Marshal(req)
	require.NoError(t, err)

	state := plan
	diags = SurveyOps{}.MapResponseToModel(context.Background(), decodeAPIResponse[httpclient.Survey](t, string(body)), &state)
	require.False(t, diags.HasError(), diags.Errors())

	for name, values := range map[string][2]core.NormalizedJSON{
		"questions_json":                      {plan.QuestionsJSON, state.QuestionsJSON},
		"conditions_json":                     {plan.ConditionsJSON, state.ConditionsJSON},
		"appearance_json":                     {plan.AppearanceJSON, state.AppearanceJSON},
		"response_sampling_daily_limits_json": {plan.ResponseSamplingDailyLimits, state.ResponseSamplingDailyLimits},
		"translations_json":                   {plan.TranslationsJSON, state.TranslationsJSON},
		"form_content_json":                   {plan.FormContentJSON, state.FormContentJSON},
	} {
		equal, diags := values[0].StringSemanticEquals(context.Background(), values[1])
		require.False(t, diags.HasError(), diags.Errors())
		assert.True(t, equal, "%s: config %s must equal state %s", name, values[0].ValueString(), values[1].ValueString())
	}
}

// PostHog assigns an id to every question sent without one, so API responses
// always carry ids. State must hold exactly the keys the config wrote for each
// question: an id the config sets is kept (and shows drift if PostHog changes
// it), an id PostHog assigned is left out.
func TestNormalizeSurveyQuestionsForState(t *testing.T) {
	question := func(id, text string) map[string]any {
		q := map[string]any{"type": "open", "question": text}
		if id != "" {
			q["id"] = id
		}
		return q
	}

	tests := []struct {
		name   string
		api    []any
		config string // questions_json as written; "" means import (no config yet)
		want   string
		drift  bool // state differs from config, so the next plan shows an update
	}{
		{
			name:   "import leaves out ids PostHog assigned",
			api:    []any{question("assigned-by-posthog", "How are you?")},
			config: "",
			want:   `[{"type":"open","question":"How are you?"}]`,
		},
		{
			name:   "config without ids leaves out ids PostHog assigned",
			api:    []any{question("assigned-by-posthog", "How are you?")},
			config: `[{"type":"open","question":"How are you?"}]`,
			want:   `[{"type":"open","question":"How are you?"}]`,
		},
		{
			name:   "id set in config is kept, id PostHog assigned to the next question is left out",
			api:    []any{question("set-in-config", "First"), question("assigned-by-posthog", "Second")},
			config: `[{"id":"set-in-config","type":"open","question":"First"},{"type":"open","question":"Second"}]`,
			want:   `[{"id":"set-in-config","type":"open","question":"First"},{"type":"open","question":"Second"}]`,
		},
		{
			name:   "id changed in PostHog shows as drift",
			api:    []any{question("changed-in-ui", "How are you?")},
			config: `[{"id":"set-in-config","type":"open","question":"How are you?"}]`,
			want:   `[{"id":"changed-in-ui","type":"open","question":"How are you?"}]`,
			drift:  true,
		},
		{
			name:   "question added in PostHog shows as drift, including its id",
			api:    []any{question("set-in-config", "First"), question("assigned-by-posthog", "Added in UI")},
			config: `[{"id":"set-in-config","type":"open","question":"First"}]`,
			want:   `[{"id":"set-in-config","type":"open","question":"First"},{"id":"assigned-by-posthog","type":"open","question":"Added in UI"}]`,
			drift:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeSurveyQuestionsForState(tt.api, tt.config)
			require.NoError(t, err)
			assert.JSONEq(t, tt.want, got)
			if tt.config == "" {
				return // import: nothing to compare against
			}
			var gotV, configV any
			require.NoError(t, json.Unmarshal([]byte(got), &gotV))
			require.NoError(t, json.Unmarshal([]byte(tt.config), &configV))
			if tt.drift {
				assert.NotEqual(t, configV, gotV, "state should differ from config")
			} else {
				assert.Equal(t, configV, gotV, "state should match config")
			}
		})
	}
}

func TestSurveyBuildRequest_Targeting(t *testing.T) {
	const filtersJSON = `{"groups":[{"properties":[],"rollout_percentage":50}]}`
	filters := map[string]any{"groups": []any{map[string]any{"properties": []any{}, "rollout_percentage": float64(50)}}}
	surveyFlag := types.StringValue(`{"id":11,"key":"survey-targeting-afc3bbeb1f"}`)
	userFlag := types.StringValue(`{"id":11,"key":"beta-users"}`)
	none, id, configured := types.Int64Null(), types.Int64Value(11), types.StringValue(filtersJSON)

	tests := map[string]struct {
		planID       types.Int64
		planFilters  types.String
		stateID      types.Int64
		stateFilters types.String
		stateFlag    types.String
		wantFilters  any
		wantRemove   bool
		wantError    string
	}{
		"never set":                       {none, types.StringNull(), none, types.StringNull(), types.StringNull(), nil, false, ""},
		"filters added":                   {none, configured, none, types.StringNull(), types.StringNull(), filters, false, ""},
		"filters kept":                    {none, configured, none, configured, surveyFlag, filters, false, ""},
		"filters removed":                 {none, types.StringNull(), none, configured, surveyFlag, nil, true, ""},
		"filters added after import":      {none, configured, id, types.StringNull(), surveyFlag, filters, false, ""},
		"state written by an old version": {none, configured, id, configured, surveyFlag, filters, false, ""},
		"flag removed":                    {none, types.StringNull(), id, types.StringNull(), userFlag, nil, true, ""},
		"flag and filters":                {id, configured, id, configured, userFlag, filters, false, ""},
		"filters removed, flag kept":      {id, types.StringNull(), id, configured, userFlag, nil, false, ""},
		"flag replaced with filters":      {none, configured, id, types.StringNull(), userFlag, nil, false, "feature flag 11"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			questions := core.NewNormalizedJSONValue(testSurveyQuestionJSON)
			plan := SurveyTFModel{QuestionsJSON: questions, TargetingFlagID: tt.planID, TargetingFlagFiltersJSON: tt.planFilters}
			state := SurveyTFModel{QuestionsJSON: questions, TargetingFlagID: tt.stateID, TargetingFlagFiltersJSON: tt.stateFilters, TargetingFlagJSON: tt.stateFlag}

			req, diags := SurveyOps{}.BuildUpdateRequest(context.Background(), plan, state)

			if tt.wantError != "" {
				require.True(t, diags.HasError())
				assert.Contains(t, diags.Errors()[0].Detail(), tt.wantError)
			} else {
				require.False(t, diags.HasError(), diags.Errors())
			}
			assert.Equal(t, tt.wantFilters, req.TargetingFlagFilters)
			assert.Equal(t, tt.wantRemove, req.RemoveTargetingFlag != nil)
		})
	}
}

func TestSurveyMapResponseToModel_Targeting(t *testing.T) {
	flag := map[string]any{"id": float64(11), "key": "survey-targeting-afc3bbeb1f"}
	none, id := types.Int64Null(), types.Int64Value(11)
	noFilters, configured, empty := types.StringNull(), types.StringValue(`{"groups":[]}`), types.StringValue(`{ }`)

	tests := map[string]struct {
		beforeID      types.Int64
		beforeFilters types.String
		flag          map[string]any
		afterID       types.Int64
		afterFilters  types.String
	}{
		"import or flag added in PostHog":   {none, noFilters, flag, id, noFilters},
		"configured flag":                   {id, noFilters, flag, id, noFilters},
		"configured flag gone in PostHog":   {id, noFilters, nil, none, noFilters},
		"filters":                           {none, configured, flag, none, configured},
		"filters with a configured flag":    {id, configured, flag, id, configured},
		"filters, flag gone in PostHog":     {none, configured, nil, none, noFilters},
		"empty filters, which make no flag": {none, empty, nil, none, empty},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := SurveyTFModel{TargetingFlagID: tt.beforeID, TargetingFlagFiltersJSON: tt.beforeFilters}
			diags := SurveyOps{}.MapResponseToModel(context.Background(), httpclient.Survey{TargetingFlag: tt.flag}, &model)
			require.False(t, diags.HasError(), diags.Errors())
			assert.Equal(t, tt.afterID, model.TargetingFlagID)
			assert.Equal(t, tt.afterFilters, model.TargetingFlagFiltersJSON)
		})
	}
}

// PostHog keeps the old value of a field the update leaves out, so removing an attribute
// from the config must send null. appearance_json and response_sampling_daily_limits_json
// are computed, so an update plans them as unknown while they are unset and sends nothing.
func TestSurveyBuildUpdateRequest_RemovedJSON(t *testing.T) {
	value := core.NewNormalizedJSONValue(`{"a":1}`)
	questions := core.NewNormalizedJSONValue(testSurveyQuestionJSON)
	requestBody := func(plan, state SurveyTFModel) map[string]any {
		plan.QuestionsJSON, state.QuestionsJSON = questions, questions
		req, diags := SurveyOps{}.BuildUpdateRequest(context.Background(), plan, state)
		require.False(t, diags.HasError(), diags.Errors())
		data, err := json.Marshal(req)
		require.NoError(t, err)
		var body map[string]any
		require.NoError(t, json.Unmarshal(data, &body))
		return body
	}

	removed := requestBody(
		SurveyTFModel{
			AppearanceJSON:              core.NewNormalizedJSONUnknown(),
			ResponseSamplingDailyLimits: core.NewNormalizedJSONUnknown(),
		},
		SurveyTFModel{
			ConditionsJSON:              value,
			TranslationsJSON:            value,
			FormContentJSON:             value,
			AppearanceJSON:              value,
			ResponseSamplingDailyLimits: value,
		},
	)
	for _, key := range []string{"conditions", "translations", "form_content"} {
		got, present := removed[key]
		assert.True(t, present, key)
		assert.Nil(t, got, key)
	}
	assert.NotContains(t, removed, "appearance")
	assert.NotContains(t, removed, "response_sampling_daily_limits")

	neverSet := requestBody(SurveyTFModel{}, SurveyTFModel{})
	for _, key := range []string{"conditions", "translations", "form_content", "appearance", "response_sampling_daily_limits"} {
		assert.NotContains(t, neverSet, key)
	}
}

func TestSurveySchema_ComputedByPostHog(t *testing.T) {
	attributes := SurveyOps{}.Schema().Attributes
	for _, name := range []string{"appearance_json", "response_sampling_daily_limits_json"} {
		assert.True(t, attributes[name].IsOptional(), name)
		assert.True(t, attributes[name].IsComputed(), name)
	}
}

func TestSurveyBuildRequestAllFields(t *testing.T) {
	ops := SurveyOps{}
	plan := SurveyTFModel{
		Name:                         types.StringValue(testSurveyName),
		Description:                  types.StringValue(testSurveyDescription),
		Type:                         types.StringValue(testSurveyType),
		Schedule:                     types.StringValue("recurring"),
		LinkedFlagID:                 types.Int64Value(11),
		LinkedInsightID:              types.Int64Value(12),
		TargetingFlagID:              types.Int64Value(13),
		TargetingFlagFiltersJSON:     types.StringValue(`{"groups":[{"rollout_percentage":100}]}`),
		QuestionsJSON:                core.NewNormalizedJSONValue(testSurveyQuestionJSON),
		ConditionsJSON:               core.NewNormalizedJSONValue(`{"url":"https://example.com"}`),
		AppearanceJSON:               core.NewNormalizedJSONValue(`{"border_color":"#000000"}`),
		StartDate:                    types.StringValue(testSurveyStartDate),
		EndDate:                      types.StringValue(testSurveyEndDate),
		Archived:                     types.BoolValue(true),
		ResponsesLimit:               types.Int64Value(10),
		IterationCount:               types.Int64Value(2),
		IterationFrequencyDays:       types.Int64Value(7),
		ResponseSamplingStartDate:    types.StringValue(testSurveySamplingStartDate),
		ResponseSamplingIntervalType: types.StringValue("week"),
		ResponseSamplingInterval:     types.Int64Value(2),
		ResponseSamplingLimit:        types.Int64Value(20),
		ResponseSamplingDailyLimits:  core.NewNormalizedJSONValue(`{"monday":5}`),
		EnablePartialResponses:       types.BoolValue(true),
		EnableIframeEmbedding:        types.BoolValue(true),
		TranslationsJSON:             core.NewNormalizedJSONValue(`{"fr":{"name":"Bonjour"}}`),
		CreateInFolder:               types.StringValue("folder-123"),
		FormContentJSON:              core.NewNormalizedJSONValue(`{"submitButtonText":"Send"}`),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), plan)
	require.False(t, diags.HasError())

	require.NotNil(t, req.Description)
	assert.Equal(t, testSurveyDescription, *req.Description)
	require.NotNil(t, req.Schedule)
	assert.Equal(t, "recurring", *req.Schedule)
	assert.Equal(t, int64(11), *req.LinkedFlagID)
	assert.Equal(t, int64(12), *req.LinkedInsightID)
	assert.Equal(t, int64(13), *req.TargetingFlagID)
	assert.Len(t, req.Questions, 1)
	assert.Equal(t, map[string]interface{}{"url": "https://example.com"}, req.Conditions)
	assert.Equal(t, map[string]interface{}{"border_color": "#000000"}, req.Appearance)
	assert.Equal(t, map[string]interface{}{"groups": []interface{}{map[string]interface{}{"rollout_percentage": float64(100)}}}, req.TargetingFlagFilters)
	assert.Equal(t, map[string]interface{}{"monday": float64(5)}, req.ResponseSamplingDailyLimits)
	assert.Equal(t, map[string]interface{}{"fr": map[string]interface{}{"name": "Bonjour"}}, req.Translations)
	assert.Equal(t, map[string]interface{}{"submitButtonText": "Send"}, req.FormContent)
	assert.Equal(t, testSurveyStartDate, *req.StartDate)
	assert.Equal(t, testSurveyEndDate, *req.EndDate)
	assert.True(t, *req.Archived)
	assert.Equal(t, int64(10), *req.ResponsesLimit)
	assert.Equal(t, int64(2), *req.IterationCount)
	assert.Equal(t, int64(7), *req.IterationFrequencyDays)
	assert.Equal(t, testSurveySamplingStartDate, *req.ResponseSamplingStartDate)
	assert.Equal(t, "week", *req.ResponseSamplingIntervalType)
	assert.Equal(t, int64(2), *req.ResponseSamplingInterval)
	assert.Equal(t, int64(20), *req.ResponseSamplingLimit)
	assert.True(t, *req.EnablePartialResponses)
	assert.True(t, *req.EnableIframeEmbedding)
	assert.Equal(t, "folder-123", *req.CreateInFolder)
}

func TestSurveyBuildRequestClearsDescriptionAndTargetingFlag(t *testing.T) {
	ops := SurveyOps{}
	plan := SurveyTFModel{
		Name:          types.StringValue(testSurveyName),
		Type:          types.StringValue(testSurveyType),
		QuestionsJSON: core.NewNormalizedJSONValue(testSurveyQuestionJSON),
	}
	state := SurveyTFModel{
		Description:              types.StringValue("previous"),
		TargetingFlagID:          types.Int64Value(44),
		TargetingFlagFiltersJSON: types.StringValue(`{"groups":[{"rollout_percentage":100}]}`),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError())
	require.NotNil(t, req.Description)
	assert.Equal(t, "", *req.Description)
	require.NotNil(t, req.RemoveTargetingFlag)
	assert.True(t, *req.RemoveTargetingFlag)
	assert.Nil(t, req.TargetingFlagID)
	assert.Nil(t, req.TargetingFlagFilters)
}

func TestSurveyBuildRequestRejectsInvalidJSON(t *testing.T) {
	ops := SurveyOps{}
	plan := SurveyTFModel{
		Name:           types.StringValue(testSurveyName),
		Type:           types.StringValue(testSurveyType),
		QuestionsJSON:  core.NewNormalizedJSONValue(testSurveyQuestionJSON),
		ConditionsJSON: core.NewNormalizedJSONValue(`{`),
	}

	_, diags := ops.BuildCreateRequest(context.Background(), plan)
	require.True(t, diags.HasError())
}

func TestSurveyMapResponseToModel(t *testing.T) {
	ops := SurveyOps{}
	model := SurveyTFModel{
		QuestionsJSON:               core.NewNormalizedJSONValue(testSurveyQuestionJSON),
		ConditionsJSON:              core.NewNormalizedJSONValue(`{"url":"https://example.com"}`),
		AppearanceJSON:              core.NewNormalizedJSONValue(`{"border_color":"#000000"}`),
		ResponseSamplingDailyLimits: core.NewNormalizedJSONValue(`{"monday":5}`),
		TranslationsJSON:            core.NewNormalizedJSONValue(`{"fr":{"name":"Bonjour"}}`),
		FormContentJSON:             core.NewNormalizedJSONValue(`{"submitButtonText":"Send"}`),
		LinkedInsightID:             types.Int64Unknown(),
		TargetingFlagFiltersJSON:    types.StringUnknown(),
	}
	resp := httpclient.Survey{
		ID:                           testSurveyID,
		Name:                         util.StringPtr("Survey"),
		Description:                  util.StringPtr("Description"),
		Type:                         util.StringPtr(testSurveyType),
		Schedule:                     util.StringPtr("once"),
		LinkedFlag:                   map[string]interface{}{"id": float64(11)},
		TargetingFlag:                map[string]interface{}{"id": float64(12)},
		InternalTargetingFlag:        map[string]interface{}{"id": float64(13)},
		Questions:                    []interface{}{map[string]interface{}{"id": "generated", "type": "open", "question": "How are you?"}},
		Conditions:                   map[string]interface{}{"url": "https://example.com"},
		Appearance:                   map[string]interface{}{"border_color": "#000000"},
		CreatedAt:                    util.StringPtr(testSurveyStartDate),
		CreatedBy:                    map[string]interface{}{"uuid": "user-123"},
		StartDate:                    util.StringPtr(testSurveyStartDate),
		EndDate:                      util.StringPtr(testSurveyEndDate),
		Archived:                     util.BoolPtr(true),
		ResponsesLimit:               util.Int64Ptr(10),
		IterationCount:               util.Int64Ptr(2),
		IterationFrequencyDays:       util.Int64Ptr(7),
		ResponseSamplingStartDate:    util.StringPtr(testSurveySamplingStartDate),
		ResponseSamplingIntervalType: util.StringPtr("week"),
		ResponseSamplingInterval:     util.Int64Ptr(2),
		ResponseSamplingLimit:        util.Int64Ptr(20),
		ResponseSamplingDailyLimits:  map[string]interface{}{"monday": float64(5)},
		EnablePartialResponses:       util.BoolPtr(true),
		EnableIframeEmbedding:        util.BoolPtr(true),
		Translations:                 map[string]interface{}{"fr": map[string]interface{}{"name": "Bonjour"}},
		FormContent:                  map[string]interface{}{"submitButtonText": "Send"},
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())

	assert.Equal(t, testSurveyID, model.ID.ValueString())
	assert.Equal(t, "Survey", model.Name.ValueString())
	assert.Equal(t, "Description", model.Description.ValueString())
	assert.Equal(t, testSurveyType, model.Type.ValueString())
	assert.Equal(t, "once", model.Schedule.ValueString())
	assert.Equal(t, int64(11), model.LinkedFlagID.ValueInt64())
	assert.Equal(t, int64(12), model.TargetingFlagID.ValueInt64())
	assert.True(t, model.LinkedInsightID.IsNull())
	assert.Equal(t, testSurveyQuestionNormalized, model.QuestionsJSON.ValueString())
	assert.JSONEq(t, `{"url":"https://example.com"}`, model.ConditionsJSON.ValueString())
	assert.JSONEq(t, `{"border_color":"#000000"}`, model.AppearanceJSON.ValueString())
	assert.JSONEq(t, `{"monday":5}`, model.ResponseSamplingDailyLimits.ValueString())
	assert.JSONEq(t, `{"fr":{"name":"Bonjour"}}`, model.TranslationsJSON.ValueString())
	assert.JSONEq(t, `{"submitButtonText":"Send"}`, model.FormContentJSON.ValueString())
	assert.True(t, model.TargetingFlagFiltersJSON.IsNull())
	assert.Equal(t, testSurveyStartDate, model.StartDate.ValueString())
	assert.Equal(t, testSurveyEndDate, model.EndDate.ValueString())
	assert.True(t, model.Archived.ValueBool())
	assert.Equal(t, int64(10), model.ResponsesLimit.ValueInt64())
	assert.Equal(t, int64(2), model.IterationCount.ValueInt64())
	assert.Equal(t, int64(7), model.IterationFrequencyDays.ValueInt64())
	assert.Equal(t, testSurveySamplingStartDate, model.ResponseSamplingStartDate.ValueString())
	assert.Equal(t, "week", model.ResponseSamplingIntervalType.ValueString())
	assert.Equal(t, int64(2), model.ResponseSamplingInterval.ValueInt64())
	assert.Equal(t, int64(20), model.ResponseSamplingLimit.ValueInt64())
	assert.True(t, model.EnablePartialResponses.ValueBool())
	assert.True(t, model.EnableIframeEmbedding.ValueBool())
	assert.JSONEq(t, `{"uuid":"user-123"}`, model.CreatedByJSON.ValueString())
	assert.JSONEq(t, `{"id":11}`, model.LinkedFlagJSON.ValueString())
	assert.JSONEq(t, `{"id":12}`, model.TargetingFlagJSON.ValueString())
	assert.JSONEq(t, `{"id":13}`, model.InternalTargetingFlagJSON.ValueString())
}

func TestSurveyCRUDWrapperMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == testSurveyCollectionPath:
			_ = json.NewEncoder(w).Encode(httpclient.Survey{ID: testSurveyID})
		case r.Method == http.MethodGet && r.URL.Path == testSurveyResourcePath:
			_ = json.NewEncoder(w).Encode(httpclient.Survey{ID: testSurveyID})
		case r.Method == http.MethodPatch && r.URL.Path == testSurveyResourcePath:
			_ = json.NewEncoder(w).Encode(httpclient.Survey{ID: testSurveyID})
		case r.Method == http.MethodDelete && r.URL.Path == testSurveyResourcePath:
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	ops := SurveyOps{}
	model := SurveyTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testSurveyID)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testSurveyProjectID)},
	}
	req := httpclient.SurveyRequest{Name: testSurveyName, Type: testSurveyType}

	created, err := ops.Create(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, testSurveyID, created.ID)

	read, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSurveyID, read.ID)

	updated, status, err := ops.Update(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSurveyID, updated.ID)

	status, err = ops.Delete(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusNoContent), status)
}

func TestSurveyHelperFallbacks(t *testing.T) {
	questions, diags := parseSurveyQuestions(core.NewNormalizedJSONNull())
	require.False(t, diags.HasError())
	assert.Nil(t, questions)

	normalized, ok := normalizedJSONString(nil, "")
	assert.True(t, ok)
	assert.True(t, normalized.IsNull())

	assert.True(t, jsonStringValue(nil).IsNull())
	assert.True(t, jsonStringValue(map[string]interface{}(nil)).IsNull(), "typed-nil maps should surface as Terraform null, not the literal string \"null\"")
	assert.True(t, jsonStringValue([]interface{}(nil)).IsNull(), "typed-nil slices should surface as Terraform null, not the literal string \"null\"")
	assert.True(t, int64ValueFromMapOrNull(map[string]interface{}{}).IsNull())
	assert.Equal(t, int64(7), int64ValueFromMapOrNull(map[string]interface{}{"id": 7}).ValueInt64())
}

// TestSurveyClearNullableFields verifies that removing a nullable attribute from config sends an explicit JSON null to the API, which PostHog
// interprets as "clear this column." See SurveyRequest tag comments for details.
func TestSurveyClearNullableFields(t *testing.T) {
	ops := SurveyOps{}
	plan := SurveyTFModel{
		Name:          types.StringValue(testSurveyName),
		Type:          types.StringValue(testSurveyType),
		QuestionsJSON: core.NewNormalizedJSONValue(testSurveyQuestionJSON),
	}
	state := SurveyTFModel{
		LinkedFlagID:              types.Int64Value(11),
		LinkedInsightID:           types.Int64Value(12),
		ResponsesLimit:            types.Int64Value(100),
		IterationCount:            types.Int64Value(3),
		IterationFrequencyDays:    types.Int64Value(7),
		ResponseSamplingStartDate: types.StringValue(testSurveySamplingStartDate),
		ResponseSamplingInterval:  types.Int64Value(2),
		ResponseSamplingLimit:     types.Int64Value(50),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError())

	body, err := json.Marshal(req)
	require.NoError(t, err)

	var decoded map[string]interface{}
	require.NoError(t, json.Unmarshal(body, &decoded))

	// All of these fields must be present with a JSON null value
	// so the server actually clears them; an omitted field would leave the
	// existing value untouched.
	for _, key := range httpclient.SurveyNullableJSONFields {
		raw, present := decoded[key]
		assert.True(t, present, "%s must be present in the request body", key)
		assert.Nil(t, raw, "%s must serialize as JSON null to clear", key)
	}
}
