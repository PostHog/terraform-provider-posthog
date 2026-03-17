# Look up an externally created role by name
data "posthog_role" "engineers" {
  organization_id = "your-organization-uuid"
  name            = "Engineers"
}

# Use the role ID to grant project access
resource "posthog_project_member" "engineers" {
  project_id   = "your-project-id"
  role         = data.posthog_role.engineers.id
  access_level = "member"
}
