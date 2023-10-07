variable access_key {
  type        = string
}

variable secret_key {
  type        = string
}

variable region {
  type        = string
}

variable lambda_name {
    type      = string 
}

variable env_vars {
    type = list(object({
        name = string 
        value = string
    }))
}

variable lambda_iam_actions {
    type = list(string)
}

variable lambda_iam_resources {
    type = list(string)
}