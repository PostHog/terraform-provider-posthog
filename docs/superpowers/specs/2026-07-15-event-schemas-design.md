# Event Schemas support — design spec

Date: 2026-07-15
Status: approved (design), pending implementation plan
Context: [Issue #104](https://github.com/PostHog/terraform-provider-posthog/issues/104) — request for Event Schemas as Terraform resources.

## Goal

Let users declare PostHog event schemas as code: define named groups of typed
properties and attach them to events, so the schema contract lives in the same
repo as the code that emits the events.

Two new resources:

1. `posthog_schema_property_group` — a named group of typed properties.
2. `posthog_event_schema` — attaches a property group to an event definition.

## Background / API facts (verified against posthog@c10f860ef3a, 2026-07-14)

- Backing feature ("Schema Management", owner Sandy Spicer) is behind the
  `schema-management` flag, rolled out at 100% to all organizations since
  2025-12-03. API surface unchanged for several months — considered stable
  enough to build against.
- Endpoints (both project-scoped, UUID ids):
  - `GET/POST/PATCH/DELETE /api/projects/:id/schema_property_groups/`
  - `GET/POST/PATCH/DELETE /api/projects/:id/event_schemas/` — **no retrieve
    endpoint** (list/create/update/destroy mixins only).
- `SchemaPropertyGroupProperty`: `name`, `property_type`
  (`DateTime | String | Numeric | Boolean | Object`), `is_required`,
  `is_optional_in_types`, `description`. Names are unique per group and the
  server returns properties ordered by name (`Meta.ordering = ["name"]`).
- Property group PATCH treats the `properties` list as authoritative:
  entries without a server id are created, entries whose ids are missing are
  deleted, all in one transaction.
- `EventSchema` is a join record: `event_definition` (UUID) +
  `property_group_id` (write-only on create/update; responses embed the full
  `property_group`). Duplicate (event, group) pairs 400.
- `enforcement_mode` lives on the event definition, values `allow` (default)
  | `reject`. Setting `reject` is gated server-side on the
  `schema-enforcement-reject` flag, which does not exist yet — it 400s for
  every org today. **Out of scope for v1** (see Non-goals).

## Resource 1: `posthog_schema_property_group`

```hcl
resource "posthog_schema_property_group" "checkout" {
  name        = "Checkout"
  description = "Properties shared by checkout events"

  properties = [
    { name = "cart_value", property_type = "Numeric", is_required = true },
    { name = "currency",   property_type = "String" },
  ]
}
```

- Built on `core.NewGenericResource` with `core.BaseStringIdentifiable`
  (UUID id) + `core.BaseProjectID`, following `alert` / `external_data_source`.
- Attributes:
  - `name` — required string.
  - `description` — optional string.
  - `properties` — optional `SetNestedAttribute`. A set, not a list: the
    server sorts by name and enforces per-group name uniqueness, so order is
    not meaningful and a set avoids ordering phantom-diffs. Nested fields:
    - `name` — required.
    - `property_type` — required, enum-validated at plan time:
      `DateTime`, `String`, `Numeric`, `Boolean`, `Object`.
    - `is_required` — optional bool, default `false`.
    - `is_optional_in_types` — optional bool, default `false`.
    - `description` — optional string.
  - Computed: `id`, `created_at`, `updated_at`.
- Per-property server UUIDs are **not** tracked in state. Updates PATCH the
  full properties list without ids, which the API applies as a transactional
  replace-all. Nothing references property ids, so id churn is harmless, and
  this keeps nested computed-id phantom-diff handling out of the model.
- API 400s for duplicate group name / duplicate property name are surfaced
  as clear diagnostics.
- New client file: `internal/httpclient/schema_property_group.go`.

## Resource 2: `posthog_event_schema`

```hcl
resource "posthog_event_schema" "checkout_completed" {
  event             = "checkout_completed"
  property_group_id = posthog_schema_property_group.checkout.id
}
```

- Attributes:
  - `event` — required string, the event **name**.
  - `property_group_id` — required string (UUID of the group).
  - Computed: `id`, `event_definition_id`, `created_at`, `updated_at`.
- Name resolution: the provider resolves `event` to an event-definition UUID
  via `GET /api/projects/:id/event_definitions/?search=<name>` plus exact
  client-side match on `name`. Search results are ordered
  `length(name) ASC`, so an exact match appears on the first page; hidden
  events are included by default (`exclude_hidden` defaults to false).
  Unresolvable name → actionable error: the event must have been ingested at
  least once before a schema can be attached.
- Create: `POST /event_schemas/` with `event_definition` +
  `property_group_id`. Changing `event` or `property_group_id` updates in
  place via PATCH — no ForceNew.
- Read: no retrieve endpoint, so Read lists
  `GET /event_schemas/?event_definition=<uuid>` and finds our id; import
  (`project_id/resource_id`, the provider's standard format) falls back to
  an unfiltered list scan. Id missing from the list is treated as deleted
  (resource removed from state).
- Duplicate attachment (same event + group) 400s → clear diagnostic.
- New client file: `internal/httpclient/event_schema.go` (plus an
  event-definition list/search helper for name resolution).

## Error handling

- Name-resolution failure gets a purpose-built actionable error (event must
  have been ingested at least once). Duplicate-name and duplicate-attachment
  400s surface the API's already-descriptive validation message through the
  standard generic-resource diagnostic — no bespoke error mapping.
- Standard generic-resource handling covers 404-on-read (remove from state)
  and auth/permission errors.

## Testing

- Unit tests per Ops struct following existing `*_test.go` table patterns:
  request building (create/update, clear-on-update), response mapping,
  name-resolution error paths.
- Acceptance tests:
  - Property group: full lifecycle including add/remove/retype of
    properties, import.
  - Event schema: create, reattach to a different group, import. Test setup
    creates an event definition directly via the event-definitions API in a
    PreCheck/setup step (no new provider surface needed for that).

## Documentation

- Docs pages for both resources, generated like existing ones, with a
  combined example showing group → attachment chaining.
- Notes: schema management is currently observational (no ingestion-time
  rejection until PostHog enables enforcement); events must have been seen
  at least once before attachment.

## Non-goals (v2 candidates)

- `enforcement_mode` management (`posthog_event_definition_settings`) —
  deferred until PostHog enables the `schema-enforcement-reject` capability;
  today it could only ever write the server default.
- `posthog_event_definition` data source.
- Event-definition metadata (verified / hidden / tags / description).
