# Import using: project_id/logs_alert_uuid
terraform import posthog_logs_alert.example 12345/your-logs-alert-uuid

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_logs_alert.example your-logs-alert-uuid
