package resource

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

// NewEarlyAccessFeature builds the posthog_early_access_feature resource, a
// feature users opt into from PostHog's in-app feature previews.
//
// Every early access feature is backed by a feature flag, linked once on
// create and never changed afterwards, so feature_flag_id forces replacement.
// Moving to an active stage (alpha, beta, general-availability) makes PostHog
// mark the flag with filters.feature_enrollment, which posthog_feature_flag
// ignores by default.
func NewEarlyAccessFeature() resource.Resource {
	return core.NewGenericResource[EarlyAccessFeatureTFModel, httpclient.EarlyAccessFeatureRequest, httpclient.EarlyAccessFeature](
		EarlyAccessFeatureOps{},
		core.ProjectScopedImportParser[EarlyAccessFeatureTFModel](),
	)
}

var earlyAccessFeatureStages = []string{"draft", "concept", "alpha", "beta", "general-availability", "archived"}

type EarlyAccessFeatureTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name             types.String         `tfsdk:"name"`
	Stage            types.String         `tfsdk:"stage"`
	Description      types.String         `tfsdk:"description"`
	DocumentationURL types.String         `tfsdk:"documentation_url"`
	Payload          jsontypes.Normalized `tfsdk:"payload"`
	FeatureFlagID    types.Int64          `tfsdk:"feature_flag_id"`
	FeatureFlagKey   types.String         `tfsdk:"feature_flag_key"`
	CreatedAt        types.String         `tfsdk:"created_at"`
}

type EarlyAccessFeatureOps struct{}

func (o EarlyAccessFeatureOps) ResourceName() string {
	return "early_access_feature"
}

func (o EarlyAccessFeatureOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "A PostHog early access feature: a feature users can opt into from your app's feature previews list. " +
			"Each one is backed by a feature flag. Moving it to an active stage (`alpha`, `beta`, or `general-availability`) turns on opt-in " +
			"for that flag; any other stage turns it off. `general-availability` remains opt-in gated: PostHog's separate, one-time " +
			"rollout-to-all action is not managed by this resource. Destroying the feature deletes it and turns off opt-in, but leaves the flag in place.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Early access feature ID (UUID).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name shown to users in the opt-in list (at most 200 characters).",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
				},
			},
			"stage": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Lifecycle stage: `draft`, `concept`, `alpha`, `beta`, `general-availability`, or `archived`. " +
					"In `alpha`, `beta`, and `general-availability`, users who opt in get the flag enabled. " +
					"Changing to `general-availability` does not roll the flag out to everyone; that is a separate one-time PostHog action. " +
					"`concept` lists the feature so users can register interest without enabling anything.",
				Validators: []validator.String{
					stringvalidator.OneOf(earlyAccessFeatureStages...),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description shown to users in the opt-in list.",
			},
			"documentation_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Link to documentation for the feature, shown to users in the opt-in list.",
			},
			"payload": schema.StringAttribute{
				CustomType:          jsontypes.NormalizedType{},
				Optional:            true,
				MarkdownDescription: "Arbitrary JSON metadata for the feature, e.g. `jsonencode({ theme = \"dark\" })`. Compared semantically, so key ordering and whitespace do not produce a diff.",
			},
			"feature_flag_id": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "ID of the feature flag backing this feature, typically `posthog_feature_flag.<name>.id`. " +
					"The flag must not be multivariate or group-based, and must not already back another early access feature, survey, or experiment. " +
					"When omitted, PostHog creates a flag keyed by the slugified `name`. That flag is left behind on destroy, so re-creating a feature " +
					"with the same name fails until you delete it; managing the flag with `posthog_feature_flag` avoids this. " +
					"The flag can only be linked on create, so changing it replaces the feature.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplace(),
				},
			},
			"feature_flag_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Key of the feature flag backing this feature.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the feature was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o EarlyAccessFeatureOps) BuildCreateRequest(_ context.Context, model EarlyAccessFeatureTFModel) (httpclient.EarlyAccessFeatureRequest, diag.Diagnostics) {
	req := httpclient.EarlyAccessFeatureRequest{
		Name:  model.Name.ValueString(),
		Stage: model.Stage.ValueString(),
	}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		req.Description = util.StringPtr(model.Description.ValueString())
	}
	if !model.DocumentationURL.IsNull() && !model.DocumentationURL.IsUnknown() {
		req.DocumentationURL = util.StringPtr(model.DocumentationURL.ValueString())
	}
	req.Payload = rawFromNormalized(model.Payload)

	if !model.FeatureFlagID.IsNull() && !model.FeatureFlagID.IsUnknown() {
		flagID := model.FeatureFlagID.ValueInt64()
		req.FeatureFlagID = &flagID
	}

	return req, nil
}

func (o EarlyAccessFeatureOps) BuildUpdateRequest(ctx context.Context, plan, state EarlyAccessFeatureTFModel) (httpclient.EarlyAccessFeatureRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	// PostHog only links the flag on create; changing it replaces the resource.
	req.FeatureFlagID = nil

	// PATCH leaves omitted fields alone, so a field removed from config is cleared explicitly.
	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = util.StringPtr("")
	}
	if core.ShouldClearString(plan.DocumentationURL, state.DocumentationURL) {
		req.DocumentationURL = util.StringPtr("")
	}
	if req.Payload == nil && !state.Payload.IsNull() {
		req.Payload = json.RawMessage(`{}`)
	}

	return req, diags
}

func (o EarlyAccessFeatureOps) MapResponseToModel(_ context.Context, resp httpclient.EarlyAccessFeature, model *EarlyAccessFeatureTFModel) diag.Diagnostics {
	model.ID = types.StringValue(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.Stage = types.StringValue(resp.Stage)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.DocumentationURL = core.PtrToStringNullIfEmptyTrimmed(resp.DocumentationURL)
	model.Payload = earlyAccessFeaturePayloadForState(resp.Payload, model.Payload)
	model.CreatedAt = types.StringValue(resp.CreatedAt)

	if resp.FeatureFlag != nil {
		model.FeatureFlagID = types.Int64Value(resp.FeatureFlag.ID)
		model.FeatureFlagKey = types.StringValue(resp.FeatureFlag.Key)
	} else {
		model.FeatureFlagID = types.Int64Null()
		model.FeatureFlagKey = types.StringNull()
	}

	return nil
}

// earlyAccessFeaturePayloadForState keeps the API's payload as-is so a change made outside
// Terraform shows as drift. PostHog returns {} for "no payload"; that reads as unset unless
// the config itself declares an empty payload.
func earlyAccessFeaturePayloadForState(raw json.RawMessage, current jsontypes.Normalized) jsontypes.Normalized {
	if isEmptyJSONObjectOrNull(raw) {
		if !current.IsNull() && !current.IsUnknown() && isEmptyJSONObjectOrNull(json.RawMessage(current.ValueString())) {
			return current
		}
		return jsontypes.NewNormalizedNull()
	}
	var compacted bytes.Buffer
	if err := json.Compact(&compacted, raw); err != nil {
		return jsontypes.NewNormalizedValue(string(raw))
	}
	return jsontypes.NewNormalizedValue(compacted.String())
}

func isEmptyJSONObjectOrNull(raw json.RawMessage) bool {
	trimmed := strings.TrimSpace(string(raw))
	if trimmed == "" || trimmed == "null" {
		return true
	}
	var obj map[string]interface{}
	return json.Unmarshal(raw, &obj) == nil && len(obj) == 0
}

func (o EarlyAccessFeatureOps) Create(ctx context.Context, client httpclient.PosthogClient, model EarlyAccessFeatureTFModel, req httpclient.EarlyAccessFeatureRequest) (httpclient.EarlyAccessFeature, error) {
	return client.CreateEarlyAccessFeature(ctx, model.GetEffectiveProjectID(), req)
}

func (o EarlyAccessFeatureOps) Read(ctx context.Context, client httpclient.PosthogClient, model EarlyAccessFeatureTFModel) (httpclient.EarlyAccessFeature, httpclient.HTTPStatusCode, error) {
	return client.GetEarlyAccessFeature(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o EarlyAccessFeatureOps) Update(ctx context.Context, client httpclient.PosthogClient, model EarlyAccessFeatureTFModel, req httpclient.EarlyAccessFeatureRequest) (httpclient.EarlyAccessFeature, httpclient.HTTPStatusCode, error) {
	return client.UpdateEarlyAccessFeature(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o EarlyAccessFeatureOps) Delete(ctx context.Context, client httpclient.PosthogClient, model EarlyAccessFeatureTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteEarlyAccessFeature(ctx, model.GetEffectiveProjectID(), model.GetID())
}
