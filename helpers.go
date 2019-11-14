package sdm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	apiV1 "github.com/strongdm/strongdm-sdk-go"
)

type apiCrudOperation func(d *schema.ResourceData, client *apiV1.Client) error

func wrapCrudOperation(op apiCrudOperation) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		client := m.(*apiV1.Client)
		return op(d, client)
	}
}

func stringFromMap(m map[string]interface{}, key string) string {
	value := m[key]
	if value == nil {
		return ""
	}
	return value.(string)
}

func stringFromResourceData(d *schema.ResourceData, key string) string {
	value := d.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func boolFromResourceData(d *schema.ResourceData, key string) bool {
	value := d.Get(key)
	if value == nil {
		return false
	}
	return value.(bool)
}