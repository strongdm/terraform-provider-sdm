package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceAccountGrant() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceAccountGrantList),
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
			"account_grants": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func accountGrantFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if d.Id() != "" {
		filter += "id:? "
		args = append(args, d.Id())
	}
	if v, ok := d.GetOk("resource_id"); ok {
		filter += "resource_id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("account_id"); ok {
		filter += "account_id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("start_from"); ok {
		filter += "start_from:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("valid_until"); ok {
		filter += "valid_until:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceAccountGrantList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := accountGrantFilterFromResourceData(d)
	resp, err := cc.AccountGrants().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list AccountGrant %s: %w", d.Id(), err)
	}
	vList := []string{}
	for resp.Next() {
		v := resp.Value()
		vList = append(vList, v.ID)
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", err)
	}

	err = d.Set("account_grants", vList)
	if err != nil {
		return fmt.Errorf("cannot set vList: %w", err)
	}
	d.SetId("AccountGrant" + filter + fmt.Sprint(args...))
	return nil
}
