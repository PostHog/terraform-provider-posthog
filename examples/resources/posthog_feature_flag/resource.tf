resource "posthog_feature_flag" "backend_rollout" {
  key                = "backend-rollout"
  name               = ""
  evaluation_runtime = "server"
  rollout_percentage = 50
}
