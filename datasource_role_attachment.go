package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceRoleAttachmentRead),
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

func dataSourceRoleAttachmentRead(d *schema.ResourceData, cc *apiv1.Client) error {
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
