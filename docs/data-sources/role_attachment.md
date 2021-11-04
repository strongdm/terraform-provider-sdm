---
page_title: "SDM: sdm_role_attachment"
description: |-
  Query for existing RoleAttachments instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-role-attachment"
---
# Data Source: sdm_role_attachment

A RoleAttachment assigns a role to a composite role.
## Example Usage

```hcl
data "sdm_role_attachment" "role_attachment_query" {
    composite_role_id = "r-233332245"
}
```
## Argument Reference
The following arguments are supported by a RoleAttachments data source:
* `attached_role_id` - (Optional) The id of the attached role of this RoleAttachment.
* `composite_role_id` - (Optional) The id of the composite role of this RoleAttachment.
* `id` - (Optional) Unique identifier of the RoleAttachment.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a RoleAttachments data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_role_attachment ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `role_attachments` - A list where each element has the following attributes:
	* `attached_role_id` - The id of the attached role of this RoleAttachment.
	* `composite_role_id` - The id of the composite role of this RoleAttachment.
	* `id` - Unique identifier of the RoleAttachment.
