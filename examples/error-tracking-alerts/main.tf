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

# ---------------------------------------------------------------------------
# Error tracking alerts (issue #131)
#
# PostHog error-tracking alerts are just Hog functions of
# `type = "internal_destination"` that subscribe to one of PostHog's internal
# error-tracking events and forward it to a destination (Slack, webhook, ...).
#
# There is no dedicated resource: the EXISTING `posthog_hog_function` resource
# already models every field an alert needs, so you can manage error-tracking
# alerts as code today.
#
# The internal events you can subscribe to:
#   - $error_tracking_issue_created   (a brand-new issue was seen)
#   - $error_tracking_issue_reopened  (a resolved issue started erroring again)
#   - $error_tracking_issue_spiking   (an issue's volume spiked)
#
# The alert is wired up with three pieces:
#   1. type        = "internal_destination"
#   2. filters_json subscribing to the internal event above
#   3. template_id + inputs_json selecting and configuring the destination
# ---------------------------------------------------------------------------

# Example 1: Slack alert on newly created error-tracking issues.
#
# Uses the built-in Slack destination template (`template-slack`). The
# `slack_workspace` input references an existing PostHog Slack integration by
# its numeric ID (find it under Data pipelines -> Sources / the integrations
# API). The PostHog app must be installed in the workspace and, for private
# channels, be a member of the channel.
resource "posthog_hog_function" "error_issue_created_slack" {
  name        = "Error tracking: new issue -> Slack"
  description = "Posts to Slack whenever a new error-tracking issue is created"
  type        = "internal_destination"
  enabled     = true
  template_id = "template-slack"

  # Subscribe to the internal error-tracking event. Swap the id to
  # "$error_tracking_issue_reopened" or "$error_tracking_issue_spiking" to
  # alert on those lifecycle events instead.
  filters_json = jsonencode({
    events = [{
      id   = "$error_tracking_issue_created"
      type = "events"
    }]
  })

  inputs_json = jsonencode({
    # Reference an existing Slack integration by its numeric ID.
    slack_workspace = {
      value = 1
    }
    # Channel ID (e.g. "C0123ABC") is preferred; "#channel-name" also works.
    channel = {
      value = "#error-tracking"
    }
    blocks = {
      value = [
        {
          type = "section"
          text = {
            type = "mrkdwn"
            text = ":rotating_light: *New error-tracking issue:* {event.properties.name}"
          }
        },
        {
          type = "actions"
          elements = [{
            type = "button"
            text = {
              type = "plain_text"
              text = "View issue in PostHog"
            }
            url = "{event.properties.issue_url}"
          }]
        }
      ]
      templating = "hog"
    }
  })
}

# Example 2: Generic webhook alert on newly created error-tracking issues.
#
# Uses the built-in HTTP Webhook template (`template-webhook`). This variant
# needs no external integration - just a URL - which makes it the most portable
# way to fan alerts out to PagerDuty, Opsgenie, a Lambda, etc.
resource "posthog_hog_function" "error_issue_created_webhook" {
  name        = "Error tracking: new issue -> webhook"
  description = "POSTs to a webhook whenever a new error-tracking issue is created"
  type        = "internal_destination"
  enabled     = true
  template_id = "template-webhook"

  filters_json = jsonencode({
    events = [{
      id   = "$error_tracking_issue_created"
      type = "events"
    }]
  })

  inputs_json = jsonencode({
    url = {
      value      = "https://example.com/hooks/posthog-error-tracking"
      templating = "hog"
    }
    method = {
      value = "POST"
    }
    body = {
      value = {
        issue_id = "{event.properties.issue_id}"
        name     = "{event.properties.name}"
        event    = "{event.event}"
      }
      templating = "hog"
    }
  })
}
