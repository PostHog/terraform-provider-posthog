package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccAccessControlPreCheck(t *testing.T) {
	t.Helper()
	testAccBasePreCheck(t)
	skipIfNoOrganizationID(t)
}

func TestAccessControl_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "id"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "resource", "feature_flag"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "editor"),
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "role"),
				),
			},
		},
	})
}

func TestAccessControl_Update(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlWithLevel(orgID, rName, "viewer"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "viewer"),
				),
			},
			{
				Config: testAccAccessControlWithLevel(orgID, rName, "editor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "editor"),
				),
			},
		},
	})
}

func TestAccessControl_OrganizationMember(t *testing.T) {
	skipIfNotAcceptance(t)

	orgID := getOrganizationID()
	userEmail := getTestUserEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t); skipIfNoTestUserEmail(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlWithMember(orgID, userEmail),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "id"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "resource", "dashboard"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "editor"),
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "organization_member"),
				),
			},
		},
	})
}

func TestAccessControl_None(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlWithLevel(orgID, rName, "editor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "editor"),
				),
			},
			{
				Config: testAccAccessControlWithLevel(orgID, rName, "none"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "none"),
				),
			},
		},
	})
}

func TestAccessControl_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlBasic(orgID, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "id"),
				),
			},
			{
				ResourceName:            "posthog_access_control.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func TestAccessControl_ResourceInstance(t *testing.T) {
	skipIfNotAcceptance(t)

	rName := acctest.RandomWithPrefix("tf-acc-role")
	dashboardName := acctest.RandomWithPrefix("tf-acc-dashboard")
	orgID := getOrganizationID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlResourceInstance(orgID, rName, dashboardName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "id"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "resource", "dashboard"),
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "resource_id"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "viewer"),
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "role"),
				),
			},
		},
	})
}

func testAccAccessControlBasic(orgID, roleName string) string {
	return testAccAccessControlWithLevel(orgID, roleName, "editor")
}

func testAccAccessControlWithLevel(orgID, roleName, accessLevel string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_role" "test" {
  organization_id = %q
  name            = %q
}

resource "posthog_access_control" "test" {
  resource     = "feature_flag"
  access_level = %q
  role         = posthog_role.test.id
}
`, orgID, roleName, accessLevel)
}

func testAccAccessControlWithMember(orgID, userEmail string) string {
	return fmt.Sprintf(`
provider "posthog" {}

data "posthog_user" "test" {
  organization_id = %q
  email           = %q
}

resource "posthog_organization_member" "test" {
  organization_id    = %q
  user_uuid          = data.posthog_user.test.uuid
  retain_on_destroy  = true
}

resource "posthog_access_control" "test" {
  resource            = "dashboard"
  access_level        = "editor"
  organization_member = posthog_organization_member.test.id
}
`, orgID, userEmail, orgID)
}

func testAccAccessControlResourceInstance(orgID, roleName, dashboardName string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_role" "test" {
  organization_id = %q
  name            = %q
}

resource "posthog_dashboard" "test" {
  name = %q
}

resource "posthog_access_control" "test" {
  resource     = "dashboard"
  resource_id  = posthog_dashboard.test.id
  access_level = "viewer"
  role         = posthog_role.test.id
}
`, orgID, roleName, dashboardName)
}

func TestAccessControl_ProjectDefault(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlProjectDefault("editor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "id"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "resource", "survey"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "editor"),
					// Verify neither role nor organization_member is set
					resource.TestCheckNoResourceAttr("posthog_access_control.test", "role"),
					resource.TestCheckNoResourceAttr("posthog_access_control.test", "organization_member"),
				),
			},
			{
				Config: testAccAccessControlProjectDefault("viewer"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "viewer"),
				),
			},
		},
	})
}

func TestAccessControl_ProjectDefaultImport(t *testing.T) {
	skipIfNotAcceptance(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlProjectDefault("editor"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "id"),
				),
			},
			{
				ResourceName:            "posthog_access_control.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func TestAccessControl_ProjectDefaultResourceInstance(t *testing.T) {
	skipIfNotAcceptance(t)

	dashboardName := acctest.RandomWithPrefix("tf-acc-dashboard")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccAccessControlPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessControlProjectDefaultResourceInstance(dashboardName, "none"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "id"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "resource", "dashboard"),
					resource.TestCheckResourceAttrSet("posthog_access_control.test", "resource_id"),
					resource.TestCheckResourceAttr("posthog_access_control.test", "access_level", "none"),
					// Verify neither role nor organization_member is set
					resource.TestCheckNoResourceAttr("posthog_access_control.test", "role"),
					resource.TestCheckNoResourceAttr("posthog_access_control.test", "organization_member"),
				),
			},
		},
	})
}

func testAccAccessControlProjectDefault(accessLevel string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_access_control" "test" {
  resource     = "survey"
  access_level = %q
}
`, accessLevel)
}

func testAccAccessControlProjectDefaultResourceInstance(dashboardName, accessLevel string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_dashboard" "test" {
  name = %q
}

resource "posthog_access_control" "test" {
  resource     = "dashboard"
  resource_id  = posthog_dashboard.test.id
  access_level = %q
}
`, dashboardName, accessLevel)
}
