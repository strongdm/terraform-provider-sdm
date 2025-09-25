---
page_title: "SDM: sdm_group"
description: |-
  Query for existing Groups instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-group"
---
# Data Source: sdm_group

A Group is a named set of principals.
## Example Usage

```hcl
data "sdm_group" "security_team" {
  name = "Security Team"
}

data "sdm_group" "administrators" {
  id = "group-1234567890abcdef"
}

data "sdm_group" "admin_teams" {
  tags = {
    admin = "true"
  }
}
```
## Argument Reference
The following arguments are supported by a Groups data source:
* `description` - (Optional) Description of the Group.
* `id` - (Optional) Unique identifier of the Group.
* `name` - (Optional) Unique human-readable name of the Group.
* `tags` - (Optional) Tags is a map of key/value pairs that can be attached to a Group.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Groups data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_group ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `groups` - A list where each element has the following attributes:
	* `description` - Description of the Group.
	* `id` - Unique identifier of the Group.
	* `name` - Unique human-readable name of the Group.
	* `source` - Source is a read only field for what service manages this group, e.g. StrongDM, Okta, Azure.
	* `tags` - Tags is a map of key/value pairs that can be attached to a Group.
