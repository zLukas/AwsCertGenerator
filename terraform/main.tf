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
    zip_file = var.zip_file
    lambda_iam_resources = [module.certTable.arn]
    env_vars = var.env_vars
    lambda_iam_actions = var.lambda_iam_actions

}