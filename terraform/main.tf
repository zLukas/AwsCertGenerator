data "aws_caller_identity" "current" {}
locals {
    account_id = data.aws_caller_identity.current.account_id
	lambda_name = "CertGen"
	table_name = "certificates"
}

module certTable {
    source ="./modules/dynamodb"
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
    source = "./modules/lambda"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    lambda_name = local.lambda_name
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
resource "aws_lambda_function_url" "certLambda_latest" {
  function_name      = local.lambda_name
  authorization_type = "AWS_IAM"
}