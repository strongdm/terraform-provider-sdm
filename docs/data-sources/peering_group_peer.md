---
page_title: "SDM: sdm_peering_group_peer"
description: |-
  Query for existing PeeringGroupPeers instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-peering-group-peer"
---
# Data Source: sdm_peering_group_peer

PeeringGroupPeer represents the link between two PeeringGroups
## Argument Reference
The following arguments are supported by a PeeringGroupPeers data source:
* `group_id` - (Optional) Group ID from which the link will originate.
* `id` - (Optional) Unique identifier of the Attachment.
* `peers_with_group_id` - (Optional) Peering Group ID to which Group ID will link.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a PeeringGroupPeers data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_peering_group_peer ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `peering_group_peers` - A list where each element has the following attributes:
	* `group_id` - Group ID from which the link will originate.
	* `id` - Unique identifier of the Attachment.
	* `peers_with_group_id` - Peering Group ID to which Group ID will link.
