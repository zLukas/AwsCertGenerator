data "aws_caller_identity" "current" {}
locals {
    account_id = data.aws_caller_identity.current.account_id
}

module certTable {
    source ="./modules/dynamodb"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    table_name = "certificates"
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
    lambda_name = "CertGen"
    zip_file = "certgen.zip"
    handler = "bootstrap"
    runtime = "go1.x"
    lambda_iam_resources = [module.certTable.arn]
    env_vars = {
		    TABLE_NAME = "certificates"
            DB_REGION = "eu-central-1"
    }
    lambda_iam_actions = ["dynamodb:TagResource",
				          "dynamodb:PutItem",
				          "dynamodb:DescribeTable",
				          "dynamodb:DeleteItem",
				          "dynamodb:UpdateItem"]

}

module userLambda {
    source = "./modules/lambda"
    access_key = var.access_key
    secret_key = var.secret_key
    region = var.region
    lambda_name = "usersLambda"
    zip_file = "users.zip"
    handler = "handler"
    runtime = "python3.9"
    lambda_iam_actions = [
				"iam:GetPolicyVersion",
				"iam:DeleteAccessKey",
				"iam:UpdateGroup",
				"iam:DeletePolicy",
				"iam:CreateRole",
				"iam:PutRolePolicy",
				"iam:CreateUser",
				"iam:CreateAccessKey",
				"iam:AddUserToGroup",
				"iam:DetachRolePolicy",
				"iam:DetachGroupPolicy",
				"iam:ListRolePolicies",
				"iam:ListAccessKeys",
				"iam:GetRole",
				"iam:CreateGroup",
				"iam:GetPolicy",
				"iam:UpdateUser",
				"iam:PutUserPermissionsBoundary",
				"iam:DeleteUserPolicy",
				"iam:AttachUserPolicy",
				"iam:UpdateAccessKey",
				"iam:DeleteUser",
				"iam:DeleteUserPermissionsBoundary",
				"iam:ListUserPolicies",
				"iam:TagUser",
				"iam:CreatePolicy",
				"iam:UntagUser",
				"iam:GetUserPolicy",
				"iam:ListGroupsForUser",
				"iam:UntagPolicy",
				"iam:UpdateRole",
				"iam:GetGroupPolicy",
				"iam:GetUser",
				"iam:DeletePolicyVersion",
				"iam:ListUserTags"
			]
    lambda_iam_resources = [
				"arn:aws:iam::${locals.account_id}:role/*",
				"arn:aws:iam::${locals.account_id}:group/*",
				"arn:aws:iam::${locals.account_id}:policy/*",
				"arn:aws:iam::${locals.account_id}:user/*"
			]
}