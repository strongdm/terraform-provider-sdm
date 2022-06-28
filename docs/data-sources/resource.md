---
page_title: "SDM: sdm_resource"
description: |-
  Query for existing Resources instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-resource"
---
# Data Source: sdm_resource

A Resource is a database, server, cluster, website, or cloud that strongDM
 delegates access to.
## Example Usage

```hcl
data "sdm_resource" "mysql_datasources" {
    name = "us-west*"
    type = "mysql"
    tags = {
        region = "us-west"
        env = "dev"
    }    
}
```
## Argument Reference
The following arguments are supported by a Resources data source:
* `type` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.
* `hostname` - (Optional) 
* `id` - (Optional) Unique identifier of the Resource.
* `name` - (Optional) Unique human-readable name of the Resource.
* `port` - (Optional) 
* `tags` - (Optional) Tags is a map of key, value pairs.
* `username` - (Optional) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Resources data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_resource ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `resources` - A single element list containing a map, where each key lists one of the following objects:
	* aks:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `remote_identity_group_id` - 
		* `remote_identity_healthcheck_username` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* aks_basic_auth:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* aks_service_account:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `remote_identity_group_id` - 
		* `remote_identity_healthcheck_username` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `token` - 
	* aks_service_account_user_impersonation:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `token` - 
	* aks_user_impersonation:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* amazon_eks:
		* `access_key` - 
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `cluster_name` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `region` - 
		* `remote_identity_group_id` - 
		* `remote_identity_healthcheck_username` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `secret_access_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* amazon_eks_user_impersonation:
		* `access_key` - 
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `cluster_name` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `region` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `secret_access_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* amazon_es:
		* `access_key` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - 
		* `region` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `secret_access_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* amazonmq_amqp_091:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* athena:
		* `access_key` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `output` - 
		* `port_override` - 
		* `region` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `secret_access_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* aurora_mysql:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* aurora_postgres:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* aws:
		* `access_key` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_region` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `role_arn` - 
		* `role_external_id` - 
		* `secret_access_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* azure:
		* `app_id` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tenant_id` - 
	* azure_certificate:
		* `app_id` - 
		* `bind_interface` - Bind interface
		* `client_certificate` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tenant_id` - 
	* azure_postgres:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* big_query:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - 
		* `private_key` - 
		* `project` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* cassandra:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* citus:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* clustrix:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* cockroach:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* db_2_i:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* db_2_luw:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* document_db_host:
		* `auth_database` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* document_db_replica_set:
		* `auth_database` - 
		* `bind_interface` - Bind interface
		* `connect_to_replica` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port_override` - 
		* `replica_set` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* druid:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* dynamo_db:
		* `access_key` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - 
		* `region` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `secret_access_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* elastic:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* elasticache_redis:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
	* gcp:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `keyfile` - 
		* `name` - Unique human-readable name of the Resource.
		* `scopes` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* google_gke:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `remote_identity_group_id` - 
		* `remote_identity_healthcheck_username` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `service_account_key` - 
		* `tags` - Tags is a map of key, value pairs.
	* google_gke_user_impersonation:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `service_account_key` - 
		* `tags` - Tags is a map of key, value pairs.
	* greenplum:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* http_auth:
		* `auth_header` - 
		* `bind_interface` - Bind interface
		* `default_path` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `headers_blacklist` - 
		* `healthcheck_path` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - 
		* `tags` - Tags is a map of key, value pairs.
		* `url` - 
	* http_basic_auth:
		* `bind_interface` - Bind interface
		* `default_path` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `headers_blacklist` - 
		* `healthcheck_path` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - 
		* `tags` - Tags is a map of key, value pairs.
		* `url` - 
		* `username` - 
	* http_no_auth:
		* `bind_interface` - Bind interface
		* `default_path` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `headers_blacklist` - 
		* `healthcheck_path` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - 
		* `tags` - Tags is a map of key, value pairs.
		* `url` - 
	* kubernetes:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `remote_identity_group_id` - 
		* `remote_identity_healthcheck_username` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* kubernetes_basic_auth:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* kubernetes_service_account:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `remote_identity_group_id` - 
		* `remote_identity_healthcheck_username` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `token` - 
	* kubernetes_service_account_user_impersonation:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `token` - 
	* kubernetes_user_impersonation:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* maria:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* memcached:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* memsql:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* mongo_host:
		* `auth_database` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* mongo_legacy_host:
		* `auth_database` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `replica_set` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* mongo_legacy_replicaset:
		* `auth_database` - 
		* `bind_interface` - Bind interface
		* `connect_to_replica` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `replica_set` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* mongo_replica_set:
		* `auth_database` - 
		* `bind_interface` - Bind interface
		* `connect_to_replica` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `replica_set` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* mongo_sharded_cluster:
		* `auth_database` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* mtls_mysql:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `server_name` - 
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* mtls_postgres:
		* `bind_interface` - Bind interface
		* `certificate_authority` - 
		* `client_certificate` - 
		* `client_key` - 
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `server_name` - 
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* mysql:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* neptune:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* neptune_iam:
		* `access_key` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `port_override` - 
		* `region` - 
		* `role_arn` - 
		* `role_external_id` - 
		* `secret_access_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* oracle:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* postgres:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* presto:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* rabbitmq_amqp_091:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - 
		* `username` - 
	* raw_tcp:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* rdp:
		* `bind_interface` - Bind interface
		* `downgrade_nla_connections` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* redis:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
	* redshift:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* single_store:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* snowflake:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port_override` - 
		* `schema` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* sql_server:
		* `bind_interface` - Bind interface
		* `database` - 
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - 
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `schema` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* ssh:
		* `allow_deprecated_key_exchanges` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `key_type` - 
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `port_forwarding` - 
		* `port_override` - 
		* `public_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* ssh_cert:
		* `allow_deprecated_key_exchanges` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `port_forwarding` - 
		* `port_override` - 
		* `remote_identity_group_id` - 
		* `remote_identity_healthcheck_username` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* ssh_customer_key:
		* `allow_deprecated_key_exchanges` - 
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - 
		* `port_forwarding` - 
		* `port_override` - 
		* `private_key` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* sybase:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* sybase_iq:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
	* teradata:
		* `bind_interface` - Bind interface
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - 
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - 
		* `port` - 
		* `port_override` - 
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `tags` - Tags is a map of key, value pairs.
		* `username` - 
