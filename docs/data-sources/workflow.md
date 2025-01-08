---
page_title: "SDM: sdm_workflow"
description: |-
  Query for existing Workflows instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-workflow"
---
# Data Source: sdm_workflow

Workflows are the collection of rules that define the resources to which access can be requested,
 the users that can request that access, and the mechanism for approving those requests which can either
 but automatic approval or a set of users authorized to approve the requests.
## Example Usage

```hcl
data "sdm_workflow" "workflow_query" {
    name = "workflow example"
    approval_mode = "automatic"
}
```
## Argument Reference
The following arguments are supported by a Workflows data source:
* `approval_flow_id` - (Optional) Optional approval flow ID identifies an approval flow that linked to the workflow
* `auto_grant` - (Optional) Optional auto grant setting to automatically approve requests or not, defaults to false.
* `description` - (Optional) Optional description of the Workflow.
* `enabled` - (Optional) Optional enabled state for workflow. This setting may be overridden by the system if the workflow doesn't meet the requirements to be enabled or if other conditions prevent enabling the workflow. The requirements to enable a workflow are that the workflow must be either set up for with auto grant enabled or have one or more WorkflowApprovers created for the workflow.
* `id` - (Optional) Unique identifier of the Workflow.
* `name` - (Optional) Unique human-readable name of the Workflow.
* `weight` - (Optional) Optional weight for workflow to specify it's priority in matching a request.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a Workflows data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_workflow ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `workflows` - A list where each element has the following attributes:
	* `access_request_fixed_duration` - Fixed Duration of access requests bound to this workflow. If fixed duration is provided, max duration must be empty. If neither max nor fixed duration are provided, requests that bind to this workflow will use the organization-level settings.
	* `access_request_max_duration` - Maximum Duration of access requests bound to this workflow. If max duration is provided, fixed duration must be empty. If neither max nor fixed duration are provided, requests that bind to this workflow will use the organization-level settings.
	* `access_rules` - AccessRules is a list of access rules defining the resources this Workflow provides access to.
	* `approval_flow_id` - Optional approval flow ID identifies an approval flow that linked to the workflow
	* `auto_grant` - Optional auto grant setting to automatically approve requests or not, defaults to false.
	* `description` - Optional description of the Workflow.
	* `enabled` - Optional enabled state for workflow. This setting may be overridden by the system if the workflow doesn't meet the requirements to be enabled or if other conditions prevent enabling the workflow. The requirements to enable a workflow are that the workflow must be either set up for with auto grant enabled or have one or more WorkflowApprovers created for the workflow.
	* `id` - Unique identifier of the Workflow.
	* `name` - Unique human-readable name of the Workflow.
	* `weight` - Optional weight for workflow to specify it's priority in matching a request.
