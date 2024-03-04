package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMApprovalWorkflowApprover_Get(t *testing.T) {
	initAcceptanceTest(t)
	approvers, err := createApprovalWorkflowApproversWithPrifix("approval-workflow-approver-get", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow approver: ", err)
	}
	approver := approvers[0]
	flowID := approver.ApprovalFlowID
	stepID := approver.ApprovalStepID

	dsName := randomWithPrefix("af-test-query")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowApproverDataSourceGetFilterConfig(dsName, flowID, stepID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow_approver."+dsName, "approval_workflow_approvers.0.approval_flow_id", flowID),
					resource.TestCheckResourceAttr("data.sdm_approval_workflow_approver."+dsName, "approval_workflow_approvers.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflowApprover_GetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	approvers, err := createApprovalWorkflowApproversWithPrifix("multi-approver-test", 3)
	if err != nil {
		t.Fatal("failed to create approval workflow approver: ", err)
	}

	flowID := approvers[0].ApprovalFlowID
	stepID := approvers[0].ApprovalStepID

	dsName := randomWithPrefix("multi-approver-test")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowApproverDataSourceGetFilterConfig(dsName, flowID, stepID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow_approver."+dsName, "approval_workflow_approvers.#", "3"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflowApprover_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createApprovalWorkflowApproversWithPrifix("none-test", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow approver: ", err)
	}

	flowID := "af-000333000"
	stepID := "afs-000333000"

	dsName := randomWithPrefix("none-test")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowApproverDataSourceGetFilterConfig(dsName, flowID, stepID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow_approver."+dsName, "approval_workflow_approvers.#", "0"),
				),
			},
		},
	})
}

func testAccSDMApprovalWorkflowApproverDataSourceGetFilterConfig(approverSourceName, flowID, stepID string) string {
	return fmt.Sprintf(`
	data "sdm_approval_workflow_approver" "%s" {
		approval_flow_id = "%s"
		approval_step_id = "%s"
	}`, approverSourceName, flowID, stepID)
}
