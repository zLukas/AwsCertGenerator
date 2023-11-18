output "dynamodbTable"{
    value = module.certTable.arn
}

output "Lambda_url"{
    value = aws_lambda_function_url.certLambda_latest.function_url
}