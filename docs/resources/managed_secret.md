---
page_title: "SDM: sdm_managed_secret"
description: |-
  Provides settings for ManagedSecret.
layout: “sdm”
sidebar_current: “docs-sdm-resource-managed-secret"
---
# Resource: sdm_managed_secret

ManagedSecret contains details about managed secret
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the ManagedSecret resource:
* `name` - (Required) Unique human-readable name of the Managed Secret.
* `policy` - (Required) Password and rotation policy for the secret
* `secret_engine_id` - (Required) An ID of a Secret Engine linked with the Managed Secret.
* `tags` - (Optional) Tags is a map of key, value pairs.
* `value` - (Optional) Sensitive value of the secret.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the ManagedSecret resource:
* `id` - A unique identifier for the ManagedSecret resource.
## Import
A ManagedSecret can be imported using the id, e.g.,

```
$ terraform import sdm_managed_secret.example ms-12345678
```
