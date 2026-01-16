package core

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

type ProviderDefaults struct {
	ProjectID      string
	OrganizationID string
}

// ResourceOperations defines the resource-specific operations that must be provided.
// This interface bridges the generic CRUD flow and resource-specific logic.
type ResourceOperations[TFModel Identifiable, APIRequest, APIResponse any] interface {
	ResourceName() string
	Schema() schema.Schema

	// BuildCreateRequest converts the plan model to an API request for creation.
	BuildCreateRequest(ctx context.Context, model TFModel) (APIRequest, diag.Diagnostics)

	// BuildUpdateRequest converts the plan model to an API request for update.
	// Both plan and current state are provided in case differential updates are needed.
	BuildUpdateRequest(ctx context.Context, plan, state TFModel) (APIRequest, diag.Diagnostics)

	// MapResponseToModel maps an API response to the Terraform resource model.
	MapResponseToModel(ctx context.Context, resp APIResponse, model *TFModel) diag.Diagnostics

	Create(ctx context.Context, client httpclient.PosthogClient, model TFModel, req APIRequest) (APIResponse, error)
	Read(ctx context.Context, client httpclient.PosthogClient, model TFModel) (APIResponse, httpclient.HTTPStatusCode, error)
	Update(ctx context.Context, client httpclient.PosthogClient, model TFModel, req APIRequest) (APIResponse, httpclient.HTTPStatusCode, error)
	Delete(ctx context.Context, client httpclient.PosthogClient, model TFModel) (httpclient.HTTPStatusCode, error)
}

// GenericResource implements resource.Resource using the provided operations.
type GenericResource[TFModel Identifiable, APIRequest, APIResponse any] struct {
	client   httpclient.PosthogClient
	defaults ProviderDefaults
	ops      ResourceOperations[TFModel, APIRequest, APIResponse]
}

var (
	_ resource.Resource                = (*GenericResource[Identifiable, any, any])(nil)
	_ resource.ResourceWithImportState = (*GenericResource[Identifiable, any, any])(nil)
)

func NewGenericResource[TFModel Identifiable, APIRequest, APIResponse any](
	ops ResourceOperations[TFModel, APIRequest, APIResponse],
) resource.Resource {
	return &GenericResource[TFModel, APIRequest, APIResponse]{
		ops: ops,
	}
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_%s", req.ProviderTypeName,
		strings.ReplaceAll(strings.ToLower(r.ops.ResourceName()), " ", "_"),
	)
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = r.ops.Schema()
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
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
	r.defaults = ProviderDefaults{
		ProjectID: providerData.DefaultProjectID,
	}
}

// initializeScope sets provider defaults on the model's embedded structs and
// populates the state fields. Call once early in each CRUD operation.
func (r *GenericResource[TFModel, APIRequest, APIResponse]) initializeScope(model *TFModel) {
	if init, ok := any(model).(ProjectIDInitializer); ok {
		init.InitializeProjectID(r.defaults.ProjectID)
	}
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan TFModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build request using resource-specific logic
	request, diags := r.ops.BuildCreateRequest(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating PostHog %s", r.ops.ResourceName()))

	r.initializeScope(&plan)
	response, err := r.ops.Create(ctx, r.client, plan, request)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating %s", r.ops.ResourceName()),
			err.Error(),
		)
		return
	}

	// Map response to state using resource-specific logic
	resp.Diagnostics.Append(r.ops.MapResponseToModel(ctx, response, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Created PostHog %s", r.ops.ResourceName()), map[string]any{
		"id": plan.GetID(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state TFModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.HasValidID() {
		resp.State.RemoveResource(ctx)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Reading PostHog %s", r.ops.ResourceName()), map[string]any{
		"id": state.GetID(),
	})

	r.initializeScope(&state)
	response, statusCode, err := r.ops.Read(ctx, r.client, state)
	if err != nil {
		if statusCode == http.StatusNotFound {
			tflog.Warn(ctx, "Resource not found, removing from state", map[string]any{"id": state.GetID()})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error reading %s", r.ops.ResourceName()),
			err.Error(),
		)
		return
	}

	// Map response to state using resource-specific logic
	resp.Diagnostics.Append(r.ops.MapResponseToModel(ctx, response, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan, state TFModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.HasValidID() {
		resp.Diagnostics.AddError("Invalid resource ID", "Resource ID is absent in state")
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating PostHog %s", r.ops.ResourceName()), map[string]any{
		"id": state.GetID(),
	})

	// Set provider defaults on embedded structs (use state for project_id as it shouldn't change)
	r.initializeScope(&state)
	r.initializeScope(&plan)

	// Build request using resource-specific logic
	request, diags := r.ops.BuildUpdateRequest(ctx, plan, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use state's project_id for the API call (shouldn't change during update)
	response, statusCode, err := r.ops.Update(ctx, r.client, state, request)
	if err != nil {
		if statusCode == http.StatusNotFound {
			tflog.Warn(ctx, "Resource not found during update, removing from state", map[string]any{"id": state.GetID()})
			resp.Diagnostics.AddWarning(
				"Resource not found",
				fmt.Sprintf("%s with ID %s was deleted externally and will be recreated.", r.ops.ResourceName(), state.GetID()),
			)
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error updating %s", r.ops.ResourceName()),
			err.Error(),
		)
		return
	}

	// Map response to state using resource-specific logic
	resp.Diagnostics.Append(r.ops.MapResponseToModel(ctx, response, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Updated PostHog %s", r.ops.ResourceName()), map[string]any{
		"id": plan.GetID(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state TFModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.HasValidID() {
		resp.Diagnostics.AddError("Invalid resource ID", "Resource ID is absent in state")
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting PostHog %s", r.ops.ResourceName()), map[string]any{
		"id": state.GetID(),
	})

	r.initializeScope(&state)
	statusCode, err := r.ops.Delete(ctx, r.client, state)
	if err != nil {
		if statusCode == http.StatusNotFound {
			tflog.Warn(ctx, "Resource already deleted, removing from state", map[string]any{"id": state.GetID()})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting %s", r.ops.ResourceName()),
			err.Error(),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleted PostHog %s", r.ops.ResourceName()), map[string]any{
		"id": state.GetID(),
	})
}

func (r *GenericResource[TFModel, APIRequest, APIResponse]) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	var state TFModel
	var projectID, resourceID string

	// Project-scoped: parse "project_id/resource_id" or just "resource_id"
	projectID = r.defaults.ProjectID
	resourceID = req.ID

	if parts := strings.SplitN(req.ID, "/", 2); len(parts) == 2 {
		projectID = parts[0]
		resourceID = parts[1]
		tflog.Debug(ctx, "Importing with explicit project_id", map[string]any{
			"project_id":  projectID,
			"resource_id": resourceID,
		})
	} else if projectID == "" {
		resp.Diagnostics.AddError(
			"Missing project_id for import",
			"Cannot import resource: no project_id specified. "+
				"Use import ID format 'project_id/resource_id' or set project_id in the provider configuration.",
		)
		return
	}

	// Set the IDs on the state
	if setter, ok := any(&state).(IDSetter); ok {
		setter.SetID(resourceID)
	}
	if init, ok := any(&state).(ProjectIDInitializer); ok {
		init.InitializeProjectID(projectID)
	}

	response, _, err := r.ops.Read(ctx, r.client, state)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error reading %s during import", r.ops.ResourceName()),
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.ops.MapResponseToModel(ctx, response, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
