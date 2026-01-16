package core

import (
	"fmt"
	"strings"
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

		orgSetter, ok := any(&model).(OrganizationIDSetter)
		if !ok {
			return model, fmt.Errorf("model %T does not implement OrganizationIDSetter", model)
		}
		orgSetter.SetOrganizationID(organizationID)

		return model, nil
	}
}
