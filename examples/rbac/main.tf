# Complete RBAC Example
#
# This example demonstrates the full role-based access control workflow in PostHog:
# 1. Create a project
# 2. Create roles
# 3. Set restrictive project defaults (deny by default)
# 4. Grant roles access to the project
# 5. Configure fine-grained resource access controls
# 6. Optionally: look up users and assign to roles / specific permissions

variable "organization_id" {
  description = "PostHog organization UUID"
  type        = string
}

variable "project_name" {
  description = "Name for the new project"
  type        = string
  default     = "Production Analytics"
}

# =============================================================================
# Step 1: Create the project
# =============================================================================

resource "posthog_project" "main" {
  organization_id = var.organization_id
  name            = var.project_name
  timezone        = "UTC"
}

# =============================================================================
# Step 2: Create roles
# =============================================================================

resource "posthog_role" "engineering" {
  organization_id = var.organization_id
  name            = "Engineering"
}

resource "posthog_role" "product_managers" {
  organization_id = var.organization_id
  name            = "Product Managers"
}

resource "posthog_role" "support" {
  organization_id = var.organization_id
  name            = "Support"
}

# =============================================================================
# Step 3: Set project default access to "none" (most restrictive)
# Only explicitly granted roles/members will have access
# =============================================================================

resource "posthog_project_default_access" "restrictive" {
  project_id   = posthog_project.main.id
  access_level = "none"
}

# =============================================================================
# Step 4: Grant roles access to the project
# =============================================================================

# Engineering team gets admin access to the project
resource "posthog_project_member" "engineering" {
  project_id   = posthog_project.main.id
  role         = posthog_role.engineering.id
  access_level = "admin"
}

# Product managers get member access
resource "posthog_project_member" "product_managers" {
  project_id   = posthog_project.main.id
  role         = posthog_role.product_managers.id
  access_level = "member"
}

# Support team gets member access
resource "posthog_project_member" "support" {
  project_id   = posthog_project.main.id
  role         = posthog_role.support.id
  access_level = "member"
}

# =============================================================================
# Step 5: Configure fine-grained resource access controls
# Access levels: none, viewer, editor
# =============================================================================

# --- Feature Flags: Only engineering can edit ---

resource "posthog_access_control" "feature_flags_engineering" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.engineering.id
  access_level = "editor"
}

resource "posthog_access_control" "feature_flags_product" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.product_managers.id
  access_level = "viewer"
}

resource "posthog_access_control" "feature_flags_support" {
  project_id   = posthog_project.main.id
  resource     = "feature_flag"
  role         = posthog_role.support.id
  access_level = "none"
}

# --- Dashboards: Only view permissions ---

resource "posthog_access_control" "dashboards_engineering" {
  project_id   = posthog_project.main.id
  resource     = "dashboard"
  role         = posthog_role.engineering.id
  access_level = "viewer"
}

resource "posthog_access_control" "dashboards_support" {
  project_id   = posthog_project.main.id
  resource     = "dashboard"
  role         = posthog_role.support.id
  access_level = "viewer"
}

# --- Experiments: Product managers can edit ---

resource "posthog_access_control" "experiments_product" {
  project_id   = posthog_project.main.id
  resource     = "experiment"
  role         = posthog_role.product_managers.id
  access_level = "editor"
}

resource "posthog_access_control" "experiments_engineering" {
  project_id   = posthog_project.main.id
  resource     = "experiment"
  role         = posthog_role.engineering.id
  access_level = "viewer"
}

# --- Session Recordings: Support and engineering have no access (explicitly configured) ---

resource "posthog_access_control" "recordings_support" {
  project_id   = posthog_project.main.id
  resource     = "session_recording"
  role         = posthog_role.support.id
  access_level = "none"
}

resource "posthog_access_control" "recordings_engineering" {
  project_id   = posthog_project.main.id
  resource     = "session_recording"
  role         = posthog_role.engineering.id
  access_level = "none"
}

# =============================================================================
# Step 6 (Optional): Look up users and assign to roles
# =============================================================================

# Uncomment to use:

# # Look up the user
# data "posthog_user" "alice" {
#   organization_id = var.organization_id
#   email           = "alice@example.com"
# }
#
# # Option A: Add user to a role (inherits role's permissions)
# resource "posthog_role_membership" "alice_engineering" {
#   organization_id = var.organization_id
#   role_id         = posthog_role.engineering.id
#   user_uuid       = data.posthog_user.alice.uuid
# }
#
# # Option B: Give user direct project access (independent of roles)
# resource "posthog_project_member" "alice_direct" {
#   project_id          = posthog_project.main.id
#   organization_member = data.posthog_user.alice.uuid
#   access_level        = "admin"  # or "member", "none"
# }
#
# # Option C: Give user specific resource-type permissions
# resource "posthog_access_control" "alice_feature_flags" {
#   project_id          = posthog_project.main.id
#   resource            = "feature_flag"
#   organization_member = data.posthog_user.alice.uuid
#   access_level        = "editor"
# }
#
# resource "posthog_access_control" "alice_dashboards" {
#   project_id          = posthog_project.main.id
#   resource            = "dashboard"
#   organization_member = data.posthog_user.alice.uuid
#   access_level        = "viewer"
# }
#
# # Option D: Give user access to a specific resource instance
# resource "posthog_access_control" "alice_exec_dashboard" {
#   project_id          = posthog_project.main.id
#   resource            = "dashboard"
#   resource_id         = posthog_dashboard.executive_summary.id  # specific dashboard
#   organization_member = data.posthog_user.alice.uuid
#   access_level        = "editor"
# }


# =============================================================================
# Outputs
# =============================================================================

output "project" {
  description = "Created project details"
  value = {
    id       = posthog_project.main.id
    name     = posthog_project.main.name
    timezone = posthog_project.main.timezone
  }
}

output "roles" {
  description = "Created role IDs"
  value = {
    engineering      = posthog_role.engineering.id
    product_managers = posthog_role.product_managers.id
    support          = posthog_role.support.id
  }
}

output "project_access" {
  description = "Project access configuration"
  value = {
    default_access = posthog_project_default_access.restrictive.access_level
    engineering    = posthog_project_member.engineering.access_level
    product        = posthog_project_member.product_managers.access_level
    support        = posthog_project_member.support.access_level
  }
}
