package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceAccountGrant() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceAccountGrantRead),
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
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func dataSourceAccountGrantRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.AccountGrants().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read AccountGrant %s: %w", d.Id(), err)
	}
	v := resp.AccountGrant
	d.Set("resource_id", v.ResourceID)
	d.Set("account_id", v.AccountID)
	return nil
}
