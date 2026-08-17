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

// TestLogsAlert_BlockedWindowsLifecycle walks quiet hours through add, change, remove and
// re-add. The removal step is the one with teeth: schedule_restriction is sent without
// omitempty so an absent set clears it, and this asserts PostHog agrees.
func TestLogsAlert_BlockedWindowsLifecycle(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "22:00", end = "06:00" },
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "blocked_windows.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_logs_alert.test", "blocked_windows.*", map[string]string{
						"start": "22:00",
						"end":   "06:00",
					}),
				),
			},
			// Replan the same config: PostHog returns windows in its own order, so a list
			// would show a spurious diff here where a set does not.
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "22:00", end = "06:00" },
`),
				PlanOnly: true,
			},
			// Two windows, neither crossing midnight and with a gap between them. PostHog
			// re-derives the stored windows from a merged daily timeline, so windows that
			// touch or wrap midnight alongside another window come back reshaped.
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "01:00", end = "05:00" },
    { start = "12:00", end = "13:00" },
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "blocked_windows.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_logs_alert.test", "blocked_windows.*", map[string]string{
						"start": "12:00",
						"end":   "13:00",
					}),
				),
			},
			// Drop the attribute entirely.
			{
				Config: testAccLogsAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("posthog_logs_alert.test", "blocked_windows.#"),
					testAccCheckLogsAlertQuietHoursCleared("posthog_logs_alert.test"),
				),
			},
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "01:00", end = "05:00" },
`),
				Check: resource.TestCheckResourceAttr("posthog_logs_alert.test", "blocked_windows.#", "1"),
			},
		},
	})
}

// TestLogsAlert_EnabledToggle covers pausing an alert and bringing it back, the usual
// reason to touch a logs alert after creating it.
func TestLogsAlert_EnabledToggle(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertWithEnabled(rName, true),
				Check:  resource.TestCheckResourceAttr("posthog_logs_alert.test", "enabled", "true"),
			},
			{
				Config: testAccLogsAlertWithEnabled(rName, false),
				Check:  resource.TestCheckResourceAttr("posthog_logs_alert.test", "enabled", "false"),
			},
			{
				Config: testAccLogsAlertWithEnabled(rName, true),
				Check:  resource.TestCheckResourceAttr("posthog_logs_alert.test", "enabled", "true"),
			},
		},
	})
}

// TestLogsAlert_FilterGroupLifecycle adds, edits and removes the JSON escape hatch. The
// replan steps are the point: PostHog annotates saved filters with its own defaults, and
// without the response being projected back onto the declared fields those annotations
// would show up as permanent drift.
func TestLogsAlert_FilterGroupLifecycle(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	withStatusCode := testAccLogsAlertWithFilterGroup(rName, "500")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			// Start without a filter group, so the next step exercises adding one.
			{
				Config: testAccLogsAlertBasic(rName),
				Check:  resource.TestCheckNoResourceAttr("posthog_logs_alert.test", "filter_group_json"),
			},
			{
				Config: withStatusCode,
				Check:  resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "filter_group_json"),
			},
			{
				Config:   withStatusCode,
				PlanOnly: true,
			},
			// Change a value inside the JSON.
			{
				Config: testAccLogsAlertWithFilterGroup(rName, "503"),
				Check: resource.TestMatchResourceAttr(
					"posthog_logs_alert.test", "filter_group_json", regexp.MustCompile(`503`)),
			},
			// Remove it, leaving severity_levels as the only filter.
			{
				Config: testAccLogsAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("posthog_logs_alert.test", "filter_group_json"),
					testAccCheckLogsAlertFilterGroupCleared("posthog_logs_alert.test"),
				),
			},
		},
	})
}

// TestLogsAlert_UpdateName renames in place. name is Optional+Computed, so this also
// confirms a rename is an update rather than a replace.
func TestLogsAlert_UpdateName(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertBasic(rName),
				Check:  resource.TestCheckResourceAttr("posthog_logs_alert.test", "name", rName),
			},
			{
				Config: testAccLogsAlertBasic(rName + "-renamed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "name", rName+"-renamed"),
					resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "id"),
				),
			},
		},
	})
}

// TestLogsAlert_ExternalDeletion deletes the alert outside Terraform, as someone would
// from the PostHog UI. Read must treat the 404 as "gone" and plan a fresh create rather
// than erroring.
func TestLogsAlert_ExternalDeletion(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"acceptance-test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	var alertID string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "id"),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["posthog_logs_alert.test"]
						if !ok {
							return fmt.Errorf("resource not found: posthog_logs_alert.test")
						}
						alertID = rs.Primary.ID
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					if _, err := (&client).DeleteLogsAlert(context.Background(), projectID, alertID); err != nil {
						t.Fatalf("deleting logs alert %s externally: %v", alertID, err)
					}
				},
				Config:             testAccLogsAlertBasic(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccLogsAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "id"),
				),
			},
		},
	})
}

// TestLogsAlert_RejectsShortBlockedWindow verifies the >=30-minute quiet-hours rule is
// caught at plan time rather than mid-apply.
func TestLogsAlert_RejectsShortBlockedWindow(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "02:00", end = "02:15" },
`),
				ExpectError: regexp.MustCompile(`Quiet-hours window is too short`),
			},
		},
	})
}

// TestLogsAlert_RejectsTouchingBlockedWindows verifies that windows which merely touch are
// rejected. PostHog merges on `next.start <= prev.end`, so 01:00-02:00 and 02:00-03:00
// would be saved as a single 01:00-03:00 window and the apply would fail.
func TestLogsAlert_RejectsTouchingBlockedWindows(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "01:00", end = "02:00" },
    { start = "02:00", end = "03:00" },
`),
				ExpectError: regexp.MustCompile(`overlap`),
			},
		},
	})
}

// TestLogsAlert_RejectsWrappedWindowAlongsideAnother verifies the midnight-crossing rule.
// PostHog re-encodes a crossing window as one window only when it is the whole
// configuration; next to another window it is stored as two, so the alert would read back
// with three windows where two were configured.
func TestLogsAlert_RejectsWrappedWindowAlongsideAnother(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "22:00", end = "07:00" },
    { start = "12:00", end = "13:00" },
`),
				ExpectError: regexp.MustCompile(`must be the only window`),
			},
		},
	})
}

// TestLogsAlert_RejectsEmptyFilterGroupJSON verifies jsonencode({}) is rejected at plan
// time. jsontypes.Normalized only checks that the string is well-formed JSON, so without
// the extra validator this reaches the API as an empty filter group.
func TestLogsAlert_RejectsEmptyFilterGroupJSON(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name              = %[1]q
  severity_levels   = ["error"]
  filter_group_json = jsonencode({})
}
`, rName),
				ExpectError: regexp.MustCompile(`must be a non-empty JSON object`),
			},
		},
	})
}

// testAccCheckLogsAlertQuietHoursCleared asserts PostHog holds no blocked windows for the
// alert. Checking state alone would pass even if the PATCH never reached the server.
func testAccCheckLogsAlertQuietHoursCleared(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		alert, err := testAccLogsAlertFromAPI(s, resourceName)
		if err != nil {
			return err
		}
		if alert.ScheduleRestriction != nil && len(alert.ScheduleRestriction.BlockedWindows) > 0 {
			return fmt.Errorf("logs alert %s still has %d blocked window(s) server-side",
				alert.ID, len(alert.ScheduleRestriction.BlockedWindows))
		}
		return nil
	}
}

// testAccCheckLogsAlertFilterGroupCleared asserts removing filter_group_json from the
// config actually cleared it server-side, which is what the whole-object filter replace
// in LogsAlertFilters exists to guarantee.
func testAccCheckLogsAlertFilterGroupCleared(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		alert, err := testAccLogsAlertFromAPI(s, resourceName)
		if err != nil {
			return err
		}
		if alert.Filters != nil && len(alert.Filters.FilterGroup) > 0 {
			return fmt.Errorf("logs alert %s still has a filter group server-side: %v",
				alert.ID, alert.Filters.FilterGroup)
		}
		return nil
	}
}

// testAccLogsAlertFromAPI re-reads the alert behind a state address straight from PostHog.
func testAccLogsAlertFromAPI(s *terraform.State, resourceName string) (httpclient.LogsAlert, error) {
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return httpclient.LogsAlert{}, fmt.Errorf("resource not found: %s", resourceName)
	}

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)

	alert, _, err := client.GetLogsAlert(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID)
	if err != nil {
		return httpclient.LogsAlert{}, fmt.Errorf("reading logs alert %s: %w", rs.Primary.ID, err)
	}
	return alert, nil
}

func testAccLogsAlertWithBlockedWindows(name, windows string) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]

  blocked_windows = [
%[2]s
  ]
}
`, name, windows)
}

func testAccLogsAlertWithEnabled(name string, enabled bool) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
  enabled         = %[2]t
}
`, name, enabled)
}

func testAccLogsAlertWithFilterGroup(name, statusCode string) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]

  filter_group_json = jsonencode({
    type = "AND"
    values = [{
      type = "AND"
      values = [{
        type     = "log_attribute"
        key      = "status_code"
        operator = "exact"
        value    = [%[2]q]
      }]
    }]
  })
}
`, name, statusCode)
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
