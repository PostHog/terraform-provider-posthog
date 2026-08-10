# Import using: project_id/subscription_id
terraform import posthog_subscription.daily_slack_digest 12345/678

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_subscription.daily_slack_digest 678
