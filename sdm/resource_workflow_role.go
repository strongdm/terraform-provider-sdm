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

func resourceWorkflowRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: wrapCrudOperation(resourceWorkflowRoleCreate),
		ReadContext:   wrapCrudOperation(resourceWorkflowRoleRead),
		DeleteContext: wrapCrudOperation(resourceWorkflowRoleDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"role_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The role id.",
			},
			"workflow_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The workflow id.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertWorkflowRoleToPlumbing(d *schema.ResourceData) *sdm.WorkflowRole {
	return &sdm.WorkflowRole{
		ID:         d.Id(),
		RoleID:     convertStringToPlumbing(d.Get("role_id")),
		WorkflowID: convertStringToPlumbing(d.Get("workflow_id")),
	}
}

func resourceWorkflowRoleCreate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertWorkflowRoleToPlumbing(d)
	resp, err := cc.WorkflowRoles().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create WorkflowRole: %w", err)
	}
	d.SetId(resp.WorkflowRole.ID)
	v := resp.WorkflowRole
	d.Set("role_id", (v.RoleID))
	d.Set("workflow_id", (v.WorkflowID))
	return nil
}

func resourceWorkflowRoleRead(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertWorkflowRoleToPlumbing(d)
	_ = localVersion
	resp, err := cc.WorkflowRoles().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read WorkflowRole %s: %w", d.Id(), err)
	}
	v := resp.WorkflowRole
	d.Set("role_id", (v.RoleID))
	d.Set("workflow_id", (v.WorkflowID))
	return nil
}
func resourceWorkflowRoleDelete(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	var errNotFound *sdm.NotFoundError
	_, err := cc.WorkflowRoles().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}