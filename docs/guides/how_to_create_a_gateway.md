# How to Create a strongDM Gateway

strongDM creates a secure software defined network that allows you to access your databases and servers. The node resource is used to create the members of that network. Nodes can be either gateways or relays.

- **Gateways** are the entry points into network. They listen for connection from the strongDM client, and provide access to databases and servers.
- **Relays** are used to extend the strongDM network into segmented subnets. They provide access to databases and servers but do not listen for incoming connections.

See the [strongDM Network Architecture](https://www.strongdm.com/docs/architecture/relays/) documentation for more details.


## Example Usage

```hcl
# Create a gateway
resource "sdm_node" "my_gateway" {
  gateway {
    name           = "gw1"
    listen_address = "sdmlocal:5000"
  }
}

# Create a relay
resource "sdm_node" "my_relay" {
  relay {
    name = "relay1"
  }
}

# Output the tokens for use in other modules
output "gateway_token" {
  value = sdm_node.my_gateway..gateway.token
  sensitive = true
}

output "relay_token" {
  value = sdm_node.my_relay.relay.token
  sensitive = true
}
```