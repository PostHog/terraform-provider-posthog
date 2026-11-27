package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testEndpointQuery = `{"kind":"HogQLQuery","query":"SELECT count() FROM events"}`

func stringMap(t *testing.T, values map[string]string) types.Map {
	t.Helper()
	m, diags := types.MapValueFrom(context.Background(), types.StringType, values)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	return m
}

// endpointState is a model as it sits in state after an apply.
func endpointState(t *testing.T) EndpointResourceTFModel {
	t.Helper()
	return EndpointResourceTFModel{
		Name:                        types.StringValue("daily-signups"),
		QueryJSON:                   core.NewNormalizedJSONValue(testEndpointQuery),
		Description:                 types.StringValue("Signups per day"),
		DataFreshnessSeconds:        types.Int64Value(3600),
		IsActive:                    types.BoolValue(true),
		IsMaterialized:              types.BoolValue(true),
		BucketOverrides:             stringMap(t, map[string]string{"timestamp": "hour"}),
		Tags:                        stringSet(t, "growth"),
		OptionalBreakdownProperties: types.SetNull(types.StringType),
	}
}

func TestEndpointBuildCreateRequest(t *testing.T) {
	model := endpointState(t)

	req, diags := EndpointOps{}.BuildCreateRequest(context.Background(), model)

	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	assert.Equal(t, util.StringPtr("daily-signups"), req.Name)
	assert.Equal(t, map[string]interface{}{"kind": "HogQLQuery", "query": "SELECT count() FROM events"}, req.Query)
	assert.Equal(t, util.StringPtr("Signups per day"), req.Description)
	assert.Equal(t, util.Int64Ptr(3600), req.DataFreshnessSeconds)
	assert.Equal(t, util.BoolPtr(true), req.IsActive)
	assert.Equal(t, util.BoolPtr(true), req.IsMaterialized)
	require.NotNil(t, req.BucketOverrides)
	assert.Equal(t, map[string]string{"timestamp": "hour"}, *req.BucketOverrides)
	require.NotNil(t, req.Tags)
	assert.Equal(t, []string{"growth"}, *req.Tags)
	assert.Nil(t, req.OptionalBreakdownProperties, "an unset attribute must leave PostHog's default alone")
}

func TestEndpointBuildCreateRequest_RejectsQueryThatIsNotAnObject(t *testing.T) {
	model := endpointState(t)
	model.QueryJSON = core.NewNormalizedJSONValue(`["SELECT 1"]`)

	_, diags := EndpointOps{}.BuildCreateRequest(context.Background(), model)

	require.True(t, diags.HasError())
	assert.Equal(t, "Invalid query_json", diags.Errors()[0].Summary())
}

func TestEndpointBuildUpdateRequest(t *testing.T) {
	tests := map[string]struct {
		change func(t *testing.T, plan *EndpointResourceTFModel)
		check  func(t *testing.T, req httpclient.EndpointRequest)
	}{
		"an unchanged plan sends nothing, so PostHog neither versions nor re-materializes": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				assert.Equal(t, httpclient.EndpointRequest{}, req)
			},
		},
		"a query that differs only in formatting is not sent": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {
				plan.QueryJSON = core.NewNormalizedJSONValue(`{ "query": "SELECT count() FROM events", "kind": "HogQLQuery" }`)
			},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				assert.Nil(t, req.Query)
			},
		},
		"a changed query is sent, which creates a new version": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {
				plan.QueryJSON = core.NewNormalizedJSONValue(`{"kind":"HogQLQuery","query":"SELECT 1"}`)
			},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				assert.Equal(t, map[string]interface{}{"kind": "HogQLQuery", "query": "SELECT 1"}, req.Query)
				assert.Nil(t, req.IsMaterialized)
				assert.Nil(t, req.DataFreshnessSeconds)
			},
		},
		"a removed description is cleared with an empty string": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {
				plan.Description = types.StringNull()
			},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				assert.Equal(t, util.StringPtr(""), req.Description)
			},
		},
		"removed tags are cleared with an empty list": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {
				plan.Tags = types.SetNull(types.StringType)
			},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				require.NotNil(t, req.Tags)
				assert.Empty(t, *req.Tags)
			},
		},
		"removed bucket overrides are cleared with an empty map": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {
				plan.BucketOverrides = types.MapNull(types.StringType)
			},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				require.NotNil(t, req.BucketOverrides)
				assert.Empty(t, *req.BucketOverrides)
			},
		},
		"changed scalars are sent and unchanged ones are not": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {
				plan.DataFreshnessSeconds = types.Int64Value(86400)
				plan.IsMaterialized = types.BoolValue(false)
			},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				assert.Equal(t, util.Int64Ptr(86400), req.DataFreshnessSeconds)
				assert.Equal(t, util.BoolPtr(false), req.IsMaterialized)
				assert.Nil(t, req.IsActive)
			},
		},
		"added optional breakdown properties are sent": {
			change: func(t *testing.T, plan *EndpointResourceTFModel) {
				plan.OptionalBreakdownProperties = stringSet(t, "$browser")
			},
			check: func(t *testing.T, req httpclient.EndpointRequest) {
				require.NotNil(t, req.OptionalBreakdownProperties)
				assert.Equal(t, []string{"$browser"}, *req.OptionalBreakdownProperties)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			state := endpointState(t)
			plan := endpointState(t)
			tc.change(t, &plan)

			req, diags := EndpointOps{}.BuildUpdateRequest(context.Background(), plan, state)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Nil(t, req.Name, "PostHog ignores a name on update")
			tc.check(t, req)
		})
	}
}

func TestEndpointMapResponseToModel(t *testing.T) {
	variableID := "0196f8c0-0000-0000-0000-00000000a001"
	configuredQuery := `{"kind":"HogQLQuery","query":"SELECT {variables.event_name}","variables":{"` + variableID +
		`":{"variableId":"` + variableID + `","code_name":"event_name"}}}`
	// PostHog stores every optional query field, mostly null, and fills in the
	// variable's value from its default.
	serverQuery := map[string]interface{}{
		"kind":          "HogQLQuery",
		"query":         "SELECT {variables.event_name}",
		"connectionId":  nil,
		"explain":       nil,
		"modifiers":     nil,
		"sendRawQuery":  nil,
		"tags":          nil,
		"response":      nil,
		"name":          nil,
		"filters":       nil,
		"values":        nil,
		"variables":     map[string]interface{}{variableID: map[string]interface{}{"variableId": variableID, "code_name": "event_name", "value": "$pageview", "isNull": false}},
		"version":       nil,
		"refreshPolicy": nil,
	}

	tests := map[string]struct {
		config    core.NormalizedJSON
		materials *httpclient.EndpointMaterialization
		overrides map[string]string
		wantQuery string
		wantMat   bool
		wantOver  types.Map
	}{
		"server-added query fields are ignored when the configuration sets the query": {
			config:    core.NewNormalizedJSONValue(configuredQuery),
			materials: &httpclient.EndpointMaterialization{Enabled: util.BoolPtr(false)},
			wantQuery: configuredQuery,
			wantOver:  types.MapNull(types.StringType),
		},
		"import keeps every non-null query field": {
			config:    core.NewNormalizedJSONNull(),
			materials: &httpclient.EndpointMaterialization{Enabled: util.BoolPtr(false)},
			wantQuery: `{"kind":"HogQLQuery","query":"SELECT {variables.event_name}","variables":{"` + variableID +
				`":{"variableId":"` + variableID + `","code_name":"event_name","value":"$pageview","isNull":false}}}`,
			wantOver: types.MapNull(types.StringType),
		},
		"is_materialized follows materialization.enabled, not whether the first run finished": {
			config:    core.NewNormalizedJSONValue(configuredQuery),
			materials: &httpclient.EndpointMaterialization{Enabled: util.BoolPtr(true)},
			overrides: map[string]string{"timestamp": "hour"},
			wantQuery: configuredQuery,
			wantMat:   true,
			wantOver:  stringMap(t, map[string]string{"timestamp": "hour"}),
		},
		"bucket overrides of an endpoint that is not materialized read as null": {
			config:    core.NewNormalizedJSONValue(configuredQuery),
			materials: &httpclient.EndpointMaterialization{Enabled: util.BoolPtr(false)},
			overrides: map[string]string{"timestamp": "hour"},
			wantQuery: configuredQuery,
			wantOver:  types.MapNull(types.StringType),
		},
		"a missing materialization object reads as not materialized": {
			config:    core.NewNormalizedJSONValue(configuredQuery),
			wantQuery: configuredQuery,
			wantOver:  types.MapNull(types.StringType),
		},
		// PostHog fills in query.variables from the code names in the SQL when the
		// configuration does not declare them.
		"variables PostHog fills in for an undeclared code name are ignored": {
			config:    core.NewNormalizedJSONValue(`{"kind":"HogQLQuery","query":"SELECT {variables.event_name}"}`),
			materials: &httpclient.EndpointMaterialization{Enabled: util.BoolPtr(false)},
			wantQuery: `{"kind":"HogQLQuery","query":"SELECT {variables.event_name}"}`,
			wantOver:  types.MapNull(types.StringType),
		},
		// Clearing bucket overrides with {} leaves {} stored on the version.
		"cleared bucket overrides on a materialized endpoint read as null": {
			config:    core.NewNormalizedJSONValue(configuredQuery),
			materials: &httpclient.EndpointMaterialization{Enabled: util.BoolPtr(true)},
			overrides: map[string]string{},
			wantQuery: configuredQuery,
			wantMat:   true,
			wantOver:  types.MapNull(types.StringType),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := EndpointResourceTFModel{
				QueryJSON:                   tc.config,
				BucketOverrides:             types.MapNull(types.StringType),
				Tags:                        types.SetNull(types.StringType),
				OptionalBreakdownProperties: types.SetNull(types.StringType),
			}
			resp := httpclient.Endpoint{
				ID:                          "0196f8c0-0000-0000-0000-000000000001",
				Name:                        "daily-signups",
				Description:                 util.StringPtr(""),
				Query:                       serverQuery,
				IsActive:                    util.BoolPtr(true),
				DataFreshnessSeconds:        util.Int64Ptr(86400),
				Materialization:             tc.materials,
				BucketOverrides:             tc.overrides,
				CurrentVersion:              util.Int64Ptr(2),
				URL:                         util.StringPtr("https://us.posthog.com/api/projects/1/endpoints/daily-signups/run"),
				OptionalBreakdownProperties: []string{},
			}

			diags := EndpointOps{}.MapResponseToModel(context.Background(), resp, &model)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, "daily-signups", model.ID.ValueString(), "the id is the name PostHog addresses the endpoint by")
			assert.Equal(t, "daily-signups", model.Name.ValueString())
			assert.True(t, model.Description.IsNull(), "an empty description reads as null")
			assert.JSONEq(t, tc.wantQuery, model.QueryJSON.ValueString())
			assert.Equal(t, tc.wantMat, model.IsMaterialized.ValueBool())
			assert.Equal(t, tc.wantOver, model.BucketOverrides)
			assert.Equal(t, int64(2), model.CurrentVersion.ValueInt64())
			assert.Equal(t, int64(86400), model.DataFreshnessSeconds.ValueInt64())
			assert.True(t, model.Tags.IsNull())
			assert.True(t, model.OptionalBreakdownProperties.IsNull())
		})
	}
}

func TestEndpointMapResponseToModel_FailsWhenMaterializationOfANewVersionFailed(t *testing.T) {
	model := endpointState(t)
	resp := httpclient.Endpoint{
		Name:                 "daily-signups",
		MaterializationError: util.StringPtr("query cannot be materialized"),
	}

	diags := EndpointOps{}.MapResponseToModel(context.Background(), resp, &model)

	require.True(t, diags.HasError())
	assert.Equal(t, "Endpoint materialization failed", diags.Errors()[0].Summary())
	assert.Contains(t, diags.Errors()[0].Detail(), "query cannot be materialized")
}

func TestValidateEndpointPlan(t *testing.T) {
	tests := map[string]struct {
		plan      EndpointResourceTFModel
		config    EndpointResourceTFModel
		wantError string
	}{
		"defaults are valid": {
			plan:   EndpointResourceTFModel{IsActive: types.BoolUnknown(), IsMaterialized: types.BoolUnknown(), BucketOverrides: types.MapNull(types.StringType)},
			config: EndpointResourceTFModel{IsActive: types.BoolNull(), IsMaterialized: types.BoolNull()},
		},
		"materializing an inactive endpoint is rejected": {
			plan:      EndpointResourceTFModel{IsActive: types.BoolValue(false), IsMaterialized: types.BoolValue(true), BucketOverrides: types.MapNull(types.StringType)},
			config:    EndpointResourceTFModel{IsActive: types.BoolValue(false), IsMaterialized: types.BoolValue(true)},
			wantError: "Cannot materialize an inactive endpoint",
		},
		"bucket overrides without materialization are rejected": {
			plan:      EndpointResourceTFModel{IsActive: types.BoolUnknown(), IsMaterialized: types.BoolUnknown(), BucketOverrides: stringMap(t, map[string]string{"timestamp": "hour"})},
			config:    EndpointResourceTFModel{IsActive: types.BoolNull(), IsMaterialized: types.BoolNull()},
			wantError: "bucket_overrides needs materialization",
		},
		"bucket overrides on a materialized endpoint are valid": {
			plan:   EndpointResourceTFModel{IsActive: types.BoolValue(true), IsMaterialized: types.BoolValue(true), BucketOverrides: stringMap(t, map[string]string{"timestamp": "hour"})},
			config: EndpointResourceTFModel{IsActive: types.BoolNull(), IsMaterialized: types.BoolValue(true)},
		},
		"a materialization flag not known until apply is not judged": {
			plan:   EndpointResourceTFModel{IsActive: types.BoolUnknown(), IsMaterialized: types.BoolUnknown(), BucketOverrides: stringMap(t, map[string]string{"timestamp": "hour"})},
			config: EndpointResourceTFModel{IsActive: types.BoolNull(), IsMaterialized: types.BoolUnknown()},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			diags := validateEndpointPlan(tc.plan, tc.config)

			if tc.wantError == "" {
				assert.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
				return
			}
			require.True(t, diags.HasError())
			assert.Equal(t, tc.wantError, diags.Errors()[0].Summary())
		})
	}
}

func TestEndpointQueryChanged(t *testing.T) {
	tests := map[string]struct {
		plan, state core.NormalizedJSON
		want        bool
	}{
		"same content in another key order is unchanged": {
			plan:  core.NewNormalizedJSONValue(`{"query":"SELECT 1","kind":"HogQLQuery"}`),
			state: core.NewNormalizedJSONValue(`{"kind":"HogQLQuery","query":"SELECT 1"}`),
		},
		"different SQL is a change": {
			plan:  core.NewNormalizedJSONValue(`{"kind":"HogQLQuery","query":"SELECT 2"}`),
			state: core.NewNormalizedJSONValue(`{"kind":"HogQLQuery","query":"SELECT 1"}`),
			want:  true,
		},
		"a query not known until apply counts as a change": {
			plan:  core.NewNormalizedJSONUnknown(),
			state: core.NewNormalizedJSONValue(`{"kind":"HogQLQuery","query":"SELECT 1"}`),
			want:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			changed, diags := endpointQueryChanged(context.Background(), tc.plan, tc.state)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, tc.want, changed)
		})
	}
}
