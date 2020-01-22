package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func resourceRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceRoleAttachmentCreate),
		Read:   wrapCrudOperation(resourceRoleAttachmentRead),
		Update: wrapCrudOperation(resourceRoleAttachmentUpdate),
		Delete: wrapCrudOperation(resourceRoleAttachmentDelete),
		Schema: map[string]*schema.Schema{
			"composite_role_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the composite role of this RoleAttachment.",
			},
			"attached_role_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the attached role of this RoleAttachment.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func roleAttachmentFromResourceData(d *schema.ResourceData) *apiv1.RoleAttachment {
	return &apiv1.RoleAttachment{
		ID:              d.Id(),
		CompositeRoleID: stringFromResourceData(d, "composite_role_id"),
		AttachedRoleID:  stringFromResourceData(d, "attached_role_id"),
	}
}

func resourceRoleAttachmentCreate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.RoleAttachments().Create(ctx, roleAttachmentFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create RoleAttachment %s: %w", "", err)
	}
	d.SetId(resp.RoleAttachment.ID)
	return resourceRoleAttachmentRead(d, cc)
}

func resourceRoleAttachmentRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.RoleAttachments().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read RoleAttachment %s: %w", d.Id(), err)
	}
	v := resp.RoleAttachment
	d.Set("composite_role_id", v.CompositeRoleID)
	d.Set("attached_role_id", v.AttachedRoleID)
	return nil
}

func resourceRoleAttachmentUpdate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	model := roleAttachmentFromResourceData(d)
	id := model.ID
	_, err := cc.RoleAttachments().Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot update RoleAttachment %s: %w", d.Id(), err)
	}
	resp, err := cc.RoleAttachments().Create(ctx, model)
	if err != nil {
		return fmt.Errorf("cannot update RoleAttachment %s: %w", d.Id(), err)
	}
	d.SetId(resp.RoleAttachment.ID)
	return resourceRoleAttachmentRead(d, cc)
}

func resourceRoleAttachmentDelete(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.RoleAttachments().Delete(ctx, d.Id())
	return err
}
