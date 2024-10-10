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
	resource.AddTestSweepers("sdm_approval_workflow", &resource.Sweeper{
		Name:         "sdm_approval_workflow",
		Dependencies: []string{"sdm_approval_workflow_step", "sdm_approval_workflow_approver"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.ApprovalWorkflows().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing approval workflows: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.ApprovalWorkflows().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting approval workflow %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing approval workflows %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMApprovalWorkflow_Create(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("af-create")
	approvalWorkflowName := randomWithPrefix("af-create")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowConfig(resourceName, approvalWorkflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+resourceName, "name", approvalWorkflowName),
					doubleCheckApprovalWorkflow(resourceName, approvalWorkflowName),
				),
			},
			{
				ResourceName:      "sdm_approval_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_Update(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("wf-create")
	approvalWorkflowName := randomWithPrefix("wf-create")
	updatedApprovalWorkflowName := randomWithPrefix("wf-create-updated")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowConfig(resourceName, approvalWorkflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+resourceName, "name", approvalWorkflowName),
					doubleCheckApprovalWorkflow(resourceName, approvalWorkflowName),
				),
			},
			{
				ResourceName:      "sdm_approval_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMApprovalWorkflowConfig(resourceName, updatedApprovalWorkflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+resourceName, "name", updatedApprovalWorkflowName),
					doubleCheckApprovalWorkflow(resourceName, updatedApprovalWorkflowName),
				),
			},
			{
				ResourceName:      "sdm_approval_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createApprovalWorkflowsWithPrifix(prefix, approvalMode string, count int) ([]*sdm.ApprovalWorkflow, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	approvalWorkflows := []*sdm.ApprovalWorkflow{}
	for i := 0; i < count; i++ {
		createResp, err := client.ApprovalWorkflows().Create(ctx, &sdm.ApprovalWorkflow{
			Name:         randomWithPrefix(prefix),
			ApprovalMode: approvalMode,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create approval_workflow: %w", err)
		}
		approvalWorkflows = append(approvalWorkflows, createResp.ApprovalWorkflow)
	}
	return approvalWorkflows, nil
}

func testAccSDMApprovalWorkflowConfig(resourceName, workflowName string) string {
	return fmt.Sprintf(`
	resource "sdm_approval_workflow" "%s" {
		name = "%s"
		approval_mode = "automatic"
	}`, resourceName, workflowName)
}

func doubleCheckApprovalWorkflow(resourceName, approvalWorkflowName string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		id, err := testCreatedID(s, "sdm_approval_workflow", resourceName)
		if err != nil {
			return err
		}
		client := testClient()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		resp, err := client.ApprovalWorkflows().Get(ctx, id)
		if err != nil {
			return fmt.Errorf("failed to get created resource: %w", err)
		}
		if resp.ApprovalWorkflow.Name != approvalWorkflowName {
			return fmt.Errorf("unexpected name '%s', expected '%s'", resp.ApprovalWorkflow.Name, approvalWorkflowName)
		}

		return nil
	}
}
