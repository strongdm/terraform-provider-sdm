---
page_title: "SDM: sdm_secret_store"
description: |-
  Provides settings for SecretStore.
layout: “sdm”
sidebar_current: “docs-sdm-resource-secret-store"
---
# Resource: sdm_secret_store

A SecretStore is a server where resource secrets (passwords, keys) are stored.
 Coming soon support for HashiCorp Vault and AWS Secret Store.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the SecretStore resource:
* aws:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `region` - (Required) The AWS region to target e.g. us-east-1
	* `tags` - (Optional) Tags is a map of key, value pairs.
* azure_store:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `vault_uri` - (Required) The URI of the key vault to target e.g. https://myvault.vault.azure.net
* cyberark_conjur:
	* `app_url` - (Required) The URL of the Cyberark instance
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* cyberark_pam:
	* `app_url` - (Required) The URL of the Cyberark instance
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* cyberark_pam_experimental:
	* `app_url` - (Required) The URL of the Cyberark instance
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* delinea_store:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `server_url` - (Optional) The URL of the Delinea instance
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_name` - (Optional) The tenant name to target
* gcp_store:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `project_id` - (Required) The GCP project ID to target.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* gcp_cert_x_509_store:
	* `ca_id` - (Optional) The ID of the target CA
	* `ca_pool_id` - (Required) The ID of the target CA pool
	* `location` - (Required) The Region for the CA in GCP format e.g. us-west1
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `project_id` - (Required) The GCP project ID to target.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_approle:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_approle_cert_ssh:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `ssh_mount_point` - (Required) The mount point of the SSH engine configured with the desired CA
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_approle_cert_x509:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `pki_mount_point` - (Required) The mount point of the PKI engine configured with the desired CA
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_tls:
	* `ca_cert_path` - (Optional) A path to a CA file accessible by a Node
	* `client_cert_path` - (Required) A path to a client certificate file accessible by a Node
	* `client_key_path` - (Required) A path to a client key file accessible by a Node
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_tls_cert_ssh:
	* `ca_cert_path` - (Optional) A path to a CA file accessible by a Node
	* `client_cert_path` - (Required) A path to a client certificate file accessible by a Node
	* `client_key_path` - (Required) A path to a client key file accessible by a Node
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `ssh_mount_point` - (Required) The mount point of the SSH engine configured with the desired CA
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_tls_cert_x509:
	* `ca_cert_path` - (Optional) A path to a CA file accessible by a Node
	* `client_cert_path` - (Required) A path to a client certificate file accessible by a Node
	* `client_key_path` - (Required) A path to a client key file accessible by a Node
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `pki_mount_point` - (Required) The mount point of the PKI engine configured with the desired CA
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_token:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_token_cert_ssh:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `ssh_mount_point` - (Required) The mount point of the SSH engine configured with the desired CA
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_token_cert_x509:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `pki_mount_point` - (Required) The mount point of the PKI engine configured with the desired CA
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the SecretStore resource:
* `id` - A unique identifier for the SecretStore resource.
## Import
A SecretStore can be imported using the id, e.g.,

```
$ terraform import sdm_secret_store.example se-12345678
```
