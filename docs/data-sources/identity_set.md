---
page_title: "SDM: sdm_identity_set"
description: |-
  Query for existing IdentitySets instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-identity-set"
---
# Data Source: sdm_identity_set

A IdentitySet defines a group of identity aliases.
## Example Usage

```hcl
data "sdm_identity_set" "default" {
    name = "default"
}
```
## Argument Reference
The following arguments are supported by a IdentitySets data source:
* `id` - (Optional) Unique identifier of the IdentitySet.
* `name` - (Optional) Unique human-readable name of the IdentitySet.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a IdentitySets data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_identity_set ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `identity_sets` - A list where each element has the following attributes:
	* `id` - Unique identifier of the IdentitySet.
	* `name` - Unique human-readable name of the IdentitySet.
