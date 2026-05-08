# Import using: project_id/external_data_source_id
terraform import posthog_external_data_source.example 12345/01900000-0000-7000-8000-000000000000

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_external_data_source.example 01900000-0000-7000-8000-000000000000
