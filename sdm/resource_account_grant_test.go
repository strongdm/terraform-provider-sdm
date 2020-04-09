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
	resource.AddTestSweepers("sdm_account_grant", &resource.Sweeper{
		Name: "sdm_account_grant",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			accountResp, err := client.Accounts().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing accounts: %s", err)
			}

			for accountResp.Next() {
				account := accountResp.Value()
				agResp, err := client.AccountGrants().List(ctx, "accountid:?", account.GetID())
				if err != nil {
					return fmt.Errorf("Error listing account grants: %s", err)
				}
				for agResp.Next() {
					v := agResp.Value()
					_, err := client.AccountGrants().Delete(ctx, v.ID)
					if err != nil {
						fmt.Printf("error deleting account grant %s %s\n", v.ID, err)
					}
				}
				if agResp.Err() != nil {
					fmt.Printf("error after listing account grants %s", agResp.Err())
				}
			}
			if accountResp.Err() != nil {
				return fmt.Errorf("error listing accounts %s", accountResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMAccountGrant_Create(t *testing.T) {
	grantRsName := randomWithPrefix("test-grant")
	redisRsName := randomWithPrefix("test-redis")
	accountRsName := randomWithPrefix("test-account")
	accountID := ""
	resourceID := ""
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGrantConfig(grantRsName, redisRsName, accountRsName),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						var err error
						accountID, err = testCreatedID(s, "sdm_account", accountRsName)
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
						if err := resource.TestCheckResourceAttr("sdm_account_grant."+grantRsName, "account_id", accountID)(s); err != nil {
							return err
						}
						if err := resource.TestCheckResourceAttr("sdm_account_grant."+grantRsName, "resource_id", resourceID)(s); err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {

						id, err := testCreatedID(s, "sdm_account_grant", grantRsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.AccountGrants().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.AccountGrant.AccountID != accountID {
							return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.AccountGrant.AccountID, accountID)
						}

						if resp.AccountGrant.ResourceID != resourceID {
							return fmt.Errorf("unexpected resourceID '%s', expected '%s'", resp.AccountGrant.ResourceID, resourceID)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_account_grant." + grantRsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMAccountGrant_Update(t *testing.T) {
	grantRsName := randomWithPrefix("test-update-grant")
	redisRsName := randomWithPrefix("test-redis")
	accountRsName := randomWithPrefix("test-account")
	altAccountRsName := randomWithPrefix("test-updated-account")

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGrantConfig(grantRsName, redisRsName, accountRsName),
				Check: func(s *terraform.State) error {
					accountID, err := testCreatedID(s, "sdm_account", accountRsName)
					if err != nil {
						return err
					}

					resourceID, err := testCreatedID(s, "sdm_resource", redisRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_account_grant", grantRsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.AccountGrants().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created resource: %w", err)
					}

					if resp.AccountGrant.AccountID != accountID {
						return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.AccountGrant.AccountID, accountID)
					}

					if resp.AccountGrant.ResourceID != resourceID {
						return fmt.Errorf("unexpected resourceID '%s', expected '%s'", resp.AccountGrant.ResourceID, resourceID)
					}

					return nil
				},
			},
			{
				ResourceName:      "sdm_account_grant." + grantRsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMAccountGrantConfig(grantRsName, redisRsName, altAccountRsName),
				Check: func(s *terraform.State) error {
					accountID, err := testCreatedID(s, "sdm_account", altAccountRsName)
					if err != nil {
						return err
					}

					resourceID, err := testCreatedID(s, "sdm_resource", redisRsName)
					if err != nil {
						return err
					}

					id, err := testCreatedID(s, "sdm_account_grant", grantRsName)
					if err != nil {
						return err
					}

					// check if it was actually created
					client := testClient()
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					resp, err := client.AccountGrants().Get(ctx, id)
					if err != nil {
						return fmt.Errorf("failed to get created resource: %w", err)
					}

					if resp.AccountGrant.AccountID != accountID {
						return fmt.Errorf("unexpected accountID '%s', expected '%s'", resp.AccountGrant.AccountID, accountID)
					}

					if resp.AccountGrant.ResourceID != resourceID {
						return fmt.Errorf("unexpected resourceID '%s', expected '%s'", resp.AccountGrant.ResourceID, resourceID)
					}

					return nil
				},
			},
			{
				ResourceName:      "sdm_account_grant." + grantRsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSDMAccountGrantConfig(grantResourceName, redisResourceName, accResourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			port_override = %d
		}
	}

	resource "sdm_account" "%s" {
		user {
			first_name = "%s"
			last_name = "%s"
			email = "%s"
		}
	}

	resource "sdm_account_grant" "%s" {
		account_id = sdm_account.%s.id
		resource_id = sdm_resource.%s.id
	}
	`,
		redisResourceName,
		randomWithPrefix("redis"),
		portOverride.Count(),
		accResourceName,
		randomWithPrefix("first-name"),
		randomWithPrefix("last-name"),
		randomWithPrefix("email"),
		grantResourceName,
		accResourceName,
		redisResourceName,
	)
}

func createAccountGrantsWithPrefix(prefix string, count int) ([]*sdm.AccountGrant, error) {
	resources, err := createRedisesWithPrefix(prefix, count)
	if err != nil {
		return nil, fmt.Errorf("failed to create accounts: %w", err)
	}
	users, err := createUsersWithPrefix(prefix, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}
	accountGrants, err := grantAccount(users[0], resources)
	if err != nil {
		return nil, fmt.Errorf("failed to grant accounts: %w", err)
	}

	return accountGrants, nil
}

func grantAccount(account sdm.Account, resources []sdm.Resource) ([]*sdm.AccountGrant, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	grants := []*sdm.AccountGrant{}
	for _, res := range resources {
		createResp, err := client.AccountGrants().Create(ctx, &sdm.AccountGrant{
			AccountID:  account.GetID(),
			ResourceID: res.GetID(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create account grant: %w", err)
		}
		grants = append(grants, createResp.AccountGrant)
	}
	return grants, nil
}
