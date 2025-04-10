---
page_title: "SDM: sdm_approval_workflow"
description: |-
  Query for existing ApprovalWorkflows instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-approval-workflow"
---
# Data Source: sdm_approval_workflow

ApprovalWorkflows are the mechanism by which requests for access can be viewed by authorized
 approvers and be approved or denied.
## Example Usage

```hcl
data "sdm_approval_workflow" "manual_approval_workflow_query" {
    name = "approval workflow manual"
    approval_mode = "manual"
}

data "sdm_approval_workflow" "auto_grant_approval_workflow_query" {
    name = "approval workflow auto"
    approval_mode = "automatic"
}
```
## Argument Reference
The following arguments are supported by a ApprovalWorkflows data source:
* `approval_mode` - (Optional) Approval mode of the ApprovalWorkflow
* `approval_step` - (Optional) The approval steps of this approval workflow
* `description` - (Optional) Optional description of the ApprovalWorkflow.
* `id` - (Optional) Unique identifier of the ApprovalWorkflow.
* `name` - (Optional) Unique human-readable name of the ApprovalWorkflow.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a ApprovalWorkflows data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_approval_workflow ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `approval_workflows` - A list where each element has the following attributes:
	* `approval_mode` - Approval mode of the ApprovalWorkflow
	* `approval_step` - The approval steps of this approval workflow
	* `description` - Optional description of the ApprovalWorkflow.
	* `id` - Unique identifier of the ApprovalWorkflow.
	* `name` - Unique human-readable name of the ApprovalWorkflow.
