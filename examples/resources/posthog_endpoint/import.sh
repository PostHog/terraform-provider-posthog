# Import using: project_id/endpoint_name
terraform import posthog_endpoint.example 12345/daily-signups

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_endpoint.example daily-signups
