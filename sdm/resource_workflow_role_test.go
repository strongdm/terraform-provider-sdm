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
	resource.AddTestSweepers("sdm_workflow_role", &resource.Sweeper{
		Name: "sdm_workflow_role",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.WorkflowRoles().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing workflow roles: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.WorkflowRoles().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting workflow role %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing workflow roles %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMWorkflowRole_Create(t *testing.T) {
	initAcceptanceTest(t)
	rs1Name := randomWithPrefix("test-workflow-role1")
	rs2Name := randomWithPrefix("test-workflow-role2")
	workflowName := randomWithPrefix("test-workflow")
	role1Name := randomWithPrefix("test-role1")
	role2Name := randomWithPrefix("test-role2")
	workflowID := ""
	role1ID := ""
	role2ID := ""
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowRolesConfig(rs1Name, rs2Name, workflowName, role1Name, role2Name),
				Check: resource.ComposeTestCheckFunc(
					// Ensure and capture the mapping IDs
					func(s *terraform.State) error {
						var err error
						workflowID, err = testCreatedID(s, "sdm_workflow", workflowName)
						if err != nil {
							return err
						}
						role1ID, err = testCreatedID(s, "sdm_role", role1Name)
						if err != nil {
							return err
						}
						role2ID, err = testCreatedID(s, "sdm_role", role2Name)
						if err != nil {
							return err
						}
						return nil
					},
					// Ensure and check the WorkflowRole creations
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_workflow_role."+rs1Name, "workflow_id", workflowID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_workflow_role."+rs1Name, "role_id", role1ID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_workflow_role."+rs2Name, "workflow_id", workflowID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_workflow_role."+rs2Name, "role_id", role2ID)(s); err != nil {
							return err
						}
						return nil
					},
					// Ensure the resources were actually created
					func(s *terraform.State) error {
						id1, err := testCreatedID(s, "sdm_workflow_role", rs1Name)
						if err != nil {
							return err
						}
						id2, err := testCreatedID(s, "sdm_workflow_role", rs2Name)
						if err != nil {
							return err
						}

						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						ids := map[string]string{id1: role1ID, id2: role2ID}
						for workflowRoleID, roleID := range ids {
							resp, err := client.WorkflowRoles().Get(ctx, workflowRoleID)
							if err != nil {
								return err
							}
							if resp.WorkflowRole.WorkflowID != workflowID {
								return fmt.Errorf("unexpected workflowID '%s', expected '%s'", resp.WorkflowRole.WorkflowID, workflowID)
							}
							if resp.WorkflowRole.RoleID != roleID {
								return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.WorkflowRole.RoleID, roleID)
							}
						}
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_workflow_role." + rs1Name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "sdm_workflow_role." + rs2Name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createWorkflowRolesWithPrefix(prefix string, count int) ([]*sdm.WorkflowRole, error) {
	roles, err := createRolesWithPrefix(prefix, count)
	if err != nil {
		return nil, fmt.Errorf("failed to create roles: %w", err)
	}
	workflows, err := createWorkflowsWithPrefix(prefix, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create workflow: %w", err)
	}
	workflowRoles, err := addRolesToWorkflow(roles, workflows[0])
	if err != nil {
		return nil, fmt.Errorf("failed to add roles: %w", err)
	}
	return workflowRoles, nil
}

func addRolesToWorkflow(roles []*sdm.Role, workflow *sdm.Workflow) ([]*sdm.WorkflowRole, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	workflowRoles := []*sdm.WorkflowRole{}
	for _, r := range roles {
		createResp, err := client.WorkflowRoles().Create(ctx, &sdm.WorkflowRole{
			WorkflowID: workflow.ID,
			RoleID:     r.ID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create workflow role: %w", err)
		}
		workflowRoles = append(workflowRoles, createResp.WorkflowRole)
	}
	return workflowRoles, nil
}

func testAccSDMWorkflowRolesConfig(name1, name2, workflowName, role1Name, role2Name string) string {
	return fmt.Sprintf(
		`
	resource "sdm_workflow" "%[1]s" {
		name = "%[2]s"
		weight = "10"
	}

	resource "sdm_role" "%[3]s" {
		name = "%[4]s"
	}

	resource "sdm_role" "%[5]s" {
		name = "%[6]s"
	}

	resource "sdm_workflow_role" "%[7]s" {
		workflow_id = sdm_workflow.%[1]s.id
		role_id = sdm_role.%[3]s.id
	}

	resource "sdm_workflow_role" "%[8]s" {
		workflow_id = sdm_workflow.%[1]s.id
		role_id = sdm_role.%[5]s.id
	}
	`,
		workflowName,
		randomWithPrefix("workflow-name"),
		role1Name,
		randomWithPrefix("role1-name"),
		role2Name,
		randomWithPrefix("role2-name"),
		name1,
		name2,
	)
}
