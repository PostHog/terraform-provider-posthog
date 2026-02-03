# Import format for resource-type level access control:
# project_id/resource_type/role/role_id
terraform import posthog_access_control.engineering_feature_flags 12345/feature_flags/role/abc-123-def

# Import format for resource-type level access control with organization member:
# project_id/resource_type/member/member_id
terraform import posthog_access_control.alice_dashboards 12345/dashboards/member/xyz-456-uvw

# Import format for resource-instance level access control:
# project_id/resource_type/resource_id/role/role_id
terraform import posthog_access_control.role_specific_dashboard 12345/dashboards/999/role/abc-123-def

# Import format for resource-instance level access control with organization member:
# project_id/resource_type/resource_id/member/member_id
terraform import posthog_access_control.alice_specific_dashboard 12345/dashboards/999/member/xyz-456-uvw
