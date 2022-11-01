package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func init() {
	resource.AddTestSweepers("sdm_node", &resource.Sweeper{
		Name: "sdm_node",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.Nodes().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing nodes: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.Nodes().Delete(ctx, v.GetID())
				if err != nil {
					fmt.Printf("error deleting nodes %s %s\n", v.GetID(), err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing nodes %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMNode_GatewayCreate(t *testing.T) {
	rsName := randomWithPrefix("test")
	gwName := randomWithPrefix("test")
	ip, _ := acctest.RandIpAddress("192.0.2.0/24")
	listenAddr := fmt.Sprintf("%s:21222", ip)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeGatewayConfig(rsName, gwName, listenAddr, "0.0.0.0:21222"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.name", gwName),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.#", "1"),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.bind_address", "0.0.0.0:21222"),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.listen_address", listenAddr),
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "gateway.0.token"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						gateway, ok := resp.Node.(*sdm.Gateway)
						if !ok {
							return fmt.Errorf("expected gateway, got relay")
						}
						if gateway.Name != gwName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", gateway.Name, gwName)
						}
						if gateway.ListenAddress != listenAddr {
							return fmt.Errorf("unexpected listen address '%s', expected '%s'", gateway.ListenAddress, listenAddr)
						}
						if gateway.BindAddress != "0.0.0.0:21222" {
							return fmt.Errorf("unexpected bind address '%s', expected '%s'", gateway.BindAddress, "0.0.0.0:21222")
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
		},
	})
}

func TestAccSDMNode_GatewayCreateNoNameOrBindAddress(t *testing.T) {
	rsName := randomWithPrefix("test")
	gwName := randomWithPrefix("test")
	ip, _ := acctest.RandIpAddress("192.0.2.0/24")
	listenAddr := fmt.Sprintf("%s:21222", ip)
	gwBindAddr1 := "0.0.0.0:21222"
	gwBindAddr2 := "0.0.0.0:5000"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			// no name or bind address; strongDM should generate them
			{
				Config: testAccSDMNodeGatewayConfigNoNameOrBindAddress(rsName, listenAddr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.#", "1"),
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "gateway.0.name"),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.bind_address", gwBindAddr1),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.listen_address", listenAddr),
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "gateway.0.token"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						gateway, ok := resp.Node.(*sdm.Gateway)
						if !ok {
							return fmt.Errorf("expected gateway, got relay")
						}
						if gateway.Name == "" {
							return fmt.Errorf("expected gateway name to be non-empty")
						}
						if gateway.ListenAddress != listenAddr {
							return fmt.Errorf("unexpected listen address '%s', expected '%s'", gateway.ListenAddress, listenAddr)
						}
						if gateway.BindAddress != gwBindAddr1 {
							return fmt.Errorf("unexpected bind address '%s', expected '%s'", gateway.BindAddress, gwBindAddr1)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
			// now specify a name and bind address, it should update properly
			{
				Config: testAccSDMNodeGatewayConfig(rsName, gwName, listenAddr, "0.0.0.0:5000"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.#", "1"),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.name", gwName),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.bind_address", gwBindAddr2),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.listen_address", listenAddr),
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "gateway.0.token"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						gateway, ok := resp.Node.(*sdm.Gateway)
						if !ok {
							return fmt.Errorf("expected gateway, got relay")
						}
						if gateway.Name != gwName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", gateway.Name, gwName)
						}
						if gateway.ListenAddress != listenAddr {
							return fmt.Errorf("unexpected listen address '%s', expected '%s'", gateway.ListenAddress, listenAddr)
						}
						if gateway.BindAddress != gwBindAddr2 {
							return fmt.Errorf("unexpected bind address '%s', expected '%s'", gateway.BindAddress, gwBindAddr2)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
			// now elide the name and bind address again; it should stay the same
			{
				Config: testAccSDMNodeGatewayConfigNoNameOrBindAddress(rsName, listenAddr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.#", "1"),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.name", gwName),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.bind_address", gwBindAddr2),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "gateway.0.listen_address", listenAddr),
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "gateway.0.token"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						gateway, ok := resp.Node.(*sdm.Gateway)
						if !ok {
							return fmt.Errorf("expected gateway, got relay")
						}
						if gateway.Name != gwName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", gateway.Name, gwName)
						}
						if gateway.ListenAddress != listenAddr {
							return fmt.Errorf("unexpected listen address '%s', expected '%s'", gateway.ListenAddress, listenAddr)
						}
						if gateway.BindAddress != gwBindAddr2 {
							return fmt.Errorf("unexpected bind address '%s', expected '%s'", gateway.BindAddress, gwBindAddr2)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
		},
	})
}

func TestAccSDMNode_RelayCreate(t *testing.T) {
	rsName := randomWithPrefix("test")
	relayName := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeRelayConfig(rsName, relayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.0.name", relayName),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.#", "1"),
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "relay.0.token"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						relay, ok := resp.Node.(*sdm.Relay)
						if !ok {
							return fmt.Errorf("expected relay, got gateway")
						}
						if relay.Name != relayName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", relay.Name, relayName)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
		},
	})
}

func TestAccSDMNode_RelayCreateGeneratedName(t *testing.T) {
	rsName := randomWithPrefix("test")
	relayName := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			// don't provide a name; strongDM should generate one
			{
				Config: testAccSDMNodeRelayConfigNoName(rsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.#", "1"),
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "relay.0.name"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						relay, ok := resp.Node.(*sdm.Relay)
						if !ok {
							return fmt.Errorf("expected relay, got gateway")
						}
						if relay.Name == "" {
							return fmt.Errorf("expected generated relay name")
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
			// specify a name, it should be updated
			{
				Config: testAccSDMNodeRelayConfig(rsName, relayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.#", "1"),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.0.name", relayName),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						relay, ok := resp.Node.(*sdm.Relay)
						if !ok {
							return fmt.Errorf("expected relay, got gateway")
						}
						if relay.Name != relayName {
							return fmt.Errorf("expected relay name to be '%v', got '%v'", relayName, relay.Name)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
			// elide name field again; it should stay the same
			{
				Config: testAccSDMNodeRelayConfigNoName(rsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.#", "1"),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.0.name", relayName),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						relay, ok := resp.Node.(*sdm.Relay)
						if !ok {
							return fmt.Errorf("expected relay, got gateway")
						}
						if relay.Name != relayName {
							return fmt.Errorf("expected relay name to be '%v', got '%v'", relayName, relay.Name)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
		},
	})
}

func TestAccSDMNode_Update(t *testing.T) {
	rsName := randomWithPrefix("test")
	relayName := randomWithPrefix("test")
	updatedRelayName := randomWithPrefix("test2")

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeRelayConfig(rsName, relayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.0.name", relayName),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.#", "1"),
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
			{
				Config: testAccSDMNodeRelayConfig(rsName, updatedRelayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.0.name", updatedRelayName),
					resource.TestCheckResourceAttr("sdm_node."+rsName, "relay.#", "1"),
					func(s *terraform.State) error {
						// retrieve the resource by name from state
						id, err := testCreatedID(s, "sdm_node", rsName)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						relay, ok := resp.Node.(*sdm.Relay)
						if !ok {
							return fmt.Errorf("expected relay, got gateway")
						}
						if relay.Name != updatedRelayName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", relay.Name, updatedRelayName)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
		},
	})
}

func TestAccSDMNode_UpdateTokenStays(t *testing.T) {
	rsName := randomWithPrefix("test")
	relayName := randomWithPrefix("test")
	updatedRelayName := randomWithPrefix("test2")
	createdToken := ""
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeRelayConfig(rsName, relayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "relay.0.token"),
					func(s *terraform.State) error {

						rs, ok := s.RootModule().Resources["sdm_node."+rsName]
						if !ok {
							return fmt.Errorf("resource not found")
						}
						createdToken = rs.Primary.Attributes["relay.0.token"]
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
			{
				Config: testAccSDMNodeRelayConfig(rsName, updatedRelayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdm_node."+rsName, "relay.0.token"),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["sdm_node."+rsName]
						if !ok {
							return fmt.Errorf("resource not found")
						}
						if createdToken != rs.Primary.Attributes["relay.0.token"] {
							return fmt.Errorf("token changed during update")
						}
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + rsName,
				ImportState:       true,
				ImportStateVerify: false, // can't verify because token cannot be recovered
			},
		},
	})
}

func TestAccSDMNode_Tags(t *testing.T) {
	name := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMNodeTagsConfig(name, name, sdm.Tags{"key": "value", "foo": "bar"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_node."+name, "relay.0.tags.key", "value"),
					resource.TestCheckResourceAttr("sdm_node."+name, "relay.0.tags.foo", "bar"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_node", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created node: %w", err)
						}

						if resp.Node.GetTags()["key"] != "value" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Node.GetTags()["key"])
						}
						if resp.Node.GetTags()["foo"] != "bar" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Node.GetTags()["key"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + name,
				ImportState:       true,
				ImportStateVerify: false,
			},
			{
				Config: testAccSDMNodeTagsConfig(name, name, sdm.Tags{"goat": "bananas"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sdm_node."+name, "relay.0.tags.key"),
					resource.TestCheckNoResourceAttr("sdm_node."+name, "relay.0.tags.foo"),
					resource.TestCheckResourceAttr("sdm_node."+name, "relay.0.tags.goat", "bananas"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_node", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Nodes().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created node: %w", err)
						}

						if resp.Node.GetTags()["goat"] != "bananas" {
							return fmt.Errorf("unexpected value '%s' for tag 'goat', expected 'bananas'", resp.Node.GetTags()["goat"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_node." + name,
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func testAccSDMNodeTagsConfig(resourceName string, nodeName string, tags sdm.Tags) string {
	return fmt.Sprintf(`
	resource "sdm_node" "%s" {
		relay {
			name = "%s"
			tags = {
%s
			}
		}
	}
	`, resourceName, nodeName, tagsToConfigString(tags))
}

func testAccSDMNodeGatewayConfig(resourceName, gatewayName, listenAddr string, bindAddr string) string {
	return fmt.Sprintf(`
	resource "sdm_node" "%s" {
		gateway {
			name = "%s"
			listen_address = "%s"
			bind_address = "%s"
		}
	}
	`, resourceName, gatewayName, listenAddr, bindAddr)
}

func testAccSDMNodeGatewayConfigNoNameOrBindAddress(resourceName, listenAddr string) string {
	return fmt.Sprintf(`
	resource "sdm_node" "%s" {
		gateway {
			listen_address = "%s"
		}
	}
	`, resourceName, listenAddr)
}

func testAccSDMNodeRelayConfig(resourceName, relayName string) string {
	return fmt.Sprintf(`
	resource "sdm_node" "%s" {
		relay {
			name = "%s"
		}
	}
	`, resourceName, relayName)
}

func testAccSDMNodeRelayConfigNoName(resourceName string) string {
	return fmt.Sprintf(`
	resource "sdm_node" "%s" {
		relay {
		}
	}
	`, resourceName)
}

func createRelaysWithPrefix(prefix string, count int) ([]sdm.Node, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	nodes := []sdm.Node{}
	for i := 0; i < count; i++ {
		createResp, err := client.Nodes().Create(ctx, &sdm.Relay{
			Name: randomWithPrefix(prefix),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create relay: %w", err)
		}
		nodes = append(nodes, createResp.Node)
	}
	return nodes, nil
}
