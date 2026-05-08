# Create the dashboard
resource "posthog_dashboard" "engineering" {
  name = "Engineering Metrics"
}

# Create insights and attach them to the dashboard
resource "posthog_insight" "pageviews" {
  name = "Page Views"
  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{ kind = "EventsNode", event = "$pageview", math = "total" }]
    }
  })
  dashboard_ids = [posthog_dashboard.engineering.id]
  depends_on    = [posthog_dashboard.engineering]
}

resource "posthog_insight" "signups" {
  name = "Sign-ups"
  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{ kind = "EventsNode", event = "user_signed_up", math = "total" }]
    }
  })
  dashboard_ids = [posthog_dashboard.engineering.id]
  depends_on    = [posthog_dashboard.engineering]
}

# Manage the dashboard layout: insight tiles + a text tile header
resource "posthog_dashboard_layout" "engineering" {
  dashboard_id = posthog_dashboard.engineering.id

  tiles = [
    # Text tile acting as a section header
    {
      text_body    = "## Key Metrics"
      layouts_json = jsonencode({ sm = { x = 0, y = 0, w = 12, h = 1 } })
    },
    # Insight tiles positioned below the header
    {
      insight_id       = posthog_insight.pageviews.id
      layouts_json     = jsonencode({ sm = { x = 0, y = 1, w = 6, h = 4 } })
      color            = "blue"
      show_description = false
    },
    {
      insight_id       = posthog_insight.signups.id
      layouts_json     = jsonencode({ sm = { x = 6, y = 1, w = 6, h = 4 } })
      show_description = false
    },
  ]

  depends_on = [posthog_insight.pageviews, posthog_insight.signups]
}

# Using a for expression to generate tiles from a list of objects
variable "insight_tiles" {
  description = "List of insight tiles with their layout positions"
  type = list(object({
    insight_id = number
    layouts    = map(object({ x = number, y = number, w = number, h = number }))
  }))
  default = [
    {
      insight_id = 12345
      layouts    = { sm = { x = 0, y = 0, w = 6, h = 4 } }
    },
    {
      insight_id = 67890
      layouts    = { sm = { x = 6, y = 0, w = 6, h = 4 } }
    },
  ]
}

resource "posthog_dashboard_layout" "dynamic" {
  dashboard_id = posthog_dashboard.engineering.id

  tiles = [for tile in var.insight_tiles : {
    insight_id   = tile.insight_id
    layouts_json = jsonencode(tile.layouts)
  }]
}
