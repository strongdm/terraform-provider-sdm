---
page_title: "SDM: sdm_policy"
description: |-
  Query for existing Policies instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-policy"
---
# Data Source: sdm_policy

Policy is a collection of one or more statements that enforce fine-grained access control
 for the users of an organization.
## Example Usage

```hcl
data "sdm_policy" "policy" {
    name = "policy-query"
}

```
## Argument Reference
The following arguments are supported by a Policies data source:
* `description` - (Optional) Optional description of the Policy.
* `id` - (Optional) Unique identifier of the Policy.
* `name` - (Optional) Unique human-readable name of the Policy.
* `policy` - (Optional) The content of the Policy, in Cedar policy language.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Policies data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_policy ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `policies` - A list where each element has the following attributes:
	* `description` - Optional description of the Policy.
	* `id` - Unique identifier of the Policy.
	* `name` - Unique human-readable name of the Policy.
	* `policy` - The content of the Policy, in Cedar policy language.
