# Minimal variable - a free-text input with no default
resource "posthog_insight_variable" "lookup" {
  name = "Lookup"
  type = "String"
}

# String variable with a default value. Scalars are JSON-encoded, so wrap them.
resource "posthog_insight_variable" "environment" {
  name               = "Environment"
  type               = "String"
  default_value_json = jsonencode("prod")
}

# Number variable
resource "posthog_insight_variable" "lookback_days" {
  name               = "Lookback days"
  type               = "Number"
  default_value_json = jsonencode(30)
}

# List variable. PostHog stores every entry as a string, so encode numbers as
# strings too - jsonencode([1, 2]) comes back as ["1", "2"] and shows as drift.
resource "posthog_insight_variable" "region" {
  name               = "Region"
  type               = "List"
  values_json        = jsonencode(["us", "eu"])
  default_value_json = jsonencode("us")
}

# Reference a variable from an insight query as {variables.code_name}. Because
# the query holds both the variable's UUID and its code name, build the
# variables map from the resource rather than hardcoding either.
resource "posthog_insight" "events_by_environment" {
  name = "Events by environment"

  query_json = jsonencode({
    kind = "DataVisualizationNode"
    source = {
      kind  = "HogQLQuery"
      query = "SELECT count() FROM events WHERE timestamp >= now() - INTERVAL 1 DAY AND properties.environment = {variables.environment}"
      variables = {
        (posthog_insight_variable.environment.id) = {
          code_name  = posthog_insight_variable.environment.code_name
          variableId = posthog_insight_variable.environment.id
          value      = "prod"
        }
      }
    }
  })
}
