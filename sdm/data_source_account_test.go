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
	initAcceptanceTest(t)
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
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.external_id", account.ExternalID),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.permission_level", sdm.PermissionLevelUser),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.managed_by", sdm.ProviderNone),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.0", account.GetID()),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMAccountDataSource_GetByTags(t *testing.T) {
	initAcceptanceTest(t)

	client, err := preTestClient()
	if err != nil {
		t.Fatal("failed to create test client", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tagValue := randomWithPrefix("value")
	name := randomWithPrefix("name")
	_, err = client.Accounts().Create(ctx, &sdm.User{
		FirstName: name,
		LastName:  name,
		Email:     name,
		Tags: sdm.Tags{
			"testTag": tagValue,
		},
	})
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}
	dsName := randomWithPrefix("test-account-ds")

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "sdm_account" "%s" {
						tags = {
							"testTag" = "%s"
						}
					}`, dsName, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.email", name),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.0.tags.testTag", tagValue),
				),
			},
		},
	})
}

func TestAccSDMAccount_GetSuspended(t *testing.T) {
	initAcceptanceTest(t)

	dsName := randomWithPrefix("ac_suspended_query")
	userName := randomWithPrefix("ac_suspended_query")
	_ = createUserWithSuspension(t, userName, true)

	userNameNotSuspended := randomWithPrefix("ac_suspended_query")
	notSuspended := createUserWithSuspension(t, userNameNotSuspended, false)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGetFilterSuspendedConfig(dsName, "ac_suspended_query*", false),
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

func createUserWithSuspension(t *testing.T, name string, suspended bool) *sdm.User {
	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}

	wantPermissionLevel := sdm.PermissionLevelUser
	if suspended {
		wantPermissionLevel = sdm.PermissionLevelSuspended
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	createResp, err := client.Accounts().Create(ctx, &sdm.User{
		FirstName:       name,
		LastName:        name,
		Email:           name,
		PermissionLevel: wantPermissionLevel,
	})
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}

	return createResp.Account.(*sdm.User)
}

func TestAccSDMAccount_GetMultiple(t *testing.T) {
	initAcceptanceTest(t)
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
	initAcceptanceTest(t)
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

func TestAccSDMAccount_GetMultipleTypes(t *testing.T) {
	initAcceptanceTest(t)
	_, err := createUsersWithPrefix("account-multi-type", 2)
	if err != nil {
		t.Fatal("failed to create users in setup: ", err)
	}
	_, err = createServicesWithPrefix("account-multi-type", 2)
	if err != nil {
		t.Fatal("failed to create services in setup: ", err)
	}
	_, err = createTokensWithPrefix("account-multi-type", "api", 2)
	if err != nil {
		t.Fatal("failed to create api tokens in setup: ", err)
	}
	_, err = createTokensWithPrefix("account-multi-type", "admin-token", 2)
	if err != nil {
		t.Fatal("failed to create admin-token tokens in setup: ", err)
	}

	dsName := randomWithPrefix("ac_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGetFilterNameConfig(dsName, "account-multi-type*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "ids.#", "8"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.user.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.service.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName, "accounts.0.token.#", "4"),
				),
			},
		},
	})

	dsName2 := randomWithPrefix("ac_test_query")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountGetFilterNameTypeConfig(dsName2, "account-multi-type*", "api"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account."+dsName2, "ids.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_account."+dsName2, "accounts.0.token.#", "2"),
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

func testAccSDMAccountGetFilterSuspendedConfig(dsName, firstNameFilter string, suspendedFilter bool) string {
	return fmt.Sprintf(`
	data "sdm_account" "%s" {
		first_name = "%s"
		suspended = %t
	}`, dsName, firstNameFilter, suspendedFilter)
}

func testAccSDMAccountGetFilterNameConfig(dsName, nameFilter string) string {
	return fmt.Sprintf(`
	data "sdm_account" "%s" {
		name = "%s"
	}`, dsName, nameFilter)
}

func testAccSDMAccountGetFilterNameTypeConfig(dsName, nameFilter, typeFilter string) string {
	return fmt.Sprintf(`
	data "sdm_account" "%s" {
		name = "%s"
		type = "%s"
	}`, dsName, nameFilter, typeFilter)
}
