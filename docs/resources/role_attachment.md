---
page_title: "SDM: sdm_role_attachment"
description: |-
  Provides settings for RoleAttachment.
layout: “sdm”
sidebar_current: “docs-sdm-resource-role-attachment"
---
# Resource: sdm_role_attachment

A RoleAttachment assigns a role to a composite role.
## Example Usage

```hcl
resource "sdm_role_attachment" "test_role_attachment" {
    composite_role_id = "r-12345123"
    attached_role_id = "r-3453433"
}
```
## Argument Reference
The following arguments are supported by the RoleAttachment resource:
* `composite_role_id` - (Required) The id of the composite role of this RoleAttachment.
* `attached_role_id` - (Required) The id of the attached role of this RoleAttachment.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the RoleAttachment resource:
* `id` - A unique identifier for the RoleAttachment resource.
