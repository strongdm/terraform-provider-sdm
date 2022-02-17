# Migrating to Access Rules

To increase flexibility when managing thousands of Resources, Role Grants have
been deprecated in favor of Access Rules, which allow you to grant access based
on Resource Tags and Type. Role Grants will be removed from the Terraform
provider by March 31, 2022.

The following examples demonstrate the deprecated Role Grants, Dynamic Access
Rules with Tags and Resource Types, and Static Access Rules for backwards
compatibility with Role Grants.


## Important Versioning Caveat
This guide only applies to version 1.0.40 and up of the Terraform Provider.
Prior to 1.0.40, the provider includes only rudimentary beta support for access
rules. Prior to 1.0.27, it does not support access rules at all. We strongly
recommend upgrading to 1.0.40 when it is available.

## Role Grants (deprecated)

Previously, you would grant a role access to specific resources by ID via role
grants:

```hcl
resource "sdm_resource" "redis-test" {
  redis {
    name = "redis-test"
    hostname = "example.com"
    port_override = 4020
    tags = {
      region = "us-west"
      env = "dev"
    }
  }
}

resource "sdm_resource" "postgres-test" {
  postgres {
    name = "postgres-test"
    hostname = "example.com"
    database = "my-db"
    username = "admin"
    password = "hunter2"
    port = 5432
    tags = {
      region = "us-west"
      env = "dev"
    }
  }
}

resource "sdm_role" "engineering" {
  name = "engineering"
}

resource "sdm_role_grant" "engineering-redis" {
  resource_id = sdm_resource.redis-test.id
  role_id = sdm_role.engineering.id
}

resource "sdm_role_grant" "engineering-postgres" {
  resource_id = sdm_resource.postgres-test.id
  role_id = sdm_role.engineering.id
}
```

## Dynamic Access Rules

When using Access Rules the best practice is to grant Resources access based on
Type and Tags.

```hcl
resource "sdm_role" "engineering" {
  name = "engineering"

  # grant access to all dev environment resources in us-west
  access_rule {
    tags = {
      env = "dev"
      region = "us-west"
    }
  }

  # grant access to all postgres resources
  access_rule {
    type = "postgres"
  }

  # grant access to all redis resources in us-east
  access_rule {
    type = "redis"
    tags = {
      region = "us-east"
    }
  }
}
```

## Static Access Rules

If it is _necessary_ to grant access to specific Resources in the same way as
RoleGrants did, you can use Resource IDs directly in Access Rules.

```hcl
resource "sdm_role" "engineering" {
  name = "engineering"
  access_rule {
    ids = [sdm_resource.redis-test.id, sdm_resource.postgres-test.id]
  }
}
```
