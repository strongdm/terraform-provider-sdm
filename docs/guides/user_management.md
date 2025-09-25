---
page_title: "Manage Users"
---

# Manage Users

In the StrongDM Admin UI, you can create two different types of accounts:

* [Users](https://www.strongdm.com/docs/admin-ui-guide/access/users/) are people within your organization that can log in to the StrongDM Desktop application and CLI. Each user may be added to one or more roles, which control what StrongDM resources users can access.
* [Service accounts](https://www.strongdm.com/docs/admin-ui-guide/access/service-accounts) provide programmatic access to StrongDM resources and authenticate machines. They can be used to automate processes such as ETL jobs or BI tools.

Both account types are managed using the `sdm_account` resource.

## Create a User

The following example [creates a user](https://www.strongdm.com/docs/admin-ui-guide/access/users/) in the StrongDM Admin UI using the Terraform provider. For argument and attribute references, see the [sdm_account](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/account) page.

### Example Usage

```hcl
# Create user in StrongDM
resource "sdm_account" "jane_doe" {
  user {
      first_name = "Jane"
      last_name = "Doe"
      email = "jane.doe@yourcompany.com"
  }
}
```

## Create a Service Account

The following example [creates a service account](https://www.strongdm.com/docs/admin-ui-guide/access/service-accounts/) in the StrongDM Admin UI using the Terraform provider. For argument and attribute references, see the [sdm_account](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/account) page.

### Example Usage

In order to use the service account, you must capture the access token. This token can be exported or used elsewhere in your Terraform configuration.

```hcl
# Create service account in StrongDM
resource "sdm_account" "dev_ci_pipeline" {
  service {
    # Display name
    name = "dev-pipeline-access"
  }
}

# Output service account token for use in other modules
# Accessible to other modules with module.<MODULE-NAME>.dev_ci_pipeline_token
output "dev_ci_pipeline_token" {
  value = sdm_account.dev_ci_pipeline.service[0].token
  sensitive = true
}
```

## Grant Access to StrongDM Resources

StrongDM uses role-based privileges to control access to infrastructure resources. You can grant users and service accounts access in two ways:

1. **Direct Role Assignment**: [Assign users directly to roles](https://www.strongdm.com/docs/admin-ui-guide/access/roles/) using account attachments
2. **Group-Based Assignment**: Add users to [groups](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/groups), then assign roles to the entire group

Group-based assignment is recommended for larger organizations as it simplifies access management and makes it easier to maintain permissions at scale.

For argument and attribute references, see [sdm_account](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/account), [sdm_role](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/role), [sdm_account_attachment](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/account_attachment), [sdm_group](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/group), [sdm_account_group](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/account_group), and [sdm_group_role](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/group_role).

### Example Usage

The following example illustrates how to create a user, add a role, and attach the user to that role using the Terraform provider. For more information about tagging resources, check our [Tags](https://www.strongdm.com/docs/automation/getting-started/tags/) documentation.

```hcl
# Create user in StrongDM
resource "sdm_account" "jane_doe" {
  user {
    first_name = "Jane"
    last_name  = "Doe"
    email      = "jane.doe@yourcompany.com"
  }
}

# Create role in StrongDM
resource "sdm_role" "developers" {
  name = "developers"
  # Give role access to all resources with `env=dev` tag
  access_rules = jsonencode([
    {
      "tags": { "env": "dev" }
    }
  ])
}

# Add user as member of role using account and role IDs
resource "sdm_account_attachment" "developers_jane_doe" {
  account_id  = sdm_account.jane_doe.id
  role_id     = sdm_role.developers.id
}
```

## Group-Based Access Management

For larger organizations, group-based access management provides a more scalable approach. Instead of assigning users to roles individually, you can organize users into groups and assign roles to the entire group.

### Example Usage

```hcl
# Create users
resource "sdm_account" "alice_smith" {
  user {
    first_name = "Alice"
    last_name  = "Smith"
    email      = "alice.smith@company.com"
  }
}

resource "sdm_account" "bob_jones" {
  user {
    first_name = "Bob"
    last_name  = "Jones"
    email      = "bob.jones@company.com"
  }
}

# Create a group for developers
resource "sdm_group" "developers" {
  name = "Developers"
}

# Add users to the group
resource "sdm_account_group" "alice_developers" {
  account_id = sdm_account.alice_smith.id
  group_id   = sdm_group.developers.id
}

resource "sdm_account_group" "bob_developers" {
  account_id = sdm_account.bob_jones.id
  group_id   = sdm_group.developers.id
}

# Create role for development resources
resource "sdm_role" "dev_resources" {
  name = "Development Resources"
  access_rules = jsonencode([
    {
      "tags": { "env": "development" }
    }
  ])
}

# Assign the role to the entire group (all group members inherit this access)
resource "sdm_group_role" "developers_access" {
  group_id = sdm_group.developers.id
  role_id  = sdm_role.dev_resources.id
}
```

This approach makes it easy to:
- Add new developers to the group and they automatically get the appropriate access
- Change permissions for all developers by modifying the group's role assignments
- Remove a developer's access by removing them from the group

For a comprehensive guide on group management, see the [Groups Guide](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/groups).
