# Manage the flag yourself and link it, so destroying the feature
# does not leave a flag behind that PostHog created for it.
resource "posthog_feature_flag" "new_editor" {
  key                = "new-editor"
  name               = "New editor beta"
  rollout_percentage = 0
}

# Users who opt in from the feature previews list get the flag enabled
# while the stage is alpha, beta, or general-availability.
resource "posthog_early_access_feature" "new_editor" {
  name              = "New editor"
  stage             = "beta"
  description       = "A faster editor with live collaboration."
  documentation_url = "https://example.com/docs/new-editor"
  feature_flag_id   = posthog_feature_flag.new_editor.id
}
