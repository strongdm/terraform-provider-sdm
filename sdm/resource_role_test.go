package sdm

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
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
			{
				ResourceName:      "sdm_role." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
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
				ResourceName:      "sdm_role." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
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
			{
				ResourceName:      "sdm_role." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMRole_UpdateComposite(t *testing.T) {
	if os.Getenv("SDM_ACCESS_OVERHAUL") == "yes" {
		t.Skip("skipping legacy test")
	}
	rsName := randomWithPrefix("test-role")
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
				),
			},
			{
				ResourceName:      "sdm_role." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMRoleConfig(rsName, roleName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+rsName, "name", roleName),
					resource.TestCheckResourceAttr("sdm_role."+rsName, "composite", "true"),
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

						if resp.Role.Composite != true {
							return fmt.Errorf("unexpected composite '%t'", resp.Role.Composite)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMRole_Tags(t *testing.T) {
	name := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleTagsConfig(name, name, false, sdm.Tags{"key": "value", "foo": "bar"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+name, "tags.key", "value"),
					resource.TestCheckResourceAttr("sdm_role."+name, "tags.foo", "bar"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Roles().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created role: %w", err)
						}

						if resp.Role.Tags["key"] != "value" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Role.Tags["key"])
						}
						if resp.Role.Tags["foo"] != "bar" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Role.Tags["key"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMRoleTagsConfig(name, name, false, sdm.Tags{"goat": "bananas"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sdm_role."+name, "tags.key"),
					resource.TestCheckNoResourceAttr("sdm_role."+name, "tags.foo"),
					resource.TestCheckResourceAttr("sdm_role."+name, "tags.goat", "bananas"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Roles().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created role: %w", err)
						}

						if resp.Role.Tags["goat"] != "bananas" {
							return fmt.Errorf("unexpected value '%s' for tag 'goat', expected 'bananas'", resp.Role.Tags["goat"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMRole_AccessRules(t *testing.T) {
	if os.Getenv("SDM_ACCESS_OVERHAUL") != "yes" {
		t.Skip("skipping access overhaul test")
	}
	name := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleAccessRulesConfig(name, name, false, `[ { "type": "redis" } ]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+name, "access_rules", `[{"type":"redis"}]`),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Roles().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created role: %w", err)
						}

						if resp.Role.AccessRules != `[{"type":"redis"}]` {
							return fmt.Errorf(`unexpected value '%s' for access_rules, expected '[{"type":"redis"}]'`, resp.Role.AccessRules)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMRoleAccessRulesConfig(name, name, false, `[ { "tags": { "a": "b" } } ]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+name, "access_rules", `[{"tags":{"a":"b"}}]`),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Roles().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created role: %w", err)
						}

						if resp.Role.AccessRules != `[{"tags":{"a":"b"}}]` {
							return fmt.Errorf(`unexpected value '%s' for access_rules, expected '[{"tags":{"a":"b"}}]'`, resp.Role.AccessRules)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + name,
				ImportState:       true,
				ImportStateVerify: true,
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

func testAccSDMRoleTagsConfig(resourceName string, roleName string, composite bool, tags sdm.Tags) string {
	return fmt.Sprintf(`
	resource "sdm_role" "%s" {
		name = "%s"
		composite = %t
		tags = {
%s
		}
	}
	`, resourceName, roleName, composite, tagsToConfigString(tags))
}

func testAccSDMRoleAccessRulesConfig(resourceName string, roleName string, composite bool, accessRules string) string {
	return fmt.Sprintf(`
	resource "sdm_role" "%s" {
		name = "%s"
		composite = %t
		access_rules = %s
	}
	`, resourceName, roleName, composite, strconv.Quote(accessRules))
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

func createRolesWithAccessRules(prefix string, count int, composite bool, accessRules string) ([]*sdm.Role, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	roles := []*sdm.Role{}
	for i := 0; i < count; i++ {
		createResp, err := client.Roles().Create(ctx, &sdm.Role{
			Name:        randomWithPrefix(prefix),
			Composite:   composite,
			AccessRules: accessRules,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create role: %w", err)
		}
		roles = append(roles, createResp.Role)
	}
	return roles, nil
}
