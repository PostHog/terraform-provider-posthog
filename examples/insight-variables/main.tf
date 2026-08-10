# Insight Variables Example
#
# This example demonstrates PostHog SQL variables (insight variables): typed
# inputs that a HogQL query references as {variables.code_name}. Any dashboard
# holding an insight that uses a variable renders an input for it, which turns a
# static insight into a lookup tool.
#
# The variable's code_name is generated from its name at creation and does not
# change when you rename it, so queries keep working across a rename.

terraform {
  required_providers {
    posthog = {
      source = "posthog/posthog"
    }
  }
}

provider "posthog" {
  # Configuration can be provided via:
  # - Environment variables: POSTHOG_API_KEY, POSTHOG_PROJECT_ID, POSTHOG_HOST
  # - Or explicitly in the provider block:
  # api_key    = "your-api-key"
  # project_id = "12345"
  # host       = "https://us.posthog.com"  # Optional, defaults to US cloud
}

# =============================================================================
# Example 1: Free-text variable, no default
# =============================================================================

resource "posthog_insight_variable" "search" {
  name = "Search"
  type = "String"
}

# =============================================================================
# Example 2: Variable with a default value
#
# Values are JSON-encoded, so a bare scalar needs jsonencode.
# =============================================================================

resource "posthog_insight_variable" "environment" {
  name               = "Environment"
  type               = "String"
  default_value_json = jsonencode("prod")
}

resource "posthog_insight_variable" "lookback_days" {
  name               = "Lookback days"
  type               = "Number"
  default_value_json = jsonencode(30)
}

resource "posthog_insight_variable" "include_internal" {
  name               = "Include internal users"
  type               = "Boolean"
  default_value_json = jsonencode(false)
}

# =============================================================================
# Example 3: List variable, which constrains the input to known options
#
# PostHog stores every entry as a string. Encode numbers as strings too, since
# jsonencode([1, 2]) is returned as ["1", "2"] and reads as drift on the next plan.
# =============================================================================

resource "posthog_insight_variable" "region" {
  name               = "Region"
  type               = "List"
  values_json        = jsonencode(["us", "eu"])
  default_value_json = jsonencode("us")
}

# =============================================================================
# Example 4: An insight driven by the variables
#
# The query's variables map is keyed by variable UUID, and each entry repeats
# the UUID and the code name. Build it from the resources so a rename or a
# re-created variable can't leave a stale reference behind.
# =============================================================================

resource "posthog_insight" "events_by_environment" {
  name        = "Events by environment"
  description = "Scoped by the Environment, Lookback days and Region variables"

  query_json = jsonencode({
    kind = "DataVisualizationNode"
    source = {
      kind = "HogQLQuery"
      query = trimspace(<<-SQL
        SELECT
            properties.environment AS environment,
            count()                AS events
        FROM events
        WHERE timestamp >= now() - toIntervalDay({variables.lookback_days})
          AND properties.environment = {variables.environment}
          AND properties.region = {variables.region}
        GROUP BY environment
        ORDER BY events DESC
      SQL
      )
      variables = {
        for variable in [
          posthog_insight_variable.environment,
          posthog_insight_variable.lookback_days,
          posthog_insight_variable.region,
          ] : variable.id => {
          code_name  = variable.code_name
          variableId = variable.id
        }
      }
    }
  })

  dashboard_ids = [posthog_dashboard.lookup.id]
}

resource "posthog_dashboard" "lookup" {
  name        = "Environment lookup"
  description = "Set the variables at the top of the dashboard to scope every tile"
}

# =============================================================================
# Outputs
# =============================================================================

output "variable_code_names" {
  description = "Code names to reference in HogQL as {variables.code_name}"
  value = {
    search           = posthog_insight_variable.search.code_name
    environment      = posthog_insight_variable.environment.code_name
    lookback_days    = posthog_insight_variable.lookback_days.code_name
    include_internal = posthog_insight_variable.include_internal.code_name
    region           = posthog_insight_variable.region.code_name
  }
}
