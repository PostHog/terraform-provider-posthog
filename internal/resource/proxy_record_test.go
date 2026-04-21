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
	testNormalizedProxyRecordDomain       = "proxy.example.com"
	testProxyRecordOrganizationID         = "org-123"
	testProxyRecordIDValue                = "proxy-1"
	testProxyRecordTargetCNAME            = "abc.proxyhog.com."
	testProxyRecordStatus                 = "waiting"
	testProxyRecordCreatedAt              = "2026-04-21T07:13:42.440644Z"
	testProxyRecordUpdatedAt              = "2026-04-21T07:13:42.440655Z"
	testProxyRecordAPIKey                 = "test-key"
	testProxyRecordClientVersion          = "test"
	testProxyRecordCreatorID        int64 = 432419
	testProxyRecordAPIPathPrefix          = "/api/organizations/"
	testProxyRecordAPIPathSuffix          = "/proxy_records/"
)

func TestNormalizeProxyRecordDomain(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"trim spaces and lowercase": {
			input:    "  Proxy.Example.COM  ",
			expected: testNormalizedProxyRecordDomain,
		},
		"trim trailing dot": {
			input:    testNormalizedProxyRecordDomain + ".",
			expected: testNormalizedProxyRecordDomain,
		},
		"keep already normalized domain": {
			input:    testNormalizedProxyRecordDomain,
			expected: testNormalizedProxyRecordDomain,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, normalizeProxyRecordDomain(tc.input))
		})
	}
}

func TestProxyRecordBuildCreateRequest(t *testing.T) {
	ops := ProxyRecordOps{}
	req, diags := ops.BuildCreateRequest(context.Background(), ProxyRecordTFModel{
		Domain: types.StringValue("  Proxy.Example.COM. "),
	})

	assert.False(t, diags.HasError())
	assert.Equal(t, testNormalizedProxyRecordDomain, req.Domain)
}

func TestProxyRecordBuildUpdateRequest(t *testing.T) {
	ops := ProxyRecordOps{}
	_, diags := ops.BuildUpdateRequest(context.Background(), ProxyRecordTFModel{}, ProxyRecordTFModel{})

	assert.True(t, diags.HasError())
	assert.Contains(t, diags.Errors()[0].Summary(), "Update not supported")
}

func TestProxyRecordResourceNameAndSchema(t *testing.T) {
	ops := ProxyRecordOps{}
	s := ops.Schema()

	assert.Equal(t, "proxy_record", ops.ResourceName())
	assert.Contains(t, s.MarkdownDescription, "organization-scoped")

	domainAttr, ok := s.Attributes["domain"].(schema.StringAttribute)
	require.True(t, ok, "domain must be a string attribute")
	assert.True(t, domainAttr.Required)
	require.Len(t, domainAttr.PlanModifiers, 1)

	targetAttr, ok := s.Attributes["target_cname"].(schema.StringAttribute)
	require.True(t, ok, "target_cname must be a string attribute")
	assert.True(t, targetAttr.Computed)

	createdByAttr, ok := s.Attributes["created_by"].(schema.Int64Attribute)
	require.True(t, ok, "created_by must be an int64 attribute")
	assert.True(t, createdByAttr.Computed)
}

func TestProxyRecordMapResponseToModel(t *testing.T) {
	ops := ProxyRecordOps{}
	resp := httpclient.ProxyRecord{
		ID:          testProxyRecordIDValue,
		Domain:      testNormalizedProxyRecordDomain,
		TargetCNAME: testProxyRecordTargetCNAME,
		Status:      testProxyRecordStatus,
		CreatedAt:   stringPtr(testProxyRecordCreatedAt),
		UpdatedAt:   stringPtr(testProxyRecordUpdatedAt),
		CreatedBy:   int64Ptr(testProxyRecordCreatorID),
	}

	var model ProxyRecordTFModel
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	assert.Equal(t, testProxyRecordIDValue, model.ID.ValueString())
	assert.Equal(t, testNormalizedProxyRecordDomain, model.Domain.ValueString())
	assert.Equal(t, testProxyRecordTargetCNAME, model.TargetCNAME.ValueString())
	assert.Equal(t, testProxyRecordStatus, model.Status.ValueString())
	assert.Equal(t, testProxyRecordCreatedAt, model.CreatedAt.ValueString())
	assert.Equal(t, testProxyRecordUpdatedAt, model.UpdatedAt.ValueString())
	assert.Equal(t, testProxyRecordCreatorID, model.CreatedBy.ValueInt64())
	assert.True(t, model.Message.IsNull())
}

func TestProxyRecordMapResponseToModel_WithMessageAndNilCreator(t *testing.T) {
	ops := ProxyRecordOps{}
	message := "Waiting on DNS"
	resp := httpclient.ProxyRecord{
		ID:          testProxyRecordIDValue,
		Domain:      testNormalizedProxyRecordDomain,
		TargetCNAME: testProxyRecordTargetCNAME,
		Status:      testProxyRecordStatus,
		Message:     &message,
	}

	var model ProxyRecordTFModel
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	assert.Equal(t, message, model.Message.ValueString())
	assert.True(t, model.CreatedAt.IsNull())
	assert.True(t, model.UpdatedAt.IsNull())
	assert.True(t, model.CreatedBy.IsNull())
}

func TestProxyRecordCRUD(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			assert.Equal(t, testProxyRecordCollectionPath(), r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(httpclient.ProxyRecord{
				ID:          testProxyRecordIDValue,
				Domain:      testNormalizedProxyRecordDomain,
				TargetCNAME: testProxyRecordTargetCNAME,
				Status:      testProxyRecordStatus,
			}))
		case http.MethodGet:
			assert.Equal(t, testProxyRecordItemPath(), r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(httpclient.ProxyRecord{
				ID:          testProxyRecordIDValue,
				Domain:      testNormalizedProxyRecordDomain,
				TargetCNAME: testProxyRecordTargetCNAME,
				Status:      "valid",
			}))
		case http.MethodDelete:
			assert.Equal(t, testProxyRecordItemPath(), r.URL.Path)
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, testProxyRecordAPIKey, testProxyRecordClientVersion)
	ops := ProxyRecordOps{}
	model := ProxyRecordTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testProxyRecordIDValue)},
		BaseOrganizationID:     core.BaseOrganizationID{OrganizationID: types.StringValue(testProxyRecordOrganizationID)},
		Domain:                 types.StringValue(testNormalizedProxyRecordDomain),
	}

	created, err := ops.Create(context.Background(), client, model, httpclient.ProxyRecordRequest{Domain: testNormalizedProxyRecordDomain})
	require.NoError(t, err)
	assert.Equal(t, testProxyRecordIDValue, created.ID)

	read, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "valid", read.Status)

	status, err = ops.Delete(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
}

func TestProxyRecordUpdateReturnsError(t *testing.T) {
	ops := ProxyRecordOps{}

	_, status, err := ops.Update(context.Background(), httpclient.PosthogClient{}, ProxyRecordTFModel{}, httpclient.ProxyRecordRequest{})

	require.Error(t, err)
	assert.Zero(t, status)
	assert.Contains(t, err.Error(), "do not support updates")
}

func stringPtr(v string) *string {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}

func testProxyRecordCollectionPath() string {
	return testProxyRecordAPIPathPrefix + testProxyRecordOrganizationID + testProxyRecordAPIPathSuffix
}

func testProxyRecordItemPath() string {
	return testProxyRecordCollectionPath() + testProxyRecordIDValue + "/"
}
