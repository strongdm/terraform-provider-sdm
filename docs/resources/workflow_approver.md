---
page_title: "SDM: sdm_workflow_approver"
description: |-
  Provides settings for WorkflowApprover.
layout: “sdm”
sidebar_current: “docs-sdm-resource-workflow-approver"
---
# Resource: sdm_workflow_approver

WorkflowApprover is an account or a role with the ability to approve requests bound to a workflow.
## Example Usage

```hcl
resource "sdm_workflow_approver" "workflow_approver_account_example" {
    workflow_id = "aw-6799234"
    account_id = "a-234605"
}

resource "sdm_workflow_approver" "workflow_approver_role_example" {
    workflow_id = "aw-1935694"
    role_id = "r-542982"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the WorkflowApprover resource:
* `account_id` - (Optional) The approver account id.
* `role_id` - (Optional) The approver role id
* `workflow_id` - (Required) The workflow id.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the WorkflowApprover resource:
* `id` - A unique identifier for the WorkflowApprover resource.
## Import
A WorkflowApprover can be imported using the id, e.g.,

```
$ terraform import sdm_workflow_approver.example nt-12345678
```
