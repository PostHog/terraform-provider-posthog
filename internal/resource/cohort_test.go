package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PostHog annotates a saved cohort's property values with compiled fields
// (bytecode, conditionHash) and defaults the caller never wrote. Those must not
// reach state, or every plan shows drift on an unchanged cohort.
func TestCohortMapResponseToModel_DropsServerComputedFilterFields(t *testing.T) {
	userFilters := `{"properties":{"type":"AND","values":[{"type":"person","key":"email","operator":"icontains","value":"@example.com"}]}}`

	model := CohortTFModel{
		Filters: jsontypes.NewNormalizedValue(userFilters),
	}

	resp := httpclient.Cohort{
		ID:   42,
		Name: util.StringPtr("Internal users"),
		Filters: map[string]interface{}{
			"filterTestAccounts": false,
			"properties": map[string]interface{}{
				"type": "AND",
				"values": []interface{}{
					map[string]interface{}{
						"type":           "person",
						"key":            "email",
						"operator":       "icontains",
						"value":          "@example.com",
						"bytecode":       []interface{}{"_H", float64(1)},
						"bytecode_error": nil,
						"conditionHash":  "5f2e1c",
						"negation":       false,
					},
				},
			},
		},
	}

	diags := CohortOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.JSONEq(t, userFilters, model.Filters.ValueString())
	assert.Equal(t, types.Int64Value(42), model.ID)
	assert.Equal(t, types.StringValue("Internal users"), model.Name)
	// Absent from the response, so both must land as false rather than unknown.
	assert.Equal(t, types.BoolValue(false), model.IsStatic)
	assert.Equal(t, types.BoolValue(false), model.Deleted)
}

// On import there is no prior config to filter against, so the whole API filters
// blob is adopted as canonical JSON.
func TestCohortMapResponseToModel_ImportAdoptsFullFilters(t *testing.T) {
	model := CohortTFModel{}

	resp := httpclient.Cohort{
		ID:       7,
		Name:     util.StringPtr("Beta testers"),
		IsStatic: util.BoolPtr(true),
		Filters: map[string]interface{}{
			"properties": map[string]interface{}{
				"type":   "OR",
				"values": []interface{}{},
			},
		},
	}

	diags := CohortOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.JSONEq(t, `{"properties":{"type":"OR","values":[]}}`, model.Filters.ValueString())
	assert.Equal(t, types.BoolValue(true), model.IsStatic)
}

func TestCohortMapResponseToModel_NoFiltersIsNull(t *testing.T) {
	model := CohortTFModel{}

	diags := CohortOps{}.MapResponseToModel(context.Background(), httpclient.Cohort{
		ID:   9,
		Name: util.StringPtr("Static only"),
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.True(t, model.Filters.IsNull())
}

func TestCohortBuildCreateRequest_RejectsInvalidFiltersJSON(t *testing.T) {
	model := CohortTFModel{
		Name:    types.StringValue("Broken"),
		Filters: jsontypes.NewNormalizedValue(`{"properties":`),
	}

	_, diags := CohortOps{}.BuildCreateRequest(context.Background(), model)
	assert.True(t, diags.HasError(), "expected a diagnostic for unparseable filters")
}

func TestCohortBuildUpdateRequest_ClearsDescriptionAndCarriesDeleted(t *testing.T) {
	plan := CohortTFModel{
		Name:    types.StringValue("Internal users"),
		Deleted: types.BoolValue(true),
	}
	state := CohortTFModel{
		Name:        types.StringValue("Internal users"),
		Description: types.StringValue("was set"),
	}

	req, diags := CohortOps{}.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	require.NotNil(t, req.Description)
	assert.Equal(t, "", *req.Description, "removing description from config should clear it server-side")
	require.NotNil(t, req.Deleted)
	assert.True(t, *req.Deleted, "deleted must be forwarded so destroy soft-deletes the cohort")
}
