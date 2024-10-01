// Code generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func resourcePeeringGroupResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: wrapCrudOperation(resourcePeeringGroupResourceCreate),
		ReadContext:   wrapCrudOperation(resourcePeeringGroupResourceRead),
		DeleteContext: wrapCrudOperation(resourcePeeringGroupResourceDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Peering Group ID to which the resource will be attached to.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Resource ID to be attached.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
			Create:  schema.DefaultTimeout(60 * time.Second),
			Read:    schema.DefaultTimeout(60 * time.Second),
			Delete:  schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertPeeringGroupResourceToPlumbing(d *schema.ResourceData) *sdm.PeeringGroupResource {
	return &sdm.PeeringGroupResource{
		ID:         d.Id(),
		GroupID:    convertStringToPlumbing(d.Get("group_id")),
		ResourceID: convertStringToPlumbing(d.Get("resource_id")),
	}
}

func resourcePeeringGroupResourceCreate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertPeeringGroupResourceToPlumbing(d)
	resp, err := cc.PeeringGroupResources().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create PeeringGroupResource: %w", err)
	}
	d.SetId(resp.PeeringGroupResource.ID)
	v := resp.PeeringGroupResource
	d.Set("group_id", (v.GroupID))
	d.Set("resource_id", (v.ResourceID))
	return nil
}

func resourcePeeringGroupResourceRead(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertPeeringGroupResourceToPlumbing(d)
	_ = localVersion
	resp, err := cc.PeeringGroupResources().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read PeeringGroupResource %s: %w", d.Id(), err)
	}
	v := resp.PeeringGroupResource
	d.Set("group_id", (v.GroupID))
	d.Set("resource_id", (v.ResourceID))
	return nil
}
func resourcePeeringGroupResourceDelete(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	var errNotFound *sdm.NotFoundError
	_, err := cc.PeeringGroupResources().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
