package httpclient

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CurrentOrganizationAlias is the special identifier PostHog accepts to refer to
// the authenticated user's organization.
const CurrentOrganizationAlias = "@current"

// uuidPattern matches a canonical UUID (8-4-4-4-12 hex groups). PostHog
// organization IDs are UUIDs, so anything matching this is treated as an ID and
// passed through without an API call; anything else is treated as a slug.
var uuidPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// Organization is a subset of the PostHog organization resource, containing the
// fields needed to resolve a user-supplied slug or "@current" to a UUID.
type Organization struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// ListOrganizations returns all organizations the authenticated user can access.
func (c *PosthogClient) ListOrganizations(ctx context.Context) ([]Organization, error) {
	return listAll[Organization](c, ctx, "/api/organizations/")
}

// ResolveOrganizationID resolves a user-supplied organization identifier to a
// UUID suitable for building API paths. It accepts:
//   - a UUID, which is returned unchanged with no API call;
//   - the literal "@current", resolved via GET /api/organizations/@current/;
//   - an organization slug, resolved by listing organizations and matching slug.
//
// Resolution is purely for constructing request URLs — the caller must never
// persist the resolved UUID back into Terraform state, otherwise a config that
// uses a slug would show a permanent diff against the resolved UUID in state.
//
// Results are cached on the client (shared across copies via a pointer) so that
// repeated calls across resources within a run do not refetch. The cache lock is
// not held across the HTTP fetch, so a concurrent cold start (Terraform applies
// resources in parallel) may issue a few redundant lookups for the same
// identifier before the cache is populated; the result is idempotent (all writes
// converge on the same UUID), so this is left uncollapsed rather than guarded
// with single-flight, which would be disproportionate for a short-lived CLI run.
func (c *PosthogClient) ResolveOrganizationID(ctx context.Context, idOrSlug string) (string, error) {
	// An empty identifier means no organization is configured; pass it through
	// unchanged and let the API surface the error, preserving prior behavior.
	if idOrSlug == "" {
		return "", nil
	}

	// UUIDs are already valid path segments; no resolution needed.
	if uuidPattern.MatchString(idOrSlug) {
		return idOrSlug, nil
	}

	if cached, ok := c.cachedOrganizationID(idOrSlug); ok {
		tflog.Trace(ctx, "resolved organization id from cache", map[string]any{"key": idOrSlug, "resolved_id": cached})
		return cached, nil
	}

	var resolved string
	if idOrSlug == CurrentOrganizationAlias {
		tflog.Debug(ctx, "resolving @current organization via API", map[string]any{"key": idOrSlug})
		org, _, err := doGet[Organization](c, ctx, "/api/organizations/@current/")
		if err != nil {
			return "", fmt.Errorf("failed to resolve %q organization: %w", CurrentOrganizationAlias, err)
		}
		resolved = org.ID
	} else {
		tflog.Debug(ctx, "resolving organization slug via API", map[string]any{"key": idOrSlug})
		orgs, err := c.ListOrganizations(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to list organizations to resolve slug %q: %w", idOrSlug, err)
		}
		for _, org := range orgs {
			if org.Slug == idOrSlug {
				resolved = org.ID
				break
			}
		}
		if resolved == "" {
			return "", fmt.Errorf("no organization found matching slug %q; pass a valid organization UUID, slug, or %q", idOrSlug, CurrentOrganizationAlias)
		}
	}

	c.storeOrganizationID(idOrSlug, resolved)
	tflog.Debug(ctx, "resolved and cached organization id", map[string]any{"key": idOrSlug, "resolved_id": resolved})
	return resolved, nil
}

// orgPath resolves orgIDOrSlug (a UUID, slug, or "@current") to a UUID and formats
// an org-scoped API path from it, with the resolved id as the first `%s`. It is the
// single place client methods perform org resolution, so adding a new org-scoped
// endpoint is a one-line call rather than a repeated resolve-and-check block.
func (c *PosthogClient) orgPath(ctx context.Context, orgIDOrSlug, format string, args ...any) (string, error) {
	orgID, err := c.ResolveOrganizationID(ctx, orgIDOrSlug)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(format, append([]any{orgID}, args...)...), nil
}

func (c *PosthogClient) cachedOrganizationID(key string) (string, bool) {
	if c.orgCache == nil {
		return "", false
	}
	c.orgCache.mu.Lock()
	defer c.orgCache.mu.Unlock()
	id, ok := c.orgCache.byKey[key]
	return id, ok
}

func (c *PosthogClient) storeOrganizationID(key, id string) {
	if c.orgCache == nil {
		return
	}
	c.orgCache.mu.Lock()
	defer c.orgCache.mu.Unlock()
	c.orgCache.byKey[key] = id
}
