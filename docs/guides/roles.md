# Managing Roles

Roles in strongDM are the primary method of providing access to datasources and servers. A role is a collection of permissions that are then granted to the users that are assigned to that role.

[Users and Roles](https://www.strongdm.com/docs/architecture/users-and-roles/)

## Creating Roles

### Regular Role

```hcl
resource "sdm_role" "IT_team" {
  name      = "IT users"
}
```

### Composite Role

```hcl
resource "sdm_role" "DevOps" {
  name      = "DevOps"
  composite = true
}
resource "sdm_role" "Dev" {
  name      = "Devs"
  composite = false
}
resource "sdm_role" "Ops" {
  name      = "Ops"
  composite = false
}
resource "sdm_role_attachment" "Dev" {
  composite_role_id = sdm_role.DevOps.id
  attached_role_id  = sdm_role.Dev.id
}
resource "sdm_role_attachment" "Ops" {
  composite_role_id = sdm_role.DevOps.id
  attached_role_id  = sdm_role.Ops.id
}
```

## Assigning Resources to Roles

```hcl
resource "sdm_role" "ops" {
  name      = "Network Operations"
}

data "sdm_resource" "prod_db" {
  name = "Prod DB"
  type = "postgres"
}

data "sdm_resource" "bastion" {
  name = "Bastion Host"
  type = "ssh"
}

resource "sdm_role_grant" "prod_db" {
  role_id = sdm_role.ops.id
  resource_id = data.sdm_resource.prod_db.ids[0]
}

resource "sdm_role_grant" "bastion" {
  role_id = sdm_role.ops.id
  resource_id = data.sdm_resource.bastion.ids[0]
}
```

## Assigning Users to a Roles

```hcl
resource "sdm_role" "ops" {
  name      = "Network Operations"
}

data "sdm_account" "dave" {
  email = "davefromops@example.com"
}

resource "sdm_role_attachment" "dave_ops" {
  role_id = sdm_role.ops.id
  account_id = data.sdm_accounts.dave.ids[0]
}
```
