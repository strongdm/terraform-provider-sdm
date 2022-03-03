// This file was generated by protogen. DO NOT EDIT.

// Package sdm is a terraform provider for strongDM
package sdm

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

type apiCrudOperation func(ctx context.Context, d *schema.ResourceData, client *sdm.Client) error

func wrapCrudOperation(op apiCrudOperation) func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		client := meta.(*sdm.Client)
		err := op(ctx, d, client)
		return diag.FromErr(err)
	}
}

func convertAccessRulesToMap(accessRules sdm.AccessRules) string {
	return string(accessRules)
}

func convertAccessRulesFromResourceData(d *schema.ResourceData, key string) sdm.AccessRules {
	value := d.Get(key)
	if value == nil {
		return sdm.AccessRules("")
	}
	return sdm.AccessRules(value.(string))
}

func convertTagsFromMap(m map[string]interface{}, key string) sdm.Tags {
	value := m[key]
	if value == nil {
		return nil
	}
	tags, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	return tagsFromMap(tags)
}

func convertTagsFromResourceData(d *schema.ResourceData, key string) sdm.Tags {
	value := d.Get(key)
	if value == nil {
		return nil
	}
	tags, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	return tagsFromMap(tags)
}

func convertTagsToMap(tags sdm.Tags) map[string]interface{} {
	res := map[string]interface{}{}
	for key, value := range tags {
		res[key] = value
	}
	return res
}

func tagsFromMap(m map[string]interface{}) sdm.Tags {
	res := sdm.Tags{}
	for key, value := range m {
		str, ok := value.(string)
		if !ok {
			continue
		}
		res[key] = str
	}
	return res
}

var tagsElemType = &schema.Schema{
	Type: schema.TypeString,
}

func accessRulesDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	var oldRules []interface{}
	if err := json.Unmarshal([]byte(old), &oldRules); err != nil {
		return false
	}
	var newRules []interface{}
	if err := json.Unmarshal([]byte(new), &newRules); err != nil {
		return false
	}
	oldNormalized, err := json.Marshal(oldRules)
	if err != nil {
		return false
	}
	newNormalized, err := json.Marshal(newRules)
	if err != nil {
		return false
	}
	if !bytes.Equal(oldNormalized, newNormalized) {
		return false
	}
	return true
}

func convertStringFromMap(m map[string]interface{}, key string) string {
	value := m[key]
	if value == nil {
		return ""
	}
	return value.(string)
}

func convertStringFromResourceData(d *schema.ResourceData, key string) string {
	value := d.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func convertInt32FromMap(m map[string]interface{}, key string) int32 {
	value := m[key]
	if value == nil {
		return 0
	}
	return int32(value.(int))
}

func convertInt32FromResourceData(d *schema.ResourceData, key string) int32 {
	value := d.Get(key)
	if value == nil {
		return 0
	}
	return int32(value.(int))
}

func convertBoolFromMap(m map[string]interface{}, key string) bool {
	value := m[key]
	if value == nil {
		return false
	}
	return value.(bool)
}

func convertBoolFromResourceData(d *schema.ResourceData, key string) bool {
	value := d.Get(key)
	if value == nil {
		return false
	}
	return value.(bool)
}

func fullSecretStorePath(raw map[string]interface{}, cred string) string {
	path := convertStringFromMap(raw, "secret_store_"+cred+"_path")
	key := convertStringFromMap(raw, "secret_store_"+cred+"_key")
	if key != "" {
		return path + "?key=" + key
	}
	return path
}
