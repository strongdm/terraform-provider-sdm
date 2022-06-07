package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMRemoteIdentityGroupDataSource_GetByName(t *testing.T) {
	t.Parallel()

	dsName := randomWithPrefix("test-role-ds")
	name := "default"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRemoteIdentityGroupDataSourceGetFilterConfig(dsName, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_remote_identity_group."+dsName, "remote_identity_groups.0.name", name),
				),
			},
		},
	})
}

func testAccSDMRemoteIdentityGroupDataSourceGetFilterConfig(dsName, filter string) string {
	return fmt.Sprintf(`
	data "sdm_remote_identity_group" "%s" {
		name = "%s"
	}`, dsName, filter)
}
