package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSDMRoleAttachmentDataSource_Get(t *testing.T) {
	t.Parallel()

	roleAttachments, err := createRoleAttachmentsWithPrefix("role", 1)
	if err != nil {
		t.Fatal("failed to create role attachment: ", err)
	}
	roleAttachment := roleAttachments[0]
	rsName := randomWithPrefix("roleattachment")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleAttachmentDataSourceGetFilterConfig(rsName, roleAttachment.CompositeRoleID, roleAttachment.AttachedRoleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+rsName, "role_attachments.0.composite_role_id", roleAttachment.CompositeRoleID),
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+rsName, "role_attachments.0.attached_role_id", roleAttachment.AttachedRoleID),
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+rsName, "role_attachments.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMRoleAttachmentDataSource_GetMultiple(t *testing.T) {
	t.Parallel()

	// create three roles
	roles, err := createRolesWithPrefix("test", 3, false)
	if err != nil {
		t.Fatal("failed to create roles in setup: ", err)
	}

	compositeRoles, err := createRolesWithPrefix("test", 1, true)
	if err != nil {
		t.Fatal("failed to create role in setup: ", err)
	}
	compositeRole := compositeRoles[0]

	// attach two of the three roles to the composite role
	_, err = attachRoles(compositeRole, roles[1:])
	if err != nil {
		t.Fatal("failed to create attachments in setup: ", err)
	}

	dsName := randomWithPrefix("roleattachment")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleAttachmentDataSourceGetFilterConfig(dsName, compositeRole.ID, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+dsName, "role_attachments.0.composite_role_id", compositeRole.ID),
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+dsName, "role_attachments.1.composite_role_id", compositeRole.ID),
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+dsName, "role_attachments.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMRoleAttachmentDataSource_GetNone(t *testing.T) {
	t.Parallel()
	_, err := createRoleAttachmentsWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create role attachment: ", err)
	}
	dsName := randomWithPrefix("roleattachment")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMRoleAttachmentDataSourceGetFilterConfig(dsName, "", "r-00333000"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+dsName, "ids.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_role_attachment."+dsName, "role_attachments.#", "0"),
				),
			},
		},
	})
}

func testAccSDMRoleAttachmentDataSourceGetFilterConfig(roleAttachmentDataSourceName, compositeRoleID, roleID string) string {
	return fmt.Sprintf(`
	data "sdm_role_attachment" "%s" {
		composite_role_id = "%s"
		attached_role_id = "%s"
	}`,
		roleAttachmentDataSourceName,
		compositeRoleID,
		roleID,
	)
}
