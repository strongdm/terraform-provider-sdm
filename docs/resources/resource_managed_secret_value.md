---
page_title: "SDM: sdm_managed_secret_value"
description: |-
  Provides encrypted value for ManagedSecret.
layout: “sdm”
sidebar_current: “docs-sdm-resource-managed-secret-value"
---
# Resource: sdm_managed_secret_value

sdm_managed_secret_value is used to provide encrypted value data for ManagedSecret resource.

## Argument Reference
The following arguments are supported by the ManagedSecret resource:
* `public_key` - (Required) A PEM-encoded Secret Engine's public key. Secret engine's `public_key` attribute should be used
* `value` - (Required) Map of attributes that are going to be stored in the secret engine

## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the ManagedSecret resource:
* `encrypted` - An encrypted value that should be assigned to `value` attribute of ManagedSecret

## Example

```

data "sdm_secret_engine" "engine" {
        name = "ldap"
}

resource "sdm_managed_secret_value" "secretVal" {
  value = {
    user_dn = "cn=John A. Zoidberg,ou=people,dc=planetexpress,dc=com"
  }
  public_key = data.sdm_secret_engine.engine.secret_engines[0].active_directory[0].public_key
}

resource "sdm_managed_secret" "mySecret" {
        name = "zoidberg_ad"
        secret_engine_id = data.sdm_secret_engine.localEngine.ids[0]
        value = sdm_managed_secret_value.secretVal.encrypted
 }
```
