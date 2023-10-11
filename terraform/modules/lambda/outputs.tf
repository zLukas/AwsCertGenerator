output lambda_invoke_arn {
    value = aws_lambda_function.lambda.invoke_arn
}

output lambda_arn {
    value = aws_lambda_function.lambda.arn
}