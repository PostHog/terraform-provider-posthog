package core

import (
	"fmt"
	"strings"

	"github.com/posthog/terraform-provider/internal/httpclient"
)

// ProjectScopedImportParser returns an import parser for project-scoped resources.
// Import format: "project_id/resource_id" or just "resource_id" (uses default project_id).
func ProjectScopedImportParser[TFModel Identifiable]() ImportIDParser[TFModel] {
	return func(importID string, defaults ProviderDefaults) (TFModel, error) {
		var model TFModel

		projectID := defaults.ProjectID
		resourceID := importID

		if parts := strings.SplitN(importID, "/", 2); len(parts) == 2 {
			projectID = parts[0]
			resourceID = parts[1]
		} else if projectID == "" {
			return model, fmt.Errorf(
				"no project_id specified: use import ID format 'project_id/resource_id' or set project_id in the provider configuration",
			)
		}

		setter, ok := any(&model).(IDSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement IDSetter", model)
		}
		if err := setter.SetID(resourceID); err != nil {
			return model, fmt.Errorf("invalid resource ID %q: %w", resourceID, err)
		}

		init, ok := any(&model).(ProjectIDInitializer)
		if !ok {
			return model, fmt.Errorf("model %T does not implement ProjectIDInitializer", model)
		}
		init.InitializeProjectID(projectID)

		return model, nil
	}
}

// OrganizationScopedImportParser returns an import parser for organization-scoped resources.
// Import format: "organization_id/resource_id" (organization_id is mandatory).
func OrganizationScopedImportParser[TFModel Identifiable]() ImportIDParser[TFModel] {
	return func(importID string, defaults ProviderDefaults) (TFModel, error) {
		var model TFModel

		parts := strings.SplitN(importID, "/", 2)
		if len(parts) != 2 {
			return model, fmt.Errorf(
				"invalid import ID format %q: expected 'organization_id/resource_id'", importID,
			)
		}

		organizationID := parts[0]
		resourceID := parts[1]

		setter, ok := any(&model).(IDSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement IDSetter", model)
		}
		if err := setter.SetID(resourceID); err != nil {
			return model, fmt.Errorf("invalid resource ID %q: %w", resourceID, err)
		}

		orgInit, ok := any(&model).(OrganizationIDInitializer)
		if !ok {
			return model, fmt.Errorf("model %T does not implement OrganizationIDInitializer", model)
		}
		orgInit.InitializeOrganizationID(organizationID)

		return model, nil
	}
}

// OrganizationMemberImportParser returns an import parser for organization member resources.
// Import format: "organization_id/user_uuid".
func OrganizationMemberImportParser[TFModel Identifiable]() ImportIDParser[TFModel] {
	return func(importID string, defaults ProviderDefaults) (TFModel, error) {
		var model TFModel

		parts := strings.SplitN(importID, "/", 2)
		if len(parts) != 2 {
			return model, fmt.Errorf(
				"invalid import ID format %q: expected 'organization_id/user_uuid'", importID,
			)
		}

		organizationID := parts[0]
		userUUID := parts[1]

		orgInit, ok := any(&model).(OrganizationIDInitializer)
		if !ok {
			return model, fmt.Errorf("model %T does not implement OrganizationIDInitializer", model)
		}
		orgInit.InitializeOrganizationID(organizationID)

		userSetter, ok := any(&model).(UserUUIDSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement UserUUIDSetter", model)
		}
		userSetter.SetUserUUID(userUUID)

		return model, nil
	}
}

// RoleMembershipImportParser returns an import parser for role membership resources.
// Import format: "organization_id/role_id/membership_id".
func RoleMembershipImportParser[TFModel Identifiable]() ImportIDParser[TFModel] {
	return func(importID string, defaults ProviderDefaults) (TFModel, error) {
		var model TFModel

		parts := strings.SplitN(importID, "/", 3)
		if len(parts) != 3 {
			return model, fmt.Errorf(
				"invalid import ID format %q: expected 'organization_id/role_id/membership_id'", importID,
			)
		}

		organizationID := parts[0]
		roleID := parts[1]
		membershipID := parts[2]

		orgInit, ok := any(&model).(OrganizationIDInitializer)
		if !ok {
			return model, fmt.Errorf("model %T does not implement OrganizationIDInitializer", model)
		}
		orgInit.InitializeOrganizationID(organizationID)

		roleSetter, ok := any(&model).(RoleIDSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement RoleIDSetter", model)
		}
		roleSetter.SetRoleID(roleID)

		setter, ok := any(&model).(IDSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement IDSetter", model)
		}
		if err := setter.SetID(membershipID); err != nil {
			return model, fmt.Errorf("invalid membership ID %q: %w", membershipID, err)
		}

		return model, nil
	}
}

// AccessControlFieldsSetter is implemented by models that need access control fields set during import.
type AccessControlFieldsSetter interface {
	SetAccessControlFields(resourceType, resourceID, targetType, targetID string)
}

// ProjectMemberFieldsSetter is implemented by models that need project member fields set during import.
type ProjectMemberFieldsSetter interface {
	SetProjectMemberFields(targetType, targetID string)
}

// ProjectSingletonImportParser returns an import parser for singleton project-scoped resources.
// These resources have one instance per project (e.g., project_default_access).
// Import format: just "project_id" (no resource_id needed).
func ProjectSingletonImportParser[TFModel Identifiable]() ImportIDParser[TFModel] {
	return func(importID string, _ ProviderDefaults) (TFModel, error) {
		var model TFModel

		init, ok := any(&model).(ProjectIDInitializer)
		if !ok {
			return model, fmt.Errorf("model %T does not implement ProjectIDInitializer", model)
		}
		init.InitializeProjectID(importID)

		return model, nil
	}
}

// ProjectMemberImportParser returns an import parser for project member resources.
//
// Import formats:
//   - project_id/role/<role_id>
//   - project_id/member/<member_id>
func ProjectMemberImportParser[TFModel Identifiable]() ImportIDParser[TFModel] {
	return func(importID string, defaults ProviderDefaults) (TFModel, error) {
		var model TFModel

		parts := strings.Split(importID, "/")
		if len(parts) != 3 {
			return model, fmt.Errorf(
				"invalid import ID format %q: expected 'project_id/role/<role_id>' or 'project_id/member/<member_id>'", importID,
			)
		}

		projectID := parts[0]
		targetType := parts[1]
		targetID := parts[2]

		if targetType != httpclient.AccessControlTargetRole && targetType != httpclient.AccessControlTargetMember {
			return model, fmt.Errorf(
				"invalid target type %q in import ID: expected %q or %q", targetType, httpclient.AccessControlTargetRole, httpclient.AccessControlTargetMember,
			)
		}

		init, ok := any(&model).(ProjectIDInitializer)
		if !ok {
			return model, fmt.Errorf("model %T does not implement ProjectIDInitializer", model)
		}
		init.InitializeProjectID(projectID)

		pmSetter, ok := any(&model).(ProjectMemberFieldsSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement ProjectMemberFieldsSetter", model)
		}
		pmSetter.SetProjectMemberFields(targetType, targetID)

		return model, nil
	}
}

// AccessControlImportParser returns an import parser for access control resources.
//
// Import formats:
//   - project_id/resource_type/default (project default for resource type)
//   - project_id/resource_type/resource_id/default (project default for specific resource)
//   - project_id/resource_type/role/role_id
//   - project_id/resource_type/member/member_id
//   - project_id/resource_type/resource_id/role/role_id
//   - project_id/resource_type/resource_id/member/member_id
func AccessControlImportParser[TFModel Identifiable]() ImportIDParser[TFModel] {
	return func(importID string, defaults ProviderDefaults) (TFModel, error) {
		var model TFModel

		parts := strings.Split(importID, "/")

		var projectID, resourceType, resourceID, targetType, targetID string

		// 3 parts: project_id/resource/default (project default, no resource_id)
		// 4 parts: project_id/resource/role|member/target_id (no resource_id)
		//      OR: project_id/resource/resource_id/default (project default with resource_id)
		// 5 parts: project_id/resource/resource_id/role|member/target_id (with resource_id)
		switch len(parts) {
		case 3:
			// project_id/resource/default
			projectID = parts[0]
			resourceType = parts[1]
			targetType = parts[2]
			if targetType != httpclient.AccessControlTargetDefault {
				return model, fmt.Errorf(
					"invalid import ID format %q: 3-part format must end with '%s'", importID, httpclient.AccessControlTargetDefault,
				)
			}
		case 4:
			projectID = parts[0]
			resourceType = parts[1]
			// Could be: resource/default (with resource_id) OR role|member/target_id (without resource_id)
			if parts[3] == httpclient.AccessControlTargetDefault {
				// project_id/resource/resource_id/default
				resourceID = parts[2]
				targetType = parts[3]
			} else {
				// project_id/resource/role|member/target_id
				targetType = parts[2]
				targetID = parts[3]
			}
		case 5:
			projectID = parts[0]
			resourceType = parts[1]
			resourceID = parts[2]
			targetType = parts[3]
			targetID = parts[4]
		default:
			return model, fmt.Errorf(
				"invalid import ID format %q: expected 'project_id/resource/default', "+
					"'project_id/resource/role/role_id', or 'project_id/resource/resource_id/role/role_id' "+
					"(use 'member' instead of 'role' for member-specific, or 'default' for project defaults)", importID,
			)
		}

		if targetType != httpclient.AccessControlTargetRole && targetType != httpclient.AccessControlTargetMember && targetType != httpclient.AccessControlTargetDefault {
			return model, fmt.Errorf(
				"invalid target type %q in import ID: expected %q, %q, or %q",
				targetType, httpclient.AccessControlTargetRole, httpclient.AccessControlTargetMember, httpclient.AccessControlTargetDefault,
			)
		}

		init, ok := any(&model).(ProjectIDInitializer)
		if !ok {
			return model, fmt.Errorf("model %T does not implement ProjectIDInitializer", model)
		}
		init.InitializeProjectID(projectID)

		acSetter, ok := any(&model).(AccessControlFieldsSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement AccessControlFieldsSetter", model)
		}
		acSetter.SetAccessControlFields(resourceType, resourceID, targetType, targetID)

		return model, nil
	}
}
