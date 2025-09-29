---
page_title: "SDM: sdm_secret_engine"
description: |-
  Provides settings for SecretEngine.
layout: “sdm”
sidebar_current: “docs-sdm-resource-secret-engine"
---
# Resource: sdm_secret_engine

A SecretEngine is managing secrets in SecretStores.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the SecretEngine resource:
* active_directory:
	* `binddn` - (Required) Distinguished name of object to bind when performing user and group search. Example: cn=vault,ou=Users,dc=example,dc=com
	* `bindpass` - (Required) Password to use along with binddn when performing user search.
	* `certificate` - (Optional) CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.
	* `connection_timeout` - (Optional) Timeout, in seconds, when attempting to connect to the LDAP server before trying the next URL in the configuration.
	* `do_not_validate_timestamps` - (Optional) If set to true this will prevent password change timestamp validation in Active Directory when validating credentials
	* `insecure_tls` - (Optional) If true, skips LDAP server SSL certificate verification - insecure, use with caution!
	* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
	* `name` - (Required) Unique human-readable name of the Secret Engine.
	* `request_timeout` - (Optional) Timeout, in seconds, for the connection when making requests against the server before returning back an error.
	* `secret_store_id` - (Required) Backing secret store identifier
	* `secret_store_root_path` - (Required) Backing Secret Store root path where managed secrets are going to be stored
	* `start_tls` - (Optional) If true, issues a StartTLS command after establishing an unencrypted connection.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `upndomain` - (Optional) The domain (userPrincipalDomain) used to construct a UPN string for authentication.
	* `url` - (Required) The LDAP server to connect to.
	* `userdn` - (Optional) Base DN under which to perform user search. Example: ou=Users,dc=example,dc=com
* key_value:
	* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
	* `name` - (Required) Unique human-readable name of the Secret Engine.
	* `secret_store_id` - (Required) Backing secret store identifier
	* `secret_store_root_path` - (Required) Backing Secret Store root path where managed secrets are going to be stored
	* `tags` - (Optional) Tags is a map of key, value pairs.
* postgres_secret_engine:
	* `hostname` - (Required) Hostname is the hostname or IP address of the Postgres server.
	* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
	* `name` - (Required) Unique human-readable name of the Secret Engine.
	* `password` - (Required) Password is the password to connect to the Postgres server.
	* `port` - (Required) Port is the port number of the Postgres server.
	* `secret_store_id` - (Required) Backing secret store identifier
	* `secret_store_root_path` - (Required) Backing Secret Store root path where managed secrets are going to be stored
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required) Username is the username to connect to the Postgres server.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the SecretEngine resource:
* `id` - A unique identifier for the SecretEngine resource.
* active_directory:
	* `public_key` - Public key linked with a secret engine
* key_value:
	* `public_key` - Public key linked with a secret engine
* postgres_secret_engine:
	* `public_key` - Public key linked with a secret engine
## Import
A SecretEngine can be imported using the id, e.g.,

```
$ terraform import sdm_secret_engine.example eng-12345678
```
