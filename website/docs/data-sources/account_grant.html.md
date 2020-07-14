---
page_title: "SDM: sdm_account_grant"
description: |-
  Query for existing AccountGrants instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-account-grant"
---
# Data Source: sdm_account_grant

AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
## Example Usage

```hcl
data "sdm_account_grant" "account_grant_query" {
    account_id = "a-00000054"
}
```
## Argument Reference
The following arguments are supported by a AccountGrants data source:
* `id` - (Optional) Unique identifier of the AccountGrant.
* `resource_id` - (Optional) The id of the composite role of this AccountGrant.
* `account_id` - (Optional) The id of the attached role of this AccountGrant.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a AccountGrants data source:
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `account_grants` - A list where each element has the following attributes:
	* `id` - Unique identifier of the AccountGrant.
	* `resource_id` - The id of the composite role of this AccountGrant.
	* `account_id` - The id of the attached role of this AccountGrant.
