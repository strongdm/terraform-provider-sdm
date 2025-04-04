---
page_title: "SDM: sdm_approval_workflow_step"
description: |-
  Provides settings for ApprovalWorkflowStep.
layout: “sdm”
sidebar_current: “docs-sdm-resource-approval-workflow-step"
---
# Resource: sdm_approval_workflow_step

ApprovalWorkflowStep links an approval workflow step to an ApprovalWorkflow
## Example Usage

```hcl
resource "sdm_approval_workflow_step" "approval_workflow_step_example" {
    approval_flow_id = "af-343865"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the ApprovalWorkflowStep resource:
* `approval_flow_id` - (Required) The approval flow id specified the approval workflow that this step belongs to
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the ApprovalWorkflowStep resource:
* `id` - A unique identifier for the ApprovalWorkflowStep resource.
## Import
A ApprovalWorkflowStep can be imported using the id, e.g.,

```
$ terraform import sdm_approval_workflow_step.example afs-12345678
```
