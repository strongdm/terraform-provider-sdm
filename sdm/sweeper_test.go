package sdm

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func TestMain(m *testing.M) {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"sdm": testAccProvider,
	}
	resource.TestMain(m)
}

func testCreatedResource(s *terraform.State, resourceType, resourceName string) (*terraform.ResourceState, error) {
	fullResourceName := fmt.Sprintf("%s.%s", resourceType, resourceName)
	rs, ok := s.RootModule().Resources[fullResourceName]
	if !ok {
		return nil, fmt.Errorf("not found: %s", fullResourceName)
	}
	return rs, nil
}

func testCreatedID(s *terraform.State, resourceType, resourceName string) (string, error) {
	rs, err := testCreatedResource(s, resourceType, resourceName)
	if err != nil {
		return "", err
	}

	if rs.Primary.ID == "" {
		return "", fmt.Errorf("resource ID not set")
	}

	return rs.Primary.ID, nil
}

func testClient() *sdm.Client {
	return testAccProvider.Meta().(*sdm.Client)
}

var preTestClientOnce sync.Once
var defaultPreTestClient *sdm.Client

func preTestClient() (*sdm.Client, error) {
	var err error
	preTestClientOnce.Do(func() {
		host := os.Getenv("SDM_API_HOST")
		id := os.Getenv("SDM_API_ACCESS_KEY")
		secretID := os.Getenv("SDM_API_SECRET_KEY")
		opts := []sdm.ClientOption{sdm.WithHost(host)}
		if strings.HasPrefix(host, "localhost:") ||
			strings.HasPrefix(host, "127.0.0.1:") ||
			strings.HasPrefix(host, "host.docker.internal:") {
			opts = append(opts, sdm.WithInsecure())
		}
		defaultPreTestClient, err = sdm.New(id, secretID, opts...)
	})
	return defaultPreTestClient, err
}

func testCheckDestroy(s *terraform.State) error {
	client := testClient()

	for _, rs := range s.RootModule().Resources {
		var err error
		switch rs.Type {
		case "sdm_account":
			_, err = client.Accounts().Get(context.Background(), rs.Primary.ID)
		case "sdm_account_attachment":
			_, err = client.AccountAttachments().Get(context.Background(), rs.Primary.ID)
		case "sdm_account_grant":
			_, err = client.AccountGrants().Get(context.Background(), rs.Primary.ID)
		case "sdm_node":
			_, err = client.Nodes().Get(context.Background(), rs.Primary.ID)
		case "sdm_resource":
			_, err = client.Resources().Get(context.Background(), rs.Primary.ID)
		case "sdm_role":
			_, err = client.Roles().Get(context.Background(), rs.Primary.ID)
		case "sdm_role_attachment":
			_, err = client.RoleAttachments().Get(context.Background(), rs.Primary.ID)
		case "sdm_role_grant":
			_, err = client.RoleGrants().Get(context.Background(), rs.Primary.ID)
		case "sdm_secret_store":
			_, err = client.SecretStores().Get(context.Background(), rs.Primary.ID)
		default:
			return fmt.Errorf("undefined resource type in testCheckDestroy: %s", rs.Type)
		}

		if err == nil {
			return fmt.Errorf("expected resource '%s' to be deleted", rs.Primary.ID)
		}
		var errNotFound *sdm.NotFoundError
		if !errors.As(err, &errNotFound) {
			return fmt.Errorf("unexpected error type: %v", err)
		}
	}

	return nil
}
