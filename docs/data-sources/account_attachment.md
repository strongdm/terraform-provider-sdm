---
page_title: "SDM: sdm_account_attachment"
description: |-
  Query for existing AccountAttachments instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-account-attachment"
---
# Data Source: sdm_account_attachment

AccountAttachments assign an account to a role or composite role.
## Example Usage

```hcl
data "sdm_account_attachment" "account_attachment_query" {
    account_id = "a-00000054"
}
```
## Argument Reference
The following arguments are supported by a AccountAttachments data source:
* `id` - (Optional) Unique identifier of the AccountAttachment.
* `account_id` - (Optional) The id of the account of this AccountAttachment.
* `role_id` - (Optional) The id of the attached role of this AccountAttachment.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a AccountAttachments data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_account_attachment ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `account_attachments` - A list where each element has the following attributes:
	* `id` - Unique identifier of the AccountAttachment.
	* `account_id` - The id of the account of this AccountAttachment.
	* `role_id` - The id of the attached role of this AccountAttachment.
