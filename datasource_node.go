package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceNode() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceNodeList),
		Schema: map[string]*schema.Schema{
			"relay": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Relay represents a StrongDM CLI installation running in relay mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique human-readable name of the Relay.",
						},
					},
				},
			},
			"gateway": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Gateway represents a StrongDM CLI installation running in gateway mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique human-readable name of the Relay.",
						},
						"listen_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The public hostname/port tuple at which the gateway will be accessible to clients.",
						},
						"bind_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The hostname/port tuple which the gateway daemon will bind to.",
						},
					},
				},
			},
			"nodes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func nodeFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if d.Id() != "" {
		filter += "id:? "
		args = append(args, d.Id())
	}
	// todo
	return filter, args
}

func dataSourceNodeList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := nodeFilterFromResourceData(d)
	resp, err := cc.Nodes().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Node %s: %w", d.Id(), err)
	}
	vList := []string{}
	for resp.Next() {
		v := resp.Value()
		vList = append(vList, v.GetID())
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", err)
	}

	err = d.Set("nodes", vList)
	if err != nil {
		return fmt.Errorf("cannot set vList: %w", err)
	}
	d.SetId("Node" + filter + fmt.Sprint(args...))
	return nil
}
