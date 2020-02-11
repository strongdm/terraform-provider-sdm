package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func resourceAccountGrant() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceAccountGrantCreate),
		Read:   wrapCrudOperation(resourceAccountGrantRead),
		Update: wrapCrudOperation(resourceAccountGrantUpdate),
		Delete: wrapCrudOperation(resourceAccountGrantDelete),
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the composite role of this AccountGrant.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the attached role of this AccountGrant.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func accountGrantFromResourceData(d *schema.ResourceData) *apiv1.AccountGrant {
	return &apiv1.AccountGrant{
		ID:         d.Id(),
		ResourceID: stringFromResourceData(d, "resource_id"),
		AccountID:  stringFromResourceData(d, "account_id"),
	}
}

func resourceAccountGrantCreate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.AccountGrants().Create(ctx, accountGrantFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create AccountGrant %s: %w", "", err)
	}
	d.SetId(resp.AccountGrant.ID)
	return resourceAccountGrantRead(d, cc)
}

func resourceAccountGrantRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.AccountGrants().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read AccountGrant %s: %w", d.Id(), err)
	}
	v := resp.AccountGrant
	d.Set("resource_id", v.ResourceID)
	d.Set("account_id", v.AccountID)
	return nil
}

func resourceAccountGrantUpdate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	model := accountGrantFromResourceData(d)
	id := model.ID
	_, err := cc.AccountGrants().Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot update AccountGrant %s: %w", d.Id(), err)
	}
	resp, err := cc.AccountGrants().Create(ctx, model)
	if err != nil {
		return fmt.Errorf("cannot update AccountGrant %s: %w", d.Id(), err)
	}
	d.SetId(resp.AccountGrant.ID)
	return resourceAccountGrantRead(d, cc)
}

func resourceAccountGrantDelete(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.AccountGrants().Delete(ctx, d.Id())
	return err
}
