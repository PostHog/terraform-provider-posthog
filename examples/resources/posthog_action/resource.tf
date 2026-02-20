# Minimal action - only requires a name
resource "posthog_action" "simple" {
  name = "User Signed Up"
}

# Action with tags and description
resource "posthog_action" "pageview_tracking" {
  name        = "Visited Pricing Page"
  description = "Tracks when a user visits the pricing page"
  tags        = ["marketing", "conversion"]

  steps_json = jsonencode([
    {
      event        = "$pageview"
      url          = "https://example.com/pricing"
      url_matching = "exact"
    }
  ])
}

# Action with multiple steps matching different user interactions
# Any matching step triggers the action (OR logic)
resource "posthog_action" "checkout_funnel" {
  name        = "Checkout Interaction"
  description = "Tracks key interactions in the checkout flow"
  tags        = ["conversion", "checkout"]

  steps_json = jsonencode([
    {
      event        = "$pageview"
      url          = "/checkout"
      url_matching = "contains"
    },
    {
      event    = "$autocapture"
      tag_name = "button"
      text     = "Place Order"
    },
    {
      event = "purchase_completed"
      properties = [
        {
          key      = "amount"
          type     = "event"
          value    = "0"
          operator = "gt"
        }
      ]
    }
  ])
}

# Autocapture action matching a CSS selector
resource "posthog_action" "cta_click" {
  name = "Clicked CTA Button"

  steps_json = jsonencode([
    {
      event    = "$autocapture"
      selector = "button.cta-primary"
    }
  ])
}
