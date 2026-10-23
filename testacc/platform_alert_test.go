package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

const platformAlertAddress = "posthog_platform_alert.test"

// skipIfNoPlatformAlerts skips when the platform alerts API is not available to this organization.
// PostHog gates it behind the `platform-alerts` feature flag and answers 403 until it is enabled.
func skipIfNoPlatformAlerts(t *testing.T) {
	t.Helper()

	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")
	_, status, err := client.GetPlatformAlert(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), "00000000-0000-0000-0000-000000000000")
	if err != nil && status == httpclient.HTTPStatusCode(http.StatusForbidden) {
		t.Skip("Skipping test: the platform-alerts feature flag is not enabled for this organization, or the key lacks alert:write")
	}
}

func testAccCheckPlatformAlertDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_platform_alert" {
			continue
		}
		_, status, err := client.GetPlatformAlert(context.Background(), projectID, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("platform alert %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
		// Deleting a logs alert through the platform removes its logs row too.
		if legacyID := rs.Primary.Attributes["legacy_configuration_id"]; legacyID != "" {
			_, status, err := client.GetLogsAlert(context.Background(), projectID, legacyID)
			if err == nil {
				return fmt.Errorf("logs alert %s behind platform alert %s still exists", legacyID, rs.Primary.ID)
			}
			if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
				return fmt.Errorf("expected 404 for logs alert %s, got %d", legacyID, status)
			}
		}
	}
	return testAccCheckLogsAlertDestroy(s)
}

func TestPlatformAlert_Basic(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoPlatformAlerts(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPlatformAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPlatformAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(platformAlertAddress, "id"),
					resource.TestCheckResourceAttrSet(platformAlertAddress, "legacy_configuration_id"),
					resource.TestCheckResourceAttr(platformAlertAddress, "name", rName),
					resource.TestCheckResourceAttr(platformAlertAddress, "source_kind", "logs"),
					resource.TestCheckResourceAttr(platformAlertAddress, "enabled", "true"),
					resource.TestCheckTypeSetElemAttr(platformAlertAddress, "logs.filters.severity_levels.*", "error"),
					// The declared defaults are what PostHog stores.
					resource.TestCheckResourceAttr(platformAlertAddress, "logs.condition.threshold_count", "100"),
					resource.TestCheckResourceAttr(platformAlertAddress, "logs.condition.threshold_operator", "above"),
					resource.TestCheckResourceAttr(platformAlertAddress, "logs.condition.window_minutes", "5"),
					resource.TestCheckResourceAttr(platformAlertAddress, "evaluation_periods", "1"),
					resource.TestCheckResourceAttr(platformAlertAddress, "datapoints_to_alarm", "1"),
					resource.TestCheckResourceAttr(platformAlertAddress, "cooldown_minutes", "0"),
					resource.TestCheckResourceAttr(platformAlertAddress, "grouping.mode", "single"),
					resource.TestCheckResourceAttr(platformAlertAddress, "grouping.keys.#", "0"),
					resource.TestCheckResourceAttr(platformAlertAddress, "grouping.max_instances", "100"),
				),
			},
			{
				ResourceName:      platformAlertAddress,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestPlatformAlert_UpdateInPlace(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoPlatformAlerts(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPlatformAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPlatformAlertBasic(rName),
			},
			{
				Config: testAccPlatformAlertAllFields(rName + "-renamed"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(platformAlertAddress, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(platformAlertAddress, "name", rName+"-renamed"),
					resource.TestCheckResourceAttr(platformAlertAddress, "logs.condition.threshold_count", "10"),
					resource.TestCheckResourceAttr(platformAlertAddress, "logs.condition.window_minutes", "10"),
					resource.TestCheckResourceAttr(platformAlertAddress, "evaluation_periods", "3"),
					resource.TestCheckResourceAttr(platformAlertAddress, "datapoints_to_alarm", "2"),
					resource.TestCheckResourceAttr(platformAlertAddress, "cooldown_minutes", "30"),
					resource.TestCheckResourceAttr(platformAlertAddress, "blocked_windows.#", "1"),
					resource.TestCheckResourceAttr(platformAlertAddress, "grouping.mode", "by_result_labels"),
					resource.TestCheckResourceAttr(platformAlertAddress, "grouping.keys.0", "service_name"),
					resource.TestCheckResourceAttr(platformAlertAddress, "grouping.max_instances", "20"),
					resource.TestCheckResourceAttrSet(platformAlertAddress, "logs.filters.filter_group_json"),
				),
			},
			// Removing settings returns them to their defaults rather than keeping the last value.
			{
				Config: testAccPlatformAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(platformAlertAddress, "logs.condition.threshold_count", "100"),
					resource.TestCheckResourceAttr(platformAlertAddress, "cooldown_minutes", "0"),
					resource.TestCheckResourceAttr(platformAlertAddress, "grouping.mode", "single"),
					resource.TestCheckNoResourceAttr(platformAlertAddress, "blocked_windows"),
					resource.TestCheckNoResourceAttr(platformAlertAddress, "logs.filters.filter_group_json"),
				),
			},
		},
	})
}

func TestPlatformAlert_DisabledDraftWithoutFilters(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoPlatformAlerts(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPlatformAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "posthog_platform_alert" "test" {
  name        = %[1]q
  source_kind = "logs"
  enabled     = false
  logs        = {}
}
`, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(platformAlertAddress, "enabled", "false"),
					resource.TestCheckNoResourceAttr(platformAlertAddress, "logs.filters"),
				),
			},
		},
	})
}

// A logs alert made through this resource keeps its logs row, so the existing destination resource
// attaches to it through legacy_configuration_id.
func TestPlatformAlert_LogsAlertDestinationAttaches(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoPlatformAlerts(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPlatformAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPlatformAlertBasic(rName) + `
resource "posthog_logs_alert_destination" "test" {
  alert_id    = posthog_platform_alert.test.legacy_configuration_id
  type        = "webhook"
  webhook_url = "https://example.com/hooks/platform"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("posthog_logs_alert_destination.test", "alert_id", platformAlertAddress, "legacy_configuration_id"),
					resource.TestCheckResourceAttrSet("posthog_logs_alert_destination.test", "id"),
				),
			},
		},
	})
}

// A moved block adopts a posthog_logs_alert with no destroy and no change, including the
// condition fields the logs config left to their defaults.
func TestPlatformAlert_MovedFromLogsAlert(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoPlatformAlerts(t)

	tests := map[string]struct {
		logsConfig     string
		platformConfig string
	}{
		"defaults": {
			logsConfig: `
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
}
`,
			platformConfig: testAccPlatformAlertBasic("%[1]s"),
		},
		"all fields": {
			logsConfig: `
resource "posthog_logs_alert" "test" {
  name                = %[1]q
  service_names       = ["checkout-api"]
  threshold_count     = 10
  threshold_operator  = "above"
  window_minutes      = 10
  evaluation_periods  = 3
  datapoints_to_alarm = 2
  cooldown_minutes    = 30
  filter_group_json = jsonencode({
    type = "AND"
    values = [{
      type = "AND"
      values = [{
        type     = "log_attribute"
        key      = "status_code"
        operator = "exact"
        value    = ["500"]
      }]
    }]
  })
  blocked_windows = [{
    start = "22:00"
    end   = "06:00"
  }]
}
`,
			platformConfig: `
resource "posthog_platform_alert" "test" {
  name                = %[1]q
  source_kind         = "logs"
  evaluation_periods  = 3
  datapoints_to_alarm = 2
  cooldown_minutes    = 30
  blocked_windows = [{
    start = "22:00"
    end   = "06:00"
  }]
  logs = {
    filters = {
      service_names = ["checkout-api"]
      filter_group_json = jsonencode({
        type = "AND"
        values = [{
          type = "AND"
          values = [{
            type     = "log_attribute"
            key      = "status_code"
            operator = "exact"
            value    = ["500"]
          }]
        }]
      })
    }
    condition = {
      threshold_count    = 10
      threshold_operator = "above"
      window_minutes     = 10
    }
  }
}
`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			rName := acctest.RandomWithPrefix("tf-acc-test")
			resource.Test(t, resource.TestCase{
				TerraformVersionChecks: []tfversion.TerraformVersionCheck{
					tfversion.SkipBelow(tfversion.Version1_8_0),
				},
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckPlatformAlertDestroy,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(test.logsConfig, rName),
					},
					{
						Config: fmt.Sprintf(test.platformConfig+`
moved {
  from = posthog_logs_alert.test
  to   = posthog_platform_alert.test
}
`, rName),
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectResourceAction(platformAlertAddress, plancheck.ResourceActionNoop),
							},
						},
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttrSet(platformAlertAddress, "legacy_configuration_id"),
							resource.TestCheckResourceAttr(platformAlertAddress, "name", rName),
						),
					},
				},
			})
		})
	}
}

func testAccPlatformAlertBasic(name string) string {
	return fmt.Sprintf(`
resource "posthog_platform_alert" "test" {
  name        = %[1]q
  source_kind = "logs"
  logs = {
    filters = {
      severity_levels = ["error"]
    }
  }
}
`, name)
}

func testAccPlatformAlertAllFields(name string) string {
	return fmt.Sprintf(`
resource "posthog_platform_alert" "test" {
  name                = %[1]q
  source_kind         = "logs"
  evaluation_periods  = 3
  datapoints_to_alarm = 2
  cooldown_minutes    = 30
  blocked_windows = [{
    start = "22:00"
    end   = "06:00"
  }]
  grouping = {
    mode          = "by_result_labels"
    keys          = ["service_name"]
    max_instances = 20
  }
  logs = {
    filters = {
      severity_levels = ["error", "fatal"]
      filter_group_json = jsonencode({
        type = "AND"
        values = [{
          type = "AND"
          values = [{
            type     = "log_attribute"
            key      = "status_code"
            operator = "exact"
            value    = ["500", "502"]
          }]
        }]
      })
    }
    condition = {
      threshold_count = 10
      window_minutes  = 10
    }
  }
}
`, name)
}
