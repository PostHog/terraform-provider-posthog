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

func testAccCheckSchemaPropertyGroupDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_schema_property_group" {
			continue
		}

		_, status, err := client.GetSchemaPropertyGroup(context.Background(), projectID, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("schema property group %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

func testAccSchemaPropertyGroupBasic(rName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name        = %[1]q
  description = "acceptance test group"

  properties = [
    {
      name          = "cart_value"
      property_type = "Numeric"
      is_required   = true
    },
    {
      name          = "currency"
      property_type = "String"
    },
  ]
}
`, rName)
}

func testAccSchemaPropertyGroupUpdated(rName string) string {
	return fmt.Sprintf(`
resource "posthog_schema_property_group" "test" {
  name = %[1]q

  properties = [
    {
      name          = "cart_value"
      property_type = "String"
    },
    {
      name          = "coupon_code"
      property_type = "String"
      description   = "applied coupon"
    },
  ]
}
`, rName)
}

func TestSchemaPropertyGroup_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSchemaPropertyGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSchemaPropertyGroupBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "description", "acceptance test group"),
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "properties.#", "2"),
					resource.TestCheckResourceAttrSet("posthog_schema_property_group.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_schema_property_group.test", "created_at"),
				),
			},
			{
				// retype cart_value, drop currency, add coupon_code, clear description
				Config: testAccSchemaPropertyGroupUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_schema_property_group.test", "properties.#", "2"),
					resource.TestCheckNoResourceAttr("posthog_schema_property_group.test", "description"),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_schema_property_group.test", "properties.*", map[string]string{
						"name":          "cart_value",
						"property_type": "String",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("posthog_schema_property_group.test", "properties.*", map[string]string{
						"name":        "coupon_code",
						"description": "applied coupon",
					}),
				),
			},
			{
				ResourceName:      "posthog_schema_property_group.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs := s.RootModule().Resources["posthog_schema_property_group.test"]
					return fmt.Sprintf("%s/%s", os.Getenv("POSTHOG_PROJECT_ID"), rs.Primary.ID), nil
				},
			},
		},
	})
}
