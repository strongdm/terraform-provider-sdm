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
* `type` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.
* `id` - (Optional) Unique identifier of the SecretStore.
* `name` - (Optional) Unique human-readable name of the SecretStore.
* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a SecretStores data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_secret_store ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `secret_stores` - A single element list containing a map, where each key lists one of the following objects:
	* aws:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `region` - 
		* `tags` - Tags is a map of key, value pairs.
	* azure_store:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
		* `vault_uri` - 
	* cyberark_conjur:
		* `app_url` - 
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
	* cyberark_pam:
		* `app_url` - 
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
	* cyberark_pam_experimental:
		* `app_url` - 
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
	* delinea_store:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `server_url` - 
		* `tags` - Tags is a map of key, value pairs.
		* `tenant_name` - 
	* gcp_store:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `project_id` - 
		* `tags` - Tags is a map of key, value pairs.
	* vault_approle:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `namespace` - 
		* `server_address` - 
		* `tags` - Tags is a map of key, value pairs.
	* vault_tls:
		* `ca_cert_path` - 
		* `client_cert_path` - 
		* `client_key_path` - 
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `namespace` - 
		* `server_address` - 
		* `tags` - Tags is a map of key, value pairs.
	* vault_token:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `namespace` - 
		* `server_address` - 
		* `tags` - Tags is a map of key, value pairs.
