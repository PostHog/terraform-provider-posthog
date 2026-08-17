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

// testAccCheckAlertDestroy verifies the alert has been destroyed.
func testAccCheckAlertDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_alert" {
			continue
		}

		_, status, err := client.GetAlert(context.Background(), projectID, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("alert %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

// TestAlert_Basic tests creating an alert with minimal configuration.
func TestAlert_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_type", "absolute"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "insight"),
				),
			},
		},
	})
}

// TestAlert_AllFields tests creating an alert with all optional fields.
func TestAlert_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "true"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_type", "absolute"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "10"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "100"),
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "absolute_value"),
					resource.TestCheckResourceAttr("posthog_alert.test", "check_ongoing_interval", "true"),
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "daily"),
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "id"),
				),
			},
		},
	})
}

// TestAlert_Update tests updating an alert.
func TestAlert_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccAlertWithThreshold(rName, 10, 100),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "10"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "100"),
				),
			},
			// Update thresholds
			{
				Config: testAccAlertWithThreshold(rName, 20, 200),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "20"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "200"),
				),
			},
			// Update name
			{
				Config: testAccAlertWithThreshold(rName+"-updated", 20, 200),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "20"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "200"),
				),
			},
		},
	})
}

// TestAlert_EnableDisable tests enabling and disabling an alert.
func TestAlert_EnableDisable(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccAlertWithEnabled(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "true"),
				),
			},
			// Disable
			{
				Config: testAccAlertWithEnabled(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "false"),
				),
			},
			// Re-enable
			{
				Config: testAccAlertWithEnabled(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "true"),
				),
			},
		},
	})
}

// TestAlert_ConditionTypes tests different condition types.
func TestAlert_ConditionTypes(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Absolute value
			{
				Config: testAccAlertWithCondition(rName, "absolute_value"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "absolute_value"),
				),
			},
			// Relative increase
			{
				Config: testAccAlertWithCondition(rName, "relative_increase"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "relative_increase"),
				),
			},
			// Relative decrease
			{
				Config: testAccAlertWithCondition(rName, "relative_decrease"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "relative_decrease"),
				),
			},
		},
	})
}

// TestAlert_CalculationIntervals tests different calculation intervals.
func TestAlert_CalculationIntervals(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertWithInterval(rName, "hourly"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "hourly"),
				),
			},
			{
				Config: testAccAlertWithInterval(rName, "daily"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "daily"),
				),
			},
			{
				Config: testAccAlertWithInterval(rName, "weekly"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "weekly"),
				),
			},
		},
	})
}

// TestAlert_CheckOngoingInterval tests toggling the check_ongoing_interval setting.
func TestAlert_CheckOngoingInterval(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create with check_ongoing_interval = true
			{
				Config: testAccAlertWithCheckOngoingInterval(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "check_ongoing_interval", "true"),
				),
			},
			// Update to false
			{
				Config: testAccAlertWithCheckOngoingInterval(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "check_ongoing_interval", "false"),
				),
			},
			// Update back to true
			{
				Config: testAccAlertWithCheckOngoingInterval(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "check_ongoing_interval", "true"),
				),
			},
		},
	})
}

// TestAlert_Import tests importing an existing alert by ID.
func TestAlert_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
				),
			},
			// Import
			{
				ResourceName:            "posthog_alert.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subscribed_users"},
			},
		},
	})
}

// TestAlert_InsightDeletion tests the behavior when removing an alert that references an insight.
// This verifies that:
// 1. Alert can be removed while insight remains
// 2. Insight can be deleted after alert is removed
func TestAlert_InsightDeletion(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	insightOnlyConfig := fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name = "%s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        name  = "$pageview"
        event = "$pageview"
        math  = "total"
      }]
    }
  })
}
`, rName)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create insight and alert
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "insight"),
				),
			},
			// Step 2: Remove alert, keep insight - alert should be deleted
			{
				Config: insightOnlyConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName+"-insight"),
				),
			},
			// Step 3: Remove insight too - should succeed since alert is gone
			{
				Config: `provider "posthog" {}`,
			},
			// Step 4: Recreate
			{
				Config: testAccAlertBasic(rName + "-recreated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName+"-recreated"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "insight"),
				),
			},
		},
	})
}

// TestAlert_RejectsInvalidConditionType verifies that the OneOf validator on
// `condition_type` rejects unknown values at plan time, before any API call.
func TestAlert_RejectsInvalidConditionType(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAlertWithCondition(rName, "definitely_not_a_condition"),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`(?s)condition_type.*value must be one of`),
			},
		},
	})
}

// TestAlert_RejectsNegativeSeriesIndex verifies that the AtLeast(0) validator
// on `series_index` rejects negative values at plan time, before any API call.
func TestAlert_RejectsNegativeSeriesIndex(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAlertWithSeriesIndex(rName, -1),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`(?s)series_index.*value must be at least 0`),
			},
		},
	})
}

// TestAlert_ScheduleRestrictionLifecycle walks quiet hours through the full CRUD cycle:
// add windows, change them, drop the block, then add it back. Step 3 is the one that
// matters most - `schedule_restriction` is serialized without `omitempty` so that removing
// the block sends an explicit null, and this asserts PostHog actually cleared it.
func TestAlert_ScheduleRestrictionLifecycle(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create with a single window.
			{
				Config: testAccAlertWithBlockedWindows(rName, `
      { start = "22:00", end = "23:59" },
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
						"start": "22:00",
						"end":   "23:59",
					}),
				),
			},
			// Add a second window and move the first.
			{
				Config: testAccAlertWithBlockedWindows(rName, `
      { start = "21:00", end = "23:59" },
      { start = "00:00", end = "06:00" },
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
						"start": "21:00",
						"end":   "23:59",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
						"start": "00:00",
						"end":   "06:00",
					}),
				),
			},
			// Remove the block. Quiet hours must be gone server-side, not just in state.
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#"),
					testAccCheckAlertQuietHoursCleared("posthog_alert.test"),
				),
			},
			// Add it back, to prove clearing did not leave the alert in a state that
			// rejects a later restriction.
			{
				Config: testAccAlertWithBlockedWindows(rName, `
      { start = "01:00", end = "05:00" },
`),
				Check: resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
			},
		},
	})
}

// TestAlert_ScheduleRestrictionWrapsMidnight covers the overnight window the PostHog UI
// offers as a preset. The half-open `[start, end)` reading means end < start is legal, so
// this asserts the provider does not normalize or reject it.
func TestAlert_ScheduleRestrictionWrapsMidnight(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertWithBlockedWindows(rName, `
      { start = "22:00", end = "07:00" },
`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
						"start": "22:00",
						"end":   "07:00",
					}),
				),
			},
			// A second plan against the same config must be empty: PostHog returns the
			// windows in its own order, and a list would report a spurious diff here.
			{
				Config: testAccAlertWithBlockedWindows(rName, `
      { start = "22:00", end = "07:00" },
`),
				PlanOnly: true,
			},
		},
	})
}

// TestAlert_ScheduleRestrictionImport verifies an alert with quiet hours round-trips
// through import, so an alert created in the PostHog UI can be adopted into Terraform.
func TestAlert_ScheduleRestrictionImport(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	config := testAccAlertWithBlockedWindows(rName, `
      { start = "02:00", end = "04:00" },
`)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.TestCheckResourceAttr(
					"posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
			},
			{
				Config:                  config,
				ResourceName:            "posthog_alert.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subscribed_users"},
			},
		},
	})
}

// TestAlert_RejectsInvalidBlockedWindows checks every quiet-hours rule rejects at plan
// time, before any API call. The rules themselves are covered by the unit table; what this
// adds is that the validator is wired at the right schema path and that each rule's
// user-facing message is the one a practitioner actually sees.
func TestAlert_RejectsInvalidBlockedWindows(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	tests := map[string]struct {
		windows   string
		wantError *regexp.Regexp
	}{
		"overlapping": {
			windows: `
      { start = "01:00", end = "03:00" },
      { start = "02:00", end = "04:00" },
`,
			wantError: regexp.MustCompile(`Overlapping blocked windows`),
		},
		"touching": {
			windows: `
      { start = "00:00", end = "06:00" },
      { start = "06:00", end = "09:00" },
`,
			wantError: regexp.MustCompile(`Overlapping blocked windows`),
		},
		"crossing midnight alongside another": {
			windows: `
      { start = "22:00", end = "07:00" },
      { start = "12:00", end = "13:00" },
`,
			wantError: regexp.MustCompile(`must be the only window`),
		},
		"meeting at midnight": {
			windows: `
      { start = "00:00", end = "06:00" },
      { start = "22:00", end = "00:00" },
`,
			wantError: regexp.MustCompile(`meeting at midnight`),
		},
		"shorter than thirty minutes": {
			windows:   "\n      { start = \"02:00\", end = \"02:15\" },\n",
			wantError: regexp.MustCompile(`Blocked window is too short`),
		},
		"time that does not exist": {
			windows: "\n      { start = \"24:00\", end = \"06:00\" },\n",
			// Terraform hard-wraps diagnostics and the break lands mid-sentence, so match
			// only the fragment that cannot straddle it.
			wantError: regexp.MustCompile(`24-hour time in HH:MM format`),
		},
		// Terraform hard-wraps diagnostics, so these patterns stop before the line break
		// that falls between the count and "elements".
		"empty list": {
			windows:   ``,
			wantError: regexp.MustCompile(`set must contain at least 1`),
		},
		"more than five": {
			windows: `
      { start = "00:00", end = "01:00" },
      { start = "02:00", end = "03:00" },
      { start = "04:00", end = "05:00" },
      { start = "06:00", end = "07:00" },
      { start = "08:00", end = "09:00" },
      { start = "10:00", end = "11:00" },
`,
			wantError: regexp.MustCompile(`set must contain at most 5`),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      testAccAlertWithBlockedWindows(rName, test.windows),
						PlanOnly:    true,
						ExpectError: test.wantError,
					},
				},
			})
		})
	}
}

// TestAlert_BlockedWindowsAcceptedBoundaries applies the shapes the validator deliberately
// permits but that were only ever checked against PostHog's normalizer, not the live API:
// the maximum five windows, and a window ending exactly at midnight next to another window.
// Both are places where a wrong permissive call fails every apply with an inconsistent
// result rather than a plan error.
func TestAlert_BlockedWindowsAcceptedBoundaries(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	fiveWindows := `
      { start = "00:00", end = "01:00" },
      { start = "02:00", end = "03:00" },
      { start = "04:00", end = "05:00" },
      { start = "06:00", end = "07:00" },
      { start = "08:00", end = "09:00" },
`
	endsAtMidnight := `
      { start = "19:00", end = "00:00" },
      { start = "12:00", end = "13:00" },
`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertWithBlockedWindows(rName, fiveWindows),
				Check:  resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "5"),
			},
			{
				Config:   testAccAlertWithBlockedWindows(rName, fiveWindows),
				PlanOnly: true,
			},
			{
				Config: testAccAlertWithBlockedWindows(rName, endsAtMidnight),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
						"start": "19:00",
						"end":   "00:00",
					}),
				),
			},
			{
				Config:   testAccAlertWithBlockedWindows(rName, endsAtMidnight),
				PlanOnly: true,
			},
		},
	})
}

// testAccCheckAlertQuietHoursCleared asserts PostHog itself holds no blocked windows for
// the alert. Checking only that Terraform dropped them from state would pass even if the
// PATCH never reached the server.
func testAccCheckAlertQuietHoursCleared(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		client := httpclient.NewDefaultClient(
			os.Getenv("POSTHOG_HOST"),
			os.Getenv("POSTHOG_API_KEY"),
			"test",
		)

		alert, _, err := client.GetAlert(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("reading alert %s: %w", rs.Primary.ID, err)
		}
		if alert.ScheduleRestriction != nil && len(alert.ScheduleRestriction.BlockedWindows) > 0 {
			return fmt.Errorf("alert %s still has %d blocked window(s) server-side",
				rs.Primary.ID, len(alert.ScheduleRestriction.BlockedWindows))
		}

		return nil
	}
}

// testAccAlertWithBlockedWindows builds an alert whose quiet hours are the given HCL
// object entries, so a test can vary the windows without restating the whole resource.
func testAccAlertWithBlockedWindows(name, windows string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_upper  = 100
  condition_type   = "absolute_value"
  series_index     = 0

  schedule_restriction = {
    blocked_windows = [
%s
    ]
  }

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, windows)
}

// Helper function to create the base insight that alerts monitor
func testAccAlertInsightBase(name string) string {
	return fmt.Sprintf(`
resource "posthog_insight" "test" {
  name = "%s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        name  = "$pageview"
        event = "$pageview"
        math  = "total"
      }]
    }
  })
}
`, name)
}

func testAccAlertBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_upper  = 100
  condition_type   = "absolute_value"
  series_index     = 0

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name)
}

func testAccAlertAllFields(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name                   = %q
  insight                = posthog_insight.test.id
  subscribed_users       = []
  enabled                = true
  threshold_type         = "absolute"
  threshold_lower        = 10
  threshold_upper        = 100
  condition_type         = "absolute_value"
  series_index           = 0
  check_ongoing_interval = true
  calculation_interval   = "daily"
  skip_weekend           = false

  schedule_restriction = {
    blocked_windows = [
      { start = "22:00", end = "07:00" },
    ]
  }

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name)
}

func testAccAlertWithThreshold(name string, lower, upper int) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_lower  = %d
  threshold_upper  = %d
  condition_type   = "absolute_value"
  series_index     = 0

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, lower, upper)
}

func testAccAlertWithEnabled(name string, enabled bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  enabled          = %t
  threshold_type   = "absolute"
  threshold_upper  = 100
  condition_type   = "absolute_value"
  series_index     = 0

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, enabled)
}

func testAccAlertWithCondition(name, conditionType string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_upper  = 100
  condition_type   = %q
  series_index     = 0

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, conditionType)
}

func testAccAlertWithInterval(name, interval string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name                 = %q
  insight              = posthog_insight.test.id
  subscribed_users     = []
  threshold_type       = "absolute"
  threshold_upper      = 100
  calculation_interval = %q
  condition_type       = "absolute_value"
  series_index         = 0

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, interval)
}

func testAccAlertWithCheckOngoingInterval(name string, checkOngoing bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name                   = %q
  insight                = posthog_insight.test.id
  subscribed_users       = []
  threshold_type         = "absolute"
  threshold_upper        = 100
  check_ongoing_interval = %t
  condition_type         = "absolute_value"
  series_index           = 0

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, checkOngoing)
}

func testAccAlertWithSeriesIndex(name string, seriesIndex int) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_upper  = 100
  condition_type   = "absolute_value"
  series_index     = %d

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, seriesIndex)
}
