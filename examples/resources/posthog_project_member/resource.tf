# Grant a role access to a project
resource "posthog_project_member" "engineering_team" {
  project_id   = "your-project-id"
  role         = posthog_role.engineering.id
  access_level = "member" # Options: none, member, admin
}

# Grant a role admin access to a project
resource "posthog_project_member" "team_leads" {
  project_id   = "your-project-id"
  role         = posthog_role.team_leads.id
  access_level = "admin"
}

# Grant an organization member direct access to a project
resource "posthog_project_member" "alice" {
  project_id          = "your-project-id"
  organization_member = posthog_organization_member.alice.id
  access_level        = "admin"
}
