package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
