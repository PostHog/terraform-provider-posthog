package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/posthog/terraform-provider/pkg/util"
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

	assert.Equal(t, "hf-1,hf-2,hf-3", hogFunctionGroupID(responseOrder))
	assert.Equal(t, hogFunctionGroupID([]string{"hf-3", "hf-1", "hf-2"}), hogFunctionGroupID(responseOrder))
	assert.Equal(t, []string{"hf-2", "hf-1", "hf-3"}, responseOrder, "building the id must not reorder the response")
}

func TestLogsAlertDestinationID_RoundTripsThroughHogFunctionIDsFromState(t *testing.T) {
	hogFunctionIDs := []string{"019dbe94-cec8-781b-9470-4a970cd69ebf", "019dbe94-cec8-7000-8000-000000000001"}

	model := LogsAlertDestinationTFModel{}
	require.NoError(t, model.SetID(hogFunctionGroupID(hogFunctionIDs)))

	assert.ElementsMatch(t, hogFunctionIDs, hogFunctionIDsFromID(model))
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

func TestLogsAlertDestinationMapResponseToModel_PreservesTheConfiguredWebhookURL(t *testing.T) {
	tests := map[string]struct {
		resp httpclient.LogsAlertDestination
	}{
		"read response redacts the url": {
			resp: httpclient.LogsAlertDestination{
				HogFunctionIDs:     []string{"hf-1"},
				Type:               destinationTypeWebhook,
				RedactedWebhookURL: util.StringPtr("https://example.com/…"),
			},
		},
		"create response omits the type": {
			resp: httpclient.LogsAlertDestination{
				HogFunctionIDs: []string{"hf-1"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			model := LogsAlertDestinationTFModel{
				Type:       types.StringValue(destinationTypeWebhook),
				WebhookURL: types.StringValue("https://example.com/configured"),
			}

			diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), test.resp, &model)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, types.StringValue("https://example.com/configured"), model.WebhookURL)
		})
	}
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
	assert.Contains(t, diags.Errors()[0].Summary(), "no id")
}

func TestLogsAlertDestinationID_PrefersTheRowIDAndFallsBackToTheHogFunctionGroup(t *testing.T) {
	tests := map[string]struct {
		destination httpclient.LogsAlertDestination
		wantID      string
	}{
		"a destination with a row is tracked by it": {
			destination: httpclient.LogsAlertDestination{ID: util.StringPtr("dest-1"), HogFunctionIDs: []string{"hf-2", "hf-1"}},
			wantID:      "dest-1",
		},
		"a row-only destination is tracked by its row": {
			destination: httpclient.LogsAlertDestination{ID: util.StringPtr("dest-1")},
			wantID:      "dest-1",
		},
		"a destination with no row keeps the hog function group id": {
			destination: httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-2", "hf-1"}},
			wantID:      "hf-1,hf-2",
		},
		"an empty row id counts as none": {
			destination: httpclient.LogsAlertDestination{ID: util.StringPtr(""), HogFunctionIDs: []string{"hf-1"}},
			wantID:      "hf-1",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			model := slackDestination()
			diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), test.destination, &model)

			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, types.StringValue(test.wantID), model.ID)
		})
	}
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

func destinationInState(id string) LogsAlertDestinationTFModel {
	return LogsAlertDestinationTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(id)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testLogsAlertDestinationProjectID)},
		AlertID:                types.StringValue(testLogsAlertDestinationAlertID),
	}
}

// Destinations are read off the alert itself; PostHog exposes no endpoint that lists
// them on their own.
func writeAlertWithDestinations(t *testing.T, w http.ResponseWriter, destinations ...httpclient.LogsAlertDestination) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(httpclient.LogsAlert{
		ID:           testLogsAlertDestinationAlertID,
		Destinations: destinations,
	}))
}

func TestLogsAlertDestinationRead_FindsTheGroupAmongTheAlertsDestinations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, logsAlertPath(), r.URL.Path)
		writeAlertWithDestinations(t, w,
			httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-1"}, Type: destinationTypeWebhook},
			httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-2", "hf-3"}, Type: destinationTypeTeams})
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
		require.Equal(t, logsAlertPath(), r.URL.Path)
		writeAlertWithDestinations(t, w,
			httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-9"}, Type: destinationTypeWebhook})
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

	_, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, destinationInState("hf-1"))

	require.Error(t, err)
	assert.Equal(t, http.StatusNotFound, int(status), "a destination gone from a live alert must be removed from state")
}

func TestLogsAlertDestinationRead_RefusesAGroupSplitAcrossConfigurations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, logsAlertPath(), r.URL.Path)
		writeAlertWithDestinations(t, w,
			httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-1"}, Type: destinationTypeWebhook},
			httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-2"}, Type: destinationTypeWebhook})
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	_, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, destinationInState("hf-1,hf-2"))

	require.ErrorContains(t, err, "split across 2 server groups")
	assert.NotEqual(t, http.StatusNotFound, int(status))
}

func TestLogsAlertDestinationRead_RefusesAGroupWidenedWithExtraIDs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		writeAlertWithDestinations(t, w,
			httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-1", "hf-2"}, Type: destinationTypeWebhook})
	}))
	defer server.Close()

	model := destinationInState("hf-1")
	model.Type = types.StringValue(destinationTypeWebhook)
	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	_, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, model)

	require.ErrorContains(t, err, "partially present")
	assert.NotEqual(t, http.StatusNotFound, int(status))
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

	assert.Equal(t, []string{"hf-1", "hf-2", "hf-3"}, hogFunctionIDsFromID(model))
}

func TestSharesHogFunction(t *testing.T) {
	group := []string{"hf-1", "hf-2", "hf-3", "hf-4"}

	assert.True(t, sharesHogFunction(group, group), "the same group matches itself")
	assert.True(t, sharesHogFunction(group, []string{"hf-3"}), "one shared id identifies the group, which is what lets an import name a single id")
	assert.False(t, sharesHogFunction(group, []string{"hf-9"}), "a different destination must not match")
	assert.False(t, sharesHogFunction(group, nil), "nothing in state matches nothing")
}

// stateAfterApply is what this provider records: the id and, beside it, the hog functions
// the destination had, if any.
func stateAfterApply(t *testing.T, id string, hogFunctionIDs ...string) LogsAlertDestinationTFModel {
	t.Helper()
	model := destinationInState(id)
	model.Type = types.StringValue(destinationTypeWebhook)
	model.HogFunctionIDs = stringSet(t, hogFunctionIDs...)
	return model
}

func importedState(id string) LogsAlertDestinationTFModel {
	model := destinationInState(id)
	model.HogFunctionIDs = types.SetNull(types.StringType)
	return model
}

// Read, then map, as a refresh does, so each case shows the id the state ends up with.
func TestLogsAlertDestinationRefresh_MatchesByRowIDThenHogFunctions(t *testing.T) {
	converted := httpclient.LogsAlertDestination{ID: util.StringPtr("dest-1"), HogFunctionIDs: []string{"hf-1", "hf-2"}, Type: destinationTypeWebhook}
	rowOnly := httpclient.LogsAlertDestination{ID: util.StringPtr("dest-2"), Type: destinationTypeWebhook}
	noRow := httpclient.LogsAlertDestination{HogFunctionIDs: []string{"hf-3", "hf-4"}, Type: destinationTypeWebhook}
	replacement := httpclient.LogsAlertDestination{ID: util.StringPtr("dest-new"), HogFunctionIDs: []string{"hf-1", "hf-2"}, Type: destinationTypeWebhook}

	tests := map[string]struct {
		server  []httpclient.LogsAlertDestination
		state   LogsAlertDestinationTFModel
		wantID  string
		missing bool
	}{
		"a row id finds its row": {
			server: []httpclient.LogsAlertDestination{noRow, converted},
			state:  stateAfterApply(t, "dest-1", "hf-1", "hf-2"),
			wantID: "dest-1",
		},
		"a row-only destination finds its row": {
			server: []httpclient.LogsAlertDestination{converted, rowOnly},
			state:  stateAfterApply(t, "dest-2"),
			wantID: "dest-2",
		},
		"state from an older release moves to the row id": {
			server: []httpclient.LogsAlertDestination{converted},
			state:  stateAfterApply(t, "hf-1,hf-2", "hf-1", "hf-2"),
			wantID: "dest-1",
		},
		"state from an older release keeps its id while the destination has no row": {
			server: []httpclient.LogsAlertDestination{noRow},
			state:  stateAfterApply(t, "hf-3,hf-4", "hf-3", "hf-4"),
			wantID: "hf-3,hf-4",
		},
		"a replaced row is found by the hog functions recorded beside its id": {
			server: []httpclient.LogsAlertDestination{replacement},
			state:  stateAfterApply(t, "dest-old", "hf-1", "hf-2"),
			wantID: "dest-new",
		},
		"an import by row id": {
			server: []httpclient.LogsAlertDestination{converted, rowOnly},
			state:  importedState("dest-2"),
			wantID: "dest-2",
		},
		"an import by one hog function id moves to the row id": {
			server: []httpclient.LogsAlertDestination{converted},
			state:  importedState("hf-2"),
			wantID: "dest-1",
		},
		"a deleted row-only destination is missing": {
			server:  []httpclient.LogsAlertDestination{converted},
			state:   stateAfterApply(t, "dest-2"),
			missing: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				writeAlertWithDestinations(t, w, test.server...)
			}))
			defer server.Close()
			client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

			model := test.state
			destination, status, err := LogsAlertDestinationOps{}.Read(context.Background(), client, model)
			if test.missing {
				require.Error(t, err)
				assert.Equal(t, http.StatusNotFound, int(status))
				return
			}
			require.NoError(t, err)

			diags := LogsAlertDestinationOps{}.MapResponseToModel(context.Background(), destination, &model)
			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			assert.Equal(t, test.wantID, model.GetID())
		})
	}
}

func TestLogsAlertDestinationDelete_NamesARowByItsIDAndAGroupByItsHogFunctions(t *testing.T) {
	tests := map[string]struct {
		state    LogsAlertDestinationTFModel
		wantBody map[string]any
	}{
		"a row id": {
			state:    stateAfterApply(t, "dest-1", "hf-1", "hf-2"),
			wantBody: map[string]any{"destination_ids": []any{"dest-1"}},
		},
		"a row-only destination": {
			state:    stateAfterApply(t, "dest-2"),
			wantBody: map[string]any{"destination_ids": []any{"dest-2"}},
		},
		"a group from an older release": {
			state:    stateAfterApply(t, "hf-1,hf-2", "hf-1", "hf-2"),
			wantBody: map[string]any{"hog_function_ids": []any{"hf-1", "hf-2"}},
		},
		"a one-function group from an older release, which looks like a row id": {
			state:    stateAfterApply(t, "hf-1", "hf-1"),
			wantBody: map[string]any{"hog_function_ids": []any{"hf-1"}},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var body map[string]any
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, logsAlertDestinationsPath()+"delete", r.URL.Path)
				require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
				w.WriteHeader(http.StatusNoContent)
			}))
			defer server.Close()
			client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

			_, err := LogsAlertDestinationOps{}.Delete(context.Background(), client, test.state)

			require.NoError(t, err)
			assert.Equal(t, test.wantBody, body)
		})
	}
}
