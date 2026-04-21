# Create a custom-domain proxy record for an organization
resource "posthog_proxy_record" "analytics" {
  organization_id = "your-organization-uuid"
  domain          = "analytics.example.com"
}

# Point your DNS provider at the generated target_cname
output "analytics_proxy_target_cname" {
  value = posthog_proxy_record.analytics.target_cname
}
