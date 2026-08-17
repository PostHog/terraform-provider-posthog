package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func blockedWindowsSet(t *testing.T, windows ...[2]string) types.Set {
	t.Helper()

	models := make([]AlertBlockedWindowModel, len(windows))
	for i, window := range windows {
		models[i] = AlertBlockedWindowModel{
			Start: types.StringValue(window[0]),
			End:   types.StringValue(window[1]),
		}
	}

	set, diags := types.SetValueFrom(context.Background(), alertBlockedWindowObjectType, models)
	require.False(t, diags.HasError(), "%v", diags)
	return set
}

func alertModelWithWindows(t *testing.T, windows ...[2]string) AlertResourceTFModel {
	t.Helper()

	restriction, diags := types.ObjectValueFrom(context.Background(), alertScheduleRestrictionAttrTypes, AlertScheduleRestrictionModel{
		BlockedWindows: blockedWindowsSet(t, windows...),
	})
	require.False(t, diags.HasError(), "%v", diags)

	return AlertResourceTFModel{
		Insight:             types.Int64Value(1),
		ThresholdType:       types.StringValue("absolute"),
		ConditionType:       types.StringValue("absolute_value"),
		SeriesIndex:         types.Int64Value(0),
		SubscribedUsers:     types.SetNull(types.Int64Type),
		ScheduleRestriction: restriction,
	}
}

func TestBuildCreateRequestScheduleRestriction(t *testing.T) {
	ctx := context.Background()

	t.Run("windows are sent", func(t *testing.T) {
		req, diags := AlertOps{}.BuildCreateRequest(ctx, alertModelWithWindows(t, [2]string{"22:00", "23:59"}, [2]string{"00:00", "06:00"}))
		require.False(t, diags.HasError(), "%v", diags)
		require.NotNil(t, req.ScheduleRestriction)
		assert.ElementsMatch(t, []httpclient.AlertBlockedWindow{
			{Start: "22:00", End: "23:59"},
			{Start: "00:00", End: "06:00"},
		}, req.ScheduleRestriction.BlockedWindows)
	})

	// A null restriction has to reach the API as an explicit null, otherwise removing the
	// block from config would leave quiet hours in place on the PATCH.
	t.Run("absent restriction serializes as null", func(t *testing.T) {
		model := alertModelWithWindows(t)
		model.ScheduleRestriction = types.ObjectNull(alertScheduleRestrictionAttrTypes)

		req, diags := AlertOps{}.BuildCreateRequest(ctx, model)
		require.False(t, diags.HasError(), "%v", diags)
		assert.Nil(t, req.ScheduleRestriction)

		body, err := json.Marshal(req)
		require.NoError(t, err)
		assert.Contains(t, string(body), `"schedule_restriction":null`)
	})
}

func TestMapResponseToModelScheduleRestriction(t *testing.T) {
	ctx := context.Background()

	t.Run("populated", func(t *testing.T) {
		model := alertModelWithWindows(t)
		resp := httpclient.Alert{
			ID:      "01a000df-f6cc-0000-9779-2ebb16417c3e",
			Insight: httpclient.AlertInsight{ID: 1},
			ScheduleRestriction: &httpclient.AlertScheduleRestriction{
				BlockedWindows: []httpclient.AlertBlockedWindow{{Start: "00:00", End: "06:00"}},
			},
		}

		diags := AlertOps{}.MapResponseToModel(ctx, resp, &model)
		require.False(t, diags.HasError(), "%v", diags)
		require.False(t, model.ScheduleRestriction.IsNull())

		var restriction AlertScheduleRestrictionModel
		require.False(t, model.ScheduleRestriction.As(ctx, &restriction, basetypes.ObjectAsOptions{}).HasError())

		var windows []AlertBlockedWindowModel
		require.False(t, restriction.BlockedWindows.ElementsAs(ctx, &windows, false).HasError())
		assert.Equal(t, []AlertBlockedWindowModel{
			{Start: types.StringValue("00:00"), End: types.StringValue("06:00")},
		}, windows)
	})

	t.Run("absent", func(t *testing.T) {
		model := alertModelWithWindows(t, [2]string{"00:00", "06:00"})
		resp := httpclient.Alert{
			ID:      "01a000df-f6cc-0000-9779-2ebb16417c3e",
			Insight: httpclient.AlertInsight{ID: 1},
		}

		diags := AlertOps{}.MapResponseToModel(ctx, resp, &model)
		require.False(t, diags.HasError(), "%v", diags)
		assert.True(t, model.ScheduleRestriction.IsNull())
	})
}

func TestBlockedWindowsValidator(t *testing.T) {
	const (
		tooShort  = "Blocked window is too short"
		overlaps  = "Overlapping blocked windows"
		crossesMN = "Blocked window crossing midnight must be the only window"
		meetsAtMN = "Blocked windows meeting at midnight are stored as one"
	)

	// wantSummary is the diagnostic summary expected, empty meaning the config is valid.
	// Asserting the summary rather than a bool stops a case passing because a different
	// rule happened to reject it.
	tests := map[string]struct {
		windows     [][2]string
		wantSummary string
	}{
		"separate windows":                           {windows: [][2]string{{"00:00", "06:00"}, {"22:00", "23:59"}}},
		"windows with a gap between them":            {windows: [][2]string{{"01:00", "05:00"}, {"12:00", "13:00"}}},
		"lone wrapped window":                        {windows: [][2]string{{"22:00", "07:00"}}},
		"window ending at midnight plus daytime":     {windows: [][2]string{{"19:00", "00:00"}, {"12:00", "13:00"}}},
		"malformed times are ignored":                {windows: [][2]string{{"nonsense", "06:00"}, {"05:00", "09:00"}}},
		"equal bounds are left to the API":           {windows: [][2]string{{"02:00", "02:00"}}},
		"exactly thirty minutes":                     {windows: [][2]string{{"02:00", "02:30"}}},
		"wrapped window is measured across midnight": {windows: [][2]string{{"23:50", "00:30"}}},
		"shorter than thirty minutes":                {windows: [][2]string{{"02:00", "02:15"}}, wantSummary: tooShort},
		"wrapped window shorter than thirty minutes": {windows: [][2]string{{"23:50", "00:10"}}, wantSummary: tooShort},
		"touching windows are merged":                {windows: [][2]string{{"00:00", "06:00"}, {"06:00", "09:00"}}, wantSummary: overlaps},
		"overlapping windows":                        {windows: [][2]string{{"00:00", "06:00"}, {"05:00", "09:00"}}, wantSummary: overlaps},
		"contained window":                           {windows: [][2]string{{"00:00", "09:00"}, {"02:00", "03:00"}}, wantSummary: overlaps},
		"wrapped window overlaps morning":            {windows: [][2]string{{"22:00", "07:00"}, {"06:00", "08:00"}}, wantSummary: overlaps},
		"wrapped window plus daytime":                {windows: [][2]string{{"22:00", "07:00"}, {"12:00", "13:00"}}, wantSummary: crossesMN},
		"midnight blocked from both sides":           {windows: [][2]string{{"00:00", "06:00"}, {"22:00", "00:00"}}, wantSummary: meetsAtMN},
		// PostHog only rejoins a midnight pair while it is the whole timeline. A third
		// window anywhere in the day leaves all three stored exactly as written.
		"midnight pair with a third window": {windows: [][2]string{{"22:00", "00:00"}, {"00:00", "07:00"}, {"12:00", "13:00"}}},
		"midnight pair with two others":     {windows: [][2]string{{"00:00", "06:00"}, {"08:00", "09:00"}, {"12:00", "13:00"}, {"19:00", "00:00"}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := validator.SetRequest{
				Path:        path.Root("schedule_restriction").AtName("blocked_windows"),
				ConfigValue: blockedWindowsSet(t, test.windows...),
			}
			resp := &validator.SetResponse{}

			blockedWindowsValidator{}.ValidateSet(context.Background(), req, resp)

			if test.wantSummary == "" {
				assert.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
				return
			}
			require.True(t, resp.Diagnostics.HasError(), "expected a diagnostic")
			assert.Equal(t, test.wantSummary, resp.Diagnostics.Errors()[0].Summary())
		})
	}
}
