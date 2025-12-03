// Package tests contains comprehensive integration tests for the PostHog Terraform provider.
// These tests require a running PostHog instance and valid API credentials.
//
// Run with: make testacc-integration
// Or: POSTHOG_API_KEY=xxx POSTHOG_PROJECT_ID=xxx POSTHOG_HOST=xxx TF_ACC=1 go test -v -timeout 30m ./tests/...
package tests

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/posthog/terraform-provider/internal/provider"
)

// testAccProtoV6ProviderFactories is used to instantiate a provider during acceptance testing.
// The provider is registered as "posthog" to match the terraform-provider-posthog naming.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"posthog": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// testAccPreCheck validates required environment variables before running tests.
func testAccPreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("POSTHOG_API_KEY") == "" {
		t.Fatal("POSTHOG_API_KEY must be set for acceptance tests")
	}
	if os.Getenv("POSTHOG_PROJECT_ID") == "" {
		t.Fatal("POSTHOG_PROJECT_ID must be set for acceptance tests")
	}
	if os.Getenv("POSTHOG_HOST") == "" {
		t.Fatal("POSTHOG_HOST must be set for acceptance tests")
	}
}

// skipIfNotAcceptance skips the test if TF_ACC is not set.
func skipIfNotAcceptance(t *testing.T) {
	t.Helper()
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless TF_ACC=1")
	}
}
