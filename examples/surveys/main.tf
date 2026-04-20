terraform {
  required_providers {
    posthog = {
      source = "posthog/posthog"
    }
  }
}

provider "posthog" {}

resource "posthog_survey" "customer_feedback" {
  name        = "Customer feedback survey"
  description = "Basic in-product feedback survey managed by Terraform"
  type        = "popover"
  schedule    = "once"

  questions_json = jsonencode([
    {
      type          = "rating"
      question      = "How satisfied are you with Ansyo?"
      scale         = 10
      buttonText    = "Submit"
      isNpsQuestion = true
    }
  ])

  conditions_json = jsonencode({
    url = "https://app.ansyo.co/*"
  })
}

output "survey_id" {
  description = "UUID of the survey"
  value       = posthog_survey.customer_feedback.id
}
