---
"terraform-provider-posthog": minor
---

**`posthog_project_settings`:** New `events_retention_months` attribute sets how many months of events a project keeps, overriding the organization's period.
  - Only organization admins can change it. A shorter period permanently deletes older events at PostHog's next monthly cleanup, so a plan that may shorten it shows a warning, including the first apply and a period only known at apply.
  - Terraform only sends the period when it differs from PostHog's value. Removing the attribute leaves the period as it is; to return a project to the organization's period, clear it in PostHog's project settings.
  - Every refresh of `posthog_project_settings` now also reads the period from PostHog's events retention endpoint. PostHog versions without that endpoint read it as null, and setting it there fails before any write.
  - PostHog checks the period when it's set (it currently accepts `1`, `3`, `12`, `24`, `36`, `60` and `84`), so a period PostHog adds later works without a provider update.
  - If someone changes the period in PostHog after a plan is made, applying that plan stops before any write and asks for a new plan, so it never deletes events the plan didn't warn about.
