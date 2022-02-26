package sdm

import (
	"fmt"
	"testing"

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
