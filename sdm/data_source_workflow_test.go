package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMWorkflow_Get(t *testing.T) {
	initAcceptanceTest(t)
	workflows, err := createWorkflowsWithPrefix("workflow-get", 1)
	if err != nil {
		t.Fatal("failed to create workflow: ", err)
	}
	workflow := workflows[0]

	dsName := randomWithPrefix("wf-test-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowDataSourceGetFilterConfig(dsName, "workflow-get*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow."+dsName, "workflows.0.name", workflow.Name),
					resource.TestCheckResourceAttr("data.sdm_workflow."+dsName, "workflows.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMWorkflow_GetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createWorkflowsWithPrefix("wf-multi-test", 2)
	if err != nil {
		t.Fatal("failed to create workflow: ", err)
	}
	_, err = createWorkflowsWithPrefix("nomatch", 1)
	if err != nil {
		t.Fatal("failed to create workflow: ", err)
	}

	dsName := randomWithPrefix("wf-multi-test-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowDataSourceGetFilterConfig(dsName, "wf-multi-test*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow."+dsName, "workflows.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMWorkflow_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createWorkflowsWithPrefix("dontfind", 1)
	if err != nil {
		t.Fatal("failed to create workflow: ", err)
	}

	dsName := randomWithPrefix("wf-none-query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowDataSourceGetFilterConfig(dsName, "none*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow."+dsName, "workflows.#", "0"),
				),
			},
		},
	})
}

func testAccSDMWorkflowDataSourceGetFilterConfig(workflowDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_workflow" "%s" {
		name = "%s"
	}`, workflowDataSourceName, nameFilter)
}
