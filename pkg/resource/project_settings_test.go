package resource

import (
	"context"
	"encoding/json"
	"maps"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
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
	// Only heatmaps configured; omitted computed settings are unknown.
	model.HeatmapsOptIn = types.BoolValue(true)
	model.AutocaptureExceptionsOptIn = types.BoolUnknown()
	model.SessionRecordingOptIn = types.BoolUnknown()
	model.SurveysOptIn = types.BoolUnknown()
	model.CookielessServerHashMode = types.Int64Unknown()
	model.AutocaptureWebVitalsOptIn = types.BoolUnknown()
	model.AnonymizeIps = types.BoolUnknown()
	model.CapturePerformanceOptIn = types.BoolUnknown()
	model.NetworkPayloadCapture = types.ObjectUnknown(networkPayloadCaptureAttrTypes)
	model.AppURLs = types.ListUnknown(types.StringType)
	model.RecordingDomains = types.ListUnknown(types.StringType)
	model.TestAccountFilters = core.NewNormalizedJSONUnknown()
	model.TestAccountFiltersDefaultChecked = types.BoolUnknown()

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
	assert.Nil(t, req.SessionRecordingNetworkPayloadCaptureConfig)
	assert.Nil(t, req.CapturePerformanceOptIn)
	assert.Nil(t, req.AppURLs)
	assert.Nil(t, req.RecordingDomains)
	assert.Nil(t, req.TestAccountFilters)
	assert.Nil(t, req.TestAccountFiltersDefaultChecked)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.NotContains(t, string(body), "session_recording_network_payload_capture_config")
	assert.NotContains(t, string(body), "capture_performance_opt_in")
	assert.NotContains(t, string(body), "test_account_filters")
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

func TestProjectSettingsBuildCreateRequest_NetworkPayloadCaptureWarnsWithoutPerformanceCapture(t *testing.T) {
	ops := ProjectSettingsOps{}

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
		model.NetworkPayloadCapture = types.ObjectUnknown(networkPayloadCaptureAttrTypes)

		_, diags := ops.BuildCreateRequest(context.Background(), model)

		assert.Equal(t, 0, diags.WarningsCount())
	})
}

func TestProjectSettingsBuildUpdateRequest_NetworkCaptureWarning(t *testing.T) {
	for _, test := range []struct {
		name        string
		prior       types.Bool
		planned     types.Bool
		warnings    int
		wantCapture *bool
	}{
		{"enabled in state", types.BoolValue(true), types.BoolUnknown(), 0, nil},
		{"disabled in state", types.BoolValue(false), types.BoolUnknown(), 1, nil},
		{"unset in state", types.BoolNull(), types.BoolUnknown(), 1, nil},
		{"explicitly disabled", types.BoolValue(true), types.BoolValue(false), 1, new(false)},
	} {
		t.Run(test.name, func(t *testing.T) {
			plan := newProjectSettingsModel()
			plan.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)
			plan.CapturePerformanceOptIn = test.planned
			plan.HeatmapsOptIn = types.BoolValue(true)
			plan.AutocaptureExceptionsOptIn = types.BoolUnknown()
			plan.SessionRecordingOptIn = types.BoolUnknown()
			plan.SurveysOptIn = types.BoolUnknown()
			plan.CookielessServerHashMode = types.Int64Unknown()
			plan.AutocaptureWebVitalsOptIn = types.BoolUnknown()
			plan.AnonymizeIps = types.BoolUnknown()
			plan.AppURLs = types.ListUnknown(types.StringType)
			plan.RecordingDomains = types.ListUnknown(types.StringType)
			plan.TestAccountFilters = core.NewNormalizedJSONUnknown()
			plan.TestAccountFiltersDefaultChecked = types.BoolUnknown()
			state := newProjectSettingsModel()
			state.CapturePerformanceOptIn = test.prior

			req, diags := (ProjectSettingsOps{}).BuildUpdateRequest(context.Background(), plan, state)

			require.False(t, diags.HasError())
			require.Equal(t, test.warnings, diags.WarningsCount())
			if test.warnings > 0 {
				assert.Contains(t, diags.Warnings()[0].Summary(), "capture_performance_opt_in")
			}
			assert.Equal(t, httpclient.EnvironmentSettingsRequest{
				HeatmapsOptIn:           new(true),
				CapturePerformanceOptIn: test.wantCapture,
				SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
					RecordHeaders: new(true),
					RecordBody:    new(false),
				},
			}, req.EnvironmentSettingsRequest)
		})
	}
}

func TestProjectSettingsModifyPlan(t *testing.T) {
	ctx := context.Background()
	ops := ProjectSettingsOps{}
	s := ops.Schema()
	objectType := s.Type().TerraformType(ctx).(tftypes.Object)
	for _, test := range []struct {
		name          string
		attribute     string
		value         tftypes.Value
		expected      attr.Value
		restoresState bool
	}{
		{"unknown setting", "anonymize_ips", tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue), types.BoolUnknown(), false},
		{"project replacement", "project_id", tftypes.NewValue(tftypes.String, "456"), types.StringValue("456"), false},
		{"unchanged settings", "anonymize_ips", tftypes.NewValue(tftypes.Bool, false), nil, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			prior := make(map[string]tftypes.Value)
			for name, attributeType := range objectType.AttributeTypes {
				prior[name] = tftypes.NewValue(attributeType, nil)
			}
			prior["id"] = tftypes.NewValue(tftypes.String, "123")
			prior["project_id"] = tftypes.NewValue(tftypes.String, "123")
			prior["anonymize_ips"] = tftypes.NewValue(tftypes.Bool, false)
			if test.restoresState {
				prior["session_recording_network_payload_capture_config"] = tftypes.NewValue(
					objectType.AttributeTypes["session_recording_network_payload_capture_config"],
					map[string]tftypes.Value{
						"record_headers": tftypes.NewValue(tftypes.Bool, true),
						"record_body":    tftypes.NewValue(tftypes.Bool, false),
					},
				)
			}
			configured := maps.Clone(prior)
			configured["id"] = tftypes.NewValue(tftypes.String, nil)
			configured["session_recording_network_payload_capture_config"] = tftypes.NewValue(objectType.AttributeTypes["session_recording_network_payload_capture_config"], nil)
			configured[test.attribute] = test.value
			planned := maps.Clone(prior)
			for name, attributeType := range objectType.AttributeTypes {
				if name != "id" && name != "project_id" && configured[name].IsNull() {
					planned[name] = tftypes.NewValue(attributeType, tftypes.UnknownValue)
				}
			}
			planned[test.attribute] = test.value
			req := tfresource.ModifyPlanRequest{
				Config: tfsdk.Config{Schema: s, Raw: tftypes.NewValue(objectType, configured)},
				State:  tfsdk.State{Schema: s, Raw: tftypes.NewValue(objectType, prior)},
				Plan:   tfsdk.Plan{Schema: s, Raw: tftypes.NewValue(objectType, planned)},
			}
			resp := tfresource.ModifyPlanResponse{Plan: req.Plan}

			ops.ModifyResourcePlan(ctx, req, &resp)

			require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
			if test.restoresState {
				assert.True(t, resp.Plan.Raw.Equal(req.State.Raw), "unchanged configuration must restore the full prior state")
				return
			}
			var result types.Object
			require.False(t, resp.Plan.Get(ctx, &result).HasError())
			assert.Equal(t, test.expected, result.Attributes()[test.attribute])
		})
	}
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
	diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: resp}, &model)

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
	diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: resp}, &model)

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
	diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: resp}, &model)

	assert.False(t, diags.HasError())
	assert.True(t, model.HeatmapsOptIn.IsNull())
	assert.True(t, model.AutocaptureExceptionsOptIn.IsNull())
	assert.True(t, model.SessionRecordingOptIn.IsNull())
	assert.True(t, model.SurveysOptIn.IsNull())
	assert.True(t, model.CookielessServerHashMode.IsNull())
	assert.True(t, model.AutocaptureWebVitalsOptIn.IsNull())
	assert.True(t, model.AnonymizeIps.IsNull())
	assert.True(t, model.CapturePerformanceOptIn.IsNull())
	assert.True(t, model.NetworkPayloadCapture.IsNull())
	assert.True(t, model.TestAccountFilters.IsNull())
	assert.True(t, model.TestAccountFiltersDefaultChecked.IsNull())
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
	diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: resp}, &model)

	assert.False(t, diags.HasError())
	assert.True(t, model.CapturePerformanceOptIn.ValueBool())
	require.False(t, model.NetworkPayloadCapture.IsNull())
	var cfg NetworkPayloadCaptureModel
	require.False(t, model.NetworkPayloadCapture.As(context.Background(), &cfg, basetypes.ObjectAsOptions{}).HasError())
	assert.True(t, cfg.RecordHeaders.ValueBool())
	require.False(t, cfg.RecordBody.IsNull(), "explicit false must map to a known value, not null")
	assert.False(t, cfg.RecordBody.ValueBool())
}

// TestProjectSettingsMapResponseToModel_DivergenceWarning verifies the
// plan-gated warning names the new attributes when PostHog returns different
// values than the ones the user configured.
func TestProjectSettingsMapResponseToModel_DivergenceWarning(t *testing.T) {
	ops := ProjectSettingsOps{}

	t.Run("network payload capture ignored (server null)", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)

		diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: httpclient.Environment{ID: 123}}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "session_recording_network_payload_capture_config")
	})

	t.Run("network payload capture child value diverged", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, true)

		diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: httpclient.Environment{
			ID: 123,
			SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
				RecordHeaders: util.BoolPtr(true),
				RecordBody:    util.BoolPtr(false),
			},
		}}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "session_recording_network_payload_capture_config")
	})

	t.Run("capture_performance_opt_in diverged", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.CapturePerformanceOptIn = types.BoolValue(true)

		diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: httpclient.Environment{
			ID:                      123,
			CapturePerformanceOptIn: util.BoolPtr(false),
		}}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "capture_performance_opt_in")
	})

	t.Run("anonymize_ips diverged", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.AnonymizeIps = types.BoolValue(true)

		diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: httpclient.Environment{
			ID:           123,
			AnonymizeIps: util.BoolPtr(false),
		}}, &model)

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

		diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: httpclient.Environment{
			ID: 123,
			SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
				RecordHeaders: util.BoolPtr(true),
				RecordBody:    util.BoolPtr(true),
			},
		}}, &model)

		require.Equal(t, 1, diags.WarningsCount())
		assert.Contains(t, diags.Warnings()[0].Detail(), "session_recording_network_payload_capture_config")
	})

	t.Run("no warning when server matches", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.CapturePerformanceOptIn = types.BoolValue(true)
		model.NetworkPayloadCapture = networkPayloadCaptureObject(t, true, false)

		diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: httpclient.Environment{
			ID:                      123,
			CapturePerformanceOptIn: util.BoolPtr(true),
			SessionRecordingNetworkPayloadCaptureConfig: &httpclient.NetworkPayloadCaptureConfig{
				RecordHeaders: util.BoolPtr(true),
				RecordBody:    util.BoolPtr(false),
			},
		}}, &model)

		assert.Equal(t, 0, diags.WarningsCount())
	})
}

func TestProjectSettingsBuildCreateRequest_TestAccountFilters(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.TestAccountFilters = core.NewNormalizedJSONValue(cohortFilterJSON)
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
	model.TestAccountFilters = core.NewNormalizedJSONValue(`[]`)

	req, diags := ops.BuildCreateRequest(context.Background(), model)

	assert.False(t, diags.HasError())
	require.NotNil(t, req.TestAccountFilters, "explicit empty array must be a non-nil pointer to clear the value")
	assert.Empty(t, *req.TestAccountFilters)
}

func TestProjectSettingsBuildCreateRequest_TestAccountFiltersInvalidJSON(t *testing.T) {
	ops := ProjectSettingsOps{}
	model := newProjectSettingsModel()
	model.TestAccountFilters = core.NewNormalizedJSONValue(`{not valid`)

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
	model.TestAccountFilters = core.NewNormalizedJSONValue(cohortFilterJSON)
	diags := ops.MapResponseToModel(context.Background(), projectSettingsResponse{Environment: resp}, &model)

	assert.False(t, diags.HasError())
	require.False(t, model.TestAccountFilters.IsNull())

	// State must match the user config semantically (cohort_name stripped): no drift.
	stateEqualsConfig, d := model.TestAccountFilters.StringSemanticEquals(
		context.Background(),
		core.NewNormalizedJSONValue(cohortFilterJSON),
	)
	assert.False(t, d.HasError())
	assert.True(t, stateEqualsConfig, "cohort_name must be stripped so the filter round-trips without drift")
	assert.NotContains(t, model.TestAccountFilters.ValueString(), "cohort_name")

	assert.True(t, model.TestAccountFiltersDefaultChecked.ValueBool())
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

// fakeProjectSettingsAPI serves the environment and events retention endpoints
// for one project and records the calls made to them.
type fakeProjectSettingsAPI struct {
	t *testing.T
	// retentionMonths is the project's events retention period on the server.
	retentionMonths *int64
	// getStatus and postStatus, when set, are returned by the retention read and set instead of a result.
	getStatus  int
	postStatus int
	calls      []string
	posted     []int64
}

func (f *fakeProjectSettingsAPI) client() (httpclient.PosthogClient, func()) {
	envPath := "/api/environments/" + testProjectSettingsProjectID + "/"
	retentionPath := "/api/projects/" + testProjectSettingsProjectID + "/events_retention/"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.calls = append(f.calls, r.Method+" "+r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == envPath && (r.Method == http.MethodPatch || r.Method == http.MethodGet):
			assert.NoError(f.t, json.NewEncoder(w).Encode(httpclient.Environment{ID: 123, HeatmapsOptIn: util.BoolPtr(true)}))
		case r.URL.Path == retentionPath && r.Method == http.MethodGet && f.getStatus != 0:
			w.WriteHeader(f.getStatus)
			_, _ = w.Write([]byte(`{"detail":"retention unavailable"}`))
		case r.URL.Path == retentionPath && r.Method == http.MethodPost && f.postStatus != 0:
			w.WriteHeader(f.postStatus)
			_, _ = w.Write([]byte(`{"detail":"Your organization access level is insufficient."}`))
		case r.URL.Path == retentionPath && r.Method == http.MethodPost:
			var body httpclient.EventsRetentionRequest
			assert.NoError(f.t, json.NewDecoder(r.Body).Decode(&body))
			f.posted = append(f.posted, body.RetentionMonths)
			f.retentionMonths = &body.RetentionMonths
			assert.NoError(f.t, json.NewEncoder(w).Encode(httpclient.EventsRetention{ProjectRetentionMonths: f.retentionMonths}))
		case r.URL.Path == retentionPath && r.Method == http.MethodGet:
			assert.NoError(f.t, json.NewEncoder(w).Encode(httpclient.EventsRetention{ProjectRetentionMonths: f.retentionMonths}))
		default:
			f.t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusTeapot)
		}
	}))
	return httpclient.NewClient(server.Client(), server.URL, "test-key", "test"), server.Close
}

func TestProjectSettingsCRUDFlow(t *testing.T) {
	patchEnv := "PATCH /api/environments/" + testProjectSettingsProjectID + "/"
	getRetention := "GET /api/projects/" + testProjectSettingsProjectID + "/events_retention/"
	postRetention := "POST /api/projects/" + testProjectSettingsProjectID + "/events_retention/"
	// setTo12 sets the period to 12 months, from a plan made when the project
	// used the organization's period (unknown on create).
	setTo12 := func(plannedFrom types.Int64) projectSettingsRequest {
		return projectSettingsRequest{EventsRetentionMonths: new(int64(12)), PlannedFromEventsRetentionMonths: plannedFrom}
	}

	t.Run("create without a period reads it and does not set it", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, retentionMonths: new(int64(24))}
		client, closeServer := api.client()
		defer closeServer()

		created, err := ProjectSettingsOps{}.Create(context.Background(), client, newProjectSettingsModel(), projectSettingsRequest{})
		require.NoError(t, err)
		assert.Equal(t, []string{getRetention, patchEnv}, api.calls)
		assert.Equal(t, new(int64(24)), created.EventsRetentionMonths)
		require.NotNil(t, created.HeatmapsOptIn)
	})

	t.Run("a different period is set after the environment", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t}
		client, closeServer := api.client()
		defer closeServer()

		updated, _, err := ProjectSettingsOps{}.Update(context.Background(), client, newProjectSettingsModel(), setTo12(types.Int64Null()))
		require.NoError(t, err)
		assert.Equal(t, []string{getRetention, patchEnv, postRetention}, api.calls)
		assert.Equal(t, []int64{12}, api.posted)
		assert.Equal(t, new(int64(12)), updated.EventsRetentionMonths)
	})

	t.Run("a period PostHog already has is not sent again", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, retentionMonths: new(int64(12))}
		client, closeServer := api.client()
		defer closeServer()

		created, err := ProjectSettingsOps{}.Create(context.Background(), client, newProjectSettingsModel(), setTo12(types.Int64Unknown()))
		require.NoError(t, err)
		assert.Empty(t, api.posted)
		assert.Equal(t, new(int64(12)), created.EventsRetentionMonths)
	})

	for _, test := range []struct {
		name        string
		plannedFrom types.Int64
		want        string
	}{
		{"a plan that keeps the period", types.Int64Value(12), "events_retention_months is 84 months in PostHog, but this plan was made when it was 12 months"},
		{"a plan that changes the period", types.Int64Value(24), "events_retention_months is 84 months in PostHog, but this plan was made when it was 24 months"},
		{"a plan made against the organization's period", types.Int64Null(), "but this plan was made when it was the organization's period"},
	} {
		t.Run(test.name+" stops before any write when the period changed in PostHog after the plan", func(t *testing.T) {
			api := &fakeProjectSettingsAPI{t: t, retentionMonths: new(int64(84))}
			client, closeServer := api.client()
			defer closeServer()

			_, _, err := ProjectSettingsOps{}.Update(context.Background(), client, newProjectSettingsModel(), setTo12(test.plannedFrom))
			require.ErrorContains(t, err, test.want)
			assert.Equal(t, []string{getRetention}, api.calls)
		})
	}

	t.Run("a period changed in PostHog to the one the plan sets is not an error", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, retentionMonths: new(int64(12))}
		client, closeServer := api.client()
		defer closeServer()

		_, _, err := ProjectSettingsOps{}.Update(context.Background(), client, newProjectSettingsModel(), setTo12(types.Int64Value(24)))
		require.NoError(t, err)
		assert.Equal(t, []string{getRetention, patchEnv}, api.calls)
	})

	t.Run("a create applies to whatever period PostHog has", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, retentionMonths: new(int64(84))}
		client, closeServer := api.client()
		defer closeServer()

		_, err := ProjectSettingsOps{}.Create(context.Background(), client, newProjectSettingsModel(), setTo12(types.Int64Unknown()))
		require.NoError(t, err)
		assert.Equal(t, []int64{12}, api.posted)
	})

	t.Run("read without a retention endpoint reads no period", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, getStatus: http.StatusNotFound}
		client, closeServer := api.client()
		defer closeServer()

		read, status, err := ProjectSettingsOps{}.Read(context.Background(), client, newProjectSettingsModel())
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, int(status))
		assert.Nil(t, read.EventsRetentionMonths)
		require.NotNil(t, read.HeatmapsOptIn)
	})

	t.Run("a configured period fails without a retention endpoint", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, getStatus: http.StatusNotFound, postStatus: http.StatusNotFound}
		client, closeServer := api.client()
		defer closeServer()

		_, err := ProjectSettingsOps{}.Create(context.Background(), client, newProjectSettingsModel(), setTo12(types.Int64Unknown()))
		require.ErrorContains(t, err, "PostHog returned 404 for /api/projects/"+testProjectSettingsProjectID+"/events_retention/, which events_retention_months needs")
		assert.Equal(t, []string{getRetention}, api.calls, "a configured period must fail before any write")
	})

	t.Run("a refused set fails after the environment applied", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, postStatus: http.StatusForbidden}
		client, closeServer := api.client()
		defer closeServer()

		_, _, err := ProjectSettingsOps{}.Update(context.Background(), client, newProjectSettingsModel(), setTo12(types.Int64Null()))
		require.ErrorContains(t, err, "the other settings were saved, but setting events_retention_months to 12 failed")
		require.ErrorContains(t, err, "Your organization access level is insufficient.")
		assert.Equal(t, []string{getRetention, patchEnv, postRetention}, api.calls)
	})

	t.Run("read fails when the retention read fails", func(t *testing.T) {
		api := &fakeProjectSettingsAPI{t: t, getStatus: http.StatusInternalServerError}
		client, closeServer := api.client()
		defer closeServer()

		_, _, err := ProjectSettingsOps{}.Read(context.Background(), client, newProjectSettingsModel())
		require.ErrorContains(t, err, "failed to read events retention")
	})
}

func TestProjectSettingsEventsRetentionMonths(t *testing.T) {
	ctx := context.Background()
	ops := ProjectSettingsOps{}

	t.Run("a configured period is requested", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.EventsRetentionMonths = types.Int64Value(12)
		req, diags := ops.BuildCreateRequest(ctx, model)
		require.False(t, diags.HasError())
		assert.Equal(t, new(int64(12)), req.EventsRetentionMonths)
	})

	t.Run("an omitted period is not requested", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.EventsRetentionMonths = types.Int64Unknown()
		req, diags := ops.BuildUpdateRequest(ctx, model, newProjectSettingsModel())
		require.False(t, diags.HasError())
		assert.Nil(t, req.EventsRetentionMonths)
	})

	t.Run("an update records the period its plan was made against", func(t *testing.T) {
		plan, state := newProjectSettingsModel(), newProjectSettingsModel()
		plan.EventsRetentionMonths, state.EventsRetentionMonths = types.Int64Value(12), types.Int64Value(24)
		req, diags := ops.BuildUpdateRequest(ctx, plan, state)
		require.False(t, diags.HasError())
		assert.Equal(t, new(int64(12)), req.EventsRetentionMonths)
		assert.Equal(t, types.Int64Value(24), req.PlannedFromEventsRetentionMonths)
	})

	t.Run("a create has no period its plan was made against", func(t *testing.T) {
		model := newProjectSettingsModel()
		model.EventsRetentionMonths = types.Int64Value(12)
		req, diags := ops.BuildCreateRequest(ctx, model)
		require.False(t, diags.HasError())
		assert.True(t, req.PlannedFromEventsRetentionMonths.IsUnknown())
	})

	t.Run("the period maps to state", func(t *testing.T) {
		model := newProjectSettingsModel()
		diags := ops.MapResponseToModel(ctx, projectSettingsResponse{EventsRetentionMonths: new(int64(24))}, &model)
		require.False(t, diags.HasError())
		assert.Equal(t, types.Int64Value(24), model.EventsRetentionMonths)
	})

	t.Run("the organization's period maps to null", func(t *testing.T) {
		model := newProjectSettingsModel()
		diags := ops.MapResponseToModel(ctx, projectSettingsResponse{}, &model)
		require.False(t, diags.HasError())
		assert.True(t, model.EventsRetentionMonths.IsNull())
	})
}

func TestProjectSettingsEventsRetentionMonthsValidator(t *testing.T) {
	ctx := context.Background()
	attribute := ProjectSettingsOps{}.Schema().Attributes["events_retention_months"].(schema.Int64Attribute)
	for _, test := range []struct {
		months int64
		valid  bool
	}{
		// PostHog decides which periods it accepts, so any positive value plans.
		{1, true}, {6, true}, {84, true}, {120, true},
		{0, false}, {-1, false},
	} {
		t.Run(strconv.FormatInt(test.months, 10), func(t *testing.T) {
			resp := &validator.Int64Response{}
			for _, v := range attribute.Validators {
				v.ValidateInt64(ctx, validator.Int64Request{Path: path.Root("events_retention_months"), ConfigValue: types.Int64Value(test.months)}, resp)
			}
			assert.Equal(t, !test.valid, resp.Diagnostics.HasError())
		})
	}
}

func TestProjectSettingsModifyPlanRetentionWarning(t *testing.T) {
	ctx := context.Background()
	ops := ProjectSettingsOps{}
	s := ops.Schema()
	objectType := s.Type().TerraformType(ctx).(tftypes.Object)
	settingsFor := func(projectID string, retention any) tftypes.Value {
		values := make(map[string]tftypes.Value)
		for name, attributeType := range objectType.AttributeTypes {
			values[name] = tftypes.NewValue(attributeType, nil)
		}
		values["id"] = tftypes.NewValue(tftypes.String, projectID)
		values["project_id"] = tftypes.NewValue(tftypes.String, projectID)
		values["events_retention_months"] = tftypes.NewValue(tftypes.Number, retention)
		return tftypes.NewValue(objectType, values)
	}
	settings := func(retention any) tftypes.Value { return settingsFor("123", retention) }
	for _, test := range []struct {
		name  string
		state tftypes.Value

		planned any
		want    string
	}{
		{"first apply", tftypes.NewValue(objectType, nil), int64(12), "replacing its current period in PostHog"},
		{"update", settings(int64(24)), int64(12), "from 24 to 12 months"},
		{"period only known at apply", settings(int64(24)), tftypes.UnknownValue, "a period only known at apply"},
		{"replacement by another project", settingsFor("456", int64(24)), int64(12), "replacing its current period in PostHog"},
	} {
		t.Run(test.name, func(t *testing.T) {
			req := tfresource.ModifyPlanRequest{
				Config: tfsdk.Config{Schema: s, Raw: settings(test.planned)},
				State:  tfsdk.State{Schema: s, Raw: test.state},
				Plan:   tfsdk.Plan{Schema: s, Raw: settings(test.planned)},
			}
			resp := tfresource.ModifyPlanResponse{Plan: req.Plan}

			ops.ModifyResourcePlan(ctx, req, &resp)

			require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
			require.Equal(t, 1, resp.Diagnostics.WarningsCount())
			assert.Contains(t, resp.Diagnostics.Warnings()[0].Detail(), test.want)
		})
	}
}

func TestShorterRetentionWarning(t *testing.T) {
	for _, test := range []struct {
		name       string
		configured types.Int64
		have       types.Int64
		want       string
	}{
		{"shorter period", types.Int64Value(12), types.Int64Value(24), "from 24 to 12 months"},
		{"replaces the organization's period", types.Int64Value(12), types.Int64Null(), "replacing the organization's period"},
		{"first apply, current period unknown", types.Int64Value(12), types.Int64Unknown(), "replacing its current period in PostHog"},
		{"longer period", types.Int64Value(24), types.Int64Value(12), ""},
		{"unchanged period", types.Int64Value(12), types.Int64Value(12), ""},
		{"omitted period", types.Int64Null(), types.Int64Value(12), ""},
		{"period only known at apply", types.Int64Unknown(), types.Int64Value(12), "a period only known at apply"},
		{"omitted on first apply", types.Int64Null(), types.Int64Unknown(), ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			diags := shorterRetentionWarning(test.configured, test.have)
			if test.want == "" {
				assert.Zero(t, diags.WarningsCount())
				return
			}
			require.Equal(t, 1, diags.WarningsCount())
			assert.Contains(t, diags.Warnings()[0].Detail(), test.want)
		})
	}
}
