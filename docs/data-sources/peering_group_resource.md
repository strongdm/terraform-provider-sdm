---
page_title: "SDM: sdm_peering_group_resource"
description: |-
  Query for existing PeeringGroupResources instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-peering-group-resource"
---
# Data Source: sdm_peering_group_resource

PeeringGroupResource represents the attachment between a PeeringGroup and a Resource.
## Argument Reference
The following arguments are supported by a PeeringGroupResources data source:
* `group_id` - (Optional) Peering Group ID to which the resource will be attached to.
* `id` - (Optional) Unique identifier of the Attachment.
* `resource_id` - (Optional) Resource ID to be attached.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a PeeringGroupResources data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_peering_group_resource ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `peering_group_resources` - A list where each element has the following attributes:
	* `group_id` - Peering Group ID to which the resource will be attached to.
	* `id` - Unique identifier of the Attachment.
	* `resource_id` - Resource ID to be attached.
