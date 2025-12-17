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

// testAccCheckHogFunctionDestroy verifies the hog function has been destroyed (soft deleted).
func testAccCheckHogFunctionDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		os.Getenv("POSTHOG_PROJECT_ID"),
		"test",
	)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_hog_function" {
			continue
		}

		_, status, err := client.GetHogFunction(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("hog_function %s still exists", rs.Primary.ID)
		}
		if status != httpclient.HTTPStatusCode(http.StatusNotFound) {
			return fmt.Errorf("expected 404, got %d", status)
		}
	}

	return nil
}

// TestHogFunction_WebhookDestination tests creating a webhook destination hog function.
func TestHogFunction_WebhookDestination(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHogFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHogFunctionWebhook(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "true"),
					resource.TestCheckResourceAttrSet("posthog_hog_function.test", "id"),
				),
			},
		},
	})
}

// TestHogFunction_WithFilters tests creating a hog function with event filters.
func TestHogFunction_WithFilters(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHogFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHogFunctionWithFilters(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_hog_function.test", "filters_json"),
				),
			},
		},
	})
}

// TestHogFunction_Update tests updating a hog function.
func TestHogFunction_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHogFunctionDestroy,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccHogFunctionWebhook(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "true"),
				),
			},
			// Update name
			{
				Config: testAccHogFunctionWebhook(rName + "-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "true"),
				),
			},
			// Disable
			{
				Config: testAccHogFunctionDisabled(rName + "-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "false"),
				),
			},
		},
	})
}

// TestHogFunction_EnableDisable tests enabling and disabling a hog function.
func TestHogFunction_EnableDisable(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHogFunctionDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccHogFunctionWithEnabled(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "true"),
				),
			},
			// Disable
			{
				Config: testAccHogFunctionWithEnabled(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "false"),
				),
			},
			// Re-enable
			{
				Config: testAccHogFunctionWithEnabled(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "true"),
				),
			},
		},
	})
}

// TestHogFunction_Import tests importing an existing hog function by ID.
func TestHogFunction_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHogFunctionDestroy,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccHogFunctionWebhook(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName),
				),
			},
			// Import
			{
				ResourceName:            "posthog_hog_function.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"hog", "inputs_json", "filters_json", "inputs_schema_json"},
			},
		},
	})
}

// TestHogFunction_ImportWithHogCode tests importing a hog function with explicit hog code.
// This verifies that the hog attribute is properly populated during import.
func TestHogFunction_ImportWithHogCode(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHogFunctionDestroy,
		Steps: []resource.TestStep{
			// Create with explicit hog code
			{
				Config: testAccHogFunctionWithHogCode(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_hog_function.test", "hog"),
				),
			},
			// Import and verify hog is preserved (not ignored)
			{
				ResourceName:            "posthog_hog_function.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"inputs_json", "filters_json", "inputs_schema_json"},
			},
		},
	})
}

// TestHogFunction_AlertWebhookIntegration tests the full chain:
// insight → alert → hog_function webhook notification when alert fires.
func TestHogFunction_AlertWebhookIntegration(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHogFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHogFunctionAlertWebhookIntegration(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify insight
					resource.TestCheckResourceAttr("posthog_insight.test", "name", rName+"-insight"),
					resource.TestCheckResourceAttrSet("posthog_insight.test", "id"),
					// Verify alert
					resource.TestCheckResourceAttr("posthog_alert.test", "name", rName+"-alert"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_type", "absolute"),
					resource.TestCheckResourceAttr("posthog_alert.test", "threshold_upper", "1000"),
					resource.TestCheckResourceAttr("posthog_alert.test", "calculation_interval", "daily"),
					resource.TestCheckResourceAttr("posthog_alert.test", "skip_weekend", "false"),
					resource.TestCheckResourceAttrSet("posthog_alert.test", "id"),
					// Verify hog function
					resource.TestCheckResourceAttr("posthog_hog_function.test", "name", rName+"-webhook"),
					resource.TestCheckResourceAttr("posthog_hog_function.test", "type", "internal_destination"),
					resource.TestCheckResourceAttr("posthog_hog_function.test", "enabled", "true"),
					resource.TestCheckResourceAttrSet("posthog_hog_function.test", "id"),
				),
			},
		},
	})
}

func testAccHogFunctionWebhook(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_hog_function" "test" {
  name        = %q
  description = "Test webhook destination"
  type        = "destination"
  enabled     = true
  template_id = "template-webhook"

  inputs_json = jsonencode({
    url = {
      value      = "https://example.com/webhook"
      templating = "hog"
    }
    method = {
      value      = "POST"
      templating = "hog"
    }
    body = {
      value = {
        event     = "{event.event}"
        timestamp = "{event.timestamp}"
      }
      templating = "hog"
    }
    headers = {
      value = {
        "Content-Type" = "application/json"
      }
      templating = "hog"
    }
  })

  filters_json = jsonencode({
    source = "events"
    events = [{
      id   = "$pageview"
      name = "$pageview"
      type = "events"
    }]
    filter_test_accounts = false
  })
}
`, name)
}

func testAccHogFunctionWithFilters(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_hog_function" "test" {
  name        = %q
  description = "Test webhook with filters"
  type        = "destination"
  enabled     = true
  template_id = "template-webhook"

  inputs_json = jsonencode({
    url = {
      value      = "https://example.com/webhook"
      templating = "hog"
    }
    method = {
      value      = "POST"
      templating = "hog"
    }
  })

  filters_json = jsonencode({
    source = "events"
    events = [
      {
        id   = "$pageview"
        name = "$pageview"
        type = "events"
      },
      {
        id   = "custom_event"
        name = "custom_event"
        type = "events"
      }
    ]
    filter_test_accounts = false
  })
}
`, name)
}

func testAccHogFunctionDisabled(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_hog_function" "test" {
  name        = %q
  description = "Test webhook destination"
  type        = "destination"
  enabled     = false
  template_id = "template-webhook"

  inputs_json = jsonencode({
    url = {
      value      = "https://example.com/webhook"
      templating = "hog"
    }
    method = {
      value      = "POST"
      templating = "hog"
    }
  })

  filters_json = jsonencode({
    source = "events"
    events = [{
      id   = "$pageview"
      name = "$pageview"
      type = "events"
    }]
    filter_test_accounts = false
  })
}
`, name)
}

func testAccHogFunctionWithEnabled(name string, enabled bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_hog_function" "test" {
  name        = %q
  description = "Test webhook destination"
  type        = "destination"
  enabled     = %t
  template_id = "template-webhook"

  inputs_json = jsonencode({
    url = {
      value      = "https://example.com/webhook"
      templating = "hog"
    }
    method = {
      value      = "POST"
      templating = "hog"
    }
  })

  filters_json = jsonencode({
    source = "events"
    events = [{
      id   = "$pageview"
      name = "$pageview"
      type = "events"
    }]
    filter_test_accounts = false
  })
}
`, name, enabled)
}

func testAccHogFunctionAlertWebhookIntegration(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

# Step 1: Create an insight to monitor
resource "posthog_insight" "test" {
  name = "%[1]s-insight"

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

# Step 2: Create an alert on the insight
resource "posthog_alert" "test" {
  name                 = "%[1]s-alert"
  insight              = posthog_insight.test.id
  subscribed_users     = []
  threshold_type       = "absolute"
  threshold_upper      = 1000
  condition_type       = "absolute_value"
  calculation_interval = "daily"
  skip_weekend         = false

  depends_on = [posthog_insight.test]
}

# Step 3: Create a webhook that fires when the alert triggers
resource "posthog_hog_function" "test" {
  name        = "%[1]s-webhook"
  description = "Webhook notification when alert fires"
  type        = "internal_destination"
  enabled     = true
  template_id = "template-webhook"

  inputs_json = jsonencode({
    url = {
      value      = "https://example.com/alert-webhook"
      templating = "hog"
    }
    method = {
      value      = "POST"
      templating = "hog"
    }
    body = {
      value = {
        alert_name = "{event.properties.alert_name}"
        alert_id   = "{event.properties.alert_id}"
        value      = "{event.properties.alert_calculated_value}"
      }
      templating = "hog"
    }
    headers = {
      value = {
        "Content-Type" = "application/json"
      }
      templating = "hog"
    }
    debug = {
      value      = false
      templating = "hog"
    }
  })

  filters_json = jsonencode({
    source = "events"
    events = [{
      id   = "$insight_alert_firing"
      type = "events"
    }]
    properties = [{
      key      = "alert_id"
      type     = "event"
      value    = posthog_alert.test.id
      operator = "exact"
    }]
  })

  depends_on = [posthog_alert.test]
}
`, name)
}

func testAccHogFunctionWithHogCode(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_hog_function" "test" {
  name        = %q
  description = "Test hog function with explicit hog code"
  type        = "destination"
  enabled     = true
  template_id = "template-webhook"

  hog = <<-EOT
    let payload := {
      'headers': inputs.headers,
      'body': inputs.body,
      'method': inputs.method
    }

    if (inputs.debug) {
      print('Request', inputs.url, payload)
    }

    let res := fetch(inputs.url, payload);

    if (res.status >= 400) {
      throw Error(f'Webhook failed with status {res.status}: {res.body}');
    }

    if (inputs.debug) {
      print('Response', res.status, res.body);
    }
  EOT

  inputs_json = jsonencode({
    url = {
      value      = "https://example.com/webhook"
      templating = "hog"
    }
    method = {
      value      = "POST"
      templating = "hog"
    }
    body = {
      value = {
        event     = "{event.event}"
        timestamp = "{event.timestamp}"
      }
      templating = "hog"
    }
    headers = {
      value = {
        "Content-Type" = "application/json"
      }
      templating = "hog"
    }
  })

  filters_json = jsonencode({
    source = "events"
    events = [{
      id   = "$pageview"
      name = "$pageview"
      type = "events"
    }]
    filter_test_accounts = false
  })
}
`, name)
}
