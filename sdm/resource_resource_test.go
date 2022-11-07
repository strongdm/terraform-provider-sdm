package sdm

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

var portOverride = NewAtomicCounter(32000)

func init() {
	resource.AddTestSweepers("sdm_resource", &resource.Sweeper{
		Name:         "sdm_resource",
		Dependencies: []string{"sdm_account_grant", "sdm_role_grant"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.Resources().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing resources: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.Resources().Delete(ctx, v.GetID())
				if err != nil {
					fmt.Printf("error deleting resource %s %s\n", v.GetID(), err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing resource %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMResource_Create(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	port := portOverride.Count()
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceRedisConfig(name, name, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.name", name),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttrSet("sdm_resource."+name, "redis.0.port_override"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.bind_interface", "127.0.0.1"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.Redis).Name != name {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Resource.(*sdm.Redis).Name, name)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_CreateWithBindInterface(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	port := portOverride.Count()
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "sdm_resource" "%s" {
					redis {
						name = "%s"
						hostname = "test.com"
						port = %d
						bind_interface = "127.0.0.2"
					}
				}
				`, name, name, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.name", name),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttrSet("sdm_resource."+name, "redis.0.port_override"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.bind_interface", "127.0.0.2"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.Redis).Name != name {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Resource.(*sdm.Redis).Name, name)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_CreateWithSecretStore(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	port := portOverride.Count()

	vaults, err := createVaultTokenStoresWithPrefix("vaultTest", 1)
	if err != nil {
		t.Fatalf("failed to create secret store: %v", err)
	}
	vault := vaults[0]
	path := "/path/to/secret"
	key := "password2"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceRedisSecretStoreConfig(name, name, port, vault.GetID(), path, key),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.name", name),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.secret_store_password_path", path),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.secret_store_password_key", key),
					resource.TestCheckResourceAttrSet("sdm_resource."+name, "redis.0.port_override"),
					resource.TestCheckResourceAttrSet("sdm_resource."+name, "redis.0.secret_store_id"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.Redis).SecretStoreID != vault.GetID() {
							return fmt.Errorf("unexpected secret store id '%s', expected '%s'", resp.Resource.(*sdm.Redis).SecretStoreID, vault.GetID())
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"redis.0.secret_store_password_path",
					"redis.0.secret_store_password_key",
				},
			},
		},
	})
}

func TestAccSDMResource_CreateWithSecretStoreNoSecretStoreID(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	port := portOverride.Count()

	path := "/path/to/secret"
	key := "password2"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccSDMResourceRedisSecretStoreConfig(name, name, port, "", path, key),
				ExpectError: regexp.MustCompile("secret store credential secret_store_password_path must be combined with secret_store_id"),
			},
		},
	})
}

func TestAccSDMResource_CreateWithSecretStoreNoRawCredentials(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	port := portOverride.Count()

	vaults, err := createVaultTokenStoresWithPrefix("vaultTest", 1)
	if err != nil {
		t.Fatalf("failed to create secret store: %v", err)
	}
	vault := vaults[0]

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccSDMResourceRedisSecretStoreRawCredentialConfig(name, name, port, vault.GetID(), "password"),
				ExpectError: regexp.MustCompile("raw credential password cannot be combined with secret_store_id"),
			},
		},
	})
}

func TestAccSDMResource_Tags(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceTagsConfig(name, name, sdm.Tags{"key": "value", "foo": "bar"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.tags.key", "value"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.tags.foo", "bar"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.GetTags()["key"] != "value" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Resource.GetTags()["key"])
						}
						if resp.Resource.GetTags()["foo"] != "bar" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Resource.GetTags()["key"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMResourceTagsConfig(name, name, sdm.Tags{"goat": "bananas"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sdm_resource."+name, "redis.0.tags.key"),
					resource.TestCheckNoResourceAttr("sdm_resource."+name, "redis.0.tags.foo"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.tags.goat", "bananas"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.GetTags()["goat"] != "bananas" {
							return fmt.Errorf("unexpected value '%s' for tag 'goat', expected 'bananas'", resp.Resource.GetTags()["goat"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_CreateSSH(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	port := portOverride.Count()
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceSSHConfig(name, name, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.name", name),
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.username", "user"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.port", fmt.Sprint(port)),
					resource.TestCheckResourceAttrSet("sdm_resource."+name, "ssh.0.public_key"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.SSH).Name != name {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Resource.(*sdm.SSH).Name, name)
						}

						return nil
					},
				),
			},
			{
				ResourceName: "sdm_resource." + name,
				ImportState:  true,
				// We can't verify the import state because the username on ssh that gets returned
				// could either be a secret store path or a raw username
				//ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_Update(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("test")

	redisName := randomWithPrefix("test")
	updatedRedisName := randomWithPrefix("test2")

	port := portOverride.Count()
	updatedPort := portOverride.Count()

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "sdm_resource" "%s" {
					redis {
						name = "%s"
						hostname = "test.com"
						port = %d
					}
				}
				`, resourceName, redisName, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port", fmt.Sprint(port)),
					resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.1"),
				),
			},
			{
				ResourceName:      "sdm_resource." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
				resource "sdm_resource" "%s" {
					redis {
						name = "%s"
						hostname = "test.com"
						port = %d
						bind_interface = "127.0.0.2"
					}
				}
				`, resourceName, updatedRedisName, updatedPort),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", updatedRedisName),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port", fmt.Sprint(updatedPort)),
					resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.2"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", resourceName)
						if err != nil {
							return err
						}

						// check if it was actually updated
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.Redis).Name != updatedRedisName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Resource.(*sdm.Redis).Name, updatedRedisName)
						}

						if resp.Resource.(*sdm.Redis).Port != updatedPort {
							return fmt.Errorf("unexpected port '%d', expected '%d'", resp.Resource.(*sdm.Redis).Port, updatedPort)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_UpdatePortOverrides(t *testing.T) {
	initAcceptanceTest(t)
	t.Run("present->present", func(t *testing.T) {
		initAcceptanceTest(t)
		resourceName := randomWithPrefix("test")
		redisName := randomWithPrefix("test")
		portOverride, updatedPortOverride := portOverride.Count(), portOverride.Count()

		resource.Test(t, resource.TestCase{
			Providers:    testAccProviders,
			CheckDestroy: testCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
							port_override = %v
						}
					}
					`, resourceName, redisName, portOverride),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port_override", fmt.Sprint(portOverride)),
						resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.1"),
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
							bind_interface = "127.0.0.2"
							port_override = %v
						}
					}
					`, resourceName, redisName, updatedPortOverride),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port_override", fmt.Sprint(updatedPortOverride)),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.2"),
						func(s *terraform.State) error {
							id, err := testCreatedID(s, "sdm_resource", resourceName)
							if err != nil {
								return err
							}

							// check if it was actually updated
							client := testClient()
							ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
							defer cancel()
							resp, err := client.Resources().Get(ctx, id)
							if err != nil {
								return fmt.Errorf("failed to get created resource: %w", err)
							}

							if resp.Resource.(*sdm.Redis).PortOverride != updatedPortOverride {
								return fmt.Errorf("unexpected port override '%d', expected '%d'", resp.Resource.(*sdm.Redis).PortOverride, updatedPortOverride)
							}

							return nil
						},
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("absent->present", func(t *testing.T) {
		initAcceptanceTest(t)
		resourceName := randomWithPrefix("test")
		redisName := randomWithPrefix("test")
		updatedPortOverride := portOverride.Count()

		resource.Test(t, resource.TestCase{
			Providers:    testAccProviders,
			CheckDestroy: testCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
						}
					}
					`, resourceName, redisName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.1"),
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
							bind_interface = "127.0.0.2"
							port_override = %v
						}
					}
					`, resourceName, redisName, updatedPortOverride),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port_override", fmt.Sprint(updatedPortOverride)),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.2"),
						func(s *terraform.State) error {
							id, err := testCreatedID(s, "sdm_resource", resourceName)
							if err != nil {
								return err
							}

							// check if it was actually updated
							client := testClient()
							ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
							defer cancel()
							resp, err := client.Resources().Get(ctx, id)
							if err != nil {
								return fmt.Errorf("failed to get created resource: %w", err)
							}

							if resp.Resource.(*sdm.Redis).PortOverride != updatedPortOverride {
								return fmt.Errorf("unexpected port override '%d', expected '%d'", resp.Resource.(*sdm.Redis).PortOverride, updatedPortOverride)
							}

							return nil
						},
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("present->absent", func(t *testing.T) {
		initAcceptanceTest(t)
		resourceName := randomWithPrefix("test")
		redisName := randomWithPrefix("test")
		portOverride := portOverride.Count()

		resource.Test(t, resource.TestCase{
			Providers:    testAccProviders,
			CheckDestroy: testCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
							port_override = %v
						}
					}
					`, resourceName, redisName, portOverride),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port_override", fmt.Sprint(portOverride)),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.1"),
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
							bind_interface = "127.0.0.2"
						}
					}
					`, resourceName, redisName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port_override", fmt.Sprint(portOverride)),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.2"),
						func(s *terraform.State) error {
							id, err := testCreatedID(s, "sdm_resource", resourceName)
							if err != nil {
								return err
							}

							// check if it was actually updated
							client := testClient()
							ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
							defer cancel()
							resp, err := client.Resources().Get(ctx, id)
							if err != nil {
								return fmt.Errorf("failed to get created resource: %w", err)
							}

							if resp.Resource.(*sdm.Redis).PortOverride != portOverride {
								return fmt.Errorf("unexpected port override '%d', expected '%d'", resp.Resource.(*sdm.Redis).PortOverride, portOverride)
							}

							return nil
						},
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("absent->absent", func(t *testing.T) {
		initAcceptanceTest(t)
		resourceName := randomWithPrefix("test")
		redisName := randomWithPrefix("test")

		resource.Test(t, resource.TestCase{
			Providers:    testAccProviders,
			CheckDestroy: testCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
						}
					}
					`, resourceName, redisName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.1"),
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					Config: fmt.Sprintf(`
					resource "sdm_resource" "%s" {
						redis {
							name = "%s"
							hostname = "test.com"
							port = 6379
							bind_interface = "127.0.0.2"
						}
					}
					`, resourceName, redisName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
						resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
						resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.bind_interface", "127.0.0.2"),
						func(s *terraform.State) error {
							id, err := testCreatedID(s, "sdm_resource", resourceName)
							if err != nil {
								return err
							}

							// check if it was actually updated
							client := testClient()
							ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
							defer cancel()
							resp, err := client.Resources().Get(ctx, id)
							if err != nil {
								return fmt.Errorf("failed to get created resource: %w", err)
							}

							if resp.Resource.(*sdm.Redis).PortOverride == 0 {
								return fmt.Errorf("missing port override '%d'", resp.Resource.(*sdm.Redis).PortOverride)
							}

							return nil
						},
					),
				},
				{
					ResourceName:      "sdm_resource." + resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func testAccSDMResourceAnyConfig(resourceName, resourceType string, pairs [][2]string) string {
	s := fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		%s {`, resourceName, resourceType)
	for _, p := range pairs {
		s += "\n"
		s += p[0] + " = " + p[1]
	}
	s += "\n"
	s += `}
	}`
	return s
}

func testAccSDMResourceRedisConfig(resourceName string, sdmResourceName string, port int32) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			port = %d
		}
	}
	`, resourceName, sdmResourceName, port)
}

func testAccSDMResourceRedisSecretStoreConfig(resourceName, sdmResourceName string, port int32, seID, path, key string) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			port = %d
			secret_store_id = "%s"
			secret_store_password_path = "%s"
			secret_store_password_key = "%s"
		}
	}
	`, resourceName, sdmResourceName, port, seID, path, key)
}

func testAccSDMResourceRedisSecretStoreRawCredentialConfig(resourceName, sdmResourceName string, port int32, seID, password string) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			port = %d
			secret_store_id = "%s"
			password = "%s"
		}
	}
	`, resourceName, sdmResourceName, port, seID, password)
}

func tagsToConfigString(tags sdm.Tags) string {
	tagString := ""
	for key, value := range tags {
		tagString += fmt.Sprintf("\t\t\t\t%s = \"%s\"\n", key, value)
	}
	return tagString
}

func testAccSDMResourceTagsConfig(resourceName string, sdmResourceName string, tags sdm.Tags) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			tags = {
%s
			}
		}
	}
	`, resourceName, sdmResourceName, tagsToConfigString(tags))
}

func testAccSDMResourceSSHConfig(resourceName string, sdmResourceName string, port int32) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		ssh {
			name = "%s"
			hostname = "test.com"
			username = "user"
			port = %d
			key_type = "ed25519"
		}
	}
	`, resourceName, sdmResourceName, port)
}

func createRedisesWithPrefix(prefix string, count int) ([]sdm.Resource, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resources := []sdm.Resource{}
	for i := 0; i < count; i++ {
		port := portOverride.Count()
		createResp, err := client.Resources().Create(ctx, &sdm.Redis{
			Name:         randomWithPrefix(prefix),
			Hostname:     randomWithPrefix(prefix),
			Port:         port,
			PortOverride: port,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create redis: %w", err)
		}
		resources = append(resources, createResp.Resource)
	}
	return resources, nil
}

func createSSHesWithPrefix(prefix string, count int) ([]sdm.Resource, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resources := []sdm.Resource{}
	for i := 0; i < count; i++ {
		port := portOverride.Count()
		createResp, err := client.Resources().Create(ctx, &sdm.SSH{
			Name:     randomWithPrefix(prefix),
			Hostname: randomWithPrefix(prefix),
			Username: randomWithPrefix(prefix),
			Port:     port,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create ssh: %w", err)
		}
		resources = append(resources, createResp.Resource)
	}
	return resources, nil
}
