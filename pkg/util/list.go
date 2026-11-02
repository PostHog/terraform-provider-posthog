package util

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StringListToSlice converts a Terraform List of strings into a Go []string.
// Null or unknown lists produce a nil slice; type mismatches surface as
// diagnostics on the returned diag.Diagnostics.
func StringListToSlice(ctx context.Context, list types.List) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	if list.IsNull() || list.IsUnknown() {
		return nil, diags
	}
	var out []string
	diags.Append(list.ElementsAs(ctx, &out, false)...)
	return out, diags
}

// StringListToSlicePtr converts a Terraform List of strings into a *[]string for
// a request body. A null or unknown list returns nil so the field is omitted; a
// non-null list returns a pointer to the (possibly empty) slice, so an explicit
// empty list clears the value server-side rather than being silently dropped.
func StringListToSlicePtr(ctx context.Context, list types.List) (*[]string, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return nil, nil
	}
	out := []string{}
	diags := list.ElementsAs(ctx, &out, false)
	return &out, diags
}

// StringSlicePtrToList converts a *[]string from an API response back into a
// Terraform List: nil -> null list (the server returned no value), otherwise the
// (possibly empty) list of strings.
func StringSlicePtrToList(ctx context.Context, s *[]string) (types.List, diag.Diagnostics) {
	if s == nil {
		return types.ListNull(types.StringType), nil
	}
	return types.ListValueFrom(ctx, types.StringType, *s)
}
