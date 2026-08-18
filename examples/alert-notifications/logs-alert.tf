# ---------------------------------------------------------------------------
# Log-alert Slack / webhook notifications
#
# Same shape as the insight-alert chain in main.tf, one event name different.
# When a `posthog_logs_alert` fires, PostHog emits the internal event
# `$logs_alert_firing`. Delivery is not a field on the alert: it is a separate
# Hog function of `type = "internal_destination"` subscribed to that event and
# filtered down to the firing alert's id.
#
# Destinations attached from the PostHog UI are hog functions too, built from
# the same templates. Declaring them here is what puts them under Terraform.
#   1. posthog_logs_alert   - the threshold on log volume
#   2. posthog_hog_function - the delivery, filtered by
#      properties[alert_id] == <the alert's id> so only THIS alert triggers it
# ---------------------------------------------------------------------------

# 1. An alert on checkout API errors.
resource "posthog_logs_alert" "checkout_errors" {
  name            = "Checkout API errors"
  severity_levels = ["error", "fatal"]
  service_names   = ["checkout-api"]

  threshold_count    = 10
  threshold_operator = "above"
  window_minutes     = 10

  # Only page once 2 of the last 3 checks breach, then stay quiet for 30 minutes.
  evaluation_periods  = 3
  datapoints_to_alarm = 2
  cooldown_minutes    = 30
}

# 2. Deliver it to Slack.
resource "posthog_hog_function" "checkout_errors_slack" {
  name        = "Checkout API errors -> Slack"
  description = "Posts to Slack when the checkout-errors log alert fires"
  type        = "internal_destination"
  enabled     = true
  template_id = "template-slack"

  filters_json = jsonencode({
    events = [{
      id   = "$logs_alert_firing"
      type = "events"
    }]
    properties = [{
      key      = "alert_id"
      value    = posthog_logs_alert.checkout_errors.id
      operator = "exact"
      type     = "event"
    }]
  })

  inputs_json = jsonencode({
    # Reference an existing Slack integration by its numeric ID.
    slack_workspace = {
      value = 1
    }
    channel = {
      value = "#alerts"
    }
    blocks = {
      value = [{
        type = "section"
        text = {
          type = "mrkdwn"
          text = ":rotating_light: *Log alert firing:* {event.properties.alert_name}"
        }
      }]
    }
    text = {
      value = "Log alert firing: {event.properties.alert_name}"
    }
  })
}

# 3. Or deliver it to a webhook. PostHog also emits `$logs_alert_resolved`,
#    `$logs_alert_broken` and `$logs_alert_errored`; subscribe to whichever
#    kinds you want by listing them all in `events`.
resource "posthog_hog_function" "checkout_errors_webhook" {
  name        = "Checkout API errors -> webhook"
  type        = "internal_destination"
  enabled     = true
  template_id = "template-webhook"

  filters_json = jsonencode({
    events = [
      { id = "$logs_alert_firing", type = "events" },
      { id = "$logs_alert_resolved", type = "events" },
    ]
    properties = [{
      key      = "alert_id"
      value    = posthog_logs_alert.checkout_errors.id
      operator = "exact"
      type     = "event"
    }]
  })

  inputs_json = jsonencode({
    url = {
      value = "https://example.com/hooks/posthog-logs-alert"
    }
    method = {
      value = "POST"
    }
    body = {
      value = {
        alert_id   = "{event.properties.alert_id}"
        alert_name = "{event.properties.alert_name}"
        result     = "{event.properties.result_count}"
        threshold  = "{event.properties.threshold_count}"
        logs_url   = "{project.url}/logs?{event.properties.logs_url_params}"
      }
    }
  })
}
