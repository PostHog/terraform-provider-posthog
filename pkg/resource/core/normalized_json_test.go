package core

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizedJSON_StringSemanticEquals(t *testing.T) {
	tests := map[string]struct {
		a, b  string
		equal bool
	}{
		"key order and whitespace": {`{"b": 1, "a": [true, null]}`, `{"a":[true,null],"b":1}`, true},
		"escaped characters":       {`{"q": "a > b & c"}`, `{"q":"a \u003e b \u0026 c"}`, true},
		"whole-number float":       {`{"n": 2.0}`, `{"n":2}`, true},
		"exponent":                 {`{"n": 1e3}`, `{"n":1000}`, true},
		"trailing zeros":           {`{"n": 0.10}`, `{"n":0.1}`, true},
		"negative zero":            {`{"n": -0}`, `{"n":0}`, true},
		"different numbers":        {`{"n": 1}`, `{"n":2}`, false},
		"beyond float64 precision": {`{"n": 9007199254740993}`, `{"n":9007199254740992}`, false},
		"number vs string":         {`{"n": 1}`, `{"n":"1"}`, false},
		"null vs missing key":      {`{"a": null}`, `{}`, false},
		"extra key":                {`{"a": 1}`, `{"a":1,"b":2}`, false},
		"array order matters":      {`[1, 2]`, `[2, 1]`, false},
		"array length":             {`[1]`, `[1, 1]`, false},
		"nested float":             {`{"a": [{"b": 5.00}]}`, `{"a":[{"b":5}]}`, true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			equal, diags := NewNormalizedJSONValue(tt.a).StringSemanticEquals(context.Background(), NewNormalizedJSONValue(tt.b))
			require.False(t, diags.HasError(), diags.Errors())
			assert.Equal(t, tt.equal, equal)
		})
	}
}

func TestNormalizedJSON_StringSemanticEqualsInvalidJSON(t *testing.T) {
	_, diags := NewNormalizedJSONValue(`{"a":1}`).StringSemanticEquals(context.Background(), NewNormalizedJSONValue(`{`))
	assert.True(t, diags.HasError())
}

// The framework builds values through the type, so the round trip must keep
// the custom value type (and with it the semantic equality above).
func TestNormalizedJSONType_ValueFromTerraform(t *testing.T) {
	ctx := context.Background()
	typ := NormalizedJSONType{}

	v, err := typ.ValueFromTerraform(ctx, tftypes.NewValue(tftypes.String, `{"n":2.0}`))
	require.NoError(t, err)
	got, ok := v.(NormalizedJSON)
	require.True(t, ok, "got %T", v)
	assert.Equal(t, `{"n":2.0}`, got.ValueString())
	assert.True(t, typ.Equal(got.Type(ctx)))

	null, err := typ.ValueFromTerraform(ctx, tftypes.NewValue(tftypes.String, nil))
	require.NoError(t, err)
	assert.True(t, null.(NormalizedJSON).IsNull())

}
