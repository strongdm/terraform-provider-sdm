---
page_title: "StrongDM Provider"
---

# StrongDM Provider

You can use the StrongDM Terraform provider to configure and manage your StrongDM account and resources. This project allows you to leverage Terraform to complete the following tasks in StrongDM:

* Create and register gateways or relays
* Enroll databases, servers, clusters, and websites
* Create, manage, and suspend users
* Grant user and resource access
* Manage roles

For more detailed examples using the StrongDM provider and Amazon Web Services (AWS) or Microsoft Azure, check the following resources:

* [Quick Start StrongDM with Terraform and AWS](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-aws/)
* [Quick Start StrongDM with Terraform and Azure](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-azure/)

## Example Usage

To install the StrongDM provider, copy and paste this code into your Terraform configuration. The following example applies to Terraform 0.13 and later.

```hcl
# Install StrongDM provider
terraform {
  required_providers {
    sdm = {
      source = "strongdm/sdm"
      version = ">=3.3.0"
    }
  }
}

# Configure StrongDM provider
provider "sdm" {
  # Add configuration options / StrongDM API credentials
}
```

## Configuration Options

To configure the provider, you must first add the [API credentials](https://www.strongdm.com/docs/admin-ui-guide/access/api-keys/#api-credentials) you created in the Admin UI of your StrongDM account.

These credentials can be supplied to the provider directly or set using environment variables. As a best practice, we recommend using environment variables.

### Timeouts

The StrongDM provider supports the timeouts listed below, shown with their default values.

For datasources
```hcl
timeouts {
  default = "60s"
}
```

For resources:
```hcl
timeouts {
  default = "60s"
  create = "60s"
  update = "60s"
  read = "60s"
  delete = "60s"
}
```

## Environment Variables

You can provide your credentials via the SDM_API_ACCESS_KEY and SDM_API_SECRET_KEY environment variables.

### Example Usage

```hcl
provider "sdm" {}
```

Add the values for your API access and secret keys from the StrongDM Admin UI. Then store the API credentials as environmental variables in your terminal:

```shell
# Set environment variables
$ export SDM_API_ACCESS_KEY="<ACCESS_KEY>"
$ export SDM_API_SECRET_KEY="<SECRET_KEY>"
$ terraform plan
```

## Direct Credentials

If you prefer to use the API credentials directly, see the following example. Add the values for your API access and secret keys from the StrongDM Admin UI to the corresponding arguments.

### Example Usage

```hcl
# Configure StrongDM provider
provider "sdm" {
  # Add API access key and secret key from Admin UI    
  api_access_key = "<ACCESS_KEY>"
  api_secret_key = "<SECRET_KEY>"
}

# Create StrongDM gateway
resource "sdm_node" "example_gateway" {
  gateway {
    name = "example-gateway"
    listen_address = "gateway.example.com:5555"
  }
}
```

## Argument Reference

* `api_access_key` - (Required) Your StrongDM API access key. It can also be sourced from the SDM_API_ACCESS_KEY environment variable.
* `api_secret_key` - (Required) Your StrongDM API secret key. It can also be sourced from the SDM_API_SECRET_KEY environment variable.
* `host` - Your StrongDM API host. It can also be sourced from the SDM_API_HOST environment variable.
