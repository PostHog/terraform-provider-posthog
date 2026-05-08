package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

type TileTFModel struct {
	TileID          types.Int64          `tfsdk:"tile_id"`
	InsightID       types.Int64          `tfsdk:"insight_id"`
	TextBody        types.String         `tfsdk:"text_body"`
	Color           types.String         `tfsdk:"color"`
	ShowDescription types.Bool           `tfsdk:"show_description"`
	LayoutsJSON     jsontypes.Normalized `tfsdk:"layouts_json"`
}

func (t TileTFModel) IsInsightTile() bool {
	return !t.InsightID.IsNull() && !t.InsightID.IsUnknown()
}

func (t TileTFModel) IsTextTile() bool {
	return !t.TextBody.IsNull() && !t.TextBody.IsUnknown()
}

type DashboardLayoutTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	DashboardID types.Int64 `tfsdk:"dashboard_id"`
	Tiles       types.List  `tfsdk:"tiles"` // element type: TileTFModel
}

var tilesObjectType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"tile_id":          types.Int64Type,
		"insight_id":       types.Int64Type,
		"text_body":        types.StringType,
		"color":            types.StringType,
		"show_description": types.BoolType,
		"layouts_json":     jsontypes.NormalizedType{},
	},
}

type DashboardLayoutOps struct {
	// planTilesByDashboard passes declared tiles from BuildUpdateRequest to Update,
	// keyed by dashboard ID. A sync.Map is used because Terraform may process
	// multiple resources of the same type concurrently; each dashboard ID is
	// written by one goroutine and consumed (LoadAndDelete) by the same one.
	planTilesByDashboard sync.Map // map[int64][]TileTFModel
}

type stateTileLookups struct {
	insightToTileID  map[int64]int64
	bodyToTileIDs    map[string][]int64
	tileIDToStateIdx map[int64]int
}

func NewDashboardLayout() resource.Resource {
	return core.NewGenericResource[DashboardLayoutTFModel, httpclient.DashboardLayoutPatchRequest, httpclient.DashboardLayoutResponse](
		&DashboardLayoutOps{},
		core.ProjectScopedImportParser[DashboardLayoutTFModel](),
	)
}

func (o *DashboardLayoutOps) ResourceName() string {
	return "dashboard_layout"
}

func (o *DashboardLayoutOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages the tile layout of a PostHog dashboard. This resource is fully authoritative: it manages all tiles on the dashboard. Unmanaged insight tiles will have their layouts cleared; unmanaged text tiles will be deleted. On destroy, all text tiles are removed and insight tile layouts are cleared.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Resource ID (same value as dashboard_id).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"dashboard_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "PostHog dashboard ID.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"tiles": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "Ordered list of tiles to manage on the dashboard.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"tile_id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Server-assigned tile ID. Populated after the first apply.",
						},
						"insight_id": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "ID of the insight to display. Exactly one of insight_id or text_body must be set.",
							Validators: []validator.Int64{
								int64validator.ExactlyOneOf(
									path.MatchRelative().AtParent().AtName("text_body"),
								),
							},
						},
						"text_body": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Markdown body for a text tile (max 4000 characters). Exactly one of insight_id or text_body must be set.",
							Validators: []validator.String{
								stringvalidator.LengthAtMost(4000),
							},
						},
						"color": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Background color of the tile. Valid values are defined by the PostHog API; see [InsightColor in types.ts](https://github.com/PostHog/posthog/blob/master/frontend/src/types.ts#L2154) for the current list.",
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"show_description": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Whether to show the insight description on the tile. Omit the field to clear it back to the PostHog API default (`null`).",
						},
						"layouts_json": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							CustomType:          jsontypes.NormalizedType{},
							MarkdownDescription: "JSON object with breakpoint keys `sm` and/or `xs`, each containing position properties: `x`, `y`, `w`, `h` (required), and optionally `minW`, `minH` (e.g. `{\"sm\":{\"x\":0,\"y\":0,\"w\":6,\"h\":5},\"xs\":{\"x\":0,\"y\":0,\"w\":1,\"h\":5}}`). Semantic JSON equality is used to suppress phantom diffs.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

// ModifyResourcePlan copies tile_id from state into the plan for existing tiles.
// tile_id is Computed-only (server-assigned), so without this it would be Unknown in
// every plan. Resolving it early means buildLayoutPatch can match text tiles by tile_id
// instead of falling back to body matching, which enables in-place updates when a text
// tile's body changes and avoids ambiguity with duplicate bodies.
func (o *DashboardLayoutOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip during destroy.
	if req.Plan.Raw.IsNull() {
		return
	}
	// Skip during create — no prior state to match against.
	if req.State.Raw.IsNull() {
		return
	}

	var plan DashboardLayoutTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state DashboardLayoutTFModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	planTiles, d := extractTilesFromModel(ctx, plan)
	resp.Diagnostics.Append(d...)
	stateTiles, d := extractTilesFromModel(ctx, state)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() || planTiles == nil || stateTiles == nil {
		return
	}

	if resolved := resolveTileIDs(planTiles, stateTiles); resolved != nil {
		tilesList, d := types.ListValueFrom(ctx, tilesObjectType, resolved)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("tiles"), tilesList)...)
	}
}

// resolveTileIDs matches plan tiles to state tiles by semantic identity and copies
// the tile_id from state into the plan. It returns a new slice with resolved tile_ids,
// or nil if no tile_id was resolved.
func resolveTileIDs(planTiles, stateTiles []TileTFModel) []TileTFModel {
	lookups := buildStateTileLookups(stateTiles)

	result := make([]TileTFModel, len(planTiles))
	copy(result, planTiles)

	// Pass 1: Match by semantic identity (insight_id, text_body).
	matchedStateIdx := make(map[int]bool, len(stateTiles))
	modified := false
	for i, pt := range planTiles {
		if !pt.TileID.IsUnknown() {
			continue
		}
		tileID, found := resolveTileIDFromState(pt, &lookups)
		if found && tileID != 0 {
			result[i].TileID = types.Int64Value(tileID)
			modified = true
			matchedStateIdx[lookups.tileIDToStateIdx[tileID]] = true
		}
	}

	// Pass 2: Positional fallback for remaining unmatched text tiles.
	unmatchedStateTileIDs := collectUnmatchedStateTextTileIDs(stateTiles, matchedStateIdx)
	positionalIdx := 0
	for i := range result {
		if !result[i].TileID.IsUnknown() || !result[i].IsTextTile() {
			continue
		}
		if positionalIdx >= len(unmatchedStateTileIDs) {
			break
		}
		tileID := unmatchedStateTileIDs[positionalIdx]
		positionalIdx++
		if tileID != 0 {
			result[i].TileID = types.Int64Value(tileID)
			modified = true
		}
	}

	if !modified {
		return nil
	}
	return result
}

func buildStateTileLookups(stateTiles []TileTFModel) stateTileLookups {
	lookups := stateTileLookups{
		insightToTileID:  make(map[int64]int64, len(stateTiles)),
		bodyToTileIDs:    make(map[string][]int64, len(stateTiles)),
		tileIDToStateIdx: make(map[int64]int, len(stateTiles)),
	}

	for si, st := range stateTiles {
		tid := st.TileID.ValueInt64()
		lookups.tileIDToStateIdx[tid] = si
		if st.IsInsightTile() {
			lookups.insightToTileID[st.InsightID.ValueInt64()] = tid
			continue
		}
		if st.IsTextTile() {
			body := st.TextBody.ValueString()
			lookups.bodyToTileIDs[body] = append(lookups.bodyToTileIDs[body], tid)
		}
	}

	return lookups
}

func resolveTileIDFromState(tile TileTFModel, lookups *stateTileLookups) (int64, bool) {
	if tile.IsInsightTile() {
		tileID, found := lookups.insightToTileID[tile.InsightID.ValueInt64()]
		return tileID, found
	}
	if tile.IsTextTile() {
		return lookups.takeNextTextTileID(tile.TextBody.ValueString())
	}
	return 0, false
}

func (l *stateTileLookups) takeNextTextTileID(body string) (int64, bool) {
	ids := l.bodyToTileIDs[body]
	if len(ids) == 0 {
		return 0, false
	}
	tileID := ids[0]
	l.bodyToTileIDs[body] = ids[1:]
	return tileID, true
}

func collectUnmatchedStateTextTileIDs(stateTiles []TileTFModel, matchedStateIdx map[int]bool) []int64 {
	var unmatchedStateTileIDs []int64
	for si, st := range stateTiles {
		if matchedStateIdx[si] || !st.IsTextTile() {
			continue
		}
		unmatchedStateTileIDs = append(unmatchedStateTileIDs, st.TileID.ValueInt64())
	}
	return unmatchedStateTileIDs
}

func (o *DashboardLayoutOps) BuildCreateRequest(_ context.Context, _ DashboardLayoutTFModel) (httpclient.DashboardLayoutPatchRequest, diag.Diagnostics) {
	// BuildCreateRequest returns an empty request because Create performs its own
	// GET-reconcile-PATCH flow using the HTTP client directly.
	return httpclient.DashboardLayoutPatchRequest{}, nil
}

func (o *DashboardLayoutOps) BuildUpdateRequest(ctx context.Context, plan, _ DashboardLayoutTFModel) (httpclient.DashboardLayoutPatchRequest, diag.Diagnostics) {
	tiles, diags := extractTilesFromModel(ctx, plan)
	if diags.HasError() {
		return httpclient.DashboardLayoutPatchRequest{}, diags
	}
	o.planTilesByDashboard.Store(plan.DashboardID.ValueInt64(), tiles)
	return httpclient.DashboardLayoutPatchRequest{}, diags
}

func (o *DashboardLayoutOps) MapResponseToModel(ctx context.Context, resp httpclient.DashboardLayoutResponse, model *DashboardLayoutTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Set ID and DashboardID from response. ID is required for GenericResource to function;
	// DashboardID is set here to handle the import case where it was not provided by the import parser.
	model.ID = types.Int64Value(resp.ID)
	model.DashboardID = types.Int64Value(resp.ID)

	configTiles, d := extractTilesFromModel(ctx, *model)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	stateTiles, d := mapTilesToState(resp.Tiles, configTiles)
	diags.Append(d...)

	if stateTiles == nil {
		stateTiles = []TileTFModel{}
	}

	tilesList, d := types.ListValueFrom(ctx, tilesObjectType, stateTiles)
	diags.Append(d...)

	model.Tiles = tilesList
	return diags
}

func (o *DashboardLayoutOps) Create(ctx context.Context, client httpclient.PosthogClient, model DashboardLayoutTFModel, _ httpclient.DashboardLayoutPatchRequest) (httpclient.DashboardLayoutResponse, error) {
	declaredTiles, diags := extractTilesFromModel(ctx, model)
	if diags.HasError() {
		return httpclient.DashboardLayoutResponse{}, fmt.Errorf("extracting tiles from config: %s", diags.Errors()[0].Detail())
	}

	resp, _, err := reconcileAndPatch(ctx, &client, model.GetEffectiveProjectID(), model.DashboardID.ValueInt64(), declaredTiles)
	return resp, err
}

func (o *DashboardLayoutOps) Read(ctx context.Context, client httpclient.PosthogClient, model DashboardLayoutTFModel) (httpclient.DashboardLayoutResponse, httpclient.HTTPStatusCode, error) {
	projectID := model.GetEffectiveProjectID()

	// Import fallback: during import, DashboardID may not be set yet. Fall back to the
	// resource ID (which is set by the import parser from the last path segment).
	dashboardID := strconv.FormatInt(model.DashboardID.ValueInt64(), 10)
	if model.DashboardID.IsNull() || model.DashboardID.IsUnknown() || model.DashboardID.ValueInt64() == 0 {
		dashboardID = model.GetID()
	}

	c := &client
	return c.GetDashboardLayout(ctx, projectID, dashboardID)
}

func (o *DashboardLayoutOps) Update(ctx context.Context, client httpclient.PosthogClient, model DashboardLayoutTFModel, _ httpclient.DashboardLayoutPatchRequest) (httpclient.DashboardLayoutResponse, httpclient.HTTPStatusCode, error) {
	dashboardID := model.DashboardID.ValueInt64()
	raw, ok := o.planTilesByDashboard.LoadAndDelete(dashboardID)
	if !ok {
		return httpclient.DashboardLayoutResponse{}, 0, fmt.Errorf("plan tiles not found for dashboard %d; this is a provider bug", dashboardID)
	}
	return reconcileAndPatch(ctx, &client, model.GetEffectiveProjectID(), dashboardID, raw.([]TileTFModel))
}

// reconcileAndPatch performs a GET-reconcile-PATCH flow.
// It reads the current dashboard state, correlates declared tiles with API tiles,
// re-associates any missing insights, and applies the resulting patch.
func reconcileAndPatch(ctx context.Context, client *httpclient.PosthogClient, projectID string, dashboardIDInt int64, declaredTiles []TileTFModel) (httpclient.DashboardLayoutResponse, httpclient.HTTPStatusCode, error) {
	dashboardID := strconv.FormatInt(dashboardIDInt, 10)

	apiResp, statusCode, err := client.GetDashboardLayout(ctx, projectID, dashboardID)
	if err != nil {
		return httpclient.DashboardLayoutResponse{}, statusCode, fmt.Errorf("reading dashboard %s: %w", dashboardID, err)
	}

	patchItems, missingInsightIDs := buildLayoutPatch(ctx, declaredTiles, apiResp.Tiles)

	if len(missingInsightIDs) > 0 {
		tflog.Info(ctx, "Re-associating insights removed from dashboard", map[string]any{
			"insight_ids":  missingInsightIDs,
			"dashboard_id": dashboardID,
		})

		for _, insightID := range missingInsightIDs {
			if err := reAssociateInsight(ctx, client, projectID, dashboardIDInt, insightID); err != nil {
				return httpclient.DashboardLayoutResponse{}, 0, fmt.Errorf("re-associating insight %d with dashboard %s: %w", insightID, dashboardID, err)
			}
		}

		// Re-GET after re-association so the new tile entries are visible.
		apiResp, statusCode, err = client.GetDashboardLayout(ctx, projectID, dashboardID)
		if err != nil {
			return httpclient.DashboardLayoutResponse{}, statusCode, fmt.Errorf("reading dashboard %s after re-association: %w", dashboardID, err)
		}

		patchItems, missingInsightIDs = buildLayoutPatch(ctx, declaredTiles, apiResp.Tiles)
		if len(missingInsightIDs) > 0 {
			return httpclient.DashboardLayoutResponse{}, 0, fmt.Errorf("insights %v still not found on dashboard %s after re-association", missingInsightIDs, dashboardID)
		}
	}

	if len(patchItems) == 0 {
		return apiResp, 200, nil
	}

	return client.UpdateDashboardLayout(ctx, projectID, dashboardID, httpclient.DashboardLayoutPatchRequest{Tiles: patchItems})
}

func (o *DashboardLayoutOps) Delete(ctx context.Context, client httpclient.PosthogClient, model DashboardLayoutTFModel) (httpclient.HTTPStatusCode, error) {
	projectID := model.GetEffectiveProjectID()
	dashboardID := strconv.FormatInt(model.DashboardID.ValueInt64(), 10)

	c := &client

	apiResp, statusCode, err := c.GetDashboardLayout(ctx, projectID, dashboardID)
	if err != nil {
		return statusCode, fmt.Errorf("reading dashboard %s for delete: %w", dashboardID, err)
	}

	patchItems := buildDeletePatch(apiResp.Tiles)
	if len(patchItems) == 0 {
		return 200, nil
	}

	_, statusCode, err = c.UpdateDashboardLayout(ctx, projectID, dashboardID, httpclient.DashboardLayoutPatchRequest{Tiles: patchItems})
	return statusCode, err
}

// extractTilesFromModel extracts tiles from a DashboardLayoutTFModel.
func extractTilesFromModel(ctx context.Context, model DashboardLayoutTFModel) ([]TileTFModel, diag.Diagnostics) {
	if model.Tiles.IsNull() || model.Tiles.IsUnknown() {
		return nil, nil
	}
	var tiles []TileTFModel
	diags := model.Tiles.ElementsAs(ctx, &tiles, false)
	return tiles, diags
}

// buildTileLookupMaps indexes API tiles by position: insight tiles keyed by insight ID,
// and text tiles keyed by tile ID. Values are indices into the apiTiles slice.
func buildTileLookupMaps(apiTiles []httpclient.DashboardTile) (insightIdx, textIdx map[int64]int) {
	insightIdx = make(map[int64]int, len(apiTiles))
	textIdx = make(map[int64]int, len(apiTiles))
	for i, t := range apiTiles {
		if t.Insight != nil {
			insightIdx[t.Insight.ID] = i
		}
		if t.Text != nil {
			textIdx[t.ID] = i
		}
	}
	return insightIdx, textIdx
}

// matchDeclaredTilesToAPI correlates declared config tiles with API tiles.
// For insight tiles, it matches by insight_id. For text tiles, it tries tile_id
// first, then falls back to body matching (skipping already-matched API tiles).
// Returns one result per declared tile (same order) and the set of matched API tile IDs.
func matchDeclaredTilesToAPI(declaredTiles []TileTFModel, apiTiles []httpclient.DashboardTile) ([]*httpclient.DashboardTile, map[int64]bool) {
	insightIdxMap, textIdxMap := buildTileLookupMaps(apiTiles)
	matchedIDs := make(map[int64]bool)
	matches := make([]*httpclient.DashboardTile, len(declaredTiles))

	for i, declared := range declaredTiles {
		match := findDeclaredTileMatch(declared, apiTiles, insightIdxMap, textIdxMap, matchedIDs)
		if match != nil {
			matchedIDs[match.ID] = true
			matches[i] = match
		}
	}

	return matches, matchedIDs
}

func findDeclaredTileMatch(declared TileTFModel, apiTiles []httpclient.DashboardTile, insightIdxMap, textIdxMap map[int64]int, matchedIDs map[int64]bool) *httpclient.DashboardTile {
	if declared.IsInsightTile() {
		return matchInsightTileByInsightID(declared, apiTiles, insightIdxMap)
	}
	if declared.IsTextTile() {
		return matchTextTileByTileIDOrBody(declared, apiTiles, textIdxMap, matchedIDs)
	}
	return nil
}

func matchInsightTileByInsightID(declared TileTFModel, apiTiles []httpclient.DashboardTile, insightIdxMap map[int64]int) *httpclient.DashboardTile {
	idx, ok := insightIdxMap[declared.InsightID.ValueInt64()]
	if !ok {
		return nil
	}
	return &apiTiles[idx]
}

func matchTextTileByTileIDOrBody(declared TileTFModel, apiTiles []httpclient.DashboardTile, textIdxMap map[int64]int, matchedIDs map[int64]bool) *httpclient.DashboardTile {
	tileID := declared.TileID.ValueInt64()
	if tileID != 0 {
		if idx, ok := textIdxMap[tileID]; ok {
			return &apiTiles[idx]
		}
	}
	return matchTextTileByBody(declared.TextBody.ValueString(), apiTiles, matchedIDs)
}

func matchTextTileByBody(body string, apiTiles []httpclient.DashboardTile, matchedIDs map[int64]bool) *httpclient.DashboardTile {
	for i := range apiTiles {
		if apiTiles[i].Text == nil || matchedIDs[apiTiles[i].ID] {
			continue
		}
		if apiTiles[i].Text.Body == body {
			return &apiTiles[i]
		}
	}
	return nil
}

// buildLayoutPatch builds the set of tile patch items for a PATCH request.
// It correlates declared config tiles with API tiles and enforces authoritative
// behavior: unmanaged text tiles are soft-deleted, unmanaged insight tiles have
// their layouts cleared. Returns the patch items and a list of insight IDs that
// were declared but not found in the API.
func buildLayoutPatch(ctx context.Context, declaredTiles []TileTFModel, apiTiles []httpclient.DashboardTile) ([]httpclient.DashboardTilePatchItem, []int64) {
	matches, matchedIDs := matchDeclaredTilesToAPI(declaredTiles, apiTiles)

	var patchItems []httpclient.DashboardTilePatchItem
	var missingInsightIDs []int64

	for i, declared := range declaredTiles {
		if matches[i] != nil {
			patchItems = append(patchItems, buildTilePatchItem(ctx, matches[i].ID, declared))
		} else if declared.IsInsightTile() {
			missingInsightIDs = append(missingInsightIDs, declared.InsightID.ValueInt64())
		} else if declared.IsTextTile() {
			patchItems = append(patchItems, buildNewTextTilePatchItem(ctx, declared))
		}
	}

	// Authoritative enforcement: unmanaged text tiles are soft-deleted (they
	// are owned by this resource), unmanaged insight tiles just have their
	// layouts cleared (insights exist independently).
	for _, t := range apiTiles {
		if matchedIDs[t.ID] {
			continue
		}
		if t.Text != nil {
			deleted := true
			patchItems = append(patchItems, httpclient.DashboardTilePatchItem{
				ID:      t.ID,
				Deleted: &deleted,
			})
		} else if len(t.Layouts) > 0 {
			emptyLayouts := map[string]interface{}{}
			patchItems = append(patchItems, httpclient.DashboardTilePatchItem{
				ID:      t.ID,
				Layouts: &emptyLayouts,
			})
		}
	}

	return patchItems, missingInsightIDs
}

// buildDeletePatch builds the set of tile patch items to clean up the dashboard when
// the resource is destroyed. This is fully authoritative: ALL text tiles on the
// dashboard are soft-deleted and ALL insight tiles have their layouts cleared.
func buildDeletePatch(apiTiles []httpclient.DashboardTile) []httpclient.DashboardTilePatchItem {
	var patchItems []httpclient.DashboardTilePatchItem
	for _, t := range apiTiles {
		if t.Text != nil {
			deleted := true
			patchItems = append(patchItems, httpclient.DashboardTilePatchItem{
				ID:      t.ID,
				Deleted: &deleted,
			})
		} else if len(t.Layouts) > 0 {
			emptyLayouts := map[string]interface{}{}
			patchItems = append(patchItems, httpclient.DashboardTilePatchItem{
				ID:      t.ID,
				Layouts: &emptyLayouts,
			})
		}
	}
	return patchItems
}

// buildTilePatchItem constructs a DashboardTilePatchItem from a declared tile and a known API tile ID.
func buildTilePatchItem(ctx context.Context, tileID int64, declared TileTFModel) httpclient.DashboardTilePatchItem {
	item := httpclient.DashboardTilePatchItem{
		ID: tileID,
	}

	if !declared.LayoutsJSON.IsNull() && !declared.LayoutsJSON.IsUnknown() {
		var layouts map[string]interface{}
		if err := json.Unmarshal([]byte(declared.LayoutsJSON.ValueString()), &layouts); err == nil {
			item.Layouts = &layouts
		} else {
			tflog.Warn(ctx, "Failed to unmarshal layouts_json; leaving layouts unset", map[string]any{
				"error":        err.Error(),
				"layouts_json": declared.LayoutsJSON.ValueString(),
			})
		}
	}

	// Always send color in the PATCH so that removing color from config
	// actually clears it on the API side. The PostHog model uses
	// CharField(null=True, blank=True), so empty string is accepted and
	// resets the tile to the default (no color).
	if !declared.Color.IsUnknown() {
		c := declared.Color.ValueString() // "" when null
		item.Color = &c
	}

	// Omitted config is authoritative: reset to the API default by sending null.
	if !declared.ShowDescription.IsUnknown() {
		if declared.ShowDescription.IsNull() {
			item.ClearShowDescription = true
		} else {
			showDescription := declared.ShowDescription.ValueBool()
			item.ShowDescription = &showDescription
		}
	}

	if declared.IsTextTile() {
		item.Text = &httpclient.DashboardTileTextPatch{Body: declared.TextBody.ValueString()}
	}

	return item
}

// buildNewTextTilePatchItem constructs a patch item for a new text tile.
// ID=0 is omitted from the JSON payload by omitempty, which tells the API to create a new tile.
func buildNewTextTilePatchItem(ctx context.Context, declared TileTFModel) httpclient.DashboardTilePatchItem {
	return buildTilePatchItem(ctx, 0, declared)
}

// apiTileToTFModel converts a single API tile to the Terraform model representation.
func apiTileToTFModel(t httpclient.DashboardTile) TileTFModel {
	tile := TileTFModel{
		TileID: types.Int64Value(t.ID),
	}

	if t.Insight != nil {
		tile.InsightID = types.Int64Value(t.Insight.ID)
		tile.TextBody = types.StringNull()
	} else if t.Text != nil {
		tile.InsightID = types.Int64Null()
		tile.TextBody = types.StringValue(t.Text.Body)
	}

	if t.Color != nil && *t.Color != "" {
		tile.Color = types.StringValue(*t.Color)
	} else {
		tile.Color = types.StringNull()
	}

	tile.ShowDescription = core.PtrToBool(t.ShowDescription)

	if len(t.Layouts) > 0 {
		bytes, err := json.Marshal(t.Layouts)
		if err == nil {
			tile.LayoutsJSON = jsontypes.NewNormalizedValue(string(bytes))
		} else {
			tile.LayoutsJSON = jsontypes.NewNormalizedNull()
		}
	} else {
		tile.LayoutsJSON = jsontypes.NewNormalizedNull()
	}

	return tile
}

// mapTilesToState maps API tiles to Terraform state, preserving config order to avoid
// phantom diffs from ListNestedAttribute positional comparison.
func mapTilesToState(apiTiles []httpclient.DashboardTile, configTiles []TileTFModel) ([]TileTFModel, diag.Diagnostics) {
	if configTiles == nil {
		// Import case: return all tiles in API order.
		result := make([]TileTFModel, 0, len(apiTiles))
		for _, t := range apiTiles {
			result = append(result, apiTileToTFModel(t))
		}
		return result, nil
	}

	matches, _ := matchDeclaredTilesToAPI(configTiles, apiTiles)

	var result []TileTFModel
	for i := range configTiles {
		if matches[i] != nil {
			result = append(result, apiTileToTFModel(*matches[i]))
		}
		// Tile not found in API: omit from state (will trigger recreation on next plan).
	}

	return result, nil
}

// reAssociateInsight ensures an insight is associated with the given dashboard.
// It GETs the current insight, checks if the dashboard is already in its list,
// and if not, appends the dashboard ID and PATCHes the insight.
// This preserves all existing dashboard associations.
func reAssociateInsight(ctx context.Context, client *httpclient.PosthogClient, projectID string, dashboardID int64, insightID int64) error {
	insightIDStr := strconv.FormatInt(insightID, 10)

	insight, _, err := client.GetInsight(ctx, projectID, insightIDStr)
	if err != nil {
		return fmt.Errorf("getting insight %s: %w", insightIDStr, err)
	}

	// Check if already associated with this dashboard.
	for _, d := range insight.Dashboards {
		if d == int32(dashboardID) {
			return nil // Already on this dashboard.
		}
	}

	dashboards := append(insight.Dashboards, int32(dashboardID))
	_, _, err = client.UpdateInsight(ctx, projectID, insightIDStr, httpclient.InsightRequest{
		Dashboards: dashboards,
	})
	return err
}
