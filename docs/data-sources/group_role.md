---
page_title: "SDM: sdm_group_role"
description: |-
  Query for existing GroupsRoles instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-group-role"
---
# Data Source: sdm_group_role

A GroupRole assigns a Group to a Role.
## Argument Reference
The following arguments are supported by a GroupsRoles data source:
* `group_id` - (Optional) The assigned Group ID.
* `id` - (Optional) Unique identifier of the GroupRole.
* `role_id` - (Optional) The assigned Role ID.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a GroupsRoles data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_group_role ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `groups_roles` - A list where each element has the following attributes:
	* `group_id` - The assigned Group ID.
	* `id` - Unique identifier of the GroupRole.
	* `role_id` - The assigned Role ID.
