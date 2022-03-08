# Migrating from Role Grants to Access Rules

To increase flexibility when managing thousands of Resources, Role Grants have
been deprecated in favor of Access Rules, which allow you to grant access based
on Resource Tags and Type.

The following examples demonstrate the deprecated Role Grants, Dynamic Access
Rules with Tags and Resource Types, and Static Access Rules for backwards
compatibility with Role Grants.

## Important Versioning Caveat
This guide only applies to version 2.0 and up of the Terraform Provider.
Prior to 2.0, the provider includes only rudimentary beta support for access
rules. Prior to 1.0.27, it does not support access rules at all. We strongly
recommend upgrading to 2.0 when it is available.

Furthermore, before you can use access rules, your organization must undergo the
"Access Overhaul" migration to enable the new UI and a myriad of other features.
Contact support@strongdm.com to learn more.

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

  access_rules = jsonencode([
    # grant access to all dev environment resources in us-west
    {
      tags = {
        env = "dev"
        region = "us-west"
      }
    },

    # grant access to all postgres resources
    {
      type = "postgres"
    },

    # grant access to all redis resources in us-east
    {
      type = "redis"
      tags = {
        region = "us-east"
      }
    },
  ])
}
```

## Static Access Rules

If it is _necessary_ to grant access to specific Resources in the same way as
RoleGrants did, you can use Resource IDs directly in Access Rules.

```hcl
resource "sdm_role" "engineering" {
  name = "engineering"
  access_rules = jsonencode([
    {
      ids = [sdm_resource.redis-test.id, sdm_resource.postgres-test.id]
    }
  ])
}
```

## Note on Deleting Access Rules

The `sdm_role.access_rules` field is an Optional Computed field. For example,
let's say you specify access rules like this:

```hcl
resource "sdm_role" "test" {
  name = "test"
  access_rules = jsonencode([
    {
      type = "postgres"
    }
  ])
}
```

If you then remove the field, the access rules will be preserved in strongDM:

```hcl
resource "sdm_role" "test" {
  name = "test"
}
```

To delete the access rules from the role, you must explicitly set the field:

```hcl
resource "sdm_role" "test" {
  name = "test"
  access_rules = jsonencode([])
}
```

## Raw JSON

If you want, you can also write access rules in raw JSON.

```hcl
resource "sdm_role" "engineering" {
  name = "engineering"
  access_rules = <<-EOF
  [
    {
      "type": "redis"
    }
  ]
  EOF
}
```
