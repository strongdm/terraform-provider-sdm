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
* active_directory_store:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `server_address` - (Required) Hostname of server that is hosting NDES (Network Device Enrollment Services). Often this is the same host as Active Directory Certificate Services
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aws:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `region` - (Required) The AWS region to target e.g. us-east-1
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aws_cert_x509:
	* `ca_arn` - (Required) The ARN of the CA in AWS Private CA
	* `certificate_template_arn` - (Required) The ARN of the AWS certificate template for requested certificates. Must allow SAN, key usage, and ext key usage passthrough from CSR
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA represented in minutes.
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `region` - (Required) The AWS region to target e.g. us-east-1
	* `signing_algo` - (Required) The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key. e.g. SHA256WITHRSA
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
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA represented in minutes.
	* `location` - (Required) The Region for the CA in GCP format e.g. us-west1
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `project_id` - (Required) The GCP project ID to target.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* keyfactor_x_509_store:
	* `ca_file_path` - (Optional) Path to the root CA that signed the certificate passed to the client for HTTPS connection. This is not required if the CA is trusted by the host operating system. This should be a PEM formatted certificate, and doesn't necessarily have to be the CA that signed CertificateFile.
	* `certificate_file_path` - (Required) Path to client certificate in PEM format. This certificate must contain a client certificate that is recognized by the EJBCA instance represented by Hostname. This PEM file may also contain the private key associated with the certificate, but KeyFile can also be set to configure the private key.
	* `default_certificate_authority_name` - (Required) Name of EJBCA certificate authority that will enroll CSR.
	* `default_certificate_profile_name` - (Required) Certificate profile name that EJBCA will enroll the CSR with.
	* `default_end_entity_profile_name` - (Required) End entity profile that EJBCA will enroll the CSR with.
	* `enrollment_code_env_var` - (Optional) code used by EJBCA during enrollment. May be left blank if no code is required.
	* `enrollment_username_env_var` - (Optional) username that used by the EJBCA during enrollment. This can be left out.  If so, the username must be auto-generated on the Keyfactor side.
	* `key_file_path` - (Optional) Path to private key in PEM format. This file should contain the private key associated with the client certificate configured in CertificateFile.
	* `key_password_env_var` - (Optional) optional environment variable housing the password that is used to decrypt the key file.
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `server_address` - (Required) the host of the Key Factor CA
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_approle:
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_approle_cert_ssh:
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA represented in minutes.
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `ssh_mount_point` - (Required) The mount point of the SSH engine configured with the desired CA
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_approle_cert_x509:
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA in minutes. Recommended value is 5.
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
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA represented in minutes.
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
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA represented in minutes.
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
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA in minutes. Recommended value is 5.
	* `name` - (Required) Unique human-readable name of the SecretStore.
	* `namespace` - (Optional) The namespace to make requests within
	* `server_address` - (Required) The URL of the Vault to target
	* `signing_role` - (Required) The signing role to be used for signing certificates
	* `ssh_mount_point` - (Required) The mount point of the SSH engine configured with the desired CA
	* `tags` - (Optional) Tags is a map of key, value pairs.
* vault_token_cert_x509:
	* `issued_cert_ttl_minutes` - (Required) The lifetime of certificates issued by this CA represented in minutes.
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
