---
"terraform-provider-posthog": patch
---

**`posthog_logs_alert_destination`:** A destination's `id` is now its own id, `destinations[].id` in the logs alerts API, instead of its `hog_function_ids` sorted and joined by commas. PostHog is moving logs alert destinations off hog functions, and new destinations will have no `hog_function_ids` to track. Releases before this one cannot create destinations once that happens: PostHog refuses the create with a `terraform_provider_upgrade_required` error and writes nothing.
  - Existing destinations need no action. The next refresh finds each one by its old `id` and records the new one, without replacing it, and the plan shows no changes. A destination PostHog has not given an id yet keeps the old form until it has one.
  - `terraform import` now takes `project_id/alert_id/destination_id`, where `destination_id` is the destination's id or, as before, any one of its hog function ids.
