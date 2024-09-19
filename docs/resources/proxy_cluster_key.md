---
page_title: "SDM: sdm_proxy_cluster_key"
description: |-
  Provides settings for ProxyClusterKey.
layout: “sdm”
sidebar_current: “docs-sdm-resource-proxy-cluster-key"
---
# Resource: sdm_proxy_cluster_key

Proxy Cluster Keys are authentication keys for all proxies within a cluster.
 The proxies within a cluster share the same key. One cluster can have
 multiple keys in order to facilitate key rotation.
## Example Usage

```hcl
resource "sdm_proxy_cluster_key" "test_proxy_cluster_key" {
    proxy_cluster_id = "n-12345123"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the ProxyClusterKey resource:
* `proxy_cluster_id` - (Required) The ID of the proxy cluster which this key authenticates to.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the ProxyClusterKey resource:
* `id` - A unique identifier for the ProxyClusterKey resource.
## Import
A ProxyClusterKey can be imported using the id, e.g.,

```
$ terraform import sdm_proxy_cluster_key.example gk-12345678
```
