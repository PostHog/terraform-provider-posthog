package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testOrgUUID = "11111111-2222-3333-4444-555555555555"
	testOrgSlug = "acme-corp"
)

func TestResolveOrganizationID_UUIDPassthrough(t *testing.T) {
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		t.Errorf("unexpected HTTP request for UUID passthrough: %s", r.URL.Path)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	resolved, err := client.ResolveOrganizationID(context.Background(), testOrgUUID)

	require.NoError(t, err, "resolving a UUID should not error")
	assert.Equal(t, testOrgUUID, resolved, "a UUID should be returned unchanged")
	assert.Equal(t, int32(0), atomic.LoadInt32(&requestCount), "resolving a UUID must not make any HTTP request")
}

func TestResolveOrganizationID_CurrentAlias(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/organizations/@current/", r.URL.Path, "@current should hit the @current endpoint")
		writeJSONResponse(t, w, Organization{ID: testOrgUUID, Slug: testOrgSlug, Name: "Acme Corp"})
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	resolved, err := client.ResolveOrganizationID(context.Background(), CurrentOrganizationAlias)

	require.NoError(t, err, "resolving @current should not error")
	assert.Equal(t, testOrgUUID, resolved, "@current should resolve to the returned organization ID")
}

func TestResolveOrganizationID_SlugResolution(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/organizations/", r.URL.Path, "slug resolution should list organizations")
		resp := PaginatedResponse[Organization]{
			Results: []Organization{
				{ID: "99999999-8888-7777-6666-555555555555", Slug: "other-org", Name: "Other"},
				{ID: testOrgUUID, Slug: testOrgSlug, Name: "Acme Corp"},
			},
		}
		writeJSONResponse(t, w, resp)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	resolved, err := client.ResolveOrganizationID(context.Background(), testOrgSlug)

	require.NoError(t, err, "resolving a known slug should not error")
	assert.Equal(t, testOrgUUID, resolved, "slug should resolve to the matching organization ID")
}

func TestResolveOrganizationID_SlugNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := PaginatedResponse[Organization]{
			Results: []Organization{
				{ID: testOrgUUID, Slug: testOrgSlug, Name: "Acme Corp"},
			},
		}
		writeJSONResponse(t, w, resp)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	_, err := client.ResolveOrganizationID(context.Background(), "does-not-exist")

	require.Error(t, err, "resolving an unknown slug should return an error")
	assert.Contains(t, err.Error(), "does-not-exist", "error should name the unresolved slug")
}

func TestResolveOrganizationID_CurrentAliasHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/organizations/@current/", r.URL.Path)
		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	_, err := client.ResolveOrganizationID(context.Background(), CurrentOrganizationAlias)

	require.Error(t, err, "a non-2xx response from the @current endpoint should surface an error")
	assert.Contains(t, err.Error(), CurrentOrganizationAlias, "error should reference the @current alias")
}

func TestResolveOrganizationID_SlugListHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/organizations/", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	_, err := client.ResolveOrganizationID(context.Background(), testOrgSlug)

	require.Error(t, err, "an error listing organizations should surface during slug resolution")
	assert.Contains(t, err.Error(), testOrgSlug, "error should reference the slug being resolved")
}

func TestResolveOrganizationID_CachesSlugResolution(t *testing.T) {
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		resp := PaginatedResponse[Organization]{
			Results: []Organization{
				{ID: testOrgUUID, Slug: testOrgSlug, Name: "Acme Corp"},
			},
		}
		writeJSONResponse(t, w, resp)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)

	first, err := client.ResolveOrganizationID(context.Background(), testOrgSlug)
	require.NoError(t, err, "first resolution should not error")
	assert.Equal(t, testOrgUUID, first)

	second, err := client.ResolveOrganizationID(context.Background(), testOrgSlug)
	require.NoError(t, err, "second resolution should not error")
	assert.Equal(t, testOrgUUID, second)

	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount), "second resolution should be served from cache without a new HTTP request")
}

func TestResolveOrganizationID_CachesCurrentAlias(t *testing.T) {
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		writeJSONResponse(t, w, Organization{ID: testOrgUUID, Slug: testOrgSlug, Name: "Acme Corp"})
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)

	first, err := client.ResolveOrganizationID(context.Background(), CurrentOrganizationAlias)
	require.NoError(t, err, "first @current resolution should not error")
	assert.Equal(t, testOrgUUID, first)

	second, err := client.ResolveOrganizationID(context.Background(), CurrentOrganizationAlias)
	require.NoError(t, err, "second @current resolution should not error")
	assert.Equal(t, testOrgUUID, second)

	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount), "second @current resolution should be served from cache without a new HTTP request")
}

func TestResolveOrganizationID_CacheSharedAcrossClientCopies(t *testing.T) {
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		resp := PaginatedResponse[Organization]{
			Results: []Organization{{ID: testOrgUUID, Slug: testOrgSlug, Name: "Acme Corp"}},
		}
		writeJSONResponse(t, w, resp)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	_, err := client.ResolveOrganizationID(context.Background(), testOrgSlug)
	require.NoError(t, err)

	// PosthogClient is passed around by value; a copy must share the same cache.
	clientCopy := client
	resolved, err := clientCopy.ResolveOrganizationID(context.Background(), testOrgSlug)
	require.NoError(t, err)
	assert.Equal(t, testOrgUUID, resolved)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount), "a value copy of the client must reuse the shared resolution cache")
}

func TestResolveOrganizationID_EmptyPassthrough(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected HTTP request for empty organization id: %s", r.URL.Path)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
	resolved, err := client.ResolveOrganizationID(context.Background(), "")

	require.NoError(t, err, "empty organization id should pass through without error")
	assert.Equal(t, "", resolved, "empty organization id should be returned unchanged")
}
