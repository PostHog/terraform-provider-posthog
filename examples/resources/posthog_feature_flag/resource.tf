resource "posthog_feature_flag" "backend_rollout" {
  key                 = "backend-rollout"
  name                = ""
  evaluation_runtime  = "server"
  evaluation_contexts = ["api"]
  rollout_percentage  = 50
}

resource "posthog_feature_flag" "signup_flow" {
  key                          = "signup-flow"
  bucketing_identifier         = "device_id"
  ensure_experience_continuity = false
  rollout_percentage           = 25
}

resource "posthog_feature_flag" "retired" {
  key      = "retired-checkout"
  active   = false
  archived = true
}

resource "posthog_feature_flag" "app_configuration" {
  key                     = "app-configuration"
  is_remote_configuration = true
  active                  = true
  filters = jsonencode({
    groups = [{ rollout_percentage = 100 }]
    # Payloads are keyed by the value the flag returns: "true" for a boolean flag
    # like this one. Each payload is a JSON string, so it is jsonencoded on its own.
    payloads = { "true" = jsonencode({ theme = "blue" }) }
  })
}

# A multivariate flag returns a variant key instead of true, so its payloads are
# keyed by variant. The variants' rollout percentages must add up to 100.
resource "posthog_feature_flag" "checkout_button" {
  key = "checkout-button"
  filters = jsonencode({
    groups = [{ rollout_percentage = 100 }]
    multivariate = {
      variants = [
        { key = "control", rollout_percentage = 50 },
        { key = "green", rollout_percentage = 50 },
      ]
    }
    payloads = {
      control = jsonencode({ color = "blue", label = "Buy now" })
      green   = jsonencode({ color = "green", label = "Buy now" })
    }
  })
}
