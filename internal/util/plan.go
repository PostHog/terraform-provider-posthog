package util

import "github.com/hashicorp/terraform-plugin-framework/types"

// Resolving the effective value of an Optional+Computed attribute needs both the plan and
// the config, because neither alone is enough.
//
// The plan carries the value whenever it is known. On update Terraform copies the prior
// state into the plan for an Optional+Computed attribute the config leaves out, so an
// omitted attribute keeps its last applied value rather than falling back to the default.
// The plan is unknown only on create, or when the config points at something not resolved
// yet. A null config means the attribute was omitted on create, so the server default is
// what the resource will end up with.
//
// The second return value reports whether the value is knowable at all. Callers use it to
// skip a cross-field check rather than guess.

// ResolveInt64 reports the int64 the server will end up with, and whether it is knowable.
func ResolveInt64(plan, config types.Int64, def int64) (int64, bool) {
	if !plan.IsUnknown() && !plan.IsNull() {
		return plan.ValueInt64(), true
	}
	if config.IsNull() {
		return def, true
	}
	return 0, false
}

// ResolveString reports the string the server will end up with, and whether it is knowable.
func ResolveString(plan, config types.String, def string) (string, bool) {
	if !plan.IsUnknown() && !plan.IsNull() {
		return plan.ValueString(), true
	}
	if config.IsNull() {
		return def, true
	}
	return "", false
}

// ResolveBool reports the bool the server will end up with, and whether it is knowable.
func ResolveBool(plan, config types.Bool, def bool) (bool, bool) {
	if !plan.IsUnknown() && !plan.IsNull() {
		return plan.ValueBool(), true
	}
	if config.IsNull() {
		return def, true
	}
	return false, false
}

// IsEmptySet reports whether a set contributes no values. An unknown set is not empty, its
// elements are just not resolvable yet, and Elements() returns nothing for an unknown set,
// so the unknown check comes first.
func IsEmptySet(v types.Set) bool {
	if v.IsUnknown() {
		return false
	}
	return v.IsNull() || len(v.Elements()) == 0
}
