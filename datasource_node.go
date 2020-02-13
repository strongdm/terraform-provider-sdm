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
			"bind_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listen_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bind_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"listen_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
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

func nodeFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOk("bind_address"); ok {
		filter += "bind_address:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("listen_address"); ok {
		filter += "listen_address:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("state"); ok {
		filter += "state:? "
		args = append(args, v)
	}
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
	vList := []map[string]interface{}{}
	for resp.Next() {
		switch v := resp.Value().(type) {
		case *apiv1.Relay:
			vList = append(vList,
				map[string]interface{}{
					"type": "relay",
					"id":   v.ID,
					"name": v.Name,
				})
		case *apiv1.Gateway:
			vList = append(vList,
				map[string]interface{}{
					"type":           "gateway",
					"id":             v.ID,
					"name":           v.Name,
					"listen_address": v.ListenAddress,
					"bind_address":   v.BindAddress,
				})
		}
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
