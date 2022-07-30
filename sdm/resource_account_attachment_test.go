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
	resource.AddTestSweepers("sdm_account_attachment", &resource.Sweeper{
		Name: "sdm_account_attachment",
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
				aaResp, err := client.AccountAttachments().List(ctx, "roleid:?", r.ID)
				if err != nil {
					return fmt.Errorf("Error listing account attachments: %s", err)
				}
				for aaResp.Next() {
					v := aaResp.Value()
					_, err := client.AccountAttachments().Delete(ctx, v.ID)
					if err != nil {
						fmt.Printf("error deleting account attachment %s %s\n", v.ID, err)
					}
				}
				if aaResp.Err() != nil {
					fmt.Printf("error listing account attachments %s", aaResp.Err())
				}
			}
			if roleResp.Err() != nil {
				return fmt.Errorf("error listing roles %s", roleResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMAccountAttachment_Create(t *testing.T) {
	rsName := randomWithPrefix("test-attachment-create")
	roleRsName := randomWithPrefix("test-role")
	accountRsName := randomWithPrefix("test-account")
	accountID := ""
	roleID := ""
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountAttachmentConfig(rsName, roleRsName, accountRsName),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						var err error
						accountID, err = testCreatedID(s, "sdm_account", accountRsName)
						if err != nil {
							return err
						}

						roleID, err = testCreatedID(s, "sdm_role", roleRsName)
						if err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_account_attachment."+rsName, "account_id", accountID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_account_attachment."+rsName, "role_id", roleID)(s); err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {

						id, err := testCreatedID(s, "sdm_account_attachment", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.AccountAttachments().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created role: %w", err)
						}

						if resp.AccountAttachment.AccountID != accountID {
							return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.AccountAttachment.AccountID, accountID)
						}

						if resp.AccountAttachment.RoleID != roleID {
							return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.AccountAttachment.RoleID, roleID)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_account_attachment." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMAccountAttachment_Update(t *testing.T) {
	rsName := randomWithPrefix("test-update-attachment")
	roleRsName := randomWithPrefix("test-role")
	accountRsName := randomWithPrefix("test-account")
	altAccountRsName := randomWithPrefix("test-account-updated")

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountAttachmentConfig(rsName, roleRsName, accountRsName),
				Check: func(s *terraform.State) error {
					accountID, err := testCreatedID(s, "sdm_account", accountRsName)
					if err != nil {
						return err
					}

					roleID, err := testCreatedID(s, "sdm_role", roleRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_account_attachment", rsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.AccountAttachments().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created role: %w", err)
					}

					if resp.AccountAttachment.AccountID != accountID {
						return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.AccountAttachment.AccountID, accountID)
					}

					if resp.AccountAttachment.RoleID != roleID {
						return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.AccountAttachment.RoleID, roleID)
					}

					return nil
				},
			},
			{
				ResourceName:      "sdm_account_attachment." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMAccountAttachmentConfig(rsName, roleRsName, altAccountRsName),
				Check: func(s *terraform.State) error {
					accountID, err := testCreatedID(s, "sdm_account", altAccountRsName)
					if err != nil {
						return err
					}

					roleID, err := testCreatedID(s, "sdm_role", roleRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_account_attachment", rsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.AccountAttachments().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created role: %w", err)
					}

					if resp.AccountAttachment.AccountID != accountID {
						return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.AccountAttachment.AccountID, accountID)
					}

					if resp.AccountAttachment.RoleID != roleID {
						return fmt.Errorf("unexpected roleID '%s', expected '%s'", resp.AccountAttachment.RoleID, roleID)
					}

					return nil
				},
			},
			{
				ResourceName:      "sdm_account_attachment." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSDMAccountAttachmentConfig(attachmentResourceName, roleResourceName, accResourceName string) string {
	return fmt.Sprintf(
		`
	resource "sdm_role" "%[1]s" {
		name = "%[2]s"
	}

	resource "sdm_account" "%[3]s" {
		user {
			first_name = "%[4]s"
			last_name = "%[5]s"
			email = "%[6]s"
		}
	}

	resource "sdm_account_attachment" "%[7]s" {
		account_id = sdm_account.%[3]s.id
		role_id = sdm_role.%[1]s.id
	}
	`,
		roleResourceName,
		randomWithPrefix("role-name"),
		accResourceName,
		randomWithPrefix("first-name"),
		randomWithPrefix("last-name"),
		randomWithPrefix("email"),
		attachmentResourceName,
	)
}

func createAccountAttachmentsWithPrefix(prefix string, count int) ([]*sdm.AccountAttachment, error) {
	acs, err := createUsersWithPrefix(prefix, count)
	if err != nil {
		return nil, fmt.Errorf("failed to create accounts: %w", err)
	}
	roles, err := createRolesWithPrefix(prefix, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}
	acAttachments, err := attachAccounts(roles[0], acs)
	if err != nil {
		return nil, fmt.Errorf("failed to attach accounts: %w", err)
	}

	return acAttachments, nil
}

func attachAccounts(role *sdm.Role, accounts []sdm.Account) ([]*sdm.AccountAttachment, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	attachments := []*sdm.AccountAttachment{}
	for _, a := range accounts {
		createResp, err := client.AccountAttachments().Create(ctx, &sdm.AccountAttachment{
			AccountID: a.GetID(),
			RoleID:    role.ID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create account attachment: %w", err)
		}
		attachments = append(attachments, createResp.AccountAttachment)
	}
	return attachments, nil
}
