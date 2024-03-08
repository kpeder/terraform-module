## Terraform Module with Terratest
An example Terraform module that includes Go tests using the Terratest library

### Decision Records
This repository uses architecture decision records to record design decisions
about important elements of the solution.

The ADR index is available [here](./docs/decisions/index.md).

### Requirements
Tested on Go version 1.21 with external modules:
```
"flag"
"os"
"runtime"
"strings"
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

Terraform can be automatically installed via script:
```
user@host:~/Projects/terraform-module$ cd fixtures/scripts
user@host:~/Projects/terraform-module/fixtures/scripts$ chmod +x install_terraform.sh && sudo ./install_terraform.sh
Detected OS Identifier: linux
Reading ./../versions.yaml
--2023-10-13 14:43:37--  https://releases.hashicorp.com/terraform/1.5.7/terraform_1.5.7_linux_amd64.zip
Resolving releases.hashicorp.com (releases.hashicorp.com)... 108.139.47.112, 108.139.47.90, 108.139.47.78, ...
Connecting to releases.hashicorp.com (releases.hashicorp.com)|108.139.47.112|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 21019880 (20M) [application/zip]
Saving to: ‘terraform.zip’

terraform.zip                                     100%[===========================================================================================================>]  20.05M  13.9MB/s    in 1.4s

2023-10-13 14:43:38 (13.9 MB/s) - ‘terraform.zip’ saved [21019880/21019880]

Archive:  terraform.zip
  inflating: /usr/local/bin/terraform
user@host:~/Projects/terraform-module/fixtures/scripts$ terraform version
Terraform v1.5.7
on linux_amd64
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5.0 |
| <a name="requirement_random"></a> [random](#requirement\_random) | ~> 3.5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_random"></a> [random](#provider\_random) | ~> 3.5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [random_pet.main](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/pet) | resource |

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
