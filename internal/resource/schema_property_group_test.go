package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSPGResourceProjectID  = "project-123"
	testSPGResourceID         = "spg-123"
	testSPGResourceCollection = "/api/projects/project-123/schema_property_groups/"
	testSPGResourcePathFull   = "/api/projects/project-123/schema_property_groups/spg-123/"
)

func testPropertyObject(t *testing.T, name, propertyType string, isRequired bool) attr.Value {
	obj, diags := types.ObjectValue(schemaPropertyAttrTypes, map[string]attr.Value{
		"name":                 types.StringValue(name),
		"property_type":        types.StringValue(propertyType),
		"is_required":          types.BoolValue(isRequired),
		"is_optional_in_types": types.BoolValue(false),
		"description":          types.StringNull(),
	})
	require.False(t, diags.HasError())
	return obj
}

func testPropertiesSet(t *testing.T, elems ...attr.Value) types.Set {
	set, diags := types.SetValue(types.ObjectType{AttrTypes: schemaPropertyAttrTypes}, elems)
	require.False(t, diags.HasError())
	return set
}

func TestSchemaPropertyGroupMetadataAndSchema(t *testing.T) {
	require.NotNil(t, NewSchemaPropertyGroup())

	ops := SchemaPropertyGroupOps{}
	assert.Equal(t, "Schema Property Group", ops.ResourceName())

	s := ops.Schema()
	nameAttr, ok := s.Attributes["name"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, nameAttr.Required)

	propsAttr, ok := s.Attributes["properties"].(schema.SetNestedAttribute)
	require.True(t, ok)
	assert.True(t, propsAttr.Optional)

	typeAttr, ok := propsAttr.NestedObject.Attributes["property_type"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, typeAttr.Required)
}

func TestSchemaPropertyGroupBuildCreateRequest(t *testing.T) {
	ops := SchemaPropertyGroupOps{}
	model := SchemaPropertyGroupTFModel{
		Name:        types.StringValue("Checkout"),
		Description: types.StringValue("desc"),
		Properties: testPropertiesSet(t,
			testPropertyObject(t, "cart_value", "Numeric", true),
		),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())
	assert.Equal(t, "Checkout", *req.Name)
	assert.Equal(t, "desc", *req.Description)
	require.NotNil(t, req.Properties)
	require.Len(t, *req.Properties, 1)
	prop := (*req.Properties)[0]
	assert.Equal(t, "cart_value", prop.Name)
	assert.Equal(t, "Numeric", prop.PropertyType)
	require.NotNil(t, prop.IsRequired)
	assert.True(t, *prop.IsRequired)
	assert.Nil(t, prop.ID, "requests must not send property ids (replace-all contract)")
}

func TestSchemaPropertyGroupBuildCreateRequestNoProperties(t *testing.T) {
	ops := SchemaPropertyGroupOps{}
	model := SchemaPropertyGroupTFModel{
		Name:       types.StringValue("Checkout"),
		Properties: types.SetNull(types.ObjectType{AttrTypes: schemaPropertyAttrTypes}),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())
	assert.Nil(t, req.Properties, "null properties must be omitted, not sent as an empty replace-all list")
	assert.Nil(t, req.Description)
}

func TestSchemaPropertyGroupBuildUpdateRequestClears(t *testing.T) {
	ops := SchemaPropertyGroupOps{}
	state := SchemaPropertyGroupTFModel{
		Name:        types.StringValue("Checkout"),
		Description: types.StringValue("desc"),
		Properties: testPropertiesSet(t,
			testPropertyObject(t, "cart_value", "Numeric", true),
		),
	}
	plan := SchemaPropertyGroupTFModel{
		Name:        types.StringValue("Checkout"),
		Description: types.StringNull(),
		Properties:  types.SetNull(types.ObjectType{AttrTypes: schemaPropertyAttrTypes}),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError())
	require.NotNil(t, req.Description)
	assert.Equal(t, "", *req.Description, "removed description must be cleared with empty string")
	require.NotNil(t, req.Properties)
	assert.Empty(t, *req.Properties, "removed properties must be cleared with empty list")
}

func TestSchemaPropertyGroupMapResponseToModel(t *testing.T) {
	ops := SchemaPropertyGroupOps{}
	required := true
	desc := "group desc"
	resp := httpclient.SchemaPropertyGroup{
		ID:          testSPGResourceID,
		Name:        "Checkout",
		Description: &desc,
		Properties: []httpclient.SchemaPropertyGroupProperty{
			{Name: "cart_value", PropertyType: "Numeric", IsRequired: &required},
		},
	}

	model := SchemaPropertyGroupTFModel{
		Properties: testPropertiesSet(t, testPropertyObject(t, "cart_value", "Numeric", true)),
	}
	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())
	assert.Equal(t, testSPGResourceID, model.ID.ValueString())
	assert.Equal(t, "Checkout", model.Name.ValueString())
	assert.Equal(t, "group desc", model.Description.ValueString())

	var props []schemaPropertyTFModel
	require.False(t, model.Properties.ElementsAs(context.Background(), &props, false).HasError())
	require.Len(t, props, 1)
	assert.Equal(t, "cart_value", props[0].Name.ValueString())
	assert.True(t, props[0].IsRequired.ValueBool())
	assert.False(t, props[0].IsOptionalInTypes.ValueBool())

	// no properties on server + null in config stays null (not empty set)
	nullModel := SchemaPropertyGroupTFModel{
		Properties: types.SetNull(types.ObjectType{AttrTypes: schemaPropertyAttrTypes}),
	}
	diags = ops.MapResponseToModel(context.Background(), httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Empty"}, &nullModel)
	require.False(t, diags.HasError())
	assert.True(t, nullModel.Properties.IsNull())

	// user cleared all properties: model non-null, server returns none — state
	// must become an empty set, not null, to match the configured empty list
	clearedModel := SchemaPropertyGroupTFModel{
		Properties: testPropertiesSet(t),
	}
	diags = ops.MapResponseToModel(context.Background(), httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Cleared"}, &clearedModel)
	require.False(t, diags.HasError())
	require.False(t, clearedModel.Properties.IsNull(), "explicit empty must stay an empty set, not become null")
	assert.Empty(t, clearedModel.Properties.Elements())
}

func TestSchemaPropertyGroupCRUDWrapperMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == testSPGResourceCollection:
			_ = json.NewEncoder(w).Encode(httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Checkout"})
		case r.Method == http.MethodGet && r.URL.Path == testSPGResourcePathFull:
			_ = json.NewEncoder(w).Encode(httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Checkout"})
		case r.Method == http.MethodPatch && r.URL.Path == testSPGResourcePathFull:
			_ = json.NewEncoder(w).Encode(httpclient.SchemaPropertyGroup{ID: testSPGResourceID, Name: "Checkout"})
		case r.Method == http.MethodDelete && r.URL.Path == testSPGResourcePathFull:
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	ops := SchemaPropertyGroupOps{}
	model := SchemaPropertyGroupTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testSPGResourceID)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testSPGResourceProjectID)},
	}
	name := "Checkout"
	req := httpclient.SchemaPropertyGroupRequest{Name: &name}

	created, err := ops.Create(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, testSPGResourceID, created.ID)

	read, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSPGResourceID, read.ID)

	_, status, err = ops.Update(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)

	status, err = ops.Delete(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusNoContent), status)
}
