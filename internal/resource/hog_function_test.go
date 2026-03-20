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
	tests := map[string]struct {
		inputsJSON          string
		sensitiveInputsJSON string
		expectedKeys        []string
		expectedURL         string
	}{
		"only inputs_json": {
			inputsJSON:          `{"url":{"value":"https://example.com"},"method":{"value":"POST"}}`,
			sensitiveInputsJSON: "",
			expectedKeys:        []string{"url", "method"},
			expectedURL:         "https://example.com",
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
		},
	}

	ops := HogFunctionOps{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := HogFunctionResourceTFModel{}
			if tt.inputsJSON != "" {
				model.InputsJSON = types.StringValue(tt.inputsJSON)
			} else {
				model.InputsJSON = types.StringNull()
			}
			if tt.sensitiveInputsJSON != "" {
				model.SensitiveInputsJSON = types.StringValue(tt.sensitiveInputsJSON)
			} else {
				model.SensitiveInputsJSON = types.StringNull()
			}

			req, diags := ops.BuildCreateRequest(context.Background(), model)
			assert.False(t, diags.HasError(), "should not have errors")

			for _, key := range tt.expectedKeys {
				assert.Contains(t, req.Inputs, key)
			}

			if tt.expectedURL != "" {
				urlInput := req.Inputs["url"].(map[string]interface{})
				assert.Equal(t, tt.expectedURL, urlInput["value"])
			}

			// Verify sensitive precedence
			if name == "sensitive takes precedence on key conflict" {
				apiKeyInput := req.Inputs["api_key"].(map[string]interface{})
				assert.Equal(t, "secret", apiKeyInput["value"])
			}
		})
	}
}

func TestMapResponseToModel_SplitsInputsBack(t *testing.T) {
	tests := map[string]struct {
		inputsJSON          types.String
		sensitiveInputsJSON types.String
		respInputs          map[string]interface{}
		expectedInputs      string
		expectedSensitive   string
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
			expectedInputs:    `{"url":{"value":"https://example.com"}}`,
			expectedSensitive: "",
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
	}

	ops := HogFunctionOps{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := &HogFunctionResourceTFModel{
				InputsJSON:          tt.inputsJSON,
				SensitiveInputsJSON: tt.sensitiveInputsJSON,
			}

			resp := httpclient.HogFunction{
				ID:     "test-id",
				Inputs: tt.respInputs,
			}

			diags := ops.MapResponseToModel(context.Background(), resp, model)
			assert.False(t, diags.HasError(), "should not have errors")

			assert.Equal(t, tt.expectedInputs, model.InputsJSON.ValueString())

			if tt.expectedSensitive != "" {
				assert.Equal(t, tt.expectedSensitive, model.SensitiveInputsJSON.ValueString())
			} else {
				assert.True(t, model.SensitiveInputsJSON.IsNull())
			}
		})
	}
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
