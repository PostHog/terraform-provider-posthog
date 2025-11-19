package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	internaldata "github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/posthog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &Insight{}
	_ resource.ResourceWithImportState = &Insight{}
)

func NewInsight() resource.Resource {
	return &Insight{}
}

type Insight struct {
	client    posthog.Client
	projectId string
}

// InsightResourceModel describes the resource data model.
type InsightResourceModel struct {
	ID             types.Int64    `tfsdk:"id"`
	Name           types.String   `tfsdk:"name"`
	DerivedName    types.String   `tfsdk:"derived_name"`
	Description    types.String   `tfsdk:"description"`
	QueryJSON      types.String   `tfsdk:"query_json"`
	Tags           []types.String `tfsdk:"tags"`
	CreateInFolder types.String   `tfsdk:"create_in_folder"`
}

func (m InsightResourceModel) ToInsightRequest() (posthog.InsightRequest, error) {
	query, err := parseQueryJSON(m.QueryJSON)
	if err != nil {
		return posthog.InsightRequest{}, fmt.Errorf("invalid query_json: %w", err)
	}

	request := posthog.InsightRequest{
		Query: query,
	}
	if !m.Name.IsNull() && !m.Name.IsUnknown() {
		request.Name = m.Name.ValueString()
	}
	if !m.DerivedName.IsNull() && !m.DerivedName.IsUnknown() {
		request.DerivedName = m.DerivedName.ValueString()
	}
	if !m.Description.IsNull() && !m.Description.IsUnknown() {
		request.Description = m.Description.ValueString()
	}
	if len(m.Tags) > 0 {
		for _, tag := range m.Tags {
			if !tag.IsNull() && !tag.IsUnknown() {
				request.Tags = append(request.Tags, tag.ValueString())
			}
		}
	}
	if !m.CreateInFolder.IsNull() && !m.CreateInFolder.IsUnknown() {
		request.CreateInFolder = m.CreateInFolder.ValueString()
	}

	return request, nil
}

func (r *Insight) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_insight"
}

func (r *Insight) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage PostHog insights via the insights endpoints.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Numeric ID of the insight.",
			},
			"name": schema.StringAttribute{
				Required:            true,
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
			"tags": schema.ListAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "List of tags to apply to the insight.",
			},
			"create_in_folder": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The folder where the insight is created.",
			},
		},
	}
}

func (r *Insight) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(internaldata.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected resource.ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = data.Client
	r.projectId = data.ProjectID
}

func (r *Insight) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InsightResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	request, err := data.ToInsightRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error creating insight request", err.Error())
		return
	}

	insight, err := r.client.CreateInsight(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error creating insight", err.Error())
		return
	}

	tflog.Info(ctx, "created PostHog insight", map[string]any{
		"project_id": r.projectId,
		"id":         insight.ID,
	})
	resp.Diagnostics.AddWarning("created PostHog insight", strconv.FormatInt(insight.ID, 10))
	r.setStateFromInsight(&data, insight, request.Query)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Insight) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InsightResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ID.IsNull() || data.ID.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	insight, err := r.client.GetInsight(ctx, data.ID.ValueInt64())
	if err != nil {
		// TODO do we want to deal with not founds?
		//var apiErr *posthog.APIError
		//if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
		//	tflog.Warn(ctx, "insight not found, removing from state", map[string]any{
		//		"project_id": projectID,
		//		"id":         data.ID.ValueInt64(),
		//	})
		//	resp.State.RemoveResource(ctx)
		//	return
		//}

		resp.Diagnostics.AddError("Error reading insight", err.Error())
		return
	}

	r.setStateFromInsight(&data, insight, nil)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Insight) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data InsightResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ID.IsNull() || data.ID.IsUnknown() {
		resp.Diagnostics.AddError("Missing ID", "Resource ID is absent in state")
		return
	}

	request, err := data.ToInsightRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error creating insight request", err.Error())
		return
	}

	insight, err := r.client.UpdateInsight(ctx, data.ID.ValueInt64(), request)
	if err != nil {
		resp.Diagnostics.AddError("Error updating insight", err.Error())
		return
	}

	tflog.Info(ctx, "updated PostHog insight", map[string]any{
		"project_id": r.projectId,
		"id":         data.ID.ValueInt64(),
	})

	r.setStateFromInsight(&data, insight, request.Query)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Insight) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data InsightResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ID.IsNull() || data.ID.IsUnknown() {
		return
	}

	err := r.client.DeleteInsight(ctx, data.ID.ValueInt64())
	if err != nil {
		//var apiErr *posthog.APIError
		//if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
		//	return
		//}
		resp.Diagnostics.AddError("Error deleting insight", err.Error())
	}
}

func (r *Insight) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}

func (r *Insight) setStateFromInsight(model *InsightResourceModel, insight posthog.Insight, fallbackQuery json.RawMessage) {
	model.ID = types.Int64Value(insight.ID)
	model.Name = stringValue(insight.Name)
	model.Description = stringValue(insight.Description)

	query := insight.Query
	if len(query) == 0 {
		query = fallbackQuery
	}

	if len(query) > 0 {
		model.QueryJSON = types.StringValue(string(query))
	} else {
		model.QueryJSON = types.StringNull()
	}
}

func stringValue(value string) types.String {
	if strings.TrimSpace(value) == "" {
		return types.StringNull()
	}

	return types.StringValue(value)
}

func parseQueryJSON(value types.String) (json.RawMessage, error) {
	if value.IsNull() || value.IsUnknown() {
		return nil, fmt.Errorf("query_json must be configured")
	}

	raw := strings.TrimSpace(value.ValueString())
	if raw == "" {
		return nil, fmt.Errorf("query_json cannot be empty")
	}

	if !json.Valid([]byte(raw)) {
		return nil, fmt.Errorf("query_json must contain valid JSON")
	}

	return json.RawMessage(raw), nil
}
