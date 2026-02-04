# Import format for project default (all resources of a type):
# project_id/resource_type/default
terraform import posthog_access_control.surveys_project_default 12345/survey/default

# Import format for project default on a specific resource:
# project_id/resource_type/resource_id/default
terraform import posthog_access_control.analytics_dashboard_default 12345/dashboard/999/default

# Import format for role-based access control (resource type level):
# project_id/resource_type/role/role_id
terraform import posthog_access_control.engineering_feature_flags 12345/feature_flag/role/abc-123-def

# Import format for member-based access control (resource type level):
# project_id/resource_type/member/member_id
terraform import posthog_access_control.alice_dashboards 12345/dashboard/member/xyz-456-uvw

# Import format for role-based access control (specific resource):
# project_id/resource_type/resource_id/role/role_id
terraform import posthog_access_control.role_specific_dashboard 12345/dashboard/999/role/abc-123-def

# Import format for member-based access control (specific resource):
# project_id/resource_type/resource_id/member/member_id
terraform import posthog_access_control.alice_specific_dashboard 12345/dashboard/999/member/xyz-456-uvw
