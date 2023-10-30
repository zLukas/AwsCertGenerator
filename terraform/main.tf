module certTable {
    source ="./modules/dynamodb"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    table_name = var.table_name
    table_main_key = var.table_main_key

}

module certLambda {
    source = "./modules/lambda"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    lambda_name = var.lambda_name
    handler = var.lambda_handler
    zip_file = var.zip_file
    lambda_iam_resources = [module.certTable.arn]
    env_vars = var.env_vars
    lambda_iam_actions = var.lambda_iam_actions

}

resource "aws_iam_user" "admin_user" {
  name = var.admin_user
  path = "/certadmin/"
}

resource "aws_iam_access_key" "admin_key" {
  user = aws_iam_user.admin_user.name
}

module "admin_iam-group-with-policies" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-group-with-policies"
  version = "5.30.0"

  name = "certAdmin"
  group_users = [var.admin_user]
  path = "/CertAdmin/"
  custom_group_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaRole",
    "arn:aws:iam::aws:policy/IAMFullAccess"
    ]
}

module "client_iam-group-with-policies" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-group-with-policies"
  version = "5.30.0"
  name = "certClient"
  path = "/CertClient/"
    custom_group_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonDynamoDBReadOnlyAccess"
    ]
}




