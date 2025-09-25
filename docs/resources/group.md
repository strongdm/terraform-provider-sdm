---
page_title: "SDM: sdm_group"
description: |-
  Provides settings for Group.
layout: “sdm”
sidebar_current: “docs-sdm-resource-group"
---
# Resource: sdm_group

A Group is a named set of principals.
## Example Usage

```hcl
resource "sdm_group" "security_team" {
  name = "Security Team"
  tags = {
    admin = "true"
  }
}

resource "sdm_group" "administrators" {
  name = "Administrators"
}

resource "sdm_group" "devops_team" {
  name = "DevOps Team"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the Group resource:
* `description` - (Optional) Description of the Group.
* `name` - (Required) Unique human-readable name of the Group.
* `tags` - (Optional) Tags is a map of key/value pairs that can be attached to a Group.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Group resource:
* `id` - A unique identifier for the Group resource.
## Import
A Group can be imported using the id, e.g.,

```
$ terraform import sdm_group.example group-12345678
```
