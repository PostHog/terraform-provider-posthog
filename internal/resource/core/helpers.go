// Package core provides helper functions for Terraform resources.
package core

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

func PtrToStringNullIfEmptyTrimmed(v *string) types.String {
	if v == nil || strings.TrimSpace(*v) == "" {
		return types.StringNull()
	}
	return types.StringValue(*v)
}

// PtrToStringTrimmed trims trailing whitespace from a string pointer.
// Returns null if the pointer is nil or the trimmed result is empty.
// Useful for multi-line code fields where trailing newlines can differ.
func PtrToStringTrimmed(v *string) types.String {
	if v == nil {
		return types.StringNull()
	}
	trimmed := strings.TrimRight(*v, " \t\n\r")
	if trimmed == "" {
		return types.StringNull()
	}
	return types.StringValue(trimmed)
}

func PtrToBool(v *bool) types.Bool {
	if v == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*v)
}

// ExtractTags extracts []string from types.Set. Returns nil if null/unknown.
func ExtractTags(ctx context.Context, tags types.Set) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	if tags.IsNull() || tags.IsUnknown() {
		return nil, diags
	}
	var result []string
	diags.Append(tags.ElementsAs(ctx, &result, false)...)
	return result, diags
}

// TagsToSet converts []string to types.Set. Returns null set if empty.
func TagsToSet(ctx context.Context, tags []string) (types.Set, diag.Diagnostics) {
	var diags diag.Diagnostics
	if len(tags) == 0 {
		return types.SetNull(types.StringType), diags
	}
	result, d := types.SetValueFrom(ctx, types.StringType, tags)
	diags.Append(d...)
	return result, diags
}

// TagsToSetPreserveEmpty converts []string to types.Set.
// If tags is empty but currentModel is not null, returns an empty set (preserving the user's intent).
// If tags is empty and currentModel is null, returns null set.
func TagsToSetPreserveEmpty(ctx context.Context, tags []string, currentModel types.Set) (types.Set, diag.Diagnostics) {
	var diags diag.Diagnostics
	if len(tags) > 0 {
		result, d := types.SetValueFrom(ctx, types.StringType, tags)
		diags.Append(d...)
		return result, diags
	}
	if !currentModel.IsNull() {
		result, d := types.SetValueFrom(ctx, types.StringType, []string{})
		diags.Append(d...)
		return result, diags
	}
	return types.SetNull(types.StringType), diags
}

// Int32SetPreserveEmpty converts []int32 to types.Set.
// If values is empty but currentModel is not null, returns an empty set (preserving the user's intent).
// If values is empty and currentModel is null, returns null set.
func Int32SetPreserveEmpty(ctx context.Context, values []int32, currentModel types.Set) (types.Set, diag.Diagnostics) {
	var diags diag.Diagnostics
	if len(values) > 0 {
		result, d := types.SetValueFrom(ctx, types.Int32Type, values)
		diags.Append(d...)
		return result, diags
	}
	if !currentModel.IsNull() {
		result, d := types.SetValueFrom(ctx, types.Int32Type, []int32{})
		diags.Append(d...)
		return result, diags
	}
	return types.SetNull(types.Int32Type), diags
}

// ShouldClearString returns true if the plan value is null but state has a value,
// indicating the user wants to clear the field by sending an empty string.
func ShouldClearString(plan, state types.String) bool {
	return plan.IsNull() && !state.IsNull()
}

// StringPtr returns a pointer to the given string.
func StringPtr(s string) *string {
	return &s
}

func ProjectIDSchemaAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Project ID (environment) for this resource. Overrides the provider-level project_id.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
	}
}

func OrganizationIDSchemaAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The identifier of the organization this resource belongs to.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	}
}

// SoftDeletable is implemented by API response types that support soft deletion.
// Some PostHog resources use soft deletes where the API returns HTTP 200 with Deleted=true
// instead of returning HTTP 404 when a resource is deleted.
// This interface enables generic drift detection for such resources.
//
// When to use: Only implement this interface for resources where the PostHog API returns
// HTTP 200 with a "deleted": true field instead of HTTP 404 when fetching a deleted resource.
type SoftDeletable interface {
	IsSoftDeleted() bool
}

func CheckSoftDeleted[T SoftDeletable](response T, originalStatusCode httpclient.HTTPStatusCode) (T, httpclient.HTTPStatusCode, error) {
	if response.IsSoftDeleted() {
		// Use reflection to get the type name for the error message
		typeName := reflect.TypeOf(response).Name()
		// Return a 404 so that the generic resource framework removes it from the TF state
		return response, http.StatusNotFound, fmt.Errorf("%s was deleted externally", typeName)
	}
	return response, originalStatusCode, nil
}
