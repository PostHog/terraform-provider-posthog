package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
)

// TestFeatureFlagsDataSource_ListsExistingFlags creates a project holding one
// flag and one archived flag through the API, simulating flags made in the
// PostHog UI, and checks that the data source finds both by project_id and by
// project_name, and finds nothing for a project name that does not exist. The
// project is created for the test so that its name is unique in the
// organization.
func TestFeatureFlagsDataSource_ListsExistingFlags(t *testing.T) {
	skipIfNotAcceptance(t)
	skipIfNoOrganizationID(t)

	host, apiKey, orgID := os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), getOrganizationID()
	client := httpclient.NewDefaultClient(host, apiKey, "acceptance-test")

	projectName := acctest.RandomWithPrefix("tf-acc-flags-ds")
	project, err := client.CreateProject(context.Background(), orgID, httpclient.ProjectRequest{Name: util.StringPtr(projectName)})
	if err != nil {
		t.Fatalf("failed to create project via API: %s", err)
	}
	projectID := strconv.FormatInt(project.ID, 10)
	t.Cleanup(func() {
		if _, err := client.DeleteProject(context.Background(), orgID, projectID); err != nil {
			t.Logf("failed to delete test project %s, delete it by hand: %s", projectID, err)
		}
	})

	key := acctest.RandomWithPrefix("tf-acc-flags-ds")
	archivedKey := key + "-archived"
	flag := createFeatureFlagForDataSourceTest(t, client, projectID, key, true)
	archived := createFeatureFlagForDataSourceTest(t, client, projectID, archivedKey, false)
	if err := archiveFeatureFlagRaw(host, apiKey, projectID, archived.ID); err != nil {
		t.Fatalf("failed to archive feature flag: %s", err)
	}

	flagID := strconv.FormatInt(flag.ID, 10)
	archivedID := strconv.FormatInt(archived.ID, 10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
provider "posthog" {}

data "posthog_feature_flags" "test" {
  project_id = %q
}
`, projectID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.posthog_feature_flags.test", "project_id", projectID),
					resource.TestCheckTypeSetElemNestedAttrs("data.posthog_feature_flags.test", "flags.*", map[string]string{
						"id": flagID, "key": key, "name": "Data source test flag", "active": "true", "archived": "false",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("data.posthog_feature_flags.test", "flags.*", map[string]string{
						"id": archivedID, "key": archivedKey, "active": "false", "archived": "true",
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
provider "posthog" {}

data "posthog_feature_flags" "test" {
  organization_id = %q
  project_name    = %q
}
`, orgID, strings.ToUpper(projectName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.posthog_feature_flags.test", "project_id", projectID),
					resource.TestCheckResourceAttr("data.posthog_feature_flags.test", "flags.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("data.posthog_feature_flags.test", "flags.*", map[string]string{
						"id": flagID, "key": key,
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
provider "posthog" {}

data "posthog_feature_flags" "test" {
  organization_id = %q
  project_name    = %q
}
`, orgID, projectName+"-no-such-project"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.posthog_feature_flags.test", "project_id"),
					resource.TestCheckResourceAttr("data.posthog_feature_flags.test", "flags.#", "0"),
				),
			},
		},
	})
}

// TestFeatureFlagsDataSource_AdoptsExistingFlagsWithForEachImport runs the
// adoption flow the data source exists for: one resource declares an existing
// flag and a new one, a for_each import block fed by the data source adopts the
// existing flag, and the new one is created. The plan the test framework runs
// after apply proves the import is a no-op once the flags are in state.
//
// The resource uses count rather than the for_each a real configuration would
// use, because the test framework cannot read for_each instances out of state.
// The import block does the same thing either way.
func TestFeatureFlagsDataSource_AdoptsExistingFlagsWithForEachImport(t *testing.T) {
	skipIfNotAcceptance(t)

	host, apiKey, projectID := os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), getProjectID()
	client := httpclient.NewDefaultClient(host, apiKey, "acceptance-test")

	existingKey := acctest.RandomWithPrefix("tf-acc-flags-adopt")
	newKey := existingKey + "-new"
	existing := createFeatureFlagForDataSourceTest(t, client, projectID, existingKey, true)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// for_each in an import block arrived in Terraform 1.7.
			tfversion.SkipBelow(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
provider "posthog" {}

locals {
  feature_flag_keys = [%[2]q, %[3]q]
}

data "posthog_feature_flags" "existing" {
  project_id = %[1]q
}

resource "posthog_feature_flag" "this" {
  count      = length(local.feature_flag_keys)
  project_id = %[1]q
  key        = local.feature_flag_keys[count.index]
  name       = "Data source test flag"
  active     = true
}

import {
  for_each = {
    for flag in data.posthog_feature_flags.existing.flags : flag.key => {
      index = index(local.feature_flag_keys, flag.key)
      id    = flag.id
    }
    if contains(local.feature_flag_keys, flag.key)
  }
  to = posthog_feature_flag.this[each.value.index]
  id = "${data.posthog_feature_flags.existing.project_id}/${each.value.id}"
}
`, projectID, existingKey, newKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("posthog_feature_flag.this.0", "id", strconv.FormatInt(existing.ID, 10)),
					resource.TestCheckResourceAttr("posthog_feature_flag.this.0", "key", existingKey),
					resource.TestCheckResourceAttr("posthog_feature_flag.this.1", "key", newKey),
					resource.TestCheckResourceAttrSet("posthog_feature_flag.this.1", "id"),
					testCheckDifferentFlagIDs("posthog_feature_flag.this.0", "posthog_feature_flag.this.1"),
				),
			},
		},
	})
}

// TestFeatureFlagsDataSource_DefaultsAndConflicts checks the data source with
// no arguments, which lists the provider-level project, and that setting
// project_id with project_name, or organization_id without it, is rejected
// before any request is made. The rejected configurations go first, because the test framework destroys using
// the last step's configuration.
func TestFeatureFlagsDataSource_DefaultsAndConflicts(t *testing.T) {
	skipIfNotAcceptance(t)

	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "acceptance-test")
	projectID := getProjectID()
	key := acctest.RandomWithPrefix("tf-acc-flags-default")
	flag := createFeatureFlagForDataSourceTest(t, client, projectID, key, true)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
provider "posthog" {}

data "posthog_feature_flags" "test" {
  project_id   = %q
  project_name = "anything"
}
`, projectID),
				ExpectError: regexp.MustCompile(`(?s)Invalid Attribute Combination.*project_name`),
			},
			{
				Config: fmt.Sprintf(`
provider "posthog" {}

data "posthog_feature_flags" "test" {
  project_id      = %q
  organization_id = "anything"
}
`, projectID),
				ExpectError: regexp.MustCompile(`(?s)Invalid Attribute Combination.*project_name`),
			},
			{
				Config: `
provider "posthog" {}

data "posthog_feature_flags" "test" {}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.posthog_feature_flags.test", "project_id", projectID),
					resource.TestCheckTypeSetElemNestedAttrs("data.posthog_feature_flags.test", "flags.*", map[string]string{
						"id": strconv.FormatInt(flag.ID, 10), "key": key,
					}),
				),
			},
		},
	})
}

func createFeatureFlagForDataSourceTest(t *testing.T, client httpclient.PosthogClient, projectID, key string, active bool) httpclient.FeatureFlag {
	t.Helper()
	flag, err := client.CreateFeatureFlag(context.Background(), projectID, httpclient.FeatureFlagRequest{
		Key:    key,
		Name:   util.StringPtr("Data source test flag"),
		Active: util.BoolPtr(active),
	})
	if err != nil {
		t.Fatalf("failed to create feature flag %q via API: %s", key, err)
	}
	// Terraform may already have deleted the flag by the time this runs, so the
	// error is not checked.
	t.Cleanup(func() {
		_, _ = client.DeleteFeatureFlag(context.Background(), projectID, strconv.FormatInt(flag.ID, 10))
	})
	return flag
}

// archiveFeatureFlagRaw archives a flag. The provider's client has no archive
// call because the resource does not manage archived state.
func archiveFeatureFlagRaw(host, apiKey, projectID string, flagID int64) error {
	body, _ := json.Marshal(map[string]interface{}{"archived": true})
	url := fmt.Sprintf("%s/api/projects/%s/feature_flags/%d/", strings.TrimRight(host, "/"), projectID, flagID)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("archive feature flag returned %d: %s", resp.StatusCode, string(respBody))
	}
	return nil
}

func testCheckDifferentFlagIDs(nameA, nameB string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		a := s.RootModule().Resources[nameA]
		b := s.RootModule().Resources[nameB]
		if a == nil || b == nil {
			return fmt.Errorf("expected both %s and %s in state", nameA, nameB)
		}
		if a.Primary.ID == b.Primary.ID {
			return fmt.Errorf("adopted and created flags share id %s", a.Primary.ID)
		}
		return nil
	}
}
