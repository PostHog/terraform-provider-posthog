package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testLogsAlertDestinationProjectID = "123"
	testLogsAlertDestinationAlertID   = "019dbe94-cec8-781b-9470-4a970cd69ebf"
)

// slackDestination is the model a valid Slack configuration produces, as the starting point
// for cases that drop or add one attribute.
func slackDestination() LogsAlertDestinationTFModel {
	return LogsAlertDestinationTFModel{
		AlertID:          types.StringValue(testLogsAlertDestinationAlertID),
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
}

// webhook_url cannot round-trip: the read redacts it to scheme and host. Adopting the
// redacted value would show drift on every plan, and the next apply would offer the mask to
// PostHog as if it were the real URL.
func TestLogsAlertDestinationMapResponseToModel_NeverAdoptsRedactedWebhookURL(t *testing.T) {
	tests := map[string]types.String{
		"configured": types.StringValue("https://example.com/hook?token=s3cret"),
		// An imported destination has no configured value, and nothing can supply one.
		"unset on import": types.StringNull(),
	}

	for name, configured := range tests {
		t.Run(name, func(t *testing.T) {
			model := LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeWebhook),
				WebhookURL: configured,
			}

			resp := httpclient.LogsAlertDestination{
				HogFunctionIDs: []string{"hf-1"},
				Type:           destinationTypeWebhook,
				WebhookURL:     util.StringPtr("https://example.com/…"),
			}

			diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), resp, &model)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, configured, model.WebhookURL)
		})
	}
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

func logsAlertPath() string {
	return fmt.Sprintf("/api/projects/%s/logs/alerts/%s/", testLogsAlertDestinationProjectID, testLogsAlertDestinationAlertID)
}

func logsAlertDestinationsPath() string {
	return logsAlertPath() + "destinations/"
}

// destinationInState is the state a read starts from: where to look, and the hog function ids
// that identify the destination among the alert's.
func destinationInState(id string) LogsAlertDestinationTFModel {
	return LogsAlertDestinationTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(id)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testLogsAlertDestinationProjectID)},
		AlertID:                types.StringValue(testLogsAlertDestinationAlertID),
	}
}

func writeDestinationPage(t *testing.T, w http.ResponseWriter, next any, destinations ...httpclient.LogsAlertDestination) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(map[string]any{
		"count":    len(destinations),
		"next":     next,
		"previous": nil,
		"results":  destinations,
	}))
}

// The read scans the alert's destinations for its own hog functions, so it has to see every
// page. Stopping at the first would report a destination on a later page as deleted, and
// Terraform would drop a live destination from state.
func TestLogsAlertDestinationRead_FindsADestinationOnALaterPage(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, logsAlertDestinationsPath(), r.URL.Path)

		if r.URL.Query().Get("offset") == "" {
			writeDestinationPage(t, w, server.URL+logsAlertDestinationsPath()+"?limit=1&offset=1",
				httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-1"}, Type: destinationTypeWebhook})
			return
		}
		writeDestinationPage(t, w, nil,
			httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-2", "hf-3"}, Type: destinationTypeTeams})
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

	destination, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, destinationInState("hf-2,hf-3"))

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, []string{"hf-2", "hf-3"}, destination.HogFunctionIDs)
}

// A 404 from the destinations list is ambiguous: every non-2xx is an error, so a PostHog
// without the endpoint answers exactly like a deleted alert. Only the deletion may reach the
// generic resource as a 404, which is what removes the resource from state.
func TestLogsAlertDestinationRead_TellsAMissingEndpointFromADeletedAlert(t *testing.T) {
	tests := map[string]struct {
		alertStatus int
		wantRemoval bool
		wantMessage string
	}{
		"alert gone, so its destinations went with it": {
			alertStatus: http.StatusNotFound,
			wantRemoval: true,
		},
		"alert present, so the endpoint is the thing that is missing": {
			alertStatus: http.StatusOK,
			wantMessage: "does not serve GET",
		},
		// Nothing can be concluded, so erroring is the only safe answer.
		"alert lookup itself failed": {
			alertStatus: http.StatusInternalServerError,
			wantMessage: "cannot tell a deleted destination from an unavailable endpoint",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == logsAlertDestinationsPath() {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				require.Equal(t, logsAlertPath(), r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.alertStatus)
				require.NoError(t, json.NewEncoder(w).Encode(httpclient.LogsAlert{ID: testLogsAlertDestinationAlertID}))
			}))
			defer server.Close()

			client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

			_, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, destinationInState("hf-1"))

			require.Error(t, err)
			if test.wantRemoval {
				assert.Equal(t, http.StatusNotFound, int(status), "a deleted alert must remove the resource from state")
				return
			}
			assert.NotEqual(t, http.StatusNotFound, int(status), "only a deletion may remove the resource from state")
			assert.Contains(t, err.Error(), test.wantMessage)
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
