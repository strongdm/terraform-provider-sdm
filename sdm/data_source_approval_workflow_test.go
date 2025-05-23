// Code generated by protogen. DO NOT EDIT.
// See pkg/api/v1/templates/terraform-provider-sdm/sdm/*_test.go.tpl

package sdm

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMApprovalWorkflow_Get(t *testing.T) {
	initAcceptanceTest(t)
	workflows, err := createApprovalWorkflowsWithPrefix("approval-workflow-get-automatic", "automatic", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}
	workflow := workflows[0]

	dsName := randomWithPrefix("af-test-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "approval-workflow-get-automatic*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.0.name", workflow.Name),
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_GetNested(t *testing.T) {
	initAcceptanceTest(t)
	workflows, err := createApprovalWorkflowsNestedWithPrefix("approval-workflow-get-nested", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}
	workflow := workflows[0]

	dsName := randomWithPrefix("af-test-query-nested")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "approval-workflow-get-nested*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.0.name", workflow.Name),
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.0.approval_step.#", "2"),
					validateApprovalWorkflows(dsName, workflows),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_GetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createApprovalWorkflowsWithPrefix("multistep-test-get-multiple", "automatic", 2)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}
	_, err = createApprovalWorkflowsWithPrefix("nomatch", "automatic", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}

	dsName := randomWithPrefix("af-multistep-test-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "multistep-test-get-multiple*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_GetMultipleNested(t *testing.T) {
	initAcceptanceTest(t)
	workflows, err := createApprovalWorkflowsNestedWithPrefix("multistep-test-nested", 3)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}
	_, err = createApprovalWorkflowsNestedWithPrefix("nomatch", 2)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}

	dsName := randomWithPrefix("af-multistep-test-query-nested")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "multistep-test-nested*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "3"),
					validateApprovalWorkflows(dsName, workflows),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createApprovalWorkflowsWithPrefix("dontfind", "automatic", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}

	dsName := randomWithPrefix("af-none-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "zero*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "0"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_GetNoneNested(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createApprovalWorkflowsNestedWithPrefix("dontfind-nested", 3)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}

	dsName := randomWithPrefix("af-none-query-nested")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "zero*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "0"),
				),
			},
		},
	})
}

func testAccSDMApprovalWorkflowDataSourceGetFilterConfig(workflowDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_approval_workflow" "%s" {
		name = "%s"
	}`, workflowDataSourceName, nameFilter)
}

func SortApprovalWorkflows(workflows []*sdm.ApprovalWorkflow) {
	sort.Slice(workflows, func(i, j int) bool {
		return workflows[i].ID < workflows[j].ID
	})
}

// validateApprovalWorkflows compares the Terraform provider response against expected ApprovalWorkflow values.
func validateApprovalWorkflows(dsName string, expectedWorkflows []*sdm.ApprovalWorkflow) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["data.sdm_approval_workflow."+dsName]
		if !ok {
			return fmt.Errorf("Data source not found: data.sdm_approval_workflow.%s", dsName)
		}

		// Get the number of returned workflows
		expectedCount := len(expectedWorkflows)
		actualCount, err := strconv.Atoi(rs.Primary.Attributes["approval_workflows.#"])
		if err != nil {
			return fmt.Errorf("Invalid count value for approval workflows: %v", err)
		}

		// Check if the count matches
		if actualCount != expectedCount {
			return fmt.Errorf("Expected %d approval workflows, but got %d", expectedCount, actualCount)
		}

		// Extract actual workflows from Terraform state
		var actualWorkflows []*sdm.ApprovalWorkflow
		for i := range actualCount {
			prefix := fmt.Sprintf("approval_workflows.%d.", i)

			workflow := &sdm.ApprovalWorkflow{
				ID:           rs.Primary.Attributes[prefix+"id"],
				Name:         rs.Primary.Attributes[prefix+"name"],
				Description:  rs.Primary.Attributes[prefix+"description"],
				ApprovalMode: rs.Primary.Attributes[prefix+"approval_mode"],
			}

			// Get approval step count
			stepCount, err := strconv.Atoi(rs.Primary.Attributes[prefix+"approval_step.#"])
			if err != nil {
				return fmt.Errorf("Invalid approval step count: %v", err)
			}

			// Extract steps
			for j := range stepCount {
				stepPrefix := fmt.Sprintf("%sapproval_step.%d.", prefix, j)

				skipAfter, err := time.ParseDuration(rs.Primary.Attributes[stepPrefix+"skip_after"])
				if err != nil {
					return fmt.Errorf("Invalid skip_after value: %v", err)
				}

				step := &sdm.ApprovalFlowStep{
					Quantifier: rs.Primary.Attributes[stepPrefix+"quantifier"],
					SkipAfter:  skipAfter,
				}

				// Get approver count
				approverCount, err := strconv.Atoi(rs.Primary.Attributes[stepPrefix+"approvers.#"])
				if err != nil {
					return fmt.Errorf("Invalid approver count: %v", err)
				}

				// Extract approvers
				for k := range approverCount {
					approverPrefix := fmt.Sprintf("%sapprovers.%d.", stepPrefix, k)

					approver := &sdm.ApprovalFlowApprover{}
					accountID := rs.Primary.Attributes[approverPrefix+"account_id"]
					roleID := rs.Primary.Attributes[approverPrefix+"role_id"]
					reference := rs.Primary.Attributes[approverPrefix+"reference"]
					if accountID != "" {
						approver.AccountID = accountID
					}
					if roleID != "" {
						approver.RoleID = roleID
					}
					if reference != "" {
						approver.Reference = reference
					}
					step.Approvers = append(step.Approvers, approver)
				}

				workflow.ApprovalWorkflowSteps = append(workflow.ApprovalWorkflowSteps, step)
			}

			actualWorkflows = append(actualWorkflows, workflow)
		}

		// Sort both expected and actual workflows by ID before comparing
		SortApprovalWorkflows(expectedWorkflows)
		SortApprovalWorkflows(actualWorkflows)

		// Compare sorted workflows

		// Element-wise comparison
		for i, expected := range expectedWorkflows {
			actual := actualWorkflows[i]

			// Compare basic attributes
			if expected.ID != actual.ID {
				return fmt.Errorf("Mismatch in workflow ID at index %d: expected %s, got %s", i, expected.ID, actual.ID)
			}
			if expected.Name != actual.Name {
				return fmt.Errorf("Mismatch in workflow Name at index %d: expected %s, got %s", i, expected.Name, actual.Name)
			}
			if expected.Description != actual.Description {
				return fmt.Errorf("Mismatch in workflow Description at index %d: expected %s, got %s", i, expected.Description, actual.Description)
			}
			if expected.ApprovalMode != actual.ApprovalMode {
				return fmt.Errorf("Mismatch in workflow ApprovalMode at index %d: expected %s, got %s", i, expected.ApprovalMode, actual.ApprovalMode)
			}

			// Compare approval step count
			if len(expected.ApprovalWorkflowSteps) != len(actual.ApprovalWorkflowSteps) {
				return fmt.Errorf("Mismatch in approval step count for workflow %s: expected %d, got %d",
					expected.ID, len(expected.ApprovalWorkflowSteps), len(actual.ApprovalWorkflowSteps))
			}

			// Compare approval steps
			for j, expectedStep := range expected.ApprovalWorkflowSteps {
				actualStep := actual.ApprovalWorkflowSteps[j]

				if expectedStep.Quantifier != actualStep.Quantifier {
					return fmt.Errorf("Mismatch in approval step Quantifier at workflow %s, step %d: expected %s, got %s",
						expected.ID, j, expectedStep.Quantifier, actualStep.Quantifier)
				}

				// Convert skipAfter to seconds for comparison
				if expectedStep.SkipAfter != actualStep.SkipAfter {
					return fmt.Errorf("Mismatch in approval step SkipAfter at workflow %s, step %d: expected %s, got %s",
						expected.ID, j, expectedStep.SkipAfter.String(), actualStep.SkipAfter.String())
				}

				// Compare approvers count
				if len(expectedStep.Approvers) != len(actualStep.Approvers) {
					return fmt.Errorf("Mismatch in approver count at workflow %s, step %d: expected %d, got %d",
						expected.ID, j, len(expectedStep.Approvers), len(actualStep.Approvers))
				}

				// Compare approvers
				for k, expectedApprover := range expectedStep.Approvers {
					actualApprover := actualStep.Approvers[k]

					if expectedApprover.AccountID != actualApprover.AccountID {
						return fmt.Errorf("Mismatch in approver AccountID at workflow %s, step %d, approver %d: expected %s, got %s",
							expected.ID, j, k, expectedApprover.AccountID, actualApprover.AccountID)
					}
					if expectedApprover.RoleID != actualApprover.RoleID {
						return fmt.Errorf("Mismatch in approver RoleID at workflow %s, step %d, approver %d: expected %s, got %s",
							expected.ID, j, k, expectedApprover.RoleID, actualApprover.RoleID)
					}
					if expectedApprover.Reference != actualApprover.Reference {
						return fmt.Errorf("Mismatch in approver Reference at workflow %s, step %d, approver %d: expected %s, got %s",
							expected.ID, j, k, expectedApprover.Reference, actualApprover.Reference)
					}
				}
			}
		}

		// If we reach this point, all elements match
		return nil
	}
}
