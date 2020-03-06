package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	sdm "github.com/strongdm/strongdm-sdk-go"
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
