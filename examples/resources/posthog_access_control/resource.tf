# Grant a role editor access to all feature flags in the project
resource "posthog_access_control" "engineering_feature_flags" {
  resource     = "feature_flags"
  access_level = "editor"
  role         = posthog_role.engineering.id
}

# Grant a role viewer access to all dashboards
resource "posthog_access_control" "support_dashboards" {
  resource     = "dashboards"
  access_level = "viewer"
  role         = posthog_role.support.id
}

# Grant a role viewer access to a specific dashboard
resource "posthog_access_control" "support_analytics_dashboard" {
  resource     = "dashboards"
  resource_id  = posthog_dashboard.analytics.id
  access_level = "viewer"
  role         = posthog_role.support.id
}

# Grant a specific user editor access to a specific dashboard
resource "posthog_access_control" "alice_analytics_dashboard" {
  resource            = "dashboards"
  resource_id         = posthog_dashboard.analytics.id
  access_level        = "editor"
  organization_member = posthog_organization_member.alice.id
}

# Explicitly deny a role access to experiments (access_level = "none")
resource "posthog_access_control" "support_no_experiments" {
  resource     = "experiments"
  access_level = "none"
  role         = posthog_role.support.id
}
