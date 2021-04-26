# Terraform Provider for strongDM

The strongDM provider is used to configure and manage your strongDM account. It provides resources and datasources that allow you to:

* Register gateways
* Enroll databases and servers
* Grant access
* Create and suspend users
* Manage roles

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) 0.12+

## Authentication

If you don't already have them you will need to generate a set of API keys, instructions are here: [API Credentials](https://www.strongdm.com/docs/admin-guide/api-credentials/)

### Environment Variables

You can provide your credentials via the SDM_API_ACCESS_KEY and SDM_API_SECRET_KEY environment variables.

```sh
$ export SDM_API_ACCESS_KEY="<access-key>"
$ export SDM_API_SECRET_KEY="<secret-key>"
$ terraform plan
```

Provider statement
```hcl
provider "sdm" {}
```

### Direct

If it is not possible to use environment variables add the api keys directly to the provider statement as so:

```hcl
provider "sdm" {
	api_access_key = "ACCESS_KEY"
	api_secret_key = "SECRET_KEY"
}
```

## Basic Example

Create a Gateway

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

* Documentation: [terraform strongDM provider](https://registry.terraform.io/providers/strongdm/sdm/latest/docs)
* Examples: [GitHub - strongdm/terraform-provider-sdm-examples](https://github.com/strongdm/terraform-provider-sdm-examples)
	1. [Managing Resources](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/1_managing_resources)
	2. [Managing Gateways](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/4_managing_gateways)
	3. [Managing Roles](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/3_managing_roles)
	4. [Managing Accounts](https://github.com/strongdm/terraform-provider-sdm-examples/tree/master/2_managing_accounts)

## Contributions

Currently, we are not accepting pull requests directly to this repository, but our users are some of the most resourceful and ambitious folks out there. So, if you have something to contribute, find a bug, or just want to give us some feedback, please email <support@strongdm.com>.


## Running integration tests

_**Important:** These are integration tests. They will create and destroy real resources in your account!_

First clone this repository.

In order to run the tests you must set these environment variables so that the provider can authenticate:
```sh
$ export SDM_API_ACCESS_KEY="<access-key>"
$ export SDM_API_SECRET_KEY="<secret-key>"
```

From the cloned repo, run: 
```sh
$ cd terraform-provider-sdm
$ TF_ACC=yes go test ./sdm -v -count=1 -mod=vendor
```

You must set `TF_ACC=yes` in order to run tests. IF `TF_ACC=yes` is not set all tests will pass without any action.

## License

[Apache 2](https://github.com/strongdm/strongdm-sdk-python/blob/master/LICENSE)
