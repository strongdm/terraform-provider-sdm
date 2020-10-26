---
page_title: "SDM: sdm_secret_store"
description: |-
  Query for existing SecretStores instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-secret-store"
---
# Data Source: sdm_secret_store

A SecretStore is a server where resource secrets (passwords, keys) are stored. 
 Coming soon support for HashiCorp Vault and AWS Secret Store. Contact support@strongdm.com to request access to the beta.
## Argument Reference
The following arguments are supported by a SecretStores data source:
* `id` - (Optional) option (grpc.gateway.protoc_gen_swagger.options.openapiv2_schema) = {
     example: { value: '{ "id": "r-7", "name": "happy-goat"}' }
 };
 Unique identifier of the SecretStore.
* `name` - (Optional) Unique human-readable name of the SecretStore.
* `server_address` - (Optional) 
* `kind` - (Optional) 
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a SecretStores data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_secret_store ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `secret_stores` - A list where each element has the following attributes:
	* `id` - option (grpc.gateway.protoc_gen_swagger.options.openapiv2_schema) = {
     example: { value: '{ "id": "r-7", "name": "happy-goat"}' }
 };
 Unique identifier of the SecretStore.
	* `name` - Unique human-readable name of the SecretStore.
	* `server_address` - 
	* `kind` - 
	* `tags` - Tags is a map of key, value pairs.
