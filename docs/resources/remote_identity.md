---
page_title: "SDM: sdm_remote_identity"
description: |-
  Provides settings for RemoteIdentity.
layout: “sdm”
sidebar_current: “docs-sdm-resource-remote-identity"
---
# Resource: sdm_remote_identity

RemoteIdentities define the username to be used for a specific account
 when connecting to a remote resource using that group.
## Example Usage

```hcl
resource "sdm_remote_identity" "user" {
    id = "i-0900909"
    username = "user"
}

```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the RemoteIdentity resource:
* `account_id` - (Required) The account for this remote identity.
* `remote_identity_group_id` - (Required) The remote identity group.
* `username` - (Required) The username to be used as the remote identity for this account.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the RemoteIdentity resource:
* `id` - A unique identifier for the RemoteIdentity resource.
## Import
A RemoteIdentity can be imported using the id, e.g.,

```
$ terraform import sdm_remote_identity.example i-12345678
```
