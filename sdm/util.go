package sdm

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// accessRulesJSONDiffSuppress returns true if the old and new access policies
// provided are semantically equivalent, or if the new value is an empty string.
// this indicates the diff should be suppressed.
func accessRulesJSONDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if new == "" {
		// this is an optional field. for a required field, an empty string
		// would be an invalid value. but let's assume it's a no-op here.
		return true
	}
	oldRules, err := parseAccessRulesJSON(old)
	if err != nil {
		return false
	}
	newRules, err := parseAccessRulesJSON(new)
	if err != nil {
		return false
	}
	if accessRulesEqual(oldRules, newRules) {
		return true
	}
	return false
}

// accessRule is the user-visible JSON representation of an access rule.
type accessRule struct {
	ResourceIDs  []string          `json:"ids,omitempty"  yaml:"ids,omitempty,flow"`
	ResourceType string            `json:"type,omitempty" yaml:"type,omitempty"`
	Tags         map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func parseAccessRulesJSON(rawJSON string) ([]accessRule, error) {
	var rules []accessRule
	err := json.Unmarshal([]byte(rawJSON), &rules)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal access rules: %w", err)
	}
	return rules, nil
}

// accessRulesEqual returns true if the two given sets of access rules are
// semantically equal
func accessRulesEqual(before []accessRule, after []accessRule) bool {
	beforeMap := make(map[string]struct{})
	for _, ar := range before {
		beforeMap[ar.Key()] = struct{}{}
	}
	afterMap := make(map[string]struct{})
	for _, ar := range after {
		afterMap[ar.Key()] = struct{}{}
	}
	if len(afterMap) != len(after) {
		// contains duplicate access rules
		return false
	}
	for key := range beforeMap {
		if _, ok := afterMap[key]; !ok {
			return false
		}
	}
	for key := range afterMap {
		if _, ok := beforeMap[key]; !ok {
			return false
		}
	}
	return true
}

func parseResourceID(id string) string {
	const resourcePrefix = "rs-"
	if id == "" {
		return resourcePrefix + fmt.Sprintf("%016x", uint64(0))
	}
	if !strings.HasPrefix(id, resourcePrefix) {
		return id
	}
	d, err := strconv.ParseUint(id[len(resourcePrefix):], 16, 64)
	if err != nil {
		return id
	}
	return resourcePrefix + fmt.Sprintf("%016x", uint64(d))
}

func (ar accessRule) Key() string {
	tags := make([]string, 0, len(ar.Tags))
	for k, v := range ar.Tags {
		k, _ := json.Marshal(k)
		v, _ := json.Marshal(v)
		tags = append(tags, string(k)+"="+string(v))
	}
	resourceIDs := make([]string, 0, len(ar.ResourceIDs))
	for _, id := range ar.ResourceIDs {
		quotedID, _ := json.Marshal(parseResourceID(id))
		resourceIDs = append(resourceIDs, string(quotedID))
	}
	sort.Slice(tags, func(a, b int) bool { return tags[a] < tags[b] })
	sort.Slice(resourceIDs, func(a, b int) bool { return resourceIDs[a] < resourceIDs[b] })
	return ar.ResourceType + ":" + strings.Join(resourceIDs, ",") + ":" + strings.Join(tags, ",")
}
