package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func resourceRoleGrant() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceRoleGrantCreate),
		Read:   wrapCrudOperation(resourceRoleGrantRead),
		Update: wrapCrudOperation(resourceRoleGrantUpdate),
		Delete: wrapCrudOperation(resourceRoleGrantDelete),
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the resource of this RoleGrant.",
			},
			"role_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the attached role of this RoleGrant.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func roleGrantFromResourceData(d *schema.ResourceData) *apiv1.RoleGrant {
	return &apiv1.RoleGrant{
		ID:         d.Id(),
		ResourceID: stringFromResourceData(d, "resource_id"),
		RoleID:     stringFromResourceData(d, "role_id"),
	}
}

func resourceRoleGrantCreate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.RoleGrants().Create(ctx, roleGrantFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create RoleGrant %s: %w", "", err)
	}
	d.SetId(resp.RoleGrant.ID)
	return resourceRoleGrantRead(d, cc)
}

func resourceRoleGrantRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.RoleGrants().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read RoleGrant %s: %w", d.Id(), err)
	}
	v := resp.RoleGrant
	d.Set("resource_id", v.ResourceID)
	d.Set("role_id", v.RoleID)
	return nil
}

func resourceRoleGrantUpdate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	model := roleGrantFromResourceData(d)
	id := model.ID
	_, err := cc.RoleGrants().Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot update RoleGrant %s: %w", d.Id(), err)
	}
	resp, err := cc.RoleGrants().Create(ctx, model)
	if err != nil {
		return fmt.Errorf("cannot update RoleGrant %s: %w", d.Id(), err)
	}
	d.SetId(resp.RoleGrant.ID)
	return resourceRoleGrantRead(d, cc)
}

func resourceRoleGrantDelete(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.RoleGrants().Delete(ctx, d.Id())
	return err
}
