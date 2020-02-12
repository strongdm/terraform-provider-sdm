package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func resourceAccountAttachment() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceAccountAttachmentCreate),
		Read:   wrapCrudOperation(resourceAccountAttachmentRead),
		Update: wrapCrudOperation(resourceAccountAttachmentUpdate),
		Delete: wrapCrudOperation(resourceAccountAttachmentDelete),
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the account of this AccountAttachment.",
			},
			"role_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the attached role of this AccountAttachment.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func accountAttachmentFromResourceData(d *schema.ResourceData) *apiv1.AccountAttachment {
	return &apiv1.AccountAttachment{
		ID:        d.Id(),
		AccountID: stringFromResourceData(d, "account_id"),
		RoleID:    stringFromResourceData(d, "role_id"),
	}
}

func resourceAccountAttachmentCreate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.AccountAttachments().Create(ctx, accountAttachmentFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create AccountAttachment %s: %w", "", err)
	}
	d.SetId(resp.AccountAttachment.ID)
	return resourceAccountAttachmentRead(d, cc)
}

func resourceAccountAttachmentRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.AccountAttachments().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read AccountAttachment %s: %w", d.Id(), err)
	}
	v := resp.AccountAttachment
	d.Set("account_id", v.AccountID)
	d.Set("role_id", v.RoleID)
	return nil
}

func resourceAccountAttachmentUpdate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	model := accountAttachmentFromResourceData(d)
	id := model.ID
	_, err := cc.AccountAttachments().Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot update AccountAttachment %s: %w", d.Id(), err)
	}
	resp, err := cc.AccountAttachments().Create(ctx, model)
	if err != nil {
		return fmt.Errorf("cannot update AccountAttachment %s: %w", d.Id(), err)
	}
	d.SetId(resp.AccountAttachment.ID)
	return resourceAccountAttachmentRead(d, cc)
}

func resourceAccountAttachmentDelete(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.AccountAttachments().Delete(ctx, d.Id())
	return err
}
