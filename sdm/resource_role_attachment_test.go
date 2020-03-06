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
	resource.AddTestSweepers("sdm_role_attachment", &resource.Sweeper{
		Name: "sdm_role_attachment",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			roleResp, err := client.Roles().List(ctx, "name:*-test-acc, composite:true")
			if err != nil {
				return fmt.Errorf("Error listing roles: %s", err)
			}

			for roleResp.Next() {
				r := roleResp.Value()
				raResp, err := client.RoleAttachments().List(ctx, "compositeroleid:?", r.ID)
				if err != nil {
					return fmt.Errorf("Error listing role attachments: %s", err)
				}
				for raResp.Next() {
					v := raResp.Value()
					_, err := client.RoleAttachments().Delete(ctx, v.ID)
					if err != nil {
						fmt.Printf("error deleting role attachment %s %s\n", v.ID, err)
					}
				}
				if raResp.Err() != nil {
					fmt.Printf("error listing role attachments %s", raResp.Err())
				}
			}
			if roleResp.Err() != nil {
				return fmt.Errorf("error listing roles %s", roleResp.Err())
			}
			return nil
		},
	})
}
func TestAccSDMRoleAttachment_Create(t *testing.T) {
	rsName := randomWithPrefix("test-create-attachment")
	compRoleRsName := randomWithPrefix("test-comprole")
	roleRsName := randomWithPrefix("test-role")
	compositeID := ""
	attachedID := ""
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleAttachmentConfig(rsName, compRoleRsName, roleRsName),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						var err error
						compositeID, err = testCreatedID(s, "sdm_role", compRoleRsName)
						if err != nil {
							return err
						}

						attachedID, err = testCreatedID(s, "sdm_role", roleRsName)
						if err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_role_attachment."+rsName, "composite_role_id", compositeID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_role_attachment."+rsName, "attached_role_id", attachedID)(s); err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {

						id, err := testCreatedID(s, "sdm_role_attachment", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.RoleAttachments().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.RoleAttachment.CompositeRoleID != compositeID {
							return fmt.Errorf("unexpected compositeID '%s', expected '%s'", resp.RoleAttachment.CompositeRoleID, compositeID)
						}

						if resp.RoleAttachment.AttachedRoleID != attachedID {
							return fmt.Errorf("unexpected attachedID '%s', expected '%s'", resp.RoleAttachment.AttachedRoleID, attachedID)
						}

						return nil
					},
				),
			},
		},
	})
}

// TODO: can we improve the readability of this test? hard to follow what is actually be tested for.
func TestAccSDMRoleAttachment_Update(t *testing.T) {
	rsName := randomWithPrefix("test-update-attachment")
	compRoleRsName := randomWithPrefix("test-comprole")
	roleRsName := randomWithPrefix("test-role")
	altCompRoleRsName := randomWithPrefix("test-comprole2")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleAttachmentConfig(rsName, compRoleRsName, roleRsName),
				Check: func(s *terraform.State) error {
					compositeID, err := testCreatedID(s, "sdm_role", compRoleRsName)
					if err != nil {
						return err
					}

					attachedID, err := testCreatedID(s, "sdm_role", roleRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_role_attachment", rsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.RoleAttachments().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created resource: %w", err)
					}

					if resp.RoleAttachment.CompositeRoleID != compositeID {
						return fmt.Errorf("unexpected compositeID '%s', expected '%s'", resp.RoleAttachment.CompositeRoleID, compositeID)
					}

					if resp.RoleAttachment.AttachedRoleID != attachedID {
						return fmt.Errorf("unexpected attachedID '%s', expected '%s'", resp.RoleAttachment.AttachedRoleID, attachedID)
					}

					return nil
				},
			},
			{
				Config: testAccSDMRoleAttachmentConfig(rsName, altCompRoleRsName, roleRsName),
				Check: func(s *terraform.State) error {
					compositeID, err := testCreatedID(s, "sdm_role", altCompRoleRsName)
					if err != nil {
						return err
					}

					attachedID, err := testCreatedID(s, "sdm_role", roleRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_role_attachment", rsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.RoleAttachments().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created resource: %w", err)
					}

					if resp.RoleAttachment.CompositeRoleID != compositeID {
						return fmt.Errorf("unexpected compositeID '%s', expected '%s'", resp.RoleAttachment.CompositeRoleID, compositeID)
					}

					if resp.RoleAttachment.AttachedRoleID != attachedID {
						return fmt.Errorf("unexpected attachedID '%s', expected '%s'", resp.RoleAttachment.AttachedRoleID, attachedID)
					}

					return nil
				},
			},
		},
	})
}

func testAccSDMRoleAttachmentConfig(roleAttachmentResourceName, compositeRoleResourceName, roleResourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_role" "%[1]s" {
		name = "%[2]s"
		composite = false
	}

	resource "sdm_role" "%[3]s" {
		name = "%[4]s"
		composite = true
	}

	resource "sdm_role_attachment" "%[5]s" {
		composite_role_id = sdm_role.%[6]s.id
		attached_role_id = sdm_role.%[7]s.id
	}`,
		roleResourceName,
		randomWithPrefix("roleName"),
		compositeRoleResourceName,
		randomWithPrefix("compRoleName"),
		roleAttachmentResourceName,
		compositeRoleResourceName,
		roleResourceName,
	)
}

func createRoleAttachmentsWithPrefix(prefix string, count int) ([]*sdm.RoleAttachment, error) {
	roles, err := createRolesWithPrefix(prefix, count, false)
	if err != nil {
		return nil, fmt.Errorf("failed to create accounts: %w", err)
	}
	compositeRoles, err := createRolesWithPrefix(prefix, 1, true)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}
	roleAttachments, err := attachRoles(compositeRoles[0], roles)
	if err != nil {
		return nil, fmt.Errorf("failed to attach accounts: %w", err)
	}

	return roleAttachments, nil
}

func attachRoles(compositeRole *sdm.Role, attachedRoles []*sdm.Role) ([]*sdm.RoleAttachment, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	attachments := []*sdm.RoleAttachment{}
	for _, ar := range attachedRoles {
		createResp, err := client.RoleAttachments().Create(ctx, &sdm.RoleAttachment{
			CompositeRoleID: compositeRole.ID,
			AttachedRoleID:  ar.ID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create account attachment: %w", err)
		}
		attachments = append(attachments, createResp.RoleAttachment)
	}
	return attachments, nil
}
