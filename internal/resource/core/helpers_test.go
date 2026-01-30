package core

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptr[T any](v T) *T {
	return &v
}

func TestPtrToStringNullIfEmptyTrimmed(t *testing.T) {
	tests := map[string]struct {
		input    *string
		wantNull bool
		wantVal  string
	}{
		"nil returns null": {
			input:    nil,
			wantNull: true,
		},
		"empty string returns null": {
			input:    ptr(""),
			wantNull: true,
		},
		"whitespace-only returns null": {
			input:    ptr("   "),
			wantNull: true,
		},
		"tabs and newlines only returns null": {
			input:    ptr("\t\n  \t"),
			wantNull: true,
		},
		"non-empty string returns value": {
			input:   ptr("hello"),
			wantVal: "hello",
		},
		"string with surrounding whitespace returns full value": {
			input:   ptr("  hello  "),
			wantVal: "  hello  ",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PtrToStringNullIfEmptyTrimmed(tc.input)
			if tc.wantNull {
				assert.True(t, got.IsNull(), "expected null")
				return
			}
			require.False(t, got.IsNull(), "expected non-null")
			assert.Equal(t, tc.wantVal, got.ValueString())
		})
	}
}

func TestPtrToBool(t *testing.T) {
	tests := map[string]struct {
		input    *bool
		wantNull bool
		wantVal  bool
	}{
		"nil returns null": {
			input:    nil,
			wantNull: true,
		},
		"true returns true": {
			input:   ptr(true),
			wantVal: true,
		},
		"false returns false": {
			input:   ptr(false),
			wantVal: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PtrToBool(tc.input)
			if tc.wantNull {
				assert.True(t, got.IsNull(), "expected null")
				return
			}
			require.False(t, got.IsNull(), "expected non-null")
			assert.Equal(t, tc.wantVal, got.ValueBool())
		})
	}
}

func TestExtractTags(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		input    types.Set
		wantNil  bool
		wantTags []string
	}{
		"null set returns nil": {
			input:   types.SetNull(types.StringType),
			wantNil: true,
		},
		"unknown set returns nil": {
			input:   types.SetUnknown(types.StringType),
			wantNil: true,
		},
		"empty set returns empty slice": {
			input: func() types.Set {
				s, _ := types.SetValueFrom(ctx, types.StringType, []string{})
				return s
			}(),
			wantTags: []string{},
		},
		"single tag returns slice with one element": {
			input: func() types.Set {
				s, _ := types.SetValueFrom(ctx, types.StringType, []string{"tag1"})
				return s
			}(),
			wantTags: []string{"tag1"},
		},
		"multiple tags returns all tags": {
			input: func() types.Set {
				s, _ := types.SetValueFrom(ctx, types.StringType, []string{"a", "b", "c"})
				return s
			}(),
			wantTags: []string{"a", "b", "c"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, diags := ExtractTags(ctx, tc.input)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)
			if tc.wantNil {
				assert.Nil(t, got, "expected null")
				return
			}
			assert.ElementsMatch(t, tc.wantTags, got)
		})
	}
}

func TestTagsToSet(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		input    []string
		wantNull bool
		wantLen  int
	}{
		"nil returns null set": {
			input:    nil,
			wantNull: true,
		},
		"empty slice returns null set": {
			input:    []string{},
			wantNull: true,
		},
		"single tag returns set with one element": {
			input:   []string{"tag1"},
			wantLen: 1,
		},
		"multiple tags returns set with all elements": {
			input:   []string{"a", "b", "c"},
			wantLen: 3,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, diags := TagsToSet(ctx, tc.input)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)
			if tc.wantNull {
				assert.True(t, got.IsNull(), "expected null set")
				return
			}
			require.False(t, got.IsNull(), "expected non-null set")
			assert.Len(t, got.Elements(), tc.wantLen)
		})
	}
}

func TestTagsToSetPreserveEmpty(t *testing.T) {
	ctx := context.Background()
	nullSet := types.SetNull(types.StringType)
	emptySet, _ := types.SetValueFrom(ctx, types.StringType, []string{})
	nonEmptySet, _ := types.SetValueFrom(ctx, types.StringType, []string{"existing"})

	tests := map[string]struct {
		tags         []string
		currentModel types.Set
		wantNull     bool
		wantEmpty    bool
		wantLen      int
	}{
		"tags with values, null model - returns set with values": {
			tags:         []string{"a", "b"},
			currentModel: nullSet,
			wantLen:      2,
		},
		"tags with values, non-null model - returns set with values": {
			tags:         []string{"a", "b"},
			currentModel: nonEmptySet,
			wantLen:      2,
		},
		"empty tags, null model - returns null": {
			tags:         []string{},
			currentModel: nullSet,
			wantNull:     true,
		},
		"nil tags, null model - returns null": {
			tags:         nil,
			currentModel: nullSet,
			wantNull:     true,
		},
		"empty tags, non-null model - returns empty set (preserves intent)": {
			tags:         []string{},
			currentModel: nonEmptySet,
			wantEmpty:    true,
		},
		"empty tags, empty model - returns empty set (preserves intent)": {
			tags:         []string{},
			currentModel: emptySet,
			wantEmpty:    true,
		},
		"single value returns set with one element": {
			tags:         []string{"test"},
			currentModel: nullSet,
			wantLen:      1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, diags := TagsToSetPreserveEmpty(ctx, tc.tags, tc.currentModel)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)
			if tc.wantNull {
				assert.True(t, got.IsNull(), "expected null")
				return
			}
			require.False(t, got.IsNull(), "expected non-null")
			if tc.wantEmpty {
				assert.Empty(t, got.Elements(), "expected empty set")
				return
			}
			assert.Len(t, got.Elements(), tc.wantLen)
		})
	}
}

func TestInt32SetPreserveEmpty(t *testing.T) {
	ctx := context.Background()
	nullSet := types.SetNull(types.Int32Type)
	emptySet, _ := types.SetValueFrom(ctx, types.Int32Type, []int32{})
	nonEmptySet, _ := types.SetValueFrom(ctx, types.Int32Type, []int32{123})

	tests := map[string]struct {
		values       []int32
		currentModel types.Set
		wantNull     bool
		wantEmpty    bool
		wantLen      int
	}{
		"values present, null model - returns set with values": {
			values:       []int32{1, 2, 3},
			currentModel: nullSet,
			wantLen:      3,
		},
		"values present, non-null model - returns set with values": {
			values:       []int32{1, 2},
			currentModel: nonEmptySet,
			wantLen:      2,
		},
		"empty values, null model - returns null": {
			values:       []int32{},
			currentModel: nullSet,
			wantNull:     true,
		},
		"nil values, null model - returns null": {
			values:       nil,
			currentModel: nullSet,
			wantNull:     true,
		},
		"empty values, non-null model - returns empty set (preserves intent)": {
			values:       []int32{},
			currentModel: nonEmptySet,
			wantEmpty:    true,
		},
		"empty values, empty model - returns empty set (preserves intent)": {
			values:       []int32{},
			currentModel: emptySet,
			wantEmpty:    true,
		},
		"single value returns set with one element": {
			values:       []int32{42},
			currentModel: nullSet,
			wantLen:      1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, diags := Int32SetPreserveEmpty(ctx, tc.values, tc.currentModel)
			require.False(t, diags.HasError(), "unexpected error: %v", diags)
			if tc.wantNull {
				assert.True(t, got.IsNull(), "expected null")
				return
			}
			require.False(t, got.IsNull(), "expected non-null")
			if tc.wantEmpty {
				assert.Empty(t, got.Elements(), "expected empty set")
				return
			}
			assert.Len(t, got.Elements(), tc.wantLen)
		})
	}
}

func TestShouldClearString(t *testing.T) {
	tests := map[string]struct {
		plan  types.String
		state types.String
		want  bool
	}{
		"plan null, state has value - should clear": {
			plan:  types.StringNull(),
			state: types.StringValue("existing"),
			want:  true,
		},
		"plan null, state null - nothing to clear": {
			plan:  types.StringNull(),
			state: types.StringNull(),
			want:  false,
		},
		"plan has value, state has value - no clearing needed": {
			plan:  types.StringValue("new"),
			state: types.StringValue("old"),
			want:  false,
		},
		"plan has value, state null - no clearing needed": {
			plan:  types.StringValue("new"),
			state: types.StringNull(),
			want:  false,
		},
		"plan unknown, state has value - unknown handled elsewhere": {
			plan:  types.StringUnknown(),
			state: types.StringValue("existing"),
			want:  false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ShouldClearString(tc.plan, tc.state)
			assert.Equal(t, tc.want, got)
		})
	}
}
