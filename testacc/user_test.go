package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccUserPreCheck(t *testing.T) {
	t.Helper()
	testAccBasePreCheck(t)
	skipIfNoOrganizationID(t)
	skipIfNoTestUserEmail(t)
}

// TestUserDataSource_Basic tests the user data source.
func TestUserDataSource_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccUserPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSource(orgID, userEmail),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.posthog_user.test", "organization_id", orgID),
					resource.TestCheckResourceAttr("data.posthog_user.test", "email", userEmail),
					resource.TestCheckResourceAttrSet("data.posthog_user.test", "uuid"),
				),
			},
		},
	})
}

func testAccUserDataSource(orgID, userEmail string) string {
	return fmt.Sprintf(`
provider "posthog" {}

data "posthog_user" "test" {
  organization_id = %q
  email           = %q
}
`, orgID, userEmail)
}
