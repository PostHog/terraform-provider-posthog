variable "stripe_secret_key" {
  type      = string
  sensitive = true
}

variable "pg_password" {
  type      = string
  sensitive = true
}

# Stripe warehouse source
resource "posthog_external_data_source" "stripe" {
  source_type    = "Stripe"
  prefix         = "stripe_"
  sync_frequency = "6hour"
  schemas        = ["charges", "customers", "invoices"]

  job_inputs_json = jsonencode({
    stripe_account_id = "acct_123"
    stripe_secret_key = var.stripe_secret_key
  })
}

# Postgres warehouse source
resource "posthog_external_data_source" "prod_pg" {
  source_type    = "Postgres"
  sync_frequency = "day"
  schemas        = ["users", "orders"]

  job_inputs_json = jsonencode({
    host     = "db.internal"
    port     = 5432
    database = "app"
    user     = "readonly"
    password = var.pg_password
    schema   = "public"
  })
}
