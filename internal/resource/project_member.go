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

func NewProjectMember() resource.Resource {
	return core.NewGenericResource[ProjectMemberModel, httpclient.ProjectAccessControlRequest, httpclient.ProjectAccessControl](
		ProjectMemberOps{},
		core.ProjectMemberImportParser[ProjectMemberModel](),
	)
}

// ProjectMemberModel manages a role or organization member's access to a project.
type ProjectMemberModel struct {
	core.BaseProjectID
	// Composite ID: project_id/role/<role_id> or project_id/member/<member_id>
	ID                 types.String `tfsdk:"id"`
	Role               types.String `tfsdk:"role"`
	OrganizationMember types.String `tfsdk:"organization_member"`
	AccessLevel        types.String `tfsdk:"access_level"`
	CreatedAt          types.String `tfsdk:"created_at"`
	UpdatedAt          types.String `tfsdk:"updated_at"`
}

func (m ProjectMemberModel) GetID() string {
	return m.ID.ValueString()
}

func (m ProjectMemberModel) HasValidID() bool {
	hasRole := !m.Role.IsNull() && !m.Role.IsUnknown() && m.Role.ValueString() != ""
	hasMember := !m.OrganizationMember.IsNull() && !m.OrganizationMember.IsUnknown() && m.OrganizationMember.ValueString() != ""
	return hasRole || hasMember
}

func (m *ProjectMemberModel) SetID(id string) error {
	m.ID = types.StringValue(id)
	return nil
}

func (m ProjectMemberModel) buildCompositeID(projectID string) string {
	role := ""
	if !m.Role.IsNull() && !m.Role.IsUnknown() {
		role = m.Role.ValueString()
	}
	member := ""
	if !m.OrganizationMember.IsNull() && !m.OrganizationMember.IsUnknown() {
		member = m.OrganizationMember.ValueString()
	}
	return httpclient.BuildProjectMemberCompositeID(projectID, role, member)
}

// SetProjectMemberFields sets the role or organization_member field based on import parsing.
func (m *ProjectMemberModel) SetProjectMemberFields(targetType, targetID string) {
	switch targetType {
	case core.AccessControlTargetRole:
		m.Role = types.StringValue(targetID)
	case core.AccessControlTargetMember:
		m.OrganizationMember = types.StringValue(targetID)
	}
}

type ProjectMemberOps struct{}

func (o ProjectMemberOps) ResourceName() string {
	return "project_member"
}

func (o ProjectMemberOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: `Manages a role or organization member's access to a PostHog project.

This resource grants a specific role or organization member access to a project with a defined access level.

~> **Note:** You must specify either ` + "`role`" + ` or ` + "`organization_member`" + `, but not both.` + core.EnterpriseRBACNote,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Composite identifier for the project member access.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"role": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The UUID of the role to grant project access to. Mutually exclusive with `organization_member`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRoot("organization_member"),
					}...),
				},
			},
			"organization_member": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The UUID of the organization member to grant project access to. Mutually exclusive with `role`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"access_level": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The access level to grant. Valid values are `none`, `member`, or `admin`.",
				Validators: []validator.String{
					stringvalidator.OneOf(ProjectAccessLevels...),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the access was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the access was last updated.",
			},
		},
	}
}

func (o ProjectMemberOps) BuildCreateRequest(_ context.Context, model ProjectMemberModel) (httpclient.ProjectAccessControlRequest, diag.Diagnostics) {
	return buildProjectMemberRequest(model), nil
}

func (o ProjectMemberOps) BuildUpdateRequest(_ context.Context, plan, _ ProjectMemberModel) (httpclient.ProjectAccessControlRequest, diag.Diagnostics) {
	return buildProjectMemberRequest(plan), nil
}

func (o ProjectMemberOps) MapResponseToModel(_ context.Context, resp httpclient.ProjectAccessControl, model *ProjectMemberModel) diag.Diagnostics {
	model.ID = types.StringValue(model.buildCompositeID(model.GetEffectiveProjectID()))

	if resp.AccessLevel != nil {
		model.AccessLevel = types.StringValue(*resp.AccessLevel)
	}
	if resp.Role != nil {
		model.Role = types.StringValue(*resp.Role)
	}
	if resp.OrganizationMember != nil {
		model.OrganizationMember = types.StringValue(*resp.OrganizationMember)
	}
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.UpdatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.UpdatedAt)

	return nil
}

func (o ProjectMemberOps) Create(ctx context.Context, client httpclient.PosthogClient, model ProjectMemberModel, req httpclient.ProjectAccessControlRequest) (httpclient.ProjectAccessControl, error) {
	result, _, err := client.UpsertProjectAccessControl(ctx, model.GetEffectiveProjectID(), req)
	return result, err
}

func (o ProjectMemberOps) Read(ctx context.Context, client httpclient.PosthogClient, model ProjectMemberModel) (httpclient.ProjectAccessControl, httpclient.HTTPStatusCode, error) {
	result, status, err := client.GetProjectAccessControls(ctx, model.GetEffectiveProjectID())
	if err != nil {
		return httpclient.ProjectAccessControl{}, status, err
	}

	// Find the matching access control by comparing composite IDs
	projectID := model.GetEffectiveProjectID()
	expectedID := model.buildCompositeID(projectID)
	for _, ac := range result.AccessControls {
		if ac.BuildCompositeID(projectID) == expectedID {
			return ac, status, nil
		}
	}

	return httpclient.ProjectAccessControl{}, http.StatusNotFound, fmt.Errorf("project member access not found for %s", expectedID)
}

func (o ProjectMemberOps) Update(ctx context.Context, client httpclient.PosthogClient, model ProjectMemberModel, req httpclient.ProjectAccessControlRequest) (httpclient.ProjectAccessControl, httpclient.HTTPStatusCode, error) {
	return client.UpsertProjectAccessControl(ctx, model.GetEffectiveProjectID(), req)
}

func (o ProjectMemberOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ProjectMemberModel) (httpclient.HTTPStatusCode, error) {
	req := buildProjectMemberRequest(model)
	return client.DeleteProjectAccessControl(ctx, model.GetEffectiveProjectID(), req)
}

func buildProjectMemberRequest(model ProjectMemberModel) httpclient.ProjectAccessControlRequest {
	accessLevel := model.AccessLevel.ValueString()
	req := httpclient.ProjectAccessControlRequest{
		AccessLevel: &accessLevel,
	}

	if !model.Role.IsNull() && !model.Role.IsUnknown() && model.Role.ValueString() != "" {
		role := model.Role.ValueString()
		req.Role = &role
	}

	if !model.OrganizationMember.IsNull() && !model.OrganizationMember.IsUnknown() && model.OrganizationMember.ValueString() != "" {
		member := model.OrganizationMember.ValueString()
		req.OrganizationMember = &member
	}

	return req
}
