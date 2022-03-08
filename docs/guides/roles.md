# Managing Roles

Roles in strongDM are the primary method of providing access to datasources and
servers. A role is a collection of permissions that are then granted to the
users that are assigned to that role.

[Users and Roles](https://www.strongdm.com/docs/architecture/users-and-roles/)

## Creating Roles

```hcl
resource "sdm_role" "IT_team" {
  name      = "IT users"
}
```

## Adding Access Rules to Roles

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

You can also use Resource IDs directly in Access Rules.

```hcl
resource "sdm_role" "engineering" {
  name = "engineering"
  access_rules = jsonencode([
    {
      ids = [sdm_resource.prod_db.id, sdm_resource.bastion.id]
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