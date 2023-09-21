# Terraform Provider for StrongDM

You can use the StrongDM Terraform provider to configure and manage your StrongDM account and resources. This project allows you to leverage Terraform to complete the following tasks in StrongDM:

* Create and register gateways or relays
* Enroll databases, servers, clusters, and websites
* Create, manage, and suspend users
* Grant user and resource access
* Manage roles
* Manage access workflows

For more detailed examples using the StrongDM provider and Amazon Web Services (AWS) or Microsoft Azure, check the following resources:

* [Quick Start StrongDM with Terraform and AWS](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-aws/)
* [Quick Start StrongDM with Terraform and Azure](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-azure/)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13+

## Versioning

StrongDM uses [semantic versioning](https://semver.org/). We do not guarantee
compatibility between major versions. Be sure to use version constraints to pin
your dependency to the desired major version of the StrongDM Terraform provider.

## Authentication

If you do not already have them, you must [generate a set of API keys](https://www.strongdm.com/docs/admin-guide/api-credentials/).

### Environment Variables

You can provide your credentials via the SDM_API_ACCESS_KEY and
SDM_API_SECRET_KEY environment variables.

```sh
$ export SDM_API_ACCESS_KEY="<ACCESS_KEY>"
$ export SDM_API_SECRET_KEY="<SECRET_KEY>"
$ terraform plan
```

Provider statement

```hcl
provider "sdm" {}
```

### Direct

If it is not possible to use environment variables, add the API keys directly to
the provider statement as follows:

```hcl
provider "sdm" {
    api_access_key = "<ACCESS_KEY>"
    api_secret_key = "<SECRET_KEY>"
}
```

## Basic Example

Use the following basic example to [create a gateway](https://www.strongdm.com/docs/admin-ui-guide/network/gateways/) in StrongDM with the Terraform provider.

```hcl
resource "sdm_node" "example_gateway" {
    gateway {
        name = "example gateway"
        listen_address = "localhost:5000"
        bind_address = "0.0.0.0:5000"
    }
}
```

```shell
$ terraform plan
$ terraform apply
```

## Useful Links

- [StrongDM Terraform provider](https://registry.terraform.io/providers/strongdm/sdm/latest/docs) documentation
- [Migrating from v2 to v3](https://github.com/strongdm/terraform-provider-sdm/releases/tag/v3.0.0)
- [Migrating from Role Grants to Access Rules](https://registry.terraform.io/providers/strongdm/sdm/2.0.0/docs/guides/migrating_from_role_grants_to_access_rules)
- Examples: [GitHub - strongdm/terraform-provider-sdm-examples](https://github.com/strongdm/terraform-provider-sdm-examples)
  - [Managing Resources](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/1_managing_resources)
  - [Managing Gateways](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/4_managing_gateways)
  - [Managing Roles](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/3_managing_roles)
  - [Managing Accounts](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/2_managing_accounts)
  - [Managing Access Workflows](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/5_managing_workflows)
  - [Quick Start StrongDM with Terraform and AWS](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-aws/)
  - [Quick Start StrongDM with Terraform and Azure](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-azure/)

## Contributions

Currently, we are not accepting pull requests directly to this repository, but
our users are some of the most resourceful and ambitious folks out there. If
you have something to contribute, find a bug, or just want to give us some
feedback, please email <support@strongdm.com>.

## Running Integration Tests

_**Important:** These are integration tests. They will create and destroy real
resources in your account!_

1. First, clone this repository.

2. In order to run the tests you must set these environment variables so that the
provider can authenticate:

```sh
$ export SDM_API_ACCESS_KEY="<ACCESS_KEY>"
$ export SDM_API_SECRET_KEY="<SECRET_KEY>"
```

3. From the cloned repo, run:

```sh
$ cd terraform-provider-sdm
$ TF_ACC=yes go test ./sdm -v -count=1 -mod=vendor
```

You must set `TF_ACC=yes` in order to run tests. IF `TF_ACC=yes` is not set, all
tests will pass without any action.

## License

[Mozilla Public License 2.0](https://github.com/strongdm/terraform-provider-sdm/blob/master/LICENSE)
