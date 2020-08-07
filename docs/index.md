---
page_title: "strongDM Provider"
---

# strongDM Provider

The strongDM provider is used to configure and manage your strongDM account and resources. It allows you to:

* Register gateways
* Enroll databases and servers
* Grant access
* Create and suspend users
* Manage roles

## Configuration

This provider needs to be configured with API credentials from your strongDM account. These credentials can be supplied to the provider directly or set using environment variables. We recommend useing environmental variables as best practice.

## Environment Variables
You can provide your credentials via the SDM_API_ACCESS_KEY and SDM_API_SECRET_KEY environment variables.

### Example Usage

```hcl
provider "sdm" {}
```

Terminal:

```shell
$ export SDM_API_ACCESS_KEY="<access-key>"
$ export SDM_API_SECRET_KEY="<secret-key>"
$ terraform plan
```

## Direct Credentials
You can also provide your credentials directly to the provider.

### Example Usage

```hcl
provider "sdm" {
   	api_access_key = "ACCESS_KEY"
	api_secret_key = "SECRET_KEY"
}

resource "sdm_node" "gateway" {
    gateway {
        name = "test-gateway"
        listen_address = "165.23.40.1:21222"
    }
}
```

### Argument Reference
* `api_access_key` - (Required) Your strongDM API access key. It can also be sourced from the `SDM_API_ACCESS_KEY` environment variable.

* `api_secret_key` - (Required) Your strongDM API secret key. It can also be sourced from the `SDM_API_SECRET_KEY` environment variable.

## Documentation

* [Guides](./guides/)
* [Data Sources](./data-sources/)
* [Resources](./resources/)

