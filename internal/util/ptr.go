package util

import "github.com/hashicorp/terraform-plugin-framework/types"

// PtrToString safely dereferences a string pointer, returning empty string if nil.
func PtrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// StringPtr returns a pointer to the given string.
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to the given int.
func IntPtr(i int) *int {
	return &i
}

// Int64Ptr returns a pointer to the given int64.
func Int64Ptr(i int64) *int64 {
	return &i
}

// BoolPtr returns a pointer to the given bool.
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtrFromValue converts a Terraform String to *string. Returns nil for null/unknown.
func StringPtrFromValue(value types.String) *string {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	v := value.ValueString()
	return &v
}

// Int64PtrFromValue converts a Terraform Int64 to *int64. Returns nil for null/unknown.
func Int64PtrFromValue(value types.Int64) *int64 {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	v := value.ValueInt64()
	return &v
}

// BoolPtrFromValue converts a Terraform Bool to *bool. Returns nil for null/unknown.
func BoolPtrFromValue(value types.Bool) *bool {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	v := value.ValueBool()
	return &v
}

// ValueStringOrEmpty returns the underlying string value, or "" for null/unknown.
func ValueStringOrEmpty(value types.String) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}
	return value.ValueString()
}

// PtrToInt64 converts a *int64 to types.Int64, returning Null for nil pointers.
func PtrToInt64(v *int64) types.Int64 {
	if v == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*v)
}
