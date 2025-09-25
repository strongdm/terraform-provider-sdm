---
page_title: "SDM: sdm_account_group"
description: |-
  Provides settings for AccountGroup.
layout: “sdm”
sidebar_current: “docs-sdm-resource-account-group"
---
# Resource: sdm_account_group

An AccountGroup is a link between an Account and a Group.
## Example Usage

```hcl
resource "sdm_account_group" "security_lead_to_security_team" {
  account_id = sdm_account.security_lead.id
  group_id   = sdm_group.security_team.id
}

resource "sdm_account_group" "admin_user_to_administrators" {
  account_id = "a-1234567890abcdef"
  group_id   = "group-1234567890abcdef"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the AccountGroup resource:
* `account_id` - (Required) Unique identifier of the Account.
* `group_id` - (Required) Unique identifier of the Group.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the AccountGroup resource:
* `id` - A unique identifier for the AccountGroup resource.
## Import
A AccountGroup can be imported using the id, e.g.,

```
$ terraform import sdm_account_group.example accountgroup-12345678
```
