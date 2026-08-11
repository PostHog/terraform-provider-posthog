package datasource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapInsightToState_AllFields(t *testing.T) {
	insight := httpclient.Insight{
		ID:          123,
		ShortID:     util.StringPtr("Ty9UeOA2"),
		Name:        util.StringPtr("Weekly signups"),
		DerivedName: util.StringPtr("Weekly signups (derived)"),
		Description: util.StringPtr("Signups over time"),
		Tags:        []string{"growth", "weekly"},
		Dashboards:  []int32{11, 22},
		Favorited:   util.BoolPtr(true),
		Query: map[string]interface{}{
			"kind":      "InsightVizNode",
			"result":    []interface{}{1, 2, 3}, // server-injected
			"hogql":     "SELECT 1",             // server-injected
			"is_cached": true,                   // server-injected
			"source": map[string]interface{}{
				"kind":    "TrendsQuery",
				"version": float64(2), // server-injected
			},
		},
	}

	var model InsightDataSourceModel
	require.False(t, mapInsightToState(context.Background(), insight, &model).HasError())

	assert.Equal(t, int64(123), model.ID.ValueInt64())
	assert.Equal(t, "Ty9UeOA2", model.ShortID.ValueString())
	assert.Equal(t, "Weekly signups", model.Name.ValueString())
	assert.Equal(t, "Weekly signups (derived)", model.DerivedName.ValueString())
	assert.Equal(t, "Signups over time", model.Description.ValueString())
	assert.True(t, model.Favorited.ValueBool())
	assert.ElementsMatch(t, []string{"growth", "weekly"}, setToStrings(t, model.Tags))
	assert.ElementsMatch(t, []int32{11, 22}, setToInt32s(t, model.DashboardIDs))
	// Server-injected query fields are stripped from the normalized output.
	assert.JSONEq(t, `{"kind":"InsightVizNode","source":{"kind":"TrendsQuery"}}`, model.QueryJSON.ValueString())
}

func TestMapInsightToState_EmptyOptionalFieldsAreNull(t *testing.T) {
	insight := httpclient.Insight{
		ID:      7,
		ShortID: util.StringPtr("abc123"),
		Name:    util.StringPtr("   "), // whitespace-only -> null
	}

	var model InsightDataSourceModel
	require.False(t, mapInsightToState(context.Background(), insight, &model).HasError())

	assert.Equal(t, int64(7), model.ID.ValueInt64())
	assert.Equal(t, "abc123", model.ShortID.ValueString())
	assert.True(t, model.Name.IsNull(), "whitespace-only name should map to null")
	assert.True(t, model.DerivedName.IsNull(), "missing derived_name should map to null")
	assert.True(t, model.Description.IsNull(), "missing description should map to null")
	assert.True(t, model.Tags.IsNull(), "missing tags should map to null")
	assert.True(t, model.DashboardIDs.IsNull(), "missing dashboards should map to null")
	assert.True(t, model.Favorited.IsNull(), "missing favorited should map to null")
	assert.True(t, model.QueryJSON.IsNull(), "missing query should map to null")
}

func setToStrings(t *testing.T, s types.Set) []string {
	t.Helper()
	var out []string
	require.False(t, s.ElementsAs(context.Background(), &out, false).HasError())
	return out
}

func setToInt32s(t *testing.T, s types.Set) []int32 {
	t.Helper()
	var out []int32
	require.False(t, s.ElementsAs(context.Background(), &out, false).HasError())
	return out
}
