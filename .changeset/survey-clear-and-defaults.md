---
"terraform-provider-posthog": patch
---

**`posthog_survey`:** fixes four cases that failed the apply with "Provider produced inconsistent result after apply". A survey that an earlier version left tainted after this error is replaced on the next apply unless you run `terraform untaint <address>` first.
  - Leaving `appearance_json` unset works again. PostHog adds a default thank-you message to new surveys, which is now read into state. Removing `appearance_json` from a config keeps the survey's current appearance; set it to `jsonencode({})` to clear it.
  - Removing `conditions_json`, `translations_json` or `form_content_json` now clears it in PostHog. If one of these was set in PostHog but not in your config, the plan already showed its removal, and applying that plan now succeeds in removing it.
  - `targeting_flag_filters_json` now applies. `targeting_flag_id` stays unset, the feature flag PostHog creates is in `targeting_flag_json`, and removing the filters deletes that flag. If the flag is deleted in PostHog, the next plan adds the filters back. Replacing `targeting_flag_id` with filters is rejected, because PostHog would overwrite the filters of the linked flag.
  - `response_sampling_*` attributes now apply without `response_sampling_daily_limits_json`, which PostHog calculates. Removing them clears the sampling start date too.
