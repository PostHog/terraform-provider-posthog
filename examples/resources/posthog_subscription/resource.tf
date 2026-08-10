# A dashboard and two insights to feed the digest.
resource "posthog_dashboard" "key_metrics" {
  name = "Key metrics"
}

resource "posthog_insight" "signups" {
  name          = "Signups"
  dashboard_ids = [posthog_dashboard.key_metrics.id]
  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{ kind = "EventsNode", event = "sign up", math = "total" }]
    }
  })
}

resource "posthog_insight" "revenue" {
  name          = "Revenue"
  dashboard_ids = [posthog_dashboard.key_metrics.id]
  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{ kind = "EventsNode", event = "purchase", math = "sum" }]
    }
  })
}

# Headline flow: a daily Slack dashboard digest.
# target_value is the canonical Slack form "<channelId>|#<channel-name>" and is stored
# verbatim. integration_id is required for Slack: connect Slack under Settings -> Integrations
# in the PostHog UI, then list this project's integrations via
# `GET /api/projects/<project_id>/integrations/` and use the id of the entry with kind "slack".
resource "posthog_subscription" "daily_slack_digest" {
  target_type               = "slack"
  target_value              = "C0B9A53J8RF|#reports"
  integration_id            = 1
  dashboard_id              = posthog_dashboard.key_metrics.id
  dashboard_export_insights = [posthog_insight.signups.id, posthog_insight.revenue.id]
  frequency                 = "daily"
  interval                  = 1
  start_date                = "2026-08-17T07:00:00Z"
  title                     = "Daily key metrics"
}

# A weekly email subscription for a single insight (integration_id omitted for email).
resource "posthog_subscription" "weekly_insight_email" {
  target_type  = "email"
  target_value = "team@example.com"
  insight_id   = posthog_insight.signups.id # mutually exclusive with dashboard_id
  frequency    = "weekly"
  interval     = 1
  byweekday    = ["monday"]
  start_date   = "2026-08-17T07:00:00Z"
}

# A monthly Slack digest on the first Monday of the month (rrule bysetpos).
resource "posthog_subscription" "monthly_first_monday" {
  target_type               = "slack"
  target_value              = "C0B9A53J8RF|#reports"
  integration_id            = 1
  dashboard_id              = posthog_dashboard.key_metrics.id
  dashboard_export_insights = [posthog_insight.revenue.id]
  frequency                 = "monthly"
  interval                  = 1
  byweekday                 = ["monday"]
  bysetpos                  = 1 # first Monday of the month
  start_date                = "2026-08-17T07:00:00Z"
  enabled                   = true
}

# A weekly AI-summary ("prompt") subscription: PostHog runs the natural-language prompt
# over your data and emails the result. No dashboard/insight subject; set ai_prompt instead.
resource "posthog_subscription" "weekly_prompt_digest" {
  target_type  = "email"
  target_value = "growth@example.com"
  ai_prompt    = "Top 5 events by volume this week, with counts and unique users for each."
  # Optional analysis window (rrule-independent); omit for the prompt's own default range.
  ai_prompt_config = jsonencode({
    window = { mode = "last_n_days", start_days_ago = 7, end_days_ago = null }
  })
  frequency  = "weekly"
  interval   = 1
  byweekday  = ["monday"]
  start_date = "2026-08-17T07:00:00Z"
  title      = "Weekly AI digest"
}
