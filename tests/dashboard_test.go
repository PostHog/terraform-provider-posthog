package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestDashboard_Basic tests creating a dashboard with only the required field (name).
func TestDashboard_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDashboardBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName),
					resource.TestCheckResourceAttrSet("posthog_dashboard.test", "id"),
					resource.TestCheckNoResourceAttr("posthog_dashboard.test", "description"),
				),
			},
		},
	})
}

// TestDashboard_AllFields tests creating a dashboard with all optional fields.
func TestDashboard_AllFields(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDashboardAllFields(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", "Test dashboard description"),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "pinned", "true"),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("posthog_dashboard.test", "id"),
				),
			},
		},
	})
}

// TestDashboard_Update tests updating each dashboard field individually.
func TestDashboard_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccDashboardWithDescription(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", "Initial description"),
				),
			},
			// Update name
			{
				Config: testAccDashboardWithDescription(rName+"-updated", "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", "Initial description"),
				),
			},
			// Update description
			{
				Config: testAccDashboardWithDescription(rName+"-updated", "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName+"-updated"),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", "Updated description"),
				),
			},
		},
	})
}

// TestDashboard_Tags tests creating, updating, and removing tags.
func TestDashboard_Tags(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with tags
			{
				Config: testAccDashboardWithTags(rName, []string{"tag1", "tag2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "tags.#", "2"),
				),
			},
			// Add more tags
			{
				Config: testAccDashboardWithTags(rName, []string{"tag1", "tag2", "tag3"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "tags.#", "3"),
				),
			},
			// Remove tags
			{
				Config: testAccDashboardWithTags(rName, []string{"tag1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "tags.#", "1"),
				),
			},
		},
	})
}

// TestDashboard_Pinned tests toggling the pinned state.
func TestDashboard_Pinned(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create unpinned
			{
				Config: testAccDashboardPinned(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "pinned", "false"),
				),
			},
			// Pin it
			{
				Config: testAccDashboardPinned(rName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "pinned", "true"),
				),
			},
			// Unpin it
			{
				Config: testAccDashboardPinned(rName, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "pinned", "false"),
				),
			},
		},
	})
}

// TestDashboard_Import tests importing an existing dashboard by ID.
func TestDashboard_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccDashboardWithDescription(rName, "Import test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName),
				),
			},
			// Import
			{
				ResourceName:      "posthog_dashboard.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestDashboard_Recreate tests deleting and recreating a dashboard with the same name.
func TestDashboard_Recreate(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create first dashboard
			{
				Config: testAccDashboardBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName),
				),
			},
			// Remove resource from config (destroys it)
			{
				Config: `provider "posthog" {}`,
			},
			// Recreate
			{
				Config: testAccDashboardWithDescription(rName, "Recreated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", "Recreated"),
				),
			},
		},
	})
}

// TestDashboard_EmptyDescription tests handling of empty vs null description.
func TestDashboard_EmptyDescription(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with description
			{
				Config: testAccDashboardWithDescription(rName, "Has description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", "Has description"),
				),
			},
			// Remove description (set to empty)
			{
				Config: testAccDashboardBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_dashboard.test", "name", rName),
					resource.TestCheckResourceAttr("posthog_dashboard.test", "description", ""),
				),
			},
		},
	})
}

// Config generators

func testAccDashboardBasic(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %q
}
`, name)
}

func testAccDashboardWithDescription(name, description string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name        = %q
  description = %q
}
`, name, description)
}

func testAccDashboardAllFields(name string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name        = %q
  description = "Test dashboard description"
  pinned      = true
  tags        = ["managed-by:terraform", "env:test"]
}
`, name)
}

func testAccDashboardWithTags(name string, tags []string) string {
	tagsStr := ""
	for i, tag := range tags {
		if i > 0 {
			tagsStr += ", "
		}
		tagsStr += fmt.Sprintf("%q", tag)
	}

	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %q
  tags = [%s]
}
`, name, tagsStr)
}

func testAccDashboardPinned(name string, pinned bool) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name   = %q
  pinned = %t
}
`, name, pinned)
}
