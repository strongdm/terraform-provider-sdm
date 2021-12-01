---
page_title: "SDM: sdm_resource"
description: |-
  Provides settings for Resource.
layout: “sdm”
sidebar_current: “docs-sdm-resource-resource"
---
# Resource: sdm_resource

A Resource is a database or server for which strongDM manages access.
## Example Usage

```hcl
resource "sdm_resource" "redis-test" {
    redis {
        name = "redis-test"
        hostname = "example.com"
        port_override = 4020
        tags = {
            region = "us-west"
            env = "dev"
        }
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
        tags = {
            region = "us-west"
            env = "dev"
        }
    }
}
```
## Argument Reference
The following arguments are supported by the Resource resource:
* aks:
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aks_basic_auth:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* aks_service_account:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
* aks_service_account_user_impersonation:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
* aks_user_impersonation:
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_user_impersonation:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_es:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazonmq_amqp_091:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* athena:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `output` - (Required) 
	* `region` - (Optional) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aurora_mysql:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* aurora_postgres:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* aws:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_region` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* azure_postgres:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* big_query:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `private_key` - (Optional) 
	* `secret_store_private_key_path` - (Optional)
	* `secret_store_private_key_key` - (Optional)
	* `project` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
* cassandra:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* citus:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* clustrix:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* cockroach:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* db_2_i:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* db_2_luw:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* document_db_host:
	* `auth_database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* document_db_replica_set:
	* `auth_database` - (Required) 
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* druid:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* dynamo_db:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* elastic:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* elasticache_redis:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
* google_gke:
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Optional) 
	* `secret_store_service_account_key_path` - (Optional)
	* `secret_store_service_account_key_key` - (Optional)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* google_gke_user_impersonation:
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Optional) 
	* `secret_store_service_account_key_path` - (Optional)
	* `secret_store_service_account_key_key` - (Optional)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* greenplum:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* http_auth:
	* `auth_header` - (Optional) 
	* `secret_store_auth_header_path` - (Optional)
	* `secret_store_auth_header_key` - (Optional)
	* `default_path` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) 
	* `healthcheck_path` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) 
* http_basic_auth:
	* `default_path` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) 
	* `healthcheck_path` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* http_no_auth:
	* `default_path` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) 
	* `healthcheck_path` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) 
* kubernetes:
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* kubernetes_basic_auth:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* kubernetes_service_account:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
* kubernetes_service_account_user_impersonation:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
* kubernetes_user_impersonation:
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* maria:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* memcached:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* memsql:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_host:
	* `auth_database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_legacy_host:
	* `auth_database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `replica_set` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_legacy_replicaset:
	* `auth_database` - (Required) 
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_replica_set:
	* `auth_database` - (Required) 
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mysql:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* neptune:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* neptune_iam:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* oracle:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* postgres:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* presto:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* rabbitmq_amqp_091:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* raw_tcp:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* rdp:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* redis:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* redshift:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* single_store:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* snowflake:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `schema` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* sql_server:
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `schema` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* ssh:
	* `allow_deprecated_key_exchanges` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* ssh_cert:
	* `allow_deprecated_key_exchanges` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* ssh_customer_key:
	* `allow_deprecated_key_exchanges` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `private_key` - (Optional) 
	* `secret_store_private_key_path` - (Optional)
	* `secret_store_private_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* sybase:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* sybase_iq:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* teradata:
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Resource resource:
* `id` - A unique identifier for the Resource resource.
* amazon_es:
	* `port_override` - 
* amazonmq_amqp_091:
	* `port_override` - 
* athena:
	* `port_override` - 
* aurora_mysql:
	* `port_override` - 
* aurora_postgres:
	* `port_override` - 
* azure_postgres:
	* `port_override` - 
* big_query:
	* `port_override` - 
* cassandra:
	* `port_override` - 
* citus:
	* `port_override` - 
* clustrix:
	* `port_override` - 
* cockroach:
	* `port_override` - 
* db_2_i:
	* `port_override` - 
* db_2_luw:
	* `port_override` - 
* document_db_host:
	* `port_override` - 
* document_db_replica_set:
	* `port_override` - 
* druid:
	* `port_override` - 
* dynamo_db:
	* `port_override` - 
* elastic:
	* `port_override` - 
* elasticache_redis:
	* `port_override` - 
* greenplum:
	* `port_override` - 
* maria:
	* `port_override` - 
* memcached:
	* `port_override` - 
* memsql:
	* `port_override` - 
* mongo_host:
	* `port_override` - 
* mongo_legacy_host:
	* `port_override` - 
* mongo_legacy_replicaset:
	* `port_override` - 
* mongo_replica_set:
	* `port_override` - 
* mysql:
	* `port_override` - 
* neptune:
	* `port_override` - 
* neptune_iam:
	* `port_override` - 
* oracle:
	* `port_override` - 
* postgres:
	* `port_override` - 
* presto:
	* `port_override` - 
* rabbitmq_amqp_091:
	* `port_override` - 
* raw_tcp:
	* `port_override` - 
* rdp:
	* `port_override` - 
* redis:
	* `port_override` - 
* redshift:
	* `port_override` - 
* single_store:
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
