package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAction_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccActionBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_action.test", "id"),
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
		Steps: []resource.TestStep{
			{
				Config: testAccActionAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_action.test", "description", "Test action description"),
					resource.TestCheckResourceAttr("posthog_action.test", "post_to_slack", "false"),
					resource.TestCheckResourceAttr("posthog_action.test", "tags.#", "1"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "id"),
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
		Steps: []resource.TestStep{
			{
				Config: testAccActionBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName),
				),
			},
			{
				Config: testAccActionUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_action.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_action.test", "description", "Updated description"),
					resource.TestCheckResourceAttrSet("posthog_action.test", "steps_json"),
				),
			},
		},
	})
}

func TestAction_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
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
  tags           = ["analytics"]
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
