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

			configurable, ok := tc.resource.(tfresource.ResourceWithConfigure)
			require.True(t, ok)
			configureResp := &tfresource.ConfigureResponse{}
			configurable.Configure(ctx, tfresource.ConfigureRequest{ProviderData: data.ProviderData{
				Client:                httpclient.NewClient(server.Client(), server.URL, testProxyRecordAPIKey, testProxyRecordClientVersion),
				DefaultOrganizationID: testLegacyOrganizationID,
			}}, configureResp)
			require.False(t, configureResp.Diagnostics.HasError(), "configure: %v", configureResp.Diagnostics)

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
