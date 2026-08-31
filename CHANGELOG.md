# Changelog

## Unreleased

### Features

- **New Resource:** `posthog_logs_alert` - Threshold-based alerting on logs, with severity, service, and attribute filters, repeat-breach thresholds, and quiet hours (snoozing and notification destinations are managed as code). Window length and count limits are enforced by PostHog and reported on apply; the provider only rejects configurations PostHog would silently reshape.
- **New Resource:** `posthog_logs_alert_destination` - Slack, webhook, and Microsoft Teams notification destinations for a `posthog_logs_alert`. PostHog has no update endpoint for destinations, so every attribute forces replacement. `slack_channel_name` is write-only: PostHog uses it to label the destination and never stores it, so it is never read back and an imported destination has it unset.
- **`posthog_alert`:** New `schedule_restriction` attribute for quiet hours - blocked local time windows during which the alert is not evaluated. Window length and count limits are enforced by PostHog and reported on apply; the provider only rejects configurations PostHog would silently reshape.

### Fixes

- **HTTP retries:** a read that fails by timing out is now retried. The 30-second timeout was set on the HTTP client, which bounds the whole retry loop including the waits between attempts, so the first retry of a timed-out request was cancelled before it began and the apply failed with `(attempt: 1) sleeping: context deadline exceeded`. Each attempt now gets its own 30-second budget instead. `POST` and `PATCH` are the exception, for the reason in the next entry: a timeout cannot tell you whether PostHog already applied the write, so they are not sent again. ([#156](https://github.com/PostHog/terraform-provider-posthog/issues/156))
- **HTTP retries:** `POST` and `PATCH` requests are no longer replayed after a 500, 502, 503, 504, a timeout, or a dropped connection. None of those failures say whether PostHog had already committed the write, so retrying a create could leave behind a second resource that Terraform never records and will not clean up. They are still retried on 429, which PostHog returns before processing the request at all. `GET`, `HEAD`, `PUT`, and `DELETE` are safe to send twice and are unaffected, as is any request whose connection was never established, since nothing was sent. ([#156](https://github.com/PostHog/terraform-provider-posthog/issues/156))
- **HTTP retries:** a failed request now says which deadline expired and how many attempts it took, instead of a bare `context deadline exceeded` that was indistinguishable from cancelling the run yourself. `TF_LOG=DEBUG` also records each retry and the reason the provider stopped, so a request that was deliberately not replayed can be told apart from one that ran out of attempts. ([#156](https://github.com/PostHog/terraform-provider-posthog/issues/156))

### Internal

- `NewRetryTransport` fills in any unset retry tunable, so a hand-built `RetryConfig` cannot end up with no request deadline or a zero backoff that retries with no spacing. `JitterFactor` and `MaxRetries` are left alone, because zero is a meaningful value for both.
- The `WithNoRetry` and `WithRetryConfig` client options replace an installed retry transport instead of stacking a second one on it. Stacked, the outer transport's per-attempt deadline had to cover the inner transport's whole retry loop, and `WithNoRetry` did not actually disable retries.
- Quiet-hours window validation is shared by `posthog_alert` and `posthog_logs_alert` (`internal/resource/core/quiethours.go`) instead of being implemented twice. Diagnostic wording is unified on "Quiet-hours ..." across both resources, so `posthog_alert`'s messages change text but not meaning.

### Upgrade notes

- An unresponsive PostHog endpoint now takes up to about two minutes to fail a read, rather than 30 seconds: the four attempts each get the full 30-second budget that previously covered all of them together. This is the same ceiling the three configured retries were always meant to have. It is per request, so a read that pages through a long list can take that long per page, and a host that redirects spends it again on each hop.
- **`posthog_alert`:** quiet hours set outside Terraform on an alert this provider manages will show as a removal on the next plan, because the provider now sends `schedule_restriction` on every update. Add them to your configuration to keep them. Rarely, PostHog stores a shape it will not accept back: splitting an overnight window at midnight can leave a piece shorter than its own minimum, which it then refuses on apply. Widen or drop that window.

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
