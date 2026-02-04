package resource

import (
	"testing"

	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLevelFromString(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"member level": {
			input:    "member",
			expected: MembershipLevelMember.Value,
		},
		"admin level": {
			input:    "admin",
			expected: MembershipLevelAdmin.Value,
		},
		"owner level": {
			input:    "owner",
			expected: MembershipLevelOwner.Value,
		},
		"unknown defaults to member": {
			input:    "unknown",
			expected: MembershipLevelMember.Value,
		},
		"empty string defaults to member": {
			input:    "",
			expected: MembershipLevelMember.Value,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := levelFromString(tc.input)
			require.NotNil(t, result, "expected non-nil result")
			assert.Equal(t, tc.expected, *result, "level value should match expected")
		})
	}
}

func TestLevelToString(t *testing.T) {
	tests := map[string]struct {
		input    *int
		expected string
	}{
		"member level": {
			input:    util.IntPtr(MembershipLevelMember.Value),
			expected: "member",
		},
		"admin level": {
			input:    util.IntPtr(MembershipLevelAdmin.Value),
			expected: "admin",
		},
		"owner level": {
			input:    util.IntPtr(MembershipLevelOwner.Value),
			expected: "owner",
		},
		"nil defaults to member": {
			input:    nil,
			expected: "member",
		},
		"unknown int defaults to member": {
			input:    util.IntPtr(999),
			expected: "member",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := levelToString(tc.input)
			assert.Equal(t, tc.expected, result, "level name should match expected")
		})
	}
}
