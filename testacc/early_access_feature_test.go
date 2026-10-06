package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

const earlyAccessFeatureAddress = "posthog_early_access_feature.test"

func testAccEarlyAccessFeatureClient() httpclient.PosthogClient {
	return httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")
}

// testAccCheckEarlyAccessFeatureDestroy also soft-deletes the flag each feature was backed by:
// PostHog keeps it on destroy, and a flag it created itself would otherwise pile up in the
// test project. A flag Terraform already deleted is skipped.
func testAccCheckEarlyAccessFeatureDestroy(s *terraform.State) error {
	return testAccCheckEarlyAccessFeatureDestroyWithFlagRequirement(s, false)
}

func testAccCheckAutoCreatedFeatureFlagRetained(s *terraform.State) error {
	return testAccCheckEarlyAccessFeatureDestroyWithFlagRequirement(s, true)
}

func testAccCheckEarlyAccessFeatureDestroyWithFlagRequirement(s *terraform.State, requireFlagRetained bool) error {
	client := testAccEarlyAccessFeatureClient()
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_early_access_feature" {
			continue
		}

		_, status, err := client.GetEarlyAccessFeature(context.Background(), projectID, rs.Primary.ID)
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			if err != nil {
				return fmt.Errorf("unexpected error checking early access feature %s: %w", rs.Primary.ID, err)
			}
			return fmt.Errorf("early access feature %s still exists", rs.Primary.ID)
		}

		flagID := rs.Primary.Attributes["feature_flag_id"]
		if flagID == "" {
			continue
		}
		flag, status, err := client.GetFeatureFlag(context.Background(), projectID, flagID)
		if status == httpclient.HTTPStatusCode(http.StatusNotFound) || (err == nil && flag.Deleted != nil && *flag.Deleted) {
			if requireFlagRetained {
				return fmt.Errorf("auto-created feature flag %s was deleted with early access feature %s", flagID, rs.Primary.ID)
			}
			continue
		}
		if err != nil {
			return fmt.Errorf("reading feature flag %s for cleanup: %w", flagID, err)
		}
		if _, err := client.DeleteFeatureFlag(context.Background(), projectID, flagID); err != nil {
			return fmt.Errorf("cleaning up feature flag %s: %w", flagID, err)
		}
	}

	return nil
}

// testAccCheckFlagFeatureEnrollment reads the linked flag straight from the API.
func testAccCheckFlagFeatureEnrollment(want bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[earlyAccessFeatureAddress]
		if !ok {
			return fmt.Errorf("%s not found in state", earlyAccessFeatureAddress)
		}
		flagID := rs.Primary.Attributes["feature_flag_id"]

		client := testAccEarlyAccessFeatureClient()
		flag, _, err := client.GetFeatureFlag(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), flagID)
		if err != nil {
			return fmt.Errorf("reading feature flag %s: %w", flagID, err)
		}
		got, _ := flag.Filters["feature_enrollment"].(bool)
		if got != want {
			return fmt.Errorf("feature flag %s: feature_enrollment = %v, want %v (filters: %v)", flagID, flag.Filters["feature_enrollment"], want, flag.Filters)
		}
		return nil
	}
}

func testAccEarlyAccessFeatureLinked(key, flagName, stage, description string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                = %q
  name               = %q
  rollout_percentage = 0
}

resource "posthog_early_access_feature" "test" {
  name              = %q
  stage             = %q
  description       = %q
  documentation_url = "https://example.com/docs"
  payload           = jsonencode({ theme = "dark" })
  feature_flag_id   = posthog_feature_flag.test.id
}
`, key, flagName, key, stage, description)
}

func testAccEarlyAccessFeatureLinkedToSecondFlag(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                = %q
  name               = "flag renamed"
  rollout_percentage = 0
}

resource "posthog_feature_flag" "second" {
  key                = %q
  name               = "second flag"
  rollout_percentage = 0
}

resource "posthog_early_access_feature" "test" {
  name              = %q
  stage             = "archived"
  description       = "second"
  documentation_url = "https://example.com/docs"
  payload           = jsonencode({ theme = "dark" })
  feature_flag_id   = posthog_feature_flag.second.id
}
`, key, key+"-second", key)
}

func TestEarlyAccessFeature_LinkedFlagLifecycle(t *testing.T) {
	skipIfNotAcceptance(t)

	key := strings.ReplaceAll(randomTestPrefix(), "_", "-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEarlyAccessFeatureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEarlyAccessFeatureLinked(key, "flag", "draft", "first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(earlyAccessFeatureAddress, "id"),
					resource.TestCheckResourceAttrSet(earlyAccessFeatureAddress, "created_at"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "name", key),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "stage", "draft"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "description", "first"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "documentation_url", "https://example.com/docs"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "payload", `{"theme":"dark"}`),
					resource.TestCheckResourceAttrPair(earlyAccessFeatureAddress, "feature_flag_id", "posthog_feature_flag.test", "id"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "feature_flag_key", key),
					testAccCheckFlagFeatureEnrollment(false),
				),
			},
			{
				// An active stage stamps feature_enrollment onto the flag. The framework's
				// post-apply plan must still be empty, so the Terraform-managed flag must not
				// show it as drift.
				Config: testAccEarlyAccessFeatureLinked(key, "flag", "beta", "second"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(earlyAccessFeatureAddress, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "stage", "beta"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "description", "second"),
					testAccCheckFlagFeatureEnrollment(true),
				),
			},
			{
				// Updating the flag sends filters without the enrollment marker; the apply must
				// succeed and PostHog must keep the marker.
				Config: testAccEarlyAccessFeatureLinked(key, "flag renamed", "beta", "second"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("posthog_feature_flag.test", plancheck.ResourceActionUpdate),
						plancheck.ExpectResourceAction(earlyAccessFeatureAddress, plancheck.ResourceActionNoop),
					},
				},
				Check: testAccCheckFlagFeatureEnrollment(true),
			},
			{
				// GA keeps opt-in gating unless the separate PostHog rollout-to-all action
				// is requested. This resource manages the stage, not that one-time action.
				Config: testAccEarlyAccessFeatureLinked(key, "flag renamed", "general-availability", "second"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(earlyAccessFeatureAddress, plancheck.ResourceActionUpdate),
						plancheck.ExpectResourceAction("posthog_feature_flag.test", plancheck.ResourceActionNoop),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "stage", "general-availability"),
					testAccCheckFlagFeatureEnrollment(true),
				),
			},
			{
				Config: testAccEarlyAccessFeatureLinked(key, "flag renamed", "archived", "second"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "stage", "archived"),
					testAccCheckFlagFeatureEnrollment(false),
				),
			},
			{
				Config: testAccEarlyAccessFeatureLinkedToSecondFlag(key),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(earlyAccessFeatureAddress, plancheck.ResourceActionReplace),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(earlyAccessFeatureAddress, "feature_flag_id", "posthog_feature_flag.second", "id"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "feature_flag_key", key+"-second"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", key),
					testAccCheckFlagFeatureEnrollment(false),
				),
			},
			{
				ResourceName:      earlyAccessFeatureAddress,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccEarlyAccessFeatureAutoFlag(name, description string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_early_access_feature" "test" {
  name        = %q
  stage       = "concept"
  description = %q
}
`, name, description)
}

// Without feature_flag_id PostHog creates the flag itself; later applies must keep it linked
// rather than plan a replacement.
func TestEarlyAccessFeature_AutoCreatedFlag(t *testing.T) {
	skipIfNotAcceptance(t)

	name := strings.ReplaceAll(randomTestPrefix(), "_", "-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoCreatedFeatureFlagRetained,
		Steps: []resource.TestStep{
			{
				Config: testAccEarlyAccessFeatureAutoFlag(name, "first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(earlyAccessFeatureAddress, "feature_flag_id"),
					resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "feature_flag_key", strings.ToLower(name)),
					resource.TestCheckNoResourceAttr(earlyAccessFeatureAddress, "payload"),
					resource.TestCheckNoResourceAttr(earlyAccessFeatureAddress, "documentation_url"),
				),
			},
			{
				Config: testAccEarlyAccessFeatureAutoFlag(name, "second"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(earlyAccessFeatureAddress, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.TestCheckResourceAttr(earlyAccessFeatureAddress, "description", "second"),
			},
		},
	})
}

func TestEarlyAccessFeature_RejectsNonObjectPayload(t *testing.T) {
	skipIfNotAcceptance(t)

	name := strings.ReplaceAll(randomTestPrefix(), "_", "-")
	config := fmt.Sprintf(`
provider "posthog" {}

resource "posthog_early_access_feature" "test" {
  name    = %q
  stage   = "draft"
  payload = jsonencode([])
}
`, name)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config:      config,
			PlanOnly:    true,
			ExpectError: regexp.MustCompile(`payload must be a JSON object`),
		}},
	})
}
