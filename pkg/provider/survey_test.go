package provider

import (
	"encoding/json/v2"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Exercise Terraform's plan/state consistency checks with both configured and
// server-generated question IDs, without creating resources in PostHog Cloud.
func TestSurveyConfiguredQuestionID(t *testing.T) {
	var mu sync.Mutex
	var survey map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/projects/123/surveys/" && r.URL.Path != "/api/projects/123/surveys/survey-id/" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		switch r.Method {
		case http.MethodPost, http.MethodPatch:
			var input map[string]any
			if err := json.UnmarshalRead(r.Body, &input); err != nil {
				t.Errorf("decode survey: %s", err)
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			if survey == nil {
				survey = map[string]any{"id": "survey-id", "archived": false}
			}
			for key, value := range input {
				survey[key] = value
			}
			questions, ok := survey["questions"].([]any)
			if !ok {
				t.Error("survey request has no questions array")
				http.Error(w, "missing questions", http.StatusBadRequest)
				return
			}
			for i, item := range questions {
				question := item.(map[string]any)
				if _, configured := question["id"]; !configured {
					question["id"] = fmt.Sprintf("generated-%d", i)
				}
			}
		case http.MethodGet:
			if survey == nil {
				http.NotFound(w, r)
				return
			}
		case http.MethodDelete:
			survey = nil
			w.WriteHeader(http.StatusNoContent)
			return
		default:
			t.Errorf("unexpected method: %s", r.Method)
			http.Error(w, "unexpected method", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.MarshalWrite(w, survey); err != nil {
			t.Errorf("encode survey: %s", err)
		}
	}))
	defer server.Close()

	config := func(question string) string {
		return fmt.Sprintf(`
provider "posthog" {
  host            = %q
  api_key         = "test-api-key"
  project_id      = "123"
  organization_id = ""
}
resource "posthog_survey" "test" {
  name     = "Question id regression"
  type     = "api"
  schedule = "always"
  questions_json = jsonencode([
    { id = "aaaaaaaa-1111-4111-8111-aaaaaaaaaaaa", type = "open", question = %q },
    { type = "open", question = "Second question" }
  ])
}
`, server.URL, question)
	}
	const address = "posthog_survey.test"
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"posthog": providerserver.NewProtocol6WithError(New("test")()),
		},
		Steps: []resource.TestStep{
			{
				Config: config("How is it going?"),
				Check:  resource.TestCheckResourceAttr(address, "questions_json", `[{"id":"aaaaaaaa-1111-4111-8111-aaaaaaaaaaaa","question":"How is it going?","type":"open"},{"question":"Second question","type":"open"}]`),
			},
			{
				Config: config("Updated question"),
				Check:  resource.TestCheckResourceAttr(address, "questions_json", `[{"id":"aaaaaaaa-1111-4111-8111-aaaaaaaaaaaa","question":"Updated question","type":"open"},{"question":"Second question","type":"open"}]`),
			},
			{
				Config:   config("Updated question"),
				PlanOnly: true,
			},
		},
	})
}
