output "cognito_pool_id" {
    value = aws_cognito_user_pool.pool.id
}

output "estimated_number_of_users" {
    value = aws_cognito_user_pool.pool.estimated_number_of_users
}