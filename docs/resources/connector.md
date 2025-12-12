---
page_title: "SDM: sdm_connector"
description: |-
  Provides settings for Connector.
layout: “sdm”
sidebar_current: “docs-sdm-resource-connector"
---
# Resource: sdm_connector

A Connector configures scanning for a given system.
This resource can be imported using the [import](https://www.terraform.io/docs/cli/commands/import.html) command.
## Argument Reference
The following arguments are supported by the Connector resource:
* aws:
	* `account_ids` - (Optional) AccountIds is the list of AWS Accounts to scan
	* `description` - (Optional) Description of the Connector.
	* `name` - (Required) Unique human-readable name of the Connector.
	* `role_name` - (Optional) RoleName is the Role we're assuming into for an account
	* `scan_period` - (Optional) ScanPeriod identifies which remote system this Connector discovers
	* `services` - (Optional) Services is a list of services this connector should scan.
* azure:
	* `client_id` - (Optional) ClientId is the ID of the Application / Service Account we're acting as
	* `description` - (Optional) Description of the Connector.
	* `name` - (Required) Unique human-readable name of the Connector.
	* `scan_period` - (Optional) ScanPeriod identifies which remote system this Connector discovers
	* `services` - (Optional) Services is a list of services this connector should scan.
	* `subscription_ids` - (Optional) SubscriptionIds are the targets of discovery.
	* `tenant_id` - (Optional) TenantId is the Azure Tenant we're discovering in
* gcp:
	* `description` - (Optional) Description of the Connector.
	* `name` - (Required) Unique human-readable name of the Connector.
	* `pool_id` - (Optional) PoolId is the GCP Workload Pool Identifier used to authenticate our JWT
	* `project_ids` - (Optional) ProjectIds is the list of GCP Projects the connector will scan
	* `project_number` - (Optional) ProjectNumber is the GCP Project the Workload Pool is defined in
	* `provider_id` - (Optional) ProviderId is the GCP Workload Provider Identifier used to authenticate our JWT
	* `scan_period` - (Optional) ScanPeriod identifies which remote system this Connector discovers
	* `services` - (Optional) Services is a list of services this connector should scan.
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by the Connector resource:
* `id` - A unique identifier for the Connector resource.
## Import
A Connector can be imported using the id, e.g.,

```
$ terraform import sdm_connector.example con-12345678
```
