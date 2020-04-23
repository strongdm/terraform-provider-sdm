---
page_title: "SDM: sdm_resource"
description: |-
  Query for existing Resources instances.
---
# Data Source: sdm_resource

A Resource is a server or service which clients connect to through relays.
## Example Usage

```hcl
data "sdm_resource" "aurora_datasources" {
    name = "us-west*"
    type = "aurora_mysql"
}
```
## Argument Reference
The following arguments are supported by a Resources data source:
* `type` - (Optional) a filter to query only one subtype. See Attribute Reference for all subtypes.
* `hostname` - (Optional) 
* `id` - (Optional) Unique identifier of the Resource.
* `name` - (Optional) Unique human-readable name of the Resource.
* `port` - (Optional) 
* `username` - (Optional) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Resources data source:
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `resources` - A single element list containing a map, where each key lists one of the following objects:
	* athena:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `access_key` - 
		* `secret_access_key` - 
		* `output` - 
		* `port_override` - 
		* `region` - 
	* big_query:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `private_key` - 
		* `project` - 
		* `port_override` - 
		* `endpoint` - 
		* `username` - 
	* cassandra:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
		* `tls_required` - 
	* druid:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port_override` - 
		* `username` - 
		* `password` - 
		* `port` - 
	* dynamo_db:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `access_key` - 
		* `secret_access_key` - 
		* `region` - 
		* `endpoint` - 
		* `port_override` - 
	* amazon_es:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `region` - 
		* `secret_access_key` - 
		* `endpoint` - 
		* `access_key` - 
		* `port_override` - 
	* elastic:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
		* `tls_required` - 
	* http_basic_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `url` - 
		* `healthcheck_path` - 
		* `username` - 
		* `password` - 
		* `headers_blacklist` - 
		* `default_path` - 
	* http_no_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `url` - 
		* `healthcheck_path` - 
		* `headers_blacklist` - 
		* `default_path` - 
	* http_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `url` - 
		* `healthcheck_path` - 
		* `auth_header` - 
		* `headers_blacklist` - 
		* `default_path` - 
	* kubernetes:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port` - 
		* `certificate_authority` - 
		* `certificate_authority_filename` - 
		* `client_certificate` - 
		* `client_certificate_filename` - 
		* `client_key` - 
		* `client_key_filename` - 
	* kubernetes_basic_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port` - 
		* `username` - 
		* `password` - 
	* kubernetes_service_account:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port` - 
		* `token` - 
	* amazon_eks:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `endpoint` - 
		* `access_key` - 
		* `secret_access_key` - 
		* `certificate_authority` - 
		* `certificate_authority_filename` - 
		* `region` - 
		* `cluster_name` - 
		* `role_arn` - 
	* google_gke:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `endpoint` - 
		* `certificate_authority` - 
		* `certificate_authority_filename` - 
		* `service_account_key` - 
		* `service_account_key_filename` - 
	* aks:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port` - 
		* `certificate_authority` - 
		* `certificate_authority_filename` - 
		* `client_certificate` - 
		* `client_certificate_filename` - 
		* `client_key` - 
		* `client_key_filename` - 
	* aks_basic_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port` - 
		* `username` - 
		* `password` - 
	* aks_service_account:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port` - 
		* `token` - 
	* memcached:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port_override` - 
		* `port` - 
	* mongo_legacy_host:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `auth_database` - 
		* `port_override` - 
		* `username` - 
		* `password` - 
		* `port` - 
		* `replica_set` - 
		* `tls_required` - 
	* mongo_legacy_replicaset:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `auth_database` - 
		* `port_override` - 
		* `username` - 
		* `password` - 
		* `port` - 
		* `replica_set` - 
		* `connect_to_replica` - 
		* `tls_required` - 
	* mongo_host:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `auth_database` - 
		* `port_override` - 
		* `username` - 
		* `password` - 
		* `port` - 
		* `tls_required` - 
	* mongo_replica_set:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `auth_database` - 
		* `port_override` - 
		* `username` - 
		* `password` - 
		* `port` - 
		* `replica_set` - 
		* `connect_to_replica` - 
		* `tls_required` - 
	* mysql:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* aurora_mysql:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* clustrix:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* maria:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* memsql:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* oracle:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port` - 
		* `port_override` - 
		* `tls_required` - 
	* postgres:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `override_database` - 
	* aurora_postgres:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `override_database` - 
	* greenplum:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `override_database` - 
	* cockroach:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `override_database` - 
	* redshift:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `override_database` - 
	* presto:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `username` - 
		* `tls_required` - 
	* rdp:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
	* redis:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port_override` - 
		* `password` - 
		* `port` - 
	* elasticache_redis:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `port_override` - 
		* `password` - 
		* `port` - 
		* `tls_required` - 
	* snowflake:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `schema` - 
		* `port_override` - 
	* sql_server:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `schema` - 
		* `port` - 
		* `override_database` - 
	* ssh:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `port` - 
		* `public_key` - 
		* `port_forwarding` - 
	* sybase:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `port_override` - 
		* `port` - 
		* `password` - 
	* sybase_iq:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `port_override` - 
		* `port` - 
		* `password` - 
	* teradata:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
