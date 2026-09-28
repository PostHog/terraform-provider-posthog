package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// Required env vars for S3 warehouse table tests:
//
//	POSTHOG_TEST_S3_URL_PATTERN    (e.g. https://bucket.s3.amazonaws.com/data/*.parquet)
//	POSTHOG_TEST_S3_ACCESS_KEY
//	POSTHOG_TEST_S3_ACCESS_SECRET
//	POSTHOG_TEST_S3_FORMAT         (optional, defaults to Parquet)
func skipIfNoS3Creds(t *testing.T) {
	t.Helper()
	for _, v := range []string{
		"POSTHOG_TEST_S3_URL_PATTERN",
		"POSTHOG_TEST_S3_ACCESS_KEY",
		"POSTHOG_TEST_S3_ACCESS_SECRET",
	} {
		if os.Getenv(v) == "" {
			t.Skipf("Skipping test: %s not set", v)
		}
	}
}

func testAccCheckWarehouseTableDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_warehouse_table" {
			continue
		}

		_, status, err := client.GetWarehouseTable(context.Background(), projectID, rs.Primary.ID)
		if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
			continue
		}
		if err != nil {
			return fmt.Errorf("unexpected error checking warehouse table %s: %w", rs.Primary.ID, err)
		}
		return fmt.Errorf("warehouse table %s still exists", rs.Primary.ID)
	}

	return nil
}

// TestWarehouseTable_InvalidFormat verifies that an unknown format is rejected
// client-side without any API call. Requires no bucket credentials.
func TestWarehouseTable_InvalidFormat(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
provider "posthog" {}

resource "posthog_warehouse_table" "test" {
  name          = "tf_acc_invalid"
  format        = "Excel"
  url_pattern   = "https://bucket.s3.amazonaws.com/*.xlsx"
  access_key    = "key"
  access_secret = "secret"
}
`,
				ExpectError: regexp.MustCompile(`(?i)value must be one of`),
			},
		},
	})
}

func TestWarehouseTable_S3_Basic(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoS3Creds(t)

	name := "tf_acc_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	renamed := name + "_v2"
	format := os.Getenv("POSTHOG_TEST_S3_FORMAT")
	if format == "" {
		format = "Parquet"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWarehouseTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWarehouseTableS3Config(name, format),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_warehouse_table.test", "id"),
					resource.TestCheckResourceAttr("posthog_warehouse_table.test", "name", name),
					resource.TestCheckResourceAttr("posthog_warehouse_table.test", "format", format),
					resource.TestCheckResourceAttrSet("posthog_warehouse_table.test", "hogql_name"),
				),
			},
			{
				Config: testAccWarehouseTableS3Config(renamed, format),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("posthog_warehouse_table.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.TestCheckResourceAttr("posthog_warehouse_table.test", "name", renamed),
			},
			{
				ResourceName:            "posthog_warehouse_table.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"access_key", "access_secret"},
			},
		},
	})
}

func testAccWarehouseTableS3Config(name, format string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_warehouse_table" "test" {
  name          = %q
  format        = %q
  url_pattern   = %q
  access_key    = %q
  access_secret = %q
}
`, name, format,
		os.Getenv("POSTHOG_TEST_S3_URL_PATTERN"),
		os.Getenv("POSTHOG_TEST_S3_ACCESS_KEY"),
		os.Getenv("POSTHOG_TEST_S3_ACCESS_SECRET"),
	)
}
