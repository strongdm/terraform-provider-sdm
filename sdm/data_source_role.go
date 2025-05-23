// Code generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: wrapCrudOperation(dataSourceRoleList),
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the Role.",
			},
			"managed_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Managed By is a read only field for what service manages this role, e.g. StrongDM, Okta, Azure.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique human-readable name of the Role.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Elem:        tagsElemType,
				Optional:    true,
				Description: "Tags is a map of key, value pairs.",
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_rules": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AccessRules is a list of access rules defining the resources this Role has access to.",
						},
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique identifier of the Role.",
						},
						"managed_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Managed By is a read only field for what service manages this role, e.g. StrongDM, Okta, Azure.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique human-readable name of the Role.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Elem:        tagsElemType,
							Optional:    true,
							Description: "Tags is a map of key, value pairs.",
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

func init() {
	dataSourcesMap["sdm_role"] = dataSourceRole
}

func convertRoleFilterToPlumbing(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOkExists("access_rules"); ok {
		filter += "accessrules:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("managed_by"); ok {
		filter += "managedby:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("tags"); ok {
		tags := convertTagsToPlumbing(v)
		for kk, vv := range tags {
			filter += "tag:?=? "
			args = append(args, kk, vv)
		}
	}
	return filter, args
}

func dataSourceRoleList(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	filter, args := convertRoleFilterToPlumbing(d)
	resp, err := cc.Roles().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Roles %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]entity, 0)
	for resp.Next() {
		v := resp.Value()
		ids = append(ids, v.ID)
		output = append(output,
			entity{
				"access_rules": convertAccessRulesToPorcelain(v.AccessRules),
				"id":           (v.ID),
				"managed_by":   (v.ManagedBy),
				"name":         (v.Name),
				"tags":         convertTagsToPorcelain(v.Tags),
			})
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("roles", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("Role" + filter + fmt.Sprint(args...))
	return nil
}
