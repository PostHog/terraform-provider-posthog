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
	testProxyRecordID          = "proxy-1"
	testProxyRecordDomain      = "proxy.example.com"
	testProxyRecordTargetCNAME = "abc.proxyhog.com."
	testOrganizationID         = "org-123"
	testAPIKey                 = "test-key"
	testClientVersion          = "test"
	jsonContentTypeHeader      = "Content-Type"
	jsonContentTypeValue       = "application/json"
	apiOrganizationsPrefix     = "/api/organizations/"
	proxyRecordsPathSuffix     = "/proxy_records/"
)

func TestListProxyRecords(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, testProxyRecordsCollectionPath(), r.URL.Path)

		resp := ProxyRecordListResponse{
			MaxProxyRecords: int64Ptr(10),
			Results: []ProxyRecord{
				{ID: testProxyRecordID, Domain: "a.example.com", TargetCNAME: "cname-1.proxyhog.com.", Status: "waiting"},
				{ID: "proxy-2", Domain: "b.example.com", TargetCNAME: "cname-2.proxyhog.com.", Status: "valid"},
			},
		}
		writeJSONResponse(t, w, resp)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	records, err := client.ListProxyRecords(context.Background(), testOrganizationID)

	require.NoError(t, err)
	assert.Len(t, records, 2)
	assert.Equal(t, testProxyRecordID, records[0].ID)
	assert.Equal(t, "valid", records[1].Status)
}

func TestCreateProxyRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testProxyRecordsCollectionPath(), r.URL.Path)

		var req ProxyRecordRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, testProxyRecordDomain, req.Domain)

		resp := ProxyRecord{
			ID:          testProxyRecordID,
			Domain:      testProxyRecordDomain,
			TargetCNAME: testProxyRecordTargetCNAME,
			Status:      "waiting",
		}
		writeJSONResponse(t, w, resp)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	record, err := client.CreateProxyRecord(context.Background(), testOrganizationID, ProxyRecordRequest{Domain: testProxyRecordDomain})

	require.NoError(t, err)
	assert.Equal(t, testProxyRecordID, record.ID)
	assert.Equal(t, testProxyRecordTargetCNAME, record.TargetCNAME)
	assert.Equal(t, "waiting", record.Status)
}

func TestGetProxyRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testProxyRecordItemPath(testProxyRecordID), r.URL.Path)

		resp := ProxyRecord{
			ID:          testProxyRecordID,
			Domain:      testProxyRecordDomain,
			TargetCNAME: testProxyRecordTargetCNAME,
			Status:      "valid",
		}
		writeJSONResponse(t, w, resp)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	record, status, err := client.GetProxyRecord(context.Background(), testOrganizationID, testProxyRecordID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "valid", record.Status)
}

func TestDeleteProxyRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, testProxyRecordItemPath(testProxyRecordID), r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	status, err := client.DeleteProxyRecord(context.Background(), testOrganizationID, testProxyRecordID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
}

func newTestPosthogClient(server *httptest.Server) PosthogClient {
	return NewClient(server.Client(), server.URL, testAPIKey, testClientVersion)
}

func writeJSONResponse(t *testing.T, w http.ResponseWriter, body any) {
	t.Helper()
	w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
	require.NoError(t, json.NewEncoder(w).Encode(body))
}

func int64Ptr(v int64) *int64 {
	return &v
}

func testProxyRecordsCollectionPath() string {
	return apiOrganizationsPrefix + testOrganizationID + proxyRecordsPathSuffix
}

func testProxyRecordItemPath(proxyRecordID string) string {
	return testProxyRecordsCollectionPath() + proxyRecordID + "/"
}
