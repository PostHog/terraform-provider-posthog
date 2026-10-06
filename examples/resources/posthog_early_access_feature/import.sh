# Import using: project_id/early_access_feature_id
terraform import posthog_early_access_feature.example 12345/01900000-0000-7000-8000-000000000000

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_early_access_feature.example 01900000-0000-7000-8000-000000000000
