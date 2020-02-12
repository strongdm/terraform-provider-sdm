package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceAccountRead),
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
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func dataSourceAccountRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Accounts().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Account %s: %w", d.Id(), err)
	}
	switch v := resp.Account.(type) {
	case *apiv1.User:
		d.Set("user", []map[string]interface{}{
			{
				"email":      v.Email,
				"first_name": v.FirstName,
				"last_name":  v.LastName,
			},
		})
	case *apiv1.Service:
		d.Set("service", []map[string]interface{}{
			{
				"name": v.Name,
			},
		})
	}
	return nil
}
