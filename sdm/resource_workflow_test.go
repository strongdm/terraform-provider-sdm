package sdm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func init() {
	resource.AddTestSweepers("sdm_workflow", &resource.Sweeper{
		Name:         "sdm_workflow",
		Dependencies: []string{"sdm_workflow_role", "sdm_workflow_approver"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.Workflows().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing workflows: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.Workflows().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting workflow %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing workflows %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMWorkflow_Create(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("wf-create")
	weight := "10"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowConfig(name, name, weight),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_workflow."+name, "name", name),
					resource.TestCheckResourceAttr("sdm_workflow."+name, "weight", weight),
					doubleCheckWorkflow(name, weight),
				),
			},
			{
				ResourceName:      "sdm_workflow." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createWorkflowsWithPrifix(prefix string, count int) ([]*sdm.Workflow, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	workflows := []*sdm.Workflow{}
	for i := 0; i < count; i++ {
		createResp, err := client.Workflows().Create(ctx, &sdm.Workflow{
			Name: randomWithPrefix(prefix),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create workflow: %w", err)
		}
		workflows = append(workflows, createResp.Workflow)
	}
	return workflows, nil
}

func testAccSDMWorkflowConfig(workflowName, sdmResourceName, weight string) string {
	return fmt.Sprintf(`
	resource "sdm_workflow" "%s" {
		name = "%s"
		weight = "%s"
	}`, workflowName, sdmResourceName, weight)
}

func doubleCheckWorkflow(name, weight string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		id, err := testCreatedID(s, "sdm_workflow", name)
		if err != nil {
			return err
		}
		client := testClient()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		resp, err := client.Workflows().Get(ctx, id)
		if err != nil {
			return fmt.Errorf("failed to get created resource: %w", err)
		}
		if resp.Workflow.Name != name {
			return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Workflow.Name, name)
		}
		w, err := strconv.ParseInt(weight, 10, 64)
		if err != nil {
			return &resource.TimeoutError{}
		}
		if resp.Workflow.Weight != w {
			return fmt.Errorf("unexpected weight '%d', expected '%d'", resp.Workflow.Weight, w)
		}

		return nil
	}
}
