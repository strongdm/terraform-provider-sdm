---
page_title: "SDM: sdm_resource"
description: |-
  Provides settings for Resource.
---
# sdm_resource

A Resource is a server or service which clients connect to through relays.
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
	* `port_override` - (Required) 
	* `region` - (Optional) 
* big_query:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `private_key` - (Required) 
	* `project` - (Required) 
	* `port_override` - (Required) 
	* `endpoint` - (Required) 
	* `username` - (Optional) 
* cassandra:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* druid:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port_override` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
* dynamo_db:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `access_key` - (Required) 
	* `secret_access_key` - (Required) 
	* `region` - (Required) 
	* `endpoint` - (Required) 
	* `port_override` - (Required) 
* amazon_es:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `secret_access_key` - (Optional) 
	* `endpoint` - (Optional) 
	* `access_key` - (Optional) 
	* `port_override` - (Required) 
* elastic:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port_override` - (Required) 
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
* http_no_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
* http_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `auth_header` - (Required) 
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
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
* kubernetes_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
* kubernetes_service_account:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Required) 
* amazon_eks:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `endpoint` - (Required) 
	* `access_key` - (Optional) 
	* `secret_access_key` - (Optional) 
	* `certificate_authority` - (Required) 
	* `certificate_authority_filename` - (Optional) 
	* `region` - (Required) 
	* `cluster_name` - (Required) 
	* `role_arn` - (Optional) 
* google_gke:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `endpoint` - (Required) 
	* `certificate_authority` - (Required) 
	* `certificate_authority_filename` - (Optional) 
	* `service_account_key` - (Required) 
	* `service_account_key_filename` - (Optional) 
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
* aks_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
* aks_service_account:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Required) 
* memcached:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
* mongo_legacy_host:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `port_override` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `replica_set` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_legacy_replicaset:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `port_override` - (Required) 
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
	* `port_override` - (Required) 
	* `username` - (Optional) 
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_replica_set:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `port_override` - (Required) 
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
	* `port_override` - (Required) 
	* `port` - (Optional) 
* aurora_mysql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
* clustrix:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
* maria:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
* memsql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
* oracle:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port` - (Required) 
	* `port_override` - (Required) 
	* `tls_required` - (Optional) 
* postgres:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* aurora_postgres:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* greenplum:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* cockroach:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* redshift:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* presto:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `username` - (Optional) 
	* `tls_required` - (Optional) 
* rdp:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
* redis:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port_override` - (Required) 
	* `password` - (Optional) 
	* `port` - (Optional) 
* elasticache_redis:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `port_override` - (Required) 
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
	* `port_override` - (Required) 
* sql_server:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `database` - (Required) 
	* `port_override` - (Required) 
	* `schema` - (Optional) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* ssh:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
* sybase:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `password` - (Optional) 
* sybase_iq:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
	* `password` - (Optional) 
* teradata:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `hostname` - (Required) 
	* `username` - (Required) 
	* `password` - (Required) 
	* `port_override` - (Required) 
	* `port` - (Optional) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Resource resource:
* `id` - A unique identifier for the Resource resource.
* ssh:
	* `public_key` - 
