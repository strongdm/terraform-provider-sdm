---
page_title: "SDM: sdm_group_role"
description: |-
  Provides settings for GroupRole.
layout: “sdm”
sidebar_current: “docs-sdm-resource-group-role"
---
# Resource: sdm_group_role

A GroupRole assigns a Group to a Role.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the GroupRole resource:
* `group_id` - (Required) The assigned Group ID.
* `role_id` - (Required) The assigned Role ID.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the GroupRole resource:
* `id` - A unique identifier for the GroupRole resource.
## Import
A GroupRole can be imported using the id, e.g.,

```
$ terraform import sdm_group_role.example grouprole-12345678
```
