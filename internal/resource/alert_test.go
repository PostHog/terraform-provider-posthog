package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

// The rule matrix lives in core.TestValidateQuietHoursWindows now that posthog_logs_alert
// shares it. What only this layer can show is that the adapter reaches those rules from
// this resource's nested shape, and how it handles element shapes core never sees.
func TestBlockedWindowsValidator(t *testing.T) {
	t.Run("reaches the shared rules", func(t *testing.T) {
		req := validator.SetRequest{
			Path:        path.Root("schedule_restriction").AtName("blocked_windows"),
			ConfigValue: blockedWindowsSet(t, [2]string{"00:00", "06:00"}, [2]string{"06:00", "09:00"}),
		}
		resp := &validator.SetResponse{}

		blockedWindowsValidator{}.ValidateSet(context.Background(), req, resp)

		require.True(t, resp.Diagnostics.HasError(), "touching windows must be rejected")
		assert.Equal(t, "Quiet-hours windows overlap", resp.Diagnostics.Errors()[0].Summary())
		withPath, ok := resp.Diagnostics.Errors()[0].(diag.DiagnosticWithPath)
		require.True(t, ok, "diagnostic must be attribute-scoped")
		assert.Equal(t, req.Path.String(), withPath.Path().String(),
			"the diagnostic must point at this resource's attribute")
	})

	t.Run("valid windows produce nothing", func(t *testing.T) {
		req := validator.SetRequest{
			Path:        path.Root("schedule_restriction").AtName("blocked_windows"),
			ConfigValue: blockedWindowsSet(t, [2]string{"01:00", "05:00"}, [2]string{"12:00", "13:00"}),
		}
		resp := &validator.SetResponse{}

		blockedWindowsValidator{}.ValidateSet(context.Background(), req, resp)

		assert.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
	})

	// A null or unknown value cannot be reflected into a plain struct, so converting one
	// would fail the plan with the framework's "report this to the provider developer"
	// error for what is a configuration shape.
	t.Run("null and unknown sets are left alone", func(t *testing.T) {
		for name, value := range map[string]types.Set{
			"null":    types.SetNull(alertBlockedWindowObjectType),
			"unknown": types.SetUnknown(alertBlockedWindowObjectType),
		} {
			t.Run(name, func(t *testing.T) {
				resp := &validator.SetResponse{}
				blockedWindowsValidator{}.ValidateSet(context.Background(), validator.SetRequest{
					Path:        path.Root("schedule_restriction").AtName("blocked_windows"),
					ConfigValue: value,
				}, resp)
				assert.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
			})
		}
	})

	t.Run("an unknown element is skipped, not reported as a provider bug", func(t *testing.T) {
		set, diags := types.SetValue(alertBlockedWindowObjectType, []attr.Value{
			types.ObjectValueMust(alertBlockedWindowObjectType.AttrTypes, map[string]attr.Value{
				"start": types.StringValue("01:00"),
				"end":   types.StringValue("05:00"),
			}),
			types.ObjectUnknown(alertBlockedWindowObjectType.AttrTypes),
		})
		require.False(t, diags.HasError(), "%v", diags)

		resp := &validator.SetResponse{}
		blockedWindowsValidator{}.ValidateSet(context.Background(), validator.SetRequest{
			Path:        path.Root("schedule_restriction").AtName("blocked_windows"),
			ConfigValue: set,
		}, resp)

		assert.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
	})
}
