---
page_title: "Manage Groups"
---

# Manage Groups

Groups in StrongDM provide a way to organize users and manage permissions at scale. Instead of assigning users to roles individually, you can create groups, assign users to those groups, and then assign roles to the entire group. This makes access management more efficient and easier to maintain.

## Overview

Groups enable you to:

* **Organize Users**: Create logical groupings based on teams, departments, or functions
* **Simplify Role Management**: Assign roles to groups instead of individual users
* **Streamline Approval Workflows**: Use groups as approvers in approval workflows
* **Scale Access Management**: Easily manage permissions for large numbers of users

## Group Resources

The StrongDM Terraform provider includes the following resources for group management:

* `sdm_group` - Create and manage groups
* `sdm_account_group` - Manage relationships between accounts (users) and groups
* `sdm_group_role` - Manage relationships between groups and roles

## Create Groups

Create groups to organize your users. Groups are simply named containers that can hold multiple users.

### Example Usage

```hcl
# Create department-based groups
resource "sdm_group" "security_team" {
  name = "Security Team"
}

resource "sdm_group" "developers" {
  name = "Developers"
}

resource "sdm_group" "administrators" {
  name = "Administrators"
}

resource "sdm_group" "qa_team" {
  name = "QA Team"
}
```

## Add Users to Groups

Use the `sdm_account_group` resource to create relationships between users and groups. Users can be members of multiple groups.

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

resource "sdm_account" "carol_wilson" {
  user {
    first_name = "Carol"
    last_name  = "Wilson"
    email      = "carol.wilson@company.com"
  }
}

# Add Alice to Security Team
resource "sdm_account_group" "alice_security" {
  account_id = sdm_account.alice_smith.id
  group_id   = sdm_group.security_team.id
}

# Add Alice to Administrators (users can be in multiple groups)
resource "sdm_account_group" "alice_admin" {
  account_id = sdm_account.alice_smith.id
  group_id   = sdm_group.administrators.id
}

# Add Bob to Developers
resource "sdm_account_group" "bob_dev" {
  account_id = sdm_account.bob_jones.id
  group_id   = sdm_group.developers.id
}

# Add Carol to QA Team
resource "sdm_account_group" "carol_qa" {
  account_id = sdm_account.carol_wilson.id
  group_id   = sdm_group.qa_team.id
}
```

## Assign Roles to Groups

Use the `sdm_group_role` resource to assign roles to entire groups. All users in the group will inherit the permissions from these roles.

### Example Usage

```hcl
# Create roles
resource "sdm_role" "database_access" {
  name = "Database Access"
  access_rules = jsonencode([
    {
      "tags": { "type": "database" }
    }
  ])
}

resource "sdm_role" "production_access" {
  name = "Production Access"
  access_rules = jsonencode([
    {
      "tags": { "env": "production" }
    }
  ])
}

resource "sdm_role" "admin_access" {
  name = "Admin Access"
  access_rules = jsonencode([
    {
      "tags": { "*": "*" }
    }
  ])
}

# Assign roles to groups
resource "sdm_group_role" "developers_db_access" {
  group_id = sdm_group.developers.id
  role_id  = sdm_role.database_access.id
}

resource "sdm_group_role" "security_prod_access" {
  group_id = sdm_group.security_team.id
  role_id  = sdm_role.production_access.id
}

resource "sdm_group_role" "admins_full_access" {
  group_id = sdm_group.administrators.id
  role_id  = sdm_role.admin_access.id
}

# Groups can have multiple roles
resource "sdm_group_role" "admins_db_access" {
  group_id = sdm_group.administrators.id
  role_id  = sdm_role.database_access.id
}
```

## Groups in Approval Workflows

Groups can be used as approvers in manual approval workflows. This allows any member of the group to approve access requests.

### Example Usage

```hcl
# Create approval workflow using groups as approvers
resource "sdm_approval_workflow" "production_access_approval" {
  name          = "Production Access Approval"
  description   = "Requires approval from security team and administrators"
  approval_mode = "manual"

  # First step: Any member of Security Team can approve
  approval_step {
    quantifier = "any"
    skip_after = "2h0m0s"
    approvers {
      group_id = sdm_group.security_team.id
    }
  }

  # Second step: All specified groups must approve
  approval_step {
    quantifier = "all"
    skip_after = "24h0m0s"
    approvers {
      group_id = sdm_group.administrators.id
    }
    approvers {
      group_id = sdm_group.security_team.id
    }
  }
}

# Mixed approval workflow combining groups with other approver types
resource "sdm_approval_workflow" "mixed_approval" {
  name          = "Mixed Approval Workflow"
  description   = "Combines group and individual approvers"
  approval_mode = "manual"

  approval_step {
    quantifier = "any"
    skip_after = "4h0m0s"
    # Group approver
    approvers {
      group_id = sdm_group.security_team.id
    }
    # Individual approver
    approvers {
      account_id = sdm_account.alice_smith.id
    }
    # Manager reference
    approvers {
      reference = "manager-of-requester"
    }
  }
}
```

## Complete Example: Department-Based Access Management

This comprehensive example shows how to set up a complete department-based access management system using groups:

```hcl
# Department Groups
resource "sdm_group" "engineering" {
  name = "Engineering"
}

resource "sdm_group" "security" {
  name = "Security"
}

resource "sdm_group" "operations" {
  name = "Operations"
}

# Users
resource "sdm_account" "dev_lead" {
  user {
    first_name = "Dev"
    last_name  = "Lead"
    email      = "dev.lead@company.com"
  }
}

resource "sdm_account" "security_analyst" {
  user {
    first_name = "Security"
    last_name  = "Analyst"
    email      = "security.analyst@company.com"
  }
}

resource "sdm_account" "ops_manager" {
  user {
    first_name = "Ops"
    last_name  = "Manager"
    email      = "ops.manager@company.com"
  }
}

# Group Memberships
resource "sdm_account_group" "dev_lead_engineering" {
  account_id = sdm_account.dev_lead.id
  group_id   = sdm_group.engineering.id
}

resource "sdm_account_group" "security_analyst_security" {
  account_id = sdm_account.security_analyst.id
  group_id   = sdm_group.security.id
}

resource "sdm_account_group" "ops_manager_operations" {
  account_id = sdm_account.ops_manager.id
  group_id   = sdm_group.operations.id
}

# Cross-department membership
resource "sdm_account_group" "ops_manager_security" {
  account_id = sdm_account.ops_manager.id
  group_id   = sdm_group.security.id
}

# Roles
resource "sdm_role" "dev_resources" {
  name = "Development Resources"
  access_rules = jsonencode([
    {
      "tags": { "env": "development" }
    }
  ])
}

resource "sdm_role" "prod_resources" {
  name = "Production Resources"
  access_rules = jsonencode([
    {
      "tags": { "env": "production" }
    }
  ])
}

resource "sdm_role" "security_audit" {
  name = "Security Audit"
  access_rules = jsonencode([
    {
      "tags": { "audit": "true" }
    }
  ])
}

# Group-Role Assignments
resource "sdm_group_role" "engineering_dev" {
  group_id = sdm_group.engineering.id
  role_id  = sdm_role.dev_resources.id
}

resource "sdm_group_role" "operations_prod" {
  group_id = sdm_group.operations.id
  role_id  = sdm_role.prod_resources.id
}

resource "sdm_group_role" "security_audit_access" {
  group_id = sdm_group.security.id
  role_id  = sdm_role.security_audit.id
}

# Approval Workflow
resource "sdm_approval_workflow" "production_changes" {
  name          = "Production Changes"
  description   = "Approval required for production access"
  approval_mode = "manual"

  approval_step {
    quantifier = "all"
    skip_after = "8h0m0s"
    approvers {
      group_id = sdm_group.security.id
    }
    approvers {
      group_id = sdm_group.operations.id
    }
  }
}
```

## Best Practices

### Group Naming
* Use clear, descriptive names for groups
* Consider using consistent naming conventions (e.g., "Team - Department" or "Role - Function")
* Avoid overly generic names that might cause confusion

### Group Organization
* Align groups with your organizational structure
* Create groups based on job functions, departments, or access needs
* Consider creating both fine-grained and broad groups for flexibility

### Role Assignment
* Assign roles to groups rather than individual users when possible
* Use groups to implement the principle of least privilege
* Regularly review and audit group-role assignments

### Approval Workflows
* Use groups in approval workflows to distribute approval responsibilities
* Consider using different quantifiers ("any" vs "all") based on your security requirements
* Implement timeout periods to prevent indefinite pending approvals

## Data Sources

You can also use data sources to reference existing groups in your Terraform configurations:

```hcl
# Reference an existing group
data "sdm_group" "existing_team" {
  name = "Existing Team Name"
}

# Reference by ID
data "sdm_group" "by_id" {
  id = "g-1234567890abcdef"
}

# Use in other resources
resource "sdm_account_group" "add_to_existing" {
  account_id = sdm_account.new_user.id
  group_id   = data.sdm_group.existing_team.id
}
```

## Further Reading

* [StrongDM Groups Documentation](https://www.strongdm.com/docs)
* [Role-Based Access Control Best Practices](https://www.strongdm.com/docs)
* [Approval Workflows Guide](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/approval_workflows)
* [User Management Guide](https://registry.terraform.io/providers/strongdm/sdm/latest/docs/guides/user_management)