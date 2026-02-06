# Create a role for the Engineering team
resource "posthog_role" "engineering" {
  organization_id = "your-organization-uuid"
  name            = "Engineering"
}

# Create a role for the Support team
resource "posthog_role" "support" {
  organization_id = "your-organization-uuid"
  name            = "Support"
}
