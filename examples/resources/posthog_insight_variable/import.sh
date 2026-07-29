# Import using: project_id/insight_variable_id
terraform import posthog_insight_variable.example 12345/019fae99-842f-0000-f983-cc2171e27686

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_insight_variable.example 019fae99-842f-0000-f983-cc2171e27686
