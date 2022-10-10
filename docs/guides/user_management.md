---
page_title: "Manage Users"
---

# Manage Users

In the StrongDM Admin UI, you can create two different types of accounts:

* [Users](https://www.strongdm.com/docs/admin-ui-guide/access/users/) are people within your organization that can log in to the StrongDM Desktop application and CLI. Each user may be added to one or more roles, which control what StrongDM resources users can access.
* [Service accounts](https://www.strongdm.com/docs/admin-ui-guide/access/service-accounts) provide programmatic access to StrongDM resources and authenticate machines. They can be used to automate processes such as ETL jobs or BI tools.

Both account types are managed using the `sdm_account` resource.

## Create a User

The following example [creates a user](https://www.strongdm.com/docs/admin-ui-guide/access/users/) in the StrongDM Admin UI using the Terraform provider. For argument and attribute references, see the [sdm_account](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/resource) page.

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

The following example [creates a service account](https://www.strongdm.com/docs/admin-ui-guide/access/service-accounts/) in the StrongDM Admin UI using the Terraform provider. For argument and attribute references, see the [sdm_account](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/resource) page.

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

StrongDM uses role-based privileges to control access to infrastructure resources. You can grant users and service accounts access by [assigning each to a role](https://www.strongdm.com/docs/admin-ui-guide/access/roles/). After the role assignment, each user entity can access the StrongDM resources attached to the specific role.

For argument and attribute references, see [sdm_account](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/resource), [sdm_role](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/role), and [sdm_account_attachment](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/resources/account_attachment).

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
