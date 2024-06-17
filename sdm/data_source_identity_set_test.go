package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMIdentitySetDataSource_GetByName(t *testing.T) {
	initAcceptanceTest(t)
	identitySet := createIdentitySet(t)
	dsName := randomWithPrefix("test-identity-set-ds")

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMIdentitySetDataSourceGetFilterConfig(dsName, identitySet.Name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_identity_set."+dsName, "identity_sets.0.name", identitySet.Name),
				),
			},
		},
	})
}

func testAccSDMIdentitySetDataSourceGetFilterConfig(dsName, filter string) string {
	return fmt.Sprintf(`
	data "sdm_identity_set" "%s" {
		name = "%s"
	}`, dsName, filter)
}
