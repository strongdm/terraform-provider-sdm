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
	resource.AddTestSweepers("sdm_account", &resource.Sweeper{
		Name:         "sdm_account",
		Dependencies: []string{"sdm_account_attachment", "sdm_account_grant"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.Accounts().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing resources: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.Accounts().Delete(ctx, v.GetID())
				if err != nil {
					fmt.Printf("error deleting resource %s %s\n", v.GetID(), err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing accounts %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMAccount_ServiceCreate(t *testing.T) {
	rsName := randomWithPrefix("test")
	serviceName := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountServiceConfig(rsName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_account."+rsName, "service.0.name", serviceName),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "service.#", "1"),
					resource.TestCheckResourceAttrSet("sdm_account."+rsName, "service.0.token"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_account", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Accounts().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						service, ok := resp.Account.(*sdm.Service)
						if !ok {
							return fmt.Errorf("expected service, got user")
						}
						if service.Name != serviceName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", service.Name, serviceName)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_account." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
		},
	})
}

func TestAccSDMAccount_UserCreate(t *testing.T) {
	rsName := randomWithPrefix("test")
	firstName := randomWithPrefix("ursula")
	lastName := randomWithPrefix("leguin")
	email := randomWithPrefix("testsuites@strongdm.com")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountUserConfig(rsName, firstName, lastName, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.first_name", firstName),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.last_name", lastName),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.email", email),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.#", "1"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_account", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Accounts().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						user, ok := resp.Account.(*sdm.User)
						if !ok {
							return fmt.Errorf("expected user, got service")
						}
						if user.FirstName != firstName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", user.FirstName, firstName)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_account." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMAccount_Update(t *testing.T) {
	rsName := randomWithPrefix("test")

	firstName := randomWithPrefix("ursula")
	lastName := randomWithPrefix("leguin")
	email := randomWithPrefix("testsuites@strongdm.com")

	firstName2 := randomWithPrefix("cheese")
	lastName2 := randomWithPrefix("cake")
	email2 := randomWithPrefix("testsuites+1@strongdm.com")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountUserConfig(rsName, firstName, lastName, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.first_name", firstName),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.last_name", lastName),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.email", email),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.#", "1"),
				),
			},
			{
				ResourceName:      "sdm_account." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMAccountUserConfig(rsName, firstName2, lastName2, email2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.first_name", firstName2),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.last_name", lastName2),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.0.email", email2),
					resource.TestCheckResourceAttr("sdm_account."+rsName, "user.#", "1"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_account", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Accounts().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						user, ok := resp.Account.(*sdm.User)
						if !ok {
							return fmt.Errorf("expected user, got service")
						}
						if user.FirstName != firstName2 {
							return fmt.Errorf("unexpected name '%s', expected '%s'", user.FirstName, firstName2)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_account." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMAccount_Tags(t *testing.T) {
	name := randomWithPrefix("test")
	firstName := randomWithPrefix("ursula")
	lastName := randomWithPrefix("leguin")
	email := randomWithPrefix("testsuites@strongdm.com")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountTagsConfig(name, firstName, lastName, email, sdm.Tags{"key": "value", "foo": "bar"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_account."+name, "user.0.tags.key", "value"),
					resource.TestCheckResourceAttr("sdm_account."+name, "user.0.tags.foo", "bar"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_account", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Accounts().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created account: %w", err)
						}

						if resp.Account.GetTags()["key"] != "value" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Account.GetTags()["key"])
						}
						if resp.Account.GetTags()["foo"] != "bar" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Account.GetTags()["key"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_account." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMAccountTagsConfig(name, firstName, lastName, email, sdm.Tags{"goat": "bananas"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sdm_account."+name, "user.0.tags.key"),
					resource.TestCheckNoResourceAttr("sdm_account."+name, "user.0.tags.foo"),
					resource.TestCheckResourceAttr("sdm_account."+name, "user.0.tags.goat", "bananas"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_account", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Accounts().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created account: %w", err)
						}

						if resp.Account.GetTags()["goat"] != "bananas" {
							return fmt.Errorf("unexpected value '%s' for tag 'goat', expected 'bananas'", resp.Account.GetTags()["goat"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_account." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSDMAccountTagsConfig(resourceName string, firstName, lastName, email string, tags sdm.Tags) string {
	return fmt.Sprintf(`
	resource "sdm_account" "%s" {
		user {
			first_name = "%s"
			last_name = "%s"
			email = "%s"
			tags = {
%s
			}
		}
	}
	`, resourceName, firstName, lastName, email, tagsToConfigString(tags))
}

func testAccSDMAccountUserConfig(resourceName, firstName, lastName, email string) string {

	return fmt.Sprintf(`
	resource "sdm_account" "%s" {
		user {
			first_name = "%s"
			last_name = "%s"
			email = "%s"
		}
	}
	`, resourceName, firstName, lastName, email)
}

func testAccSDMAccountServiceConfig(resourceName, serviceName string) string {
	return fmt.Sprintf(`
	resource "sdm_account" "%s" {
		service {
			name = "%s"
		}
	}`, resourceName, serviceName)
}

func createUsersWithPrefix(prefix string, count int) ([]sdm.Account, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	accounts := []sdm.Account{}
	for i := 0; i < count; i++ {
		createResp, err := client.Accounts().Create(ctx, &sdm.User{
			FirstName: randomWithPrefix(prefix),
			LastName:  randomWithPrefix(prefix),
			Email:     randomWithPrefix(prefix),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create account: %w", err)
		}
		accounts = append(accounts, createResp.Account)
	}
	return accounts, nil
}
