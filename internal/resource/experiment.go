package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

// Experiment lifecycle states as reported by the API's derived `status` field and used in the
// `status.state` config attribute. `exposure_frozen` is a server state that is out of scope
// for v1; it is passed through on read but has no forward transition.
const (
	stateDraft   = "draft"
	stateRunning = "running"
	statePaused  = "paused"
	stateStopped = "stopped"
)

func NewExperiment() resource.Resource {
	return core.NewGenericResource[ExperimentTFModel, experimentAPIRequest, httpclient.Experiment](
		ExperimentOps{},
		core.ProjectScopedImportParser[ExperimentTFModel](),
	)
}

// ExperimentTFModel is the Terraform state model for a PostHog experiment.
type ExperimentTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	Name                    types.String             `tfsdk:"name"`
	Description             types.String             `tfsdk:"description"`
	FeatureFlagKey          types.String             `tfsdk:"feature_flag_key"`
	Variant                 []ExperimentVariantModel `tfsdk:"variant"`
	Metrics                 jsontypes.Normalized     `tfsdk:"metrics"`
	MetricsSecondary        jsontypes.Normalized     `tfsdk:"metrics_secondary"`
	ExposureCriteria        jsontypes.Normalized     `tfsdk:"exposure_criteria"`
	HoldoutID               types.Int64              `tfsdk:"holdout_id"`
	AllowUnknownEvents      types.Bool               `tfsdk:"allow_unknown_events"`
	UpdateFeatureFlagParams types.Bool               `tfsdk:"update_feature_flag_params"`
	Status                  *ExperimentStatusModel   `tfsdk:"status"`
}

// ExperimentVariantModel is one entry of the multivariate split on the backing feature flag.
type ExperimentVariantModel struct {
	Key               types.String `tfsdk:"key"`
	Name              types.String `tfsdk:"name"`
	RolloutPercentage types.Int64  `tfsdk:"rollout_percentage"`
}

// ExperimentStatusModel is the `status` lifecycle block. `state` is the desired lifecycle
// state; `stopped` carries the metadata that only the stopped state uses.
type ExperimentStatusModel struct {
	State   types.String            `tfsdk:"state"`
	Stopped *ExperimentStoppedModel `tfsdk:"stopped"`
}

// ExperimentStoppedModel is the metadata for a stopped experiment. `ship_variant`,
// `release_to_everyone` and `conclusion_comment` are config-only (the API does not echo them
// back off the experiment object); `conclusion` is readable.
type ExperimentStoppedModel struct {
	ShipVariant       types.String `tfsdk:"ship_variant"`
	ReleaseToEveryone types.Bool   `tfsdk:"release_to_everyone"`
	Conclusion        types.String `tfsdk:"conclusion"`
	ConclusionComment types.String `tfsdk:"conclusion_comment"`
}

// experimentAPIRequest is the resource-layer request carried between BuildCreate/UpdateRequest
// and Ops.Create/Update. It bundles the REST body with the encoded lifecycle transition so the
// Ops methods can PATCH/POST the definition and then run the lifecycle sub-actions in order.
type experimentAPIRequest struct {
	body       httpclient.ExperimentRequest
	transition statusTransition
}

// statusTransition is the ordered list of lifecycle sub-actions to run after the definition
// write, plus the payload the terminal action (end/ship) carries.
type statusTransition struct {
	actions           []string // ordered subset of: launch, pause, resume, end, ship
	conclusion        *string
	conclusionComment *string
	shipVariant       string
	releaseToEveryone bool
}

type ExperimentOps struct{}

func (o ExperimentOps) ResourceName() string {
	return "Experiment"
}

func (o ExperimentOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage PostHog experiments (A/B tests) — definition, variants, metrics, and the " +
			"draft → running → paused → stopped lifecycle. Creating an experiment auto-creates its backing " +
			"feature flag from `feature_flag_key`; that flag is owned by the experiment and should not also be " +
			"managed by a separate `posthog_feature_flag` resource. Business rules (rollout sums, variant counts, " +
			"transition legality, metric schema) are enforced by the PostHog API at apply time, not by the provider.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Experiment ID.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Experiment name.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Experiment description.",
			},
			"feature_flag_key": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Key of the backing feature flag. Must be a new, unused key — the flag is " +
					"created for you and owned by the experiment; pointing at a feature flag that already exists is " +
					"not supported (the API rejects it). Changing this forces a new experiment (a linked flag cannot " +
					"be re-keyed in place).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"holdout_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "ID of an existing holdout group to reference. Holdout management itself is out of scope.",
			},
			"metrics": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
				Computed:   true,
				MarkdownDescription: "Primary metrics as a JSON array. Compared semantically, so key ordering and " +
					"whitespace differences from the PostHog API do not produce a diff. Only fields you declare are " +
					"tracked; server-computed fields (e.g. metric fingerprints) are ignored.",
			},
			"metrics_secondary": schema.StringAttribute{
				CustomType:          jsontypes.NormalizedType{},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Secondary metrics as a JSON array. Same semantic-compare handling as `metrics`.",
			},
			"exposure_criteria": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
				Computed:   true,
				MarkdownDescription: "Exposure criteria as a JSON object (this is where `filterTestAccounts` lives). " +
					"Same semantic-compare handling as `metrics`.",
			},
			"allow_unknown_events": schema.BoolAttribute{
				Optional: true,
				MarkdownDescription: "Opt-in bypass of the API validation that rejects metrics referencing not-yet-ingested " +
					"events. Create-time only and not read back from the API.",
			},
			"update_feature_flag_params": schema.BoolAttribute{
				Optional: true,
				MarkdownDescription: "Set to `true` to allow editing the variant split (`variant` blocks) on a " +
					"launched experiment. The API rejects flag-config edits on a running experiment unless this is set. " +
					"Not read back from the API.",
			},
		},
		Blocks: map[string]schema.Block{
			"variant": schema.ListNestedBlock{
				MarkdownDescription: "Multivariate split for the backing feature flag. Declare at least one; rollout " +
					"percentages must sum to 100 and the API enforces the 2–20 variant range. Variants are " +
					"authoritative from config — they are not re-read from the API after apply, so an out-of-band " +
					"edit (or the 100/0 split left by shipping a winner) is not tracked as drift.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Variant key (e.g. `control`, `test`).",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Human-readable variant name.",
						},
						"rollout_percentage": schema.Int64Attribute{
							Required:            true,
							MarkdownDescription: "Percentage of traffic assigned to this variant (0-100).",
						},
					},
				},
			},
			"status": schema.SingleNestedBlock{
				MarkdownDescription: "Desired lifecycle state (required). Declare this block to drive the experiment " +
					"through `draft` → `running` → `paused` → `stopped`. The provider maps the desired `state` (vs. the " +
					"current state) to the matching launch/pause/resume/end/ship sub-action. Backward transitions have no " +
					"API call and return an error. Note: a transition spanning two sub-actions in one apply (e.g. creating " +
					"directly as `paused` = launch+pause, or `stopped` = launch+end) is not atomic — if the second action " +
					"fails the experiment may be left mid-transition; re-apply to reconcile, or advance one state at a time.",
				Attributes: map[string]schema.Attribute{
					"state": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Desired lifecycle state — one of `draft`, `running`, `paused`, or " +
							"`stopped`. Required (the `status` block and its `state` must be declared explicitly); " +
							"use `draft` for a not-yet-launched experiment. The lifecycle is forward-only.",
					},
				},
				Blocks: map[string]schema.Block{
					"stopped": schema.SingleNestedBlock{
						MarkdownDescription: "Metadata applied when stopping the experiment (read only when " +
							"`state = \"stopped\"`, ignored otherwise).",
						Attributes: map[string]schema.Attribute{
							"ship_variant": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: "Key of the winning variant to ship. Rewrites the flag so this " +
									"variant gets 100% and ends the experiment. Destructive and irreversible via the API. " +
									"Config-only (not read back); re-ships only when this value changes. Clearing it does not " +
									"un-ship.",
							},
							"release_to_everyone": schema.BoolAttribute{
								Optional: true,
								MarkdownDescription: "When shipping, release to everyone (catch-all) instead of " +
									"preserving the flag's release conditions. Defaults to `false`. Config-only.",
							},
							"conclusion": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: "Conclusion recorded when the experiment is stopped (e.g. `won`, `lost`, " +
									"`inconclusive`). Write-once: applied by the stop/ship action and not read back or " +
									"updated afterward (editing it on an already-stopped experiment is a no-op).",
							},
							"conclusion_comment": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "Free-text note recorded alongside the conclusion. Config-only.",
							},
						},
					},
				},
			},
		},
	}
}

// applyCommonExperimentFields maps the fields shared by create and update onto the request body.
func applyCommonExperimentFields(body *httpclient.ExperimentRequest, model ExperimentTFModel) {
	body.Name = util.StringPtrFromValue(model.Name)
	body.Description = util.StringPtrFromValue(model.Description)
	body.Metrics = rawFromNormalized(model.Metrics)
	body.MetricsSecondary = rawFromNormalized(model.MetricsSecondary)
	body.ExposureCriteria = rawFromNormalized(model.ExposureCriteria)
	body.HoldoutID = util.Int64PtrFromValue(model.HoldoutID)
}

func (o ExperimentOps) BuildCreateRequest(_ context.Context, model ExperimentTFModel) (experimentAPIRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	body := httpclient.ExperimentRequest{}
	applyCommonExperimentFields(&body, model)
	body.FeatureFlagKey = util.StringPtrFromValue(model.FeatureFlagKey)
	body.AllowUnknownEvents = util.BoolPtrFromValue(model.AllowUnknownEvents)
	body.FeatureFlag = buildFeatureFlagConfig(model.Variant)

	// A create always yields a draft; compute the forward transition here so Create and Update
	// both simply run req.transition (a draft has no backward move, so this never errors).
	transition, err := computeTransition(stateDraft, model.Status, "")
	if err != nil {
		diags.AddError("Invalid experiment status transition", err.Error())
		return experimentAPIRequest{body: body}, diags
	}

	return experimentAPIRequest{body: body, transition: transition}, diags
}

func (o ExperimentOps) BuildUpdateRequest(_ context.Context, plan, state ExperimentTFModel) (experimentAPIRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	body := httpclient.ExperimentRequest{}
	applyCommonExperimentFields(&body, plan)

	// feature_flag_key is RequiresReplace, so it never changes during an update — omit it to
	// sidestep the API's "reject explicit flag config if the key already exists" guard.

	// Only send the variant/flag config when it actually changed, so a status-only update to a
	// running experiment doesn't trip the API guard that rejects flag edits without opt-in.
	if variantsChanged(plan.Variant, state.Variant) {
		body.FeatureFlag = buildFeatureFlagConfig(plan.Variant)
		body.UpdateFeatureFlagParams = util.BoolPtrFromValue(plan.UpdateFeatureFlagParams)
	}

	from := stateDraft
	priorShipVariant := ""
	if state.Status != nil {
		from = normalizeState(state.Status.State.ValueString())
		priorShipVariant = shipVariantOf(state.Status)
	}

	transition, err := computeTransition(from, plan.Status, priorShipVariant)
	if err != nil {
		diags.AddError("Invalid experiment status transition", err.Error())
		return experimentAPIRequest{body: body}, diags
	}

	return experimentAPIRequest{body: body, transition: transition}, diags
}

func (o ExperimentOps) MapResponseToModel(_ context.Context, resp httpclient.Experiment, model *ExperimentTFModel) diag.Diagnostics {
	model.ID = types.Int64Value(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.Description = core.PtrToStringNullIfEmptyTrimmed(resp.Description)

	// feature_flag_key is Required (non-computed); only overwrite when the API returns it so we
	// never clobber the configured value with an empty string.
	if resp.FeatureFlagKey != "" {
		model.FeatureFlagKey = types.StringValue(resp.FeatureFlagKey)
	}

	model.HoldoutID = util.PtrToInt64(resp.HoldoutID)

	// Normalize the JSON blobs so reordered keys and server-computed fields don't perpetual-diff.
	model.Metrics = normalizeRawForState(resp.Metrics, model.Metrics)
	model.MetricsSecondary = normalizeRawForState(resp.MetricsSecondary, model.MetricsSecondary)
	model.ExposureCriteria = normalizeRawForState(resp.ExposureCriteria, model.ExposureCriteria)

	// Variants are authoritative from config: only populate them from the API when the model has
	// none (i.e. import). Keeping declared variants stable in state avoids an inconsistent-result
	// error after ship_variant rewrites the flag distribution (winner→100%), and after PostHog
	// auto-creates a default split. Out-of-band variant edits are therefore not tracked as drift.
	if len(model.Variant) == 0 {
		if variants := readVariants(resp.Parameters); variants != nil {
			model.Variant = variants
		}
	}

	// Status: map the server-derived state string. Create the block on import (empty model).
	if model.Status == nil {
		model.Status = &ExperimentStatusModel{}
	}
	model.Status.State = types.StringValue(normalizeState(resp.Status))

	// The whole stopped block, plus allow_unknown_events and update_feature_flag_params, are
	// config-only: left as passed in, never read back. conclusion/conclusion_comment are sent at
	// stop time by the end/ship action; reading conclusion back would perpetual-diff a post-stop
	// edit (there is no PATCH path for such an edit, so it is a documented no-op).

	return nil
}

func (o ExperimentOps) Create(ctx context.Context, client httpclient.PosthogClient, model ExperimentTFModel, req experimentAPIRequest) (httpclient.Experiment, error) {
	projectID := model.GetEffectiveProjectID()

	exp, err := client.CreateExperiment(ctx, projectID, req.body)
	if err != nil {
		return exp, err
	}

	// A create always yields a draft; reconcile forward to the desired status. The lifecycle
	// sub-actions return the full, authoritative experiment (same serializer as GET), so the
	// last one's response is the final state — no extra fetch needed.
	if len(req.transition.actions) == 0 {
		return exp, nil
	}
	id := strconv.FormatInt(exp.ID, 10)
	final, _, err := runTransition(ctx, client, projectID, id, req.transition)
	if err != nil {
		return exp, err
	}
	return final, nil
}

func (o ExperimentOps) Read(ctx context.Context, client httpclient.PosthogClient, model ExperimentTFModel) (httpclient.Experiment, httpclient.HTTPStatusCode, error) {
	exp, code, err := client.GetExperiment(ctx, model.GetEffectiveProjectID(), model.GetID())
	if err != nil {
		return exp, code, err
	}
	// A soft-deleted experiment is still returned by GET (200) with deleted=true. Surface it as
	// not-found so the generic Read removes it from state and a subsequent plan recreates it,
	// rather than silently tracking a deleted resource as present.
	if exp.Deleted != nil && *exp.Deleted {
		return exp, http.StatusNotFound, fmt.Errorf("experiment %d is soft-deleted", exp.ID)
	}
	return exp, code, nil
}

func (o ExperimentOps) Update(ctx context.Context, client httpclient.PosthogClient, model ExperimentTFModel, req experimentAPIRequest) (httpclient.Experiment, httpclient.HTTPStatusCode, error) {
	projectID := model.GetEffectiveProjectID()
	id := model.GetID()

	exp, code, err := client.UpdateExperiment(ctx, projectID, id, req.body)
	if err != nil {
		return exp, code, err
	}

	if len(req.transition.actions) == 0 {
		return exp, code, nil
	}

	// The lifecycle sub-actions return the full, authoritative experiment, so use the last
	// action's response directly rather than issuing a redundant GET.
	final, c, err := runTransition(ctx, client, projectID, id, req.transition)
	if err != nil {
		return exp, c, err
	}
	return final, c, nil
}

func (o ExperimentOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ExperimentTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteExperiment(ctx, model.GetEffectiveProjectID(), model.GetID())
}

// ModifyResourcePlan requires the `variant` and `status` blocks to be declared explicitly.
// Both are non-computed blocks that MapResponseToModel always populates from the server (variants
// default to a control/test split; status derives from the lifecycle). Omitting either leaves it
// null in the plan, so the provider materializing it after apply trips Terraform's
// "inconsistent result / report a bug" error. These plan-time checks give a clear message
// instead. (Skipped on destroy, where the plan is null.)
func (o ExperimentOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan ExperimentTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if len(plan.Variant) == 0 {
		resp.Diagnostics.AddError(
			"Missing experiment variants",
			"At least one `variant` block is required. A PostHog experiment needs an explicit "+
				"multivariate split; the API enforces the 2–20 variant range.",
		)
	}
	if plan.Status == nil || plan.Status.State.IsNull() || plan.Status.State.IsUnknown() || plan.Status.State.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Missing experiment status",
			`A `+"`status`"+` block with an explicit `+"`state`"+` is required — declare the desired `+
				`lifecycle state, e.g. `+"`status { state = \"draft\" }`"+`.`,
		)
	}
}

// computeTransition maps a from-state and desired status to the ordered lifecycle sub-actions.
// This is action dispatch, not validation: a transition with no forward API call (any backward
// move) returns a structural error. priorShipVariant supports ship-idempotency on an
// already-stopped experiment (ship again only when the variant changes).
func computeTransition(from string, to *ExperimentStatusModel, priorShipVariant string) (statusTransition, error) {
	from = normalizeState(from)
	if to == nil {
		return statusTransition{}, nil
	}
	toState := normalizeState(to.State.ValueString())

	if from == toState {
		if toState == stateStopped {
			if ship := shipVariantOf(to); ship != "" && ship != priorShipVariant {
				t := terminalTransition(to)
				return t, nil
			}
		}
		return statusTransition{}, nil
	}

	switch from {
	case stateDraft:
		switch toState {
		case stateRunning:
			return statusTransition{actions: []string{"launch"}}, nil
		case statePaused:
			return statusTransition{actions: []string{"launch", "pause"}}, nil
		case stateStopped:
			t := terminalTransition(to)
			t.actions = append([]string{"launch"}, t.actions...)
			return t, nil
		}
	case stateRunning:
		switch toState {
		case statePaused:
			return statusTransition{actions: []string{"pause"}}, nil
		case stateStopped:
			return terminalTransition(to), nil
		}
	case statePaused:
		switch toState {
		case stateRunning:
			return statusTransition{actions: []string{"resume"}}, nil
		case stateStopped:
			return terminalTransition(to), nil
		}
	}

	return statusTransition{}, fmt.Errorf(
		"no transition available from %q to %q: the experiment lifecycle is forward-only "+
			"(draft → running → paused → stopped, with pause/resume toggling running↔paused). "+
			"To restart from an earlier state, delete and recreate the experiment",
		from, toState,
	)
}

// terminalTransition builds the stop action: ship when a ship_variant is set, else end. Both
// carry the conclusion/comment payload.
func terminalTransition(to *ExperimentStatusModel) statusTransition {
	t := statusTransition{}
	if to != nil && to.Stopped != nil {
		t.conclusion = util.StringPtrFromValue(to.Stopped.Conclusion)
		t.conclusionComment = util.StringPtrFromValue(to.Stopped.ConclusionComment)
	}
	if ship := shipVariantOf(to); ship != "" {
		t.shipVariant = ship
		t.releaseToEveryone = to.Stopped.ReleaseToEveryone.ValueBool()
		t.actions = []string{"ship"}
		return t
	}
	t.actions = []string{"end"}
	return t
}

func shipVariantOf(to *ExperimentStatusModel) string {
	if to == nil || to.Stopped == nil {
		return ""
	}
	return util.ValueStringOrEmpty(to.Stopped.ShipVariant)
}

func normalizeState(s string) string {
	if s == "" {
		return stateDraft
	}
	return s
}

// runTransition runs the encoded sub-actions in order and returns the experiment from the last
// action performed.
func runTransition(ctx context.Context, client httpclient.PosthogClient, projectID, id string, t statusTransition) (httpclient.Experiment, httpclient.HTTPStatusCode, error) {
	var (
		last httpclient.Experiment
		code httpclient.HTTPStatusCode = http.StatusOK
		err  error
	)
	for _, action := range t.actions {
		switch action {
		case "launch":
			last, code, err = client.LaunchExperiment(ctx, projectID, id)
		case "pause":
			last, code, err = client.PauseExperiment(ctx, projectID, id)
		case "resume":
			last, code, err = client.ResumeExperiment(ctx, projectID, id)
		case "end":
			last, code, err = client.EndExperiment(ctx, projectID, id, httpclient.ExperimentEndRequest{
				Conclusion:        t.conclusion,
				ConclusionComment: t.conclusionComment,
			})
		case "ship":
			last, code, err = client.ShipVariant(ctx, projectID, id, httpclient.ExperimentShipVariantRequest{
				VariantKey:        t.shipVariant,
				ReleaseToEveryone: t.releaseToEveryone,
				Conclusion:        t.conclusion,
				ConclusionComment: t.conclusionComment,
			})
		default:
			// Guards against a future action-name mismatch between computeTransition/
			// terminalTransition and this dispatcher silently no-opping.
			return last, code, fmt.Errorf("unknown lifecycle action %q", action)
		}
		if err != nil {
			// Name the failed action so a partial multi-step transition (e.g. launch succeeded,
			// pause failed) is diagnosable from the error alone.
			return last, code, fmt.Errorf("%q action failed: %w", action, err)
		}
	}
	return last, code, err
}

// variantEntries converts typed variant blocks into the API's variant maps ({key,
// rollout_percentage, and name when set}). Shared by the flag-config writer and the
// change-detection signature so the two can never drift out of sync.
func variantEntries(variants []ExperimentVariantModel) []map[string]interface{} {
	entries := make([]map[string]interface{}, 0, len(variants))
	for _, v := range variants {
		entry := map[string]interface{}{
			"key":                v.Key.ValueString(),
			"rollout_percentage": v.RolloutPercentage.ValueInt64(),
		}
		if !v.Name.IsNull() && !v.Name.IsUnknown() {
			entry["name"] = v.Name.ValueString()
		}
		entries = append(entries, entry)
	}
	return entries
}

// buildFeatureFlagConfig serializes typed variant blocks into the flag write shape
// {filters:{multivariate:{variants:[...]}}}. Returns nil when no variants are configured.
func buildFeatureFlagConfig(variants []ExperimentVariantModel) json.RawMessage {
	if len(variants) == 0 {
		return nil
	}
	config := map[string]interface{}{
		"filters": map[string]interface{}{
			"multivariate": map[string]interface{}{
				"variants": variantEntries(variants),
			},
		},
	}
	raw, err := json.Marshal(config)
	if err != nil {
		return nil
	}
	return raw
}

// readVariants reads the variant split back from parameters.feature_flag_variants.
func readVariants(parameters json.RawMessage) []ExperimentVariantModel {
	if len(parameters) == 0 {
		return nil
	}
	var params struct {
		FeatureFlagVariants []struct {
			Key               string   `json:"key"`
			Name              *string  `json:"name"`
			RolloutPercentage *float64 `json:"rollout_percentage"`
		} `json:"feature_flag_variants"`
	}
	if err := json.Unmarshal(parameters, &params); err != nil {
		return nil
	}
	if len(params.FeatureFlagVariants) == 0 {
		return nil
	}
	variants := make([]ExperimentVariantModel, 0, len(params.FeatureFlagVariants))
	for _, v := range params.FeatureFlagVariants {
		entry := ExperimentVariantModel{Key: types.StringValue(v.Key)}
		if v.Name != nil && *v.Name != "" {
			entry.Name = types.StringValue(*v.Name)
		} else {
			entry.Name = types.StringNull()
		}
		if v.RolloutPercentage != nil {
			entry.RolloutPercentage = types.Int64Value(int64(*v.RolloutPercentage))
		} else {
			entry.RolloutPercentage = types.Int64Null()
		}
		variants = append(variants, entry)
	}
	return variants
}

// variantsChanged reports whether the desired variant split differs from the prior state.
func variantsChanged(plan, state []ExperimentVariantModel) bool {
	return variantsSignature(plan) != variantsSignature(state)
}

func variantsSignature(variants []ExperimentVariantModel) string {
	raw, _ := json.Marshal(variantEntries(variants))
	return string(raw)
}

// rawFromNormalized converts a jsontypes.Normalized field into a JSON request body value,
// returning nil (omitted from the request) when null/unknown/empty.
func rawFromNormalized(n jsontypes.Normalized) json.RawMessage {
	if n.IsNull() || n.IsUnknown() {
		return nil
	}
	trimmed := strings.TrimSpace(n.ValueString())
	if trimmed == "" {
		return nil
	}
	return json.RawMessage(trimmed)
}

// normalizeRawForState normalizes an API JSON blob against the configured value so that
// reordered keys and server-computed fields do not surface as a perpetual diff. Reuses the
// shared whitelist normalizer (keep only user-declared fields), the same family used by
// survey/insight/hog_function.
func normalizeRawForState(raw json.RawMessage, current jsontypes.Normalized) jsontypes.Normalized {
	if len(raw) == 0 || string(raw) == "null" {
		return jsontypes.NewNormalizedNull()
	}
	var apiData interface{}
	if err := json.Unmarshal(raw, &apiData); err != nil {
		return jsontypes.NewNormalizedNull()
	}
	userJSON := ""
	if !current.IsNull() && !current.IsUnknown() {
		userJSON = current.ValueString()
	}
	normalized, err := normalizeJSONForState(apiData, userJSON)
	if err != nil || normalized == "" {
		return jsontypes.NewNormalizedNull()
	}
	return jsontypes.NewNormalizedValue(normalized)
}
