package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceRoleList),
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique human-readable name of the Role.",
			},
			"composite": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if the Role is a composite role.",
			},
			"roles": {
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

func roleFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if d.Id() != "" {
		filter += "id:? "
		args = append(args, d.Id())
	}
	if v, ok := d.GetOk("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("composite"); ok {
		filter += "composite:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceRoleList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := roleFilterFromResourceData(d)
	resp, err := cc.Roles().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Role %s: %w", d.Id(), err)
	}
	vList := []string{}
	for resp.Next() {
		v := resp.Value()
		vList = append(vList, v.ID)
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", err)
	}

	err = d.Set("roles", vList)
	if err != nil {
		return fmt.Errorf("cannot set vList: %w", err)
	}
	d.SetId("Role" + filter + fmt.Sprint(args...))
	return nil
}
