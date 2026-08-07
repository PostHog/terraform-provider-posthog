# Changelog

## Unreleased

### Features

- **New Resource:** `posthog_experiment` - Manage experiments (A/B tests) with typed variant blocks, JSON-normalized `metrics`/`metrics_secondary`/`exposure_criteria`, and the full draft → running → paused → stopped lifecycle (including shipping a winning variant on stop) via a `status` block. Provide `variant` blocks to auto-create the backing feature flag, or omit them to link an existing multivariate flag (e.g. `feature_flag_key = posthog_feature_flag.foo.key`)

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
