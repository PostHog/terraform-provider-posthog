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

// TestAlert_ScheduleRestrictionLifecycle walks quiet hours through add, change, remove and
// re-add. Step 3 carries the weight: removing the block sends an explicit null, and this
// checks PostHog really cleared it.
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
			// Add it back, to show clearing did not leave the alert unable to take one.
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
// offers as a preset. An end before the start is legal, so the provider must not rewrite
// or reject it.
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
			// A second plan must be empty, showing the crossing window round-trips
			// unchanged. Ordering is covered by the five-window replan in
			// BlockedWindowsAcceptedBoundaries, since one window has no order to get
			// wrong.
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

// TestAlert_RejectsInvalidBlockedWindows checks the plan-time rejections the unit table
// cannot reach, plus one representative rule.
//
// The rule matrix lives in TestBlockedWindowsValidator and runs without an instance.
// Repeating it here would cost a full plan per rule for the same result. What only this
// layer shows is that the validators are wired at the right schema path. The two midnight
// rows are kept on purpose, because those rules were wrong once.
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
			wantError: regexp.MustCompile(`Quiet-hours windows overlap`),
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
		"time that does not exist": {
			windows: "\n      { start = \"24:00\", end = \"06:00\" },\n",
			// Terraform wraps diagnostics mid-sentence, so match a fragment that cannot
			// span the break.
			wantError: regexp.MustCompile(`24-hour time in HH:MM format`),
		},
		// Terraform wraps diagnostics, so these patterns stop before the break between the
		// count and "elements".
		// Not a PostHog limit. An empty list becomes a null restriction, so the alert would
		// read back different from the configured block.
		"empty list": {
			windows:   ``,
			wantError: regexp.MustCompile(`set must contain at least 1`),
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

// TestAlert_BlockedWindowsAcceptedBoundaries applies the shapes the validator permits but
// that were only ever checked against PostHog's normalizer. If either call is wrong, every
// apply fails instead of the plan.
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
	// The shortest window PostHog accepts. Only a live apply shows whether its floor is
	// really 30 minutes.
	exactlyThirtyMinutes := `
      { start = "02:00", end = "02:30" },
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
			{
				Config: testAccAlertWithBlockedWindows(rName, exactlyThirtyMinutes),
				Check:  resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
			},
			{
				Config:   testAccAlertWithBlockedWindows(rName, exactlyThirtyMinutes),
				PlanOnly: true,
			},
		},
	})
}

// TestAlert_UnrelatedUpdateKeepsQuietHours changes an unrelated attribute while quiet
// hours are set. Update re-sends the whole request, so a regression would clear the windows
// on an edit that has nothing to do with them.
func TestAlert_UnrelatedUpdateKeepsQuietHours(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	withWindows := func(upper int) string {
		return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_upper  = %d
  condition_type   = "absolute_value"
  series_index     = 0

  schedule_restriction = {
    blocked_windows = [
      { start = "02:00", end = "04:00" },
    ]
  }

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(rName), rName, upper)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: withWindows(100),
				Check:  resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
			},
			{
				Config: withWindows(250),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "250"),
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
						"start": "02:00",
						"end":   "04:00",
					}),
				),
			},
		},
	})
}

// TestAlert_ScheduleRestrictionDrift edits quiet hours outside Terraform, as someone would
// from the PostHog UI, and checks the change shows up as a plan. An empty window list maps
// to null, so a clear on the server must read as drift rather than agreement.
func TestAlert_ScheduleRestrictionDrift(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"acceptance-test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	config := testAccAlertWithBlockedWindows(rName, `
      { start = "02:00", end = "04:00" },
`)

	var alertID string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["posthog_alert.test"]
						if !ok {
							return fmt.Errorf("resource not found: posthog_alert.test")
						}
						alertID = rs.Primary.ID
						return nil
					},
				),
			},
			// Clear quiet hours behind Terraform's back.
			{
				PreConfig: func() {
					alert, _, err := client.GetAlert(context.Background(), projectID, alertID)
					if err != nil {
						t.Fatalf("reading alert %s: %v", alertID, err)
					}
					if _, _, err := client.UpdateAlert(context.Background(), projectID, alertID, httpclient.AlertRequest{
						Insight:         alert.Insight.ID,
						Threshold:       alert.Threshold,
						Condition:       alert.Condition,
						Config:          alert.Config,
						SubscribedUsers: []int64{},
						// Explicit null is how quiet hours are cleared.
						ScheduleRestriction: nil,
					}); err != nil {
						t.Fatalf("clearing quiet hours on %s: %v", alertID, err)
					}
				},
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// Applying restores what the configuration asks for.
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
						"start": "02:00",
						"end":   "04:00",
					}),
				),
			},
			// Move the window rather than clearing it. This is the other drift shape: the
			// server holds windows, just different ones, so it exercises the populated
			// branch of the response mapper instead of the null one.
			{
				PreConfig: func() {
					alert, _, err := client.GetAlert(context.Background(), projectID, alertID)
					if err != nil {
						t.Fatalf("reading alert %s: %v", alertID, err)
					}
					if _, _, err := client.UpdateAlert(context.Background(), projectID, alertID, httpclient.AlertRequest{
						Insight:         alert.Insight.ID,
						Threshold:       alert.Threshold,
						Condition:       alert.Condition,
						Config:          alert.Config,
						SubscribedUsers: []int64{},
						ScheduleRestriction: &httpclient.AlertScheduleRestriction{
							BlockedWindows: []httpclient.AlertBlockedWindow{{Start: "03:00", End: "05:00"}},
						},
					}); err != nil {
						t.Fatalf("moving quiet hours on %s: %v", alertID, err)
					}
				},
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: config,
				Check: resource.TestCheckTypeSetElemNestedAttrs("posthog_alert.test", "schedule_restriction.blocked_windows.*", map[string]string{
					"start": "02:00",
					"end":   "04:00",
				}),
			},
			// Drop quiet hours from the configuration, then have PostHog gain some out of
			// band. This is the shape the CHANGELOG upgrade note describes: a config with
			// no schedule_restriction against a server that has one. It is the only drift
			// direction where the state field goes from null to populated on refresh.
			{
				Config: testAccAlertBasic(rName),
				Check:  resource.TestCheckNoResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#"),
			},
			{
				PreConfig: func() {
					alert, _, err := client.GetAlert(context.Background(), projectID, alertID)
					if err != nil {
						t.Fatalf("reading alert %s: %v", alertID, err)
					}
					if _, _, err := client.UpdateAlert(context.Background(), projectID, alertID, httpclient.AlertRequest{
						Insight:         alert.Insight.ID,
						Threshold:       alert.Threshold,
						Condition:       alert.Condition,
						Config:          alert.Config,
						SubscribedUsers: []int64{},
						ScheduleRestriction: &httpclient.AlertScheduleRestriction{
							BlockedWindows: []httpclient.AlertBlockedWindow{{Start: "08:00", End: "09:00"}},
						},
					}); err != nil {
						t.Fatalf("adding quiet hours to %s out of band: %v", alertID, err)
					}
				},
				Config:             testAccAlertBasic(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// Applying removes them again, which is what the upgrade note promises.
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("posthog_alert.test", "schedule_restriction.blocked_windows.#"),
					testAccCheckAlertQuietHoursCleared("posthog_alert.test"),
				),
			},
		},
	})
}

// TestAlert_ServerEnforcedWindowLimits pins the rules the provider deliberately does NOT
// duplicate. Window length and count are PostHog's constants, so repeating them here would
// mean a provider release whenever PostHog changes one. These configs must therefore reach
// the API and be rejected by it. If PostHog ever stops rejecting them, this fails and tells
// us the delegation assumption has moved.
//
// Whole-day coverage is deliberately absent: covering 1440 minutes requires windows that
// touch, so the reshape rules reject it at plan time and it never reaches the API.
func TestAlert_ServerEnforcedWindowLimits(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	tests := map[string]string{
		"window shorter than PostHog allows": `
      { start = "02:00", end = "02:15" },
`,
		"more windows than PostHog stores": `
      { start = "00:00", end = "01:00" },
      { start = "02:00", end = "03:00" },
      { start = "04:00", end = "05:00" },
      { start = "06:00", end = "07:00" },
      { start = "08:00", end = "09:00" },
      { start = "10:00", end = "11:00" },
`,
	}

	for name, windows := range tests {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckAlertDestroy,
				Steps: []resource.TestStep{
					{
						Config: testAccAlertWithBlockedWindows(rName, windows),
						// This endpoint flattens every schedule error to one message, which
						// is the cost of delegating rather than checking at plan time.
						// Terraform hard-wraps the body, so match a fragment that cannot
						// straddle the break.
						ExpectError: regexp.MustCompile(`Invalid schedule`),
					},
				},
			})
		})
	}
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
