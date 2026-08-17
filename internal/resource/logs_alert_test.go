package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// emptyStringSet builds an explicitly empty (non-null) set. Note that calling the
// variadic stringSet with no arguments yields a *null* set instead, since a
// zero-argument variadic call passes a nil slice.
func emptyStringSet(t *testing.T) types.Set {
	t.Helper()
	set, diags := types.SetValueFrom(context.Background(), types.StringType, []string{})
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	return set
}

func blockedWindowSet(t *testing.T, windows ...BlockedWindowTFModel) types.Set {
	t.Helper()
	objectType := types.ObjectType{AttrTypes: map[string]attr.Type{
		"start": types.StringType,
		"end":   types.StringType,
	}}
	set, diags := types.SetValueFrom(context.Background(), objectType, windows)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	return set
}

// PostHog echoes back only the filter keys that are set, so the filters the user
// never configured must land as null rather than as empty collections.
func TestLogsAlertMapResponseToModel_AbsentFilterKeysAreNull(t *testing.T) {
	model := LogsAlertTFModel{}

	resp := httpclient.LogsAlert{
		ID:      "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Name:    util.StringPtr("Contour Errors"),
		Enabled: util.BoolPtr(true),
		State:   util.StringPtr("not_firing"),
		Filters: &httpclient.LogsAlertFilters{
			ServiceNames: []string{"contour"},
		},
		ThresholdCount:    util.Int64Ptr(10),
		ThresholdOperator: util.StringPtr("above"),
		WindowMinutes:     util.Int64Ptr(10),
	}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.Equal(t, types.StringValue("019dbe94-cec8-781b-9470-4a970cd69ebf"), model.ID)
	assert.Equal(t, stringSet(t, "contour"), model.ServiceNames)
	assert.True(t, model.SeverityLevels.IsNull(), "severity_levels was never configured, so it must be null")
	assert.True(t, model.FilterGroupJSON.IsNull(), "filter_group_json was never configured, so it must be null")
	assert.True(t, model.BlockedWindows.IsNull(), "no schedule_restriction means quiet hours are off")
	assert.Equal(t, types.Int64Value(10), model.ThresholdCount)
	assert.Equal(t, types.StringValue("not_firing"), model.State)
}

// A user who writes `severity_levels = []` must keep an empty set in state; flipping
// it to null would show as drift on every plan.
func TestLogsAlertMapResponseToModel_PreservesConfiguredEmptySet(t *testing.T) {
	model := LogsAlertTFModel{
		SeverityLevels: emptyStringSet(t),
	}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID:      "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Filters: &httpclient.LogsAlertFilters{ServiceNames: []string{"api"}},
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.False(t, model.SeverityLevels.IsNull())
	assert.Empty(t, model.SeverityLevels.Elements())
}

// filter_group_json must survive a create -> read round trip without a semantic diff,
// even though PostHog is free to reorder keys.
func TestLogsAlertFilterGroupJSONRoundTrips(t *testing.T) {
	filterGroup := `{"type":"AND","values":[{"type":"AND","values":[{"type":"log_attribute","key":"status_code","operator":"exact","value":["500"]}]}]}`

	model := LogsAlertTFModel{
		FilterGroupJSON: jsontypes.NewNormalizedValue(filterGroup),
	}

	req, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	require.NotNil(t, req.Filters)
	require.NotNil(t, req.Filters.FilterGroup)

	mapped := LogsAlertTFModel{}
	diags = LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID:      "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Filters: &httpclient.LogsAlertFilters{FilterGroup: req.Filters.FilterGroup},
	}, &mapped)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.JSONEq(t, filterGroup, mapped.FilterGroupJSON.ValueString())
}

func TestLogsAlertBuildCreateRequest_RejectsInvalidFilterGroupJSON(t *testing.T) {
	model := LogsAlertTFModel{
		FilterGroupJSON: jsontypes.NewNormalizedValue(`{"type":`),
	}

	_, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), model)
	assert.True(t, diags.HasError(), "expected a diagnostic for unparseable filter_group_json")
}

// filters is a whole-object replace, so it must always be sent — otherwise a draft
// alert can never have its filters cleared.
func TestLogsAlertBuildCreateRequest_AlwaysSendsFilters(t *testing.T) {
	req, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), LogsAlertTFModel{})
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	require.NotNil(t, req.Filters)
	assert.Nil(t, req.Filters.SeverityLevels)
	assert.Nil(t, req.Filters.ServiceNames)
	assert.Nil(t, req.Filters.FilterGroup)
	assert.Nil(t, req.ScheduleRestriction, "no blocked_windows must send null to clear quiet hours")
}

func TestLogsAlertBuildCreateRequest_MapsBlockedWindows(t *testing.T) {
	model := LogsAlertTFModel{
		SeverityLevels: stringSet(t, "error"),
		BlockedWindows: blockedWindowSet(t, BlockedWindowTFModel{
			Start: types.StringValue("22:00"),
			End:   types.StringValue("23:30"),
		}),
	}

	req, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	require.NotNil(t, req.ScheduleRestriction)
	require.Len(t, req.ScheduleRestriction.BlockedWindows, 1)
	assert.Equal(t, "22:00", req.ScheduleRestriction.BlockedWindows[0].Start)
	assert.Equal(t, "23:30", req.ScheduleRestriction.BlockedWindows[0].End)
	assert.Equal(t, []string{"error"}, req.Filters.SeverityLevels)
}

func TestLogsAlertMapResponseToModel_MapsBlockedWindows(t *testing.T) {
	model := LogsAlertTFModel{}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID: "019dbe94-cec8-781b-9470-4a970cd69ebf",
		ScheduleRestriction: &httpclient.LogsAlertSchedule{
			BlockedWindows: []httpclient.LogsAlertBlockedWindow{{Start: "22:00", End: "23:30"}},
		},
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	expected := blockedWindowSet(t, BlockedWindowTFModel{
		Start: types.StringValue("22:00"),
		End:   types.StringValue("23:30"),
	})
	assert.Equal(t, expected, model.BlockedWindows)
}
