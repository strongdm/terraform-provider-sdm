package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceAccountList),
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A User can connect to resources they are granted directly, or granted\n via roles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"service": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Service is a service account that can connect to resources they are granted\n directly, or granted via roles. Services are typically automated jobs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique human-readable name of the Service.",
						},
					},
				},
			},
			"accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A User can connect to resources they are granted directly, or granted\n via roles.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
								},
							},
						},
						"service": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A Service is a service account that can connect to resources they are granted\n directly, or granted via roles. Services are typically automated jobs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Service.",
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
	if d.Id() != "" {
		filter += "id:? "
		args = append(args, d.Id())
	}
	// todo
	return filter, args
}

func dataSourceAccountList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := accountFilterFromResourceData(d)
	resp, err := cc.Accounts().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Account %s: %w", d.Id(), err)
	}
	vList := []map[string]interface{}{}
	for resp.Next() {
		v := resp.Value()
		// TODO: fix it!
		fmt.Println(v)
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", err)
	}

	err = d.Set("accounts", vList)
	if err != nil {
		return fmt.Errorf("cannot set vList: %w", err)
	}
	d.SetId("Account" + filter + fmt.Sprint(args...))
	return nil
}
