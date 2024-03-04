package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func init() {
	resource.AddTestSweepers("sdm_approval_workflow_step", &resource.Sweeper{
		Name:         "sdm_approval_workflow_step",
		Dependencies: []string{"sdm_approval_workflow_approver"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.ApprovalWorkflowSteps().List(ctx, "")
			if err != nil {
				return fmt.Errorf("Error listing approval workflow steps: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.ApprovalWorkflowSteps().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting approval workflow step %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing approval workflow steps %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMApprovalWorkflowStep_Create(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("af-create")
	workflowName := randomWithPrefix("af-create")
	stepName := randomWithPrefix("afs-create")
	workflowID := ""
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowConfig(resourceName, workflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+resourceName, "name", workflowName),
					doubleCheckApprovalWorkflow(resourceName, workflowName),
				),
			},
			{
				ResourceName:      "sdm_approval_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMApprovalWorkflowStepConfig(resourceName, workflowName, stepName),
				Check: resource.ComposeTestCheckFunc(
					// Ensure and capture the mapping IDs
					func(s *terraform.State) error {
						var err error
						workflowID, err = testCreatedID(s, "sdm_approval_workflow", resourceName)
						if err != nil {
							return err
						}
						return nil
					},
					// Ensure and check the ApprovalWorkflowStep creations
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_approval_workflow_step."+stepName, "approval_flow_id", workflowID)(s); err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {
						stepID, err := testCreatedID(s, "sdm_approval_workflow_step", stepName)
						if err != nil {
							return err
						}

						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()

						resp, err := client.ApprovalWorkflowSteps().Get(ctx, stepID)
						if err != nil {
							return err
						}
						if resp.ApprovalWorkflowStep.ApprovalFlowID != workflowID {
							return fmt.Errorf("unexpected approval workflowID '%s', expected '%s'", resp.ApprovalWorkflowStep.ApprovalFlowID, workflowID)
						}
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_approval_workflow_step." + stepName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createApprovalWorkflowStepWithPrifix(prefix string) (*sdm.ApprovalWorkflowStep, error) {
	workflows, err := createApprovalWorkflowsWithPrifix(prefix, "manual", 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create approval workflow: %w", err)
	}
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	createResp, err := client.ApprovalWorkflowSteps().Create(ctx, &sdm.ApprovalWorkflowStep{
		ApprovalFlowID: workflows[0].ID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create approval workflow step: %w", err)
	}
	return createResp.ApprovalWorkflowStep, nil
}

func testAccSDMApprovalWorkflowStepConfig(workflowResourceName, workflowName, stepResourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_approval_workflow_step" "%[3]s" {
		approval_flow_id = sdm_approval_workflow.%[1]s.id
	}

	resource "sdm_approval_workflow" "%[1]s" {
		name = "%[2]s"
		approval_mode = "manual"
	}
	`, workflowResourceName, workflowName, stepResourceName)
}
