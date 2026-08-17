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

  # Quiet hours, in the project's timezone. A window may cross midnight and must
  # span at least 30 minutes. Windows must not overlap — PostHog merges overlapping
  # windows on save, so the provider rejects them at plan time.
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
