variable "s3_access_key" {
  type      = string
  sensitive = true
}

variable "s3_access_secret" {
  type      = string
  sensitive = true
}

# Parquet files in an S3 bucket, read as one table.
# `format` and `url_pattern` are immutable: a change replaces the table,
# so PostHog infers the columns again from the new files.
resource "posthog_warehouse_table" "orders" {
  name        = "s3_orders"
  format      = "Parquet"
  url_pattern = "https://my-bucket.s3.us-east-1.amazonaws.com/orders/*.parquet"

  access_key    = var.s3_access_key
  access_secret = var.s3_access_secret
}

# CSV files with a header row that quote fields with doubled quotes.
resource "posthog_warehouse_table" "customers" {
  name        = "s3_customers"
  format      = "CSVWithNames"
  url_pattern = "https://my-bucket.s3.us-east-1.amazonaws.com/customers/*.csv"

  access_key    = var.s3_access_key
  access_secret = var.s3_access_secret

  csv_allow_double_quotes = true
}
