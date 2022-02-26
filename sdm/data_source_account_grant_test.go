package sdm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMAccountGrant_DataSourceGet(t *testing.T) {
	if os.Getenv("SDM_ACCESS_OVERHAUL") == "yes" {
		t.Skip("skipping legacy test")
	}
	t.Parallel()

	grants, err := createAccountGrantsWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create account grant: ", err)
	}
	grant := grants[0]

	grantDataName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGrantGetFilterConfig(grantDataName, grant.AccountID, grant.ResourceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "account_grants.0.account_id", grant.AccountID),
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "account_grants.0.resource_id", grant.ResourceID),
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "account_grants.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMAccountGrant_DataSourceGetMultiple(t *testing.T) {
	if os.Getenv("SDM_ACCESS_OVERHAUL") == "yes" {
		t.Skip("skipping legacy test")
	}
	t.Parallel()

	// create three redises
	resources, err := createRedisesWithPrefix("test", 3)
	if err != nil {
		t.Fatal("failed to create accounts: ", err)
	}
	users, err := createUsersWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create role: ", err)
	}
	// only grant two of the created resources
	_, err = grantAccount(users[0], resources[1:])
	if err != nil {
		t.Fatal("failed to attach accounts: ", err)
	}
	accountID := users[0].GetID()

	grantDataName := randomWithPrefix("ag_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGrantGetFilterConfig(grantDataName, accountID, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "account_grants.0.account_id", accountID),
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "account_grants.1.account_id", accountID),
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "account_grants.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "ids.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMAccountGrant_DataSourceGetNone(t *testing.T) {
	if os.Getenv("SDM_ACCESS_OVERHAUL") == "yes" {
		t.Skip("skipping legacy test")
	}
	t.Parallel()

	_, err := createAccountGrantsWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create account grant: ", err)
	}

	grantDataName := randomWithPrefix("ag_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGrantGetFilterConfig(grantDataName, "", "rs-00333000"), // this must not be a valid ID
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "account_grants.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_account_grant."+grantDataName, "ids.#", "0"),
				),
			},
		},
	})
}

func testAccSDMAccountGrantGetFilterConfig(accountDataSourceName, accountIDFilter, resourceIDFilter string) string {
	return fmt.Sprintf(
		`
	data "sdm_account_grant" "%s" {
		account_id = "%s"
		resource_id = "%s"
	}`,
		accountDataSourceName,
		accountIDFilter,
		resourceIDFilter,
	)
}
