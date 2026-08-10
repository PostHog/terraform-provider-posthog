# HogQL SQL insight loaded from a .sql file. The provider wraps the SQL in a
# DataVisualizationNode/HogQLQuery query for you. query_sql and query_json are
# mutually exclusive - exactly one must be set.
resource "posthog_insight" "active_users" {
  name        = "Active users (last 7 days)"
  description = "Daily active users over the past week"

  query_sql = file("${path.module}/queries/active_users.sql")
}

# The same, inline via a heredoc instead of a file.
resource "posthog_insight" "event_count" {
  name = "Total events"

  query_sql = <<-SQL
    SELECT count()
    FROM events
    WHERE timestamp >= now() - INTERVAL 7 DAY
  SQL
}

# For chart/column display options or non-HogQL insight types, use query_json.
resource "posthog_insight" "pageviews_trend" {
  name = "Pageviews trend"

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
}
