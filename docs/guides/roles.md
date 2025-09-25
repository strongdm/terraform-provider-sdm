---
page_title: "Manage Roles"
---

# Manage Roles

[Roles](https://www.strongdm.com/docs/admin-ui-guide/access/roles/) serve as the primary method of providing human and machine access to StrongDM infrastructure resources, such as databases and servers.

A role consists of a collection of permissions that leverages [static or dynamic access rules](https://www.strongdm.com/docs/admin-ui-guide/access/rules/) to manage access to your StrongDM environment. When added to the role, members gain access to the specific resources defined by a particular role.

Roles can be assigned to users in two ways:
1. **Direct Assignment**: Assign roles directly to individual accounts using [sdm_account_attachment](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/account_attachment)
2. **Group-Based Assignment**: Assign roles to [groups](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/groups) using [sdm_group_role](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/group_role), and all group members inherit the role's permissions

Group-based assignment is recommended for larger organizations as it simplifies role management and makes it easier to maintain access permissions at scale.

## Create a Role

Use the following example to [create a role](https://www.strongdm.com/docs/admin-ui-guide/access/roles/) in StrongDM with the Terraform provider. See the [sdm_role](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/role) page for additional information about arguments and attributes.

### Example Usage

```hcl
# Create role in StrongDM
resource "sdm_role" "IT_team" {
  name = "IT users"
}
```

## Add Dynamic Access Rules to Roles

[Access rules](https://www.strongdm.com/docs/admin-ui-guide/access/rules/) serve as the building blocks for roles. Static access rules grant access to specific StrongDM resources, while dynamic access rules accomplish the same goal by using resource properties. As a best practice, we recommend utilizing the `type` and `tags` arguments to provision your StrongDM environment.

See the [sdm_role](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/role) page for additional information about arguments and attributes.

### Example Usage

```hcl
# Create role in StrongDM
resource "sdm_role" "engineering" {
  name = "engineering"

  # Add dynamic access rule to role
  access_rules = jsonencode([
    # Grant access to all dev environment resources in us-west
    {
      "tags": {
        "env": "dev",
        "region": "us-west"
      }
    },

    # Grant access to all Postgres resources
    {
      "type": "postgres"
    },

    # Grant access to all Redis resources in us-east
    {
      "type": "redis",
      "tags": {
        "region": "us-east"
      }
    }
  ])
}
```

## Assign Roles to Groups

Instead of assigning roles to individual users, you can assign roles to groups. All members of the group automatically inherit the role's permissions. This approach scales better for large organizations.

### Example Usage

```hcl
# Create a group for developers
resource "sdm_group" "developers" {
  name = "Developers"
}

# Create users
resource "sdm_account" "alice_developer" {
  user {
    first_name = "Alice"
    last_name  = "Developer"
    email      = "alice@company.com"
  }
}

resource "sdm_account" "bob_developer" {
  user {
    first_name = "Bob"
    last_name  = "Developer"
    email      = "bob@company.com"
  }
}

# Add users to the developers group
resource "sdm_account_group" "alice_to_developers" {
  account_id = sdm_account.alice_developer.id
  group_id   = sdm_group.developers.id
}

resource "sdm_account_group" "bob_to_developers" {
  account_id = sdm_account.bob_developer.id
  group_id   = sdm_group.developers.id
}

# Create a role for development resources
resource "sdm_role" "dev_access" {
  name = "Development Access"
  access_rules = jsonencode([
    {
      "tags": {
        "env": "development"
      }
    }
  ])
}

# Assign the role to the entire group
resource "sdm_group_role" "developers_dev_access" {
  group_id = sdm_group.developers.id
  role_id  = sdm_role.dev_access.id
}
```

This approach provides several advantages:
- **Simplified Management**: Add new developers to the group and they automatically get appropriate access
- **Consistent Permissions**: All group members have the same access level
- **Easy Updates**: Modify the group's role assignments to change access for all members
- **Audit Trail**: Clear visibility into which groups have which roles

For more information on group management, see the [Groups Guide](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/groups).

If you prefer, you can also write the access rules using raw JSON.

### Example Usage

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

## Add Static Access Rules to Roles

You can also use StrongDM resource IDs directly to create a static access rule for a specific role. See the [sdm_role](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/role) page for additional information about arguments and attributes.

### Example Usage

```hcl
# Retrieve existing MySQL datasource from StrongDM
data "sdm_resource" "mysql_database" {
  name = "terraform-sdm-mysql-admin"
}

# Create role in StrongDM
resource "sdm_role" "engineering" {
  name = "engineering"
  # Add static access rule to give direct access to existing MySQL datasource
  access_rules = jsonencode([
    {
      # Retrieve ID for existing MySQL datasource
      "ids": [data.sdm_resource.mysql_database.ids[0]]
    }
  ])
}
```

## Delete an Access Rule

The `sdm_role.accesss_rules` field is an optional, computed field. When deleting an access rule, you must explicitly set the list of access rules to an empty array.

### Example Usage

Let us say you specify access rules in this manner:

```hcl
# Create role in StrongDM
resource "sdm_role" "engineering" {
  name = "engineering"
  access_rules = jsonencode([
    {
      "type": "postgres"
    }
  ])
}
```

If you remove the argument using Terraform, the access rule for the role is preserved in StrongDM:

```hcl
resource "sdm_role" "engineering" {
  name = "engineering"
}
```

To completely remove the access rule from the role, explicitly set the `access_rules` argument to an empty array:

```hcl
# Re-create role in StrongDM and remove existing access rule
resource "sdm_role" "engineering" {
  name = "engineering"
  access_rules = jsonencode([])
}
```

## Migrate From Role Grants to Access Rules

->**Important Versioning Note** The information in this section only applies to **version 2.0 and up** of the Terraform provider. Prior to version 2.0, the provider includes only rudimentary beta support for access rules. Prior to version 1.0.27, the usage of access rules is not supported. We strongly recommend upgrading to version 2.0 when possible.

To improve flexibility when managing thousands of StrongDM resources, we recently deprecated role grants in favor of access rules. The usage of the latter allows you to take advantage of [dynamic access rules](https://www.strongdm.com/docs/admin-ui-guide/access/rules/#dynamic-access-rules) based on resource `type` and `tags`. As a best practice, we recommend leveraging [dynamic access rules](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/roles#add-dynamic-access-rules-to-roles) to manage your StrongDM roles.

In this section, we demonstrate how deprecated role grants work, while providing a static access rules example for backwards compatibility with role grants.

### Role Grants (Deprecated)

As illustrated in the following example, role grants previously allowed you to give a role access to specific StrongDM resources by using resource ID.

#### Example Usage

```hcl
# Create Redis db in StrongDM
resource "sdm_resource" "example_redis_db" {
  redis {
    name = "example-redis-db"
    hostname = "example.com"
    port_override = 4020
    tags = {
      region = "us-west"
      env = "dev"
    }
  }
}

# Create Postgres db in StrongDM
resource "sdm_resource" "example_postgres_db" {
  postgres {
    name = "example-postgres-db"
    hostname = "example.com"
    database = "my-postgres-db"
    username = "admin"
    password = "password"
    port = 5432
    tags = {
      region = "us-west"
      env = "dev"
    }
  }
}

# Create role in StrongDM
resource "sdm_role" "engineering" {
  name = "engineering"
}

# Use role grant to give engineering access to Redis db
resource "sdm_role_grant" "engineering_redis" {
  resource_id = sdm_resource.example_redis_db.id
  role_id = sdm_role.engineering.id
}

# Use role grant to give engineering access to Postgres db
resource "sdm_role_grant" "engineering_postgres" {
  resource_id = sdm_resource.example_postgres_db.id
  role_id = sdm_role.engineering.id
}
```

Instead, we recommend using [dynamic access rules](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/roles#add-dynamic-access-rules-to-roles) to achieve the same goal.

If you need to mimic the behavior of role grants, you can use resource IDs directly to add an access rule to a role. This example is similar to leveraging [static access rules](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/roles#add-static-access-rules-to-roles) to manage your StrongDM roles.

```hcl
# Create role in StrongDM
resource "sdm_role" "engineering" {
  name = "engineering"
  # Add static access rule to give direct access to existing datasources
  access_rules = jsonencode([
    {
      "ids": [sdm_resource.example_redis_db.id, sdm_resource.example_postgres_db.id]
    }
  ])
}
```
