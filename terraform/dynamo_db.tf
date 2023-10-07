resource "aws_dynamodb_table" "CertTable" {
  name           = "Certificates"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Name"
  attribute {
    name = "Name"
    type = "S"
  }
}
