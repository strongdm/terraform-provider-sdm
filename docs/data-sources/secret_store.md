---
page_title: "SDM: sdm_secret_store"
description: |-
  Query for existing SecretStores instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-secret-store"
---
# Data Source: sdm_secret_store

A SecretStore is a server where resource secrets (passwords, keys) are stored.
 Coming soon support for HashiCorp Vault and AWS Secret Store.
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
		* `region` - The AWS region to target e.g. us-east-1
		* `tags` - Tags is a map of key, value pairs.
	* azure_store:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
		* `vault_uri` - The URI of the key vault to target e.g. https://myvault.vault.azure.net
	* cyberark_conjur:
		* `app_url` - The URL of the Cyberark instance
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
	* cyberark_pam:
		* `app_url` - The URL of the Cyberark instance
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
	* cyberark_pam_experimental:
		* `app_url` - The URL of the Cyberark instance
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `tags` - Tags is a map of key, value pairs.
	* delinea_store:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `server_url` - The URL of the Delinea instance
		* `tags` - Tags is a map of key, value pairs.
		* `tenant_name` - The tenant name to target
	* gcp_store:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `project_id` - The GCP project ID to target.
		* `tags` - Tags is a map of key, value pairs.
	* vault_approle:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `namespace` - The namespace to make requests within
		* `server_address` - The URL of the Vault to target
		* `tags` - Tags is a map of key, value pairs.
	* vault_tls:
		* `ca_cert_path` - A path to a CA file accessible by a Node
		* `client_cert_path` - A path to a client certificate file accessible by a Node
		* `client_key_path` - A path to a client key file accessible by a Node
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `namespace` - The namespace to make requests within
		* `server_address` - The URL of the Vault to target
		* `tags` - Tags is a map of key, value pairs.
	* vault_token:
		* `id` - Unique identifier of the SecretStore.
		* `name` - Unique human-readable name of the SecretStore.
		* `namespace` - The namespace to make requests within
		* `server_address` - The URL of the Vault to target
		* `tags` - Tags is a map of key, value pairs.
