---
page_title: "SDM: sdm_remote_identity"
description: |-
  Query for existing RemoteIdentities instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-remote-identity"
---
# Data Source: sdm_remote_identity

RemoteIdentities define the username to be used for a specific account
 when connecting to a remote resource using that group.
## Example Usage

```hcl

```
## Argument Reference
The following arguments are supported by a RemoteIdentities data source:
* `account_id` - (Optional) The account for this remote identity.
* `id` - (Optional) Unique identifier of the RemoteIdentity.
* `remote_identity_group_id` - (Optional) The remote identity group.
* `username` - (Optional) The username to be used as the remote identity for this account.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a RemoteIdentities data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_remote_identity ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `remote_identities` - A list where each element has the following attributes:
	* `account_id` - The account for this remote identity.
	* `id` - Unique identifier of the RemoteIdentity.
	* `remote_identity_group_id` - The remote identity group.
	* `username` - The username to be used as the remote identity for this account.
