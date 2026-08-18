package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

const logsAlertDestinationAddress = "posthog_logs_alert_destination.test"

// skipIfNoSlackWorkspace skips when no Slack integration is available. A Slack destination
// needs a workspace connected to the project through PostHog's OAuth flow, which an
// acceptance test cannot create, so the id has to be supplied.
func skipIfNoSlackWorkspace(t *testing.T) {
	t.Helper()

	if os.Getenv("POSTHOG_TEST_SLACK_WORKSPACE_ID") == "" {
		t.Skip("Skipping test: POSTHOG_TEST_SLACK_WORKSPACE_ID not set")
	}
}

func getSlackWorkspaceID() string {
	return os.Getenv("POSTHOG_TEST_SLACK_WORKSPACE_ID")
}

func getSlackChannelID() string {
	if channelID := os.Getenv("POSTHOG_TEST_SLACK_CHANNEL_ID"); channelID != "" {
		return channelID
	}
	return "C0123456789"
}

// testAccCheckLogsAlertDestinationDestroy verifies the destination is gone from its alert.
// The alert is usually destroyed in the same run, which answers 404 and counts as destroyed.
func testAccCheckLogsAlertDestinationDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_logs_alert_destination" {
			continue
		}

		destinations, status, err := client.ListLogsAlertDestinations(context.Background(), projectID, rs.Primary.Attributes["alert_id"])
		if err != nil {
			if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
				continue
			}
			return fmt.Errorf("listing destinations for alert %s: %w", rs.Primary.Attributes["alert_id"], err)
		}

		stateIDs := strings.Split(rs.Primary.ID, ",")
		for _, destination := range destinations {
			for _, id := range destination.HogFunctionIDs {
				if slices.Contains(stateIDs, id) {
					return fmt.Errorf("logs alert destination %s still exists", rs.Primary.ID)
				}
			}
		}
	}

	return nil
}

// testAccCheckLogsAlertDestinationExists asserts PostHog holds the destination Terraform
// recorded. State alone would pass even if the POST never reached the server.
func testAccCheckLogsAlertDestinationExists(resourceName string) resource.TestCheckFunc {
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

		alertID := rs.Primary.Attributes["alert_id"]
		destinations, _, err := client.ListLogsAlertDestinations(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), alertID)
		if err != nil {
			return fmt.Errorf("listing destinations for alert %s: %w", alertID, err)
		}

		for _, destination := range destinations {
			// The provider builds the id by sorting the group's hog function ids and joining
			// them, so rebuilding it here is what proves the two describe the same group.
			sorted := slices.Sorted(slices.Values(destination.HogFunctionIDs))
			if strings.Join(sorted, ",") == rs.Primary.ID {
				return nil
			}
		}
		return fmt.Errorf("logs alert destination %s not found on alert %s", rs.Primary.ID, alertID)
	}
}

// TestLogsAlertDestination_Webhook covers the create, read and delete cycle for the type
// that needs no external integration.
func TestLogsAlertDestination_Webhook(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertDestinationWebhook(rName, "https://example.com/hooks/first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "type", "webhook"),
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "webhook_url", "https://example.com/hooks/first"),
					resource.TestCheckResourceAttrSet(logsAlertDestinationAddress, "id"),
					// PostHog builds one hog function per alert transition, so a destination
					// is always a group rather than a single function.
					resource.TestCheckResourceAttrSet(logsAlertDestinationAddress, "hog_function_ids.#"),
					resource.TestCheckResourceAttrPair(logsAlertDestinationAddress, "alert_id", "posthog_logs_alert.test", "id"),
					testAccCheckLogsAlertDestinationExists(logsAlertDestinationAddress),
				),
			},
			// Nothing changed, so nothing should be planned. A read that mishandles the
			// group would show drift here.
			{
				Config:   testAccLogsAlertDestinationWebhook(rName, "https://example.com/hooks/first"),
				PlanOnly: true,
			},
		},
	})
}

// TestLogsAlertDestination_Teams covers the third type, which shares the webhook shape but
// is a distinct type server-side.
func TestLogsAlertDestination_Teams(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogsAlertDestinationTeams(rName, "https://outlook.office.com/webhook/first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "type", "teams"),
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "webhook_url", "https://outlook.office.com/webhook/first"),
					testAccCheckLogsAlertDestinationExists(logsAlertDestinationAddress),
				),
			},
			{
				Config:   testAccLogsAlertDestinationTeams(rName, "https://outlook.office.com/webhook/first"),
				PlanOnly: true,
			},
		},
	})
}

// TestLogsAlertDestination_Slack needs a Slack workspace connected to the project, which
// only an operator can set up, so it is skipped unless one is supplied.
func TestLogsAlertDestination_Slack(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)
	skipIfNoSlackWorkspace(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	config := testAccLogsAlertDestinationSlack(rName, getSlackChannelID(), "#alerts")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "type", "slack"),
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "slack_workspace_id", getSlackWorkspaceID()),
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "slack_channel_id", getSlackChannelID()),
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "slack_channel_name", "#alerts"),
					resource.TestCheckNoResourceAttr(logsAlertDestinationAddress, "webhook_url"),
					testAccCheckLogsAlertDestinationExists(logsAlertDestinationAddress),
				),
			},
			// slack_channel_name is write-only: PostHog never stores it, so a read cannot
			// return it. If the response mapper touched the attribute this plan would not be
			// empty, and would stay non-empty on every subsequent plan.
			{
				Config:   config,
				PlanOnly: true,
			},
		},
	})
}

// TestLogsAlertDestination_ReplaceOnChange pins the consequence of there being no update
// endpoint: changing any attribute destroys and recreates the destination, which gives it a
// new set of hog functions and so a new id.
func TestLogsAlertDestination_ReplaceOnChange(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	tests := map[string]struct {
		before string
		after  string
	}{
		"webhook url": {
			before: testAccLogsAlertDestinationWebhook(rName, "https://example.com/hooks/first"),
			after:  testAccLogsAlertDestinationWebhook(rName, "https://example.com/hooks/second"),
		},
		"teams url": {
			before: testAccLogsAlertDestinationTeams(rName, "https://outlook.office.com/webhook/first"),
			after:  testAccLogsAlertDestinationTeams(rName, "https://outlook.office.com/webhook/second"),
		},
		// Switching type also swaps which attributes apply, so this is the case most likely
		// to expose a stale attribute surviving the replace.
		"webhook to teams": {
			before: testAccLogsAlertDestinationWebhook(rName, "https://example.com/hooks/first"),
			after:  testAccLogsAlertDestinationTeams(rName, "https://outlook.office.com/webhook/first"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var firstID string

			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				CheckDestroy:             testAccCheckLogsAlertDestinationDestroy,
				Steps: []resource.TestStep{
					{
						Config: test.before,
						Check: func(s *terraform.State) error {
							rs, ok := s.RootModule().Resources[logsAlertDestinationAddress]
							if !ok {
								return fmt.Errorf("resource not found: %s", logsAlertDestinationAddress)
							}
							firstID = rs.Primary.ID
							return nil
						},
					},
					{
						Config: test.after,
						ConfigPlanChecks: resource.ConfigPlanChecks{
							PreApply: []plancheck.PlanCheck{
								plancheck.ExpectResourceAction(logsAlertDestinationAddress, plancheck.ResourceActionReplace),
							},
						},
						Check: resource.ComposeAggregateTestCheckFunc(
							testAccCheckLogsAlertDestinationExists(logsAlertDestinationAddress),
							func(s *terraform.State) error {
								rs, ok := s.RootModule().Resources[logsAlertDestinationAddress]
								if !ok {
									return fmt.Errorf("resource not found: %s", logsAlertDestinationAddress)
								}
								if rs.Primary.ID == firstID {
									return fmt.Errorf("destination kept id %s across a replace, so the old hog functions were reused", firstID)
								}
								return nil
							},
						),
					},
				},
			})
		})
	}
}

// TestLogsAlertDestination_Import covers the project_id/alert_id/hog_function_id form
// documented in examples/resources/posthog_logs_alert_destination/import.sh. The import
// names one hog function and the read has to find the whole group behind it.
func TestLogsAlertDestination_Import(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	config := testAccLogsAlertDestinationWebhook(rName, "https://example.com/hooks/first")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				Config:       config,
				ResourceName: logsAlertDestinationAddress,
				ImportState:  true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[logsAlertDestinationAddress]
					if !ok {
						return "", fmt.Errorf("resource not found: %s", logsAlertDestinationAddress)
					}
					// One hog function id out of the group, which is all the import ID names.
					hogFunctionID := strings.Split(rs.Primary.ID, ",")[0]
					return fmt.Sprintf("%s/%s/%s", getProjectID(), rs.Primary.Attributes["alert_id"], hogFunctionID), nil
				},
				ImportStateVerify: true,
				// slack_channel_name is write-only, so an imported destination cannot have
				// one. This config sets no channel name either, but the ignore documents why
				// a comparison would be wrong.
				ImportStateVerifyIgnore: []string{"slack_channel_name"},
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) != 1 {
						return fmt.Errorf("expected 1 imported state, got %d", len(states))
					}
					// The import names one hog function; the read rewrites the id to the
					// whole group, so the imported id must list more than what was given.
					if got := states[0].Attributes["hog_function_ids.#"]; got == "" || got == "0" {
						return fmt.Errorf("imported destination has no hog_function_ids")
					}
					return nil
				},
			},
		},
	})
}

// TestLogsAlertDestination_ExternalDeletion deletes the destination outside Terraform, as
// someone would from the PostHog UI. The read must treat it as gone and plan a fresh create
// rather than erroring, and the alert it hangs off must survive.
func TestLogsAlertDestination_ExternalDeletion(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	config := testAccLogsAlertDestinationWebhook(rName, "https://example.com/hooks/first")

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"acceptance-test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	var alertID string
	var hogFunctionIDs []string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLogsAlertDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[logsAlertDestinationAddress]
					if !ok {
						return fmt.Errorf("resource not found: %s", logsAlertDestinationAddress)
					}
					alertID = rs.Primary.Attributes["alert_id"]
					hogFunctionIDs = strings.Split(rs.Primary.ID, ",")
					return nil
				},
			},
			{
				PreConfig: func() {
					if _, err := (&client).DeleteLogsAlertDestination(context.Background(), projectID, alertID, hogFunctionIDs); err != nil {
						t.Fatalf("deleting logs alert destination %v externally: %v", hogFunctionIDs, err)
					}
				},
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(logsAlertDestinationAddress, "webhook_url", "https://example.com/hooks/first"),
					testAccCheckLogsAlertDestinationExists(logsAlertDestinationAddress),
				),
			},
		},
	})
}

// TestLogsAlertDestination_RejectsInvalidConfigs checks that the type and its settings are
// matched at plan time, before any API call. The rules themselves are covered by the unit
// tables; what this adds is that ModifyResourcePlan and the schema validators are wired into
// the registered resource, and that each message is the one a practitioner sees.
func TestLogsAlertDestination_RejectsInvalidConfigs(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoLogsAlerting(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	tests := map[string]struct {
		body      string
		wantError *regexp.Regexp
	}{
		"slack without a workspace": {
			body:      "type = \"slack\"\n  slack_channel_id = \"C0123456789\"",
			wantError: regexp.MustCompile(`Missing Slack destination settings`),
		},
		"slack without a channel": {
			body:      "type = \"slack\"\n  slack_workspace_id = 1",
			wantError: regexp.MustCompile(`Missing Slack destination settings`),
		},
		"slack with a webhook url": {
			body: "type = \"slack\"\n  slack_workspace_id = 1\n  slack_channel_id = \"C0123456789\"\n" +
				"  webhook_url = \"https://example.com/hooks/first\"",
			wantError: regexp.MustCompile(`does not apply to this destination type`),
		},
		"webhook without a url": {
			body:      `type = "webhook"`,
			wantError: regexp.MustCompile(`Missing destination URL`),
		},
		"teams without a url": {
			body:      `type = "teams"`,
			wantError: regexp.MustCompile(`Missing destination URL`),
		},
		"webhook with slack settings": {
			body: "type = \"webhook\"\n  webhook_url = \"https://example.com/hooks/first\"\n" +
				"  slack_channel_id = \"C0123456789\"",
			wantError: regexp.MustCompile(`does not apply to this destination type`),
		},
		"webhook with a slack channel name": {
			body: "type = \"webhook\"\n  webhook_url = \"https://example.com/hooks/first\"\n" +
				"  slack_channel_name = \"#alerts\"",
			wantError: regexp.MustCompile(`does not apply to this destination type`),
		},
		"unknown type": {
			body:      `type = "carrier-pigeon"`,
			wantError: regexp.MustCompile(`Attribute type value must be one of`),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      testAccLogsAlertDestinationConfig(rName, test.body),
						ExpectError: test.wantError,
					},
				},
			})
		})
	}
}

// testAccLogsAlertDestinationConfig wraps a destination body in the alert it hangs off.
func testAccLogsAlertDestinationConfig(name, body string) string {
	return fmt.Sprintf(`
resource "posthog_logs_alert" "test" {
  name            = %[1]q
  severity_levels = ["error"]
}

resource "posthog_logs_alert_destination" "test" {
  alert_id = posthog_logs_alert.test.id
  %[2]s
}
`, name, body)
}

func testAccLogsAlertDestinationWebhook(name, url string) string {
	return testAccLogsAlertDestinationConfig(name, fmt.Sprintf("type = \"webhook\"\n  webhook_url = %[1]q", url))
}

func testAccLogsAlertDestinationTeams(name, url string) string {
	return testAccLogsAlertDestinationConfig(name, fmt.Sprintf("type = \"teams\"\n  webhook_url = %[1]q", url))
}

func testAccLogsAlertDestinationSlack(name, channelID, channelName string) string {
	return testAccLogsAlertDestinationConfig(name, fmt.Sprintf(
		"type = \"slack\"\n  slack_workspace_id = %[1]s\n  slack_channel_id = %[2]q\n  slack_channel_name = %[3]q",
		getSlackWorkspaceID(), channelID, channelName))
}
