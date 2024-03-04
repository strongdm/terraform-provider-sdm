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

func resourceApprovalWorkflowApprover() *schema.Resource {
	return &schema.Resource{
		CreateContext: wrapCrudOperation(resourceApprovalWorkflowApproverCreate),
		ReadContext:   wrapCrudOperation(resourceApprovalWorkflowApproverRead),
		DeleteContext: wrapCrudOperation(resourceApprovalWorkflowApproverDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The approver account id.",
			},
			"approval_flow_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The approval flow id specified the approval workflow that this approver belongs to",
			},
			"approval_step_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The approval step id specified the approval flow step that this approver belongs to",
			},
			"role_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The approver role id",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertApprovalWorkflowApproverToPlumbing(d *schema.ResourceData) *sdm.ApprovalWorkflowApprover {
	return &sdm.ApprovalWorkflowApprover{
		ID:             d.Id(),
		AccountID:      convertStringToPlumbing(d.Get("account_id")),
		ApprovalFlowID: convertStringToPlumbing(d.Get("approval_flow_id")),
		ApprovalStepID: convertStringToPlumbing(d.Get("approval_step_id")),
		RoleID:         convertStringToPlumbing(d.Get("role_id")),
	}
}

func resourceApprovalWorkflowApproverCreate(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertApprovalWorkflowApproverToPlumbing(d)
	resp, err := cc.ApprovalWorkflowApprovers().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create ApprovalWorkflowApprover: %w", err)
	}
	d.SetId(resp.ApprovalWorkflowApprover.ID)
	v := resp.ApprovalWorkflowApprover
	d.Set("account_id", (v.AccountID))
	d.Set("approval_flow_id", (v.ApprovalFlowID))
	d.Set("approval_step_id", (v.ApprovalStepID))
	d.Set("role_id", (v.RoleID))
	return nil
}

func resourceApprovalWorkflowApproverRead(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	localVersion := convertApprovalWorkflowApproverToPlumbing(d)
	_ = localVersion
	resp, err := cc.ApprovalWorkflowApprovers().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read ApprovalWorkflowApprover %s: %w", d.Id(), err)
	}
	v := resp.ApprovalWorkflowApprover
	d.Set("account_id", (v.AccountID))
	d.Set("approval_flow_id", (v.ApprovalFlowID))
	d.Set("approval_step_id", (v.ApprovalStepID))
	d.Set("role_id", (v.RoleID))
	return nil
}
func resourceApprovalWorkflowApproverDelete(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	var errNotFound *sdm.NotFoundError
	_, err := cc.ApprovalWorkflowApprovers().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}