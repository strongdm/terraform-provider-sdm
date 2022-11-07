package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMRemoteIdentityGroupDataSource_GetByName(t *testing.T) {
	initAcceptanceTest(t)

	dsName := randomWithPrefix("test-remote-identity-group-ds")
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

func getDefaultRemoteIdentityGroup(t *testing.T) *sdm.RemoteIdentityGroup {
	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := "name:default"
	listResp, err := client.RemoteIdentityGroups().List(ctx, filter)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if !listResp.Next() {
		t.Fatalf("list empty")
	}
	return listResp.Value()
}
