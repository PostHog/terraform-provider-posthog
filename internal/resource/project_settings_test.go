package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testProjectSettingsProjectID = "123"

// cohortFilterJSON is a single cohort-referencing test_account_filters entry, reused
// across the parse and cohort_name-stripping tests.
const cohortFilterJSON = `[{"key":"id","type":"cohort","value":2,"operator":"in"}]`

func newProjectSettingsModel() ProjectSettingsModel {
	m := ProjectSettingsModel{}
	m.InitializeProjectID(testProjectSettingsProjectID)
	return m
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
	model.AnonymizeIps = types.BoolValue(true)

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
	require.NotNil(t, req.AnonymizeIps)
	assert.True(t, *req.AnonymizeIps)
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
	assert.Nil(t, req.AnonymizeIps, "unset fields must serialize as nil (omitted)")
}

// TestProjectSettingsBuildCreateRequest_EmptyListClears guards the clearing path:
// an explicit empty list must serialize to a non-nil pointer to an empty slice
// (which clears the value server-side), not nil — otherwise clearing app_urls or
// recording_domains would silently become a no-op.
func TestProjectSettingsBuildCreateRequest_EmptyListClears(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	emptyList, d := types.ListValueFrom(context.Background(), types.StringType, []string{})
	require.False(t, d.HasError())
	model.AppURLs = emptyList
	model.RecordingDomains = emptyList

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.AppURLs, "explicit empty app_urls must be a non-nil pointer to clear the value")
	assert.Empty(t, *req.AppURLs)
	require.NotNil(t, req.RecordingDomains, "explicit empty recording_domains must be a non-nil pointer to clear the value")
	assert.Empty(t, *req.RecordingDomains)
}

func networkPayloadCaptureObject(t *testing.T, headers, body bool) types.Object {
	t.Helper()
	obj, d := types.ObjectValueFrom(context.Background(), networkPayloadCaptureAttrTypes, NetworkPayloadCaptureModel{
		RecordHeaders: types.BoolValue(headers),
		RecordBody:    types.BoolValue(body),
	})
	require.False(t, d.HasError())
	return obj
}

func TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureSet(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.CapturePerformanceOptIn = types.BoolValue(true)
	model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, true)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.CapturePerformanceOptIn)
	assert.True(t, *req.CapturePerformanceOptIn)
	require.NotNil(t, req.SessionRecordingNetworkPayloadCaptureConfig)
	require.NotNil(t, req.SessionRecordingNetworkPayloadCaptureConfig.RecordHeaders)
	assert.True(t, *req.SessionRecordingNetworkPayloadCaptureConfig.RecordHeaders)
	require.NotNil(t, req.SessionRecordingNetworkPayloadCaptureConfig.RecordBody)
	assert.True(t, *req.SessionRecordingNetworkPayloadCaptureConfig.RecordBody)
}

// TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureExplicitFalse guards
// the zero-value path: explicit false on both children must survive into the
// serialized PATCH body (camelCase keys), not be dropped by omitempty.
func TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureExplicitFalse(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.NetworkPayloadCapture = networkPayloadCaptureObject(t, false, false)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.SessionRecordingNetworkPayloadCaptureConfig)
	require.NotNil(t, req.SessionRecordingNetworkPayloadCaptureConfig.RecordHeaders)
	assert.False(t, *req.SessionRecordingNetworkPayloadCaptureConfig.RecordHeaders)
	require.NotNil(t, req.SessionRecordingNetworkPayloadCaptureConfig.RecordBody)
	assert.False(t, *req.SessionRecordingNetworkPayloadCaptureConfig.RecordBody)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"session_recording_network_payload_capture_config":{"recordHeaders":false,"recordBody":false}`)
}

// TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureAbsent guards the
// leave-untouched contract: an unset block must be omitted from the PATCH body
// entirely, not sent as null or {} (the API replaces the whole JSON blob).
// TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureWarnsWithoutPerformanceCapture
// guards the no-op warning: payload capture only takes effect when network
// performance capture is enabled, so configuring the block without a known-true
// capture_performance_opt_in must warn (and only then).
func TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureWarnsWithoutPerformanceCapture(t *testing.T) {
	ops := ProjectSettingsOps{}

	t.Run("warns when capture_performance_opt_in is unset", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)

		_, diags := ops.BuildCreateRequest(context.Background(), model)

		assert.False(t, diags.HasError())
		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Summary(), "capture_performance_opt_in")
	})

	t.Run("warns when capture_performance_opt_in is false", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.CapturePerformanceOptIn = types.BoolValue(false)
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)

		_, diags := ops.BuildCreateRequest(context.Background(), model)

		require.Equal(t, 1, diags.WarningsCount())
	})

	t.Run("no warning when capture_performance_opt_in is true", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.CapturePerformanceOptIn = types.BoolValue(true)
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)

		_, diags := ops.BuildCreateRequest(context.Background(), model)

		assert.Equal(t, 0, diags.WarningsCount())
	})

	t.Run("no warning when the block is not configured", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.CapturePerformanceOptIn = types.BoolValue(false)

		_, diags := ops.BuildCreateRequest(context.Background(), model)

		assert.Equal(t, 0, diags.WarningsCount())
	})
}

func TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureAbsent(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.HeatmapsOptIn = types.BoolValue(true)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	assert.Nil(t, req.SessionRecordingNetworkPayloadCaptureConfig)
	assert.Nil(t, req.CapturePerformanceOptIn)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.NotContains(t, string(body), "session_recording_network_payload_capture_config")
	assert.NotContains(t, string(body), "capture_performance_opt_in")
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
		AnonymizeIps:               util.BoolPtr(true),
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
	assert.True(t, model.AnonymizeIps.ValueBool())
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
	assert.True(t, model.AnonymizeIps.IsNull())
}

func TestProjectSettingsMapResponseToModel_NetworkPayloadCapturePresent(t *testing.T) {
	ops := ProjectSettingsOps{}
	resp := httpclient.Environment{
		ID:                      123,
		CapturePerformanceOptIn: util.BoolPtr(true),
		SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
			RecordHeaders: util.BoolPtr(true),
			RecordBody:    util.BoolPtr(false),
		},
	}

	model := newProjectSettingsModel()
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	assert.True(t, model.CapturePerformanceOptIn.ValueBool())
	require.False(t, model.NetworkPayloadCapture.IsNull())
	var cfg NetworkPayloadCaptureModel
	require.False(t, model.NetworkPayloadCapture.As(context.Background(), &cfg, basetypes.ObjectAsOptions{}).HasError())
	assert.True(t, cfg.RecordHeaders.ValueBool())
	require.False(t, cfg.RecordBody.IsNull(), "explicit false must map to a known value, not null")
	assert.False(t, cfg.RecordBody.ValueBool())
}

// TestProjectSettingsMapResponseToModel_NetworkPayloadCaptureNull covers the
// server default: PostHog stores null until the setting is first written, and
// that must map to a null object (not an object of null bools).
func TestProjectSettingsMapResponseToModel_NetworkPayloadCaptureNull(t *testing.T) {
	ops := ProjectSettingsOps{}
	resp := httpclient.Environment{ID: 123}

	model := newProjectSettingsModel()
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	assert.True(t, model.CapturePerformanceOptIn.IsNull())
	assert.True(t, model.NetworkPayloadCapture.IsNull())
}

// TestProjectSettingsMapResponseToModel_DivergenceWarning verifies the
// plan-gated warning names the new attributes when PostHog returns different
// values than the ones the user configured.
func TestProjectSettingsMapResponseToModel_DivergenceWarning(t *testing.T) {
	ops := ProjectSettingsOps{}

	t.Run("network payload capture ignored (server null)", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)

		diags := ops.MapResponseToModel(context.Background(), httpclient.Environment{ID: 123}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "session_recording_network_payload_capture_config")
	})

	t.Run("network payload capture child value diverged", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, true)

		diags := ops.MapResponseToModel(context.Background(), httpclient.Environment{
			ID: 123,
			SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
				RecordHeaders: util.BoolPtr(true),
				RecordBody:    util.BoolPtr(false),
			},
		}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "session_recording_network_payload_capture_config")
	})

	t.Run("capture_performance_opt_in diverged", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.CapturePerformanceOptIn = types.BoolValue(true)

		diags := ops.MapResponseToModel(context.Background(), httpclient.Environment{
			ID:                      123,
			CapturePerformanceOptIn: util.BoolPtr(false),
		}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "capture_performance_opt_in")
	})

	t.Run("anonymize_ips diverged", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.AnonymizeIps = types.BoolValue(true)

		diags := ops.MapResponseToModel(context.Background(), httpclient.Environment{
			ID:           123,
			AnonymizeIps: util.BoolPtr(false),
		}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "anonymize_ips")
	})

	t.Run("unparseable configured object counts as diverged", func(t *testing.T) {
		model := newProjectSettingsModel()
		// Wrong child type: As() into NetworkPayloadCaptureModel must fail, and
		// the attribute must still be named in the warning (not suppressed).
		model.NetworkPayloadCapture = types.ObjectValueMust(
			map[string]attr.Type{"record_headers": types.StringType, "record_body": types.BoolType},
			map[string]attr.Value{"record_headers": types.StringValue("not-a-bool"), "record_body": types.BoolValue(true)},
		)

		diags := ops.MapResponseToModel(context.Background(), httpclient.Environment{
			ID: 123,
			SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
				RecordHeaders: util.BoolPtr(true),
				RecordBody:    util.BoolPtr(true),
			},
		}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "session_recording_network_payload_capture_config")
	})

	t.Run("no warning when server matches", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.CapturePerformanceOptIn = types.BoolValue(true)
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)

		diags := ops.MapResponseToModel(context.Background(), httpclient.Environment{
			ID:                      123,
			CapturePerformanceOptIn: util.BoolPtr(true),
			SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
				RecordHeaders: util.BoolPtr(true),
				RecordBody:    util.BoolPtr(false),
			},
		}, &model)

		assert.Equal(t, 0, diags.WarningsCount())
	})
}

func TestProjectSettingsBuildCreateRequest_TestAccountFilters(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.TestAccountFilters = jsontypes.NewNormalizedValue(cohortFilterJSON)
	model.TestAccountFiltersDefaultChecked = types.BoolValue(true)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.TestAccountFilters)
	require.Len(t, *req.TestAccountFilters, 1)
	filter, ok := (*req.TestAccountFilters)[0].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "cohort", filter["type"])
	assert.Equal(t, float64(2), filter["value"]) // JSON numbers decode to float64
	require.NotNil(t, req.TestAccountFiltersDefaultChecked)
	assert.True(t, *req.TestAccountFiltersDefaultChecked)
}

// TestProjectSettingsBuildCreateRequest_TestAccountFiltersEmptyClears guards the
// clearing path: an explicit empty array must serialize to a non-nil pointer to an
// empty slice so the filters can be cleared server-side rather than left untouched.
func TestProjectSettingsBuildCreateRequest_TestAccountFiltersEmptyClears(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.TestAccountFilters = jsontypes.NewNormalizedValue(`[]`)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.TestAccountFilters, "explicit empty array must be a non-nil pointer to clear the value")
	assert.Empty(t, *req.TestAccountFilters)
}

func TestProjectSettingsBuildCreateRequest_TestAccountFiltersUnset(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.HeatmapsOptIn = types.BoolValue(true)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	assert.Nil(t, req.TestAccountFilters)
	assert.Nil(t, req.TestAccountFiltersDefaultChecked)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.NotContains(t, string(body), "test_account_filters")
}

func TestProjectSettingsBuildCreateRequest_TestAccountFiltersInvalidJSON(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.TestAccountFilters = jsontypes.NewNormalizedValue(`{not valid`)

	_, diags := ops.BuildCreateRequest(context.Background(), model)

	require.True(t, diags.HasError())
	assert.Contains(t, diags.Errors()[0].Summary(), "test_account_filters")
}

// TestProjectSettingsMapResponseToModel_TestAccountFiltersStripsCohortName is the
// key drift proof: PostHog injects a server-computed cohort_name into a
// cohort-referencing filter object, and it must be stripped from state so the
// configured filter round-trips without a perpetual diff (same class as #136).
func TestProjectSettingsMapResponseToModel_TestAccountFiltersStripsCohortName(t *testing.T) {
	ops := ProjectSettingsOps{}
	resp := httpclient.Environment{
		ID: 123,
		TestAccountFilters: &[]interface{}{
			map[string]interface{}{
				"key":         "id",
				"type":        "cohort",
				"value":       float64(2),
				"operator":    "in",
				"cohort_name": "Internal users", // server-injected enrichment
			},
		},
		TestAccountFiltersDefaultChecked: util.BoolPtr(true),
	}

	model := newProjectSettingsModel()
	// The user configured the filter WITHOUT cohort_name.
	model.TestAccountFilters = jsontypes.NewNormalizedValue(cohortFilterJSON)
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	require.False(t, model.TestAccountFilters.IsNull())

	// State must match the user config semantically (cohort_name stripped): no drift.
	stateEqualsConfig, d := model.TestAccountFilters.StringSemanticEquals(
		context.Background(),
		jsontypes.NewNormalizedValue(cohortFilterJSON),
	)
	assert.False(t, d.HasError())
	assert.True(t, stateEqualsConfig, "cohort_name must be stripped so the filter round-trips without drift")
	assert.NotContains(t, model.TestAccountFilters.ValueString(), "cohort_name")

	assert.True(t, model.TestAccountFiltersDefaultChecked.ValueBool())
}

func TestProjectSettingsMapResponseToModel_TestAccountFiltersNil(t *testing.T) {
	ops := ProjectSettingsOps{}
	resp := httpclient.Environment{ID: 123}

	model := newProjectSettingsModel()
	diags := ops.MapResponseToModel(context.Background(), resp, &model)

	assert.False(t, diags.HasError())
	assert.True(t, model.TestAccountFilters.IsNull())
	assert.True(t, model.TestAccountFiltersDefaultChecked.IsNull())
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
