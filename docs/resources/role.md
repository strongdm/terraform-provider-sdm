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

resource "sdm_role" "k8s-admin" {
    name = "k8s-admin"
    access_rules = jsonencode([
        {
            "tags": {
                "env": "production"
            },
            "privileges": {
                "k8s": {
                    "groups": ["system:masters"]
                }
            }
        }
    ])
}

resource "sdm_role" "k8s-developers" {
    name = "k8s-developers"
    access_rules = jsonencode([
        {
            "type": "amazon_eks",
            "tags": {
                "env": "dev"
            },
            "privileges": {
                "k8s": {
                    "groups": ["developers", "viewers"]
                }
            }
        },
        {
            "type": "kubernetes",
            "tags": {
                "region": "us-west"
            },
            "privileges": {
                "k8s": {
                    "groups": ["edit"]
                }
            }
        }
    ])
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the Role resource:
* `access_rules` - (Optional) AccessRules is a list of access rules defining the resources this Role has access to.
* `name` - (Required) Unique human-readable name of the Role.
* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Role resource:
* `id` - A unique identifier for the Role resource.
## Import
A Role can be imported using the id, e.g.,

```
$ terraform import sdm_role.example r-12345678
```
