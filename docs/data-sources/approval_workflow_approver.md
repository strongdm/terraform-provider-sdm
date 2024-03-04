---
page_title: "SDM: sdm_approval_workflow_approver"
description: |-
  Query for existing ApprovalWorkflowApprovers instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-approval-workflow-approver"
---
# Data Source: sdm_approval_workflow_approver

ApprovalWorkflowApprover links an approval workflow approver to an ApprovalWorkflowStep
## Example Usage

```hcl
data "sdm_approval_workflow_approver" "approval_workflow_approver_account_query" {
    approval_flow_id = "af-6799234"
    approval_step_id = "afs-2956266"
    account_id = "a-234605"
}

data "sdm_approval_workflow_approver" "approval_workflow_approver_role_query" {
    approval_flow_id = "af-1935694"
    approval_step_id = "afs-9245942"
    role_id = "r-542982"
}
```
## Argument Reference
The following arguments are supported by a ApprovalWorkflowApprovers data source:
* `account_id` - (Optional) The approver account id.
* `approval_flow_id` - (Optional) The approval flow id specified the approval workflow that this approver belongs to
* `approval_step_id` - (Optional) The approval step id specified the approval flow step that this approver belongs to
* `id` - (Optional) Unique identifier of the ApprovalWorkflowApprover.
* `role_id` - (Optional) The approver role id
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a ApprovalWorkflowApprovers data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_approval_workflow_approver ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `approval_workflow_approvers` - A list where each element has the following attributes:
	* `account_id` - The approver account id.
	* `approval_flow_id` - The approval flow id specified the approval workflow that this approver belongs to
	* `approval_step_id` - The approval step id specified the approval flow step that this approver belongs to
	* `id` - Unique identifier of the ApprovalWorkflowApprover.
	* `role_id` - The approver role id
