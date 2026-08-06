resource "posthog_schema_property_group" "checkout" {
  name = "Checkout"

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
      is_required   = true
    },
  ]
}

# The event must have been ingested at least once so its definition exists.
resource "posthog_event_schema" "checkout_completed" {
  event             = "checkout_completed"
  property_group_id = posthog_schema_property_group.checkout.id
}
