---
page_title: "SDM: sdm_account_grant"
description: |-
  Provides settings for AccountGrant.
layout: “sdm”
sidebar_current: “docs-sdm-resource-account-grant"
---
# Resource: sdm_account_grant

AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
## Example Usage

```hcl
resource "sdm_account_grant" "test_account_grant" {
    account_id = "a-00000054"
    resource_id = "rs-12355562"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.

## Argument Reference
The following arguments are supported by the AccountGrant resource:
* `account_id` - (Required) The id of the attached role of this AccountGrant.
* `resource_id` - (Required) The id of the composite role of this AccountGrant.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the AccountGrant resource:
* `id` - A unique identifier for the AccountGrant resource.
