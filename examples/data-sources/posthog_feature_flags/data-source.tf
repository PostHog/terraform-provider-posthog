# List the flags in a project.
data "posthog_feature_flags" "all" {
  project_id = "12345"
}

output "flag_keys" {
  value = [for flag in data.posthog_feature_flags.all.flags : flag.key]
}

# Adopt flags that already exist, and create the rest, without writing any IDs.
#
# Each environment has its own project, and every project declares the same
# flags. Where a flag already exists (for example, created in the PostHog UI),
# the import block adopts it; where it does not, Terraform creates it. Once a
# flag is in state its import is a no-op, so the block can stay in place.
#
# The project is looked up by name rather than by ID, because a new
# environment's project ID is not known until apply. Until that project exists
# the data source finds no flags, so nothing is imported and every flag is
# created.
#
# If the environment already has a project with this name, import that project
# into posthog_project.this first instead of creating a second one. Otherwise
# the name matches the old project, and adopting its flags into resources that
# point at the new project plans to replace them. Terraform deletes the old
# flags before PostHog refuses the second project (names are unique within an
# organization), so the flags are lost even though the apply fails.
# prevent_destroy on the flags turns that into an error at plan time.
variable "environment" {
  type = string
}

locals {
  feature_flags = {
    "new-onboarding" = "New onboarding flow"
    "beta-dashboard" = "Beta dashboard"
  }
}

resource "posthog_project" "this" {
  name = var.environment
}

data "posthog_feature_flags" "existing" {
  project_name = var.environment
}

resource "posthog_feature_flag" "this" {
  for_each   = local.feature_flags
  project_id = posthog_project.this.id
  key        = each.key
  name       = each.value

  lifecycle {
    prevent_destroy = true
  }
}

# for_each in an import block needs Terraform 1.7 or later.
import {
  for_each = {
    for flag in data.posthog_feature_flags.existing.flags : flag.key => flag.id
    if contains(keys(local.feature_flags), flag.key)
  }
  to = posthog_feature_flag.this[each.key]
  id = "${data.posthog_feature_flags.existing.project_id}/${each.value}"
}
