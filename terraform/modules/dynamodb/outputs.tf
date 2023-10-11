output "name" {
  value = aws_dynamodb_table.table.name
}

output "arn" {
  value = aws_dynamodb_table.table.arn
}

output "main_key" {
    value = aws_dynamodb_table.table.hash_key
}