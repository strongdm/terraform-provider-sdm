# Terraform Provider for strongDM

The strongDM provider is used to configure and manage your strongDM account. It provides resources and datasources that allow you to:

* Register gateways
* Enroll databases and servers
* Grant access
* Create and suspend users
* Manage roles

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) 0.12+

## Example Usage

```hcl
provider "sdm" {
   	api_access_key = "ACCESS_KEY"
	api_secret_key = "SECRET_KEY"
}

resource "sdm_node" "example_gateway" {
    gateway {
        name = "example gateway"
        listen_address = "localhost:21222"
        bind_address = "0.0.0.0:0"
    }
}

```

## Argument Reference
* `api_access_key` - (Required) Your strongDM API access key. It can also be sourced from the `SDM_API_ACCESS_KEY` environment variable.

* `api_secret_key` - (Required) Your strongDM API secret key. It can also be sourced from the `SDM_API_SECRET_KEY` environment variable.

## Environment Variables
You can provide your credentials via the SDM_API_ACCESS_KEY and SDM_API_SECRET_KEY environment variables.

```hcl
provider "sdm" {}
```

Usage:

```shell
$ export SDM_API_ACCESS_KEY="<access-key>"
$ export SDM_API_SECRET_KEY="<secret-key>"
$ terraform plan
```

### Running the tests

_**Important:** These are integration tests. They will create and destroy real resources in your account!_

First clone this repository.

In order to run the tests you must set these environment variables so that the provider can authenticate:
```shell
$ export SDM_API_ACCESS_KEY="<access-key>"
$ export SDM_API_SECRET_KEY="<secret-key>"
```

From the cloned repo, run: 
```shell
$ cd terraform-provider-sdm
$ TF_ACC=yes go test ./sdm -v -count=1 -mod=vendor
```

You must set `TF_ACC=yes` in order to run tests. IF `TF_ACC=yes` is not set all tests will pass without any action.