resource "posthog_schema_property_group" "checkout" {
  name        = "Checkout"
  description = "Properties shared by checkout events"

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
      is_required   = true
    },
    {
      name          = "currency"
      property_type = "String"
    },
  ]
}
