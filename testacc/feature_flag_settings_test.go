package tests

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

func testAccFeatureFlagSettings(key, settings string) string {
	return fmt.Sprintf(`provider "posthog" {}
resource "posthog_feature_flag" "test" {
 key = %q
 %s
}`, key, settings)
}

func TestFeatureFlag_Archived(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-test-archived")
	var steps []resource.TestStep
	for _, archived := range []bool{true, false, true} {
		config := testAccFeatureFlagSettings(key, fmt.Sprintf("active = false\n archived = %t", archived))
		steps = append(steps,
			resource.TestStep{Config: config, Check: resource.TestCheckResourceAttr("posthog_feature_flag.test", "archived", fmt.Sprint(archived))},
			resource.TestStep{Config: config, PlanOnly: true},
			resource.TestStep{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"}},
		)
	}
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories, Steps: steps})
}

func TestFeatureFlag_BucketingIdentifier(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-test-bucketing")
	var steps []resource.TestStep
	for _, identifier := range []string{"device_id", "distinct_id", "device_id"} {
		config := testAccFeatureFlagSettings(key, fmt.Sprintf("bucketing_identifier = %q", identifier))
		steps = append(steps,
			resource.TestStep{Config: config, Check: resource.TestCheckResourceAttr("posthog_feature_flag.test", "bucketing_identifier", identifier)},
			resource.TestStep{Config: config, PlanOnly: true},
			resource.TestStep{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"}},
		)
	}
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories, Steps: steps})
}

func TestFeatureFlag_EvaluationContexts(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfEvaluationContextsDisabled(t)
	key := acctest.RandomWithPrefix("tf-acc-test-contexts")
	var steps []resource.TestStep
	for _, contexts := range []struct{ config, count string }{
		{`["web", "checkout"]`, "2"}, {`["mobile"]`, "1"}, {`[]`, "0"},
	} {
		config := testAccFeatureFlagSettings(key, "evaluation_contexts = "+contexts.config)
		steps = append(steps,
			resource.TestStep{Config: config, Check: resource.TestCheckResourceAttr("posthog_feature_flag.test", "evaluation_contexts.#", contexts.count)},
			resource.TestStep{Config: config, PlanOnly: true},
			resource.TestStep{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"}},
		)
	}
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories, Steps: steps})
}

func TestFeatureFlag_RemoteConfiguration(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-test-remote-config")
	config := testAccFeatureFlagSettings(key, "is_remote_configuration = true")
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "is_remote_configuration", "true"),
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "rollout_percentage", "100"),
			)},
			{Config: config, PlanOnly: true},
			{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"filters"}},
			{Config: testAccFeatureFlagSettings(key, "is_remote_configuration = false"), Check: resource.TestCheckResourceAttr("posthog_feature_flag.test", "is_remote_configuration", "false")},
			{Config: config},
			{Config: config, PlanOnly: true},
			{Config: testAccFeatureFlagSettings(key, "is_remote_configuration = true\n rollout_percentage = 0"), Check: resource.TestCheckResourceAttr("posthog_feature_flag.test", "rollout_percentage", "0")},
			{Config: testAccFeatureFlagSettings(key, "is_remote_configuration = true\n rollout_percentage = 0"), PlanOnly: true},
		},
	})
}

func TestFeatureFlag_RemoteConfigurationPayload(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-test-remote-payload")
	config := testAccFeatureFlagSettings(key, `is_remote_configuration = true
filters = jsonencode({
 groups = [{rollout_percentage = 100}]
 payloads = {"true" = jsonencode({color = "blue"})}
})`)
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config},
			{Config: config, PlanOnly: true},
			{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true},
		},
	})
}

func TestFeatureFlag_SettingsValidation(t *testing.T) {
	skipIfNotAcceptance(t)
	for _, tc := range []struct{ name, settings, message string }{
		{"archive_active", "archived = true\n active = true", `Cannot\s+archive\s+an\s+enabled\s+feature\s+flag`},
		{"device_persistence", "bucketing_identifier = \"device_id\"\n ensure_experience_continuity = true", `Cannot\s+enable\s+'persist\s+across\s+authentication\s+steps'`},
		{"invalid_bucketing", `bucketing_identifier = "session_id"`, "Attribute bucketing_identifier value must be one of"},
		{"context_case", `evaluation_contexts = ["Web"]`, "Invalid evaluation context"},
		{"context_whitespace", `evaluation_contexts = [" web "]`, "Invalid evaluation context"},
		{"context_empty", `evaluation_contexts = [""]`, "Invalid Attribute Value Length"},
		{"remote_zero_rollout", "is_remote_configuration = true\n rollout_percentage = 0", "Invalid remote configuration rollout"},
		{"remote_payload_only", "is_remote_configuration = true\n filters = jsonencode({payloads = {\"true\" = \"{}\"}})", "Invalid remote configuration filters"},
		{"remote_payload_with_rollout", "is_remote_configuration = true\n rollout_percentage = 100\n filters = jsonencode({payloads = {\"true\" = \"{}\"}})", "Invalid remote configuration filters"},
		{"remote_conflicting_rollout", "is_remote_configuration = true\n rollout_percentage = 100\n filters = jsonencode({groups = [{rollout_percentage = 50}]})", "Invalid remote configuration filters"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    []resource.TestStep{{Config: testAccFeatureFlagSettings(acctest.RandomWithPrefix("tf-acc-test-invalid"), tc.settings), ExpectError: regexp.MustCompile(tc.message)}},
			})
		})
	}
}

func TestFeatureFlag_SettingsSavedPlan(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfEvaluationContextsDisabled(t)
	key := acctest.RandomWithPrefix("tf-acc-test-settings-plan")
	var flagID string
	config := testAccFeatureFlagSettings(key, "active = false\n name = \"unrelated update\"")
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: testAccFeatureFlagSettings(key, "active = false"), Check: func(s *terraform.State) error {
				flagID = s.RootModule().Resources["posthog_feature_flag.test"].Primary.ID
				return nil
			}},
			{Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{runAfterPlan(func(ctx context.Context) error { return changeFeatureFlagSettings(ctx, flagID, key) })}},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "archived", "true"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "bucketing_identifier", "device_id"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "is_remote_configuration", "true"),
					resource.TestCheckTypeSetElemAttr("posthog_feature_flag.test", "evaluation_contexts.*", "web"),
				)},
			{Config: config, PlanOnly: true},
			{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"}},
			{Config: testAccFeatureFlagSettings(key, `active = false
archived = false
bucketing_identifier = "distinct_id"
evaluation_contexts = []
is_remote_configuration = false`), Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "archived", "false"),
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "bucketing_identifier", "distinct_id"),
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "evaluation_contexts.#", "0"),
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "is_remote_configuration", "false"),
			)},
		},
	})
}

// skipIfEvaluationContextsDisabled skips a test that needs PostHog to store evaluation contexts.
// PostHog ignores them unless evaluation contexts are enabled for the organization, through the
// flag-evaluation-tags feature flag, which a local stack may not have.
func skipIfEvaluationContextsDisabled(t *testing.T) {
	t.Helper()
	ctx, projectID := context.Background(), os.Getenv("POSTHOG_PROJECT_ID")
	client := testAccAPIClient()
	contexts := []string{"web"}
	flag, err := client.CreateFeatureFlag(ctx, projectID, httpclient.FeatureFlagRequest{
		Key:                acctest.RandomWithPrefix("tf-acc-contexts-probe"),
		EvaluationContexts: &contexts,
	})
	if err != nil {
		t.Fatalf("failed to check whether PostHog stores evaluation contexts: %s", err)
	}
	if _, err := client.DeleteFeatureFlag(ctx, projectID, fmt.Sprintf("%d", flag.ID)); err != nil {
		t.Logf("failed to delete probe flag %d, delete it by hand: %s", flag.ID, err)
	}
	if len(flag.EvaluationContexts) == 0 {
		t.Skip("PostHog does not store evaluation contexts for this organization; enable the flag-evaluation-tags feature flag to run this test")
	}
}

// runAfterPlan runs its function between plan and apply, standing in for a change made in
// PostHog after a saved plan was made.
type runAfterPlan func(context.Context) error

func (f runAfterPlan) CheckPlan(ctx context.Context, _ plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	resp.Error = f(ctx)
}

func changeFeatureFlagSettings(ctx context.Context, id, key string) error {
	client := testAccAPIClient()
	enabled, disabled, identifier, contexts := true, false, "device_id", []string{"web"}
	_, _, err := client.UpdateFeatureFlag(ctx, os.Getenv("POSTHOG_PROJECT_ID"), id, httpclient.FeatureFlagRequest{
		Key: key, Active: &disabled, Archived: &enabled, BucketingIdentifier: &identifier,
		EvaluationContexts: &contexts, IsRemoteConfiguration: &enabled,
	})
	return err
}

func TestFeatureFlag_SettingsDrift(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-test-settings-drift")
	var flagID string
	config := testAccFeatureFlagSettings(key, `active = false
archived = false
bucketing_identifier = "distinct_id"
evaluation_contexts = []
is_remote_configuration = false`)
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, Check: func(s *terraform.State) error {
				flagID = s.RootModule().Resources["posthog_feature_flag.test"].Primary.ID
				return nil
			}},
			{Config: config, PreConfig: func() {
				if err := changeFeatureFlagSettings(context.Background(), flagID, key); err != nil {
					t.Fatal(err)
				}
			}, PlanOnly: true, ExpectNonEmptyPlan: true},
			{Config: config, Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "archived", "false"),
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "bucketing_identifier", "distinct_id"),
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "evaluation_contexts.#", "0"),
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "is_remote_configuration", "false"),
			)},
			{Config: config, PlanOnly: true},
		},
	})
}
