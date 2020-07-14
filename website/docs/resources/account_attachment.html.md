---
page_title: "SDM: sdm_account_attachment"
description: |-
  Provides settings for AccountAttachment.
layout: “sdm”
sidebar_current: “docs-sdm-resource-account-attachment"
---
# sdm_account_attachment

AccountAttachments assign an account to a role.
## Example Usage

```hcl
resource "sdm_account_attachment" "test_account_attachment" {
    account_id = "a-00000054"
    role_id = "r-12355562"
}
```
## Argument Reference
The following arguments are supported by the AccountAttachment resource:
* `account_id` - (Required) The id of the account of this AccountAttachment.
* `role_id` - (Required) The id of the attached role of this AccountAttachment.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the AccountAttachment resource:
* `id` - A unique identifier for the AccountAttachment resource.
