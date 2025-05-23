// Code generated by protogen. DO NOT EDIT.
// See pkg/api/v1/templates/terraform-provider-sdm/sdm/*_test.go.tpl

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

func TestAccSDMApprovalWorkflow_CreateNested(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("af-create")
	approvalWorkflowName := randomWithPrefix("af-create")
	accountName := randomWithPrefix("af-create-account")
	roleName := randomWithPrefix("af-create-role")
	approvalMode := "manual"
	step1Quantifier := "all"
	step1SkipAfter := "0s"
	step2Quantifier := "any"
	step2SkipAfter := "1h0m0s"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowNestedConfig(resourceName, approvalWorkflowName, approvalMode, step1Quantifier, step1SkipAfter, accountName, step2Quantifier, step2SkipAfter, roleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+resourceName, "name", approvalWorkflowName),
					doubleCheckApprovalWorkflow(resourceName, approvalWorkflowName),
					func(s *terraform.State) error {
						var err error
						_, err = testCreatedID(s, "sdm_account", accountName)
						if err != nil {
							return err
						}
						_, err = testCreatedID(s, "sdm_role", roleName)
						if err != nil {
							return err
						}
						return nil
					},
					checkApprovalSteps(resourceName, approvalMode, step1Quantifier, step1SkipAfter, accountName, step2Quantifier, step2SkipAfter, roleName),
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

func TestAccSDMApprovalWorkflow_UpdateNested(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("af-update")
	approvalWorkflowName := randomWithPrefix("af-update")
	accountName := randomWithPrefix("af-update-account")
	roleName := randomWithPrefix("af-update-role")
	approvalMode := "manual"
	step1Quantifier := "all"
	step1SkipAfter := "0s"
	step2Quantifier := "any"
	step2SkipAfter := "1h0m0s"

	accountNameUpdated := randomWithPrefix("af-update-account")
	roleNameUpdated := randomWithPrefix("af-update-role")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccSDMApprovalWorkflowNestedConfig(resourceName, approvalWorkflowName, approvalMode, step1Quantifier, step1SkipAfter, accountName, step2Quantifier, step2SkipAfter, roleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+resourceName, "name", approvalWorkflowName),
					doubleCheckApprovalWorkflow(resourceName, approvalWorkflowName),
					func(s *terraform.State) error {
						var err error
						_, err = testCreatedID(s, "sdm_account", accountName)
						if err != nil {
							return err
						}
						_, err = testCreatedID(s, "sdm_role", roleName)
						if err != nil {
							return err
						}
						return nil
					},
					checkApprovalSteps(resourceName, approvalMode, step1Quantifier, step1SkipAfter, accountName, step2Quantifier, step2SkipAfter, roleName),
				),
			},
			{
				ResourceName:      "sdm_approval_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMApprovalWorkflowUpdateNestedConfig(resourceName, approvalWorkflowName, approvalMode, step2Quantifier, step2SkipAfter, accountNameUpdated, step1Quantifier, step1SkipAfter, roleNameUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_approval_workflow."+resourceName, "name", approvalWorkflowName),
					doubleCheckApprovalWorkflow(resourceName, approvalWorkflowName),
					checkApprovalStepsUpdate(resourceName, approvalMode, step2Quantifier, step2SkipAfter, accountNameUpdated, step1Quantifier, step1SkipAfter, roleNameUpdated),
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

func createApprovalWorkflowsWithPrefix(prefix, approvalMode string, count int) ([]*sdm.ApprovalWorkflow, error) {
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

func createApprovalWorkflowsNestedWithPrefix(prefix string, count int) ([]*sdm.ApprovalWorkflow, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	users, err := createUsersWithPrefix(prefix, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to create users for approvers: %w", err)
	}
	roles, err := createRolesWithPrefix(prefix, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create roles for approvers: %w", err)
	}
	approvalWorkflows := []*sdm.ApprovalWorkflow{}
	for i := 0; i < count; i++ {
		createResp, err := client.ApprovalWorkflows().Create(ctx, &sdm.ApprovalWorkflow{
			Name:         randomWithPrefix(prefix),
			ApprovalMode: "manual",
			ApprovalWorkflowSteps: []*sdm.ApprovalFlowStep{
				{
					Quantifier: "any",
					SkipAfter:  time.Hour,
					Approvers: []*sdm.ApprovalFlowApprover{
						{AccountID: users[0].GetID()},
						{RoleID: roles[0].ID},
					},
				},
				{
					Quantifier: "all",
					Approvers: []*sdm.ApprovalFlowApprover{
						{AccountID: users[1].GetID()},
						{AccountID: users[2].GetID()},
						{Reference: sdm.ApproverReferenceManagerOfRequester},
						{Reference: sdm.ApproverReferenceManagerOfManagerOfRequester},
					},
				},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create approval_workflow: %w", err)
		}
		approvalWorkflows = append(approvalWorkflows, createResp.ApprovalWorkflow)
	}
	return approvalWorkflows, nil
}

func testAccSDMApprovalWorkflowNestedConfig(resourceName, workflowName, approvalMode, step1Quantifier, step1SkipAfter, step1AccountName, step2Quantifier, step2SkipAfter, step2RoleName string) string {
	return fmt.Sprintf(`
	resource "sdm_account" "%[6]s" {
		user {
			first_name = "FIRST"
			last_name = "LAST"
			email = "%[6]s@sdm.com"
		}
	}

	resource "sdm_role" "%[9]s" {
		name = "role-%[9]s"
	}

	resource "sdm_approval_workflow" "%[1]s" {
		name = "%[2]s"
		approval_mode = "%[3]s"
	 	approval_step {
			quantifier = "%[4]s"
			skip_after = "%[5]s"
			approvers {
				account_id = sdm_account.%[6]s.id
			}
		}
		approval_step {
			quantifier = "%[7]s"
			skip_after = "%[8]s"
			approvers {
				role_id = sdm_role.%[9]s.id
			}
			approvers {
				reference = "manager-of-requester"
			}
			approvers {
				reference = "manager-of-manager-of-requester"
			}
		}
	}`, resourceName, workflowName, approvalMode, step1Quantifier, step1SkipAfter, step1AccountName, step2Quantifier, step2SkipAfter, step2RoleName)
}

func testAccSDMApprovalWorkflowUpdateNestedConfig(resourceName, workflowName, approvalMode, step1Quantifier, step1SkipAfter, step1AccountName, step2Quantifier, step2SkipAfter, step2RoleName string) string {
	return fmt.Sprintf(`
	resource "sdm_account" "%[6]s" {
		user {
			first_name = "FIRST"
			last_name = "LAST"
			email = "%[6]s@sdm.com"
		}
	}

	resource "sdm_role" "%[9]s" {
		name = "role-%[9]s"
	}

	resource "sdm_approval_workflow" "%[1]s" {
		name = "%[2]s"
		approval_mode = "%[3]s"
	 	approval_step {
			quantifier = "%[4]s"
			skip_after = "%[5]s"
			approvers {
				account_id = sdm_account.%[6]s.id
			}
			approvers {
				reference = "manager-of-requester"
			}
			approvers {
				reference = "manager-of-manager-of-requester"
			}
		}
		approval_step {
			quantifier = "%[7]s"
			skip_after = "%[8]s"
			approvers {
				role_id = sdm_role.%[9]s.id
			}
		}
	}`, resourceName, workflowName, approvalMode, step1Quantifier, step1SkipAfter, step1AccountName, step2Quantifier, step2SkipAfter, step2RoleName)
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

func checkApprovalSteps(resourceName, approvalMode, step1Quantifier, step1SkipAfter, step1AccountName, step2Quantifier, step2SkipAfter, step2RoleName string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		id, err := testCreatedID(s, "sdm_approval_workflow", resourceName)
		if err != nil {
			return err
		}
		step1ApproverID, err := testCreatedID(s, "sdm_account", step1AccountName)
		if err != nil {
			return err
		}
		step2ApproverID, err := testCreatedID(s, "sdm_role", step2RoleName)
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
		if resp.ApprovalWorkflow.ApprovalMode != approvalMode {
			return fmt.Errorf("unexpected approval mode '%s', expected '%s'", resp.ApprovalWorkflow.ApprovalMode, approvalMode)
		}
		approvalSteps := resp.ApprovalWorkflow.ApprovalWorkflowSteps
		if len(approvalSteps) != 2 {
			return fmt.Errorf("expected 2 approval steps, got %v", len(approvalSteps))
		}
		step1 := approvalSteps[0]
		if step1.Quantifier != step1Quantifier {
			return fmt.Errorf("unexpected quantifier '%s', expected '%s'", step1.Quantifier, step1Quantifier)
		}

		step1SkipAfterDuration, err := time.ParseDuration(step1SkipAfter)
		if err != nil {
			return fmt.Errorf("error parsing duration")
		}
		if step1.SkipAfter != step1SkipAfterDuration {
			return fmt.Errorf("unexpected skip_after '%v', expected '%v'", step1.SkipAfter, step1SkipAfterDuration)
		}
		if len(step1.Approvers) != 1 {
			return fmt.Errorf("expected 1 approver, got %v", len(step1.Approvers))
		}
		if step1.Approvers[0].AccountID != step1ApproverID {
			return fmt.Errorf("unexpected approver id '%v', expected '%v'", step1.Approvers[0].AccountID, step1ApproverID)
		}

		step2 := approvalSteps[1]
		if step2.Quantifier != step2Quantifier {
			return fmt.Errorf("unexpected quantifier '%s', expected '%s'", resp.ApprovalWorkflow.ApprovalMode, step2.Quantifier)
		}

		step2SkipAfterDuration, err := time.ParseDuration(step2SkipAfter)
		if err != nil {
			return fmt.Errorf("error parsing duration")
		}
		if step2.SkipAfter != step2SkipAfterDuration {
			return fmt.Errorf("unexpected skip_after '%v', expected '%v'", step2.SkipAfter, step2SkipAfterDuration)
		}
		if len(step2.Approvers) != 3 {
			return fmt.Errorf("expected 3 approvers, got %v", len(step2.Approvers))
		}
		if step2.Approvers[0].RoleID != step2ApproverID {
			return fmt.Errorf("unexpected approver id '%v', expected '%v'", step2.Approvers[0].RoleID, step2ApproverID)
		}
		if step2.Approvers[1].Reference != sdm.ApproverReferenceManagerOfRequester {
			return fmt.Errorf("unexpected approver reference '%v', expected '%v'", step2.Approvers[1].Reference, sdm.ApproverReferenceManagerOfRequester)
		}
		if step2.Approvers[2].Reference != sdm.ApproverReferenceManagerOfManagerOfRequester {
			return fmt.Errorf("unexpected approver reference '%v', expected '%v'", step2.Approvers[2].Reference, sdm.ApproverReferenceManagerOfManagerOfRequester)
		}

		return nil
	}
}

func checkApprovalStepsUpdate(resourceName, approvalMode, step1Quantifier, step1SkipAfter, step1AccountName, step2Quantifier, step2SkipAfter, step2RoleName string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		id, err := testCreatedID(s, "sdm_approval_workflow", resourceName)
		if err != nil {
			return err
		}
		step1ApproverID, err := testCreatedID(s, "sdm_account", step1AccountName)
		if err != nil {
			return err
		}
		step2ApproverID, err := testCreatedID(s, "sdm_role", step2RoleName)
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
		if resp.ApprovalWorkflow.ApprovalMode != approvalMode {
			return fmt.Errorf("unexpected approval mode '%s', expected '%s'", resp.ApprovalWorkflow.ApprovalMode, approvalMode)
		}
		approvalSteps := resp.ApprovalWorkflow.ApprovalWorkflowSteps
		if len(approvalSteps) != 2 {
			return fmt.Errorf("expected 2 approval steps, got %v", len(approvalSteps))
		}
		step1 := approvalSteps[0]
		if step1.Quantifier != step1Quantifier {
			return fmt.Errorf("unexpected quantifier '%s', expected '%s'", step1.Quantifier, step1Quantifier)
		}

		step1SkipAfterDuration, err := time.ParseDuration(step1SkipAfter)
		if err != nil {
			return fmt.Errorf("error parsing duration")
		}
		if step1.SkipAfter != step1SkipAfterDuration {
			return fmt.Errorf("unexpected skip_after '%v', expected '%v'", step1.SkipAfter, step1SkipAfterDuration)
		}
		if len(step1.Approvers) != 3 {
			return fmt.Errorf("expected 3 approver, got %v", len(step1.Approvers))
		}
		if step1.Approvers[0].AccountID != step1ApproverID {
			return fmt.Errorf("unexpected approver id '%v', expected '%v'", step1.Approvers[0].AccountID, step1ApproverID)
		}
		if step1.Approvers[1].Reference != sdm.ApproverReferenceManagerOfRequester {
			return fmt.Errorf("unexpected approver reference '%v', expected '%v'", step1.Approvers[1].Reference, sdm.ApproverReferenceManagerOfRequester)
		}
		if step1.Approvers[2].Reference != sdm.ApproverReferenceManagerOfManagerOfRequester {
			return fmt.Errorf("unexpected approver reference '%v', expected '%v'", step1.Approvers[2].Reference, sdm.ApproverReferenceManagerOfManagerOfRequester)
		}

		step2 := approvalSteps[1]
		if step2.Quantifier != step2Quantifier {
			return fmt.Errorf("unexpected quantifier '%s', expected '%s'", resp.ApprovalWorkflow.ApprovalMode, step2.Quantifier)
		}

		step2SkipAfterDuration, err := time.ParseDuration(step2SkipAfter)
		if err != nil {
			return fmt.Errorf("error parsing duration")
		}
		if step2.SkipAfter != step2SkipAfterDuration {
			return fmt.Errorf("unexpected skip_after '%v', expected '%v'", step2.SkipAfter, step2SkipAfterDuration)
		}
		if len(step2.Approvers) != 1 {
			return fmt.Errorf("expected 1 approvers, got %v", len(step2.Approvers))
		}
		if step2.Approvers[0].RoleID != step2ApproverID {
			return fmt.Errorf("unexpected approver id '%v', expected '%v'", step2.Approvers[0].RoleID, step2ApproverID)
		}

		return nil
	}
}
