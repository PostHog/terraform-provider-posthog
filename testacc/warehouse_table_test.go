package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

const warehouseTableAddress = "posthog_warehouse_table.test"

// Required env vars for warehouse table tests:
//
//	POSTHOG_TEST_S3_URL_PATTERN    HTTPS URL of one Parquet file, e.g.
//	                               https://<bucket>.s3.<region>.amazonaws.com/tf-acc/orders.parquet
//	POSTHOG_TEST_S3_ACCESS_KEY     key that can list the bucket and read the file
//	POSTHOG_TEST_S3_ACCESS_SECRET
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

func testAccWarehouseTableConfig(name, urlPattern, accessKey, accessSecret string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_warehouse_table" "test" {
  name          = %q
  format        = "Parquet"
  url_pattern   = %q
  access_key    = %q
  access_secret = %q
}
`, name, urlPattern, accessKey, accessSecret)
}

// wildcardOf swaps the file name at the end of a URL for `*.parquet`, which
// still matches the test file but is a different url_pattern.
func wildcardOf(urlPattern string) string {
	return urlPattern[:strings.LastIndex(urlPattern, "/")+1] + "*.parquet"
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

func TestWarehouseTable_Lifecycle(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoS3Creds(t)

	name := randomTestPrefix()
	urlPattern := os.Getenv("POSTHOG_TEST_S3_URL_PATTERN")
	accessKey := os.Getenv("POSTHOG_TEST_S3_ACCESS_KEY")
	accessSecret := os.Getenv("POSTHOG_TEST_S3_ACCESS_SECRET")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWarehouseTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWarehouseTableConfig(name, urlPattern, accessKey, accessSecret),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(warehouseTableAddress, "id"),
					resource.TestCheckResourceAttrSet(warehouseTableAddress, "created_at"),
					resource.TestCheckResourceAttr(warehouseTableAddress, "name", name),
					resource.TestCheckResourceAttr(warehouseTableAddress, "format", "Parquet"),
					resource.TestCheckResourceAttr(warehouseTableAddress, "url_pattern", urlPattern),
				),
			},
			{
				Config: testAccWarehouseTableConfig(name+"_renamed", urlPattern, accessKey, accessSecret),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(warehouseTableAddress, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.TestCheckResourceAttr(warehouseTableAddress, "name", name+"_renamed"),
			},
			{
				Config: testAccWarehouseTableConfig(name+"_renamed", wildcardOf(urlPattern), accessKey, accessSecret),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(warehouseTableAddress, plancheck.ResourceActionReplace),
					},
				},
				Check: resource.TestCheckResourceAttr(warehouseTableAddress, "url_pattern", wildcardOf(urlPattern)),
			},
			{
				ResourceName:      warehouseTableAddress,
				ImportState:       true,
				ImportStateVerify: true,
				// PostHog never returns the credential.
				ImportStateVerifyIgnore: []string{"access_key", "access_secret"},
			},
		},
	})
}

// PostHog reads the files on create, so a key that cannot read them fails the
// apply instead of leaving behind a table no query can use.
func TestWarehouseTable_UnreadableFilesFailTheApply(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoS3Creds(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWarehouseTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWarehouseTableConfig(
					randomTestPrefix(),
					os.Getenv("POSTHOG_TEST_S3_URL_PATTERN"),
					"AKIAIOSFODNN7EXAMPLE",
					"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				),
				ExpectError: regexp.MustCompile(`Could\s+not\s+read\s+the\s+files`),
			},
		},
	})
}
