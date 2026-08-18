resource "posthog_logs_alert" "checkout_errors" {
  name            = "Checkout API errors"
  severity_levels = ["error", "fatal"]
  service_names   = ["checkout-api"]

  threshold_count = 10
  window_minutes  = 10
}

# Post into Slack. slack_workspace_id is the PostHog Slack integration, created by
# connecting Slack to the project in the PostHog UI. The channel is addressed by ID rather
# than by name so that renaming it in Slack cannot silently repoint the alert.
resource "posthog_logs_alert_destination" "checkout_errors_slack" {
  alert_id           = posthog_logs_alert.checkout_errors.id
  type               = "slack"
  slack_workspace_id = 1
  slack_channel_id   = "C0123456789"

  # Display name only, used to label the destination in the PostHog UI. PostHog never
  # stores it, so Terraform cannot read it back and will not report drift on it.
  slack_channel_name = "#checkout-alerts"
}

# POST the notification to an arbitrary endpoint, such as an on-call router.
resource "posthog_logs_alert_destination" "checkout_errors_oncall" {
  alert_id    = posthog_logs_alert.checkout_errors.id
  type        = "webhook"
  webhook_url = "https://example.com/hooks/oncall"
}

# Microsoft Teams takes the same incoming webhook URL, under its own type.
resource "posthog_logs_alert_destination" "checkout_errors_teams" {
  alert_id    = posthog_logs_alert.checkout_errors.id
  type        = "teams"
  webhook_url = "https://example.webhook.office.com/webhookb2/00000000-0000-0000-0000-000000000000"
}

# An alert can hold several destinations. Each one is a separate resource, and changing any
# attribute of one replaces it: PostHog has no update endpoint for destinations.
