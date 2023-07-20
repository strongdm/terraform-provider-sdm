---
page_title: "SDM: sdm_peering_group_node"
description: |-
  Provides settings for PeeringGroupNode.
layout: “sdm”
sidebar_current: “docs-sdm-resource-peering-group-node"
---
# Resource: sdm_peering_group_node

PeeringGroupNode represents the attachment between a PeeringGroup and a Node.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the PeeringGroupNode resource:
* `group_id` - (Optional) Peering Group ID to which the node will be attached to.
* `node_id` - (Optional) Node ID to be attached.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the PeeringGroupNode resource:
* `id` - A unique identifier for the PeeringGroupNode resource.
## Import
PeeringGroupNode can be imported using the id, e.g.,

```
$ terraform import sdm_peering_group_node.example gn-12345678
```
