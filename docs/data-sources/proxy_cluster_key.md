---
page_title: "SDM: sdm_proxy_cluster_key"
description: |-
  Query for existing ProxyClusterKeys instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-proxy-cluster-key"
---
# Data Source: sdm_proxy_cluster_key

Proxy Cluster Keys are authentication keys for all proxies within a cluster.
 The proxies within a cluster share the same key. One cluster can have
 multiple keys in order to facilitate key rotation.
## Example Usage

```hcl
data "sdm_proxy_cluster_key" "proxy_cluster_key_query" {
    proxy_cluster_id = "n-233332245"
}
```
## Argument Reference
The following arguments are supported by a ProxyClusterKeys data source:
* `id` - (Optional) Unique identifier of the Relay.
* `proxy_cluster_id` - (Optional) The ID of the proxy cluster which this key authenticates to.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a ProxyClusterKeys data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_proxy_cluster_key ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `proxy_cluster_keys` - A list where each element has the following attributes:
	* `id` - Unique identifier of the Relay.
	* `proxy_cluster_id` - The ID of the proxy cluster which this key authenticates to.
