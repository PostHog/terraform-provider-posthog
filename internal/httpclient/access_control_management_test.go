package httpclient

import (
	"encoding/json"
	"testing"

	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
)

// A release is expressed as an explicit null. `omitempty` on ManagedBy would drop the key, the
// API would read the body as a claim, and a released rule would stay locked forever.
func TestAccessControlClaimRequest_ReleaseSendsExplicitNull(t *testing.T) {
	body, err := json.Marshal(AccessControlClaimRequest{
		Resource:   "dashboard",
		ResourceID: util.StringPtr("42"),
		ManagedBy:  nil,
	})

	assert.NoError(t, err)
	assert.JSONEq(t, `{"resource":"dashboard","resource_id":"42","managed_by":null}`, string(body))
}

func TestAccessControlClaimRequest_OmitsUnsetTargets(t *testing.T) {
	managedBy := ManagerTerraform
	body, err := json.Marshal(AccessControlClaimRequest{
		Resource:  "feature_flag",
		Role:      util.StringPtr("role-1"),
		ManagedBy: &managedBy,
	})

	assert.NoError(t, err)
	assert.JSONEq(t, `{"resource":"feature_flag","role":"role-1","managed_by":"terraform"}`, string(body))
}
