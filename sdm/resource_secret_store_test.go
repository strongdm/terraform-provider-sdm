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
	resource.AddTestSweepers("sdm_secret_store", &resource.Sweeper{
		Name: "sdm_secret_store",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.SecretStores().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing secret stores: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.SecretStores().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting secret stores %s %s\n", v.ID, err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing secret stores %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMSecretStore_Create(t *testing.T) {
	rsName := randomWithPrefix("test-secretstore-resource")
	seName := randomWithPrefix("test-secretstore")
	address := randomWithPrefix("test-secretstore")
	kind := "vault"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreConfig(rsName, seName, kind, address),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "name", seName),
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "kind", kind),
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "server_address", address),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_secret_store", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.SecretStores().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.SecretStore.Name != seName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.Name, seName)
						}
						if resp.SecretStore.Kind != kind {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.Name, kind)
						}
						if resp.SecretStore.ServerAddress != address {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.Name, address)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_secret_store." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMSecretStore_Update(t *testing.T) {
	rsName := randomWithPrefix("test-secretstore")
	seName := randomWithPrefix("test-secretstore")
	updatedSecretStoreName := randomWithPrefix("test-secretstore-updated")
	address := randomWithPrefix("test-secretstore")
	kind := "vault"
	updatedAddress := randomWithPrefix("test-secretstore-updated")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreConfig(rsName, seName, kind, address),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "name", seName),
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "kind", kind),
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "server_address", address),
				),
			},
			{
				ResourceName:      "sdm_secret_store." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMSecretStoreConfig(rsName, updatedSecretStoreName, kind, updatedAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "name", updatedSecretStoreName),
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "kind", kind),
					resource.TestCheckResourceAttr("sdm_secret_store."+rsName, "server_address", updatedAddress),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_secret_store", rsName)
						if err != nil {
							return err
						}

						// check if it was actually updated
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.SecretStores().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.SecretStore.Name != updatedSecretStoreName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.Name, updatedSecretStoreName)
						}
						if resp.SecretStore.Kind != kind {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.Name, kind)
						}
						if resp.SecretStore.ServerAddress != updatedAddress {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.Name, updatedAddress)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_secret_store." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMSecretStore_Tags(t *testing.T) {
	name := randomWithPrefix("test")
	address := randomWithPrefix("test-secretstore")
	kind := "vault"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreTagsConfig(name, name, kind, address, sdm.Tags{"key": "value", "foo": "bar"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "tags.key", "value"),
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "tags.foo", "bar"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_secret_store", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.SecretStores().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created secret store: %w", err)
						}

						if resp.SecretStore.Tags["key"] != "value" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.SecretStore.Tags["key"])
						}
						if resp.SecretStore.Tags["foo"] != "bar" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.SecretStore.Tags["key"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_secret_store." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMSecretStoreTagsConfig(name, name, kind, address, sdm.Tags{"goat": "bananas"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sdm_secret_store."+name, "tags.key"),
					resource.TestCheckNoResourceAttr("sdm_secret_store."+name, "tags.foo"),
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "tags.goat", "bananas"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_secret_store", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.SecretStores().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created secret store: %w", err)
						}

						if resp.SecretStore.Tags["goat"] != "bananas" {
							return fmt.Errorf("unexpected value '%s' for tag 'goat', expected 'bananas'", resp.SecretStore.Tags["goat"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_secret_store." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSDMSecretStoreConfig(resourceName, seName, kind, serverAddress string) string {
	return fmt.Sprintf(`
	resource "sdm_secret_store" "%s" {
		name = "%s"
		kind = "%s"
		server_address = "%s"
	}
	`, resourceName, seName, kind, serverAddress)
}

func testAccSDMSecretStoreTagsConfig(resourceName string, seName, kind, serverAddress string, tags sdm.Tags) string {
	return fmt.Sprintf(`
	resource "sdm_secret_store" "%s" {
		name = "%s"
		kind = "%s"
		server_address = "%s"
		tags = {
%s
		}
	}
	`, resourceName, seName, kind, serverAddress, tagsToConfigString(tags))
}

func createSecretStoresWithPrefix(prefix string, count int) ([]*sdm.SecretStore, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	secretStores := []*sdm.SecretStore{}
	for i := 0; i < count; i++ {
		createResp, err := client.SecretStores().Create(ctx, &sdm.SecretStore{
			Name:          randomWithPrefix(prefix),
			Kind:          "vault",
			ServerAddress: randomWithPrefix(prefix),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create secret store: %w", err)
		}
		secretStores = append(secretStores, createResp.SecretStore)
	}
	return secretStores, nil
}
