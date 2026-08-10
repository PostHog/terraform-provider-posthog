package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func int64Set(t *testing.T, values ...int64) types.Set {
	t.Helper()
	s, diags := types.SetValueFrom(context.Background(), types.Int64Type, values)
	require.False(t, diags.HasError(), diags.Errors())
	return s
}

func stringSet(t *testing.T, values ...string) types.Set {
	t.Helper()
	s, diags := types.SetValueFrom(context.Background(), types.StringType, values)
	require.False(t, diags.HasError(), diags.Errors())
	return s
}

// TestSubscriptionBuildCreateRequest_SlackDashboard covers the headline flow: a Slack
// dashboard digest with the full set of required + optional fields.
func TestSubscriptionBuildCreateRequest_SlackDashboard(t *testing.T) {
	ops := SubscriptionOps{}
	model := SubscriptionResourceTFModel{
		TargetType:              types.StringValue("slack"),
		TargetValue:             types.StringValue("C0B9A53J8RF|#reports"),
		IntegrationID:           types.Int64Value(1),
		DashboardID:             types.Int64Value(7),
		DashboardExportInsights: int64Set(t, 11, 22),
		Frequency:               types.StringValue("daily"),
		Interval:                types.Int64Value(1),
		StartDate:               types.StringValue("2026-08-17T07:00:00Z"),
		Title:                   types.StringValue("Daily key metrics"),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())

	assert.Equal(t, "slack", req.TargetType)
	assert.Equal(t, "C0B9A53J8RF|#reports", req.TargetValue)
	require.NotNil(t, req.IntegrationID)
	assert.Equal(t, int64(1), *req.IntegrationID)
	require.NotNil(t, req.Dashboard)
	assert.Equal(t, int64(7), *req.Dashboard)
	assert.Nil(t, req.Insight)
	assert.ElementsMatch(t, []int64{11, 22}, req.DashboardExportInsights)
	assert.Equal(t, "daily", req.Frequency)
	assert.Equal(t, int64(1), req.Interval)
	assert.Equal(t, "2026-08-17T07:00:00Z", req.StartDate)
	require.NotNil(t, req.Title)
	assert.Equal(t, "Daily key metrics", *req.Title)
}

// TestSubscriptionBuildCreateRequest_SendsIntervalUnconditionally guards that interval is
// always on the wire (the API 400s if omitted, and it has no server default).
func TestSubscriptionBuildCreateRequest_SendsIntervalUnconditionally(t *testing.T) {
	ops := SubscriptionOps{}
	model := SubscriptionResourceTFModel{
		TargetType:  types.StringValue("email"),
		TargetValue: types.StringValue("team@example.com"),
		InsightID:   types.Int64Value(9),
		Frequency:   types.StringValue("weekly"),
		Interval:    types.Int64Value(1),
		StartDate:   types.StringValue("2026-08-17T07:00:00Z"),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())
	assert.Equal(t, int64(1), req.Interval)
	assert.Nil(t, req.IntegrationID)
	assert.Nil(t, req.Dashboard)
	require.NotNil(t, req.Insight)
	assert.Equal(t, int64(9), *req.Insight)
}

// TestSubscriptionBuildCreateRequest_UnsetOptionals verifies how unset optionals reach the
// wire: the pointer optionals stay nil (integration_id/bysetpos marshal to null, which the
// API accepts on create), while the two slices default to a non-nil empty slice (they drop
// ,omitempty so an emptied set clears on update, and the API rejects a null slice on create).
func TestSubscriptionBuildCreateRequest_UnsetOptionals(t *testing.T) {
	ops := SubscriptionOps{}
	model := SubscriptionResourceTFModel{
		TargetType:              types.StringValue("email"),
		TargetValue:             types.StringValue("team@example.com"),
		InsightID:               types.Int64Value(9),
		Frequency:               types.StringValue("weekly"),
		Interval:                types.Int64Value(1),
		StartDate:               types.StringValue("2026-08-17T07:00:00Z"),
		IntegrationID:           types.Int64Null(),
		DashboardID:             types.Int64Null(),
		DashboardExportInsights: types.SetNull(types.Int64Type),
		ByWeekday:               types.SetNull(types.StringType),
		BySetPos:                types.Int64Null(),
		Enabled:                 types.BoolNull(),
		Title:                   types.StringNull(),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())
	assert.Nil(t, req.IntegrationID)
	assert.Nil(t, req.BySetPos)
	assert.Nil(t, req.Enabled)
	assert.Nil(t, req.Title)
	// Slices are explicit empty (non-nil) so they marshal as [] rather than being dropped.
	assert.NotNil(t, req.ByWeekday)
	assert.Empty(t, req.ByWeekday)
	assert.NotNil(t, req.DashboardExportInsights)
	assert.Empty(t, req.DashboardExportInsights)
}

// TestSubscriptionBuildCreateRequest_RruleFields covers the recurrence-tuning optionals.
func TestSubscriptionBuildCreateRequest_RruleFields(t *testing.T) {
	ops := SubscriptionOps{}
	enabled := true
	model := SubscriptionResourceTFModel{
		TargetType:              types.StringValue("slack"),
		TargetValue:             types.StringValue("C0B9A53J8RF|#reports"),
		IntegrationID:           types.Int64Value(1),
		DashboardID:             types.Int64Value(7),
		DashboardExportInsights: int64Set(t, 3),
		Frequency:               types.StringValue("monthly"),
		Interval:                types.Int64Value(1),
		StartDate:               types.StringValue("2026-08-17T07:00:00Z"),
		ByWeekday:               stringSet(t, "monday"),
		BySetPos:                types.Int64Value(1),
		Enabled:                 types.BoolValue(enabled),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())
	assert.ElementsMatch(t, []string{"monday"}, req.ByWeekday)
	require.NotNil(t, req.BySetPos)
	assert.Equal(t, int64(1), *req.BySetPos)
	require.NotNil(t, req.Enabled)
	assert.True(t, *req.Enabled)
}

// TestSubscriptionBuildUpdateRequest_ClearsTitle verifies removing title from config
// sends an empty string so the API clears it.
func TestSubscriptionBuildUpdateRequest_ClearsTitle(t *testing.T) {
	ops := SubscriptionOps{}
	base := SubscriptionResourceTFModel{
		TargetType:  types.StringValue("email"),
		TargetValue: types.StringValue("team@example.com"),
		InsightID:   types.Int64Value(9),
		Frequency:   types.StringValue("weekly"),
		Interval:    types.Int64Value(1),
		StartDate:   types.StringValue("2026-08-17T07:00:00Z"),
	}
	plan := base
	plan.Title = types.StringNull()
	state := base
	state.Title = types.StringValue("old title")

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError(), diags.Errors())
	require.NotNil(t, req.Title)
	assert.Equal(t, "", *req.Title)
}

// TestSubscriptionBuildUpdateRequest_ClearsOptionals verifies that removing the optional
// integration_id / bysetpos / byweekday / dashboard_export_insights from config produces an
// explicit clearing value on the wire (null for the int64s, [] for the slices), rather than
// being dropped and leaving the stale server value (the "inconsistent result after apply" bug).
func TestSubscriptionBuildUpdateRequest_ClearsOptionals(t *testing.T) {
	ops := SubscriptionOps{}
	base := SubscriptionResourceTFModel{
		TargetType:  types.StringValue("email"),
		TargetValue: types.StringValue("team@example.com"),
		InsightID:   types.Int64Value(9),
		Frequency:   types.StringValue("monthly"),
		Interval:    types.Int64Value(1),
		StartDate:   types.StringValue("2026-08-17T07:00:00Z"),
	}
	state := base
	state.IntegrationID = types.Int64Value(1)
	state.BySetPos = types.Int64Value(1)
	state.ByWeekday = stringSet(t, "monday")
	state.DashboardExportInsights = int64Set(t, 11, 22)

	plan := base
	plan.IntegrationID = types.Int64Null()
	plan.BySetPos = types.Int64Null()
	plan.ByWeekday = types.SetNull(types.StringType)
	plan.DashboardExportInsights = types.SetNull(types.Int64Type)

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError(), diags.Errors())

	// int64s: nil pointer marshals to null (,omitempty is dropped on these fields).
	assert.Nil(t, req.IntegrationID)
	assert.Nil(t, req.BySetPos)
	// slices: explicit non-nil empty marshals to [] rather than being omitted.
	assert.NotNil(t, req.ByWeekday)
	assert.Empty(t, req.ByWeekday)
	assert.NotNil(t, req.DashboardExportInsights)
	assert.Empty(t, req.DashboardExportInsights)
}

// TestNormalizeDateTimeToModel covers both branches: a parseable value is canonicalized to
// UTC "Z" form, and an unparseable value is passed through verbatim with a warning diagnostic.
func TestNormalizeDateTimeToModel(t *testing.T) {
	t.Run("canonicalizes parseable", func(t *testing.T) {
		var diags diag.Diagnostics
		got := normalizeDateTimeToModel("2026-08-17T07:00:00+00:00", &diags)
		require.False(t, diags.HasError(), diags.Errors())
		assert.Empty(t, diags.Warnings())
		assert.Equal(t, "2026-08-17T07:00:00Z", got.ValueString())
	})

	t.Run("warns and passes through unparseable", func(t *testing.T) {
		var diags diag.Diagnostics
		got := normalizeDateTimeToModel("not-a-date", &diags)
		require.False(t, diags.HasError(), diags.Errors())
		require.Len(t, diags.Warnings(), 1)
		assert.Contains(t, diags.Warnings()[0].Summary(), "Unparseable datetime")
		// value stored verbatim, not dropped or mangled.
		assert.Equal(t, "not-a-date", got.ValueString())
	})
}

// TestReconcileStartDate covers the Required start_date reconciliation: a config value
// that denotes the same instant as the API's must be preserved verbatim (else a Required
// attribute trips "provider produced inconsistent result"), a genuine drift adopts the
// API value, and a null config (import) is normalized.
func TestReconcileStartDate(t *testing.T) {
	t.Run("preserves the configured value when it is the same instant", func(t *testing.T) {
		var diags diag.Diagnostics
		got := reconcileStartDate(types.StringValue("2026-08-17T09:00:00+02:00"), "2026-08-17T07:00:00Z", &diags)
		require.False(t, diags.HasError(), diags.Errors())
		assert.Equal(t, "2026-08-17T09:00:00+02:00", got.ValueString())
	})

	t.Run("adopts the API value on a genuine drift", func(t *testing.T) {
		var diags diag.Diagnostics
		got := reconcileStartDate(types.StringValue("2026-08-17T07:00:00Z"), "2026-09-01T07:00:00Z", &diags)
		require.False(t, diags.HasError(), diags.Errors())
		assert.Equal(t, "2026-09-01T07:00:00Z", got.ValueString())
	})

	t.Run("normalizes when config is null (import)", func(t *testing.T) {
		var diags diag.Diagnostics
		got := reconcileStartDate(types.StringNull(), "2026-08-17T07:00:00+00:00", &diags)
		require.False(t, diags.HasError(), diags.Errors())
		assert.Equal(t, "2026-08-17T07:00:00Z", got.ValueString())
	})
}

// TestAIPromptConfigToModel covers the empty/{}/null -> null mapping and the config-aware
// whitelisting (server-added keys are filtered so they don't surface as drift).
func TestAIPromptConfigToModel(t *testing.T) {
	t.Run("empty, {} and null map to null", func(t *testing.T) {
		for _, in := range []string{"", "   ", "{}", "null"} {
			got := aiPromptConfigToModel(json.RawMessage(in), jsontypes.NewNormalizedNull())
			assert.True(t, got.IsNull(), "input %q should map to a null value", in)
		}
	})

	t.Run("filters the API blob to the user's configured keys", func(t *testing.T) {
		current := jsontypes.NewNormalizedValue(`{"tone":"formal"}`)
		got := aiPromptConfigToModel(json.RawMessage(`{"tone":"formal","server_added":true}`), current)
		require.False(t, got.IsNull())
		assert.JSONEq(t, `{"tone":"formal"}`, got.ValueString())
	})
}

// TestSubscriptionMapResponseToModel_RoundTrip covers the verbatim target_value, datetime
// normalization, computed read-only fields, and set mapping.
func TestSubscriptionMapResponseToModel_RoundTrip(t *testing.T) {
	ops := SubscriptionOps{}
	model := SubscriptionResourceTFModel{
		// start_date configured in canonical form; API echoes an offset form.
		StartDate: types.StringValue("2026-08-17T07:00:00Z"),
	}
	resp := httpclient.Subscription{
		ID:                      42,
		TargetType:              "slack",
		TargetValue:             "C0B9A53J8RF|#reports",
		IntegrationID:           util.Int64Ptr(1),
		Dashboard:               util.Int64Ptr(7),
		DashboardExportInsights: []int64{11, 22},
		Frequency:               "daily",
		Interval:                1,
		StartDate:               "2026-08-17T07:00:00+00:00",
		ByWeekday:               []string{"monday"},
		BySetPos:                util.Int64Ptr(1),
		Enabled:                 util.BoolPtr(true),
		Title:                   util.StringPtr("Daily key metrics"),
		ResourceType:            util.StringPtr("dashboard"),
		Summary:                 util.StringPtr("sent every day"),
		NextDeliveryDate:        util.StringPtr("2026-08-18T07:00:00Z"),
		CreatedAt:               util.StringPtr("2026-08-10T00:00:00Z"),
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assert.Equal(t, int64(42), model.ID.ValueInt64())
	assert.Equal(t, "slack", model.TargetType.ValueString())
	// target_value verbatim, no split/reconstruct.
	assert.Equal(t, "C0B9A53J8RF|#reports", model.TargetValue.ValueString())
	assert.Equal(t, int64(1), model.IntegrationID.ValueInt64())
	assert.Equal(t, int64(7), model.DashboardID.ValueInt64())
	assert.True(t, model.InsightID.IsNull())
	// start_date normalized to canonical RFC3339 (matches configured start_date, no diff).
	assert.Equal(t, "2026-08-17T07:00:00Z", model.StartDate.ValueString())
	assert.ElementsMatch(t, []string{"monday"}, setStrings(t, model.ByWeekday))
	assert.ElementsMatch(t, []int64{11, 22}, setInt64s(t, model.DashboardExportInsights))
	assert.Equal(t, int64(1), model.BySetPos.ValueInt64())
	assert.True(t, model.Enabled.ValueBool())
	assert.Equal(t, "Daily key metrics", model.Title.ValueString())
	assert.Equal(t, "dashboard", model.ResourceType.ValueString())
	assert.Equal(t, "sent every day", model.Summary.ValueString())
	assert.Equal(t, "2026-08-18T07:00:00Z", model.NextDeliveryDate.ValueString())
	assert.Equal(t, "2026-08-10T00:00:00Z", model.CreatedAt.ValueString())
}

// TestSubscriptionMapResponseToModel_NullOptionals verifies absent optionals map to null.
func TestSubscriptionMapResponseToModel_NullOptionals(t *testing.T) {
	ops := SubscriptionOps{}
	model := SubscriptionResourceTFModel{}
	resp := httpclient.Subscription{
		ID:          9,
		TargetType:  "email",
		TargetValue: "team@example.com",
		Insight:     util.Int64Ptr(5),
		Frequency:   "weekly",
		Interval:    1,
		StartDate:   "2026-08-17T07:00:00Z",
		// no byweekday, integration, dashboard, title, etc.
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assert.True(t, model.ByWeekday.IsNull())
	assert.True(t, model.DashboardExportInsights.IsNull())
	assert.True(t, model.IntegrationID.IsNull())
	assert.True(t, model.DashboardID.IsNull())
	assert.True(t, model.BySetPos.IsNull())
	assert.True(t, model.Title.IsNull())
	assert.True(t, model.Enabled.IsNull())
	assert.Equal(t, int64(5), model.InsightID.ValueInt64())
}

func setStrings(t *testing.T, s types.Set) []string {
	t.Helper()
	var out []string
	diags := s.ElementsAs(context.Background(), &out, false)
	require.False(t, diags.HasError(), diags.Errors())
	return out
}

func setInt64s(t *testing.T, s types.Set) []int64 {
	t.Helper()
	var out []int64
	diags := s.ElementsAs(context.Background(), &out, false)
	require.False(t, diags.HasError(), diags.Errors())
	return out
}
