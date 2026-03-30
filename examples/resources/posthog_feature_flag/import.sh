# Import using: project_id/feature_flag_id
terraform import posthog_feature_flag.example 12345/678

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_feature_flag.example 678
