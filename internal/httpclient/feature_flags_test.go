package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListFeatureFlags_IncludesArchivedFlagsAndFollowsPages(t *testing.T) {
	var queries []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/projects/123/feature_flags/", r.URL.Path)
		queries = append(queries, r.URL.RawQuery)

		switch r.URL.RawQuery {
		case "archived=false":
			next := "/api/projects/123/feature_flags/?archived=false&limit=100&offset=100"
			writeJSONResponse(t, w, PaginatedResponse[FeatureFlag]{Results: []FeatureFlag{{ID: 1, Key: "first-page"}}, Next: &next})
		case "archived=false&limit=100&offset=100":
			writeJSONResponse(t, w, PaginatedResponse[FeatureFlag]{Results: []FeatureFlag{{ID: 2, Key: "second-page"}}})
		case "archived=true":
			writeJSONResponse(t, w, PaginatedResponse[FeatureFlag]{Results: []FeatureFlag{{ID: 3, Key: "archived", Archived: util.BoolPtr(true)}}})
		default:
			t.Errorf("unexpected query %q", r.URL.RawQuery)
		}
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	flags, err := client.ListFeatureFlags(context.Background(), "123")

	require.NoError(t, err)
	assert.Equal(t, []string{"archived=false", "archived=false&limit=100&offset=100", "archived=true"}, queries)
	require.Len(t, flags, 3)
	assert.Equal(t, []int64{1, 2, 3}, []int64{flags[0].ID, flags[1].ID, flags[2].ID})
	assert.True(t, *flags[2].Archived)
}

func TestListFeatureFlags_DeduplicatesWhenTheArchivedFilterIsIgnored(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSONResponse(t, w, PaginatedResponse[FeatureFlag]{Results: []FeatureFlag{{ID: 1, Key: "a"}, {ID: 2, Key: "b"}}})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	flags, err := client.ListFeatureFlags(context.Background(), "123")

	require.NoError(t, err)
	assert.Len(t, flags, 2, "both passes returned the same flags, which must appear once")
}

func TestListFeatureFlags_FailsWholeWhenTheArchivedPassFails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("archived") == "true" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		writeJSONResponse(t, w, PaginatedResponse[FeatureFlag]{Results: []FeatureFlag{{ID: 1, Key: "a"}}})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	flags, err := client.ListFeatureFlags(context.Background(), "123")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "archived=true", "the error names the pass that failed")
	assert.Nil(t, flags, "half a list would look like flags that do not exist")
}

func TestListFeatureFlags_ReturnsTheListError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	_, err := client.ListFeatureFlags(context.Background(), "123")

	require.Error(t, err)
}
