---
page_title: "SDM: sdm_remote_identity_group"
description: |-
  Provides settings for RemoteIdentityGroup.
layout: “sdm”
sidebar_current: “docs-sdm-resource-remote-identity-group"
---
# Resource: sdm_remote_identity_group

A RemoteIdentityGroup defines a group of remote identities.
## Example Usage

```hcl

```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.

## Argument Reference
The following arguments are supported by the RemoteIdentityGroup resource:
* `name` - (Required) Unique human-readable name of the RemoteIdentityGroup.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the RemoteIdentityGroup resource:
* `id` - A unique identifier for the RemoteIdentityGroup resource.
## Import
RemoteIdentityGroup can be imported using the id, e.g.,

```
$ terraform import sdm_remote_identity_group.example ig-12345678
```
