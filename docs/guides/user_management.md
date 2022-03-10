# Managing Users

strongDM has two different types of user accounts

1. **Users:** humans who are authenticated through username and password or SSO.
2. **Service Accounts:** machines that are authenticated using a service token.

Both types of accounts are managed using the `sdm_account` resource.

## Creating users

### How to create a User

```hcl
resource "sdm_account" "jane_doe" {
  user {
    first_name = "Jane"
    last_name  = "Doe"
    email      = "Jane.doe@yourcompany.com"
  }
}
```

### How to create a Service Account

```hcl
resource "sdm_account" "dev_ci_pipeline" {
  service {
    name = "dev pipeline access"
  }
}

output "dev_ci_pipeline_token" {
  value = sdm_account.dev_ci_pipeline.service[0].token
}
```

*In order to use the service user you must make sure to capture the access
token. It can be exported or used elsewhere in your Terraform configuration.*

## Granting access

Users and Serivce Accounts can be granted access to resources by assigning them
to a role. Users then have access to all of the resources attached to the role.

### How to assign a user to a role

```hcl

resource "sdm_account" "jane_doe" {
  user {
    first_name = "Jane"
    last_name  = "Doe"
    email      = "Jane.doe@yourcompany.com"
  }
}

resource "sdm_role" "developers" {
  name = "developers"
  access_rules = jsonencode([
    {
      "tags": { "env": "dev" }
    }
  ])
}

resource "sdm_account_attachment" "developers_jane_doe" {
  account_id  = sdm_account.jane_doe.id
  role_id     = sdm_role.developers.id
}
```
