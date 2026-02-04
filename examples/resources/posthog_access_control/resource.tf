# --- Project Defaults (no role or organization_member) ---

# Set project-wide default: everyone can view surveys
resource "posthog_access_control" "surveys_project_default" {
  resource     = "survey"
  access_level = "viewer"
}

# Set default for a specific dashboard: everyone can view this dashboard
resource "posthog_access_control" "analytics_dashboard_default" {
  resource     = "dashboard"
  resource_id  = posthog_dashboard.analytics.id
  access_level = "viewer"
}

# --- Role-based Access ---

# Grant a role editor access to all feature flags in the project
resource "posthog_access_control" "engineering_feature_flags" {
  resource     = "feature_flag"
  access_level = "editor"
  role         = posthog_role.engineering.id
}

# Grant a role viewer access to all dashboards
resource "posthog_access_control" "support_dashboards" {
  resource     = "dashboard"
  access_level = "viewer"
  role         = posthog_role.support.id
}

# Grant a role viewer access to a specific dashboard
resource "posthog_access_control" "support_analytics_dashboard" {
  resource     = "dashboard"
  resource_id  = posthog_dashboard.analytics.id
  access_level = "viewer"
  role         = posthog_role.support.id
}

# Explicitly deny a role access to experiments (access_level = "none")
resource "posthog_access_control" "support_no_experiments" {
  resource     = "experiment"
  access_level = "none"
  role         = posthog_role.support.id
}

# --- User-specific Access ---

# Grant a specific user editor access to a specific dashboard
resource "posthog_access_control" "alice_analytics_dashboard" {
  resource            = "dashboard"
  resource_id         = posthog_dashboard.analytics.id
  access_level        = "editor"
  organization_member = posthog_organization_member.alice.id
}
