---
page_title: "SDM: sdm_peering_group_node"
description: |-
  Query for existing PeeringGroupNodes instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-peering-group-node"
---
# Data Source: sdm_peering_group_node

PeeringGroupNode represents the attachment between a PeeringGroup and a Node.
## Argument Reference
The following arguments are supported by a PeeringGroupNodes data source:
* `group_id` - (Optional) Peering Group ID to which the node will be attached to.
* `id` - (Optional) Unique identifier of the Attachment.
* `node_id` - (Optional) Node ID to be attached.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a PeeringGroupNodes data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_peering_group_node ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `peering_group_nodes` - A list where each element has the following attributes:
	* `group_id` - Peering Group ID to which the node will be attached to.
	* `id` - Unique identifier of the Attachment.
	* `node_id` - Node ID to be attached.
