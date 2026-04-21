package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func skipIfNoProxyRecordDNSHarness(t *testing.T) proxyRecordDNSHarnessConfig {
	t.Helper()

	cfg, err := loadProxyRecordDNSHarnessConfigFromEnv()
	if err != nil {
		t.Skipf("Skipping proxy_record DNS harness test: %v", err)
	}

	return cfg
}

func TestProxyRecordDNSHarness_CloudflareCNAME(t *testing.T) {
	skipIfNotAcceptance(t)

	cfg := skipIfNoProxyRecordDNSHarness(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.PropagationTimeout+time.Minute)
	defer cancel()

	harness, err := newCloudflareDNSHarness(ctx, cfg)
	if err != nil {
		t.Fatalf("failed to initialize DNS harness: %v", err)
	}

	fqdn := harness.nextDomain("proxy-record")

	if err := harness.deleteByName(ctx, fqdn); err != nil {
		t.Fatalf("failed to ensure clean DNS name: %v", err)
	}

	record, err := harness.createCNAME(ctx, fqdn, testProxyRecordSmokeCNAMETarget)
	if err != nil {
		t.Fatalf("failed to create CNAME record: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		_ = harness.deleteRecord(cleanupCtx, record.ID)
	})

	if got := normalizeFQDN(record.Name); got != normalizeFQDN(fqdn) {
		t.Fatalf("unexpected record name: got %s want %s", got, normalizeFQDN(fqdn))
	}

	if err := harness.waitForPublicCNAME(ctx, fqdn, testProxyRecordSmokeCNAMETarget); err != nil {
		t.Fatalf("failed waiting for public CNAME propagation: %v", err)
	}
}

func TestCloudflareDNSHarness_CreateAndDeleteCNAME(t *testing.T) {
	const (
		expectedZoneID = "zone-123"
		expectedRecord = "record-456"
	)

	var seen []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, fmt.Sprintf("%s %s?%s", r.Method, r.URL.Path, r.URL.RawQuery))
		if got := r.Header.Get("Authorization"); got != "Bearer token-123" {
			t.Fatalf("unexpected auth header: %s", got)
		}

		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/zones":
			_, _ = w.Write([]byte(`{"success":true,"result":[{"id":"zone-123"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/zones/zone-123/dns_records":
			_, _ = w.Write([]byte(`{"success":true,"result":{"id":"record-456","type":"CNAME","name":"proxy-record.tfacc.example.com","content":"example.com"}}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/zones/zone-123/dns_records/record-456":
			_, _ = w.Write([]byte(`{"success":true,"result":{}}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	cfg := proxyRecordDNSHarnessConfig{
		Provider:           "cloudflare",
		ZoneName:           "example.com",
		BaseDomain:         "tfacc.example.com",
		ResolverAddr:       "1.1.1.1:53",
		PropagationTimeout: 10 * time.Second,
		CloudflareAPIToken: "token-123",
		CloudflareBaseURL:  server.URL,
	}

	ctx := context.Background()
	harness, err := newCloudflareDNSHarness(ctx, cfg)
	if err != nil {
		t.Fatalf("failed to create harness: %v", err)
	}

	record, err := harness.createCNAME(ctx, "proxy-record.tfacc.example.com", testProxyRecordSmokeCNAMETarget)
	if err != nil {
		t.Fatalf("failed to create record: %v", err)
	}

	if record.ID != expectedRecord {
		t.Fatalf("unexpected record ID: got %s want %s", record.ID, expectedRecord)
	}

	if err := harness.deleteRecord(ctx, record.ID); err != nil {
		t.Fatalf("failed to delete record: %v", err)
	}

	expected := []string{
		"GET /zones?name=example.com",
		"POST /zones/zone-123/dns_records?",
		"DELETE /zones/zone-123/dns_records/record-456?",
	}
	if len(seen) != len(expected) {
		t.Fatalf("unexpected number of requests: got %d want %d (%v)", len(seen), len(expected), seen)
	}
	for idx := range expected {
		if seen[idx] != expected[idx] {
			t.Fatalf("unexpected request[%d]: got %q want %q", idx, seen[idx], expected[idx])
		}
	}
}

func TestCloudflareDNSHarness_WaitForPublicCNAME(t *testing.T) {
	harness := &cloudflareDNSHarness{
		propagationTimeout: 200 * time.Millisecond,
		retryInterval:      10 * time.Millisecond,
	}

	attempts := 0
	harness.lookupCNAME = func(_ context.Context, _ string) (string, error) {
		attempts++
		if attempts < 3 {
			return "", fmt.Errorf("not ready")
		}
		return normalizeFQDN(testProxyRecordSmokeCNAMETarget), nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := harness.waitForPublicCNAME(ctx, "proxy-record.tfacc.example.com", testProxyRecordSmokeCNAMETarget); err != nil {
		t.Fatalf("expected resolver to converge: %v", err)
	}
}
