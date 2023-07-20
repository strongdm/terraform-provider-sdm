---
page_title: "SDM: sdm_peering_group_resource"
description: |-
  Provides settings for PeeringGroupResource.
layout: “sdm”
sidebar_current: “docs-sdm-resource-peering-group-resource"
---
# Resource: sdm_peering_group_resource

PeeringGroupResource represents the attachment between a PeeringGroup and a Resource.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the PeeringGroupResource resource:
* `group_id` - (Optional) Peering Group ID to which the resource will be attached to.
* `resource_id` - (Optional) Resource ID to be attached.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the PeeringGroupResource resource:
* `id` - A unique identifier for the PeeringGroupResource resource.
## Import
PeeringGroupResource can be imported using the id, e.g.,

```
$ terraform import sdm_peering_group_resource.example gr-12345678
```
