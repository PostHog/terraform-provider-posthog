---
"terraform-provider-posthog": patch
---

**`posthog_insight_variable`:** Compare `values_json` and `default_value_json` by their JSON content so formatting differences in the API response no longer cause inconsistent state errors.
