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
* `hostname` - (Optional) The host to dial to initiate a connection from the egress node to this resource.
* `id` - (Optional) Unique identifier of the Resource.
* `name` - (Optional) Unique human-readable name of the Resource.
* `port` - (Optional) The port to dial to initiate a connection from the egress node to this resource.
* `tags` - (Optional) Tags is a map of key, value pairs.
* `username` - (Optional) The username to authenticate with.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Resources data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_resource ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `resources` - A single element list containing a map, where each key lists one of the following objects:
	* aks:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `client_certificate` - The certificate to authenticate TLS connections with.
		* `client_key` - The key to authenticate TLS connections with.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* aks_basic_auth:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* aks_service_account:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `token` - The API token to authenticate with.
	* aks_service_account_user_impersonation:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `token` - The API token to authenticate with.
	* aks_user_impersonation:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `client_certificate` - The certificate to authenticate TLS connections with.
		* `client_key` - The key to authenticate TLS connections with.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* amazon_eks:
		* `access_key` - The Access Key ID to use to authenticate.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `cluster_name` - The name of the cluster to connect to.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to e.g. us-east-1.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to use to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* amazon_eks_instance_profile:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `cluster_name` - The name of the cluster to connect to.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to e.g. us-east-1.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* amazon_eks_instance_profile_user_impersonation:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `cluster_name` - The name of the cluster to connect to.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to e.g. us-east-1.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* amazon_eks_user_impersonation:
		* `access_key` - The Access Key ID to use to authenticate.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `cluster_name` - The name of the cluster to connect to.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to e.g. us-east-1.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to use to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* amazon_es:
		* `access_key` - The Access Key ID to use to authenticate.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial e.g. search-?.region.es.amazonaws.com"
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to e.g. us-east-1.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to use to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* amazonmq_amqp_091:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* athena:
		* `access_key` - The Access Key ID to use to authenticate.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `output` - The AWS S3 output location.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to e.g. us-east-1.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to use to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* aurora_mysql:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks. Does not affect client requests
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* aurora_postgres:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* aws:
		* `access_key` - The Access Key ID to use to authenticate.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_region` - The AWS region healthcheck requests should attempt to connect to.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to use to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* aws_console:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `enable_env_variables` - If true, prefer environment variables to authenticate connection even if EC2 roles are configured.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `session_expiry` - The length of time in seconds AWS console sessions will live before needing to reauthenticate.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* aws_console_static_key_pair:
		* `access_key` - The Access Key ID to authenticate with.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to authenticate with.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `session_expiry` - The length of time in seconds AWS console sessions will live before needing to reauthenticate.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* azure:
		* `app_id` - The application ID to authenticate with.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tenant_id` - The tenant ID to authenticate to.
	* azure_certificate:
		* `app_id` - The application ID to authenticate with.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `client_certificate` - The service Principal certificate file, both private and public key included.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tenant_id` - The tenant ID to authenticate to.
	* azure_mysql:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks. Does not affect client requests.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* azure_postgres:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* azure_postgres_managed_identity:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* big_query:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `private_key` - The JSON Private key to authenticate with.
		* `project` - The project to connect to.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* cassandra:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* citus:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* clustrix:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks. Does not affect client requests.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* cockroach:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* db_2_i:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* db_2_luw:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* document_db_host:
		* `auth_database` - The authentication database to use.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* document_db_replica_set:
		* `auth_database` - The authentication database to use.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `connect_to_replica` - Set to connect to a replica instead of the primary node.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - Hostname must contain the hostname/port pairs of all instances in the replica set separated by commas.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port_override` - The local port used by clients to connect to this resource.
		* `replica_set` - The name of the mongo replicaset.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* druid:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* dynamo_db:
		* `access_key` - The Access Key ID to use to authenticate.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial e.g. dynamodb.region.amazonaws.com
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The region to authenticate requests against e.g. us-east-1
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to use to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* elastic:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* elasticache_redis:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* gcp:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `id` - Unique identifier of the Resource.
		* `keyfile` - The service account keyfile to authenticate with.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `scopes` - Space separated scopes that this login should assume into when authenticating.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* google_gke:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `service_account_key` - The service account key to authenticate with.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* google_gke_user_impersonation:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The endpoint to dial.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `service_account_key` - The service account key to authenticate with.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* greenplum:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* http_auth:
		* `auth_header` - The content to set as the authorization header.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `default_path` - Automatically redirect to this path upon connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `headers_blacklist` - Header names (e.g. Authorization), to omit from logs.
		* `healthcheck_path` - This path will be used to check the health of your site.
		* `host_override` - The host header will be overwritten with this field if provided.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `url` - The base address of your website without the path.
	* http_basic_auth:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `default_path` - Automatically redirect to this path upon connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `headers_blacklist` - Header names (e.g. Authorization), to omit from logs.
		* `healthcheck_path` - This path will be used to check the health of your site.
		* `host_override` - The host header will be overwritten with this field if provided.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `url` - The base address of your website without the path.
		* `username` - The username to authenticate with.
	* http_no_auth:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `default_path` - Automatically redirect to this path upon connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `headers_blacklist` - Header names (e.g. Authorization), to omit from logs.
		* `healthcheck_path` - This path will be used to check the health of your site.
		* `host_override` - The host header will be overwritten with this field if provided.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `url` - The base address of your website without the path.
	* kubernetes:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `client_certificate` - The certificate to authenticate TLS connections with.
		* `client_key` - The key to authenticate TLS connections with.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* kubernetes_basic_auth:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* kubernetes_service_account:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `token` - The API token to authenticate with.
	* kubernetes_service_account_user_impersonation:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `token` - The API token to authenticate with.
	* kubernetes_user_impersonation:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `client_certificate` - The certificate to authenticate TLS connections with.
		* `client_key` - The key to authenticate TLS connections with.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_namespace` - The path used to check the health of your connection.  Defaults to `default`.  This field is required, and is only marked as optional for backwards compatibility.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* maria:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks. Does not affect client requests.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* memcached:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* memsql:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks. Does not affect client requests.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* mongo_host:
		* `auth_database` - The authentication database to use.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* mongo_legacy_host:
		* `auth_database` - The authentication database to use.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `replica_set` - The name of the mongo replicaset.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* mongo_legacy_replicaset:
		* `auth_database` - The authentication database to use.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `connect_to_replica` - Set to connect to a replica instead of the primary node.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `replica_set` - The name of the mongo replicaset.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* mongo_replica_set:
		* `auth_database` - The authentication database to use.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `connect_to_replica` - Set to connect to a replica instead of the primary node.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `replica_set` - The name of the mongo replicaset.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* mongo_sharded_cluster:
		* `auth_database` - The authentication database to use.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* mtls_mysql:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `client_certificate` - The certificate to authenticate TLS connections with.
		* `client_key` - The key to authenticate TLS connections with.
		* `database` - The database for healthchecks. Does not affect client requests.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `server_name` - Server name for TLS verification (unverified by StrongDM if empty)
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* mtls_postgres:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `certificate_authority` - The CA to authenticate TLS connections with.
		* `client_certificate` - The certificate to authenticate TLS connections with.
		* `client_key` - The key to authenticate TLS connections with.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `server_name` - Server name for TLS verification (unverified by StrongDM if empty)
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* mysql:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks. Does not affect client requests.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* neptune:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The neptune endpoint to connect to as in endpoint.region.neptune.amazonaws.com
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* neptune_iam:
		* `access_key` - The Access Key ID to use to authenticate.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `endpoint` - The neptune endpoint to connect to as in endpoint.region.neptune.amazonaws.com
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `region` - The AWS region to connect to.
		* `role_arn` - The role to assume after logging in.
		* `role_external_id` - The external ID to associate with assume role requests. Does nothing if a role ARN is not provided.
		* `secret_access_key` - The Secret Access Key to use to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* oracle:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* postgres:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* presto:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* rabbitmq_amqp_091:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* raw_tcp:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* rdp:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `downgrade_nla_connections` - When set, network level authentication will not be used. May resolve unexpected authentication errors to older servers. When set, healthchecks cannot detect if a provided username / password pair is correct.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* redis:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tls_required` - If set, TLS must be used to connect to this resource.
		* `username` - The username to authenticate with.
	* redshift:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* single_store:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks. Does not affect client requests.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `require_native_auth` - Whether native auth (mysql_native_password) is used for all connections (for backwards compatibility)
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `use_azure_single_server_usernames` - If true, appends the hostname to the username when hitting a database.azure.com address
		* `username` - The username to authenticate with.
	* snowflake:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port_override` - The local port used by clients to connect to this resource.
		* `schema` - The schema to provide on authentication.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* snowsight:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `healthcheck_username` - The StrongDM user email to use for healthchecks.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `saml_metadata` - The Metadata for your snowflake IDP integration
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
	* sql_server:
		* `allow_deprecated_encryption` - Whether to allow deprecated encryption protocols to be used for this resource. For example, TLS 1.0.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks, and used for clients if Override Default Database is true.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `schema` - The Schema to use to direct initial requests.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* sql_server_azure_ad:
		* `allow_deprecated_encryption` - Whether to allow deprecated encryption protocols to be used for this resource. For example, TLS 1.0.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `client_id` - The Azure AD application (client) ID with which to authenticate.
		* `database` - The database for healthchecks, and used for clients if Override Default Database is true.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `schema` - The Schema to use to direct initial requests.
		* `secret` - The Azure AD client secret (application password) with which to authenticate.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `tenant_id` - The Azure AD directory (tenant) ID with which to authenticate.
	* sql_server_kerberos_ad:
		* `allow_deprecated_encryption` - Whether to allow deprecated encryption protocols to be used for this resource. For example, TLS 1.0.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The database for healthchecks, and used for clients if Override Default Database is true.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `keytab` - The keytab file in base64 format containing an entry with the principal name (username@realm) and key version number with which to authenticate.
		* `krb_config` - The Kerberos 5 configuration file (krb5.conf) specifying the Active Directory server (KDC) for the configured realm.
		* `name` - Unique human-readable name of the Resource.
		* `override_database` - If set, the database configured cannot be changed by users. This setting is not recommended for most use cases, as some clients will insist their database has changed when it has not, leading to user confusion.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `realm` - The Active Directory domain (realm) to which the configured username belongs.
		* `schema` - The Schema to use to direct initial requests.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `server_spn` - The Service Principal Name of the Microsoft SQL Server instance in Active Directory.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* ssh:
		* `allow_deprecated_key_exchanges` - Whether deprecated, insecure key exchanges are allowed for use to connect to the target ssh server.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `key_type` - The key type to use e.g. rsa-2048 or ed25519
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_forwarding` - Whether port forwarding is allowed through this server.
		* `port_override` - The local port used by clients to connect to this resource.
		* `public_key` - The public key to append to a server's authorized keys. This will be generated after resource creation.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* ssh_cert:
		* `allow_deprecated_key_exchanges` - Whether deprecated, insecure key exchanges are allowed for use to connect to the target ssh server.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `key_type` - The key type to use e.g. rsa-2048 or ed25519
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_forwarding` - Whether port forwarding is allowed through this server.
		* `port_override` - The local port used by clients to connect to this resource.
		* `remote_identity_group_id` - The ID of the remote identity group to use for remote identity connections.
		* `remote_identity_healthcheck_username` - The username to use for healthchecks, when clients otherwise connect with their own remote identity username.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* ssh_customer_key:
		* `allow_deprecated_key_exchanges` - Whether deprecated, insecure key exchanges are allowed for use to connect to the target ssh server.
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_forwarding` - Whether port forwarding is allowed through this server.
		* `port_override` - The local port used by clients to connect to this resource.
		* `private_key` - The private key used to authenticate with the server.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* sybase:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* sybase_iq:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* teradata:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
	* trino:
		* `bind_interface` - The bind interface is the IP address to which the port override of a resource is bound (for example, 127.0.0.1). It is automatically generated if not provided.
		* `database` - The initial database to connect to. This setting does not by itself prevent switching to another database after connecting.
		* `egress_filter` - A filter applied to the routing logic to pin datasource to nodes.
		* `hostname` - The host to dial to initiate a connection from the egress node to this resource.
		* `id` - Unique identifier of the Resource.
		* `name` - Unique human-readable name of the Resource.
		* `password` - The password to authenticate with.
		* `port` - The port to dial to initiate a connection from the egress node to this resource.
		* `port_override` - The local port used by clients to connect to this resource.
		* `secret_store_id` - ID of the secret store containing credentials for this resource, if any.
		* `subdomain` - Subdomain is the local DNS address.  (e.g. app-prod1 turns into app-prod1.your-org-name.sdm.network)
		* `tags` - Tags is a map of key, value pairs.
		* `username` - The username to authenticate with.
