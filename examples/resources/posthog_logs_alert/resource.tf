# Fire when a service logs more than 10 errors in a 10-minute window.
resource "posthog_logs_alert" "checkout_errors" {
  name            = "Checkout API errors"
  severity_levels = ["error", "fatal"]
  service_names   = ["checkout-api"]

  threshold_count    = 10
  threshold_operator = "above"
  window_minutes     = 10

  # Noise reduction: only fire once 2 of the last 3 check periods have breached
  # the threshold, then wait 30 minutes before notifying again.
  evaluation_periods  = 3
  datapoints_to_alarm = 2
  cooldown_minutes    = 30

  # Quiet hours, in the project's timezone. Each window must span at least 30
  # minutes, and windows must not overlap or touch. A window may cross midnight
  # only when it is the only window. PostHog stores blocked windows on a single
  # merged 24-hour timeline, so anything it would reshape is rejected at plan time.
  blocked_windows = [{
    start = "22:00"
    end   = "06:00"
  }]
}

# Alert on a log attribute rather than just severity and service. Use
# `filter_group_json` for anything the typed attributes do not cover.
resource "posthog_logs_alert" "server_errors" {
  name = "5xx responses"

  filter_group_json = jsonencode({
    type = "AND"
    values = [{
      type = "AND"
      values = [{
        type     = "log_attribute"
        key      = "status_code"
        operator = "exact"
        value    = ["500", "502", "503"]
      }]
    }]
  })

  threshold_count = 50
  window_minutes  = 15
}

# A draft alert. Filters are only optional while `enabled = false`.
resource "posthog_logs_alert" "draft" {
  name    = "Work in progress"
  enabled = false
}

# An alert on its own evaluates but notifies nobody. Delivery is a separate hog
# function subscribed to the internal event PostHog emits when the alert fires,
# filtered to this alert's id so no other alert triggers it. This is the same
# chain insight alerts use; see examples/alert-notifications/ for Slack and
# webhook variants side by side.
resource "posthog_hog_function" "checkout_errors_slack" {
  name        = "Checkout API errors -> Slack"
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
    slack_workspace = { value = 1 }
    channel         = { value = "#alerts" }
    text            = { value = "Log alert firing: {event.properties.alert_name}" }
  })
}
