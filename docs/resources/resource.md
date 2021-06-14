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
* rabbitmq_amqp_091:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `tls_required` - (Optional) 
* amazonmq_amqp_091:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `tls_required` - (Optional) 
* athena:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `output` - (Required) 
	* `region` - (Optional) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
* aws:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `healthcheck_region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
* big_query:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `private_key` - (Optional) 
	* `secret_store_private_key_path` - (Optional)
	* `secret_store_private_key_key` - (Optional)
	* `project` - (Required) 
	* `endpoint` - (Required) 
	* `username` - (Optional) 
* cassandra:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* db_2_i:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* db_2_luw:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
* druid:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
* dynamo_db:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `region` - (Required) 
	* `endpoint` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
* amazon_es:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `region` - (Required) 
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `endpoint` - (Optional) 
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
* elastic:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* http_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
	* `subdomain` - (Required) 
* http_no_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
	* `subdomain` - (Required) 
* http_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `url` - (Required) 
	* `healthcheck_path` - (Required) 
	* `auth_header` - (Optional) 
	* `secret_store_auth_header_path` - (Optional)
	* `secret_store_auth_header_key` - (Optional)
	* `headers_blacklist` - (Optional) 
	* `default_path` - (Optional) 
	* `subdomain` - (Required) 
* kubernetes:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* kubernetes_user_impersonation:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* kubernetes_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* kubernetes_service_account:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* kubernetes_service_account_user_impersonation:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* amazon_eks:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `region` - (Required) 
	* `cluster_name` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* amazon_eks_user_impersonation:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `secret_access_key` - (Optional) 
	* `secret_store_secret_access_key_path` - (Optional)
	* `secret_store_secret_access_key_key` - (Optional)
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `region` - (Required) 
	* `cluster_name` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* google_gke:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `service_account_key` - (Optional) 
	* `secret_store_service_account_key_path` - (Optional)
	* `secret_store_service_account_key_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* google_gke_user_impersonation:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `service_account_key` - (Optional) 
	* `secret_store_service_account_key_path` - (Optional)
	* `secret_store_service_account_key_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* aks:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* aks_user_impersonation:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Optional) 
	* `secret_store_client_certificate_path` - (Optional)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Optional) 
	* `secret_store_client_key_path` - (Optional)
	* `secret_store_client_key_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* aks_basic_auth:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* aks_service_account:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* aks_service_account_user_impersonation:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Required) 
	* `token` - (Optional) 
	* `secret_store_token_path` - (Optional)
	* `secret_store_token_key` - (Optional)
	* `healthcheck_namespace` - (Optional) 
* memcached:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `port` - (Optional) 
* mongo_legacy_host:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `replica_set` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_legacy_replicaset:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `replica_set` - (Required) 
	* `connect_to_replica` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_host:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* mongo_replica_set:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `auth_database` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `replica_set` - (Required) 
	* `connect_to_replica` - (Optional) 
	* `tls_required` - (Optional) 
* mysql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
* aurora_mysql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
* clustrix:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
* maria:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
* memsql:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
* oracle:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Required) 
	* `tls_required` - (Optional) 
* postgres:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* aurora_postgres:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* greenplum:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* cockroach:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* redshift:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* citus:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* presto:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `port` - (Optional) 
	* `username` - (Optional) 
	* `tls_required` - (Optional) 
* rdp:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
* redis:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
* elasticache_redis:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `tls_required` - (Optional) 
* snowflake:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `schema` - (Required) 
* sql_server:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `database` - (Required) 
	* `schema` - (Optional) 
	* `port` - (Optional) 
	* `override_database` - (Optional) 
* ssh:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `allow_deprecated_key_exchanges` - (Optional) 
* ssh_cert:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `allow_deprecated_key_exchanges` - (Optional) 
* ssh_customer_key:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `port` - (Required) 
	* `private_key` - (Optional) 
	* `secret_store_private_key_path` - (Optional)
	* `secret_store_private_key_key` - (Optional)
	* `port_forwarding` - (Optional) 
	* `allow_deprecated_key_exchanges` - (Optional) 
* sybase:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `port` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
* sybase_iq:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `port` - (Optional) 
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
* teradata:
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Resource resource:
* `id` - A unique identifier for the Resource resource.
* rabbitmq_amqp_091:
	* `port_override` - 
* amazonmq_amqp_091:
	* `port_override` - 
* athena:
	* `port_override` - 
* big_query:
	* `port_override` - 
* cassandra:
	* `port_override` - 
* db_2_i:
	* `port_override` - 
* db_2_luw:
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
