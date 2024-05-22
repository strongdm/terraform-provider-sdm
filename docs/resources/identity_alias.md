---
page_title: "SDM: sdm_identity_alias"
description: |-
  Provides settings for IdentityAlias.
layout: “sdm”
sidebar_current: “docs-sdm-resource-identity-alias"
---
# Resource: sdm_identity_alias

IdentityAliases define the username to be used for a specific account
 when connecting to a remote resource using that identity set.
## Example Usage

```hcl
resource "sdm_identity_alias" "user" {
    id = "i-0900909"
    username = "user"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the IdentityAlias resource:
* `account_id` - (Required) The account for this identity alias.
* `identity_set_id` - (Required) The identity set.
* `username` - (Required) The username to be used as the identity alias for this account.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the IdentityAlias resource:
* `id` - A unique identifier for the IdentityAlias resource.
## Import
A IdentityAlias can be imported using the id, e.g.,

```
$ terraform import sdm_identity_alias.example i-12345678
```
