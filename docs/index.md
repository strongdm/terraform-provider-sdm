---
page_title: "strongDM Provider"
---

# strongDM Provider

The strongDM provider is used to configure and manage your strongDM account. It provides resources and datasources that allow you to:

* Register gateways
* Enroll databases and servers
* Grant access
* Create and suspend users
* Manage roles

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

## Documentation

* [Guides](./guides/)
* [Data Sources](./data-sources/)
* [Resources](./resources/)

