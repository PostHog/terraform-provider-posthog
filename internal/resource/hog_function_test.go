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

func TestHogFunctionSchema_InputsSchemaJSONExists(t *testing.T) {
	ops := HogFunctionOps{}
	s := ops.Schema()

	attr, ok := s.Attributes["inputs_schema_json"]
	require.True(t, ok, "inputs_schema_json attribute must exist in schema")

	strAttr, ok := attr.(schema.StringAttribute)
	require.True(t, ok, "inputs_schema_json must be a StringAttribute")

	assert.True(t, strAttr.Optional, "inputs_schema_json must be optional")
	assert.False(t, strAttr.Sensitive, "inputs_schema_json should not be sensitive")
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

func TestBuildCreateRequest_InputsSchemaJSON(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	tests := map[string]struct {
		inputsSchemaJSON string
		expectedLen      int
		expectedKey      string
	}{
		"parses inputs_schema_json into request": {
			inputsSchemaJSON: `[{"key":"apiKey","type":"string","secret":true,"required":true}]`,
			expectedLen:      1,
			expectedKey:      "apiKey",
		},
		"multiple schema entries": {
			inputsSchemaJSON: `[{"key":"apiKey","type":"string","secret":true},{"key":"region","type":"string"}]`,
			expectedLen:      2,
			expectedKey:      "apiKey",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := HogFunctionResourceTFModel{
				InputsSchemaJSON: types.StringValue(tc.inputsSchemaJSON),
			}

			req, diags := ops.BuildCreateRequest(ctx, model)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)

			require.Len(t, req.InputsSchema, tc.expectedLen)
			assert.Equal(t, tc.expectedKey, req.InputsSchema[0]["key"])
		})
	}
}

func TestBuildCreateRequest_InputsSchemaJSON_Invalid(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	model := HogFunctionResourceTFModel{
		InputsSchemaJSON: types.StringValue(`not valid json`),
	}

	_, diags := ops.BuildCreateRequest(ctx, model)
	require.True(t, diags.HasError(), "expected error for invalid JSON")
}

func TestBuildCreateRequest_InputsSchemaJSON_Null(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	model := HogFunctionResourceTFModel{
		InputsSchemaJSON: types.StringNull(),
	}

	req, diags := ops.BuildCreateRequest(ctx, model)
	require.False(t, diags.HasError())
	assert.Nil(t, req.InputsSchema)
}

func TestBuildCreateRequest_InputsSchemaWithInputsAndSensitiveInputs(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	t.Run("all three fields populate request correctly", func(t *testing.T) {
		model := HogFunctionResourceTFModel{
			InputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			SensitiveInputsJSON: types.StringValue(`{"apiKey":{"value":"secret-token"}}`),
			InputsSchemaJSON:    types.StringValue(`[{"key":"apiKey","type":"string","secret":true,"required":true}]`),
		}

		req, diags := ops.BuildCreateRequest(ctx, model)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		// Schema is set
		require.Len(t, req.InputsSchema, 1)
		assert.Equal(t, "apiKey", req.InputsSchema[0]["key"])
		assert.Equal(t, true, req.InputsSchema[0]["secret"])

		// Both inputs merged (sensitive takes precedence)
		assert.Contains(t, req.Inputs, "url")
		assert.Contains(t, req.Inputs, "apiKey")
		apiKey := req.Inputs["apiKey"].(map[string]interface{})
		assert.Equal(t, "secret-token", apiKey["value"])
	})

	t.Run("schema with inputs_json only (no sensitive)", func(t *testing.T) {
		model := HogFunctionResourceTFModel{
			InputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"},"region":{"value":"us-east-1"}}`),
			SensitiveInputsJSON: types.StringNull(),
			InputsSchemaJSON:    types.StringValue(`[{"key":"region","type":"string"}]`),
		}

		req, diags := ops.BuildCreateRequest(ctx, model)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		require.Len(t, req.InputsSchema, 1)
		assert.Equal(t, "region", req.InputsSchema[0]["key"])
		assert.Contains(t, req.Inputs, "url")
		assert.Contains(t, req.Inputs, "region")
	})

	t.Run("schema with sensitive_inputs_json only (no inputs_json)", func(t *testing.T) {
		model := HogFunctionResourceTFModel{
			InputsJSON:          types.StringNull(),
			SensitiveInputsJSON: types.StringValue(`{"apiKey":{"value":"secret-token"}}`),
			InputsSchemaJSON:    types.StringValue(`[{"key":"apiKey","type":"string","secret":true}]`),
		}

		req, diags := ops.BuildCreateRequest(ctx, model)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		require.Len(t, req.InputsSchema, 1)
		assert.Equal(t, "apiKey", req.InputsSchema[0]["key"])
		assert.Contains(t, req.Inputs, "apiKey")
	})
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
		"duplicate key in both fields: inputs_json gets API value, sensitive_inputs_json keeps prior state": {
			inputsJSON:          types.StringValue(`{"shared_key":{"value":"from_inputs"}}`),
			sensitiveInputsJSON: types.StringValue(`{"shared_key":{"value":"from_sensitive"}}`),
			respInputs: map[string]interface{}{
				"shared_key": map[string]interface{}{"value": "api_value"},
			},
			// inputs_json uses the API value for drift detection.
			// sensitive_inputs_json preserves prior state because the API never returns
			// real secret values (it returns placeholders), so we must not overwrite state.
			expectedInputs:    `{"shared_key":{"value":"api_value"}}`,
			expectedSensitive: `{"shared_key":{"value":"from_sensitive"}}`,
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

func TestMapResponseToModel_InputsSchemaJSON(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	tests := map[string]struct {
		inputsSchemaJSON types.String
		respInputsSchema []map[string]interface{}
		expectedJSON     string
		expectNull       bool
	}{
		"filters to user-specified keys": {
			inputsSchemaJSON: types.StringValue(`[{"key":"apiKey","type":"string","secret":true}]`),
			respInputsSchema: []map[string]interface{}{
				{"key": "url", "type": "string", "label": "URL"},
				{"key": "apiKey", "type": "string", "secret": true, "label": "API Key"},
				{"key": "method", "type": "string", "label": "Method"},
			},
			expectedJSON: `[{"key":"apiKey","secret":true,"type":"string"}]`,
		},
		"null stays null on import": {
			inputsSchemaJSON: types.StringNull(),
			respInputsSchema: []map[string]interface{}{
				{"key": "url", "type": "string"},
			},
			expectNull: true,
		},
		"empty response with non-null model": {
			inputsSchemaJSON: types.StringValue(`[{"key":"apiKey","type":"string"}]`),
			respInputsSchema: nil,
			expectedJSON:     "[]",
		},
		"multiple user keys filtered from larger response": {
			inputsSchemaJSON: types.StringValue(`[{"key":"apiKey","type":"string"},{"key":"region","type":"string"}]`),
			respInputsSchema: []map[string]interface{}{
				{"key": "url", "type": "string"},
				{"key": "apiKey", "type": "string", "secret": true},
				{"key": "method", "type": "string"},
				{"key": "region", "type": "string", "default": "us-east-1"},
			},
			expectedJSON: `[{"key":"apiKey","type":"string"},{"key":"region","type":"string"}]`,
		},
		"API missing a user-specified key uses correct field templates": {
			inputsSchemaJSON: types.StringValue(`[{"key":"a","type":"string"},{"key":"b","type":"number","label":"B"}]`),
			respInputsSchema: []map[string]interface{}{
				{"key": "b", "type": "number", "label": "B", "order": 1},
			},
			// "a" is missing from API so it's dropped. "b" should be filtered
			// using its own user template (type+label), not "a"'s (type only).
			expectedJSON: `[{"key":"b","label":"B","type":"number"}]`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := &HogFunctionResourceTFModel{
				InputsSchemaJSON: tc.inputsSchemaJSON,
			}

			resp := httpclient.HogFunction{
				ID:           "test-id",
				InputsSchema: tc.respInputsSchema,
			}

			diags := ops.MapResponseToModel(ctx, resp, model)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)

			if tc.expectNull {
				assert.True(t, model.InputsSchemaJSON.IsNull())
			} else {
				assert.Equal(t, tc.expectedJSON, model.InputsSchemaJSON.ValueString())
			}
		})
	}
}

func TestMapResponseToModel_InputsSchemaJSON_InvalidState(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	model := &HogFunctionResourceTFModel{
		InputsSchemaJSON: types.StringValue(`not valid json`),
	}

	resp := httpclient.HogFunction{
		ID: "test-id",
		InputsSchema: []map[string]interface{}{
			{"key": "apiKey", "type": "string"},
		},
	}

	diags := ops.MapResponseToModel(ctx, resp, model)
	require.True(t, diags.HasError(), "expected error for invalid JSON in state")
}

func TestMapResponseToModel_InputsSchemaWithInputsAndSensitiveInputs(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	t.Run("all three fields round-trip correctly", func(t *testing.T) {
		model := &HogFunctionResourceTFModel{
			InputsJSON:          types.StringValue(`{"url":{"value":"https://example.com"}}`),
			SensitiveInputsJSON: types.StringValue(`{"apiKey":{"value":"secret-token"}}`),
			InputsSchemaJSON:    types.StringValue(`[{"key":"apiKey","type":"string","secret":true}]`),
		}

		resp := httpclient.HogFunction{
			ID: "test-id",
			InputsSchema: []map[string]interface{}{
				{"key": "url", "type": "string", "label": "URL"},
				{"key": "method", "type": "string", "label": "Method"},
				{"key": "apiKey", "type": "string", "secret": true, "label": "API Key"},
			},
			Inputs: map[string]interface{}{
				"url":    map[string]interface{}{"value": "https://example.com"},
				"apiKey": map[string]interface{}{"value": "secret-token"},
			},
		}

		diags := ops.MapResponseToModel(ctx, resp, model)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		// Schema filtered to user-specified fields only (label stripped since user didn't specify it)
		assert.Equal(t, `[{"key":"apiKey","secret":true,"type":"string"}]`, model.InputsSchemaJSON.ValueString())

		// inputs_json filtered to user-specified key
		assert.Equal(t, `{"url":{"value":"https://example.com"}}`, model.InputsJSON.ValueString())

		// sensitive_inputs_json filtered to user-specified key
		assert.Equal(t, `{"apiKey":{"value":"secret-token"}}`, model.SensitiveInputsJSON.ValueString())
	})
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

func TestBuildUpdateRequest_InputsSchemaJSON(t *testing.T) {
	ctx := context.Background()
	ops := HogFunctionOps{}

	t.Run("inputs_schema_json flows through update", func(t *testing.T) {
		plan := HogFunctionResourceTFModel{
			InputsSchemaJSON: types.StringValue(`[{"key":"apiKey","type":"string","secret":true}]`),
		}
		state := HogFunctionResourceTFModel{
			InputsSchemaJSON: types.StringValue(`[{"key":"oldKey","type":"string"}]`),
		}

		req, diags := ops.BuildUpdateRequest(ctx, plan, state)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		require.Len(t, req.InputsSchema, 1)
		assert.Equal(t, "apiKey", req.InputsSchema[0]["key"])
	})

	t.Run("removing inputs_schema_json sends nil", func(t *testing.T) {
		plan := HogFunctionResourceTFModel{
			InputsSchemaJSON: types.StringNull(),
		}
		state := HogFunctionResourceTFModel{
			InputsSchemaJSON: types.StringValue(`[{"key":"apiKey","type":"string"}]`),
		}

		req, diags := ops.BuildUpdateRequest(ctx, plan, state)
		require.False(t, diags.HasError(), "unexpected error: %v", diags)

		assert.Nil(t, req.InputsSchema)
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
