---
page_title: "SDM: sdm_account"
description: |-
  Query for existing Accounts instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-account"
---
# Data Source: sdm_account

Accounts are users that have access to strongDM. There are two types of accounts:
 1. **Users:** humans who are authenticated through username and password or SSO.
 2. **Service Accounts:** machines that are authenticated using a service token.
## Example Usage

```hcl
data "sdm_account" "user-queries" {
    type = "user"
    email = "*@strongdm.com"
    tags = {
        region = "us-west"
        env = "dev"
    }    
}
```
## Argument Reference
The following arguments are supported by a Accounts data source:
* `type` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.
* `email` - (Optional) The User's email address. Must be unique.
* `first_name` - (Optional) The User's first name.
* `id` - (Optional) Unique identifier of the User.
* `last_name` - (Optional) The User's last name.
* `name` - (Optional) Unique human-readable name of the Service.
* `suspended` - (Optional) The User's suspended state.
* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Accounts data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_account ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `accounts` - A single element list containing a map, where each key lists one of the following objects:
	* service:
		* `id` - Unique identifier of the Service.
		* `name` - Unique human-readable name of the Service.
		* `suspended` - The Service's suspended state.
		* `tags` - Tags is a map of key, value pairs.
	* user:
		* `email` - The User's email address. Must be unique.
		* `first_name` - The User's first name.
		* `id` - Unique identifier of the User.
		* `last_name` - The User's last name.
		* `suspended` - The User's suspended state.
		* `tags` - Tags is a map of key, value pairs.
