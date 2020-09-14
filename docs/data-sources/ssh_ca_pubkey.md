---
page_title: "SDM: sdm_ssh_ca_pubkey"
description: |-
  Query for the SSH Certificate Authority Public Key. 
layout: “sdm”
sidebar_current: “docs-sdm-datasource-ssh-ca-pub-key"
---
# Data Source: sdm_ssh_ca_pubkey

The SSH CA Pubkey is a public key used for setting up SSH resources.
## Example Usage

```hcl
data "sdm_ssh_ca_pubkey" "ssh_pubkey_query" {
}
```
## Argument Reference
This datasource has no arguments.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a SSH CA Public Key data source:
* `id` - a generated id representing this request.
* `public_key` - the SSH Certificate Authority public key.