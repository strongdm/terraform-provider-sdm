package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMIdentityAliasDataSource_Get(t *testing.T) {
	initAcceptanceTest(t)

	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	accounts, err := createUsersWithPrefix("identity-alias-get", 1)
	if err != nil {
		t.Fatalf("failed to create account in setup: %v", err)
	}
	account := accounts[0].(*sdm.User)
	identitySet := createIdentitySet(t)
	username := randomWithPrefix("bobby")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// test setup
	created, err := client.IdentityAliases().Create(ctx, &sdm.IdentityAlias{
		AccountID:     account.ID,
		IdentitySetID: identitySet.ID,
		Username:      username,
	})
	if err != nil {
		t.Fatalf("failed to create identity alias in setup: %v", err)
	}
	dsName := randomWithPrefix("test-identity-ds")
	config := fmt.Sprintf(`
	data "sdm_identity_alias" "%s" {
		id = "%s"
		account_id = "%s"
		identity_set_id = "%s"
	}`, dsName, created.IdentityAlias.ID, account.ID, identitySet.ID)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					// TOOO fix s/remoteidentitys/remoteidentities/ (pkg/api/v1/templates/terraform-provider-sdm/sdm/data_source.go.tpl)
					resource.TestCheckResourceAttr("data.sdm_identity_alias."+dsName, "identity_aliases.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_identity_alias."+dsName, "identity_aliases.0.id", created.IdentityAlias.ID),
					resource.TestCheckResourceAttr("data.sdm_identity_alias."+dsName, "identity_aliases.0.account_id", created.IdentityAlias.AccountID),
					resource.TestCheckResourceAttr("data.sdm_identity_alias."+dsName, "identity_aliases.0.identity_set_id", created.IdentityAlias.IdentitySetID),
					resource.TestCheckResourceAttr("data.sdm_identity_alias."+dsName, "identity_aliases.0.username", created.IdentityAlias.Username),
				),
			},
		},
	})
}
