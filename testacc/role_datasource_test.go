package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// TestRoleDataSource_Basic creates a role via the API (simulating SCIM provisioning),
// then looks it up by name using only the data source.
func TestRoleDataSource_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role-ds")
	orgID := getOrganizationID()

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)

	// Create the role via the API, simulating external creation (e.g. SCIM).
	role, err := client.CreateRole(context.Background(), orgID, httpclient.RoleRequest{Name: rName})
	if err != nil {
		t.Fatalf("failed to create role via API: %s", err)
	}
	t.Cleanup(func() {
		_, _ = client.DeleteRole(context.Background(), orgID, role.ID)
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccRolePreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleDataSourceConfig(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.posthog_role.test", "organization_id", orgID),
					resource.TestCheckResourceAttr("data.posthog_role.test", "name", rName),
					resource.TestCheckResourceAttr("data.posthog_role.test", "id", role.ID),
				),
			},
		},
	})
}

func testAccRoleDataSourceConfig(orgID, name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

data "posthog_role" "test" {
  organization_id = %q
  name            = %q
}
`, orgID, name)
}
