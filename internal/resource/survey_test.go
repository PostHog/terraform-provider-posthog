package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
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

func TestNormalizeSurveyQuestionsForState(t *testing.T) {
	apiQuestions := []interface{}{
		map[string]interface{}{
			"id":       "generated-id",
			"type":     "open",
			"question": "How are you?",
		},
	}

	t.Run("strips generated ids during import", func(t *testing.T) {
		result, err := normalizeSurveyQuestionsForState(apiQuestions, "")
		require.NoError(t, err)
		assert.Equal(t, testSurveyQuestionNormalized, result)
	})

	t.Run("preserves user field shape during read", func(t *testing.T) {
		result, err := normalizeSurveyQuestionsForState(apiQuestions, testSurveyQuestionJSON)
		require.NoError(t, err)
		assert.Equal(t, testSurveyQuestionNormalized, result)
	})
}

func TestSurveyBuildRequestPreservesWriteOnlyFilters(t *testing.T) {
	ops := SurveyOps{}

	plan := SurveyTFModel{
		Name:          types.StringValue(testSurveyName),
		Type:          types.StringValue(testSurveyType),
		QuestionsJSON: types.StringValue(testSurveyQuestionJSON),
	}
	state := SurveyTFModel{
		TargetingFlagFiltersJSON: types.StringValue(`{"groups":[{"properties":[{"key":"email","operator":"icontains","type":"person","value":"@example.com"}],"rollout_percentage":100}]}`),
	}

	req, diags := ops.buildRequest(context.Background(), plan, state)
	require.False(t, diags.HasError())
	require.NotNil(t, req.TargetingFlagFilters)
	assert.Equal(t,
		map[string]interface{}{
			"groups": []interface{}{
				map[string]interface{}{
					"properties": []interface{}{
						map[string]interface{}{
							"key":      "email",
							"operator": "icontains",
							"type":     "person",
							"value":    "@example.com",
						},
					},
					"rollout_percentage": float64(100),
				},
			},
		},
		req.TargetingFlagFilters,
	)
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
		QuestionsJSON:                types.StringValue(testSurveyQuestionJSON),
		ConditionsJSON:               types.StringValue(`{"url":"https://example.com"}`),
		AppearanceJSON:               types.StringValue(`{"border_color":"#000000"}`),
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
		ResponseSamplingDailyLimits:  types.StringValue(`{"monday":5}`),
		EnablePartialResponses:       types.BoolValue(true),
		EnableIframeEmbedding:        types.BoolValue(true),
		TranslationsJSON:             types.StringValue(`{"fr":{"name":"Bonjour"}}`),
		CreateInFolder:               types.StringValue("folder-123"),
		FormContentJSON:              types.StringValue(`{"submitButtonText":"Send"}`),
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
		QuestionsJSON: types.StringValue(testSurveyQuestionJSON),
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
		QuestionsJSON:  types.StringValue(testSurveyQuestionJSON),
		ConditionsJSON: types.StringValue(`{`),
	}

	_, diags := ops.BuildCreateRequest(context.Background(), plan)
	require.True(t, diags.HasError())
}

func TestSurveyMapResponseToModel(t *testing.T) {
	ops := SurveyOps{}
	model := SurveyTFModel{
		QuestionsJSON:               types.StringValue(testSurveyQuestionJSON),
		ConditionsJSON:              types.StringValue(`{"url":"https://example.com"}`),
		AppearanceJSON:              types.StringValue(`{"border_color":"#000000"}`),
		ResponseSamplingDailyLimits: types.StringValue(`{"monday":5}`),
		TranslationsJSON:            types.StringValue(`{"fr":{"name":"Bonjour"}}`),
		FormContentJSON:             types.StringValue(`{"submitButtonText":"Send"}`),
		LinkedInsightID:             types.Int64Unknown(),
		TargetingFlagFiltersJSON:    types.StringUnknown(),
	}
	resp := httpclient.Survey{
		ID:                           testSurveyID,
		Name:                         stringPtr("Survey"),
		Description:                  stringPtr("Description"),
		Type:                         stringPtr(testSurveyType),
		Schedule:                     stringPtr("once"),
		LinkedFlag:                   map[string]interface{}{"id": float64(11)},
		TargetingFlag:                map[string]interface{}{"id": float64(12)},
		InternalTargetingFlag:        map[string]interface{}{"id": float64(13)},
		Questions:                    []interface{}{map[string]interface{}{"id": "generated", "type": "open", "question": "How are you?"}},
		Conditions:                   map[string]interface{}{"url": "https://example.com"},
		Appearance:                   map[string]interface{}{"border_color": "#000000"},
		CreatedAt:                    stringPtr(testSurveyStartDate),
		CreatedBy:                    map[string]interface{}{"uuid": "user-123"},
		StartDate:                    stringPtr(testSurveyStartDate),
		EndDate:                      stringPtr(testSurveyEndDate),
		Archived:                     boolPtr(true),
		ResponsesLimit:               int64Ptr(10),
		IterationCount:               int64Ptr(2),
		IterationFrequencyDays:       int64Ptr(7),
		ResponseSamplingStartDate:    stringPtr(testSurveySamplingStartDate),
		ResponseSamplingIntervalType: stringPtr("week"),
		ResponseSamplingInterval:     int64Ptr(2),
		ResponseSamplingLimit:        int64Ptr(20),
		ResponseSamplingDailyLimits:  map[string]interface{}{"monday": float64(5)},
		EnablePartialResponses:       boolPtr(true),
		EnableIframeEmbedding:        boolPtr(true),
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
		case r.Method == http.MethodPut && r.URL.Path == testSurveyResourcePath:
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
	questions, diags := parseSurveyQuestions(types.StringNull())
	require.False(t, diags.HasError())
	assert.Nil(t, questions)

	normalized, ok := normalizedJSONString(nil, "")
	assert.True(t, ok)
	assert.True(t, normalized.IsNull())

	assert.True(t, jsonStringValue(nil).IsNull())
	assert.True(t, int64ValueOrNull(nil).IsNull())
	assert.True(t, int64ValueFromMapOrNull(map[string]interface{}{}).IsNull())
	assert.Equal(t, int64(7), int64ValueFromMapOrNull(map[string]interface{}{"id": 7}).ValueInt64())
	assert.True(t, valueStringOrEmpty(types.StringNull()) == "")
	assert.False(t, boolPtrFromValue(types.BoolNull()) != nil)
	assert.False(t, stringPtrFromValue(types.StringNull()) != nil)
}

func stringPtr(value string) *string {
	return &value
}

func boolPtr(value bool) *bool {
	return &value
}

func int64Ptr(value int64) *int64 {
	return &value
}
