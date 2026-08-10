# Complete Experiment (A/B test) Example
#
# This walks through the full lifecycle of an experiment the way you'd actually run one.
# Terraform is declarative, so you don't run these as separate commands — you edit the
# `status` block on the experiment and re-apply. The stages below show what that config
# looks like at each point in the journey:
#
#   Stage 1 — Draft:    create the backing flag + the experiment, still a draft (not live).
#   Stage 2 — Running:  attach metrics and launch it to start collecting data.  <-- live state below
#   Stage 3 — Complete: stop the experiment and ship the winning variant.
#
# The committed config below sits at Stage 2 (running with metrics) — the state an
# experiment spends most of its life in. Stages 1 and 3 are shown as inline alternatives
# on the `status` block so you can see the whole progression in one place.

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
# The backing feature flag
#
# An experiment runs on an existing multivariate feature flag — the flag owns the
# variant split (here a 50/50 control/test), the experiment references it by key.
# One variant MUST be keyed "control".
#
# `lifecycle.ignore_changes = [filters]` matters for Stage 3: shipping a winner
# rewrites the flag's distribution (winner -> 100%). Without this, the flag resource
# would revert that change on the next apply and fight the experiment. Set it up front
# so completing the experiment later is drift-free.
# =============================================================================

resource "posthog_feature_flag" "onboarding_checklist" {
  key = "onboarding-checklist-test"
  filters = jsonencode({
    multivariate = { variants = [
      { key = "control", name = "No checklist", rollout_percentage = 50 },
      { key = "test", name = "Guided checklist", rollout_percentage = 50 },
    ] }
    groups = [{ properties = [], rollout_percentage = 100 }]
  })

  lifecycle {
    ignore_changes = [filters] # the experiment owns the live distribution once it ships
  }
}

# -----------------------------------------------------------------------------
# Alternative: link a flag by bare key (no posthog_feature_flag resource)
#
# feature_flag_key is just a string, so you can point at a flag you don't manage
# in Terraform — one you created in the PostHog UI, or a brand-new key that
# PostHog auto-creates a default control/test flag for. Then you drop the flag
# resource above entirely and reference the key directly:
#
#   resource "posthog_experiment" "onboarding_checklist" {
#     name             = "Onboarding checklist"
#     feature_flag_key = "onboarding-checklist-test"  # bare key, not a resource ref
#     status { state = "draft" }
#   }
#
# Trade-offs vs. the managed flag above:
#   + Simpler to ship: since Terraform doesn't own the flag, nothing reverts the
#     shipped distribution — no lifecycle.ignore_changes needed.
#   - The flag must already be multivariate with a "control" variant (or you
#     accept PostHog's auto-created default split — you don't control the keys).
#   - On `terraform destroy` the experiment is removed but the flag is left in
#     place (correct for a UI-managed flag; an orphan for an auto-created one).
# -----------------------------------------------------------------------------

# =============================================================================
# The experiment
#
# Metrics and exposure criteria are JSON. They're compared semantically, so key
# ordering and server-computed fields (metric uuid/fingerprint) never produce a diff.
# You typically add these when you launch (Stage 2) — an idea worth testing is worth
# defining a metric for.
# =============================================================================

resource "posthog_experiment" "onboarding_checklist" {
  name             = "Onboarding checklist"
  description      = "Does a guided setup checklist lift new-user activation?"
  feature_flag_key = posthog_feature_flag.onboarding_checklist.key

  # Primary metric: did the guided checklist move activation?
  metrics = jsonencode([
    {
      kind        = "ExperimentMetric"
      metric_type = "mean"
      name        = "Activated within 7 days"
      source      = { kind = "EventsNode", event = "user_activated", math = "total" }
    }
  ])

  # A secondary funnel metric: how far new users get through setup.
  metrics_secondary = jsonencode([
    {
      kind        = "ExperimentMetric"
      metric_type = "funnel"
      name        = "Setup funnel"
      series = [
        { kind = "EventsNode", event = "signed_up" },
        { kind = "EventsNode", event = "completed_setup" },
      ]
    }
  ])

  # filter_test_accounts lives inside exposure_criteria as filterTestAccounts (camelCase).
  exposure_criteria = jsonencode({ filterTestAccounts = true })

  # ---------------------------------------------------------------------------
  # Lifecycle: edit this `status` block and re-apply to move the experiment along.
  # ---------------------------------------------------------------------------

  # Stage 1 — Draft (created but not live). Start here:
  #   status {
  #     state = "draft"
  #   }

  # Stage 2 — Running (live, collecting data). The current committed state:
  status {
    state = "running"
  }

  # Stage 3 — Complete and ship the winner. Swap the block above for this once you
  # have a result. `ship_variant` rolls the winning variant out to 100% on the flag
  # (which is why the flag ignores `filters` changes above):
  #   status {
  #     state = "stopped"
  #     stopped {
  #       ship_variant       = "test"
  #       conclusion         = "won"
  #       conclusion_comment = "Guided checklist lifted 7-day activation by 8%."
  #     }
  #   }
  #
  # To stop WITHOUT shipping a variant (e.g. an inconclusive result), omit ship_variant:
  #   status {
  #     state = "stopped"
  #     stopped {
  #       conclusion = "inconclusive"
  #     }
  #   }
}
