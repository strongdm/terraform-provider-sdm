package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSDMControlPanel_SSHCAPublicKeyGet(t *testing.T) {
	t.Parallel()

	dataName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMControlPanelSSHCAPublicKeyGetConfig(dataName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.sdm_ssh_ca_pubkey."+dataName, "public_key"),
				),
			},
		},
	})
}

func testAccSDMControlPanelSSHCAPublicKeyGetConfig(dataName string) string {
	return fmt.Sprintf(`
	data "sdm_ssh_ca_pubkey" "%s" {
	}`,
		dataName,
	)
}
