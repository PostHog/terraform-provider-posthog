package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/posthog"
)

var (
	_ resource.Resource                = &DashboardResource{}
	_ resource.ResourceWithImportState = &DashboardResource{}
)

func NewDashboardResource() resource.Resource {
	return &DashboardResource{}
}

type DashboardResource struct {
	client posthog.Client
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

	providerData, ok := req.ProviderData.(ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = providerData.Client
}

func (r *DashboardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DashboardResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := posthog.CreateDashboardParams{
		Name: data.Name.ValueString(),
	}

	if !data.Description.IsNull() {
		desc := data.Description.ValueString()
		params.Description = &desc
	}

	if !data.Pinned.IsNull() {
		pinned := data.Pinned.ValueBool()
		params.Pinned = &pinned
	}

	if !data.Tags.IsNull() {
		var tags []types.String
		resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &tags, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tagStrings := make([]string, len(tags))
		for i, tag := range tags {
			tagStrings[i] = tag.ValueString()
		}
		params.Tags = tagStrings
	}

	tflog.Debug(ctx, "Creating PostHog dashboard", map[string]interface{}{
		"name": data.Name.ValueString(),
	})

	dashboard, err := r.client.CreateDashboard(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating dashboard",
			fmt.Sprintf("Could not create dashboard: %s", err.Error()),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%d", dashboard.ID))
	if dashboard.Name != nil {
		data.Name = types.StringValue(*dashboard.Name)
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

	if len(dashboard.Tags) > 0 {
		tagTypes := make([]types.String, len(dashboard.Tags))
		for i, tag := range dashboard.Tags {
			tagTypes[i] = types.StringValue(tag)
		}
		tagsSet, diags := types.SetValueFrom(ctx, types.StringType, tagTypes)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Tags = tagsSet
		}
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

	dashboard, err := r.client.GetDashboard(data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading dashboard",
			fmt.Sprintf("Could not read dashboard %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%d", dashboard.ID))
	if dashboard.Name != nil {
		data.Name = types.StringValue(*dashboard.Name)
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
			tagTypes[i] = types.StringValue(tag)
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

	params := posthog.UpdateDashboardParams{}

	if !data.Name.IsNull() {
		name := data.Name.ValueString()
		params.Name = &name
	}

	if !data.Description.IsNull() {
		desc := data.Description.ValueString()
		params.Description = &desc
	}

	if !data.Pinned.IsNull() {
		pinned := data.Pinned.ValueBool()
		params.Pinned = &pinned
	}

	if !data.Tags.IsNull() {
		var tags []types.String
		resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &tags, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tagStrings := make([]string, len(tags))
		for i, tag := range tags {
			tagStrings[i] = tag.ValueString()
		}
		params.Tags = tagStrings
	}

	if !data.Deleted.IsNull() {
		deleted := data.Deleted.ValueBool()
		params.Deleted = &deleted
	}

	tflog.Debug(ctx, "Updating PostHog dashboard", map[string]interface{}{
		"id": data.ID.ValueString(),
	})

	dashboard, err := r.client.UpdateDashboard(data.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating dashboard",
			fmt.Sprintf("Could not update dashboard %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%d", dashboard.ID))
	if dashboard.Name != nil {
		data.Name = types.StringValue(*dashboard.Name)
	}
	if dashboard.Description != nil {
		data.Description = types.StringValue(*dashboard.Description)
	}
	if dashboard.Pinned != nil {
		data.Pinned = types.BoolValue(*dashboard.Pinned)
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

	err := r.client.DeleteDashboard(data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting dashboard",
			fmt.Sprintf("Could not delete dashboard %s: %s", data.ID.ValueString(), err.Error()),
		)
		return
	}

	tflog.Debug(ctx, "Deleted PostHog dashboard", map[string]interface{}{
		"id": data.ID.ValueString(),
	})
}

func (r *DashboardResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
