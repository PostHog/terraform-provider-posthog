package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DefaultBoolFalse is a plan modifier that sets the value to false if not configured.
// This is useful for soft-delete fields where we want Terraform to restore deleted
// resources by defaulting to deleted=false when not explicitly set.
// Specifically in cases where the API returns a 200 HTTP code for soft deleted resources
type DefaultBoolFalse struct{}

func (m DefaultBoolFalse) Description(_ context.Context) string {
	return "Defaults to false if not configured"
}

func (m DefaultBoolFalse) MarkdownDescription(_ context.Context) string {
	return "Defaults to false if not configured"
}

func (m DefaultBoolFalse) PlanModifyBool(_ context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.ConfigValue.IsNull() {
		resp.PlanValue = types.BoolValue(false)
	}
}
