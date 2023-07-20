---
page_title: "SDM: sdm_peering_group"
description: |-
  Provides settings for PeeringGroup.
layout: “sdm”
sidebar_current: “docs-sdm-resource-peering-group"
---
# Resource: sdm_peering_group

PeeringGroups are the building blocks used for explicit network topology making.
 They may be linked to other peering groups. Sets of PeeringGroupResource and PeeringGroupNode can be attached to a peering group.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the PeeringGroup resource:
* `name` - (Optional) Unique human-readable name of the PeeringGroup.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the PeeringGroup resource:
* `id` - A unique identifier for the PeeringGroup resource.
## Import
PeeringGroup can be imported using the id, e.g.,

```
$ terraform import sdm_peering_group.example g-12345678
```
