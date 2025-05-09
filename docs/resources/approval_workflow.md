---
page_title: "SDM: sdm_approval_workflow"
description: |-
  Provides settings for ApprovalWorkflow.
layout: “sdm”
sidebar_current: “docs-sdm-resource-approval-workflow"
---
# Resource: sdm_approval_workflow

ApprovalWorkflows are the mechanism by which requests for access can be viewed by authorized
 approvers and be approved or denied.
## Example Usage

```hcl
resource "sdm_approval_workflow" "manual_approval_workflow" {
    name = "manual approval workflow example"
    approval_mode = "manual"
    approval_step {
        quantifier = "any"
        skip_after = "1h0m0s"
        approvers {
            account_id = "a-1234abc"
        }
        approvers {
            reference = "manager-of-requester"
        }
    }
    approval_step {
        quantifier = "all"
        skip_after = "0s"
        approvers {
            role_id = "r-1234abc"
        }
        approvers {
            account_id = "a-5678def"
        }
        approvers {
            reference = "manager-of-manager-of-requester"
        }
    }
}


resource "sdm_approval_workflow" "auto_grant_approval_workflow" {
    name = "auto approval workflow example"
    approval_mode = "automatic"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the ApprovalWorkflow resource:
* `approval_mode` - (Required) Approval mode of the ApprovalWorkflow
* `approval_step` - (Optional) The approval steps of this approval workflow
* `description` - (Optional) Optional description of the ApprovalWorkflow.
* `name` - (Required) Unique human-readable name of the ApprovalWorkflow.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the ApprovalWorkflow resource:
* `id` - A unique identifier for the ApprovalWorkflow resource.
## Import
A ApprovalWorkflow can be imported using the id, e.g.,

```
$ terraform import sdm_approval_workflow.example af-12345678
```
