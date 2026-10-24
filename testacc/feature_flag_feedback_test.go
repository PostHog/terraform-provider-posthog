package tests

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

func TestFeatureFlag_NameRoundTrip(t *testing.T) {
	skipIfNotAcceptance(t)
	for _, name := range []string{"", "  padded description  ", " \t\n"} {
		t.Run(fmt.Sprintf("name=%q", name), func(t *testing.T) {
			key := acctest.RandomWithPrefix("tf-acc-name")
			config := testAccFeatureFlagWithName(key, name, true)
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{Config: config, Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("posthog_feature_flag.test", "name", name),
						checkFeatureFlagAPIField("name", strings.TrimSpace(name)),
					)},
					{Config: config, PlanOnly: true},
					{Config: testAccFeatureFlagWithName(key, "updated", true)},
					{Config: config, Check: resource.TestCheckResourceAttr("posthog_feature_flag.test", "name", name)},
				},
			})
		})
	}
}

func TestFeatureFlag_EvaluationRuntime(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-runtime")
	config := func(runtime string) string {
		return fmt.Sprintf(`provider "posthog" {}
resource "posthog_feature_flag" "test" {
 key = %q
 evaluation_runtime = %q
}`, key, runtime)
	}
	var steps []resource.TestStep
	for _, runtime := range []string{"server", "client", "all"} {
		steps = append(steps,
			resource.TestStep{Config: config(runtime), Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "evaluation_runtime", runtime),
				checkFeatureFlagAPIField("evaluation_runtime", runtime),
			)},
			resource.TestStep{Config: config(runtime), PlanOnly: true},
			resource.TestStep{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"}},
		)
	}
	resource.Test(t, resource.TestCase{PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories, Steps: steps})
}

// An omitted runtime adopts remote changes and survives unrelated updates/imports.
func TestFeatureFlag_EvaluationRuntimeDefaultAndOmitted(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-runtime-default")
	config := testAccFeatureFlagBasic(key)
	var flagID string
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "evaluation_runtime", "all"),
				func(s *terraform.State) error {
					flagID = s.RootModule().Resources["posthog_feature_flag.test"].Primary.ID
					return nil
				},
			)},
			{Config: config, PreConfig: func() { setFeatureFlagRuntime(t, flagID, key, "server") }, PlanOnly: true},
			{Config: testAccFeatureFlagWithName(key, "unrelated update", true), Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("posthog_feature_flag.test", "evaluation_runtime", "server"),
				checkFeatureFlagAPIField("evaluation_runtime", "server"),
			)},
			{ResourceName: "posthog_feature_flag.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"filters", "rollout_percentage"}},
		},
	})
}

// A configured runtime detects an out-of-band change and restores the desired value.
func TestFeatureFlag_EvaluationRuntimeDrift(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-runtime-drift")
	config := fmt.Sprintf(`provider "posthog" {}
resource "posthog_feature_flag" "test" {
 key = %q
 evaluation_runtime = "server"
}`, key)
	var flagID string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) }, ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, Check: func(s *terraform.State) error {
				flagID = s.RootModule().Resources["posthog_feature_flag.test"].Primary.ID
				return nil
			}},
			{Config: config, PreConfig: func() { setFeatureFlagRuntime(t, flagID, key, "client") }, PlanOnly: true, ExpectNonEmptyPlan: true},
			{Config: config, Check: checkFeatureFlagAPIField("evaluation_runtime", "server")},
		},
	})
}

func setFeatureFlagRuntime(t *testing.T, id, key, runtime string) {
	t.Helper()
	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")
	_, _, err := client.UpdateFeatureFlag(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), id, httpclient.FeatureFlagRequest{Key: key, EvaluationRuntime: &runtime})
	if err != nil {
		t.Fatal(err)
	}
}

func TestFeatureFlag_EvaluationRuntimeInvalid(t *testing.T) {
	skipIfNotAcceptance(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{Config: `provider "posthog" {}
resource "posthog_feature_flag" "test" {
 key = "invalid-runtime"
 evaluation_runtime = "browser"
}`, ExpectError: regexp.MustCompile(`Attribute evaluation_runtime value must be one of`)}},
	})
}

func checkFeatureFlagAPIField(field, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources["posthog_feature_flag.test"]
		client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")
		flag, _, err := client.GetFeatureFlag(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID)
		if err != nil {
			return err
		}
		var got *string
		switch field {
		case "name":
			if flag.Name != nil {
				got = flag.Name
			}
		case "evaluation_runtime":
			if flag.EvaluationRuntime != nil {
				got = flag.EvaluationRuntime
			}
		}
		if got == nil {
			return fmt.Errorf("API %s is null, want %q", field, want)
		}
		if *got != want {
			return fmt.Errorf("API %s = %q, want %q", field, *got, want)
		}
		return nil
	}
}

func TestFeatureFlag_EvaluationRuntimeSavedPlan(t *testing.T) {
	skipIfNotAcceptance(t)
	key := acctest.RandomWithPrefix("tf-acc-runtime-saved-plan")
	var flagID string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) }, ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: testAccFeatureFlagBasic(key), Check: func(s *terraform.State) error {
				flagID = s.RootModule().Resources["posthog_feature_flag.test"].Primary.ID
				return nil
			}},
			{
				Config:           testAccFeatureFlagWithName(key, "unrelated update", true),
				ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{changeFlagRuntimeAfterPlan{flagID: &flagID, key: key}}},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "evaluation_runtime", "server"),
					checkFeatureFlagAPIField("evaluation_runtime", "server"),
				),
			},
		},
	})
}

type changeFlagRuntimeAfterPlan struct {
	flagID *string
	key    string
}

func (c changeFlagRuntimeAfterPlan) CheckPlan(ctx context.Context, _ plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")
	runtime := "server"
	_, _, resp.Error = client.UpdateFeatureFlag(ctx, os.Getenv("POSTHOG_PROJECT_ID"), *c.flagID, httpclient.FeatureFlagRequest{Key: c.key, EvaluationRuntime: &runtime})
}
