---
page_title: "SDM: sdm_approval_workflow_approver"
description: |-
  Provides settings for ApprovalWorkflowApprover.
layout: “sdm”
sidebar_current: “docs-sdm-resource-approval-workflow-approver"
---
# Resource: sdm_approval_workflow_approver

ApprovalWorkflowApprover links an approval workflow approver to an ApprovalWorkflowStep
## Example Usage

```hcl
resource "sdm_approval_workflow_approver" "approval_workflow_approver_account_example" {
    approval_flow_id = "af-6799234"
    approval_step_id = "afs-2956266"
    account_id = "a-234605"
}

resource "sdm_approval_workflow_approver" "approval_workflow_approver_role_example" {
    approval_flow_id = "af-1935694"
    approval_step_id = "afs-9245942"
    role_id = "r-542982"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the ApprovalWorkflowApprover resource:
* `account_id` - (Optional) The approver account id.
* `approval_flow_id` - (Required) The approval flow id specified the approval workflow that this approver belongs to
* `approval_step_id` - (Required) The approval step id specified the approval flow step that this approver belongs to
* `role_id` - (Optional) The approver role id
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the ApprovalWorkflowApprover resource:
* `id` - A unique identifier for the ApprovalWorkflowApprover resource.
## Import
A ApprovalWorkflowApprover can be imported using the id, e.g.,

```
$ terraform import sdm_approval_workflow_approver.example afa-12345678
```
