package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMRemoteIdentity_Create(t *testing.T) {
	// setup data
	accounts, err := createUsersWithPrefix("remote-identity-user", 1)
	if err != nil {
		t.Fatalf("failed to create account in setup: %v", err)
	}
	account := accounts[0].(*sdm.User)
	group := getDefaultRemoteIdentityGroup(t)
	username := randomWithPrefix("test-username")

	rsName := randomWithPrefix("test-remote-identity-resource")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "sdm_remote_identity" "%s" {
					account_id = "%s"
					remote_identity_group_id = "%s"
					username = "%s"
				}`, rsName, account.ID, group.ID, username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_remote_identity."+rsName, "account_id", account.ID),
					resource.TestCheckResourceAttr("sdm_remote_identity."+rsName, "remote_identity_group_id", group.ID),
					resource.TestCheckResourceAttr("sdm_remote_identity."+rsName, "username", username),
				),
			},
			{
				ResourceName:      "sdm_remote_identity." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMRemoteIdentity_Update(t *testing.T) {
}

func TestAccSDMRemoteIdentity_Delete(t *testing.T) {
}
