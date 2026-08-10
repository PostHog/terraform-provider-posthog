package resource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewFeatureFlag() resource.Resource {
	return core.NewGenericResource[FeatureFlagTFModel, httpclient.FeatureFlagRequest, httpclient.FeatureFlag](
		FeatureFlagOps{},
		core.ProjectScopedImportParser[FeatureFlagTFModel](),
	)
}

type FeatureFlagTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	Key                        types.String         `tfsdk:"key"`
	Name                       types.String         `tfsdk:"name"`
	Active                     types.Bool           `tfsdk:"active"`
	Filters                    jsontypes.Normalized `tfsdk:"filters"`
	IgnoreFilterFields         types.Set            `tfsdk:"ignore_filter_fields"`
	RolloutPercentage          types.Int64          `tfsdk:"rollout_percentage"`
	Tags                       types.Set            `tfsdk:"tags"`
	Deleted                    types.Bool           `tfsdk:"deleted"`
	EnsureExperienceContinuity types.Bool           `tfsdk:"ensure_experience_continuity"`
	CreateUsageDashboard       types.Bool           `tfsdk:"create_usage_dashboard"`
}

type FeatureFlagOps struct{}

func (o FeatureFlagOps) ResourceName() string {
	return "Feature Flag"
}

func (o FeatureFlagOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "PostHog Feature Flag resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Feature Flag ID",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"key": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Feature flag key (unique identifier)",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Feature flag name/description (PostHog's UI labels this as 'Description'). The API does not expose a separate dedicated description field for feature flags.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"active": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the feature flag is active",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"ensure_experience_continuity": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to persist the flag across authentication steps (PostHog's UI labels this as 'Persist flag across authentication steps'). Flags with experience continuity enabled cannot be evaluated by server-side local evaluation; set this to false for flags that must be evaluated locally.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"filters": schema.StringAttribute{
				CustomType:          jsontypes.NormalizedType{},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Feature flag filters as JSON. Compared semantically, so key ordering and whitespace differences from the PostHog API do not produce a diff. Fields present in the API response but absent from this config are kept in state so remote changes surface as drift — except the top-level keys listed in `ignore_filter_fields`.",
			},
			"ignore_filter_fields": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				MarkdownDescription: "Top-level keys inside `filters` that Terraform does not track for drift (state mirrors config for them, so changes made outside Terraform don't show as a diff). " +
					"When unset, defaults to the keys other PostHog products wire into a flag — `[\"super_groups\", \"holdout_groups\", \"holdout\"]` (Early Access Features and Experiments). " +
					"Set to `[]` to track the entire filters blob — including any Early Access Feature or Experiment wiring, which will then show as drift if not declared in `filters` — or provide your own set to replace the default. " +
					"A key you also declare inside `filters` is always tracked (explicit config wins over the ignore list).",
			},
			"rollout_percentage": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Rollout percentage (0-100)",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Set of tags for the feature flag",
			},
			"deleted": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the feature flag is soft-deleted. Terraform will restore soft-deleted flags on apply.",
				PlanModifiers: []planmodifier.Bool{
					core.DefaultBoolFalse{},
				},
			},
			"create_usage_dashboard": schema.BoolAttribute{
				Optional: true,
				MarkdownDescription: "Whether PostHog should auto-create a \"Generated Dashboard: <key> Usage\" dashboard when the flag is created. " +
					"Defaults to `false`: PostHog's API defaults this to `true`, which silently creates one usage dashboard per flag — " +
					"for Terraform-managed flag fleets that quickly clutters the dashboard list. Set to `true` to keep the API's " +
					"auto-generation. Create-time only; changing it later has no effect and a usage dashboard can always be generated " +
					"on demand from the flag's Usage tab.",
			},
		},
	}
}
func (o FeatureFlagOps) BuildCreateRequest(ctx context.Context, model FeatureFlagTFModel) (httpclient.FeatureFlagRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.FeatureFlagRequest{
		Key: model.Key.ValueString(),
	}

	if !model.Name.IsNull() {
		name := model.Name.ValueString()
		req.Name = &name
	}

	if !model.Active.IsNull() {
		active := model.Active.ValueBool()
		req.Active = &active
	}

	if !model.EnsureExperienceContinuity.IsNull() {
		ensureExperienceContinuity := model.EnsureExperienceContinuity.ValueBool()
		req.EnsureExperienceContinuity = &ensureExperienceContinuity
	}

	// Handle filters and rollout_percentage
	var filters map[string]interface{}

	// Check both IsNull and IsUnknown since filters is Computed
	if !model.Filters.IsNull() && !model.Filters.IsUnknown() {
		if err := json.Unmarshal([]byte(model.Filters.ValueString()), &filters); err != nil {
			diags.AddError("Invalid filters JSON", fmt.Sprintf("Could not parse filters: %s", err.Error()))
			return req, diags
		}
	} else {
		// Create default filters structure
		filters = map[string]interface{}{
			"groups": []interface{}{
				map[string]interface{}{},
			},
		}
	}

	// If rollout_percentage is provided, add it to the first group in filters
	// rollout_percentage is a convenience field that maps to filters.groups[0].rollout_percentage
	// Check both IsNull and IsUnknown since rollout_percentage is Computed
	if !model.RolloutPercentage.IsNull() && !model.RolloutPercentage.IsUnknown() {
		percentage := int32(model.RolloutPercentage.ValueInt64())
		groups, ok := filters["groups"].([]interface{})
		if !ok || len(groups) == 0 {
			groups = []interface{}{map[string]interface{}{}}
		}
		firstGroup, ok := groups[0].(map[string]interface{})
		if !ok {
			firstGroup = map[string]interface{}{}
			groups[0] = firstGroup
		}
		firstGroup["rollout_percentage"] = percentage
		filters["groups"] = groups
	}

	req.Filters = filters

	if !model.Tags.IsNull() {
		tags, d := core.ExtractTags(ctx, model.Tags)
		diags.Append(d...)
		req.Tags = tags
	}

	// Always set deleted to false on create
	deleted := false
	req.Deleted = &deleted

	// Suppress PostHog's auto-generated per-flag usage dashboard unless
	// explicitly requested (the API's own default is true). Create-time only.
	createUsageDashboard := model.CreateUsageDashboard.ValueBool()
	req.ShouldCreateUsageDashboard = &createUsageDashboard

	return req, diags
}

func (o FeatureFlagOps) BuildUpdateRequest(ctx context.Context, plan, state FeatureFlagTFModel) (httpclient.FeatureFlagRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.FeatureFlagRequest{
		Key: plan.Key.ValueString(),
	}

	if !plan.Name.IsNull() {
		name := plan.Name.ValueString()
		req.Name = &name
	}

	if !plan.Active.IsNull() {
		active := plan.Active.ValueBool()
		req.Active = &active
	}

	if !plan.EnsureExperienceContinuity.IsNull() {
		ensureExperienceContinuity := plan.EnsureExperienceContinuity.ValueBool()
		req.EnsureExperienceContinuity = &ensureExperienceContinuity
	}

	// Handle filters and rollout_percentage
	var filters map[string]interface{}

	// Check both IsNull and IsUnknown since filters is Computed
	if !plan.Filters.IsNull() && !plan.Filters.IsUnknown() {
		if err := json.Unmarshal([]byte(plan.Filters.ValueString()), &filters); err != nil {
			diags.AddError("Invalid filters JSON", fmt.Sprintf("Could not parse filters: %s", err.Error()))
			return req, diags
		}
	} else if !plan.RolloutPercentage.IsNull() && !plan.RolloutPercentage.IsUnknown() {
		// If only rollout_percentage is provided, create default filters structure
		filters = map[string]interface{}{
			"groups": []interface{}{
				map[string]interface{}{},
			},
		}
	}

	// If rollout_percentage is provided, add it to the first group in filters
	// Check both IsNull and IsUnknown since rollout_percentage is Computed
	if !plan.RolloutPercentage.IsNull() && !plan.RolloutPercentage.IsUnknown() && filters != nil {
		percentage := int32(plan.RolloutPercentage.ValueInt64())
		groups, ok := filters["groups"].([]interface{})
		if !ok || len(groups) == 0 {
			groups = []interface{}{map[string]interface{}{}}
		}
		firstGroup, ok := groups[0].(map[string]interface{})
		if !ok {
			firstGroup = map[string]interface{}{}
			groups[0] = firstGroup
		}
		firstGroup["rollout_percentage"] = percentage
		filters["groups"] = groups
	}

	if filters != nil {
		req.Filters = filters
	}

	if !plan.Tags.IsNull() {
		tags, d := core.ExtractTags(ctx, plan.Tags)
		diags.Append(d...)
		req.Tags = tags
	}

	// Always include deleted - the plan modifier ensures it defaults to false
	deleted := plan.Deleted.ValueBool()
	req.Deleted = &deleted

	return req, diags
}

func (o FeatureFlagOps) MapResponseToModel(ctx context.Context, resp httpclient.FeatureFlag, model *FeatureFlagTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.Key = types.StringValue(resp.Key)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.Active = core.PtrToBool(resp.Active)
	model.EnsureExperienceContinuity = core.PtrToBool(resp.EnsureExperienceContinuity)

	// Set filters if present
	if len(resp.Filters) > 0 {
		ignoredKeys, d := resolveIgnoredFilterKeys(ctx, model.IgnoreFilterFields)
		diags.Append(d...)
		if diags.HasError() {
			return diags
		}
		normalizedFilters, err := normalizeFeatureFlagFiltersForState(resp.Filters, model.Filters.ValueString(), ignoredKeys)
		if err != nil {
			diags.AddError("Failed to normalize filters", err.Error())
			return diags
		}
		model.Filters = jsontypes.NewNormalizedValue(normalizedFilters)
	} else {
		model.Filters = jsontypes.NewNormalizedNull()
	}

	// ignore_filter_fields is config-only (never from the API). On paths that start from an
	// empty model (import), the zero-value Set has no element type; give it a typed null so
	// it converts cleanly to the schema's Set[String].
	if model.IgnoreFilterFields.IsNull() {
		model.IgnoreFilterFields = types.SetNull(types.StringType)
	}

	model.RolloutPercentage = extractRolloutPercentage(resp)

	// Set tags
	tagsSet, d := core.TagsToSet(ctx, resp.Tags)
	diags.Append(d...)
	model.Tags = tagsSet

	// Set deleted status - treat nil as false to avoid perpetual diffs
	deleted := resp.Deleted != nil && *resp.Deleted
	model.Deleted = types.BoolValue(deleted)

	return diags
}

// defaultIgnoredFilterKeys are the top-level filters keys other PostHog products wire into
// a flag rather than the author writing them (super_groups via Early Access Features;
// holdout_groups/holdout via Experiments). Tracking them would show a perpetual diff on
// every EAF- or experiment-linked flag.
var defaultIgnoredFilterKeys = []string{"super_groups", "holdout_groups", "holdout"}

// resolveIgnoredFilterKeys returns the default set when unset, else the user's set —
// including an empty set, which tracks the entire filters blob.
func resolveIgnoredFilterKeys(ctx context.Context, set types.Set) ([]string, diag.Diagnostics) {
	if set.IsNull() || set.IsUnknown() {
		return defaultIgnoredFilterKeys, nil
	}
	var keys []string
	diags := set.ElementsAs(ctx, &keys, false)
	return keys, diags
}

// normalizeFeatureFlagFiltersForState shapes the API's filters for state so remote changes
// to a flag's targeting surface as drift. Unlike the shared normalizeJSONForState whitelist
// (survey/action/hog_function/insight), which drops unconfigured fields and so hides
// remote-added ones, this keeps every API field with a value, dropping only unconfigured
// empty defaults and ignoredKeys the user hasn't configured (see defaultIgnoredFilterKeys).
// Feature-flag-specific by design.
func normalizeFeatureFlagFiltersForState(apiData map[string]interface{}, stateJSON string, ignoredKeys []string) (string, error) {
	var stateData interface{}
	if err := json.Unmarshal([]byte(stateJSON), &stateData); err != nil {
		stateData = nil // unparseable/empty prior state → nothing configured
	}
	stateMap, _ := stateData.(map[string]interface{})

	normalized := normalizeFeatureFlagFilterValue(stateData, apiData)

	// Top level only (never recursively) so a like-named nested key is safe; a key the
	// user declared in filters is kept.
	if result, ok := normalized.(map[string]interface{}); ok {
		for _, key := range ignoredKeys {
			if _, configured := stateMap[key]; !configured {
				delete(result, key)
			}
		}
	}

	return marshalJSON(normalized)
}

// serverComputedFilterKeys are read-only fields the API injects into filter property objects
// at any depth (never authored by the user). Unlike defaultIgnoredFilterKeys (top-level keys),
// these must be stripped recursively. Add future server-only enrichments here.
var serverComputedFilterKeys = map[string]struct{}{
	"cohort_name": {}, // API attaches this next to a cohort property's numeric value id
}

// normalizeFeatureFlagFilterValue walks the API tree keeping every value, dropping unconfigured
// empty defaults and server-computed keys. It mirrors insight.go's filterToOnlyIncludeUserFields
// but inverts the intent (keep-all vs whitelist) and drives off the API; kept separate deliberately.
func normalizeFeatureFlagFilterValue(stateData, apiData interface{}) interface{} {
	switch apiValue := apiData.(type) {
	case map[string]interface{}:
		stateMap, _ := stateData.(map[string]interface{})
		result := make(map[string]interface{})
		for key, apiFieldValue := range apiValue {
			stateFieldValue, configured := stateMap[key]
			// Drop unconfigured server enrichment keys (e.g. cohort_name) even when non-empty.
			if _, isServerKey := serverComputedFilterKeys[key]; !configured && isServerKey {
				continue
			}
			normalized := normalizeFeatureFlagFilterValue(stateFieldValue, apiFieldValue)
			if !configured && isEmptyFeatureFlagFilterValue(normalized) {
				continue
			}
			result[key] = normalized
		}
		return result

	case []interface{}:
		stateSlice, _ := stateData.([]interface{})
		result := make([]interface{}, len(apiValue))
		for i, apiItem := range apiValue {
			var stateItem interface{}
			if i < len(stateSlice) {
				stateItem = stateSlice[i]
			}
			result[i] = normalizeFeatureFlagFilterValue(stateItem, apiItem)
		}
		return result

	default:
		return apiData
	}
}

func isEmptyFeatureFlagFilterValue(value interface{}) bool {
	switch typed := value.(type) {
	case nil:
		return true
	case map[string]interface{}:
		return len(typed) == 0
	case []interface{}:
		return len(typed) == 0
	default:
		return false
	}
}

func (o FeatureFlagOps) Create(ctx context.Context, client httpclient.PosthogClient, model FeatureFlagTFModel, req httpclient.FeatureFlagRequest) (httpclient.FeatureFlag, error) {
	return client.CreateFeatureFlag(ctx, model.GetEffectiveProjectID(), req)
}

func (o FeatureFlagOps) Read(ctx context.Context, client httpclient.PosthogClient, model FeatureFlagTFModel) (httpclient.FeatureFlag, httpclient.HTTPStatusCode, error) {
	return client.GetFeatureFlag(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o FeatureFlagOps) Update(ctx context.Context, client httpclient.PosthogClient, model FeatureFlagTFModel, req httpclient.FeatureFlagRequest) (httpclient.FeatureFlag, httpclient.HTTPStatusCode, error) {
	return client.UpdateFeatureFlag(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o FeatureFlagOps) Delete(ctx context.Context, client httpclient.PosthogClient, model FeatureFlagTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteFeatureFlag(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func extractRolloutPercentage(resp httpclient.FeatureFlag) types.Int64 {
	// Try top-level field first
	if resp.RolloutPercentage != nil {
		return types.Int64Value(int64(*resp.RolloutPercentage))
	}

	// Fall back to extracting from filters.groups[0].rollout_percentage
	if len(resp.Filters) > 0 {
		groups, ok := resp.Filters["groups"].([]interface{})
		if !ok || len(groups) == 0 {
			return types.Int64Null()
		}

		firstGroup, ok := groups[0].(map[string]interface{})
		if !ok {
			return types.Int64Null()
		}

		if rp, ok := firstGroup["rollout_percentage"].(float64); ok {
			return types.Int64Value(int64(rp))
		}
	}

	return types.Int64Null()
}
