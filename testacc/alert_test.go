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

// testAccCheckAlertDestroy verifies the alert has been destroyed.
func testAccCheckAlertDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		os.Getenv("POSTHOG_PROJECT_ID"),
		"test",
	)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_alert" {
			continue
		}

		_, status, err := client.GetAlert(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("alert %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

// TestAlert_Basic tests creating an alert with minimal configuration.
func TestAlert_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_type", "absolute"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "id"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "insight"),
				),
			},
		},
	})
}

// TestAlert_AllFields tests creating an alert with all optional fields.
func TestAlert_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "true"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_type", "absolute"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "10"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "100"),
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "absolute_value"),
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "daily"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "id"),
				),
			},
		},
	})
}

// TestAlert_Update tests updating an alert.
func TestAlert_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccAlertWithThreshold(rName, 10, 100),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "10"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "100"),
				),
			},
			// Update thresholds
			{
				Config: testAccAlertWithThreshold(rName, 20, 200),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "20"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "200"),
				),
			},
			// Update name
			{
				Config: testAccAlertWithThreshold(rName+"-updated", 20, 200),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_lower", "20"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "200"),
				),
			},
		},
	})
}

// TestAlert_EnableDisable tests enabling and disabling an alert.
func TestAlert_EnableDisable(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccAlertWithEnabled(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "true"),
				),
			},
			// Disable
			{
				Config: testAccAlertWithEnabled(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "false"),
				),
			},
			// Re-enable
			{
				Config: testAccAlertWithEnabled(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "enabled", "true"),
				),
			},
		},
	})
}

// TestAlert_ConditionTypes tests different condition types.
func TestAlert_ConditionTypes(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Absolute value
			{
				Config: testAccAlertWithCondition(rName, "absolute_value"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "absolute_value"),
				),
			},
			// Relative increase
			{
				Config: testAccAlertWithCondition(rName, "relative_increase"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "relative_increase"),
				),
			},
			// Relative decrease
			{
				Config: testAccAlertWithCondition(rName, "relative_decrease"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "condition_type", "relative_decrease"),
				),
			},
		},
	})
}

// TestAlert_CalculationIntervals tests different calculation intervals.
func TestAlert_CalculationIntervals(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertWithInterval(rName, "hourly"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "hourly"),
				),
			},
			{
				Config: testAccAlertWithInterval(rName, "daily"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "daily"),
				),
			},
			{
				Config: testAccAlertWithInterval(rName, "weekly"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "weekly"),
				),
			},
		},
	})
}

// TestAlert_Import tests importing an existing alert by ID.
func TestAlert_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
				),
			},
			// Import
			{
				ResourceName:            "posthog_alert.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subscribed_users"},
			},
		},
	})
}

// TestAlert_InsightDeletion tests the behavior when removing an alert that references an insight.
// This verifies that:
// 1. Alert can be removed while insight remains
// 2. Insight can be deleted after alert is removed
func TestAlert_InsightDeletion(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	insightOnlyConfig := fmt.Sprintf(`
provider "posthog" {}

resource "posthog_insight" "test" {
  name = "%s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        name  = "$pageview"
        event = "$pageview"
        math  = "total"
      }]
    }
  })
}
`, rName)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create insight and alert
			{
				Config: testAccAlertBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "insight"),
				),
			},
			// Step 2: Remove alert, keep insight - alert should be deleted
			{
				Config: insightOnlyConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName+"-insight"),
					resource.TestCheckNoResourceAttr("posthog_alert.test", "name"),
				),
			},
			// Step 3: Remove insight too - should succeed since alert is gone
			{
				Config: `provider "posthog" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("posthog_insight.test", "name"),
					resource.TestCheckNoResourceAttr("posthog_alert.test", "name"),
				),
			},
			// Step 4: Recreate
			{
				Config: testAccAlertBasic(rName + "-recreated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName+"-recreated"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "insight"),
				),
			},
		},
	})
}

// Helper function to create the base insight that alerts monitor
func testAccAlertInsightBase(name string) string {
	return fmt.Sprintf(`
resource "posthog_insight" "test" {
  name = "%s-insight"

  query_json = jsonencode({
    kind   = "InsightVizNode"
    source = {
      kind   = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        name  = "$pageview"
        event = "$pageview"
        math  = "total"
      }]
    }
  })
}
`, name)
}

func testAccAlertBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_upper  = 100

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name)
}

func testAccAlertAllFields(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name                 = %q
  insight              = posthog_insight.test.id
  subscribed_users     = []
  enabled              = true
  threshold_type       = "absolute"
  threshold_lower      = 10
  threshold_upper      = 100
  condition_type       = "absolute_value"
  series_index         = 0
  calculation_interval = "daily"
  skip_weekend         = false

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name)
}

func testAccAlertWithThreshold(name string, lower, upper int) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_lower  = %d
  threshold_upper  = %d

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, lower, upper)
}

func testAccAlertWithEnabled(name string, enabled bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  enabled          = %t
  threshold_type   = "absolute"
  threshold_upper  = 100

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, enabled)
}

func testAccAlertWithCondition(name, conditionType string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name             = %q
  insight          = posthog_insight.test.id
  subscribed_users = []
  threshold_type   = "absolute"
  threshold_upper  = 100
  condition_type   = %q

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, conditionType)
}

func testAccAlertWithInterval(name, interval string) string {
	return fmt.Sprintf(`
provider "posthog" {}

%s

resource "posthog_alert" "test" {
  name                 = %q
  insight              = posthog_insight.test.id
  subscribed_users     = []
  threshold_type       = "absolute"
  threshold_upper      = 100
  calculation_interval = %q

  depends_on = [posthog_insight.test]
}
`, testAccAlertInsightBase(name), name, interval)
}
