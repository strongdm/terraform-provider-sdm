---
page_title: "SDM: sdm_workflow_role"
description: |-
  Provides settings for WorkflowRole.
layout: “sdm”
sidebar_current: “docs-sdm-resource-workflow-role"
---
# Resource: sdm_workflow_role

WorkflowRole links a role to a workflow. The linked roles indicate which roles a user must be a part of
 to request access to a resource via the workflow.
## Example Usage

```hcl
resource "sdm_workflow_role" "workflow_role_example" {
    workflow_id = sdm_workflow.manual_approval_workflow.id
    role_id = sdm_role.developers.id
}

resource "sdm_workflow_role" "workflow_role_with_ids" {
    workflow_id = "w-1234567890abcdef"
    role_id = "r-1234567890abcdef"
}
```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the WorkflowRole resource:
* `role_id` - (Required) The role id.
* `workflow_id` - (Required) The workflow id.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the WorkflowRole resource:
* `id` - A unique identifier for the WorkflowRole resource.
## Import
A WorkflowRole can be imported using the id, e.g.,

```
$ terraform import sdm_workflow_role.example wr-12345678
```
