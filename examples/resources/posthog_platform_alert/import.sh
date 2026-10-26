# Import using: project_id/platform_alert_uuid
terraform import posthog_platform_alert.example 12345/your-platform-alert-uuid

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_platform_alert.example your-platform-alert-uuid
