package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMApprovalWorkflow_Get(t *testing.T) {
	initAcceptanceTest(t)
	workflows, err := createApprovalWorkflowsWithPrifix("approval-workflow-get", "automatic", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}
	workflow := workflows[0]

	dsName := randomWithPrefix("af-test-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "approval-workflow-get*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.0.name", workflow.Name),
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_GetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createApprovalWorkflowsWithPrifix("multi-test", "automatic", 2)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}
	_, err = createApprovalWorkflowsWithPrifix("nomatch", "automatic", 1)
	if err != nil {
		t.Fatal("failed to create approval workflow: ", err)
	}

	dsName := randomWithPrefix("af-multi-test-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMApprovalWorkflowDataSourceGetFilterConfig(dsName, "multi-test*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_approval_workflow."+dsName, "approval_workflows.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMApprovalWorkflow_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createApprovalWorkflowsWithPrifix("dontfind", "automatic", 1)
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

func testAccSDMApprovalWorkflowDataSourceGetFilterConfig(workflowDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_approval_workflow" "%s" {
		name = "%s"
	}`, workflowDataSourceName, nameFilter)
}
