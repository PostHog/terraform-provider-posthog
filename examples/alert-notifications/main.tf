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

# ---------------------------------------------------------------------------
# Insight-alert Slack / webhook notifications (issue #130)
#
# When a `posthog_alert` fires, PostHog emits the internal event
# `$insight_alert_firing`. Routing that to Slack, a webhook, Discord, Teams,
# etc. is NOT a field on the alert API - it's a separate Hog function of
# `type = "internal_destination"` that subscribes to `$insight_alert_firing`
# and is filtered down to the firing alert's id.
#
# There is no dedicated resource: the EXISTING `posthog_hog_function` already
# models every field a notification needs, so you can wire alert notifications
# as code today. The chain is three resources:
#   1. posthog_insight   - the thing being measured
#   2. posthog_alert     - the threshold that fires on that insight
#   3. posthog_hog_function (internal_destination) - the delivery, filtered by
#      properties[alert_id] == <the alert's id> so it only fires for THIS alert
# ---------------------------------------------------------------------------

# 1. An insight to monitor.
resource "posthog_insight" "pageviews" {
  name = "Daily pageviews"

  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        name  = "$pageview"
        event = "$pageview"
        math  = "total"
      }]
    }
  })
}

# 2. An alert that fires when the insight crosses a threshold.
resource "posthog_alert" "pageviews_spike" {
  name                 = "Pageviews spike"
  insight              = posthog_insight.pageviews.id
  subscribed_users     = [] # in-app subscribers; delivery below is separate
  threshold_type       = "absolute"
  threshold_upper      = 10000
  condition_type       = "absolute_value"
  series_index         = 0
  calculation_interval = "daily"
}

# 3. Deliver the alert to Slack.
#
# Uses the built-in Slack destination template (`template-slack`). The
# `slack_workspace` input references an existing PostHog Slack integration by
# its numeric ID (find it under Data pipelines -> Integrations). The PostHog
# app must be installed in the workspace and, for private channels, be a member
# of the channel.
resource "posthog_hog_function" "pageviews_spike_slack" {
  name        = "Pageviews spike -> Slack"
  description = "Posts to Slack when the pageviews-spike alert fires"
  type        = "internal_destination"
  enabled     = true
  template_id = "template-slack"

  # Fire on the internal alert-firing event, scoped to THIS alert's id so
  # other alerts don't trigger this destination.
  filters_json = jsonencode({
    events = [{
      id   = "$insight_alert_firing"
      type = "events"
    }]
    properties = [{
      key      = "alert_id"
      value    = posthog_alert.pageviews_spike.id
      operator = "exact"
      type     = "event"
    }]
  })

  inputs_json = jsonencode({
    # Reference an existing Slack integration by its numeric ID.
    slack_workspace = {
      value = 1
    }
    # Channel ID (e.g. "C0123ABC") is preferred; "#channel-name" also works.
    channel = {
      value = "#alerts"
    }
    blocks = {
      value = [
        {
          type = "section"
          text = {
            type = "mrkdwn"
            text = ":rotating_light: *Alert firing:* {event.properties.alert_name}"
          }
        },
        {
          type = "actions"
          elements = [{
            type = "button"
            text = {
              type = "plain_text"
              text = "View insight in PostHog"
            }
            url = "{event.properties.insight_url}"
          }]
        }
      ]
      templating = "hog"
    }
  })

  depends_on = [posthog_alert.pageviews_spike]
}

# 3b. Webhook variant (uncomment to use instead of / alongside Slack).
#
# Uses the built-in HTTP Webhook template (`template-webhook`). Needs no
# external integration - just a URL - which makes it the most portable way to
# fan alerts out to PagerDuty, Opsgenie, a Lambda, etc.
#
# resource "posthog_hog_function" "pageviews_spike_webhook" {
#   name        = "Pageviews spike -> webhook"
#   description = "POSTs to a webhook when the pageviews-spike alert fires"
#   type        = "internal_destination"
#   enabled     = true
#   template_id = "template-webhook"
#
#   filters_json = jsonencode({
#     events = [{
#       id   = "$insight_alert_firing"
#       type = "events"
#     }]
#     properties = [{
#       key      = "alert_id"
#       value    = posthog_alert.pageviews_spike.id
#       operator = "exact"
#       type     = "event"
#     }]
#   })
#
#   inputs_json = jsonencode({
#     url = {
#       value      = "https://example.com/hooks/posthog-alert"
#       templating = "hog"
#     }
#     method = {
#       value = "POST"
#     }
#     body = {
#       value = {
#         alert_id   = "{event.properties.alert_id}"
#         alert_name = "{event.properties.alert_name}"
#         value      = "{event.properties.alert_calculated_value}"
#       }
#       templating = "hog"
#     }
#   })
#
#   depends_on = [posthog_alert.pageviews_spike]
# }
