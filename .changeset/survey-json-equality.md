---
"terraform-provider-posthog": patch
---

**`posthog_survey`:** `questions_json`, `conditions_json`, `appearance_json`, `translations_json`, `form_content_json` and `response_sampling_daily_limits_json` are now compared by content, so key order, whitespace, escaping and number formatting (`2.0` and `2`) no longer fail the apply with "Provider produced inconsistent result after apply". Before, only JSON formatted like `jsonencode` output worked; a heredoc, a file, or Pulumi Python's `json.dumps` did not. A survey that an earlier version left tainted after this error is replaced on the next apply unless you run `terraform untaint <address>` first.
  - JSON that doesn't parse, including `""`, is now rejected by `terraform validate` and `terraform plan` instead of failing during the apply.
