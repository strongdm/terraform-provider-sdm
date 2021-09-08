---
page_title: "SDM: sdm_resource"
description: |-
  Query for existing Resources instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-resource"
---
# Data Source: sdm_resource

A Resource is a database or server for which strongDM manages access.
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
* `id` - a generated id representing this request, unrelated to input id and sdm_resource ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `resources` - A single element list containing a map, where each key lists one of the following objects:
	* rabbitmq_amqp_091:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port_override` - 
		* `port` - 
		* `username` - 
		* `password` - 
		* `tls_required` - 
	* amazonmq_amqp_091:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port_override` - 
		* `port` - 
		* `username` - 
		* `password` - 
		* `tls_required` - 
	* athena:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `access_key` - 
		* `secret_access_key` - 
		* `output` - 
		* `port_override` - 
		* `region` - 
		* `role_arn` - 
		* `role_external_id` - 
	* aws:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `access_key` - 
		* `secret_access_key` - 
		* `healthcheck_region` - 
		* `role_arn` - 
		* `role_external_id` - 
	* big_query:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `private_key` - 
		* `project` - 
		* `port_override` - 
		* `endpoint` - 
		* `username` - 
	* cassandra:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
		* `tls_required` - 
	* db_2_i:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
		* `tls_required` - 
	* db_2_luw:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* druid:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port_override` - 
		* `username` - 
		* `password` - 
		* `port` - 
	* dynamo_db:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `access_key` - 
		* `secret_access_key` - 
		* `region` - 
		* `endpoint` - 
		* `port_override` - 
		* `role_arn` - 
		* `role_external_id` - 
	* amazon_es:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `region` - 
		* `secret_access_key` - 
		* `endpoint` - 
		* `access_key` - 
		* `port_override` - 
		* `role_arn` - 
		* `role_external_id` - 
	* elastic:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
		* `tls_required` - 
	* http_basic_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `url` - 
		* `healthcheck_path` - 
		* `username` - 
		* `password` - 
		* `headers_blacklist` - 
		* `default_path` - 
		* `subdomain` - 
	* http_no_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `url` - 
		* `healthcheck_path` - 
		* `headers_blacklist` - 
		* `default_path` - 
		* `subdomain` - 
	* http_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `url` - 
		* `healthcheck_path` - 
		* `auth_header` - 
		* `headers_blacklist` - 
		* `default_path` - 
		* `subdomain` - 
	* kubernetes:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* kubernetes_user_impersonation:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* kubernetes_basic_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `username` - 
		* `password` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* kubernetes_service_account:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `token` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* kubernetes_service_account_user_impersonation:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `token` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* amazon_eks:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `access_key` - 
		* `secret_access_key` - 
		* `certificate_authority` - 
		* `region` - 
		* `cluster_name` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* amazon_eks_user_impersonation:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `access_key` - 
		* `secret_access_key` - 
		* `certificate_authority` - 
		* `region` - 
		* `cluster_name` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* google_gke:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `certificate_authority` - 
		* `service_account_key` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* google_gke_user_impersonation:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `certificate_authority` - 
		* `service_account_key` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* aks:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* aks_user_impersonation:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* aks_basic_auth:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `username` - 
		* `password` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* aks_service_account:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `token` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* aks_service_account_user_impersonation:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port` - 
		* `token` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.
	* memcached:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port_override` - 
		* `port` - 
	* mongo_legacy_host:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* aurora_mysql:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* clustrix:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* maria:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* memsql:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* single_store:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
	* oracle:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `override_database` - 
	* citus:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `password` - 
		* `database` - 
		* `port_override` - 
		* `port` - 
		* `username` - 
		* `tls_required` - 
	* raw_tcp:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port_override` - 
		* `port` - 
	* rdp:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
	* redis:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port_override` - 
		* `password` - 
		* `port` - 
	* elasticache_redis:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `port_override` - 
		* `password` - 
		* `port` - 
		* `tls_required` - 
	* snowflake:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `database` - 
		* `schema` - 
		* `port_override` - 
	* sql_server:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
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
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `port` - 
		* `public_key` - 
		* `port_forwarding` - 
		* `allow_deprecated_key_exchanges` - 
	* ssh_cert:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `port` - 
		* `port_forwarding` - 
		* `allow_deprecated_key_exchanges` - 
	* ssh_customer_key:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `port` - 
		* `private_key` - 
		* `port_forwarding` - 
		* `allow_deprecated_key_exchanges` - 
	* sybase:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `port_override` - 
		* `port` - 
		* `password` - 
	* sybase_iq:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `port_override` - 
		* `port` - 
		* `password` - 
	* teradata:
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `tags` - Tags is a map of key, value pairs.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `username` - 
		* `password` - 
		* `port_override` - 
		* `port` - 
