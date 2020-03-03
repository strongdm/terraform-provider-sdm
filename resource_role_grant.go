package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

func resourceRoleGrant() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceRoleGrantCreate),
		Read:   wrapCrudOperation(resourceRoleGrantRead),
		Delete: wrapCrudOperation(resourceRoleGrantDelete),
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the resource of this RoleGrant.",
			},
			"role_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the attached role of this RoleGrant.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func roleGrantFromResourceData(d *schema.ResourceData) *sdm.RoleGrant {
	return &sdm.RoleGrant{
		ID:         d.Id(),
		ResourceID: stringFromResourceData(d, "resource_id"),
		RoleID:     stringFromResourceData(d, "role_id"),
	}
}

func resourceRoleGrantCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.RoleGrants().Create(ctx, roleGrantFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create RoleGrant %s: %w", "", err)
	}
	d.SetId(resp.RoleGrant.ID)
	v := resp.RoleGrant
	d.Set("resource_id", v.ResourceID)
	d.Set("role_id", v.RoleID)
	return nil
}

func resourceRoleGrantRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.RoleGrants().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
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
func resourceRoleGrantDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.RoleGrants().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
