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
	initAcceptanceTest(t)
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
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.bind_interface", "127.0.0.1"),
				),
			},
		},
	})
}

// lintignore:AT012
func TestAccSDMResourceDataSource_GetByTags(t *testing.T) {
	initAcceptanceTest(t)

	client, err := preTestClient()
	if err != nil {
		t.Fatal("failed to create test client", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tagValue := randomWithPrefix("value")
	name := randomWithPrefix("name")
	_, err = client.Resources().Create(ctx, &sdm.Redis{
		Name:         name,
		Hostname:     "hostname",
		PortOverride: portOverride.Count(),
		Tags: sdm.Tags{
			"testTag":  tagValue,
			"testTag2": tagValue,
		},
	})
	if err != nil {
		t.Fatalf("failed to create resource: %v", err)
	}
	dsName := randomWithPrefix("test-resource-ds")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "sdm_resource" "%[1]s" {
						tags = {
							"testTag" = "%[2]s"
							"testTag2" = "%[2]s"
						}
					}`, dsName, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.name", name),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.tags.testTag", tagValue),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.redis.0.tags.testTag2", tagValue),
				),
			},
		},
	})
}

func TestAccSDMResource_SecretStoreGet(t *testing.T) {
	initAcceptanceTest(t)

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
	initAcceptanceTest(t)

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
	initAcceptanceTest(t)
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
	initAcceptanceTest(t)

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
	initAcceptanceTest(t)

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

// lintignore:AT012
func TestAccSDMResourceWithIdentitySet_Get(t *testing.T) {
	initAcceptanceTest(t)

	identitySet := createIdentitySet(t)
	resources, err := createSSHCertsWithPrefix("resource-identity-get-test", *identitySet, 1)
	if err != nil {
		t.Fatal("failed to create redis in setup: ", err)
	}
	sshCert := resources[0].(*sdm.SSHCert)

	dsName := randomWithPrefix("rs_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceFilterConfig(dsName, "resource-identity-get-test*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh_cert.0.name", sshCert.Name),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh_cert.0.hostname", sshCert.Hostname),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh_cert.0.port", fmt.Sprintf("%d", sshCert.Port)),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "ids.0", sshCert.ID),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh_cert.0.identity_set_id", identitySet.ID),
					resource.TestCheckResourceAttr("data.sdm_resource."+dsName, "resources.0.ssh_cert.0.identity_alias_healthcheck_username", "healthcheck"),
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

func createSSHCertsWithPrefix(prefix string, identitySet sdm.IdentitySet, count int) ([]sdm.Resource, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resources := []sdm.Resource{}
	for i := 0; i < count; i++ {
		port := portOverride.Count()
		createResp, err := client.Resources().Create(ctx, &sdm.SSHCert{
			Name:                             randomWithPrefix(prefix),
			Hostname:                         randomWithPrefix(prefix),
			Username:                         randomWithPrefix(prefix),
			Port:                             port,
			IdentitySetID:                    identitySet.ID,
			IdentityAliasHealthcheckUsername: "healthcheck",
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create ssh: %w", err)
		}
		resources = append(resources, createResp.Resource)
	}
	return resources, nil
}
