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
  inputs = yamldecode(file("./inputs.yaml"))
}

module "random_pet" {
  source = "../../."

  length = coalesce(local.inputs.length, var.length)
  prefix = try(coalesce(local.inputs.prefix, var.prefix), null)
}

output "random_pet" {
  value = module.random_pet.random_pet
}
