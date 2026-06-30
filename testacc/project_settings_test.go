package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
