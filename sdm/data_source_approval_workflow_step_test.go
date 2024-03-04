package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMApprovalWorkflowStep_Get(t *testing.T) {
	initAcceptanceTest(t)
	step, err := createApprovalWorkflowStepWithPrifix("approval-workflow-step-get")
	if err != nil {
		t.Fatal("failed to create approval workflow step: ", err)
	}
	flowId := step.ApprovalFlowID

	dsName := randomWithPrefix("af-test-query")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowStepDataSourceGetFilterConfig(dsName, flowId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow_step."+dsName, "approval_workflow_steps.0.approval_flow_id", flowId),
					resource.TestCheckResourceAttr("data.sdm_approval_workflow_step."+dsName, "approval_workflow_steps.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflowStep_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createApprovalWorkflowStepWithPrifix("none-test")
	if err != nil {
		t.Fatal("failed to create approval workflow step: ", err)
	}

	flowID := "af-000333000"

	dsName := randomWithPrefix("none-test")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowStepDataSourceGetFilterConfig(dsName, flowID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow_step."+dsName, "approval_workflow_steps.#", "0"),
				),
			},
		},
	})
}

func testAccSDMApprovalWorkflowStepDataSourceGetFilterConfig(stepSourceName, flowId string) string {
	return fmt.Sprintf(`
	data "sdm_approval_workflow_step" "%s" {
		approval_flow_id = "%s"
	}`, stepSourceName, flowId)
}
