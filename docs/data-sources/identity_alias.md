---
page_title: "SDM: sdm_identity_alias"
description: |-
  Query for existing IdentityAliases instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-identity-alias"
---
# Data Source: sdm_identity_alias

IdentityAliases define the username to be used for a specific account
 when connecting to a remote resource using that identity set.
## Example Usage

```hcl
data "sdm_identity_alias" "user" {
    id = "i-0900909"
    username = "user"
}
```
## Argument Reference
The following arguments are supported by a IdentityAliases data source:
* `account_id` - (Optional) The account for this identity alias.
* `id` - (Optional) Unique identifier of the IdentityAlias.
* `identity_set_id` - (Optional) The identity set.
* `username` - (Optional) The username to be used as the identity alias for this account.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a IdentityAliases data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_identity_alias ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `identity_aliases` - A list where each element has the following attributes:
	* `account_id` - The account for this identity alias.
	* `id` - Unique identifier of the IdentityAlias.
	* `identity_set_id` - The identity set.
	* `username` - The username to be used as the identity alias for this account.
