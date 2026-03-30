# Import using: project_id/alert_uuid
terraform import posthog_alert.example 12345/your-alert-uuid

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_alert.example your-alert-uuid
