output "dynamodbTable"{
    value = module.certTable.arn
}

output "Lambda"{
    value = module.certLambda.arn
}

output "adminUsers" {
    value = aws_iam_user.admin_user.name
}

output "adminUsers_access_key" {
    value = aws_iam_access_key.admin_key.id
    sensitive = true
}

output "adminUsers_secret_key" {
    value = aws_iam_access_key.admin_key.secret
    sensitive = true
}