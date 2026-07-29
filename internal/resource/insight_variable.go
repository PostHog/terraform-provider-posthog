package resource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

// InsightVariableTypes are the variable types PostHog accepts.
var InsightVariableTypes = []string{"String", "Number", "Boolean", "List", "Date"}

// insightVariableTypeList is the one type PostHog coerces values for.
const insightVariableTypeList = "List"

func NewInsightVariable() resource.Resource {
	return core.NewGenericResource[InsightVariableResourceTFModel, httpclient.InsightVariableRequest, httpclient.InsightVariable](
		InsightVariableOps{},
		core.ProjectScopedImportParser[InsightVariableResourceTFModel](),
	)
}

type InsightVariableResourceTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name             types.String `tfsdk:"name"`
	Type             types.String `tfsdk:"type"`
	CodeName         types.String `tfsdk:"code_name"`
	DefaultValueJSON types.String `tfsdk:"default_value_json"`
	ValuesJSON       types.String `tfsdk:"values_json"`
	CreatedAt        types.String `tfsdk:"created_at"`
}

type InsightVariableOps struct{}

func (o InsightVariableOps) ResourceName() string {
	return "insight_variable"
}

func (o InsightVariableOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog SQL variables (also called insight variables). A HogQL query references a variable as `{variables.code_name}`, " +
			"and every dashboard holding an insight that uses it renders an input for the value.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the SQL variable. Insight queries key their `variables` map on this value.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Human-readable name for the SQL variable, shown above the input.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 400),
				},
			},
			"type": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Variable type, which controls how the value is rendered and substituted into HogQL: " +
					"`String`, `Number`, `Boolean`, `List`, or `Date`.",
				Validators: []validator.String{
					stringvalidator.OneOf(InsightVariableTypes...),
				},
			},
			"code_name": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Code-safe name generated from `name`, referenced in HogQL as `{variables.code_name}`. " +
					"PostHog derives it once at creation and keeps it if you later rename the variable, so queries do not break on a rename. " +
					"It must be unique within the project.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"default_value_json": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "JSON-encoded default value used when a query references the variable and no value is supplied. " +
					"The value is a bare JSON scalar for most types, so wrap it: `jsonencode(\"prod\")`, `jsonencode(30)`, `jsonencode(true)`. " +
					"For `List` variables PostHog stores the default as a string, so encode it as one.",
			},
			"values_json": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "JSON-encoded array of allowed values, for `List` variables only. " +
					"PostHog stores every entry as a string, so encode them as strings: `jsonencode([\"1\", \"2\"])` rather than `jsonencode([1, 2])`. " +
					"Otherwise the value PostHog returns will not match your configuration. Leave unset for other variable types.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the SQL variable was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o InsightVariableOps) BuildCreateRequest(ctx context.Context, model InsightVariableResourceTFModel) (httpclient.InsightVariableRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := model.Name.ValueString()
	variableType := model.Type.ValueString()
	req := httpclient.InsightVariableRequest{
		Name: &name,
		Type: &variableType,
	}

	defaultValue, d := util.ParseJSONStringRaw("default_value_json", model.DefaultValueJSON)
	diags.Append(d...)
	req.DefaultValue = defaultValue

	values, d := util.ParseJSONStringRaw("values_json", model.ValuesJSON)
	diags.Append(d...)
	req.Values = values

	if diags.HasError() {
		return req, diags
	}

	diags.Append(validateInsightVariableListShape(variableType, req)...)

	return req, diags
}

func (o InsightVariableOps) BuildUpdateRequest(ctx context.Context, plan, state InsightVariableResourceTFModel) (httpclient.InsightVariableRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	// Clear values removed from configuration. A PATCH that omits the field
	// leaves the stored value in place, so clearing has to be explicit.
	if plan.DefaultValueJSON.IsNull() && !state.DefaultValueJSON.IsNull() {
		req.DefaultValue = util.JSONNull()
	}
	if plan.ValuesJSON.IsNull() && !state.ValuesJSON.IsNull() {
		req.Values = util.JSONNull()
	}

	return req, diags
}

func (o InsightVariableOps) MapResponseToModel(ctx context.Context, resp httpclient.InsightVariable, model *InsightVariableResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.Type = core.PtrToStringNullIfEmptyTrimmed(resp.Type)
	model.CodeName = core.PtrToStringNullIfEmptyTrimmed(resp.CodeName)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)

	defaultValue, err := jsonValueToState(resp.DefaultValue, model.DefaultValueJSON)
	if err != nil {
		diags.AddError("Failed to read default_value", err.Error())
		return diags
	}
	model.DefaultValueJSON = defaultValue

	values, err := jsonValueToState(resp.Values, model.ValuesJSON)
	if err != nil {
		diags.AddError("Failed to read values", err.Error())
		return diags
	}
	model.ValuesJSON = values

	return diags
}

func (o InsightVariableOps) Create(ctx context.Context, client httpclient.PosthogClient, model InsightVariableResourceTFModel, req httpclient.InsightVariableRequest) (httpclient.InsightVariable, error) {
	return client.CreateInsightVariable(ctx, model.GetEffectiveProjectID(), req)
}

func (o InsightVariableOps) Read(ctx context.Context, client httpclient.PosthogClient, model InsightVariableResourceTFModel) (httpclient.InsightVariable, httpclient.HTTPStatusCode, error) {
	return client.GetInsightVariable(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o InsightVariableOps) Update(ctx context.Context, client httpclient.PosthogClient, model InsightVariableResourceTFModel, req httpclient.InsightVariableRequest) (httpclient.InsightVariable, httpclient.HTTPStatusCode, error) {
	return client.UpdateInsightVariable(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o InsightVariableOps) Delete(ctx context.Context, client httpclient.PosthogClient, model InsightVariableResourceTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteInsightVariable(ctx, model.GetEffectiveProjectID(), model.GetID())
}

// jsonValueToState converts a raw JSON value from the API into canonical JSON
// for state, matching what jsonencode produces so an unchanged variable plans
// clean.
//
// PostHog represents an unset value inconsistently: null for most variable
// types, and "" (default) or [] (values) for List variables. All of those read
// back as unset unless the configuration explicitly asked for that value.
func jsonValueToState(raw *json.RawMessage, config types.String) (types.String, error) {
	if raw == nil || len(*raw) == 0 {
		return types.StringNull(), nil
	}

	var decoded interface{}
	if err := json.Unmarshal(*raw, &decoded); err != nil {
		return types.StringNull(), fmt.Errorf("failed to parse JSON returned by PostHog: %w", err)
	}

	if config.IsNull() && isUnsetJSONValue(decoded) {
		return types.StringNull(), nil
	}

	canonical, err := marshalJSON(decoded)
	if err != nil {
		return types.StringNull(), err
	}
	return types.StringValue(canonical), nil
}

// isUnsetJSONValue reports whether a decoded JSON value is one of the shapes
// PostHog stores for "no value": null, an empty string, or an empty array.
func isUnsetJSONValue(v interface{}) bool {
	switch value := v.(type) {
	case nil:
		return true
	case string:
		return value == ""
	case []interface{}:
		return len(value) == 0
	default:
		return false
	}
}

// validateInsightVariableListShape rejects List values PostHog would silently
// rewrite. It coerces list entries and the default to strings, which Terraform
// would then report as a provider inconsistency rather than as the API quirk it
// is, so fail with an actionable message instead.
func validateInsightVariableListShape(variableType string, req httpclient.InsightVariableRequest) diag.Diagnostics {
	var diags diag.Diagnostics
	if variableType != insightVariableTypeList {
		return diags
	}

	if req.Values != nil {
		var values []interface{}
		if err := json.Unmarshal(*req.Values, &values); err != nil {
			diags.AddError(
				"Invalid values_json",
				"values_json must be a JSON array for List variables, for example jsonencode([\"a\", \"b\"]).",
			)
			return diags
		}
		for i, value := range values {
			if _, ok := value.(string); !ok {
				diags.AddError(
					"Invalid values_json",
					fmt.Sprintf("PostHog stores List values as strings, and the entry at index %d is not one. "+
						"Quote it, for example jsonencode([\"1\", \"2\"]) rather than jsonencode([1, 2]).", i),
				)
				return diags
			}
		}
	}

	if req.DefaultValue != nil {
		var defaultValue interface{}
		if err := json.Unmarshal(*req.DefaultValue, &defaultValue); err == nil {
			if _, ok := defaultValue.(string); !ok && defaultValue != nil {
				diags.AddError(
					"Invalid default_value_json",
					"PostHog stores the default of a List variable as a string. "+
						"Quote it, for example jsonencode(\"1\") rather than jsonencode(1).",
				)
			}
		}
	}

	return diags
}
