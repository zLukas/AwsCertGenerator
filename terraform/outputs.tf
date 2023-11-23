output "dynamodbTable"{
    value = module.certTable.arn
}

output "CertLambda_url"{
    value = aws_lambda_function_url.certLambda.function_url
}

output "UserLambda_url"{
    value = aws_lambda_function_url.userLambda.function_url
}