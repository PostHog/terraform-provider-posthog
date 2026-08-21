# Import using: project_id/alert_id/hog_function_id
#
# A destination has no ID of its own: PostHog builds it as a group of hog functions, one per
# alert transition. Name any one of them and Terraform adopts the whole group. List them
# in the PostHog UI or through the project's generic Hog Functions list API.
terraform import posthog_logs_alert_destination.example 12345/your-logs-alert-uuid/your-hog-function-uuid

# slack_channel_name is write-only, so an imported Slack destination has it unset. Add it to
# your configuration to name the destination in the PostHog UI, which replaces it.

# PostHog returns webhook_url through the generic Hog Function API, so webhook and teams
# destinations are adopted in place. Terraform stores the imported URL as sensitive state.
