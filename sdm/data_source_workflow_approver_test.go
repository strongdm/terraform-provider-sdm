package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMWorkflowApprovers_GetAccounts(t *testing.T) {
	initAcceptanceTest(t)

	workflowApprovers, err := createWorkflowApproverAccountsWithPrefix("wa-test-get-account", 1)
	if err != nil {
		t.Fatal("failed to create workflow approvers: ", err)
	}

	workflowID := workflowApprovers[0].WorkflowID
	approverID := workflowApprovers[0].AccountID
	workflowApproverName := randomWithPrefix("wa-test-get-account")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowApproverAccountGetFilterConfig(workflowApproverName, workflowID, approverID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.0.workflow_id", workflowID),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.0.account_id", approverID),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMWorkflowApprovers_GetRoles(t *testing.T) {
	initAcceptanceTest(t)

	workflowApprovers, err := createWorkflowApproverRolesWithPrefix("wa-test-get-role", 1)
	if err != nil {
		t.Fatal("failed to create workflow approvers: ", err)
	}

	workflowID := workflowApprovers[0].WorkflowID
	approverID := workflowApprovers[0].RoleID
	workflowApproverName := randomWithPrefix("wa-test-get-role")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowApproverRoleGetFilterConfig(workflowApproverName, workflowID, approverID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.0.workflow_id", workflowID),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.0.role_id", approverID),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMWorkflowApprovers_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createWorkflowApproverAccountsWithPrefix("wa-test-get-none", 1)
	if err != nil {
		t.Fatal("failed to create workflow approvers: ", err)
	}

	approverID := "a-000333000"
	workflowApproverName := randomWithPrefix("wa-test-get-none")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowApproverAccountGetFilterConfig(workflowApproverName, "", approverID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "ids.#", "0"),
				),
			},
		},
	})
}

func TestAccSDMWorkflowApprovers_GetMulti(t *testing.T) {
	initAcceptanceTest(t)

	workflowApprovers, err := createWorkflowApproverAccountsWithPrefix("wa-test-get-multi", 3)
	if err != nil {
		t.Fatal("failed to create workflow approvers: ", err)
	}
	workflowID := workflowApprovers[0].WorkflowID

	workflowApproverName := randomWithPrefix("wa-test-get-multi")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowApproverAccountGetFilterConfig(workflowApproverName, workflowID, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "workflow_approvers.#", "3"),
					resource.TestCheckResourceAttr("data.sdm_workflow_approver."+workflowApproverName, "ids.#", "3"),
				),
			},
		},
	})
}

func testAccSDMWorkflowApproverAccountGetFilterConfig(workflowApproverDataSourceName, workflowID, approverID string) string {
	return fmt.Sprintf(
		`
	data "sdm_workflow_approver" "%s" {
		workflow_id = "%s"
		account_id = "%s"
	}`,
		workflowApproverDataSourceName,
		workflowID,
		approverID,
	)
}

func testAccSDMWorkflowApproverRoleGetFilterConfig(workflowApproverDataSourceName, workflowID, approverID string) string {
	return fmt.Sprintf(
		`
	data "sdm_workflow_approver" "%s" {
		workflow_id = "%s"
		role_id = "%s"
	}`,
		workflowApproverDataSourceName,
		workflowID,
		approverID,
	)
}
