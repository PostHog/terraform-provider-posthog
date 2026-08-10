package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

func NewInsight() datasource.DataSource {
	return &InsightDataSource{}
}

type InsightDataSource struct {
	client           httpclient.PosthogClient
	defaultProjectID string
}

type InsightDataSourceModel struct {
	ProjectID    types.String `tfsdk:"project_id"`
	ShortID      types.String `tfsdk:"short_id"`
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	DerivedName  types.String `tfsdk:"derived_name"`
	Description  types.String `tfsdk:"description"`
	QueryJSON    types.String `tfsdk:"query_json"`
	Tags         types.Set    `tfsdk:"tags"`
	DashboardIDs types.Set    `tfsdk:"dashboard_ids"`
	Favorited    types.Bool   `tfsdk:"favorited"`
}

var (
	_ datasource.DataSource              = (*InsightDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*InsightDataSource)(nil)
)

func (d *InsightDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_insight", req.ProviderTypeName)
}

func (d *InsightDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Look up an existing insight by its `short_id` (the value in the insight's URL). " +
			"Use this to reference insights created outside Terraform, or to resolve the numeric `id` needed to adopt an " +
			"existing insight into a `posthog_insight` resource via a one-time `import {}` block.",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project ID (environment) to look up the insight in. Defaults to the provider-level project_id.",
			},
			"short_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The short ID of the insight to look up. This is the value shown in the insight's URL.",
			},
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Numeric ID of the insight. Use this as the `id` in an `import {}` block to adopt the insight into a `posthog_insight` resource.",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Insight name.",
			},
			"derived_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Auto-generated insight name (set by PostHog when the name is empty).",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Insight description.",
			},
			"query_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The insight query as a normalized JSON string, with server-injected fields removed.",
			},
			"tags": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Tags applied to the insight.",
			},
			"dashboard_ids": schema.SetAttribute{
				Computed:            true,
				ElementType:         types.Int32Type,
				MarkdownDescription: "IDs of the dashboards the insight appears on.",
			},
			"favorited": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the insight is marked as favorited.",
			},
		},
	}
}

func (d *InsightDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(data.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = providerData.Client
	d.defaultProjectID = providerData.DefaultProjectID
}

func (d *InsightDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config InsightDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use provider default if project_id not explicitly set
	projectID := config.ProjectID.ValueString()
	if config.ProjectID.IsNull() || config.ProjectID.IsUnknown() || projectID == "" {
		projectID = d.defaultProjectID
		config.ProjectID = types.StringValue(projectID)
	}

	if projectID == "" {
		resp.Diagnostics.AddError(
			"Missing project_id",
			"project_id must be set either in the data source or at the provider level",
		)
		return
	}

	shortID := strings.TrimSpace(config.ShortID.ValueString())

	tflog.Debug(ctx, "Looking up insight by short_id", map[string]any{
		"project_id": projectID,
		"short_id":   shortID,
	})

	// GetInsight's detail path resolves a short_id as well as a numeric id
	// (verified against the API and by the acceptance test).
	insight, statusCode, err := d.client.GetInsight(ctx, projectID, shortID)
	if err != nil {
		if statusCode == http.StatusNotFound {
			resp.Diagnostics.AddError(
				"Insight not found",
				fmt.Sprintf("No insight with short_id %q found in project %s", shortID, projectID),
			)
			return
		}
		resp.Diagnostics.AddError("Error reading insight", err.Error())
		return
	}

	// A soft-deleted insight is treated as not found — a data source should not
	// resolve a deleted insight.
	if insight.Deleted != nil && *insight.Deleted {
		resp.Diagnostics.AddError(
			"Insight not found",
			fmt.Sprintf("No insight with short_id %q found in project %s", shortID, projectID),
		)
		return
	}

	resp.Diagnostics.Append(mapInsightToState(ctx, insight, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// mapInsightToState maps an API insight response onto the data source model.
func mapInsightToState(ctx context.Context, insight httpclient.Insight, model *InsightDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(insight.ID)
	model.ShortID = core.PtrToStringNullIfEmptyTrimmed(insight.ShortID)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(insight.Name)
	model.DerivedName = core.PtrToStringNullIfEmptyTrimmed(insight.DerivedName)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(insight.Description)
	model.Favorited = core.PtrToBool(insight.Favorited)

	tags, d := core.TagsToSet(ctx, insight.Tags)
	diags.Append(d...)
	model.Tags = tags

	dashboards, d := core.Int32SetPreserveEmpty(ctx, insight.Dashboards, types.SetNull(types.Int32Type))
	diags.Append(d...)
	model.DashboardIDs = dashboards

	if insight.Query != nil {
		cleaned := util.StripFields(insight.Query, httpclient.InsightQueryServerFields)
		queryJSON, err := json.Marshal(cleaned)
		if err != nil {
			diags.AddError("Failed to normalize query", err.Error())
			return diags
		}
		model.QueryJSON = types.StringValue(string(queryJSON))
	} else {
		model.QueryJSON = types.StringNull()
	}

	return diags
}
