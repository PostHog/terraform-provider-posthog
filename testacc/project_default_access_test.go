package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccProjectDefaultAccessPreCheck(t *testing.T) {
	t.Helper()
	testAccPreCheck(t)
}

func TestProjectDefaultAccess_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectDefaultAccessPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectDefaultAccessBasic("member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_project_default_access.test", "id"),
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "member"),
				),
			},
		},
	})
}

func TestProjectDefaultAccess_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectDefaultAccessPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectDefaultAccessBasic("member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "member"),
				),
			},
			{
				Config: testAccProjectDefaultAccessBasic("admin"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "admin"),
				),
			},
			{
				Config: testAccProjectDefaultAccessBasic("none"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "none"),
				),
			},
		},
	})
}

func TestProjectDefaultAccess_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	projectID := getProjectID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectDefaultAccessPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectDefaultAccessBasic("member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_project_default_access.test", "id"),
				),
			},
			{
				ResourceName:            "posthog_project_default_access.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           projectID,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func testAccProjectDefaultAccessBasic(accessLevel string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_project_default_access" "test" {
  access_level = %q
}
`, accessLevel)
}
