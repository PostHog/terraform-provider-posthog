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
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

// The destination types PostHog accepts. Named so the schema validator, the request builder
// and the plan-time check cannot drift apart.
const (
	destinationTypeSlack   = "slack"
	destinationTypeWebhook = "webhook"
	destinationTypeTeams   = "teams"
)

// hogFunctionIDSeparator joins the hog function ids of a destination group into the
// Terraform id. A UUID never contains a comma, so the join is reversible.
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

// SetAlertID satisfies core.AlertIDSetter so the import parser can populate alert_id.
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
			"PostHog implements one destination as a group of hog functions, one per alert transition " +
			"(firing, resolved, errored, auto-disabled), sharing the configuration below. The group has no id " +
			"of its own, so this resource's `id` is the group's `hog_function_ids`, sorted and joined by " +
			"commas. Those hog functions are owned by the alert: `posthog_hog_function` cannot create, read or " +
			"delete them.\n\n" +
			"~> Managing destinations needs a PostHog instance whose logs alerts API exposes " +
			"`GET .../destinations/`. Older instances can only create and delete them, which leaves Terraform " +
			"unable to refresh what it created.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "The destination's `hog_function_ids`, sorted and joined by commas. " +
					"PostHog gives a destination no id of its own.",
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
				MarkdownDescription: "Where the notification goes: `slack`, `webhook`, or `teams` (Microsoft " +
					"Teams). A `slack` destination needs `slack_workspace_id` and `slack_channel_id`; `webhook` " +
					"and `teams` need `webhook_url`.",
				Validators: []validator.String{
					stringvalidator.OneOf(destinationTypeSlack, destinationTypeWebhook, destinationTypeTeams),
				},
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
				Optional: true,
				MarkdownDescription: "URL to POST the notification to. Required when `type` is `webhook` or " +
					"`teams`. For `teams`, this is the Microsoft Teams incoming webhook URL.",
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
					"alert transition. Useful for finding the destination in the PostHog UI.",
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

func (o LogsAlertDestinationOps) MapResponseToModel(ctx context.Context, resp httpclient.LogsAlertDestination, model *LogsAlertDestinationTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Without the ids there is no identity to write into state, and the next refresh would
	// drop the resource and the one after it would create a second destination.
	if len(resp.HogFunctionIDs) == 0 {
		diags.AddError(
			"Log alert destination has no hog functions",
			"PostHog returned a destination with no hog_function_ids, which Terraform cannot track. "+
				"The destination may exist on the alert; check it in the PostHog UI before applying again.",
		)
		return diags
	}

	model.ID = types.StringValue(logsAlertDestinationID(resp.HogFunctionIDs))

	hogFunctionIDs, d := core.TagsToSet(ctx, resp.HogFunctionIDs)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.HogFunctionIDs = hogFunctionIDs

	// A create response carries only the hog function ids, so the configured values are left
	// alone. A read returns the whole group and is the only response that can correct drift.
	if resp.Type != "" {
		model.Type = types.StringValue(resp.Type)
		model.SlackWorkspaceID = util.PtrToInt64(resp.SlackWorkspaceID)
		model.SlackChannelID = core.PtrToStringNullIfEmptyTrimmed(resp.SlackChannelID)
		model.WebhookURL = core.PtrToStringNullIfEmptyTrimmed(resp.WebhookURL)
	}

	// slack_channel_name is deliberately absent. PostHog only uses it to build the display
	// name and never stores it, so touching it here would null out the configured value and
	// show drift on every plan.

	return diags
}

// logsAlertDestinationID builds the resource id from a destination group. Sorted so the
// same group always yields the same id whatever order the API returns the ids in.
func logsAlertDestinationID(hogFunctionIDs []string) string {
	sorted := slices.Clone(hogFunctionIDs)
	slices.Sort(sorted)
	return strings.Join(sorted, hogFunctionIDSeparator)
}

// ModifyResourcePlan rejects configurations PostHog would refuse, so they fail at plan time
// with a message naming the problem.
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

// validateLogsAlertDestinationPlan checks that the configured attributes match the
// destination type. It is separate from ModifyResourcePlan so it can be unit tested against
// a model.
//
// The schema rejects 0 and the empty string as configured values, so the zero value each
// util.Resolve call falls back to means the attribute was omitted. An attribute that is not
// resolvable yet is skipped rather than guessed at, and the API gets the final say.
func validateLogsAlertDestinationPlan(plan, config LogsAlertDestinationTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	destinationType, typeKnown := util.ResolveString(plan.Type, config.Type, "")
	if !typeKnown || destinationType == "" {
		return diags
	}

	workspaceID, workspaceKnown := util.ResolveInt64(plan.SlackWorkspaceID, config.SlackWorkspaceID, 0)
	channelID, channelKnown := util.ResolveString(plan.SlackChannelID, config.SlackChannelID, "")
	channelName, channelNameKnown := util.ResolveString(plan.SlackChannelName, config.SlackChannelName, "")
	webhookURL, webhookKnown := util.ResolveString(plan.WebhookURL, config.WebhookURL, "")

	if destinationType == destinationTypeSlack {
		if workspaceKnown && workspaceID == 0 {
			diags.AddAttributeError(
				path.Root("slack_workspace_id"),
				"Missing Slack destination settings",
				`type = "slack" requires slack_workspace_id, the ID of the Slack integration connected to `+
					`this project.`,
			)
		}
		if channelKnown && channelID == "" {
			diags.AddAttributeError(
				path.Root("slack_channel_id"),
				"Missing Slack destination settings",
				`type = "slack" requires slack_channel_id, the ID of the channel to post into, such as `+
					`C0123456789.`,
			)
		}
		if webhookKnown && webhookURL != "" {
			diags.Append(destinationAttributeDoesNotApply("webhook_url", destinationType)...)
		}
		return diags
	}

	// webhook and teams. Any other value is rejected by the schema before this runs.
	if webhookKnown && webhookURL == "" {
		diags.AddAttributeError(
			path.Root("webhook_url"),
			"Missing destination URL",
			fmt.Sprintf("type = %q requires webhook_url, the URL to POST the notification to.", destinationType),
		)
	}
	if workspaceKnown && workspaceID != 0 {
		diags.Append(destinationAttributeDoesNotApply("slack_workspace_id", destinationType)...)
	}
	if channelKnown && channelID != "" {
		diags.Append(destinationAttributeDoesNotApply("slack_channel_id", destinationType)...)
	}
	if channelNameKnown && channelName != "" {
		diags.Append(destinationAttributeDoesNotApply("slack_channel_name", destinationType)...)
	}

	return diags
}

// destinationAttributeDoesNotApply reports an attribute that belongs to a different
// destination type. PostHog ignores it rather than failing, so a config that sets a Slack
// channel on a webhook destination would otherwise apply cleanly and notify the wrong place.
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

// Read finds this destination among the alert's destinations. The API lists them all and
// offers no way to fetch one, since a destination has no id to fetch it by.
func (o LogsAlertDestinationOps) Read(ctx context.Context, client httpclient.PosthogClient, model LogsAlertDestinationTFModel) (httpclient.LogsAlertDestination, httpclient.HTTPStatusCode, error) {
	destinations, status, err := client.ListLogsAlertDestinations(ctx, model.GetEffectiveProjectID(), model.GetAlertID())
	if err != nil {
		return httpclient.LogsAlertDestination{}, status, err
	}

	stateIDs := hogFunctionIDsFromState(model)
	for _, destination := range destinations {
		if sharesHogFunction(destination.HogFunctionIDs, stateIDs) {
			return destination, status, nil
		}
	}

	// The alert is still there but this destination is not, so report it the way every
	// deleted resource is reported and let the generic resource drop it from state.
	return httpclient.LogsAlertDestination{}, http.StatusNotFound, fmt.Errorf(
		"log alert destination %s not found on alert %s", model.GetID(), model.GetAlertID())
}

func (o LogsAlertDestinationOps) Update(_ context.Context, _ httpclient.PosthogClient, _ LogsAlertDestinationTFModel, _ httpclient.LogsAlertDestinationRequest) (httpclient.LogsAlertDestination, httpclient.HTTPStatusCode, error) {
	return httpclient.LogsAlertDestination{}, 0, fmt.Errorf("log alert destinations do not support updates; this is a bug if you see this error")
}

func (o LogsAlertDestinationOps) Delete(ctx context.Context, client httpclient.PosthogClient, model LogsAlertDestinationTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteLogsAlertDestination(ctx, model.GetEffectiveProjectID(), model.GetAlertID(), hogFunctionIDsFromState(model))
}

// hogFunctionIDsFromState recovers the group from the id. The id is built by joining those
// same ids, so this is exact, and it works during an import where hog_function_ids has not
// been read yet.
func hogFunctionIDsFromState(model LogsAlertDestinationTFModel) []string {
	return strings.Split(model.GetID(), hogFunctionIDSeparator)
}

// sharesHogFunction reports whether two groups have a hog function in common. One shared id
// is enough to identify a destination, because the ids are UUIDs owned by a single group.
// Matching on one rather than on the whole set is what lets an import name a single id, and
// keeps the resource attached if PostHog ever changes how many hog functions it builds.
func sharesHogFunction(destinationIDs, stateIDs []string) bool {
	return slices.ContainsFunc(destinationIDs, func(id string) bool {
		return slices.Contains(stateIDs, id)
	})
}
