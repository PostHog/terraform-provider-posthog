package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccOrganizationMemberPreCheck(t *testing.T) {
	t.Helper()
	testAccBasePreCheck(t)
	skipIfNoOrganizationID(t)
	skipIfNoTestUserEmail(t)
}

// TestOrganizationMember_Basic tests managing an organization member's level.
func TestOrganizationMember_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccOrganizationMemberPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationMemberBasic(orgID, userEmail, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_organization_member.test", "organization_id", orgID),
					resource.TestCheckResourceAttr("posthog_organization_member.test", "level", "member"),
					resource.TestCheckResourceAttrSet("posthog_organization_member.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_organization_member.test", "user_uuid"),
					resource.TestCheckResourceAttrSet("posthog_organization_member.test", "email"),
				),
			},
		},
	})
}

// TestOrganizationMember_UpdateLevel tests updating a member's level.
func TestOrganizationMember_UpdateLevel(t *testing.T) {
	skipIfNotAcceptance(t)

	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccOrganizationMemberPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationMemberBasic(orgID, userEmail, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_organization_member.test", "level", "member"),
				),
			},
			{
				Config: testAccOrganizationMemberBasic(orgID, userEmail, "admin"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_organization_member.test", "level", "admin"),
				),
			},
			// Revert back to member
			{
				Config: testAccOrganizationMemberBasic(orgID, userEmail, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_organization_member.test", "level", "member"),
				),
			},
		},
	})
}

// TestOrganizationMember_Import tests importing an existing organization member.
func TestOrganizationMember_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccOrganizationMemberPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationMemberBasic(orgID, userEmail, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_organization_member.test", "id"),
				),
			},
			{
				ResourceName:      "posthog_organization_member.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccOrganizationMemberImportStateIdFunc("posthog_organization_member.test"),
			},
		},
	})
}

func testAccOrganizationMemberImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["organization_id"]
		userUUID := rs.Primary.Attributes["user_uuid"]

		return fmt.Sprintf("%s/%s", orgID, userUUID), nil
	}
}

func testAccOrganizationMemberBasic(orgID, userEmail, level string) string {
	return fmt.Sprintf(`
provider "posthog" {}

data "posthog_user" "test" {
  organization_id = %q
  email           = %q
}

resource "posthog_organization_member" "test" {
  organization_id = %q
  user_uuid       = data.posthog_user.test.uuid
  level           = %q
}
`, orgID, userEmail, orgID, level)
}
