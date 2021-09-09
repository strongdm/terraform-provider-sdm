package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMSecretStore_Get(t *testing.T) {
	t.Parallel()

	stores, err := createVaultTokenStoresWithPrefix("secret-store-get-test", 1)
	if err != nil {
		t.Fatal("failed to create redis in setup: ", err)
	}
	store := stores[0].(*sdm.VaultTokenStore)

	dsName := randomWithPrefix("se_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreFilterConfig(dsName, "secret-store-get-test*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.0.vault_token.0.name", store.Name),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.0.vault_token.0.server_address", store.ServerAddress),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "ids.0", store.ID),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMSecretStore_GetMultiple(t *testing.T) {
	t.Parallel()
	_, err := createVaultTokenStoresWithPrefix("secret-store-multiple", 2)
	if err != nil {
		t.Fatal("failed to create stores in setup: ", err)
	}
	_, err = createVaultTokenStoresWithPrefix("nomatch", 1)
	if err != nil {
		t.Fatal("failed to create store in setup: ", err)
	}

	dsName := randomWithPrefix("se_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreFilterConfig(dsName, "secret-store-multiple*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.0.vault_token.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "ids.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMSecretStore_GetNone(t *testing.T) {
	t.Parallel()

	_, err := createVaultTokenStoresWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create store in setup: ", err)
	}

	dsName := randomWithPrefix("rs_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreFilterConfig(dsName, "nothingWillEverMatch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.0.vault_token.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "ids.#", "0"),
				),
			},
		},
	})
}

func TestAccSDMSecretStore_GetTags(t *testing.T) {
	t.Parallel()

	name := randomWithPrefix("test")
	address := randomWithPrefix("test")

	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = client.SecretStores().Create(ctx, &sdm.VaultTokenStore{
		Name:          name,
		ServerAddress: address,
		Tags:          sdm.Tags{"foo": "bar", "key": "value"},
	})
	if err != nil {
		t.Fatalf("failed to create store: %v", err)
	}
	dsName := randomWithPrefix("se_test_query")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreFilterConfig(dsName, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.0.vault_token.0.tags.key", "value"),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.0.vault_token.0.tags.foo", "bar"),
				),
			},
		},
	})

}

func testAccSDMSecretStoreFilterConfig(resourceDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
		data "sdm_secret_store" "%s" {
			name = "%s"
		}`, resourceDataSourceName, nameFilter)
}
