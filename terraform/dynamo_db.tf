resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "Certificates"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Name"
  attribute {
    name = "Name"
    type = "S"
  }
}
