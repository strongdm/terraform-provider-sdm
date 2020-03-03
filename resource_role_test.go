package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	sdm "github.com/strongdm/strongdm-sdk-go"
)

func init() {
	resource.AddTestSweepers("sdm_role", &resource.Sweeper{
		Name:         "sdm_role",
		Dependencies: []string{"sdm_account_attachment", "sdm_role_attachment", "sdm_role_grant"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.Roles().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing roles: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.Roles().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting roles %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing roles %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMRole_Create(t *testing.T) {
	rsName := randomWithPrefix("test-role-resource")
	roleName := randomWithPrefix("test-role")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleConfig(rsName, roleName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+rsName, "name", roleName),
					resource.TestCheckResourceAttr("sdm_role."+rsName, "composite", "false"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Roles().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Role.Name != roleName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Role.Name, roleName)
						}

						return nil
					},
				),
			},
		},
	})
}

func TestAccSDMRole_Update(t *testing.T) {
	rsName := randomWithPrefix("test-role")
	roleName := randomWithPrefix("test-role")
	updatedRoleName := randomWithPrefix("test-role-updated")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleConfig(rsName, roleName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+rsName, "name", roleName),
					resource.TestCheckResourceAttr("sdm_role."+rsName, "composite", "false"),
				),
			},
			{
				Config: testAccSDMRoleConfig(rsName, updatedRoleName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+rsName, "name", updatedRoleName),
					resource.TestCheckResourceAttr("sdm_role."+rsName, "composite", "false"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", rsName)
						if err != nil {
							return err
						}

						// check if it was actually updated
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Roles().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Role.Name != updatedRoleName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Role.Name, updatedRoleName)
						}

						return nil
					},
				),
			},
		},
	})
}

func testAccSDMRoleConfig(resourceName, roleName string, composite bool) string {
	return fmt.Sprintf(`
	resource "sdm_role" "%s" {
		name = "%s"
		composite = %t
	}
	`, resourceName, roleName, composite)
}

func createRolesWithPrefix(prefix string, count int, composite bool) ([]*sdm.Role, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	roles := []*sdm.Role{}
	for i := 0; i < count; i++ {
		createResp, err := client.Roles().Create(ctx, &sdm.Role{
			Name:      randomWithPrefix(prefix),
			Composite: composite,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create role: %w", err)
		}
		roles = append(roles, createResp.Role)
	}
	return roles, nil
}
