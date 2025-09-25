---
page_title: "SDM: sdm_account_group"
description: |-
  Query for existing AccountsGroups instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-account-group"
---
# Data Source: sdm_account_group

An AccountGroup is a link between an Account and a Group.
## Example Usage

```hcl
data "sdm_account_group" "security_lead_groups" {
  account_id = sdm_account.security_lead.id
}

data "sdm_account_group" "specific_relationship" {
  id = "accountgroup-1234567890abcdef"
}

data "sdm_account_group" "users_in_group" {
  group_id = sdm_group.devops_group.id
}
```
## Argument Reference
The following arguments are supported by a AccountsGroups data source:
* `account_id` - (Optional) Unique identifier of the Account.
* `group_id` - (Optional) Unique identifier of the Group.
* `id` - (Optional) Unique identifier of the AccountGroup.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a AccountsGroups data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_account_group ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `accounts_groups` - A list where each element has the following attributes:
	* `account_id` - Unique identifier of the Account.
	* `group_id` - Unique identifier of the Group.
	* `id` - Unique identifier of the AccountGroup.
