package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
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
		config   types.String
		expected types.String
	}{
		"absent value is null": {
			raw:      nil,
			config:   types.StringNull(),
			expected: types.StringNull(),
		},
		"scalar string round-trips": {
			raw:      rawJSON(`"prod"`),
			config:   types.StringValue(`"prod"`),
			expected: types.StringValue(`"prod"`),
		},
		"number round-trips": {
			raw:      rawJSON(`30`),
			config:   types.StringValue(`30`),
			expected: types.StringValue(`30`),
		},
		"whitespace is canonicalized away": {
			raw:      rawJSON(`[ "a",  "b" ]`),
			config:   types.StringValue(`["a","b"]`),
			expected: types.StringValue(`["a","b"]`),
		},
		// PostHog writes back "" for an unset List default and [] for unset
		// List values. Without this, an unconfigured attribute would come back
		// non-null and Terraform would report provider inconsistency.
		"empty string is unset when not configured": {
			raw:      rawJSON(`""`),
			config:   types.StringNull(),
			expected: types.StringNull(),
		},
		"empty array is unset when not configured": {
			raw:      rawJSON(`[]`),
			config:   types.StringNull(),
			expected: types.StringNull(),
		},
		"null is unset when not configured": {
			raw:      rawJSON(`null`),
			config:   types.StringNull(),
			expected: types.StringNull(),
		},
		"explicitly configured empty string is preserved": {
			raw:      rawJSON(`""`),
			config:   types.StringValue(`""`),
			expected: types.StringValue(`""`),
		},
		"explicitly configured empty array is preserved": {
			raw:      rawJSON(`[]`),
			config:   types.StringValue(`[]`),
			expected: types.StringValue(`[]`),
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
	_, err := jsonValueToState(rawJSON(`{`), types.StringNull())
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
		DefaultValueJSON: types.StringValue(`"prod"`),
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
		DefaultValueJSON: types.StringValue(`"a"`),
		ValuesJSON:       types.StringValue(`["a","b"]`),
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
		ValuesJSON: types.StringValue(`["a","b"]`),
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
		DefaultValueJSON: types.StringValue(`"prod"`),
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
	assert.Equal(t, types.StringValue(`"prod"`), model.DefaultValueJSON)
	assert.Equal(t, types.StringNull(), model.ValuesJSON)
}
