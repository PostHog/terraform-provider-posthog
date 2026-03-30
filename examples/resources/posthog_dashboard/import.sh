# Import using: project_id/dashboard_id
terraform import posthog_dashboard.example 12345/678

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_dashboard.example 678
