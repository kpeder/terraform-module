terraform {
  required_version = "~> 1.5"

  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "~> 3.5"
    }
  }
}

variable "length" {
  description = "Length of the random name, in words"
  type        = number

  default = 2
}

variable "prefix" {
  description = "First word in the random name, if not random"
  type        = string

  default = null
}

locals {
  separator = "-"
}

resource "random_pet" "main" {
  length    = var.length
  prefix    = var.prefix
  separator = local.separator
}

output "random_pet" {
  value = random_pet.main.id
}
