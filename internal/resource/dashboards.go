package resource

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/data"
	posthogapi "github.com/posthog/terraform-provider/internal/posthog/swagger"
)

var (
	_ resource.Resource                = &DashboardResource{}
	_ resource.ResourceWithImportState = &DashboardResource{}
)

func NewDashboard() resource.Resource {
	return &DashboardResource{}
}

type DashboardResource struct {
	client    *posthogapi.APIClient
	projectID string
}

type DashboardResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Pinned      types.Bool   `tfsdk:"pinned"`
	Tags        types.Set    `tfsdk:"tags"`
	Deleted     types.Bool   `tfsdk:"deleted"`
}

func (r *DashboardResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dashboard"
}

func (r *DashboardResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PostHog Dashboard resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Dashboard ID",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Dashboard name",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Dashboard description",
			},
			"pinned": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Whether the dashboard is pinned",
			},
			"tags": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Set of tags for the dashboard",
			},
			"deleted": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Whether the dashboard is deleted (soft delete)",
			},
		},
	}
}

func (r *DashboardResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(data.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = providerData.SwaggerClient
	r.projectID = providerData.ProjectID
}

func (r *DashboardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DashboardResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dashboard := posthogapi.NewDashboardWithDefaults()
	dashboard.SetName(data.Name.ValueString())

	if !data.Description.IsNull() {
		desc := data.Description.ValueString()
		dashboard.SetDescription(desc)
	}

	if !data.Pinned.IsNull() {
		pinned := data.Pinned.ValueBool()
		dashboard.SetPinned(pinned)
	}

	if !data.Tags.IsNull() {
		var tags []types.String
		resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &tags, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tagInterfaces := make([]interface{}, len(tags))
		for i, tag := range tags {
			tagInterfaces[i] = tag.ValueString()
		}
		dashboard.SetTags(tagInterfaces)
	}

	tflog.Debug(ctx, "Creating PostHog dashboard", map[string]interface{}{
		"name": data.Name.ValueString(),
	})

	apiReq := r.client.DashboardsAPI.DashboardsCreate(ctx, r.projectID)
	apiReq = apiReq.Dashboard(*dashboard)

	created, httpResp, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating dashboard",
			fmt.Sprintf("Could not create dashboard: %s", err.Error()),
		)
		return
	}
	defer httpResp.Body.Close()

	data.ID = types.StringValue(strconv.Itoa(int(created.Id)))

	if created.Name.IsSet() {
		name := created.Name.Get()
		data.Name = types.StringValue(*name)
	}

	// Only set description if it was provided in plan or if API returns non-empty value
	if created.Description != nil && *created.Description != "" {
		data.Description = types.StringValue(*created.Description)
	} else if data.Description.IsNull() {
		data.Description = types.StringNull()
	}

	// Only set pinned if it was provided in plan or if API returns true
	if created.Pinned != nil && *created.Pinned {
		data.Pinned = types.BoolValue(*created.Pinned)
	} else if data.Pinned.IsNull() {
		data.Pinned = types.BoolNull()
	}

	// Only set tags if there are actual tags
	if len(created.Tags) > 0 {
		tagTypes := make([]types.String, len(created.Tags))
		for i, tag := range created.Tags {
			if tagStr, ok := tag.(string); ok {
				tagTypes[i] = types.StringValue(tagStr)
			}
		}
		tagsSet, diags := types.SetValueFrom(ctx, types.StringType, tagTypes)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Tags = tagsSet
		}
	} else if data.Tags.IsNull() {
		data.Tags = types.SetNull(types.StringType)
	}

	tflog.Debug(ctx, "Created PostHog dashboard", map[string]interface{}{
		"id": data.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DashboardResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DashboardResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading PostHog dashboard", map[string]interface{}{
		"id": data.ID.ValueString(),
	})

	dashboardID, err := strconv.ParseInt(data.ID.ValueString(), 10, 32)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid dashboard ID",
			fmt.Sprintf("Could not parse dashboard ID %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}

	apiReq := r.client.DashboardsAPI.DashboardsRetrieve(ctx, int32(dashboardID), r.projectID)
	dashboard, httpResp, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading dashboard",
			fmt.Sprintf("Could not read dashboard %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}
	defer httpResp.Body.Close()

	data.ID = types.StringValue(strconv.Itoa(int(dashboard.Id)))

	if dashboard.Name.IsSet() {
		name := dashboard.Name.Get()
		data.Name = types.StringValue(*name)
	} else {
		data.Name = types.StringNull()
	}

	if dashboard.Description != nil {
		data.Description = types.StringValue(*dashboard.Description)
	} else {
		data.Description = types.StringNull()
	}

	if dashboard.Pinned != nil {
		data.Pinned = types.BoolValue(*dashboard.Pinned)
	} else {
		data.Pinned = types.BoolNull()
	}

	if dashboard.Deleted != nil {
		data.Deleted = types.BoolValue(*dashboard.Deleted)
	} else {
		data.Deleted = types.BoolNull()
	}

	if len(dashboard.Tags) > 0 {
		tagTypes := make([]types.String, len(dashboard.Tags))
		for i, tag := range dashboard.Tags {
			if tagStr, ok := tag.(string); ok {
				tagTypes[i] = types.StringValue(tagStr)
			}
		}
		tagsSet, diags := types.SetValueFrom(ctx, types.StringType, tagTypes)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Tags = tagsSet
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DashboardResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DashboardResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dashboardID, err := strconv.ParseInt(data.ID.ValueString(), 10, 32)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid dashboard ID",
			fmt.Sprintf("Could not parse dashboard ID %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}

	patchedDashboard := posthogapi.NewPatchedDashboard()

	if !data.Name.IsNull() {
		name := data.Name.ValueString()
		patchedDashboard.SetName(name)
	}

	if !data.Description.IsNull() {
		desc := data.Description.ValueString()
		patchedDashboard.SetDescription(desc)
	}

	if !data.Pinned.IsNull() {
		pinned := data.Pinned.ValueBool()
		patchedDashboard.SetPinned(pinned)
	}

	if !data.Tags.IsNull() {
		var tags []types.String
		resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &tags, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tagInterfaces := make([]interface{}, len(tags))
		for i, tag := range tags {
			tagInterfaces[i] = tag.ValueString()
		}
		patchedDashboard.SetTags(tagInterfaces)
	}

	if !data.Deleted.IsNull() {
		deleted := data.Deleted.ValueBool()
		patchedDashboard.SetDeleted(deleted)
	}

	tflog.Debug(ctx, "Updating PostHog dashboard", map[string]interface{}{
		"id": data.ID.ValueString(),
	})

	apiReq := r.client.DashboardsAPI.DashboardsPartialUpdate(ctx, int32(dashboardID), r.projectID)
	apiReq = apiReq.PatchedDashboard(*patchedDashboard)

	updated, httpResp, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating dashboard",
			fmt.Sprintf("Could not update dashboard %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}
	defer httpResp.Body.Close()

	data.ID = types.StringValue(strconv.Itoa(int(updated.Id)))

	if updated.Name.IsSet() {
		name := updated.Name.Get()
		data.Name = types.StringValue(*name)
	}

	if updated.Description != nil {
		data.Description = types.StringValue(*updated.Description)
	}

	if updated.Pinned != nil {
		data.Pinned = types.BoolValue(*updated.Pinned)
	}

	tflog.Debug(ctx, "Updated PostHog dashboard", map[string]interface{}{
		"id": data.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DashboardResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DashboardResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting PostHog dashboard (soft delete)", map[string]interface{}{
		"id": data.ID.ValueString(),
	})

	dashboardID, err := strconv.ParseInt(data.ID.ValueString(), 10, 32)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid dashboard ID",
			fmt.Sprintf("Could not parse dashboard ID %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}

	// Soft delete by updating with deleted=true
	patchedDashboard := posthogapi.NewPatchedDashboard()
	patchedDashboard.SetDeleted(true)

	apiReq := r.client.DashboardsAPI.DashboardsPartialUpdate(ctx, int32(dashboardID), r.projectID)
	apiReq = apiReq.PatchedDashboard(*patchedDashboard)

	_, httpResp, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting dashboard",
			fmt.Sprintf("Could not delete dashboard %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}
	defer httpResp.Body.Close()

	tflog.Debug(ctx, "Deleted PostHog dashboard", map[string]interface{}{
		"id": data.ID.ValueString(),
	})
}

func (r *DashboardResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
