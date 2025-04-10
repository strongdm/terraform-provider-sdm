---
page_title: "SDM: sdm_managed_secret"
description: |-
  Query for existing ManagedSecrets instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-managed-secret"
---
# Data Source: sdm_managed_secret

ManagedSecret contains details about managed secret
## Argument Reference
The following arguments are supported by a ManagedSecrets data source:
* `id` - (Optional) Unique identifier of the Managed Secret.
* `name` - (Optional) Unique human-readable name of the Managed Secret.
* `policy` - (Optional) Password and rotation policy for the secret
* `secret_engine_id` - (Optional) An ID of a Secret Engine linked with the Managed Secret.
* `tags` - (Optional) Tags is a map of key, value pairs.
* `value` - (Optional) Sensitive value of the secret.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a ManagedSecrets data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_managed_secret ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `managed_secrets` - A list where each element has the following attributes:
	* `config` - public part of the secret value
	* `expires_at` - Timestamp of when secret is going to be rotated
	* `id` - Unique identifier of the Managed Secret.
	* `last_rotated_at` - Timestamp of when secret was last rotated
	* `name` - Unique human-readable name of the Managed Secret.
	* `secret_engine_id` - An ID of a Secret Engine linked with the Managed Secret.
	* `secret_store_path` - Path in a secret store.
	* `tags` - Tags is a map of key, value pairs.
	* `value` - Sensitive value of the secret.
