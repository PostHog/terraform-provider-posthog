# Changelog

## Unreleased

### Features

- **New Resource:** `posthog_logs_alert` - Threshold-based alerting on logs, with severity, service, and attribute filters, repeat-breach thresholds, and quiet hours (notification destinations and snoozing are managed as code; see `examples/alert-notifications/logs-alert.tf`)
- **`posthog_alert`:** New `schedule_restriction` attribute for quiet hours - blocked local time windows during which the alert is not evaluated

### Internal

- Quiet-hours window validation is shared by `posthog_alert` and `posthog_logs_alert` (`internal/resource/core/quiethours.go`) instead of being implemented twice. Diagnostic wording is unified on "Quiet-hours ..." across both resources.

### Upgrade notes

- **`posthog_alert`:** quiet hours set outside Terraform on an alert this provider manages will show as a removal on the next plan, because the provider now sends `schedule_restriction` on every update. Add them to your configuration to keep them. Rarely, PostHog stores a shape it will not accept back: splitting an overnight window at midnight can leave a sub-30-minute piece, which then fails validation. Widen or drop that window.

## 1.0.0

### Features

- **New Resource:** `posthog_insight` - Create and manage insights with full query JSON support (Trends, Funnels, etc.)
- **New Resource:** `posthog_dashboard` - Manage dashboards with name, description, tags, and pinned status
- **New Resource:** `posthog_feature_flag` - Feature flags with filters, rollout percentages, multivariate variants, and payloads
- **New Resource:** `posthog_alert` - Threshold-based alerting on insights with configurable intervals and notifications
- **New Resource:** `posthog_hog_function` - Hog functions for destinations, webhooks, and transformations

### Provider Features

- Support for US and EU PostHog Cloud, plus self-hosted instances
- Automatic retry with exponential backoff for rate limits and transient errors
- Import support for all resources
