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

const (
	destinationTypeSlack   = "slack"
	destinationTypeWebhook = "webhook"
	destinationTypeTeams   = "teams"
)

var destinationTypes = []string{destinationTypeSlack, destinationTypeWebhook, destinationTypeTeams}

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
			"Each resource represents one user-visible destination. PostHog implements that destination as a " +
			"group of hog functions, one per alert transition " +
			"(firing, resolved, errored, auto-disabled), sharing the configuration below. The group has no id " +
			"of its own, so this resource's `id` is the group's `hog_function_ids`, sorted and joined by " +
			"commas. Those hog functions are owned by the alert: `posthog_hog_function` cannot create, update, or " +
			"delete them. Terraform reads them through the same generic Hog Function list API as the PostHog UI, " +
			"then uses the alert destinations API to create and delete each managed group.",
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
					stringvalidator.OneOf(destinationTypes...),
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
				Optional:  true,
				Sensitive: true,
				MarkdownDescription: "URL to POST the notification to. Required when `type` is `webhook` or " +
					"`teams`. For `teams`, this is the Microsoft Teams incoming webhook URL. Marked sensitive " +
					"because the secret is in the URL: whoever holds it can post to the channel.\n\n" +
					"PostHog returns the URL through the generic Hog Function API. Terraform stores it as a " +
					"sensitive value, detects changes made outside Terraform, and adopts it during import.",
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

type writeOnlyDestinationAttributes struct {
	SlackChannelName types.String
}

func writeOnlyAttributesOf(model LogsAlertDestinationTFModel) writeOnlyDestinationAttributes {
	return writeOnlyDestinationAttributes{
		SlackChannelName: model.SlackChannelName,
	}
}

func (w writeOnlyDestinationAttributes) restoreTo(model *LogsAlertDestinationTFModel) {
	model.SlackChannelName = w.SlackChannelName
}

func destinationIncludesConfiguration(resp httpclient.LogsAlertDestination) bool {
	return resp.Type != ""
}

func (o LogsAlertDestinationOps) MapResponseToModel(ctx context.Context, resp httpclient.LogsAlertDestination, model *LogsAlertDestinationTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if len(resp.HogFunctionIDs) == 0 {
		diags.AddError(
			"Log alert destination has no hog functions",
			"PostHog returned a destination with no hog_function_ids, which Terraform cannot track. "+
				"The destination may exist on the alert; check it in the PostHog UI before applying again.",
		)
		return diags
	}

	writeOnly := writeOnlyAttributesOf(*model)

	model.ID = types.StringValue(logsAlertDestinationID(resp.HogFunctionIDs))

	hogFunctionIDs, d := core.TagsToSet(ctx, resp.HogFunctionIDs)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	model.HogFunctionIDs = hogFunctionIDs

	if destinationIncludesConfiguration(resp) {
		model.Type = types.StringValue(resp.Type)
		model.SlackWorkspaceID = util.PtrToInt64(resp.SlackWorkspaceID)
		model.SlackChannelID = core.PtrToStringNullIfEmptyTrimmed(resp.SlackChannelID)
		model.WebhookURL = core.PtrToStringNullIfEmptyTrimmed(resp.WebhookURL)
	}

	writeOnly.restoreTo(model)

	return diags
}

func logsAlertDestinationID(hogFunctionIDs []string) string {
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

type effectiveValue[T comparable] struct {
	value T
	known bool
}

func effectiveString(plan, config types.String) effectiveValue[string] {
	value, known := util.ResolveString(plan, config, "")
	return effectiveValue[string]{value: value, known: known}
}

func effectiveInt64(plan, config types.Int64) effectiveValue[int64] {
	value, known := util.ResolveInt64(plan, config, 0)
	return effectiveValue[int64]{value: value, known: known}
}

func (e effectiveValue[T]) isSet() bool {
	var unset T
	return e.known && e.value != unset
}

func (e effectiveValue[T]) isOmitted() bool {
	var unset T
	return e.known && e.value == unset
}

type destinationPlanAttributes struct {
	slackWorkspaceID effectiveValue[int64]
	slackChannelID   effectiveValue[string]
	slackChannelName effectiveValue[string]
	webhookURL       effectiveValue[string]
}

func resolveDestinationPlanAttributes(plan, config LogsAlertDestinationTFModel) destinationPlanAttributes {
	return destinationPlanAttributes{
		slackWorkspaceID: effectiveInt64(plan.SlackWorkspaceID, config.SlackWorkspaceID),
		slackChannelID:   effectiveString(plan.SlackChannelID, config.SlackChannelID),
		slackChannelName: effectiveString(plan.SlackChannelName, config.SlackChannelName),
		webhookURL:       effectiveString(plan.WebhookURL, config.WebhookURL),
	}
}

func validateLogsAlertDestinationPlan(plan, config LogsAlertDestinationTFModel) diag.Diagnostics {
	destinationType := effectiveString(plan.Type, config.Type)
	if !destinationType.isSet() {
		return nil
	}

	attributes := resolveDestinationPlanAttributes(plan, config)
	if destinationType.value == destinationTypeSlack {
		return attributes.slackDestinationErrors()
	}
	return attributes.webhookOrTeamsDestinationErrors(destinationType.value)
}

func (a destinationPlanAttributes) slackDestinationErrors() diag.Diagnostics {
	var diags diag.Diagnostics

	if a.slackWorkspaceID.isOmitted() {
		diags.AddAttributeError(
			path.Root("slack_workspace_id"),
			"Missing Slack destination settings",
			`type = "slack" requires slack_workspace_id, the ID of the Slack integration connected to `+
				`this project.`,
		)
	}
	if a.slackChannelID.isOmitted() {
		diags.AddAttributeError(
			path.Root("slack_channel_id"),
			"Missing Slack destination settings",
			`type = "slack" requires slack_channel_id, the ID of the channel to post into, such as `+
				`C0123456789.`,
		)
	}
	if a.webhookURL.isSet() {
		diags.Append(destinationAttributeDoesNotApply("webhook_url", destinationTypeSlack)...)
	}

	return diags
}

func (a destinationPlanAttributes) webhookOrTeamsDestinationErrors(destinationType string) diag.Diagnostics {
	var diags diag.Diagnostics

	if a.webhookURL.isOmitted() {
		diags.AddAttributeError(
			path.Root("webhook_url"),
			"Missing destination URL",
			fmt.Sprintf("type = %q requires webhook_url, the URL to POST the notification to.", destinationType),
		)
	}
	if a.slackWorkspaceID.isSet() {
		diags.Append(destinationAttributeDoesNotApply("slack_workspace_id", destinationType)...)
	}
	if a.slackChannelID.isSet() {
		diags.Append(destinationAttributeDoesNotApply("slack_channel_id", destinationType)...)
	}
	if a.slackChannelName.isSet() {
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

	stateIDs := hogFunctionIDsFromState(model)
	for _, destination := range destinations {
		if sharesHogFunction(destination.HogFunctionIDs, stateIDs) {
			return destination, status, nil
		}
	}

	return httpclient.LogsAlertDestination{}, http.StatusNotFound, fmt.Errorf(
		"log alert destination %s not found on alert %s", model.GetID(), model.GetAlertID())
}

func (o LogsAlertDestinationOps) Update(_ context.Context, _ httpclient.PosthogClient, _ LogsAlertDestinationTFModel, _ httpclient.LogsAlertDestinationRequest) (httpclient.LogsAlertDestination, httpclient.HTTPStatusCode, error) {
	return httpclient.LogsAlertDestination{}, 0, fmt.Errorf("log alert destinations do not support updates; this is a bug if you see this error")
}

func (o LogsAlertDestinationOps) Delete(ctx context.Context, client httpclient.PosthogClient, model LogsAlertDestinationTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteLogsAlertDestination(ctx, model.GetEffectiveProjectID(), model.GetAlertID(), hogFunctionIDsFromState(model))
}

func hogFunctionIDsFromState(model LogsAlertDestinationTFModel) []string {
	return strings.Split(model.GetID(), hogFunctionIDSeparator)
}

func sharesHogFunction(destinationIDs, stateIDs []string) bool {
	return slices.ContainsFunc(destinationIDs, func(id string) bool {
		return slices.Contains(stateIDs, id)
	})
}
