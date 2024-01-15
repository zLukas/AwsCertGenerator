resource "aws_cognito_user_pool" "pool" {
  name = var.pool_name
}

resource "aws_cognito_user_group" "main" {
  name         = "default-group"
  user_pool_id = aws_cognito_user_pool.pool.id
  description  = "default user group"
  role_arn     = aws_iam_role.cognito_role.arn
}