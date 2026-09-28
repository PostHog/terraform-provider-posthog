package util

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseJSONStringMap(t *testing.T) {
	tests := map[string]struct {
		input        types.String
		expectError  bool
		expectKeys   []string
		expectEmpty  bool
		expectNilMap bool
	}{
		"valid json": {
			input:      types.StringValue(`{"host":"localhost","port":5432}`),
			expectKeys: []string{"host", "port"},
		},
		"empty string": {
			input:       types.StringValue(""),
			expectEmpty: true,
		},
		"null": {
			input:        types.StringNull(),
			expectNilMap: true,
		},
		"unknown": {
			input:        types.StringUnknown(),
			expectNilMap: true,
		},
		"invalid": {
			input:       types.StringValue(`not-json`),
			expectError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, diags := ParseJSONStringMap("test_field", tc.input)
			if tc.expectError {
				require.True(t, diags.HasError())
				assert.Contains(t, diags.Errors()[0].Summary(), "Invalid test_field")
				return
			}
			require.False(t, diags.HasError())
			if tc.expectNilMap {
				assert.Nil(t, got)
				return
			}
			if tc.expectEmpty {
				assert.Empty(t, got)
				return
			}
			for _, k := range tc.expectKeys {
				assert.Contains(t, got, k)
			}
		})
	}
}

func TestStripFieldsWalksTypedMapSlice(t *testing.T) {
	denylist := map[string]struct{}{
		"bytecode": {},
		"order":    {},
	}
	input := []map[string]interface{}{
		{
			"name": "Page view",
			"filters": map[string]interface{}{
				"events": []interface{}{
					map[string]interface{}{"id": "$pageview", "type": "events"},
				},
				"bytecode": []interface{}{"_H", 1},
			},
			"inputs": map[string]interface{}{
				"eventName": map[string]interface{}{
					"value":    "PageView",
					"bytecode": []interface{}{"_H"},
					"order":    0,
				},
			},
		},
	}

	cleaned := StripFields(input, denylist)
	encoded, err := json.Marshal(cleaned)
	require.NoError(t, err)
	assert.NotContains(t, string(encoded), "bytecode")
	assert.NotContains(t, string(encoded), `"order"`)
	assert.Contains(t, string(encoded), "PageView")
	assert.Contains(t, string(encoded), "$pageview")
}
