# Actions Example
#
# This example demonstrates creating PostHog actions to track meaningful
# user behaviors by combining raw events (pageviews, autocapture, custom)
# into named actions that can be used in insights, cohorts, and alerts.
#
# Actions match on steps - each step defines a condition, and the action
# fires when ANY step matches (OR logic).

terraform {
  required_providers {
    posthog = {
      source = "posthog/posthog"
    }
  }
}

provider "posthog" {
  # Configuration can be provided via:
  # - Environment variables: POSTHOG_API_KEY, POSTHOG_PROJECT_ID, POSTHOG_HOST
  # - Or explicitly in the provider block:
  # api_key    = "your-api-key"
  # project_id = "12345"
  # host       = "https://us.posthog.com"  # Optional, defaults to US cloud
}

# =============================================================================
# Example 1: Simple custom event action
# =============================================================================

resource "posthog_action" "user_signed_up" {
  name        = "User Signed Up"
  description = "Tracks successful user registrations"
  tags        = ["activation", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      event = "user_signed_up"
    }
  ])
}

# =============================================================================
# Example 2: Pageview action with exact URL matching
# =============================================================================

resource "posthog_action" "visited_pricing" {
  name        = "Visited Pricing Page"
  description = "Tracks visits to the pricing page"
  tags        = ["marketing", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      event        = "$pageview"
      url          = "https://example.com/pricing"
      url_matching = "exact"
    }
  ])
}

# =============================================================================
# Example 3: Pageview action with regex URL matching
# =============================================================================

resource "posthog_action" "viewed_docs" {
  name = "Viewed Documentation"
  tags = ["product", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      event        = "$pageview"
      url          = "^/docs/.*"
      url_matching = "regex"
    }
  ])
}

# =============================================================================
# Example 4: Autocapture action with CSS selector
# =============================================================================

resource "posthog_action" "clicked_cta" {
  name = "Clicked CTA Button"
  tags = ["conversion", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      event    = "$autocapture"
      selector = "button.cta-primary"
    }
  ])
}

# =============================================================================
# Example 5: Autocapture action matching button text
# =============================================================================

resource "posthog_action" "clicked_upgrade" {
  name = "Clicked Upgrade"
  tags = ["conversion", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      event    = "$autocapture"
      tag_name = "button"
      text     = "Upgrade"
    }
  ])
}

# =============================================================================
# Example 6: Multi-step action combining different event types
# Any matching step triggers the action (OR logic)
# =============================================================================

resource "posthog_action" "checkout_interaction" {
  name        = "Checkout Interaction"
  description = "Tracks key interactions in the checkout flow"
  tags        = ["conversion", "checkout", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      # Step 1: User views the checkout page
      event        = "$pageview"
      url          = "/checkout"
      url_matching = "contains"
    },
    {
      # Step 2: User clicks "Place Order" button
      event    = "$autocapture"
      tag_name = "button"
      text     = "Place Order"
    },
    {
      # Step 3: Backend confirms purchase with amount > 0
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

# =============================================================================
# Example 7: Action with event property filters
# =============================================================================

resource "posthog_action" "mobile_signup" {
  name        = "Mobile Signup"
  description = "Signups originating from mobile devices"
  tags        = ["mobile", "activation", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      event = "user_signed_up"
      properties = [
        {
          key      = "$os"
          type     = "event"
          value    = ["iOS", "Android"]
          operator = "exact"
        }
      ]
    }
  ])
}

# =============================================================================
# Example 8: Action with href matching (link clicks)
# =============================================================================

resource "posthog_action" "clicked_external_link" {
  name = "Clicked External Link"
  tags = ["engagement", "managed-by:terraform"]

  steps_json = jsonencode([
    {
      event         = "$autocapture"
      tag_name      = "a"
      href          = "https://external-site.com"
      href_matching = "contains"
    }
  ])
}

# =============================================================================
# Outputs
# =============================================================================

output "action_ids" {
  description = "Created action IDs"
  value = {
    user_signed_up       = posthog_action.user_signed_up.id
    visited_pricing      = posthog_action.visited_pricing.id
    viewed_docs          = posthog_action.viewed_docs.id
    checkout_interaction = posthog_action.checkout_interaction.id
    mobile_signup        = posthog_action.mobile_signup.id
  }
}
