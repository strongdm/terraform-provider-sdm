---
page_title: "SDM: sdm_role_grant"
description: |-
  Provides settings for RoleGrant.
layout: “sdm”
sidebar_current: “docs-sdm-resource-role-grant"
---
# Resource: sdm_role_grant

A RoleGrant connects a resource to a role, granting members of the role access to that resource.

## Example Usage

```hcl
resource "sdm_role_grant" "test_role_grant" {
    role_id = "r-243561"
    resource_id = "rs-21522"
}
```

## Argument Reference
The following arguments are supported by the RoleGrant resource:
* `resource_id` - (Required) The id of the resource of this RoleGrant.
* `role_id` - (Required) The id of the attached role of this RoleGrant.

## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the RoleGrant resource:
* `id` - A unique identifier for the RoleGrant resource.
