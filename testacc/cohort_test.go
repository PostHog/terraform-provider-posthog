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

func testAccCheckCohortDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_cohort" {
			continue
		}

		cohort, status, err := client.GetCohort(context.Background(), projectID, rs.Primary.ID)
		if err != nil {
			if status == httpclient.HTTPStatusCode(http.StatusNotFound) {
				continue
			}
			return fmt.Errorf("unexpected error checking cohort %s: %w", rs.Primary.ID, err)
		}
		if cohort.Deleted == nil || !*cohort.Deleted {
			return fmt.Errorf("cohort %s still exists and is not soft-deleted", rs.Primary.ID)
		}
	}

	return nil
}

func TestCohort_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCohortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCohortBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_cohort.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_cohort.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_cohort.test", "created_at"),
					resource.TestCheckResourceAttr("posthog_cohort.test", "is_static", "false"),
					resource.TestCheckResourceAttr("posthog_cohort.test", "deleted", "false"),
				),
			},
		},
	})
}

// TestCohort_Filters verifies that a filters blob survives a round-trip without
// drift, despite PostHog annotating property values with computed fields.
func TestCohort_Filters(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCohortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCohortWithFilters(rName, "@example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_cohort.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_cohort.test", "description", "Acceptance test cohort"),
					resource.TestCheckResourceAttrSet("posthog_cohort.test", "filters"),
				),
			},
			// Re-applying the same config must be a no-op.
			{
				Config:   testAccCohortWithFilters(rName, "@example.com"),
				PlanOnly: true,
			},
		},
	})
}

func TestCohort_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCohortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCohortWithFilters(rName, "@example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_cohort.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_cohort.test", "description", "Acceptance test cohort"),
				),
			},
			// Change the name and the filter value.
			{
				Config: testAccCohortWithFilters(rName+"-updated", "@updated.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_cohort.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttrSet("posthog_cohort.test", "filters"),
				),
			},
			// Drop the optional description.
			{
				Config: testAccCohortBasic(rName + "-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_cohort.test", "name", rName+"-updated"),
				),
			},
		},
	})
}

func TestCohort_Static(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCohortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCohortStatic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_cohort.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_cohort.test", "is_static", "true"),
				),
			},
		},
	})
}

func TestCohort_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCohortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCohortBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_cohort.test", "name", rName),
				),
			},
			{
				ResourceName:            "posthog_cohort.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"filters"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["posthog_cohort.test"]
					if !ok {
						return "", fmt.Errorf("resource not found: posthog_cohort.test")
					}
					return fmt.Sprintf("%s/%s", projectID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccCohortBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_cohort" "test" {
  name = %q
}
`, name)
}

func testAccCohortWithFilters(name, emailFragment string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_cohort" "test" {
  name        = %q
  description = "Acceptance test cohort"

  filters = jsonencode({
    properties = {
      type = "AND"
      values = [
        {
          type     = "person"
          key      = "email"
          operator = "icontains"
          value    = %q
        }
      ]
    }
  })
}
`, name, emailFragment)
}

func testAccCohortStatic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_cohort" "test" {
  name      = %q
  is_static = true
}
`, name)
}
