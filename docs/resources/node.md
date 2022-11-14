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
	* `device` - (Optional) Device is a read only device name uploaded by the gateway process when  it comes online.
	* `gateway_filter` - (Optional) GatewayFilter can be used to restrict the peering between relays and gateways.
	* `listen_address` - (Required) The public hostname/port tuple at which the gateway will be accessible to clients.
	* `location` - (Optional) Location is a read only network location uploaded by the gateway process when it comes online.
	* `name` - (Optional) Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `version` - (Optional) Version is a read only sdm binary version uploaded by the gateway process when it comes online.
* relay:
	* `device` - (Optional) Device is a read only device name uploaded by the gateway process when  it comes online.
	* `gateway_filter` - (Optional) GatewayFilter can be used to restrict the peering between relays and gateways.
	* `location` - (Optional) Location is a read only network location uploaded by the gateway process when it comes online.
	* `name` - (Optional) Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.
	* `tags` - (Optional) Tags is a map of key, value pairs.
	* `version` - (Optional) Version is a read only sdm binary version uploaded by the gateway process when it comes online.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Node resource:
* `id` - A unique identifier for the Node resource.
## Import
Node can be imported using the id, e.g.,

```
$ terraform import sdm_node.example n-12345678
```
