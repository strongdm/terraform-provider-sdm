package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceRoleRead),
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
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func dataSourceRoleRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Roles().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Role %s: %w", d.Id(), err)
	}
	v := resp.Role
	d.Set("name", v.Name)
	d.Set("composite", v.Composite)
	return nil
}
