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
				Resource: "feature_flags",
				Role:     strPtr("role-456"),
			},
			expected: "proj-123/feature_flags/role/role-456",
		},
		"role with resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:   "dashboards",
				ResourceID: strPtr("dash-789"),
				Role:       strPtr("role-456"),
			},
			expected: "proj-123/dashboards/dash-789/role/role-456",
		},
		"organization_member without resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:           "experiments",
				OrganizationMember: strPtr("member-abc"),
			},
			expected: "proj-123/experiments/member/member-abc",
		},
		"organization_member with resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:           "dashboards",
				ResourceID:         strPtr("dash-789"),
				OrganizationMember: strPtr("member-abc"),
			},
			expected: "proj-123/dashboards/dash-789/member/member-abc",
		},
		"empty resource_id is ignored": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:   "feature_flags",
				ResourceID: strPtr(""),
				Role:       strPtr("role-456"),
			},
			expected: "proj-123/feature_flags/role/role-456",
		},
		"nil role and member returns empty": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource: "dashboards",
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
