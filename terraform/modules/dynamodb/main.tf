resource "aws_dynamodb_table" "table" {
  name           = var.table_name
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = var.table_main_key.name
  attribute {
    name = var.table_main_key.name
    type = var.table_main_key.type
  }
}