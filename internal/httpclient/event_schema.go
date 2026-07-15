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
