---
page_title: "SDM: sdm_peering_group_peer"
description: |-
  Provides settings for PeeringGroupPeer.
layout: “sdm”
sidebar_current: “docs-sdm-resource-peering-group-peer"
---
# Resource: sdm_peering_group_peer

PeeringGroupPeer represents the link between two PeeringGroups
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the PeeringGroupPeer resource:
* `group_id` - (Optional) Group ID from which the link will originate.
* `peers_with_group_id` - (Optional) Peering Group ID to which Group ID will link.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the PeeringGroupPeer resource:
* `id` - A unique identifier for the PeeringGroupPeer resource.
## Import
PeeringGroupPeer can be imported using the id, e.g.,

```
$ terraform import sdm_peering_group_peer.example gp-12345678
```
