---
page_title: "SDM: sdm_rdp_ca_pubkey"
description: |-
  Query for the RDP Certificate Authority Public Key. 
layout: “sdm”
sidebar_current: “docs-sdm-datasource-rdp-ca-pub-key"
---
# Data Source: sdm_rdp_ca_pubkey

The RDP CA Pubkey is a public key used for setting up a trusted CA on Active Directiory Domain Controllers.
## Example Usage

```hcl
data "sdm_rdp_ca_pubkey" "rdp_pubkey_query" {
}

output "rdpca" {
  value = data.sdm_rdp_ca_pubkey.rdp_pubkey_query.public_key
}
```
## Argument Reference
This datasource has no arguments.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a RDP CA Public Key data source:
* `id` - a generated id representing this request.
* `public_key` - the RDP Certificate Authority public key.