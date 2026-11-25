---
"terraform-provider-posthog": patch
---

**`posthog_hog_function`:** Creating a function from a template without `description` no longer fails with "Provider produced inconsistent result after apply"; the function gets the template's description. A function that an earlier version left tainted after this error is replaced on the next apply unless you run `terraform untaint <address>` first. ([#206](https://github.com/PostHog/terraform-provider-posthog/issues/206))
  - Leaving out `description` now keeps the description in PostHog. Removing it from the configuration used to clear it, and a description edited in PostHog used to show as a change. To clear it, set `description = ""`. An omitted description shows as `(known after apply)` on updates.
  - Creating a function from a template with `description = ""` still fails, because PostHog fills in the template's description. Create it without `description`, then set `""`. Replacing such a function fails the same way; run `terraform untaint <address>` and apply again to recover.
