# Proxy Record Design

## Decision

Defer upstream implementation of `posthog_proxy_record` until the provider has a DNS-backed acceptance strategy that can reliably drive the PostHog custom-domain lifecycle to `valid`.

The API surface is sufficient to model a Terraform resource, but it is not sufficient to validate that resource upstream with the current provider test harness. Shipping the resource without that validation path would make the provider brittle around the most important behavior: the asynchronous transition from initial creation to an operational domain.

## Observed API Contract

The organization-scoped API currently exposes:

- `GET /api/organizations/{organization_id}/proxy_records/`
- `POST /api/organizations/{organization_id}/proxy_records/`
- `GET /api/organizations/{organization_id}/proxy_records/{id}/`
- `DELETE /api/organizations/{organization_id}/proxy_records/{id}/`
- `POST /api/organizations/{organization_id}/proxy_records/{id}/retry/`

The list endpoint returns an envelope with `max_proxy_records` and `results[]`.
Each record includes at least:

- `id`
- `domain`
- `target_cname`
- `status`
- `message`
- `created_at`
- `updated_at`
- `created_by`

Observed statuses in the Ansyo test organization include:

- `valid`
- `timed_out`

There is no `PUT` or `PATCH` endpoint, so the API should be treated as create/delete plus asynchronous server-side reconciliation.

## Implemented Terraform Contract

The provider resource is modeled as an organization-scoped immutable resource:

- resource name: `posthog_proxy_record`
- import format: `<organization_id>/<proxy_record_uuid>`
- required `organization_id`
- required `domain` with `RequiresReplace`
- computed `target_cname`
- computed `status`
- computed `message`
- computed timestamps and creator metadata when useful

Recommended lifecycle behavior:

- `Create`: create the record and persist the immediately returned `target_cname` plus the current `status`.
- `Read`: refresh computed fields and drop state on `404`.
- `Update`: unsupported; domain changes should force replacement.
- `Delete`: delete the record and rely on `404` handling for any later refresh.

`retry` should not be part of the declarative CRUD contract. It is an operational remediation endpoint to be used after the user fixes DNS, not a stable desired-state transition.
The live API probe on `2026-04-21` also showed that `POST /retry/` returns `403 permission_denied` when called with a Personal API Key, so it cannot be a reliable part of provider CRUD with the current authentication model.

## Why Acceptance Is The Real Blocker

The provider can only discover `target_cname` after `POST`. PostHog then expects the user-controlled domain to point at that generated CNAME target before the record can converge to `valid`.

That means a realistic acceptance test needs all of the following:

- a dedicated delegated test subdomain
- programmable DNS credentials for that zone
- test logic that can create a unique domain name per test run
- test logic that can create a matching CNAME to the returned `target_cname`
- a follow-up wait for PostHog to observe the new DNS state

The provider `Create` path must not wait for `valid`, because that would deadlock a Terraform graph where another DNS provider resource depends on the computed `target_cname` emitted by `posthog_proxy_record`.

## Recommendation

- immutable create/delete semantics
- import support using `<organization_id>/<proxy_record_uuid>`
- read-time status refresh only
- explicit documentation that `target_cname` is provider-computed and must be wired into DNS outside Terraform unless the acceptance harness manages it

Acceptance should keep the DNS orchestration outside the resource itself: create the PostHog record, create the matching CNAME via the dedicated harness, then wait until `GET` reports `valid`.
