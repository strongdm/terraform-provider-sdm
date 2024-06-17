package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMIdentityAlias_Create(t *testing.T) {
	initAcceptanceTest(t)
	// setup data
	accounts, err := createUsersWithPrefix("identity-alias-user", 1)
	if err != nil {
		t.Fatalf("failed to create account in setup: %v", err)
	}
	account := accounts[0].(*sdm.User)
	identitySet := createIdentitySet(t)
	username := randomWithPrefix("test-username")

	rsName := randomWithPrefix("test-identity-alias-resource")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "sdm_identity_alias" "%s" {
					account_id = "%s"
					identity_set_id = "%s"
					username = "%s"
				}`, rsName, account.ID, identitySet.ID, username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "account_id", account.ID),
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "identity_set_id", identitySet.ID),
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "username", username),
				),
			},
			{
				ResourceName:      "sdm_identity_alias." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMIdentityAlias_Update(t *testing.T) {
	initAcceptanceTest(t)
	// setup data
	accounts, err := createUsersWithPrefix("identity-alias-user", 1)
	if err != nil {
		t.Fatalf("failed to create account in setup: %v", err)
	}
	account := accounts[0].(*sdm.User)
	identitySet := createIdentitySet(t)
	username := randomWithPrefix("test-username")

	rsName := randomWithPrefix("test-identity-alias-resource")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "sdm_identity_alias" "%s" {
					account_id = "%s"
					identity_set_id = "%s"
					username = "%s"
				}`, rsName, account.ID, identitySet.ID, username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "account_id", account.ID),
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "identity_set_id", identitySet.ID),
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "username", username),
				),
			},
			{
				ResourceName:      "sdm_identity_alias." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
				resource "sdm_identity_alias" "%s" {
					account_id = "%s"
					identity_set_id = "%s"
					username = "%s-update"
				}`, rsName, account.ID, identitySet.ID, username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "account_id", account.ID),
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "identity_set_id", identitySet.ID),
					resource.TestCheckResourceAttr("sdm_identity_alias."+rsName, "username", username+"-update"),
				),
			},
			{
				ResourceName:      "sdm_identity_alias." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
