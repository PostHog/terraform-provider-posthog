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

## Recommended Terraform Contract

If implemented later, the resource should be modeled as an organization-scoped immutable resource:

- resource name: `posthog_proxy_record`
- import format: `<organization_id>/<proxy_record_uuid>`
- required `organization_id`
- required `domain` with `RequiresReplace`
- computed `target_cname`
- computed `status`
- computed `message`
- computed timestamps and creator metadata when useful

Recommended lifecycle behavior:

- `Create`: create the record, then poll `GET` until a terminal status is reached or a timeout expires.
- `Read`: refresh computed fields and drop state on `404`.
- `Update`: unsupported; domain changes should force replacement.
- `Delete`: delete the record, then poll `GET` until `404` or deletion timeout.

`retry` should not be part of the declarative CRUD contract. It is an operational remediation endpoint to be used after the user fixes DNS, not a stable desired-state transition.

## Why Acceptance Is The Real Blocker

The provider can only discover `target_cname` after `POST`. PostHog then expects the user-controlled domain to point at that generated CNAME target before the record can converge to `valid`.

That means a realistic acceptance test needs all of the following:

- a dedicated delegated test subdomain
- programmable DNS credentials for that zone
- test logic that can create a unique domain name per test run
- test logic that can create a matching CNAME to the returned `target_cname`
- a follow-up wait or explicit `retry` after DNS propagation

Without that harness, the provider can still create records, but tests cannot deterministically prove convergence. In practice, records drift into statuses like `timed_out`, which is exactly the failure mode we must avoid upstream.

## Recommendation

Do not implement `posthog_proxy_record` upstream yet.

First create a DNS-capable acceptance strategy, then implement the resource with:

- immutable create/delete semantics
- import support using `<organization_id>/<proxy_record_uuid>`
- async polling around create and delete
- explicit documentation that `target_cname` is provider-computed and must be wired into DNS outside Terraform unless the acceptance harness manages it

Once a DNS-backed test harness exists, the resource can be implemented cleanly and proposed upstream with confidence.
