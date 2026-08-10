# Look up an existing insight by its short_id (the value in the insight's URL).
data "posthog_insight" "existing" {
  short_id = "Ty9UeOA2"
}

# Reference the existing insight's attributes elsewhere in your configuration.
# Exposed: id, short_id, name, derived_name, description, query_json, tags,
# dashboard_ids, and favorited.
output "existing_insight_name" {
  value = data.posthog_insight.existing.name
}

output "existing_insight_tags" {
  value = data.posthog_insight.existing.tags
}

output "existing_insight_dashboards" {
  value = data.posthog_insight.existing.dashboard_ids
}

# Adopt an insight that was created outside Terraform (via the UI or a script)
# into a managed posthog_insight resource, without creating a duplicate.
#
# 1. Author the managed configuration to match the existing insight. Copy the
#    data source's normalized query (data.posthog_insight.existing.query_json)
#    as your starting point, then commit it as config below — a managed
#    resource's query should be its own source of truth, not a live reference
#    to the data source (which you remove after adoption).
# 2. The one-time import block below tells Terraform to adopt the existing
#    insight (resolved by the data source's numeric id) instead of creating a
#    new one on the first apply.
# 3. After the first successful apply, remove the import block. The insight is
#    now fully managed by this resource.
resource "posthog_insight" "adopted" {
  name        = data.posthog_insight.existing.name
  description = data.posthog_insight.existing.description

  # Author the query to match the existing insight (the data source exposes the
  # current query as data.posthog_insight.existing.query_json for reference).
  query_json = jsonencode({
    kind = "InsightVizNode"
    source = {
      kind = "TrendsQuery"
      series = [{
        kind  = "EventsNode"
        event = "$pageview"
        math  = "total"
      }]
    }
  })
}

import {
  to = posthog_insight.adopted
  id = data.posthog_insight.existing.id
}
