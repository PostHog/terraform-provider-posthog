---
"terraform-provider-posthog": patch
---

**`posthog_survey`:** fixes four cases that failed the apply with "Provider produced inconsistent result after apply". A survey that an earlier version left tainted after this error is replaced on the next apply unless you run `terraform untaint <address>` first.
  - Leaving `appearance_json` unset works again. PostHog adds a default thank-you message to new surveys, which is now read into state. Removing `appearance_json` from a config keeps the survey's current appearance; change it to `jsonencode({})` to clear it.
  - Removing `conditions_json`, `translations_json` or `form_content_json` now clears it in PostHog. If one of these was set in PostHog but not in your config, the plan already showed its removal, and applying that plan now succeeds in removing it.
  - `targeting_flag_filters_json` now applies. `targeting_flag_id` stays unset, the feature flag PostHog creates is in `targeting_flag_json`, and removing the filters or setting them to `jsonencode({})` deletes that flag. If the survey loses its targeting flag in PostHog, the next plan adds the filters back. Filters are rejected when the survey targets a flag whose key does not start with `survey-targeting-`, unless `targeting_flag_id` keeps that flag, because PostHog would overwrite its filters.
  - `response_sampling_*` attributes now apply without `response_sampling_daily_limits_json`, which PostHog calculates. Removing them clears the sampling start date too.
