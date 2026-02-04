package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccProjectMemberPreCheck(t *testing.T) {
	t.Helper()
	testAccPreCheck(t)
	skipIfNoOrganizationID(t)
}

func TestProjectMember_BasicWithRole(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectMemberPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectMemberWithRole(orgID, rName, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_project_member.test", "id"),
					resource.TestCheckResourceAttr("posthog_project_member.test", "access_level", "member"),
					resource.TestCheckResourceAttrSet("posthog_project_member.test", "role"),
				),
			},
		},
	})
}

func TestProjectMember_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectMemberPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectMemberWithRole(orgID, rName, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_member.test", "access_level", "member"),
				),
			},
			{
				Config: testAccProjectMemberWithRole(orgID, rName, "admin"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_member.test", "access_level", "admin"),
				),
			},
		},
	})
}

func TestProjectMember_WithOrganizationMember(t *testing.T) {
	skipIfNotAcceptance(t)

	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectMemberPreCheck(t); skipIfNoTestUserEmail(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectMemberWithOrgMember(orgID, userEmail),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_project_member.test", "id"),
					resource.TestCheckResourceAttr("posthog_project_member.test", "access_level", "member"),
					resource.TestCheckResourceAttrSet("posthog_project_member.test", "organization_member"),
				),
			},
		},
	})
}

func TestProjectMember_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectMemberPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectMemberWithRole(orgID, rName, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_project_member.test", "id"),
				),
			},
			{
				ResourceName:            "posthog_project_member.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccProjectMemberImportStateIdFunc("posthog_project_member.test"),
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func testAccProjectMemberImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		projectID := rs.Primary.Attributes["project_id"]
		roleID := rs.Primary.Attributes["role"]

		return fmt.Sprintf("%s/role/%s", projectID, roleID), nil
	}
}

func testAccProjectMemberWithRole(orgID, roleName, accessLevel string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_role" "test" {
  organization_id = %q
  name            = %q
}

resource "posthog_project_member" "test" {
  role         = posthog_role.test.id
  access_level = %q
}
`, orgID, roleName, accessLevel)
}

func testAccProjectMemberWithOrgMember(orgID, userEmail string) string {
	return fmt.Sprintf(`
provider "posthog" {}

data "posthog_user" "test" {
  organization_id = %q
  email           = %q
}

resource "posthog_organization_member" "test" {
  organization_id   = %q
  user_uuid         = data.posthog_user.test.uuid
  retain_on_destroy = true
}

resource "posthog_project_member" "test" {
  organization_member = posthog_organization_member.test.id
  access_level        = "member"
}
`, orgID, userEmail, orgID)
}
