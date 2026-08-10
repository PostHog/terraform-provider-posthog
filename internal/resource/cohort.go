package resource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
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
	"github.com/posthog/terraform-provider/internal/util"
)

func NewCohort() resource.Resource {
	return core.NewGenericResource[CohortTFModel, httpclient.CohortRequest, httpclient.Cohort](
		CohortOps{},
		core.ProjectScopedImportParser[CohortTFModel](),
	)
}

type CohortTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	Name        types.String         `tfsdk:"name"`
	Description types.String         `tfsdk:"description"`
	Filters     jsontypes.Normalized `tfsdk:"filters"`
	IsStatic    types.Bool           `tfsdk:"is_static"`
	CreatedAt   types.String         `tfsdk:"created_at"`
	Deleted     types.Bool           `tfsdk:"deleted"`
}

type CohortOps struct{}

func (o CohortOps) ResourceName() string {
	return "Cohort"
}

func (o CohortOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "PostHog Cohort resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Cohort ID",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Cohort name",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Cohort description",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"filters": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
				Computed:   true,
				MarkdownDescription: "Cohort membership rules as JSON, matching the `filters` object of the " +
					"[cohorts API](https://posthog.com/docs/api/cohorts). Compared semantically, so key ordering and " +
					"whitespace differences from the PostHog API do not produce a diff. Only the fields declared here " +
					"are tracked: PostHog annotates saved cohorts with computed fields (`bytecode`, `conditionHash`) " +
					"that would otherwise surface as perpetual drift.",
			},
			"is_static": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Whether the cohort is static (a fixed list of persons) rather than dynamic " +
					"(recomputed by PostHog from `filters`). Defaults to false. Membership of a static cohort is not " +
					"managed by this resource.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the cohort was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"deleted": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				MarkdownDescription: "Whether the cohort is soft-deleted. PostHog does not permit hard deletion of " +
					"cohorts, so destroying this resource marks it deleted. Terraform will restore soft-deleted " +
					"cohorts on apply.",
				PlanModifiers: []planmodifier.Bool{
					core.DefaultBoolFalse{},
				},
			},
		},
	}
}

func (o CohortOps) BuildCreateRequest(_ context.Context, model CohortTFModel) (httpclient.CohortRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := model.Name.ValueString()
	req := httpclient.CohortRequest{
		Name: &name,
	}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}

	if !model.IsStatic.IsNull() && !model.IsStatic.IsUnknown() {
		isStatic := model.IsStatic.ValueBool()
		req.IsStatic = &isStatic
	}

	if !model.Filters.IsNull() && !model.Filters.IsUnknown() {
		var filters map[string]interface{}
		if err := json.Unmarshal([]byte(model.Filters.ValueString()), &filters); err != nil {
			diags.AddError("Invalid filters JSON", fmt.Sprintf("Could not parse filters: %s", err.Error()))
			return req, diags
		}
		req.Filters = filters
	}

	deleted := false
	req.Deleted = &deleted

	return req, diags
}

func (o CohortOps) BuildUpdateRequest(ctx context.Context, plan, state CohortTFModel) (httpclient.CohortRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = util.StringPtr("")
	}

	deleted := plan.Deleted.ValueBool()
	req.Deleted = &deleted

	return req, diags
}

func (o CohortOps) MapResponseToModel(_ context.Context, resp httpclient.Cohort, model *CohortTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)

	if len(resp.Filters) > 0 {
		normalized, err := normalizeJSONForState(resp.Filters, model.Filters.ValueString())
		if err != nil {
			diags.AddError("Failed to normalize filters", err.Error())
			return diags
		}
		model.Filters = jsontypes.NewNormalizedValue(normalized)
	} else {
		model.Filters = jsontypes.NewNormalizedNull()
	}

	// Treat a missing is_static/deleted as false so an omitted key cannot leave
	// an Optional+Computed attribute unknown after apply.
	model.IsStatic = types.BoolValue(resp.IsStatic != nil && *resp.IsStatic)
	model.Deleted = types.BoolValue(resp.Deleted != nil && *resp.Deleted)

	return diags
}

func (o CohortOps) Create(ctx context.Context, client httpclient.PosthogClient, model CohortTFModel, req httpclient.CohortRequest) (httpclient.Cohort, error) {
	return client.CreateCohort(ctx, model.GetEffectiveProjectID(), req)
}

func (o CohortOps) Read(ctx context.Context, client httpclient.PosthogClient, model CohortTFModel) (httpclient.Cohort, httpclient.HTTPStatusCode, error) {
	return client.GetCohort(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o CohortOps) Update(ctx context.Context, client httpclient.PosthogClient, model CohortTFModel, req httpclient.CohortRequest) (httpclient.Cohort, httpclient.HTTPStatusCode, error) {
	return client.UpdateCohort(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o CohortOps) Delete(ctx context.Context, client httpclient.PosthogClient, model CohortTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteCohort(ctx, model.GetEffectiveProjectID(), model.GetID())
}
