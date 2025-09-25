---
page_title: "SDM: sdm_workflow"
description: |-
  Provides settings for Workflow.
layout: “sdm”
sidebar_current: “docs-sdm-resource-workflow"
---
# Resource: sdm_workflow

Workflows are the collection of rules that define the resources to which access can be requested,
 the users that can request that access, and the mechanism for approving those requests which can either
 but automatic approval or a set of users authorized to approve the requests.
## Example Usage

```hcl
# Create approval workflows first
resource "sdm_approval_workflow" "auto_grant" {
  name = "Auto Grant Example"
  approval_mode = "automatic"
}

resource "sdm_approval_workflow" "manual_approval" {
  name = "Manual Approval Example"
  approval_mode = "manual"
  approval_step {
    quantifier = "any"
    skip_after = "2h0m0s"
    approvers {
      reference = "manager-of-requester"
    }
  }
}

# Create workflows that reference the approval workflows
resource "sdm_workflow" "auto_grant_workflow" {
    name = "auto grant workflow example"
    approval_flow_id = sdm_approval_workflow.auto_grant.id
    enabled = true
    access_rules = jsonencode([
    {
      "type" : "redis",
      "tags" : { "region" : "us-east" }
    }
  ])
}

resource "sdm_workflow" "manual_approval_workflow" {
    name = "manual approval workflow example"
    approval_flow_id = sdm_approval_workflow.manual_approval.id
    enabled = true
    access_rules = jsonencode([
    {
      "type" : "redis",
      "tags" : { "region" : "us-east" }
    }
  ])
}

```
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the Workflow resource:
* `access_request_fixed_duration` - (Optional) Fixed Duration of access requests bound to this workflow. If fixed duration is provided, max duration must be empty. If neither max nor fixed duration are provided, requests that bind to this workflow will use the organization-level settings.
* `access_request_max_duration` - (Optional) Maximum Duration of access requests bound to this workflow. If max duration is provided, fixed duration must be empty. If neither max nor fixed duration are provided, requests that bind to this workflow will use the organization-level settings.
* `access_rules` - (Optional) AccessRules is a list of access rules defining the resources this Workflow provides access to.
* `approval_flow_id` - (Optional) Optional approval flow ID identifies an approval flow that linked to the workflow
* `auto_grant` - (Optional) Optional auto grant setting to automatically approve requests or not, defaults to false.
* `description` - (Optional) Optional description of the Workflow.
* `enabled` - (Optional) Optional enabled state for workflow. This setting may be overridden by the system if the workflow doesn't meet the requirements to be enabled or if other conditions prevent enabling the workflow. The requirements to enable a workflow are that the workflow must be either set up for with auto grant enabled or have one or more WorkflowApprovers created for the workflow.
* `name` - (Required) Unique human-readable name of the Workflow.
* `weight` - (Optional) Optional weight for workflow to specify it's priority in matching a request.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Workflow resource:
* `id` - A unique identifier for the Workflow resource.
## Import
A Workflow can be imported using the id, e.g.,

```
$ terraform import sdm_workflow.example aw-12345678
```
