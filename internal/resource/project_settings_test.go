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
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testProjectSettingsProjectID = "123"

func newProjectSettingsModel() ProjectSettingsModel {
	m := ProjectSettingsModel{}
	m.InitializeProjectID(testProjectSettingsProjectID)
	return m
}

func TestProjectSettingsResourceNameAndSchema(t *testing.T) {
	ops := ProjectSettingsOps{}
	s := ops.Schema()

	assert.Equal(t, "project_settings", ops.ResourceName())

	for _, name := range []string{
		"heatmaps_opt_in", "autocapture_exceptions_opt_in", "session_recording_opt_in",
		"surveys_opt_in", "autocapture_web_vitals_opt_in",
	} {
		attr, ok := s.Attributes[name].(schema.BoolAttribute)
		require.Truef(t, ok, "%s must be a bool attribute", name)
		assert.Truef(t, attr.Optional, "%s should be optional", name)
		assert.Truef(t, attr.Computed, "%s should be computed", name)
	}

	hashAttr, ok := s.Attributes["cookieless_server_hash_mode"].(schema.Int64Attribute)
	require.True(t, ok, "cookieless_server_hash_mode must be an int64 attribute")
	assert.True(t, hashAttr.Optional)
	assert.True(t, hashAttr.Computed)

	for _, name := range []string{"app_urls", "recording_domains"} {
		attr, ok := s.Attributes[name].(schema.ListAttribute)
		require.Truef(t, ok, "%s must be a list attribute", name)
		assert.Truef(t, attr.Optional, "%s should be optional", name)
		assert.Truef(t, attr.Computed, "%s should be computed", name)
		assert.Equalf(t, types.StringType, attr.ElementType, "%s should be a list of strings", name)
	}
}

func TestProjectSettingsBuildCreateRequest_AllFields(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.HeatmapsOptIn = types.BoolValue(true)
	model.AutocaptureExceptionsOptIn = types.BoolValue(false)
	model.SessionRecordingOptIn = types.BoolValue(true)
	model.SurveysOptIn = types.BoolValue(false)
	model.CookielessServerHashMode = types.Int64Value(2)
	model.AutocaptureWebVitalsOptIn = types.BoolValue(true)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.HeatmapsOptIn)
	assert.True(t, *req.HeatmapsOptIn)
	require.NotNil(t, req.AutocaptureExceptionsOptIn)
	assert.False(t, *req.AutocaptureExceptionsOptIn)
	require.NotNil(t, req.SessionRecordingOptIn)
	assert.True(t, *req.SessionRecordingOptIn)
	require.NotNil(t, req.SurveysOptIn)
	assert.False(t, *req.SurveysOptIn)
	require.NotNil(t, req.CookielessServerHashMode)
	assert.Equal(t, int64(2), *req.CookielessServerHashMode)
	require.NotNil(t, req.AutocaptureWebVitalsOptIn)
	assert.True(t, *req.AutocaptureWebVitalsOptIn)
}

// TestProjectSettingsBuildCreateRequest_ZeroValues guards the zero-value path:
// explicit false / 0 must be sent (non-nil pointer), not treated as unset.
func TestProjectSettingsBuildCreateRequest_ZeroValues(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.HeatmapsOptIn = types.BoolValue(false)
	model.SessionRecordingOptIn = types.BoolValue(false)
	model.CookielessServerHashMode = types.Int64Value(0)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.HeatmapsOptIn, "explicit false must produce a non-nil pointer")
	assert.False(t, *req.HeatmapsOptIn)
	require.NotNil(t, req.SessionRecordingOptIn)
	assert.False(t, *req.SessionRecordingOptIn)
	require.NotNil(t, req.CookielessServerHashMode, "explicit 0 must produce a non-nil pointer")
	assert.Equal(t, int64(0), *req.CookielessServerHashMode)
}

func TestProjectSettingsBuildCreateRequest_SubsetSet(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	// Only heatmaps configured; everything else null/unknown.
	model.HeatmapsOptIn = types.BoolValue(true)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.HeatmapsOptIn)
	assert.True(t, *req.HeatmapsOptIn)
	assert.Nil(t, req.AutocaptureExceptionsOptIn, "unset fields must serialize as nil (omitted)")
	assert.Nil(t, req.SessionRecordingOptIn)
	assert.Nil(t, req.SurveysOptIn)
	assert.Nil(t, req.CookielessServerHashMode)
	assert.Nil(t, req.AutocaptureWebVitalsOptIn)
}

func TestProjectSettingsBuildUpdateRequest(t *testing.T) {
	ops := ProjectSettingsOps{}
	plan := newProjectSettingsModel()
	plan.SurveysOptIn = types.BoolValue(true)

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, newProjectSettingsModel())

	assert.False(t, diags.HasError())
	require.NotNil(t, req.SurveysOptIn)
	assert.True(t, *req.SurveysOptIn)
}

func TestProjectSettingsMapResponseToModel_AllFields(t *testing.T) {
	ops := ProjectSettingsOps{}
	resp := httpclient.Environment{
		ID:                         123,
		HeatmapsOptIn:              util.BoolPtr(true),
		AutocaptureExceptionsOptIn: util.BoolPtr(false),
		SessionRecordingOptIn:      util.BoolPtr(true),
		SurveysOptIn:               util.BoolPtr(false),
		CookielessServerHashMode:   util.Int64Ptr(1),
		AutocaptureWebVitalsOptIn:  util.BoolPtr(true),
	}

	model := newProjectSettingsModel()
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	assert.Equal(t, testProjectSettingsProjectID, model.ID.ValueString())
	assert.True(t, model.HeatmapsOptIn.ValueBool())
	assert.False(t, model.AutocaptureExceptionsOptIn.ValueBool())
	assert.True(t, model.SessionRecordingOptIn.ValueBool())
	assert.False(t, model.SurveysOptIn.ValueBool())
	assert.Equal(t, int64(1), model.CookielessServerHashMode.ValueInt64())
	assert.True(t, model.AutocaptureWebVitalsOptIn.ValueBool())
}

// TestProjectSettingsMapResponseToModel_ZeroValues ensures a returned 0 / false
// round-trips as a known value, not null (distinct from the nil/unset case).
func TestProjectSettingsMapResponseToModel_ZeroValues(t *testing.T) {
	ops := ProjectSettingsOps{}
	resp := httpclient.Environment{
		ID:                       123,
		HeatmapsOptIn:            util.BoolPtr(false),
		CookielessServerHashMode: util.Int64Ptr(0),
	}

	model := newProjectSettingsModel()
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	require.False(t, model.HeatmapsOptIn.IsNull(), "explicit false must map to a known value, not null")
	assert.False(t, model.HeatmapsOptIn.ValueBool())
	require.False(t, model.CookielessServerHashMode.IsNull(), "explicit 0 must map to a known value, not null")
	assert.Equal(t, int64(0), model.CookielessServerHashMode.ValueInt64())
}

func TestProjectSettingsMapResponseToModel_NilFieldsBecomeNull(t *testing.T) {
	ops := ProjectSettingsOps{}
	resp := httpclient.Environment{ID: 123}

	model := newProjectSettingsModel()
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	assert.True(t, model.HeatmapsOptIn.IsNull())
	assert.True(t, model.AutocaptureExceptionsOptIn.IsNull())
	assert.True(t, model.SessionRecordingOptIn.IsNull())
	assert.True(t, model.SurveysOptIn.IsNull())
	assert.True(t, model.CookielessServerHashMode.IsNull())
	assert.True(t, model.AutocaptureWebVitalsOptIn.IsNull())
}

func TestProjectSettingsDeleteIsNoOp(t *testing.T) {
	ops := ProjectSettingsOps{}
	// A nil client would panic if any HTTP call were attempted; the no-op must not call out.
	status, err := ops.Delete(context.Background(), httpclient.PosthogClient{}, newProjectSettingsModel())

	require.NoError(t, err)
	assert.Zero(t, int(status))
}

func TestProjectSettingsHasValidID(t *testing.T) {
	model := newProjectSettingsModel()
	assert.True(t, model.HasValidID())

	empty := ProjectSettingsModel{}
	assert.False(t, empty.HasValidID())
}

func TestProjectSettingsCRUDFlow(t *testing.T) {
	// Verifies Create (PATCH) and Update (PATCH) hit the environment endpoint and
	// Read (GET) maps the response back.
	envPath := "/api/environments/" + testProjectSettingsProjectID + "/"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, envPath, r.URL.Path)
		switch r.Method {
		case http.MethodPatch, http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(httpclient.Environment{
				ID:            123,
				HeatmapsOptIn: util.BoolPtr(true),
			}))
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.HeatmapsOptIn = types.BoolValue(true)

	created, err := ops.Create(context.Background(), client, model, httpclient.EnvironmentSettingsRequest{HeatmapsOptIn: util.BoolPtr(true)})
	require.NoError(t, err)
	require.NotNil(t, created.HeatmapsOptIn)
	assert.True(t, *created.HeatmapsOptIn)

	read, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	require.NotNil(t, read.HeatmapsOptIn)
	assert.True(t, *read.HeatmapsOptIn)

	updated, status, err := ops.Update(context.Background(), client, model, httpclient.EnvironmentSettingsRequest{HeatmapsOptIn: util.BoolPtr(true)})
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	require.NotNil(t, updated.HeatmapsOptIn)
}
