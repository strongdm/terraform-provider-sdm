package sdm

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

// randomWithPrefix is used for sweeper clean up.
func randomWithPrefix(prefix string) string {
	return acctest.RandomWithPrefix(prefix) + "-test-acc"
}

// initAcceptanceTest skips the test unless the TF_ACC environment variable is
// set. the built-in resource.Test function does this as well, but we sometimes
// do some setup steps before calling that. This function ensures everything
// gets skipped properly.
func initAcceptanceTest(t testing.TB) {
	t.Helper()
	// replicate Terraform behavior: "We only run acceptance tests if an env var
	// is set because they're slow and generally require some outside
	// configuration."
	if os.Getenv(resource.TestEnvVar) == "" {
		t.Skipf("Acceptance tests skipped unless env '%s' set", resource.TestEnvVar)
	}
}

type AtomicCounter struct {
	mu    sync.Mutex
	count int32
}

func NewAtomicCounter(start int32) *AtomicCounter {
	return &AtomicCounter{
		count: start,
	}
}

func (c *AtomicCounter) Count() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	count := c.count
	c.count++
	return count
}

func createIdentitySet(t testing.TB) *sdm.IdentitySet {
	t.Helper()
	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	createResp, err := client.IdentitySets().Create(context.Background(), &sdm.IdentitySet{
		Name: randomWithPrefix("test-identity-set"),
	})
	if err != nil {
		t.Fatalf("failed to create test identity set: %v", err)
	}
	return createResp.IdentitySet
}
