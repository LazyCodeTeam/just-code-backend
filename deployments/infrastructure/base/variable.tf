variable "env" {
  type = string

  validation {
    condition     = contains(["dev", "prod"], var.env)
    error_message = "env must be one of [dev, prod]"
  }
}

variable "app_name" {
  type = string
}

variable "region" {
  type = string
}

variable "image_tag" {
  type = string
}
