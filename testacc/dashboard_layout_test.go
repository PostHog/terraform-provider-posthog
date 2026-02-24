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

// TestDashboardLayout_CRUD tests the full create/update/destroy lifecycle of a
// posthog_dashboard_layout resource: create with a single insight tile, add a
// text tile, update the insight tile position, change the text tile body
// (in-place update via positional fallback), then implicitly destroy.
func TestDashboardLayout_CRUD(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix("tf-acc-test")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create — single insight tile
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_dashboard_layout.test", "id"),
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.#", "1"),
					resource.TestCheckResourceAttrSet("posthog_dashboard_layout.test", "tiles.0.tile_id"),
					resource.TestCheckResourceAttrPair("posthog_dashboard_layout.test", "dashboard_id", "posthog_dashboard.test", "id"),
				),
			},
			// Step 2: Update — add a text tile (now 2 tiles total)
			{
				Config: testAccDashboardLayoutInsightAndText(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.#", "2"),
				),
			},
			// Step 3: Update — change layout position of the insight tile
			{
				Config: testAccDashboardLayoutUpdatedPosition(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.#", "2"),
				),
			},
			// Step 4: Update — change text tile body (should update in place, not orphan)
			{
				Config: testAccDashboardLayoutChangedTextBody(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.#", "2"),
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.0.text_body", "## Updated Header"),
				),
			},
			// Step 5: Update — remove text tile, back to insight-only
			{
				Config: testAccDashboardLayoutSingleInsightShifted(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.#", "1"),
				),
			},
			// Implicit destroy at end of resource.Test
		},
	})
}

// TestDashboardLayout_Import tests that a posthog_dashboard_layout can be imported
// and that the resulting state matches the configuration.
func TestDashboardLayout_Import(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix("tf-acc-test")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_dashboard_layout.test", "id"),
				),
			},
			// Step 2: Import
			{
				ResourceName:      "posthog_dashboard_layout.test",
				ImportState:       true,
				ImportStateVerify: true,
				// tiles is ignored because import returns API-order tiles, which may differ
				// from config order. The tile count and existence is validated by the framework.
				ImportStateVerifyIgnore: []string{"tiles"},
			},
		},
	})
}

// TestDashboardLayout_AuthoritativeTextTileDelete verifies that text tiles added
// outside of Terraform are soft-deleted by apply. The resource is fully authoritative:
// unmanaged text tiles are deleted, unmanaged insight tiles have layouts cleared.
func TestDashboardLayout_AuthoritativeTextTileDelete(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix("tf-acc-test")

	host := os.Getenv("POSTHOG_HOST")
	apiKey := os.Getenv("POSTHOG_API_KEY")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")
	client := httpclient.NewDefaultClient(host, apiKey, "acceptance-test")

	var dashboardID string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create layout with a single insight tile
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.#", "1"),
					// Capture dashboard ID for API calls
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["posthog_dashboard.test"]
						if !ok {
							return fmt.Errorf("resource not found: posthog_dashboard.test")
						}
						dashboardID = rs.Primary.ID
						return nil
					},
				),
			},
			// Step 2: Inject an external text tile via direct API call, then re-apply same config
			{
				PreConfig: func() {
					// Inject a text tile with layouts directly via the PostHog API (outside Terraform).
					externalLayouts := map[string]interface{}{
						"sm": map[string]interface{}{"x": float64(0), "y": float64(10), "w": float64(12), "h": float64(2)},
					}
					_, _, err := (&client).UpdateDashboardLayout(
						context.Background(), projectID, dashboardID,
						httpclient.DashboardLayoutPatchRequest{
							Tiles: []httpclient.DashboardTilePatchItem{
								{
									Text:    &httpclient.DashboardTileTextPatch{Body: "Externally added tile"},
									Layouts: &externalLayouts,
								},
							},
						},
					)
					if err != nil {
						t.Fatalf("inject external tile via API: %v", err)
					}
				},
				Config: testAccDashboardLayoutSingleInsightShifted(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Terraform state should still show only 1 managed tile
					resource.TestCheckResourceAttr("posthog_dashboard_layout.test", "tiles.#", "1"),
					// Verify the external text tile was soft-deleted by authoritative enforcement
					func(s *terraform.State) error {
						resp, _, err := (&client).GetDashboardLayout(
							context.Background(), projectID, dashboardID,
						)
						if err != nil {
							return fmt.Errorf("get dashboard layout: %w", err)
						}
						// The external text tile should have been soft-deleted.
						for _, tile := range resp.Tiles {
							if tile.Text != nil && tile.Text.Body == "Externally added tile" {
								if tile.Deleted == nil || !*tile.Deleted {
									return fmt.Errorf("external text tile should have been soft-deleted by authoritative enforcement")
								}
							}
						}
						return nil
					},
				),
			},
		},
	})
}

// testAccDashboardLayoutSingleInsight returns HCL for a dashboard + insight + layout
// with a single insight tile at position sm={x:0, y:0, w:6, h:4}.
func testAccDashboardLayoutSingleInsight(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %[1]q
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$pageview"
        math  = "total"
      }]
    }
  })

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}

resource "posthog_dashboard_layout" "test" {
  dashboard_id = posthog_dashboard.test.id

  tiles = [
    {
      insight_id   = posthog_insight.test.id
      layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 6, h = 4 } })
    },
  ]

  depends_on = [posthog_insight.test]
}
`, name)
}

// testAccDashboardLayoutChangedTextBody returns HCL with the text tile body changed
// from "## Test Header" to "## Updated Header". Verifies in-place update via positional fallback.
func testAccDashboardLayoutChangedTextBody(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %[1]q
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$pageview"
        math  = "total"
      }]
    }
  })

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}

resource "posthog_dashboard_layout" "test" {
  dashboard_id = posthog_dashboard.test.id

  tiles = [
    {
      text_body    = "## Updated Header"
      layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
    },
    {
      insight_id   = posthog_insight.test.id
      layouts_json = jsonencode({ sm = { x = 6, y = 1, w = 6, h = 4 } })
    },
  ]

  depends_on = [posthog_insight.test]
}
`, name)
}

// testAccDashboardLayoutSingleInsightShifted returns the same layout as
// testAccDashboardLayoutSingleInsight but with the insight tile at x=3 instead of x=0,
// forcing an Update when used after the base config.
func testAccDashboardLayoutSingleInsightShifted(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %[1]q
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$pageview"
        math  = "total"
      }]
    }
  })

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}

resource "posthog_dashboard_layout" "test" {
  dashboard_id = posthog_dashboard.test.id

  tiles = [
    {
      insight_id   = posthog_insight.test.id
      layouts_json = jsonencode({ sm = { x = 3, y = 0, w = 6, h = 4 } })
    },
  ]

  depends_on = [posthog_insight.test]
}
`, name)
}

// testAccDashboardLayoutInsightAndText returns HCL for a layout with the insight tile
// moved to y=1 and a new text tile added at sm={x:0, y:0, w:12, h:1}.
func testAccDashboardLayoutInsightAndText(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %[1]q
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$pageview"
        math  = "total"
      }]
    }
  })

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}

resource "posthog_dashboard_layout" "test" {
  dashboard_id = posthog_dashboard.test.id

  tiles = [
    {
      text_body    = "## Test Header"
      layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
    },
    {
      insight_id   = posthog_insight.test.id
      layouts_json = jsonencode({ sm = { x = 0, y = 1, w = 6, h = 4 } })
    },
  ]

  depends_on = [posthog_insight.test]
}
`, name)
}

// testAccDashboardLayoutUpdatedPosition returns HCL for the same layout as
// testAccDashboardLayoutInsightAndText but with the insight tile moved to x=6
// (shifted right).
func testAccDashboardLayoutUpdatedPosition(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %[1]q
}

resource "posthog_insight" "test" {
  name = "%[1]s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$pageview"
        math  = "total"
      }]
    }
  })

  dashboard_ids = [posthog_dashboard.test.id]
  depends_on    = [posthog_dashboard.test]
}

resource "posthog_dashboard_layout" "test" {
  dashboard_id = posthog_dashboard.test.id

  tiles = [
    {
      text_body    = "## Test Header"
      layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
    },
    {
      insight_id   = posthog_insight.test.id
      layouts_json = jsonencode({ sm = { x = 6, y = 1, w = 6, h = 4 } })
    },
  ]

  depends_on = [posthog_insight.test]
}
`, name)
}
