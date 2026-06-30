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

  # Authorized URLs / permitted domains (toolbar + project domain allowlist).
  app_urls = ["https://app.example.com", "https://www.example.com"]
  # Authorized domains for session replay.
  recording_domains = ["https://app.example.com"]
}

# Use the provider-level project_id and manage only a subset of settings.
resource "posthog_project_settings" "minimal" {
  session_recording_opt_in = true
}
