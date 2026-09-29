# Import using: project_id/warehouse_table_id
terraform import posthog_warehouse_table.example 12345/01900000-0000-7000-8000-000000000000

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_warehouse_table.example 01900000-0000-7000-8000-000000000000

# PostHog never returns the access key or secret, so the first apply after an
# import sends the configured values as an in-place update.
