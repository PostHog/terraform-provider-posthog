# Dashboard Layout Example
#
# This example demonstrates the dashboard layout resource in PostHog:
# 1. Create a dashboard
# 2. Create insights to display on the dashboard
# 3. Arrange tiles (insights + text) with precise grid positions
# 4. Optionally: add color to highlight specific tiles
# 5. Optionally: update a text tile body in place
# 6. Optionally: add a new tile
# 7. Optionally: remove a tile and reposition

# =============================================================================
# Step 1: Create the dashboard
# =============================================================================

resource "posthog_dashboard" "demo" {
  name        = "Production Metrics"
  description = "Key metrics managed by Terraform"
}

# =============================================================================
# Step 2: Create insights
# =============================================================================

resource "posthog_insight" "pageviews" {
  name = "Pageview Trends"
  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$pageview"
        math  = "total"
      }]
    }
  })
  dashboard_ids = [posthog_dashboard.demo.id]
  depends_on    = [posthog_dashboard.demo]
}

resource "posthog_insight" "sessions" {
  name = "Session Trends"
  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$session_start"
        math  = "total"
      }]
    }
  })
  dashboard_ids = [posthog_dashboard.demo.id]
  depends_on    = [posthog_dashboard.demo]
}

# =============================================================================
# Step 3: Arrange tiles on the dashboard
#
# The layout is authoritative: any tiles not listed here will have their
# layouts cleared (insight tiles) or be soft-deleted (text tiles).
#
# Grid: 12 columns wide. Each tile needs x, y, w, h in at least the "sm"
# breakpoint.
# =============================================================================

resource "posthog_dashboard_layout" "demo" {
  dashboard_id = posthog_dashboard.demo.id

  tiles = [
    {
      text_body    = "## Production Metrics"
      layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
    },
    {
      insight_id   = posthog_insight.pageviews.id
      layouts_json = jsonencode({ sm = { x = 0, y = 1, w = 6, h = 5 } })
    },
    {
      insight_id   = posthog_insight.sessions.id
      layouts_json = jsonencode({ sm = { x = 6, y = 1, w = 6, h = 5 } })
    },
  ]

  depends_on = [posthog_insight.pageviews, posthog_insight.sessions]
}

# =============================================================================
# Step 4 (Optional): Add color to highlight a tile
# =============================================================================

# Uncomment to use:

# resource "posthog_dashboard_layout" "demo" {
#   dashboard_id = posthog_dashboard.demo.id
#
#   tiles = [
#     {
#       text_body    = "## Production Metrics"
#       layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
#     },
#     {
#       insight_id   = posthog_insight.pageviews.id
#       color        = "blue"
#       layouts_json = jsonencode({ sm = { x = 0, y = 1, w = 6, h = 5 } })
#     },
#     {
#       insight_id   = posthog_insight.sessions.id
#       layouts_json = jsonencode({ sm = { x = 6, y = 1, w = 6, h = 5 } })
#     },
#   ]
#
#   depends_on = [posthog_insight.pageviews, posthog_insight.sessions]
# }

# =============================================================================
# Step 5 (Optional): Update text tile body in place
#
# Changing a text tile's body updates the existing tile (matched by position).
# The tile_id stays stable — no destroy/recreate.
# =============================================================================

# Uncomment to use (comment out Steps 3-4 first):

# resource "posthog_dashboard_layout" "demo" {
#   dashboard_id = posthog_dashboard.demo.id
#
#   tiles = [
#     {
#       text_body    = "## Updated Production Metrics"
#       layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
#     },
#     {
#       insight_id   = posthog_insight.pageviews.id
#       layouts_json = jsonencode({ sm = { x = 0, y = 1, w = 6, h = 5 } })
#     },
#     {
#       insight_id   = posthog_insight.sessions.id
#       layouts_json = jsonencode({ sm = { x = 6, y = 1, w = 6, h = 5 } })
#     },
#   ]
#
#   depends_on = [posthog_insight.pageviews, posthog_insight.sessions]
# }

# =============================================================================
# Step 6 (Optional): Add a footer text tile
#
# Adding a tile increases the count. Existing tiles keep their tile_ids.
# =============================================================================

# Uncomment to use (comment out Steps 3-5 first):

# resource "posthog_dashboard_layout" "demo" {
#   dashboard_id = posthog_dashboard.demo.id
#
#   tiles = [
#     {
#       text_body    = "## Updated Production Metrics"
#       layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
#     },
#     {
#       insight_id   = posthog_insight.pageviews.id
#       layouts_json = jsonencode({ sm = { x = 0, y = 1, w = 6, h = 5 } })
#     },
#     {
#       insight_id   = posthog_insight.sessions.id
#       layouts_json = jsonencode({ sm = { x = 6, y = 1, w = 6, h = 5 } })
#     },
#     {
#       text_body    = "_Managed by Terraform_"
#       layouts_json = jsonencode({ sm = { x = 0, y = 6, w = 12, h = 1 } })
#     },
#   ]
#
#   depends_on = [posthog_insight.pageviews, posthog_insight.sessions]
# }

# =============================================================================
# Step 7 (Optional): Remove an insight and reposition
#
# Removing a tile from the list clears its layout (insight tiles) or
# soft-deletes it (text tiles). Remaining tiles can be repositioned freely.
# =============================================================================

# Uncomment to use (comment out Steps 3-6 first):

# resource "posthog_dashboard_layout" "demo" {
#   dashboard_id = posthog_dashboard.demo.id
#
#   tiles = [
#     {
#       text_body    = "## Updated Production Metrics"
#       layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
#     },
#     {
#       insight_id   = posthog_insight.pageviews.id
#       layouts_json = jsonencode({ sm = { x = 0, y = 1, w = 12, h = 5 } })
#     },
#   ]
#
#   depends_on = [posthog_insight.pageviews]
# }

# =============================================================================
# Outputs
# =============================================================================

output "dashboard" {
  description = "Created dashboard details"
  value = {
    id   = posthog_dashboard.demo.id
    name = posthog_dashboard.demo.name
  }
}

output "layout" {
  description = "Dashboard layout ID"
  value = {
    id           = posthog_dashboard_layout.demo.id
    dashboard_id = posthog_dashboard_layout.demo.dashboard_id
    tile_count   = length(posthog_dashboard_layout.demo.tiles)
  }
}
