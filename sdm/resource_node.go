// Code generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		CreateContext: wrapCrudOperation(resourceNodeCreate),
		ReadContext:   wrapCrudOperation(resourceNodeRead),
		UpdateContext: wrapCrudOperation(resourceNodeUpdate),
		DeleteContext: wrapCrudOperation(resourceNodeDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"gateway": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Gateway represents a StrongDM CLI installation running in gateway mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bind_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The hostname/port tuple which the gateway daemon will bind to. If not provided on create, set to \"0.0.0.0:listen_address_port\".",
						},
						"device": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Device is a read only device name uploaded by the gateway process when it comes online.",
						},
						"gateway_filter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GatewayFilter can be used to restrict the peering between relays and gateways. Deprecated.",
						},
						"listen_address": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The public hostname/port tuple at which the gateway will be accessible to clients.",
						},
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location is a read only network location uploaded by the gateway process when it comes online.",
						},
						"maintenance_window": {
							Type:        schema.TypeList,
							Elem:        nodeMaintenanceWindowElemType,
							Optional:    true,
							Description: "Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Unique human-readable name of the Gateway. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version is a read only sdm binary version uploaded by the gateway process when it comes online.",
						},
						"token": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
					},
				},
			},
			"proxy_cluster": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "ProxyCluster represents a cluster of StrongDM proxies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The public hostname/port tuple at which the proxy cluster will be accessible to clients.",
						},
						"maintenance_window": {
							Type:        schema.TypeList,
							Elem:        nodeMaintenanceWindowElemType,
							Optional:    true,
							Description: "Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Unique human-readable name of the proxy cluster. Names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
					},
				},
			},
			"relay": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Relay represents a StrongDM CLI installation running in relay mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Device is a read only device name uploaded by the gateway process when it comes online.",
						},
						"gateway_filter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "GatewayFilter can be used to restrict the peering between relays and gateways. Deprecated.",
						},
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location is a read only network location uploaded by the gateway process when it comes online.",
						},
						"maintenance_window": {
							Type:        schema.TypeList,
							Elem:        nodeMaintenanceWindowElemType,
							Optional:    true,
							Description: "Maintenance Windows define when this node is allowed to restart. If a node is requested to restart, it will check each window to determine if any of them permit it to restart, and if any do, it will. This check is repeated per window until the restart is successfully completed.  If not set here, may be set on the command line or via an environment variable on the process itself; any server setting will take precedence over local settings. This setting is ineffective for nodes below version 38.44.0.  If this setting is not applied via this remote configuration or via local configuration, the default setting is used: always allow restarts if serving no connections, and allow a restart even if serving connections between 7-8 UTC, any day.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Unique human-readable name of the Relay. Node names must include only letters, numbers, and hyphens (no spaces, underscores, or other special characters). Generated if not provided on create.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version is a read only sdm binary version uploaded by the gateway process when it comes online.",
						},
						"token": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertNodeToPlumbing(d *schema.ResourceData) sdm.Node {
	if list := d.Get("gateway").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Gateway{}
		}
		out := &sdm.Gateway{
			ID:                 d.Id(),
			BindAddress:        convertStringToPlumbing(raw["bind_address"]),
			GatewayFilter:      convertStringToPlumbing(raw["gateway_filter"]),
			ListenAddress:      convertStringToPlumbing(raw["listen_address"]),
			MaintenanceWindows: convertRepeatedNodeMaintenanceWindowToPlumbing(raw["maintenance_window"]),
			Name:               convertStringToPlumbing(raw["name"]),
			Tags:               convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	if list := d.Get("proxy_cluster").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.ProxyCluster{}
		}
		out := &sdm.ProxyCluster{
			ID:                 d.Id(),
			Address:            convertStringToPlumbing(raw["address"]),
			MaintenanceWindows: convertRepeatedNodeMaintenanceWindowToPlumbing(raw["maintenance_window"]),
			Name:               convertStringToPlumbing(raw["name"]),
			Tags:               convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	if list := d.Get("relay").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Relay{}
		}
		out := &sdm.Relay{
			ID:                 d.Id(),
			GatewayFilter:      convertStringToPlumbing(raw["gateway_filter"]),
			MaintenanceWindows: convertRepeatedNodeMaintenanceWindowToPlumbing(raw["maintenance_window"]),
			Name:               convertStringToPlumbing(raw["name"]),
			Tags:               convertTagsToPlumbing(raw["tags"]),
		}
		return out
	}
	return nil
}

func resourceNodeCreate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertNodeToPlumbing(d)
	resp, err := cc.Nodes().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create Node: %w", err)
	}
	d.SetId(resp.Node.GetID())
	switch v := resp.Node.(type) {
	case *sdm.Gateway:
		localV, _ := localVersion.(*sdm.Gateway)
		_ = localV
		d.Set("gateway", []map[string]interface{}{
			{
				"bind_address":       (v.BindAddress),
				"device":             (v.Device),
				"gateway_filter":     (v.GatewayFilter),
				"listen_address":     (v.ListenAddress),
				"location":           (v.Location),
				"maintenance_window": convertRepeatedNodeMaintenanceWindowToPorcelain(v.MaintenanceWindows),
				"name":               (v.Name),
				"tags":               convertTagsToPorcelain(v.Tags),
				"version":            (v.Version),
				"token":              resp.Token,
			},
		})
	case *sdm.ProxyCluster:
		localV, _ := localVersion.(*sdm.ProxyCluster)
		_ = localV
		d.Set("proxy_cluster", []map[string]interface{}{
			{
				"address":            (v.Address),
				"maintenance_window": convertRepeatedNodeMaintenanceWindowToPorcelain(v.MaintenanceWindows),
				"name":               (v.Name),
				"tags":               convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.Relay:
		localV, _ := localVersion.(*sdm.Relay)
		_ = localV
		d.Set("relay", []map[string]interface{}{
			{
				"device":             (v.Device),
				"gateway_filter":     (v.GatewayFilter),
				"location":           (v.Location),
				"maintenance_window": convertRepeatedNodeMaintenanceWindowToPorcelain(v.MaintenanceWindows),
				"name":               (v.Name),
				"tags":               convertTagsToPorcelain(v.Tags),
				"version":            (v.Version),
				"token":              resp.Token,
			},
		})
	}
	return nil
}

func resourceNodeRead(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertNodeToPlumbing(d)
	_ = localVersion
	resp, err := cc.Nodes().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Node %s: %w", d.Id(), err)
	}
	switch v := resp.Node.(type) {
	case *sdm.Gateway:
		localV, ok := localVersion.(*sdm.Gateway)
		if !ok {
			localV = &sdm.Gateway{}
		}
		_ = localV
		d.Set("gateway", []map[string]interface{}{
			{
				"bind_address":       (v.BindAddress),
				"device":             (v.Device),
				"gateway_filter":     (v.GatewayFilter),
				"listen_address":     (v.ListenAddress),
				"location":           (v.Location),
				"maintenance_window": convertRepeatedNodeMaintenanceWindowToPorcelain(v.MaintenanceWindows),
				"name":               (v.Name),
				"tags":               convertTagsToPorcelain(v.Tags),
				"version":            (v.Version),
				"token":              d.Get("gateway.0.token"),
			},
		})
	case *sdm.ProxyCluster:
		localV, ok := localVersion.(*sdm.ProxyCluster)
		if !ok {
			localV = &sdm.ProxyCluster{}
		}
		_ = localV
		d.Set("proxy_cluster", []map[string]interface{}{
			{
				"address":            (v.Address),
				"maintenance_window": convertRepeatedNodeMaintenanceWindowToPorcelain(v.MaintenanceWindows),
				"name":               (v.Name),
				"tags":               convertTagsToPorcelain(v.Tags),
			},
		})
	case *sdm.Relay:
		localV, ok := localVersion.(*sdm.Relay)
		if !ok {
			localV = &sdm.Relay{}
		}
		_ = localV
		d.Set("relay", []map[string]interface{}{
			{
				"device":             (v.Device),
				"gateway_filter":     (v.GatewayFilter),
				"location":           (v.Location),
				"maintenance_window": convertRepeatedNodeMaintenanceWindowToPorcelain(v.MaintenanceWindows),
				"name":               (v.Name),
				"tags":               convertTagsToPorcelain(v.Tags),
				"version":            (v.Version),
				"token":              d.Get("relay.0.token"),
			},
		})
	}
	return nil
}
func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	resp, err := cc.Nodes().Update(ctx, convertNodeToPlumbing(d))
	if err != nil {
		return fmt.Errorf("cannot update Node %s: %w", d.Id(), err)
	}
	d.SetId(resp.Node.GetID())
	return resourceNodeRead(ctx, d, cc)
}
func resourceNodeDelete(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	var errNotFound *sdm.NotFoundError
	_, err := cc.Nodes().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
