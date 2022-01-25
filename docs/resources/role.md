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
    tags = {
        region = "us-west"
        env = "dev"
    }    
}

resource "sdm_role" "redis-only" {
    name = "redis-role"
    tags = {
        region = "us-west"
        env = "dev"
    }    
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.

## Argument Reference
The following arguments are supported by the Role resource:
* `access_rules` - (Optional) AccessRules JSON encoded access rules data.
* `composite` - (Optional) True if the Role is a composite role.
* `name` - (Required) Unique human-readable name of the Role.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Role resource:
* `id` - A unique identifier for the Role resource.
