package resource

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewEventSchema() resource.Resource {
	return core.NewGenericResource[EventSchemaTFModel, httpclient.EventSchemaRequest, httpclient.EventSchema](
		EventSchemaOps{},
		core.ProjectScopedImportParser[EventSchemaTFModel](),
	)
}

type EventSchemaTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Event             types.String `tfsdk:"event"`
	PropertyGroupID   types.String `tfsdk:"property_group_id"`
	EventDefinitionID types.String `tfsdk:"event_definition_id"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
}

type EventSchemaOps struct{}

func (o EventSchemaOps) ResourceName() string {
	return "Event Schema"
}

func (o EventSchemaOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Attach a `posthog_schema_property_group` to an event. The event must have been " +
			"ingested at least once so its event definition exists.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the event schema attachment.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"event": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the event to attach the property group to (e.g. `checkout_completed`).",
			},
			"property_group_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "UUID of the `posthog_schema_property_group` to attach.",
			},
			// No UseStateForUnknown: when `event` changes the id must re-resolve.
			"event_definition_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the event definition the event name resolved to.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the event schema was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the event schema was last updated.",
			},
		},
	}
}

func (o EventSchemaOps) BuildCreateRequest(ctx context.Context, model EventSchemaTFModel) (httpclient.EventSchemaRequest, diag.Diagnostics) {
	// Name resolution needs the API client, which build methods don't get:
	// carry the name and let Create/Update resolve it.
	return httpclient.EventSchemaRequest{
		EventName:       model.Event.ValueString(),
		PropertyGroupID: model.PropertyGroupID.ValueString(),
	}, nil
}

func (o EventSchemaOps) BuildUpdateRequest(ctx context.Context, plan, state EventSchemaTFModel) (httpclient.EventSchemaRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o EventSchemaOps) MapResponseToModel(ctx context.Context, resp httpclient.EventSchema, model *EventSchemaTFModel) diag.Diagnostics {
	model.ID = types.StringValue(resp.ID)
	model.EventDefinitionID = types.StringValue(resp.EventDefinition)
	if resp.EventName != "" {
		model.Event = types.StringValue(resp.EventName)
	}
	if resp.PropertyGroup != nil {
		model.PropertyGroupID = types.StringValue(resp.PropertyGroup.ID)
	}
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.UpdatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.UpdatedAt)
	return nil
}

func (o EventSchemaOps) resolveEvent(ctx context.Context, client httpclient.PosthogClient, projectID string, req *httpclient.EventSchemaRequest) (httpclient.EventDefinition, error) {
	def, err := client.FindEventDefinitionByName(ctx, projectID, req.EventName)
	if err != nil {
		return httpclient.EventDefinition{}, err
	}
	req.EventDefinition = def.ID
	return def, nil
}

func (o EventSchemaOps) Create(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel, req httpclient.EventSchemaRequest) (httpclient.EventSchema, error) {
	def, err := o.resolveEvent(ctx, client, model.GetEffectiveProjectID(), &req)
	if err != nil {
		return httpclient.EventSchema{}, err
	}
	resp, err := client.CreateEventSchema(ctx, model.GetEffectiveProjectID(), req)
	if err != nil {
		return resp, err
	}
	resp.EventName = def.Name
	return resp, nil
}

func (o EventSchemaOps) Read(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel) (httpclient.EventSchema, httpclient.HTTPStatusCode, error) {
	// On import event_definition_id is not in state yet; fall back to an
	// unfiltered list scan.
	eventDefID := ""
	if !model.EventDefinitionID.IsNull() && !model.EventDefinitionID.IsUnknown() {
		eventDefID = model.EventDefinitionID.ValueString()
	}

	resp, status, err := client.GetEventSchema(ctx, model.GetEffectiveProjectID(), model.GetID(), eventDefID)
	if err != nil {
		return resp, status, err
	}

	def, defStatus, err := client.GetEventDefinition(ctx, model.GetEffectiveProjectID(), resp.EventDefinition)
	if err != nil {
		return resp, defStatus, err
	}
	resp.EventName = def.Name

	return resp, http.StatusOK, nil
}

func (o EventSchemaOps) Update(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel, req httpclient.EventSchemaRequest) (httpclient.EventSchema, httpclient.HTTPStatusCode, error) {
	def, err := o.resolveEvent(ctx, client, model.GetEffectiveProjectID(), &req)
	if err != nil {
		return httpclient.EventSchema{}, 0, err
	}
	resp, status, err := client.UpdateEventSchema(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
	if err != nil {
		return resp, status, err
	}
	resp.EventName = def.Name
	return resp, status, nil
}

func (o EventSchemaOps) Delete(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteEventSchema(ctx, model.GetEffectiveProjectID(), model.GetID())
}
