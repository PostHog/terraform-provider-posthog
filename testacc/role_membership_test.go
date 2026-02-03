package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccRoleMembershipPreCheck(t *testing.T) {
	t.Helper()
	testAccBasePreCheck(t)
	skipIfNoOrganizationID(t)
	skipIfNoTestUserEmail(t)
}

// TestRoleMembership_Basic tests creating a role membership.
func TestRoleMembership_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccRoleMembershipPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleMembershipBasic(orgID, rName, userEmail),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_role_membership.test", "organization_id", orgID),
					resource.TestCheckResourceAttrSet("posthog_role_membership.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_role_membership.test", "role_id"),
					resource.TestCheckResourceAttrSet("posthog_role_membership.test", "user_uuid"),
					resource.TestCheckResourceAttrSet("posthog_role_membership.test", "joined_at"),
				),
			},
		},
	})
}

// TestRoleMembership_Import tests importing an existing role membership.
func TestRoleMembership_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccRoleMembershipPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleMembershipBasic(orgID, rName, userEmail),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_role_membership.test", "id"),
				),
			},
			{
				ResourceName:      "posthog_role_membership.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccRoleMembershipImportStateIdFunc("posthog_role_membership.test"),
			},
		},
	})
}

func testAccRoleMembershipImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["organization_id"]
		roleID := rs.Primary.Attributes["role_id"]
		membershipID := rs.Primary.Attributes["id"]

		return fmt.Sprintf("%s/%s/%s", orgID, roleID, membershipID), nil
	}
}

func testAccRoleMembershipBasic(orgID, roleName, userEmail string) string {
	return fmt.Sprintf(`
provider "posthog" {}

data "posthog_user" "test" {
  organization_id = %q
  email           = %q
}

resource "posthog_role" "test" {
  organization_id = %q
  name            = %q
}

resource "posthog_role_membership" "test" {
  organization_id = %q
  role_id         = posthog_role.test.id
  user_uuid       = data.posthog_user.test.uuid
}
`, orgID, userEmail, orgID, roleName, orgID)
}
