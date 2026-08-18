package util

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

// The three Resolve helpers share one rule: the plan wins when known, a null config means
// the attribute was omitted so the server default applies, and anything else is not yet
// knowable. Int64 stands in for all three since the logic is identical.
func TestResolveInt64(t *testing.T) {
	tests := map[string]struct {
		plan, config types.Int64
		want         int64
		known        bool
	}{
		"plan known":                {plan: types.Int64Value(7), config: types.Int64Value(7), want: 7, known: true},
		"omitted on create":         {plan: types.Int64Unknown(), config: types.Int64Null(), want: 100, known: true},
		"carried forward on update": {plan: types.Int64Value(10), config: types.Int64Null(), want: 10, known: true},
		"config not yet resolvable": {plan: types.Int64Unknown(), config: types.Int64Unknown()},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, known := ResolveInt64(test.plan, test.config, 100)
			assert.Equal(t, test.known, known)
			if test.known {
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestResolveStringAndBool(t *testing.T) {
	got, known := ResolveString(types.StringUnknown(), types.StringNull(), "above")
	assert.True(t, known)
	assert.Equal(t, "above", got)

	enabled, known := ResolveBool(types.BoolValue(false), types.BoolValue(false), true)
	assert.True(t, known)
	assert.False(t, enabled)
}

func TestIsEmptySet(t *testing.T) {
	empty, diags := types.SetValueFrom(t.Context(), types.StringType, []string{})
	assert.False(t, diags.HasError())
	populated, diags := types.SetValueFrom(t.Context(), types.StringType, []string{"error"})
	assert.False(t, diags.HasError())

	assert.True(t, IsEmptySet(types.SetNull(types.StringType)))
	assert.True(t, IsEmptySet(empty))
	assert.False(t, IsEmptySet(populated))
	// Unknown is not empty: the elements are simply not resolvable yet.
	assert.False(t, IsEmptySet(types.SetUnknown(types.StringType)))
}
