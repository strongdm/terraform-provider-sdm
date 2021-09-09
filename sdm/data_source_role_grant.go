// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func dataSourceRoleGrant() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceRoleGrantList),
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the RoleGrant.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the resource of this RoleGrant.",
			},
			"role_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the attached role of this RoleGrant.",
			},
			"role_grants": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique identifier of the RoleGrant.",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The id of the resource of this RoleGrant.",
						},
						"role_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The id of the attached role of this RoleGrant.",
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

func convertRoleGrantFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("resource_id"); ok {
		filter += "resourceid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("role_id"); ok {
		filter += "roleid:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceRoleGrantList(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := convertRoleGrantFilterFromResourceData(d)
	resp, err := cc.RoleGrants().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list RoleGrants %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]entity, 0)
	for resp.Next() {
		v := resp.Value()
		ids = append(ids, v.ID)
		output = append(output,
			entity{
				"id":          (v.ID),
				"resource_id": (v.ResourceID),
				"role_id":     (v.RoleID),
			})
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("role_grants", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("RoleGrant" + filter + fmt.Sprint(args...))
	return nil
}
