# A running two-variant experiment with a metric.
# The backing feature flag "pricing-page-test" is auto-created and owned by the experiment.
resource "posthog_experiment" "pricing" {
  name             = "Pricing page test"
  description      = "Does the new pricing layout lift conversion?"
  feature_flag_key = "pricing-page-test"

  variant {
    key                = "control"
    name               = "Original"
    rollout_percentage = 50
  }
  variant {
    key                = "test"
    name               = "Redesign"
    rollout_percentage = 50
  }

  # Metrics and exposure are JSON-normalized: compared semantically, so key ordering and
  # server-computed fields (metric uuid/fingerprint) do not produce a diff.
  metrics = jsonencode([
    {
      kind        = "ExperimentMetric"
      metric_type = "mean"
      name        = "Revenue per user"
      source = {
        kind  = "EventsNode"
        event = "purchase"
        math  = "sum"
      }
    }
  ])

  # filter_test_accounts lives inside the exposure_criteria blob as filterTestAccounts (camelCase),
  # not a top-level field.
  exposure_criteria = jsonencode({
    filterTestAccounts = true
  })

  status {
    state = "running"
  }
}

# A stopped experiment shipping the winning variant to everyone.
resource "posthog_experiment" "onboarding" {
  name             = "Onboarding checklist"
  feature_flag_key = "onboarding-checklist"

  variant {
    key                = "control"
    rollout_percentage = 50
  }
  variant {
    key                = "test"
    rollout_percentage = 50
  }

  status {
    state = "stopped"
    stopped {
      ship_variant        = "test"
      release_to_everyone = false # distribution-only: preserves the flag's release conditions
      conclusion          = "won"
      conclusion_comment  = "New checklist lifted activation 8%"
    }
  }
}

# Linking an existing feature flag instead of auto-creating one: omit the variant blocks and
# point feature_flag_key at a flag managed by a posthog_feature_flag resource. The variant split
# is owned by that flag.
resource "posthog_feature_flag" "checkout" {
  key = "checkout-flow"
  filters = jsonencode({
    multivariate = { variants = [
      { key = "control", rollout_percentage = 50 },
      { key = "test", rollout_percentage = 50 },
    ] }
    groups = [{ properties = [], rollout_percentage = 100 }]
  })
}

resource "posthog_experiment" "checkout" {
  name             = "Checkout flow test"
  feature_flag_key = posthog_feature_flag.checkout.key # link the existing flag
  status { state = "draft" }
}
