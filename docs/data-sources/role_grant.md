---
page_title: "SDM: sdm_role_grant"
description: |-
  Query for existing RoleGrants instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-role-grant"
---
# Data Source: sdm_role_grant

A RoleGrant connects a resource to a role, granting members of the role access to that resource.

## Example Usage

```hcl
data "sdm_role_grant" "role_grant_query" {
    role_id = "r-0009994"
}
```

## Argument Reference
The following arguments are supported by a RoleGrants data source:
* `id` - (Optional) Unique identifier of the RoleGrant.
* `resource_id` - (Optional) The id of the resource of this RoleGrant.
* `role_id` - (Optional) The id of the attached role of this RoleGrant.

## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a RoleGrants data source:
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `role_grants` - A list where each element has the following attributes:
	* `id` - Unique identifier of the RoleGrant.
	* `resource_id` - The id of the resource of this RoleGrant.
	* `role_id` - The id of the attached role of this RoleGrant.
