// Code generated by protogen. DO NOT EDIT.
// See pkg/api/v1/templates/terraform-provider-sdm/sdm/*_test.go.tpl

package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMWorkflowRoles_Get(t *testing.T) {
	initAcceptanceTest(t)

	workflowRoles, err := createWorkflowRolesWithPrefix("wr-test-get", 1)
	if err != nil {
		t.Fatal("failed to create workflow roles: ", err)
	}

	workflowID := workflowRoles[0].WorkflowID
	roleID := workflowRoles[0].RoleID
	workflowRoleName := randomWithPrefix("wr-test-get")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowRoleGetFilterConfig(workflowRoleName, workflowID, roleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "workflow_roles.0.workflow_id", workflowID),
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "workflow_roles.0.role_id", roleID),
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "workflow_roles.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMWorkflowRoles_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createWorkflowRolesWithPrefix("wr-test-get-none", 1)
	if err != nil {
		t.Fatal("failed to create workflow roles: ", err)
	}
	roleID := "r-000333000"

	workflowRoleName := randomWithPrefix("wr-test-get-none")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowRoleGetFilterConfig(workflowRoleName, "", roleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "workflow_roles.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "ids.#", "0"),
				),
			},
		},
	})
}

func TestAccSDMWorkflowRoles_GetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	workflowRoles, err := createWorkflowRolesWithPrefix("wr-test-get-multi", 3)
	if err != nil {
		t.Fatal("failed to create workflow roles: ", err)
	}
	workflowID := workflowRoles[0].WorkflowID

	workflowRoleName := randomWithPrefix("wr-test-get-multi")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMWorkflowRoleGetFilterConfig(workflowRoleName, workflowID, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "workflow_roles.#", "3"),
					resource.TestCheckResourceAttr("data.sdm_workflow_role."+workflowRoleName, "ids.#", "3"),
				),
			},
		},
	})
}

func testAccSDMWorkflowRoleGetFilterConfig(workflowRoleDataSourceName, workflowID, roleID string) string {
	return fmt.Sprintf(
		`
	data "sdm_workflow_role" "%s" {
		workflow_id = "%s"
		role_id = "%s"
	}`,
		workflowRoleDataSourceName,
		workflowID,
		roleID,
	)
}
