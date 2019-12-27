package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func resourceResource() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceResourceCreate),
		Read:   wrapCrudOperation(resourceResourceRead),
		Update: wrapCrudOperation(resourceResourceUpdate),
		Delete: wrapCrudOperation(resourceResourceDelete),
		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique human-readable name of the Resource.",
			},
			"port_override": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Port number override.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func resourceFromResourceData(d *schema.ResourceData) *apiv1.Resource {
	return &apiv1.Resource{
		ID:           d.Id(),
		Name:         stringFromResourceData(d, "name"),
		PortOverride: int32FromResourceData(d, "port_override"),
	}
}

func resourceResourceCreate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.Resources().Create(ctx, resourceFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create Resource %s: %w", "", err)
	}
	d.SetId(resp.Resource.ID)
	return resourceResourceRead(d, cc)
}

func resourceResourceRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Resources().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Resource %s: %w", d.Id(), err)
	}
	v := resp.Resource

	d.Set("name", v.Name)
	d.Set("port_override", v.PortOverride)
	return nil
}

func resourceResourceUpdate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	resp, err := cc.Resources().Update(ctx, resourceFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Resource %s: %w", d.Id(), err)
	}
	d.SetId(resp.Resource.ID)
	return resourceResourceRead(d, cc)
}

func resourceResourceDelete(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.Resources().Delete(ctx, d.Id())
	return err
}
