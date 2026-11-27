# Minimal endpoint. Run it with
# POST /api/projects/{project_id}/endpoints/daily-signups/run/
resource "posthog_endpoint" "daily_signups" {
  name        = "daily-signups"
  description = "Signups per day over the last 30 days"

  query_json = jsonencode({
    kind  = "HogQLQuery"
    query = "SELECT toDate(timestamp) AS day, count() AS signups FROM events WHERE event = 'user signed up' AND timestamp >= now() - INTERVAL 30 DAY GROUP BY day ORDER BY day"
  })
}

# Endpoint with a SQL variable. The query names the variable as
# {variables.code_name} and declares it under `variables`, keyed by the
# variable's id. Building both from the posthog_insight_variable resource
# also makes Terraform create the variable first. Callers override the
# default in the run body: {"variables": {"event_name": "$pageview"}}.
resource "posthog_insight_variable" "event_name" {
  name               = "Event name"
  type               = "String"
  default_value_json = jsonencode("$pageview")
}

resource "posthog_endpoint" "event_count" {
  name                   = "event-count"
  data_freshness_seconds = 3600
  tags                   = ["public-api"]

  query_json = jsonencode({
    kind  = "HogQLQuery"
    query = "SELECT count() AS total FROM events WHERE event = {variables.${posthog_insight_variable.event_name.code_name}} AND timestamp >= now() - INTERVAL 7 DAY"
    variables = {
      (posthog_insight_variable.event_name.id) = {
        variableId = posthog_insight_variable.event_name.id
        code_name  = posthog_insight_variable.event_name.code_name
      }
    }
  })
}

# Materialized endpoint. PostHog precomputes the results every
# data_freshness_seconds instead of running the query on each call.
resource "posthog_endpoint" "weekly_active_users" {
  name                   = "weekly-active-users"
  is_materialized        = true
  data_freshness_seconds = 21600

  query_json = jsonencode({
    kind  = "HogQLQuery"
    query = "SELECT count(DISTINCT person_id) AS users FROM events WHERE timestamp >= now() - INTERVAL 7 DAY"
  })
}
