package resource

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

var (
	_ resource.Resource                = &projectResource{}
	_ resource.ResourceWithConfigure   = &projectResource{}
	_ resource.ResourceWithImportState = &projectResource{}
)

func NewProject() resource.Resource {
	return &projectResource{}
}

type projectResource struct {
	client httpclient.PosthogClient
}

type ProjectTFModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	OrganizationID types.String `tfsdk:"organization_id"`
	APIToken       types.String `tfsdk:"api_token"`
	Timezone       types.String `tfsdk:"timezone"`
}

func (r *projectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *projectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a PostHog project within an organization.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The numeric identifier of the project.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the project.",
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The identifier of the organization this project belongs to.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"api_token": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The API token for this project. This is used to send events to PostHog.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"timezone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The timezone for this project (e.g., 'UTC', 'America/New_York', 'Europe/London'). Defaults to 'UTC'.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *projectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProjectTFModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	organizationID := plan.OrganizationID.ValueString()
	name := plan.Name.ValueString()

	tflog.Debug(ctx, "Creating PostHog Project", map[string]any{
		"organization_id": organizationID,
		"name":            name,
	})

	apiReq := httpclient.ProjectRequest{
		Name: &name,
	}

	if !plan.Timezone.IsNull() && !plan.Timezone.IsUnknown() {
		tz := plan.Timezone.ValueString()
		apiReq.Timezone = &tz
	}

	project, err := r.client.CreateProject(ctx, organizationID, apiReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating project", err.Error())
		return
	}

	r.mapResponseToModel(&project, &plan)

	tflog.Debug(ctx, "Created PostHog Project", map[string]any{
		"id":              plan.ID.ValueInt64(),
		"organization_id": organizationID,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *projectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProjectTFModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.IsNull() || state.ID.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	organizationID := state.OrganizationID.ValueString()
	projectID := fmt.Sprintf("%d", state.ID.ValueInt64())

	tflog.Debug(ctx, "Reading PostHog Project", map[string]any{
		"id":              projectID,
		"organization_id": organizationID,
	})

	project, statusCode, err := r.client.GetProject(ctx, organizationID, projectID)
	if err != nil {
		if statusCode == http.StatusNotFound {
			tflog.Warn(ctx, "Project not found, removing from state", map[string]any{"id": projectID})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading project", err.Error())
		return
	}

	r.mapResponseToModel(&project, &state)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *projectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ProjectTFModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	organizationID := state.OrganizationID.ValueString()
	projectID := fmt.Sprintf("%d", state.ID.ValueInt64())

	tflog.Debug(ctx, "Updating PostHog Project", map[string]any{
		"id":              projectID,
		"organization_id": organizationID,
	})

	name := plan.Name.ValueString()
	apiReq := httpclient.ProjectRequest{
		Name: &name,
	}

	if !plan.Timezone.IsNull() && !plan.Timezone.IsUnknown() {
		tz := plan.Timezone.ValueString()
		apiReq.Timezone = &tz
	}

	project, statusCode, err := r.client.UpdateProject(ctx, organizationID, projectID, apiReq)
	if err != nil {
		if statusCode == http.StatusNotFound {
			tflog.Warn(ctx, "Project not found during update, removing from state", map[string]any{"id": projectID})
			resp.Diagnostics.AddWarning(
				"Resource not found",
				fmt.Sprintf("Project with ID %s was deleted externally and will be recreated.", projectID),
			)
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error updating project", err.Error())
		return
	}

	r.mapResponseToModel(&project, &plan)

	tflog.Debug(ctx, "Updated PostHog Project", map[string]any{
		"id":              plan.ID.ValueInt64(),
		"organization_id": organizationID,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *projectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProjectTFModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.IsNull() || state.ID.IsUnknown() {
		resp.Diagnostics.AddError("Invalid resource ID", "Resource ID is absent in state")
		return
	}

	organizationID := state.OrganizationID.ValueString()
	projectID := fmt.Sprintf("%d", state.ID.ValueInt64())

	tflog.Debug(ctx, "Deleting PostHog Project", map[string]any{
		"id":              projectID,
		"organization_id": organizationID,
	})

	statusCode, err := r.client.DeleteProject(ctx, organizationID, projectID)
	if err != nil {
		if statusCode == http.StatusNotFound {
			tflog.Warn(ctx, "Project already deleted, removing from state", map[string]any{"id": projectID})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error deleting project", err.Error())
		return
	}

	tflog.Debug(ctx, "Deleted PostHog Project", map[string]any{
		"id":              projectID,
		"organization_id": organizationID,
	})
}

func (r *projectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: organization_id:project_id
	parts := strings.Split(req.ID, ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Import ID must be in format 'organization_id:project_id', got: %s", req.ID),
		)
		return
	}

	organizationID := parts[0]
	projectID := parts[1]

	tflog.Debug(ctx, "Importing PostHog Project", map[string]any{
		"organization_id": organizationID,
		"project_id":      projectID,
	})

	project, _, err := r.client.GetProject(ctx, organizationID, projectID)
	if err != nil {
		resp.Diagnostics.AddError("Error reading project during import", err.Error())
		return
	}

	var state ProjectTFModel
	state.OrganizationID = types.StringValue(organizationID)
	r.mapResponseToModel(&project, &state)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *projectResource) mapResponseToModel(project *httpclient.Project, model *ProjectTFModel) {
	model.ID = types.Int64Value(project.ID)
	model.Name = core.PtrToStringNullIfEmptyTrimmed(project.Name)
	model.APIToken = core.PtrToStringNullIfEmptyTrimmed(project.APIToken)
	model.Timezone = core.PtrToStringNullIfEmptyTrimmed(project.Timezone)
}
