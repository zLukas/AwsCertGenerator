variable access_key {
  type        = string
}

variable secret_key {
  type        = string
}

variable region {
  type        = string
}

 variable policy {
  type = object({
    actions = list(string)
    resource = list(string)
  })
 }

variable policy_name {
    type       =  string
}

variable pool {
    type        = string
}
