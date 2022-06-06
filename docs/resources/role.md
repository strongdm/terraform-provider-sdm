---
page_title: "SDM: sdm_role"
description: |-
  Provides settings for Role.
layout: “sdm”
sidebar_current: “docs-sdm-resource-role"
---
# Resource: sdm_role

A Role has a list of access rules which determine which Resources the members
 of the Role have access to. An Account can be a member of multiple Roles via
 AccountAttachments.
## Example Usage

```hcl
resource "sdm_role" "engineers" {
    name = "engineers"
    tags = {
        foo = "bar"
    }
}

resource "sdm_role" "example-role" {
    name = "example-role"
    access_rules = jsonencode([
        {
            "tags": {
                "env": "staging"
            }
        },
        {
            "type": "postgres",
            "tags": {
                "region": "us-west",
                "env": "dev"
            }
        },
        {
            "ids": ["rs-093e6f3061eb4dad"]
        }
    ])
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.

## Argument Reference
The following arguments are supported by the Role resource:
* `access_rules` - (Optional) AccessRules is a list of access rules defining the resources this Role has access to.
* `composite` - (Optional) Composite is true if the Role is a composite role.  Deprecated: composite roles are deprecated, use multi-role via AccountAttachments instead.
* `name` - (Required) Unique human-readable name of the Role.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Role resource:
* `id` - A unique identifier for the Role resource.
## Import
Role can be imported using the id, e.g.,

```
$ terraform import Role.example r-12345678
```
