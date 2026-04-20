package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

const testSurveyResourceName = "posthog_survey.test"

func testAccCheckSurveyDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_survey" {
			continue
		}

		_, status, err := client.GetSurvey(context.Background(), projectID, rs.Primary.ID)
		if err != nil {
			if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
				continue
			}
			return fmt.Errorf("unexpected error checking survey %s: %w", rs.Primary.ID, err)
		}

		return fmt.Errorf("survey %s still exists", rs.Primary.ID)
	}

	return nil
}

func TestSurvey_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-survey")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSurveyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSurveyBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testSurveyResourceName, "name", rName),
					resource.TestCheckResourceAttr(testSurveyResourceName, "type", "popover"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "schedule", "once"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "questions_json", `[{"question":"How satisfied are you?","type":"open"}]`),
					resource.TestCheckResourceAttrSet(testSurveyResourceName, "id"),
					resource.TestCheckResourceAttrSet(testSurveyResourceName, "created_at"),
				),
			},
			{
				Config: testAccSurveyUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testSurveyResourceName, "description", "Updated description"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "enable_partial_responses", "true"),
				),
			},
		},
	})
}

func TestSurvey_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-survey")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSurveyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSurveyBasic(rName),
			},
			{
				ResourceName:            testSurveyResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"linked_insight_id", "targeting_flag_filters_json"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[testSurveyResourceName]
					if !ok {
						return "", fmt.Errorf("resource not found: %s", testSurveyResourceName)
					}
					return fmt.Sprintf("%s/%s", projectID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccSurveyBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_survey" "test" {
  name  = %q
  type  = "popover"
  questions_json = jsonencode([
    {
      type     = "open"
      question = "How satisfied are you?"
    }
  ])
}
`, name)
}

func testAccSurveyUpdated(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_survey" "test" {
  name                     = %q
  type                     = "popover"
  description              = "Updated description"
  enable_partial_responses = true

  questions_json = jsonencode([
    {
      type     = "open"
      question = "How satisfied are you?"
    }
  ])
}
`, name)
}
