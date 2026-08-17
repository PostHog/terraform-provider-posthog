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

// skipIfNoLogsAlerting skips when the logs alerts API is not available to this
// organization. PostHog gates the endpoint behind the `logs-alerting` feature flag and
// answers 403 until it is enabled, which would otherwise surface as every test in this
// file failing on a raw API error rather than skipping.
func skipIfNoLogsAlerting(t *testing.T) {
	t.Helper()

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)

	_, status, err := client.GetLogsAlert(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), "capability-probe")
	if err != nil && status == httpclient.HTTPStatusCode(http.StatusForbidden) {
		t.Skip("Skipping test: the logs-alerting feature flag is not enabled for this organization")
	}
}

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
	skipIfNoLogsAlerting(t)

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
	skipIfNoLogsAlerting(t)

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
	skipIfNoLogsAlerting(t)

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
	skipIfNoLogsAlerting(t)

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
	skipIfNoLogsAlerting(t)

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

// TestLogsAlert_ClearsFilter removes a previously-set filter and asserts it is actually
// gone server-side — the whole-object replace semantics the client relies on.
func TestLogsAlert_ClearsFilter(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("posthog_logs_alert.test", "service_names.#"),
					// State alone would pass even if the PATCH never reached the server.
					// service_names goes through the same whole-object Filters replace as
					// the filter group, so it needs the same server-side assertion.
					testAccCheckLogsAlertServiceNamesCleared("posthog_logs_alert.test"),
				),
			},
		},
	})
}

// TestLogsAlert_ClearsFilterWithEmptyList covers the other removal form the schema
// documents: setting the attribute to an empty list rather than omitting it. The provider
// preserves an explicitly empty set instead of flipping it to null, so this must stay a
// clean no-op plan rather than drifting back to the configured value.
func TestLogsAlert_ClearsFilterWithEmptyList(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	emptyLists := fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
  service_names   = []
  blocked_windows = []
}
`, rName)

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

  blocked_windows = [{ start = "22:00", end = "06:00" }]
}
`, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "service_names.#", "1"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "blocked_windows.#", "1"),
				),
			},
			{
				Config: emptyLists,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "service_names.#", "0"),
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "blocked_windows.#", "0"),
					testAccCheckLogsAlertServiceNamesCleared("posthog_logs_alert.test"),
					testAccCheckLogsAlertQuietHoursCleared("posthog_logs_alert.test"),
				),
			},
		},
	})
}

// TestLogsAlert_ImportComposite covers the `project_id/uuid` import form documented in
// examples/resources/posthog_logs_alert/import.sh, alongside an alert that carries a
// filter group. Import adopts the API's filter JSON verbatim while apply projects it onto
// the declared fields, so filter_group_json is verified separately rather than compared.
func TestLogsAlert_ImportComposite(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	config := testAccLogsAlertWithFilterGroup(rName, "500")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  resource.TestCheckResourceAttrSet("posthog_logs_alert.test", "filter_group_json"),
			},
			{
				Config:       config,
				ResourceName: "posthog_logs_alert.test",
				ImportState:  true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["posthog_logs_alert.test"]
					if !ok {
						return "", fmt.Errorf("resource not found: posthog_logs_alert.test")
					}
					return fmt.Sprintf("%s/%s", getProjectID(), rs.Primary.ID), nil
				},
				ImportStateVerify: true,
				// Import adopts PostHog's stored filter group verbatim while apply
				// projects it onto the declared fields, so the two legitimately differ.
				// Assert the imported value is present rather than comparing it.
				ImportStateVerifyIgnore: []string{"filter_group_json"},
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) != 1 {
						return fmt.Errorf("expected 1 imported state, got %d", len(states))
					}
					if got := states[0].Attributes["filter_group_json"]; got == "" {
						return fmt.Errorf("imported alert has no filter_group_json")
					}
					return nil
				},
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
	skipIfNoLogsAlerting(t)

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

// TestLogsAlert_BlockedWindowsLifecycle walks quiet hours through add, change, remove and
// re-add. The removal step is the one with teeth: schedule_restriction is sent without
// omitempty so an absent set clears it, and this asserts PostHog agrees.
func TestLogsAlert_BlockedWindowsLifecycle(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

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
			// Replan with two windows configured, which is where ordering can actually
			// vary: PostHog returns them sorted by start time regardless of config order,
			// so a list would show a spurious diff here where a set does not.
			{
				Config: testAccLogsAlertWithBlockedWindows(rName, `
    { start = "12:00", end = "13:00" },
    { start = "01:00", end = "05:00" },
`),
				PlanOnly: true,
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
	skipIfNoLogsAlerting(t)

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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "enabled", "false"),
					// The schema documents that disabling resets the alert state.
					resource.TestCheckResourceAttr("posthog_logs_alert.test", "state", "not_firing"),
				),
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
	skipIfNoLogsAlerting(t)

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
	skipIfNoLogsAlerting(t)

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
	skipIfNoLogsAlerting(t)

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

// TestLogsAlert_RejectsInvalidConfigs checks every plan-time rule rejects before any API
// call. The rules themselves are covered exhaustively by the unit tables; what this adds
// is that ModifyResourcePlan and the schema validators are actually wired into the
// registered resource, and that each rule's user-facing message is the one a practitioner
// sees. One table rather than one function per rule keeps that visible on a single screen.
func TestLogsAlert_RejectsInvalidConfigs(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	tests := map[string]struct {
		body      string
		wantError *regexp.Regexp
	}{
		"datapoints exceed evaluation periods": {
			body:      "evaluation_periods = 2\n  datapoints_to_alarm = 5",
			wantError: regexp.MustCompile(`Alert can never fire`),
		},
		"below zero can never be satisfied": {
			body:      `threshold_operator = "below"` + "\n  threshold_count = 0",
			wantError: regexp.MustCompile(`Alert can never fire`),
		},
		"invalid threshold operator": {
			body:      `threshold_operator = "sideways"`,
			wantError: regexp.MustCompile(`Attribute threshold_operator value must be one of`),
		},
		"empty filter group json": {
			body:      "filter_group_json = jsonencode({})",
			wantError: regexp.MustCompile(`must be a non-empty JSON object`),
		},
		"overlapping windows": {
			body:      `blocked_windows = [{ start = "01:00", end = "03:00" }, { start = "02:00", end = "04:00" }]`,
			wantError: regexp.MustCompile(`overlap`),
		},
		"touching windows": {
			body:      `blocked_windows = [{ start = "01:00", end = "02:00" }, { start = "02:00", end = "03:00" }]`,
			wantError: regexp.MustCompile(`overlap`),
		},
		"window shorter than thirty minutes": {
			body:      `blocked_windows = [{ start = "02:00", end = "02:15" }]`,
			wantError: regexp.MustCompile(`Quiet-hours window is too short`),
		},
		"crossing window alongside another": {
			body:      `blocked_windows = [{ start = "22:00", end = "07:00" }, { start = "12:00", end = "13:00" }]`,
			wantError: regexp.MustCompile(`must be the only window`),
		},
		"windows meeting at midnight": {
			body:      `blocked_windows = [{ start = "00:00", end = "06:00" }, { start = "22:00", end = "00:00" }]`,
			wantError: regexp.MustCompile(`meeting at midnight`),
		},
		"malformed window time": {
			body: `blocked_windows = [{ start = "24:00", end = "06:00" }]`,
			// Terraform hard-wraps diagnostics and the break lands mid-sentence, so match
			// only the fragment that cannot straddle it.
			wantError: regexp.MustCompile(`24-hour time in HH:MM format`),
		},
		"more than five windows": {
			body: `blocked_windows = [
    { start = "00:00", end = "01:00" }, { start = "02:00", end = "03:00" },
    { start = "04:00", end = "05:00" }, { start = "06:00", end = "07:00" },
    { start = "08:00", end = "09:00" }, { start = "10:00", end = "11:00" },
  ]`,
			// Terraform hard-wraps diagnostics, so the pattern stops before the line
			// break that falls between the count and "elements".
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
						Config: fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
  %[2]s
}
`, rName, test.body),
						ExpectError: test.wantError,
					},
				},
			})
		})
	}
}

// TestLogsAlert_RejectsEnabledAlertWithoutFilters is separate from the table above because
// it is the one rule whose config must omit severity_levels, which the table always sets.
func TestLogsAlert_RejectsEnabledAlertWithoutFilters(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf("resource \"posthog_logs_alert\" \"test\" {\n  name = %[1]q\n}\n", rName),
				ExpectError: regexp.MustCompile(`no filters`),
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

// testAccCheckLogsAlertServiceNamesCleared asserts the service filter is gone server-side.
// It goes through the same whole-object Filters replace as the filter group, so a
// state-only check would pass even if the PATCH never reached PostHog.
func testAccCheckLogsAlertServiceNamesCleared(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		alert, err := testAccLogsAlertFromAPI(s, resourceName)
		if err != nil {
			return err
		}
		if alert.Filters != nil && len(alert.Filters.ServiceNames) > 0 {
			return fmt.Errorf("logs alert %s still has service names server-side: %v",
				alert.ID, alert.Filters.ServiceNames)
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
