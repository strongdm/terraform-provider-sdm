package sdm

import (
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

// randomWithPrefix is used for sweeper clean up.
func randomWithPrefix(prefix string) string {
	return acctest.RandomWithPrefix(prefix) + "-test-acc"
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
