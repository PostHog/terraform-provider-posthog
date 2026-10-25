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
    groups   = [{ rollout_percentage = 100 }]
    payloads = { "true" = jsonencode({ theme = "blue" }) }
  })
}
