---
page_title: "SDM: sdm_account"
description: |-
  Provides settings for Account.
layout: “sdm”
sidebar_current: “docs-sdm-resource-account"
---
# Resource: sdm_account

Accounts are users that have access to strongDM. There are two types of accounts:
 1. **Users:** humans who are authenticated through username and password or SSO.
 2. **Service Accounts:** machines that are authenticated using a service token.
## Example Usage

```hcl
resource "sdm_account" "test-user" {
    user {
        first_name = "al"
        last_name = "bob"
        email = "albob@strongdm.com"
    }
}

resource "sdm_account" "test-service" {
    service {
        name = "test-service"
    }
}
```
## Argument Reference
The following arguments are supported by the Account resource:
* user:
	* `email` - (Required) The User's email address. Must be unique.
	* `first_name` - (Required) The User's first name.
	* `last_name` - (Required) The User's last name.
	* `suspended` - (Optional) The User's suspended state.
* service:
	* `name` - (Required) Unique human-readable name of the Service.
	* `suspended` - (Optional) The Service's suspended state.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Account resource:
* `id` - A unique identifier for the Account resource.
