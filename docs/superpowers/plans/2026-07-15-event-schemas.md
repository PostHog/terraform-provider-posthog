# Event Schemas Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add `posthog_schema_property_group` and `posthog_event_schema` resources so users can declare PostHog event schemas as code (spec: `docs/superpowers/specs/2026-07-15-event-schemas-design.md`).

**Architecture:** Both resources ride the existing `core.NewGenericResource` framework (Ops struct + TF model + httpclient methods). Property groups are plain CRUD against `/api/projects/:id/schema_property_groups/`. Event schemas are a join record against `/api/projects/:id/event_schemas/` with two quirks: the API has **no retrieve endpoint** (Read = list + find by id) and the resource takes an event **name** which the provider resolves to an event-definition UUID at Create/Update time.

**Tech Stack:** Go, terraform-plugin-framework v1.19.0, terraform-plugin-testing (acceptance), testify + httptest (unit).

## Global Constraints

- API paths use the **projects** prefix: `/api/projects/%s/...` (NOT `/api/environments/` — that's the alerts prefix, don't copy it).
- `property_type` enum (exact strings, verified against PostHog `SchemaPropertyType`): `DateTime`, `String`, `Numeric`, `Boolean`, `Object`.
- Import ID format is `project_id/resource_id` via `core.ProjectScopedImportParser` (slash, not colon).
- Conventional commits (`feat:`, `test:`, `docs:`); NO `Co-Authored-By` lines, ever.
- Gates before every push: `make fmt`, `make lint`, `make test`, `make build`. `make generate` after schema/example changes.
- Before final push: `grep -rn` the repo for every new identifier to catch stale references (including `docs/` and `examples/`).
- Module path is `github.com/posthog/terraform-provider` (no `-posthog` suffix).

---

### Task 1: httpclient — schema property groups

**Files:**
- Create: `internal/httpclient/schema_property_group.go`
- Test: `internal/httpclient/schema_property_group_test.go`

**Interfaces:**
- Consumes: `doPost[T]`, `doGet[T]`, `doPatch[T]`, `doDelete` from `internal/httpclient/client.go`.
- Produces (used by Task 3):
  - `type SchemaPropertyGroupProperty struct` (fields below)
  - `type SchemaPropertyGroup struct` (fields below)
  - `type SchemaPropertyGroupRequest struct` (fields below)
  - `CreateSchemaPropertyGroup(ctx, projectID string, input SchemaPropertyGroupRequest) (SchemaPropertyGroup, error)`
  - `GetSchemaPropertyGroup(ctx, projectID, id string) (SchemaPropertyGroup, HTTPStatusCode, error)`
  - `UpdateSchemaPropertyGroup(ctx, projectID, id string, input SchemaPropertyGroupRequest) (SchemaPropertyGroup, HTTPStatusCode, error)`
  - `DeleteSchemaPropertyGroup(ctx, projectID, id string) (HTTPStatusCode, error)`

- [ ] **Step 1: Write the failing test**

`internal/httpclient/schema_property_group_test.go`:

```go
package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSPGProjectID      = "project-123"
	testSPGID             = "0196-aaaa"
	testSPGCollectionPath = "/api/projects/project-123/schema_property_groups/"
	testSPGResourcePath   = "/api/projects/project-123/schema_property_groups/0196-aaaa/"
)

func TestSchemaPropertyGroupCRUD(t *testing.T) {
	var lastCreateBody SchemaPropertyGroupRequest

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == testSPGCollectionPath:
			require.NoError(t, json.NewDecoder(r.Body).Decode(&lastCreateBody))
			_ = json.NewEncoder(w).Encode(SchemaPropertyGroup{ID: testSPGID, Name: "Checkout"})
		case r.Method == http.MethodGet && r.URL.Path == testSPGResourcePath:
			_ = json.NewEncoder(w).Encode(SchemaPropertyGroup{ID: testSPGID, Name: "Checkout"})
		case r.Method == http.MethodPatch && r.URL.Path == testSPGResourcePath:
			_ = json.NewEncoder(w).Encode(SchemaPropertyGroup{ID: testSPGID, Name: "Checkout v2"})
		case r.Method == http.MethodDelete && r.URL.Path == testSPGResourcePath:
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	ctx := context.Background()

	name := "Checkout"
	required := true
	props := []SchemaPropertyGroupProperty{{Name: "cart_value", PropertyType: "Numeric", IsRequired: &required}}
	created, err := client.CreateSchemaPropertyGroup(ctx, testSPGProjectID, SchemaPropertyGroupRequest{Name: &name, Properties: &props})
	require.NoError(t, err)
	assert.Equal(t, testSPGID, created.ID)
	require.NotNil(t, lastCreateBody.Properties)
	assert.Equal(t, "cart_value", (*lastCreateBody.Properties)[0].Name)
	assert.Equal(t, "Numeric", (*lastCreateBody.Properties)[0].PropertyType)

	got, status, err := client.GetSchemaPropertyGroup(ctx, testSPGProjectID, testSPGID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSPGID, got.ID)

	updated, status, err := client.UpdateSchemaPropertyGroup(ctx, testSPGProjectID, testSPGID, SchemaPropertyGroupRequest{Name: &name})
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, "Checkout v2", updated.Name)

	status, err = client.DeleteSchemaPropertyGroup(ctx, testSPGProjectID, testSPGID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNoContent), status)
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./internal/httpclient/ -run TestSchemaPropertyGroupCRUD -v`
Expected: FAIL to compile — `undefined: SchemaPropertyGroupRequest` etc.

- [ ] **Step 3: Write minimal implementation**

`internal/httpclient/schema_property_group.go`:

```go
package httpclient

import (
	"context"
	"fmt"
)

// SchemaPropertyGroupProperty is one typed property inside a schema property group.
// ID is server-assigned; requests omit it, which makes the API treat updates as
// replace-all (existing properties whose ids are absent get deleted, incoming
// entries without ids get created — transactionally).
type SchemaPropertyGroupProperty struct {
	ID                *string `json:"id,omitempty"`
	Name              string  `json:"name"`
	PropertyType      string  `json:"property_type"`
	IsRequired        *bool   `json:"is_required,omitempty"`
	IsOptionalInTypes *bool   `json:"is_optional_in_types,omitempty"`
	Description       *string `json:"description,omitempty"`
}

type SchemaPropertyGroup struct {
	ID          string                        `json:"id"`
	Name        string                        `json:"name"`
	Description *string                       `json:"description,omitempty"`
	Properties  []SchemaPropertyGroupProperty `json:"properties,omitempty"`
	CreatedAt   *string                       `json:"created_at,omitempty"`
	UpdatedAt   *string                       `json:"updated_at,omitempty"`
}

// SchemaPropertyGroupRequest uses a pointer-to-slice for Properties so callers
// can distinguish "leave untouched" (nil) from "delete all" (empty slice).
type SchemaPropertyGroupRequest struct {
	Name        *string                        `json:"name,omitempty"`
	Description *string                        `json:"description,omitempty"`
	Properties  *[]SchemaPropertyGroupProperty `json:"properties,omitempty"`
}

func (c *PosthogClient) CreateSchemaPropertyGroup(ctx context.Context, projectID string, input SchemaPropertyGroupRequest) (SchemaPropertyGroup, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/", projectID)
	result, _, err := doPost[SchemaPropertyGroup](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetSchemaPropertyGroup(ctx context.Context, projectID, id string) (SchemaPropertyGroup, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/%s/", projectID, id)
	return doGet[SchemaPropertyGroup](c, ctx, path)
}

func (c *PosthogClient) UpdateSchemaPropertyGroup(ctx context.Context, projectID, id string, input SchemaPropertyGroupRequest) (SchemaPropertyGroup, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/%s/", projectID, id)
	return doPatch[SchemaPropertyGroup](c, ctx, path, input)
}

func (c *PosthogClient) DeleteSchemaPropertyGroup(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/%s/", projectID, id)
	return doDelete(c, ctx, path)
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./internal/httpclient/ -run TestSchemaPropertyGroupCRUD -v`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add internal/httpclient/schema_property_group.go internal/httpclient/schema_property_group_test.go
git commit -m "feat(httpclient): add schema property group CRUD"
```

---

### Task 2: httpclient — event schemas + event definition resolution

**Files:**
- Create: `internal/httpclient/event_schema.go`
- Test: `internal/httpclient/event_schema_test.go`

**Interfaces:**
- Consumes: `doPost[T]`, `doGet[T]`, `doPatch[T]`, `doDelete`, `listAll[T]` from `client.go`; `SchemaPropertyGroup` from Task 1.
- Produces (used by Tasks 4 and 6):
  - `type EventSchema struct` — incl. `EventName string \`json:"-"\`` (populated by callers after resolving `EventDefinition`, never sent on the wire)
  - `type EventSchemaRequest struct` — incl. `EventName string \`json:"-"\`` (carries the un-resolved name from BuildCreateRequest to Ops.Create/Update)
  - `type EventDefinition struct { ID, Name string }`
  - `CreateEventSchema(ctx, projectID string, input EventSchemaRequest) (EventSchema, error)`
  - `ListEventSchemas(ctx, projectID, eventDefinitionID string) ([]EventSchema, error)` — empty `eventDefinitionID` = unfiltered
  - `GetEventSchema(ctx, projectID, id, eventDefinitionID string) (EventSchema, HTTPStatusCode, error)` — list+find; returns `http.StatusNotFound` + error when absent
  - `UpdateEventSchema(ctx, projectID, id string, input EventSchemaRequest) (EventSchema, HTTPStatusCode, error)`
  - `DeleteEventSchema(ctx, projectID, id string) (HTTPStatusCode, error)`
  - `GetEventDefinition(ctx, projectID, id string) (EventDefinition, HTTPStatusCode, error)`
  - `FindEventDefinitionByName(ctx, projectID, name string) (EventDefinition, error)` — exact match; actionable not-found error
  - `CreateEventDefinition(ctx, projectID, name string) (EventDefinition, error)` — for acceptance-test setup
  - `DeleteEventDefinition(ctx, projectID, id string) (HTTPStatusCode, error)` — for acceptance-test cleanup

- [ ] **Step 1: Write the failing test**

`internal/httpclient/event_schema_test.go`:

```go
package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testESProjectID   = "project-123"
	testESID          = "es-1"
	testESDefID       = "def-1"
	testESListPath    = "/api/projects/project-123/event_schemas/"
	testESDefListPath = "/api/projects/project-123/event_definitions/"
)

func newEventSchemaTestServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == testESListPath:
			var req EventSchemaRequest
			require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
			assert.Equal(t, testESDefID, req.EventDefinition)
			_ = json.NewEncoder(w).Encode(EventSchema{ID: testESID, EventDefinition: testESDefID})
		case r.Method == http.MethodGet && r.URL.Path == testESListPath:
			assert.Equal(t, testESDefID, r.URL.Query().Get("event_definition"))
			_ = json.NewEncoder(w).Encode(PaginatedResponse[EventSchema]{
				Results: []EventSchema{{ID: "other"}, {ID: testESID, EventDefinition: testESDefID}},
			})
		case r.Method == http.MethodPatch && r.URL.Path == testESListPath+testESID+"/":
			_ = json.NewEncoder(w).Encode(EventSchema{ID: testESID, EventDefinition: testESDefID})
		case r.Method == http.MethodDelete && r.URL.Path == testESListPath+testESID+"/":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == testESDefListPath:
			// search returns fuzzy matches; exact-name filtering is the client's job
			_ = json.NewEncoder(w).Encode(PaginatedResponse[EventDefinition]{
				Results: []EventDefinition{
					{ID: "def-2", Name: "checkout_completed_v2"},
					{ID: testESDefID, Name: "checkout_completed"},
				},
			})
		case r.Method == http.MethodGet && r.URL.Path == testESDefListPath+testESDefID+"/":
			_ = json.NewEncoder(w).Encode(EventDefinition{ID: testESDefID, Name: "checkout_completed"})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
}

func TestEventSchemaCRUD(t *testing.T) {
	server := newEventSchemaTestServer(t)
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	ctx := context.Background()

	created, err := client.CreateEventSchema(ctx, testESProjectID, EventSchemaRequest{
		EventDefinition: testESDefID,
		PropertyGroupID: "pg-1",
	})
	require.NoError(t, err)
	assert.Equal(t, testESID, created.ID)

	got, status, err := client.GetEventSchema(ctx, testESProjectID, testESID, testESDefID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testESID, got.ID)

	_, status, err = client.GetEventSchema(ctx, testESProjectID, "missing-id", testESDefID)
	require.Error(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNotFound), status)

	_, status, err = client.UpdateEventSchema(ctx, testESProjectID, testESID, EventSchemaRequest{
		EventDefinition: testESDefID, PropertyGroupID: "pg-2",
	})
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)

	status, err = client.DeleteEventSchema(ctx, testESProjectID, testESID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNoContent), status)
}

func TestFindEventDefinitionByName(t *testing.T) {
	server := newEventSchemaTestServer(t)
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")

	def, err := client.FindEventDefinitionByName(context.Background(), testESProjectID, "checkout_completed")
	require.NoError(t, err)
	assert.Equal(t, testESDefID, def.ID, "must exact-match, not take the first fuzzy result")

	_, err = client.FindEventDefinitionByName(context.Background(), testESProjectID, "never_ingested")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "never_ingested")
	assert.Contains(t, err.Error(), "ingested at least once")
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./internal/httpclient/ -run 'TestEventSchemaCRUD|TestFindEventDefinitionByName' -v`
Expected: FAIL to compile — `undefined: EventSchemaRequest` etc.

- [ ] **Step 3: Write minimal implementation**

`internal/httpclient/event_schema.go`:

```go
package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// EventSchema attaches a schema property group to an event definition.
// The API embeds the full property group on reads; writes reference it by id.
type EventSchema struct {
	ID              string               `json:"id"`
	EventDefinition string               `json:"event_definition"`
	PropertyGroup   *SchemaPropertyGroup `json:"property_group,omitempty"`
	CreatedAt       *string              `json:"created_at,omitempty"`
	UpdatedAt       *string              `json:"updated_at,omitempty"`

	// EventName is not part of the API payload. Resource Ops populate it after
	// resolving EventDefinition so state can surface the human-readable name.
	EventName string `json:"-"`
}

// EventSchemaRequest carries the write payload. EventName is the un-resolved
// event name from the Terraform config; Ops resolve it into EventDefinition
// before calling the API (the field is never serialized).
type EventSchemaRequest struct {
	EventDefinition string `json:"event_definition"`
	PropertyGroupID string `json:"property_group_id"`
	EventName       string `json:"-"`
}

type EventDefinition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *PosthogClient) CreateEventSchema(ctx context.Context, projectID string, input EventSchemaRequest) (EventSchema, error) {
	path := fmt.Sprintf("/api/projects/%s/event_schemas/", projectID)
	result, _, err := doPost[EventSchema](c, ctx, path, input)
	return result, err
}

// ListEventSchemas returns all event schemas, optionally filtered to one event
// definition. Empty eventDefinitionID lists everything (used on import, when
// the definition id is not yet in state).
func (c *PosthogClient) ListEventSchemas(ctx context.Context, projectID, eventDefinitionID string) ([]EventSchema, error) {
	path := fmt.Sprintf("/api/projects/%s/event_schemas/", projectID)
	if eventDefinitionID != "" {
		path += "?event_definition=" + url.QueryEscape(eventDefinitionID)
	}
	return listAll[EventSchema](c, ctx, path)
}

// GetEventSchema emulates a retrieve endpoint (the API only exposes list) by
// listing and matching on id. Absence maps to 404 so the generic resource
// removes the state entry.
func (c *PosthogClient) GetEventSchema(ctx context.Context, projectID, id, eventDefinitionID string) (EventSchema, HTTPStatusCode, error) {
	schemas, err := c.ListEventSchemas(ctx, projectID, eventDefinitionID)
	if err != nil {
		return EventSchema{}, 0, err
	}
	for _, s := range schemas {
		if s.ID == id {
			return s, http.StatusOK, nil
		}
	}
	return EventSchema{}, http.StatusNotFound, fmt.Errorf("event schema %s not found", id)
}

func (c *PosthogClient) UpdateEventSchema(ctx context.Context, projectID, id string, input EventSchemaRequest) (EventSchema, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/event_schemas/%s/", projectID, id)
	return doPatch[EventSchema](c, ctx, path, input)
}

func (c *PosthogClient) DeleteEventSchema(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/event_schemas/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

func (c *PosthogClient) GetEventDefinition(ctx context.Context, projectID, id string) (EventDefinition, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/event_definitions/%s/", projectID, id)
	return doGet[EventDefinition](c, ctx, path)
}

// FindEventDefinitionByName resolves an event name to its definition via the
// search endpoint (fuzzy) plus an exact client-side match.
func (c *PosthogClient) FindEventDefinitionByName(ctx context.Context, projectID, name string) (EventDefinition, error) {
	path := fmt.Sprintf("/api/projects/%s/event_definitions/?search=%s", projectID, url.QueryEscape(name))
	defs, err := listAll[EventDefinition](c, ctx, path)
	if err != nil {
		return EventDefinition{}, err
	}
	for _, def := range defs {
		if def.Name == name {
			return def, nil
		}
	}
	return EventDefinition{}, fmt.Errorf(
		"no event definition named %q found in project %s: events must have been ingested at least once before a schema can be attached",
		name, projectID,
	)
}

// CreateEventDefinition registers an event definition without ingesting an
// event. Used by acceptance tests; not exposed as a Terraform resource.
func (c *PosthogClient) CreateEventDefinition(ctx context.Context, projectID, name string) (EventDefinition, error) {
	path := fmt.Sprintf("/api/projects/%s/event_definitions/", projectID)
	result, _, err := doPost[EventDefinition](c, ctx, path, map[string]string{"name": name})
	return result, err
}

func (c *PosthogClient) DeleteEventDefinition(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/event_definitions/%s/", projectID, id)
	return doDelete(c, ctx, path)
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./internal/httpclient/ -run 'TestEventSchemaCRUD|TestFindEventDefinitionByName' -v`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add internal/httpclient/event_schema.go internal/httpclient/event_schema_test.go
git commit -m "feat(httpclient): add event schema CRUD and event definition resolution"
```

---

### Task 3: resource — posthog_schema_property_group

**Files:**
- Create: `internal/resource/schema_property_group.go`
- Modify: `internal/provider/provider.go` (Resources list, alphabetical: `NewSchemaPropertyGroup` between `NewRoleMembership` and `NewSurvey`)
- Test: `internal/resource/schema_property_group_test.go`

**Interfaces:**
- Consumes (Task 1): `httpclient.SchemaPropertyGroup{,Property,Request}`, `Create/Get/Update/DeleteSchemaPropertyGroup`. Core: `core.BaseStringIdentifiable`, `core.BaseProjectID`, `core.ProjectIDSchemaAttribute()`, `core.ProjectScopedImportParser`, `core.PtrToStringNullIfEmptyTrimmed`, `core.ShouldClearString`.
- Produces: `NewSchemaPropertyGroup() resource.Resource`, `SchemaPropertyGroupTFModel`, `SchemaPropertyGroupOps`, `schemaPropertyTFModel`, `schemaPropertyAttrTypes` (Task 4's tests don't depend on these; docs task does via provider registration).

- [ ] **Step 1: Write the failing test**

`internal/resource/schema_property_group_test.go`:

```go
package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSPGResourceProjectID  = "project-123"
	testSPGResourceID         = "spg-123"
	testSPGResourceCollection = "/api/projects/project-123/schema_property_groups/"
	testSPGResourcePathFull   = "/api/projects/project-123/schema_property_groups/spg-123/"
)

func testPropertyObject(t *testing.T, name, propertyType string, isRequired bool) attr.Value {
	obj, diags := types.ObjectValue(schemaPropertyAttrTypes, map[string]attr.Value{
		"name":                 types.StringValue(name),
		"property_type":        types.StringValue(propertyType),
		"is_required":          types.BoolValue(isRequired),
		"is_optional_in_types": types.BoolValue(false),
		"description":          types.StringNull(),
	})
	require.False(t, diags.HasError())
	return obj
}

func testPropertiesSet(t *testing.T, elems ...attr.Value) types.Set {
	set, diags := types.SetValue(types.ObjectType{AttrTypes: schemaPropertyAttrTypes}, elems)
	require.False(t, diags.HasError())
	return set
}

func TestSchemaPropertyGroupMetadataAndSchema(t *testing.T) {
	require.NotNil(t, NewSchemaPropertyGroup())

	ops := SchemaPropertyGroupOps{}
	assert.Equal(t, "Schema Property Group", ops.ResourceName())

	s := ops.Schema()
	nameAttr, ok := s.Attributes["name"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, nameAttr.Required)

	propsAttr, ok := s.Attributes["properties"].(schema.SetNestedAttribute)
	require.True(t, ok)
	assert.True(t, propsAttr.Optional)

	typeAttr, ok := propsAttr.NestedObject.Attributes["property_type"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, typeAttr.Required)
}

func TestSchemaPropertyGroupBuildCreateRequest(t *testing.T) {
	ops := SchemaPropertyGroupOps{}
	model := SchemaPropertyGroupTFModel{
		Name:        types.StringValue("Checkout"),
		Description: types.StringValue("desc"),
		Properties: testPropertiesSet(t,
			testPropertyObject(t, "cart_value", "Numeric", true),
		),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())
	assert.Equal(t, "Checkout", *req.Name)
	assert.Equal(t, "desc", *req.Description)
	require.NotNil(t, req.Properties)
	require.Len(t, *req.Properties, 1)
	prop := (*req.Properties)[0]
	assert.Equal(t, "cart_value", prop.Name)
	assert.Equal(t, "Numeric", prop.PropertyType)
	require.NotNil(t, prop.IsRequired)
	assert.True(t, *prop.IsRequired)
	assert.Nil(t, prop.ID, "requests must not send property ids (replace-all contract)")
}

func TestSchemaPropertyGroupBuildUpdateRequestClears(t *testing.T) {
	ops := SchemaPropertyGroupOps{}
	state := SchemaPropertyGroupTFModel{
		Name:        types.StringValue("Checkout"),
		Description: types.StringValue("desc"),
		Properties: testPropertiesSet(t,
			testPropertyObject(t, "cart_value", "Numeric", true),
		),
	}
	plan := SchemaPropertyGroupTFModel{
		Name:        types.StringValue("Checkout"),
		Description: types.StringNull(),
		Properties:  types.SetNull(types.ObjectType{AttrTypes: schemaPropertyAttrTypes}),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError())
	require.NotNil(t, req.Description)
	assert.Equal(t, "", *req.Description, "removed description must be cleared with empty string")
	require.NotNil(t, req.Properties)
	assert.Empty(t, *req.Properties, "removed properties must be cleared with empty list")
}

func TestSchemaPropertyGroupMapResponseToModel(t *testing.T) {
	ops := SchemaPropertyGroupOps{}
	required := true
	desc := "group desc"
	resp := httpclient.SchemaPropertyGroup{
		ID:          testSPGResourceID,
		Name:        "Checkout",
		Description: &desc,
		Properties: []httpclient.SchemaPropertyGroupProperty{
			{Name: "cart_value", PropertyType: "Numeric", IsRequired: &required},
		},
	}

	model := SchemaPropertyGroupTFModel{
		Properties: testPropertiesSet(t, testPropertyObject(t, "cart_value", "Numeric", true)),
	}
	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())
	assert.Equal(t, testSPGResourceID, model.ID.ValueString())
	assert.Equal(t, "Checkout", model.Name.ValueString())
	assert.Equal(t, "group desc", model.Description.ValueString())

	var props []schemaPropertyTFModel
	require.False(t, model.Properties.ElementsAs(context.Background(), &props, false).HasError())
	require.Len(t, props, 1)
	assert.Equal(t, "cart_value", props[0].Name.ValueString())
	assert.True(t, props[0].IsRequired.ValueBool())
	assert.False(t, props[0].IsOptionalInTypes.ValueBool())

	// no properties on server + null in config stays null (not empty set)
	nullModel := SchemaPropertyGroupTFModel{
		Properties: types.SetNull(types.ObjectType{AttrTypes: schemaPropertyAttrTypes}),
	}
	diags = ops.MapResponseToModel(context.Background(), httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Empty"}, &nullModel)
	require.False(t, diags.HasError())
	assert.True(t, nullModel.Properties.IsNull())
}

func TestSchemaPropertyGroupCRUDWrapperMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == testSPGResourceCollection:
			_ = json.NewEncoder(w).Encode(httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Checkout"})
		case r.Method == http.MethodGet && r.URL.Path == testSPGResourcePathFull:
			_ = json.NewEncoder(w).Encode(httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Checkout"})
		case r.Method == http.MethodPatch && r.URL.Path == testSPGResourcePathFull:
			_ = json.NewEncoder(w).Encode(httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Checkout"})
		case r.Method == http.MethodDelete && r.URL.Path == testSPGResourcePathFull:
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	ops := SchemaPropertyGroupOps{}
	model := SchemaPropertyGroupTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testSPGResourceID)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testSPGResourceProjectID)},
	}
	name := "Checkout"
	req := httpclient.SchemaPropertyGroupRequest{Name: &name}

	created, err := ops.Create(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, testSPGResourceID, created.ID)

	read, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSPGResourceID, read.ID)

	_, status, err = ops.Update(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)

	status, err = ops.Delete(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusNoContent), status)
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./internal/resource/ -run TestSchemaPropertyGroup -v`
Expected: FAIL to compile — `undefined: NewSchemaPropertyGroup` etc.

- [ ] **Step 3: Write the implementation**

`internal/resource/schema_property_group.go`:

```go
package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

func NewSchemaPropertyGroup() resource.Resource {
	return core.NewGenericResource[SchemaPropertyGroupTFModel, httpclient.SchemaPropertyGroupRequest, httpclient.SchemaPropertyGroup](
		SchemaPropertyGroupOps{},
		core.ProjectScopedImportParser[SchemaPropertyGroupTFModel](),
	)
}

type SchemaPropertyGroupTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Properties  types.Set    `tfsdk:"properties"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

type schemaPropertyTFModel struct {
	Name              types.String `tfsdk:"name"`
	PropertyType      types.String `tfsdk:"property_type"`
	IsRequired        types.Bool   `tfsdk:"is_required"`
	IsOptionalInTypes types.Bool   `tfsdk:"is_optional_in_types"`
	Description       types.String `tfsdk:"description"`
}

var schemaPropertyAttrTypes = map[string]attr.Type{
	"name":                 types.StringType,
	"property_type":        types.StringType,
	"is_required":          types.BoolType,
	"is_optional_in_types": types.BoolType,
	"description":          types.StringType,
}

type SchemaPropertyGroupOps struct{}

func (o SchemaPropertyGroupOps) ResourceName() string {
	return "Schema Property Group"
}

func (o SchemaPropertyGroupOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog schema property groups: named groups of typed properties that can be " +
			"attached to events via `posthog_event_schema`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the property group.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the property group. Unique per project.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the property group.",
			},
			"properties": schema.SetNestedAttribute{
				Optional: true,
				MarkdownDescription: "Typed properties in this group. Property names are unique within a group. " +
					"Updates replace the full property list server-side.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Property name.",
						},
						"property_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Property type: `DateTime`, `String`, `Numeric`, `Boolean`, or `Object`.",
							Validators: []validator.String{
								stringvalidator.OneOf("DateTime", "String", "Numeric", "Boolean", "Object"),
							},
						},
						"is_required": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
							MarkdownDescription: "Whether events must include this property to conform to the schema.",
						},
						"is_optional_in_types": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
							MarkdownDescription: "Whether the property is marked optional in generated type definitions.",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Property description.",
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the property group was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the property group was last updated.",
			},
		},
	}
}

func (o SchemaPropertyGroupOps) BuildCreateRequest(ctx context.Context, model SchemaPropertyGroupTFModel) (httpclient.SchemaPropertyGroupRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := model.Name.ValueString()
	req := httpclient.SchemaPropertyGroupRequest{Name: &name}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}

	if !model.Properties.IsNull() && !model.Properties.IsUnknown() {
		var props []schemaPropertyTFModel
		diags.Append(model.Properties.ElementsAs(ctx, &props, false)...)
		if diags.HasError() {
			return req, diags
		}

		apiProps := make([]httpclient.SchemaPropertyGroupProperty, 0, len(props))
		for _, p := range props {
			apiProp := httpclient.SchemaPropertyGroupProperty{
				Name:         p.Name.ValueString(),
				PropertyType: p.PropertyType.ValueString(),
			}
			if !p.IsRequired.IsNull() && !p.IsRequired.IsUnknown() {
				v := p.IsRequired.ValueBool()
				apiProp.IsRequired = &v
			}
			if !p.IsOptionalInTypes.IsNull() && !p.IsOptionalInTypes.IsUnknown() {
				v := p.IsOptionalInTypes.ValueBool()
				apiProp.IsOptionalInTypes = &v
			}
			if !p.Description.IsNull() && !p.Description.IsUnknown() {
				v := p.Description.ValueString()
				apiProp.Description = &v
			}
			apiProps = append(apiProps, apiProp)
		}
		req.Properties = &apiProps
	}

	return req, diags
}

func (o SchemaPropertyGroupOps) BuildUpdateRequest(ctx context.Context, plan, state SchemaPropertyGroupTFModel) (httpclient.SchemaPropertyGroupRequest, diag.Diagnostics) {
	req, diags := o.BuildCreateRequest(ctx, plan)

	if core.ShouldClearString(plan.Description, state.Description) {
		req.Description = util.StringPtr("")
	}

	// Clear properties if removed from config but previously set
	if plan.Properties.IsNull() && !state.Properties.IsNull() {
		empty := []httpclient.SchemaPropertyGroupProperty{}
		req.Properties = &empty
	}

	return req, diags
}

func (o SchemaPropertyGroupOps) MapResponseToModel(ctx context.Context, resp httpclient.SchemaPropertyGroup, model *SchemaPropertyGroupTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.UpdatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.UpdatedAt)

	if len(resp.Properties) == 0 && model.Properties.IsNull() {
		model.Properties = types.SetNull(types.ObjectType{AttrTypes: schemaPropertyAttrTypes})
		return diags
	}

	elems := make([]schemaPropertyTFModel, 0, len(resp.Properties))
	for _, p := range resp.Properties {
		elems = append(elems, schemaPropertyTFModel{
			Name:              types.StringValue(p.Name),
			PropertyType:      types.StringValue(p.PropertyType),
			IsRequired:        types.BoolValue(p.IsRequired != nil && *p.IsRequired),
			IsOptionalInTypes: types.BoolValue(p.IsOptionalInTypes != nil && *p.IsOptionalInTypes),
			Description:       core.PtrToStringNullIfEmptyTrimmed(p.Description),
		})
	}

	set, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: schemaPropertyAttrTypes}, elems)
	diags.Append(d...)
	model.Properties = set

	return diags
}

func (o SchemaPropertyGroupOps) Create(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel, req httpclient.SchemaPropertyGroupRequest) (httpclient.SchemaPropertyGroup, error) {
	return client.CreateSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), req)
}

func (o SchemaPropertyGroupOps) Read(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel) (httpclient.SchemaPropertyGroup, httpclient.HTTPStatusCode, error) {
	return client.GetSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o SchemaPropertyGroupOps) Update(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel, req httpclient.SchemaPropertyGroupRequest) (httpclient.SchemaPropertyGroup, httpclient.HTTPStatusCode, error) {
	return client.UpdateSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o SchemaPropertyGroupOps) Delete(ctx context.Context, client httpclient.PosthogClient, model SchemaPropertyGroupTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteSchemaPropertyGroup(ctx, model.GetEffectiveProjectID(), model.GetID())
}
```

In `internal/provider/provider.go`, add to the `Resources` list (alphabetical):

```go
		posthogresource.NewRoleMembership,
		posthogresource.NewSchemaPropertyGroup,
		posthogresource.NewSurvey,
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `go test ./internal/resource/ -run TestSchemaPropertyGroup -v && go build ./...`
Expected: PASS, clean build

- [ ] **Step 5: Commit**

```bash
git add internal/resource/schema_property_group.go internal/resource/schema_property_group_test.go internal/provider/provider.go
git commit -m "feat: add posthog_schema_property_group resource"
```

---

### Task 4: resource — posthog_event_schema

**Files:**
- Create: `internal/resource/event_schema.go`
- Modify: `internal/provider/provider.go` (Resources list, alphabetical: `NewEventSchema` between `NewDashboardLayout` and `NewExternalDataSource`)
- Test: `internal/resource/event_schema_test.go`

**Interfaces:**
- Consumes (Task 2): `httpclient.EventSchema`, `httpclient.EventSchemaRequest` (incl. `EventName` carrier field), `httpclient.EventDefinition`, `CreateEventSchema`, `GetEventSchema`, `UpdateEventSchema`, `DeleteEventSchema`, `GetEventDefinition`, `FindEventDefinitionByName`.
- Produces: `NewEventSchema() resource.Resource`, `EventSchemaTFModel`, `EventSchemaOps`.

Key mechanics (read before implementing):
- `BuildCreateRequest` cannot call the API (no client), so it stashes the event **name** in `req.EventName`. `Ops.Create`/`Ops.Update` resolve it via `FindEventDefinitionByName`, fill `req.EventDefinition`, call the API, then set `resp.EventName` so `MapResponseToModel` can populate `event` in state.
- `Ops.Read` looks up the schema via the filtered list (or unfiltered on import, when `event_definition_id` is not in state), then fetches the event definition to recover the name.
- `event_definition_id` is Computed WITHOUT `UseStateForUnknown` — when `event` changes, the id must re-resolve, so a stale-carrying plan modifier would be wrong.

- [ ] **Step 1: Write the failing test**

`internal/resource/event_schema_test.go`:

```go
package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testESResourceProjectID = "project-123"
	testESResourceID        = "es-123"
	testESResourceDefID     = "def-123"
	testESResourceGroupID   = "pg-123"
	testESResourceEvent     = "checkout_completed"
)

func TestEventSchemaMetadataAndSchema(t *testing.T) {
	require.NotNil(t, NewEventSchema())

	ops := EventSchemaOps{}
	assert.Equal(t, "Event Schema", ops.ResourceName())

	s := ops.Schema()
	eventAttr, ok := s.Attributes["event"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, eventAttr.Required)

	groupAttr, ok := s.Attributes["property_group_id"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, groupAttr.Required)

	defAttr, ok := s.Attributes["event_definition_id"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, defAttr.Computed)
	assert.Empty(t, defAttr.PlanModifiers, "must re-resolve when event changes; no UseStateForUnknown")
}

func TestEventSchemaBuildCreateRequest(t *testing.T) {
	ops := EventSchemaOps{}
	model := EventSchemaTFModel{
		Event:           types.StringValue(testESResourceEvent),
		PropertyGroupID: types.StringValue(testESResourceGroupID),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())
	assert.Equal(t, testESResourceEvent, req.EventName)
	assert.Equal(t, testESResourceGroupID, req.PropertyGroupID)
	assert.Empty(t, req.EventDefinition, "resolution happens in Create/Update, not in request building")
}

func TestEventSchemaMapResponseToModel(t *testing.T) {
	ops := EventSchemaOps{}
	resp := httpclient.EventSchema{
		ID:              testESResourceID,
		EventDefinition: testESResourceDefID,
		PropertyGroup:   &httpclient.SchemaPropertyGroup{ID: testESResourceGroupID, Name: "Checkout"},
		EventName:       testESResourceEvent,
	}

	var model EventSchemaTFModel
	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())
	assert.Equal(t, testESResourceID, model.ID.ValueString())
	assert.Equal(t, testESResourceDefID, model.EventDefinitionID.ValueString())
	assert.Equal(t, testESResourceGroupID, model.PropertyGroupID.ValueString())
	assert.Equal(t, testESResourceEvent, model.Event.ValueString())
}

// eventSchemaTestServer covers create/list/patch/delete plus the event
// definition search and retrieve endpoints used for name resolution.
func eventSchemaTestServer(t *testing.T) *httptest.Server {
	listPath := "/api/projects/project-123/event_schemas/"
	defListPath := "/api/projects/project-123/event_definitions/"

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodGet && r.URL.Path == defListPath:
			assert.Equal(t, testESResourceEvent, r.URL.Query().Get("search"))
			_ = json.NewEncoder(w).Encode(httpclient.PaginatedResponse[httpclient.EventDefinition]{
				Results: []httpclient.EventDefinition{{ID: testESResourceDefID, Name: testESResourceEvent}},
			})
		case r.Method == http.MethodGet && r.URL.Path == defListPath+testESResourceDefID+"/":
			_ = json.NewEncoder(w).Encode(httpclient.EventDefinition{ID: testESResourceDefID, Name: testESResourceEvent})
		case r.Method == http.MethodPost && r.URL.Path == listPath:
			var req httpclient.EventSchemaRequest
			require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
			assert.Equal(t, testESResourceDefID, req.EventDefinition, "Create must send the resolved UUID")
			_ = json.NewEncoder(w).Encode(httpclient.EventSchema{ID: testESResourceID, EventDefinition: testESResourceDefID})
		case r.Method == http.MethodGet && r.URL.Path == listPath:
			_ = json.NewEncoder(w).Encode(httpclient.PaginatedResponse[httpclient.EventSchema]{
				Results: []httpclient.EventSchema{{ID: testESResourceID, EventDefinition: testESResourceDefID}},
			})
		case r.Method == http.MethodPatch && r.URL.Path == listPath+testESResourceID+"/":
			_ = json.NewEncoder(w).Encode(httpclient.EventSchema{ID: testESResourceID, EventDefinition: testESResourceDefID})
		case r.Method == http.MethodDelete && r.URL.Path == listPath+testESResourceID+"/":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
}

func TestEventSchemaOpsResolveAndCRUD(t *testing.T) {
	server := eventSchemaTestServer(t)
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	ops := EventSchemaOps{}
	model := EventSchemaTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testESResourceID)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testESResourceProjectID)},
		Event:                  types.StringValue(testESResourceEvent),
		EventDefinitionID:      types.StringValue(testESResourceDefID),
	}
	req := httpclient.EventSchemaRequest{EventName: testESResourceEvent, PropertyGroupID: testESResourceGroupID}

	created, err := ops.Create(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, testESResourceID, created.ID)
	assert.Equal(t, testESResourceEvent, created.EventName, "Create must backfill the resolved event name")

	read, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testESResourceEvent, read.EventName, "Read must recover the event name for state")

	updated, status, err := ops.Update(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testESResourceEvent, updated.EventName)

	status, err = ops.Delete(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusNoContent), status)
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./internal/resource/ -run TestEventSchema -v`
Expected: FAIL to compile — `undefined: NewEventSchema` etc.

- [ ] **Step 3: Write the implementation**

`internal/resource/event_schema.go`:

```go
package resource

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewEventSchema() resource.Resource {
	return core.NewGenericResource[EventSchemaTFModel, httpclient.EventSchemaRequest, httpclient.EventSchema](
		EventSchemaOps{},
		core.ProjectScopedImportParser[EventSchemaTFModel](),
	)
}

type EventSchemaTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Event             types.String `tfsdk:"event"`
	PropertyGroupID   types.String `tfsdk:"property_group_id"`
	EventDefinitionID types.String `tfsdk:"event_definition_id"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
}

type EventSchemaOps struct{}

func (o EventSchemaOps) ResourceName() string {
	return "Event Schema"
}

func (o EventSchemaOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Attach a `posthog_schema_property_group` to an event. The event must have been " +
			"ingested at least once so its event definition exists.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the event schema attachment.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"event": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the event to attach the property group to (e.g. `checkout_completed`).",
			},
			"property_group_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "UUID of the `posthog_schema_property_group` to attach.",
			},
			// No UseStateForUnknown: when `event` changes the id must re-resolve.
			"event_definition_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the event definition the event name resolved to.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the event schema was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the event schema was last updated.",
			},
		},
	}
}

func (o EventSchemaOps) BuildCreateRequest(ctx context.Context, model EventSchemaTFModel) (httpclient.EventSchemaRequest, diag.Diagnostics) {
	// Name resolution needs the API client, which build methods don't get:
	// carry the name and let Create/Update resolve it.
	return httpclient.EventSchemaRequest{
		EventName:       model.Event.ValueString(),
		PropertyGroupID: model.PropertyGroupID.ValueString(),
	}, nil
}

func (o EventSchemaOps) BuildUpdateRequest(ctx context.Context, plan, state EventSchemaTFModel) (httpclient.EventSchemaRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o EventSchemaOps) MapResponseToModel(ctx context.Context, resp httpclient.EventSchema, model *EventSchemaTFModel) diag.Diagnostics {
	model.ID = types.StringValue(resp.ID)
	model.EventDefinitionID = types.StringValue(resp.EventDefinition)
	if resp.EventName != "" {
		model.Event = types.StringValue(resp.EventName)
	}
	if resp.PropertyGroup != nil {
		model.PropertyGroupID = types.StringValue(resp.PropertyGroup.ID)
	}
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.UpdatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.UpdatedAt)
	return nil
}

func (o EventSchemaOps) resolveEvent(ctx context.Context, client httpclient.PosthogClient, projectID string, req *httpclient.EventSchemaRequest) (httpclient.EventDefinition, error) {
	def, err := client.FindEventDefinitionByName(ctx, projectID, req.EventName)
	if err != nil {
		return httpclient.EventDefinition{}, err
	}
	req.EventDefinition = def.ID
	return def, nil
}

func (o EventSchemaOps) Create(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel, req httpclient.EventSchemaRequest) (httpclient.EventSchema, error) {
	def, err := o.resolveEvent(ctx, client, model.GetEffectiveProjectID(), &req)
	if err != nil {
		return httpclient.EventSchema{}, err
	}
	resp, err := client.CreateEventSchema(ctx, model.GetEffectiveProjectID(), req)
	if err != nil {
		return resp, err
	}
	resp.EventName = def.Name
	return resp, nil
}

func (o EventSchemaOps) Read(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel) (httpclient.EventSchema, httpclient.HTTPStatusCode, error) {
	// On import event_definition_id is not in state yet; fall back to an
	// unfiltered list scan.
	eventDefID := ""
	if !model.EventDefinitionID.IsNull() && !model.EventDefinitionID.IsUnknown() {
		eventDefID = model.EventDefinitionID.ValueString()
	}

	resp, status, err := client.GetEventSchema(ctx, model.GetEffectiveProjectID(), model.GetID(), eventDefID)
	if err != nil {
		return resp, status, err
	}

	def, defStatus, err := client.GetEventDefinition(ctx, model.GetEffectiveProjectID(), resp.EventDefinition)
	if err != nil {
		return resp, defStatus, err
	}
	resp.EventName = def.Name

	return resp, http.StatusOK, nil
}

func (o EventSchemaOps) Update(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel, req httpclient.EventSchemaRequest) (httpclient.EventSchema, httpclient.HTTPStatusCode, error) {
	def, err := o.resolveEvent(ctx, client, model.GetEffectiveProjectID(), &req)
	if err != nil {
		return httpclient.EventSchema{}, 0, err
	}
	resp, status, err := client.UpdateEventSchema(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
	if err != nil {
		return resp, status, err
	}
	resp.EventName = def.Name
	return resp, status, nil
}

func (o EventSchemaOps) Delete(ctx context.Context, client httpclient.PosthogClient, model EventSchemaTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteEventSchema(ctx, model.GetEffectiveProjectID(), model.GetID())
}
```

In `internal/provider/provider.go`, add to the `Resources` list (alphabetical):

```go
		posthogresource.NewDashboardLayout,
		posthogresource.NewEventSchema,
		posthogresource.NewExternalDataSource,
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `go test ./internal/resource/ -run TestEventSchema -v && go build ./...`
Expected: PASS, clean build

- [ ] **Step 5: Commit**

```bash
git add internal/resource/event_schema.go internal/resource/event_schema_test.go internal/provider/provider.go
git commit -m "feat: add posthog_event_schema resource"
```

---

### Task 5: examples, generated docs, spec fix

**Files:**
- Create: `examples/resources/posthog_schema_property_group/resource.tf`
- Create: `examples/resources/posthog_schema_property_group/import.sh`
- Create: `examples/resources/posthog_event_schema/resource.tf`
- Create: `examples/resources/posthog_event_schema/import.sh`
- Generated: `docs/resources/schema_property_group.md`, `docs/resources/event_schema.md` (via `make generate`)

**Interfaces:**
- Consumes: registered resources from Tasks 3–4 (tfplugindocs reads the provider schema).
- Produces: rendered docs pages; no code interfaces.

- [ ] **Step 1: Check the example-directory conventions**

Run: `ls examples/resources/posthog_alert/`
Match whatever file set exists there (`resource.tf`, and `import.sh` if present). If `import.sh` is not a convention in this repo, skip creating it.

- [ ] **Step 2: Write the examples**

`examples/resources/posthog_schema_property_group/resource.tf`:

```hcl
resource "posthog_schema_property_group" "checkout" {
  name        = "Checkout"
  description = "Properties shared by checkout events"

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
      is_required   = true
    },
    {
      name          = "currency"
      property_type = "String"
    },
  ]
}
```

`examples/resources/posthog_schema_property_group/import.sh` (if the convention exists):

```bash
terraform import posthog_schema_property_group.checkout <project_id>/<property_group_uuid>
```

`examples/resources/posthog_event_schema/resource.tf`:

```hcl
resource "posthog_schema_property_group" "checkout" {
  name = "Checkout"

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
      is_required   = true
    },
  ]
}

# The event must have been ingested at least once so its definition exists.
resource "posthog_event_schema" "checkout_completed" {
  event             = "checkout_completed"
  property_group_id = posthog_schema_property_group.checkout.id
}
```

`examples/resources/posthog_event_schema/import.sh` (if the convention exists):

```bash
terraform import posthog_event_schema.checkout_completed <project_id>/<event_schema_uuid>
```

- [ ] **Step 3: Generate docs and verify**

Run: `make generate`
Expected: `docs/resources/schema_property_group.md` and `docs/resources/event_schema.md` created; `git diff docs/` shows only those additions (plus any legitimate regeneration).

- [ ] **Step 4: Commit**

```bash
git add examples/resources/posthog_schema_property_group examples/resources/posthog_event_schema docs/resources/schema_property_group.md docs/resources/event_schema.md
git commit -m "docs: add examples and generated docs for event schema resources"
```

---

### Task 6: acceptance tests

**Files:**
- Create: `testacc/schema_property_group_test.go`
- Create: `testacc/event_schema_test.go`

**Interfaces:**
- Consumes: `skipIfNotAcceptance(t)`, `testAccPreCheck(t)`, `testAccProtoV6ProviderFactories` from `testacc/main_test.go`; httpclient methods from Tasks 1–2 (incl. `CreateEventDefinition`/`DeleteEventDefinition` for setup).
- Produces: nothing downstream; final verification layer.

Existing patterns: mirror `testacc/alert_test.go` (destroy check via `httpclient.NewDefaultClient` + env vars, `acctest.RandomWithPrefix("tf-acc-test")`, `resource.Test` steps).

- [ ] **Step 1: Write the property group acceptance test**

`testacc/schema_property_group_test.go`:

```go
package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

func testAccCheckSchemaPropertyGroupDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_schema_property_group" {
			continue
		}

		_, status, err := client.GetSchemaPropertyGroup(context.Background(), projectID, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("schema property group %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

func testAccSchemaPropertyGroupBasic(rName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name        = %[1]q
  description = "acceptance test group"

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
      is_required   = true
    },
    {
      name          = "currency"
      property_type = "String"
    },
  ]
}
`, rName)
}

func testAccSchemaPropertyGroupUpdated(rName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name = %[1]q

  properties = [
    {
      name          = "cart_value"
      property_type = "String"
    },
    {
      name          = "coupon_code"
      property_type = "String"
      description   = "applied coupon"
    },
  ]
}
`, rName)
}

func TestSchemaPropertyGroup_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSchemaPropertyGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSchemaPropertyGroupBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "description", "acceptance test group"),
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "properties.#", "2"),
					resource.TestCheckResourceAttrSet("posthog_schema_property_group.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_schema_property_group.test", "created_at"),
				),
			},
			{
				// retype cart_value, drop currency, add coupon_code, clear description
				Config: testAccSchemaPropertyGroupUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "properties.#", "2"),
					resource.TestCheckNoResourceAttr("posthog_schema_property_group.test", "description"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_schema_property_group.test", "properties.*", map[string]string{
						"name":          "cart_value",
						"property_type": "String",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_schema_property_group.test", "properties.*", map[string]string{
						"name":        "coupon_code",
						"description": "applied coupon",
					}),
				),
			},
			{
				ResourceName:      "posthog_schema_property_group.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs := s.RootModule().Resources["posthog_schema_property_group.test"]
					return fmt.Sprintf("%s/%s", os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID), nil
				},
			},
		},
	})
}
```

- [ ] **Step 2: Write the event schema acceptance test**

`testacc/event_schema_test.go`:

```go
package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// createTestEventDefinition registers an event definition so the event name
// resolves during the test (schemas can only attach to existing definitions).
func createTestEventDefinition(t *testing.T, name string) {
	t.Helper()

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	def, err := client.CreateEventDefinition(context.Background(), projectID, name)
	if err != nil {
		t.Fatalf("failed to create test event definition: %v", err)
	}
	t.Cleanup(func() {
		_, _ = client.DeleteEventDefinition(context.Background(), projectID, def.ID)
	})
}

func testAccCheckEventSchemaDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_event_schema" {
			continue
		}

		_, status, err := client.GetEventSchema(context.Background(), projectID, rs.Primary.ID, "")
		if err == nil {
			return fmt.Errorf("event schema %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

func testAccEventSchemaBasic(groupName, eventName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name = %[1]q

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
    },
  ]
}

resource "posthog_event_schema" "test" {
  event             = %[2]q
  property_group_id = posthog_schema_property_group.test.id
}
`, groupName, eventName)
}

func testAccEventSchemaReattached(groupName, eventName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name = %[1]q

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
    },
  ]
}

resource "posthog_schema_property_group" "second" {
  name = "%[1]s-second"

  properties = [
    {
      name          = "currency"
      property_type = "String"
    },
  ]
}

resource "posthog_event_schema" "test" {
  event             = %[2]q
  property_group_id = posthog_schema_property_group.second.id
}
`, groupName, eventName)
}

func TestEventSchema_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	eventName := acctest.RandomWithPrefix("tf-acc-event")
	createTestEventDefinition(t, eventName)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEventSchemaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEventSchemaBasic(rName, eventName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_event_schema.test", "event", eventName),
					resource.TestCheckResourceAttrSet("posthog_event_schema.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_event_schema.test", "event_definition_id"),
					resource.TestCheckResourceAttrPair(
						"posthog_event_schema.test", "property_group_id",
						"posthog_schema_property_group.test", "id",
					),
				),
			},
			{
				Config: testAccEventSchemaReattached(rName, eventName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"posthog_event_schema.test", "property_group_id",
						"posthog_schema_property_group.second", "id",
					),
				),
			},
			{
				ResourceName:      "posthog_event_schema.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs := s.RootModule().Resources["posthog_event_schema.test"]
					return fmt.Sprintf("%s/%s", os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID), nil
				},
			},
		},
	})
}
```

- [ ] **Step 3: Compile-check without running**

Run: `go vet ./testacc/`
Expected: clean (acceptance tests skip without `TF_ACC=1`).

- [ ] **Step 4: Run the acceptance tests against the local dev instance**

Per `testacc-setup` memory: check `curl -s $POSTHOG_HOST/_health` is `ok` first; env vars `POSTHOG_HOST`, `POSTHOG_API_KEY`, `POSTHOG_PROJECT_ID`, `POSTHOG_ORGANIZATION_ID` must be set.

Run: `TF_ACC=1 go test ./testacc/ -run 'TestSchemaPropertyGroup_Basic|TestEventSchema_Basic' -v -timeout 10m`
Expected: PASS both. Known risk to watch: if the set-nested `booldefault` produces "inconsistent result after apply", the fallback is documented in the plan header of Task 3 — replace defaults with plain Optional bools and normalize in MapResponseToModel by matching plan elements by property name.

- [ ] **Step 5: Commit**

```bash
git add testacc/schema_property_group_test.go testacc/event_schema_test.go
git commit -m "test: add acceptance tests for event schema resources"
```

---

### Task 7: final gates and PR

**Files:** none new.

- [ ] **Step 1: Run all gates**

```bash
make fmt && make lint && make test && make build
```
Expected: all green. `make fmt` must produce no diff.

- [ ] **Step 2: Grep sweep for stale references**

```bash
grep -rn "schema_property_group\|event_schema\|SchemaPropertyGroup\|EventSchema\|EventDefinition" --include="*.go" --include="*.tf" --include="*.md" . | grep -v ".claude/worktrees" | grep -v "docs/superpowers"
```
Expected: hits only in the files this plan created/modified (httpclient, resource, provider registration, testacc, examples, generated docs). No orphans.

- [ ] **Step 3: Submit the branch as a draft PR**

```bash
gt submit --draft
```

PR title: `feat: add posthog_schema_property_group and posthog_event_schema resources (#104)`

PR body must include: what the two resources do, the combined HCL example, the design decisions (set-nested properties, replace-all updates, name resolution, no-retrieve-endpoint Read, enforcement deferred to v2 with the flag-gating reason), a test plan (unit + acceptance results), and a mermaid diagram of the event-name resolution flow. Reference issue #104. End body with the standard generated-with-Claude-Code footer per harness rules.
