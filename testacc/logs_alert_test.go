package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// testAccCheckLogsAlertDestroy verifies the log alert has been destroyed.
func testAccCheckLogsAlertDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_logs_alert" {
			continue
		}

		_, status, err := client.GetLogsAlert(context.Background(), projectID, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("logs alert %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

// TestLogsAlert_Basic tests creating a log alert with minimal configuration.
func TestLogsAlert_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "severity_levels.#", "1"),
					resource.TestCheckTypeSetElemAttr("posthog_logs_alert.test", "severity_levels.*", "error"),
					resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "state"),
					// Server defaults must land in state rather than staying unknown.
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "threshold_count", "100"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "threshold_operator", "above"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "window_minutes", "5"),
				),
			},
		},
	})
}

// TestLogsAlert_AllFields tests creating a log alert with all optional fields.
func TestLogsAlert_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "enabled", "true"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "threshold_count", "10"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "threshold_operator", "above"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "window_minutes", "10"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "evaluation_periods", "3"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "datapoints_to_alarm", "2"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "cooldown_minutes", "30"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "service_names.#", "1"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "blocked_windows.#", "1"),
					resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "filter_group_json"),
				),
			},
		},
	})
}

// TestLogsAlert_Update tests updating a log alert's threshold and filters in place.
func TestLogsAlert_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertWithThreshold(rName, 10, "error"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "threshold_count", "10"),
					resource.TestCheckTypeSetElemAttr("posthog_logs_alert.test", "severity_levels.*", "error"),
				),
			},
			{
				Config: testAccLogsAlertWithThreshold(rName, 250, "warn"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "threshold_count", "250"),
					resource.TestCheckTypeSetElemAttr("posthog_logs_alert.test", "severity_levels.*", "warn"),
				),
			},
		},
	})
}

// TestLogsAlert_Import tests importing an existing log alert.
func TestLogsAlert_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertBasic(rName),
			},
			{
				ResourceName:      "posthog_logs_alert.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestLogsAlert_OmittedName covers the config shape the schema explicitly permits: no
// name at all, adopting PostHog's "Untitled alert" server default. This is the path that
// fails with "inconsistent result after apply" if name is not Computed.
func TestLogsAlert_OmittedName(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource "posthog_logs_alert" "test" {
  severity_levels = ["error"]
  enabled         = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "name", "Untitled alert"),
					resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "id"),
				),
			},
		},
	})
}

// TestLogsAlert_RejectsNeverFiringConfig verifies the N-of-M cross-field check runs at
// plan time rather than producing an alert that can never fire.
func TestLogsAlert_RejectsNeverFiringConfig(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name                = %[1]q
  severity_levels     = ["error"]
  evaluation_periods  = 2
  datapoints_to_alarm = 5
}
`, rName),
				ExpectError: regexp.MustCompile(`Alert can never fire`),
			},
		},
	})
}

// TestLogsAlert_RejectsOverlappingWindows verifies quiet-hours overlap is caught at plan
// time, since PostHog would otherwise merge the windows and break every subsequent apply.
func TestLogsAlert_RejectsOverlappingWindows(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]

  blocked_windows = [
    { start = "01:00", end = "03:00" },
    { start = "02:00", end = "04:00" },
  ]
}
`, rName),
				ExpectError: regexp.MustCompile(`overlap`),
			},
		},
	})
}

// TestLogsAlert_RejectsEnabledAlertWithoutFilters verifies the at-least-one-filter rule is
// enforced at plan time instead of surfacing as a mid-apply API error.
func TestLogsAlert_RejectsEnabledAlertWithoutFilters(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name = %[1]q
}
`, rName),
				ExpectError: regexp.MustCompile(`no filters`),
			},
		},
	})
}

// TestLogsAlert_ClearsFilter removes a previously-set filter and asserts it is actually
// gone server-side — the whole-object replace semantics the client relies on.
func TestLogsAlert_ClearsFilter(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
  service_names   = ["checkout-api"]
}
`, rName),
				Check: resource.TestCheckResourceAttr("posthog_logs_alert.test", "service_names.#", "1"),
			},
			{
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
}
`, rName),
				Check: resource.TestCheckNoResourceAttr("posthog_logs_alert.test", "service_names.#"),
			},
		},
	})
}

// TestLogsAlert_UpdateOmittingComputedAttribute removes an Optional+Computed attribute from
// a config that keeps a dependent one. The attribute retains its last applied value, so the
// apply must succeed — validating it against the server default instead would reject a valid
// config with "Alert can never fire: datapoints_to_alarm (5) must not exceed
// evaluation_periods (1)".
func TestLogsAlert_UpdateOmittingComputedAttribute(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name                = %[1]q
  severity_levels     = ["error"]
  evaluation_periods  = 10
  datapoints_to_alarm = 5
}
`, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "evaluation_periods", "10"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "datapoints_to_alarm", "5"),
				),
			},
			{
				// evaluation_periods dropped from config; it keeps the applied value of 10.
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name                = %[1]q
  severity_levels     = ["error"]
  datapoints_to_alarm = 5
}
`, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "evaluation_periods", "10"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "datapoints_to_alarm", "5"),
				),
			},
		},
	})
}

// TestLogsAlert_RejectsInvalidOperator verifies schema validation runs before any API call.
func TestLogsAlert_RejectsInvalidOperator(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccLogsAlertInvalidOperator(rName),
				ExpectError: regexp.MustCompile(`Attribute threshold_operator value must be one of`),
			},
		},
	})
}

func testAccLogsAlertBasic(name string) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
}
`, name)
}

func testAccLogsAlertAllFields(name string) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name                = %[1]q
  enabled             = true
  severity_levels     = ["error", "fatal"]
  service_names       = ["checkout-api"]

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

  threshold_count     = 10
  threshold_operator  = "above"
  window_minutes      = 10
  evaluation_periods  = 3
  datapoints_to_alarm = 2
  cooldown_minutes    = 30

  blocked_windows = [{
    start = "22:00"
    end   = "06:00"
  }]
}
`, name)
}

func testAccLogsAlertWithThreshold(name string, threshold int, severity string) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = [%[3]q]
  threshold_count = %[2]d
  window_minutes  = 10
}
`, name, threshold, severity)
}

func testAccLogsAlertInvalidOperator(name string) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name               = %[1]q
  severity_levels    = ["error"]
  threshold_operator = "sideways"
}
`, name)
}
