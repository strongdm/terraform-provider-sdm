package sdm

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMRoleDataSource_Get(t *testing.T) {
	if os.Getenv("SDM_ACCESS_OVERHAUL") != "yes" {
		t.Skip("skipping access overhaul test")
	}
	t.Parallel()

	roles, err := createRolesWithAccessRules("test-role", 1, false, `[{"type":"redis"}]`)
	if err != nil {
		t.Fatal("failed to create role: ", err)
	}
	role := roles[0]

	rsName := randomWithPrefix("test-role-ds")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleDataSourceGetFilterConfig(rsName, role.Name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.0.name", role.Name),
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.0.composite", "false"),
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.0.access_rules", `[{"type":"redis"}]`),
				),
			},
		},
	})
}

func TestAccSDMRoleDataSource_GetByTags(t *testing.T) {
	t.Parallel()

	client, err := preTestClient()
	if err != nil {
		t.Fatal("failed to create test client", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	createResp, err := client.Roles().Create(ctx, &sdm.Role{
		Name: randomWithPrefix("role-tags"),
		Tags: sdm.Tags{"envTags": "devTags"},
	})
	if err != nil {
		t.Fatal("failed to create role", err)
	}
	role := createResp.Role
	rsName := randomWithPrefix("test-role-ds")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "sdm_role" "%s" {
						tags = {
							envTags = "devTags"
						}
					}`, rsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.0.name", role.Name),
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.0.tags.envTags", "devTags"),
				),
			},
		},
	})
}

func TestAccSDMRoleDataSource_GetMultiple(t *testing.T) {
	t.Parallel()

	_, err := createRolesWithPrefix("multi-test", 2, false)
	if err != nil {
		t.Fatal("failed to create roles: ", err)
	}
	_, err = createRolesWithPrefix("nomatch", 1, false)
	if err != nil {
		t.Fatal("failed to create role: ", err)
	}

	rsName := randomWithPrefix("test-role")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleDataSourceGetFilterConfig(rsName, "multi-test*", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.0.composite", "false"),
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.1.composite", "false"),
				),
			},
		},
	})
}

func TestAccSDMRoleDataSource_GetNone(t *testing.T) {
	t.Parallel()

	_, err := createRolesWithPrefix("dontfind", 1, false)
	if err != nil {
		t.Fatal("failed to create role: ", err)
	}

	dsName := randomWithPrefix("role-ds")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleDataSourceGetFilterConfig(dsName, "neverWillMatch", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+dsName, "roles.#", "0"),
				),
			},
		},
	})
}

func TestAccSDMRoleDataSource_GetByID(t *testing.T) {
	roleName := randomWithPrefix("role")
	dsName := randomWithPrefix("role")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleDataSourceGetByIDConfig(roleName, dsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+dsName, "roles.#", "1"),
				),
			},
		},
	})
}

func testAccSDMRoleDataSourceGetFilterConfig(roleDataSourceName, nameFilter string, composite bool) string {
	return fmt.Sprintf(`
	data "sdm_role" "%s" {
		name = "%s"
		composite = %t
	}`, roleDataSourceName, nameFilter, composite)
}

func testAccSDMRoleDataSourceGetByIDConfig(roleName string, roleDataSourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_role" "test_role" {
		name = "%s"
	}

	data "sdm_role" "%s" {
		id = sdm_role.test_role.id
	}
	`, roleName, roleDataSourceName)
}
