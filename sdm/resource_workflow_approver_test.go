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
	resource.AddTestSweepers("sdm_workflow_approver", &resource.Sweeper{
		Name: "sdm_workflow_approver",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.WorkflowApprovers().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing workflow approvers: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.WorkflowApprovers().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting workflow approver %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing workflow approvers %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMWorkflowApprover_Create(t *testing.T) {
	initAcceptanceTest(t)
	rs1Name := randomWithPrefix("test-workflow-approver1")
	rs2Name := randomWithPrefix("test-workflow-approver2")
	workflowName := randomWithPrefix("test-workflow")
	approver1Name := randomWithPrefix("test-approver1")
	approver2Name := randomWithPrefix("test-approver2")
	workflowID := ""
	approver1ID := ""
	approver2ID := ""
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowApproversConfig(rs1Name, rs2Name, workflowName, approver1Name, approver2Name),
				Check: resource.ComposeTestCheckFunc(
					// Ensure and capture the mapping IDs
					func(s *terraform.State) error {
						var err error
						workflowID, err = testCreatedID(s, "sdm_workflow", workflowName)
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
					// Ensure and check the WorkflowApprover creations
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_workflow_approver."+rs1Name, "workflow_id", workflowID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_workflow_approver."+rs1Name, "account_id", approver1ID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_workflow_approver."+rs2Name, "workflow_id", workflowID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_workflow_approver."+rs2Name, "role_id", approver2ID)(s); err != nil {
							return err
						}
						return nil
					},
					// Ensure the resources were actually created
					func(s *terraform.State) error {
						id1, err := testCreatedID(s, "sdm_workflow_approver", rs1Name)
						if err != nil {
							return err
						}
						id2, err := testCreatedID(s, "sdm_workflow_approver", rs2Name)
						if err != nil {
							return err
						}

						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						ids := map[string]string{id1: approver1ID, id2: approver2ID}
						for workflowApproverID, approverID := range ids {
							resp, err := client.WorkflowApprovers().Get(ctx, workflowApproverID)
							if err != nil {
								return err
							}
							if resp.WorkflowApprover.WorkflowID != workflowID {
								return fmt.Errorf("unexpected workflowID '%s', expected '%s'", resp.WorkflowApprover.WorkflowID, workflowID)
							}
							if workflowApproverID == id1 && resp.WorkflowApprover.AccountID != approverID {
								return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.WorkflowApprover.AccountID, approverID)
							}
							if workflowApproverID == id2 && resp.WorkflowApprover.RoleID != approverID {
								return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.WorkflowApprover.RoleID, approverID)
							}
						}
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_workflow_approver." + rs1Name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "sdm_workflow_approver." + rs2Name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createWorkflowApproverAccountsWithPrefix(prefix string, count int) ([]*sdm.WorkflowApprover, error) {
	approvers, err := createUsersWithPrefix(prefix, count)
	if err != nil {
		return nil, fmt.Errorf("failed to create approvers: %w", err)
	}
	workflows, err := createWorkflowsWithPrifix(prefix, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create workflow: %w", err)
	}
	workflowApprovers, err := addApproverAccountsToWorkflow(approvers, workflows[0])
	if err != nil {
		return nil, fmt.Errorf("failed to add approvers: %w", err)
	}
	return workflowApprovers, nil
}

func addApproverAccountsToWorkflow(approvers []sdm.Account, workflow *sdm.Workflow) ([]*sdm.WorkflowApprover, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	workflowApprovers := []*sdm.WorkflowApprover{}
	for _, a := range approvers {
		createResp, err := client.WorkflowApprovers().Create(ctx, &sdm.WorkflowApprover{
			WorkflowID: workflow.ID,
			AccountID:  a.GetID(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create workflow approver: %w", err)
		}
		workflowApprovers = append(workflowApprovers, createResp.WorkflowApprover)
	}
	return workflowApprovers, nil
}

func createWorkflowApproverRolesWithPrefix(prefix string, count int) ([]*sdm.WorkflowApprover, error) {
	approvers, err := createRolesWithPrefix(prefix, count)
	if err != nil {
		return nil, fmt.Errorf("failed to create approvers: %w", err)
	}
	workflows, err := createWorkflowsWithPrifix(prefix, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create workflow: %w", err)
	}
	workflowApprovers, err := addApproverRolesToWorkflow(approvers, workflows[0])
	if err != nil {
		return nil, fmt.Errorf("failed to add approvers: %w", err)
	}
	return workflowApprovers, nil
}

func addApproverRolesToWorkflow(approvers []*sdm.Role, workflow *sdm.Workflow) ([]*sdm.WorkflowApprover, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	workflowApprovers := []*sdm.WorkflowApprover{}
	for _, a := range approvers {
		createResp, err := client.WorkflowApprovers().Create(ctx, &sdm.WorkflowApprover{
			WorkflowID: workflow.ID,
			RoleID:     a.ID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create workflow approver: %w", err)
		}
		workflowApprovers = append(workflowApprovers, createResp.WorkflowApprover)
	}
	return workflowApprovers, nil
}

func testAccSDMWorkflowApproversConfig(name1, name2, workflowName, approver1Name, approver2Name string) string {
	return fmt.Sprintf(
		`
	resource "sdm_workflow" "%[1]s" {
		name = "%[2]s"
		weight = "10"
	}

	resource "sdm_account" "%[3]s" {
		user {
			first_name = "%[4]s"
			last_name = "%[5]s"
			email = "%[6]s"
		}
	}

	resource "sdm_role" "%[7]s" {
		name = "%[8]s"
	}

	resource "sdm_workflow_approver" "%[9]s" {
		workflow_id = sdm_workflow.%[1]s.id
		account_id = sdm_account.%[3]s.id
	}

	resource "sdm_workflow_approver" "%[10]s" {
		workflow_id = sdm_workflow.%[1]s.id
		role_id = sdm_role.%[7]s.id
	}
	`,
		workflowName,
		randomWithPrefix("workflow-name"),
		approver1Name,
		randomWithPrefix("first-name"),
		randomWithPrefix("last-name"),
		randomWithPrefix("email"),
		approver2Name,
		randomWithPrefix("name"),
		name1,
		name2,
	)
}
