package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// hogFunctionJSONTestCases lists every hog_function attribute that holds a raw
// JSON document, together with a config-style JSON value, a semantically equal
// variant serialised the way the PostHog API returns it (alphabetised keys,
// different whitespace), and a semantically different value.
var hogFunctionJSONTestCases = []struct {
	attribute     string
	configJSON    string
	reorderedJSON string
	differentJSON string
}{
	{
		attribute:     "inputs_json",
		configJSON:    `{"pixelId":{"value":"123","templating":"hog"},"eventName":{"value":"Purchase","templating":"hog"}}`,
		reorderedJSON: `{"eventName": {"templating": "hog", "value": "Purchase"}, "pixelId": {"templating": "hog", "value": "123"}}`,
		differentJSON: `{"pixelId":{"value":"456","templating":"hog"},"eventName":{"value":"Purchase","templating":"hog"}}`,
	},
	{
		attribute:     "sensitive_inputs_json",
		configJSON:    `{"apiKey":{"value":"secret-token","templating":"hog"}}`,
		reorderedJSON: `{"apiKey": {"templating": "hog", "value": "secret-token"}}`,
		differentJSON: `{"apiKey":{"value":"other-token","templating":"hog"}}`,
	},
	{
		attribute:     "filters_json",
		configJSON:    `{"events":[{"id":"$pageview","type":"events"}],"filter_test_accounts":true}`,
		reorderedJSON: `{"filter_test_accounts": true, "events": [{"type": "events", "id": "$pageview"}]}`,
		differentJSON: `{"events":[{"id":"$autocapture","type":"events"}],"filter_test_accounts":true}`,
	},
	{
		attribute:     "inputs_schema_json",
		configJSON:    `[{"key":"apiKey","type":"string","secret":true,"required":true}]`,
		reorderedJSON: `[{"required": true, "secret": true, "key": "apiKey", "type": "string"}]`,
		differentJSON: `[{"key":"apiKey","type":"string","secret":false,"required":true}]`,
	},
	{
		attribute:     "masking_json",
		configJSON:    `{"ttl":3600,"threshold":5,"hash":"{person.id}"}`,
		reorderedJSON: `{"hash": "{person.id}", "threshold": 5, "ttl": 3600}`,
		differentJSON: `{"ttl":7200,"threshold":5,"hash":"{person.id}"}`,
	},
	{
		attribute:     "mappings_json",
		configJSON:    `[{"name":"Purchase","filters":{"events":[{"id":"$pageview","type":"events"}]}}]`,
		reorderedJSON: `[{"filters": {"events": [{"type": "events", "id": "$pageview"}]}, "name": "Purchase"}]`,
		differentJSON: `[{"name":"Checkout","filters":{"events":[{"id":"$pageview","type":"events"}]}}]`,
	},
}

// hogFunctionJSONValue builds the given attribute's value from a raw JSON
// string using the type declared in the schema, and requires that the type
// supports semantic string equality. With plain types.String the framework
// compares config and state byte-for-byte, which is exactly the perpetual-diff
// bug from issue #120.
func hogFunctionJSONValue(t *testing.T, attribute, rawJSON string) basetypes.StringValuableWithSemanticEquals {
	t.Helper()

	attr, ok := HogFunctionOps{}.Schema().Attributes[attribute]
	require.True(t, ok, "attribute %s not found in schema", attribute)

	strTypable, ok := attr.GetType().(basetypes.StringTypable)
	require.True(t, ok, "attribute %s must be string-based", attribute)

	value, diags := strTypable.ValueFromString(context.Background(), basetypes.NewStringValue(rawJSON))
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	semantic, ok := value.(basetypes.StringValuableWithSemanticEquals)
	require.True(t, ok,
		"attribute %s is compared as a raw string; it must use a type with semantic JSON equality (jsontypes.Normalized) or key reordering by the API causes a perpetual diff",
		attribute)

	return semantic
}

// Semantically equal JSON (same document, different key order and whitespace,
// as returned by the PostHog API) must compare as equal, so Terraform reports
// no diff.
func TestHogFunctionJSONAttributes_SemanticallyEqualJSONNoDiff(t *testing.T) {
	for _, tc := range hogFunctionJSONTestCases {
		t.Run(tc.attribute, func(t *testing.T) {
			configValue := hogFunctionJSONValue(t, tc.attribute, tc.configJSON)
			stateValue := hogFunctionJSONValue(t, tc.attribute, tc.reorderedJSON)

			equal, diags := configValue.StringSemanticEquals(context.Background(), stateValue)
			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.True(t, equal,
				"config %s and API-serialised %s are semantically equal and must not diff",
				tc.configJSON, tc.reorderedJSON)
		})
	}
}

// Semantically different JSON must still be detected as a change.
func TestHogFunctionJSONAttributes_SemanticallyDifferentJSONStillDiffs(t *testing.T) {
	for _, tc := range hogFunctionJSONTestCases {
		t.Run(tc.attribute, func(t *testing.T) {
			configValue := hogFunctionJSONValue(t, tc.attribute, tc.configJSON)
			changedValue := hogFunctionJSONValue(t, tc.attribute, tc.differentJSON)

			equal, diags := configValue.StringSemanticEquals(context.Background(), changedValue)
			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.False(t, equal,
				"%s and %s differ semantically and must produce a diff",
				tc.configJSON, tc.differentJSON)
		})
	}
}
