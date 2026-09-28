package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testWarehouseTableID  = "01900000-0000-7000-8000-000000000001"
	testWarehouseTableURL = "https://bucket.s3.amazonaws.com/orders/*.parquet"
)

func testWarehouseTableModel() WarehouseTableTFModel {
	return WarehouseTableTFModel{
		Name:                 types.StringValue("s3_orders"),
		Format:               types.StringValue("Parquet"),
		URLPattern:           types.StringValue(testWarehouseTableURL),
		AccessKey:            types.StringValue("AKIA"),
		AccessSecret:         types.StringValue("secret"),
		CSVAllowDoubleQuotes: types.BoolNull(),
	}
}

func TestWarehouseTableResourceNameAndSchema(t *testing.T) {
	ops := WarehouseTableOps{}
	s := ops.Schema()

	assert.Equal(t, "warehouse_table", ops.ResourceName())

	// PostHog infers columns only on create, so an in-place change to either
	// attribute would leave a stale column snapshot.
	for _, name := range []string{"format", "url_pattern"} {
		attr, ok := s.Attributes[name].(schema.StringAttribute)
		require.Truef(t, ok, "%s must be a string attribute", name)
		require.Lenf(t, attr.PlanModifiers, 1, "%s should carry exactly one plan modifier (RequiresReplace)", name)
		assert.Containsf(t, planModifierDescription(t, attr.PlanModifiers[0]), "destroy and recreate", "%s plan modifier should be RequiresReplace", name)
	}

	nameAttr, ok := s.Attributes["name"].(schema.StringAttribute)
	require.True(t, ok)
	assert.Empty(t, nameAttr.PlanModifiers, "name can be updated in place")

	for _, name := range []string{"access_key", "access_secret"} {
		attr, ok := s.Attributes[name].(schema.StringAttribute)
		require.Truef(t, ok, "%s must be a string attribute", name)
		assert.Truef(t, attr.Required, "%s must be required: PostHog rejects a table without a credential", name)
		assert.Truef(t, attr.Sensitive, "%s must be sensitive", name)
	}
}

func TestWarehouseTableBuildCreateRequest(t *testing.T) {
	ops := WarehouseTableOps{}
	model := testWarehouseTableModel()
	model.Format = types.StringValue("CSVWithNames")
	model.CSVAllowDoubleQuotes = types.BoolValue(true)

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())

	assert.Equal(t, "s3_orders", req["name"])
	assert.Equal(t, "CSVWithNames", req["format"])
	assert.Equal(t, testWarehouseTableURL, req["url_pattern"])
	assert.Equal(t, map[string]any{"access_key": "AKIA", "access_secret": "secret"}, req["credential"])
	assert.Equal(t, map[string]any{"csv_allow_double_quotes": true}, req["options"])
}

func TestWarehouseTableBuildCreateRequest_NoOptions(t *testing.T) {
	ops := WarehouseTableOps{}

	req, diags := ops.BuildCreateRequest(context.Background(), testWarehouseTableModel())
	require.False(t, diags.HasError())
	assert.Equal(t, map[string]any{}, req["options"])
}

func TestWarehouseTableBuildUpdateRequest_UnchangedCredential(t *testing.T) {
	ops := WarehouseTableOps{}
	state := testWarehouseTableModel()
	plan := testWarehouseTableModel()
	plan.Name = types.StringValue("s3_orders_v2")

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError())

	assert.Equal(t, "s3_orders_v2", req["name"])
	_, hasCredential := req["credential"]
	assert.False(t, hasCredential, "PATCH must not resend an unchanged credential")
	_, hasURL := req["url_pattern"]
	assert.False(t, hasURL, "url_pattern is RequiresReplace and must not be sent on update")
}

func TestWarehouseTableBuildUpdateRequest_RotatedSecret(t *testing.T) {
	ops := WarehouseTableOps{}
	state := testWarehouseTableModel()
	plan := testWarehouseTableModel()
	plan.AccessSecret = types.StringValue("new-secret")

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError())
	assert.Equal(t, map[string]any{"access_secret": "new-secret"}, req["credential"])
}

func TestWarehouseTableBuildUpdateRequest_AfterImport(t *testing.T) {
	ops := WarehouseTableOps{}
	state := testWarehouseTableModel()
	state.AccessKey = types.StringNull()
	state.AccessSecret = types.StringNull()

	req, diags := ops.BuildUpdateRequest(context.Background(), testWarehouseTableModel(), state)
	require.False(t, diags.HasError())
	assert.Equal(t, map[string]any{"access_key": "AKIA", "access_secret": "secret"}, req["credential"])
}

func TestWarehouseTableMapResponseToModel(t *testing.T) {
	ops := WarehouseTableOps{}
	model := testWarehouseTableModel()

	resp := httpclient.WarehouseTable{
		ID:         testWarehouseTableID,
		Name:       "s3_orders",
		HogQLName:  util.StringPtr("s3_orders"),
		Format:     "Parquet",
		URLPattern: testWarehouseTableURL,
		Options:    map[string]any{"csv_allow_double_quotes": false},
		CreatedAt:  util.StringPtr("2026-01-01T00:00:00Z"),
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())

	assert.Equal(t, testWarehouseTableID, model.ID.ValueString())
	assert.Equal(t, "s3_orders", model.HogQLName.ValueString())
	assert.Equal(t, types.BoolValue(false), model.CSVAllowDoubleQuotes)
	// PostHog never returns the credential, so the planned values must stay.
	assert.Equal(t, "AKIA", model.AccessKey.ValueString())
	assert.Equal(t, "secret", model.AccessSecret.ValueString())
}

func TestWarehouseTableMapResponseToModel_NoOptions(t *testing.T) {
	ops := WarehouseTableOps{}
	model := testWarehouseTableModel()

	diags := ops.MapResponseToModel(context.Background(), httpclient.WarehouseTable{ID: testWarehouseTableID}, &model)
	require.False(t, diags.HasError())
	assert.True(t, model.CSVAllowDoubleQuotes.IsNull())
}
