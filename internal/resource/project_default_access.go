package resource

import (
	"context"

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
)

// Project access levels (different from resource-type access levels).
const (
	projectAccessLevelNone   = "none"
	projectAccessLevelMember = "member"
	projectAccessLevelAdmin  = "admin"
)

var projectAccessLevels = []string{projectAccessLevelNone, projectAccessLevelMember, projectAccessLevelAdmin}

func NewProjectDefaultAccess() resource.Resource {
	return core.NewGenericResource[ProjectDefaultAccessModel, httpclient.ProjectAccessControlRequest, httpclient.ProjectAccessControlListResponse](
		ProjectDefaultAccessOps{},
		core.ProjectSingletonImportParser[ProjectDefaultAccessModel](),
	)
}

// ProjectDefaultAccessModel manages the default access level for a project.
// This is a singleton per project - there's only one default access level.
type ProjectDefaultAccessModel struct {
	core.BaseProjectID
	// ID is the project_id since there's only one default per project
	ID          types.String `tfsdk:"id"`
	AccessLevel types.String `tfsdk:"access_level"`
}

func (m ProjectDefaultAccessModel) GetID() string {
	return m.ID.ValueString()
}

func (m ProjectDefaultAccessModel) HasValidID() bool {
	return m.GetEffectiveProjectID() != ""
}

type ProjectDefaultAccessOps struct{}

func (o ProjectDefaultAccessOps) ResourceName() string {
	return "project_default_access"
}

func (o ProjectDefaultAccessOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: `Sets the default access level for a PostHog project.

This determines the baseline access level for all organization members who don't have explicit access granted via roles or direct membership.

~> **Note:** There is only one default access level per project. This resource is a singleton - creating multiple instances for the same project will cause conflicts.

~> **Destroy Behavior:** Destroying this resource resets the default access level to ` + "`none`" + ` (most restrictive). It does not restore the previous value.` + core.EnterpriseRBACNote,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The project ID (used as the resource identifier).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"access_level": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The default access level for the project. Valid values are `none`, `member`, or `admin`.",
				Validators: []validator.String{
					stringvalidator.OneOf(projectAccessLevels...),
				},
			},
		},
	}
}

func (o ProjectDefaultAccessOps) BuildCreateRequest(_ context.Context, model ProjectDefaultAccessModel) (httpclient.ProjectAccessControlRequest, diag.Diagnostics) {
	accessLevel := model.AccessLevel.ValueString()
	return httpclient.ProjectAccessControlRequest{
		AccessLevel: &accessLevel,
	}, nil
}

func (o ProjectDefaultAccessOps) BuildUpdateRequest(ctx context.Context, plan, _ ProjectDefaultAccessModel) (httpclient.ProjectAccessControlRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o ProjectDefaultAccessOps) MapResponseToModel(_ context.Context, resp httpclient.ProjectAccessControlListResponse, model *ProjectDefaultAccessModel) diag.Diagnostics {
	model.ID = types.StringValue(model.GetEffectiveProjectID())

	// Find the "everyone" access control entry (no role or organization_member)
	for _, ac := range resp.AccessControls {
		if ac.Role == nil && ac.OrganizationMember == nil && ac.AccessLevel != nil {
			model.AccessLevel = types.StringValue(*ac.AccessLevel)
			return nil
		}
	}

	// If no explicit "everyone" entry exists, use the system default
	model.AccessLevel = types.StringValue(resp.DefaultAccessLevel)
	return nil
}

func (o ProjectDefaultAccessOps) Create(ctx context.Context, client httpclient.PosthogClient, model ProjectDefaultAccessModel, req httpclient.ProjectAccessControlRequest) (httpclient.ProjectAccessControlListResponse, error) {
	_, _, err := client.UpsertProjectAccessControl(ctx, model.GetEffectiveProjectID(), req)
	if err != nil {
		return httpclient.ProjectAccessControlListResponse{}, err
	}
	// Return the full response by reading back
	result, _, err := client.GetProjectAccessControls(ctx, model.GetEffectiveProjectID())
	return result, err
}

func (o ProjectDefaultAccessOps) Read(ctx context.Context, client httpclient.PosthogClient, model ProjectDefaultAccessModel) (httpclient.ProjectAccessControlListResponse, httpclient.HTTPStatusCode, error) {
	return client.GetProjectAccessControls(ctx, model.GetEffectiveProjectID())
}

func (o ProjectDefaultAccessOps) Update(ctx context.Context, client httpclient.PosthogClient, model ProjectDefaultAccessModel, req httpclient.ProjectAccessControlRequest) (httpclient.ProjectAccessControlListResponse, httpclient.HTTPStatusCode, error) {
	_, status, err := client.UpsertProjectAccessControl(ctx, model.GetEffectiveProjectID(), req)
	if err != nil {
		return httpclient.ProjectAccessControlListResponse{}, status, err
	}
	return client.GetProjectAccessControls(ctx, model.GetEffectiveProjectID())
}

func (o ProjectDefaultAccessOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ProjectDefaultAccessModel) (httpclient.HTTPStatusCode, error) {
	// Can't truly delete a default - reset to "none" (most restrictive)
	noneLevel := projectAccessLevelNone
	req := httpclient.ProjectAccessControlRequest{
		AccessLevel: &noneLevel,
	}
	_, status, err := client.UpsertProjectAccessControl(ctx, model.GetEffectiveProjectID(), req)
	return status, err
}
