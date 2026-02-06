package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// TestRBACWorkflow_FullIntegration tests the complete RBAC workflow:
// 1. Create a project
// 2. Create roles
// 3. Set restrictive project default access
// 4. Grant roles access to the project
// 5. Configure resource-level access controls
// Heavily inspired by examples/rbac/main.tf
func TestRBACWorkflow_FullIntegration(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-rbac-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			skipIfNoOrganizationID(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRBACWorkflowConfig(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify project created
					resource.TestCheckResourceAttrSet("posthog_project.main", "id"),
					resource.TestCheckResourceAttr("posthog_project.main", "name", rName+"-project"),

					// Verify roles created
					resource.TestCheckResourceAttrSet("posthog_role.engineering", "id"),
					resource.TestCheckResourceAttrSet("posthog_role.product_managers", "id"),
					resource.TestCheckResourceAttrSet("posthog_role.support", "id"),
					resource.TestCheckResourceAttr("posthog_role.engineering", "name", rName+"-engineering"),
					resource.TestCheckResourceAttr("posthog_role.product_managers", "name", rName+"-product"),
					resource.TestCheckResourceAttr("posthog_role.support", "name", rName+"-support"),

					// Verify project default access is restrictive
					resource.TestCheckResourceAttr("posthog_project_default_access.restrictive", "access_level", "none"),

					// Verify project member access
					resource.TestCheckResourceAttr("posthog_project_member.engineering", "access_level", "admin"),
					resource.TestCheckResourceAttr("posthog_project_member.product_managers", "access_level", "member"),
					resource.TestCheckResourceAttr("posthog_project_member.support", "access_level", "member"),

					// Verify resource-level access controls - feature flags
					resource.TestCheckResourceAttr("posthog_access_control.feature_flags_engineering", "access_level", "editor"),
					resource.TestCheckResourceAttr("posthog_access_control.feature_flags_product", "access_level", "viewer"),
					resource.TestCheckResourceAttr("posthog_access_control.feature_flags_support", "access_level", "none"),

					// Verify resource-level access controls - dashboards
					resource.TestCheckResourceAttr("posthog_access_control.dashboards_engineering", "access_level", "viewer"),
					resource.TestCheckResourceAttr("posthog_access_control.dashboards_support", "access_level", "viewer"),

					// Verify resource-level access controls - experiments
					resource.TestCheckResourceAttr("posthog_access_control.experiments_product", "access_level", "editor"),
					resource.TestCheckResourceAttr("posthog_access_control.experiments_engineering", "access_level", "viewer"),

					// Verify resource-level access controls - session recordings
					resource.TestCheckResourceAttr("posthog_access_control.recordings_support", "access_level", "none"),
					resource.TestCheckResourceAttr("posthog_access_control.recordings_engineering", "access_level", "none"),
				),
			},
			// Test updates work across the workflow
			{
				Config: testAccRBACWorkflowConfigUpdated(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Project default changed to member
					resource.TestCheckResourceAttr("posthog_project_default_access.restrictive", "access_level", "member"),
					// Engineering demoted to member
					resource.TestCheckResourceAttr("posthog_project_member.engineering", "access_level", "member"),
					// Product managers promoted to admin
					resource.TestCheckResourceAttr("posthog_project_member.product_managers", "access_level", "admin"),
					// Feature flag access updated
					resource.TestCheckResourceAttr("posthog_access_control.feature_flags_engineering", "access_level", "viewer"),
					resource.TestCheckResourceAttr("posthog_access_control.feature_flags_product", "access_level", "editor"),
				),
			},
		},
	})
}

// TestProjectDefaultAccess_SingletonBehavior verifies the singleton behaves correctly:
// - Only one can exist per project
// - Updates are in-place (not replace)
func TestProjectDefaultAccess_SingletonBehavior(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-singleton-test")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			skipIfNoOrganizationID(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with "member" access
			{
				Config: testAccProjectDefaultAccessSingleton(orgID, rName, "member"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "member"),
				),
			},
			// Update to "admin" - verify it's an in-place update, not replace
			{
				Config: testAccProjectDefaultAccessSingleton(orgID, rName, "admin"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "admin"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("posthog_project_default_access.test", plancheck.ResourceActionUpdate),
					},
				},
			},
			// Update to "none"
			{
				Config: testAccProjectDefaultAccessSingleton(orgID, rName, "none"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "none"),
				),
			},
		},
	})
}

// TestProjectDefaultAccess_DestroyResetsToNone verifies that destroying the
// project_default_access resource resets the access level to "none" (most restrictive).
// This is the expected singleton behavior - the default always exists, we just manage its value.
func TestProjectDefaultAccess_DestroyResetsToNone(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-destroy-test")
	orgID := getOrganizationID()

	var projectID string

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			skipIfNoOrganizationID(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			return testAccCheckProjectDefaultAccessResetsToNone(projectID)
		},
		Steps: []resource.TestStep{
			// Create with "admin" access
			{
				Config: testAccProjectDefaultAccessForDestroy(orgID, rName, "admin"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_default_access.test", "access_level", "admin"),
					// Capture the project ID for CheckDestroy
					resource.TestCheckResourceAttrWith("posthog_project.test", "id", func(value string) error {
						projectID = value
						return nil
					}),
				),
			},
			// After destroy, CheckDestroy will verify it reset to "none"
		},
	})
}

func testAccRBACWorkflowConfig(orgID, rName string) string {
	return fmt.Sprintf(`
provider "posthog" {}

# Step 1: Create the project
resource "posthog_project" "main" {
  organization_id = %[1]q
  name            = "%[2]s-project"
}

# Step 2: Create roles
resource "posthog_role" "engineering" {
  organization_id = %[1]q
  name            = "%[2]s-engineering"
}

resource "posthog_role" "product_managers" {
  organization_id = %[1]q
  name            = "%[2]s-product"
}

resource "posthog_role" "support" {
  organization_id = %[1]q
  name            = "%[2]s-support"
}

# Step 3: Set restrictive project default
resource "posthog_project_default_access" "restrictive" {
  project_id   = posthog_project.main.id
  access_level = "none"
}

# Step 4: Grant roles project access
resource "posthog_project_member" "engineering" {
  project_id   = posthog_project.main.id
  role         = posthog_role.engineering.id
  access_level = "admin"
}

resource "posthog_project_member" "product_managers" {
  project_id   = posthog_project.main.id
  role         = posthog_role.product_managers.id
  access_level = "member"
}

resource "posthog_project_member" "support" {
  project_id   = posthog_project.main.id
  role         = posthog_role.support.id
  access_level = "member"
}

# Step 5: Configure resource-level access

# Feature flags
resource "posthog_access_control" "feature_flags_engineering" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.engineering.id
  access_level = "editor"
}

resource "posthog_access_control" "feature_flags_product" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.product_managers.id
  access_level = "viewer"
}

resource "posthog_access_control" "feature_flags_support" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.support.id
  access_level = "none"
}

# Dashboards
resource "posthog_access_control" "dashboards_engineering" {
  project_id   = posthog_project.main.id
  resource     = "dashboard"
  role         = posthog_role.engineering.id
  access_level = "viewer"
}

resource "posthog_access_control" "dashboards_support" {
  project_id   = posthog_project.main.id
  resource     = "dashboard"
  role         = posthog_role.support.id
  access_level = "viewer"
}

# Experiments
resource "posthog_access_control" "experiments_product" {
  project_id   = posthog_project.main.id
  resource     = "experiment"
  role         = posthog_role.product_managers.id
  access_level = "editor"
}

resource "posthog_access_control" "experiments_engineering" {
  project_id   = posthog_project.main.id
  resource     = "experiment"
  role         = posthog_role.engineering.id
  access_level = "viewer"
}

# Session recordings
resource "posthog_access_control" "recordings_support" {
  project_id   = posthog_project.main.id
  resource     = "session_recording"
  role         = posthog_role.support.id
  access_level = "none"
}

resource "posthog_access_control" "recordings_engineering" {
  project_id   = posthog_project.main.id
  resource     = "session_recording"
  role         = posthog_role.engineering.id
  access_level = "none"
}
`, orgID, rName)
}

func testAccRBACWorkflowConfigUpdated(orgID, rName string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_project" "main" {
  organization_id = %[1]q
  name            = "%[2]s-project"
}

resource "posthog_role" "engineering" {
  organization_id = %[1]q
  name            = "%[2]s-engineering"
}

resource "posthog_role" "product_managers" {
  organization_id = %[1]q
  name            = "%[2]s-product"
}

resource "posthog_role" "support" {
  organization_id = %[1]q
  name            = "%[2]s-support"
}

# Updated: default access changed to "member"
resource "posthog_project_default_access" "restrictive" {
  project_id   = posthog_project.main.id
  access_level = "member"
}

# Updated: engineering demoted to member
resource "posthog_project_member" "engineering" {
  project_id   = posthog_project.main.id
  role         = posthog_role.engineering.id
  access_level = "member"
}

# Updated: product managers promoted to admin
resource "posthog_project_member" "product_managers" {
  project_id   = posthog_project.main.id
  role         = posthog_role.product_managers.id
  access_level = "admin"
}

resource "posthog_project_member" "support" {
  project_id   = posthog_project.main.id
  role         = posthog_role.support.id
  access_level = "member"
}

# Updated: engineering can only view feature flags now
resource "posthog_access_control" "feature_flags_engineering" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.engineering.id
  access_level = "viewer"
}

# Updated: product managers can edit feature flags now
resource "posthog_access_control" "feature_flags_product" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.product_managers.id
  access_level = "editor"
}

resource "posthog_access_control" "feature_flags_support" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.support.id
  access_level = "none"
}

resource "posthog_access_control" "dashboards_engineering" {
  project_id   = posthog_project.main.id
  resource     = "dashboard"
  role         = posthog_role.engineering.id
  access_level = "viewer"
}

resource "posthog_access_control" "dashboards_support" {
  project_id   = posthog_project.main.id
  resource     = "dashboard"
  role         = posthog_role.support.id
  access_level = "viewer"
}

resource "posthog_access_control" "experiments_product" {
  project_id   = posthog_project.main.id
  resource     = "experiment"
  role         = posthog_role.product_managers.id
  access_level = "editor"
}

resource "posthog_access_control" "experiments_engineering" {
  project_id   = posthog_project.main.id
  resource     = "experiment"
  role         = posthog_role.engineering.id
  access_level = "viewer"
}

resource "posthog_access_control" "recordings_support" {
  project_id   = posthog_project.main.id
  resource     = "session_recording"
  role         = posthog_role.support.id
  access_level = "none"
}

resource "posthog_access_control" "recordings_engineering" {
  project_id   = posthog_project.main.id
  resource     = "session_recording"
  role         = posthog_role.engineering.id
  access_level = "none"
}
`, orgID, rName)
}

func testAccProjectDefaultAccessSingleton(orgID, rName, accessLevel string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_project" "test" {
  organization_id = %[1]q
  name            = "%[2]s-project"
}

resource "posthog_project_default_access" "test" {
  project_id   = posthog_project.test.id
  access_level = %[3]q
}
`, orgID, rName, accessLevel)
}

func testAccProjectDefaultAccessForDestroy(orgID, rName, accessLevel string) string {
	return testAccProjectDefaultAccessSingleton(orgID, rName, accessLevel)
}

// testAccCheckProjectDefaultAccessResetsToNone verifies that after destroying
// the project_default_access resource, the API returns "none" as the default access level.
func testAccCheckProjectDefaultAccessResetsToNone(projectID string) error {
	if projectID == "" {
		return fmt.Errorf("project id is not set")
	}

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)

	// Get project access controls
	result, _, err := client.GetProjectAccessControls(context.Background(), projectID)
	if err != nil {
		return fmt.Errorf("error whilst retrieving access controls: %w", err)
	}

	// Find the default access level (where role and organization_member are both nil)
	for _, ac := range result.AccessControls {
		if ac.Role == nil && ac.OrganizationMember == nil {
			if ac.AccessLevel == nil || *ac.AccessLevel != "none" {
				return fmt.Errorf("expected default access level to be reset to 'none' after destroy, got '%s'", *ac.AccessLevel)
			}
			// Found and verified - success
			return nil
		}
	}

	return fmt.Errorf("no default found")
}
