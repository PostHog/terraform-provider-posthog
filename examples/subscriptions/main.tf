# Complete Subscriptions Example
#
# This example demonstrates the three ways PostHog delivers a recurring digest:
# 1. A daily Slack digest of a dashboard (with selected insights)
# 2. A weekly email subscription for a single insight
# 3. A weekly AI-summary ("prompt") subscription (no dashboard/insight subject)
#
# Each subscription targets exactly one subject: dashboard_id, insight_id, or ai_prompt.

terraform {
  required_providers {
    posthog = {
      source = "posthog/posthog"
    }
  }
}

provider "posthog" {
  # Configuration can be provided via:
  # - Environment variables: POSTHOG_API_KEY, POSTHOG_PROJECT_ID, POSTHOG_HOST
  # - Or explicitly in the provider block:
  # api_key    = "your-api-key"
  # project_id = "12345"
  # host       = "https://us.posthog.com"  # Optional, defaults to US cloud
}

variable "slack_integration_id" {
  # Connect Slack under Settings -> Integrations in the PostHog UI, then list this project's
  # integrations via `GET /api/projects/<project_id>/integrations/` and use the id of the
  # entry whose kind is "slack".
  description = "ID of the Slack integration to deliver through (kind=slack in GET /api/projects/<project_id>/integrations/)"
  type        = number
  default     = 1
}

# =============================================================================
# Subjects: a dashboard with two insights to feed the digests
# =============================================================================

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

# =============================================================================
# 1. Daily Slack digest of a dashboard
# integration_id is required for Slack; target_value is the canonical
# "<channelId>|#<channel-name>" form and is stored verbatim.
# dashboard_export_insights selects which of the dashboard's insights to include.
# =============================================================================

resource "posthog_subscription" "daily_slack_dashboard" {
  target_type               = "slack"
  target_value              = "C0B9A53J8RF|#reports"
  integration_id            = var.slack_integration_id
  dashboard_id              = posthog_dashboard.key_metrics.id
  dashboard_export_insights = [posthog_insight.signups.id, posthog_insight.revenue.id]
  frequency                 = "daily"
  interval                  = 1
  start_date                = "2026-08-17T07:00:00Z"
  title                     = "Daily key metrics"
}

# =============================================================================
# 2. Weekly email subscription for a single insight
# integration_id is omitted for email; insight_id is mutually exclusive with dashboard_id.
# =============================================================================

resource "posthog_subscription" "weekly_email_insight" {
  target_type  = "email"
  target_value = "team@example.com"
  insight_id   = posthog_insight.signups.id
  frequency    = "weekly"
  interval     = 1
  byweekday    = ["monday"]
  start_date   = "2026-08-17T07:00:00Z"
  title        = "Weekly signups"
}

# =============================================================================
# 3. Weekly AI-summary ("prompt") subscription
# No dashboard/insight subject: set ai_prompt instead. PostHog runs the prompt over
# your project data. ai_prompt_config optionally scopes the analysis window; the
# summary_* options steer an AI-generated summary of the result.
# =============================================================================

resource "posthog_subscription" "weekly_prompt_digest" {
  target_type  = "email"
  target_value = "growth@example.com"
  ai_prompt    = "Top 5 events by volume this week, with counts and unique users for each."
  ai_prompt_config = jsonencode({
    window = { mode = "last_n_days", start_days_ago = 7, end_days_ago = null }
  })
  summary_enabled      = true
  summary_prompt_guide = "Focus on week-over-week growth and call out anomalies."
  frequency            = "weekly"
  interval             = 1
  byweekday            = ["monday"]
  start_date           = "2026-08-17T07:00:00Z"
  title                = "Weekly AI digest"
}

# =============================================================================
# Outputs
# =============================================================================

output "subscriptions" {
  description = "Created subscription IDs by flow"
  value = {
    daily_slack_dashboard = posthog_subscription.daily_slack_dashboard.id
    weekly_email_insight  = posthog_subscription.weekly_email_insight.id
    weekly_prompt_digest  = posthog_subscription.weekly_prompt_digest.id
  }
}
