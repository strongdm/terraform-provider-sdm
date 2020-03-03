# Roles

Roles are the primary method in which users are granted access to a datasource. Each role is a logical grouping of access, this could be based on Job Title or

[Users and Roles](https://www.strongdm.com/docs/architecture/users-and-roles/)

## What is a role?

Roles in strongDM are the primary method of providing access to datasources and servers. A role is a collection of permissions that are then granted to the users that are assigned to that role.

### Regular Role

    resource "sdm_role" "IT_team" {
      name      = "IT users"
    }

### Composite Role

    resource "sdm_role" "DevOps" {
      name      = "DevOps"
      composite = true
    # It would be nice if roles were listed in the resource object.
    #	roles     = [sdm_role.Dev.id, sdm_role.Ops.id]
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

## Assigning Resources to Roles

## Assigning Users to a Roles

