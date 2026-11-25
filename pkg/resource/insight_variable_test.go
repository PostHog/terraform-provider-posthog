package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func rawJSON(s string) *json.RawMessage {
	msg := json.RawMessage(s)
	return &msg
}

func TestJSONValueToState(t *testing.T) {
	tests := map[string]struct {
		raw      *json.RawMessage
		config   core.NormalizedJSON
		expected core.NormalizedJSON
	}{
		"absent value is null": {
			raw:      nil,
			config:   core.NewNormalizedJSONNull(),
			expected: core.NewNormalizedJSONNull(),
		},
		"scalar string round-trips": {
			raw:      rawJSON(`"prod"`),
			config:   core.NewNormalizedJSONValue(`"prod"`),
			expected: core.NewNormalizedJSONValue(`"prod"`),
		},
		"number round-trips": {
			raw:      rawJSON(`30`),
			config:   core.NewNormalizedJSONValue(`30`),
			expected: core.NewNormalizedJSONValue(`30`),
		},
		"whitespace is canonicalized away": {
			raw:      rawJSON(`[ "a",  "b" ]`),
			config:   core.NewNormalizedJSONValue(`["a","b"]`),
			expected: core.NewNormalizedJSONValue(`["a","b"]`),
		},
		// PostHog writes back "" for an unset List default and [] for unset
		// List values. Without this, an unconfigured attribute would come back
		// non-null and Terraform would report provider inconsistency.
		"empty string is unset when not configured": {
			raw:      rawJSON(`""`),
			config:   core.NewNormalizedJSONNull(),
			expected: core.NewNormalizedJSONNull(),
		},
		"empty array is unset when not configured": {
			raw:      rawJSON(`[]`),
			config:   core.NewNormalizedJSONNull(),
			expected: core.NewNormalizedJSONNull(),
		},
		"null is unset when not configured": {
			raw:      rawJSON(`null`),
			config:   core.NewNormalizedJSONNull(),
			expected: core.NewNormalizedJSONNull(),
		},
		"explicitly configured empty string is preserved": {
			raw:      rawJSON(`""`),
			config:   core.NewNormalizedJSONValue(`""`),
			expected: core.NewNormalizedJSONValue(`""`),
		},
		"explicitly configured empty array is preserved": {
			raw:      rawJSON(`[]`),
			config:   core.NewNormalizedJSONValue(`[]`),
			expected: core.NewNormalizedJSONValue(`[]`),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := jsonValueToState(tc.raw, tc.config)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestJSONValueToState_InvalidJSON(t *testing.T) {
	_, err := jsonValueToState(rawJSON(`{`), core.NewNormalizedJSONNull())
	assert.Error(t, err)
}

func TestValidateInsightVariableListShape(t *testing.T) {
	tests := map[string]struct {
		variableType string
		values       *json.RawMessage
		defaultValue *json.RawMessage
		expectError  bool
	}{
		"string list is accepted": {
			variableType: "List",
			values:       rawJSON(`["a","b"]`),
			defaultValue: rawJSON(`"a"`),
		},
		"numeric list entries are rejected": {
			variableType: "List",
			values:       rawJSON(`[1,2]`),
			expectError:  true,
		},
		"object list entries are rejected": {
			variableType: "List",
			values:       rawJSON(`[{"value":"a"}]`),
			expectError:  true,
		},
		"non-array values are rejected": {
			variableType: "List",
			values:       rawJSON(`"a"`),
			expectError:  true,
		},
		"numeric list default is rejected": {
			variableType: "List",
			defaultValue: rawJSON(`1`),
			expectError:  true,
		},
		// Only List variables get coerced server-side, so every other type is
		// free to hold whatever JSON the user configured.
		"numbers are fine for a Number variable": {
			variableType: "Number",
			defaultValue: rawJSON(`30`),
		},
		"arrays are not inspected for other types": {
			variableType: "String",
			values:       rawJSON(`[1,2]`),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			diags := validateInsightVariableListShape(tc.variableType, httpclient.InsightVariableRequest{
				Values:       tc.values,
				DefaultValue: tc.defaultValue,
			})
			assert.Equal(t, tc.expectError, diags.HasError())
		})
	}
}

func TestInsightVariableBuildCreateRequest(t *testing.T) {
	model := InsightVariableResourceTFModel{
		Name:             types.StringValue("Lookup"),
		Type:             types.StringValue("String"),
		DefaultValueJSON: core.NewNormalizedJSONValue(`"prod"`),
	}

	req, diags := InsightVariableOps{}.BuildCreateRequest(context.Background(), model)

	require.False(t, diags.HasError())
	require.NotNil(t, req.Name)
	assert.Equal(t, "Lookup", *req.Name)
	require.NotNil(t, req.Type)
	assert.Equal(t, "String", *req.Type)
	require.NotNil(t, req.DefaultValue)
	assert.JSONEq(t, `"prod"`, string(*req.DefaultValue))
	assert.Nil(t, req.Values, "unset attributes must be omitted from the request")
}

func TestInsightVariableBuildUpdateRequest_ClearsRemovedValues(t *testing.T) {
	plan := InsightVariableResourceTFModel{
		Name: types.StringValue("Lookup"),
		Type: types.StringValue("List"),
	}
	state := InsightVariableResourceTFModel{
		Name:             types.StringValue("Lookup"),
		Type:             types.StringValue("List"),
		DefaultValueJSON: core.NewNormalizedJSONValue(`"a"`),
		ValuesJSON:       core.NewNormalizedJSONValue(`["a","b"]`),
	}

	req, diags := InsightVariableOps{}.BuildUpdateRequest(context.Background(), plan, state)

	require.False(t, diags.HasError())
	require.NotNil(t, req.DefaultValue, "removing an attribute must send an explicit null, not omit it")
	assert.Equal(t, "null", string(*req.DefaultValue))
	require.NotNil(t, req.Values)
	assert.Equal(t, "null", string(*req.Values))
}

func TestInsightVariableBuildUpdateRequest_KeepsUnchangedValues(t *testing.T) {
	model := InsightVariableResourceTFModel{
		Name:       types.StringValue("Lookup"),
		Type:       types.StringValue("List"),
		ValuesJSON: core.NewNormalizedJSONValue(`["a","b"]`),
	}

	req, diags := InsightVariableOps{}.BuildUpdateRequest(context.Background(), model, model)

	require.False(t, diags.HasError())
	require.NotNil(t, req.Values)
	assert.JSONEq(t, `["a","b"]`, string(*req.Values))
}

func TestInsightVariableMapResponseToModel(t *testing.T) {
	name := "Lookup"
	variableType := "String"
	codeName := "lookup"
	createdAt := "2026-07-29T15:58:44.275733Z"

	model := InsightVariableResourceTFModel{
		DefaultValueJSON: core.NewNormalizedJSONValue(`"prod"`),
	}

	diags := InsightVariableOps{}.MapResponseToModel(context.Background(), httpclient.InsightVariable{
		ID:           "019fae99-842f-0000-f983-cc2171e27686",
		Name:         &name,
		Type:         &variableType,
		CodeName:     &codeName,
		CreatedAt:    &createdAt,
		DefaultValue: rawJSON(`"prod"`),
	}, &model)

	require.False(t, diags.HasError())
	assert.Equal(t, types.StringValue("019fae99-842f-0000-f983-cc2171e27686"), model.ID)
	assert.Equal(t, types.StringValue("Lookup"), model.Name)
	assert.Equal(t, types.StringValue("String"), model.Type)
	assert.Equal(t, types.StringValue("lookup"), model.CodeName)
	assert.Equal(t, types.StringValue(createdAt), model.CreatedAt)
	assert.Equal(t, core.NewNormalizedJSONValue(`"prod"`), model.DefaultValueJSON)
	assert.Equal(t, core.NewNormalizedJSONNull(), model.ValuesJSON)
}

func TestInsightVariableJSONSemanticEquality(t *testing.T) {
	ctx := context.Background()
	ops := InsightVariableOps{}
	tests := []struct {
		name      string
		attribute string
		planned   string
		response  string
		equal     bool
	}{
		{"list whitespace", "values_json", `["us", "eu"]`, `["us","eu"]`, true},
		{"changed list item", "values_json", `["us", "eu"]`, `["us","apac"]`, false},
		{"reordered list", "values_json", `["us", "eu"]`, `["eu","us"]`, false},
		{"default number formatting", "default_value_json", `2.0`, `2`, true},
		{"changed default number", "default_value_json", `2.0`, `3`, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			valueType := ops.Schema().Attributes[tc.attribute].GetType()
			planned, err := valueType.ValueFromTerraform(ctx, tftypes.NewValue(tftypes.String, tc.planned))
			require.NoError(t, err)

			var model InsightVariableResourceTFModel
			response := httpclient.InsightVariable{Values: rawJSON(tc.response)}
			if tc.attribute == "default_value_json" {
				response = httpclient.InsightVariable{DefaultValue: rawJSON(tc.response)}
			}
			require.False(t, ops.MapResponseToModel(ctx, response, &model).HasError())

			var actual basetypes.StringValuable = model.ValuesJSON
			if tc.attribute == "default_value_json" {
				actual = model.DefaultValueJSON
			}
			equal := planned.Equal(actual)
			if comparable, ok := planned.(basetypes.StringValuableWithSemanticEquals); !equal && ok {
				var diags diag.Diagnostics
				equal, diags = comparable.StringSemanticEquals(ctx, actual)
				require.False(t, diags.HasError(), "%v", diags)
			}
			assert.Equal(t, tc.equal, equal)
		})
	}
}
