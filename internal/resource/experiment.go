package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

// Experiment lifecycle states. draft/running/paused/stopped are settable via `status.state`;
// exposure_frozen is a server-only state (enrollment frozen) that is out of scope for v1 — it is
// passed through on read but has no forward transition, so it can't be driven from config.
const (
	stateDraft          = "draft"
	stateRunning        = "running"
	statePaused         = "paused"
	stateStopped        = "stopped"
	stateExposureFrozen = "exposure_frozen"
)

// experimentStates are the values accepted for status.state (exposure_frozen is server-only).
var experimentStates = []string{stateDraft, stateRunning, statePaused, stateStopped}

// experimentConclusions mirrors the PostHog API's conclusion enum.
var experimentConclusions = []string{"won", "lost", "inconclusive", "stopped_early", "invalid"}

// lifecycleAction is one API sub-action in a status transition; the typed constants keep
// computeTransition (which builds them) and runTransition (which dispatches them) from drifting.
type lifecycleAction string

const (
	actionLaunch lifecycleAction = "launch"
	actionPause  lifecycleAction = "pause"
	actionResume lifecycleAction = "resume"
	actionEnd    lifecycleAction = "end"
	actionShip   lifecycleAction = "ship"
)

func NewExperiment() resource.Resource {
	return core.NewGenericResource[ExperimentTFModel, experimentAPIRequest, httpclient.Experiment](
		ExperimentOps{},
		core.ProjectScopedImportParser[ExperimentTFModel](),
	)
}

// ExperimentTFModel is the Terraform state model for a PostHog experiment. The backing feature
// flag is a separate resource referenced by feature_flag_key — the experiment does not own the
// variant split.
type ExperimentTFModel struct {
	core.BaseInt64Identifiable
	core.BaseProjectID
	Name               types.String           `tfsdk:"name"`
	Description        types.String           `tfsdk:"description"`
	FeatureFlagKey     types.String           `tfsdk:"feature_flag_key"`
	Metrics            jsontypes.Normalized   `tfsdk:"metrics"`
	MetricsSecondary   jsontypes.Normalized   `tfsdk:"metrics_secondary"`
	ExposureCriteria   jsontypes.Normalized   `tfsdk:"exposure_criteria"`
	HoldoutID          types.Int64            `tfsdk:"holdout_id"`
	AllowUnknownEvents types.Bool             `tfsdk:"allow_unknown_events"`
	Status             *ExperimentStatusModel `tfsdk:"status"`
}

// ExperimentStatusModel is the `status` lifecycle block. `state` is the desired lifecycle
// state; `stopped` carries the metadata that only the stopped state uses.
type ExperimentStatusModel struct {
	State   types.String            `tfsdk:"state"`
	Stopped *ExperimentStoppedModel `tfsdk:"stopped"`
}

// ExperimentStoppedModel is the metadata applied when stopping an experiment. conclusion and
// conclusion_comment are fully managed (sent on stop/update and read back). ship_variant and
// release_to_everyone are config-only ship instructions the API does not echo back.
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
	actions           []lifecycleAction
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
		MarkdownDescription: "Manage PostHog experiments (A/B tests) — definition, metrics, and the " +
			"draft → running → paused → stopped lifecycle. The backing feature flag is a separate " +
			"`posthog_feature_flag` resource referenced by `feature_flag_key`; the experiment attaches metrics " +
			"and drives the lifecycle (including shipping a winning variant on stop). Business rules " +
			"(transition legality, metric schema) are enforced by the PostHog API at apply time, not by the provider.",
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
				MarkdownDescription: "Key of the multivariate feature flag this experiment runs on. Reference a " +
					"`posthog_feature_flag` resource — e.g. `feature_flag_key = posthog_feature_flag.<name>.key`. The flag " +
					"must already exist and be multivariate (2–20 variants, one keyed `control`). Changing this forces a " +
					"new experiment (a linked flag cannot be re-keyed in place).",
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
					"events. Sent on every write (so a draft with such a metric can still be launched/edited); config-only, " +
					"not read back from the API.",
			},
		},
		Blocks: map[string]schema.Block{
			"status": schema.SingleNestedBlock{
				MarkdownDescription: "Desired lifecycle state. The `status` block and its `state` are **required** " +
					"(the block is schema-optional only because Terraform has no required-block modifier; a plan without " +
					"it errors). Declare it to drive the experiment through `draft` → `running` → `paused` → `stopped`; the " +
					"provider maps the desired `state` (vs. the current state) to the matching launch/pause/resume/end/ship " +
					"sub-action. The lifecycle is forward-only — backward transitions error. Note: a transition spanning two " +
					"sub-actions in one apply (e.g. creating directly as `paused` = launch+pause, or `stopped` = launch+end) " +
					"is not atomic — if the second action fails the experiment is left mid-transition; the error names its " +
					"live state so you can reconcile. The server-only `exposure_frozen` state cannot be managed here; " +
					"resume or end such an experiment in the PostHog UI first.",
				Attributes: map[string]schema.Attribute{
					"state": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Desired lifecycle state — one of `draft`, `running`, `paused`, or " +
							"`stopped` (required; use `draft` for a not-yet-launched experiment). The lifecycle is forward-only.",
						Validators: []validator.String{
							stringvalidator.OneOf(experimentStates...),
						},
					},
				},
				Blocks: map[string]schema.Block{
					"stopped": schema.SingleNestedBlock{
						MarkdownDescription: "Metadata applied when stopping the experiment (read only when " +
							"`state = \"stopped\"`, ignored otherwise).",
						Attributes: map[string]schema.Attribute{
							"ship_variant": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: "Key of the winning variant to ship. Rewrites the linked flag's " +
									"distribution so this variant gets 100% and ends the experiment. If the flag is managed by a " +
									"`posthog_feature_flag` resource, set `lifecycle { ignore_changes = [filters] }` on it so it does " +
									"not revert the shipped distribution. Config-only; re-ships only when this value changes; clearing " +
									"it does not un-ship.",
							},
							"release_to_everyone": schema.BoolAttribute{
								Optional: true,
								MarkdownDescription: "When shipping, prepend a catch-all release condition (roll out to " +
									"everyone) instead of only flipping the variant distribution. Defaults to `false`. Config-only.",
							},
							"conclusion": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: "Conclusion recorded when the experiment is stopped — one of `won`, " +
									"`lost`, `inconclusive`, `stopped_early`, `invalid`. Fully managed: applied at stop time and " +
									"editable afterwards (an edit on an already-stopped experiment is PATCHed and read back).",
								Validators: []validator.String{
									stringvalidator.OneOf(experimentConclusions...),
								},
							},
							"conclusion_comment": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: "Free-text note recorded alongside the conclusion. Fully managed " +
									"(sent on stop/update and read back).",
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
	// The API re-validates metrics on every write, so the unknown-event bypass must ride along on
	// updates too — otherwise an experiment whose metric references a not-yet-ingested event could be
	// created but never edited or launched (the definition PATCH would re-trip the validation).
	body.AllowUnknownEvents = util.BoolPtrFromValue(model.AllowUnknownEvents)
}

func (o ExperimentOps) BuildCreateRequest(_ context.Context, model ExperimentTFModel) (experimentAPIRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	body := httpclient.ExperimentRequest{}
	applyCommonExperimentFields(&body, model)
	body.FeatureFlagKey = util.StringPtrFromValue(model.FeatureFlagKey)
	// No feature_flag config is sent: the experiment links the existing flag named by
	// feature_flag_key as-is (the flag is owned by a posthog_feature_flag resource).

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

	// feature_flag_key is RequiresReplace, so it never changes during an update; the linked flag's
	// config is managed by its own posthog_feature_flag resource, not from here.

	from := stateDraft
	priorShipVariant := ""
	if state.Status != nil {
		from = normalizeState(state.Status.State.ValueString())
		priorShipVariant = shipVariantOf(state.Status)
	}

	// When the experiment is already stopped there is no lifecycle action to carry a conclusion
	// edit, so send conclusion/comment in the PATCH body (they are applied and read back — two-way).
	if from == stateStopped && plan.Status != nil && plan.Status.Stopped != nil {
		body.Conclusion = util.StringPtrFromValue(plan.Status.Stopped.Conclusion)
		body.ConclusionComment = util.StringPtrFromValue(plan.Status.Stopped.ConclusionComment)
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

	// Status: map the server-derived state string. Create the block on import (empty model).
	if model.Status == nil {
		model.Status = &ExperimentStatusModel{}
	}
	model.Status.State = types.StringValue(normalizeState(resp.Status))

	// conclusion/conclusion_comment are two-way (read back so out-of-band and post-stop edits
	// round-trip). ship_variant/release_to_everyone/allow_unknown_events are config-only, so unmapped.
	if model.Status.Stopped != nil {
		model.Status.Stopped.Conclusion = core.PtrToStringNullIfEmptyTrimmed(resp.Conclusion)
		model.Status.Stopped.ConclusionComment = core.PtrToStringNullIfEmptyTrimmed(resp.ConclusionComment)
	}

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
	final, _, err := runTransition(ctx, &client, projectID, id, req.transition)
	if err != nil {
		// The experiment was created and partially transitioned, but Terraform won't record it in
		// state (create failed). Name it and its live state so the user can import or delete it
		// rather than have a re-apply create a second, orphaned live experiment.
		return exp, fmt.Errorf("experiment %s was created (now %s) but a lifecycle action failed; import it "+
			"(terraform import <resource> %s/%s) or delete it in the PostHog UI before retrying: %w",
			id, liveState(ctx, client, projectID, id), projectID, id, err)
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
	final, c, err := runTransition(ctx, &client, projectID, id, req.transition)
	if err != nil {
		// Multi-step transitions aren't atomic; name the live state so a mid-sequence failure is
		// recoverable (re-apply reconciles from the real state).
		return exp, c, fmt.Errorf("experiment %s is now %s after a partial transition: %w",
			id, liveState(ctx, client, projectID, id), err)
	}
	return final, c, nil
}

// liveState fetches the experiment's current server state for error messages; returns a placeholder
// if the fetch fails so diagnostics never mask the original error.
func liveState(ctx context.Context, client httpclient.PosthogClient, projectID, id string) string {
	got, _, err := client.GetExperiment(ctx, projectID, id)
	if err != nil {
		return "in an unknown state"
	}
	return fmt.Sprintf("in state %q", normalizeState(got.Status))
}

func (o ExperimentOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ExperimentTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteExperiment(ctx, model.GetEffectiveProjectID(), model.GetID())
}

// ModifyResourcePlan requires the `status` block to be declared explicitly. It is a non-computed
// block that MapResponseToModel always populates from the server (it derives from the lifecycle),
// so omitting it would leave it null in the plan and trip Terraform's "inconsistent result /
// report a bug" error; this check gives a clear message instead. (Skipped on destroy, where the
// plan is null.)
func (o ExperimentOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan ExperimentTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
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
			return statusTransition{actions: []lifecycleAction{actionLaunch}}, nil
		case statePaused:
			return statusTransition{actions: []lifecycleAction{actionLaunch, actionPause}}, nil
		case stateStopped:
			t := terminalTransition(to)
			t.actions = append([]lifecycleAction{actionLaunch}, t.actions...)
			return t, nil
		}
	case stateRunning:
		switch toState {
		case statePaused:
			return statusTransition{actions: []lifecycleAction{actionPause}}, nil
		case stateStopped:
			return terminalTransition(to), nil
		}
	case statePaused:
		switch toState {
		case stateRunning:
			return statusTransition{actions: []lifecycleAction{actionResume}}, nil
		case stateStopped:
			return terminalTransition(to), nil
		}
	case stateExposureFrozen:
		return statusTransition{}, fmt.Errorf(
			"experiment is in the server-only state %q (enrollment frozen), which Terraform cannot drive; "+
				"resume or end it in the PostHog UI, then re-apply", stateExposureFrozen)
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
		t.actions = []lifecycleAction{actionShip}
		return t
	}
	t.actions = []lifecycleAction{actionEnd}
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

// experimentLifecycleClient is the subset of the PostHog client that runTransition drives; taking
// an interface lets the transition sequencer be unit-tested with a stub.
type experimentLifecycleClient interface {
	LaunchExperiment(ctx context.Context, projectID, id string) (httpclient.Experiment, httpclient.HTTPStatusCode, error)
	PauseExperiment(ctx context.Context, projectID, id string) (httpclient.Experiment, httpclient.HTTPStatusCode, error)
	ResumeExperiment(ctx context.Context, projectID, id string) (httpclient.Experiment, httpclient.HTTPStatusCode, error)
	EndExperiment(ctx context.Context, projectID, id string, input httpclient.ExperimentEndRequest) (httpclient.Experiment, httpclient.HTTPStatusCode, error)
	ShipVariant(ctx context.Context, projectID, id string, input httpclient.ExperimentShipVariantRequest) (httpclient.Experiment, httpclient.HTTPStatusCode, error)
}

// runTransition runs the encoded sub-actions in order and returns the experiment from the last
// action performed.
func runTransition(ctx context.Context, client experimentLifecycleClient, projectID, id string, t statusTransition) (httpclient.Experiment, httpclient.HTTPStatusCode, error) {
	var (
		last httpclient.Experiment
		code httpclient.HTTPStatusCode = http.StatusOK
		err  error
	)
	for _, action := range t.actions {
		switch action {
		case actionLaunch:
			last, code, err = client.LaunchExperiment(ctx, projectID, id)
		case actionPause:
			last, code, err = client.PauseExperiment(ctx, projectID, id)
		case actionResume:
			last, code, err = client.ResumeExperiment(ctx, projectID, id)
		case actionEnd:
			last, code, err = client.EndExperiment(ctx, projectID, id, httpclient.ExperimentEndRequest{
				Conclusion:        t.conclusion,
				ConclusionComment: t.conclusionComment,
			})
		case actionShip:
			last, code, err = client.ShipVariant(ctx, projectID, id, httpclient.ExperimentShipVariantRequest{
				VariantKey:        t.shipVariant,
				ReleaseToEveryone: t.releaseToEveryone,
				Conclusion:        t.conclusion,
				ConclusionComment: t.conclusionComment,
			})
		default:
			// Unreachable with the typed constants; defensive guard.
			return last, code, fmt.Errorf("unknown lifecycle action %q", action)
		}
		if err != nil {
			// Name the failed action so a partial multi-step transition (e.g. launch succeeded,
			// pause failed) is diagnosable from the error alone.
			return last, code, fmt.Errorf("%q action failed: %w", string(action), err)
		}
	}
	return last, code, err
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

// normalizeRawForState normalizes an API JSON blob against the configured value so reordered keys
// and server-computed fields do not surface as a perpetual diff (whitelist to user-declared fields).
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
