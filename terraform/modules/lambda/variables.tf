variable access_key {
  type        = string
}

variable secret_key {
  type        = string
}

variable region {
  type        = string
}

variable timeout {
  type = number
  default = 30
}


variable handler {
  type = string
  default = 30
}

variable runtime {
  type = string
  default = 30
}

variable memory_size {
  type = number
  default = 512
}

variable lambda_name {
    type      = string 
}

variable env_vars {
    type = map(string)
    default = {
      ENVIROMENT = "LAMBDA"
    }
}

variable lambda_iam_actions {
    type = list(string)
}

variable lambda_iam_resources {
    type = list(string)
}

variable zip_file {
  type = string 
}