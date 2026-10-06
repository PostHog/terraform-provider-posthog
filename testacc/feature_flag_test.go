package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
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

// TestFeatureFlag_IgnoreFilterFieldsRoundTripsAndUpdates exercises the config→apply→state
// round-trip of the ignore_filter_fields Set attribute (the schema-wiring path that had the
// import bug), plus updating the set. No wiring keys are present here — the point is that the
// attribute itself applies, stays in state without a perpetual diff, and can be changed.
func TestFeatureFlag_IgnoreFilterFieldsRoundTripsAndUpdates(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagWithIgnore(rKey, `["payloads", "super_groups"]`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "ignore_filter_fields.#", "2"),
				),
			},
			{
				// The Set attribute round-trips with no perpetual diff.
				Config:   testAccFeatureFlagWithIgnore(rKey, `["payloads", "super_groups"]`),
				PlanOnly: true,
			},
			{
				// Updating the set applies cleanly.
				Config: testAccFeatureFlagWithIgnore(rKey, `["holdout"]`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "ignore_filter_fields.#", "1"),
				),
			},
		},
	})
}

// TestFeatureFlag_DefaultIgnoresServerWiredKeys is the end-to-end test of the feature's
// core value. A flag carrying Early Access Feature / Experiment wiring (super_groups,
// holdout_groups) — the kind PostHog attaches server-side and which a plain update strips —
// is created OUTSIDE Terraform, then imported. The config manages only groups and leaves
// ignore_filter_fields at its default. Because the default ignores those wiring keys, the
// imported flag matches the config and a re-plan is empty; without the feature it would
// show a perpetual diff.
func TestFeatureFlag_DefaultIgnoresServerWiredKeys(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")
	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "acceptance-test")

	var flagID int64
	createWiredFlag := func() {
		active := true
		f, err := (&client).CreateFeatureFlag(context.Background(), projectID, httpclient.FeatureFlagRequest{
			Key:    rKey,
			Active: &active,
			Filters: map[string]interface{}{
				"groups":         []interface{}{map[string]interface{}{"rollout_percentage": 100}},
				"super_groups":   []interface{}{map[string]interface{}{"rollout_percentage": 100}},
				"holdout_groups": []interface{}{map[string]interface{}{"rollout_percentage": 50}},
			},
		})
		if err != nil {
			t.Fatalf("Failed to create server-wired flag externally: %v", err)
		}
		flagID = f.ID
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig:          createWiredFlag,
				Config:             testAccFeatureFlagGroupsOnly(rKey),
				ResourceName:       "posthog_feature_flag.test",
				ImportState:        true,
				ImportStatePersist: true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return fmt.Sprintf("%s/%d", projectID, flagID), nil
				},
			},
			{
				// Default ignore set drops super_groups/holdout_groups → no perpetual diff.
				Config:   testAccFeatureFlagGroupsOnly(rKey),
				PlanOnly: true,
			},
		},
	})
}

// TestFeatureFlag_DefaultIgnoresEarlyAccessEnrollment links an Early Access Feature to a
// Terraform-managed flag outside Terraform, as the PostHog UI does. In an active stage PostHog
// marks the flag with filters.feature_enrollment and keeps it when an update omits it, so the
// default ignore set must cover it: otherwise every plan shows a diff and the apply fails with
// "Provider produced inconsistent result after apply".
func TestFeatureFlag_DefaultIgnoresEarlyAccessEnrollment(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")
	host := os.Getenv("POSTHOG_HOST")
	apiKey := os.Getenv("POSTHOG_API_KEY")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")
	client := httpclient.NewDefaultClient(host, apiKey, "acceptance-test")

	var flagID, featureID string
	captureFlagID := func(s *terraform.State) error {
		flagID = s.RootModule().Resources["posthog_feature_flag.test"].Primary.ID
		return nil
	}
	checkEnrolled := func(s *terraform.State) error {
		flag, _, err := (&client).GetFeatureFlag(context.Background(), projectID, flagID)
		if err != nil {
			return err
		}
		if flag.Filters["feature_enrollment"] != true {
			return fmt.Errorf("feature_enrollment = %v, want true", flag.Filters["feature_enrollment"])
		}
		return nil
	}
	deleteFeature := func() {
		if featureID == "" {
			return
		}
		if err := deleteEarlyAccessFeatureRaw(host, apiKey, projectID, featureID); err != nil {
			t.Errorf("Failed to delete early access feature %s: %v", featureID, err)
		}
		featureID = ""
	}
	// The last step deletes the feature, since PostHog refuses to delete a flag one still uses.
	// If an earlier step fails, destroying the flag fails too; this at least removes the feature.
	t.Cleanup(deleteFeature)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagLinkedToEarlyAccess(rKey, "linked"),
				Check:  captureFlagID,
			},
			{
				PreConfig: func() {
					id, err := createEarlyAccessFeatureRaw(host, apiKey, projectID, rKey, flagID)
					if err != nil {
						t.Fatalf("Failed to link an early access feature: %v", err)
					}
					featureID = id
				},
				Config:   testAccFeatureFlagLinkedToEarlyAccess(rKey, "linked"),
				PlanOnly: true,
			},
			{
				// An update while linked sends filters without the marker; it must apply cleanly
				// and PostHog must keep the marker.
				Config: testAccFeatureFlagLinkedToEarlyAccess(rKey, "linked renamed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "name", "linked renamed"),
					checkEnrolled,
				),
			},
			{
				// Opting out of the default ignore set must also apply: the linked EAF still
				// owns its marker, so omitting it from filters cannot remove it.
				Config: testAccFeatureFlagLinkedToEarlyAccess(rKey, "linked renamed", "[]"),
				Check:  checkEnrolled,
			},
			{
				Config:   testAccFeatureFlagLinkedToEarlyAccess(rKey, "linked renamed", "[]"),
				PlanOnly: true,
			},
			{
				PreConfig: deleteFeature,
				Config:    testAccFeatureFlagLinkedToEarlyAccess(rKey, "linked renamed"),
			},
			{
				Config:   testAccFeatureFlagLinkedToEarlyAccess(rKey, "linked renamed"),
				PlanOnly: true,
			},
		},
	})
}

func testAccFeatureFlagLinkedToEarlyAccess(key, name string, ignoreFilterFields ...string) string {
	ignore := ""
	if len(ignoreFilterFields) != 0 {
		ignore = "ignore_filter_fields = " + ignoreFilterFields[0]
	}
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = %q
  active = true
  %s

  filters = jsonencode({
    groups = [{ rollout_percentage = 100 }]
  })
}
`, key, name, ignore)
}

// createEarlyAccessFeatureRaw links a beta Early Access Feature to a flag via raw HTTP (the
// provider has no resource for it), which is what stamps feature_enrollment onto the flag.
func createEarlyAccessFeatureRaw(host, apiKey, projectID, name, flagID string) (string, error) {
	body := fmt.Sprintf(`{"name":%q,"stage":"beta","feature_flag_id":%s}`, name, flagID)
	url := fmt.Sprintf("%s/api/projects/%s/early_access_feature/", strings.TrimRight(host, "/"), projectID)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("create early access feature returned %d: %s", resp.StatusCode, string(respBody))
	}
	var parsed struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil || parsed.ID == "" {
		return "", fmt.Errorf("parse early access feature response: %v (body: %s)", err, string(respBody))
	}
	return parsed.ID, nil
}

func deleteEarlyAccessFeatureRaw(host, apiKey, projectID, featureID string) error {
	url := fmt.Sprintf("%s/api/projects/%s/early_access_feature/%s/", strings.TrimRight(host, "/"), projectID, featureID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete early access feature returned %d: %s", resp.StatusCode, string(respBody))
	}
	return nil
}

// TestFeatureFlag_FiltersOmittedRolloutNoPerpetualDiff guards the interaction between the
// drift-preserving normalizer (normalizeFeatureFlagFiltersForState) and semantic equality:
// a group that omits rollout_percentage must not drift. The normalizer keeps every non-empty
// API field, so if PostHog ever echoed a non-empty default (e.g. rollout_percentage=100) for
// a field the config omits, that field would be stored in state and surface as a permanent
// plan. Verified against the live API: PostHog returns the group WITHOUT rollout_percentage
// when it is omitted, so the re-plan stays empty. This test pins that behavior.
func TestFeatureFlag_FiltersOmittedRolloutNoPerpetualDiff(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagFiltersNoRollout(rKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "filters"),
				),
			},
			{
				// Re-plan with the identical config whose group omits rollout_percentage. A
				// non-empty plan would mean the normalizer kept a server-injected default.
				Config:   testAccFeatureFlagFiltersNoRollout(rKey),
				PlanOnly: true,
			},
		},
	})
}

// TestFeatureFlag_FiltersNoPerpetualDiff is the end-to-end regression test for issue #111.
// It configures filters as a RAW JSON string whose object keys are in a natural,
// non-alphabetical order (key, type, operator, value) — unlike jsonencode, which sorts
// keys. PostHog's API returns the same filters with alphabetically-sorted keys, so with
// the old types.String attribute the second (PlanOnly) step would report a perpetual diff.
// With jsontypes.Normalized the two are compared semantically, so the re-plan is empty.
func TestFeatureFlag_FiltersNoPerpetualDiff(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeatureFlagRawFilters(rKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "filters"),
				),
			},
			{
				// Re-plan with the identical raw, non-alphabetical config. A non-empty plan
				// here is exactly the issue #111 perpetual diff; PlanOnly fails on any change.
				Config:   testAccFeatureFlagRawFilters(rKey),
				PlanOnly: true,
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

// testAccFeatureFlagRawFilters builds a config whose filters is a RAW JSON string (not
// jsonencode) with object keys deliberately in non-alphabetical order, to reproduce the
// key-ordering drift from issue #111.
func testAccFeatureFlagRawFilters(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Raw Filters No Perpetual Diff"
  active = true

  filters = "{\"groups\":[{\"properties\":[{\"key\":\"email\",\"type\":\"person\",\"operator\":\"exact\",\"value\":[\"test@example.com\"]}],\"rollout_percentage\":100}]}"
}
`, key)
}

// testAccFeatureFlagWithIgnore manages a simple flag with an explicit ignore_filter_fields.
func testAccFeatureFlagWithIgnore(key, ignoreHCL string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                  = %q
  active               = true
  filters              = jsonencode({ groups = [{ rollout_percentage = 100 }] })
  ignore_filter_fields = %s
}
`, key, ignoreHCL)
}

// testAccFeatureFlagGroupsOnly manages just a single rollout group — no super_groups,
// holdout_groups, or other wiring — relying on the default ignore_filter_fields.
func testAccFeatureFlagGroupsOnly(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  active = true

  filters = jsonencode({
    groups = [{ rollout_percentage = 100 }]
  })
}
`, key)
}

// testAccFeatureFlagFiltersNoRollout builds a config whose single group omits
// rollout_percentage, to verify PostHog does not echo a non-empty default that the
// drift-preserving normalizer would then surface as a perpetual diff.
func testAccFeatureFlagFiltersNoRollout(key string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Filters without rollout"
  active = true

  filters = jsonencode({
    groups = [{
      properties = [{
        key      = "email"
        type     = "person"
        operator = "exact"
        value    = ["test@example.com"]
      }]
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

// TestFeatureFlag_CohortTargeting is the regression test for the cohort-targeting
// inconsistent-result bug: when a flag references a cohort, the API injects a server-computed
// "cohort_name" into the cohort property, which the keep-all normalizer used to store — absent
// from config, so create failed and re-plans drifted. The fix strips it recursively.
//
// The cohort is created via raw HTTP (no posthog_cohort resource on main) and soft-deleted in
// cleanup. It's created before the steps because TestStep.Config is an eager string (no lazy
// config func), so the id must be known when the config is rendered.
func TestFeatureFlag_CohortTargeting(t *testing.T) {
	skipIfNotAcceptance(t)

	rKey := acctest.RandomWithPrefix("tf-acc-test")
	host := os.Getenv("POSTHOG_HOST")
	apiKey := os.Getenv("POSTHOG_API_KEY")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	cohortID, err := createCohortRaw(host, apiKey, projectID, acctest.RandomWithPrefix("tf-acc-cohort"))
	if err != nil {
		t.Fatalf("Failed to create cohort externally: %v", err)
	}
	t.Cleanup(func() {
		if err := softDeleteCohortRaw(host, apiKey, projectID, cohortID); err != nil {
			t.Logf("Warning: failed to soft-delete cohort %d: %v", cohortID, err)
		}
	})

	config := testAccFeatureFlagCohortTargeting(rKey, cohortID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Apply must succeed (previously errored with inconsistent-result).
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "filters"),
				),
			},
			// Re-plan must be empty (no cohort_name drift).
			{
				Config:   config,
				PlanOnly: true,
			},
		},
	})
}

// createCohortRaw POSTs a minimal cohort via raw HTTP (no posthog_cohort resource on main).
func createCohortRaw(host, apiKey, projectID, name string) (int64, error) {
	body, _ := json.Marshal(map[string]interface{}{"name": name})
	url := fmt.Sprintf("%s/api/projects/%s/cohorts/", strings.TrimRight(host, "/"), projectID)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("create cohort returned %d: %s", resp.StatusCode, string(respBody))
	}
	var parsed struct {
		ID int64 `json:"id"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return 0, fmt.Errorf("parse cohort response: %w (body: %s)", err, string(respBody))
	}
	if parsed.ID == 0 {
		return 0, fmt.Errorf("cohort response had no id: %s", string(respBody))
	}
	return parsed.ID, nil
}

func softDeleteCohortRaw(host, apiKey, projectID string, cohortID int64) error {
	body, _ := json.Marshal(map[string]interface{}{"deleted": true})
	url := fmt.Sprintf("%s/api/projects/%s/cohorts/%d/", strings.TrimRight(host, "/"), projectID, cohortID)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("soft-delete cohort returned %d: %s", resp.StatusCode, string(respBody))
	}
	return nil
}

func testAccFeatureFlagCohortTargeting(key string, cohortID int64) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key    = %q
  name   = "Cohort Targeting"
  active = true

  filters = jsonencode({
    groups = [{
      properties = [{
        type  = "cohort"
        key   = "id"
        value = %d
      }]
      rollout_percentage = 100
    }]
  })
}
`, key, cohortID)
}

// TestFeatureFlag_ExternalDeletion tests that Terraform detects when a feature flag
// is soft-deleted externally (e.g., via the PostHog UI) and restores it on apply.
// This verifies that soft-deleted flags are kept in state and restored via update.
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
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "deleted", "false"),
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
			// Terraform should detect drift (deleted changed from false to true)
			// and plan an update to restore it
			{
				PreConfig: func() {
					// Delete the feature flag externally using the provider's client
					_, err := (&client).DeleteFeatureFlag(context.Background(), projectID, flagID)
					if err != nil {
						t.Fatalf("Failed to delete feature flag externally: %v", err)
					}
				},
				Config:             testAccFeatureFlagWithName(rKey, "External Deletion Test", true),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// Step 3: Apply to restore the soft-deleted flag
			{
				Config: testAccFeatureFlagWithName(rKey, "External Deletion Test", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rKey),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "deleted", "false"),
				),
			},
		},
	})
}

// TestFeatureFlag_CreateUsageDashboard confirms that creating a flag does not leave behind
// a saved dashboard. PostHog now renders usage charts on the flag's Usage tab instead.
func TestFeatureFlag_CreateUsageDashboard(t *testing.T) {
	skipIfNotAcceptance(t)

	key := acctest.RandomWithPrefix("tf-acc-test")
	host := os.Getenv("POSTHOG_HOST")
	apiKey := os.Getenv("POSTHOG_API_KEY")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: testAccFeatureFlagBasic(key),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", key),
				testCheckUsageDashboardCount(t, host, apiKey, projectID, key, 0),
			),
		}},
	})
}

// The removed create-time option must fail before a flag is sent to PostHog.
func TestFeatureFlag_CreateUsageDashboardOptInRejected(t *testing.T) {
	skipIfNotAcceptance(t)

	key := acctest.RandomWithPrefix("tf-acc-test")
	config := fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                    = %q
  create_usage_dashboard = true
}
`, key)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config:      config,
			PlanOnly:    true,
			ExpectError: regexp.MustCompile(`Unsupported create_usage_dashboard`),
		}},
	})
}

// A project change replaces an existing flag, so the removed create-only option
// must fail during planning before Terraform destroys the original flag.
func TestFeatureFlag_CreateUsageDashboardReplacementOptInRejected(t *testing.T) {
	skipIfNotAcceptance(t)

	key := acctest.RandomWithPrefix("tf-acc-test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")
	replacementProjectID := projectID + "9999"
	config := func(projectID string, createUsageDashboard bool) string {
		return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                    = %q
  project_id             = %q
  create_usage_dashboard = %t
}
`, key, projectID, createUsageDashboard)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config(projectID, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_feature_flag.test", "id"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "project_id", projectID),
				),
			},
			{
				Config:      config(replacementProjectID, true),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Unsupported create_usage_dashboard`),
			},
		},
	})
}

// usageDashboardName is the name of a legacy generated usage dashboard.
func usageDashboardName(flagKey string) string {
	return fmt.Sprintf("Generated Dashboard: %s Usage", flagKey)
}

// searchUsageDashboards looks for legacy dashboards via the raw dashboards endpoint
// (there is no list method on the client) and returns only those whose name exactly matches
// the auto-generated usage-dashboard name. Using ?search=<key> avoids paginating the project's
// full dashboard list; the exact-name filter guards against fuzzy search matches.
func searchUsageDashboards(host, apiKey, projectID, flagKey string) ([]httpclient.Dashboard, error) {
	endpoint := fmt.Sprintf("%s/api/projects/%s/dashboards/?search=%s", host, projectID, url.QueryEscape(flagKey))
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("dashboards list returned status %d: %s", resp.StatusCode, string(body))
	}

	var page struct {
		Results []httpclient.Dashboard `json:"results"`
	}
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, err
	}

	wantName := usageDashboardName(flagKey)
	var matched []httpclient.Dashboard
	for _, d := range page.Results {
		if d.Name != nil && *d.Name == wantName {
			matched = append(matched, d)
		}
	}
	return matched, nil
}

// testCheckUsageDashboardCount asserts the number of saved usage dashboards for a flag key.
func testCheckUsageDashboardCount(t *testing.T, host, apiKey, projectID, flagKey string, want int) resource.TestCheckFunc {
	return func(*terraform.State) error {
		found, err := searchUsageDashboards(host, apiKey, projectID, flagKey)
		if err != nil {
			return fmt.Errorf("searching usage dashboards for %q: %w", flagKey, err)
		}
		if len(found) != want {
			return fmt.Errorf("expected %d usage dashboard(s) named %q, found %d", want, usageDashboardName(flagKey), len(found))
		}
		t.Logf("usage dashboards for %q: found %d (want %d)", flagKey, len(found), want)
		return nil
	}
}
