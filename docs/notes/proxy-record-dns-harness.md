# Proxy Record DNS Harness

## Purpose

`posthog_proxy_record` cannot be validated safely upstream unless the test suite can control a disposable DNS zone. The provider only learns `target_cname` after creating the proxy record, so a real acceptance flow must be able to publish a matching CNAME before PostHog can converge the record to `valid`.

This note defines the harness introduced by `KEF-179`.

## Scope

The harness does not implement `posthog_proxy_record`.

It only proves the DNS half of the future acceptance workflow:

1. generate a unique disposable FQDN
2. create a matching CNAME in a programmable zone
3. wait until that CNAME resolves publicly
4. delete the DNS record during cleanup

Once that loop is stable, a later proxy record resource test can add the PostHog side around it:

1. create `posthog_proxy_record`
2. read `target_cname`
3. publish a CNAME to `target_cname`
4. retry or wait until PostHog reports `valid`
5. destroy both the PostHog record and the DNS record

## Current Provider Support

The initial supported DNS provider is Cloudflare.

Required environment variables:

- `CLOUDFLARE_API_TOKEN`
- `POSTHOG_TEST_DNS_PROVIDER=cloudflare`
- `POSTHOG_TEST_DNS_ZONE_NAME`
- `POSTHOG_TEST_DNS_BASE_DOMAIN`

Optional:

- `POSTHOG_TEST_DNS_RESOLVER` default `1.1.1.1:53`
- `POSTHOG_TEST_DNS_PROPAGATION_TIMEOUT` default `3m`

`POSTHOG_TEST_DNS_BASE_DOMAIN` should be a delegated disposable namespace inside the configured zone, for example `tfacc.example.com` inside the `example.com` zone.

## Runnable Entry Point

Use:

```shell
make testacc-proxy-dns
```

This runs `TestProxyRecordDNSHarness_CloudflareCNAME`, which:

- creates a unique CNAME pointing to `example.com`
- waits for public resolution via the configured resolver
- deletes the record during cleanup

## Why This Is Enough For Now

This ticket exists to remove the main blocker identified in `KEF-177`: no deterministic DNS automation path.

After this harness exists, the remaining work to implement `posthog_proxy_record` becomes straightforward engineering:

- add the organization-scoped resource
- reuse the DNS harness in acceptance tests
- wire the dynamic PostHog `target_cname`
- decide whether create should wait directly or call `retry` after propagation
