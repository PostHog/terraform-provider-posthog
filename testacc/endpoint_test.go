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
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/posthog/terraform-provider/pkg/httpclient"
)

func testAccCheckEndpointDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_endpoint" {
			continue
		}

		_, status, err := client.GetEndpoint(context.Background(), projectID, rs.Primary.ID)
		if err != nil {
			if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
				continue
			}
			return fmt.Errorf("unexpected error checking endpoint %s: %w", rs.Primary.ID, err)
		}
		return fmt.Errorf("endpoint %s still exists", rs.Primary.ID)
	}

	return nil
}

// endpointTestName returns a name PostHog accepts: it must start with a letter
// and hold only letters, digits, hyphens and underscores.
func endpointTestName() string {
	return acctest.RandomWithPrefix("tf-acc-test")
}

func TestEndpoint_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := endpointTestName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointBasic(rName, "SELECT count() FROM events"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_endpoint.test", "id", rName),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "current_version", "1"),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "is_active", "true"),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "is_materialized", "false"),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "data_freshness_seconds", "86400"),
					resource.TestCheckResourceAttrSet("posthog_endpoint.test", "run_url"),
					resource.TestCheckResourceAttrSet("posthog_endpoint.test", "created_at"),
				),
			},
			{
				Config:   testAccEndpointBasic(rName, "SELECT count() FROM events"),
				PlanOnly: true,
			},
			{
				ResourceName:      "posthog_endpoint.test",
				ImportState:       true,
				ImportStateId:     rName,
				ImportStateVerify: true,
				// Without a configuration to filter by, import keeps every field
				// PostHog stored in the query, so it is checked by the plan below instead.
				ImportStateVerifyIgnore: []string{"query_json"},
			},
		},
	})
}

// TestEndpoint_QueryChangeCreatesVersion pins that a changed query updates the
// endpoint in place and PostHog saves it as the next version.
func TestEndpoint_QueryChangeCreatesVersion(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := endpointTestName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointBasic(rName, "SELECT count() FROM events"),
				Check:  resource.TestCheckResourceAttr("posthog_endpoint.test", "current_version", "1"),
			},
			{
				Config: testAccEndpointBasic(rName, "SELECT count() FROM events WHERE event = '$pageview'"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("posthog_endpoint.test", plancheck.ResourceActionUpdate),
						plancheck.ExpectUnknownValue("posthog_endpoint.test", tfjsonpath.New("current_version")),
					},
				},
				Check: resource.TestCheckResourceAttr("posthog_endpoint.test", "current_version", "2"),
			},
		},
	})
}

func TestEndpoint_RenameReplaces(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := endpointTestName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointBasic(rName, "SELECT 1"),
			},
			{
				Config: testAccEndpointBasic(rName+"-renamed", "SELECT 1"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("posthog_endpoint.test", plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.TestCheckResourceAttr("posthog_endpoint.test", "id", rName+"-renamed"),
			},
		},
	})
}

// TestEndpoint_OptionalFields sets and then removes every optional attribute,
// so removing one clears it in PostHog instead of leaving the old value.
func TestEndpoint_OptionalFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := endpointTestName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEndpointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_endpoint.test", "description", "Events in the last day"),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "data_freshness_seconds", "3600"),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "is_active", "false"),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "tags.#", "1"),
					resource.TestCheckTypeSetElemAttr("posthog_endpoint.test", "tags.*", "tf-acc"),
				),
			},
			{
				Config:   testAccEndpointAllFields(rName),
				PlanOnly: true,
			},
			{
				Config: testAccEndpointBasic(rName, "SELECT count() FROM events WHERE timestamp >= now() - INTERVAL 1 DAY"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("posthog_endpoint.test", "description"),
					resource.TestCheckNoResourceAttr("posthog_endpoint.test", "tags.#"),
					// Optional+Computed attributes keep their last value when removed.
					resource.TestCheckResourceAttr("posthog_endpoint.test", "data_freshness_seconds", "3600"),
					resource.TestCheckResourceAttr("posthog_endpoint.test", "is_active", "false"),
				),
			},
			{
				Config:   testAccEndpointBasic(rName, "SELECT count() FROM events WHERE timestamp >= now() - INTERVAL 1 DAY"),
				PlanOnly: true,
			},
		},
	})
}

// TestEndpoint_Variables pins the documented variable wiring: the query names
// the variable by code name and declares it keyed by its UUID. PostHog adds the
// variable's value to the stored query, which must not show as drift.
func TestEndpoint_Variables(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := endpointTestName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			if err := testAccCheckEndpointDestroy(s); err != nil {
				return err
			}
			return testAccCheckInsightVariableDestroy(s)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointWithVariable(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_endpoint.test", "current_version", "1"),
					resource.TestCheckResourceAttrSet("posthog_endpoint.test", "query_json"),
				),
			},
			{
				Config:   testAccEndpointWithVariable(rName),
				PlanOnly: true,
			},
		},
	})
}

func testAccEndpointBasic(name, sql string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_endpoint" "test" {
  name = %[1]q
  query_json = jsonencode({
    kind  = "HogQLQuery"
    query = %[2]q
  })
}
`, name, sql)
}

func testAccEndpointAllFields(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_endpoint" "test" {
  name                   = %[1]q
  description            = "Events in the last day"
  data_freshness_seconds = 3600
  is_active              = false
  tags                   = ["tf-acc"]
  query_json = jsonencode({
    kind  = "HogQLQuery"
    query = "SELECT count() FROM events WHERE timestamp >= now() - INTERVAL 1 DAY"
  })
}
`, name)
}

func testAccEndpointWithVariable(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight_variable" "event_name" {
  name               = "%[1]s event"
  type               = "String"
  default_value_json = jsonencode("$pageview")
}

resource "posthog_endpoint" "test" {
  name = %[1]q
  query_json = jsonencode({
    kind  = "HogQLQuery"
    query = "SELECT count() FROM events WHERE event = {variables.${posthog_insight_variable.event_name.code_name}}"
    variables = {
      (posthog_insight_variable.event_name.id) = {
        variableId = posthog_insight_variable.event_name.id
        code_name  = posthog_insight_variable.event_name.code_name
      }
    }
  })
}
`, name)
}
