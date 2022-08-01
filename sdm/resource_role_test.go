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
				Config: testAccSDMRoleConfig(rsName, roleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+rsName, "name", roleName),
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
				Config: testAccSDMRoleConfig(rsName, roleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+rsName, "name", roleName),
				),
			},
			{
				ResourceName:      "sdm_role." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMRoleConfig(rsName, updatedRoleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+rsName, "name", updatedRoleName),
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

func TestAccSDMRole_Tags(t *testing.T) {
	name := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleTagsConfig(name, name, sdm.Tags{"key": "value", "foo": "bar"}),
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
				Config: testAccSDMRoleTagsConfig(name, name, sdm.Tags{"goat": "bananas"}),
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
	resourceName := randomWithPrefix("testrs")
	roleName := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			// ID-based access rule
			{
				Config: testAccSDMRoleAccessRulesConfig(resourceName, roleName,
					fmt.Sprintf(`access_rules = jsonencode([
						{ ids = [sdm_resource.%s.id] }
					])`, resourceName),
				),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", roleName)
						if err != nil {
							return err
						}

						resourceID, err := testCreatedID(s, "sdm_resource", resourceName)
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

						expectedPolicy := sdm.AccessRules(fmt.Sprintf(`[{"ids":["%s"]}]`, resourceID))
						if resp.Role.AccessRules != expectedPolicy {
							return fmt.Errorf(`unexpected value '%+v' for access_rule, expected %+v`, resp.Role.AccessRules, expectedPolicy)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + roleName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// dbtype-based access rules
			{
				Config: testAccSDMRoleAccessRulesConfig(
					resourceName,
					roleName,
					`access_rules = jsonencode([
						{ type = "redis" },
						{ type = "postgres" }
					])`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+roleName, "access_rules", `[{"type":"redis"},{"type":"postgres"}]`),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", roleName)
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

						if resp.Role.AccessRules != `[{"type":"redis"},{"type":"postgres"}]` {
							return fmt.Errorf(`unexpected value '%+v' for access_rule, expected '[{"type":"redis"},{"type":"postgres"}]'`, resp.Role.AccessRules)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + roleName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// change order of access rules
			{
				Config: testAccSDMRoleAccessRulesConfig(
					resourceName,
					roleName,
					`access_rules = jsonencode([
						{ type = "postgres" },
						{ type = "redis" },
					])`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+roleName, "access_rules", `[{"type":"postgres"},{"type":"redis"}]`),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", roleName)
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

						if resp.Role.AccessRules != `[{"type":"postgres"},{"type":"redis"}]` {
							return fmt.Errorf(`unexpected value '%+v' for access_rule, expected '[{"type":"postgres"},{"type":"redis"}]'`, resp.Role.AccessRules)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + roleName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// tag-based access rule
			{
				Config: testAccSDMRoleAccessRulesConfig(resourceName, roleName,
					`access_rules = jsonencode([
						{ tags = { a = "b" } }
					])`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+roleName, "access_rules", `[{"tags":{"a":"b"}}]`),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", roleName)
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
							return fmt.Errorf(`unexpected value '%+v' for access_rule, expected '[{"tags":{"a":"b"}}]'`, resp.Role.AccessRules)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + roleName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// elide access rules field; access rules should stay the same (it's an optional computed field)
			{
				Config: testAccSDMRoleAccessRulesConfig(resourceName, roleName, ``),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+roleName, "access_rules", `[{"tags":{"a":"b"}}]`),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", roleName)
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
							return fmt.Errorf(`unexpected value '%+v' for access_rule, expected '[{"tags":{"a":"b"}}]'`, resp.Role.AccessRules)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + roleName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// now clear out access rules
			{
				Config: testAccSDMRoleAccessRulesConfig(resourceName, roleName, `access_rules = jsonencode([])`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_role."+roleName, "access_rules", `[]`),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_role", roleName)
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

						if resp.Role.AccessRules != "[]" {
							return fmt.Errorf(`expected no access rules, got %v`, resp.Role.AccessRules)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_role." + roleName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSDMRoleConfig(resourceName, roleName string) string {
	return fmt.Sprintf(`
	resource "sdm_role" "%s" {
		name = "%s"
	}
	`, resourceName, roleName)
}

func testAccSDMRoleTagsConfig(resourceName string, roleName string, tags sdm.Tags) string {
	return fmt.Sprintf(`
	resource "sdm_role" "%s" {
		name = "%s"
		tags = {
%s
		}
	}
	`, resourceName, roleName, tagsToConfigString(tags))
}

func testAccSDMRoleAccessRulesConfig(resourceName string, roleName string, accessRules string) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%[1]s" {
		redis {
			name = "%[1]s"
			hostname = "example.com"
		}
	}

	resource "sdm_role" "%[2]s" {
		name = "%[2]s"
		%[3]s
	}
	`, resourceName, roleName, accessRules)
}

func createRolesWithPrefix(prefix string, count int) ([]*sdm.Role, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	roles := []*sdm.Role{}
	for i := 0; i < count; i++ {
		createResp, err := client.Roles().Create(ctx, &sdm.Role{
			Name: randomWithPrefix(prefix),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create role: %w", err)
		}
		roles = append(roles, createResp.Role)
	}
	return roles, nil
}

func createRolesWithAccessRules(prefix string, count int, accessRules sdm.AccessRules) ([]*sdm.Role, error) {
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
			AccessRules: accessRules,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create role: %w", err)
		}
		roles = append(roles, createResp.Role)
	}
	return roles, nil
}
