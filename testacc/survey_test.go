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

// TestSurvey_ClearNumericLimits verifies that removing the nullable numeric
// attributes from config sends an explicit JSON null on PATCH, which PostHog
// interprets as "clear this column", and that the resulting state reflects the
// cleared values. Without the drop-omitempty fix, step 2 would still observe
// the values from step 1 because the omitempty pointer would never be sent.
func TestSurvey_ClearNumericLimits(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-survey")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSurveyDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with every nullable numeric field set so we have a
			// non-trivial baseline to clear from.
			{
				Config: testAccSurveyAllNumericLimits(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testSurveyResourceName, "schedule", "recurring"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "responses_limit", "100"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "iteration_count", "3"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "iteration_frequency_days", "7"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "response_sampling_interval_type", "week"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "response_sampling_interval", "2"),
					resource.TestCheckResourceAttr(testSurveyResourceName, "response_sampling_limit", "50"),
				),
			},
			// Step 2: drop every nullable numeric field. State must show them as
			// unset, which only happens if PostHog actually cleared the columns —
			// MapResponseToModel reads them straight from the GET response.
			{
				Config: testAccSurveyNoNumericLimits(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(testSurveyResourceName, "responses_limit"),
					resource.TestCheckNoResourceAttr(testSurveyResourceName, "iteration_count"),
					resource.TestCheckNoResourceAttr(testSurveyResourceName, "iteration_frequency_days"),
					resource.TestCheckNoResourceAttr(testSurveyResourceName, "response_sampling_interval"),
					resource.TestCheckNoResourceAttr(testSurveyResourceName, "response_sampling_limit"),
				),
			},
		},
	})
}

// TestSurvey_LinkAndUnlinkFlag is the most user-visible drop-omitempty case:
// linking a feature flag, then removing the linked_flag_id and confirming the
// survey is actually unlinked server-side (the API response no longer
// includes a linked_flag, so int64ValueFromMapOrNull surfaces null).
func TestSurvey_LinkAndUnlinkFlag(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-survey")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSurveyDestroy,
		Steps: []resource.TestStep{
			// Step 1: survey linked to a fresh feature flag.
			{
				Config: testAccSurveyWithLinkedFlag(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						testSurveyResourceName, "linked_flag_id",
						"posthog_feature_flag.test", "id",
					),
				),
			},
			// Step 2: remove linked_flag_id from the survey while keeping the
			// flag resource around. The survey should report no linked flag.
			{
				Config: testAccSurveyWithoutLinkedFlag(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(testSurveyResourceName, "linked_flag_id"),
				),
			},
		},
	})
}

func testAccSurveyAllNumericLimits(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_survey" "test" {
  name                            = %q
  type                            = "popover"
  schedule                        = "recurring"
  responses_limit                 = 100
  iteration_count                 = 3
  iteration_frequency_days        = 7
  response_sampling_interval_type = "week"
  response_sampling_interval      = 2
  response_sampling_limit         = 50

  questions_json = jsonencode([
    {
      type     = "open"
      question = "How satisfied are you?"
    }
  ])
}
`, name)
}

func testAccSurveyNoNumericLimits(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_survey" "test" {
  name = %q
  type = "popover"

  questions_json = jsonencode([
    {
      type     = "open"
      question = "How satisfied are you?"
    }
  ])
}
`, name)
}

func testAccSurveyWithLinkedFlag(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                = %[1]q
  name               = "Test flag for survey link"
  active             = true
  rollout_percentage = 100
}

resource "posthog_survey" "test" {
  name           = %[1]q
  type           = "popover"
  linked_flag_id = posthog_feature_flag.test.id

  questions_json = jsonencode([
    {
      type     = "open"
      question = "How satisfied are you?"
    }
  ])
}
`, name)
}

func testAccSurveyWithoutLinkedFlag(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "test" {
  key                = %[1]q
  name               = "Test flag for survey link"
  active             = true
  rollout_percentage = 100
}

resource "posthog_survey" "test" {
  name = %[1]q
  type = "popover"

  questions_json = jsonencode([
    {
      type     = "open"
      question = "How satisfied are you?"
    }
  ])
}
`, name)
}
