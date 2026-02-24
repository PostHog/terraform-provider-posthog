package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildLayoutPatch(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		declaredTiles  []TileTFModel
		apiTiles       []httpclient.DashboardTile
		wantPatchIDs   []int64                                                       // IDs of expected patch items (order-insensitive)
		wantMissingIDs []int64                                                       // expected missingInsightIDs
		wantPatchCount int                                                           // total expected patch items
		checkPatchItem func(t *testing.T, items []httpclient.DashboardTilePatchItem) // optional extra assertions
	}{
		"insight tile matched": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Null(),
					InsightID:   types.Int64Value(999),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedValue(`{"xs":{"x":0,"y":0,"w":6,"h":4}}`),
				},
			},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      77,
					Insight: &httpclient.DashboardTileInsight{ID: 999},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchIDs:   []int64{77},
			wantMissingIDs: nil,
			wantPatchCount: 1,
		},
		"text tile matched by tile_id": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(100),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("hello world"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      100,
					Text:    &httpclient.DashboardTileText{ID: 200, Body: "hello world"},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchIDs:   []int64{100},
			wantMissingIDs: nil,
			wantPatchCount: 1,
		},
		"missing insight returns in missing list": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Null(),
					InsightID:   types.Int64Value(888),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			apiTiles:       []httpclient.DashboardTile{},
			wantPatchIDs:   []int64{},
			wantMissingIDs: []int64{888},
			wantPatchCount: 0,
		},
		"new text tile creates without ID": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(0),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("brand new text tile"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			apiTiles:       []httpclient.DashboardTile{},
			wantPatchIDs:   []int64{0},
			wantMissingIDs: nil,
			wantPatchCount: 1,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				assert.Equal(t, int64(0), items[0].ID)
				require.NotNil(t, items[0].Text)
				assert.Equal(t, "brand new text tile", items[0].Text.Body)
			},
		},
		"existing text tile matched by body when tile_id unknown": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(0),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("existing text"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedValue(`{"sm":{"x":0,"y":0,"w":12,"h":1}}`),
				},
			},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      400,
					Text:    &httpclient.DashboardTileText{ID: 500, Body: "existing text"},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchIDs:   []int64{400},
			wantMissingIDs: nil,
			wantPatchCount: 1,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				// Should match existing tile by body, not create a new one.
				assert.Equal(t, int64(400), items[0].ID)
			},
		},
		"unmanaged tile layouts cleared": {
			declaredTiles: []TileTFModel{},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      50,
					Insight: &httpclient.DashboardTileInsight{ID: 1},
					Layouts: map[string]interface{}{"xs": map[string]interface{}{"x": float64(0)}},
				},
			},
			wantPatchIDs:   []int64{50},
			wantMissingIDs: nil,
			wantPatchCount: 1,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				assert.Equal(t, int64(50), items[0].ID)
				require.NotNil(t, items[0].Layouts)
				assert.Equal(t, map[string]interface{}{}, *items[0].Layouts)
			},
		},
		"unmanaged tile with empty layouts not patched": {
			declaredTiles: []TileTFModel{},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      60,
					Insight: &httpclient.DashboardTileInsight{ID: 2},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchIDs:   []int64{},
			wantMissingIDs: nil,
			wantPatchCount: 0,
		},
		"null color sends empty string to clear": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Null(),
					InsightID:   types.Int64Value(42),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      10,
					Insight: &httpclient.DashboardTileInsight{ID: 42},
					Layouts: map[string]interface{}{},
					Color:   strPtr("blue"),
				},
			},
			wantPatchIDs:   []int64{10},
			wantMissingIDs: nil,
			wantPatchCount: 1,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				require.NotNil(t, items[0].Color)
				assert.Equal(t, "", *items[0].Color, "null config color should send empty string to clear")
			},
		},
		"unmanaged text tile soft-deleted": {
			declaredTiles: []TileTFModel{},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      70,
					Text:    &httpclient.DashboardTileText{ID: 80, Body: "orphan text"},
					Layouts: map[string]interface{}{"sm": map[string]interface{}{"x": float64(0)}},
				},
			},
			wantPatchIDs:   []int64{70},
			wantMissingIDs: nil,
			wantPatchCount: 1,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				assert.Equal(t, int64(70), items[0].ID)
				require.NotNil(t, items[0].Deleted)
				assert.True(t, *items[0].Deleted)
			},
		},
		"unmanaged text tile with empty layouts still soft-deleted": {
			declaredTiles: []TileTFModel{},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      71,
					Text:    &httpclient.DashboardTileText{ID: 81, Body: "orphan text no layouts"},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchIDs:   []int64{71},
			wantMissingIDs: nil,
			wantPatchCount: 1,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				require.NotNil(t, items[0].Deleted)
				assert.True(t, *items[0].Deleted)
			},
		},
		"color set on patch item": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Null(),
					InsightID:   types.Int64Value(42),
					TextBody:    types.StringNull(),
					Color:       types.StringValue("blue"),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      10,
					Insight: &httpclient.DashboardTileInsight{ID: 42},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchIDs:   []int64{10},
			wantMissingIDs: nil,
			wantPatchCount: 1,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				require.NotNil(t, items[0].Color)
				assert.Equal(t, "blue", *items[0].Color)
			},
		},
		"duplicate body text tiles each match distinct API tile": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(0),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("same body"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
				{
					TileID:      types.Int64Value(0),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("same body"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      300,
					Text:    &httpclient.DashboardTileText{ID: 301, Body: "same body"},
					Layouts: map[string]interface{}{},
				},
				{
					ID:      400,
					Text:    &httpclient.DashboardTileText{ID: 401, Body: "same body"},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchIDs:   []int64{300, 400},
			wantMissingIDs: nil,
			wantPatchCount: 2,
			checkPatchItem: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				// Each config tile must target a different API tile.
				require.Len(t, items, 2)
				assert.NotEqual(t, items[0].ID, items[1].ID)
			},
		},
		"mixed insight and text tiles": {
			declaredTiles: []TileTFModel{
				{
					TileID:      types.Int64Null(),
					InsightID:   types.Int64Value(11),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
				{
					TileID:      types.Int64Value(200),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("some text"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      101,
					Insight: &httpclient.DashboardTileInsight{ID: 11},
					Layouts: map[string]interface{}{},
				},
				{
					ID:      200,
					Text:    &httpclient.DashboardTileText{ID: 300, Body: "some text"},
					Layouts: map[string]interface{}{},
				},
				{
					ID:      999,
					Insight: &httpclient.DashboardTileInsight{ID: 55},
					Layouts: map[string]interface{}{"xs": map[string]interface{}{"w": float64(6)}},
				},
			},
			wantPatchIDs:   []int64{101, 200, 999},
			wantMissingIDs: nil,
			wantPatchCount: 3,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			patchItems, missingIDs := buildLayoutPatch(ctx, tt.declaredTiles, tt.apiTiles)

			assert.Len(t, patchItems, tt.wantPatchCount, "unexpected number of patch items")

			if tt.wantMissingIDs == nil {
				assert.Empty(t, missingIDs)
			} else {
				assert.Equal(t, tt.wantMissingIDs, missingIDs)
			}

			if len(tt.wantPatchIDs) > 0 {
				gotIDs := make([]int64, len(patchItems))
				for i, item := range patchItems {
					gotIDs[i] = item.ID
				}
				assert.ElementsMatch(t, tt.wantPatchIDs, gotIDs)
			}

			if tt.checkPatchItem != nil {
				tt.checkPatchItem(t, patchItems)
			}
		})
	}
}

func TestBuildDeletePatch(t *testing.T) {
	tests := map[string]struct {
		apiTiles       []httpclient.DashboardTile
		wantPatchCount int
		checkItems     func(t *testing.T, items []httpclient.DashboardTilePatchItem)
	}{
		"insight tile layouts cleared": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      55,
					Insight: &httpclient.DashboardTileInsight{ID: 10},
					Layouts: map[string]interface{}{"xs": map[string]interface{}{"x": float64(0)}},
				},
			},
			wantPatchCount: 1,
			checkItems: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				assert.Equal(t, int64(55), items[0].ID)
				require.NotNil(t, items[0].Layouts)
				assert.Equal(t, map[string]interface{}{}, *items[0].Layouts)
				assert.Nil(t, items[0].Deleted)
			},
		},
		"insight tile with empty layouts skipped": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      55,
					Insight: &httpclient.DashboardTileInsight{ID: 10},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchCount: 0,
		},
		"text tile soft-deleted": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      300,
					Text:    &httpclient.DashboardTileText{ID: 400, Body: "goodbye"},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchCount: 1,
			checkItems: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 1)
				assert.Equal(t, int64(300), items[0].ID)
				require.NotNil(t, items[0].Deleted)
				assert.True(t, *items[0].Deleted)
			},
		},
		"empty dashboard is no-op": {
			apiTiles:       []httpclient.DashboardTile{},
			wantPatchCount: 0,
		},
		"unmanaged insight and text tiles both cleaned up": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      55,
					Insight: &httpclient.DashboardTileInsight{ID: 10},
					Layouts: map[string]interface{}{"sm": map[string]interface{}{"w": float64(6)}},
				},
				{
					ID:      99,
					Text:    &httpclient.DashboardTileText{ID: 100, Body: "orphan text"},
					Layouts: map[string]interface{}{},
				},
			},
			wantPatchCount: 2,
			checkItems: func(t *testing.T, items []httpclient.DashboardTilePatchItem) {
				t.Helper()
				require.Len(t, items, 2)
				var insightItem, textItem *httpclient.DashboardTilePatchItem
				for i := range items {
					if items[i].ID == 55 {
						insightItem = &items[i]
					}
					if items[i].ID == 99 {
						textItem = &items[i]
					}
				}
				require.NotNil(t, insightItem)
				require.NotNil(t, insightItem.Layouts)
				assert.Equal(t, map[string]interface{}{}, *insightItem.Layouts)

				require.NotNil(t, textItem, "text tile should be deleted")
				require.NotNil(t, textItem.Deleted)
				assert.True(t, *textItem.Deleted)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			patchItems := buildDeletePatch(tt.apiTiles)

			assert.Len(t, patchItems, tt.wantPatchCount)

			if tt.checkItems != nil {
				tt.checkItems(t, patchItems)
			}
		})
	}
}

func TestMapTilesToState(t *testing.T) {
	tests := map[string]struct {
		apiTiles    []httpclient.DashboardTile
		configTiles []TileTFModel
		wantLen     int
		checkResult func(t *testing.T, result []TileTFModel)
	}{
		"import case returns API order": {
			apiTiles: []httpclient.DashboardTile{
				{ID: 1, Insight: &httpclient.DashboardTileInsight{ID: 10}, Layouts: map[string]interface{}{}},
				{ID: 2, Insight: &httpclient.DashboardTileInsight{ID: 20}, Layouts: map[string]interface{}{}},
				{ID: 3, Text: &httpclient.DashboardTileText{ID: 30, Body: "hi"}, Layouts: map[string]interface{}{}},
			},
			configTiles: nil,
			wantLen:     3,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 3)
				assert.Equal(t, int64(1), result[0].TileID.ValueInt64())
				assert.Equal(t, int64(2), result[1].TileID.ValueInt64())
				assert.Equal(t, int64(3), result[2].TileID.ValueInt64())
			},
		},
		"normal read preserves config order": {
			apiTiles: []httpclient.DashboardTile{
				{ID: 11, Insight: &httpclient.DashboardTileInsight{ID: 100}, Layouts: map[string]interface{}{}},
				{ID: 22, Insight: &httpclient.DashboardTileInsight{ID: 200}, Layouts: map[string]interface{}{}},
			},
			// Config declares insight 200 first, then 100 — opposite of API order.
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(22),
					InsightID:   types.Int64Value(200),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
				{
					TileID:      types.Int64Value(11),
					InsightID:   types.Int64Value(100),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 2,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 2)
				// First result should be insight 200 (config order).
				assert.Equal(t, int64(200), result[0].InsightID.ValueInt64())
				// Second result should be insight 100.
				assert.Equal(t, int64(100), result[1].InsightID.ValueInt64())
			},
		},
		"missing API tile omitted": {
			apiTiles: []httpclient.DashboardTile{},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Null(),
					InsightID:   types.Int64Value(999),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 0,
		},
		"text tile matched by tile_id": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      100,
					Text:    &httpclient.DashboardTileText{ID: 200, Body: "some text"},
					Layouts: map[string]interface{}{},
				},
			},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(100),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("some text"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 1,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 1)
				assert.Equal(t, int64(100), result[0].TileID.ValueInt64())
				assert.Equal(t, "some text", result[0].TextBody.ValueString())
			},
		},
		"text tile with unknown tile_id matched by body": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      150,
					Text:    &httpclient.DashboardTileText{ID: 250, Body: "my note"},
					Layouts: map[string]interface{}{},
				},
			},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Unknown(),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("my note"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 1,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 1)
				assert.Equal(t, int64(150), result[0].TileID.ValueInt64())
				assert.Equal(t, "my note", result[0].TextBody.ValueString())
			},
		},
		"text tile with zero tile_id matched by body": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      160,
					Text:    &httpclient.DashboardTileText{ID: 260, Body: "zero tile"},
					Layouts: map[string]interface{}{},
				},
			},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(0),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("zero tile"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 1,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 1)
				assert.Equal(t, int64(160), result[0].TileID.ValueInt64())
				assert.Equal(t, "zero tile", result[0].TextBody.ValueString())
			},
		},
		"text tile with wrong tile_id falls back to body match": {
			// Config tile_id doesn't match any text tile in API — falls back to body.
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      500,
					Insight: &httpclient.DashboardTileInsight{ID: 42},
					Layouts: map[string]interface{}{},
				},
				{
					ID:      600,
					Text:    &httpclient.DashboardTileText{ID: 700, Body: "my text"},
					Layouts: map[string]interface{}{},
				},
			},
			configTiles: []TileTFModel{
				{
					// tile_id=500 is actually the insight tile's ID, wrongly assigned by UseStateForUnknown.
					TileID:      types.Int64Value(500),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("my text"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 1,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 1)
				// Should fall back to body match and find tile 600.
				assert.Equal(t, int64(600), result[0].TileID.ValueInt64())
				assert.Equal(t, "my text", result[0].TextBody.ValueString())
			},
		},
		"duplicate body text tiles each map to distinct API tile": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      300,
					Text:    &httpclient.DashboardTileText{ID: 301, Body: "same body"},
					Layouts: map[string]interface{}{},
				},
				{
					ID:      400,
					Text:    &httpclient.DashboardTileText{ID: 401, Body: "same body"},
					Layouts: map[string]interface{}{},
				},
			},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(0),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("same body"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
				{
					TileID:      types.Int64Value(0),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("same body"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 2,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 2)
				// Each config tile must map to a different API tile.
				assert.NotEqual(t, result[0].TileID.ValueInt64(), result[1].TileID.ValueInt64())
				ids := []int64{result[0].TileID.ValueInt64(), result[1].TileID.ValueInt64()}
				assert.ElementsMatch(t, []int64{300, 400}, ids)
			},
		},
		"api color suppressed when config color is null": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      42,
					Insight: &httpclient.DashboardTileInsight{ID: 100},
					Layouts: map[string]interface{}{},
					Color:   strPtr("blue"),
				},
			},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(42),
					InsightID:   types.Int64Value(100),
					TextBody:    types.StringNull(),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 1,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 1)
				assert.True(t, result[0].Color.IsNull(), "API color should be suppressed when config is null")
			},
		},
		"api color preserved when config color is set": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      42,
					Insight: &httpclient.DashboardTileInsight{ID: 100},
					Layouts: map[string]interface{}{},
					Color:   strPtr("blue"),
				},
			},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(42),
					InsightID:   types.Int64Value(100),
					TextBody:    types.StringNull(),
					Color:       types.StringValue("blue"),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 1,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 1)
				assert.Equal(t, "blue", result[0].Color.ValueString())
			},
		},
		"text tile matched by tile_id among duplicates": {
			apiTiles: []httpclient.DashboardTile{
				{
					ID:      190,
					Text:    &httpclient.DashboardTileText{ID: 290, Body: "same body"},
					Layouts: map[string]interface{}{},
				},
				{
					ID:      200,
					Text:    &httpclient.DashboardTileText{ID: 300, Body: "same body"},
					Layouts: map[string]interface{}{},
				},
			},
			configTiles: []TileTFModel{
				{
					TileID:      types.Int64Value(200),
					InsightID:   types.Int64Null(),
					TextBody:    types.StringValue("same body"),
					Color:       types.StringNull(),
					LayoutsJSON: jsontypes.NewNormalizedNull(),
				},
			},
			wantLen: 1,
			checkResult: func(t *testing.T, result []TileTFModel) {
				t.Helper()
				require.Len(t, result, 1)
				// Should match by tile_id=200.
				assert.Equal(t, int64(200), result[0].TileID.ValueInt64())
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, diags := mapTilesToState(tt.apiTiles, tt.configTiles)
			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Len(t, result, tt.wantLen)

			if tt.checkResult != nil {
				tt.checkResult(t, result)
			}
		})
	}
}

func TestApiTileToTFModel(t *testing.T) {
	tests := map[string]struct {
		apiTile     httpclient.DashboardTile
		checkResult func(t *testing.T, result TileTFModel)
	}{
		"insight tile": {
			apiTile: httpclient.DashboardTile{
				ID:      42,
				Insight: &httpclient.DashboardTileInsight{ID: 42},
				Layouts: map[string]interface{}{"xs": map[string]interface{}{"x": float64(0), "y": float64(0)}},
				Color:   strPtr("blue"),
			},
			checkResult: func(t *testing.T, result TileTFModel) {
				t.Helper()
				assert.Equal(t, int64(42), result.TileID.ValueInt64())
				assert.Equal(t, int64(42), result.InsightID.ValueInt64())
				assert.True(t, result.TextBody.IsNull(), "TextBody should be null for insight tile")
				assert.Equal(t, "blue", result.Color.ValueString())
				assert.False(t, result.LayoutsJSON.IsNull(), "LayoutsJSON should be set")
			},
		},
		"text tile": {
			apiTile: httpclient.DashboardTile{
				ID:      10,
				Text:    &httpclient.DashboardTileText{ID: 20, Body: "hello"},
				Layouts: map[string]interface{}{},
			},
			checkResult: func(t *testing.T, result TileTFModel) {
				t.Helper()
				assert.Equal(t, int64(10), result.TileID.ValueInt64())
				assert.True(t, result.InsightID.IsNull(), "InsightID should be null for text tile")
				assert.Equal(t, "hello", result.TextBody.ValueString())
				assert.True(t, result.LayoutsJSON.IsNull(), "LayoutsJSON should be null for empty layouts")
			},
		},
		"nil color": {
			apiTile: httpclient.DashboardTile{
				ID:      5,
				Insight: &httpclient.DashboardTileInsight{ID: 5},
				Layouts: map[string]interface{}{},
				Color:   nil,
			},
			checkResult: func(t *testing.T, result TileTFModel) {
				t.Helper()
				assert.True(t, result.Color.IsNull(), "Color should be null when API color is nil")
			},
		},
		"empty color string": {
			apiTile: httpclient.DashboardTile{
				ID:      6,
				Insight: &httpclient.DashboardTileInsight{ID: 6},
				Layouts: map[string]interface{}{},
				Color:   strPtr(""),
			},
			checkResult: func(t *testing.T, result TileTFModel) {
				t.Helper()
				assert.True(t, result.Color.IsNull(), "Color should be null when API color is empty string")
			},
		},
		"empty layouts": {
			apiTile: httpclient.DashboardTile{
				ID:      7,
				Insight: &httpclient.DashboardTileInsight{ID: 7},
				Layouts: map[string]interface{}{},
			},
			checkResult: func(t *testing.T, result TileTFModel) {
				t.Helper()
				assert.True(t, result.LayoutsJSON.IsNull(), "LayoutsJSON should be null for empty layouts map")
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := apiTileToTFModel(tt.apiTile)
			tt.checkResult(t, result)
		})
	}
}

func TestReAssociateInsight(t *testing.T) {
	tests := map[string]struct {
		insightDashboards []int32 // dashboards returned by GetInsight
		targetDashboardID int64
		wantPatch         bool // whether a PATCH should be sent
		getError          bool // simulate GetInsight failure
	}{
		"already associated": {
			insightDashboards: []int32{1, 2, 3},
			targetDashboardID: 2,
			wantPatch:         false,
		},
		"not associated": {
			insightDashboards: []int32{1, 3},
			targetDashboardID: 2,
			wantPatch:         true,
		},
		"get error": {
			targetDashboardID: 2,
			getError:          true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var patchCalled bool
			var patchedDashboards []int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodGet:
					if tt.getError {
						w.WriteHeader(http.StatusInternalServerError)
						_, _ = w.Write([]byte(`{"detail":"server error"}`))
						return
					}
					resp := httpclient.Insight{
						ID:         42,
						Dashboards: tt.insightDashboards,
					}
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(resp)

				case http.MethodPatch:
					patchCalled = true
					var req httpclient.InsightRequest
					_ = json.NewDecoder(r.Body).Decode(&req)
					patchedDashboards = req.Dashboards
					resp := httpclient.Insight{
						ID:         42,
						Dashboards: req.Dashboards,
					}
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(resp)
				}
			}))
			defer server.Close()

			client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
			c := &client

			err := reAssociateInsight(context.Background(), c, "proj-1", tt.targetDashboardID, 42)

			if tt.getError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantPatch, patchCalled, "PATCH called")

			if tt.wantPatch {
				// Should contain original dashboards plus the target.
				assert.Contains(t, patchedDashboards, int32(tt.targetDashboardID))
				for _, d := range tt.insightDashboards {
					assert.Contains(t, patchedDashboards, d)
				}
			}
		})
	}
}

func TestModifyResourcePlan(t *testing.T) {
	makeTiles := func(tiles []TileTFModel) types.List {
		t.Helper()
		list, diags := types.ListValueFrom(context.Background(), tilesObjectType, tiles)
		require.False(t, diags.HasError(), "ListValueFrom: %v", diags)
		return list
	}

	tests := map[string]struct {
		planModel  *DashboardLayoutTFModel
		stateModel *DashboardLayoutTFModel
		check      func(t *testing.T, resp *tfresource.ModifyPlanResponse)
	}{
		"copies tile_id from state to plan by insight_id": {
			planModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Value(100),
						TextBody:    types.StringNull(),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			stateModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Value(777),
						InsightID:   types.Int64Value(100),
						TextBody:    types.StringNull(),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			check: func(t *testing.T, resp *tfresource.ModifyPlanResponse) {
				t.Helper()
				require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				tiles := extractPlanTiles(t, resp)
				require.Len(t, tiles, 1)
				assert.Equal(t, int64(777), tiles[0].TileID.ValueInt64())
			},
		},
		"copies tile_id from state to plan by text_body": {
			planModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("hello world"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			stateModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Value(888),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("hello world"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			check: func(t *testing.T, resp *tfresource.ModifyPlanResponse) {
				t.Helper()
				require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				tiles := extractPlanTiles(t, resp)
				require.Len(t, tiles, 1)
				assert.Equal(t, int64(888), tiles[0].TileID.ValueInt64())
			},
		},
		"leaves unknown tile_id for new tiles": {
			planModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Value(100),
						TextBody:    types.StringNull(),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("new text tile"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			stateModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Value(777),
						InsightID:   types.Int64Value(100),
						TextBody:    types.StringNull(),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			check: func(t *testing.T, resp *tfresource.ModifyPlanResponse) {
				t.Helper()
				require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				tiles := extractPlanTiles(t, resp)
				require.Len(t, tiles, 2)
				// First tile matched by insight_id.
				assert.Equal(t, int64(777), tiles[0].TileID.ValueInt64())
				// Second tile is new — tile_id stays unknown.
				assert.True(t, tiles[1].TileID.IsUnknown(), "new tile should have unknown tile_id")
			},
		},
		"handles duplicate text bodies": {
			planModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("same body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("same body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			stateModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Value(500),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("same body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
					{
						TileID:      types.Int64Value(600),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("same body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			check: func(t *testing.T, resp *tfresource.ModifyPlanResponse) {
				t.Helper()
				require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				tiles := extractPlanTiles(t, resp)
				require.Len(t, tiles, 2)
				// Both should get unique tile_ids from state (consumed FIFO).
				ids := []int64{tiles[0].TileID.ValueInt64(), tiles[1].TileID.ValueInt64()}
				assert.ElementsMatch(t, []int64{500, 600}, ids)
			},
		},
		"text_body change resolves tile_id via positional fallback": {
			planModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Value(100),
						TextBody:    types.StringNull(),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("new body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			stateModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Value(777),
						InsightID:   types.Int64Value(100),
						TextBody:    types.StringNull(),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
					{
						TileID:      types.Int64Value(500),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("old body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			check: func(t *testing.T, resp *tfresource.ModifyPlanResponse) {
				t.Helper()
				require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				tiles := extractPlanTiles(t, resp)
				require.Len(t, tiles, 2)
				assert.Equal(t, int64(777), tiles[0].TileID.ValueInt64())
				// Body changed but positional fallback resolves to the old tile.
				assert.Equal(t, int64(500), tiles[1].TileID.ValueInt64())
				assert.False(t, tiles[1].TileID.IsUnknown())
			},
		},
		"duplicate bodies with one changed resolves correctly": {
			planModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("same body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("changed body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			stateModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Value(1)},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Value(500),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("same body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
					{
						TileID:      types.Int64Value(600),
						InsightID:   types.Int64Null(),
						TextBody:    types.StringValue("same body"),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			check: func(t *testing.T, resp *tfresource.ModifyPlanResponse) {
				t.Helper()
				require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				tiles := extractPlanTiles(t, resp)
				require.Len(t, tiles, 2)
				// First tile matched by body -> gets first state tile_id.
				assert.Equal(t, int64(500), tiles[0].TileID.ValueInt64())
				// Second tile body changed, positional fallback -> gets remaining state tile_id.
				assert.Equal(t, int64(600), tiles[1].TileID.ValueInt64())
			},
		},
		"no-op when state is nil (create)": {
			planModel: &DashboardLayoutTFModel{
				BaseInt64Identifiable: core.BaseInt64Identifiable{ID: types.Int64Unknown()},
				BaseProjectID:         core.BaseProjectID{ProjectID: types.StringValue("proj-1")},
				DashboardID:           types.Int64Value(1),
				Tiles: makeTiles([]TileTFModel{
					{
						TileID:      types.Int64Unknown(),
						InsightID:   types.Int64Value(100),
						TextBody:    types.StringNull(),
						Color:       types.StringNull(),
						LayoutsJSON: jsontypes.NewNormalizedNull(),
					},
				}),
			},
			stateModel: nil, // create — no prior state
			check: func(t *testing.T, resp *tfresource.ModifyPlanResponse) {
				t.Helper()
				require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
				tiles := extractPlanTiles(t, resp)
				require.Len(t, tiles, 1)
				// tile_id should remain unknown (no state to match against).
				assert.True(t, tiles[0].TileID.IsUnknown(), "tile_id should stay unknown on create")
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ops := &DashboardLayoutOps{}
			req, resp := buildModifyPlanReqResp(t, tt.planModel, tt.stateModel)
			ops.ModifyResourcePlan(context.Background(), req, resp)
			tt.check(t, resp)
		})
	}
}

func TestResolveTileIDs(t *testing.T) {
	t.Run("does not mutate planTiles", func(t *testing.T) {
		planTiles := []TileTFModel{
			{
				TileID:      types.Int64Unknown(),
				InsightID:   types.Int64Value(100),
				TextBody:    types.StringNull(),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
			{
				TileID:      types.Int64Unknown(),
				InsightID:   types.Int64Null(),
				TextBody:    types.StringValue("hello"),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
		}
		stateTiles := []TileTFModel{
			{
				TileID:      types.Int64Value(777),
				InsightID:   types.Int64Value(100),
				TextBody:    types.StringNull(),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
			{
				TileID:      types.Int64Value(888),
				InsightID:   types.Int64Null(),
				TextBody:    types.StringValue("hello"),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
		}

		resolved := resolveTileIDs(planTiles, stateTiles)

		require.NotNil(t, resolved, "should resolve tile IDs")
		assert.Equal(t, int64(777), resolved[0].TileID.ValueInt64())
		assert.Equal(t, int64(888), resolved[1].TileID.ValueInt64())

		// Original planTiles must not be modified.
		assert.True(t, planTiles[0].TileID.IsUnknown(), "planTiles[0].TileID should still be unknown")
		assert.True(t, planTiles[1].TileID.IsUnknown(), "planTiles[1].TileID should still be unknown")
	})

	t.Run("returns nil when no tile_id resolved", func(t *testing.T) {
		planTiles := []TileTFModel{
			{
				TileID:      types.Int64Unknown(),
				InsightID:   types.Int64Value(999),
				TextBody:    types.StringNull(),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
		}
		// State has no matching insight.
		stateTiles := []TileTFModel{
			{
				TileID:      types.Int64Value(111),
				InsightID:   types.Int64Value(222),
				TextBody:    types.StringNull(),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
		}

		resolved := resolveTileIDs(planTiles, stateTiles)
		assert.Nil(t, resolved, "should return nil when nothing resolved")
	})

	t.Run("skips tiles with known tile_id", func(t *testing.T) {
		planTiles := []TileTFModel{
			{
				TileID:      types.Int64Value(42),
				InsightID:   types.Int64Value(100),
				TextBody:    types.StringNull(),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
		}
		stateTiles := []TileTFModel{
			{
				TileID:      types.Int64Value(42),
				InsightID:   types.Int64Value(100),
				TextBody:    types.StringNull(),
				Color:       types.StringNull(),
				LayoutsJSON: jsontypes.NewNormalizedNull(),
			},
		}

		resolved := resolveTileIDs(planTiles, stateTiles)
		assert.Nil(t, resolved, "should return nil when all tile_ids already known")
	})
}

// strPtr is a helper to create a pointer to a string literal.
func strPtr(s string) *string {
	return &s
}

// buildModifyPlanReqResp constructs a ModifyPlanRequest and ModifyPlanResponse
// from plan and state models. Pass nil for stateModel to simulate a create (null state).
func buildModifyPlanReqResp(t *testing.T, planModel *DashboardLayoutTFModel, stateModel *DashboardLayoutTFModel) (tfresource.ModifyPlanRequest, *tfresource.ModifyPlanResponse) {
	t.Helper()
	ctx := context.Background()
	s := (&DashboardLayoutOps{}).Schema()

	schemaType := s.Type().TerraformType(ctx)

	plan := tfsdk.Plan{Schema: s}
	if planModel != nil {
		plan.Raw = tftypes.NewValue(schemaType, nil) // placeholder, overwritten by Set
		diags := plan.Set(ctx, planModel)
		require.False(t, diags.HasError(), "plan.Set: %v", diags)
	} else {
		plan.Raw = tftypes.NewValue(schemaType, nil)
	}

	state := tfsdk.State{Schema: s}
	if stateModel != nil {
		state.Raw = tftypes.NewValue(schemaType, nil)
		diags := state.Set(ctx, stateModel)
		require.False(t, diags.HasError(), "state.Set: %v", diags)
	} else {
		state.Raw = tftypes.NewValue(schemaType, nil)
	}

	req := tfresource.ModifyPlanRequest{
		Plan:  plan,
		State: state,
	}
	resp := &tfresource.ModifyPlanResponse{
		Plan: plan,
	}
	return req, resp
}

// extractPlanTiles extracts TileTFModel slice from a ModifyPlanResponse.
func extractPlanTiles(t *testing.T, resp *tfresource.ModifyPlanResponse) []TileTFModel {
	t.Helper()
	var model DashboardLayoutTFModel
	diags := resp.Plan.Get(context.Background(), &model)
	require.False(t, diags.HasError(), "resp.Plan.Get: %v", diags)
	tiles, diags := extractTilesFromModel(context.Background(), model)
	require.False(t, diags.HasError(), "extractTilesFromModel: %v", diags)
	return tiles
}
