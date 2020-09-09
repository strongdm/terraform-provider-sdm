---
page_title: "SDM: sdm_role"
description: |-
  Provides settings for Role.
layout: “sdm”
sidebar_current: “docs-sdm-resource-role"
---
# Resource: sdm_role

A Role is a collection of access grants, and typically corresponds to a team, Active Directory OU, or other organizational unit. Users are granted access to resources by assigning them to roles.
## Example Usage

```hcl
resource "sdm_role" "all-dbs" {
    name = "all-database-access"
    composite = true
}

resource "sdm_role" "redis-only" {
    name = "redis-role"
}
```
## Argument Reference
The following arguments are supported by the Role resource:
* `name` - (Required) Unique human-readable name of the Role.
* `composite` - (Optional) True if the Role is a composite role.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Role resource:
* `id` - A unique identifier for the Role resource.
