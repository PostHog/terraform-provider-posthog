# Minimal dynamic cohort - PostHog recomputes membership from the filter rules
resource "posthog_cohort" "internal_users" {
  name = "Internal users"

  filters = jsonencode({
    properties = {
      type = "AND"
      values = [
        {
          type     = "person"
          key      = "email"
          operator = "icontains"
          value    = "@example.com"
        }
      ]
    }
  })
}

# Cohort combining person properties with a behavioural rule
resource "posthog_cohort" "engaged_customers" {
  name        = "Engaged customers"
  description = "Paying accounts that logged in during the last week"

  filters = jsonencode({
    properties = {
      type = "AND"
      values = [
        {
          type     = "person"
          key      = "plan"
          operator = "exact"
          value    = ["pro", "enterprise"]
        },
        {
          type          = "behavioral"
          key           = "$pageview"
          value         = "performed_event"
          event_type    = "events"
          time_value    = 7
          time_interval = "day"
        }
      ]
    }
  })
}

# Static cohort - the person list is populated outside Terraform
# (via the PostHog UI or the add_persons_to_static_cohort API)
resource "posthog_cohort" "beta_testers" {
  name        = "Beta testers"
  description = "Hand-picked accounts for the beta programme"
  is_static   = true
}

# Cohorts are commonly used to target a feature flag
resource "posthog_feature_flag" "beta_feature" {
  key    = "beta-feature"
  active = true

  filters = jsonencode({
    groups = [
      {
        properties = [
          {
            type  = "cohort"
            key   = "id"
            value = posthog_cohort.beta_testers.id
          }
        ]
        rollout_percentage = 100
      }
    ]
  })
}
