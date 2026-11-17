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

## Embedding the Provider in Go

The provider's packages live under `pkg/`, so other Go modules can import them. Tools that run a Terraform provider in-process, such as Crossplane's upjet or Pulumi's Terraform bridge, start from the constructor:

```go
import "github.com/posthog/terraform-provider/pkg/provider"

p := provider.New("1.0.0")()
```

Only the Terraform interface is versioned: resource and data source names, their schemas, and state. The Go packages are not a stable API and can change in any release, including patch releases, so pin a version and check that your build passes before upgrading.

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

Slack subscription tests also need `POSTHOG_SLACK_INTEGRATION_ID` set to a Slack
integration belonging to `POSTHOG_PROJECT_ID`. They skip when it is unset;
use a non-delivering integration fixture for local contract tests.

### Proxy Record DNS Harness

`posthog_proxy_record` acceptance tests need programmable DNS so they can create a custom domain, point it at the PostHog-returned `target_cname`, and wait for the record to converge to `valid`.

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

The smoke test creates a unique CNAME under `POSTHOG_TEST_DNS_BASE_DOMAIN`, waits for public resolution through the configured resolver, and deletes the record during cleanup. The `posthog_proxy_record` resource tests reuse the same harness and replace the static CNAME target with the provider-computed PostHog `target_cname`.

## Generating Documentation

```shell
make generate
```

## Changesets

A pull request that changes what provider users see adds a changeset: a fix, a new resource, data source or attribute, a breaking change, or an upgrade note. The changeset's summary becomes that change's `CHANGELOG.md` entry, so write it for people who use the provider. Do not edit `CHANGELOG.md` by hand. Pull requests that only touch tests, CI or internal tooling do not need one.

To add one, run the following and commit the file it creates in `.changeset/` with your pull request. It needs Node.js 22.11 or newer and pnpm.

```shell
pnpm install
pnpm changeset
```

You can also write the file by hand, as `.changeset/<any-name>.md`:

```markdown
---
"terraform-provider-posthog": patch
---

**`posthog_insight`:** What changed, and what a user has to do about it.
```

Pick the bump level for the change. It also picks the `CHANGELOG.md` section the entry goes under:

- `patch`: bug fixes, documentation and internal changes, under "Fixes"
- `minor`: new resources, data sources or attributes, under "Features"
- `major`: breaking changes, under "Breaking changes"

## Releasing

Releases are semi-automatic. Merging a pull request with a changeset starts a release, but nothing is published until someone on `@PostHog/client-libraries-approvers` approves it. Approval requests are posted in the [#approvals-client-libraries](https://posthog.slack.com/archives/C0A3UEVDDNF) Slack channel.

Releases come from the changesets merged to `main`. The version lives in `package.json`, which exists only for Changesets. The provider binary takes its version from the release tag.

When a pull request with a changeset merges to `main`, the [Release workflow](.github/workflows/release.yml):

1. Waits for a maintainer to approve the run in the `Release` environment.
2. Runs `pnpm changeset version`, which bumps `package.json`, adds the `CHANGELOG.md` section and deletes the changesets, then commits that to `main`.
3. Tags that commit `vX.Y.Z`.
4. Runs the [Publish workflow](.github/workflows/publish.yml): unit tests, then GoReleaser, which signs the checksums with the registry GPG key and creates the GitHub release with the `CHANGELOG.md` section as its notes. The Terraform Registry picks up the release from GitHub.

Changesets that merge before a release is approved ship together in it. You can also start the Release workflow from the Actions tab. It releases the pending changesets, or, when there are none, the `package.json` version if it has no tag yet.

## License

See [LICENSE](LICENSE).
