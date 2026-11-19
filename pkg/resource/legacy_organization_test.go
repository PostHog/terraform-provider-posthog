package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/posthog/terraform-provider/pkg/data"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
)

const testLegacyOrganizationID = "00000000-0000-0000-0000-000000000456"

// Role memberships and proxy records have no update endpoint. The only in-place change they
// can plan is organization_id moving off "@current", so applying that plan through the
// resource's Update must succeed without writing to the API and record the provider's
// organization_id.
func TestLegacyOrganizationIDUpdateWithoutUpdateEndpoint(t *testing.T) {
	email := "alice@example.com"
	joinedAt := "2026-04-21T07:13:42Z"

	tests := map[string]struct {
		resource tfresource.Resource
		schema   schema.Schema
		state    func(organizationID types.String) any
		response any
	}{
		"role membership": {
			resource: NewRoleMembership(),
			schema:   RoleMembershipOps{}.Schema(),
			state: func(organizationID types.String) any {
				return &RoleMembershipTFModel{
					BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue("membership-1")},
					BaseOrganizationID:     core.BaseOrganizationID{OrganizationID: organizationID},
					BaseRoleID:             core.BaseRoleID{RoleID: types.StringValue("role-1")},
					UserUUID:               types.StringValue("user-1"),
					UserEmail:              types.StringValue(email),
					JoinedAt:               types.StringValue(joinedAt),
				}
			},
			response: httpclient.RoleMembership{
				ID:       "membership-1",
				RoleID:   "role-1",
				User:     &httpclient.RoleMembershipUser{UUID: "user-1", Email: &email},
				JoinedAt: &joinedAt,
			},
		},
		"proxy record": {
			resource: NewProxyRecord(),
			schema:   ProxyRecordOps{}.Schema(),
			state: func(organizationID types.String) any {
				return &ProxyRecordTFModel{
					BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testProxyRecordIDValue)},
					BaseOrganizationID:     core.BaseOrganizationID{OrganizationID: organizationID},
					Domain:                 types.StringValue(testNormalizedProxyRecordDomain),
					TargetCNAME:            types.StringValue(testNormalizedProxyRecordTargetCNAME),
					Status:                 types.StringValue("valid"),
					Message:                types.StringNull(),
					CreatedAt:              types.StringNull(),
					UpdatedAt:              types.StringNull(),
					CreatedBy:              types.Int64Null(),
				}
			},
			response: httpclient.ProxyRecord{
				ID:          testProxyRecordIDValue,
				Domain:      testNormalizedProxyRecordDomain,
				TargetCNAME: testProxyRecordTargetCNAME,
				Status:      "valid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("update must not write to the API, got %s %s", r.Method, r.URL.Path)
				}
				assert.True(t, strings.HasPrefix(r.URL.Path, "/api/organizations/"+testLegacyOrganizationID+"/"),
					"request should target the provider's organization, got %s", r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				require.NoError(t, json.NewEncoder(w).Encode(tc.response))
			}))
			defer server.Close()

			configureLegacyOrganizationResource(t, tc.resource, server, testLegacyOrganizationID)

			// The plan modifier plans an unknown organization_id when state holds "@current"
			// and the configuration leaves it unset.
			nullObject := tftypes.NewValue(tc.schema.Type().TerraformType(ctx), nil)
			state := tfsdk.State{Schema: tc.schema, Raw: nullObject}
			require.False(t, state.Set(ctx, tc.state(types.StringValue(httpclient.CurrentOrganizationAlias))).HasError())
			plan := tfsdk.Plan{Schema: tc.schema, Raw: nullObject}
			require.False(t, plan.Set(ctx, tc.state(types.StringUnknown())).HasError())

			resp := &tfresource.UpdateResponse{State: tfsdk.State{Schema: tc.schema, Raw: nullObject}}
			tc.resource.Update(ctx, tfresource.UpdateRequest{Plan: plan, State: state}, resp)
			require.False(t, resp.Diagnostics.HasError(), "update: %v", resp.Diagnostics)

			var organizationID types.String
			require.False(t, resp.State.GetAttribute(ctx, path.Root("organization_id"), &organizationID).HasError())
			assert.Equal(t, testLegacyOrganizationID, organizationID.ValueString())
		})
	}
}

// A resource whose state holds "@current" may live in another organization than the one it is
// looked up in, so a 404 must fail the operation instead of dropping the resource from state,
// which would make the next apply create a duplicate.
func TestLegacyOrganizationIDNotFoundFailsClosed(t *testing.T) {
	ctx := context.Background()
	schema := RoleMembershipOps{}.Schema()
	nullObject := tftypes.NewValue(schema.Type().TerraformType(ctx), nil)

	newResource := func(t *testing.T) tfresource.Resource {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		t.Cleanup(server.Close)
		res := NewRoleMembership()
		configureLegacyOrganizationResource(t, res, server, testLegacyOrganizationID)
		return res
	}
	stateWith := func(t *testing.T, organizationID types.String) tfsdk.State {
		state := tfsdk.State{Schema: schema, Raw: nullObject}
		require.False(t, state.Set(ctx, legacyRoleMembership(organizationID)).HasError())
		return state
	}

	t.Run("read keeps the resource in state", func(t *testing.T) {
		state := stateWith(t, types.StringValue(httpclient.CurrentOrganizationAlias))
		resp := &tfresource.ReadResponse{State: state}
		newResource(t).Read(ctx, tfresource.ReadRequest{State: state}, resp)

		require.True(t, resp.Diagnostics.HasError())
		assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "terraform apply -refresh=false")
		assert.False(t, resp.State.Raw.IsNull(), "the resource must stay in state")
	})

	t.Run("update keeps the resource in state", func(t *testing.T) {
		state := stateWith(t, types.StringValue(httpclient.CurrentOrganizationAlias))
		plan := tfsdk.Plan{Schema: schema, Raw: nullObject}
		require.False(t, plan.Set(ctx, legacyRoleMembership(types.StringValue(testLegacyOrganizationID))).HasError())
		resp := &tfresource.UpdateResponse{State: state}
		newResource(t).Update(ctx, tfresource.UpdateRequest{Plan: plan, State: state}, resp)

		require.True(t, resp.Diagnostics.HasError())
		assert.False(t, resp.State.Raw.IsNull(), "the resource must stay in state")
	})

	t.Run("delete keeps the resource in state", func(t *testing.T) {
		state := stateWith(t, types.StringValue(httpclient.CurrentOrganizationAlias))
		resp := &tfresource.DeleteResponse{State: state}
		newResource(t).Delete(ctx, tfresource.DeleteRequest{State: state}, resp)

		require.True(t, resp.Diagnostics.HasError())
		assert.False(t, resp.State.Raw.IsNull(), "the resource must stay in state")
	})

	t.Run("read of a real organization id still drops a deleted resource", func(t *testing.T) {
		state := stateWith(t, types.StringValue(testLegacyOrganizationID))
		resp := &tfresource.ReadResponse{State: state}
		newResource(t).Read(ctx, tfresource.ReadRequest{State: state}, resp)

		require.False(t, resp.Diagnostics.HasError(), "read: %v", resp.Diagnostics)
		assert.True(t, resp.State.Raw.IsNull(), "a resource deleted outside Terraform is removed from state")
	})
}

// With no organization_id on the provider or the resource, an update of a resource whose state
// holds "@current" has no organization to call, so it fails before sending any request.
func TestLegacyOrganizationIDUpdateWithoutOrganization(t *testing.T) {
	ctx := context.Background()
	schema := RoleMembershipOps{}.Schema()
	nullObject := tftypes.NewValue(schema.Type().TerraformType(ctx), nil)

	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		t.Errorf("no request expected, got %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()
	res := NewRoleMembership()
	configureLegacyOrganizationResource(t, res, server, "")

	state := tfsdk.State{Schema: schema, Raw: nullObject}
	require.False(t, state.Set(ctx, legacyRoleMembership(types.StringValue(httpclient.CurrentOrganizationAlias))).HasError())
	plan := tfsdk.Plan{Schema: schema, Raw: nullObject}
	require.False(t, plan.Set(ctx, legacyRoleMembership(types.StringUnknown())).HasError())

	resp := &tfresource.UpdateResponse{State: state}
	res.Update(ctx, tfresource.UpdateRequest{Plan: plan, State: state}, resp)

	require.True(t, resp.Diagnostics.HasError())
	assert.Equal(t, "Missing organization_id", resp.Diagnostics.Errors()[0].Summary())
}

func configureLegacyOrganizationResource(t *testing.T, res tfresource.Resource, server *httptest.Server, organizationID string) {
	t.Helper()
	configurable, ok := res.(tfresource.ResourceWithConfigure)
	require.True(t, ok)
	resp := &tfresource.ConfigureResponse{}
	configurable.Configure(context.Background(), tfresource.ConfigureRequest{ProviderData: data.ProviderData{
		Client:                httpclient.NewClient(server.Client(), server.URL, testProxyRecordAPIKey, testProxyRecordClientVersion),
		DefaultOrganizationID: organizationID,
	}}, resp)
	require.False(t, resp.Diagnostics.HasError(), "configure: %v", resp.Diagnostics)
}

func legacyRoleMembership(organizationID types.String) *RoleMembershipTFModel {
	return &RoleMembershipTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue("membership-1")},
		BaseOrganizationID:     core.BaseOrganizationID{OrganizationID: organizationID},
		BaseRoleID:             core.BaseRoleID{RoleID: types.StringValue("role-1")},
		UserUUID:               types.StringValue("user-1"),
		UserEmail:              types.StringNull(),
		JoinedAt:               types.StringNull(),
	}
}
