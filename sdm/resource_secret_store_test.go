package sdm

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	sdm "github.com/strongdm/web/pkg/api/v1/generated/go"
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
				return fmt.Errorf("Error listing secretstores: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.SecretStores().Delete(ctx, v.GetID())
				if err != nil {
					fmt.Printf("error deleting secret store %s %s\n", v.GetID(), err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing secret store %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMSecretStore_Create(t *testing.T) {
	name := randomWithPrefix("test")
	address := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreVaultTokenConfig(name, name, address),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "vault_token.0.name", name),
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "vault_token.0.server_address", address),
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
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.SecretStore.(*sdm.VaultTokenStore).Name != name {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.(*sdm.VaultTokenStore).Name, name)
						}
						if resp.SecretStore.(*sdm.VaultTokenStore).ServerAddress != address {
							return fmt.Errorf("unexpected address '%s', expected '%s'", resp.SecretStore.(*sdm.VaultTokenStore).ServerAddress, address)
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

func TestAccSDMSecretStore_Tags(t *testing.T) {
	name := randomWithPrefix("test")
	address := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreTagsConfig(name, name, address, sdm.Tags{"key": "value", "foo": "bar"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "vault_token.0.tags.key", "value"),
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "vault_token.0.tags.foo", "bar"),
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
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.SecretStore.GetTags()["key"] != "value" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.SecretStore.GetTags()["key"])
						}
						if resp.SecretStore.GetTags()["foo"] != "bar" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.SecretStore.GetTags()["key"])
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
				Config: testAccSDMSecretStoreTagsConfig(name, name, address, sdm.Tags{"goat": "bananas"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sdm_secret_store."+name, "vault_token.0.tags.key"),
					resource.TestCheckNoResourceAttr("sdm_secret_store."+name, "vault_token.0.tags.foo"),
					resource.TestCheckResourceAttr("sdm_secret_store."+name, "vault_token.0.tags.goat", "bananas"),
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
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.SecretStore.GetTags()["goat"] != "bananas" {
							return fmt.Errorf("unexpected value '%s' for tag 'goat', expected 'bananas'", resp.SecretStore.GetTags()["goat"])
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

func TestAccSDMSecretStore_Update(t *testing.T) {
	resourceName := randomWithPrefix("test")

	storeName := randomWithPrefix("test")
	updatedStoreName := randomWithPrefix("test2")

	address := randomWithPrefix("test")
	updatedAddress := randomWithPrefix("test2")

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreVaultTokenConfig(resourceName, storeName, address),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+resourceName, "vault_token.0.name", storeName),
					resource.TestCheckResourceAttr("sdm_secret_store."+resourceName, "vault_token.0.server_address", address),
				),
			},
			{
				ResourceName:      "sdm_secret_store." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMSecretStoreVaultTokenConfig(resourceName, updatedStoreName, updatedAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_secret_store."+resourceName, "vault_token.0.name", updatedStoreName),
					resource.TestCheckResourceAttr("sdm_secret_store."+resourceName, "vault_token.0.server_address", updatedAddress),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_secret_store", resourceName)
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

						if resp.SecretStore.(*sdm.VaultTokenStore).Name != updatedStoreName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.SecretStore.(*sdm.VaultTokenStore).Name, updatedStoreName)
						}

						if resp.SecretStore.(*sdm.VaultTokenStore).ServerAddress != updatedAddress {
							return fmt.Errorf("unexpected address '%s', expected '%s'", resp.SecretStore.(*sdm.VaultTokenStore).ServerAddress, updatedAddress)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_secret_store." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMSecretStore_UpdateAllTypes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		resource string
		pairs    [][2]string
	}
	tcs := []testCase{
		{
			resource: "vault_token",
			pairs: [][2]string{
				{"name", `"vault_token"`},
				{"server_address", `"ServerAddress"`},
			},
		},
		{
			resource: "vault_tls",
			pairs: [][2]string{
				{"name", `"vault_tls"`},
				{"server_address", `"ServerAddress"`},
				{"ca_cert_path", `"CACertPath"`},
				{"client_cert_path", `"ClientCertPath"`},
				{"client_key_path", `"ClientKeyPath"`},
			},
		},
	}

	resourceNameBase := randomWithPrefix("test")

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.resource, func(t *testing.T) {
			name := resourceNameBase + tc.resource
			cfg := testAccSDMSecretStoreAnyConfig(name, tc.resource, tc.pairs)

			checks := make([]resource.TestCheckFunc, len(tc.pairs))
			for i, p := range tc.pairs {
				val := p[1]
				// TF removes quotes around strings
				val = strings.Trim(val, "\"")
				// ... and converts escaped new lines into real newlines
				val = strings.Replace(val, "\\n", "\n", -1)
				checks[i] = resource.TestCheckResourceAttr("sdm_secret_store."+name, tc.resource+".0."+p[0], val)
			}

			resource.ParallelTest(t, resource.TestCase{
				Providers:    testAccProviders,
				CheckDestroy: testCheckDestroy,
				Steps: []resource.TestStep{
					{
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						// there should be no change if the resource is updated from the server
						Taint:  []string{"sdm_secret_store." + name},
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						ResourceName: "sdm_secret_store." + name,
						ImportState:  true,
					},
				},
			})
		})
	}
}

func testAccSDMSecretStoreAnyConfig(resourceName, resourceType string, pairs [][2]string) string {
	s := fmt.Sprintf(`
	resource "sdm_secret_store" "%s" {
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

func testAccSDMSecretStoreVaultTokenConfig(resourceName, sdmResourceName, address string) string {
	return fmt.Sprintf(`
	resource "sdm_secret_store" "%s" {
		vault_token {
			name = "%s"
			server_address = "%s"
		}
	}
	`, resourceName, sdmResourceName, address)
}

func testAccSDMSecretStoreTagsConfig(resourceName, sdmResourceName, address string, tags sdm.Tags) string {
	return fmt.Sprintf(`
	resource "sdm_secret_store" "%s" {
		vault_token {
			name = "%s"
			server_address = "%s"
			tags = {
%s
			}
		}
	}
	`, resourceName, sdmResourceName, address, tagsToConfigString(tags))
}

func createVaultTokenStoresWithPrefix(prefix string, count int) ([]sdm.SecretStore, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	stores := []sdm.SecretStore{}
	for i := 0; i < count; i++ {
		createResp, err := client.SecretStores().Create(ctx, &sdm.VaultTokenStore{
			Name:          randomWithPrefix(prefix),
			ServerAddress: randomWithPrefix(prefix),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create vault token store: %w", err)
		}
		stores = append(stores, createResp.SecretStore)
	}
	return stores, nil
}
