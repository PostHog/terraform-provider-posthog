# List the flags in a project.
data "posthog_feature_flags" "all" {
  project_id = "12345"
}

output "active_flag_keys" {
  value = [for flag in data.posthog_feature_flags.all.flags : flag.key if flag.active]
}

# Promote flags from dev to staging to production, adopting flags that already
# exist instead of creating duplicates, without writing any flag IDs.
#
# Each environment has its own PostHog project and its own Terraform state, and
# the same configuration is applied to each in turn, for example from CI:
#
#   terraform -chdir=flags init -backend-config=dev.backend.hcl
#   terraform -chdir=flags apply -var environment=dev
#   # ...then the same with staging, then production
#
# A flag created by hand in a project's UI is adopted by the import block on
# that environment's next apply; anywhere it does not exist yet, it is created.
variable "environment" {
  type = string

  validation {
    condition     = contains(["dev", "staging", "production"], var.environment)
    error_message = "environment must be dev, staging or production."
  }
}

locals {
  # Every flag, declared once. A feature ships to dev first and is rolled out
  # further as it moves towards production.
  feature_flags = {
    "new-onboarding" = {
      name    = "New onboarding flow"
      rollout = { dev = 100, staging = 100, production = 25 }
    }
    "beta-dashboard" = {
      name    = "Beta dashboard"
      rollout = { dev = 100, staging = 50, production = 0 }
    }
  }
}

resource "posthog_project" "this" {
  name = var.environment
}

# Looked up by name, because a new environment's project ID is not known until
# apply. Until the project exists this finds no flags, so every flag is created.
data "posthog_feature_flags" "existing" {
  project_name = var.environment
}

resource "posthog_feature_flag" "this" {
  for_each           = local.feature_flags
  project_id         = posthog_project.this.id
  key                = each.key
  name               = each.value.name
  rollout_percentage = each.value.rollout[var.environment]

  # If the environment already had a project with this name before Terraform
  # managed it, import that project first: see project_name above. This turns
  # the replace that would otherwise delete its flags into a plan-time error.
  lifecycle {
    prevent_destroy = true
  }
}

# for_each in an import block needs Terraform 1.7 or later. Once a flag is in
# state its import does nothing, so the block can stay in place.
import {
  for_each = {
    for flag in data.posthog_feature_flags.existing.flags : flag.key => flag.id
    if contains(keys(local.feature_flags), flag.key)
  }
  to = posthog_feature_flag.this[each.key]
  id = "${data.posthog_feature_flags.existing.project_id}/${each.value}"
}
