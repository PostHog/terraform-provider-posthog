package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

func NewSchemaPropertyGroup() resource.Resource {
	return core.NewGenericResource[SchemaPropertyGroupTFModel, httpclient.SchemaPropertyGroupRequest, httpclient.SchemaPropertyGroup](
		SchemaPropertyGroupOps{},
		core.ProjectScopedImportParser[SchemaPropertyGroupTFModel](),
	)
}

type SchemaPropertyGroupTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Properties  types.Set    `tfsdk:"properties"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

type schemaPropertyTFModel struct {
	Name              types.String `tfsdk:"name"`
	PropertyType      types.String `tfsdk:"property_type"`
	IsRequired        types.Bool   `tfsdk:"is_required"`
	IsOptionalInTypes types.Bool   `tfsdk:"is_optional_in_types"`
	Description       types.String `tfsdk:"description"`
}

var schemaPropertyAttrTypes = map[string]attr.Type{
	"name":                 types.StringType,
	"property_type":        types.StringType,
	"is_required":          types.BoolType,
	"is_optional_in_types": types.BoolType,
	"description":          types.StringType,
}

type SchemaPropertyGroupOps struct{}

func (o SchemaPropertyGroupOps) ResourceName() string {
	return "Schema Property Group"
}

func (o SchemaPropertyGroupOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog schema property groups: named groups of typed properties that can be " +
			"attached to events via `posthog_event_schema`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the property group.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the property group. Unique per project.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the property group.",
			},
			"properties": schema.SetNestedAttribute{
				Optional: true,
				MarkdownDescription: "Typed properties in this group. Property names are unique within a group. " +
					"Updates replace the full property list server-side.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Property name.",
						},
						"property_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Property type: `DateTime`, `String`, `Numeric`, `Boolean`, or `Object`.",
							Validators: []validator.String{
								stringvalidator.OneOf("DateTime", "String", "Numeric", "Boolean", "Object"),
							},
						},
						"is_required": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
							MarkdownDescription: "Whether events must include this property to conform to the schema.",
						},
						"is_optional_in_types": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
							MarkdownDescription: "Whether the property is marked optional in generated type definitions.",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Property description.",
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the property group was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the property group was last updated.",
			},
		},
	}
}

func (o SchemaPropertyGroupOps) BuildCreateRequest(ctx context.Context, model SchemaPropertyGroupTFModel) (httpclient.SchemaPropertyGroupRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := model.Name.ValueString()
	req := httpclient.SchemaPropertyGroupRequest{Name: &name}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}

	if !model.Properties.IsNull() && !model.Properties.IsUnknown() {
		var props []schemaPropertyTFModel
		diags.Append(model.Properties.ElementsAs(ctx, &props, false)...)
		if diags.HasError() {
			return req, diags
		}

		apiProps := make([]httpclient.SchemaPropertyGroupProperty, 0, len(props))
		for _, p := range props {
			apiProp := httpclient.SchemaPropertyGroupProperty{
				Name:         p.Name.ValueString(),
				PropertyType: p.PropertyType.ValueString(),
			}
			if !p.IsRequired.IsNull() && !p.IsRequired.IsUnknown() {
				v := p.IsRequired.ValueBool()
				apiProp.IsRequired = &v
			}
			if !p.IsOptionalInTypes.IsNull() && !p.IsOptionalInTypes.IsUnknown() {
				v := p.IsOptionalInTypes.ValueBool()
				apiProp.IsOptionalInTypes = &v
			}
			if !p.Description.IsNull() && !p.Description.IsUnknown() {
				v := p.Description.ValueString()
				apiProp.Description = &v
			}
			apiProps = append(apiProps, apiProp)
		}
		req.Properties = &apiProps
	}

	return req, diags
}

func (o SchemaPropertyGroupOps) BuildUpdateRequest(ctx context.Context, plan, state SchemaPropertyGroupTFModel) (httpclient.SchemaPropertyGroupRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = util.StringPtr("")
	}

	// Clear properties if removed from config but previously set
	if plan.Properties.IsNull() && !state.Properties.IsNull() {
		empty := []httpclient.SchemaPropertyGroupProperty{}
		req.Properties = &empty
	}

	return req, diags
}

func (o SchemaPropertyGroupOps) MapResponseToModel(ctx context.Context, resp httpclient.SchemaPropertyGroup, model *SchemaPropertyGroupTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.UpdatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.UpdatedAt)

	if len(resp.Properties) == 0 && model.Properties.IsNull() {
		model.Properties = types.SetNull(types.ObjectType{AttrTypes: schemaPropertyAttrTypes})
		return diags
	}

	elems := make([]schemaPropertyTFModel, 0, len(resp.Properties))
	for _, p := range resp.Properties {
		elems = append(elems, schemaPropertyTFModel{
			Name:              types.StringValue(p.Name),
			PropertyType:      types.StringValue(p.PropertyType),
			IsRequired:        types.BoolValue(p.IsRequired != nil && *p.IsRequired),
			IsOptionalInTypes: types.BoolValue(p.IsOptionalInTypes != nil && *p.IsOptionalInTypes),
			Description:       core.PtrToStringNullIfEmptyTrimmed(p.Description),
		})
	}

	set, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: schemaPropertyAttrTypes}, elems)
	diags.Append(d...)
	model.Properties = set

	return diags
}

func (o SchemaPropertyGroupOps) Create(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel, req httpclient.SchemaPropertyGroupRequest) (httpclient.SchemaPropertyGroup, error) {
	return client.CreateSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), req)
}

func (o SchemaPropertyGroupOps) Read(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel) (httpclient.SchemaPropertyGroup, httpclient.HTTPStatusCode, error) {
	return client.GetSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o SchemaPropertyGroupOps) Update(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel, req httpclient.SchemaPropertyGroupRequest) (httpclient.SchemaPropertyGroup, httpclient.HTTPStatusCode, error) {
	return client.UpdateSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o SchemaPropertyGroupOps) Delete(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), model.GetID())
}
