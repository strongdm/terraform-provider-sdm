---
page_title: "SDM: sdm_workflow_role"
description: |-
  Query for existing WorkflowRoles instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-workflow-role"
---
# Data Source: sdm_workflow_role

WorkflowRole links a role to a workflow. The linked roles indicate which roles a user must be a part of
 to request access to a resource via the workflow.
## Example Usage

```hcl
data "sdm_workflow_role" "workflow_role_query" {
    workflow_id = "aw-7935485"
    role_id = "r-9862923"
}
```
## Argument Reference
The following arguments are supported by a WorkflowRoles data source:
* `id` - (Optional) Unique identifier of the WorkflowRole.
* `role_id` - (Optional) The role id.
* `workflow_id` - (Optional) The workflow id.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a WorkflowRoles data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_workflow_role ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `workflow_roles` - A list where each element has the following attributes:
	* `id` - Unique identifier of the WorkflowRole.
	* `role_id` - The role id.
	* `workflow_id` - The workflow id.
