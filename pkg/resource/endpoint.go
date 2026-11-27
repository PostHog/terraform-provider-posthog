package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
)

// EndpointDataFreshnessSeconds are the data freshness values PostHog accepts.
var EndpointDataFreshnessSeconds = []int64{900, 1800, 3600, 21600, 43200, 86400, 604800}

var endpointNameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,127}$`)

const (
	defaultEndpointIsActive       = true
	defaultEndpointIsMaterialized = false
)

func NewEndpoint() resource.Resource {
	return core.NewGenericResource[EndpointResourceTFModel, httpclient.EndpointRequest, httpclient.Endpoint](
		EndpointOps{},
		core.ProjectScopedImportParser[EndpointResourceTFModel](),
	)
}

// EndpointResourceTFModel uses the endpoint name as its id, because PostHog
// addresses an endpoint by name in every URL.
type EndpointResourceTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name                        types.String        `tfsdk:"name"`
	QueryJSON                   core.NormalizedJSON `tfsdk:"query_json"`
	Description                 types.String        `tfsdk:"description"`
	DataFreshnessSeconds        types.Int64         `tfsdk:"data_freshness_seconds"`
	IsActive                    types.Bool          `tfsdk:"is_active"`
	IsMaterialized              types.Bool          `tfsdk:"is_materialized"`
	BucketOverrides             types.Map           `tfsdk:"bucket_overrides"`
	Tags                        types.Set           `tfsdk:"tags"`
	OptionalBreakdownProperties types.Set           `tfsdk:"optional_breakdown_properties"`
	CurrentVersion              types.Int64         `tfsdk:"current_version"`
	RunURL                      types.String        `tfsdk:"run_url"`
	CreatedAt                   types.String        `tfsdk:"created_at"`
}

type EndpointOps struct{}

func (o EndpointOps) ResourceName() string {
	return "endpoint"
}

func (o EndpointOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog endpoints. An endpoint turns a HogQL or insight query into a named, versioned API route " +
			"that callers run with `POST /api/projects/{project_id}/endpoints/{name}/run/`. " +
			"Changing the query creates a new version, and the endpoint serves the latest one.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The endpoint name. PostHog addresses endpoints by name, so the id and `name` are the same.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "URL-safe name, unique in the project. It must start with a letter and hold only letters, digits, hyphens " +
					"and underscores, up to 128 characters. The name is part of the run URL and PostHog cannot rename an endpoint, " +
					"so changing it replaces the endpoint.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(endpointNameRegex,
						"must start with a letter and contain only letters, digits, hyphens or underscores, up to 128 characters"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"query_json": schema.StringAttribute{
				Required:   true,
				CustomType: core.NormalizedJSONType{},
				MarkdownDescription: "JSON-encoded query the endpoint runs, for example a `HogQLQuery` or a `TrendsQuery`. " +
					"Changing it creates a new version in place. A HogQL query references a SQL variable as `{variables.code_name}`. " +
					"PostHog looks the variable up by code name and fills in `variables` itself, so the variable must exist first. " +
					"Declaring it under `variables`, keyed by the `posthog_insight_variable` id with `variableId` and `code_name`, " +
					"also makes Terraform create the variable before the endpoint. " +
					"The values and other fields PostHog adds to the stored query do not show as drift. " +
					"Callers pass variable values by code name when they run the endpoint: `{\"variables\": {\"code_name\": \"value\"}}`.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of what the endpoint returns.",
			},
			"data_freshness_seconds": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "How fresh the results must be, in seconds. It sets the cache duration, and the sync interval when " +
					"the endpoint is materialized. One of `900`, `1800`, `3600`, `21600`, `43200`, `86400` or `604800`. " +
					"PostHog defaults to `86400`. An organization's plan can set a minimum for materialized endpoints.",
				Validators: []validator.Int64{
					int64validator.OneOf(EndpointDataFreshnessSeconds...),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"is_active": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Whether callers can run the endpoint. Defaults to `true`. " +
					"Deactivating an endpoint turns off its materialization.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"is_materialized": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Whether PostHog precomputes the results on a schedule set by `data_freshness_seconds`. Defaults to `false`. " +
					"This reports whether materialization is turned on, not whether the first run has finished. " +
					"Only active endpoints whose query PostHog can materialize accept `true`.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"bucket_overrides": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				MarkdownDescription: "Bucket function per column for materialized range variables, for example `{ timestamp = \"hour\" }`. " +
					"Requires `is_materialized = true`. PostHog checks the bucket names.",
			},
			"tags": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Tags on the endpoint.",
			},
			"optional_breakdown_properties": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
				MarkdownDescription: "Breakdown properties a caller may leave out of a run, for insight queries with a `breakdownFilter`. " +
					"A run without one returns data aggregated across all its values. By default every breakdown is required.",
			},
			"current_version": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Latest version number. It goes up by one each time `query_json` changes.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"run_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Absolute URL that runs the endpoint.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the endpoint was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o EndpointOps) BuildCreateRequest(ctx context.Context, model EndpointResourceTFModel) (httpclient.EndpointRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := model.Name.ValueString()
	req := httpclient.EndpointRequest{Name: &name}

	query, d := parseEndpointQuery(model.QueryJSON)
	diags.Append(d...)
	req.Query = query

	if core.IsKnown(model.Description) {
		req.Description = util.StringPtr(model.Description.ValueString())
	}
	if core.IsKnown(model.DataFreshnessSeconds) {
		req.DataFreshnessSeconds = util.Int64Ptr(model.DataFreshnessSeconds.ValueInt64())
	}
	if core.IsKnown(model.IsActive) {
		req.IsActive = util.BoolPtr(model.IsActive.ValueBool())
	}
	if core.IsKnown(model.IsMaterialized) {
		req.IsMaterialized = util.BoolPtr(model.IsMaterialized.ValueBool())
	}
	if core.IsKnown(model.BucketOverrides) {
		overrides, d := endpointBucketOverrides(ctx, model.BucketOverrides)
		diags.Append(d...)
		req.BucketOverrides = &overrides
	}
	if core.IsKnown(model.Tags) {
		tags, d := core.ExtractTags(ctx, model.Tags)
		diags.Append(d...)
		req.Tags = &tags
	}
	if core.IsKnown(model.OptionalBreakdownProperties) {
		properties, d := core.ExtractTags(ctx, model.OptionalBreakdownProperties)
		diags.Append(d...)
		req.OptionalBreakdownProperties = &properties
	}

	return req, diags
}

// BuildUpdateRequest sends only what changed. Resending is_materialized or
// data_freshness_seconds on a materialized endpoint makes PostHog re-enable its
// materialization, and resending a query it has already stored would be a no-op
// at best.
func (o EndpointOps) BuildUpdateRequest(ctx context.Context, plan, state EndpointResourceTFModel) (httpclient.EndpointRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	req := httpclient.EndpointRequest{}

	queryChanged, d := endpointQueryChanged(ctx, plan.QueryJSON, state.QueryJSON)
	diags.Append(d...)
	if queryChanged {
		query, d := parseEndpointQuery(plan.QueryJSON)
		diags.Append(d...)
		req.Query = query
	}

	if !plan.Description.Equal(state.Description) {
		// PostHog ignores a null description, so clearing sends an empty one.
		req.Description = util.StringPtr(plan.Description.ValueString())
	}
	if core.IsKnown(plan.DataFreshnessSeconds) && !plan.DataFreshnessSeconds.Equal(state.DataFreshnessSeconds) {
		req.DataFreshnessSeconds = util.Int64Ptr(plan.DataFreshnessSeconds.ValueInt64())
	}
	if core.IsKnown(plan.IsActive) && !plan.IsActive.Equal(state.IsActive) {
		req.IsActive = util.BoolPtr(plan.IsActive.ValueBool())
	}
	if core.IsKnown(plan.IsMaterialized) && !plan.IsMaterialized.Equal(state.IsMaterialized) {
		req.IsMaterialized = util.BoolPtr(plan.IsMaterialized.ValueBool())
	}
	if !plan.BucketOverrides.IsUnknown() && !plan.BucketOverrides.Equal(state.BucketOverrides) {
		overrides := map[string]string{}
		if !plan.BucketOverrides.IsNull() {
			overrides, d = endpointBucketOverrides(ctx, plan.BucketOverrides)
			diags.Append(d...)
		}
		req.BucketOverrides = &overrides
	}
	if !plan.Tags.IsUnknown() && !plan.Tags.Equal(state.Tags) {
		tags := []string{}
		if !plan.Tags.IsNull() {
			tags, d = core.ExtractTags(ctx, plan.Tags)
			diags.Append(d...)
		}
		req.Tags = &tags
	}
	if !plan.OptionalBreakdownProperties.IsUnknown() && !plan.OptionalBreakdownProperties.Equal(state.OptionalBreakdownProperties) {
		properties := []string{}
		if !plan.OptionalBreakdownProperties.IsNull() {
			properties, d = core.ExtractTags(ctx, plan.OptionalBreakdownProperties)
			diags.Append(d...)
		}
		req.OptionalBreakdownProperties = &properties
	}

	return req, diags
}

func (o EndpointOps) MapResponseToModel(ctx context.Context, resp httpclient.Endpoint, model *EndpointResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp.MaterializationError != nil && *resp.MaterializationError != "" {
		diags.AddError(
			"Endpoint materialization failed",
			fmt.Sprintf("PostHog saved the new query of endpoint %q but could not materialize it: %s. "+
				"Fix the query or set is_materialized = false, then apply again.", resp.Name, *resp.MaterializationError),
		)
		return diags
	}

	model.ID = types.StringValue(resp.Name)
	model.Name = types.StringValue(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.DataFreshnessSeconds = types.Int64PointerValue(resp.DataFreshnessSeconds)
	model.IsActive = core.PtrToBool(resp.IsActive)
	model.CurrentVersion = types.Int64PointerValue(resp.CurrentVersion)
	model.RunURL = core.PtrToStringNullIfEmptyTrimmed(resp.URL)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)

	materialized := resp.Materialization != nil && resp.Materialization.Enabled != nil && *resp.Materialization.Enabled
	model.IsMaterialized = types.BoolValue(materialized)

	overrides, d := endpointBucketOverridesToState(ctx, resp.BucketOverrides, materialized, model.BucketOverrides)
	diags.Append(d...)
	model.BucketOverrides = overrides

	tags, d := core.TagsToSetPreserveEmpty(ctx, resp.Tags, model.Tags)
	diags.Append(d...)
	model.Tags = tags

	properties, d := core.TagsToSetPreserveEmpty(ctx, resp.OptionalBreakdownProperties, model.OptionalBreakdownProperties)
	diags.Append(d...)
	model.OptionalBreakdownProperties = properties

	userQueryJSON := ""
	if core.IsKnown(model.QueryJSON) {
		userQueryJSON = model.QueryJSON.ValueString()
	}
	query, err := endpointQueryForState(resp.Query, userQueryJSON)
	if err != nil {
		diags.AddError("Failed to normalize query", err.Error())
		return diags
	}
	model.QueryJSON = core.NewNormalizedJSONValue(query)

	return diags
}

func (o EndpointOps) Create(ctx context.Context, client httpclient.PosthogClient, model EndpointResourceTFModel, req httpclient.EndpointRequest) (httpclient.Endpoint, error) {
	return client.CreateEndpoint(ctx, model.GetEffectiveProjectID(), req)
}

func (o EndpointOps) Read(ctx context.Context, client httpclient.PosthogClient, model EndpointResourceTFModel) (httpclient.Endpoint, httpclient.HTTPStatusCode, error) {
	return client.GetEndpoint(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o EndpointOps) Update(ctx context.Context, client httpclient.PosthogClient, model EndpointResourceTFModel, req httpclient.EndpointRequest) (httpclient.Endpoint, httpclient.HTTPStatusCode, error) {
	return client.UpdateEndpoint(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o EndpointOps) Delete(ctx context.Context, client httpclient.PosthogClient, model EndpointResourceTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteEndpoint(ctx, model.GetEffectiveProjectID(), model.GetID())
}

// ModifyResourcePlan rejects combinations PostHog refuses, and plans the values
// PostHog changes on its own so the apply result matches the plan.
func (o EndpointOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan, config EndpointResourceTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(validateEndpointPlan(plan, config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Deactivating an endpoint turns off its materialization, so a carried-over
	// is_materialized = true would not survive the apply.
	if config.IsMaterialized.IsNull() && core.IsKnown(plan.IsActive) && !plan.IsActive.ValueBool() {
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("is_materialized"), types.BoolValue(false))...)
	}

	if req.State.Raw.IsNull() {
		return
	}
	var state EndpointResourceTFModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	queryChanged, d := endpointQueryChanged(ctx, plan.QueryJSON, state.QueryJSON)
	resp.Diagnostics.Append(d...)
	if queryChanged {
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("current_version"), types.Int64Unknown())...)
	}
}

// validateEndpointPlan is separate from ModifyResourcePlan so it can be unit
// tested against a model.
func validateEndpointPlan(plan, config EndpointResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	active, activeKnown := util.ResolveBool(plan.IsActive, config.IsActive, defaultEndpointIsActive)
	materialized, materializedKnown := util.ResolveBool(plan.IsMaterialized, config.IsMaterialized, defaultEndpointIsMaterialized)

	if activeKnown && materializedKnown && materialized && !active {
		diags.AddAttributeError(
			path.Root("is_materialized"),
			"Cannot materialize an inactive endpoint",
			"PostHog only materializes active endpoints. Set is_active = true, or set is_materialized = false.",
		)
	}

	hasOverrides := core.IsKnown(plan.BucketOverrides) && len(plan.BucketOverrides.Elements()) > 0
	if hasOverrides && materializedKnown && !materialized {
		diags.AddAttributeError(
			path.Root("bucket_overrides"),
			"bucket_overrides needs materialization",
			"PostHog applies bucket overrides only to a materialized endpoint. Set is_materialized = true, or remove bucket_overrides.",
		)
	}

	return diags
}

// parseEndpointQuery decodes query_json. The attribute type already rejects
// invalid JSON, so this only fails for valid JSON that is not an object.
func parseEndpointQuery(v core.NormalizedJSON) (map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	if !core.IsKnown(v) {
		return nil, diags
	}
	var query map[string]interface{}
	if err := json.Unmarshal([]byte(v.ValueString()), &query); err != nil {
		diags.AddAttributeError(path.Root("query_json"), "Invalid query_json", "query_json must be a JSON object: "+err.Error())
		return nil, diags
	}
	return query, diags
}

// endpointQueryChanged reports whether the planned query differs from state by
// content. An unknown plan counts as a change.
func endpointQueryChanged(ctx context.Context, plan, state core.NormalizedJSON) (bool, diag.Diagnostics) {
	if plan.IsUnknown() {
		return true, nil
	}
	if plan.IsNull() || state.IsNull() || state.IsUnknown() {
		return plan.IsNull() != state.IsNull(), nil
	}
	equal, diags := plan.StringSemanticEquals(ctx, state)
	return !equal, diags
}

// endpointQueryForState keeps only the query fields the configuration sets, so
// the fields PostHog adds (null defaults, variable values) do not show as drift.
// With no configuration, as on import, it keeps the fields a configuration
// would write: every non-null field, without the variable values PostHog
// filled in.
func endpointQueryForState(apiQuery map[string]interface{}, userQueryJSON string) (string, error) {
	if userQueryJSON == "" {
		query, _ := dropJSONNulls(apiQuery).(map[string]interface{})
		return marshalJSON(dropFilledVariableValues(query))
	}
	return normalizeQueryForState(apiQuery, userQueryJSON)
}

// endpointFilledVariableFields are the fields PostHog adds to each entry of a
// HogQL query's variables from the variable's default.
var endpointFilledVariableFields = map[string]struct{}{"value": {}, "isNull": {}}

// dropFilledVariableValues removes the value PostHog filled in on each
// declared variable, so an imported endpoint matches a configuration that
// declares only variableId and code_name.
func dropFilledVariableValues(query map[string]interface{}) map[string]interface{} {
	variables, ok := query["variables"].(map[string]interface{})
	if !ok {
		return query
	}
	for id, entry := range variables {
		if fields, ok := entry.(map[string]interface{}); ok {
			variables[id] = util.StripFields(fields, endpointFilledVariableFields)
		}
	}
	return query
}

// dropJSONNulls removes null object members at every depth. PostHog stores a
// query with every optional field it knows, most of them null.
func dropJSONNulls(v interface{}) interface{} {
	switch value := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(value))
		for key, item := range value {
			if item == nil {
				continue
			}
			result[key] = dropJSONNulls(item)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(value))
		for i, item := range value {
			result[i] = dropJSONNulls(item)
		}
		return result
	default:
		return v
	}
}

func endpointBucketOverrides(ctx context.Context, v types.Map) (map[string]string, diag.Diagnostics) {
	overrides := map[string]string{}
	diags := v.ElementsAs(ctx, &overrides, false)
	return overrides, diags
}

// endpointBucketOverridesToState reads bucket overrides only while the endpoint
// is materialized, because PostHog ignores them otherwise and may still return
// the last ones it stored.
func endpointBucketOverridesToState(ctx context.Context, overrides map[string]string, materialized bool, current types.Map) (types.Map, diag.Diagnostics) {
	if materialized && len(overrides) > 0 {
		return types.MapValueFrom(ctx, types.StringType, overrides)
	}
	if core.IsKnown(current) && len(current.Elements()) == 0 {
		return types.MapValueFrom(ctx, types.StringType, map[string]string{})
	}
	return types.MapNull(types.StringType), nil
}
