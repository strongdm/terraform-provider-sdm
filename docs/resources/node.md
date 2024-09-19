---
page_title: "SDM: sdm_node"
description: |-
  Provides settings for Node.
layout: “sdm”
sidebar_current: “docs-sdm-resource-node"
---
# Resource: sdm_node

Nodes make up the strongDM network, and allow your users to connect securely to your resources.
 There are two types of nodes:
 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
 1. **Gateways:** a relay that also listens for connections from strongDM clients
## Example Usage

```hcl
resource "sdm_node" "gateway" {
    gateway {
        name = "test-gateway"
        listen_address = "165.23.40.1:21222"
        bind_address = "0.0.0.0:21222"
        tags = {
            region = "us-west"
            env = "dev"
        }    
    }
}

resource "sdm_node" "relay" {
    relay {
        name = "test-relay"
        tags = {
            region = "us-west"
            env = "dev"
        }    
    }
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the Node resource:
* gateway:
	* `bind_address` - (Optional) The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to "0.0.0.0:listen_address_port".
	* `gateway_filter` - (Optional) GatewayFilter can be used to restrict the peering between relays and gateways. Deprecated.
	* `listen_address` - (Required) The public hostname/port tuple at which the gateway will be accessible to clients.
	* `maintenance_window` - (Optional) Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.
	* `name` - (Optional) Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* proxy_cluster:
	* `address` - (Required) The public hostname/port tuple at which the proxy cluster will be accessible to clients.
	* `maintenance_window` - (Optional) Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.
	* `name` - (Optional) Unique human-readable name of the proxy cluster. Names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
	* `tags` - (Optional) Tags is a map of key, value pairs.
* relay:
	* `gateway_filter` - (Optional) GatewayFilter can be used to restrict the peering between relays and gateways. Deprecated.
	* `maintenance_window` - (Optional) Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.
	* `name` - (Optional) Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
	* `tags` - (Optional) Tags is a map of key, value pairs.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Node resource:
* `id` - A unique identifier for the Node resource.
* gateway:
	* `device` - Device is a read only device name uploaded by the gateway process when it comes online.
	* `location` - Location is a read only network location uploaded by the gateway process when it comes online.
	* `version` - Version is a read only sdm binary version uploaded by the gateway process when it comes online.
* relay:
	* `device` - Device is a read only device name uploaded by the gateway process when it comes online.
	* `location` - Location is a read only network location uploaded by the gateway process when it comes online.
	* `version` - Version is a read only sdm binary version uploaded by the gateway process when it comes online.
## Import
A Node can be imported using the id, e.g.,

```
$ terraform import sdm_node.example n-12345678
```
