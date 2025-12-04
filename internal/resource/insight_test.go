package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterToUserFields(t *testing.T) {
	tests := map[string]struct {
		userData interface{}
		apiData  interface{}
		expected interface{}
	}{
		"filters out server-only fields": {
			userData: map[string]interface{}{
				"kind": "DataVisualizationNode",
				"source": map[string]interface{}{
					"kind": "TrendsQuery",
				},
			},
			apiData: map[string]interface{}{
				"kind": "DataVisualizationNode",
				"source": map[string]interface{}{
					"kind":    "TrendsQuery",
					"version": 2, // server-added
				},
				"result":    []interface{}{1, 2, 3}, // server-added
				"hogql":     "SELECT ...",           // server-added
				"is_cached": true,                   // server-added
			},
			expected: map[string]interface{}{
				"kind": "DataVisualizationNode",
				"source": map[string]interface{}{
					"kind": "TrendsQuery",
				},
			},
		},
		"complex objects should be handled correctly as well": {
			userData: map[string]interface{}{
				"kind":   "DataVisualizationNode",
				"result": []interface{}{1, 2, 4},
				"source": map[string]interface{}{
					"kind": "TrendsQuery",
				},
				"test": "foobar",
			},
			apiData: map[string]interface{}{
				"kind": "DataVisualizationNode",
				"source": map[string]interface{}{
					"kind":    "TrendsQuery",
					"version": 2, // server-added
				},
				"test": map[string]interface{}{
					"kind": "TestQuery",
				},
				"result":    []interface{}{1, 2, 3, 4},
				"hogql":     "SELECT ...", // server-added
				"is_cached": true,         // server-added
			},
			expected: map[string]interface{}{
				"kind": "DataVisualizationNode",
				"source": map[string]interface{}{
					"kind": "TrendsQuery",
				},
				"result": []interface{}{1, 2, 3, 4}, // Drift will be flagged
				"test": map[string]interface{}{
					"kind": "TestQuery",
				},
			},
		},
		"detects value changes (drift)": {
			userData: map[string]interface{}{
				"kind": "DataVisualizationNode",
			},
			apiData: map[string]interface{}{
				"kind": "InsightVizNode", // changed!
			},
			expected: map[string]interface{}{
				"kind": "InsightVizNode", // should show API value for drift detection
			},
		},
		"handles nested objects": {
			userData: map[string]interface{}{
				"source": map[string]interface{}{
					"kind": "TrendsQuery",
					"properties": map[string]interface{}{
						"type": "AND",
					},
				},
			},
			apiData: map[string]interface{}{
				"source": map[string]interface{}{
					"kind":    "TrendsQuery",
					"version": 2,
					"properties": map[string]interface{}{
						"type":   "OR",
						"values": []interface{}{}, // server-added
					},
				},
				"result": []interface{}{},
			},
			expected: map[string]interface{}{
				"source": map[string]interface{}{
					"kind": "TrendsQuery",
					"properties": map[string]interface{}{
						"type": "OR",
					},
				},
			},
		},
		"handles arrays with objects": {
			userData: map[string]interface{}{
				"series": []interface{}{
					map[string]interface{}{"event": "$pageview"},
					map[string]interface{}{"event": "$click"},
				},
			},
			apiData: map[string]interface{}{
				"series": []interface{}{
					map[string]interface{}{"event": "$pageview", "id": 1},
					map[string]interface{}{"event": "$click", "id": 2},
				},
				"result": []interface{}{100, 200},
			},
			expected: map[string]interface{}{
				"series": []interface{}{
					map[string]interface{}{"event": "$pageview"},
					map[string]interface{}{"event": "$click"},
				},
			},
		},
		"handles primitive values": {
			userData: map[string]interface{}{
				"limit": float64(100),
			},
			apiData: map[string]interface{}{
				"limit":  float64(50), // changed - should show drift
				"offset": float64(0),  // server-added
			},
			expected: map[string]interface{}{
				"limit": float64(50), // returns API value for drift detection
			},
		},
		"handles empty user data": {
			userData: map[string]interface{}{},
			apiData: map[string]interface{}{
				"result": []interface{}{1, 2, 3},
			},
			expected: map[string]interface{}{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := filterToOnlyIncludeUserFields(tt.userData, tt.apiData)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNormalizeQueryForState(t *testing.T) {
	tests := map[string]struct {
		apiQuery      map[string]interface{}
		userQueryJSON string
		expected      string
	}{
		"filters, sorts keys and normalizes": {
			apiQuery: map[string]interface{}{
				"kind": "DataVisualizationNode",
				"source": map[string]interface{}{
					"kind": "TrendsQuery",
				},
				"result": []interface{}{1, 2, 3},
			},
			userQueryJSON: `{"source": {"kind": "TrendsQuery"}, "kind": "DataVisualizationNode"}`,
			expected:      `{"kind":"DataVisualizationNode","source":{"kind":"TrendsQuery"}}`,
		},
		"sorts keys canonically": {
			apiQuery: map[string]interface{}{
				"z_field": "z",
				"a_field": "a",
			},
			userQueryJSON: `{"z_field": "z", "a_field": "a"}`,
			expected:      `{"a_field":"a","z_field":"z"}`,
		},
		"handles nil query": {
			apiQuery:      nil,
			userQueryJSON: `{}`,
			expected:      "",
		},
		"falls back on invalid user JSON": {
			apiQuery: map[string]interface{}{
				"kind": "test",
			},
			userQueryJSON: `invalid json`,
			expected:      `{"kind":"test"}`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := normalizeQueryForState(tt.apiQuery, tt.userQueryJSON)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
