package core

import (
	"testing"

	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testOrganizationScopedModel struct {
	BaseStringIdentifiable
	BaseOrganizationID
}

func TestOrganizationScopedImportParser(t *testing.T) {
	parser := OrganizationScopedImportParser[testOrganizationScopedModel]()

	tests := map[string]struct {
		importID       string
		wantErr        bool
		wantOrgID      string
		wantResourceID string
	}{
		"valid two-part format": {
			importID:       "org-123/resource-456",
			wantOrgID:      "org-123",
			wantResourceID: "resource-456",
		},
		"multi-parts includes slash in resource ID": {
			importID:       "org-123/resource/with/slashes",
			wantOrgID:      "org-123",
			wantResourceID: "resource/with/slashes",
		},
		"one part is invalid": {
			importID: "resource-456",
			wantErr:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model, err := parser(tc.importID, ProviderDefaults{})

			if tc.wantErr {
				require.Error(t, err, "expected parser to return an error for invalid import ID format")
				return
			}

			require.NoError(t, err, "parser should successfully parse valid import ID")
			assert.Equal(t, tc.wantOrgID, model.OrganizationID.ValueString(), "organization_id should be extracted from first segment")
			assert.Equal(t, tc.wantResourceID, model.ID.ValueString(), "resource_id should be extracted from remaining segments")
		})
	}
}

type testRoleMembershipModel struct {
	BaseStringIdentifiable
	BaseOrganizationID
	BaseRoleID
}

type testProjectSingletonModel struct {
	BaseProjectID
}

func (m testProjectSingletonModel) GetID() string {
	return m.GetEffectiveProjectID()
}

func (m testProjectSingletonModel) HasValidID() bool {
	return m.GetEffectiveProjectID() != ""
}

type testProjectMemberModel struct {
	BaseProjectID
	Role               string
	OrganizationMember string
}

func (m testProjectMemberModel) GetID() string {
	return ""
}

func (m testProjectMemberModel) HasValidID() bool {
	return m.Role != "" || m.OrganizationMember != ""
}

func (m *testProjectMemberModel) SetProjectMemberFields(targetType, targetID string) {
	switch targetType {
	case httpclient.AccessControlTargetRole:
		m.Role = targetID
	case httpclient.AccessControlTargetMember:
		m.OrganizationMember = targetID
	}
}

func TestRoleMembershipImportParser(t *testing.T) {
	parser := RoleMembershipImportParser[testRoleMembershipModel]()

	tests := map[string]struct {
		importID         string
		wantErr          bool
		wantOrgID        string
		wantRoleID       string
		wantMembershipID string
	}{
		"valid three-part format": {
			importID:         "org-123/role-456/membership-789",
			wantOrgID:        "org-123",
			wantRoleID:       "role-456",
			wantMembershipID: "membership-789",
		},
		"empty parts should parse": {
			importID:         "//",
			wantOrgID:        "",
			wantRoleID:       "",
			wantMembershipID: "",
		},
		"two parts is invalid": {
			importID: "org-123/role-456",
			wantErr:  true,
		},
		"one part is invalid": {
			importID: "membership-789",
			wantErr:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model, err := parser(tc.importID, ProviderDefaults{})

			if tc.wantErr {
				require.Error(t, err, "expected parser to return an error for invalid import ID format")
				return
			}

			require.NoError(t, err, "parser should successfully parse valid import ID")
			assert.Equal(t, tc.wantOrgID, model.OrganizationID.ValueString(), "organization_id should be extracted from first segment")
			assert.Equal(t, tc.wantRoleID, model.RoleID.ValueString(), "role_id should be extracted from second segment")
			assert.Equal(t, tc.wantMembershipID, model.ID.ValueString(), "membership_id should be extracted from third segment")
		})
	}
}

func TestProjectSingletonImportParser(t *testing.T) {
	parser := ProjectSingletonImportParser[testProjectSingletonModel]()

	tests := map[string]struct {
		importID      string
		wantProjectID string
	}{
		"simple project ID": {
			importID:      "project-123",
			wantProjectID: "project-123",
		},
		"UUID format": {
			importID:      "019a90e1-4e8e-0000-4fed-6d6380020c63",
			wantProjectID: "019a90e1-4e8e-0000-4fed-6d6380020c63",
		},
		"numeric ID": {
			importID:      "12345",
			wantProjectID: "12345",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model, err := parser(tc.importID, ProviderDefaults{})

			require.NoError(t, err, "parser should successfully parse import ID")
			assert.Equal(t, tc.wantProjectID, model.GetEffectiveProjectID(), "project_id should be set from import ID")
		})
	}
}

type testAccessControlModel struct {
	BaseProjectID
	Resource           string
	ResourceID         string
	TargetType         string
	Role               string
	OrganizationMember string
}

func (m testAccessControlModel) GetID() string {
	return ""
}

func (m testAccessControlModel) HasValidID() bool {
	return m.Resource != ""
}

func (m *testAccessControlModel) SetAccessControlFields(resourceType, resourceID, targetType, targetID string) {
	m.Resource = resourceType
	m.ResourceID = resourceID
	m.TargetType = targetType
	switch targetType {
	case httpclient.AccessControlTargetRole:
		m.Role = targetID
	case httpclient.AccessControlTargetMember:
		m.OrganizationMember = targetID
	}
}

func TestAccessControlImportParser(t *testing.T) {
	parser := AccessControlImportParser[testAccessControlModel]()

	tests := map[string]struct {
		importID       string
		wantErr        bool
		wantProjectID  string
		wantResource   string
		wantResID      string
		wantTargetType string
		wantRole       string
		wantMember     string
	}{
		// 3-part: project default for resource type
		"project default for resource type": {
			importID:       "project-123/survey/default",
			wantProjectID:  "project-123",
			wantResource:   "survey",
			wantTargetType: httpclient.AccessControlTargetDefault,
		},
		// 4-part: project default for specific resource
		"project default for specific resource": {
			importID:       "project-123/dashboard/999/default",
			wantProjectID:  "project-123",
			wantResource:   "dashboard",
			wantResID:      "999",
			wantTargetType: httpclient.AccessControlTargetDefault,
		},
		// 4-part: role for resource type
		"role for resource type": {
			importID:       "project-123/feature_flag/role/role-456",
			wantProjectID:  "project-123",
			wantResource:   "feature_flag",
			wantTargetType: httpclient.AccessControlTargetRole,
			wantRole:       "role-456",
		},
		// 4-part: member for resource type
		"member for resource type": {
			importID:       "project-123/dashboard/member/member-789",
			wantProjectID:  "project-123",
			wantResource:   "dashboard",
			wantTargetType: httpclient.AccessControlTargetMember,
			wantMember:     "member-789",
		},
		// 5-part: role for specific resource
		"role for specific resource": {
			importID:       "project-123/dashboard/999/role/role-456",
			wantProjectID:  "project-123",
			wantResource:   "dashboard",
			wantResID:      "999",
			wantTargetType: httpclient.AccessControlTargetRole,
			wantRole:       "role-456",
		},
		// 5-part: member for specific resource
		"member for specific resource": {
			importID:       "project-123/dashboard/999/member/member-789",
			wantProjectID:  "project-123",
			wantResource:   "dashboard",
			wantResID:      "999",
			wantTargetType: httpclient.AccessControlTargetMember,
			wantMember:     "member-789",
		},
		// UUID format
		"UUID format": {
			importID:       "019a90e1-4e8e-0000-4fed-6d6380020c63/experiment/role/019b1234-5678-0000-abcd-ef1234567890",
			wantProjectID:  "019a90e1-4e8e-0000-4fed-6d6380020c63",
			wantResource:   "experiment",
			wantTargetType: httpclient.AccessControlTargetRole,
			wantRole:       "019b1234-5678-0000-abcd-ef1234567890",
		},
		// Error cases
		"3-part without default is invalid": {
			importID: "project-123/survey/viewer",
			wantErr:  true,
		},
		"invalid target type": {
			importID: "project-123/dashboard/admin/user-456",
			wantErr:  true,
		},
		"too few parts": {
			importID: "project-123/dashboard",
			wantErr:  true,
		},
		"too many parts": {
			importID: "project-123/dashboard/999/role/role-456/extra",
			wantErr:  true,
		},
		"one part is invalid": {
			importID: "project-123",
			wantErr:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model, err := parser(tc.importID, ProviderDefaults{})

			if tc.wantErr {
				require.Error(t, err, "expected parser to return an error for invalid import ID format")
				return
			}

			require.NoError(t, err, "parser should successfully parse valid import ID")
			assert.Equal(t, tc.wantProjectID, model.GetEffectiveProjectID(), "project_id should be extracted")
			assert.Equal(t, tc.wantResource, model.Resource, "resource should be extracted")
			assert.Equal(t, tc.wantResID, model.ResourceID, "resource_id should be extracted when present")
			assert.Equal(t, tc.wantTargetType, model.TargetType, "target_type should be extracted")
			assert.Equal(t, tc.wantRole, model.Role, "role should be extracted when target type is 'role'")
			assert.Equal(t, tc.wantMember, model.OrganizationMember, "organization_member should be extracted when target type is 'member'")
		})
	}
}

func TestProjectMemberImportParser(t *testing.T) {
	parser := ProjectMemberImportParser[testProjectMemberModel]()

	tests := map[string]struct {
		importID      string
		wantErr       bool
		wantProjectID string
		wantRole      string
		wantMember    string
	}{
		"role format": {
			importID:      "project-123/role/role-456",
			wantProjectID: "project-123",
			wantRole:      "role-456",
		},
		"member format": {
			importID:      "project-123/member/member-789",
			wantProjectID: "project-123",
			wantMember:    "member-789",
		},
		"UUID format": {
			importID:      "019a90e1-4e8e-0000-4fed-6d6380020c63/role/019b1234-5678-0000-abcd-ef1234567890",
			wantProjectID: "019a90e1-4e8e-0000-4fed-6d6380020c63",
			wantRole:      "019b1234-5678-0000-abcd-ef1234567890",
		},
		"invalid target type": {
			importID: "project-123/admin/user-456",
			wantErr:  true,
		},
		"too few parts": {
			importID: "project-123/role",
			wantErr:  true,
		},
		"too many parts": {
			importID: "project-123/role/role-456/extra",
			wantErr:  true,
		},
		"one part is invalid": {
			importID: "project-123",
			wantErr:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model, err := parser(tc.importID, ProviderDefaults{})

			if tc.wantErr {
				require.Error(t, err, "expected parser to return an error for invalid import ID format")
				return
			}

			require.NoError(t, err, "parser should successfully parse valid import ID")
			assert.Equal(t, tc.wantProjectID, model.GetEffectiveProjectID(), "project_id should be extracted from first segment")
			assert.Equal(t, tc.wantRole, model.Role, "role should be extracted when target type is 'role'")
			assert.Equal(t, tc.wantMember, model.OrganizationMember, "organization_member should be extracted when target type is 'member'")
		})
	}
}
