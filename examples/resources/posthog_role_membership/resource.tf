# First, look up the user by email
data "posthog_user" "alice" {
  organization_id = "your-organization-uuid"
  email           = "alice@example.com"
}

# Create a role
resource "posthog_role" "engineering" {
  organization_id = "your-organization-uuid"
  name            = "Engineering"
}

# Assign Alice to the Engineering role
resource "posthog_role_membership" "alice_engineering" {
  organization_id = "your-organization-uuid"
  role_id         = posthog_role.engineering.id
  user_uuid       = data.posthog_user.alice.uuid
}
