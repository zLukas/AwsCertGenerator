# general settings

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

variable clients {
  type = list(string)
}

# DynamoDB module settings
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

