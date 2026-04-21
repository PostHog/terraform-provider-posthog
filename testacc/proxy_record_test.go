package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

const (
	proxyRecordValidWaitTimeout = 10 * time.Minute
	testProxyRecordResourceName = "posthog_proxy_record.test"
	proxyRecordStatusValid      = "valid"
)

func testAccProxyRecordPreCheck(t *testing.T) {
	t.Helper()
	testAccBasePreCheck(t)
	skipIfNoOrganizationID(t)
}

func testAccProxyRecordDNSPreCheck(t *testing.T) {
	t.Helper()
	testAccProxyRecordPreCheck(t)
	if _, err := loadProxyRecordDNSHarnessConfigFromEnv(); err != nil {
		t.Skipf("Skipping proxy record DNS validation: %v", err)
	}
}

func TestProxyRecord_Basic(t *testing.T) {
	skipIfNotAcceptance(t)

	orgID := getOrganizationID()
	domain := fmt.Sprintf("tf-acc-proxy-basic-%d.invalid", time.Now().UnixNano())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProxyRecordPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProxyRecordConfig(orgID, domain),
				Check:  testAccProxyRecordBaseChecks(orgID, domain),
			},
			{
				ResourceName:      testProxyRecordResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProxyRecordImportStateID(testProxyRecordResourceName),
			},
		},
	})
}

func TestProxyRecord_ValidWithDNS(t *testing.T) {
	skipIfNotAcceptance(t)

	cfg, err := loadProxyRecordDNSHarnessConfigFromEnv()
	if err != nil {
		t.Skipf("Skipping proxy record DNS validation: %v", err)
	}

	ctx := context.Background()
	harness, err := newCloudflareDNSHarness(ctx, cfg)
	if err != nil {
		t.Fatalf("initializing DNS harness: %v", err)
	}

	orgID := getOrganizationID()
	client := httpclient.NewDefaultClient(os.Getenv("POSTHOG_HOST"), os.Getenv("POSTHOG_API_KEY"), "acceptance-test")
	domain := harness.nextDomain("tf-acc-proxy")

	var proxyRecordID string
	var targetCNAME string

	t.Cleanup(testAccProxyRecordDNSCleanup(harness, domain))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccProxyRecordDNSPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProxyRecordConfig(orgID, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProxyRecordBaseChecks(orgID, domain),
					testAccCaptureProxyRecordState(&proxyRecordID, &targetCNAME),
				),
			},
			{
				PreConfig: testAccConfigureProxyRecordDNS(t, harness, domain, &targetCNAME),
				Config:    testAccProxyRecordConfig(orgID, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccWaitForProxyRecordStatusCheck(&client, orgID, &proxyRecordID, proxyRecordStatusValid),
				),
			},
			{
				Config: testAccProxyRecordConfig(orgID, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testProxyRecordResourceName, "status", proxyRecordStatusValid),
					resource.TestCheckResourceAttr(testProxyRecordResourceName, "domain", domain),
					resource.TestCheckResourceAttr(testProxyRecordResourceName, "target_cname", targetCNAME),
				),
			},
		},
	})
}

func waitForProxyRecordStatus(ctx context.Context, client *httpclient.PosthogClient, organizationID, proxyRecordID, expectedStatus string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		record, _, err := client.GetProxyRecord(ctx, organizationID, proxyRecordID)
		if err == nil {
			if record.Status == expectedStatus {
				return nil
			}
			if record.Status == "timed_out" {
				return fmt.Errorf("proxy record reached timed_out before %s", expectedStatus)
			}
		}

		select {
		case <-ctx.Done():
			if err != nil {
				return fmt.Errorf("waiting for proxy record status %s: %w", expectedStatus, err)
			}
			return fmt.Errorf("timed out waiting for proxy record %s to reach %s", proxyRecordID, expectedStatus)
		case <-ticker.C:
		}
	}
}

func testAccProxyRecordImportStateID(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["organization_id"]
		return fmt.Sprintf("%s/%s", orgID, rs.Primary.ID), nil
	}
}

func testAccProxyRecordBaseChecks(orgID, domain string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(testProxyRecordResourceName, "organization_id", orgID),
		resource.TestCheckResourceAttr(testProxyRecordResourceName, "domain", domain),
		resource.TestCheckResourceAttrSet(testProxyRecordResourceName, "id"),
		resource.TestCheckResourceAttrSet(testProxyRecordResourceName, "target_cname"),
		resource.TestCheckResourceAttrSet(testProxyRecordResourceName, "status"),
	)
}

func testAccCaptureProxyRecordState(proxyRecordID, targetCNAME *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[testProxyRecordResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", testProxyRecordResourceName)
		}

		*proxyRecordID = rs.Primary.ID
		*targetCNAME = rs.Primary.Attributes["target_cname"]
		if *proxyRecordID == "" || *targetCNAME == "" {
			return fmt.Errorf("missing proxy_record id or target_cname in state")
		}
		return nil
	}
}

func testAccConfigureProxyRecordDNS(t *testing.T, harness *cloudflareDNSHarness, domain string, targetCNAME *string) func() {
	t.Helper()

	return func() {
		if *targetCNAME == "" {
			t.Fatal("missing target_cname from previous step")
		}

		dnsCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		if err := harness.deleteByName(dnsCtx, domain); err != nil {
			t.Fatalf("cleaning DNS before create: %v", err)
		}
		if _, err := harness.createCNAME(dnsCtx, domain, *targetCNAME); err != nil {
			t.Fatalf("creating DNS CNAME: %v", err)
		}
		if err := harness.waitForPublicCNAME(dnsCtx, domain, *targetCNAME); err != nil {
			t.Fatalf("waiting for public DNS: %v", err)
		}
	}
}

func testAccWaitForProxyRecordStatusCheck(client *httpclient.PosthogClient, organizationID string, proxyRecordID *string, expectedStatus string) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		return waitForProxyRecordStatus(context.Background(), client, organizationID, *proxyRecordID, expectedStatus, proxyRecordValidWaitTimeout)
	}
}

func testAccProxyRecordDNSCleanup(harness *cloudflareDNSHarness, domain string) func() {
	return func() {
		if domain == "" {
			return
		}
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		_ = harness.deleteByName(cleanupCtx, domain)
	}
}

func testAccProxyRecordConfig(orgID, domain string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_proxy_record" "test" {
  organization_id = %q
  domain          = %q
}
`, orgID, domain)
}
