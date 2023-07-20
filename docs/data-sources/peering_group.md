---
page_title: "SDM: sdm_peering_group"
description: |-
  Query for existing PeeringGroups instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-peering-group"
---
# Data Source: sdm_peering_group

PeeringGroups are the building blocks used for explicit network topology making.
 They may be linked to other peering groups. Sets of PeeringGroupResource and PeeringGroupNode can be attached to a peering group.
## Argument Reference
The following arguments are supported by a PeeringGroups data source:
* `id` - (Optional) Unique identifier of the PeeringGroup.
* `name` - (Optional) Unique human-readable name of the PeeringGroup.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a PeeringGroups data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_peering_group ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `peering_groups` - A list where each element has the following attributes:
	* `id` - Unique identifier of the PeeringGroup.
	* `name` - Unique human-readable name of the PeeringGroup.
