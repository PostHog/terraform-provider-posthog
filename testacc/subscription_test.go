package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// slackIntegrationID returns the Slack integration ID to use in Slack subscription tests.
// Defaults to "1" (the live test stack's Slack integration) but can be overridden.
func slackIntegrationID() string {
	if v := os.Getenv("POSTHOG_SLACK_INTEGRATION_ID"); v != "" {
		return v
	}
	return "1"
}

// testAccCheckSubscriptionDestroy verifies each subscription has been soft-deleted.
// Hard DELETE is unsupported; GET still returns the row with deleted=true, so we assert
// either not-found or deleted=true.
func testAccCheckSubscriptionDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"acceptance-test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_subscription" {
			continue
		}
		sub, status, err := client.GetSubscription(context.Background(), projectID, rs.Primary.ID)
		if err != nil {
			if status == 404 {
				continue
			}
			return fmt.Errorf("unexpected error checking subscription %s: %w", rs.Primary.ID, err)
		}
		if sub.Deleted == nil || !*sub.Deleted {
			return fmt.Errorf("subscription %s still exists (not soft-deleted)", rs.Primary.ID)
		}
	}
	return nil
}

// TestSubscription_SlackDashboard is the headline flow: a daily Slack dashboard digest
// created alongside its dashboard and insights, asserting create, no post-apply drift,
// and round-trip import.
func TestSubscription_SlackDashboard(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-sub")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriptionSlackDashboard(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_subscription.test", "id"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_type", "slack"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_value", "C0B9A53J8RF|#reports"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "frequency", "daily"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "interval", "1"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "start_date", "2026-08-17T07:00:00Z"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "dashboard_export_insights.#", "2"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "title", rName),
					resource.TestCheckResourceAttr("posthog_subscription.test", "enabled", "true"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "resource_type", "dashboard"),
					resource.TestCheckResourceAttrSet("posthog_subscription.test", "summary"),
					resource.TestCheckResourceAttrPair(
						"posthog_subscription.test", "dashboard_id",
						"posthog_dashboard.test", "id",
					),
				),
			},
			// No drift: re-planning the same config must be a no-op (guards target_value
			// verbatim storage and start_date normalization against a perpetual diff).
			{
				Config:   testAccSubscriptionSlackDashboard(rName),
				PlanOnly: true,
			},
			{
				ResourceName:            "posthog_subscription.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"next_delivery_date"},
			},
		},
	})
}

// TestSubscription_EmailInsight exercises the general email + single-insight path with a
// weekly byweekday recurrence, and toggles enabled to false on update.
func TestSubscription_EmailInsight(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-sub")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriptionEmailInsight(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_type", "email"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_value", "team@example.com"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "frequency", "weekly"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "byweekday.#", "1"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "enabled", "true"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "resource_type", "insight"),
					resource.TestCheckNoResourceAttr("posthog_subscription.test", "integration_id"),
					resource.TestCheckResourceAttrPair(
						"posthog_subscription.test", "insight_id",
						"posthog_insight.a", "id",
					),
				),
			},
			// Pause (enabled=false) without deleting.
			{
				Config: testAccSubscriptionEmailInsight(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "enabled", "false"),
				),
			},
		},
	})
}

// TestSubscription_NonUTCStartDate proves a non-UTC (but valid RFC3339) start_date
// round-trips without an inconsistent-result error or a perpetual diff. start_date is
// Required and the provider normalizes the API's echoed value, so the configured offset
// value must be preserved when it denotes the same instant.
func TestSubscription_NonUTCStartDate(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-sub")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriptionNonUTCStartDate(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// The configured non-UTC value is preserved verbatim (no inconsistent result).
					resource.TestCheckResourceAttr("posthog_subscription.test", "start_date", "2026-08-17T09:00:00+02:00"),
				),
			},
			{
				// Re-plan must be clean: the offset value must not perpetually diff.
				Config:   testAccSubscriptionNonUTCStartDate(rName),
				PlanOnly: true,
			},
		},
	})
}

// TestSubscription_AIPrompt exercises the AI-summary ("prompt") path: no dashboard/insight
// subject, an ai_prompt with a JSON analysis window and summary options, asserting the
// server-inferred ai_prompt resource_type, no post-apply drift, and round-trip import.
func TestSubscription_AIPrompt(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-sub")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriptionAIPrompt(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_subscription.test", "id"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_type", "email"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_value", "growth@example.com"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "ai_prompt", "Top 5 events by volume this week"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "summary_enabled", "true"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "summary_prompt_guide", "Focus on week-over-week growth"),
					resource.TestCheckResourceAttrSet("posthog_subscription.test", "ai_prompt_config"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "frequency", "weekly"),
					// server-inferred subject type + no dashboard/insight subject
					resource.TestCheckResourceAttr("posthog_subscription.test", "resource_type", "ai_prompt"),
					resource.TestCheckNoResourceAttr("posthog_subscription.test", "dashboard_id"),
					resource.TestCheckNoResourceAttr("posthog_subscription.test", "insight_id"),
				),
			},
			// No drift: re-planning must be a no-op (guards ai_prompt_config JSON normalization
			// and the summary_enabled computed default against a perpetual diff).
			{
				Config:   testAccSubscriptionAIPrompt(rName),
				PlanOnly: true,
			},
			{
				ResourceName:            "posthog_subscription.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"next_delivery_date"},
			},
		},
	})
}

// TestSubscription_ClearsOptionalFields guards the update-clearing path: an optional field
// removed from config must actually reset server-side instead of leaving the stale value
// (which would surface as "inconsistent result after apply"). It sets integration_id,
// byweekday and bysetpos, then removes all three in a single update (a Slack->email
// transition, which is the only in-place context where integration_id may be cleared),
// asserting each clears and that the follow-up plan is empty.
//
// dashboard_export_insights is intentionally not covered: it is required non-empty on a
// dashboard subscription and rejected on any other subject, so it has no valid same-resource
// clear path (the request still emits an explicit [], turning a removal into a clean API
// validation error rather than a silent inconsistent-result).
func TestSubscription_ClearsOptionalFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-sub")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriptionClearable(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_type", "slack"),
					resource.TestCheckResourceAttrSet("posthog_subscription.test", "integration_id"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "byweekday.#", "1"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "bysetpos", "1"),
				),
			},
			// Remove integration_id, byweekday and bysetpos: each must clear (no drift).
			{
				Config: testAccSubscriptionClearable(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "target_type", "email"),
					resource.TestCheckNoResourceAttr("posthog_subscription.test", "integration_id"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "byweekday.#", "0"),
					resource.TestCheckNoResourceAttr("posthog_subscription.test", "bysetpos"),
				),
			},
			// The cleared config must re-plan as a no-op.
			{
				Config:   testAccSubscriptionClearable(rName, false),
				PlanOnly: true,
			},
		},
	})
}

// TestSubscription_UpdateValues covers an in-place update that changes fields to new,
// non-null values (not just toggling or clearing): title, interval, and byweekday all
// change, and the update must apply and then re-plan clean.
func TestSubscription_UpdateValues(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-sub")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriptionUpdatableValues(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "title", rName+" report"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "interval", "1"),
					resource.TestCheckTypeSetElemAttr("posthog_subscription.test", "byweekday.*", "monday"),
				),
			},
			{
				// In-place update: change title, interval and byweekday to new values.
				Config: testAccSubscriptionUpdatableValues(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "title", rName+" report (renamed)"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "interval", "2"),
					resource.TestCheckTypeSetElemAttr("posthog_subscription.test", "byweekday.*", "friday"),
				),
			},
			{
				Config:   testAccSubscriptionUpdatableValues(rName, true),
				PlanOnly: true,
			},
		},
	})
}

// TestSubscription_ChangeSubjectInsightToPrompt switches the subscription subject from an
// insight to an ai_prompt. The subject fields are RequiresReplace + ExactlyOneOf, so this is
// a destroy+recreate transition and asserts the server-inferred resource_type flips
// insight -> ai_prompt (and insight_id is cleared).
func TestSubscription_ChangeSubjectInsightToPrompt(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-sub")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriptionSubjectSwitch(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "resource_type", "insight"),
					resource.TestCheckResourceAttrSet("posthog_subscription.test", "insight_id"),
				),
			},
			{
				// Switching subject forces replacement (insight_id is RequiresReplace).
				Config: testAccSubscriptionSubjectSwitch(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_subscription.test", "resource_type", "ai_prompt"),
					resource.TestCheckResourceAttr("posthog_subscription.test", "ai_prompt", "Top 5 events by volume this week"),
					resource.TestCheckNoResourceAttr("posthog_subscription.test", "insight_id"),
				),
			},
			{
				Config:   testAccSubscriptionSubjectSwitch(rName, true),
				PlanOnly: true,
			},
		},
	})
}

func testAccSubscriptionUpdatableValues(name string, updated bool) string {
	title := name + " report"
	weekday := "monday"
	interval := 1
	if updated {
		title = name + " report (renamed)"
		weekday = "friday"
		interval = 2
	}
	return fmt.Sprintf(`
provider "posthog" {}

%[1]s

resource "posthog_subscription" "test" {
  target_type  = "email"
  target_value = "team@example.com"
  insight_id   = posthog_insight.a.id
  frequency    = "weekly"
  interval     = %[4]d
  byweekday    = [%[3]q]
  start_date   = "2026-08-17T07:00:00Z"
  title        = %[2]q
}
`, testAccSubscriptionInsightsBase(name), title, weekday, interval)
}

func testAccSubscriptionSubjectSwitch(name string, usePrompt bool) string {
	subject := `insight_id   = posthog_insight.a.id`
	if usePrompt {
		subject = `ai_prompt        = "Top 5 events by volume this week"
  ai_prompt_config = jsonencode({
    window = { mode = "last_n_days", start_days_ago = 7, end_days_ago = null }
  })
  summary_enabled = true`
	}
	return fmt.Sprintf(`
provider "posthog" {}

%[1]s

resource "posthog_subscription" "test" {
  target_type  = "email"
  target_value = "team@example.com"
  %[2]s
  frequency    = "weekly"
  interval     = 1
  byweekday    = ["monday"]
  start_date   = "2026-08-17T07:00:00Z"
  title        = %[3]q
}
`, testAccSubscriptionInsightsBase(name), subject, name)
}

// testAccSubscriptionClearable renders an insight subscription with (withOptionals=true) or
// without (false) the clearable optionals integration_id/byweekday/bysetpos. The insight
// subject is unchanged across both so the transition is an in-place update, not a replace.
func testAccSubscriptionClearable(name string, withOptionals bool) string {
	if withOptionals {
		return fmt.Sprintf(`
provider "posthog" {}

%[1]s

resource "posthog_subscription" "test" {
  target_type    = "slack"
  target_value   = "C0B9A53J8RF|#reports"
  integration_id = %[2]s
  insight_id     = posthog_insight.a.id
  frequency      = "monthly"
  interval       = 1
  byweekday      = ["monday"]
  bysetpos       = 1
  start_date     = "2026-08-17T07:00:00Z"
}
`, testAccSubscriptionInsightsBase(name), slackIntegrationID())
	}
	return fmt.Sprintf(`
provider "posthog" {}

%[1]s

resource "posthog_subscription" "test" {
  target_type  = "email"
  target_value = "team@example.com"
  insight_id   = posthog_insight.a.id
  frequency    = "monthly"
  interval     = 1
  start_date   = "2026-08-17T07:00:00Z"
}
`, testAccSubscriptionInsightsBase(name))
}

func testAccSubscriptionAIPrompt(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_subscription" "test" {
  target_type  = "email"
  target_value = "growth@example.com"
  ai_prompt    = "Top 5 events by volume this week"
  ai_prompt_config = jsonencode({
    window = { mode = "last_n_days", start_days_ago = 7, end_days_ago = null }
  })
  summary_enabled      = true
  summary_prompt_guide = "Focus on week-over-week growth"
  frequency            = "weekly"
  interval             = 1
  byweekday            = ["monday"]
  start_date           = "2026-08-17T07:00:00Z"
  title                = %[1]q
}
`, name)
}

// testAccSubscriptionInsightsBase creates a dashboard with two insights attached to it.
func testAccSubscriptionInsightsBase(name string) string {
	return fmt.Sprintf(`
resource "posthog_dashboard" "test" {
  name = "%[1]s-dashboard"
}

resource "posthog_insight" "a" {
  name          = "%[1]s-insight-a"
  dashboard_ids = [posthog_dashboard.test.id]

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{ kind = "EventsNode", event = "$pageview", math = "total" }]
    }
  })
}

resource "posthog_insight" "b" {
  name          = "%[1]s-insight-b"
  dashboard_ids = [posthog_dashboard.test.id]

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{ kind = "EventsNode", event = "$autocapture", math = "total" }]
    }
  })
}
`, name)
}

func testAccSubscriptionNonUTCStartDate(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%[1]s

resource "posthog_subscription" "test" {
  target_type  = "email"
  target_value = "team@example.com"
  insight_id   = posthog_insight.a.id
  frequency    = "weekly"
  interval     = 1
  byweekday    = ["monday"]
  start_date   = "2026-08-17T09:00:00+02:00"
}
`, testAccSubscriptionInsightsBase(name))
}

func testAccSubscriptionSlackDashboard(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%[1]s

resource "posthog_subscription" "test" {
  target_type               = "slack"
  target_value              = "C0B9A53J8RF|#reports"
  integration_id            = %[2]s
  dashboard_id              = posthog_dashboard.test.id
  dashboard_export_insights = [posthog_insight.a.id, posthog_insight.b.id]
  frequency                 = "daily"
  interval                  = 1
  start_date                = "2026-08-17T07:00:00Z"
  title                     = %[3]q
}
`, testAccSubscriptionInsightsBase(name), slackIntegrationID(), name)
}

func testAccSubscriptionEmailInsight(name string, enabled bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

%[1]s

resource "posthog_subscription" "test" {
  target_type  = "email"
  target_value = "team@example.com"
  insight_id   = posthog_insight.a.id
  frequency    = "weekly"
  interval     = 1
  byweekday    = ["monday"]
  start_date   = "2026-08-17T07:00:00Z"
  enabled      = %[2]t
}
`, testAccSubscriptionInsightsBase(name), enabled)
}
