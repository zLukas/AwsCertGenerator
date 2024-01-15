output "dynamodbTable"{
    value = module.certTable.arn
}

output "CertLambda_url"{
    value = aws_lambda_function_url.certLambda.function_url
}

output "cognito_pool_id" {
    value = module.cognito.cognito_pool_id
}

output "estimated_number_of_users" {
    value = module.cognito.estimated_number_of_users
}