# A cohort referenced by the test_account_filters below (internal/test users).
resource "posthog_cohort" "internal" {
  name = "Internal users"
}

# Manage environment-level settings for a project.
# Any omitted attribute is left at PostHog's current value.
resource "posthog_project_settings" "example" {
  project_id = "your-project-id"

  heatmaps_opt_in               = true
  autocapture_exceptions_opt_in = true
  session_recording_opt_in      = true
  surveys_opt_in                = true
  autocapture_web_vitals_opt_in = false
  cookieless_server_hash_mode   = 0 # 0=disabled, 1=stateless, 2=stateful

  anonymize_ips = false

  # Authorized URLs / permitted domains (toolbar + project domain allowlist).
  app_urls = ["https://app.example.com", "https://www.example.com"]
  # Authorized domains for session replay.
  recording_domains = ["https://app.example.com"]

  # Capture network performance in session replay; required for the payload
  # capture settings below to have any effect.
  capture_performance_opt_in = true
  # Record network request headers/bodies in session replay. Both keys must be
  # set together. Careful: headers and bodies can contain tokens and PII.
  session_recording_network_payload_capture_config = {
    record_headers = true
    record_body    = false
  }

  # The "internal and test users" filter list every managed insight/hog function
  # references via filter_test_accounts. Managing it here makes that shared
  # definition owned by Terraform. A filter object may reference a cohort; PostHog
  # injects a read-only cohort_name which the provider strips, so no perpetual diff.
  test_account_filters = jsonencode([
    {
      key      = "id"
      type     = "cohort"
      value    = posthog_cohort.internal.id
      operator = "in"
    }
  ])
  # Apply the filters above by default across the project.
  test_account_filters_default_checked = true
}

# Use the provider-level project_id and manage only a subset of settings.
resource "posthog_project_settings" "minimal" {
  session_recording_opt_in = true
}
