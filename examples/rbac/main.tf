# Complete RBAC Example
#
# This example demonstrates how to set up role-based access control in PostHog:
# 1. Look up existing users by email
# 2. Manage their organization membership levels
# 3. Create roles
# 4. Assign users to roles
# 5. Grant roles access to resources via access controls

variable "organization_id" {
  description = "PostHog organization UUID"
  type        = string
}

# =============================================================================
# Step 1: Look up users by email
# =============================================================================

data "posthog_user" "alice" {
  organization_id = var.organization_id
  email           = "alice@example.com"
}

data "posthog_user" "bob" {
  organization_id = var.organization_id
  email           = "bob@example.com"
}

# =============================================================================
# Step 2: Manage organization membership levels
# =============================================================================

# Make Alice an admin
resource "posthog_organization_member" "alice" {
  organization_id = var.organization_id
  user_uuid       = data.posthog_user.alice.uuid
  level           = "admin"
}

# Bob stays as a regular member
resource "posthog_organization_member" "bob" {
  organization_id = var.organization_id
  user_uuid       = data.posthog_user.bob.uuid
  level           = "member"
}

# =============================================================================
# Step 3: Create roles
# =============================================================================

resource "posthog_role" "engineering" {
  organization_id = var.organization_id
  name            = "Engineering"
}

resource "posthog_role" "support" {
  organization_id = var.organization_id
  name            = "Support"
}

# =============================================================================
# Step 4: Assign users to roles
# =============================================================================

# Alice is on the Engineering team
resource "posthog_role_membership" "alice_engineering" {
  organization_id = var.organization_id
  role_id         = posthog_role.engineering.id
  user_uuid       = data.posthog_user.alice.uuid
}

# Bob is on both Engineering and Support teams
resource "posthog_role_membership" "bob_engineering" {
  organization_id = var.organization_id
  role_id         = posthog_role.engineering.id
  user_uuid       = data.posthog_user.bob.uuid
}

resource "posthog_role_membership" "bob_support" {
  organization_id = var.organization_id
  role_id         = posthog_role.support.id
  user_uuid       = data.posthog_user.bob.uuid
}

# =============================================================================
# Step 5: Configure access controls for roles
# =============================================================================

# Engineering team can edit feature flags
resource "posthog_access_control" "engineering_feature_flags" {
  resource     = "feature_flags"
  access_level = "editor"
  role         = posthog_role.engineering.id
}

# Engineering team can edit experiments
resource "posthog_access_control" "engineering_experiments" {
  resource     = "experiments"
  access_level = "editor"
  role         = posthog_role.engineering.id
}

# Support team can view dashboards
resource "posthog_access_control" "support_dashboards" {
  resource     = "dashboards"
  access_level = "viewer"
  role         = posthog_role.support.id
}

# Support team can view session recordings
resource "posthog_access_control" "support_recordings" {
  resource     = "session_recordings"
  access_level = "viewer"
  role         = posthog_role.support.id
}

# Support team cannot access feature flags (explicit deny)
resource "posthog_access_control" "support_no_feature_flags" {
  resource     = "feature_flags"
  access_level = "none"
  role         = posthog_role.support.id
}

# =============================================================================
# Step 6: Grant individual users access to specific resources (optional)
# =============================================================================

# Give Bob editor access to a specific dashboard (overrides role permissions)
resource "posthog_access_control" "bob_specific_dashboard" {
  resource            = "dashboards"
  resource_id         = "12345" # specific dashboard ID
  access_level        = "editor"
  organization_member = posthog_organization_member.bob.id
}

# =============================================================================
# Outputs
# =============================================================================

output "engineering_role_id" {
  description = "Engineering role ID"
  value       = posthog_role.engineering.id
}

output "support_role_id" {
  description = "Support role ID"
  value       = posthog_role.support.id
}
