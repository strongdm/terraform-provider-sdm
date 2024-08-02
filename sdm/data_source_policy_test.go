package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMPolicy_Get(t *testing.T) {
	initAcceptanceTest(t)
	policies, err := createPoliciesWithPrefix("policy-get", 1)
	if err != nil {
		t.Fatal("failed to create policy in setup: ", err)
	}
	policy := policies[0]
	dsName := randomWithPrefix("policy_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMPoliciesGetFilterConfig(dsName, "policy-get*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "policies.0.name", policy.Name),
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "policies.0.id", policy.ID),
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "ids.0", policy.ID),
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMPolicy_GetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createPoliciesWithPrefix("policy-multiple", 2)
	if err != nil {
		t.Fatal("failed to create policies in setup: ", err)
	}

	_, err = createPoliciesWithPrefix("nomatch", 1)
	if err != nil {
		t.Fatal("failed to create policy in setup: ", err)
	}

	dsName := randomWithPrefix("policies_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMPoliciesGetFilterConfig(dsName, "policy-multiple*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "ids.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "policies.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMPolicy_GetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createPoliciesWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create policy in setup: ", err)
	}

	dsName := randomWithPrefix("policy_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMPoliciesGetFilterConfig(dsName, "nothingWillEverMatch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "ids.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_policy."+dsName, "policies.#", "0"),
				),
			},
		},
	})
}

func testAccSDMPoliciesGetFilterConfig(policyDataSourceName, nameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_policy" "%s" {
		name = "%s"
	}`, policyDataSourceName, nameFilter)
}
