---
page_title: "SDM: sdm_account"
description: |-
  Query for existing Accounts instances.
---
# Data Source: sdm_account

Accounts are users that have access to strongDM.
 There are two types of accounts:
 1. **Regular users:** humans who are authenticated through username and password or SSO
 2. **Service users:** machines that are authneticated using a service token
## Example Usage

```hcl
data "sdm_account" "user-queries" {
    type = "user"
    email = "*@strongdm.com"
}
```
## Argument Reference
The following arguments are supported by a Accounts data source:
* `type` - (Optional) a filter to query only one subtype. See Attribute Reference for all subtypes.
* `email` - (Optional) The User's email address. Must be unique.
* `first_name` - (Optional) The User's first name.
* `id` - (Optional) Unique identifier of the Service.
* `last_name` - (Optional) The User's last name.
* `name` - (Optional) Unique human-readable name of the Service.
* `suspended` - (Optional) The Service's suspended state.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Accounts data source:
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `accounts` - A single element list containing a map, where each key lists one of the following objects:
	* user:
		* `id` - Unique identifier of the User.
		* `email` - The User's email address. Must be unique.
		* `first_name` - The User's first name.
		* `last_name` - The User's last name.
		* `suspended` - The User's suspended state.
	* service:
		* `id` - Unique identifier of the Service.
		* `name` - Unique human-readable name of the Service.
		* `suspended` - The Service's suspended state.
