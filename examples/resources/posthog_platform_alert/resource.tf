# Fire when a service logs more than 10 errors in a 10-minute window, with one
# group per service so each service fires and resolves on its own.
resource "posthog_platform_alert" "checkout_errors" {
  name        = "Checkout API errors"
  source_kind = "logs"

  # Noise reduction: only fire once 2 of the last 3 checks have breached the
  # threshold, then wait 30 minutes before notifying again.
  evaluation_periods  = 3
  datapoints_to_alarm = 2
  cooldown_minutes    = 30

  # Quiet hours, in the project's timezone.
  blocked_windows = [{
    start = "22:00"
    end   = "06:00"
  }]

  grouping = {
    mode = "by_result_labels"
    keys = ["service_name"]
  }

  logs = {
    filters = {
      severity_levels = ["error", "fatal"]
      service_names   = ["checkout-api", "payments-api"]
    }
    condition = {
      threshold_count    = 10
      threshold_operator = "above"
      window_minutes     = 10
    }
  }
}

# A logs alert keeps a logs alert row beside it, so destinations attach through
# legacy_configuration_id.
resource "posthog_logs_alert_destination" "checkout_errors_oncall" {
  alert_id    = posthog_platform_alert.checkout_errors.legacy_configuration_id
  type        = "webhook"
  webhook_url = "https://example.com/hooks/oncall"
}

# Move an alert from posthog_logs_alert with no destroy (Terraform 1.8 and later).
# Remove the posthog_logs_alert block in the same change.
moved {
  from = posthog_logs_alert.server_errors
  to   = posthog_platform_alert.server_errors
}

resource "posthog_platform_alert" "server_errors" {
  name        = "5xx responses"
  source_kind = "logs"

  logs = {
    filters = {
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
    }
    condition = {
      threshold_count = 50
      window_minutes  = 15
    }
  }
}
