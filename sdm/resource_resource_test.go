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
	name := randomWithPrefix("test")
	port := portOverride.Count()
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceRedisConfig(name, name, port, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.name", name),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.port_override", fmt.Sprint(port)),
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
		},
	})
}

func TestAccSDMResource_CreateSSH(t *testing.T) {
	name := randomWithPrefix("test")
	port := portOverride.Count()
	resource.ParallelTest(t, resource.TestCase{
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
		},
	})
}

func TestAccSDMResource_Update(t *testing.T) {
	resourceName := randomWithPrefix("test")

	redisName := randomWithPrefix("test")
	updatedRedisName := randomWithPrefix("test2")

	port := portOverride.Count()
	updatedPort := portOverride.Count()

	pOverride := portOverride.Count()
	updatedPortOverride := portOverride.Count()

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceRedisConfig(resourceName, redisName, port, pOverride),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port", fmt.Sprint(port)),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port_override", fmt.Sprint(pOverride)),
				),
			},
			{
				Config: testAccSDMResourceRedisConfig(resourceName, updatedRedisName, updatedPort, updatedPortOverride),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", updatedRedisName),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port", fmt.Sprint(updatedPort)),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port_override", fmt.Sprint(updatedPortOverride)),
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

						if resp.Resource.(*sdm.Redis).PortOverride != updatedPortOverride {
							return fmt.Errorf("unexpected port override '%d', expected '%d'", resp.Resource.(*sdm.Redis).PortOverride, updatedPortOverride)
						}

						return nil
					},
				),
			},
		},
	})
}

func testAccSDMResourceRedisConfig(resourceName string, sdmResourceName string, port, pOverride int32) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			port = %d
			port_override = %d
		}
	}
	`, resourceName, sdmResourceName, port, pOverride)
}

func testAccSDMResourceSSHConfig(resourceName string, sdmResourceName string, port int32) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		ssh {
			name = "%s"
			hostname = "test.com"
			username = "user"
			port = %d
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
