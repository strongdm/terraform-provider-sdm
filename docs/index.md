---
page_title: "strongDM Provider"
---

# strongDM Provider

Summary of what the provider is for, including use cases and links to
app/service documentation.

## Example Usage

```hcl
provider "sdm" {
   	api_access_key = "ACCESS_KEY"
	api_secret_key = "SECRET_KEY"
	host = "api.strongdm.com:443"
}

resource "sdm_node" "example_gateway" {
    gateway {
        name = "example gateway"
        listen_address = "localhost:21222"
        bind_address = "0.0.0.0:0"
    }
}

```



## Authentication
The strongDM provider allows for two methods of authentication:
- static credentials
- environment variables


## Argument Reference
The following arguments are supported by the strongDM provider:

* `api_access_key` - (Optional) This is the strongDM API access key. Alternatively, sourced via the environment variable `SDM_API_ACCESS_KEY`.

* `api_secret_key` - (Optional) This is the strongDM API secret key. Alternatively, sourced via the environment variable `SDM_API_SECRET_KEY`.

* `host` - (Optional) This is the host that strongDM API requests will be sent to. Alternatively, sourced via the environment variable `SDM_API_HOST`. 