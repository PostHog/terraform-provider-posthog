package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHogFunctionSchema_SensitiveInputsJSONIsSensitive(t *testing.T) {
	ops := HogFunctionOps{}
	s := ops.Schema()

	attr, ok := s.Attributes["sensitive_inputs_json"]
	require.True(t, ok, "sensitive_inputs_json attribute must exist in schema")

	strAttr, ok := attr.(schema.StringAttribute)
	require.True(t, ok, "sensitive_inputs_json must be a StringAttribute")

	assert.True(t, strAttr.Sensitive,
		"sensitive_inputs_json must be marked Sensitive (issue #55)")
}

func TestHogFunctionSchema_InputsJSONIsNotSensitive(t *testing.T) {
	ops := HogFunctionOps{}
	s := ops.Schema()

	attr, ok := s.Attributes["inputs_json"]
	require.True(t, ok, "inputs_json attribute must exist in schema")

	strAttr, ok := attr.(schema.StringAttribute)
	require.True(t, ok, "inputs_json must be a StringAttribute")

	assert.False(t, strAttr.Sensitive,
		"inputs_json should NOT be sensitive - use sensitive_inputs_json for secrets")
}

func TestBuildCreateRequest_MergesSensitiveInputs(t *testing.T) {
	ctx := context.Background()
	tests := map[string]struct {
		inputsJSON          string
		sensitiveInputsJSON string
		expectedKeys        []string
		expectedInputValue  map[string]string // key -> expected "value" field
	}{
		"only inputs_json": {
			inputsJSON:          `{"url":{"value":"https://example.com"},"method":{"value":"POST"}}`,
			sensitiveInputsJSON: "",
			expectedKeys:        []string{"url", "method"},
			expectedInputValue:  map[string]string{"url": "https://example.com"},
		},
		"only sensitive_inputs_json": {
			inputsJSON:          "",
			sensitiveInputsJSON: `{"api_key":{"value":"secret123"}}`,
			expectedKeys:        []string{"api_key"},
		},
		"both fields merged": {
			inputsJSON:          `{"url":{"value":"https://example.com"}}`,
			sensitiveInputsJSON: `{"api_key":{"value":"secret123"}}`,
			expectedKeys:        []string{"url", "api_key"},
		},
		"sensitive takes precedence on key conflict": {
			inputsJSON:          `{"api_key":{"value":"visible"}}`,
			sensitiveInputsJSON: `{"api_key":{"value":"secret"}}`,
			expectedKeys:        []string{"api_key"},
			expectedInputValue:  map[string]string{"api_key": "secret"},
		},
	}

	ops := HogFunctionOps{}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := HogFunctionResourceTFModel{}
			if tc.inputsJSON != "" {
				model.InputsJSON = types.StringValue(tc.inputsJSON)
			} else {
				model.InputsJSON = types.StringNull()
			}
			if tc.sensitiveInputsJSON != "" {
				model.SensitiveInputsJSON = types.StringValue(tc.sensitiveInputsJSON)
			} else {
				model.SensitiveInputsJSON = types.StringNull()
			}

			req, diags := ops.BuildCreateRequest(ctx, model)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)

			for _, key := range tc.expectedKeys {
				assert.Contains(t, req.Inputs, key)
			}

			for key, wantVal := range tc.expectedInputValue {
				input := req.Inputs[key].(map[string]interface{})
				assert.Equal(t, wantVal, input["value"])
			}
		})
	}
}

func TestMapResponseToModel_SplitsInputsBack(t *testing.T) {
	ctx := context.Background()
	tests := map[string]struct {
		inputsJSON            types.String
		sensitiveInputsJSON   types.String
		respInputs            map[string]interface{}
		expectedInputs        string
		expectedSensitive     string
		expectInputsIsNull    bool
		expectSensitiveIsNull bool
	}{
		"splits keys to original fields": {
			inputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			sensitiveInputsJSON: types.StringValue(`{"api_key":{"value":"secret123"}}`),
			respInputs: map[string]interface{}{
				"url":     map[string]interface{}{"value": "https://example.com"},
				"api_key": map[string]interface{}{"value": "secret123"},
			},
			expectedInputs:    `{"url":{"value":"https://example.com"}}`,
			expectedSensitive: `{"api_key":{"value":"secret123"}}`,
		},
		"null sensitive_inputs_json stays null": {
			inputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			sensitiveInputsJSON: types.StringNull(),
			respInputs: map[string]interface{}{
				"url": map[string]interface{}{"value": "https://example.com"},
			},
			expectedInputs:        `{"url":{"value":"https://example.com"}}`,
			expectSensitiveIsNull: true,
		},
		"strips server fields from both": {
			inputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			sensitiveInputsJSON: types.StringValue(`{"api_key":{"value":"secret"}}`),
			respInputs: map[string]interface{}{
				"url":     map[string]interface{}{"value": "https://example.com", "bytecode": []interface{}{"_H", 1}, "order": 0},
				"api_key": map[string]interface{}{"value": "secret", "bytecode": []interface{}{"_H", 2}, "order": 1},
			},
			expectedInputs:    `{"url":{"value":"https://example.com"}}`,
			expectedSensitive: `{"api_key":{"value":"secret"}}`,
		},
		"only sensitive_inputs_json set - must not leak into inputs_json": {
			inputsJSON:          types.StringNull(),
			sensitiveInputsJSON: types.StringValue(`{"api_key":{"value":"secret123"}}`),
			respInputs: map[string]interface{}{
				"api_key": map[string]interface{}{"value": "secret123"},
			},
			expectInputsIsNull: true,
			expectedSensitive:  `{"api_key":{"value":"secret123"}}`,
		},
		"empty resp.Inputs with only sensitive_inputs_json zeroes sensitive field": {
			inputsJSON:          types.StringNull(),
			sensitiveInputsJSON: types.StringValue(`{"api_key":{"value":"secret123"}}`),
			respInputs:          map[string]interface{}{},
			expectInputsIsNull:  true,
			expectedSensitive:   `{}`,
		},
		"import: both fields null populates inputs_json with all API inputs": {
			inputsJSON:          types.StringNull(),
			sensitiveInputsJSON: types.StringNull(),
			respInputs: map[string]interface{}{
				"url":     map[string]interface{}{"value": "https://example.com"},
				"api_key": map[string]interface{}{"value": "secret123"},
			},
			expectedInputs:        `{"api_key":{"value":"secret123"},"url":{"value":"https://example.com"}}`,
			expectSensitiveIsNull: true,
		},
		"duplicate key in both fields appears in both after read": {
			inputsJSON:          types.StringValue(`{"shared_key":{"value":"from_inputs"}}`),
			sensitiveInputsJSON: types.StringValue(`{"shared_key":{"value":"from_sensitive"}}`),
			respInputs: map[string]interface{}{
				"shared_key": map[string]interface{}{"value": "api_value"},
			},
			expectedInputs:    `{"shared_key":{"value":"api_value"}}`,
			expectedSensitive: `{"shared_key":{"value":"api_value"}}`,
		},
	}

	ops := HogFunctionOps{}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := &HogFunctionResourceTFModel{
				InputsJSON:          tc.inputsJSON,
				SensitiveInputsJSON: tc.sensitiveInputsJSON,
			}

			resp := httpclient.HogFunction{
				ID:     "test-id",
				Inputs: tc.respInputs,
			}

			diags := ops.MapResponseToModel(ctx, resp, model)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)

			if tc.expectInputsIsNull {
				assert.True(t, model.InputsJSON.IsNull(), "inputs_json should remain null, got: %s", model.InputsJSON.ValueString())
			} else {
				assert.Equal(t, tc.expectedInputs, model.InputsJSON.ValueString())
			}

			if tc.expectSensitiveIsNull {
				assert.True(t, model.SensitiveInputsJSON.IsNull())
			} else {
				assert.Equal(t, tc.expectedSensitive, model.SensitiveInputsJSON.ValueString())
			}
		})
	}
}

func TestBuildUpdateRequest_SensitiveInputs(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	t.Run("removing sensitive_inputs_json excludes sensitive keys from request", func(t *testing.T) {
		plan := HogFunctionResourceTFModel{
			InputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			SensitiveInputsJSON: types.StringNull(), // user removed sensitive_inputs_json
		}
		state := HogFunctionResourceTFModel{
			InputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			SensitiveInputsJSON: types.StringValue(`{"api_key":{"value":"secret123"}}`),
		}

		req, diags := ops.BuildUpdateRequest(ctx, plan, state)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		assert.Contains(t, req.Inputs, "url", "url should be in inputs")
		assert.NotContains(t, req.Inputs, "api_key", "api_key should NOT be in inputs after removing sensitive_inputs_json")
	})

	t.Run("updating sensitive_inputs_json merges correctly", func(t *testing.T) {
		plan := HogFunctionResourceTFModel{
			InputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			SensitiveInputsJSON: types.StringValue(`{"api_key":{"value":"new_secret"}}`),
		}
		state := HogFunctionResourceTFModel{
			InputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			SensitiveInputsJSON: types.StringValue(`{"api_key":{"value":"old_secret"}}`),
		}

		req, diags := ops.BuildUpdateRequest(ctx, plan, state)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		assert.Contains(t, req.Inputs, "url")
		assert.Contains(t, req.Inputs, "api_key")
		apiKey := req.Inputs["api_key"].(map[string]interface{})
		assert.Equal(t, "new_secret", apiKey["value"], "should use the new secret value")
	})
}

func TestStripServerFields(t *testing.T) {
	tests := map[string]struct {
		input    interface{}
		expected interface{}
	}{
		"nil input": {
			input:    nil,
			expected: nil,
		},
		"primitive string": {
			input:    "hello",
			expected: "hello",
		},
		"primitive number": {
			input:    42.0,
			expected: 42.0,
		},
		"simple map without server fields": {
			input: map[string]interface{}{
				"name":  "test",
				"value": 123,
			},
			expected: map[string]interface{}{
				"name":  "test",
				"value": 123,
			},
		},
		"map with bytecode field": {
			input: map[string]interface{}{
				"name":     "test",
				"bytecode": []interface{}{"_H", 1, 32},
			},
			expected: map[string]interface{}{
				"name": "test",
			},
		},
		"map with order field": {
			input: map[string]interface{}{
				"name":  "test",
				"order": 5,
			},
			expected: map[string]interface{}{
				"name": "test",
			},
		},
		"map with both bytecode and order": {
			input: map[string]interface{}{
				"name":     "test",
				"bytecode": []interface{}{"_H", 1},
				"order":    3,
				"value":    "keep",
			},
			expected: map[string]interface{}{
				"name":  "test",
				"value": "keep",
			},
		},
		"nested map with server fields": {
			input: map[string]interface{}{
				"url": map[string]interface{}{
					"value":    "https://example.com",
					"bytecode": []interface{}{"_H", 1, 32, "https://example.com"},
					"order":    0,
				},
				"method": map[string]interface{}{
					"value": "POST",
					"order": 1,
				},
			},
			expected: map[string]interface{}{
				"url": map[string]interface{}{
					"value": "https://example.com",
				},
				"method": map[string]interface{}{
					"value": "POST",
				},
			},
		},
		"array of maps with server fields": {
			input: []interface{}{
				map[string]interface{}{
					"id":       "event1",
					"bytecode": []interface{}{"_H"},
				},
				map[string]interface{}{
					"id":    "event2",
					"order": 1,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"id": "event1",
				},
				map[string]interface{}{
					"id": "event2",
				},
			},
		},
		"deeply nested structure (like mappings)": {
			input: map[string]interface{}{
				"mappings": []interface{}{
					map[string]interface{}{
						"name": "mapping1",
						"inputs": map[string]interface{}{
							"url": map[string]interface{}{
								"value":    "https://example.com",
								"bytecode": []interface{}{"_H", 1},
								"order":    0,
							},
						},
						"filters": map[string]interface{}{
							"events":   []interface{}{"$pageview"},
							"bytecode": []interface{}{"_H", 1, 32},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"mappings": []interface{}{
					map[string]interface{}{
						"name": "mapping1",
						"inputs": map[string]interface{}{
							"url": map[string]interface{}{
								"value": "https://example.com",
							},
						},
						"filters": map[string]interface{}{
							"events": []interface{}{"$pageview"},
						},
					},
				},
			},
		},
		"empty map": {
			input:    map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		"empty array": {
			input:    []interface{}{},
			expected: []interface{}{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := stripServerFields(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
