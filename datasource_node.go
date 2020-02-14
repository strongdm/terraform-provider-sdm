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
			"ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
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
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"relay": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Relay represents a StrongDM CLI installation running in relay mode.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Relay.",
									},
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
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Relay.",
									},
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
	if v, ok := d.GetOk("type"); ok {
		filter += "type:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("bind_address"); ok {
		filter += "bindaddress:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("listen_address"); ok {
		filter += "listenaddress:? "
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
		return fmt.Errorf("cannot list Nodes %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]map[string][]entity, 1)
	output[0] = make(map[string][]entity)
	for resp.Next() {
		ids = append(ids, resp.Value().GetID())
		switch v := resp.Value().(type) {
		case *apiv1.Relay:
			output[0]["relay"] = append(output[0]["relay"], entity{
				"id":   v.ID,
				"name": v.Name,
			})
		case *apiv1.Gateway:
			output[0]["gateway"] = append(output[0]["gateway"], entity{
				"id":             v.ID,
				"name":           v.Name,
				"listen_address": v.ListenAddress,
				"bind_address":   v.BindAddress,
			})
		}
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("nodes", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("Node" + filter + fmt.Sprint(args...))
	return nil
}
