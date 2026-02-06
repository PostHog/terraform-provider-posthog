# Set the default access level for a project
# Everyone without explicit access will have this level
resource "posthog_project_default_access" "restrictive" {
  project_id   = "your-project-id"
  access_level = "none" # Options: none, member, admin
}

# Use with provider-level project_id
resource "posthog_project_default_access" "default" {
  access_level = "member"
}
