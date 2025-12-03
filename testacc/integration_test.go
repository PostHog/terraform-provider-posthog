package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestIntegration_InsightWithDashboard tests creating a dashboard and then an insight linked to it.
func TestIntegration_InsightWithDashboard(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationInsightWithDashboard(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify dashboard was created
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName+"-dashboard"),
					resource.TestCheckResourceAttrSet("posthog_dashboard.test", "id"),
					// Verify insight was created and linked
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName+"-insight"),
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "1"),
				),
			},
		},
	})
}

// TestIntegration_MultipleInsightsOnDashboard tests creating multiple insights on the same dashboard.
func TestIntegration_MultipleInsightsOnDashboard(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationMultipleInsightsOnDashboard(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify dashboard
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName+"-dashboard"),
					// Verify all insights are linked to the same dashboard
					resource.TestCheckResourceAttr("posthog_insight.test1", "dashboard_ids.#", "1"),
					resource.TestCheckResourceAttr("posthog_insight.test2", "dashboard_ids.#", "1"),
					resource.TestCheckResourceAttr("posthog_insight.test3", "dashboard_ids.#", "1"),
				),
			},
		},
	})
}

// TestIntegration_UpdateDashboardWithInsights tests updating a dashboard while insights reference it.
func TestIntegration_UpdateDashboardWithInsights(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create dashboard and insight
			{
				Config: testAccIntegrationDashboardWithInsight(rName, "Initial Dashboard", "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", "Initial Dashboard"),
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "1"),
				),
			},
			// Update dashboard name - insight should still be linked
			{
				Config: testAccIntegrationDashboardWithInsight(rName, "Updated Dashboard", "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", "Updated Dashboard"),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", "Updated description"),
					// Insight should still be linked
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "1"),
				),
			},
		},
	})
}

// TestIntegration_MoveInsightBetweenDashboards tests changing an insight's dashboard_ids.
func TestIntegration_MoveInsightBetweenDashboards(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create insight on dashboard 1
			{
				Config: testAccIntegrationInsightOnDashboard1(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "1"),
				),
			},
			// Move insight to dashboard 2
			{
				Config: testAccIntegrationInsightOnDashboard2(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "1"),
				),
			},
			// Add insight to both dashboards
			{
				Config: testAccIntegrationInsightOnBothDashboards(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "2"),
				),
			},
		},
	})
}

// TestIntegration_FullStack tests creating a dashboard, multiple insights, and a feature flag together.
func TestIntegration_FullStack(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationFullStack(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Dashboard
					resource.TestCheckResourceAttr("posthog_dashboard.main", "name", rName+"-main-dashboard"),
					resource.TestCheckResourceAttr("posthog_dashboard.main", "tags.#", "1"),
					// Insights
					resource.TestCheckResourceAttr("posthog_insight.trends", "name", rName+"-trends"),
					resource.TestCheckResourceAttr("posthog_insight.trends", "dashboard_ids.#", "1"),
					resource.TestCheckResourceAttr("posthog_insight.funnels", "name", rName+"-funnels"),
					resource.TestCheckResourceAttr("posthog_insight.funnels", "dashboard_ids.#", "1"),
					// Feature flag
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "key", rName+"-flag"),
					resource.TestCheckResourceAttr("posthog_feature_flag.test", "active", "true"),
				),
			},
		},
	})
}

func testAccIntegrationInsightWithDashboard(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name        = "%[1]s-dashboard"
  description = "Integration test dashboard"
}

resource "posthog_insight" "test" {
  name        = "%[1]s-insight"
  description = "Integration test insight"

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

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}
`, name)
}

func testAccIntegrationMultipleInsightsOnDashboard(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name        = "%[1]s-dashboard"
  description = "Dashboard with multiple insights"
}

resource "posthog_insight" "test1" {
  name = "%[1]s-insight-1"

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

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}

resource "posthog_insight" "test2" {
  name = "%[1]s-insight-2"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        name  = "$autocapture"
        event = "$autocapture"
        math  = "total"
      }]
    }
  })

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}

resource "posthog_insight" "test3" {
  name = "%[1]s-insight-3"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "FunnelsQuery"
      series = [
        { kind = "EventsNode", name = "$pageview", event = "$pageview" },
        { kind = "EventsNode", name = "$autocapture", event = "$autocapture" }
      ]
    }
  })

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}
`, name)
}

func testAccIntegrationDashboardWithInsight(name, dashboardName, description string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name        = %q
  description = %q
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

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

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}
`, dashboardName, description, name)
}

func testAccIntegrationInsightOnDashboard1(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test1" {
  name = "%[1]s-dashboard-1"
}

resource "posthog_dashboard" "test2" {
  name = "%[1]s-dashboard-2"
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

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

  dashboard_ids = [posthog_dashboard.test1.id]
  depends_on    = [posthog_dashboard.test1, posthog_dashboard.test2]
}
`, name)
}

func testAccIntegrationInsightOnDashboard2(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test1" {
  name = "%[1]s-dashboard-1"
}

resource "posthog_dashboard" "test2" {
  name = "%[1]s-dashboard-2"
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

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

  dashboard_ids = [posthog_dashboard.test2.id]
  depends_on    = [posthog_dashboard.test1, posthog_dashboard.test2]
}
`, name)
}

func testAccIntegrationInsightOnBothDashboards(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test1" {
  name = "%[1]s-dashboard-1"
}

resource "posthog_dashboard" "test2" {
  name = "%[1]s-dashboard-2"
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

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

  dashboard_ids = [posthog_dashboard.test1.id, posthog_dashboard.test2.id]
  depends_on    = [posthog_dashboard.test1, posthog_dashboard.test2]
}
`, name)
}

func testAccIntegrationFullStack(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "main" {
  name        = "%[1]s-main-dashboard"
  description = "Main integration test dashboard"
  tags        = ["integration-test"]
}

resource "posthog_insight" "trends" {
  name        = "%[1]s-trends"
  description = "Trends insight for integration test"

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

  dashboard_ids = [posthog_dashboard.main.id]
  depends_on    = [posthog_dashboard.main]
}

resource "posthog_insight" "funnels" {
  name        = "%[1]s-funnels"
  description = "Funnels insight for integration test"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "FunnelsQuery"
      series = [
        { kind = "EventsNode", name = "$pageview", event = "$pageview" },
        { kind = "EventsNode", name = "$autocapture", event = "$autocapture" }
      ]
    }
  })

  dashboard_ids = [posthog_dashboard.main.id]
  depends_on    = [posthog_dashboard.main]
}

resource "posthog_feature_flag" "test" {
  key    = "%[1]s-flag"
  name   = "Integration test feature flag"
  active = true
  tags   = ["integration-test"]

  filters = jsonencode({
    groups = [{
      properties         = []
      rollout_percentage = 100
    }]
  })
}
`, name)
}
