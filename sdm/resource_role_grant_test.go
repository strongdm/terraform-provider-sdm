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
	resource.AddTestSweepers("sdm_role_grant", &resource.Sweeper{
		Name: "sdm_role_grant",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			roleResp, err := client.Roles().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing roles: %s", err)
			}

			for roleResp.Next() {
				r := roleResp.Value()
				rgResp, err := client.RoleGrants().List(ctx, "roleid:?", r.ID)
				if err != nil {
					return fmt.Errorf("Error listing role grants: %s", err)
				}
				for rgResp.Next() {
					v := rgResp.Value()
					_, err := client.RoleGrants().Delete(ctx, v.ID)
					if err != nil {
						fmt.Printf("error deleting role grant %s %s\n", v.ID, err)
					}
				}
				if rgResp.Err() != nil {
					fmt.Printf("error after listing role grants %s", rgResp.Err())
				}
			}
			if roleResp.Err() != nil {
				return fmt.Errorf("error listing roles %s", roleResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMRoleGrant_Create(t *testing.T) {
	rsName := randomWithPrefix("test-create-role-grant")
	roleRsName := randomWithPrefix("test-role")
	redisRsName := randomWithPrefix("test-redis")
	roleID := ""
	resourceID := ""
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleGrantConfig(rsName, redisRsName, roleRsName),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						var err error
						roleID, err = testCreatedID(s, "sdm_role", roleRsName)
						if err != nil {
							return err
						}

						resourceID, err = testCreatedID(s, "sdm_resource", redisRsName)
						if err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_role_grant."+rsName, "role_id", roleID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_role_grant."+rsName, "resource_id", resourceID)(s); err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {

						id, err := testCreatedID(s, "sdm_role_grant", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.RoleGrants().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.RoleGrant.RoleID != roleID {
							return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.RoleGrant.RoleID, roleID)
						}

						if resp.RoleGrant.ResourceID != resourceID {
							return fmt.Errorf("unexpected resourceID '%s', expected '%s'", resp.RoleGrant.ResourceID, resourceID)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role_grant." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMRoleGrant_Update(t *testing.T) {
	rsName := randomWithPrefix("test-update-role-grant")
	redisRsName := randomWithPrefix("test-redis")
	roleRsName := randomWithPrefix("test-role")
	altRoleRsName := randomWithPrefix("test-role-updated")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleGrantConfig(rsName, redisRsName, roleRsName),
				Check: func(s *terraform.State) error {
					roleID, err := testCreatedID(s, "sdm_role", roleRsName)
					if err != nil {
						return err
					}

					resourceID, err := testCreatedID(s, "sdm_resource", redisRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_role_grant", rsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.RoleGrants().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created resource: %w", err)
					}

					if resp.RoleGrant.RoleID != roleID {
						return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.RoleGrant.RoleID, roleID)
					}

					if resp.RoleGrant.ResourceID != resourceID {
						return fmt.Errorf("unexpected resourceID '%s', expected '%s'", resp.RoleGrant.ResourceID, resourceID)
					}

					return nil
				},
			},
			{
				ResourceName:      "sdm_role_grant." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMRoleGrantConfig(rsName, redisRsName, altRoleRsName),
				Check: func(s *terraform.State) error {
					roleID, err := testCreatedID(s, "sdm_role", altRoleRsName)
					if err != nil {
						return err
					}

					resourceID, err := testCreatedID(s, "sdm_resource", redisRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_role_grant", rsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.RoleGrants().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created resource: %w", err)
					}

					if resp.RoleGrant.RoleID != roleID {
						return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.RoleGrant.RoleID, roleID)
					}

					if resp.RoleGrant.ResourceID != resourceID {
						return fmt.Errorf("unexpected resourceID '%s', expected '%s'", resp.RoleGrant.ResourceID, resourceID)
					}

					return nil
				},
			},
			{
				ResourceName:      "sdm_role_grant." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSDMRoleGrantConfig(roleGrantResourceName, redisResourceName, roleResourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%[1]s" {
		redis {
			name = "%[2]s"
			hostname = "test.com"
		}
	}

	resource "sdm_role" "%[3]s" {
		name = "%[4]s"
	}

	resource "sdm_role_grant" "%[5]s" {
		role_id = sdm_role.%[6]s.id
		resource_id = sdm_resource.%[7]s.id
	}
	`,
		redisResourceName,
		randomWithPrefix("redisname"),
		roleResourceName,
		randomWithPrefix("rolename"),
		roleGrantResourceName,
		roleResourceName,
		redisResourceName,
	)
}

func createRoleGrantsWithPrefix(prefix string, count int) ([]*sdm.RoleGrant, error) {
	resources, err := createRedisesWithPrefix(prefix, count)
	if err != nil {
		return nil, fmt.Errorf("failed to create roles: %w", err)
	}
	roles, err := createRolesWithPrefix(prefix, 1, false)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}
	roleGrants, err := grantRole(roles[0], resources)
	if err != nil {
		return nil, fmt.Errorf("failed to grant roles: %w", err)
	}

	return roleGrants, nil
}

func grantRole(role *sdm.Role, resources []sdm.Resource) ([]*sdm.RoleGrant, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	grants := []*sdm.RoleGrant{}
	for _, res := range resources {
		createResp, err := client.RoleGrants().Create(ctx, &sdm.RoleGrant{
			RoleID:     role.ID,
			ResourceID: res.GetID(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create role grant: %w", err)
		}
		grants = append(grants, createResp.RoleGrant)
	}
	return grants, nil
}
