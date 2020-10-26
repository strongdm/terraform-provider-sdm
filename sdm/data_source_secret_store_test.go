package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSDMSecretStoreDataSource_Get(t *testing.T) {
	t.Parallel()

	secretStores, err := createSecretStoresWithPrefix("test-secret-store", 1)
	if err != nil {
		t.Fatal("failed to create secret store: ", err)
	}
	secretStore := secretStores[0]

	rsName := randomWithPrefix("test-secret-store-ds")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreDataSourceGetFilterConfig(rsName, secretStore.Name, secretStore.Kind, secretStore.ServerAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+rsName, "secret_stores.0.name", secretStore.Name),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+rsName, "secret_stores.0.kind", secretStore.Kind),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+rsName, "secret_stores.0.server_address", secretStore.ServerAddress),
				),
			},
		},
	})
}

func TestAccSDMSecretStoreDataSource_GetMultiple(t *testing.T) {
	t.Parallel()

	_, err := createSecretStoresWithPrefix("multi-test", 2)
	if err != nil {
		t.Fatal("failed to create secret stores: ", err)
	}
	_, err = createSecretStoresWithPrefix("nomatch", 1)
	if err != nil {
		t.Fatal("failed to create secret store: ", err)
	}

	rsName := randomWithPrefix("test-secret-store")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreDataSourceGetFilterConfig(rsName, "multi-test*", "*", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+rsName, "secret_stores.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+rsName, "secret_stores.0.kind", "vault"),
					resource.TestCheckResourceAttr("data.sdm_secret_store."+rsName, "secret_stores.1.kind", "vault"),
				),
			},
		},
	})
}

func TestAccSDMSecretStoreDataSource_GetNone(t *testing.T) {
	t.Parallel()

	_, err := createSecretStoresWithPrefix("dontfind", 1)
	if err != nil {
		t.Fatal("failed to create secret store: ", err)
	}

	dsName := randomWithPrefix("secretStore-ds")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreDataSourceGetFilterConfig(dsName, "neverWillMatch", "*", "*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.#", "0"),
				),
			},
		},
	})
}

func TestAccSDMSecretStoreDataSource_GetByID(t *testing.T) {
	secretStoreName := randomWithPrefix("secret-store")
	dsName := randomWithPrefix("secret-store")
	kind := "vault"
	serverAddress := randomWithPrefix("secret-store")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMSecretStoreDataSourceGetByIDConfig(secretStoreName, kind, serverAddress, dsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_secret_store."+dsName, "secret_stores.#", "1"),
				),
			},
		},
	})
}

func testAccSDMSecretStoreDataSourceGetFilterConfig(secretStoreDataSourceName, nameFilter, kindFilter, serverAddressFilter string) string {
	return fmt.Sprintf(`
	data "sdm_secret_store" "%s" {
		name = "%s"
		kind = "%s"
		server_address = "%s"
	}`, secretStoreDataSourceName, nameFilter, kindFilter, serverAddressFilter)
}

func testAccSDMSecretStoreDataSourceGetByIDConfig(secretStoreName, kind, serverAddress string, secretStoreDataSourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_secret_store" "test_secretStore" {
		name = "%s"
		kind = "%s"
		server_address = "%s"
	}

	data "sdm_secret_store" "%s" {
		id = sdm_secret_store.test_secretStore.id
	}
	`, secretStoreName, kind, serverAddress, secretStoreDataSourceName)
}
