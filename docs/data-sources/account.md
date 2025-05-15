---
page_title: "SDM: sdm_account"
description: |-
  Query for existing Accounts instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-account"
---
# Data Source: sdm_account

Accounts are users that have access to strongDM. The types of accounts are:
 1. **Users:** humans who are authenticated through username and password or SSO.
 2. **Service Accounts:** machines that are authenticated using a service token.
 3. **Tokens** are access keys with permissions that can be used for authentication.
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

data "sdm_account" "api-key-queries" {
    type = "api"
    name = "*-dev"
}

data "sdm_account" "admin-token-queries" {
    type = "admin-token"
    name = "*-prod"
}
```
## Argument Reference
The following arguments are supported by a Accounts data source:
* `type` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters) for more information.
* `account_type` - (Optional) Corresponds to the type of token, e.g. api or admin-token.
* `email` - (Optional) The User's email address. Must be unique.
* `external_id` - (Optional) External ID is an alternative unique ID this user is represented by within an external service.
* `first_name` - (Optional) The User's first name.
* `id` - (Optional) Unique identifier of the User.
* `last_name` - (Optional) The User's last name.
* `name` - (Optional) Unique human-readable name of the Token.
* `permission_level` - (Optional) PermissionLevel is the user's permission level e.g. admin, DBA, user.
* `permissions` - (Optional) Permissions assigned to the token, e.g. role:create.
* `suspended` - (Optional) Reserved for future use.  Always false for tokens.
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
	* token:
		* `account_type` - Corresponds to the type of token, e.g. api or admin-token.
		* `deadline` - The timestamp when the Token will expire.
		* `duration` - Duration from token creation to expiration.
		* `id` - Unique identifier of the Token.
		* `name` - Unique human-readable name of the Token.
		* `permissions` - Permissions assigned to the token, e.g. role:create.
		* `rekeyed` - The timestamp when the Token was last rekeyed.
		* `suspended` - Reserved for future use.  Always false for tokens.
		* `tags` - Tags is a map of key, value pairs.
	* user:
		* `email` - The User's email address. Must be unique.
		* `external_id` - External ID is an alternative unique ID this user is represented by within an external service.
		* `first_name` - The User's first name.
		* `id` - Unique identifier of the User.
		* `last_name` - The User's last name.
		* `managed_by` - Managed By is a read only field for what service manages this user, e.g. StrongDM, Okta, Azure.
		* `permission_level` - PermissionLevel is the user's permission level e.g. admin, DBA, user.
		* `suspended` - Suspended is a read only field for the User's suspended state.
		* `tags` - Tags is a map of key, value pairs.
