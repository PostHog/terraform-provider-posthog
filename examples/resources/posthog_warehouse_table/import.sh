# Import using: project_id/warehouse_table_id
# PostHog never returns the bucket credential. After import, set access_key and
# access_secret in the configuration and run apply to store them in state.
terraform import posthog_warehouse_table.example 12345/01900000-0000-7000-8000-000000000000

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_warehouse_table.example 01900000-0000-7000-8000-000000000000
