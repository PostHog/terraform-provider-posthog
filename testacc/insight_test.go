package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestInsight_Basic tests creating an insight with minimal query_json.
func TestInsight_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "query_json"),
				),
			},
		},
	})
}

// TestInsight_AllFields tests creating an insight with all optional fields.
func TestInsight_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_insight.test", "description", "Test insight description"),
					resource.TestCheckResourceAttr("posthog_insight.test", "derived_name", "Custom derived name"),
					resource.TestCheckResourceAttr("posthog_insight.test", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "id"),
				),
			},
		},
	})
}

// TestInsight_TrendsQuery tests creating an insight with a trends query.
func TestInsight_TrendsQuery(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightTrendsQuery(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "query_json"),
				),
			},
		},
	})
}

// TestInsight_FunnelsQuery tests creating an insight with a funnels query.
func TestInsight_FunnelsQuery(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightFunnelsQuery(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "query_json"),
				),
			},
		},
	})
}

// TestInsight_RetentionQuery tests creating an insight with a retention query.
func TestInsight_RetentionQuery(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightRetentionQuery(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "query_json"),
				),
			},
		},
	})
}

// TestInsight_LifecycleQuery tests creating an insight with a lifecycle query.
func TestInsight_LifecycleQuery(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightLifecycleQuery(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "query_json"),
				),
			},
		},
	})
}

// TestInsight_PathsQuery tests creating an insight with a paths query.
func TestInsight_PathsQuery(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightPathsQuery(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "query_json"),
				),
			},
		},
	})
}

// TestInsight_Update tests updating query and metadata.
func TestInsight_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccInsightWithDescription(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_insight.test", "description", "Initial description"),
				),
			},
			// Update name
			{
				Config: testAccInsightWithDescription(rName+"-updated", "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_insight.test", "description", "Initial description"),
				),
			},
			// Update description
			{
				Config: testAccInsightWithDescription(rName+"-updated", "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_insight.test", "description", "Updated description"),
				),
			},
		},
	})
}

// TestInsight_Tags tests creating, updating, and removing tags.
// This test verifies the fix for the "Provider produced inconsistent result after apply" error
// where tags were being returned as null/empty after creation/update.
func TestInsight_Tags(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with tags - verify both count and specific values
			{
				Config: testAccInsightWithTags(rName, []string{"tag1", "tag2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "tags.#", "2"),
					resource.TestCheckTypeSetElemAttr("posthog_insight.test", "tags.*", "tag1"),
					resource.TestCheckTypeSetElemAttr("posthog_insight.test", "tags.*", "tag2"),
				),
			},
			// Add more tags
			{
				Config: testAccInsightWithTags(rName, []string{"tag1", "tag2", "tag3"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "tags.#", "3"),
					resource.TestCheckTypeSetElemAttr("posthog_insight.test", "tags.*", "tag1"),
					resource.TestCheckTypeSetElemAttr("posthog_insight.test", "tags.*", "tag2"),
					resource.TestCheckTypeSetElemAttr("posthog_insight.test", "tags.*", "tag3"),
				),
			},
			// Remove tags
			{
				Config: testAccInsightWithTags(rName, []string{"tag1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "tags.#", "1"),
					resource.TestCheckTypeSetElemAttr("posthog_insight.test", "tags.*", "tag1"),
				),
			},
		},
	})
}

// TestInsight_WithDashboard tests creating an insight linked to a dashboard.
func TestInsight_WithDashboard(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightWithDashboard(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "1"),
				),
			},
		},
	})
}

// TestInsight_MultipleDashboards tests linking an insight to multiple dashboards.
func TestInsight_MultipleDashboards(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightMultipleDashboards(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_insight.test", "dashboard_ids.#", "2"),
				),
			},
		},
	})
}

// TestInsight_Import tests importing an existing insight by ID.
func TestInsight_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccInsightBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
				),
			},
			// Import
			{
				ResourceName:            "posthog_insight.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"query_json", "create_in_folder"},
			},
		},
	})
}

// TestDashboard_EmptyDescription tests handling of empty vs null description.
func TestInsight_EmptyDescription(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccInsightWithDescription(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_insight.test", "description", "Initial description"),
				),
			},
			// Remove description
			{
				Config: testAccInsightBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName),
					resource.TestCheckNoResourceAttr("posthog_insight.test", "description"),
				),
			},
		},
	})
}

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

func testAccInsightBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name = %q

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

func testAccInsightAllFields(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name         = %q
  description  = "Test insight description"
  derived_name = "Custom derived name"
  tags         = ["managed-by:terraform", "env:test"]

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

func testAccInsightTrendsQuery(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name        = %q
  description = "Trends query insight"

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
      trendsFilter = {
        showLegend = true
        display    = "ActionsLineGraph"
      }
      interval = "day"
    }
  })
}
`, name)
}

func testAccInsightFunnelsQuery(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name        = %q
  description = "Funnels query insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "FunnelsQuery"
      series = [
        {
          kind  = "EventsNode"
          name  = "$pageview"
          event = "$pageview"
        },
        {
          kind  = "EventsNode"
          name  = "$autocapture"
          event = "$autocapture"
        }
      ]
      funnelsFilter = {
        funnelVizType = "steps"
      }
    }
  })
}
`, name)
}

func testAccInsightRetentionQuery(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name        = %q
  description = "Retention query insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind = "RetentionQuery"
      retentionFilter = {
        targetEntity = {
          id   = "$pageview"
          name = "$pageview"
          type = "events"
        }
        returningEntity = {
          id   = "$pageview"
          name = "$pageview"
          type = "events"
        }
        retentionType = "retention_first_time"
        period        = "Day"
        totalIntervals = 7
      }
    }
  })
}
`, name)
}

func testAccInsightWithDescription(name, description string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name        = %q
  description = %q

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
`, name, description)
}

func testAccInsightWithTags(name string, tags []string) string {
	tagsStr := ""
	for i, tag := range tags {
		if i > 0 {
			tagsStr += ", "
		}
		tagsStr += fmt.Sprintf("%q", tag)
	}

	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name = %q
  tags = [%s]

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
`, name, tagsStr)
}

func testAccInsightWithDashboard(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name        = "%[1]s-dashboard"
  description = "Dashboard for insight test"
}

resource "posthog_insight" "test" {
  name        = %[1]q
  description = "Insight attached to dashboard"

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

  depends_on = [posthog_dashboard.test]
}
`, name)
}

func testAccInsightMultipleDashboards(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test1" {
  name = "%[1]s-dashboard-1"
}

resource "posthog_dashboard" "test2" {
  name = "%[1]s-dashboard-2"
}

resource "posthog_insight" "test" {
  name        = %[1]q
  description = "Insight on multiple dashboards"

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

  depends_on = [posthog_dashboard.test1, posthog_dashboard.test2]
}
`, name)
}

func testAccInsightLifecycleQuery(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name        = %q
  description = "Lifecycle query insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "LifecycleQuery"
      series = [{
        kind  = "EventsNode"
        name  = "$pageview"
        event = "$pageview"
      }]
      lifecycleFilter = {}
    }
  })
}
`, name)
}

func testAccInsightPathsQuery(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name        = %q
  description = "Paths query insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind = "PathsQuery"
      pathsFilter = {
        includeEventTypes = ["$pageview"]
      }
    }
  })
}
`, name)
}
