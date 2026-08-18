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

	// An empty window list means the same thing as no restriction. This is the branch the
	// httpclient comment promises, and it is what stops a cleared restriction reading back
	// as a populated object against a null config.
	t.Run("empty list is treated as absent", func(t *testing.T) {
		model := alertModelWithWindows(t, [2]string{"00:00", "06:00"})
		resp := httpclient.Alert{
			ID:                  "01a000df-f6cc-0000-9779-2ebb16417c3e",
			Insight:             httpclient.AlertInsight{ID: 1},
			ScheduleRestriction: &httpclient.AlertScheduleRestriction{BlockedWindows: []httpclient.AlertBlockedWindow{}},
		}

		diags := AlertOps{}.MapResponseToModel(ctx, resp, &model)
		require.False(t, diags.HasError(), "%v", diags)
		assert.True(t, model.ScheduleRestriction.IsNull())
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
		overlaps  = "Overlapping blocked windows"
		crossesMN = "Blocked window crossing midnight must be the only window"
		meetsAtMN = "Blocked windows meeting at midnight are stored as one"
	)

	// wantSummaries lists every diagnostic summary expected, in order; empty means the
	// config is valid. Asserting the whole list rather than the first entry is what pins
	// the validator reporting all problems in one plan instead of stopping at the first.
	tests := map[string]struct {
		windows       [][2]string
		wantSummaries []string
	}{
		"separate windows":                       {windows: [][2]string{{"00:00", "06:00"}, {"22:00", "23:59"}}},
		"windows with a gap between them":        {windows: [][2]string{{"01:00", "05:00"}, {"12:00", "13:00"}}},
		"lone wrapped window":                    {windows: [][2]string{{"22:00", "07:00"}}},
		"window ending at midnight plus daytime": {windows: [][2]string{{"19:00", "00:00"}, {"12:00", "13:00"}}},
		"malformed times are ignored":            {windows: [][2]string{{"nonsense", "06:00"}, {"05:00", "09:00"}}},
		// The HH:MM regex and this validator run independently, so anything the regex
		// rejects can still reach here. A negative hour used to parse and index the
		// coverage array out of range.
		"negative hour is rejected, not parsed": {windows: [][2]string{{"-1:30", "06:00"}}},
		"out-of-range time is rejected":         {windows: [][2]string{{"99:99", "06:00"}}},
		"trailing junk is rejected":             {windows: [][2]string{{"12:30extra", "13:00"}}},
		// time.Parse alone would read this as 09:30, which the pattern does not allow.
		// Both layers have to agree on what a valid time is.
		"single-digit hour is rejected": {windows: [][2]string{{"9:30", "13:00"}}},
		// Window length and count are PostHog's to enforce; repeating them here would mean
		// a provider release whenever it changes one. The API rejects these on apply.
		"short window is left to the API":         {windows: [][2]string{{"02:00", "02:15"}}},
		"wrapped short window is left to the API": {windows: [][2]string{{"23:50", "00:10"}}},
		// Equal bounds block no time. Skipped rather than treated as a span, since
		// 00:00-00:00 would otherwise read as a whole-day block.
		"equal bounds are skipped":                   {windows: [][2]string{{"02:00", "02:00"}}},
		"exactly thirty minutes":                     {windows: [][2]string{{"02:00", "02:30"}}},
		"wrapped window is measured across midnight": {windows: [][2]string{{"23:50", "00:30"}}},
		"touching windows are merged":                {windows: [][2]string{{"00:00", "06:00"}, {"06:00", "09:00"}}, wantSummaries: []string{overlaps}},
		"overlapping windows":                        {windows: [][2]string{{"00:00", "06:00"}, {"05:00", "09:00"}}, wantSummaries: []string{overlaps}},
		"contained window":                           {windows: [][2]string{{"00:00", "09:00"}, {"02:00", "03:00"}}, wantSummaries: []string{overlaps}},
		// Both rules fire: the wrapped window overlaps the morning one, and it is not
		// the only window. Asserting both is what pins the validator reporting every
		// problem in one plan.
		"wrapped window overlaps morning":  {windows: [][2]string{{"22:00", "07:00"}, {"06:00", "08:00"}}, wantSummaries: []string{overlaps, crossesMN}},
		"wrapped window plus daytime":      {windows: [][2]string{{"22:00", "07:00"}, {"12:00", "13:00"}}, wantSummaries: []string{crossesMN}},
		"midnight blocked from both sides": {windows: [][2]string{{"00:00", "06:00"}, {"22:00", "00:00"}}, wantSummaries: []string{meetsAtMN}},
		// PostHog only rejoins a midnight pair while it is the whole timeline. A third
		// window anywhere in the day leaves all three stored exactly as written.
		"midnight pair with a third window": {windows: [][2]string{{"22:00", "00:00"}, {"00:00", "07:00"}, {"12:00", "13:00"}}},
		// A whole-day block necessarily touches, so the reshape rules still stop it
		// reaching the API in a form PostHog would silently rewrite.
		"windows covering the whole day": {
			windows:       [][2]string{{"00:00", "12:00"}, {"12:00", "00:00"}},
			wantSummaries: []string{overlaps, meetsAtMN},
		},
		// The meets-at-midnight check tests both orientations of the pair. Terraform does
		// not promise set elements reach ElementsAs in config order, so the reversed form
		// is the one that fires roughly half the time in practice.
		"midnight blocked from both sides, reversed": {windows: [][2]string{{"22:00", "00:00"}, {"00:00", "06:00"}}, wantSummaries: []string{meetsAtMN}},
		// Short windows now take part in the timeline like any other, so they overlap
		// their neighbours normally instead of being skipped and leaving a partial picture.
		"short window overlaps its neighbour": {
			windows:       [][2]string{{"02:00", "02:10"}, {"00:00", "06:00"}, {"22:00", "00:00"}},
			wantSummaries: []string{overlaps},
		},
		"short window does not hide a real overlap": {
			windows:       [][2]string{{"02:00", "02:10"}, {"08:00", "12:00"}, {"11:00", "14:00"}},
			wantSummaries: []string{overlaps},
		},
		"short window beside a crossing one": {
			windows:       [][2]string{{"02:00", "02:10"}, {"22:00", "07:00"}},
			wantSummaries: []string{overlaps, crossesMN},
		},
		"midnight pair with two others": {windows: [][2]string{{"00:00", "06:00"}, {"08:00", "09:00"}, {"12:00", "13:00"}, {"19:00", "00:00"}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := validator.SetRequest{
				Path:        path.Root("schedule_restriction").AtName("blocked_windows"),
				ConfigValue: blockedWindowsSet(t, test.windows...),
			}
			resp := &validator.SetResponse{}

			blockedWindowsValidator{}.ValidateSet(context.Background(), req, resp)

			var got []string
			for _, d := range resp.Diagnostics.Errors() {
				got = append(got, d.Summary())
			}
			assert.Equal(t, test.wantSummaries, got)
		})
	}
}
