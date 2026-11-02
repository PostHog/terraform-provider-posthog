package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type projectSettingsPlanCheck func(context.Context, plancheck.CheckPlanRequest, *plancheck.CheckPlanResponse)

func (check projectSettingsPlanCheck) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	check(ctx, req, resp)
}

func TestProjectSettingsSavedPlanKeepsOmittedSettings(t *testing.T) {
	var mu sync.Mutex
	settings := map[string]any{
		"id":                            123,
		"heatmaps_opt_in":               true,
		"autocapture_exceptions_opt_in": true,
		"session_recording_opt_in":      true,
		"surveys_opt_in":                true,
		"cookieless_server_hash_mode":   2,
		"autocapture_web_vitals_opt_in": true,
		"capture_performance_opt_in":    true,
		"anonymize_ips":                 false,
		"session_recording_network_payload_capture_config": map[string]any{"recordHeaders": true, "recordBody": true},
		"app_urls":                             []string{"https://example.com"},
		"recording_domains":                    []string{"https://example.com"},
		"test_account_filters":                 []any{map[string]any{"key": "email", "type": "person", "value": "@example.com", "operator": "icontains"}},
		"test_account_filters_default_checked": true,
	}
	var patches []map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/environments/123/" || r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		switch r.Method {
		case http.MethodGet:
		case http.MethodPatch:
			var patch map[string]any
			if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
				t.Errorf("decode PATCH: %s", err)
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			patches = append(patches, patch)
			for key, value := range patch {
				settings[key] = value
			}
		default:
			t.Errorf("unexpected method: %s", r.Method)
			http.Error(w, "unexpected method", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(settings); err != nil {
			t.Errorf("encode settings: %s", err)
		}
	}))
	defer server.Close()

	config := fmt.Sprintf(`
provider "posthog" {
  host = %q
  api_key = "test-api-key"
  project_id = "123"
  organization_id = ""
}
resource "posthog_project_settings" "test" {
  %%s
}
`, server.URL)
	const address = "posthog_project_settings.test"
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"posthog": providerserver.NewProtocol6WithError(New("test")()),
		},
		Steps: []resource.TestStep{
			{
				Config:             fmt.Sprintf(config, ""),
				ResourceName:       address,
				ImportState:        true,
				ImportStateId:      "123",
				ImportStatePersist: true,
			},
			{
				Config: fmt.Sprintf(config, "anonymize_ips = true\n  surveys_opt_in = true"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						projectSettingsPlanCheck(func(_ context.Context, _ plancheck.CheckPlanRequest, _ *plancheck.CheckPlanResponse) {
							// The harness has saved the plan and applies that same file next.
							mu.Lock()
							defer mu.Unlock()
							settings["session_recording_opt_in"] = false
							settings["session_recording_network_payload_capture_config"] = map[string]any{"recordHeaders": false, "recordBody": true}
						}),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(address, "session_recording_opt_in", "false"),
					resource.TestCheckResourceAttr(address, "session_recording_network_payload_capture_config.record_headers", "false"),
				),
				PostApplyFunc: func() {
					mu.Lock()
					defer mu.Unlock()
					require.Len(t, patches, 1)
					assert.Equal(t, true, patches[0]["anonymize_ips"])
					assert.Equal(t, true, patches[0]["surveys_opt_in"])
				},
			},
		},
	})
}
