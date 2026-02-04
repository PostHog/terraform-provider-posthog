# Import a role's project access using: project_id/role/<role_id>
terraform import posthog_project_member.engineering_team your-project-id/role/your-role-uuid

# Import an organization member's project access using: project_id/member/<member_id>
terraform import posthog_project_member.alice your-project-id/member/your-member-uuid
