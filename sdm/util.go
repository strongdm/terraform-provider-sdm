package sdm

import (
	"bytes"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// optionalJSONDiffSuppress returns true if the old and new values provided are
// semantically equivalent, or if the new value is an empty string. this
// indicates the diff should be suppressed.
func optionalJSONDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if new == "" {
		// this is an optional field. for a required field, an empty string
		// would be an invalid value. but let's assume it's a no-op here.
		return true
	}
	var oldCompacted bytes.Buffer
	_ = json.Compact(&oldCompacted, []byte(old))
	var newCompacted bytes.Buffer
	_ = json.Compact(&newCompacted, []byte(new))
	return bytes.Equal(oldCompacted.Bytes(), newCompacted.Bytes())
}
