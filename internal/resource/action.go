package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

func NewAction() resource.Resource {
	return core.NewGenericResource[ActionTFModel, httpclient.ActionRequest, httpclient.Action](
		ActionOps{},
		core.ProjectScopedImportParser[ActionTFModel](),
	)
}

type ActionTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	Name               types.String `tfsdk:"name"`
	Description        types.String `tfsdk:"description"`
	Tags               types.Set    `tfsdk:"tags"`
	PostToSlack        types.Bool   `tfsdk:"post_to_slack"`
	SlackMessageFormat types.String `tfsdk:"slack_message_format"`
	StepsJSON          types.String `tfsdk:"steps_json"`
	Deleted            types.Bool   `tfsdk:"deleted"`
}

type ActionOps struct{}

func (o ActionOps) ResourceName() string {
	return "Action"
}

func (o ActionOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "PostHog Action resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Action ID",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Action name",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Action description",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Set of tags for the action",
			},
			"post_to_slack": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether to post to Slack when the action is triggered",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"slack_message_format": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Slack message format string",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"steps_json": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "JSON-encoded array of action step objects",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"deleted": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the action is soft-deleted. Terraform will restore soft-deleted actions on apply.",
				PlanModifiers: []planmodifier.Bool{
					core.DefaultBoolFalse{},
				},
			},
		},
	}
}

func (o ActionOps) BuildCreateRequest(ctx context.Context, model ActionTFModel) (httpclient.ActionRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.ActionRequest{
		Name: model.Name.ValueString(),
	}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}

	if !model.PostToSlack.IsNull() && !model.PostToSlack.IsUnknown() {
		v := model.PostToSlack.ValueBool()
		req.PostToSlack = &v
	}

	if !model.SlackMessageFormat.IsNull() && !model.SlackMessageFormat.IsUnknown() {
		v := model.SlackMessageFormat.ValueString()
		req.SlackMessageFormat = &v
	}

	if !model.StepsJSON.IsNull() && !model.StepsJSON.IsUnknown() {
		raw := strings.TrimSpace(model.StepsJSON.ValueString())
		if raw != "" {
			var steps []map[string]interface{}
			if err := json.Unmarshal([]byte(raw), &steps); err != nil {
				diags.AddError("Invalid steps_json", fmt.Sprintf("steps_json must be a valid JSON array: %s", err.Error()))
				return req, diags
			}
			req.Steps = steps
		}
	}

	if !model.Tags.IsNull() {
		tags, d := core.ExtractTags(ctx, model.Tags)
		diags.Append(d...)
		req.Tags = tags
	}

	deleted := false
	req.Deleted = &deleted

	return req, diags
}

func (o ActionOps) BuildUpdateRequest(ctx context.Context, plan, state ActionTFModel) (httpclient.ActionRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = util.StringPtr("")
	}

	if core.ShouldClearString(plan.SlackMessageFormat, state.SlackMessageFormat) {
		req.SlackMessageFormat = util.StringPtr("")
	}

	deleted := plan.Deleted.ValueBool()
	req.Deleted = &deleted

	return req, diags
}

func (o ActionOps) MapResponseToModel(ctx context.Context, resp httpclient.Action, model *ActionTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.Int64Value(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.PostToSlack = core.PtrToBool(resp.PostToSlack)
	model.SlackMessageFormat = core.PtrToStringNullIfEmptyTrimmed(resp.SlackMessageFormat)

	if len(resp.Steps) > 0 {
		// Round-trip through JSON to get generic []interface{} for normalizeJSONForState
		stepsBytes, err := json.Marshal(resp.Steps)
		if err != nil {
			diags.AddError("Failed to marshal steps", err.Error())
			return diags
		}
		var stepsGeneric interface{}
		if err := json.Unmarshal(stepsBytes, &stepsGeneric); err != nil {
			diags.AddError("Failed to unmarshal steps", err.Error())
			return diags
		}
		normalized, err := normalizeJSONForState(stepsGeneric, model.StepsJSON.ValueString())
		if err != nil {
			diags.AddError("Failed to normalize steps", err.Error())
			return diags
		}
		model.StepsJSON = types.StringValue(normalized)
	} else if !model.StepsJSON.IsNull() {
		model.StepsJSON = types.StringValue("[]")
	} else {
		model.StepsJSON = types.StringNull()
	}

	tagsSet, d := core.TagsToSet(ctx, resp.Tags)
	diags.Append(d...)
	model.Tags = tagsSet

	deleted := resp.Deleted != nil && *resp.Deleted
	model.Deleted = types.BoolValue(deleted)

	return diags
}

func (o ActionOps) Create(ctx context.Context, client httpclient.PosthogClient, model ActionTFModel, req httpclient.ActionRequest) (httpclient.Action, error) {
	return client.CreateAction(ctx, model.GetEffectiveProjectID(), req)
}

func (o ActionOps) Read(ctx context.Context, client httpclient.PosthogClient, model ActionTFModel) (httpclient.Action, httpclient.HTTPStatusCode, error) {
	return client.GetAction(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o ActionOps) Update(ctx context.Context, client httpclient.PosthogClient, model ActionTFModel, req httpclient.ActionRequest) (httpclient.Action, httpclient.HTTPStatusCode, error) {
	return client.UpdateAction(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o ActionOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ActionTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteAction(ctx, model.GetEffectiveProjectID(), model.GetID())
}
