package resource

import (
	"context"
	"encoding/json"
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
)

func NewHogFunction() resource.Resource {
	return core.NewGenericResource[HogFunctionResourceTFModel, httpclient.HogFunctionRequest, httpclient.HogFunction](HogFunctionOps{})
}

type HogFunctionResourceTFModel struct {
	core.BaseStringIdentifiable
	Type             types.String `tfsdk:"type"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	Hog              types.String `tfsdk:"hog"`
	InputsSchemaJSON types.String `tfsdk:"inputs_schema_json"`
	InputsJSON       types.String `tfsdk:"inputs_json"`
	FiltersJSON      types.String `tfsdk:"filters_json"`
	MaskingJSON      types.String `tfsdk:"masking_json"`
	MappingsJSON     types.String `tfsdk:"mappings_json"`
	IconURL          types.String `tfsdk:"icon_url"`
	TemplateID       types.String `tfsdk:"template_id"`
	ExecutionOrder   types.Int64  `tfsdk:"execution_order"`
}

type HogFunctionOps struct{}

func (o HogFunctionOps) ResourceName() string {
	return "hog_function"
}

func (o HogFunctionOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog Hog functions. Hog functions enable destinations, webhooks, and transformations triggered by events or alerts.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Hog function.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Type of Hog function: `destination`, `site_destination`, `internal_destination`, `source_webhook`, `site_app`, or `transformation`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the Hog function.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the Hog function.",
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Whether the Hog function is enabled. Defaults to true.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"hog": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The Hog code to execute. Not required when using a template - the template provides the code.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"inputs_schema_json": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "JSON array defining the input schema for the Hog function. When using a template, this is provided by the template.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"inputs_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object containing the input values for the Hog function. Keys correspond to the input schema, values contain `value` and optional `templating` properties.",
			},
			"filters_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object defining filters for when the Hog function should execute. Includes `events`, `actions`, `properties`, and `filter_test_accounts` options.",
			},
			"masking_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON object configuring PII masking for the Hog function. Includes `ttl`, `threshold`, and `hash` properties.",
			},
			"mappings_json": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "JSON array of mapping configurations. Each mapping can have its own `name`, `inputs_schema`, `inputs`, and `filters`.",
			},
			"icon_url": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "URL of the icon for this Hog function.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"template_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "ID of a template to use as the basis for this Hog function. The template provides default code, inputs schema, and configuration.",
			},
			"execution_order": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Order in which this Hog function executes relative to others (0-32767).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o HogFunctionOps) BuildCreateRequest(ctx context.Context, model HogFunctionResourceTFModel) (httpclient.HogFunctionRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := httpclient.HogFunctionRequest{}

	if !model.Type.IsNull() && !model.Type.IsUnknown() {
		t := model.Type.ValueString()
		req.Type = &t
	}

	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		name := model.Name.ValueString()
		req.Name = &name
	}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}

	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		enabled := model.Enabled.ValueBool()
		req.Enabled = &enabled
	}

	if !model.Hog.IsNull() && !model.Hog.IsUnknown() {
		hog := model.Hog.ValueString()
		req.Hog = &hog
	}

	if !model.IconURL.IsNull() && !model.IconURL.IsUnknown() {
		iconURL := model.IconURL.ValueString()
		req.IconURL = &iconURL
	}

	if !model.TemplateID.IsNull() && !model.TemplateID.IsUnknown() {
		templateID := model.TemplateID.ValueString()
		req.TemplateID = &templateID
	}

	if !model.ExecutionOrder.IsNull() && !model.ExecutionOrder.IsUnknown() {
		order := int(model.ExecutionOrder.ValueInt64())
		req.ExecutionOrder = &order
	}

	if !model.InputsSchemaJSON.IsNull() && !model.InputsSchemaJSON.IsUnknown() {
		raw := strings.TrimSpace(model.InputsSchemaJSON.ValueString())
		if raw != "" {
			var inputsSchema []map[string]interface{}
			if err := json.Unmarshal([]byte(raw), &inputsSchema); err != nil {
				diags.AddError("Invalid inputs_schema_json", "inputs_schema_json must be valid JSON array: "+err.Error())
				return req, diags
			}
			req.InputsSchema = inputsSchema
		}
	}

	if !model.InputsJSON.IsNull() && !model.InputsJSON.IsUnknown() {
		raw := strings.TrimSpace(model.InputsJSON.ValueString())
		if raw != "" {
			var inputs map[string]interface{}
			if err := json.Unmarshal([]byte(raw), &inputs); err != nil {
				diags.AddError("Invalid inputs_json", "inputs_json must be valid JSON object: "+err.Error())
				return req, diags
			}
			req.Inputs = inputs
		}
	}

	if !model.FiltersJSON.IsNull() && !model.FiltersJSON.IsUnknown() {
		raw := strings.TrimSpace(model.FiltersJSON.ValueString())
		if raw != "" {
			var filters map[string]interface{}
			if err := json.Unmarshal([]byte(raw), &filters); err != nil {
				diags.AddError("Invalid filters_json", "filters_json must be valid JSON object: "+err.Error())
				return req, diags
			}
			req.Filters = filters
		}
	}

	if !model.MaskingJSON.IsNull() && !model.MaskingJSON.IsUnknown() {
		raw := strings.TrimSpace(model.MaskingJSON.ValueString())
		if raw != "" {
			var masking map[string]interface{}
			if err := json.Unmarshal([]byte(raw), &masking); err != nil {
				diags.AddError("Invalid masking_json", "masking_json must be valid JSON object: "+err.Error())
				return req, diags
			}
			req.Masking = masking
		}
	}

	if !model.MappingsJSON.IsNull() && !model.MappingsJSON.IsUnknown() {
		raw := strings.TrimSpace(model.MappingsJSON.ValueString())
		if raw != "" {
			var mappings []map[string]interface{}
			if err := json.Unmarshal([]byte(raw), &mappings); err != nil {
				diags.AddError("Invalid mappings_json", "mappings_json must be valid JSON array: "+err.Error())
				return req, diags
			}
			req.Mappings = mappings
		}
	}

	return req, diags
}

func (o HogFunctionOps) BuildUpdateRequest(ctx context.Context, plan, state HogFunctionResourceTFModel) (httpclient.HogFunctionRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	// Clear description if removed from config
	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = core.StringPtr("")
	}

	return req, diags
}

func (o HogFunctionOps) MapResponseToModel(ctx context.Context, resp httpclient.HogFunction, model *HogFunctionResourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)

	if resp.Type != nil && *resp.Type != "" {
		model.Type = types.StringValue(*resp.Type)
	} else {
		model.Type = types.StringNull()
	}

	model.Name = core.PtrToStringNullIfEmptyTrimmed(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.Enabled = core.PtrToBool(resp.Enabled)
	model.Hog = core.PtrToStringNullIfEmptyTrimmed(resp.Hog)
	model.IconURL = core.PtrToStringNullIfEmptyTrimmed(resp.IconURL)

	// TemplateID - only set if the model already has one (write-only field)
	// The API returns template object, not template_id
	if !model.TemplateID.IsNull() {
		// Keep the existing template_id from state
	} else if resp.Template != nil && resp.Template.ID != "" {
		// If we have a template in response and no template_id in state, set it
		model.TemplateID = types.StringValue(resp.Template.ID)
	} else {
		model.TemplateID = types.StringNull()
	}

	if resp.ExecutionOrder != nil {
		model.ExecutionOrder = types.Int64Value(int64(*resp.ExecutionOrder))
	} else {
		model.ExecutionOrder = types.Int64Null()
	}

	// Only update inputs_schema_json if user configured it (not template-provided)
	if !model.InputsSchemaJSON.IsNull() {
		if len(resp.InputsSchema) > 0 {
			normalized, err := normalizeJSONForState(resp.InputsSchema, model.InputsSchemaJSON.ValueString())
			if err != nil {
				diags.AddError("Failed to normalize inputs_schema", err.Error())
				return diags
			}
			model.InputsSchemaJSON = types.StringValue(normalized)
		} else {
			model.InputsSchemaJSON = types.StringValue("[]")
		}
	}

	// Only update inputs_json if user configured it (not template-provided)
	// This prevents inconsistent state when template provides default inputs
	if !model.InputsJSON.IsNull() {
		if len(resp.Inputs) > 0 {
			normalized, err := normalizeJSONForState(resp.Inputs, model.InputsJSON.ValueString())
			if err != nil {
				diags.AddError("Failed to normalize inputs", err.Error())
				return diags
			}
			model.InputsJSON = types.StringValue(normalized)
		} else {
			model.InputsJSON = types.StringValue("{}")
		}
	}

	if len(resp.Filters) > 0 {
		normalized, err := normalizeFiltersJSON(resp.Filters, model.FiltersJSON.ValueString())
		if err != nil {
			diags.AddError("Failed to normalize filters", err.Error())
			return diags
		}
		model.FiltersJSON = types.StringValue(normalized)
	} else if !model.FiltersJSON.IsNull() {
		model.FiltersJSON = types.StringValue("{}")
	} else {
		model.FiltersJSON = types.StringNull()
	}

	if len(resp.Masking) > 0 {
		normalized, err := normalizeJSONForState(resp.Masking, model.MaskingJSON.ValueString())
		if err != nil {
			diags.AddError("Failed to normalize masking", err.Error())
			return diags
		}
		model.MaskingJSON = types.StringValue(normalized)
	} else if !model.MaskingJSON.IsNull() {
		model.MaskingJSON = types.StringValue("{}")
	} else {
		model.MaskingJSON = types.StringNull()
	}

	if len(resp.Mappings) > 0 {
		normalized, err := normalizeJSONForState(resp.Mappings, model.MappingsJSON.ValueString())
		if err != nil {
			diags.AddError("Failed to normalize mappings", err.Error())
			return diags
		}
		model.MappingsJSON = types.StringValue(normalized)
	} else if !model.MappingsJSON.IsNull() {
		model.MappingsJSON = types.StringValue("[]")
	} else {
		model.MappingsJSON = types.StringNull()
	}

	return diags
}

func (o HogFunctionOps) Create(ctx context.Context, client httpclient.PosthogClient, req httpclient.HogFunctionRequest) (httpclient.HogFunction, error) {
	return client.CreateHogFunction(ctx, req)
}

func (o HogFunctionOps) Read(ctx context.Context, client httpclient.PosthogClient, id string) (httpclient.HogFunction, httpclient.HTTPStatusCode, error) {
	return client.GetHogFunction(ctx, id)
}

func (o HogFunctionOps) Update(ctx context.Context, client httpclient.PosthogClient, id string, req httpclient.HogFunctionRequest) (httpclient.HogFunction, httpclient.HTTPStatusCode, error) {
	return client.UpdateHogFunction(ctx, id, req)
}

func (o HogFunctionOps) Delete(ctx context.Context, client httpclient.PosthogClient, id string) (httpclient.HTTPStatusCode, error) {
	return client.DeleteHogFunction(ctx, id)
}

// normalizeJSONForState takes the API response data and user's configured JSON,
// filters the API response to only include fields the user specified, and returns
// canonical JSON.
func normalizeJSONForState(apiData interface{}, userJSON string) (string, error) {
	if apiData == nil {
		return "", nil
	}

	// If user didn't specify JSON, return the API data as canonical JSON
	userJSON = strings.TrimSpace(userJSON)
	if userJSON == "" {
		return marshalJSON(apiData)
	}

	// Parse user's JSON to know which fields to keep
	var userData interface{}
	if err := json.Unmarshal([]byte(userJSON), &userData); err != nil {
		// If we can't parse user's JSON, fall back to returning API data
		return marshalJSON(apiData)
	}

	filtered := filterToOnlyIncludeUserFields(userData, apiData)
	return marshalJSON(filtered)
}

// normalizeFiltersJSON normalizes the API response for filters_json,
// stripping server-computed fields like bytecode.
func normalizeFiltersJSON(apiData map[string]interface{}, userJSON string) (string, error) {
	if apiData == nil {
		return "", nil
	}

	// Remove server-computed bytecode from API response
	apiDataCleaned := make(map[string]interface{})
	for k, v := range apiData {
		if k != "bytecode" {
			apiDataCleaned[k] = v
		}
	}

	return normalizeJSONForState(apiDataCleaned, userJSON)
}
