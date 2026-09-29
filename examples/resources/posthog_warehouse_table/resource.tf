variable "warehouse_bucket_access_key" {
  type      = string
  sensitive = true
}

variable "warehouse_bucket_access_secret" {
  type      = string
  sensitive = true
}

# Query every Parquet file under orders/ as one table named `orders`.
# The key needs s3:ListBucket on the bucket and s3:GetObject on the files.
# Changing `format` or `url_pattern` replaces the table; renaming it or
# rotating the key updates it in place.
resource "posthog_warehouse_table" "orders" {
  name          = "orders"
  format        = "Parquet"
  url_pattern   = "https://your-bucket.s3.us-east-1.amazonaws.com/orders/*.parquet"
  access_key    = var.warehouse_bucket_access_key
  access_secret = var.warehouse_bucket_access_secret
}
