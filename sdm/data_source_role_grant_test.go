package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMRoleGrantDataSource_Get(t *testing.T) {
	t.Parallel()

	grants, err := createRoleGrantsWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create role grants: ", err)
	}
	grant := grants[0]

	roleGrantDataName := randomWithPrefix("roleGrant")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleGrantDataSourceGetFilterConfig(roleGrantDataName, grant.RoleID, grant.ResourceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataName, "role_grants.0.role_id", grant.RoleID),
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataName, "role_grants.0.resource_id", grant.ResourceID),
				),
			},
		},
	})
}

func TestAccSDMRoleGrantDataSource_GetMultiple(t *testing.T) {
	t.Parallel()

	// create three redises
	resources, err := createRedisesWithPrefix("test", 3)
	if err != nil {
		t.Fatal("failed to create redises: ", err)
	}
	roles, err := createRolesWithPrefix("test", 1, false)
	if err != nil {
		t.Fatal("failed to create role: ", err)
	}
	// only grant two of the created resources
	_, err = grantRole(roles[0], resources[1:])
	if err != nil {
		t.Fatal("failed to grant roles: ", err)
	}
	roleID := roles[0].ID

	roleGrantDataSourceName := randomWithPrefix("roleGrant")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleGrantDataSourceGetFilterConfig(roleGrantDataSourceName, roleID, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataSourceName, "role_grants.0.role_id", roleID),
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataSourceName, "role_grants.1.role_id", roleID),
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataSourceName, "ids.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataSourceName, "role_grants.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMRoleGrantDataSource_GetNone(t *testing.T) {
	_, err := createRoleGrantsWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create role grants: ", err)
	}
	roleGrantDataSourceName := randomWithPrefix("roleGrant")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleGrantDataSourceGetFilterConfig(roleGrantDataSourceName, "r-00333000", ""), // must not be a valid role ID
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataSourceName, "ids.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_role_grant."+roleGrantDataSourceName, "role_grants.#", "0"),
				),
			},
		},
	})
}

func testAccSDMRoleGrantDataSourceGetFilterConfig(roleGrantDataSourceName, roleID, resourceID string) string {
	return fmt.Sprintf(`
	data "sdm_role_grant" "%s" {
		role_id = "%s"
		resource_id = "%s"
	}`,
		roleGrantDataSourceName,
		roleID,
		resourceID,
	)
}
