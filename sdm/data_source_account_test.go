package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMAccount_Get(t *testing.T) {
	t.Parallel()

	accounts, err := createUsersWithPrefix("account-get", 1)
	if err != nil {
		t.Fatal("failed to create account in setup: ", err)
	}
	account := accounts[0].(*sdm.User)

	dsName := randomWithPrefix("ac_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGetFilterConfig(dsName, "account-get*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.first_name", account.FirstName),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.last_name", account.LastName),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.email", account.Email),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.0", account.GetID()),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMAccount_GetSuspended(t *testing.T) {
	t.Parallel()

	dsName := randomWithPrefix("ac_suspended_query")
	dsNameSuspended := randomWithPrefix("ac_suspended_query")
	_ = createUserWithSuspension(t, dsNameSuspended, true)

	dsNameNotSuspended := randomWithPrefix("ac_suspended_query")
	notSuspended := createUserWithSuspension(t, dsNameNotSuspended, false)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGetFilterSuspenedConfig(dsName, "ac_suspend_query*", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.#", "1"), // only one user returned
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.#", "1"),

					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.first_name", notSuspended.FirstName),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.last_name", notSuspended.LastName),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.email", notSuspended.Email),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.0", notSuspended.GetID()),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.suspended", "false"),
				),
			},
		},
	})
}

func createUserWithSuspension(t *testing.T, dsName string, suspended bool) *sdm.User {
	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	createResp, err := client.Accounts().Create(ctx, &sdm.User{
		Suspended: suspended,
		FirstName: dsName,
		LastName:  dsName,
		Email:     dsName,
	})
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}

	return createResp.Account.(*sdm.User)
}

func TestAccSDMAccount_GetMultiple(t *testing.T) {
	t.Parallel()
	_, err := createUsersWithPrefix("account-multiple", 2)
	if err != nil {
		t.Fatal("failed to create accounts in setup: ", err)
	}
	_, err = createUsersWithPrefix("noMatch", 1)
	if err != nil {
		t.Fatal("failed to create account in setup: ", err)
	}

	dsName := randomWithPrefix("ac_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGetFilterConfig(dsName, "account-multiple*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMAccount_GetNone(t *testing.T) {
	t.Parallel()
	_, err := createUsersWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create account in setup: ", err)
	}
	dsName := randomWithPrefix("ac_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGetFilterConfig(dsName, "nothingWillEverMatch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.#", "0"),
				),
			},
		},
	})
}

func testAccSDMAccountGetFilterConfig(accountDataSourceName, firstNameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_account" "%s" {
		first_name = "%s"
	}`, accountDataSourceName, firstNameFilter)
}

func testAccSDMAccountGetFilterSuspenedConfig(dsName, firstNameFilter string, suspendedFilter bool) string {
	return fmt.Sprintf(`
	data "sdm_account" "%s" {
		suspended = %v
		first_name = "%s"
	}`, dsName, firstNameFilter, suspendedFilter)
}
