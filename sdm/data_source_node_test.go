package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMNode_Get(t *testing.T) {
	t.Parallel()
	nodes, err := createRelaysWithPrefix("node-get", 1)
	if err != nil {
		t.Fatal("failed to create node in setup: ", err)
	}
	relay := nodes[0].(*sdm.Relay)
	dsName := randomWithPrefix("node_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeGetFilterConfig(dsName, "node-get*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "nodes.0.relay.0.name", relay.Name),
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "ids.0", relay.ID),
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "ids.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "nodes.0.relay.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMNodeDataSource_GetByTags(t *testing.T) {
	t.Parallel()

	client, err := preTestClient()
	if err != nil {
		t.Fatal("failed to create test client", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tagValue := randomWithPrefix("value")
	name := randomWithPrefix("name")
	createResp, err := client.Nodes().Create(ctx, &sdm.Relay{
		Name: name,
		Tags: sdm.Tags{
			"testTag": tagValue,
		},
	})
	_ = createResp.Node
	if err != nil {
		t.Fatalf("failed to create node: %v", err)
	}
	dsName := randomWithPrefix("test-node-ds")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "sdm_node" "%s" {
						tags = {
							"testTag" = "%s"
						}
					}`, dsName, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "nodes.0.relay.0.name", name),
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "nodes.0.relay.0.tags.testTag", tagValue),
				),
			},
		},
	})
}

func TestAccSDMNode_GetMultiple(t *testing.T) {
	t.Parallel()

	_, err := createRelaysWithPrefix("node-multiple", 2)
	if err != nil {
		t.Fatal("failed to create nodes in setup: ", err)
	}

	_, err = createRelaysWithPrefix("nomatch", 1)
	if err != nil {
		t.Fatal("failed to create node in setup: ", err)
	}

	dsName := randomWithPrefix("node_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeGetFilterConfig(dsName, "node-multiple*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "ids.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "nodes.0.relay.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMNode_GetNone(t *testing.T) {
	t.Parallel()

	_, err := createRelaysWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create node in setup: ", err)
	}

	dsName := randomWithPrefix("node_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeGetFilterConfig(dsName, "nothingWillEverMatch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "ids.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_node."+dsName, "nodes.0.relay.#", "0"),
				),
			},
		},
	})
}

func testAccSDMNodeGetFilterConfig(nodeDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_node" "%s" {
		name = "%s"
	}`, nodeDataSourceName, nameFilter)
}
