package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/posthog/terraform-provider/pkg/httpclient"
)

// NotCurrentOrganizationAlias rejects httpclient.CurrentOrganizationAlias as an organization_id.
type NotCurrentOrganizationAlias struct{}

func (v NotCurrentOrganizationAlias) Description(_ context.Context) string {
	return `must not be "@current"`
}

func (v NotCurrentOrganizationAlias) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v NotCurrentOrganizationAlias) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.ValueString() == httpclient.CurrentOrganizationAlias {
		resp.Diagnostics.AddAttributeError(req.Path, "Unsupported organization_id", httpclient.ErrCurrentOrganizationAlias.Error())
	}
}

// OrganizationIDPlanModifier replaces the resource when organization_id changes, and keeps the
// state value when the configuration leaves it unset.
//
// State written before httpclient.CurrentOrganizationAlias was rejected can still hold it. Moving
// such a resource to a real organization id updates it in place instead of replacing it, because
// the alias already named that resource's organization and replacing would delete it. When the
// configuration leaves organization_id unset, the value becomes unknown, so the apply records the
// provider's organization_id in place of the alias.
type OrganizationIDPlanModifier struct{}

func (m OrganizationIDPlanModifier) Description(_ context.Context) string {
	return "Replaces the resource when organization_id changes, except when the prior state holds the retired \"@current\" alias."
}

func (m OrganizationIDPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m OrganizationIDPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	legacyAlias := req.StateValue.ValueString() == httpclient.CurrentOrganizationAlias

	if req.ConfigValue.IsNull() {
		if legacyAlias {
			resp.PlanValue = types.StringUnknown()
		} else if resp.PlanValue.IsUnknown() {
			resp.PlanValue = req.StateValue
		}
		return
	}

	if !legacyAlias && !req.ConfigValue.Equal(req.StateValue) {
		resp.RequiresReplace = true
	}
}
