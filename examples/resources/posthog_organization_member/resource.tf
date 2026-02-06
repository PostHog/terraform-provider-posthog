# Look up an existing user by email
data "posthog_user" "alice" {
  organization_id = "your-organization-uuid"
  email           = "alice@example.com"
}

# Manage Alice's organization membership level
# Note: The user must already be a member of the organization (e.g., via invite)
resource "posthog_organization_member" "alice" {
  organization_id = "your-organization-uuid"
  user_uuid       = data.posthog_user.alice.uuid
  level           = "admin" # or "member" / "owner"

  # Set to true to keep the user in the organization when this resource is destroyed.
  # By default, destroying this resource removes the user from the organization entirely.
  retain_on_destroy = true
}
