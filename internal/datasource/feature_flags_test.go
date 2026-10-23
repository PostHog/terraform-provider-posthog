package datasource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testOrgID = "11111111-2222-3333-4444-555555555555"

func TestFeatureFlagsToList_DropsDeletedFlagsAndSortsByKey(t *testing.T) {
	flags := []httpclient.FeatureFlag{
		{ID: 2, Key: "zeta", Name: util.StringPtr("Zeta"), Active: util.BoolPtr(true)},
		{ID: 3, Key: "gone", Deleted: util.BoolPtr(true)},
		{ID: 1, Key: "alpha", Name: util.StringPtr("  "), Active: util.BoolPtr(false), Archived: util.BoolPtr(true)},
	}

	list, diags := featureFlagsToList(context.Background(), flags)
	require.False(t, diags.HasError())

	var got []featureFlagSummaryModel
	require.False(t, list.ElementsAs(context.Background(), &got, false).HasError())
	require.Len(t, got, 2)

	assert.Equal(t, "alpha", got[0].Key.ValueString())
	assert.Equal(t, int64(1), got[0].ID.ValueInt64())
	assert.True(t, got[0].Name.IsNull(), "whitespace-only name should map to null")
	assert.False(t, got[0].Active.ValueBool())
	assert.True(t, got[0].Archived.ValueBool())

	assert.Equal(t, "zeta", got[1].Key.ValueString())
	assert.Equal(t, "Zeta", got[1].Name.ValueString())
	assert.True(t, got[1].Active.ValueBool())
	assert.False(t, got[1].Archived.ValueBool(), "a missing archived field means not archived")
}

func TestFeatureFlagsToList_NoFlagsIsAnEmptyListNotNull(t *testing.T) {
	list, diags := featureFlagsToList(context.Background(), nil)
	require.False(t, diags.HasError())

	assert.False(t, list.IsNull(), "an empty project should give an empty list, so for_each over it works")
	assert.Empty(t, list.Elements())
}

func TestMatchProjectByName(t *testing.T) {
	projects := []httpclient.Project{
		{ID: 1, Name: util.StringPtr("Production")},
		{ID: 2, Name: util.StringPtr(" staging ")},
		{ID: 3, Name: nil},
		{ID: 4, Name: util.StringPtr("dup")},
		{ID: 5, Name: util.StringPtr("DUP")},
	}

	tests := map[string]struct {
		projects  []httpclient.Project
		name      string
		wantID    string
		wantFound bool
		wantErr   string
	}{
		"matches ignoring case":               {projects: projects, name: "production", wantID: "1", wantFound: true},
		"matches ignoring surrounding spaces": {projects: projects, name: "staging", wantID: "2", wantFound: true},
		"an unmatched name is not an error":   {projects: projects, name: "preview-123"},
		"no projects at all":                  {projects: nil, name: "production"},
		"a shared name is an error":           {projects: projects, name: "dup", wantErr: "IDs 4, 5"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			id, found, err := matchProjectByName(tt.projects, tt.name)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantFound, found)
			assert.Equal(t, tt.wantID, id)
		})
	}
}

// projectsServer answers the organization's project list. A nil projects
// list makes it fail the request instead.
func projectsServer(t *testing.T, projects []httpclient.Project) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/organizations/"+testOrgID+"/projects/", r.URL.Path)
		if projects == nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		require.NoError(t, json.NewEncoder(w).Encode(httpclient.PaginatedResponse[httpclient.Project]{Results: projects}))
	}))
	t.Cleanup(server.Close)
	return server
}

func TestResolveProjectID(t *testing.T) {
	named := []httpclient.Project{
		{ID: 7, Name: util.StringPtr("production")},
		{ID: 8, Name: util.StringPtr("dup")},
		{ID: 9, Name: util.StringPtr("dup")},
	}

	tests := map[string]struct {
		config         FeatureFlagsDataSourceModel
		defaultProject string
		defaultOrg     string
		projects       []httpclient.Project
		wantID         string
		wantErr        string
	}{
		"project_id from the data source": {
			config: FeatureFlagsDataSourceModel{ProjectID: types.StringValue("42"), ProjectName: types.StringNull()},
			wantID: "42",
		},
		"project_id from the provider": {
			config:         FeatureFlagsDataSourceModel{ProjectID: types.StringNull(), ProjectName: types.StringNull()},
			defaultProject: "43",
			wantID:         "43",
		},
		"no project at all": {
			config:  FeatureFlagsDataSourceModel{ProjectID: types.StringNull(), ProjectName: types.StringNull()},
			wantErr: "Missing project_id",
		},
		"project_name with no organization": {
			config:  FeatureFlagsDataSourceModel{ProjectName: types.StringValue("production"), OrganizationID: types.StringNull()},
			wantErr: "Missing organization_id",
		},
		"project_name found": {
			config:   FeatureFlagsDataSourceModel{ProjectName: types.StringValue("production"), OrganizationID: types.StringValue(testOrgID)},
			projects: named,
			wantID:   "7",
		},
		"project_name in the provider's organization": {
			config:     FeatureFlagsDataSourceModel{ProjectName: types.StringValue("production"), OrganizationID: types.StringNull()},
			defaultOrg: testOrgID,
			projects:   named,
			wantID:     "7",
		},
		"project_name not found": {
			config:   FeatureFlagsDataSourceModel{ProjectName: types.StringValue("preview"), OrganizationID: types.StringValue(testOrgID)},
			projects: named,
		},
		"project_name shared by two projects": {
			config:   FeatureFlagsDataSourceModel{ProjectName: types.StringValue("dup"), OrganizationID: types.StringValue(testOrgID)},
			projects: named,
			wantErr:  "Set project_id instead.",
		},
		"listing projects fails": {
			config:  FeatureFlagsDataSourceModel{ProjectName: types.StringValue("production"), OrganizationID: types.StringValue(testOrgID)},
			wantErr: "Error listing projects",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			server := projectsServer(t, tt.projects)
			d := &FeatureFlagsDataSource{
				client:                httpclient.NewClient(server.Client(), server.URL, "test-key", "test", httpclient.WithNoRetry()),
				defaultProjectID:      tt.defaultProject,
				defaultOrganizationID: tt.defaultOrg,
			}

			id, diags := d.resolveProjectID(context.Background(), tt.config)

			if tt.wantErr != "" {
				require.True(t, diags.HasError())
				summary, detail := diags.Errors()[0].Summary(), diags.Errors()[0].Detail()
				assert.Contains(t, summary+" "+detail, tt.wantErr)
				assert.True(t, id.IsNull(), "an error leaves project_id null")
				return
			}
			require.False(t, diags.HasError(), "%v", diags)
			if tt.wantID == "" {
				assert.True(t, id.IsNull(), "an unmatched project_name resolves to a null project_id")
				return
			}
			assert.Equal(t, tt.wantID, id.ValueString())
		})
	}
}
