package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessControl_BuildCompositeID(t *testing.T) {
	tests := map[string]struct {
		projectID string
		ac        AccessControl
		expected  string
	}{
		"role without resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource: "feature_flag",
				Role:     strPtr("role-456"),
			},
			expected: "proj-123/feature_flag/role/role-456",
		},
		"role with resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:   "dashboard",
				ResourceID: strPtr("dash-789"),
				Role:       strPtr("role-456"),
			},
			expected: "proj-123/dashboard/dash-789/role/role-456",
		},
		"organization_member without resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:           "experiment",
				OrganizationMember: strPtr("member-abc"),
			},
			expected: "proj-123/experiment/member/member-abc",
		},
		"organization_member with resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:           "dashboard",
				ResourceID:         strPtr("dash-789"),
				OrganizationMember: strPtr("member-abc"),
			},
			expected: "proj-123/dashboard/dash-789/member/member-abc",
		},
		"empty resource_id is ignored": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:   "feature_flag",
				ResourceID: strPtr(""),
				Role:       strPtr("role-456"),
			},
			expected: "proj-123/feature_flag/role/role-456",
		},
		"nil role and member returns empty": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource: "dashboard",
			},
			expected: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := tc.ac.BuildCompositeID(tc.projectID)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func strPtr(s string) *string {
	return &s
}
