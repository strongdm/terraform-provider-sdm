package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceRoleCreate),
		Read:   wrapCrudOperation(resourceRoleRead),
		Update: wrapCrudOperation(resourceRoleUpdate),
		Delete: wrapCrudOperation(resourceRoleDelete),
		Schema: map[string]*schema.Schema{
			
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "Unique human-readable name of the Role.",
			},
			"composite": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				Description: "True if the Role is a composite role.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func roleFromResourceData(d *schema.ResourceData) *apiv1.Role {
	return &apiv1.Role{
		ID:        d.Id(),
		Name: stringFromResourceData(d, "name"),
		Composite: boolFromResourceData(d, "composite"),
	}
}

func resourceRoleCreate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.Roles().Create(ctx, roleFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create Role %s: %w", "", err)
	}
	d.SetId(resp.Role.ID)
	return resourceRoleRead(d, cc)
}

func resourceRoleRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Roles().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Role %s: %w", d.Id(), err)
	}
	v := resp.Role
	
	d.Set("name", v.Name)
	d.Set("composite", v.Composite)
	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	resp, err := cc.Roles().Update(ctx, roleFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Role %s: %w", d.Id(), err)
	}
	d.SetId(resp.Role.ID)
	return resourceRoleRead(d, cc)
}

func resourceRoleDelete(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.Roles().Delete(ctx, d.Id())
	return err
}