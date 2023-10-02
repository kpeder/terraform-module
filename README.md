## Terraform Module with Terratest
An example Terraform module that includes Go tests using the Terratest library

### Requirements
Tested on Go version 1.21 with external modules:
```
"flag"
"os"
"testing"

"gopkg.in/yaml.v3"

"github.com/gruntwork-io/terratest/modules/terraform"
"github.com/stretchr/testify/assert"
```

Uses installed packages:
```
golangci-lint
make
pre-commit
terraform
terraform-docs
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_random"></a> [random](#requirement\_random) | 3.5.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_random"></a> [random](#provider\_random) | 3.5.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [random_pet.main](https://registry.terraform.io/providers/hashicorp/random/3.5.1/docs/resources/pet) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_length"></a> [length](#input\_length) | Length of the random name, in words | `number` | `2` | no |
| <a name="input_prefix"></a> [prefix](#input\_prefix) | First word in the random name, if not random | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_random_pet"></a> [random\_pet](#output\_random\_pet) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
