---
page_title: "SDM: sdm_approval_workflow_step"
description: |-
  Query for existing ApprovalWorkflowSteps instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-approval-workflow-step"
---
# Data Source: sdm_approval_workflow_step

ApprovalWorkflowStep links an approval workflow step to an ApprovalWorkflow
## Example Usage

```hcl
data "sdm_approval_workflow_step" "approval_workflow_step_query" {
    approval_flow_id = "af-7935485"
}
```
## Argument Reference
The following arguments are supported by a ApprovalWorkflowSteps data source:
* `approval_flow_id` - (Optional) The approval flow id specified the approval workfflow that this step belongs to
* `id` - (Optional) Unique identifier of the ApprovalWorkflowStep.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a ApprovalWorkflowSteps data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_approval_workflow_step ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `approval_workflow_steps` - A list where each element has the following attributes:
	* `approval_flow_id` - The approval flow id specified the approval workfflow that this step belongs to
	* `id` - Unique identifier of the ApprovalWorkflowStep.
