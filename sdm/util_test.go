package sdm

import (
	"sync"
	"testing"

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

func TestJSONAccessRulesDiffSuppress(t *testing.T) {
	testCases := []struct {
		Name     string
		Old      string
		New      string
		Expected bool
	}{
		{
			Name:     "EmptyStringNoop",
			Old:      `[{"type":"redis"}]`,
			New:      ``,
			Expected: true,
		},
		{
			Name:     "ResourceTypeExactMatch",
			Old:      `[{"type":"redis"}]`,
			New:      `[{"type":"redis"}]`,
			Expected: true,
		},
		{
			Name:     "ResourceTypeFormattingSuppressed",
			Old:      `[{"type":"redis"}]`,
			New:      `[ { "type": "redis" } ]`,
			Expected: true,
		},
		{
			Name:     "ResourceTypeChanged",
			Old:      `[{"type":"redis"}]`,
			New:      `[{"type":"mysql"}]`,
			Expected: false,
		},
		{
			Name:     "ResourceIDsOrderingSuppressed",
			Old:      `[{"ids":["rs-0000000000000001","rs-0000000000000002"]}]`,
			New:      `[{"ids":["rs-0000000000000002","rs-0000000000000001"]}]`,
			Expected: true,
		},
		{
			Name:     "ResourceIDFormattingSuppressed",
			Old:      `[{"ids":["rs-0000000000000001","rs-0000000000000002"]}]`,
			New:      `[{"ids":["rs-1","rs-2"]}]`,
			Expected: true,
		},
		{
			Name:     "ResourceIDsExactMatch",
			Old:      `[{"ids":["rs-0000000000000001","rs-0000000000000002"]}]`,
			New:      `[{"ids":["rs-0000000000000001","rs-0000000000000002"]}]`,
			Expected: true,
		},
		{
			Name:     "ResourceIDAdded",
			Old:      `[{"ids":["rs-0000000000000001"]}]`,
			New:      `[{"ids":["rs-0000000000000001","rs-0000000000000002"]}]`,
			Expected: false,
		},
		{
			Name:     "ResourceIDRemoved",
			Old:      `[{"ids":["rs-0000000000000001","rs-0000000000000002"]}]`,
			New:      `[{"ids":["rs-0000000000000001"]}]`,
			Expected: false,
		},
		{
			Name:     "ResourceIDChanged",
			Old:      `[{"ids":["rs-0000000000000001"]}]`,
			New:      `[{"ids":["rs-0000000000000002"]}]`,
			Expected: false,
		},
		{
			Name:     "InvalidResourceID",
			Old:      `[{"ids":["rs-0000000000000001"]}]`,
			New:      `[{"ids":["ag-0000000000000001"]}]`,
			Expected: false,
		},
		{
			Name:     "TagsExactMatch",
			Old:      `[{"tags":{"a":"b"}}]`,
			New:      `[{"tags":{"a":"b"}}]`,
			Expected: true,
		},
		{
			Name:     "TagAdded",
			Old:      `[{"tags":{"a":"b"}}]`,
			New:      `[{"tags":{"a":"b","c":"d"}}]`,
			Expected: false,
		},
		{
			Name:     "TagRemoved",
			Old:      `[{"tags":{"a":"b","c":"d"}}]`,
			New:      `[{"tags":{"a":"b"}}]`,
			Expected: false,
		},
		{
			Name:     "TagChanged",
			Old:      `[{"tags":{"a":"b"}}]`,
			New:      `[{"tags":{"a":"z"}}]`,
			Expected: false,
		},
		{
			Name:     "TagsOrderingSuppressed",
			Old:      `[{"tags":{"a":"b","c":"d"}}]`,
			New:      `[{"tags":{"c":"d","a":"b"}}]`,
			Expected: true,
		},
		{
			Name:     "ResourceTypeAndTagsOrderingSuppressed",
			Old:      `[{"tags":{"a":"b"},"type":"redis"}]`,
			New:      `[{"type":"redis","tags":{"a":"b"}}]`,
			Expected: true,
		},
		{
			Name:     "ResourceTypeChangedWithTags",
			Old:      `[{"tags":{"a":"b"},"type":"redis"}]`,
			New:      `[{"tags":{"a":"b"},"type":"mysql"}]`,
			Expected: false,
		},
		{
			Name:     "TagsChangedWithResourceType",
			Old:      `[{"tags":{"a":"b"},"type":"redis"}]`,
			New:      `[{"tags":{"a":"c"},"type":"redis"}]`,
			Expected: false,
		},
		{
			Name:     "AccessRuleOrderingSuppressed",
			Old:      `[{"tags":{"a":"b"}},{"type":"redis"}]`,
			New:      `[{"type":"redis"},{"tags":{"a":"b"}}]`,
			Expected: true,
		},
		{
			Name:     "AddAccessRule",
			Old:      `[{"type":"redis"}]`,
			New:      `[{"type":"redis"},{"tags":{"a":"b"}}]`,
			Expected: false,
		},
		{
			Name:     "RemoveAccessRule",
			Old:      `[{"type":"redis"},{"tags":{"a":"b"}}]`,
			New:      `[{"type":"redis"}]`,
			Expected: false,
		},
		{
			Name:     "AddDuplicateAccessRule",
			Old:      `[{"type":"redis"}]`,
			New:      `[{"type":"redis"},{"type":"redis"}]`,
			Expected: false,
		},
		{
			Name:     "EmptyRulesExactMatch",
			Old:      `[]`,
			New:      `[]`,
			Expected: true,
		},
		{
			Name:     "JunkSuppressed",
			Old:      `[{"type":"redis"}]`,
			New:      `[{"type":"redis","asdf":"true"}]`,
			Expected: true,
		},
		{
			Name:     "InvalidJSON",
			Old:      `[]`,
			New:      `asdf`,
			Expected: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := accessRulesJSONDiffSuppress("", testCase.Old, testCase.New, nil)
			if result != testCase.Expected {
				t.Fatalf("expected result %v, got %v", testCase.Expected, result)
			}
		})
	}
}
