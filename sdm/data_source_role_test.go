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

	roles, err := createRolesWithAccessRules("test-role", 1, `[{"type":"redis"}]`)
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
				Config: testAccSDMRoleDataSourceGetFilterConfig(rsName, role.Name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.0.name", role.Name),
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

	tagValue := randomWithPrefix("value")
	createResp, err := client.Roles().Create(ctx, &sdm.Role{
		Name: randomWithPrefix("role-tags"),
		Tags: sdm.Tags{"testTag": tagValue},
	})
	if err != nil {
		t.Fatal("failed to create role", err)
	}
	role := createResp.Role
	dsName := randomWithPrefix("test-role-ds")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "sdm_role" "%s" {
						tags = {
							"testTag" = "%s"
						}
					}`, dsName, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+dsName, "roles.0.name", role.Name),
					resource.TestCheckResourceAttr("data.sdm_role."+dsName, "roles.0.tags.testTag", tagValue),
				),
			},
		},
	})
}

func TestAccSDMRoleDataSource_GetMultiple(t *testing.T) {
	t.Parallel()

	_, err := createRolesWithPrefix("multi-test", 2)
	if err != nil {
		t.Fatal("failed to create roles: ", err)
	}
	_, err = createRolesWithPrefix("nomatch", 1)
	if err != nil {
		t.Fatal("failed to create role: ", err)
	}

	rsName := randomWithPrefix("test-role")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleDataSourceGetFilterConfig(rsName, "multi-test*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role."+rsName, "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMRoleDataSource_GetNone(t *testing.T) {
	t.Parallel()

	_, err := createRolesWithPrefix("dontfind", 1)
	if err != nil {
		t.Fatal("failed to create role: ", err)
	}

	dsName := randomWithPrefix("role-ds")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleDataSourceGetFilterConfig(dsName, "neverWillMatch"),
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

func testAccSDMRoleDataSourceGetFilterConfig(roleDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_role" "%s" {
		name = "%s"
	}`, roleDataSourceName, nameFilter)
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
