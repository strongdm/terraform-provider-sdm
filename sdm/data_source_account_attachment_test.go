package sdm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSDMAccountAttachment_DataSourceGet(t *testing.T) {
	initAcceptanceTest(t)

	attachments, err := createAccountAttachmentsWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create test attachment: ", err)
	}
	roleID := attachments[0].RoleID
	accountID := attachments[0].AccountID

	accountAttachmentName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountAttachmentGetFilterConfig(accountAttachmentName, accountID, roleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.0.account_id", accountID),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.0.role_id", roleID),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.#", "1"),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "ids.#", "1"),
				),
			},
		},
	})
}

func TestAccSDMAccountAttachment_DataSourceGetMultiple(t *testing.T) {
	initAcceptanceTest(t)

	// create three accounts
	accounts, err := createUsersWithPrefix("test", 3)
	if err != nil {
		t.Fatal("failed to create accounts in setup: ", err)
	}

	roles, err := createRolesWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create role in setup: ", err)
	}
	role := roles[0]

	// attach two of the three accounts to the role
	_, err = attachAccounts(role, accounts[1:])
	if err != nil {
		t.Fatal("failed to create attachments in setup: ", err)
	}

	accountAttachmentName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountAttachmentGetFilterConfig(accountAttachmentName, "", role.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "ids.#", "2"),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.0.role_id", role.ID),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.1.role_id", role.ID),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.1.role_id", role.ID),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.#", "2"),
				),
			},
		},
	})
}

func TestAccSDMAccountAttachment_DataSourceGetNone(t *testing.T) {
	initAcceptanceTest(t)

	_, err := createAccountAttachmentsWithPrefix("test", 1)
	if err != nil {
		t.Fatal("failed to create test attachment: ", err)
	}

	accountAttachmentName := randomWithPrefix("test")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMAccountAttachmentGetFilterConfig(accountAttachmentName, "", "r-00333000"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "account_attachments.#", "0"),
					resource.TestCheckResourceAttr("data.sdm_account_attachment."+accountAttachmentName, "ids.#", "0"),
				),
			},
		},
	})
}

func testAccSDMAccountAttachmentGetFilterConfig(attachmentDataSourceName, accountID, roleID string) string {
	return fmt.Sprintf(
		`
	data "sdm_account_attachment" "%s" {
		account_id = "%s"
		role_id = "%s"
	}`,
		attachmentDataSourceName,
		accountID,
		roleID,
	)
}
