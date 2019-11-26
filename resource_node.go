package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiV1 "github.com/strongdm/strongdm-sdk-go"
	errors_v1 "github.com/strongdm/strongdm-sdk-go/errors"
	models_v1 "github.com/strongdm/strongdm-sdk-go/models"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceNodeCreate),
		Read:   wrapCrudOperation(resourceNodeRead),
		Update: wrapCrudOperation(resourceNodeUpdate),
		Delete: wrapCrudOperation(resourceNodeDelete),
		Schema: map[string]*schema.Schema{
			"relay": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Description: "Relay represents a StrongDM CLI installation running in relay mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Description: "Unique human-readable name of the Relay.",
						},
					},
				},
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Description: "Gateway represents a StrongDM CLI installation running in gateway mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Description: "Unique human-readable name of the Relay.",
						},
						"listen_address": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Description: "The public hostname/port tuple at which the gateway will be accessible to clients.",
						},
						"bind_address": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							Description: "The hostname/port tuple which the gateway daemon will bind to.",
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

func nodeFromResourceData(d *schema.ResourceData) models_v1.Node {
	if list := d.Get("relay").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &models_v1.Relay{
			ID:        d.Id(),
			Name: stringFromMap(raw, "name"),
		}
	}
	if list := d.Get("gateway").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &models_v1.Gateway{
			ID:        d.Id(),
			Name: stringFromMap(raw, "name"),
			ListenAddress: stringFromMap(raw, "listen_address"),
			BindAddress: stringFromMap(raw, "bind_address"),
		}
	}
	return nil
}

func resourceNodeCreate(d *schema.ResourceData, cc *apiV1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.Nodes().Create(ctx, nodeFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create Node %s: %w", "", err)
	}
	d.SetId(resp.Node.GetID())
	return resourceNodeRead(d, cc)
}

func resourceNodeRead(d *schema.ResourceData, cc *apiV1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Nodes().Get(ctx, d.Id())
	var errNotFound *errors_v1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Node %s: %w", d.Id(), err)
	}
	switch v := resp.Node.(type) {
	case *models_v1.Relay:
		d.Set("relay", []map[string]interface{}{
			map[string]interface{}{
				"name": v.Name,
			},
		})
	case *models_v1.Gateway:
		d.Set("gateway", []map[string]interface{}{
			map[string]interface{}{
				"name": v.Name,
				"listen_address": v.ListenAddress,
				"bind_address": v.BindAddress,
			},
		})
	}
	return nil
}

func resourceNodeUpdate(d *schema.ResourceData, cc *apiV1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	resp, err := cc.Nodes().Update(ctx, nodeFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Node %s: %w", d.Id(), err)
	}
	d.SetId(resp.Node.GetID())
	return resourceNodeRead(d, cc)
}

func resourceNodeDelete(d *schema.ResourceData, cc *apiV1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.Nodes().Delete(ctx, d.Id())
	return err
}