package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// TestFeatureFlag_Basic tests creating a feature flag with only the required field (key).
func TestFeatureFlag_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagBasic(rKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "id"),
				),
			},
		},
	})
}

// TestFeatureFlag_AllFields tests creating a feature flag with all optional fields.
func TestFeatureFlag_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagAllFields(rKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "name", "Test Feature Flag"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "true"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "id"),
				),
			},
		},
	})
}

// TestFeatureFlag_RolloutPercentage tests using the rollout_percentage convenience field.
func TestFeatureFlag_RolloutPercentage(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagRolloutPercentage(rKey, 50),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "rollout_percentage", "50"),
				),
			},
			// Update rollout percentage
			{
				Config: testAccFeatureFlagRolloutPercentage(rKey, 100),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "rollout_percentage", "100"),
				),
			},
			// Set to 0%
			{
				Config: testAccFeatureFlagRolloutPercentage(rKey, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "rollout_percentage", "0"),
				),
			},
		},
	})
}

// TestFeatureFlag_Filters tests using raw filters JSON.
func TestFeatureFlag_Filters(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagSimpleFilters(rKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "filters"),
				),
			},
		},
	})
}

// TestFeatureFlag_FiltersWithRollout tests filters JSON with embedded rollout_percentage.
func TestFeatureFlag_FiltersWithRollout(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagFiltersWithRollout(rKey, 75),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "filters"),
					testCheckFiltersRolloutPercentage("posthog_feature_flag.test", 0, 75),
				),
			},
		},
	})
}

// testCheckFiltersRolloutPercentage verifies the rollout_percentage in a filters JSON attribute.
func testCheckFiltersRolloutPercentage(resourceName string, groupIndex int, expected float64) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(resourceName, "filters", func(value string) error {
		var filters struct {
			Groups []struct {
				RolloutPercentage *float64 `json:"rollout_percentage"`
			} `json:"groups"`
		}
		if err := json.Unmarshal([]byte(value), &filters); err != nil {
			return fmt.Errorf("failed to parse filters JSON: %w", err)
		}
		if groupIndex >= len(filters.Groups) {
			return fmt.Errorf("group index %d out of range (have %d groups)", groupIndex, len(filters.Groups))
		}
		if filters.Groups[groupIndex].RolloutPercentage == nil {
			return fmt.Errorf("rollout_percentage is nil for group %d", groupIndex)
		}
		if *filters.Groups[groupIndex].RolloutPercentage != expected {
			return fmt.Errorf("expected rollout_percentage %v, got %v", expected, *filters.Groups[groupIndex].RolloutPercentage)
		}
		return nil
	})
}

// TestFeatureFlag_Update tests updating each field individually.
func TestFeatureFlag_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccFeatureFlagWithName(rKey, "Initial Name", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "name", "Initial Name"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "true"),
				),
			},
			// Update name
			{
				Config: testAccFeatureFlagWithName(rKey, "Updated Name", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "name", "Updated Name"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "true"),
				),
			},
			// Update active
			{
				Config: testAccFeatureFlagWithName(rKey, "Updated Name", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "name", "Updated Name"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "false"),
				),
			},
		},
	})
}

// TestFeatureFlag_ToggleActive tests toggling the active state on/off.
func TestFeatureFlag_ToggleActive(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Start active
			{
				Config: testAccFeatureFlagActive(rKey, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "true"),
				),
			},
			// Deactivate
			{
				Config: testAccFeatureFlagActive(rKey, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "false"),
				),
			},
			// Reactivate
			{
				Config: testAccFeatureFlagActive(rKey, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "true"),
				),
			},
		},
	})
}

// TestFeatureFlag_Tags tests creating, updating, and removing tags.
func TestFeatureFlag_Tags(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with tags
			{
				Config: testAccFeatureFlagWithTags(rKey, []string{"tag1", "tag2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "tags.#", "2"),
				),
			},
			// Add more tags
			{
				Config: testAccFeatureFlagWithTags(rKey, []string{"tag1", "tag2", "tag3"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "tags.#", "3"),
				),
			},
			// Remove tags
			{
				Config: testAccFeatureFlagWithTags(rKey, []string{"tag1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "tags.#", "1"),
				),
			},
		},
	})
}

// TestFeatureFlag_ComplexFilters tests complex filters with properties and multiple groups.
func TestFeatureFlag_ComplexFilters(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagComplexFilters(rKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "filters"),
				),
			},
		},
	})
}

// TestFeatureFlag_Import tests importing an existing feature flag by ID.
func TestFeatureFlag_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccFeatureFlagWithName(rKey, "Import Test", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
				),
			},
			// Import
			{
				ResourceName:            "posthog_feature_flag.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"},
			},
		},
	})
}

// TestFeatureFlag_ImportWithProjectID tests importing using the project_id/resource_id format.
// This format allows importing resources without having project_id set at the provider level.
func TestFeatureFlag_ImportWithProjectID(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccFeatureFlagWithName(rKey, "Import With ProjectID Test", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
				),
			},
			// Import using project_id/resource_id format
			{
				ResourceName:            "posthog_feature_flag.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["posthog_feature_flag.test"]
					if !ok {
						return "", fmt.Errorf("resource not found: posthog_feature_flag.test")
					}
					// Use project_id/resource_id format
					return fmt.Sprintf("%s/%s", projectID, rs.Primary.ID), nil
				},
			},
			// Verify project_id is set in state after import
			{
				Config: testAccFeatureFlagWithName(rKey, "Import With ProjectID Test", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "project_id", projectID),
				),
			},
		},
	})
}

// TestFeatureFlag_MultipleGroups tests feature flags with multiple release groups.
func TestFeatureFlag_MultipleGroups(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagMultipleGroups(rKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "filters"),
				),
			},
		},
	})
}

func testAccFeatureFlagBasic(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key = %q
}
`, key)
}

func testAccFeatureFlagAllFields(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Test Feature Flag"
  active = true
  tags   = ["managed-by:terraform", "env:test"]
}
`, key)
}

func testAccFeatureFlagRolloutPercentage(key string, percentage int) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                = %q
  name               = "Rollout Test"
  active             = true
  rollout_percentage = %d
}
`, key, percentage)
}

func testAccFeatureFlagSimpleFilters(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Simple Filters"
  active = true

  filters = jsonencode({
    groups = [{
      properties         = []
      rollout_percentage = 100
    }]
  })
}
`, key)
}

func testAccFeatureFlagFiltersWithRollout(key string, percentage int) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Filters with rollout"
  active = true

  filters = jsonencode({
    groups = [{
      properties         = []
      rollout_percentage = %d
    }]
  })
}
`, key, percentage)
}

func testAccFeatureFlagWithName(key, name string, active bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = %q
  active = %t
}
`, key, name, active)
}

func testAccFeatureFlagActive(key string, active bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Active Toggle Test"
  active = %t
}
`, key, active)
}

func testAccFeatureFlagWithTags(key string, tags []string) string {
	tagsStr := ""
	for i, tag := range tags {
		if i > 0 {
			tagsStr += ", "
		}
		tagsStr += fmt.Sprintf("%q", tag)
	}

	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Tags Test"
  active = true
  tags   = [%s]
}
`, key, tagsStr)
}

func testAccFeatureFlagComplexFilters(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Complex Filters"
  active = true

  filters = jsonencode({
    groups = [{
      properties = [
        {
          key      = "email"
          type     = "person"
          value    = ["test@example.com", "admin@example.com"]
          operator = "exact"
        },
        {
          key      = "$browser"
          type     = "person"
          value    = ["Chrome", "Firefox"]
          operator = "exact"
        }
      ]
      rollout_percentage = 100
    }]
  })
}
`, key)
}

func testAccFeatureFlagMultipleGroups(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Multiple Groups"
  active = true

  filters = jsonencode({
    groups = [
      {
        properties = [
          {
            key      = "email"
            type     = "person"
            value    = ["admin@example.com"]
            operator = "exact"
          }
        ]
        rollout_percentage = 100
      },
      {
        properties         = []
        rollout_percentage = 10
      }
    ]
  })
}
`, key)
}

// TestFeatureFlag_ExternalDeletion tests that Terraform detects when a feature flag
// is deleted externally (e.g., via the PostHog UI) and properly removes it from state.
// This verifies drift detection for soft-deleted feature flags.
func TestFeatureFlag_ExternalDeletion(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	host := os.Getenv("POSTHOG_HOST")
	apiKey := os.Getenv("POSTHOG_API_KEY")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")
	client := httpclient.NewDefaultClient(host, apiKey, "acceptance-test")

	var flagID string
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create the feature flag
			{
				Config: testAccFeatureFlagWithName(rKey, "External Deletion Test", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "id"),
					// Capture the ID for external deletion
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["posthog_feature_flag.test"]
						if !ok {
							return fmt.Errorf("resource not found: posthog_feature_flag.test")
						}
						flagID = rs.Primary.ID
						return nil
					},
				),
			},
			// Step 2: Delete the flag externally via the API (soft delete)
			// Then verify Terraform detects the deletion and plans to recreate
			{
				PreConfig: func() {
					// Delete the feature flag externally using the provider's client
					_, err := (&client).DeleteFeatureFlag(context.Background(), projectID, flagID)
					if err != nil {
						t.Fatalf("Failed to delete feature flag externally: %v", err)
					}
				},
				Config: testAccFeatureFlagWithName(rKey, "External Deletion Test", true),
				// ExpectNonEmptyPlan: true means Terraform detected drift and will recreate
				ExpectNonEmptyPlan: true,
				// The plan should show the resource will be created (since it was deleted externally)
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
				),
			},
		},
	})
}
