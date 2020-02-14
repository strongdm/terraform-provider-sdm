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
			"ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the RoleAttachment.",
			},
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
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique identifier of the RoleAttachment.",
						},
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
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("composite_role_id"); ok {
		filter += "compositeroleid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("attached_role_id"); ok {
		filter += "attachedroleid:? "
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
		return fmt.Errorf("cannot list RoleAttachments %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]entity, 0)
	for resp.Next() {
		v := resp.Value()
		ids = append(ids, v.ID)
		output = append(output,
			entity{
				"id":                v.ID,
				"composite_role_id": v.CompositeRoleID,
				"attached_role_id":  v.AttachedRoleID,
			})
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("role_attachments", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("RoleAttachment" + filter + fmt.Sprint(args...))
	return nil
}
