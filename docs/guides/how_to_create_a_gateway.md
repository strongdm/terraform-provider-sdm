# How to Create a strongDM Gateway

The node resource creates either a gateway or a relay.

- Gateways act as an a secure bridge for your strongDM clients to access your databases.
- Relays act as an extension into private or secure subnets.

## Additional information

[Relays](https://www.strongdm.com/docs/architecture/relays/)

## Example Usage

    resource "sdm_node" "relay" {
      relay {
        name = "relay1"
      }
    }
    resource "sdm_node" "gateway" {
      gateway {
        name           = "gw1"
        listen_address = "sdmlocal:5000"
      }
    }
    # guessing on how outputs will look
    output "gateway_token" {
      value = sdm_node.gateway.token
      sensitive = true
    }

    output "relay_token" {
      value = sdm_node.relay.token
      sensitive = true
    }