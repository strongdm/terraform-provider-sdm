package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceAccountAttachment() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceAccountAttachmentList),
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the AccountAttachment.",
			},
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
			"account_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique identifier of the AccountAttachment.",
						},
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
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func accountAttachmentFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("account_id"); ok {
		filter += "accountid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("role_id"); ok {
		filter += "roleid:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceAccountAttachmentList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := accountAttachmentFilterFromResourceData(d)
	resp, err := cc.AccountAttachments().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list AccountAttachments %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]entity, 0)
	for resp.Next() {
		v := resp.Value()
		ids = append(ids, v.ID)
		output = append(output,
			entity{
				"id":         v.ID,
				"account_id": v.AccountID,
				"role_id":    v.RoleID,
			})
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("account_attachments", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("AccountAttachment" + filter + fmt.Sprint(args...))
	return nil
}
