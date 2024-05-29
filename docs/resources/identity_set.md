---
page_title: "SDM: sdm_identity_set"
description: |-
  Provides settings for IdentitySet.
layout: “sdm”
sidebar_current: “docs-sdm-resource-identity-set"
---
# Resource: sdm_identity_set

A IdentitySet defines a group of identity aliases.
## Example Usage

```hcl
resource "sdm_identity_set" "default" {
    name = "default"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the IdentitySet resource:
* `name` - (Required) Unique human-readable name of the IdentitySet.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the IdentitySet resource:
* `id` - A unique identifier for the IdentitySet resource.
## Import
A IdentitySet can be imported using the id, e.g.,

```
$ terraform import sdm_identity_set.example ig-12345678
```
