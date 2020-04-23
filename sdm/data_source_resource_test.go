package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	sdm "github.com/strongdm/strongdm-sdk-go"
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

func testAccSDMResourceFilterConfig(resourceDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
		data "sdm_resource" "%s" {
			name = "%s"
		}`, resourceDataSourceName, nameFilter)
}
