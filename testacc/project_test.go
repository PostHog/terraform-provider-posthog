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
	testAccBasePreCheck(t)
	skipIfNoOrganizationID(t)
}

func getOrganizationID() string {
	return os.Getenv("POSTHOG_ORGANIZATION_ID")
}

// TestProject_ProviderLevelOrganizationID tests that a project inherits organization_id from the provider.
func TestProject_ProviderLevelOrganizationID(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectWithProviderOrganizationID(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
					// Verify organization_id is set from provider default
					resource.TestCheckResourceAttr("posthog_project.test", "organization_id", orgID),
					resource.TestCheckResourceAttrSet("posthog_project.test", "id"),
				),
			},
		},
	})
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

// TestProject_Import tests importing an existing project by organization_id/project_id.
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

		return fmt.Sprintf("%s/%s", orgID, projectID), nil
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

// TestProject_UseProjectIDForResources tests the workflow where a user creates a project
// and then uses the project's ID for other resources. This validates that project_id
// is optional at the provider level and can be specified per-resource.
func TestProject_UseProjectIDForResources(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	flagKey := acctest.RandomWithPrefix("tf-acc-flag")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectWithFeatureFlag(orgID, rName, flagKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify project was created
					resource.TestCheckResourceAttr("posthog_project.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_project.test", "id"),
					// Verify feature flag was created in the new project
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", flagKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "id"),
					// Verify feature flag's project_id matches the project's id
					resource.TestCheckResourceAttrPair(
						"posthog_feature_flag.test", "project_id",
						"posthog_project.test", "id",
					),
				),
			},
		},
	})
}

func testAccProjectWithFeatureFlag(orgID, projectName, flagKey string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_project" "test" {
  organization_id = %q
  name            = %q
}

resource "posthog_feature_flag" "test" {
  project_id = posthog_project.test.id
  key        = %q
  name       = "Test flag in new project"
  active     = true
}
`, orgID, projectName, flagKey)
}

func testAccProjectWithProviderOrganizationID(orgID, name string) string {
	return fmt.Sprintf(`
provider "posthog" {
  organization_id = %q
}

resource "posthog_project" "test" {
  name = %q
}
`, orgID, name)
}
