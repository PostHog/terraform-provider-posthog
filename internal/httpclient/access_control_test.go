package httpclient

import (
	"testing"

	"github.com/posthog/terraform-provider/internal/util"
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
				Role:     util.StringPtr("role-456"),
			},
			expected: "proj-123/feature_flag/role/role-456",
		},
		"role with resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:   "dashboard",
				ResourceID: util.StringPtr("dash-789"),
				Role:       util.StringPtr("role-456"),
			},
			expected: "proj-123/dashboard/dash-789/role/role-456",
		},
		"organization_member without resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:           "experiment",
				OrganizationMember: util.StringPtr("member-abc"),
			},
			expected: "proj-123/experiment/member/member-abc",
		},
		"organization_member with resource_id": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:           "dashboard",
				ResourceID:         util.StringPtr("dash-789"),
				OrganizationMember: util.StringPtr("member-abc"),
			},
			expected: "proj-123/dashboard/dash-789/member/member-abc",
		},
		"empty resource_id is ignored": {
			projectID: "proj-123",
			ac: AccessControl{
				Resource:   "feature_flag",
				ResourceID: util.StringPtr(""),
				Role:       util.StringPtr("role-456"),
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
