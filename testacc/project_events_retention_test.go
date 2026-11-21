package tests

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/pkg/httpclient"
)

func TestProjectEventsRetention_Basic(t *testing.T) {
	skipIfNotAcceptance(t)
	testAccProjectPreCheck(t)
	ctx := context.Background()
	client := testAccAPIClient()
	orgID := getOrganizationID()
	name := acctest.RandomWithPrefix("tf-acc-retention")
	project, err := client.CreateProject(ctx, orgID, httpclient.ProjectRequest{Name: &name})
	if err != nil {
		t.Fatalf("creating test project: %v", err)
	}
	projectID := strconv.FormatInt(project.ID, 10)
	t.Cleanup(func() {
		if _, err := client.DeleteProject(ctx, orgID, projectID); err != nil {
			t.Errorf("deleting test project: %v", err)
		}
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProjectPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: testAccProjectEventsRetention(projectID, "retention_months = 12"), Check: resource.TestCheckResourceAttr("posthog_project_events_retention.test", "retention_months", "12")},
			{ResourceName: "posthog_project_events_retention.test", ImportState: true, ImportStateVerify: true, ImportStateId: projectID},
			{Config: testAccProjectEventsRetention(projectID, ""), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckNoResourceAttr("posthog_project_events_retention.test", "retention_months"), func(*terraform.State) error {
				result, _, err := client.GetProjectEventsRetention(ctx, projectID)
				if err != nil {
					return err
				}
				if result.ProjectRetentionMonths != nil {
					return fmt.Errorf("expected cleared override, got %d", *result.ProjectRetentionMonths)
				}
				return nil
			})},
		},
	})
}

func testAccProjectEventsRetention(projectID, setting string) string {
	return fmt.Sprintf(`
provider "posthog" {}
resource "posthog_project_events_retention" "test" {
  project_id = %q
  %s
}
`, projectID, setting)
}
