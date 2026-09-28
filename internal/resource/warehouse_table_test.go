package resource

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testWarehouseTableID  = "01900000-0000-7000-8000-00000000abcd"
	testWarehouseTableURL = "https://bucket.s3.us-east-2.amazonaws.com/orders/*.parquet"
)

func testWarehouseTableModel() WarehouseTableTFModel {
	return WarehouseTableTFModel{
		Name:         types.StringValue("orders"),
		Format:       types.StringValue("Parquet"),
		URLPattern:   types.StringValue(testWarehouseTableURL),
		AccessKey:    types.StringValue("AKIAEXAMPLE"),
		AccessSecret: types.StringValue("shh"),
	}
}

func testWarehouseTableRequest() httpclient.WarehouseTableRequest {
	return httpclient.WarehouseTableRequest{
		Name:       "orders",
		Format:     "Parquet",
		URLPattern: testWarehouseTableURL,
		Credential: httpclient.WarehouseTableCredential{AccessKey: "AKIAEXAMPLE", AccessSecret: "shh"},
	}
}

// PostHog reads the files to derive a table's columns on create but not on
// update, so only a change to which files are read may replace the table.
func TestWarehouseTableSchema_OnlyAChangeToTheFilesReplacesTheTable(t *testing.T) {
	attrs := WarehouseTableOps{}.Schema().Attributes

	for name, wantReplace := range map[string]bool{
		"format":        true,
		"url_pattern":   true,
		"name":          false,
		"access_key":    false,
		"access_secret": false,
	} {
		attr, ok := attrs[name].(schema.StringAttribute)
		require.Truef(t, ok, "%s must be a string attribute", name)

		replaces := false
		for _, m := range attr.PlanModifiers {
			replaces = replaces || strings.Contains(planModifierDescription(t, m), "destroy and recreate")
		}
		assert.Equalf(t, wantReplace, replaces, "%s: replaces the table", name)
	}
}

func TestWarehouseTableSchema_KeepsTheKeysOutOfPlanOutput(t *testing.T) {
	attrs := WarehouseTableOps{}.Schema().Attributes

	for _, name := range []string{"access_key", "access_secret"} {
		assert.Truef(t, attrs[name].IsSensitive(), "%s must be sensitive", name)
	}
}

func TestWarehouseTableBuildCreateRequest(t *testing.T) {
	req, diags := WarehouseTableOps{}.BuildCreateRequest(context.Background(), testWarehouseTableModel())

	require.False(t, diags.HasError(), diags.Errors())
	assert.Equal(t, testWarehouseTableRequest(), req)
}

// An imported table has no keys in state, so the first apply after an import
// is an update, and it is the only way the keys reach PostHog.
func TestWarehouseTableBuildUpdateRequest_SendsTheKeys(t *testing.T) {
	importedState := testWarehouseTableModel()
	importedState.AccessKey = types.StringNull()
	importedState.AccessSecret = types.StringNull()

	req, diags := WarehouseTableOps{}.BuildUpdateRequest(context.Background(), testWarehouseTableModel(), importedState)

	require.False(t, diags.HasError(), diags.Errors())
	assert.Equal(t, testWarehouseTableRequest(), req)
}

func TestWarehouseTableMapResponseToModel_KeepsTheConfiguredKeys(t *testing.T) {
	model := testWarehouseTableModel()
	resp := httpclient.WarehouseTable{
		ID:         testWarehouseTableID,
		Name:       "orders_renamed",
		Format:     "Parquet",
		URLPattern: testWarehouseTableURL,
		CreatedAt:  "2026-09-28T17:00:00Z",
	}

	diags := WarehouseTableOps{}.MapResponseToModel(context.Background(), resp, &model)

	require.False(t, diags.HasError(), diags.Errors())
	want := testWarehouseTableModel()
	want.ID = types.StringValue(testWarehouseTableID)
	want.Name = types.StringValue("orders_renamed")
	want.CreatedAt = types.StringValue("2026-09-28T17:00:00Z")
	assert.Equal(t, want, model)
}
