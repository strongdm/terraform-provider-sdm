---
page_title: "Add a StrongDM Gateway or Relay"
---

# Add a StrongDM Gateway or Relay

StrongDM creates a secure, software-defined network that controls access to your infrastructure resources, such as databases, servers, clusters, clouds, and websites. We use node resources to create the members of this network, and nodes come in two forms &#8212; gateways or relays.

* [Gateways](https://www.strongdm.com/docs/architecture/deployment/gateways/) serve as the entry points into your network. They listen for connections from the StrongDM client and provide access to network resources, such as databases and servers.
* [Relays](https://www.strongdm.com/docs/architecture/deployment/relays/) are also gateways, but they have a different purpose. They can be used to extend the StrongDM network into segmented subnets that do not allow ingress. Relays provide access to network resources behind your firewall, but they do not listen for incoming connections.

For AWS and Azure-specific onboarding examples, check our [StrongDM with Terraform and AWS](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-aws/) and [StrongDM with Terraform and Azure](https://www.strongdm.com/docs/automation/configuration-management-tools/terraform/quick-start-azure/) quick start guides.

## Create a Gateway

Use the following example to [create a gateway](https://www.strongdm.com/docs/admin-ui-guide/network/gateways/) in StrongDM with the Terraform provider. See the [sdm_node](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/node) page for additional information about arguments and attributes.

### Example Usage

```hcl
# Create StrongDM gateway
resource "sdm_node" "example_gateway" {
  gateway {
    # Name to display in Admin UI
    name = "example-gateway"
    # Public hostname/port to make gateway accessible to clients
    listen_address = "gateway.example.com:5555"
    # Optional hostname/port for gateway daemon to bind to
    bind_address = "0.0.0.0:5555"
  }
}
```

## Create a Relay

Use the following example to [create a relay](https://www.strongdm.com/docs/admin-ui-guide/network/relays/) in StrongDM with the Terraform provider. See the [sdm_node](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/node) page for additional information about arguments and attributes.

### Example Usage

```hcl
# Create StrongDM relay
resource "sdm_node" "example_relay" {
  relay {
    # Name to display in Admin UI
    name = "example-relay"
  }
}
```

## Output Gateway or Relay Tokens

When creating a gateway or relay, StrongDM [generates a token](https://www.strongdm.com/docs/admin-ui-guide/network/gateways/#add-a-gateway) that is needed to complete the gateway or relay installation.

Return this token in the standard output stream after the Terraform module runs with `terraform output`. Other modules can access the value of the token variable using `module.<MODULE-NAME>.<OUTPUT-NAME>`. For more information, see [Accessing Child Module Outputs](https://www.terraform.io/language/values/outputs#accessing-child-module-outputs) in the Terraform documentation.

### Example Usage

```hcl
# Create StrongDM gateway
resource "sdm_node" "example_gateway" {
  gateway {
    name = "example-gateway"
    listen_address = "gateway.example.com:5555"
  }
}

# Output gateway token for use in other modules
# Accessible to other modules with module.<MODULE-NAME>.gateway_token
output "gateway_token" {
  value = sdm_node.my_gateway.gateway[0].token
  sensitive = true
}

# Create StrongDM relay
resource "sdm_node" "example_relay" {
  relay {
    name = "example-relay"
  }
}

# Output relay token for use in other modules
# Accessible to other modules with module.<MODULE-NAME>.relay_token
output "relay_token" {
  value = sdm_node.example_relay.relay[0].token
  sensitive = true
}
```
