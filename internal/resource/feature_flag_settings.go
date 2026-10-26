package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// bucketingByDevice is the bucketing_identifier PostHog refuses to combine with new uses of
// ensure_experience_continuity.
const bucketingByDevice = "device_id"

// applyFeatureFlagSettings copies the flag settings the model knows into req. Omitted and unknown
// settings are left out of the request, so PostHog keeps its values for them.
func applyFeatureFlagSettings(ctx context.Context, model FeatureFlagTFModel, req *httpclient.FeatureFlagRequest) diag.Diagnostics {
	var diags diag.Diagnostics
	if known(model.EvaluationRuntime) {
		req.EvaluationRuntime = model.EvaluationRuntime.ValueStringPointer()
	}
	if known(model.BucketingIdentifier) {
		req.BucketingIdentifier = model.BucketingIdentifier.ValueStringPointer()
	}
	if known(model.Archived) {
		req.Archived = model.Archived.ValueBoolPointer()
	}
	if known(model.IsRemoteConfiguration) {
		req.IsRemoteConfiguration = model.IsRemoteConfiguration.ValueBoolPointer()
	}
	if known(model.EvaluationContexts) {
		// An empty set is sent as [] so it clears the flag's contexts.
		contexts := []string{}
		diags.Append(model.EvaluationContexts.ElementsAs(ctx, &contexts, false)...)
		req.EvaluationContexts = &contexts
	}
	return diags
}

func known(v attr.Value) bool {
	return !v.IsNull() && !v.IsUnknown()
}

// checkEvaluationContextsApplied reports when PostHog accepted a create or update but kept the
// flag's previous evaluation contexts. PostHog skips them without an error unless evaluation
// contexts are enabled for the organization, which otherwise surfaces only as Terraform's
// generic "Provider produced inconsistent result after apply".
func checkEvaluationContextsApplied(req httpclient.FeatureFlagRequest, flag httpclient.FeatureFlag) error {
	if req.EvaluationContexts == nil {
		return nil
	}
	requested := slices.Sorted(slices.Values(*req.EvaluationContexts))
	applied := slices.Sorted(slices.Values(flag.EvaluationContexts))
	if slices.Equal(requested, applied) {
		return nil
	}
	return fmt.Errorf("PostHog did not apply evaluation_contexts (requested %q, the flag has %q); PostHog ignores evaluation contexts unless the feature is enabled for your organization", requested, applied)
}

// Validate before sending values the API would normalize and change in state.
type canonicalFeatureFlagContextValidator struct{}

func (canonicalFeatureFlagContextValidator) Description(context.Context) string {
	return "Context names must be lowercase without surrounding whitespace."
}

func (v canonicalFeatureFlagContextValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v canonicalFeatureFlagContextValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	if value != strings.ToLower(strings.TrimSpace(value)) {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid evaluation context", v.Description(ctx))
	}
}

func validateRemoteConfigurationCreate(plan FeatureFlagTFModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if !plan.IsRemoteConfiguration.ValueBool() {
		return diags
	}
	if !plan.RolloutPercentage.IsNull() && !plan.RolloutPercentage.IsUnknown() && plan.RolloutPercentage.ValueInt64() == 0 {
		diags.AddAttributeError(path.Root("rollout_percentage"), "Invalid remote configuration rollout", "PostHog changes a zero rollout to 100 when creating remote configuration. Omit rollout_percentage or set a nonzero percentage.")
	}
	if plan.Filters.IsNull() || plan.Filters.IsUnknown() {
		return diags
	}
	var filters map[string]interface{}
	if json.Unmarshal([]byte(plan.Filters.ValueString()), &filters) != nil {
		return diags // The JSON type reports malformed JSON.
	}
	groups, _ := filters["groups"].([]interface{})
	valid := len(groups) > 0
	for index, group := range groups {
		condition, _ := group.(map[string]interface{})
		percentage, ok := condition["rollout_percentage"].(float64)
		// Both configured values must survive the API response unchanged.
		if index == 0 && !plan.RolloutPercentage.IsNull() && !plan.RolloutPercentage.IsUnknown() && percentage != float64(plan.RolloutPercentage.ValueInt64()) {
			valid = false
		}
		if !ok || percentage == 0 {
			valid = false
		}
	}
	if !valid {
		diags.AddAttributeError(path.Root("filters"), "Invalid remote configuration filters", "On creation, omit filters to use PostHog's default, or include groups with explicit nonzero rollout_percentage values (typically 100). If top-level rollout_percentage is also configured, it must match the first group.")
	}
	return diags
}
