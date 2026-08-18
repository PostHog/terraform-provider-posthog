package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// slackDestination is the model a valid Slack configuration produces, as the starting point
// for cases that drop or add one attribute.
func slackDestination() LogsAlertDestinationTFModel {
	return LogsAlertDestinationTFModel{
		AlertID:          types.StringValue("019dbe94-cec8-781b-9470-4a970cd69ebf"),
		Type:             types.StringValue(destinationTypeSlack),
		SlackWorkspaceID: types.Int64Value(1),
		SlackChannelID:   types.StringValue("C0123456789"),
	}
}

func TestLogsAlertDestinationBuildCreateRequest(t *testing.T) {
	tests := map[string]struct {
		model LogsAlertDestinationTFModel
		want  httpclient.LogsAlertDestinationRequest
	}{
		"slack with display name": {
			model: func() LogsAlertDestinationTFModel {
				m := slackDestination()
				m.SlackChannelName = types.StringValue("#alerts")
				return m
			}(),
			want: httpclient.LogsAlertDestinationRequest{
				Type:             destinationTypeSlack,
				SlackWorkspaceID: util.Int64Ptr(1),
				SlackChannelID:   util.StringPtr("C0123456789"),
				SlackChannelName: util.StringPtr("#alerts"),
			},
		},
		"webhook": {
			model: LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeWebhook),
				WebhookURL: types.StringValue("https://example.com/hook"),
			},
			want: httpclient.LogsAlertDestinationRequest{
				Type:       destinationTypeWebhook,
				WebhookURL: util.StringPtr("https://example.com/hook"),
			},
		},
		"teams": {
			model: LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeTeams),
				WebhookURL: types.StringValue("https://outlook.office.com/webhook/abc"),
			},
			want: httpclient.LogsAlertDestinationRequest{
				Type:       destinationTypeTeams,
				WebhookURL: util.StringPtr("https://outlook.office.com/webhook/abc"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, diags := LogsAlertDestinationOps{}.BuildCreateRequest(context.Background(), test.model)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, test.want, req)
		})
	}
}

// There is no update endpoint, and every attribute forces replacement, so an update path
// that ever runs is a provider bug rather than something to send to the API.
func TestLogsAlertDestinationUpdateIsRefused(t *testing.T) {
	_, diags := LogsAlertDestinationOps{}.BuildUpdateRequest(context.Background(), slackDestination(), slackDestination())
	require.True(t, diags.HasError(), "expected an update to be refused")

	_, _, err := LogsAlertDestinationOps{}.Update(context.Background(), httpclient.PosthogClient{}, slackDestination(), httpclient.LogsAlertDestinationRequest{})
	require.Error(t, err)
}

// A create response carries only the hog function ids, so everything the practitioner
// configured has to survive being mapped through it.
func TestLogsAlertDestinationMapResponseToModel_CreateKeepsConfiguredValues(t *testing.T) {
	model := slackDestination()
	model.SlackChannelName = types.StringValue("#alerts")

	resp := httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-2", "hf-1"}}

	diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), resp, &model)

	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	// The id sorts the ids so it does not depend on the order PostHog returns them in. The
	// attribute keeps the response order, which a set makes immaterial.
	assert.Equal(t, types.StringValue("hf-1,hf-2"), model.ID)
	assert.Equal(t, stringSet(t, "hf-2", "hf-1"), model.HogFunctionIDs)
	assert.Equal(t, types.StringValue(destinationTypeSlack), model.Type)
	assert.Equal(t, types.Int64Value(1), model.SlackWorkspaceID)
	assert.Equal(t, types.StringValue("C0123456789"), model.SlackChannelID)
	assert.Equal(t, types.StringValue("#alerts"), model.SlackChannelName)
}

// A read returns the whole group, and is the only response that can correct drift.
func TestLogsAlertDestinationMapResponseToModel_ReadAdoptsServerValues(t *testing.T) {
	model := slackDestination()

	resp := httpclient.LogsAlertDestination{
		HogFunctionIDs:   []string{"hf-1"},
		Type:             destinationTypeSlack,
		SlackWorkspaceID: util.Int64Ptr(7),
		SlackChannelID:   util.StringPtr("C9999999999"),
	}

	diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), resp, &model)

	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	assert.Equal(t, types.Int64Value(7), model.SlackWorkspaceID)
	assert.Equal(t, types.StringValue("C9999999999"), model.SlackChannelID)
	// Absent from a Slack group rather than emptied, so it must read back as null.
	assert.True(t, model.WebhookURL.IsNull(), "webhook_url does not apply to a Slack destination")
}

// slack_channel_name is write-only: PostHog uses it to build the display name and never
// stores it, so no response can return it. Mapping it would null out the configured value
// and every plan would show drift.
func TestLogsAlertDestinationMapResponseToModel_LeavesSlackChannelNameAlone(t *testing.T) {
	tests := map[string]struct {
		channelName types.String
		resp        httpclient.LogsAlertDestination
	}{
		"configured, create response": {
			channelName: types.StringValue("#alerts"),
			resp:        httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-1"}},
		},
		"configured, read response": {
			channelName: types.StringValue("#alerts"),
			resp: httpclient.LogsAlertDestination{
				HogFunctionIDs:   []string{"hf-1"},
				Type:             destinationTypeSlack,
				SlackWorkspaceID: util.Int64Ptr(1),
				SlackChannelID:   util.StringPtr("C0123456789"),
			},
		},
		// An imported destination has no configured value, and nothing can supply one.
		"unset, read response": {
			channelName: types.StringNull(),
			resp: httpclient.LogsAlertDestination{
				HogFunctionIDs:   []string{"hf-1"},
				Type:             destinationTypeSlack,
				SlackWorkspaceID: util.Int64Ptr(1),
				SlackChannelID:   util.StringPtr("C0123456789"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			model := slackDestination()
			model.SlackChannelName = test.channelName

			diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), test.resp, &model)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, test.channelName, model.SlackChannelName)
		})
	}
}

// Without ids there is no identity to write into state, and the resource would be silently
// forgotten and then created a second time.
func TestLogsAlertDestinationMapResponseToModel_RejectsEmptyGroup(t *testing.T) {
	model := slackDestination()

	diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlertDestination{}, &model)

	require.True(t, diags.HasError(), "expected a diagnostic")
	assert.Contains(t, diags.Errors()[0].Summary(), "no hog functions")
}

func TestValidateLogsAlertDestinationPlan(t *testing.T) {
	webhookDestination := LogsAlertDestinationTFModel{
		Type:       types.StringValue(destinationTypeWebhook),
		WebhookURL: types.StringValue("https://example.com/hook"),
	}

	// Every attribute this validates is Optional or Required, never Computed, so Terraform
	// puts the same values in the plan and the config. Cases that need them to differ set
	// both fields explicitly.
	tests := map[string]struct {
		model     LogsAlertDestinationTFModel
		expectErr string
	}{
		"slack with workspace and channel": {
			model: slackDestination(),
		},
		"slack missing workspace": {
			model: func() LogsAlertDestinationTFModel {
				m := slackDestination()
				m.SlackWorkspaceID = types.Int64Null()
				return m
			}(),
			expectErr: "Missing Slack destination settings",
		},
		"slack missing channel": {
			model: func() LogsAlertDestinationTFModel {
				m := slackDestination()
				m.SlackChannelID = types.StringNull()
				return m
			}(),
			expectErr: "Missing Slack destination settings",
		},
		"slack with a webhook url": {
			model: func() LogsAlertDestinationTFModel {
				m := slackDestination()
				m.WebhookURL = types.StringValue("https://example.com/hook")
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		"webhook with a url": {
			model: webhookDestination,
		},
		"webhook missing its url": {
			model:     LogsAlertDestinationTFModel{Type: types.StringValue(destinationTypeWebhook)},
			expectErr: "Missing destination URL",
		},
		"teams with a url": {
			model: LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeTeams),
				WebhookURL: types.StringValue("https://outlook.office.com/webhook/abc"),
			},
		},
		"teams missing its url": {
			model:     LogsAlertDestinationTFModel{Type: types.StringValue(destinationTypeTeams)},
			expectErr: "Missing destination URL",
		},
		"webhook with a slack workspace": {
			model: func() LogsAlertDestinationTFModel {
				m := webhookDestination
				m.SlackWorkspaceID = types.Int64Value(1)
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		"webhook with a slack channel": {
			model: func() LogsAlertDestinationTFModel {
				m := webhookDestination
				m.SlackChannelID = types.StringValue("C0123456789")
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		"webhook with a slack channel name": {
			model: func() LogsAlertDestinationTFModel {
				m := webhookDestination
				m.SlackChannelName = types.StringValue("#alerts")
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		// Cannot conclude anything: the reference may resolve to any type.
		"unresolved type skips every check": {
			model: LogsAlertDestinationTFModel{Type: types.StringUnknown()},
		},
		// Same for the URL itself, which may resolve to a value.
		"unresolved webhook url is left to the API": {
			model: LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeWebhook),
				WebhookURL: types.StringUnknown(),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			diags := validateLogsAlertDestinationPlan(test.model, test.model)

			if test.expectErr == "" {
				assert.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
				return
			}
			require.True(t, diags.HasError(), "expected a diagnostic")
			assert.Contains(t, diags.Errors()[0].Summary(), test.expectErr)
		})
	}
}

// The id is the whole group, so a destroy hands PostHog every hog function it created.
func TestHogFunctionIDsFromState(t *testing.T) {
	model := LogsAlertDestinationTFModel{}
	require.NoError(t, model.SetID("hf-1,hf-2,hf-3"))

	assert.Equal(t, []string{"hf-1", "hf-2", "hf-3"}, hogFunctionIDsFromState(model))
}

func TestSharesHogFunction(t *testing.T) {
	group := []string{"hf-1", "hf-2", "hf-3", "hf-4"}

	assert.True(t, sharesHogFunction(group, group), "the same group matches itself")
	// An import names a single hog function id, which still has to find its group.
	assert.True(t, sharesHogFunction(group, []string{"hf-3"}), "one shared id identifies the group")
	assert.False(t, sharesHogFunction(group, []string{"hf-9"}), "a different destination must not match")
	assert.False(t, sharesHogFunction(group, nil), "nothing in state matches nothing")
}
