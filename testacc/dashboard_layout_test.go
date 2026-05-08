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

const (
	dashboardLayoutRandomPrefix        = "tf-acc-test"
	dashboardLayoutResourceName        = "posthog_dashboard_layout.test"
	dashboardResourceName              = "posthog_dashboard.test"
	dashboardLayoutTilesCountAttr      = "tiles.#"
	dashboardLayoutFirstTileIDAttr     = "tiles.0.tile_id"
	dashboardLayoutFirstTileColorAttr  = "tiles.0.color"
	dashboardLayoutShowDescriptionAttr = "tiles.0.show_description"
	dashboardLayoutResourceNotFoundFmt = "resource not found: %s"
)

// TestDashboardLayout_CRUD tests the full create/update/destroy lifecycle of a
// posthog_dashboard_layout resource: create with a single insight tile, add a
// text tile, update the insight tile position, change the text tile body
// (in-place update via positional fallback), then implicitly destroy.
func TestDashboardLayout_CRUD(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix(dashboardLayoutRandomPrefix)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create — single insight tile
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(dashboardLayoutResourceName, "id"),
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutTilesCountAttr, "1"),
					resource.TestCheckResourceAttrSet(dashboardLayoutResourceName, dashboardLayoutFirstTileIDAttr),
					resource.TestCheckResourceAttrPair(dashboardLayoutResourceName, "dashboard_id", dashboardResourceName, "id"),
				),
			},
			// Step 2: Update — add a text tile (now 2 tiles total)
			{
				Config: testAccDashboardLayoutInsightAndText(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutTilesCountAttr, "2"),
				),
			},
			// Step 3: Update — change layout position of the insight tile
			{
				Config: testAccDashboardLayoutUpdatedPosition(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutTilesCountAttr, "2"),
				),
			},
			// Step 4: Update — change text tile body (should update in place, not orphan)
			{
				Config: testAccDashboardLayoutChangedTextBody(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutTilesCountAttr, "2"),
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, "tiles.0.text_body", "## Updated Header"),
				),
			},
			// Step 5: Update — remove text tile, back to insight-only
			{
				Config: testAccDashboardLayoutSingleInsightShifted(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutTilesCountAttr, "1"),
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
	rName := acctest.RandomWithPrefix(dashboardLayoutRandomPrefix)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(dashboardLayoutResourceName, "id"),
				),
			},
			// Step 2: Import
			{
				ResourceName:      dashboardLayoutResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// tiles is ignored because import returns API-order tiles, which may differ
				// from config order. The tile count and existence is validated by the framework.
				ImportStateVerifyIgnore: []string{"tiles"},
			},
		},
	})
}

func TestDashboardLayout_ShowDescription(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix(dashboardLayoutRandomPrefix)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDashboardLayoutInsightWithShowDescription(rName, "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutShowDescriptionAttr, "false"),
				),
			},
			{
				ResourceName:            dashboardLayoutResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{dashboardLayoutFirstTileIDAttr},
			},
			{
				Config: testAccDashboardLayoutInsightWithShowDescription(rName, "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutShowDescriptionAttr, "true"),
				),
			},
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(dashboardLayoutResourceName, dashboardLayoutShowDescriptionAttr),
				),
			},
		},
	})
}

// TestDashboardLayout_AuthoritativeTextTileDelete verifies that text tiles added
// outside of Terraform are soft-deleted by apply. The resource is fully authoritative:
// unmanaged text tiles are deleted, unmanaged insight tiles have layouts cleared.
func TestDashboardLayout_AuthoritativeTextTileDelete(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix(dashboardLayoutRandomPrefix)

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
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutTilesCountAttr, "1"),
					// Capture dashboard ID for API calls
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[dashboardResourceName]
						if !ok {
							return fmt.Errorf(dashboardLayoutResourceNotFoundFmt, dashboardResourceName)
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
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutTilesCountAttr, "1"),
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

// TestDashboardLayout_Color tests the full lifecycle of the color field on a
// dashboard layout tile: set, change, remove, re-add.
func TestDashboardLayout_Color(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix(dashboardLayoutRandomPrefix)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with color = "blue"
			{
				Config: testAccDashboardLayoutInsightWithColor(rName, "blue"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutFirstTileColorAttr, "blue"),
				),
			},
			// Step 2: Change color to "green"
			{
				Config: testAccDashboardLayoutInsightWithColor(rName, "green"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutFirstTileColorAttr, "green"),
				),
			},
			// Step 3: Remove color (omit from config)
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(dashboardLayoutResourceName, dashboardLayoutFirstTileColorAttr),
				),
			},
			// Step 4: Re-add color "purple"
			{
				Config: testAccDashboardLayoutInsightWithColor(rName, "purple"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dashboardLayoutResourceName, dashboardLayoutFirstTileColorAttr, "purple"),
				),
			},
		},
	})
}

// TestDashboardLayout_ColorDriftClearedByApply verifies that a color set via the
// API (outside Terraform) is detected as drift and cleared on the next apply when
// the config does not declare a color.
func TestDashboardLayout_ColorDriftClearedByApply(t *testing.T) {
	skipIfNotAcceptance(t)
	rName := acctest.RandomWithPrefix(dashboardLayoutRandomPrefix)

	host := os.Getenv("POSTHOG_HOST")
	apiKey := os.Getenv("POSTHOG_API_KEY")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")
	client := httpclient.NewDefaultClient(host, apiKey, "acceptance-test")

	var dashboardID string
	var tileID int64

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create layout with no color
			{
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(dashboardLayoutResourceName, dashboardLayoutFirstTileColorAttr),
					// Capture IDs for the API call in step 2
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources[dashboardResourceName]
						if !ok {
							return fmt.Errorf(dashboardLayoutResourceNotFoundFmt, dashboardResourceName)
						}
						dashboardID = rs.Primary.ID

						layout, ok := s.RootModule().Resources[dashboardLayoutResourceName]
						if !ok {
							return fmt.Errorf(dashboardLayoutResourceNotFoundFmt, dashboardLayoutResourceName)
						}
						if _, err := fmt.Sscanf(layout.Primary.Attributes[dashboardLayoutFirstTileIDAttr], "%d", &tileID); err != nil {
							return fmt.Errorf("parsing tile_id: %w", err)
						}
						return nil
					},
				),
			},
			// Step 2: Inject color via API, then re-apply same config (no color)
			{
				PreConfig: func() {
					color := "blue"
					_, _, err := (&client).UpdateDashboardLayout(
						context.Background(), projectID, dashboardID,
						httpclient.DashboardLayoutPatchRequest{
							Tiles: []httpclient.DashboardTilePatchItem{
								{
									ID:    tileID,
									Color: &color,
								},
							},
						},
					)
					if err != nil {
						t.Fatalf("inject color via API: %v", err)
					}
				},
				Config: testAccDashboardLayoutSingleInsight(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// After apply, color should be cleared back to absent
					resource.TestCheckNoResourceAttr(dashboardLayoutResourceName, dashboardLayoutFirstTileColorAttr),
					// Verify via API that the color was actually cleared
					func(s *terraform.State) error {
						resp, _, err := (&client).GetDashboardLayout(
							context.Background(), projectID, dashboardID,
						)
						if err != nil {
							return fmt.Errorf("get dashboard layout: %w", err)
						}
						for _, tile := range resp.Tiles {
							if tile.ID == tileID {
								if tile.Color != nil && *tile.Color != "" {
									return fmt.Errorf("expected color to be cleared, got %q", *tile.Color)
								}
								return nil
							}
						}
						return fmt.Errorf("tile %d not found in API response", tileID)
					},
				),
			},
		},
	})
}

// testAccDashboardLayoutInsightWithColor returns HCL for a dashboard + insight + layout
// with a single insight tile that has the given color set.
func testAccDashboardLayoutInsightWithColor(name, color string) string {
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
      color        = %[2]q
      layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 6, h = 4 } })
    },
  ]

  depends_on = [posthog_insight.test]
}
`, name, color)
}

func testAccDashboardLayoutInsightWithShowDescription(name, showDescription string) string {
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
      insight_id       = posthog_insight.test.id
      show_description = %[2]s
      layouts_json     = jsonencode({ sm = { x = 0, y = 0, w = 6, h = 4 } })
    },
  ]

  depends_on = [posthog_insight.test]
}
`, name, showDescription)
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
