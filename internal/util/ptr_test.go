package util

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestBoolDiverged(t *testing.T) {
	tests := map[string]struct {
		configured types.Bool
		server     *bool
		want       bool
	}{
		"null configured never diverges":    {types.BoolNull(), BoolPtr(true), false},
		"unknown configured never diverges": {types.BoolUnknown(), BoolPtr(false), false},
		"configured matches server":         {types.BoolValue(true), BoolPtr(true), false},
		"configured differs from server":    {types.BoolValue(true), BoolPtr(false), true},
		"configured false vs server true":   {types.BoolValue(false), BoolPtr(true), true},
		"nil server counts as divergent":    {types.BoolValue(true), nil, true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, BoolDiverged(tc.configured, tc.server))
		})
	}
}

func TestInt64Diverged(t *testing.T) {
	tests := map[string]struct {
		configured types.Int64
		server     *int64
		want       bool
	}{
		"null configured never diverges":    {types.Int64Null(), Int64Ptr(2), false},
		"unknown configured never diverges": {types.Int64Unknown(), Int64Ptr(0), false},
		"configured matches server":         {types.Int64Value(2), Int64Ptr(2), false},
		"configured differs from server":    {types.Int64Value(2), Int64Ptr(1), true},
		"zero configured vs nonzero server": {types.Int64Value(0), Int64Ptr(1), true},
		"nil server counts as divergent":    {types.Int64Value(2), nil, true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, Int64Diverged(tc.configured, tc.server))
		})
	}
}
