package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

func testAccCheckInsightVariableDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_insight_variable" {
			continue
		}

		_, status, err := client.GetInsightVariable(context.Background(), projectID, rs.Primary.ID)
		if err != nil {
			if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
				continue
			}
			return fmt.Errorf("unexpected error checking insight variable %s: %w", rs.Primary.ID, err)
		}
		return fmt.Errorf("insight variable %s still exists", rs.Primary.ID)
	}

	return nil
}

func TestInsightVariable_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInsightVariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightVariableBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "type", "String"),
					resource.TestCheckResourceAttrSet("posthog_insight_variable.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_insight_variable.test", "created_at"),
					// PostHog derives code_name by dropping everything that is not
					// alphanumeric, a space, or an underscore, turning spaces into
					// underscores, and lowercasing. The random name is hyphenated.
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "code_name", strings.ToLower(strings.ReplaceAll(rName, "-", ""))),
				),
			},
		},
	})
}

func TestInsightVariable_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInsightVariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightVariableList(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "type", "List"),
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "values_json", `["dev","staging","prod"]`),
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "default_value_json", `"prod"`),
				),
			},
		},
	})
}

// TestInsightVariable_Update covers renaming and retyping. code_name is derived
// from the name at creation only, so a rename must not change it — insights
// reference the variable by code_name and would otherwise break.
func TestInsightVariable_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInsightVariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightVariableBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_insight_variable.test", "code_name"),
				),
			},
			{
				Config: testAccInsightVariableRenamed(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "name", rName+"-renamed"),
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "type", "Number"),
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "default_value_json", "30"),
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "code_name", strings.ToLower(strings.ReplaceAll(rName, "-", ""))),
				),
			},
		},
	})
}

// TestInsightVariable_ClearValues verifies that removing default_value_json and
// values_json from configuration clears them server-side. A PATCH that omits a
// field leaves the stored value alone, so the resource has to send null.
func TestInsightVariable_ClearValues(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInsightVariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightVariableList(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_insight_variable.test", "values_json"),
					resource.TestCheckResourceAttrSet("posthog_insight_variable.test", "default_value_json"),
				),
			},
			{
				Config: testAccInsightVariableListCleared(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight_variable.test", "name", rName),
					resource.TestCheckNoResourceAttr("posthog_insight_variable.test", "values_json"),
					resource.TestCheckNoResourceAttr("posthog_insight_variable.test", "default_value_json"),
				),
			},
		},
	})
}

func TestInsightVariable_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInsightVariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInsightVariableList(rName),
			},
			{
				ResourceName:      "posthog_insight_variable.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["posthog_insight_variable.test"]
					if !ok {
						return "", fmt.Errorf("resource not found: posthog_insight_variable.test")
					}
					return fmt.Sprintf("%s/%s", projectID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccInsightVariableBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight_variable" "test" {
  name = %q
  type = "String"
}
`, name)
}

func testAccInsightVariableRenamed(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight_variable" "test" {
  name               = %q
  type               = "Number"
  default_value_json = jsonencode(30)
}
`, name+"-renamed")
}

func testAccInsightVariableList(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight_variable" "test" {
  name               = %q
  type               = "List"
  values_json        = jsonencode(["dev", "staging", "prod"])
  default_value_json = jsonencode("prod")
}
`, name)
}

func testAccInsightVariableListCleared(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight_variable" "test" {
  name = %q
  type = "List"
}
`, name)
}
