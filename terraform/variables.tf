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