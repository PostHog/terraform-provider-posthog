# Terraform Provider for PostHog

Terraform provider for managing [PostHog](https://posthog.com) resources.

## Documentation

For usage documentation and supported resources, see the [Terraform Registry](https://registry.terraform.io/providers/posthog/posthog/latest/docs).

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.24

## Building the Provider

```shell
go install
```

## Local Development

The `playground/` directory lets you test provider changes locally without publishing.

### Setup

1. Create your Terraform config in `playground/` (e.g., `playground/demo.tf`)

2. Build the provider and run terraform:

```shell
# Plan changes
make playground-plan

# Apply changes
make playground-apply
```

This builds the provider binary and configures Terraform to use your local build via `dev_overrides` - no `terraform init` required.

### Manual Setup

If you prefer to test outside the playground directory:

```shell
# Build the provider
make playground-binary

# Point Terraform to your local build
export TF_CLI_CONFIG_FILE=/path/to/terraform-provider-posthog/playground/terraformrc

# Run terraform commands in any directory
terraform plan
terraform apply
```

### Cleanup

```shell
make playground-clean
```

## Testing

### Unit Tests

```shell
go test ./...
```

### Acceptance Tests

Acceptance tests run against a real PostHog instance and create actual resources:

```shell
export POSTHOG_API_KEY="your-api-key"
export POSTHOG_PROJECT_ID="12345"
export POSTHOG_HOST="https://us.posthog.com"
export POSTHOG_ORGANIZATION_ID="your-org-uuid"   # Default for organization-scoped resources
export POSTHOG_TEST_USER_EMAIL="user@example.com" # Email of existing org member for membership tests (not the one who owns the token)

make testacc
```

### Proxy Record DNS Harness

Future `posthog_proxy_record` acceptance tests need programmable DNS so they can create a custom domain, point it at the PostHog-returned `target_cname`, and wait for the record to converge to `valid`.

The repo now includes a dedicated smoke test for that DNS leg, currently with Cloudflare support:

```shell
export CLOUDFLARE_API_TOKEN="your-cloudflare-token"
export POSTHOG_TEST_DNS_PROVIDER="cloudflare"
export POSTHOG_TEST_DNS_ZONE_NAME="example.com"
export POSTHOG_TEST_DNS_BASE_DOMAIN="tfacc.example.com"

# Optional overrides
export POSTHOG_TEST_DNS_RESOLVER="1.1.1.1:53"
export POSTHOG_TEST_DNS_PROPAGATION_TIMEOUT="3m"

make testacc-proxy-dns
```

The smoke test creates a unique CNAME under `POSTHOG_TEST_DNS_BASE_DOMAIN`, waits for public resolution through the configured resolver, and deletes the record during cleanup. The later `posthog_proxy_record` resource tests should reuse the same harness and replace the static CNAME target with the provider-computed PostHog `target_cname`.

## Generating Documentation

```shell
make generate
```

## Releasing

Releases are automated via GoReleaser when a signed tag is pushed. The Makefile provides convenience targets:

```shell
# Alpha releases (pre-release, for early testing)
make release-alpha VERSION=1.0.0 NUM=1  # creates v1.0.0-alpha.1

# Beta releases (pre-release, feature complete)
make release-beta VERSION=1.0.0 NUM=1   # creates v1.0.0-beta.1

# Stable releases
make release VERSION=1.0.0              # creates v1.0.0
```

**Requirements:**
- GPG key configured for signing (`git tag -s`)
- GPG key added to your GitHub account (for the "Verified" badge)

Pre-release versions (alpha, beta) won't be installed by default - users must explicitly pin to them in their Terraform configuration.

## License

See [LICENSE](LICENSE).
