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
	resourceName := randomWithPrefix("wf-create")
	workflowName := randomWithPrefix("wf-create")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowConfig(resourceName, workflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "name", workflowName),
					doubleCheckWorkflow(resourceName, workflowName),
				),
			},
			{
				ResourceName:      "sdm_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMWorkflow_Update(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("wf-create")
	workflowName := randomWithPrefix("wf-create")
	updatedWorkflowName := randomWithPrefix("wf-create-updated")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowConfig(resourceName, workflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "name", workflowName),
					doubleCheckWorkflow(resourceName, workflowName),
				),
			},
			{
				ResourceName:      "sdm_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMWorkflowConfig(resourceName, updatedWorkflowName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "name", updatedWorkflowName),
					doubleCheckWorkflow(resourceName, updatedWorkflowName),
				),
			},
			{
				ResourceName:      "sdm_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMWorkflow_Settings(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("wf-create")
	workflowName := randomWithPrefix("wf-create")
	maxDuration := "1h30m0s"
	fixedDuration := "1h45m0s"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowConfigWithSettings(resourceName, workflowName, maxDuration, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "name", workflowName),
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "access_request_max_duration", maxDuration),
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "access_request_fixed_duration", "0s"),
					doubleCheckWorkflow(resourceName, workflowName),
				),
			},
			{
				ResourceName:      "sdm_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMWorkflowConfigWithSettings(resourceName, workflowName, "", fixedDuration),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "name", workflowName),
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "access_request_max_duration", "0s"),
					resource.TestCheckResourceAttr("sdm_workflow."+resourceName, "access_request_fixed_duration", fixedDuration),
					doubleCheckWorkflow(resourceName, workflowName),
				),
			},
			{
				ResourceName:      "sdm_workflow." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createWorkflowsWithPrefix(prefix string, count int) ([]*sdm.Workflow, error) {
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

func testAccSDMWorkflowConfig(workflowName, sdmResourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_workflow" "%s" {
		name = "%s"
	}`, workflowName, sdmResourceName)
}

func testAccSDMWorkflowConfigWithSettings(workflowName, sdmResourceName, maxDuration, fixedDuration string) string {
	value := fmt.Sprintf(`
	resource "sdm_workflow" "%s" {
		name = "%s"`, workflowName, sdmResourceName) + "\n"
	if maxDuration != "" {
		value += fmt.Sprintf(`access_request_max_duration = "%s"`, maxDuration) + "\n"
	}
	if fixedDuration != "" {
		value += fmt.Sprintf(`access_request_fixed_duration = "%s"`, fixedDuration) + "\n"
	}
	value += `}`
	return value
}

func doubleCheckWorkflow(resourceName, workflowName string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		id, err := testCreatedID(s, "sdm_workflow", resourceName)
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
		if resp.Workflow.Name != workflowName {
			return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Workflow.Name, workflowName)
		}

		return nil
	}
}
