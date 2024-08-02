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

const defaultPolicy = `forbid( principal, action, resource );`

func init() {
	resource.AddTestSweepers("sdm_policy", &resource.Sweeper{
		Name: "sdm_policy",
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.Policies().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing policies: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				fmt.Printf("%+v     %T", v, v)
				// _, err := client.Policies().Delete(ctx, v.Id)
				// if err != nil {
				// 	fmt.Printf("error deleting policy %s %s\n", v.Id, err)
				// }
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing secret store %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMPolicy_Create(t *testing.T) {
	initAcceptanceTest(t)
	name := randomWithPrefix("test")
	description := randomWithPrefix("test")
	rawPolicy := defaultPolicy
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMPolicyConfig(name, name, description, rawPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_policy."+name, "name", name),
					resource.TestCheckResourceAttr("sdm_policy."+name, "description", description),
					resource.TestCheckResourceAttr("sdm_policy."+name, "policy", rawPolicy),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_policy", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Policies().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Policy.Name != name {
							return fmt.Errorf("unexpected name %q, expected %q", resp.Policy.Name, name)
						}
						if resp.Policy.Description != description {
							return fmt.Errorf("unexpected description %q, expected %q", resp.Policy.Description, description)
						}
						if resp.Policy.Policy != rawPolicy {
							return fmt.Errorf("unexpected policy %q, expected %q", resp.Policy.Policy, rawPolicy)
						}
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_policy." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMPolicy_Update(t *testing.T) {
	initAcceptanceTest(t)
	resourceName := randomWithPrefix("test")
	updatedName := randomWithPrefix("test-updated")

	description := randomWithPrefix("desc")
	updatedDescription := randomWithPrefix("desc-updated")

	rawPolicy := ""
	updatedPolicy := defaultPolicy

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMPolicyConfig(resourceName, resourceName, description, rawPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_policy."+resourceName, "name", resourceName),
					resource.TestCheckResourceAttr("sdm_policy."+resourceName, "description", description),
					resource.TestCheckResourceAttr("sdm_policy."+resourceName, "policy", rawPolicy),
				),
			},
			{
				ResourceName:      "sdm_policy." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMPolicyConfig(resourceName, updatedName, updatedDescription, updatedPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_policy."+resourceName, "name", updatedName),
					resource.TestCheckResourceAttr("sdm_policy."+resourceName, "description", updatedDescription),
					resource.TestCheckResourceAttr("sdm_policy."+resourceName, "policy", updatedPolicy),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_policy", resourceName)
						if err != nil {
							return err
						}

						// check if it was actually updated
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Policies().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Policy.Name != updatedName {
							return fmt.Errorf("unexpected name %q, expected %q", resp.Policy.Name, updatedName)
						}
						if resp.Policy.Description != updatedDescription {
							return fmt.Errorf("unexpected description %q, expected %q", resp.Policy.Description, updatedDescription)
						}
						if resp.Policy.Policy != updatedPolicy {
							return fmt.Errorf("unexpected policy %q, expected %q", resp.Policy.Policy, updatedPolicy)
						}
						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_policy." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSDMPolicyConfig(resourceName, sdmResourceName, description, policy string) string {
	return fmt.Sprintf(`
	resource "sdm_policy" "%s" {
    name = "%s"
    description = "%s"
    policy = %q
	}
	`, resourceName, sdmResourceName, description, policy)
}

func createPoliciesWithPrefix(prefix string, count int) ([]*sdm.Policy, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	policies := []*sdm.Policy{}
	for i := 0; i < count; i++ {
		createResp, err := client.Policies().Create(ctx, &sdm.Policy{
			Name: randomWithPrefix(prefix),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create policy: %w", err)
		}
		policies = append(policies, createResp.Policy)
	}
	return policies, nil
}
