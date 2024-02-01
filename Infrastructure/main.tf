data "aws_caller_identity" "current" {}

locals {
    account_id = data.aws_caller_identity.current.account_id
	  cert_lambda_name = "CertGen"
    user_lambda_name = "Users"
    client_lambda_name= "Clients"
	  table_name = "certificates"
}

module certTable {
    source ="./IaaS/aws/dynamodb"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    table_name = local.table_name
    table_main_key = {
                name = "Name"
                type = "S"
    }

}

module certLambda {
    source = "./IaaS/aws/lambda"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    lambda_name = local.cert_lambda_name
    zip_file = "certgen.zip"
    handler = "bootstrap"
    runtime = "go1.x"
    lambda_iam_resources = [module.certTable.arn]
    env_vars = {
		    TABLE_NAME = local.table_name
            DB_REGION = var.region
    }
    lambda_iam_actions = ["dynamodb:TagResource",
				          "dynamodb:PutItem",
				          "dynamodb:DescribeTable",
				          "dynamodb:DeleteItem",
				          "dynamodb:UpdateItem"]

}

module userLambda {
    source = "./IaaS/aws/lambda"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    lambda_name = local.user_lambda_name
    zip_file = "users.zip"
    handler = "lambda_handler"
    runtime = "python3.11"
    lambda_iam_resources = ["arn:aws:cognito::${local.account_id}:*"]
    lambda_iam_actions = ["cognito:addUser"]

}

module clientLambda {
    source = "./IaaS/aws/lambda"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    lambda_name = local.client_lambda_name
    zip_file = "client.zip"
    handler = "lambda_handler"
    runtime = "python3.11"
    lambda_iam_resources = [module.certTable.arn]
    lambda_iam_actions = ["dynamodb:GetItem"]

}



resource "aws_lambda_function_url" "certLambda" {
  function_name      = local.cert_lambda_name
  authorization_type = "AWS_IAM"
  depends_on = [
    module.certLambda
  ]
}

resource "aws_lambda_function_url" "userLambda" {
  function_name      = local.user_lambda_name
  authorization_type = "AWS_IAM"
  depends_on = [
    module.userLambda
  ]
}

resource "aws_lambda_function_url" "clientLambda" {
  function_name      = local.client_lambda_name
  authorization_type = "AWS_IAM"
  depends_on = [
    module.clientLambda
  ]
}

module "cognito"{
  source = "./IaaS/aws/cognito"
  pool_name = "CertUsers"
  access_key = var.access_key
  secret_key = var.secret_key
  region = var.region
  policy = {
        name = "CertUsers"
        actions = ["dynamodb:DescribeTable"]
        resources = [module.certTable.arn]
  }
}