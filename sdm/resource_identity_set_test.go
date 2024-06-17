package sdm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("sdm_identity_set", &resource.Sweeper{
		Name:         "sdm_identity_set",
		Dependencies: []string{"sdm_resource"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.IdentitySets().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing identity sets: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.IdentitySets().Delete(ctx, v.ID)
				if err != nil {
					fmt.Printf("error deleting identity set %s %s\n", v.ID, err)
				}
			}
			return nil
		},
	})
}

func TestAccSDMIdentitySet_Create(t *testing.T) {
	initAcceptanceTest(t)
	rsName := randomWithPrefix("test-identity-set-resource")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "sdm_identity_set" "%s" {
					name = "%s"
				}`, rsName, rsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_identity_set."+rsName, "name", rsName),
				),
			},
			{
				ResourceName:      "sdm_identity_set." + rsName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMIdentitySet_Update(t *testing.T) {
	initAcceptanceTest(t)
	rsName := randomWithPrefix("test-identity-set-resource")
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "sdm_identity_set" "%s" {
					name = "%s"
				}`, rsName, rsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_identity_set."+rsName, "name", rsName),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "sdm_identity_set" "%s" {
					name = "%s"
				}`, rsName, rsName+"-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_identity_set."+rsName, "name", rsName+"-updated"),
				),
			},
		},
	})
}
