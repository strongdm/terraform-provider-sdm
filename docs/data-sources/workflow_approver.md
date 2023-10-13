---
page_title: "SDM: sdm_workflow_approver"
description: |-
  Query for existing WorkflowApprovers instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-workflow-approver"
---
# Data Source: sdm_workflow_approver

WorkflowApprover is an account with the ability to approve requests bound to a workflow.
## Example Usage

```hcl
data "sdm_workflow_approver" "workflow_approver_query" {
    workflow_id = "aw-541894"
    approver_id = "a-2496542"
}
```
## Argument Reference
The following arguments are supported by a WorkflowApprovers data source:
* `approver_id` - (Optional) The approver id.
* `id` - (Optional) Unique identifier of the WorkflowApprover.
* `workflow_id` - (Optional) The workflow id.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a WorkflowApprovers data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_workflow_approver ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `workflow_approvers` - A list where each element has the following attributes:
	* `approver_id` - The approver id.
	* `id` - Unique identifier of the WorkflowApprover.
	* `workflow_id` - The workflow id.
