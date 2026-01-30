package resource

import (
	"context"
	"encoding/json"
	"fmt"

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
	Key               types.String `tfsdk:"key"`
	Name              types.String `tfsdk:"name"`
	Active            types.Bool   `tfsdk:"active"`
	Filters           types.String `tfsdk:"filters"`
	RolloutPercentage types.Int64  `tfsdk:"rollout_percentage"`
	Tags              types.Set    `tfsdk:"tags"`
	Deleted           types.Bool   `tfsdk:"deleted"`
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
				MarkdownDescription: "Feature flag name/description",
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
			"filters": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Feature flag filters as JSON",
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

	// Extract rollout_percentage from filters.groups[0].rollout_percentage if it exists
	// The API doesn't return a top-level rollout_percentage field
	var rolloutPercentage *int32
	if len(resp.Filters) > 0 {
		if groups, ok := resp.Filters["groups"].([]interface{}); ok && len(groups) > 0 {
			if firstGroup, ok := groups[0].(map[string]interface{}); ok {
				if rp, ok := firstGroup["rollout_percentage"].(float64); ok {
					percentage := int32(rp)
					rolloutPercentage = &percentage
				}
			}
		}
	}

	// Set filters if present
	if len(resp.Filters) > 0 {
		filtersJSON, err := json.Marshal(resp.Filters)
		if err == nil {
			model.Filters = types.StringValue(string(filtersJSON))
		}
	} else {
		model.Filters = types.StringNull()
	}

	// Set rollout_percentage if extracted from filters
	if rolloutPercentage != nil {
		model.RolloutPercentage = types.Int64Value(int64(*rolloutPercentage))
	} else {
		model.RolloutPercentage = types.Int64Null()
	}

	// Set tags
	tagsSet, d := core.TagsToSet(ctx, resp.Tags)
	diags.Append(d...)
	model.Tags = tagsSet

	// Set deleted status - treat nil as false to avoid perpetual diffs
	deleted := resp.Deleted != nil && *resp.Deleted
	model.Deleted = types.BoolValue(deleted)

	return diags
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
