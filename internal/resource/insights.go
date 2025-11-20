// Package resource provides Terraform resource implementations for PostHog entities.
//
// Architecture:
// This package implements Terraform resources using the Plugin Framework and directly
// leverages Swagger-generated API clients for type safety and consistency. The implementation
// follows these principles:
//
//   - Direct usage of Swagger-generated types (posthogapi.Insight, posthogapi.PatchedInsight)
//     eliminates custom intermediate models and reduces maintenance burden
//   - Clean separation between Terraform's type system (types.String, types.Int64) and
//     Swagger's nullable types (NullableString, pointers)
//   - Idiomatic error handling with proper context propagation
//   - Comprehensive logging for observability
//   - Server-side query normalization to prevent state drift
package resource

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	internaldata "github.com/posthog/terraform-provider/internal/data"
	posthogapi "github.com/posthog/terraform-provider/internal/posthog/swagger"
)

// Compile-time assertions to ensure provider types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &Insight{}
	_ resource.ResourceWithImportState = &Insight{}
)

// NewInsight creates a new Insight resource.
func NewInsight() resource.Resource {
	return &Insight{}
}

// Insight implements the PostHog insight resource.
type Insight struct {
	swaggerClient *posthogapi.APIClient
	projectId     string
}

// InsightResourceModel describes the Terraform resource schema for a PostHog insight.
// This model maps 1:1 with the Terraform configuration and state.
type InsightResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	DerivedName    types.String `tfsdk:"derived_name"`
	Description    types.String `tfsdk:"description"`
	QueryJSON      types.String `tfsdk:"query_json"`
	Tags           types.Set    `tfsdk:"tags"`
	CreateInFolder types.String `tfsdk:"create_in_folder"`
	DashboardIDs   types.Set    `tfsdk:"dashboard_ids"`
}

// insightAPIResponse represents the actual API response from PostHog.
// The Swagger spec has type mismatches, so we need this intermediate model.
type insightAPIResponse struct {
	ID             int64                  `json:"id"`
	Name           string                 `json:"name"`
	DerivedName    string                 `json:"derived_name"`
	Description    string                 `json:"description"`
	Query          map[string]interface{} `json:"query"`
	Tags           []string               `json:"tags"`
	CreateInFolder string                 `json:"_create_in_folder"`
	DashboardIDs   []int32                `json:"dashboards"`
}

// Metadata returns the resource type name.
func (r *Insight) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_insight"
}

// Schema defines the schema for the insight resource.
func (r *Insight) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage PostHog insights via the insights endpoints.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Numeric ID of the insight.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Insight name.",
			},
			"derived_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Insight derived name.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Insight description.",
			},
			"query_json": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Raw JSON serialized query payload accepted by PostHog (for example an `InsightVizNode` with a `TrendsQuery`).",
			},
			"tags": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "List of tags to apply to the insight.",
			},
			"create_in_folder": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The folder where the insight is created.",
			},
			"dashboard_ids": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.Int32Type,
				MarkdownDescription: "List of dashboard ids which should contain the insight.",
			},
		},
	}
}

// Configure initializes the resource with provider-level data.
func (r *Insight) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(internaldata.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected internaldata.ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.swaggerClient = data.SwaggerClient
	r.projectId = data.ProjectID
}

// Create creates a new insight resource.
func (r *Insight) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InsightResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build the Swagger API payload directly from the Terraform model
	payload, diags := buildInsightPayload(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Execute the API call
	swaggerInsight, httpResp, err := r.swaggerClient.InsightsAPI.InsightsCreate(ctx, r.projectId).Insight(*payload).Execute()

	// Handle API responses with extra fields not in Swagger spec (e.g., "filters")
	var apiResp *insightAPIResponse
	if err != nil {
		apiResp, err = unmarshalInsightFromError(err, httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Failed to create insight", formatAPIError(err, httpResp))
			return
		}
	} else {
		// Convert Swagger response to our API response model
		apiResp = swaggerInsightToAPIResponse(swaggerInsight)
	}

	tflog.Info(ctx, "created PostHog insight", map[string]any{
		"project_id": r.projectId,
		"id":         apiResp.ID,
	})

	// Update state from API response
	resp.Diagnostics.Append(setInsightState(ctx, &plan, apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read retrieves the current state of an insight resource.
func (r *Insight) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InsightResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.IsNull() || state.ID.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	insightID := int32(state.ID.ValueInt64())
	swaggerInsight, httpResp, err := r.swaggerClient.InsightsAPI.InsightsRetrieve(ctx, insightID, r.projectId).Execute()

	// Handle 404 by removing from state
	if err != nil && httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
		tflog.Warn(ctx, "insight not found, removing from state", map[string]any{
			"project_id": r.projectId,
			"id":         insightID,
		})
		resp.State.RemoveResource(ctx)
		return
	}

	// Handle API responses with extra fields not in Swagger spec (e.g., "filters")
	var apiResp *insightAPIResponse
	if err != nil {
		apiResp, err = unmarshalInsightFromError(err, httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Failed to read insight", formatAPIError(err, httpResp))
			return
		}
	} else {
		apiResp = swaggerInsightToAPIResponse(swaggerInsight)
	}

	// Update state from API response
	resp.Diagnostics.Append(setInsightState(ctx, &state, apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies an existing insight resource.
func (r *Insight) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InsightResourceModel
	var state InsightResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.IsNull() || state.ID.IsUnknown() {
		resp.Diagnostics.AddError("Missing resource ID", "Resource ID is absent in state")
		return
	}

	// Build the patch payload
	payload, diags := buildPatchedInsightPayload(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	insightID := int32(state.ID.ValueInt64())
	swaggerInsight, httpResp, err := r.swaggerClient.InsightsAPI.InsightsPartialUpdate(ctx, insightID, r.projectId).PatchedInsight(*payload).Execute()

	// Handle 404 by removing from state
	if err != nil && httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		resp.Diagnostics.AddError("Insight not found", fmt.Sprintf("Insight %d no longer exists in PostHog", insightID))
		return
	}

	// Handle API responses with extra fields not in Swagger spec (e.g., "filters")
	var apiResp *insightAPIResponse
	if err != nil {
		apiResp, err = unmarshalInsightFromError(err, httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Failed to update insight", formatAPIError(err, httpResp))
			return
		}
	} else {
		apiResp = swaggerInsightToAPIResponse(swaggerInsight)
	}

	tflog.Info(ctx, "updated PostHog insight", map[string]any{
		"project_id": r.projectId,
		"id":         insightID,
	})

	// Update state from API response
	resp.Diagnostics.Append(setInsightState(ctx, &plan, apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete soft-deletes an insight resource by setting the deleted flag.
func (r *Insight) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InsightResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.IsNull() || state.ID.IsUnknown() {
		return
	}

	insightID := int32(state.ID.ValueInt64())

	// Soft delete by setting deleted flag
	payload := posthogapi.NewPatchedInsight()
	payload.SetDeleted(true)

	_, httpResp, err := r.swaggerClient.InsightsAPI.InsightsPartialUpdate(ctx, insightID, r.projectId).PatchedInsight(*payload).Execute()

	// 404 means already deleted - this is acceptable
	if err != nil && httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
		tflog.Info(ctx, "insight already deleted", map[string]any{
			"project_id": r.projectId,
			"id":         insightID,
		})
		return
	}

	// Success response (2xx) is also acceptable even if there's an error object
	if httpResp != nil && httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		tflog.Info(ctx, "deleted PostHog insight", map[string]any{
			"project_id": r.projectId,
			"id":         insightID,
		})
		return
	}

	if err != nil {
		resp.Diagnostics.AddError("Failed to delete insight", formatAPIError(err, httpResp))
	}
}

// ImportState imports an existing insight resource using its numeric ID.
func (r *Insight) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(strings.TrimSpace(req.ID), 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("Expected numeric insight ID, got %q: %v", req.ID, err))
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

// setInsightState populates the Terraform model from the API response.
func setInsightState(ctx context.Context, model *InsightResourceModel, apiResp *insightAPIResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	// Set ID
	model.ID = types.Int64Value(apiResp.ID)

	// Set string fields
	if strings.TrimSpace(apiResp.Name) != "" {
		model.Name = types.StringValue(apiResp.Name)
	} else {
		model.Name = types.StringNull()
	}

	if strings.TrimSpace(apiResp.DerivedName) != "" {
		model.DerivedName = types.StringValue(apiResp.DerivedName)
	} else {
		model.DerivedName = types.StringNull()
	}

	if strings.TrimSpace(apiResp.Description) != "" {
		model.Description = types.StringValue(apiResp.Description)
	} else {
		model.Description = types.StringNull()
	}

	// Set tags
	if len(apiResp.Tags) > 0 {
		tagsSet, tagDiags := types.SetValueFrom(ctx, types.StringType, apiResp.Tags)
		diags.Append(tagDiags...)
		if diags.HasError() {
			return diags
		}
		model.Tags = tagsSet
	} else if !model.Tags.IsNull() {
		// Preserve configured empty set
		emptySet, tagDiags := types.SetValueFrom(ctx, types.StringType, []string{})
		diags.Append(tagDiags...)
		if diags.HasError() {
			return diags
		}
		model.Tags = emptySet
	} else {
		model.Tags = types.SetNull(types.StringType)
	}

	// Set create_in_folder if present
	if strings.TrimSpace(apiResp.CreateInFolder) != "" {
		model.CreateInFolder = types.StringValue(apiResp.CreateInFolder)
	}

	if len(apiResp.DashboardIDs) > 0 {
		dashboardIds, dashboardDiags := types.SetValueFrom(ctx, types.Int32Type, apiResp.DashboardIDs)
		diags.Append(dashboardDiags...)
		if diags.HasError() {
			return diags
		}
		model.DashboardIDs = dashboardIds
	} else if !model.DashboardIDs.IsNull() {
		emptySet, dashboardDiags := types.SetValueFrom(ctx, types.Int32Type, []int32{})
		diags.Append(dashboardDiags...)
		if diags.HasError() {
			return diags
		}
		model.DashboardIDs = emptySet
	} else {
		model.DashboardIDs = types.SetNull(types.Int32Type)
	}

	// Normalize and set query JSON
	if apiResp.Query != nil {
		normalizedQuery, err := normalizeQueryMap(apiResp.Query)
		if err != nil {
			diags.AddError("Failed to normalize query", err.Error())
			return diags
		}
		model.QueryJSON = types.StringValue(normalizedQuery)
	}

	return diags
}

// buildInsightPayload creates a Swagger Insight payload from the Terraform model.
func buildInsightPayload(ctx context.Context, model InsightResourceModel) (*posthogapi.Insight, diag.Diagnostics) {
	var diags diag.Diagnostics

	payload := posthogapi.NewInsightWithDefaults()

	// Parse and set query
	query, err := parseAndValidateQuery(model.QueryJSON)
	if err != nil {
		diags.AddError("Invalid query_json", err.Error())
		return nil, diags
	}
	payload.Query = query

	// Set nullable string fields using Swagger's NullableString setters
	applyNullableString(&payload.Name, model.Name)
	applyNullableString(&payload.DerivedName, model.DerivedName)
	applyNullableString(&payload.Description, model.Description)

	// Set tags
	if !model.Tags.IsNull() && !model.Tags.IsUnknown() {
		var tags []string
		diags.Append(model.Tags.ElementsAs(ctx, &tags, false)...)
		if diags.HasError() {
			return nil, diags
		}
		payload.Tags = convertStringsToInterfaces(tags)
	}

	// Set folder
	if !model.CreateInFolder.IsNull() && !model.CreateInFolder.IsUnknown() {
		folder := model.CreateInFolder.ValueString()
		payload.CreateInFolder = &folder
	}

	if !model.DashboardIDs.IsUnknown() {
		if model.DashboardIDs.IsNull() {
			// User wants to remove all dashboard associations - send empty array
			payload.Dashboards = []int32{}
		} else {
			// User specified dashboard IDs
			var ids []int32
			diags.Append(model.DashboardIDs.ElementsAs(ctx, &ids, false)...)
			if diags.HasError() {
				return nil, diags
			}
			payload.Dashboards = ids
		}
	}

	return payload, diags
}

// buildPatchedInsightPayload creates a Swagger PatchedInsight payload from the Terraform model.
func buildPatchedInsightPayload(ctx context.Context, model InsightResourceModel) (*posthogapi.PatchedInsight, diag.Diagnostics) {
	var diags diag.Diagnostics

	payload := posthogapi.NewPatchedInsight()

	// Parse and set query
	query, err := parseAndValidateQuery(model.QueryJSON)
	if err != nil {
		diags.AddError("Invalid query_json", err.Error())
		return nil, diags
	}
	payload.Query = query

	// Set nullable string fields
	applyNullableString(&payload.Name, model.Name)
	applyNullableString(&payload.DerivedName, model.DerivedName)
	applyNullableString(&payload.Description, model.Description)

	// Set tags
	if !model.Tags.IsNull() && !model.Tags.IsUnknown() {
		var tags []string
		diags.Append(model.Tags.ElementsAs(ctx, &tags, false)...)
		if diags.HasError() {
			return nil, diags
		}
		payload.Tags = convertStringsToInterfaces(tags)
	}

	// Set folder
	if !model.CreateInFolder.IsNull() && !model.CreateInFolder.IsUnknown() {
		folder := model.CreateInFolder.ValueString()
		payload.CreateInFolder = &folder
	}

	if !model.DashboardIDs.IsUnknown() {
		if model.DashboardIDs.IsNull() {
			// User wants to remove all dashboard associations - send empty array
			payload.Dashboards = []int32{}
		} else {
			// User specified dashboard IDs
			var ids []int32
			diags.Append(model.DashboardIDs.ElementsAs(ctx, &ids, false)...)
			if diags.HasError() {
				return nil, diags
			}
			payload.Dashboards = ids
		}
	}

	return payload, diags
}

// parseAndValidateQuery parses and validates the query JSON string.
func parseAndValidateQuery(queryJSON types.String) (map[string]interface{}, error) {
	if queryJSON.IsNull() || queryJSON.IsUnknown() {
		return nil, fmt.Errorf("query_json is required")
	}

	raw := strings.TrimSpace(queryJSON.ValueString())
	if raw == "" {
		return nil, fmt.Errorf("query_json cannot be empty")
	}

	if !json.Valid([]byte(raw)) {
		return nil, fmt.Errorf("query_json must be valid JSON")
	}

	var query map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &query); err != nil {
		return nil, fmt.Errorf("failed to parse query_json: %w", err)
	}

	return query, nil
}

// applyNullableString sets a NullableString field based on a Terraform types.String value.
func applyNullableString(dst *posthogapi.NullableString, src types.String) {
	if src.IsNull() || src.IsUnknown() {
		dst.Unset()
		return
	}

	value := strings.TrimSpace(src.ValueString())
	if value == "" {
		dst.Set(nil) // Set explicit nil for empty strings
	} else {
		dst.Set(&value)
	}
}

// nullableStringToTerraform converts a Swagger NullableString to a Terraform types.String.
func nullableStringToTerraform(ns posthogapi.NullableString) types.String {
	if !ns.IsSet() {
		return types.StringNull()
	}

	val := ns.Get()
	if val == nil || strings.TrimSpace(*val) == "" {
		return types.StringNull()
	}

	return types.StringValue(*val)
}

// extractStringTags converts []interface{} tags to []string.
func extractStringTags(tags []interface{}) []string {
	if len(tags) == 0 {
		return nil
	}

	result := make([]string, 0, len(tags))
	for _, tag := range tags {
		if str, ok := tag.(string); ok {
			result = append(result, str)
		}
	}

	return result
}

// convertStringsToInterfaces converts []string to []interface{} for API payloads.
func convertStringsToInterfaces(tags []string) []interface{} {
	if len(tags) == 0 {
		return []interface{}{}
	}

	result := make([]interface{}, len(tags))
	for i, tag := range tags {
		result[i] = tag
	}

	return result
}

// normalizeQueryMap normalizes a query map into canonical JSON string for state storage.
// This strips server-only fields and ensures consistent formatting.
func normalizeQueryMap(query map[string]interface{}) (string, error) {
	if query == nil {
		return "", nil
	}

	// Server-only fields that should be stripped from state
	serverOnlyFields := map[string]struct{}{
		"filters":                     {},
		"result":                      {},
		"hogql":                       {},
		"columns":                     {},
		"cache_target_age":            {},
		"next_allowed_client_refresh": {},
		"created_at":                  {},
		"updated_at":                  {},
		"last_refresh":                {},
		"last_modified_at":            {},
		"last_modified_by":            {},
		"last_viewed_at":              {},
		"timezone":                    {},
		"is_cached":                   {},
		"query_status":                {},
		"types":                       {},
	}

	// Create a copy and strip server fields
	normalized := make(map[string]interface{})
	for key, value := range query {
		if _, isServerField := serverOnlyFields[key]; !isServerField {
			normalized[key] = value
		}
	}

	// Marshal to canonical JSON (sorted keys)
	return marshalCanonical(normalized)
}

// marshalCanonical marshals data to JSON with sorted keys for consistency.
func marshalCanonical(v interface{}) (string, error) {
	buf := &bytes.Buffer{}
	if err := encodeCanonical(buf, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// encodeCanonical recursively encodes values to canonical JSON (sorted keys).
func encodeCanonical(buf *bytes.Buffer, v interface{}) error {
	switch val := v.(type) {
	case map[string]interface{}:
		buf.WriteByte('{')
		keys := make([]string, 0, len(val))
		for key := range val {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for i, key := range keys {
			if i > 0 {
				buf.WriteByte(',')
			}
			keyJSON, _ := json.Marshal(key)
			buf.Write(keyJSON)
			buf.WriteByte(':')
			if err := encodeCanonical(buf, val[key]); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case []interface{}:
		buf.WriteByte('[')
		for i, item := range val {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encodeCanonical(buf, item); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	default:
		// For primitives, use standard JSON marshaling
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		buf.Write(data)
	}

	return nil
}

// swaggerInsightToAPIResponse converts a Swagger Insight to our internal API response model.
func swaggerInsightToAPIResponse(src *posthogapi.Insight) *insightAPIResponse {
	if src == nil {
		return &insightAPIResponse{}
	}

	resp := &insightAPIResponse{
		ID:    int64(src.GetId()),
		Query: src.GetQuery(),
	}

	// Extract nullable strings
	if name, ok := src.GetNameOk(); ok && name != nil {
		resp.Name = *name
	}
	if derived, ok := src.GetDerivedNameOk(); ok && derived != nil {
		resp.DerivedName = *derived
	}
	if desc, ok := src.GetDescriptionOk(); ok && desc != nil {
		resp.Description = *desc
	}
	if folder, ok := src.GetCreateInFolderOk(); ok && folder != nil {
		resp.CreateInFolder = *folder
	}
	if dashboardTiles, ok := src.GetDashboardTilesOk(); ok && len(dashboardTiles) > 0 {
		result := make([]int32, len(dashboardTiles))
		for i, tile := range dashboardTiles {
			result[i] = tile.DashboardId
		}
		resp.DashboardIDs = result
	} else if dashboardIds, ok := src.GetDashboardsOk(); ok && len(dashboardIds) > 0 {
		resp.DashboardIDs = dashboardIds
	}

	// Extract tags
	if tags := src.GetTags(); len(tags) > 0 {
		resp.Tags = make([]string, 0, len(tags))
		for _, tag := range tags {
			if str, ok := tag.(string); ok {
				resp.Tags = append(resp.Tags, str)
			}
		}
	}

	return resp
}

// unmarshalInsightFromError extracts an insight from API error responses.
// The Swagger spec has type mismatches with the actual PostHog API, so we use
// a custom struct that matches the real API response.
func unmarshalInsightFromError(err error, resp *http.Response) (*insightAPIResponse, error) {
	// Only handle successful responses (2xx) that failed unmarshaling
	if resp == nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, err
	}

	// Extract body from Swagger error
	apiErr, ok := err.(*posthogapi.GenericOpenAPIError)
	if !ok || len(apiErr.Body()) == 0 {
		return nil, err
	}

	// Unmarshal directly to our custom struct (ignores extra fields automatically)
	var insight insightAPIResponse
	if err := json.Unmarshal(apiErr.Body(), &insight); err != nil {
		return nil, fmt.Errorf("failed to unmarshal insight response: %w", err)
	}

	return &insight, nil
}

// formatAPIError formats API errors into human-readable messages.
func formatAPIError(err error, resp *http.Response) string {
	if err == nil {
		return "unknown error"
	}

	var msg strings.Builder

	// Add HTTP status if available
	if resp != nil {
		msg.WriteString(fmt.Sprintf("HTTP %s", resp.Status))
	}

	// Extract error details from Swagger error
	if apiErr, ok := err.(*posthogapi.GenericOpenAPIError); ok {
		if msg.Len() > 0 {
			msg.WriteString(": ")
		}
		msg.WriteString(apiErr.Error())

		// Include response body if available
		if body := apiErr.Body(); len(body) > 0 {
			msg.WriteString(" (")
			msg.WriteString(string(body))
			msg.WriteByte(')')
		}
	} else {
		if msg.Len() > 0 {
			msg.WriteString(": ")
		}
		msg.WriteString(err.Error())
	}

	return msg.String()
}
