package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccRolePreCheck(t *testing.T) {
	t.Helper()
	testAccBasePreCheck(t)
	skipIfNoOrganizationID(t)
}

// TestRole_ProviderLevelOrganizationID tests that a role inherits organization_id from the provider.
func TestRole_ProviderLevelOrganizationID(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccRolePreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleWithProviderOrganizationID(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_role.test", "name", rName),
					// Verify organization_id is set from provider default
					resource.TestCheckResourceAttr("posthog_role.test", "organization_id", orgID),
					resource.TestCheckResourceAttrSet("posthog_role.test", "id"),
				),
			},
		},
	})
}

// TestRole_Basic tests creating a role with only the required fields.
func TestRole_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccRolePreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_role.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_role.test", "organization_id", orgID),
					resource.TestCheckResourceAttrSet("posthog_role.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_role.test", "created_at"),
				),
			},
		},
	})
}

// TestRole_Update tests updating a role's name.
func TestRole_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccRolePreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_role.test", "name", rName),
				),
			},
			{
				Config: testAccRoleBasic(orgID, rName+"-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_role.test", "name", rName+"-updated"),
				),
			},
		},
	})
}

// TestRole_Import tests importing an existing role.
func TestRole_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccRolePreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoleBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_role.test", "name", rName),
				),
			},
			{
				ResourceName:      "posthog_role.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccRoleImportStateIdFunc("posthog_role.test"),
			},
		},
	})
}

func testAccRoleImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["organization_id"]
		roleID := rs.Primary.Attributes["id"]

		return fmt.Sprintf("%s/%s", orgID, roleID), nil
	}
}

func testAccRoleBasic(orgID, name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_role" "test" {
  organization_id = %q
  name            = %q
}
`, orgID, name)
}

func testAccRoleWithProviderOrganizationID(orgID, name string) string {
	return fmt.Sprintf(`
provider "posthog" {
  organization_id = %q
}

resource "posthog_role" "test" {
  name = %q
}
`, orgID, name)
}
