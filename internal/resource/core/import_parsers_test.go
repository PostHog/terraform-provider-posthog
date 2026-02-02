package core

import (
	"testing"

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
