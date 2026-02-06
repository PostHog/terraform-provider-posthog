package resource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewAccessControl() resource.Resource {
	return core.NewGenericResource[AccessControlTFModel, httpclient.AccessControlRequest, httpclient.AccessControl](
		AccessControlOps{},
		core.AccessControlImportParser[AccessControlTFModel](),
	)
}

type AccessControlTFModel struct {
	core.BaseProjectID
	// Composite ID: project_id/resource/[resource_id/]role/role_id or similar for member
	ID                 types.String `tfsdk:"id"`
	Resource           types.String `tfsdk:"resource"`
	ResourceID         types.String `tfsdk:"resource_id"`
	AccessLevel        types.String `tfsdk:"access_level"`
	Role               types.String `tfsdk:"role"`
	OrganizationMember types.String `tfsdk:"organization_member"`
	CreatedAt          types.String `tfsdk:"created_at"`
	UpdatedAt          types.String `tfsdk:"updated_at"`
}

func (m AccessControlTFModel) GetID() string {
	return m.ID.ValueString()
}

// HasValidID returns true if we have enough information to identify the access control.
// For access controls, this means having resource. Role and organization_member are optional
// (when neither is set, it's a project-level default).
func (m AccessControlTFModel) HasValidID() bool {
	return !m.Resource.IsNull() && !m.Resource.IsUnknown() && m.Resource.ValueString() != ""
}

func (m *AccessControlTFModel) SetAccessControlFields(resourceType, resourceID, targetType, targetID string) {
	m.Resource = types.StringValue(resourceType)
	if resourceID != "" {
		m.ResourceID = types.StringValue(resourceID)
	}
	switch targetType {
	case httpclient.AccessControlTargetRole:
		m.Role = types.StringValue(targetID)
	case httpclient.AccessControlTargetMember:
		m.OrganizationMember = types.StringValue(targetID)
	}
}

// buildExpectedID constructs the composite ID from the model's fields.
func (m AccessControlTFModel) buildExpectedID(projectID string) string {
	resourcePart := m.Resource.ValueString()
	if !m.ResourceID.IsNull() && !m.ResourceID.IsUnknown() && m.ResourceID.ValueString() != "" {
		resourcePart = fmt.Sprintf("%s/%s", resourcePart, m.ResourceID.ValueString())
	}
	if !m.Role.IsNull() && !m.Role.IsUnknown() && m.Role.ValueString() != "" {
		return fmt.Sprintf("%s/%s/%s/%s", projectID, resourcePart, httpclient.AccessControlTargetRole, m.Role.ValueString())
	}
	if !m.OrganizationMember.IsNull() && !m.OrganizationMember.IsUnknown() && m.OrganizationMember.ValueString() != "" {
		return fmt.Sprintf("%s/%s/%s/%s", projectID, resourcePart, httpclient.AccessControlTargetMember, m.OrganizationMember.ValueString())
	}
	// Project default (no role or member specified)
	return fmt.Sprintf("%s/%s/%s", projectID, resourcePart, httpclient.AccessControlTargetDefault)
}

type AccessControlOps struct{}

func (o AccessControlOps) ResourceName() string {
	return "access_control"
}

func (o AccessControlOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: `Manages access control for resources within a PostHog project.

This resource allows you to set access levels for resource types (like feature flags, dashboards, etc.).

You can set permissions at three levels:
- **Project default**: Applies to everyone in the project for this resource type. Omit both ` + "`role`" + ` and ` + "`organization_member`" + `.
- **Role-specific**: Applies to members of a specific role. Set ` + "`role`" + `.
- **Member-specific**: Applies to a specific organization member. Set ` + "`organization_member`" + `.

Additionally, you can scope to:
- **Resource-type level**: Applies to all resources of a type (e.g., all dashboards). Omit ` + "`resource_id`" + `.
- **Resource-instance level**: Applies to a specific resource (e.g., one dashboard). Set ` + "`resource_id`" + `.

You can combine these: set a project default for all dashboards, or set a project default for a specific dashboard.

~> **Note:** ` + "`role`" + ` and ` + "`organization_member`" + ` are mutually exclusive - you cannot specify both. Omit both for project defaults.` + core.EnterpriseRBACNote,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Composite identifier for the access control.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"resource": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "The resource type to control access for. " +
					"Valid values include: `action`, `alert`, `annotation`, `cohort`, `dashboard`, " +
					"`experiment`, `feature_flag`, `insight`, `notebook`, `session_recording`, `survey`, etc.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"resource_id": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "The ID of a specific resource to control access for. " +
					"If omitted, the access control applies to all resources of the specified type.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"access_level": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The access level to grant. Common values are `none`, `viewer`, `editor`.",
			},
			"role": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "The UUID of the role to grant access to. " +
					"Mutually exclusive with `organization_member`. " +
					"If neither `role` nor `organization_member` is set, this becomes the project default for the resource type.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.Expressions{
						path.MatchRoot("organization_member"),
					}...),
				},
			},
			"organization_member": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "The organization member ID to grant access to (either `organization_member_id` from `posthog_user` data source, or `posthog_organization_member.<name>.id`). " +
					"Mutually exclusive with `role`. " +
					"If neither `role` nor `organization_member` is set, this becomes the project default for the resource type.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the access control was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the access control was last updated.",
			},
		},
	}
}

func (o AccessControlOps) BuildCreateRequest(_ context.Context, model AccessControlTFModel) (httpclient.AccessControlRequest, diag.Diagnostics) {
	return buildAccessControlRequest(model), nil
}

func (o AccessControlOps) BuildUpdateRequest(_ context.Context, plan, _ AccessControlTFModel) (httpclient.AccessControlRequest, diag.Diagnostics) {
	return buildAccessControlRequest(plan), nil
}

func (o AccessControlOps) MapResponseToModel(_ context.Context, resp httpclient.AccessControl, model *AccessControlTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.BuildCompositeID(model.GetEffectiveProjectID()))

	model.Resource = types.StringValue(resp.Resource)
	if resp.ResourceID != nil && *resp.ResourceID != "" {
		model.ResourceID = types.StringValue(*resp.ResourceID)
	}
	if resp.Role != nil {
		model.Role = types.StringValue(*resp.Role)
	}
	if resp.OrganizationMember != nil {
		model.OrganizationMember = types.StringValue(*resp.OrganizationMember)
	}
	if resp.AccessLevel != nil {
		model.AccessLevel = types.StringValue(*resp.AccessLevel)
	}
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.UpdatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.UpdatedAt)

	return diags
}

func (o AccessControlOps) Create(ctx context.Context, client httpclient.PosthogClient, model AccessControlTFModel, req httpclient.AccessControlRequest) (httpclient.AccessControl, error) {
	return client.UpsertAccessControl(ctx, model.GetEffectiveProjectID(), req)
}

func (o AccessControlOps) Read(ctx context.Context, client httpclient.PosthogClient, model AccessControlTFModel) (httpclient.AccessControl, httpclient.HTTPStatusCode, error) {
	var resourceID *string
	if !model.ResourceID.IsNull() && !model.ResourceID.IsUnknown() && model.ResourceID.ValueString() != "" {
		rid := model.ResourceID.ValueString()
		resourceID = &rid
	}

	result, status, err := client.GetAccessControls(ctx, model.GetEffectiveProjectID(), model.Resource.ValueString(), resourceID)
	if err != nil {
		return httpclient.AccessControl{}, status, err
	}

	// Find the matching access control by comparing composite IDs
	projectID := model.GetEffectiveProjectID()
	targetID := model.buildExpectedID(projectID)
	for _, ac := range result.AccessControls {
		if ac.BuildCompositeID(projectID) == targetID {
			return ac, status, nil
		}
	}

	return httpclient.AccessControl{}, http.StatusNotFound, fmt.Errorf("access control not found for %s", targetID)
}

func (o AccessControlOps) Update(ctx context.Context, client httpclient.PosthogClient, model AccessControlTFModel, req httpclient.AccessControlRequest) (httpclient.AccessControl, httpclient.HTTPStatusCode, error) {
	result, err := client.UpsertAccessControl(ctx, model.GetEffectiveProjectID(), req)
	return result, http.StatusOK, err
}

func (o AccessControlOps) Delete(ctx context.Context, client httpclient.PosthogClient, model AccessControlTFModel) (httpclient.HTTPStatusCode, error) {
	req := buildAccessControlRequest(model)
	return client.DeleteAccessControl(ctx, model.GetEffectiveProjectID(), req)
}

func buildAccessControlRequest(model AccessControlTFModel) httpclient.AccessControlRequest {
	accessLevel := model.AccessLevel.ValueString()
	req := httpclient.AccessControlRequest{
		AccessLevel: &accessLevel,
		Resource:    model.Resource.ValueString(),
	}

	if !model.ResourceID.IsNull() && !model.ResourceID.IsUnknown() && model.ResourceID.ValueString() != "" {
		resourceID := model.ResourceID.ValueString()
		req.ResourceID = &resourceID
	}

	if !model.Role.IsNull() && !model.Role.IsUnknown() {
		role := model.Role.ValueString()
		req.Role = &role
	}

	if !model.OrganizationMember.IsNull() && !model.OrganizationMember.IsUnknown() {
		member := model.OrganizationMember.ValueString()
		req.OrganizationMember = &member
	}

	return req
}
