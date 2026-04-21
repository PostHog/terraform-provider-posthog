package tests

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	envTestDNSProvider           = "POSTHOG_TEST_DNS_PROVIDER"
	envTestDNSZoneName           = "POSTHOG_TEST_DNS_ZONE_NAME"
	envTestDNSBaseDomain         = "POSTHOG_TEST_DNS_BASE_DOMAIN"
	envTestDNSResolver           = "POSTHOG_TEST_DNS_RESOLVER"
	envTestDNSPropagationTimeout = "POSTHOG_TEST_DNS_PROPAGATION_TIMEOUT"
	envCloudflareAPIToken        = "CLOUDFLARE_API_TOKEN"

	missingEnvTemplate             = "%s must be set"
	cloudflareRequestUnsuccessful  = "cloudflare API request was not successful"
	testProxyRecordSmokeCNAMETarget = "example.com"
)

type proxyRecordDNSHarnessConfig struct {
	Provider           string
	ZoneName           string
	BaseDomain         string
	ResolverAddr       string
	PropagationTimeout time.Duration

	CloudflareAPIToken string
	CloudflareBaseURL  string
}

type cloudflareEnvelope[T any] struct {
	Success bool `json:"success"`
	Result  T    `json:"result"`
	Errors  []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type cloudflareZone struct {
	ID string `json:"id"`
}

type cloudflareDNSRecord struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type cloudflareDNSHarness struct {
	baseURL            string
	token              string
	zoneID             string
	zoneName           string
	baseDomain         string
	resolverAddr       string
	propagationTimeout time.Duration
	retryInterval      time.Duration
	httpClient         *http.Client
	lookupCNAME        func(context.Context, string) (string, error)
}

func loadProxyRecordDNSHarnessConfigFromEnv() (proxyRecordDNSHarnessConfig, error) {
	cfg := proxyRecordDNSHarnessConfig{
		Provider:           strings.ToLower(strings.TrimSpace(os.Getenv(envTestDNSProvider))),
		ZoneName:           trimTrailingDot(os.Getenv(envTestDNSZoneName)),
		BaseDomain:         trimTrailingDot(os.Getenv(envTestDNSBaseDomain)),
		ResolverAddr:       strings.TrimSpace(os.Getenv(envTestDNSResolver)),
		CloudflareAPIToken: strings.TrimSpace(os.Getenv(envCloudflareAPIToken)),
	}

	if cfg.Provider == "" {
		cfg.Provider = "cloudflare"
	}
	if cfg.ResolverAddr == "" {
		cfg.ResolverAddr = "1.1.1.1:53"
	}
	cfg.PropagationTimeout = 3 * time.Minute
	if raw := strings.TrimSpace(os.Getenv(envTestDNSPropagationTimeout)); raw != "" {
		timeout, err := time.ParseDuration(raw)
		if err != nil {
			return cfg, fmt.Errorf("%s must be a valid duration: %w", envTestDNSPropagationTimeout, err)
		}
		cfg.PropagationTimeout = timeout
	}

	switch {
	case cfg.Provider != "cloudflare":
		return cfg, fmt.Errorf("%s=%q is not supported", envTestDNSProvider, cfg.Provider)
	case cfg.ZoneName == "":
		return cfg, fmt.Errorf(missingEnvTemplate, envTestDNSZoneName)
	case cfg.BaseDomain == "":
		return cfg, fmt.Errorf(missingEnvTemplate, envTestDNSBaseDomain)
	case !isSubdomainOf(cfg.BaseDomain, cfg.ZoneName):
		return cfg, fmt.Errorf("%s must be within zone %s", envTestDNSBaseDomain, cfg.ZoneName)
	case cfg.CloudflareAPIToken == "":
		return cfg, fmt.Errorf(missingEnvTemplate, envCloudflareAPIToken)
	}

	return cfg, nil
}

func newCloudflareDNSHarness(ctx context.Context, cfg proxyRecordDNSHarnessConfig) (*cloudflareDNSHarness, error) {
	baseURL := cfg.CloudflareBaseURL
	if baseURL == "" {
		baseURL = "https://api.cloudflare.com/client/v4"
	}

	h := &cloudflareDNSHarness{
		baseURL:            strings.TrimRight(baseURL, "/"),
		token:              cfg.CloudflareAPIToken,
		zoneName:           cfg.ZoneName,
		baseDomain:         cfg.BaseDomain,
		resolverAddr:       cfg.ResolverAddr,
		propagationTimeout: cfg.PropagationTimeout,
		retryInterval:      5 * time.Second,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	h.lookupCNAME = newPublicCNAMEResolver(h.resolverAddr)

	zoneID, err := h.lookupZoneID(ctx)
	if err != nil {
		return nil, err
	}
	h.zoneID = zoneID

	return h, nil
}

func (h *cloudflareDNSHarness) lookupZoneID(ctx context.Context) (string, error) {
	path := fmt.Sprintf("%s/zones?name=%s", h.baseURL, url.QueryEscape(h.zoneName))

	var envelope cloudflareEnvelope[[]cloudflareZone]
	if err := h.doJSON(ctx, http.MethodGet, path, nil, &envelope); err != nil {
		return "", err
	}
	if len(envelope.Result) == 0 {
		return "", fmt.Errorf("cloudflare zone not found: %s", h.zoneName)
	}

	return envelope.Result[0].ID, nil
}

func (h *cloudflareDNSHarness) createCNAME(ctx context.Context, fqdn, target string) (cloudflareDNSRecord, error) {
	var envelope cloudflareEnvelope[cloudflareDNSRecord]
	path := fmt.Sprintf("%s/zones/%s/dns_records", h.baseURL, h.zoneID)
	payload := map[string]any{
		"type":    "CNAME",
		"name":    trimTrailingDot(fqdn),
		"content": trimTrailingDot(target),
		"ttl":     60,
		"proxied": false,
	}

	if err := h.doJSON(ctx, http.MethodPost, path, payload, &envelope); err != nil {
		return cloudflareDNSRecord{}, err
	}

	return envelope.Result, nil
}

func (h *cloudflareDNSHarness) listCNAMERecords(ctx context.Context, fqdn string) ([]cloudflareDNSRecord, error) {
	var envelope cloudflareEnvelope[[]cloudflareDNSRecord]
	path := fmt.Sprintf(
		"%s/zones/%s/dns_records?type=CNAME&name=%s",
		h.baseURL,
		h.zoneID,
		url.QueryEscape(trimTrailingDot(fqdn)),
	)

	if err := h.doJSON(ctx, http.MethodGet, path, nil, &envelope); err != nil {
		return nil, err
	}

	return envelope.Result, nil
}

func (h *cloudflareDNSHarness) deleteRecord(ctx context.Context, recordID string) error {
	path := fmt.Sprintf("%s/zones/%s/dns_records/%s", h.baseURL, h.zoneID, recordID)
	return h.doJSON(ctx, http.MethodDelete, path, nil, nil)
}

func (h *cloudflareDNSHarness) deleteByName(ctx context.Context, fqdn string) error {
	records, err := h.listCNAMERecords(ctx, fqdn)
	if err != nil {
		return err
	}

	for _, record := range records {
		if err := h.deleteRecord(ctx, record.ID); err != nil {
			return err
		}
	}

	return nil
}

func (h *cloudflareDNSHarness) waitForPublicCNAME(ctx context.Context, fqdn, target string) error {
	deadline := time.Now().Add(h.propagationTimeout)
	expected := normalizeFQDN(target)

	for {
		cname, err := h.lookupCNAME(ctx, fqdn)
		if err == nil && normalizeFQDN(cname) == expected {
			return nil
		}

		if time.Now().After(deadline) {
			if err != nil {
				return fmt.Errorf("dns propagation timed out for %s -> %s: %w", fqdn, expected, err)
			}
			return fmt.Errorf("dns propagation timed out for %s -> %s (last seen %s)", fqdn, expected, cname)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(h.retryInterval):
		}
	}
}

func (h *cloudflareDNSHarness) nextDomain(prefix string) string {
	label := sanitizeDNSLabel(prefix)
	if label == "" {
		label = "proxyhog"
	}

	return fmt.Sprintf("%s-%s.%s", label, randomHex(4), h.baseDomain)
}

func (h *cloudflareDNSHarness) doJSON(ctx context.Context, method, requestURL string, body any, dest any) error {
	bodyReader, err := marshalRequestBody(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, requestURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+h.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("cloudflare API error (status %d): %s", resp.StatusCode, strings.TrimSpace(string(respBody)))
	}

	if dest == nil {
		return nil
	}

	if err := json.Unmarshal(respBody, dest); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return checkCloudflareSuccess(dest)
}

func marshalRequestBody(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	return strings.NewReader(string(bodyBytes)), nil
}

func checkCloudflareSuccess(dest any) error {
	switch envelope := dest.(type) {
	case *cloudflareEnvelope[[]cloudflareZone]:
		if !envelope.Success {
			return fmt.Errorf(cloudflareRequestUnsuccessful)
		}
	case *cloudflareEnvelope[cloudflareDNSRecord]:
		if !envelope.Success {
			return fmt.Errorf(cloudflareRequestUnsuccessful)
		}
	case *cloudflareEnvelope[[]cloudflareDNSRecord]:
		if !envelope.Success {
			return fmt.Errorf(cloudflareRequestUnsuccessful)
		}
	}

	return nil
}

func newPublicCNAMEResolver(resolverAddr string) func(context.Context, string) (string, error) {
	dialer := &net.Dialer{Timeout: 5 * time.Second}
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, _ string) (net.Conn, error) {
			return dialer.DialContext(ctx, "udp", resolverAddr)
		},
	}

	return func(ctx context.Context, fqdn string) (string, error) {
		return resolver.LookupCNAME(ctx, trimTrailingDot(fqdn))
	}
}

func trimTrailingDot(value string) string {
	return strings.TrimSuffix(strings.TrimSpace(value), ".")
}

func normalizeFQDN(value string) string {
	return strings.ToLower(trimTrailingDot(value)) + "."
}

func isSubdomainOf(subdomain, zone string) bool {
	subdomain = strings.ToLower(trimTrailingDot(subdomain))
	zone = strings.ToLower(trimTrailingDot(zone))
	return subdomain == zone || strings.HasSuffix(subdomain, "."+zone)
}

func sanitizeDNSLabel(raw string) string {
	raw = strings.ToLower(raw)
	var b strings.Builder
	lastHyphen := false
	for _, r := range raw {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			lastHyphen = false
		case !lastHyphen:
			b.WriteRune('-')
			lastHyphen = true
		}
	}
	return strings.Trim(b.String(), "-")
}

func randomHex(numBytes int) string {
	buf := make([]byte, numBytes)
	if _, err := rand.Read(buf); err != nil {
		panic(fmt.Errorf("failed to generate random suffix: %w", err))
	}
	return hex.EncodeToString(buf)
}
