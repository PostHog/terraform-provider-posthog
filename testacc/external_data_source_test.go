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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// Required env vars for Postgres external data source tests:
//
//	POSTHOG_TEST_PG_HOST
//	POSTHOG_TEST_PG_PORT       (optional, defaults to 5432)
//	POSTHOG_TEST_PG_DATABASE
//	POSTHOG_TEST_PG_USER
//	POSTHOG_TEST_PG_PASSWORD
func skipIfNoPostgresCreds(t *testing.T) {
	t.Helper()
	for _, v := range []string{
		"POSTHOG_TEST_PG_HOST",
		"POSTHOG_TEST_PG_DATABASE",
		"POSTHOG_TEST_PG_USER",
		"POSTHOG_TEST_PG_PASSWORD",
	} {
		if os.Getenv(v) == "" {
			t.Skipf("Skipping test: %s not set", v)
		}
	}
}

func testAccCheckExternalDataSourceDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_external_data_source" {
			continue
		}

		_, status, err := client.GetExternalDataSource(context.Background(), projectID, rs.Primary.ID)
		if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
			continue
		}
		if err != nil {
			return fmt.Errorf("unexpected error checking external data source %s: %w", rs.Primary.ID, err)
		}
		return fmt.Errorf("external data source %s still exists", rs.Primary.ID)
	}

	return nil
}

// TestExternalDataSource_InvalidJobInputs verifies that a malformed job_inputs_json
// is rejected client-side without any API call. Requires no credentials.
func TestExternalDataSource_InvalidJobInputs(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
provider "posthog" {}

resource "posthog_external_data_source" "test" {
  source_type     = "Stripe"
  job_inputs_json = "not-json"
  schemas         = ["charges"]
}
`,
				ExpectError: regexp.MustCompile(`(?i)invalid job_inputs_json`),
			},
		},
	})
}

func TestExternalDataSource_Postgres_Basic(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoPostgresCreds(t)

	rPrefix := acctest.RandomWithPrefix("tf_acc_")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExternalDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExternalDataSourcePostgres(rPrefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_external_data_source.test", "source_type", "Postgres"),
					resource.TestCheckResourceAttrSet("posthog_external_data_source.test", "id"),
				),
			},
		},
	})
}

func testAccExternalDataSourcePostgres(prefix string) string {
	port := os.Getenv("POSTHOG_TEST_PG_PORT")
	if port == "" {
		port = "5432"
	}

	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_external_data_source" "test" {
  source_type = "Postgres"
  prefix      = %q
  schemas     = ["test_data"]

  job_inputs_json = jsonencode({
    host     = %q
    port     = %q
    database = %q
    user     = %q
    password = %q
    schema   = "public"
  })
}
`,
		prefix,
		os.Getenv("POSTHOG_TEST_PG_HOST"),
		port,
		os.Getenv("POSTHOG_TEST_PG_DATABASE"),
		os.Getenv("POSTHOG_TEST_PG_USER"),
		os.Getenv("POSTHOG_TEST_PG_PASSWORD"),
	)
}
