---
page_title: "SDM: sdm_node"
description: |-
  Query for existing Nodes instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-node"
---
# Data Source: sdm_node

Nodes make up the strongDM network, and allow your users to connect securely to your resources.
 There are two types of nodes:
 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
 1. **Gateways:** a relay that also listens for connections from strongDM clients
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
* `type` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://www.strongdm.com/docs/automation/getting-started/filters for more information.
* `bind_address` - (Optional) The hostname/port tuple which the gateway daemon will bind to.
 If not provided on create, set to "0.0.0.0:<listen_address_port>".
* `gateway_filter` - (Optional) GatewayFilter can be used to restrict the peering between relays and
 gateways.
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
		* `bind_address` - The hostname/port tuple which the gateway daemon will bind to.
 If not provided on create, set to "0.0.0.0:<listen_address_port>".
		* `gateway_filter` - GatewayFilter can be used to restrict the peering between relays and
 gateways.
		* `id` - Unique identifier of the Gateway.
		* `listen_address` - The public hostname/port tuple at which the gateway will be accessible to clients.
		* `name` - Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
		* `tags` - Tags is a map of key, value pairs.
	* relay:
		* `gateway_filter` - GatewayFilter can be used to restrict the peering between relays and
 gateways.
		* `id` - Unique identifier of the Relay.
		* `name` - Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
		* `tags` - Tags is a map of key, value pairs.
