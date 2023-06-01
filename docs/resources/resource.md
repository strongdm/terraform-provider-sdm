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
## Providing Credentials

Credentials can be provided to resources in two forms:
- As raw text, which will not be returned to the terraform client on import or on loading state from StrongDM, but may be stored in the terraform state itself.
- As a path to a credential in a Secret Store, which will be returned on import. e.g. /path/to/secret?key=password&encoding=base64

All credentials must be either raw or Secret Store paths, depending on whether the resource has a Secret Store ID provided. In both cases, some credentials may be optional depending on the resource subtype.
## Argument Reference
The following arguments are supported by the Resource resource:
* aks:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `client_certificate` - (Optional) 
	* `client_key` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aks_basic_auth:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* aks_service_account:
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
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) 
* aks_service_account_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) 
* aks_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `client_certificate` - (Optional) 
	* `client_key` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks:
	* `access_key` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_instance_profile:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_instance_profile_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_eks_user_impersonation:
	* `access_key` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `cluster_name` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazon_es:
	* `access_key` - (Optional) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Optional) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* amazonmq_amqp_091:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required, either in plaintext, or as a secret store path) 
* athena:
	* `access_key` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `output` - (Required) 
	* `port_override` - (Optional) 
	* `region` - (Optional) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aurora_mysql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* aurora_postgres:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* aws:
	* `access_key` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_region` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
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
	* `role_arn` - (Required, either in plaintext, or as a secret store path) 
	* `role_external_id` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `session_expiry` - (Optional) 
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
* aws_console_static_key_pair:
	* `access_key` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `role_arn` - (Required, either in plaintext, or as a secret store path) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `session_expiry` - (Optional) 
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
* azure:
	* `app_id` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required, either in plaintext, or as a secret store path) 
* azure_certificate:
	* `app_id` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `client_certificate` - (Required, either in plaintext, or as a secret store path) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required, either in plaintext, or as a secret store path) 
* azure_mysql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* azure_postgres:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* big_query:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `private_key` - (Required, either in plaintext, or as a secret store path) 
	* `project` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
* cassandra:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required, either in plaintext, or as a secret store path) 
* citus:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* clustrix:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* cockroach:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* db_2_i:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required, either in plaintext, or as a secret store path) 
* db_2_luw:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* document_db_host:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
* document_db_replica_set:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port_override` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
* druid:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
* dynamo_db:
	* `access_key` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* elastic:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* elasticache_redis:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* gcp:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `keyfile` - (Required, either in plaintext, or as a secret store path) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `scopes` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* google_gke:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Required, either in plaintext, or as a secret store path) 
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* google_gke_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `name` - (Required) Unique human-readable name of the Resource.
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `service_account_key` - (Required, either in plaintext, or as a secret store path) 
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* greenplum:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* http_auth:
	* `auth_header` - (Required, either in plaintext, or as a secret store path) 
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
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Required) 
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `url` - (Required) 
	* `username` - (Optional) 
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
	* `client_certificate` - (Optional) 
	* `client_key` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `remote_identity_group_id` - (Optional) 
	* `remote_identity_healthcheck_username` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* kubernetes_basic_auth:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
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
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) 
* kubernetes_service_account_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `token` - (Required, either in plaintext, or as a secret store path) 
* kubernetes_user_impersonation:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `client_certificate` - (Optional) 
	* `client_key` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* maria:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* memcached:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* memsql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* mongo_host:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* mongo_legacy_host:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `replica_set` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* mongo_legacy_replicaset:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* mongo_replica_set:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `connect_to_replica` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `replica_set` - (Required) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* mongo_sharded_cluster:
	* `auth_database` - (Required) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* mtls_mysql:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `client_certificate` - (Required, either in plaintext, or as a secret store path) 
	* `client_key` - (Required, either in plaintext, or as a secret store path) 
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_name` - (Optional) 
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* mtls_postgres:
	* `bind_interface` - (Optional) Bind interface
	* `certificate_authority` - (Optional) 
	* `client_certificate` - (Required, either in plaintext, or as a secret store path) 
	* `client_key` - (Required, either in plaintext, or as a secret store path) 
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_name` - (Optional) 
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* mysql:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* neptune:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* neptune_iam:
	* `access_key` - (Required, either in plaintext, or as a secret store path) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `endpoint` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `region` - (Required) 
	* `role_arn` - (Optional) 
	* `role_external_id` - (Optional) 
	* `secret_access_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* oracle:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Required) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required, either in plaintext, or as a secret store path) 
* postgres:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* presto:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* rabbitmq_amqp_091:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Required, either in plaintext, or as a secret store path) 
* raw_tcp:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
* rdp:
	* `bind_interface` - (Optional) Bind interface
	* `downgrade_nla_connections` - (Optional) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* redis:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tls_required` - (Optional) 
	* `username` - (Optional) 
* redshift:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* single_store:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* snowflake:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port_override` - (Optional) 
	* `schema` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* snowsight:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `healthcheck_username` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port_override` - (Optional) 
	* `saml_metadata` - (Required, either in plaintext, or as a secret store path) 
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
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `schema` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* sql_server_azure_ad:
	* `bind_interface` - (Optional) Bind interface
	* `client_id` - (Required, either in plaintext, or as a secret store path) 
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `schema` - (Optional) 
	* `secret` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `tenant_id` - (Required, either in plaintext, or as a secret store path) 
* sql_server_kerberos_ad:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `keytab` - (Required, either in plaintext, or as a secret store path) 
	* `krb_config` - (Required, either in plaintext, or as a secret store path) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `override_database` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `realm` - (Required, either in plaintext, or as a secret store path) 
	* `schema` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `server_spn` - (Required, either in plaintext, or as a secret store path) 
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
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
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
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
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* ssh_customer_key:
	* `allow_deprecated_key_exchanges` - (Optional) 
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `port` - (Required) 
	* `port_forwarding` - (Optional) 
	* `port_override` - (Optional) 
	* `private_key` - (Required, either in plaintext, or as a secret store path) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* sybase:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* sybase_iq:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* teradata:
	* `bind_interface` - (Optional) Bind interface
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Required, either in plaintext, or as a secret store path) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Required, either in plaintext, or as a secret store path) 
* trino:
	* `bind_interface` - (Optional) Bind interface
	* `database` - (Required) 
	* `egress_filter` - (Optional) A filter applied to the routing logic to pin datasource to nodes.
	* `hostname` - (Required) 
	* `name` - (Required) Unique human-readable name of the Resource.
	* `password` - (Optional) 
	* `port` - (Optional) 
	* `port_override` - (Optional) 
	* `secret_store_id` - (Optional) ID of the secret store containing credentials for this resource, if any.
	* `subdomain` - (Optional) Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `username` - (Optional) 
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
