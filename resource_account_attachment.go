package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

func resourceAccountAttachment() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceAccountAttachmentCreate),
		Read:   wrapCrudOperation(resourceAccountAttachmentRead),
		Delete: wrapCrudOperation(resourceAccountAttachmentDelete),
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the account of this AccountAttachment.",
			},
			"role_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the attached role of this AccountAttachment.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func accountAttachmentFromResourceData(d *schema.ResourceData) *sdm.AccountAttachment {
	return &sdm.AccountAttachment{
		ID:        d.Id(),
		AccountID: stringFromResourceData(d, "account_id"),
		RoleID:    stringFromResourceData(d, "role_id"),
	}
}

func resourceAccountAttachmentCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.AccountAttachments().Create(ctx, accountAttachmentFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create AccountAttachment %s: %w", "", err)
	}
	d.SetId(resp.AccountAttachment.ID)
	v := resp.AccountAttachment
	d.Set("account_id", v.AccountID)
	d.Set("role_id", v.RoleID)
	return nil
}

func resourceAccountAttachmentRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.AccountAttachments().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
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
func resourceAccountAttachmentDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.AccountAttachments().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
