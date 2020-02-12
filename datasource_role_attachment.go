package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceRoleAttachmentList),
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
			"role_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
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
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func roleAttachmentFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if d.Id() != "" {
		filter += "id:? "
		args = append(args, d.Id())
	}
	if v, ok := d.GetOk("composite_role_id"); ok {
		filter += "composite_role_id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("attached_role_id"); ok {
		filter += "attached_role_id:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceRoleAttachmentList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := roleAttachmentFilterFromResourceData(d)
	resp, err := cc.RoleAttachments().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list RoleAttachment %s: %w", d.Id(), err)
	}
	vList := []map[string]interface{}{}
	for resp.Next() {
		v := resp.Value()
		vList = append(vList,
			map[string]interface{}{
				"composite_role_id": v.CompositeRoleID,
				"attached_role_id":  v.AttachedRoleID,
			})
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", err)
	}

	err = d.Set("role_attachments", vList)
	if err != nil {
		return fmt.Errorf("cannot set vList: %w", err)
	}
	d.SetId("RoleAttachment" + filter + fmt.Sprint(args...))
	return nil
}
