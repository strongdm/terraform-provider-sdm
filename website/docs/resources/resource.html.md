---
page_title: "SDM: sdm_resource"
description: |-
  Provides settings for Resource.
layout: “sdm”
sidebar_current: “docs-sdm-resource-resource"
---
# sdm_resource

A Resource is a database or server for which strongDM manages access.
## Example Usage

```hcl
resource "sdm_resource" "redis-test" {
    redis {
        name = "redis-test"
        hostname = "example.com"
        port_override = 4020
    }
}

resource "sdm_resource" "postgres-test" {
    postgres {
        name = "postgres-test"
        hostname = "example.com"
        database = "my-db"
        username = "admin"
        password = "hunter2"
        port = 5432
    }
}
```
## Argument Reference
The following arguments are supported by the Resource resource:
* athena:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `access_key` - (Required) 
	* `secret_access_key` - (Required) 
	* `output` - (Required) 
	* `region` - (Optional) 
* big_query:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `private_key` - (Required) 
	* `project` - (Required) 
	* `endpoint` - (Required) 
	* `username` - (Optional) 
* cassandra:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* db_2:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
* druid:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
* dynamo_db:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `access_key` - (Required) 
	* `secret_access_key` - (Required) 
	* `region` - (Required) 
	* `endpoint` - (Required) 
* amazon_es:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `secret_access_key` - (Optional) 
	* `endpoint` - (Optional) 
	* `access_key` - (Optional) 
* elastic:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* http_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
	* `subdomain` - (Required) 
* http_no_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
	* `subdomain` - (Required) 
* http_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `auth_header` - (Required) 
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
	* `subdomain` - (Required) 
* kubernetes:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `certificate_authority` - (Optional) 
	* `certificate_authority_filename` - (Optional) 
	* `client_certificate` - (Optional) 
	* `client_certificate_filename` - (Optional) 
	* `client_key` - (Optional) 
	* `client_key_filename` - (Optional) 
	* `healthcheck_namespace` - (Optional) 
* kubernetes_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `healthcheck_namespace` - (Optional) 
* kubernetes_service_account:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Required) 
	* `healthcheck_namespace` - (Optional) 
* amazon_eks:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `endpoint` - (Required) 
	* `access_key` - (Required) 
	* `secret_access_key` - (Required) 
	* `certificate_authority` - (Required) 
	* `certificate_authority_filename` - (Optional) 
	* `region` - (Required) 
	* `cluster_name` - (Required) 
	* `role_arn` - (Optional) 
	* `healthcheck_namespace` - (Optional) 
* google_gke:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `endpoint` - (Required) 
	* `certificate_authority` - (Required) 
	* `certificate_authority_filename` - (Optional) 
	* `service_account_key` - (Required) 
	* `service_account_key_filename` - (Optional) 
	* `healthcheck_namespace` - (Optional) 
* aks:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `certificate_authority` - (Optional) 
	* `certificate_authority_filename` - (Optional) 
	* `client_certificate` - (Optional) 
	* `client_certificate_filename` - (Optional) 
	* `client_key` - (Optional) 
	* `client_key_filename` - (Optional) 
	* `healthcheck_namespace` - (Optional) 
* aks_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `healthcheck_namespace` - (Optional) 
* aks_service_account:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Required) 
	* `healthcheck_namespace` - (Optional) 
* memcached:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Optional) 
* mongo_legacy_host:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `replica_set` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_legacy_replicaset:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `replica_set` - (Required) 
	* `connect_to_replica` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_host:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_replica_set:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `replica_set` - (Required) 
	* `connect_to_replica` - (Optional) 
	* `tls_required` - (Optional) 
* mysql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
* aurora_mysql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
* clustrix:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
* maria:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
* memsql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
* oracle:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Required) 
	* `tls_required` - (Optional) 
* postgres:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* aurora_postgres:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* greenplum:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* cockroach:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* redshift:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* citus:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* presto:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Optional) 
	* `username` - (Optional) 
	* `tls_required` - (Optional) 
* rdp:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port` - (Required) 
* redis:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `password` - (Optional) 
	* `port` - (Optional) 
* elasticache_redis:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* snowflake:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `schema` - (Required) 
* sql_server:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `schema` - (Optional) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* ssh:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `allow_deprecated_key_exchanges` - (Optional) 
* ssh_cert:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `allow_deprecated_key_exchanges` - (Optional) 
* sybase:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `port` - (Optional) 
	* `password` - (Optional) 
* sybase_iq:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `port` - (Optional) 
	* `password` - (Optional) 
* teradata:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port` - (Optional) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Resource resource:
* `id` - A unique identifier for the Resource resource.
* athena:
	* `port_override` - 
* big_query:
	* `port_override` - 
* cassandra:
	* `port_override` - 
* db_2:
	* `port_override` - 
* druid:
	* `port_override` - 
* dynamo_db:
	* `port_override` - 
* amazon_es:
	* `port_override` - 
* elastic:
	* `port_override` - 
* memcached:
	* `port_override` - 
* mongo_legacy_host:
	* `port_override` - 
* mongo_legacy_replicaset:
	* `port_override` - 
* mongo_host:
	* `port_override` - 
* mongo_replica_set:
	* `port_override` - 
* mysql:
	* `port_override` - 
* aurora_mysql:
	* `port_override` - 
* clustrix:
	* `port_override` - 
* maria:
	* `port_override` - 
* memsql:
	* `port_override` - 
* oracle:
	* `port_override` - 
* postgres:
	* `port_override` - 
* aurora_postgres:
	* `port_override` - 
* greenplum:
	* `port_override` - 
* cockroach:
	* `port_override` - 
* redshift:
	* `port_override` - 
* citus:
	* `port_override` - 
* presto:
	* `port_override` - 
* rdp:
	* `port_override` - 
* redis:
	* `port_override` - 
* elasticache_redis:
	* `port_override` - 
* snowflake:
	* `port_override` - 
* sql_server:
	* `port_override` - 
* ssh:
	* `public_key` - 
* sybase:
	* `port_override` - 
* sybase_iq:
	* `port_override` - 
* teradata:
	* `port_override` - 
