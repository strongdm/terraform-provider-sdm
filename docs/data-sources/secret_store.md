---
page_title: "SDM: sdm_secret_store"
description: |-
  Query for existing SecretStores instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-secret-store"
---
# Data Source: sdm_secret_store

A SecretStore is a server where resource secrets (passwords, keys) are stored. 
 Coming soon support for HashiCorp Vault and AWS Secret Store. Contact support@strongdm.com to request access to the beta.
## Argument Reference
The following arguments are supported by a SecretStores data source:
* `type` - (Optional) a filter to query only one subtype. See Attribute Reference for all subtypes.
* `ca_cert_path` - (Optional) 
* `client_cert_path` - (Optional) 
* `client_key_path` - (Optional) 
* `id` - (Optional) Unique identifier of the SecretStore.
* `name` - (Optional) Unique human-readable name of the SecretStore.
* `server_address` - (Optional) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a SecretStores data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_secret_store ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `secret_stores` - A single element list containing a map, where each key lists one of the following objects:
	* vault_tls:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `server_address` - 
		* `ca_cert_path` - 
		* `client_cert_path` - 
		* `client_key_path` - 
		* `tags` - Tags is a map of key, value pairs.
	* vault_token:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `server_address` - 
		* `tags` - Tags is a map of key, value pairs.
