---
page_title: "SDM: sdm_account"
description: |-
  Provides settings for Account.
layout: “sdm”
sidebar_current: “docs-sdm-resource-account"
---
# Resource: sdm_account

Accounts are users that have access to strongDM. The types of accounts are:
 1. **Users:** humans who are authenticated through username and password or SSO.
 2. **Service Accounts:** machines that are authenticated using a service token.
 3. **Tokens** are access keys with permissions that can be used for authentication.
## Example Usage

```hcl
resource "sdm_account" "test-user" {
    user {
        first_name = "al"
        last_name = "bob"
        email = "albob@strongdm.com"
        permission_level  = "database-admin"
        manager_id = "a-1234abc"
        tags = {
            region = "us-west"
            env = "dev"
        }
    }
}

resource "sdm_account" "test-service" {
    service {
        name = "test-service"
        tags = {
            region = "us-west"
            env = "dev"
        }    
    }
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the Account resource:
* service:
	* `name` - (Required) Unique human-readable name of the Service.
	* `suspended` - (Optional) The Service's suspended state.
	* `tags` - (Optional) Tags is a map of key, value pairs.

* user:
	* `email` - (Required) The User's email address. Must be unique.
	* `external_id` - (Optional) External ID is an alternative unique ID this user is represented by within an external service.
	* `first_name` - (Required) The User's first name.
	* `last_name` - (Required) The User's last name.
	* `manager_id` - (Optional) Manager ID is the ID of the user's manager. This field is empty when the user has no manager.
	* `permission_level` - (Optional) PermissionLevel is the user's permission level e.g. admin, DBA, user.
	* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Account resource:
* `id` - A unique identifier for the Account resource.
* service:
	* `created_at` - CreatedAt is the timestamp when the service was created

* user:
	* `scim` - SCIM contains the raw SCIM metadata for the user. This is a read-only field.
	* `created_at` - CreatedAt is the timestamp when the user was created
	* `managed_by` - Managed By is a read only field for what service manages this user, e.g. StrongDM, Okta, Azure.
	* `resolved_manager_id` - Resolved Manager ID is the ID of the user's manager derived from the manager_id, if present, or from the SCIM metadata. This is a read-only field that's only populated for get and list.
	* `suspended` - Suspended is a read only field for the User's suspended state.
## Import
A Account can be imported using the id, e.g.,

```
$ terraform import sdm_account.example a-12345678
```
