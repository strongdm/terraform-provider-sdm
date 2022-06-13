package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMRemoteIdentityDataSource_Get(t *testing.T) {
	t.Parallel()

	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	accounts, err := createUsersWithPrefix("account-get-remote-ident", 1)
	if err != nil {
		t.Fatalf("failed to create account in setup: %v", err)
	}
	account := accounts[0].(*sdm.User)
	group := getDefaultRemoteIdentityGroup(t)
	username := randomWithPrefix("bobby")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// test setup
	created, err := client.RemoteIdentities().Create(ctx, &sdm.RemoteIdentity{
		AccountID:             account.ID,
		RemoteIdentityGroupID: group.ID,
		Username:              username,
	})
	if err != nil {
		t.Fatalf("failed to create remote identity in setup: %v", err)
	}
	dsName := randomWithPrefix("test-remote-identity-ds")
	config := fmt.Sprintf(`
	data "sdm_remote_identity" "%s" {
		id = "%s"
		account_id = "%s"
		remote_identity_group_id = "%s"
	}`, dsName, created.RemoteIdentity.ID, account.ID, group.ID)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					// TOOO fix s/remoteidentitys/remoteidentities/ (pkg/api/v1/templates/terraform-provider-sdm/sdm/data_source.go.tpl)
					resource.TestCheckResourceAttr("data.sdm_remote_identity."+dsName, "remote_identities.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_remote_identity."+dsName, "remote_identities.0.id", created.RemoteIdentity.ID),
					resource.TestCheckResourceAttr("data.sdm_remote_identity."+dsName, "remote_identities.0.account_id", created.RemoteIdentity.AccountID),
					resource.TestCheckResourceAttr("data.sdm_remote_identity."+dsName, "remote_identities.0.remote_identity_group_id", created.RemoteIdentity.RemoteIdentityGroupID),
					resource.TestCheckResourceAttr("data.sdm_remote_identity."+dsName, "remote_identities.0.username", created.RemoteIdentity.Username),
				),
			},
		},
	})
}
