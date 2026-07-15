# Import using: project_id/event_schema_uuid
terraform import posthog_event_schema.checkout_completed 12345/your-event-schema-uuid

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_event_schema.checkout_completed your-event-schema-uuid
