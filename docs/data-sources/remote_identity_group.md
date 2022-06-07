---
page_title: "SDM: sdm_remote_identity_group"
description: |-
  Query for existing RemoteIdentityGroups instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-remote-identity-group"
---
# Data Source: sdm_remote_identity_group

A RemoteIdentityGroup defines a group of remote identities.
## Example Usage

```hcl

```
## Argument Reference
The following arguments are supported by a RemoteIdentityGroups data source:
* `id` - (Optional) Unique identifier of the RemoteIdentityGroup.
* `name` - (Optional) Unique human-readable name of the RemoteIdentityGroup.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a RemoteIdentityGroups data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_remote_identity_group ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `remote_identity_groups` - A list where each element has the following attributes:
	* `id` - Unique identifier of the RemoteIdentityGroup.
	* `name` - Unique human-readable name of the RemoteIdentityGroup.
