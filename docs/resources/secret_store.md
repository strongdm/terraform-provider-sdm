---
page_title: "SDM: sdm_secret_store"
description: |-
  Provides settings for SecretStore.
layout: “sdm”
sidebar_current: “docs-sdm-resource-secret-store"
---
# Resource: sdm_secret_store

A SecretStore is a server where resource secrets (passwords, keys) are stored. 
 Coming soon support for HashiCorp Vault and AWS Secret Store. Contact support@strongdm.com to request access to the beta.
## Argument Reference
The following arguments are supported by the SecretStore resource:
* vault_tls:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `server_address` - (Required) 
	* `ca_cert_path` - (Optional) 
	* `client_cert_path` - (Required) 
	* `client_key_path` - (Required) 
* vault_token:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `server_address` - (Required) 
* aws:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `region` - (Required) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the SecretStore resource:
* `id` - A unique identifier for the SecretStore resource.
