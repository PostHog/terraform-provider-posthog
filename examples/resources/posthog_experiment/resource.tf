# The backing flag is a separate posthog_feature_flag resource; the experiment references it by key.
resource "posthog_feature_flag" "pricing" {
  key = "pricing-page-test"
  filters = jsonencode({
    multivariate = { variants = [
      { key = "control", name = "Original", rollout_percentage = 50 },
      { key = "test", name = "Redesign", rollout_percentage = 50 },
    ] }
    groups = [{ properties = [], rollout_percentage = 100 }]
  })
}

# A running experiment with a metric, linked to the flag above.
resource "posthog_experiment" "pricing" {
  name             = "Pricing page test"
  description      = "Does the new pricing layout lift conversion?"
  feature_flag_key = posthog_feature_flag.pricing.key

  # Metrics and exposure are JSON-normalized: compared semantically, so key ordering and
  # server-computed fields (metric uuid/fingerprint) do not produce a diff.
  metrics = jsonencode([
    {
      kind        = "ExperimentMetric"
      metric_type = "mean"
      name        = "Revenue per user"
      source      = { kind = "EventsNode", event = "purchase", math = "sum" }
    }
  ])

  # filter_test_accounts lives inside exposure_criteria as filterTestAccounts (camelCase).
  exposure_criteria = jsonencode({ filterTestAccounts = true })

  status {
    state = "running"
  }
}

# Shipping a winning variant. ship_variant rewrites the flag's distribution to the winner, so tell
# the flag resource not to revert it with lifecycle.ignore_changes on filters.
resource "posthog_feature_flag" "onboarding" {
  key = "onboarding-checklist"
  filters = jsonencode({
    multivariate = { variants = [
      { key = "control", rollout_percentage = 50 },
      { key = "test", rollout_percentage = 50 },
    ] }
    groups = [{ properties = [], rollout_percentage = 100 }]
  })
  lifecycle {
    ignore_changes = [filters] # the experiment owns the live distribution once it ships
  }
}

resource "posthog_experiment" "onboarding" {
  name             = "Onboarding checklist"
  feature_flag_key = posthog_feature_flag.onboarding.key

  status {
    state = "stopped"
    stopped {
      ship_variant       = "test"
      conclusion         = "won"
      conclusion_comment = "New checklist lifted activation 8%"
    }
  }
}
