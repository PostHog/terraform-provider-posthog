# Changelog

## Unreleased

### Features

- **New Resource:** `posthog_logs_alert` - Threshold-based alerting on logs, with severity, service, and attribute filters, repeat-breach thresholds, and quiet hours (notification destinations and snoozing are managed as code; see `examples/alert-notifications/logs-alert.tf`). Window length and count limits are enforced by PostHog and reported on apply; the provider only rejects configurations PostHog would silently reshape.

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
