# Import using: project_id/property_group_uuid
terraform import posthog_schema_property_group.checkout 12345/your-property-group-uuid

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_schema_property_group.checkout your-property-group-uuid
