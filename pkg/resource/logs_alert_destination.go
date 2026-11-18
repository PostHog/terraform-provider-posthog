package resource

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
)

const (
	destinationTypeSlack   = "slack"
	destinationTypeWebhook = "webhook"
	destinationTypeTeams   = "teams"
)

const hogFunctionIDSeparator = ","

func NewLogsAlertDestination() resource.Resource {
	return core.NewGenericResource[LogsAlertDestinationTFModel, httpclient.LogsAlertDestinationRequest, httpclient.LogsAlertDestination](
		LogsAlertDestinationOps{},
		core.LogsAlertDestinationImportParser[LogsAlertDestinationTFModel](),
	)
}

type LogsAlertDestinationTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	AlertID          types.String `tfsdk:"alert_id"`
	Type             types.String `tfsdk:"type"`
	SlackWorkspaceID types.Int64  `tfsdk:"slack_workspace_id"`
	SlackChannelID   types.String `tfsdk:"slack_channel_id"`
	SlackChannelName types.String `tfsdk:"slack_channel_name"`
	WebhookURL       types.String `tfsdk:"webhook_url"`
	HogFunctionIDs   types.Set    `tfsdk:"hog_function_ids"`
}

func (m LogsAlertDestinationTFModel) GetAlertID() string {
	return m.AlertID.ValueString()
}

var _ core.AlertIDSetter = (*LogsAlertDestinationTFModel)(nil)

func (m *LogsAlertDestinationTFModel) SetAlertID(alertID string) {
	m.AlertID = types.StringValue(alertID)
}

type LogsAlertDestinationOps struct{}

func (o LogsAlertDestinationOps) ResourceName() string {
	return "Logs Alert Destination"
}

func (o LogsAlertDestinationOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage where a [log alert](https://posthog.com/docs/logs/alerts) sends its " +
			"notifications. A `posthog_logs_alert` with no destination still evaluates and reports its state, " +
			"but notifies nobody.\n\n" +
			"~> **Every attribute forces replacement.** PostHog has no update endpoint for destinations, so " +
			"changing a channel or a URL destroys the destination and creates a new one. There is a short " +
			"window during the apply where the alert has no destination and would notify nobody if it fired.\n\n" +
			"Each resource represents one user-visible destination, and its `id` is the destination's own " +
			"id, the `destinations[].id` the logs alerts API returns. Terraform creates and deletes the " +
			"destination through the alert's destinations API, and reads it back from the alert itself.\n\n" +
			"-> **The `id` changes once, on the next refresh after upgrading to 1.0.32.** Earlier releases " +
			"tracked a destination by the hog functions PostHog built for it, so their `id` is the " +
			"destination's `hog_function_ids`, sorted and joined by commas. PostHog now gives each " +
			"destination an id of its own and stops building hog functions for new ones. The first refresh " +
			"finds the destination by its old `id` and records the new one. Nothing is replaced, and the plan " +
			"shows no changes. A destination PostHog has not given an id yet keeps the old form until it has " +
			"one. Releases before 1.0.32 cannot create destinations: PostHog refuses the create with a " +
			"`terraform_provider_upgrade_required` error and writes nothing.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "The destination's id, `destinations[].id` in the logs alerts API. A " +
					"destination PostHog has not given an id yet uses its `hog_function_ids`, sorted and joined " +
					"by commas, until it has one.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"alert_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "UUID of the `posthog_logs_alert` this destination notifies for.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Where the notification goes. PostHog supports `slack`, `webhook`, and " +
					"`teams` (Microsoft Teams); see the [alert destinations API]" +
					"(https://posthog.com/docs/api/logs-alerts) for the current list. A `slack` destination " +
					"needs `slack_workspace_id` and `slack_channel_id`; `webhook` and `teams` need " +
					"`webhook_url`. Any other value is sent to PostHog as given, so a destination type added " +
					"to the API works without a new provider release. Terraform cannot check the settings a " +
					"type it does not know requires, so PostHog reports those at apply.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"slack_workspace_id": schema.Int64Attribute{
				Optional: true,
				MarkdownDescription: "ID of the Slack integration to post through, as created by connecting " +
					"Slack to the project in the PostHog UI. Required when `type` is `slack`.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"slack_channel_id": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "Slack channel ID to post into, such as `C0123456789`. Required when " +
					"`type` is `slack`. The channel ID rather than the name, because renaming a channel in " +
					"Slack must not silently repoint the alert.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"slack_channel_name": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "Channel name such as `#alerts`, used only to label the destination in " +
					"the PostHog UI. Optional, and only meaningful when `type` is `slack`.\n\n" +
					"Write-only: PostHog uses it to build the display name and never stores it, so nothing can " +
					"read it back. Terraform keeps whatever you configured and never reports drift on it, and " +
					"an imported destination has it unset. Changing it forces replacement, which is the only " +
					"way to change the display name.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"webhook_url": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				MarkdownDescription: "URL to POST the notification to. Required when `type` is `webhook` or " +
					"`teams`. For `teams`, this is the Microsoft Teams incoming webhook URL. Marked sensitive " +
					"because the secret is in the URL: whoever holds it can post to the channel.\n\n" +
					"PostHog redacts the URL when reading destinations, so Terraform keeps the configured " +
					"value and never compares it against the live one. A destination's URL cannot change " +
					"in place: PostHog has no update endpoint for one, and the Hog Function API refuses " +
					"writes to alert-owned destinations. Changing a URL outside Terraform means deleting " +
					"the destination and adding another, which changes its `id`, so Terraform sees the " +
					"destination as deleted and creates it again. Rotate the URL by changing it here.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"hog_function_ids": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
				MarkdownDescription: "UUIDs of the hog functions PostHog built for this destination, one per " +
					"alert transition. Unset for a destination PostHog delivers without hog functions, which " +
					"includes every destination created from 1.0.32 on.",
			},
		},
	}
}

func (o LogsAlertDestinationOps) BuildCreateRequest(_ context.Context, model LogsAlertDestinationTFModel) (httpclient.LogsAlertDestinationRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	return httpclient.LogsAlertDestinationRequest{
		Type:             model.Type.ValueString(),
		SlackWorkspaceID: util.Int64PtrFromValue(model.SlackWorkspaceID),
		SlackChannelID:   util.StringPtrFromValue(model.SlackChannelID),
		SlackChannelName: util.StringPtrFromValue(model.SlackChannelName),
		WebhookURL:       util.StringPtrFromValue(model.WebhookURL),
	}, diags
}

func (o LogsAlertDestinationOps) BuildUpdateRequest(_ context.Context, _, _ LogsAlertDestinationTFModel) (httpclient.LogsAlertDestinationRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	diags.AddError(
		"Update not supported",
		"Log alert destinations have no update endpoint, and every attribute forces replacement. "+
			"This is a bug if you see this error.",
	)
	return httpclient.LogsAlertDestinationRequest{}, diags
}

type writeOnlyDestinationAttributes struct {
	SlackChannelName types.String
	WebhookURL       types.String
}

func writeOnlyAttributesOf(model LogsAlertDestinationTFModel) writeOnlyDestinationAttributes {
	return writeOnlyDestinationAttributes{
		SlackChannelName: model.SlackChannelName,
		WebhookURL:       model.WebhookURL,
	}
}

func (w writeOnlyDestinationAttributes) restoreTo(model *LogsAlertDestinationTFModel) {
	model.SlackChannelName = w.SlackChannelName
	model.WebhookURL = w.WebhookURL
}

// Where a person fixes a destination Terraform cannot reconcile. Relative, because the
// provider is configured with an API host that may not be the one they browse.
func logsAlertNotificationsPath(projectID, alertID string) string {
	return fmt.Sprintf("/project/%s/logs/alerts/%s?tab=notifications", projectID, alertID)
}

func destinationIncludesConfiguration(resp httpclient.LogsAlertDestination) bool {
	return resp.Type != ""
}

func (o LogsAlertDestinationOps) MapResponseToModel(ctx context.Context, resp httpclient.LogsAlertDestination, model *LogsAlertDestinationTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if !hasRowID(resp) && len(resp.HogFunctionIDs) == 0 {
		diags.AddError(
			"Log alert destination has no id",
			fmt.Sprintf(
				"PostHog returned a destination with neither an id nor hog_function_ids, which Terraform cannot track. "+
					"The destination may still exist on the alert. Open %s in PostHog to see its current "+
					"state, then apply again.",
				logsAlertNotificationsPath(model.GetEffectiveProjectID(), model.GetAlertID()),
			),
		)
		return diags
	}

	writeOnly := writeOnlyAttributesOf(*model)

	model.ID = types.StringValue(logsAlertDestinationID(resp))

	hogFunctionIDs, d := core.TagsToSet(ctx, resp.HogFunctionIDs)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.HogFunctionIDs = hogFunctionIDs

	// The create response returns only the ids, so an empty Type marks a create and we
	// keep the caller's configured values. A populated Type marks a read or import, where we adopt
	// the server's canonical values. The shared MapResponseToModel signature carries no operation
	// flag, so this emptiness check stands in for it. Write-only attributes are always restored
	// from the model below, since PostHog never returns them.
	if destinationIncludesConfiguration(resp) {
		model.Type = types.StringValue(resp.Type)
		model.SlackWorkspaceID = util.PtrToInt64(resp.SlackWorkspaceID)
		model.SlackChannelID = core.PtrToStringNullIfEmptyTrimmed(resp.SlackChannelID)
	}

	writeOnly.restoreTo(model)

	return diags
}

func hasRowID(destination httpclient.LogsAlertDestination) bool {
	return destination.ID != nil && *destination.ID != ""
}

// logsAlertDestinationID is the destination's row id, or, for a destination PostHog has
// not given a row yet, the id releases before row ids used.
func logsAlertDestinationID(destination httpclient.LogsAlertDestination) string {
	if hasRowID(destination) {
		return *destination.ID
	}
	return hogFunctionGroupID(destination.HogFunctionIDs)
}

func hogFunctionGroupID(hogFunctionIDs []string) string {
	sorted := slices.Clone(hogFunctionIDs)
	slices.Sort(sorted)
	return strings.Join(sorted, hogFunctionIDSeparator)
}

func (o LogsAlertDestinationOps) ModifyResourcePlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var plan, config LogsAlertDestinationTFModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(validateLogsAlertDestinationPlan(plan, config)...)
}

type destinationPlanAttributes struct {
	slackWorkspaceID util.EffectiveValue[int64]
	slackChannelID   util.EffectiveValue[string]
	slackChannelName util.EffectiveValue[string]
	webhookURL       util.EffectiveValue[string]
}

func resolveDestinationPlanAttributes(plan, config LogsAlertDestinationTFModel) destinationPlanAttributes {
	return destinationPlanAttributes{
		slackWorkspaceID: util.EffectiveInt64(plan.SlackWorkspaceID, config.SlackWorkspaceID),
		slackChannelID:   util.EffectiveString(plan.SlackChannelID, config.SlackChannelID),
		slackChannelName: util.EffectiveString(plan.SlackChannelName, config.SlackChannelName),
		webhookURL:       util.EffectiveString(plan.WebhookURL, config.WebhookURL),
	}
}

func validateLogsAlertDestinationPlan(plan, config LogsAlertDestinationTFModel) diag.Diagnostics {
	destinationType := util.EffectiveString(plan.Type, config.Type)
	if !destinationType.IsSet() {
		return nil
	}

	attributes := resolveDestinationPlanAttributes(plan, config)
	switch destinationType.Value {
	case destinationTypeSlack:
		return attributes.slackDestinationErrors()
	case destinationTypeWebhook, destinationTypeTeams:
		return attributes.webhookOrTeamsDestinationErrors(destinationType.Value)
	}
	// A type this provider does not know about. Sending it on lets a destination type added
	// to the API work without a release, so PostHog validates its settings instead.
	return nil
}

func (a destinationPlanAttributes) slackDestinationErrors() diag.Diagnostics {
	var diags diag.Diagnostics

	if a.slackWorkspaceID.IsOmitted() {
		diags.AddAttributeError(
			path.Root("slack_workspace_id"),
			"Missing Slack destination settings",
			`type = "slack" requires slack_workspace_id, the ID of the Slack integration connected to `+
				`this project.`,
		)
	}
	if a.slackChannelID.IsOmitted() {
		diags.AddAttributeError(
			path.Root("slack_channel_id"),
			"Missing Slack destination settings",
			`type = "slack" requires slack_channel_id, the ID of the channel to post into, such as `+
				`C0123456789.`,
		)
	}
	if a.webhookURL.IsSet() {
		diags.Append(destinationAttributeDoesNotApply("webhook_url", destinationTypeSlack)...)
	}

	return diags
}

func (a destinationPlanAttributes) webhookOrTeamsDestinationErrors(destinationType string) diag.Diagnostics {
	var diags diag.Diagnostics

	if a.webhookURL.IsOmitted() {
		diags.AddAttributeError(
			path.Root("webhook_url"),
			"Missing destination URL",
			fmt.Sprintf("type = %q requires webhook_url, the URL to POST the notification to.", destinationType),
		)
	}
	if a.slackWorkspaceID.IsSet() {
		diags.Append(destinationAttributeDoesNotApply("slack_workspace_id", destinationType)...)
	}
	if a.slackChannelID.IsSet() {
		diags.Append(destinationAttributeDoesNotApply("slack_channel_id", destinationType)...)
	}
	if a.slackChannelName.IsSet() {
		diags.Append(destinationAttributeDoesNotApply("slack_channel_name", destinationType)...)
	}

	return diags
}

func destinationAttributeDoesNotApply(attribute, destinationType string) diag.Diagnostics {
	var diags diag.Diagnostics

	addressedBy := "slack_workspace_id and slack_channel_id"
	if destinationType != destinationTypeSlack {
		addressedBy = "webhook_url"
	}

	diags.AddAttributeError(
		path.Root(attribute),
		"Attribute does not apply to this destination type",
		fmt.Sprintf("%s cannot be set when type = %q. That destination is addressed by %s.",
			attribute, destinationType, addressedBy),
	)
	return diags
}

func (o LogsAlertDestinationOps) Create(ctx context.Context, client httpclient.PosthogClient, model LogsAlertDestinationTFModel, req httpclient.LogsAlertDestinationRequest) (httpclient.LogsAlertDestination, error) {
	return client.CreateLogsAlertDestination(ctx, model.GetEffectiveProjectID(), model.GetAlertID(), req)
}

func (o LogsAlertDestinationOps) Read(ctx context.Context, client httpclient.PosthogClient, model LogsAlertDestinationTFModel) (httpclient.LogsAlertDestination, httpclient.HTTPStatusCode, error) {
	destinations, status, err := client.ListLogsAlertDestinations(ctx, model.GetEffectiveProjectID(), model.GetAlertID())
	if err != nil {
		return httpclient.LogsAlertDestination{}, status, err
	}

	if destination, ok := findByRowID(destinations, model.GetID()); ok {
		return destination, status, nil
	}

	// No row has this id, so it is a hog function group from a release before row ids, an
	// import that named one hog function, or a row PostHog replaced with a new one. Match on
	// hog functions; MapResponseToModel then moves the state to the match's row id.
	stateIDs := hogFunctionIDsFromState(ctx, model)
	matches := make([]httpclient.LogsAlertDestination, 0, 1)
	for _, destination := range destinations {
		if sharesHogFunction(destination.HogFunctionIDs, stateIDs) {
			matches = append(matches, destination)
		}
	}
	if len(matches) > 1 {
		return httpclient.LogsAlertDestination{}, 0, fmt.Errorf(
			"log alert destination %s is split across %d server groups; refusing to abandon active HogFunctions. "+
				"Resolve the grouping at %s, or remove this resource from state with "+
				"'terraform state rm' and re-import the intended group",
			model.GetID(), len(matches),
			logsAlertNotificationsPath(model.GetEffectiveProjectID(), model.GetAlertID()))
	}
	if len(matches) == 1 {
		allStateIDsPresent := !slices.ContainsFunc(stateIDs, func(id string) bool {
			return !slices.Contains(matches[0].HogFunctionIDs, id)
		})
		isImport := model.Type.IsNull()
		if allStateIDsPresent && (isImport || len(stateIDs) == len(matches[0].HogFunctionIDs)) {
			return matches[0], status, nil
		}
		return httpclient.LogsAlertDestination{}, 0, fmt.Errorf(
			"log alert destination %s is only partially present in its server group; refusing to rewrite state. "+
				"Resolve the grouping at %s, or remove this resource from state with "+
				"'terraform state rm' and re-import the intended group",
			model.GetID(),
			logsAlertNotificationsPath(model.GetEffectiveProjectID(), model.GetAlertID()))
	}

	return httpclient.LogsAlertDestination{}, http.StatusNotFound, fmt.Errorf(
		"log alert destination %s not found on alert %s", model.GetID(), model.GetAlertID())
}

func (o LogsAlertDestinationOps) Update(_ context.Context, _ httpclient.PosthogClient, _ LogsAlertDestinationTFModel, _ httpclient.LogsAlertDestinationRequest) (httpclient.LogsAlertDestination, httpclient.HTTPStatusCode, error) {
	return httpclient.LogsAlertDestination{}, 0, fmt.Errorf("log alert destinations do not support updates; this is a bug if you see this error")
}

func (o LogsAlertDestinationOps) Delete(ctx context.Context, client httpclient.PosthogClient, model LogsAlertDestinationTFModel) (httpclient.HTTPStatusCode, error) {
	if trackedByRowID(ctx, model) {
		return client.DeleteLogsAlertDestination(ctx, model.GetEffectiveProjectID(), model.GetAlertID(), model.GetID())
	}
	return client.DeleteLogsAlertDestinationHogFunctions(ctx, model.GetEffectiveProjectID(), model.GetAlertID(), hogFunctionIDsFromID(model))
}

func findByRowID(destinations []httpclient.LogsAlertDestination, id string) (httpclient.LogsAlertDestination, bool) {
	if id == "" || strings.Contains(id, hogFunctionIDSeparator) {
		return httpclient.LogsAlertDestination{}, false
	}
	for _, destination := range destinations {
		if hasRowID(destination) && *destination.ID == id {
			return destination, true
		}
	}
	return httpclient.LogsAlertDestination{}, false
}

// trackedByRowID tells a row id from the hog function group id releases before row ids
// wrote. Both are UUIDs, so a one-function group looks like a row id; a group's id is always
// made from the hog_function_ids recorded beside it, and a row id never is.
func trackedByRowID(ctx context.Context, model LogsAlertDestinationTFModel) bool {
	id := model.GetID()
	if strings.Contains(id, hogFunctionIDSeparator) {
		return false
	}
	recorded := recordedHogFunctionIDs(ctx, model)
	return len(recorded) == 0 || id != hogFunctionGroupID(recorded)
}

func recordedHogFunctionIDs(ctx context.Context, model LogsAlertDestinationTFModel) []string {
	recorded, diags := core.ExtractTags(ctx, model.HogFunctionIDs)
	if diags.HasError() {
		return nil
	}
	return recorded
}

func hogFunctionIDsFromID(model LogsAlertDestinationTFModel) []string {
	return strings.Split(model.GetID(), hogFunctionIDSeparator)
}

// hogFunctionIDsFromState names the hog functions to look for when no row matches the
// state's id. Those recorded in hog_function_ids come first: for a row id they are the only
// ones, and PostHog keeps them on a row that replaces it, which is how a destination survives
// its row being replaced. An import has none recorded and names one hog function as its id.
func hogFunctionIDsFromState(ctx context.Context, model LogsAlertDestinationTFModel) []string {
	if recorded := recordedHogFunctionIDs(ctx, model); len(recorded) > 0 {
		return recorded
	}
	return hogFunctionIDsFromID(model)
}

func sharesHogFunction(destinationIDs, stateIDs []string) bool {
	return slices.ContainsFunc(destinationIDs, func(id string) bool {
		return slices.Contains(stateIDs, id)
	})
}
