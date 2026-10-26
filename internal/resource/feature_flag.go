package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	EvaluationRuntime          types.String         `tfsdk:"evaluation_runtime"`
	BucketingIdentifier        types.String         `tfsdk:"bucketing_identifier"`
	Archived                   types.Bool           `tfsdk:"archived"`
	EvaluationContexts         types.Set            `tfsdk:"evaluation_contexts"`
	IsRemoteConfiguration      types.Bool           `tfsdk:"is_remote_configuration"`
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
				MarkdownDescription: "Feature flag name/description (PostHog's UI labels this as 'Description'). The API does not expose a separate dedicated description field for feature flags. Empty strings are preserved. PostHog trims surrounding whitespace; Terraform retains the configured spelling when only that whitespace differs.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"evaluation_runtime": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Where the flag is evaluated: `server`, `client`, or `all`. PostHog defaults new flags to `all`. When omitted, the existing runtime is preserved, including on import; updates leave it unknown until the API response so a saved plan cannot overwrite a newer remote value. This controls remote evaluation; it does not restrict server-side local evaluation.",
				Validators:          []validator.String{stringvalidator.OneOf("server", "client", "all")},
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"bucketing_identifier": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Identifier used to assign rollout buckets and variants: `distinct_id` (the API default) or `device_id`. Device bucketing cannot be combined with `ensure_experience_continuity = true`. When omitted, updates preserve the remote setting.",
				Validators:          []validator.String{stringvalidator.OneOf("distinct_id", "device_id")},
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"archived": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the flag is archived and hidden from the default flag list. Archived flags must have `active = false`. Unarchiving does not enable the flag. When omitted, updates preserve the remote setting.",
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"evaluation_contexts": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Contexts in which the flag evaluates, such as `web` or `checkout`. Use lowercase names without surrounding whitespace (up to 50 names, each 1–255 characters). Matching contexts must also be configured in the SDK. Omit to preserve existing contexts, or set `[]` to clear them where project policy permits. PostHog applies contexts only when evaluation contexts are enabled for your organization; otherwise the apply fails with an error that says so.",
				Validators:          []validator.Set{setvalidator.SizeAtMost(50), setvalidator.ValueStringsAre(stringvalidator.UTF8LengthBetween(1, 255), canonicalFeatureFlagContextValidator{})},
			},
			"is_remote_configuration": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the flag delivers remote configuration. Payloads are configured through `filters.payloads`. On creation, PostHog defaults missing or zero rollout percentages to 100; when supplying filters, provide explicit nonzero group rollout percentages. A top-level `rollout_percentage` must match the first group when both are configured. When omitted, updates preserve the remote setting.",
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
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
				MarkdownDescription: "Feature flag filters as JSON. Compared semantically, so key ordering and whitespace differences from the PostHog API do not produce a diff. Fields present in the API response but absent from this config are kept in state so remote changes surface as drift — except the top-level keys listed in `ignore_filter_fields` and an undeclared `feature_enrollment` marker owned by Early Access Features.",
			},
			"ignore_filter_fields": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				MarkdownDescription: "Top-level keys inside `filters` that Terraform does not track for drift (state mirrors config for them, so changes made outside Terraform don't show as a diff). " +
					"When unset, defaults to the keys other PostHog products wire into a flag — `[\"feature_enrollment\", \"super_groups\", \"holdout_groups\", \"holdout\"]` (Early Access Features and Experiments). " +
					"Set to `[]` to track other filter keys, including Experiment wiring, or provide your own set to replace the default. " +
					"`feature_enrollment` is always ignored when absent from `filters`: PostHog owns this marker while an Early Access Feature is linked and preserves it on flag updates, so Terraform cannot clear it by omission. " +
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
				Computed:            true,
				MarkdownDescription: "Set of tags for the feature flag. When omitted, the flag keeps the tags it has in PostHog; set `[]` to remove them.",
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
				Optional:           true,
				DeprecationMessage: "PostHog no longer creates saved usage dashboards when feature flags are created. Remove this setting from new configurations.",
				MarkdownDescription: "Deprecated create-time option. PostHog no longer creates a saved usage dashboard when a flag is created, " +
					"so `true` is rejected for new flags instead of silently doing nothing. Omit this attribute for new configurations. " +
					"Existing flags with this setting remain readable, and usage charts are available on the flag's Usage tab.",
			},
		},
	}
}

// ModifyResourcePlan plans settings the config leaves out as unknown on updates, and rejects
// options PostHog cannot honor on new flags.
func (o FeatureFlagOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan, config FeatureFlagTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !req.State.Raw.IsNull() {
		// A plan that changes nothing keeps the state's values, so it stays empty.
		if !req.Plan.Raw.Equal(req.State.Raw) {
			planOmittedSettingsUnknown(ctx, config, resp)
		}
		return
	}

	resp.Diagnostics.Append(validateRemoteConfigurationCreate(plan)...)
	if plan.CreateUsageDashboard.ValueBool() {
		resp.Diagnostics.AddError(
			"Unsupported create_usage_dashboard",
			"PostHog no longer creates a saved usage dashboard when a feature flag is created. Remove create_usage_dashboard = true; usage charts are available on the flag's Usage tab.",
		)
	}
}

// planOmittedSettingsUnknown plans the settings the config leaves out as unknown, so an update
// does not send them and PostHog keeps its current values. Planning them from state instead would
// let a saved plan, or one made with -refresh=false, overwrite changes made in PostHog since.
func planOmittedSettingsUnknown(ctx context.Context, config FeatureFlagTFModel, resp *resource.ModifyPlanResponse) {
	for _, setting := range []struct {
		name                string
		configured, unknown attr.Value
	}{
		{"evaluation_runtime", config.EvaluationRuntime, types.StringUnknown()},
		{"bucketing_identifier", config.BucketingIdentifier, types.StringUnknown()},
		{"archived", config.Archived, types.BoolUnknown()},
		{"evaluation_contexts", config.EvaluationContexts, types.SetUnknown(types.StringType)},
		{"is_remote_configuration", config.IsRemoteConfiguration, types.BoolUnknown()},
		{"tags", config.Tags, types.SetUnknown(types.StringType)},
	} {
		if setting.configured.IsNull() {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root(setting.name), setting.unknown)...)
		}
	}
}

func (o FeatureFlagOps) BuildCreateRequest(ctx context.Context, model FeatureFlagTFModel) (httpclient.FeatureFlagRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.FeatureFlagRequest{
		Key: model.Key.ValueString(),
	}
	if model.CreateUsageDashboard.ValueBool() {
		diags.AddError(
			"Unsupported create_usage_dashboard",
			"PostHog no longer creates a saved usage dashboard when a feature flag is created. Remove create_usage_dashboard = true; usage charts are available on the flag's Usage tab.",
		)
		return req, diags
	}
	diags.Append(validateRemoteConfigurationCreate(model)...)
	if diags.HasError() {
		return req, diags
	}

	if !model.Name.IsNull() {
		name := model.Name.ValueString()
		req.Name = &name
	}

	diags.Append(applyFeatureFlagSettings(ctx, model, &req)...)

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
	if core.IsKnown(model.Filters) {
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
	if core.IsKnown(model.RolloutPercentage) {
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

	// Always set deleted to false on create
	deleted := false
	req.Deleted = &deleted

	// Older PostHog versions created a usage dashboard by default. False still
	// suppresses that behavior.
	createUsageDashboard := false
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

	diags.Append(applyFeatureFlagSettings(ctx, plan, &req)...)

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
	if core.IsKnown(plan.Filters) {
		if err := json.Unmarshal([]byte(plan.Filters.ValueString()), &filters); err != nil {
			diags.AddError("Invalid filters JSON", fmt.Sprintf("Could not parse filters: %s", err.Error()))
			return req, diags
		}
	} else if core.IsKnown(plan.RolloutPercentage) {
		// If only rollout_percentage is provided, create default filters structure
		filters = map[string]interface{}{
			"groups": []interface{}{
				map[string]interface{}{},
			},
		}
	}

	// If rollout_percentage is provided, add it to the first group in filters
	// Check both IsNull and IsUnknown since rollout_percentage is Computed
	if core.IsKnown(plan.RolloutPercentage) && filters != nil {
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

	// Always include deleted - the plan modifier ensures it defaults to false
	deleted := plan.Deleted.ValueBool()
	req.Deleted = &deleted

	return req, diags
}

func (o FeatureFlagOps) MapResponseToModel(ctx context.Context, resp httpclient.FeatureFlag, model *FeatureFlagTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.Key = types.StringValue(resp.Key)
	// The API trims names on write. Keep the planned/prior spelling only when
	// that normalization explains the difference; real remote edits remain drift.
	if resp.Name == nil || model.Name.IsNull() || model.Name.IsUnknown() || strings.TrimSpace(model.Name.ValueString()) != *resp.Name {
		model.Name = types.StringPointerValue(resp.Name)
	}
	model.EvaluationRuntime = types.StringPointerValue(resp.EvaluationRuntime)
	model.BucketingIdentifier = types.StringPointerValue(resp.BucketingIdentifier)
	model.Archived = types.BoolPointerValue(resp.Archived)
	model.IsRemoteConfiguration = types.BoolPointerValue(resp.IsRemoteConfiguration)
	// PostHog versions without evaluation contexts leave the field out. Read that as no contexts
	// rather than null, which would not match `evaluation_contexts = []` in a configuration.
	contextNames := resp.EvaluationContexts
	if contextNames == nil {
		contextNames = []string{}
	}
	contexts, d := types.SetValueFrom(ctx, types.StringType, contextNames)
	diags.Append(d...)
	model.EvaluationContexts = contexts
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
	// A flag without tags reads as null, unless the configuration says `tags = []`.
	configuredTags := model.Tags
	if !core.IsKnown(configuredTags) {
		configuredTags = types.SetNull(types.StringType)
	}
	tagsSet, d := core.TagsToSetPreserveEmpty(ctx, resp.Tags, configuredTags)
	diags.Append(d...)
	model.Tags = tagsSet

	// Set deleted status - treat nil as false to avoid perpetual diffs
	deleted := resp.Deleted != nil && *resp.Deleted
	model.Deleted = types.BoolValue(deleted)

	return diags
}

// defaultIgnoredFilterKeys are the top-level filters keys other PostHog products wire into
// a flag rather than the author writing them (feature_enrollment via Early Access Features,
// super_groups being its legacy form; holdout_groups/holdout via Experiments). Tracking them
// would show a perpetual diff on every EAF- or experiment-linked flag.
var defaultIgnoredFilterKeys = []string{"feature_enrollment", "super_groups", "holdout_groups", "holdout"}

// resolveIgnoredFilterKeys returns the default set when unset, else the user's set —
// including an empty set, which tracks other filter keys. An undeclared
// feature_enrollment marker is always omitted separately below.
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
// remote-added ones, this keeps every API field with a value, dropping unconfigured
// empty defaults, ignoredKeys, and PostHog-owned feature_enrollment when undeclared.
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
		// EAFs own this marker and PostHog preserves it on flag updates. Omitting it from
		// filters cannot remove it, even when ignore_filter_fields is explicitly empty.
		if _, configured := stateMap["feature_enrollment"]; !configured {
			delete(result, "feature_enrollment")
		}
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
	"cohort_name":    {}, // API attaches this next to a cohort property's numeric value id
	"cohort_deleted": {}, // API sets this when the cohort is deleted or missing, so the flag can't evaluate the condition
	"cohort_missing": {}, // API sets this when no cohort with that id exists in the project
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
	projectID := model.GetEffectiveProjectID()
	flag, err := client.CreateFeatureFlag(ctx, projectID, req)
	if err != nil {
		return flag, err
	}
	if err := checkEvaluationContextsApplied(req, flag); err != nil {
		// The flag is not saved to state, and it would be serving in every context rather than the
		// configured ones. Delete it: PostHog frees a deleted flag's key, so the next apply can
		// create it again.
		id := strconv.FormatInt(flag.ID, 10)
		if _, deleteErr := client.DeleteFeatureFlag(ctx, projectID, id); deleteErr != nil {
			return flag, fmt.Errorf("%w. PostHog created the flag anyway (ID %d), and it is serving without evaluation contexts. Deleting it failed: %v. Delete it in PostHog, or import it with ID %s/%s, before applying again", err, flag.ID, deleteErr, projectID, id)
		}
		return flag, fmt.Errorf("%w. The flag PostHog created was deleted again, so nothing was left behind", err)
	}
	return flag, nil
}

func (o FeatureFlagOps) Read(ctx context.Context, client httpclient.PosthogClient, model FeatureFlagTFModel) (httpclient.FeatureFlag, httpclient.HTTPStatusCode, error) {
	return client.GetFeatureFlag(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o FeatureFlagOps) Update(ctx context.Context, client httpclient.PosthogClient, model FeatureFlagTFModel, req httpclient.FeatureFlagRequest) (httpclient.FeatureFlag, httpclient.HTTPStatusCode, error) {
	projectID, id := model.GetEffectiveProjectID(), model.GetID()
	if contextsChange(ctx, req, model) {
		// Send the new contexts before the rest. PostHog ignores them unless evaluation
		// contexts are enabled for the organization, and in one request with the rest it would
		// still apply the rest, for example enabling the flag in every context instead of the
		// configured ones. A requested deactivation goes with them, so clearing the contexts of a
		// flag being disabled never leaves it serving everywhere in between.
		first := httpclient.FeatureFlagRequest{Key: model.Key.ValueString(), EvaluationContexts: req.EvaluationContexts}
		if req.Active != nil && !*req.Active {
			first.Active = req.Active
		}
		flag, status, err := client.UpdateFeatureFlag(ctx, projectID, id, first)
		if err != nil {
			return flag, status, err
		}
		if err := checkEvaluationContextsApplied(req, flag); err != nil {
			if first.Active != nil {
				return flag, status, fmt.Errorf("%w. The flag was disabled; nothing else in this update was applied", err)
			}
			return flag, status, fmt.Errorf("%w. Nothing else in this update was applied", err)
		}
		flag, status, err = client.UpdateFeatureFlag(ctx, projectID, id, req)
		if err != nil {
			return flag, status, fmt.Errorf("evaluation_contexts were applied, but the rest of the update failed: %w", err)
		}
		return flag, status, checkEvaluationContextsApplied(req, flag)
	}
	flag, status, err := client.UpdateFeatureFlag(ctx, projectID, id, req)
	if err != nil {
		return flag, status, err
	}
	return flag, status, checkEvaluationContextsApplied(req, flag)
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
