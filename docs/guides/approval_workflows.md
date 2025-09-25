---
page_title: "Manage Approval Workflows"
---

# Manage Approval Workflows

Approval workflows in StrongDM provide a way to require approval before users can access sensitive resources. Instead of granting immediate access, you can define approval processes that require specific individuals, groups, or roles to approve access requests.

## Overview

Approval workflows enable you to:

* **Control Sensitive Access**: Require approval for production or critical resources
* **Implement Compliance Requirements**: Meet organizational policies for privileged access
* **Distribute Approval Responsibilities**: Use multiple approvers and approval steps
* **Automate Approval Processes**: Define rules for automatic or manual approval

## Approval Workflow Types

StrongDM supports two types of approval workflows:

* **Manual Approval Workflows**: Require explicit approval from designated approvers
* **Automatic Approval Workflows**: Automatically grant access without human intervention

## Manual Approval Workflows

Manual approval workflows require one or more approval steps. Each step can have multiple approvers with different requirements.

### Basic Manual Approval Workflow

```hcl
resource "sdm_approval_workflow" "basic_manual" {
  name          = "Basic Manual Approval"
  description   = "Requires manager approval"
  approval_mode = "manual"
  
  approval_step {
    quantifier = "any"
    skip_after = "2h0m0s"
    approvers {
      reference = "manager-of-requester"
    }
  }
}
```

### Multi-Step Approval Workflow

```hcl
resource "sdm_approval_workflow" "multi_step" {
  name          = "Multi-Step Approval"
  description   = "Requires multiple levels of approval"
  approval_mode = "manual"
  
  # First step: Any security team member
  approval_step {
    quantifier = "any"
    skip_after = "4h0m0s"
    approvers {
      account_id = sdm_account.security_lead.id
    }
    approvers {
      role_id = sdm_role.security_team.id
    }
  }
  
  # Second step: Both admin approval AND manager approval
  approval_step {
    quantifier = "all"
    skip_after = "24h0m0s"
    approvers {
      role_id = sdm_role.administrators.id
    }
    approvers {
      reference = "manager-of-requester"
    }
  }
}
```

## Group-Based Approvers

Groups provide an efficient way to manage approvers in approval workflows. Instead of specifying individual users, you can designate entire groups as approvers.

### Example Usage

```hcl
# Create approver groups
resource "sdm_group" "security_team" {
  name = "Security Team"
}

resource "sdm_group" "administrators" {
  name = "Administrators"
}

resource "sdm_group" "ops_team" {
  name = "Operations Team"
}

# Add users to groups
resource "sdm_account" "security_lead" {
  user {
    first_name = "Security"
    last_name  = "Lead"
    email      = "security.lead@company.com"
  }
}

resource "sdm_account_group" "security_lead_to_security" {
  account_id = sdm_account.security_lead.id
  group_id   = sdm_group.security_team.id
}

# Approval workflow using groups
resource "sdm_approval_workflow" "group_based_approval" {
  name          = "Group-Based Approval Workflow"
  description   = "Uses groups as approvers"
  approval_mode = "manual"
  
  # Step 1: Any member of Security Team can approve
  approval_step {
    quantifier = "any"
    skip_after = "2h0m0s"
    approvers {
      group_id = sdm_group.security_team.id
    }
  }
  
  # Step 2: All specified groups must approve
  approval_step {
    quantifier = "all"
    skip_after = "24h0m0s"
    approvers {
      group_id = sdm_group.administrators.id
    }
    approvers {
      group_id = sdm_group.ops_team.id
    }
  }
}
```

## Approver Types

Approval workflows support multiple types of approvers:

### Individual Account Approvers

```hcl
approvers {
  account_id = sdm_account.specific_user.id
}
```

### Role-Based Approvers

```hcl
approvers {
  role_id = sdm_role.administrators.id
}
```

### Group-Based Approvers

```hcl
approvers {
  group_id = sdm_group.security_team.id
}
```

### Reference Approvers

```hcl
# Manager of the person making the request
approvers {
  reference = "manager-of-requester"
}

# Manager of the manager of the person making the request
approvers {
  reference = "manager-of-manager-of-requester"
}
```

## Mixed Approver Types

You can combine different approver types in a single approval step:

```hcl
resource "sdm_approval_workflow" "mixed_approvers" {
  name          = "Mixed Approver Types"
  description   = "Combines different types of approvers"
  approval_mode = "manual"
  
  approval_step {
    quantifier = "any"  # Any ONE of these can approve
    skip_after = "4h0m0s"
    
    # Specific individual
    approvers {
      account_id = sdm_account.cto.id
    }
    
    # Any member of security group
    approvers {
      group_id = sdm_group.security_team.id
    }
    
    # Anyone with admin role
    approvers {
      role_id = sdm_role.administrators.id
    }
    
    # Requester's manager
    approvers {
      reference = "manager-of-requester"
    }
  }
}
```

## Quantifiers

Each approval step has a quantifier that determines how many approvers must approve:

* **"any"**: Only one approver from the list needs to approve
* **"all"**: All approvers in the list must approve

### Example: Any vs All

```hcl
resource "sdm_approval_workflow" "quantifier_example" {
  name          = "Quantifier Examples"
  approval_mode = "manual"
  
  # Step 1: ANY security team member OR admin (just one needed)
  approval_step {
    quantifier = "any"
    skip_after = "2h0m0s"
    approvers {
      group_id = sdm_group.security_team.id
    }
    approvers {
      group_id = sdm_group.administrators.id
    }
  }
  
  # Step 2: Security team AND operations team (both groups needed)
  approval_step {
    quantifier = "all"
    skip_after = "8h0m0s"
    approvers {
      group_id = sdm_group.security_team.id
    }
    approvers {
      group_id = sdm_group.operations.id
    }
  }
}
```

## Automatic Approval Workflows

Automatic approval workflows grant access immediately without human intervention:

```hcl
resource "sdm_approval_workflow" "auto_grant" {
  name          = "Automatic Approval"
  description   = "Automatically grants access"
  approval_mode = "automatic"
  
  # No approval steps needed for automatic workflows
}
```

## Timeout Configuration

Use the `skip_after` parameter to define how long to wait before automatically approving a step:

```hcl
approval_step {
  quantifier = "any"
  skip_after = "24h0m0s"  # Auto-approve after 24 hours
  approvers {
    group_id = sdm_group.security_team.id
  }
}

# Common timeout values:
# "1h0m0s"   - 1 hour
# "8h0m0s"   - 8 hours (business day)
# "24h0m0s"  - 24 hours
# "168h0m0s" - 1 week
# "0s"       - No timeout (manual approval required)
```

## Complete Example: Department-Based Approval

This example shows a comprehensive approval workflow for a large organization:

```hcl
# Department groups
resource "sdm_group" "security" {
  name = "Security Department"
}

resource "sdm_group" "engineering" {
  name = "Engineering Department"  
}

resource "sdm_group" "operations" {
  name = "Operations Department"
}

resource "sdm_group" "executives" {
  name = "Executive Team"
}

# Production access approval workflow
resource "sdm_approval_workflow" "production_access" {
  name          = "Production Access Approval"
  description   = "Multi-level approval for production resources"
  approval_mode = "manual"
  
  # Level 1: Initial security review (any security team member)
  approval_step {
    quantifier = "any"
    skip_after = "2h0m0s"
    approvers {
      group_id = sdm_group.security.id
    }
  }
  
  # Level 2: Technical review (engineering AND operations)
  approval_step {
    quantifier = "all"
    skip_after = "8h0m0s"
    approvers {
      group_id = sdm_group.engineering.id
    }
    approvers {
      group_id = sdm_group.operations.id
    }
  }
  
  # Level 3: Executive approval for sensitive access
  approval_step {
    quantifier = "any"
    skip_after = "24h0m0s"
    approvers {
      group_id = sdm_group.executives.id
    }
    # Fallback to manager if executives unavailable
    approvers {
      reference = "manager-of-manager-of-requester"
    }
  }
}

# Emergency access workflow (faster approval)
resource "sdm_approval_workflow" "emergency_access" {
  name          = "Emergency Access"
  description   = "Expedited approval for emergency situations"
  approval_mode = "manual"
  
  approval_step {
    quantifier = "any"
    skip_after = "30m0s"  # Quick timeout for emergencies
    approvers {
      group_id = sdm_group.security.id
    }
    approvers {
      group_id = sdm_group.executives.id
    }
    approvers {
      reference = "manager-of-requester"
    }
  }
}

# Development resources (auto-approve)
resource "sdm_approval_workflow" "dev_access" {
  name          = "Development Access"
  description   = "Automatic approval for development resources"
  approval_mode = "automatic"
}
```

## Best Practices

### Workflow Design
* Design workflows that match your organizational approval processes
* Use multiple steps for highly sensitive resources
* Implement reasonable timeouts to prevent indefinite waiting
* Consider emergency access procedures

### Group Organization
* Create groups that align with your approval authority structure
* Use descriptive group names that clearly indicate approval responsibilities
* Consider creating both broad and specific approval groups

### Quantifier Selection
* Use "any" when you want flexible approval from multiple options
* Use "all" when you need consensus from multiple parties
* Combine both quantifiers across different steps for complex approval flows

### Timeout Management
* Set shorter timeouts for urgent access scenarios
* Use longer timeouts for major production changes
* Consider business hours and time zones when setting timeouts
* Use "0s" timeout when manual approval is always required

## Data Sources

Reference existing approval workflows in your Terraform configurations:

```hcl
# Reference an existing approval workflow
data "sdm_approval_workflow" "existing_workflow" {
  name = "Production Approval Process"
}

# Reference by ID
data "sdm_approval_workflow" "by_id" {
  id = "aw-1234567890abcdef"
}
```
