---
page_title: "SDM: sdm_resource"
description: |-
  Provides settings for Resource.
layout: “sdm”
sidebar_current: “docs-sdm-resource-resource"
---
# Resource: sdm_resource

A Resource is a database, server, cluster, website, or cloud that strongDM
 delegates access to.
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
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.

## Argument Reference
The following arguments are supported by the Resource resource:
* aks:
	* `bind_interface` - (Optional) Bind interface
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
	* `port_override` - (Optional) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aks_basic_auth:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* aks_service_account:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required if storing credentials directly strongDM) 
	* `secret_store_token_path` - (Required if using credentials stored in a secret store)
	* `secret_store_token_key` - (Optional)
* aks_service_account_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required if storing credentials directly strongDM) 
	* `secret_store_token_path` - (Required if using credentials stored in a secret store)
	* `secret_store_token_key` - (Optional)
* aks_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
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
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks:
	* `access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Required if storing credentials directly strongDM) 
	* `secret_store_certificate_authority_path` - (Required if using credentials stored in a secret store)
	* `secret_store_certificate_authority_key` - (Optional)
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_secret_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_instance_profile:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Required if storing credentials directly strongDM) 
	* `secret_store_certificate_authority_path` - (Required if using credentials stored in a secret store)
	* `secret_store_certificate_authority_key` - (Optional)
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_user_impersonation:
	* `access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Required if storing credentials directly strongDM) 
	* `secret_store_certificate_authority_path` - (Required if using credentials stored in a secret store)
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
	* `secret_access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_secret_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_es:
	* `access_key` - (Optional) 
	* `secret_store_access_key_path` - (Optional)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
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
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* athena:
	* `access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `output` - (Required) 
	* `port_override` - (Optional) 
	* `region` - (Optional) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_secret_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aurora_mysql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* aurora_postgres:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* aws:
	* `access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_region` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_secret_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aws_console:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `enable_env_variables` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `role_arn` - (Required if storing credentials directly strongDM) 
	* `secret_store_role_arn_path` - (Required if using credentials stored in a secret store)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `session_expiry` - (Optional) 
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aws_console_static_key_pair:
	* `access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `role_arn` - (Required if storing credentials directly strongDM) 
	* `secret_store_role_arn_path` - (Required if using credentials stored in a secret store)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_secret_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `session_expiry` - (Optional) 
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
* azure:
	* `app_id` - (Required if storing credentials directly strongDM) 
	* `secret_store_app_id_path` - (Required if using credentials stored in a secret store)
	* `secret_store_app_id_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required if storing credentials directly strongDM) 
	* `secret_store_tenant_id_path` - (Required if using credentials stored in a secret store)
	* `secret_store_tenant_id_key` - (Optional)
* azure_certificate:
	* `app_id` - (Required if storing credentials directly strongDM) 
	* `secret_store_app_id_path` - (Required if using credentials stored in a secret store)
	* `secret_store_app_id_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `client_certificate` - (Required if storing credentials directly strongDM) 
	* `secret_store_client_certificate_path` - (Required if using credentials stored in a secret store)
	* `secret_store_client_certificate_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required if storing credentials directly strongDM) 
	* `secret_store_tenant_id_path` - (Required if using credentials stored in a secret store)
	* `secret_store_tenant_id_key` - (Optional)
* azure_mysql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* azure_postgres:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* big_query:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `private_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_private_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_private_key_key` - (Optional)
	* `project` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
* cassandra:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* citus:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* clustrix:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* cockroach:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* db_2_i:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* db_2_luw:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* document_db_host:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* document_db_replica_set:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port_override` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* druid:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* dynamo_db:
	* `access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_secret_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* elastic:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* elasticache_redis:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* gcp:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `keyfile` - (Required if storing credentials directly strongDM) 
	* `secret_store_keyfile_path` - (Required if using credentials stored in a secret store)
	* `secret_store_keyfile_key` - (Optional)
	* `name` - (Required) Unique human-readable name of the Resource.
	* `scopes` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* google_gke:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Required if storing credentials directly strongDM) 
	* `secret_store_certificate_authority_path` - (Required if using credentials stored in a secret store)
	* `secret_store_certificate_authority_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_service_account_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_service_account_key_key` - (Optional)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* google_gke_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Required if storing credentials directly strongDM) 
	* `secret_store_certificate_authority_path` - (Required if using credentials stored in a secret store)
	* `secret_store_certificate_authority_key` - (Optional)
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_service_account_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_service_account_key_key` - (Optional)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* greenplum:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* http_auth:
	* `auth_header` - (Required if storing credentials directly strongDM) 
	* `secret_store_auth_header_path` - (Required if using credentials stored in a secret store)
	* `secret_store_auth_header_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `default_path` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) 
	* `healthcheck_path` - (Required) 
	* `host_override` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) 
* http_basic_auth:
	* `bind_interface` - (Optional) Bind interface
	* `default_path` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) 
	* `healthcheck_path` - (Required) 
	* `host_override` - (Optional) 
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
	* `bind_interface` - (Optional) Bind interface
	* `default_path` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `headers_blacklist` - (Optional) 
	* `healthcheck_path` - (Required) 
	* `host_override` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) 
* kubernetes:
	* `bind_interface` - (Optional) Bind interface
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
	* `port_override` - (Optional) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* kubernetes_basic_auth:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* kubernetes_service_account:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required if storing credentials directly strongDM) 
	* `secret_store_token_path` - (Required if using credentials stored in a secret store)
	* `secret_store_token_key` - (Optional)
* kubernetes_service_account_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required if storing credentials directly strongDM) 
	* `secret_store_token_path` - (Required if using credentials stored in a secret store)
	* `secret_store_token_key` - (Optional)
* kubernetes_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
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
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* maria:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* memcached:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* memsql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* mongo_host:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_legacy_host:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `replica_set` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_legacy_replicaset:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_replica_set:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mongo_sharded_cluster:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* mtls_mysql:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Required if storing credentials directly strongDM) 
	* `secret_store_client_certificate_path` - (Required if using credentials stored in a secret store)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_client_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_client_key_key` - (Optional)
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_name` - (Optional) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* mtls_postgres:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `secret_store_certificate_authority_path` - (Optional)
	* `secret_store_certificate_authority_key` - (Optional)
	* `client_certificate` - (Required if storing credentials directly strongDM) 
	* `secret_store_client_certificate_path` - (Required if using credentials stored in a secret store)
	* `secret_store_client_certificate_key` - (Optional)
	* `client_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_client_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_client_key_key` - (Optional)
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_name` - (Optional) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* mysql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* neptune:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* neptune_iam:
	* `access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_access_key_key` - (Optional)
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `secret_store_role_arn_path` - (Optional)
	* `secret_store_role_arn_key` - (Optional)
	* `role_external_id` - (Optional) 
	* `secret_store_role_external_id_path` - (Optional)
	* `secret_store_role_external_id_key` - (Optional)
	* `secret_access_key` - (Required if storing credentials directly strongDM) 
	* `secret_store_secret_access_key_path` - (Required if using credentials stored in a secret store)
	* `secret_store_secret_access_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* oracle:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* postgres:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* presto:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* rabbitmq_amqp_091:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* raw_tcp:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* rdp:
	* `bind_interface` - (Optional) Bind interface
	* `downgrade_nla_connections` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* redis:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
	* `secret_store_username_path` - (Optional)
	* `secret_store_username_key` - (Optional)
* redshift:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* single_store:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* snowflake:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port_override` - (Optional) 
	* `schema` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* snowsight:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_username` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `saml_metadata` - (Required if storing credentials directly strongDM) 
	* `secret_store_saml_metadata_path` - (Required if using credentials stored in a secret store)
	* `secret_store_saml_metadata_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
* sql_server:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `schema` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* ssh:
	* `allow_deprecated_key_exchanges` - (Optional) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `key_type` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* ssh_cert:
	* `allow_deprecated_key_exchanges` - (Optional) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `key_type` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `port_override` - (Optional) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* ssh_customer_key:
	* `allow_deprecated_key_exchanges` - (Optional) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `port_override` - (Optional) 
	* `private_key` - (Optional) 
	* `secret_store_private_key_path` - (Optional)
	* `secret_store_private_key_key` - (Optional)
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* sybase:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* sybase_iq:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `secret_store_password_path` - (Optional)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
* teradata:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required if storing credentials directly strongDM) 
	* `secret_store_password_path` - (Required if using credentials stored in a secret store)
	* `secret_store_password_key` - (Optional)
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required if storing credentials directly strongDM) 
	* `secret_store_username_path` - (Required if using credentials stored in a secret store)
	* `secret_store_username_key` - (Optional)
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Resource resource:
* `id` - A unique identifier for the Resource resource.
* ssh:
	* `public_key` - 
## Import
Resource can be imported using the id, e.g.,

```
$ terraform import sdm_resource.example rs-12345678
```
