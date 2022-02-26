package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMResource_Get(t *testing.T) {
	t.Parallel()

	resources, err := createRedisesWithPrefix("resource-get-test", 1)
	if err != nil {
		t.Fatal("failed to create redis in setup: ", err)
	}
	redis := resources[0].(*sdm.Redis)

	dsName := randomWithPrefix("rs_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceFilterConfig(dsName, "resource-get-test*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.name", redis.Name),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.hostname", redis.Hostname),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.port", fmt.Sprintf("%d", redis.Port)),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.port_override", fmt.Sprintf("%d", redis.PortOverride)),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "ids.0", redis.ID),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMResource_SecretStoreGet(t *testing.T) {
	t.Parallel()

	vaults, err := createVaultTokenStoresWithPrefix("vaultTest", 1)
	if err != nil {
		t.Fatalf("failed to create secret store: %v", err)
	}
	vault := vaults[0]

	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}

	_, err = client.Resources().Create(context.Background(), &sdm.Redis{
		Name:          randomWithPrefix("redis-secret-store-get"),
		Hostname:      "hostname",
		Password:      "/secret/path?key=password",
		PortOverride:  portOverride.Count(),
		SecretStoreID: vault.GetID(),
	})
	if err != nil {
		t.Fatalf("failed to create resource: %v", err)
	}

	dsName := randomWithPrefix("rs_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceFilterConfig(dsName, "redis-secret-store-get*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.secret_store_id", vault.GetID()),
				),
			},
		},
	})
}

func TestAccSDMResource_SSHGet(t *testing.T) {
	t.Parallel()

	resources, err := createSSHesWithPrefix("resource-get-ssh-test", 1)
	if err != nil {
		t.Fatal("failed to create redis in setup: ", err)
	}
	ssh := resources[0].(*sdm.SSH)

	dsName := randomWithPrefix("rs_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceFilterConfig(dsName, "resource-get-ssh-test*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh.0.name", ssh.Name),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh.0.hostname", ssh.Hostname),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh.0.username", ssh.Username),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh.0.port", fmt.Sprintf("%d", ssh.Port)),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh.0.public_key", ssh.PublicKey),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "ids.0", ssh.ID),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMResource_GetMultiple(t *testing.T) {
	t.Parallel()
	_, err := createRedisesWithPrefix("resource-multiple", 2)
	if err != nil {
		t.Fatal("failed to create redises in setup: ", err)
	}
	_, err = createRedisesWithPrefix("nomatch", 1)
	if err != nil {
		t.Fatal("failed to create redis in setup: ", err)
	}

	dsName := randomWithPrefix("rs_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceFilterConfig(dsName, "resource-multiple-*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "ids.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMResource_GetNone(t *testing.T) {
	t.Parallel()

	_, err := createRedisesWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create redis in setup: ", err)
	}

	dsName := randomWithPrefix("rs_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceFilterConfig(dsName, "nothingWillEverMatch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "ids.#", "0"),
				),
			},
		},
	})
}

func TestAccSDMResource_GetTags(t *testing.T) {
	t.Parallel()

	name := randomWithPrefix("test")
	port := portOverride.Count()

	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = client.Resources().Create(ctx, &sdm.Redis{
		Name:         name,
		Hostname:     "example.com",
		Port:         port,
		PortOverride: port,
		Tags:         sdm.Tags{"foo": "bar", "key": "value"},
	})
	if err != nil {
		t.Fatalf("failed to create redis: %v", err)
	}
	dsName := randomWithPrefix("rs_test_query")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceFilterConfig(dsName, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.tags.key", "value"),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.tags.foo", "bar"),
				),
			},
		},
	})

}

func testAccSDMResourceFilterConfig(resourceDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
		data "sdm_resource" "%s" {
			name = "%s"
		}`, resourceDataSourceName, nameFilter)
}
