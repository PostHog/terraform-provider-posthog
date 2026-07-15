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

// createTestEventDefinition registers an event definition so the event name
// resolves during the test (schemas can only attach to existing definitions).
func createTestEventDefinition(t *testing.T, name string) {
	t.Helper()

	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	def, err := client.CreateEventDefinition(context.Background(), projectID, name)
	if err != nil {
		t.Fatalf("failed to create test event definition: %v", err)
	}
	t.Cleanup(func() {
		_, _ = client.DeleteEventDefinition(context.Background(), projectID, def.ID)
	})
}

func testAccCheckEventSchemaDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_event_schema" {
			continue
		}

		_, status, err := client.GetEventSchema(context.Background(), projectID, rs.Primary.ID, "")
		if err == nil {
			return fmt.Errorf("event schema %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

func testAccEventSchemaBasic(groupName, eventName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name = %[1]q

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
    },
  ]
}

resource "posthog_event_schema" "test" {
  event             = %[2]q
  property_group_id = posthog_schema_property_group.test.id
}
`, groupName, eventName)
}

func testAccEventSchemaReattached(groupName, eventName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name = %[1]q

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
    },
  ]
}

resource "posthog_schema_property_group" "second" {
  name = "%[1]s-second"

  properties = [
    {
      name          = "currency"
      property_type = "String"
    },
  ]
}

resource "posthog_event_schema" "test" {
  event             = %[2]q
  property_group_id = posthog_schema_property_group.second.id
}
`, groupName, eventName)
}

func TestEventSchema_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	eventName := acctest.RandomWithPrefix("tf-acc-event")
	createTestEventDefinition(t, eventName)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEventSchemaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEventSchemaBasic(rName, eventName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_event_schema.test", "event", eventName),
					resource.TestCheckResourceAttrSet("posthog_event_schema.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_event_schema.test", "event_definition_id"),
					resource.TestCheckResourceAttrPair(
						"posthog_event_schema.test", "property_group_id",
						"posthog_schema_property_group.test", "id",
					),
				),
			},
			{
				Config: testAccEventSchemaReattached(rName, eventName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"posthog_event_schema.test", "property_group_id",
						"posthog_schema_property_group.second", "id",
					),
				),
			},
			{
				ResourceName:      "posthog_event_schema.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs := s.RootModule().Resources["posthog_event_schema.test"]
					return fmt.Sprintf("%s/%s", os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID), nil
				},
			},
		},
	})
}
