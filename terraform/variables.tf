variable access_key {
  type        = string
  sensitive   = true
}

variable secret_key {
  type        = string
  sensitive = true
}

variable region {
  type        = string
  default = "eu-central-1"
}

variable table_name {
    type = string 
    default = "certificates"
}

variable table_main_key {
    type = object({
        name = string
        type = string
    })
    default = {
        name = "Name"
        type = "S"
}
}

variable lambda_name {
    type      = string 
    default = "CertLambda"
}

variable env_vars {
    type = map(string)
    default = {
		    ENVIROMENT = "LAMBDA"
		    TABLE_NAME = "certificates"
        DB_REGION = "eu-central-1"
    }
}

variable lambda_iam_actions {
    type = list(string)
    default = ["dynamodb:TagResource",
				        "dynamodb:PutItem",
				        "dynamodb:DescribeTable",
				        "dynamodb:DeleteItem",
				        "dynamodb:UpdateItem"]
}

variable zip_file {
  type = string
  default = "bootstrap.zip"
}
