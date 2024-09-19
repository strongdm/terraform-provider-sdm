package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMProxyClusterKey_DataSourceGet(t *testing.T) {
	initAcceptanceTest(t)

	proxyCluster, err := createProxyClusterWithPrefix("test")
	if err != nil {
		t.Fatal("failed to create proxy cluster:", err)
	}

	if _, err := createProxyClusterKey(proxyCluster.ID); err != nil {
		t.Fatal("failed to create proxy cluster key: ", err)
	}

	dataSourceName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMProxyClusterKeyGetFilterConfig(dataSourceName, proxyCluster.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "proxy_cluster_keys.0.proxy_cluster_id", proxyCluster.ID),
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "proxy_cluster_keys.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMProxyClusterKey_DataSourceGetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	proxyCluster1, err := createProxyClusterWithPrefix("test")
	if err != nil {
		t.Fatal("failed to create proxy cluster:", err)
	}
	proxyCluster2, err := createProxyClusterWithPrefix("test")
	if err != nil {
		t.Fatal("failed to create proxy cluster:", err)
	}

	if _, err := createProxyClusterKey(proxyCluster1.ID); err != nil {
		t.Fatal("failed to create proxy cluster key: ", err)
	}
	if _, err := createProxyClusterKey(proxyCluster1.ID); err != nil {
		t.Fatal("failed to create proxy cluster key: ", err)
	}
	if _, err := createProxyClusterKey(proxyCluster2.ID); err != nil {
		t.Fatal("failed to create proxy cluster key: ", err)
	}

	dataSourceName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMProxyClusterKeyGetFilterConfig(dataSourceName, proxyCluster1.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "ids.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "proxy_cluster_keys.0.proxy_cluster_id", proxyCluster1.ID),
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "proxy_cluster_keys.1.proxy_cluster_id", proxyCluster1.ID),
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "proxy_cluster_keys.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMProxyClusterKey_DataSourceGetNone(t *testing.T) {
	initAcceptanceTest(t)

	proxyCluster, err := createProxyClusterWithPrefix("test")
	if err != nil {
		t.Fatal("failed to create proxy cluster:", err)
	}
	_, err = createProxyClusterKey(proxyCluster.ID)
	if err != nil {
		t.Fatal("failed to create test proxyClusterKey: ", err)
	}

	dataSourceName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMProxyClusterKeyGetFilterConfig(dataSourceName, "n-00333000"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "proxy_cluster_keys.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_proxy_cluster_key."+dataSourceName, "ids.#", "0"),
				),
			},
		},
	})
}

func testAccSDMProxyClusterKeyGetFilterConfig(dataSourceName string, proxyClusterID string) string {
	return fmt.Sprintf(
		`
	data "sdm_proxy_cluster_key" "%s" {
		proxy_cluster_id = "%s"
	}`,
		dataSourceName,
		proxyClusterID,
	)
}
