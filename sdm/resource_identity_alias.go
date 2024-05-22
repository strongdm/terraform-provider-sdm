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

func resourceIdentityAlias() *schema.Resource {
	return &schema.Resource{
		CreateContext: wrapCrudOperation(resourceIdentityAliasCreate),
		ReadContext:   wrapCrudOperation(resourceIdentityAliasRead),
		UpdateContext: wrapCrudOperation(resourceIdentityAliasUpdate),
		DeleteContext: wrapCrudOperation(resourceIdentityAliasDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account for this identity alias.",
			},
			"identity_set_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The identity set.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username to be used as the identity alias for this account.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertIdentityAliasToPlumbing(d *schema.ResourceData) *sdm.IdentityAlias {
	return &sdm.IdentityAlias{
		ID:            d.Id(),
		AccountID:     convertStringToPlumbing(d.Get("account_id")),
		IdentitySetID: convertStringToPlumbing(d.Get("identity_set_id")),
		Username:      convertStringToPlumbing(d.Get("username")),
	}
}

func resourceIdentityAliasCreate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertIdentityAliasToPlumbing(d)
	resp, err := cc.IdentityAliases().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create IdentityAlias: %w", err)
	}
	d.SetId(resp.IdentityAlias.ID)
	v := resp.IdentityAlias
	d.Set("account_id", (v.AccountID))
	d.Set("identity_set_id", (v.IdentitySetID))
	d.Set("username", (v.Username))
	return nil
}

func resourceIdentityAliasRead(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertIdentityAliasToPlumbing(d)
	_ = localVersion
	resp, err := cc.IdentityAliases().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read IdentityAlias %s: %w", d.Id(), err)
	}
	v := resp.IdentityAlias
	d.Set("account_id", (v.AccountID))
	d.Set("identity_set_id", (v.IdentitySetID))
	d.Set("username", (v.Username))
	return nil
}
func resourceIdentityAliasUpdate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	resp, err := cc.IdentityAliases().Update(ctx, convertIdentityAliasToPlumbing(d))
	if err != nil {
		return fmt.Errorf("cannot update IdentityAlias %s: %w", d.Id(), err)
	}
	d.SetId(resp.IdentityAlias.ID)
	return resourceIdentityAliasRead(ctx, d, cc)
}
func resourceIdentityAliasDelete(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	var errNotFound *sdm.NotFoundError
	_, err := cc.IdentityAliases().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
