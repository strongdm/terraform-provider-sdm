---
page_title: "SDM: sdm_connector"
description: |-
  Query for existing DiscoveryConnectors instances.
layout: “sdm”
sidebar_current: “docs-sdm-datasource-connector"
---
# Data Source: sdm_connector

A Connector configures scanning for a given system.
## Argument Reference
The following arguments are supported by a DiscoveryConnectors data source:
* `cloud` - (Optional) a filter to select all items of a certain subtype. See the [filter documentation](https://docs.strongdm.com/references/cli/filters/) for more information.
* `account_ids` - (Optional) AccountIds is the list of AWS Accounts to scan
* `client_id` - (Optional) ClientId is the ID of the Application / Service Account we're acting as
* `description` - (Optional) Description of the Connector.
* `id` - (Optional) Unique identifier of the Connector.
* `name` - (Optional) Unique human-readable name of the Connector.
* `project_ids` - (Optional) ProjectIds is the list of GCP Projects the connector will scan
* `role_name` - (Optional) RoleName is the Role we're assuming into for an account
* `scan_period` - (Optional) ScanPeriod identifies which remote system this Connector discovers
* `services` - (Optional) Services is a list of services this connector should scan.
* `subscription_ids` - (Optional) SubscriptionIds are the targets of discovery.
* `tenant_id` - (Optional) TenantId is the Azure Tenant we're discovering in
* `workload_pool_id` - (Optional) WorkloadPoolId is the GCP Workload Pool Identifier used to authenticate our JWT
* `workload_project_id` - (Optional) WorkloadProjectId is the GCP Project ID where the Workload Pool is defined
* `workload_project_number` - (Optional) WorkloadProjectNumber is the GCP Project Number where the Workload Pool is defined
* `workload_provider_id` - (Optional) WorkloadProviderId is the GCP Workload Provider Identifier used to authenticate our JWT
## Attribute Reference
In addition to provided arguments above, the following attributes are returned by a DiscoveryConnectors data source:
* `id` - a generated id representing this request, unrelated to input id and sdm_connector ids.
* `ids` - a list of strings of ids of data sources that match the given arguments.
* `discovery_connectors` - A single element list containing a map, where each key lists one of the following objects:
	* aws:
		* `account_ids` - AccountIds is the list of AWS Accounts to scan
		* `description` - Description of the Connector.
		* `id` - Unique identifier of the Connector.
		* `name` - Unique human-readable name of the Connector.
		* `role_name` - RoleName is the Role we're assuming into for an account
		* `scan_period` - ScanPeriod identifies which remote system this Connector discovers
		* `services` - Services is a list of services this connector should scan.
	* azure:
		* `client_id` - ClientId is the ID of the Application / Service Account we're acting as
		* `description` - Description of the Connector.
		* `id` - Unique identifier of the Connector.
		* `name` - Unique human-readable name of the Connector.
		* `scan_period` - ScanPeriod identifies which remote system this Connector discovers
		* `services` - Services is a list of services this connector should scan.
		* `subscription_ids` - SubscriptionIds are the targets of discovery.
		* `tenant_id` - TenantId is the Azure Tenant we're discovering in
	* gcp:
		* `description` - Description of the Connector.
		* `id` - Unique identifier of the Connector.
		* `name` - Unique human-readable name of the Connector.
		* `project_ids` - ProjectIds is the list of GCP Projects the connector will scan
		* `scan_period` - ScanPeriod identifies which remote system this Connector discovers
		* `services` - Services is a list of services this connector should scan.
		* `workload_pool_id` - WorkloadPoolId is the GCP Workload Pool Identifier used to authenticate our JWT
		* `workload_project_id` - WorkloadProjectId is the GCP Project ID where the Workload Pool is defined
		* `workload_project_number` - WorkloadProjectNumber is the GCP Project Number where the Workload Pool is defined
		* `workload_provider_id` - WorkloadProviderId is the GCP Workload Provider Identifier used to authenticate our JWT
