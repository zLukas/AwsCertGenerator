output "dynamodbTable"{
    value = module.certTable.arn
}

output "CertLambda_url"{
    value = aws_lambda_function_url.certLambda.function_url
}

output "activeClients" {
  value = [ for key  in aws_iam_access_key.clents_acces_keys: format("Client Name: %s, access key ID: %s, secret key: %s", key.user, key.id, key.secret )]
  sensitive = true
}