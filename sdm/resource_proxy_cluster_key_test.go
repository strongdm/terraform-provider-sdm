package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func init() {
	resource.AddTestSweepers("sdm_proxy_cluster_key", &resource.Sweeper{
		Name: "sdm_proxy_cluster_key",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			list, err := client.Nodes().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("error listing nodes: %s", err)
			}

			for list.Next() {
				node := list.Value()
				if _, ok := node.(*sdm.ProxyCluster); !ok {
					continue
				}
				_, err := client.Nodes().Delete(ctx, node.GetID())
				if err != nil {
					return fmt.Errorf("failed to clean up proxy cluster: %w", err)
				}
			}
			if list.Err() != nil {
				return fmt.Errorf("error listing nodes %s", list.Err())
			}
			return nil
		},
	})
}

func TestAccSDMProxyClusterKey_Create(t *testing.T) {
	initAcceptanceTest(t)
	proxyClusterName := randomWithPrefix("test-proxy-cluster")
	proxyClusterKeyName := randomWithPrefix("test-proxy-cluster-key")
	proxyClusterID := ""
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMProxyClusterKeyConfig(proxyClusterName, proxyClusterKeyName),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						var err error
						proxyClusterID, err = testCreatedID(s, "sdm_node", proxyClusterName)
						if err != nil {
							return err
						}
						return nil
					},
					resource.TestCheckResourceAttrSet("sdm_proxy_cluster_key."+proxyClusterKeyName, "secret_key"),
					func(s *terraform.State) error {
						if err := resource.TestCheckResourceAttr("sdm_proxy_cluster_key."+proxyClusterKeyName, "proxy_cluster_id", proxyClusterID)(s); err != nil {
							return err
						}
						return nil
					},
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_proxy_cluster_key", proxyClusterKeyName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.ProxyClusterKeys().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created role: %w", err)
						}

						if resp.ProxyClusterKey.ProxyClusterID != proxyClusterID {
							return fmt.Errorf("unexpected proxy cluster ID '%s', expected '%s'", resp.ProxyClusterKey.ProxyClusterID, proxyClusterID)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_proxy_cluster_key." + proxyClusterKeyName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because secret key cannot be recovered
			},
		},
	})
}

func testAccSDMProxyClusterKeyConfig(proxyClusterName string, proxyClusterKeyName string) string {
	return fmt.Sprintf(
		`
	resource "sdm_node" "%[1]s" {
		proxy_cluster {
			name = "%[1]s"
			address = "proxy:8443"
		}
	}

	resource "sdm_proxy_cluster_key" "%[2]s" {
		proxy_cluster_id = sdm_node.%[1]s.id
	}
	`,
		proxyClusterName,
		proxyClusterKeyName,
	)
}

func createProxyClusterKey(proxyClusterID string) (*sdm.ProxyClusterKey, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resp, err := client.ProxyClusterKeys().Create(ctx, &sdm.ProxyClusterKey{
		ProxyClusterID: proxyClusterID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ProxyClusterKey, nil
}
