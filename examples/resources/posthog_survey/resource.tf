# NPS-style feedback survey delivered via the in-product popover.
resource "posthog_survey" "nps" {
  name        = "PostHog NPS"
  description = "Quarterly NPS survey for PostHog users."
  type        = "popover"
  schedule    = "once"

  questions_json = jsonencode([
    {
      type          = "rating"
      question      = "How would you rate your PostHog experience?"
      scale         = 10
      buttonText    = "Submit"
      isNpsQuestion = true
    }
  ])

  conditions_json = jsonencode({
    url = "https://us.posthog.com/*"
  })
}

# Open-ended follow-up triggered only for users on a specific feature flag.
resource "posthog_survey" "feature_feedback" {
  name     = "New onboarding feedback"
  type     = "popover"
  schedule = "always"

  # Show this survey to users in the variant of an existing flag.
  linked_flag_id = 12345

  questions_json = jsonencode([
    {
      type        = "open"
      question    = "What was the most confusing part of the new onboarding?"
      description = "We read every response."
    }
  ])

  appearance_json = jsonencode({
    submitButtonText = "Send feedback"
  })
}
