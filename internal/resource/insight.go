package resource

import (
	"bytes"
	"context"
	"encoding/json"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewInsight() resource.Resource {
	return core.NewGenericResource[InsightResourceTFModel, httpclient.InsightRequest, httpclient.Insight](InsightOps{})
}

type InsightResourceTFModel struct {
	core.BaseInt64Identifiable
	Name           types.String `tfsdk:"name"`
	DerivedName    types.String `tfsdk:"derived_name"`
	Description    types.String `tfsdk:"description"`
	QueryJSON      types.String `tfsdk:"query_json"`
	Tags           types.Set    `tfsdk:"tags"`
	CreateInFolder types.String `tfsdk:"create_in_folder"`
	DashboardIDs   types.Set    `tfsdk:"dashboard_ids"`
	Deleted        types.Bool   `tfsdk:"deleted"`
}

type InsightOps struct{}

func (o InsightOps) ResourceName() string {
	return "Insight"
}

func (o InsightOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog insights via the insights endpoints.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Numeric ID of the insight.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
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
			"deleted": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the insight is deleted (soft delete)",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o InsightOps) BuildCreateRequest(ctx context.Context, model InsightResourceTFModel) (httpclient.InsightRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	input := httpclient.InsightRequest{}

	// Parse query JSON
	if !model.QueryJSON.IsNull() && !model.QueryJSON.IsUnknown() {
		raw := strings.TrimSpace(model.QueryJSON.ValueString())
		if raw == "" {
			diags.AddError("Invalid query_json", "query_json cannot be empty")
			return input, diags
		}
		if !json.Valid([]byte(raw)) {
			diags.AddError("Invalid query_json", "query_json must be valid JSON")
			return input, diags
		}
		var query map[string]interface{}
		if err := json.Unmarshal([]byte(raw), &query); err != nil {
			diags.AddError("Invalid query_json", err.Error())
			return input, diags
		}
		input.Query = query
	}

	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		name := model.Name.ValueString()
		input.Name = &name
	}
	if !model.DerivedName.IsNull() && !model.DerivedName.IsUnknown() {
		derivedName := model.DerivedName.ValueString()
		input.DerivedName = &derivedName
	}
	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		desc := model.Description.ValueString()
		input.Description = &desc
	}
	if !model.CreateInFolder.IsNull() && !model.CreateInFolder.IsUnknown() {
		folder := model.CreateInFolder.ValueString()
		input.CreateInFolder = &folder
	}

	// Tags
	if !model.Tags.IsNull() && !model.Tags.IsUnknown() {
		tags, d := core.ExtractTags(ctx, model.Tags)
		diags.Append(d...)
		if diags.HasError() {
			return input, diags
		}
		input.Tags = tags
	}

	// Dashboard IDs
	if !model.DashboardIDs.IsUnknown() {
		if model.DashboardIDs.IsNull() {
			input.Dashboards = []int32{}
		} else {
			var ids []int32
			diags.Append(model.DashboardIDs.ElementsAs(ctx, &ids, false)...)
			if diags.HasError() {
				return input, diags
			}
			input.Dashboards = ids
		}
	}

	return input, diags
}

func (o InsightOps) BuildUpdateRequest(ctx context.Context, plan, state InsightResourceTFModel) (httpclient.InsightRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	// Clear description if removed from config
	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = core.StringPtr("")
	}

	if !plan.Deleted.IsNull() {
		deleted := plan.Deleted.ValueBool()
		req.Deleted = &deleted
	}

	return req, diags
}

func (o InsightOps) MapResponseToModel(ctx context.Context, resp httpclient.Insight, model *InsightResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.Deleted = core.PtrToBool(resp.Deleted)

	// String fields - convert empty/whitespace to null
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.DerivedName = core.PtrToStringNullIfEmptyTrimmed(resp.DerivedName)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.CreateInFolder = core.PtrToStringNullIfEmptyTrimmed(resp.CreateInFolder)

	// Tags - preserve empty set if configured
	tagsSet, d := core.TagsToSetPreserveEmpty(ctx, resp.Tags, model.Tags)
	diags.Append(d...)
	model.Tags = tagsSet

	// Dashboard IDs - preserve empty set if configured
	dashSet, d := core.Int32SetPreserveEmpty(ctx, resp.Dashboards, model.DashboardIDs)
	diags.Append(d...)
	model.DashboardIDs = dashSet

	// Query - normalize to prevent state drift
	if resp.Query != nil {
		normalizedQuery, err := normalizeQueryMap(resp.Query)
		if err != nil {
			diags.AddError("Failed to normalize query", err.Error())
			return diags
		}
		model.QueryJSON = types.StringValue(normalizedQuery)
	}

	return diags
}

func (o InsightOps) Create(ctx context.Context, client httpclient.PosthogClient, req httpclient.InsightRequest) (httpclient.Insight, error) {
	return client.CreateInsight(ctx, req)
}

func (o InsightOps) Read(ctx context.Context, client httpclient.PosthogClient, id int64) (httpclient.Insight, httpclient.HTTPStatusCode, error) {
	return client.GetInsight(ctx, id)
}

func (o InsightOps) Update(ctx context.Context, client httpclient.PosthogClient, id int64, req httpclient.InsightRequest) (httpclient.Insight, httpclient.HTTPStatusCode, error) {
	return client.UpdateInsight(ctx, id, req)
}

func (o InsightOps) Delete(ctx context.Context, client httpclient.PosthogClient, id int64) (httpclient.HTTPStatusCode, error) {
	return client.DeleteInsight(ctx, id)
}

// normalizeQueryMap normalizes a query map to canonical JSON, stripping server-only fields.
func normalizeQueryMap(query map[string]interface{}) (string, error) {
	if query == nil {
		return "", nil
	}

	// Server-only fields to strip at top level
	topLevelServerOnlyFields := map[string]struct{}{
		"filters": {}, "result": {}, "hogql": {}, "columns": {},
		"cache_target_age": {}, "next_allowed_client_refresh": {},
		"created_at": {}, "updated_at": {}, "last_refresh": {},
		"last_modified_at": {}, "last_modified_by": {}, "last_viewed_at": {},
		"timezone": {}, "is_cached": {}, "query_status": {}, "types": {},
		"version": {},
	}

	// Server-only fields to strip from the "source" object
	sourceServerOnlyFields := map[string]struct{}{
		"version": {},
	}

	normalized := make(map[string]interface{})
	for key, value := range query {
		if _, skip := topLevelServerOnlyFields[key]; skip {
			continue
		}

		// If this is the "source" object, strip server-only fields from it too
		if key == "source" {
			if sourceMap, ok := value.(map[string]interface{}); ok {
				normalizedSource := make(map[string]interface{})
				for sourceKey, sourceValue := range sourceMap {
					if _, skip := sourceServerOnlyFields[sourceKey]; !skip {
						normalizedSource[sourceKey] = sourceValue
					}
				}
				normalized[key] = normalizedSource
				continue
			}
		}

		normalized[key] = value
	}

	return marshalCanonical(normalized)
}

// marshalCanonical marshals to JSON with sorted keys.
func marshalCanonical(v interface{}) (string, error) {
	buf := &bytes.Buffer{}
	if err := encodeCanonical(buf, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

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
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		buf.Write(data)
	}
	return nil
}
