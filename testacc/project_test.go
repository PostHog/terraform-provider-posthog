package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccProjectPreCheck(t *testing.T) {
	t.Helper()
	testAccPreCheck(t)
	if os.Getenv("POSTHOG_ORGANIZATION_ID") == "" {
		t.Fatal("POSTHOG_ORGANIZATION_ID must be set for project acceptance tests")
	}
}

func getOrganizationID() string {
	return os.Getenv("POSTHOG_ORGANIZATION_ID")
}

// TestProject_Basic tests creating a project with only the required fields.
func TestProject_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_project.test", "organization_id", orgID),
					resource.TestCheckResourceAttrSet("posthog_project.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_project.test", "api_token"),
				),
			},
		},
	})
}

// TestProject_Update tests updating the project name.
func TestProject_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccProjectBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
				),
			},
			// Update name
			{
				Config: testAccProjectBasic(orgID, rName+"-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName+"-updated"),
				),
			},
		},
	})
}

// TestProject_Import tests importing an existing project by organization_id:project_id.
func TestProject_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccProjectBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
				),
			},
			// Import
			{
				ResourceName:            "posthog_project.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccProjectImportStateIdFunc("posthog_project.test"),
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// TestProject_Recreate tests deleting and recreating a project with the same name.
func TestProject_Recreate(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create first project
			{
				Config: testAccProjectBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
				),
			},
			// Remove resource from config (destroys it)
			{
				Config: `provider "posthog" {}`,
			},
			// Recreate
			{
				Config: testAccProjectBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
				),
			},
		},
	})
}

func testAccProjectImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["organization_id"]
		projectID := rs.Primary.Attributes["id"]

		return fmt.Sprintf("%s:%s", orgID, projectID), nil
	}
}

// TestProject_WithTimezone tests creating a project with an explicit timezone.
func TestProject_WithTimezone(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectWithTimezone(orgID, rName, "Europe/London"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_project.test", "timezone", "Europe/London"),
				),
			},
		},
	})
}

func testAccProjectBasic(orgID, name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_project" "test" {
  organization_id = %q
  name            = %q
}
`, orgID, name)
}

func testAccProjectWithTimezone(orgID, name, timezone string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_project" "test" {
  organization_id = %q
  name            = %q
  timezone        = %q
}
`, orgID, name, timezone)
}
