package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testExperimentID      = "42"
	testExperimentProject = "proj-1"
	experimentsAPI        = "/api/projects/"
	experimentsSub        = "/experiments/"
)

func experimentCollectionPath() string {
	return experimentsAPI + testExperimentProject + experimentsSub
}

func experimentItemPath() string {
	return experimentCollectionPath() + testExperimentID + "/"
}

func experimentSubActionPath(action string) string {
	return experimentItemPath() + action + "/"
}

func TestCreateExperiment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, experimentCollectionPath(), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "Pricing test", body["name"])
		assert.Equal(t, "pricing-test", body["feature_flag_key"])

		writeJSONResponse(t, w, Experiment{
			ID:             42,
			Name:           "Pricing test",
			FeatureFlagKey: "pricing-test",
			Status:         "draft",
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	name := "Pricing test"
	key := "pricing-test"
	exp, err := client.CreateExperiment(context.Background(), testExperimentProject, ExperimentRequest{
		Name:           &name,
		FeatureFlagKey: &key,
	})

	require.NoError(t, err)
	assert.Equal(t, int64(42), exp.ID)
	assert.Equal(t, "draft", exp.Status)
}

func TestGetExperiment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, experimentItemPath(), r.URL.Path)

		writeJSONResponse(t, w, Experiment{ID: 42, Name: "Pricing test", Status: "running"})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	exp, status, err := client.GetExperiment(context.Background(), testExperimentProject, testExperimentID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "running", exp.Status)
}

func TestUpdateExperiment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, experimentItemPath(), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "New name", body["name"])

		writeJSONResponse(t, w, Experiment{ID: 42, Name: "New name", Status: "draft"})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	name := "New name"
	exp, status, err := client.UpdateExperiment(context.Background(), testExperimentProject, testExperimentID, ExperimentRequest{Name: &name})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "New name", exp.Name)
}

// TestDeleteExperiment verifies deletion is a soft-delete: a PATCH setting deleted=true
// (hard DELETE is forbidden by the API), sending only that field.
func TestDeleteExperiment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, experimentItemPath(), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, true, body["deleted"])
		_, hasName := body["name"]
		assert.False(t, hasName, "soft-delete must not send other fields")

		writeJSONResponse(t, w, Experiment{ID: 42, Deleted: util.BoolPtr(true)})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	status, err := client.DeleteExperiment(context.Background(), testExperimentProject, testExperimentID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
}

func TestLaunchExperiment(t *testing.T) {
	assertNoBodyAction(t, "launch", func(client PosthogClient) (Experiment, HTTPStatusCode, error) {
		return client.LaunchExperiment(context.Background(), testExperimentProject, testExperimentID)
	}, "running")
}

func TestPauseExperiment(t *testing.T) {
	assertNoBodyAction(t, "pause", func(client PosthogClient) (Experiment, HTTPStatusCode, error) {
		return client.PauseExperiment(context.Background(), testExperimentProject, testExperimentID)
	}, "paused")
}

func TestResumeExperiment(t *testing.T) {
	assertNoBodyAction(t, "resume", func(client PosthogClient) (Experiment, HTTPStatusCode, error) {
		return client.ResumeExperiment(context.Background(), testExperimentProject, testExperimentID)
	}, "running")
}

func TestEndExperiment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, experimentSubActionPath("end"), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "won", body["conclusion"])

		writeJSONResponse(t, w, Experiment{ID: 42, Status: "stopped", Conclusion: util.StringPtr("won")})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	exp, status, err := client.EndExperiment(context.Background(), testExperimentProject, testExperimentID, ExperimentEndRequest{
		Conclusion: util.StringPtr("won"),
	})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "stopped", exp.Status)
}

func TestShipVariant(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, experimentSubActionPath("ship_variant"), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "test", body["variant_key"])
		assert.Equal(t, false, body["release_to_everyone"])

		writeJSONResponse(t, w, Experiment{ID: 42, Status: "stopped"})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	exp, status, err := client.ShipVariant(context.Background(), testExperimentProject, testExperimentID, ExperimentShipVariantRequest{
		VariantKey:        "test",
		ReleaseToEveryone: false,
	})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "stopped", exp.Status)
}

// assertNoBodyAction exercises a lifecycle sub-action that takes no meaningful body and
// asserts it POSTs to the right sub-action path and returns the updated experiment.
func assertNoBodyAction(t *testing.T, action string, call func(PosthogClient) (Experiment, HTTPStatusCode, error), wantStatus string) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, experimentSubActionPath(action), r.URL.Path)
		writeJSONResponse(t, w, Experiment{ID: 42, Status: wantStatus})
	}))
	defer server.Close()

	exp, status, err := call(newTestPosthogClient(server))
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, wantStatus, exp.Status)
}
