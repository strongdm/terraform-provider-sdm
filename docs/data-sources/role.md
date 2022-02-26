---
page_title: "SDM: sdm_role"
description: |-
  Query for existing Roles instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-role"
---
# Data Source: sdm_role

A Role is a collection of access grants, and typically corresponds to a team, Active Directory OU, or other organizational unit. Users are granted access to resources by assigning them to roles.
## Example Usage

```hcl
data "sdm_role" "composite_role_query" {
    composite = true
    tags = {
        region = "us-west"
        env = "dev"
    }    
}
```
## Argument Reference
The following arguments are supported by a Roles data source:
* `composite` - (Optional) Composite is true if the Role is a composite role.

 Deprecated: composite roles are deprecated, use multi-role instead.
* `id` - (Optional) Unique identifier of the Role.
* `name` - (Optional) Unique human-readable name of the Role.
* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Roles data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_role ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `roles` - A list where each element has the following attributes:
	* `composite` - Composite is true if the Role is a composite role.

 Deprecated: composite roles are deprecated, use multi-role instead.
	* `id` - Unique identifier of the Role.
	* `name` - Unique human-readable name of the Role.
	* `tags` - Tags is a map of key, value pairs.
