---
"terraform-provider-posthog": patch
---

**`posthog_logs_alert_destination`: upgrade required.** PostHog is moving logs alert destinations off hog functions. Once that ships, PostHog refuses destination creates from releases before this one with a `terraform_provider_upgrade_required` error and writes nothing. Upgrade to this release to keep creating logs alert destinations.
  - A destination's `id` is now its own id, `destinations[].id` in the logs alerts API, instead of its `hog_function_ids` sorted and joined by commas.
  - Existing destinations need no config change. The next refresh finds each one by its old `id` and records the new one, without replacing it, and the plan shows no changes. A destination PostHog has not given an id yet keeps the old form until it has one.
  - `terraform import` now takes `project_id/alert_id/destination_id`, where `destination_id` is the destination's id or, as before, any one of its hog function ids.
