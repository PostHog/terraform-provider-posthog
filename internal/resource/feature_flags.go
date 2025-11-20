package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/posthog"
)

var (
	_ resource.Resource                = &FeatureFlagResource{}
	_ resource.ResourceWithImportState = &FeatureFlagResource{}
)

func NewFeatureFlag() resource.Resource {
	return &FeatureFlagResource{}
}

type FeatureFlagResource struct {
	client posthog.Client
}

type FeatureFlagResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	Key               types.String `tfsdk:"key"`
	Name              types.String `tfsdk:"name"`
	Active            types.Bool   `tfsdk:"active"`
	Filters           types.String `tfsdk:"filters"`
	RolloutPercentage types.Int64  `tfsdk:"rollout_percentage"`
	Tags              types.Set    `tfsdk:"tags"`
}

func (r *FeatureFlagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feature_flag"
}

func (r *FeatureFlagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PostHog Feature Flag resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Feature Flag ID",
			},
			"key": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Feature flag key (unique identifier)",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Feature flag name/description",
			},
			"active": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Whether the feature flag is active",
			},
			"filters": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Feature flag filters as JSON",
			},
			"rollout_percentage": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Rollout percentage (0-100)",
			},
			"tags": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Set of tags for the feature flag",
			},
		},
	}
}

func (r *FeatureFlagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = providerData.Client
}

func (r *FeatureFlagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FeatureFlagResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build feature flag request
	input := posthog.FeatureFlagRequest{
		Key: data.Key.ValueString(),
	}

	if !data.Name.IsNull() {
		name := data.Name.ValueString()
		input.Name = &name
	}

	if !data.Active.IsNull() {
		active := data.Active.ValueBool()
		input.Active = &active
	}

	// Handle filters and rollout_percentage
	// rollout_percentage is a convenience field that goes into filters.groups[0].rollout_percentage
	var filters map[string]interface{}

	if !data.Filters.IsNull() {
		if err := json.Unmarshal([]byte(data.Filters.ValueString()), &filters); err != nil {
			resp.Diagnostics.AddError(
				"Invalid filters JSON",
				fmt.Sprintf("Could not parse filters: %s", err.Error()),
			)
			return
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
	if !data.RolloutPercentage.IsNull() {
		percentage := int32(data.RolloutPercentage.ValueInt64())
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

	input.Filters = filters

	if !data.Tags.IsNull() {
		var tags []string
		resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &tags, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		input.Tags = tags
	}

	tflog.Debug(ctx, "Creating PostHog feature flag", map[string]interface{}{
		"key": data.Key.ValueString(),
	})

	created, err := r.client.CreateFeatureFlag(ctx, input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating feature flag",
			fmt.Sprintf("Could not create feature flag: %s", err.Error()),
		)
		return
	}

	data.ID = types.Int64Value(created.ID)
	data.Key = types.StringValue(created.Key)

	if created.Name != nil {
		data.Name = types.StringValue(*created.Name)
	}

	if created.Active != nil {
		data.Active = types.BoolValue(*created.Active)
	}

	// Only set filters if it was provided in plan
	if !data.Filters.IsNull() && created.Filters != nil {
		filtersJSON, err := json.Marshal(created.Filters)
		if err == nil {
			data.Filters = types.StringValue(string(filtersJSON))
		}
	}

	// Only set rollout_percentage if it was provided in plan
	if !data.RolloutPercentage.IsNull() && created.RolloutPercentage != nil {
		data.RolloutPercentage = types.Int64Value(int64(*created.RolloutPercentage))
	}

	if len(created.Tags) > 0 {
		tagsSet, diags := types.SetValueFrom(ctx, types.StringType, created.Tags)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			data.Tags = tagsSet
		}
	} else {
		data.Tags = types.SetNull(types.StringType)
	}

	tflog.Debug(ctx, "Created PostHog feature flag", map[string]interface{}{
		"id": data.ID.ValueInt64(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeatureFlagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FeatureFlagResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Track what fields were originally configured
	hasFiltersInState := !data.Filters.IsNull()
	hasRolloutInState := !data.RolloutPercentage.IsNull()
	hasTagsInState := !data.Tags.IsNull()

	tflog.Debug(ctx, "Reading PostHog feature flag", map[string]interface{}{
		"id": data.ID.ValueInt64(),
	})

	flag, err := r.client.GetFeatureFlag(ctx, data.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading feature flag",
			fmt.Sprintf("Could not read feature flag %d: %s", data.ID.ValueInt64(), err.Error()),
		)
		return
	}

	data.ID = types.Int64Value(flag.ID)
	data.Key = types.StringValue(flag.Key)

	if flag.Name != nil {
		data.Name = types.StringValue(*flag.Name)
	} else {
		data.Name = types.StringNull()
	}

	if flag.Active != nil {
		data.Active = types.BoolValue(*flag.Active)
	} else {
		data.Active = types.BoolNull()
	}

	// Only set filters if it was originally in state
	if hasFiltersInState {
		if flag.Filters != nil && len(flag.Filters) > 0 {
			filtersJSON, err := json.Marshal(flag.Filters)
			if err == nil {
				data.Filters = types.StringValue(string(filtersJSON))
			}
		} else {
			data.Filters = types.StringNull()
		}
	} else {
		data.Filters = types.StringNull()
	}

	// Only set rollout_percentage if it was originally in state
	if hasRolloutInState {
		if flag.RolloutPercentage != nil {
			data.RolloutPercentage = types.Int64Value(int64(*flag.RolloutPercentage))
		} else {
			data.RolloutPercentage = types.Int64Null()
		}
	} else {
		data.RolloutPercentage = types.Int64Null()
	}

	// Only set tags if they were originally in state
	if hasTagsInState {
		if len(flag.Tags) > 0 {
			tagsSet, diags := types.SetValueFrom(ctx, types.StringType, flag.Tags)
			resp.Diagnostics.Append(diags...)
			if !resp.Diagnostics.HasError() {
				data.Tags = tagsSet
			}
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
	} else {
		data.Tags = types.SetNull(types.StringType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeatureFlagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FeatureFlagResourceModel
	var state FeatureFlagResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use the ID from state, but other values from plan
	input := posthog.FeatureFlagRequest{
		Key: plan.Key.ValueString(),
	}

	if !plan.Name.IsNull() {
		name := plan.Name.ValueString()
		input.Name = &name
	}

	if !plan.Active.IsNull() {
		active := plan.Active.ValueBool()
		input.Active = &active
	}

	// Handle filters and rollout_percentage (same logic as Create)
	var filters map[string]interface{}

	if !plan.Filters.IsNull() {
		if err := json.Unmarshal([]byte(plan.Filters.ValueString()), &filters); err != nil {
			resp.Diagnostics.AddError(
				"Invalid filters JSON",
				fmt.Sprintf("Could not parse filters: %s", err.Error()),
			)
			return
		}
	} else if !plan.RolloutPercentage.IsNull() {
		// If only rollout_percentage is provided, create default filters structure
		filters = map[string]interface{}{
			"groups": []interface{}{
				map[string]interface{}{},
			},
		}
	}

	// If rollout_percentage is provided, add it to the first group in filters
	if !plan.RolloutPercentage.IsNull() && filters != nil {
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
		input.Filters = filters
	}

	if !plan.Tags.IsNull() {
		var tags []string
		resp.Diagnostics.Append(plan.Tags.ElementsAs(ctx, &tags, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		input.Tags = tags
	}

	tflog.Debug(ctx, "Updating PostHog feature flag", map[string]interface{}{
		"id": state.ID.ValueInt64(),
	})

	updated, err := r.client.UpdateFeatureFlag(ctx, state.ID.ValueInt64(), input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating feature flag",
			fmt.Sprintf("Could not update feature flag %d: %s", state.ID.ValueInt64(), err.Error()),
		)
		return
	}

	// Build the new state from the updated response
	// Only include attributes that were in the plan to avoid "inconsistent result" errors
	var newState FeatureFlagResourceModel
	newState.ID = types.Int64Value(updated.ID)
	newState.Key = types.StringValue(updated.Key)

	if updated.Name != nil {
		newState.Name = types.StringValue(*updated.Name)
	}

	if updated.Active != nil {
		newState.Active = types.BoolValue(*updated.Active)
	}

	// Only set filters if it was in the plan
	if !plan.Filters.IsNull() && updated.Filters != nil && len(updated.Filters) > 0 {
		filtersJSON, err := json.Marshal(updated.Filters)
		if err == nil {
			newState.Filters = types.StringValue(string(filtersJSON))
		}
	}

	// Only set rollout_percentage if it was in the plan
	if !plan.RolloutPercentage.IsNull() && updated.RolloutPercentage != nil {
		newState.RolloutPercentage = types.Int64Value(int64(*updated.RolloutPercentage))
	}

	// Only set tags if they were in the plan
	if !plan.Tags.IsNull() {
		if len(updated.Tags) > 0 {
			tags := make([]string, len(updated.Tags))
			for i, tag := range updated.Tags {
				tags[i] = tag
			}
			tagsSet, diags := types.SetValueFrom(ctx, types.StringType, tags)
			resp.Diagnostics.Append(diags...)
			if !resp.Diagnostics.HasError() {
				newState.Tags = tagsSet
			}
		} else {
			// Empty set if plan had tags but API returned none
			newState.Tags, _ = types.SetValueFrom(ctx, types.StringType, []string{})
		}
	} else {
		// Tags not in plan - preserve from state or set null with proper type
		if !state.Tags.IsNull() {
			newState.Tags = state.Tags
		} else {
			newState.Tags = types.SetNull(types.StringType)
		}
	}

	tflog.Debug(ctx, "Updated PostHog feature flag", map[string]interface{}{
		"id": newState.ID.ValueInt64(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *FeatureFlagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FeatureFlagResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting PostHog feature flag (soft delete)", map[string]interface{}{
		"id": data.ID.ValueInt64(),
	})

	err := r.client.DeleteFeatureFlag(ctx, data.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting feature flag",
			fmt.Sprintf("Could not delete feature flag %d: %s", data.ID.ValueInt64(), err.Error()),
		)
		return
	}

	tflog.Debug(ctx, "Deleted PostHog feature flag", map[string]interface{}{
		"id": data.ID.ValueInt64(),
	})
}

func (r *FeatureFlagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Expected numeric feature flag ID, got %q: %v", req.ID, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
