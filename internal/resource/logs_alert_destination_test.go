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

func TestLogsAlertDestinationUpdateIsRefused(t *testing.T) {
	_, diags := LogsAlertDestinationOps{}.BuildUpdateRequest(context.Background(), slackDestination(), slackDestination())
	require.True(t, diags.HasError(), "expected an update to be refused")

	_, _, err := LogsAlertDestinationOps{}.Update(context.Background(), httpclient.PosthogClient{}, slackDestination(), httpclient.LogsAlertDestinationRequest{})
	require.Error(t, err)
}

func TestLogsAlertDestinationID_DoesNotDependOnTheOrderPostHogReturnsTheIDsIn(t *testing.T) {
	responseOrder := []string{"hf-2", "hf-1", "hf-3"}

	assert.Equal(t, "hf-1,hf-2,hf-3", logsAlertDestinationID(responseOrder))
	assert.Equal(t, logsAlertDestinationID([]string{"hf-3", "hf-1", "hf-2"}), logsAlertDestinationID(responseOrder))
	assert.Equal(t, []string{"hf-2", "hf-1", "hf-3"}, responseOrder, "building the id must not reorder the response")
}

func TestLogsAlertDestinationID_RoundTripsThroughHogFunctionIDsFromState(t *testing.T) {
	hogFunctionIDs := []string{"019dbe94-cec8-781b-9470-4a970cd69ebf", "019dbe94-cec8-7000-8000-000000000001"}

	model := LogsAlertDestinationTFModel{}
	require.NoError(t, model.SetID(logsAlertDestinationID(hogFunctionIDs)))

	assert.ElementsMatch(t, hogFunctionIDs, hogFunctionIDsFromState(model))
}

func TestLogsAlertDestinationMapResponseToModel_CreateKeepsConfiguredValues(t *testing.T) {
	model := slackDestination()
	model.SlackChannelName = types.StringValue("#alerts")

	resp := httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-2", "hf-1"}}

	diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), resp, &model)

	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	assert.Equal(t, types.StringValue("hf-1,hf-2"), model.ID)
	assert.Equal(t, stringSet(t, "hf-2", "hf-1"), model.HogFunctionIDs)
	assert.Equal(t, types.StringValue(destinationTypeSlack), model.Type)
	assert.Equal(t, types.Int64Value(1), model.SlackWorkspaceID)
	assert.Equal(t, types.StringValue("C0123456789"), model.SlackChannelID)
	assert.Equal(t, types.StringValue("#alerts"), model.SlackChannelName)
}

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

func TestLogsAlertDestinationMapResponseToModel_AdoptsTheReturnedWebhookURL(t *testing.T) {
	model := LogsAlertDestinationTFModel{Type: types.StringValue(destinationTypeWebhook)}
	resp := httpclient.LogsAlertDestination{
		HogFunctionIDs: []string{"hf-1"},
		Type:           destinationTypeWebhook,
		WebhookURL:     util.StringPtr("https://example.com/from-api"),
	}

	diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), resp, &model)

	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	assert.Equal(t, types.StringValue("https://example.com/from-api"), model.WebhookURL)
}

func TestLogsAlertDestinationMapResponseToModel_KeepsTheWriteOnlySlackChannelName(t *testing.T) {
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
		"unset on import, read response": {
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

func TestLogsAlertDestinationMapResponseToModel_RejectsADestinationWithNoIdentity(t *testing.T) {
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

	tests := map[string]struct {
		planAndConfig LogsAlertDestinationTFModel
		expectErr     string
	}{
		"slack with workspace and channel": {
			planAndConfig: slackDestination(),
		},
		"slack missing workspace": {
			planAndConfig: func() LogsAlertDestinationTFModel {
				m := slackDestination()
				m.SlackWorkspaceID = types.Int64Null()
				return m
			}(),
			expectErr: "Missing Slack destination settings",
		},
		"slack missing channel": {
			planAndConfig: func() LogsAlertDestinationTFModel {
				m := slackDestination()
				m.SlackChannelID = types.StringNull()
				return m
			}(),
			expectErr: "Missing Slack destination settings",
		},
		"slack with a webhook url": {
			planAndConfig: func() LogsAlertDestinationTFModel {
				m := slackDestination()
				m.WebhookURL = types.StringValue("https://example.com/hook")
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		"webhook with a url": {
			planAndConfig: webhookDestination,
		},
		"webhook missing its url": {
			planAndConfig: LogsAlertDestinationTFModel{Type: types.StringValue(destinationTypeWebhook)},
			expectErr:     "Missing destination URL",
		},
		"teams with a url": {
			planAndConfig: LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeTeams),
				WebhookURL: types.StringValue("https://outlook.office.com/webhook/abc"),
			},
		},
		"teams missing its url": {
			planAndConfig: LogsAlertDestinationTFModel{Type: types.StringValue(destinationTypeTeams)},
			expectErr:     "Missing destination URL",
		},
		"webhook with a slack workspace": {
			planAndConfig: func() LogsAlertDestinationTFModel {
				m := webhookDestination
				m.SlackWorkspaceID = types.Int64Value(1)
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		"webhook with a slack channel": {
			planAndConfig: func() LogsAlertDestinationTFModel {
				m := webhookDestination
				m.SlackChannelID = types.StringValue("C0123456789")
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		"webhook with a slack channel name": {
			planAndConfig: func() LogsAlertDestinationTFModel {
				m := webhookDestination
				m.SlackChannelName = types.StringValue("#alerts")
				return m
			}(),
			expectErr: "Attribute does not apply to this destination type",
		},
		"a type that may still resolve to anything skips every check": {
			planAndConfig: LogsAlertDestinationTFModel{Type: types.StringUnknown()},
		},
		"a webhook url that may still resolve to a value is left to the API": {
			planAndConfig: LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeWebhook),
				WebhookURL: types.StringUnknown(),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			diags := validateLogsAlertDestinationPlan(test.planAndConfig, test.planAndConfig)

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

func hogFunctionsPath() string {
	return fmt.Sprintf("/api/projects/%s/hog_functions/", testLogsAlertDestinationProjectID)
}

func destinationInState(id string) LogsAlertDestinationTFModel {
	return LogsAlertDestinationTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(id)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testLogsAlertDestinationProjectID)},
		AlertID:                types.StringValue(testLogsAlertDestinationAlertID),
	}
}

func writeHogFunctionPage(t *testing.T, w http.ResponseWriter, next any, hogFunctions ...httpclient.HogFunction) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(map[string]any{
		"count":    len(hogFunctions),
		"next":     next,
		"previous": nil,
		"results":  hogFunctions,
	}))
}

func TestLogsAlertDestinationRead_FindsADestinationOnALaterPage(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, hogFunctionsPath(), r.URL.Path)

		if r.URL.Query().Get("offset") == "" {
			writeHogFunctionPage(t, w, server.URL+hogFunctionsPath()+"?limit=1&offset=1",
				httpclient.HogFunction{ID: "hf-1", TemplateID: util.StringPtr("template-webhook"), Inputs: map[string]interface{}{
					"url": map[string]interface{}{"value": "https://example.com/first"},
				}})
			return
		}
		writeHogFunctionPage(t, w, nil,
			httpclient.HogFunction{ID: "hf-2", TemplateID: util.StringPtr("template-microsoft-teams"), Inputs: map[string]interface{}{
				"webhookUrl": map[string]interface{}{"value": "https://example.com/teams"},
			}},
			httpclient.HogFunction{ID: "hf-3", TemplateID: util.StringPtr("template-microsoft-teams"), Inputs: map[string]interface{}{
				"webhookUrl": map[string]interface{}{"value": "https://example.com/teams"},
			}})
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

	destination, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, destinationInState("hf-2,hf-3"))

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, []string{"hf-2", "hf-3"}, destination.HogFunctionIDs)
}

func TestLogsAlertDestinationRead_ReportsADestinationMissingFromALiveAlertAsDeleted(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, hogFunctionsPath(), r.URL.Path)
		writeHogFunctionPage(t, w, nil,
			httpclient.HogFunction{ID: "hf-9", TemplateID: util.StringPtr("template-webhook"), Inputs: map[string]interface{}{
				"url": map[string]interface{}{"value": "https://example.com/other"},
			}})
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

	_, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, destinationInState("hf-1"))

	require.Error(t, err)
	assert.Equal(t, http.StatusNotFound, int(status), "a destination gone from a live alert must be removed from state")
}

func TestLogsAlertDestinationDelete_TreatsA404AsAlreadyDeletedWithoutLookingTheAlertUp(t *testing.T) {
	var requestedPaths []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPaths = append(requestedPaths, r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

	status, err := LogsAlertDestinationOps{}.Delete(context.Background(), client, destinationInState("hf-1,hf-2"))

	require.Error(t, err)
	assert.Equal(t, http.StatusNotFound, int(status), "a 404 from delete means the destination is already gone")
	assert.Equal(t, []string{logsAlertDestinationsPath() + "delete"}, requestedPaths)
}

func TestHogFunctionIDsFromState_RecoversTheWholeGroupFromTheID(t *testing.T) {
	model := LogsAlertDestinationTFModel{}
	require.NoError(t, model.SetID("hf-1,hf-2,hf-3"))

	assert.Equal(t, []string{"hf-1", "hf-2", "hf-3"}, hogFunctionIDsFromState(model))
}

func TestSharesHogFunction(t *testing.T) {
	group := []string{"hf-1", "hf-2", "hf-3", "hf-4"}

	assert.True(t, sharesHogFunction(group, group), "the same group matches itself")
	assert.True(t, sharesHogFunction(group, []string{"hf-3"}), "one shared id identifies the group, which is what lets an import name a single id")
	assert.False(t, sharesHogFunction(group, []string{"hf-9"}), "a different destination must not match")
	assert.False(t, sharesHogFunction(group, nil), "nothing in state matches nothing")
}
