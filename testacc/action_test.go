package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/pkg/httpclient"
)

func testAccCheckActionDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_action" {
			continue
		}

		action, status, err := client.GetAction(context.Background(), projectID, rs.Primary.ID)
		if err != nil {
			if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
				continue
			}
			return fmt.Errorf("unexpected error checking action %s: %w", rs.Primary.ID, err)
		}
		if action.Deleted == nil || !*action.Deleted {
			return fmt.Errorf("action %s still exists and is not soft-deleted", rs.Primary.ID)
		}
	}

	return nil
}

func TestAction_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccActionBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_action.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "created_at"),
					resource.TestCheckResourceAttr("posthog_action.test", "deleted", "false"),
				),
			},
		},
	})
}

func TestAction_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccActionAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_action.test", "description", "Test action description"),
					resource.TestCheckResourceAttr("posthog_action.test", "post_to_slack", "false"),
					resource.TestCheckResourceAttr("posthog_action.test", "tags.#", "2"),
					resource.TestCheckTypeSetElemAttr("posthog_action.test", "tags.*", "analytics"),
					resource.TestCheckTypeSetElemAttr("posthog_action.test", "tags.*", "tracking"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "created_at"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "steps_json"),
				),
			},
		},
	})
}

func TestAction_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckActionDestroy,
		Steps: []resource.TestStep{
			// Create with all fields
			{
				Config: testAccActionAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_action.test", "description", "Test action description"),
					resource.TestCheckResourceAttr("posthog_action.test", "tags.#", "2"),
				),
			},
			// Update name, description, and steps
			{
				Config: testAccActionUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_action.test", "description", "Updated description"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "steps_json"),
				),
			},
			// Clear optional fields by removing them
			{
				Config: testAccActionBasic(rName + "-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName+"-updated"),
				),
			},
		},
	})
}

// TestAction_Tags tests creating, updating, and reducing tags.
func TestAction_Tags(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckActionDestroy,
		Steps: []resource.TestStep{
			// Create with two tags
			{
				Config: testAccActionWithTags(rName, []string{"tag1", "tag2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "tags.#", "2"),
					resource.TestCheckTypeSetElemAttr("posthog_action.test", "tags.*", "tag1"),
					resource.TestCheckTypeSetElemAttr("posthog_action.test", "tags.*", "tag2"),
				),
			},
			// Add a third tag
			{
				Config: testAccActionWithTags(rName, []string{"tag1", "tag2", "tag3"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "tags.#", "3"),
					resource.TestCheckTypeSetElemAttr("posthog_action.test", "tags.*", "tag3"),
				),
			},
			// Reduce to one tag
			{
				Config: testAccActionWithTags(rName, []string{"tag1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "tags.#", "1"),
					resource.TestCheckTypeSetElemAttr("posthog_action.test", "tags.*", "tag1"),
				),
			},
		},
	})
}

// TestAction_ClearTagsAndSteps verifies that tags and steps_json can be cleared.
// Tags (Optional-only) can be cleared by removing them from config.
// Steps (Optional+Computed) must be explicitly set to an empty array — removing
// the attribute preserves the prior value due to UseStateForUnknown.
func TestAction_ClearTagsAndSteps(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckActionDestroy,
		Steps: []resource.TestStep{
			// Create with tags and steps
			{
				Config: testAccActionAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "steps_json"),
				),
			},
			// Clear: remove tags from config, explicitly empty steps
			{
				Config: testAccActionCleared(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
					resource.TestCheckNoResourceAttr("posthog_action.test", "tags.#"),
					resource.TestCheckResourceAttr("posthog_action.test", "steps_json", "[]"),
				),
			},
		},
	})
}

// TestAction_StepsJSONFormatting verifies that steps_json is compared by content, so
// JSON that is not formatted like jsonencode output applies and plans clean, while a
// changed step still shows as a change.
func TestAction_StepsJSONFormatting(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccActionStepsJSONHeredoc(rName, "https://example.com/"),
			},
			{
				Config:   testAccActionStepsJSONHeredoc(rName, "https://example.com/"),
				PlanOnly: true,
			},
			{
				Config: testAccActionStepsJSONHeredoc(rName, "https://example.com/changed"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("posthog_action.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: testAccCheckActionStepURL("https://example.com/changed"),
			},
		},
	})
}

// testAccCheckActionStepURL checks the URL of the action's first step in PostHog.
func testAccCheckActionStepURL(want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["posthog_action.test"]
		if !ok {
			return fmt.Errorf("resource not found: posthog_action.test")
		}
		client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "test")
		action, _, err := client.GetAction(context.Background(), os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID)
		if err != nil {
			return err
		}
		if len(action.Steps) != 1 || action.Steps[0]["url"] != want {
			return fmt.Errorf("expected one step with url %q, got %v", want, action.Steps)
		}
		return nil
	}
}

func TestAction_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccActionBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
				),
			},
			{
				ResourceName:            "posthog_action.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"steps_json"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["posthog_action.test"]
					if !ok {
						return "", fmt.Errorf("resource not found: posthog_action.test")
					}
					return fmt.Sprintf("%s/%s", projectID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccActionBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_action" "test" {
  name = %q
}
`, name)
}

func testAccActionAllFields(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_action" "test" {
  name           = %q
  description    = "Test action description"
  tags           = ["analytics", "tracking"]
  post_to_slack  = false

  steps_json = jsonencode([
    {
      event        = "$pageview"
      url          = "https://example.com/"
      url_matching = "exact"
    }
  ])
}
`, name)
}

func testAccActionStepsJSONHeredoc(name, url string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_action" "test" {
  name = %q

  steps_json = <<-EOT
    [
      {
        "url_matching": "exact",
        "url": %q,
        "event": "$pageview"
      }
    ]
  EOT
}
`, name, url)
}

func testAccActionUpdated(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_action" "test" {
  name        = %q
  description = "Updated description"

  steps_json = jsonencode([
    {
      event = "$autocapture"
    }
  ])
}
`, name+"-updated")
}

func testAccActionCleared(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_action" "test" {
  name       = %q
  steps_json = jsonencode([])
}
`, name)
}

func testAccActionWithTags(name string, tags []string) string {
	tagsStr := ""
	for i, tag := range tags {
		if i > 0 {
			tagsStr += ", "
		}
		tagsStr += fmt.Sprintf("%q", tag)
	}

	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_action" "test" {
  name = %q
  tags = [%s]
}
`, name, tagsStr)
}
