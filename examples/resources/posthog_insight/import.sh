# Import using: project_id/insight_id
terraform import posthog_insight.example 12345/678

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_insight.example 678
