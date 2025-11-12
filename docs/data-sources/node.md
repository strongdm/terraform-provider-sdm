---
page_title: "SDM: sdm_node"
description: |-
  Query for existing Nodes instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-node"
---
# Data Source: sdm_node

Nodes make up the StrongDM network, and allow your users to connect securely to your resources.
 There are three types of nodes:
 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
 2. **Gateway:** a relay that also listens for connections from StrongDM clients
 3. **Proxy Cluster:** a cluster of workers that together mediate access from clients to resources
## Example Usage

```hcl
data "sdm_node" "gateway_query" {
    type = "gateway"
    tags = {
        region = "us-west"
        env = "dev"
    }    
}
```
## Argument Reference
The following arguments are supported by a Nodes data source:
* `type` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://docs.strongdm.com/references/cli/filters/) for more information.
* `bind_address` - (Optional) The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:listen_address_port".
* `id` - (Optional) Unique identifier of the Relay.
* `listen_address` - (Optional) The public hostname/port tuple at which the gateway will be accessible to clients.
* `name` - (Optional) Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Nodes data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_node ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `nodes` - A single element list containing a map, where each key lists one of the following objects:
	* gateway:
		* `bind_address` - The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:listen_address_port".
		* `device` - Device is a read only device name uploaded by the gateway process when it comes online.
		* `gateway_filter` - GatewayFilter can be used to restrict the peering between relays and gateways. Deprecated.
		* `id` - Unique identifier of the Gateway.
		* `listen_address` - The public hostname/port tuple at which the gateway will be accessible to clients.
		* `location` - Location is a read only network location uploaded by the gateway process when it comes online.
		* `maintenance_window` - Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.
		* `name` - Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
		* `tags` - Tags is a map of key, value pairs.
		* `version` - Version is a read only sdm binary version uploaded by the gateway process when it comes online.
	* proxy_cluster:
		* `address` - The public hostname/port tuple at which the proxy cluster will be accessible to clients.
		* `id` - Unique identifier of the Proxy Cluster.
		* `maintenance_window` - Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.
		* `name` - Unique human-readable name of the proxy cluster. Names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
		* `tags` - Tags is a map of key, value pairs.
	* relay:
		* `device` - Device is a read only device name uploaded by the gateway process when it comes online.
		* `gateway_filter` - GatewayFilter can be used to restrict the peering between relays and gateways. Deprecated.
		* `id` - Unique identifier of the Relay.
		* `location` - Location is a read only network location uploaded by the gateway process when it comes online.
		* `maintenance_window` - Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.
		* `name` - Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
		* `tags` - Tags is a map of key, value pairs.
		* `version` - Version is a read only sdm binary version uploaded by the gateway process when it comes online.
