variable access_key {
  type        = string
}

variable secret_key {
  type        = string
}

variable region {
  type        = string
}

variable table_name {
    type = string 
}

variable table_main_key {
    type = object({
        name = string
        type = string
    })
}