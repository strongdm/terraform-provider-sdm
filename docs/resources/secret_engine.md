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
	* `binddn` - (Required) Distinguished name of object to bind when performing user and group search. Example: cn=vault,ou=Users,dc=example,dc=com. With path_credentials it is a secret store path in the resource format path?key=username, otherwise it is the literal username.
	* `bindpass` - (Required) Bindpass is the password to use along with binddn when performing user search. With path_credentials it is a secret store path in the resource format path?key=password, otherwise it is the literal password.
	* `certificate` - (Optional) CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.
	* `connection_timeout` - (Optional) Timeout, in seconds, when attempting to connect to the LDAP server before trying the next URL in the configuration.
	* `do_not_validate_timestamps` - (Optional) If set to true this will prevent password change timestamp validation in Active Directory when validating credentials
	* `insecure_tls` - (Optional) If true, skips LDAP server SSL certificate verification - insecure, use with caution!
	* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
	* `name` - (Required) Unique human-readable name of the Secret Engine.
	* `node_selector` - (Optional) node selector is used to narrow down the nodes used to communicate with with secret engine
	* `path_credentials` - (Optional) PathCredentials is true when username and password are secret store paths, which the control plane stores. When false, they are literal values and the password is rotated and kept in the engine's secret store.
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
	* `node_selector` - (Optional) node selector is used to narrow down the nodes used to communicate with with secret engine
	* `secret_store_id` - (Required) Backing secret store identifier
	* `secret_store_root_path` - (Required) Backing Secret Store root path where managed secrets are going to be stored
	* `tags` - (Optional) Tags is a map of key, value pairs.
* mysql_secret_engine:
	* `database` - (Required) Database is the database to verify credential against.
	* `hostname` - (Required) Hostname is the hostname or IP address of the MySQL server.
	* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
	* `name` - (Required) Unique human-readable name of the Secret Engine.
	* `node_selector` - (Optional) node selector is used to narrow down the nodes used to communicate with with secret engine
	* `password` - (Required) Password is the password to connect to the MySQL server. With path_credentials it is a secret store path in the resource format path?key=password, otherwise it is the literal password.
	* `path_credentials` - (Optional) PathCredentials is true when username and password are secret store paths, which the control plane stores. When false, they are literal values and the password is rotated and kept in the engine's secret store.
	* `port` - (Required) Port is the port number of the MySQL server.
	* `secret_store_id` - (Required) Backing secret store identifier
	* `secret_store_root_path` - (Required) Backing Secret Store root path where managed secrets are going to be stored
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls` - (Optional) TLS enables TLS/SSL when connecting to the MySQL server.
	* `tls_skip_verify` - (Optional) TLS disable certificate verification
	* `username` - (Required) Username is the username to connect to the MySQL server. With path_credentials it is a secret store path in the resource format path?key=username, otherwise it is the literal username.
* postgres_secret_engine:
	* `database` - (Required) Database is the database to verify credential against.
	* `hostname` - (Required) Hostname is the hostname or IP address of the Postgres server.
	* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
	* `name` - (Required) Unique human-readable name of the Secret Engine.
	* `node_selector` - (Optional) node selector is used to narrow down the nodes used to communicate with with secret engine
	* `password` - (Required) Password is the password to connect to the Postgres server. With path_credentials it is a secret store path in the resource format path?key=password, otherwise it is the literal password.
	* `path_credentials` - (Optional) PathCredentials is true when username and password are secret store paths, which the control plane stores. When false, they are literal values and the password is rotated and kept in the engine's secret store.
	* `port` - (Required) Port is the port number of the Postgres server.
	* `secret_store_id` - (Required) Backing secret store identifier
	* `secret_store_root_path` - (Required) Backing Secret Store root path where managed secrets are going to be stored
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls` - (Optional) TLS enables TLS/SSL when connecting to the Postgres server.
	* `username` - (Required) Username is the username to connect to the Postgres server. With path_credentials it is a secret store path in the resource format path?key=username, otherwise it is the literal username.
* sqlserver_secret_engine:
	* `database` - (Required) Database is the database to verify credential against.
	* `hostname` - (Required) Hostname is the hostname or IP address of the SQL Server.
	* `key_rotation_interval_days` - (Optional) An interval of public/private key rotation for secret engine in days
	* `name` - (Required) Unique human-readable name of the Secret Engine.
	* `node_selector` - (Optional) node selector is used to narrow down the nodes used to communicate with with secret engine
	* `password` - (Required) Password is the password to connect to the SQL Server server. With path_credentials it is a secret store path in the resource format path?key=password, otherwise it is the literal password.
	* `path_credentials` - (Optional) PathCredentials is true when username and password are secret store paths, which the control plane stores. When false, they are literal values and the password is rotated and kept in the engine's secret store.
	* `port` - (Required) Port is the port number of the SQL Server server.
	* `secret_store_id` - (Required) Backing secret store identifier
	* `secret_store_root_path` - (Required) Backing Secret Store root path where managed secrets are going to be stored
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls` - (Optional) TLS enables TLS/SSL when connecting to the SQL Server server.
	* `tls_skip_verify` - (Optional) TLS disable certificate verification
	* `username` - (Required) Username is the username to connect to the SQL Server. With path_credentials it is a secret store path in the resource format path?key=username, otherwise it is the literal username.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the SecretEngine resource:
* `id` - A unique identifier for the SecretEngine resource.
* active_directory:
	* `public_key` - Public key linked with a secret engine
* key_value:
	* `public_key` - Public key linked with a secret engine
* mysql_secret_engine:
	* `public_key` - Public key linked with a secret engine
* postgres_secret_engine:
	* `public_key` - Public key linked with a secret engine
* sqlserver_secret_engine:
	* `public_key` - Public key linked with a secret engine
## Import
A SecretEngine can be imported using the id, e.g.,

```
$ terraform import sdm_secret_engine.example eng-12345678
```
