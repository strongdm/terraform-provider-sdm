// Code generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func dataSourceWorkflowApprover() *schema.Resource {
	return &schema.Resource{
		ReadContext: wrapCrudOperation(dataSourceWorkflowApproverList),
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"approver_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The approver id.",
			},
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the WorkflowApprover.",
			},
			"workflow_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The workflow id.",
			},
			"workflow_approvers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"approver_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The approver id.",
						},
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique identifier of the WorkflowApprover.",
						},
						"workflow_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The workflow id.",
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

func convertWorkflowApproverFilterToPlumbing(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOkExists("approver_id"); ok {
		filter += "approverid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("workflow_id"); ok {
		filter += "workflowid:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceWorkflowApproverList(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	filter, args := convertWorkflowApproverFilterToPlumbing(d)
	resp, err := cc.WorkflowApprovers().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list WorkflowApprovers %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]entity, 0)
	for resp.Next() {
		v := resp.Value()
		ids = append(ids, v.ID)
		output = append(output,
			entity{
				"approver_id": (v.ApproverID),
				"id":          (v.ID),
				"workflow_id": (v.WorkflowID),
			})
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("workflow_approvers", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("WorkflowApprover" + filter + fmt.Sprint(args...))
	return nil
}
