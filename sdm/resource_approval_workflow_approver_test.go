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
	resource.AddTestSweepers("sdm_approval_workflow_approver", &resource.Sweeper{
		Name: "sdm_approval_workflow_approver",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.ApprovalWorkflowApprovers().List(ctx, "")
			if err != nil {
				return fmt.Errorf("Error listing approval workflow approvers: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.ApprovalWorkflowApprovers().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting approval workflow approver %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing approval workflow approvers %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMApprovalWorkflowApprover_Create(t *testing.T) {
	initAcceptanceTest(t)
	workflowResourceName := randomWithPrefix("af-create")
	workflowName := randomWithPrefix("af-create")
	stepName := randomWithPrefix("afs-create")
	approver1Name := randomWithPrefix("a-create")
	approver2Name := randomWithPrefix("r-create")
	flowApprover1Name := randomWithPrefix("afa-create-account")
	flowApprover2Name := randomWithPrefix("afa-create-role")
	workflowID := ""
	stepID := ""
	approver1ID := ""
	approver2ID := ""
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowConfig(workflowResourceName, workflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+workflowResourceName, "name", workflowName),
					doubleCheckApprovalWorkflow(workflowResourceName, workflowName),
				),
			},
			{
				ResourceName:      "sdm_approval_workflow." + workflowResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMApprovalWorkflowStepConfig(workflowResourceName, workflowName, stepName),
				Check: resource.ComposeTestCheckFunc(
					// Ensure and capture the mapping IDs
					func(s *terraform.State) error {
						var err error
						workflowID, err = testCreatedID(s, "sdm_approval_workflow", workflowResourceName)
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
			{
				Config: testAccSDMApprovalWorkflowApproverConfig(workflowResourceName, workflowName, stepName, approver1Name, approver2Name, flowApprover1Name, flowApprover2Name),
				Check: resource.ComposeTestCheckFunc(
					// Ensure and capture the mapping IDs
					func(s *terraform.State) error {
						var err error
						workflowID, err = testCreatedID(s, "sdm_approval_workflow", workflowResourceName)
						if err != nil {
							return err
						}
						stepID, err = testCreatedID(s, "sdm_approval_workflow_step", stepName)
						if err != nil {
							return err
						}
						approver1ID, err = testCreatedID(s, "sdm_account", approver1Name)
						if err != nil {
							return err
						}
						approver2ID, err = testCreatedID(s, "sdm_role", approver2Name)
						if err != nil {
							return err
						}
						return nil
					},
					// Ensure and check the ApprovalWorkflowStep creations
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_approval_workflow_approver."+flowApprover1Name, "approval_flow_id", workflowID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_approval_workflow_approver."+flowApprover1Name, "account_id", approver1ID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_approval_workflow_approver."+flowApprover1Name, "approval_step_id", stepID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_approval_workflow_approver."+flowApprover2Name, "approval_flow_id", workflowID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_approval_workflow_approver."+flowApprover2Name, "role_id", approver2ID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_approval_workflow_approver."+flowApprover2Name, "approval_step_id", stepID)(s); err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {
						id1, err := testCreatedID(s, "sdm_approval_workflow_approver", flowApprover1Name)
						if err != nil {
							return err
						}
						id2, err := testCreatedID(s, "sdm_approval_workflow_approver", flowApprover2Name)
						if err != nil {
							return err
						}

						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						ids := map[string]string{id1: approver1ID, id2: approver2ID}
						for workflowApproverID, approverID := range ids {
							resp, err := client.ApprovalWorkflowApprovers().Get(ctx, workflowApproverID)
							if err != nil {
								return err
							}
							if resp.ApprovalWorkflowApprover.ApprovalFlowID != workflowID {
								return fmt.Errorf("unexpected approval workflowID '%s', expected '%s'", resp.ApprovalWorkflowApprover.ApprovalFlowID, workflowID)
							}
							if workflowApproverID == id1 && resp.ApprovalWorkflowApprover.AccountID != approverID {
								return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.ApprovalWorkflowApprover.AccountID, approverID)
							}
							if workflowApproverID == id2 && resp.ApprovalWorkflowApprover.RoleID != approverID {
								return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.ApprovalWorkflowApprover.RoleID, approverID)
							}
						}
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_approval_workflow_approver." + flowApprover1Name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "sdm_approval_workflow_approver." + flowApprover2Name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createApprovalWorkflowApproversWithPrifix(prefix string, count int) ([]*sdm.ApprovalWorkflowApprover, error) {
	approvers, err := createUsersWithPrefix(prefix, count)
	if err != nil {
		return nil, fmt.Errorf("failed to create approvers: %w", err)
	}

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
	step := createResp.ApprovalWorkflowStep

	approvalWorkflowApprovers := []*sdm.ApprovalWorkflowApprover{}
	for i := 0; i < count; i++ {
		createResp, err := client.ApprovalWorkflowApprovers().Create(ctx, &sdm.ApprovalWorkflowApprover{
			ApprovalFlowID: workflows[0].ID,
			ApprovalStepID: step.ID,
			AccountID:      approvers[i].GetID(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create approval workflow approver: %w", err)
		}
		approvalWorkflowApprovers = append(approvalWorkflowApprovers, createResp.ApprovalWorkflowApprover)
	}
	return approvalWorkflowApprovers, nil
}

func testAccSDMApprovalWorkflowApproverConfig(workflowResourceName, workflowName, stepResourceName, accountName, roleName, approver1Name, approver2Name string) string {
	return fmt.Sprintf(`
	resource "sdm_approval_workflow" "%[1]s" {
		name = "%[2]s"
		approval_mode = "manual"
	}

	resource "sdm_approval_workflow_step" "%[3]s" {
		approval_flow_id = sdm_approval_workflow.%[1]s.id
	}

	resource "sdm_account" "%[4]s" {
		user {
			first_name = "%[5]s"
			last_name = "%[6]s"
			email = "%[7]s"
		}
	}

	resource "sdm_role" "%[8]s" {
		name = "%[9]s"
	}

	resource "sdm_approval_workflow_approver" "%[10]s" {
		approval_flow_id = sdm_approval_workflow.%[1]s.id
		approval_step_id = sdm_approval_workflow_step.%[3]s.id
		account_id = sdm_account.%[4]s.id
	}

	resource "sdm_approval_workflow_approver" "%[11]s" {
		approval_flow_id = sdm_approval_workflow.%[1]s.id
		approval_step_id = sdm_approval_workflow_step.%[3]s.id
		role_id = sdm_role.%[8]s.id
	}
	`,
		workflowResourceName,
		workflowName,
		stepResourceName,
		accountName,
		randomWithPrefix("first-name"),
		randomWithPrefix("last-name"),
		randomWithPrefix("email"),
		roleName,
		randomWithPrefix("name"),
		approver1Name,
		approver2Name,
	)
}
