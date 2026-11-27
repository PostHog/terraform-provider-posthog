---
"terraform-provider-posthog": minor
---

**`posthog_endpoint`:** New resource that manages PostHog endpoints, which turn a HogQL or insight query into a named, versioned API route.
  - The endpoint name is its id and part of its run URL. PostHog cannot rename an endpoint, so changing `name` replaces it. Import by name: `terraform import posthog_endpoint.example <project_id>/<name>`.
  - Changing `query_json` updates the endpoint in place, and PostHog saves the query as a new version. `current_version` shows the latest version.
  - A query references a SQL variable as `{variables.code_name}` and declares it under `variables`, keyed by the `posthog_insight_variable` id. The resource docs show an example.
  - `is_materialized` reports whether materialization is turned on, not whether its first run has finished. A plan that materializes an inactive endpoint, or sets `bucket_overrides` without materialization, fails before any write.
