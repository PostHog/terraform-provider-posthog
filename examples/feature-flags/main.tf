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

# Example 1: Simple Boolean Feature Flag
# A basic on/off toggle that's enabled for all users
resource "posthog_feature_flag" "simple_boolean" {
  key    = "simple-boolean-flag"
  name   = "Simple Boolean Flag"
  active = true

  filters = jsonencode({
    groups = [{
      properties         = []
      rollout_percentage = 100
    }]
  })

  tags = ["managed-by:terraform"]
}

# Example 2: Gradual Rollout with Percentage
# Roll out a feature to 25% of users using the convenience field
resource "posthog_feature_flag" "gradual_rollout" {
  key                = "gradual-rollout-25-percent"
  name               = "Gradual Rollout 25%"
  active             = true
  rollout_percentage = 25
  tags               = ["managed-by:terraform"]
}

# Example 3: Multivariate Feature Flag (A/B/C Test)
# Test multiple variants with different rollout percentages
resource "posthog_feature_flag" "multivariate_test" {
  key    = "multivariate-abc-test"
  name   = "Multivariate A/B/C Test"
  active = true

  filters = jsonencode({
    groups = [{
      properties         = []
      rollout_percentage = 100
    }]
    multivariate = {
      variants = [
        {
          key                = "control"
          name               = "Control"
          rollout_percentage = 34
        },
        {
          key                = "variant-a"
          name               = "Variant A"
          rollout_percentage = 33
        },
        {
          key                = "variant-b"
          name               = "Variant B"
          rollout_percentage = 33
        }
      ]
    }
  })

  tags = ["managed-by:terraform"]
}

# Example 4: Feature Flag with User Property Targeting
# Enable for users with specific properties (email domain and plan type)
resource "posthog_feature_flag" "user_property_targeting" {
  key    = "user-property-targeting"
  name   = "User Property Targeting"
  active = true

  filters = jsonencode({
    groups = [
      {
        # Group 1: Target enterprise customers
        properties = [
          {
            key      = "plan"
            type     = "person"
            value    = ["enterprise", "business"]
            operator = "exact"
          }
        ]
        rollout_percentage = 100
      },
      {
        # Group 2: Gradual rollout to all other users
        properties         = []
        rollout_percentage = 10
      }
    ]
  })

  tags = ["managed-by:terraform"]
}

# Example 5: Feature Flag Dependency
# Enable a feature only for users who have another feature flag enabled
resource "posthog_feature_flag" "flag_dependency" {
  key    = "flag-dependency-example"
  name   = "Flag Dependency Example"
  active = true

  filters = jsonencode({
    groups = [{
      properties = [{
        key      = "$feature/${posthog_feature_flag.simple_boolean.key}"
        type     = "person"
        value    = ["true"]
        operator = "exact"
      }]
      rollout_percentage = 100
    }]
  })

  tags = ["managed-by:terraform"]
}

# Example 6: Remote Configuration Flag
# Feature flag with a JSON payload for dynamic configuration
resource "posthog_feature_flag" "remote_config" {
  key    = "remote-config-payload"
  name   = "Remote Config with Payload"
  active = true

  filters = jsonencode({
    groups = [{
      properties         = []
      rollout_percentage = 100
    }]
    payloads = {
      "true" = jsonencode({
        api_endpoint = "https://api.example.com/v2"
        timeout_ms   = 5000
        max_retries  = 3
        features = {
          caching_enabled = true
          rate_limit      = 1000
        }
      })
    }
  })

  tags = ["managed-by:terraform"]
}

# Outputs to reference the created feature flags
output "simple_boolean_flag_id" {
  description = "ID of the simple boolean feature flag"
  value       = posthog_feature_flag.simple_boolean.id
}

output "multivariate_test_flag_id" {
  description = "ID of the multivariate test feature flag"
  value       = posthog_feature_flag.multivariate_test.id
}
