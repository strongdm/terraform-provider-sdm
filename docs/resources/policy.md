---
page_title: "SDM: sdm_policy"
description: |-
  Provides settings for Policy.
layout: “sdm”
sidebar_current: “docs-sdm-resource-policy"
---
# Resource: sdm_policy

Policy is a collection of one or more statements that enforce fine-grained access control
 for the users of an organization.
## Example Usage

```hcl
resource "sdm_policy" "permit_everything" {
    name = "permit-everything"
    description = "this policy permits everything"
    policy = <<EOP
permit(principal, action, resource);
EOP
}

```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the Policy resource:
* `description` - (Optional) Optional description of the Policy.
* `name` - (Required) Unique human-readable name of the Policy.
* `policy` - (Optional) The content of the Policy, in Cedar policy language.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Policy resource:
* `id` - A unique identifier for the Policy resource.
## Import
A Policy can be imported using the id, e.g.,

```
$ terraform import sdm_policy.example po-12345678
```
