package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceAccountList),
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"suspended": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"accounts": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"user": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A User can connect to resources they are granted directly, or granted\n via roles.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the User.",
									},
									"email": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The User's email address. Must be unique.",
									},
									"first_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The User's first name.",
									},
									"last_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The User's last name.",
									},
									"suspended": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "The User's suspended state.",
									},
								},
							},
						},
						"service": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A Service is a service account that can connect to resources they are granted\n directly, or granted via roles. Services are typically automated jobs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Service.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Service.",
									},
									"suspended": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "The Service's suspended state.",
									},
								},
							},
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

func accountFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOk("type"); ok {
		filter += "type:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("email"); ok {
		filter += "email:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("first_name"); ok {
		filter += "firstname:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("last_name"); ok {
		filter += "lastname:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("suspended"); ok {
		filter += "suspended:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceAccountList(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := accountFilterFromResourceData(d)
	resp, err := cc.Accounts().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Accounts %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]map[string][]entity, 1)
	output[0] = map[string][]entity{
		"user": {},
	}
	for resp.Next() {
		ids = append(ids, resp.Value().GetID())
		switch v := resp.Value().(type) {
		case *sdm.User:
			output[0]["user"] = append(output[0]["user"], entity{
				"id":         v.ID,
				"email":      v.Email,
				"first_name": v.FirstName,
				"last_name":  v.LastName,
				"suspended":  v.Suspended,
			})
		case *sdm.Service:
			output[0]["service"] = append(output[0]["service"], entity{
				"id":        v.ID,
				"name":      v.Name,
				"suspended": v.Suspended,
			})
		}
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("accounts", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("Account" + filter + fmt.Sprint(args...))
	return nil
}