# Import using: project_id/alert_id/destination_id
#
# destination_id is the destination's id, `destinations[].id` on the alert in the logs alerts
# API. A destination PostHog built from hog functions can also be imported by any one of
# them; Terraform adopts the whole destination and records its id.
terraform import posthog_logs_alert_destination.example 12345/your-logs-alert-uuid/your-destination-uuid

# slack_channel_name is write-only, so an imported Slack destination has it unset. Add it to
# your configuration to name the destination in the PostHog UI, which replaces it.

# PostHog redacts webhook_url on read, so imported webhook and teams destinations have it
# unset. Add the real URL to configuration after import; the first plan replaces the destination.
