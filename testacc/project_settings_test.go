package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

func TestProjectSettings_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectSettingsConfig(true, false, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_project_settings.test", "id"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "heatmaps_opt_in", "true"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "session_recording_opt_in", "false"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "surveys_opt_in", "true"),
				),
			},
		},
	})
}

func TestProjectSettings_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectSettingsConfig(true, false, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_settings.test", "heatmaps_opt_in", "true"),
				),
			},
			{
				Config: testAccProjectSettingsConfig(false, true, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_settings.test", "heatmaps_opt_in", "false"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "session_recording_opt_in", "true"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "surveys_opt_in", "false"),
				),
			},
			{
				// Exercise all six attributes end-to-end against the live API.
				Config: testAccProjectSettingsConfigAllFields(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_settings.test", "heatmaps_opt_in", "true"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "autocapture_exceptions_opt_in", "true"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "session_recording_opt_in", "true"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "surveys_opt_in", "true"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "autocapture_web_vitals_opt_in", "false"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "cookieless_server_hash_mode", "0"),
				),
			},
		},
	})
}

func TestProjectSettings_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	projectID := getProjectID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectSettingsConfig(true, false, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_project_settings.test", "id"),
				),
			},
			{
				ResourceName:            "posthog_project_settings.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           projectID,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

// TestProjectSettings_DeleteIsNoOp verifies the documented Delete semantics:
// destroying the resource stops Terraform managing the settings but does NOT
// reset them on PostHog. To avoid a coincidental match with PostHog's default,
// it sets surveys_opt_in to the OPPOSITE of the current server value, then after
// destroy asserts the value is still the one Terraform applied (a reset would
// have reverted it).
func TestProjectSettings_DeleteIsNoOp(t *testing.T) {
	skipIfNotAcceptance(t)

	projectID := getProjectID()
	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")

	before, _, err := client.GetEnvironment(context.Background(), projectID)
	if err != nil {
		t.Fatalf("reading environment before test: %v", err)
	}
	// Pick the value distinct from whatever PostHog currently reports.
	target := before.SurveysOptIn == nil || !*before.SurveysOptIn

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: func(*terraform.State) error {
			env, _, err := client.GetEnvironment(context.Background(), projectID)
			if err != nil {
				return fmt.Errorf("fetching environment after destroy: %w", err)
			}
			if env.SurveysOptIn == nil || *env.SurveysOptIn != target {
				return fmt.Errorf("expected surveys_opt_in to persist as %t after destroy (no-op delete), got %v", target, env.SurveysOptIn)
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccProjectSettingsConfig(true, false, target),
				Check:  resource.TestCheckResourceAttr("posthog_project_settings.test", "surveys_opt_in", fmt.Sprintf("%t", target)),
			},
		},
	})
}

// TestProjectSettings_Domains exercises the app_urls and recording_domains list
// attributes end-to-end, including ordering (app_urls has two entries). The
// implicit post-apply plan check confirms the lists round-trip with no ordering
// drift.
func TestProjectSettings_Domains(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectSettingsConfigDomains(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_project_settings.test", "app_urls.#", "2"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "app_urls.0", "https://app.example.com"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "app_urls.1", "https://www.example.com"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "recording_domains.#", "1"),
					resource.TestCheckResourceAttr("posthog_project_settings.test", "recording_domains.0", "https://app.example.com"),
				),
			},
		},
	})
}

func testAccProjectSettingsConfig(heatmaps, sessionRecording, surveys bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_project_settings" "test" {
  heatmaps_opt_in          = %t
  session_recording_opt_in = %t
  surveys_opt_in           = %t
}
`, heatmaps, sessionRecording, surveys)
}

func testAccProjectSettingsConfigAllFields() string {
	return `
provider "posthog" {}

resource "posthog_project_settings" "test" {
  heatmaps_opt_in               = true
  autocapture_exceptions_opt_in = true
  session_recording_opt_in      = true
  surveys_opt_in                = true
  autocapture_web_vitals_opt_in = false
  cookieless_server_hash_mode   = 0
}
`
}

func testAccProjectSettingsConfigDomains() string {
	return `
provider "posthog" {}

resource "posthog_project_settings" "test" {
  app_urls          = ["https://app.example.com", "https://www.example.com"]
  recording_domains = ["https://app.example.com"]
}
`
}
