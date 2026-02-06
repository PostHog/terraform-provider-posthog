# Look up a user by their email address
data "posthog_user" "alice" {
  organization_id = "your-organization-uuid"
  email           = "alice@example.com"
}

# Use the user's UUID in other resources
output "alice_uuid" {
  value = data.posthog_user.alice.uuid
}

output "alice_name" {
  value = "${data.posthog_user.alice.first_name} ${data.posthog_user.alice.last_name}"
}
