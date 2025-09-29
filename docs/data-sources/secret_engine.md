---
page_title: "SDM: sdm_secret_engine"
description: |-
  Query for existing SecretEngines instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-secret-engine"
---
# Data Source: sdm_secret_engine

A SecretEngine is managing secrets in SecretStores.
## Argument Reference
The following arguments are supported by a SecretEngines data source:
* `type` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://docs.strongdm.com/references/cli/filters/) for more information.
* `binddn` - (Optional) Distinguished name of object to bind when performing user and group search. Example: cn=vault,ou=Users,dc=example,dc=com
* `bindpass` - (Optional) Password to use along with binddn when performing user search.
* `certificate` - (Optional) CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.
* `connection_timeout` - (Optional) Timeout, in seconds, when attempting to connect to the LDAP server before trying the next URL in the configuration.
* `do_not_validate_timestamps` - (Optional) If set to true this will prevent password change timestamp validation in Active Directory when validating credentials
* `hostname` - (Optional) Hostname is the hostname or IP address of the Postgres server.
* `id` - (Optional) Unique identifier of the Secret Engine.
* `insecure_tls` - (Optional) If true, skips LDAP server SSL certificate verification - insecure, use with caution!
* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
* `name` - (Optional) Unique human-readable name of the Secret Engine.
* `password` - (Optional) Password is the password to connect to the Postgres server.
* `port` - (Optional) Port is the port number of the Postgres server.
* `request_timeout` - (Optional) Timeout, in seconds, for the connection when making requests against the server before returning back an error.
* `secret_store_id` - (Optional) Backing secret store identifier
* `secret_store_root_path` - (Optional) Backing Secret Store root path where managed secrets are going to be stored
* `start_tls` - (Optional) If true, issues a StartTLS command after establishing an unencrypted connection.
* `tags` - (Optional) Tags is a map of key, value pairs.
* `upndomain` - (Optional) The domain (userPrincipalDomain) used to construct a UPN string for authentication.
* `url` - (Optional) The LDAP server to connect to.
* `userdn` - (Optional) Base DN under which to perform user search. Example: ou=Users,dc=example,dc=com
* `username` - (Optional) Username is the username to connect to the Postgres server.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a SecretEngines data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_secret_engine ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `secret_engines` - A single element list containing a map, where each key lists one of the following objects:
	* active_directory:
		* `after_read_ttl` - The default time-to-live duration of the password after it's read. Once the ttl has passed, a password will be rotated.
		* `binddn` - Distinguished name of object to bind when performing user and group search. Example: cn=vault,ou=Users,dc=example,dc=com
		* `bindpass` - Password to use along with binddn when performing user search.
		* `certificate` - CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.
		* `connection_timeout` - Timeout, in seconds, when attempting to connect to the LDAP server before trying the next URL in the configuration.
		* `do_not_validate_timestamps` - If set to true this will prevent password change timestamp validation in Active Directory when validating credentials
		* `id` - Unique identifier of the Secret Engine.
		* `insecure_tls` - If true, skips LDAP server SSL certificate verification - insecure, use with caution!
		* `key_rotation_interval_days` - An interval of public/private key rotation for secret engine in days
		* `max_backoff_duration` - The maximum retry duration in case of automatic failure. On failed ttl rotation attempt it will be retried in an increasing intervals until it reaches max_backoff_duration
		* `name` - Unique human-readable name of the Secret Engine.
		* `public_key` - Public key linked with a secret engine
		* `request_timeout` - Timeout, in seconds, for the connection when making requests against the server before returning back an error.
		* `secret_store_id` - Backing secret store identifier
		* `secret_store_root_path` - Backing Secret Store root path where managed secrets are going to be stored
		* `start_tls` - If true, issues a StartTLS command after establishing an unencrypted connection.
		* `tags` - Tags is a map of key, value pairs.
		* `ttl` - The default password time-to-live duration. Once the ttl has passed, a password will be rotated the next time it's requested.
		* `upndomain` - The domain (userPrincipalDomain) used to construct a UPN string for authentication.
		* `url` - The LDAP server to connect to.
		* `userdn` - Base DN under which to perform user search. Example: ou=Users,dc=example,dc=com
	* key_value:
		* `id` - Unique identifier of the Secret Engine.
		* `key_rotation_interval_days` - An interval of public/private key rotation for secret engine in days
		* `name` - Unique human-readable name of the Secret Engine.
		* `public_key` - Public key linked with a secret engine
		* `secret_store_id` - Backing secret store identifier
		* `secret_store_root_path` - Backing Secret Store root path where managed secrets are going to be stored
		* `tags` - Tags is a map of key, value pairs.
	* postgres_secret_engine:
		* `hostname` - Hostname is the hostname or IP address of the Postgres server.
		* `id` - Unique identifier of the Secret Engine.
		* `key_rotation_interval_days` - An interval of public/private key rotation for secret engine in days
		* `name` - Unique human-readable name of the Secret Engine.
		* `password` - Password is the password to connect to the Postgres server.
		* `port` - Port is the port number of the Postgres server.
		* `public_key` - Public key linked with a secret engine
		* `secret_store_id` - Backing secret store identifier
		* `secret_store_root_path` - Backing Secret Store root path where managed secrets are going to be stored
		* `tags` - Tags is a map of key, value pairs.
		* `username` - Username is the username to connect to the Postgres server.
