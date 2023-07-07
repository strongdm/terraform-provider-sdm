// Code generated by protogen. DO NOT EDIT.

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

func convertAccessRulesToPorcelain(plumbing sdm.AccessRules) string {
	return string(plumbing)
}

func convertAccessRulesToPlumbing(porcelain interface{}) sdm.AccessRules {
	if porcelain == nil {
		return sdm.AccessRules("")
	}
	return sdm.AccessRules(porcelain.(string))
}

func convertTagsToPlumbing(porcelain interface{}) sdm.Tags {
	if porcelain == nil {
		return nil
	}
	tags, ok := porcelain.(map[string]interface{})
	if !ok {
		return nil
	}
	res := sdm.Tags{}
	for key, porcelain := range tags {
		str, ok := porcelain.(string)
		if !ok {
			continue
		}
		res[key] = str
	}
	return res
}

func convertTagsToPorcelain(tags sdm.Tags) map[string]interface{} {
	res := map[string]interface{}{}
	for key, value := range tags {
		res[key] = value
	}
	return res
}

var tagsElemType = &schema.Schema{
	Type: schema.TypeString,
}

func convertRepeatedNodeMaintenanceWindowToPlumbing(porcelain interface{}) []*sdm.NodeMaintenanceWindow {
	if porcelain == nil {
		return nil
	}
	windows, ok := porcelain.([]interface{})
	if !ok {
		return nil
	}
	ws := make([]*sdm.NodeMaintenanceWindow, len(windows))
	for i, w := range windows {
		window, ok := w.(map[string]interface{})
		if !ok {
			return nil
		}
		ri, ok := window["require_idleness"].(bool)
		if !ok {
			return nil
		}
		cs, ok := window["cron_schedule"].(string)
		if !ok {
			return nil
		}
		ws[i] = &sdm.NodeMaintenanceWindow{
			CronSchedule:    cs,
			RequireIdleness: ri,
		}
	}
	return ws
}

func convertRepeatedNodeMaintenanceWindowToPorcelain(ws []*sdm.NodeMaintenanceWindow) interface{} {
	windows := make([]map[string]interface{}, len(ws))
	for i, w := range ws {
		windows[i] = map[string]interface{}{
			"require_idleness": w.RequireIdleness,
			"cron_schedule":    w.CronSchedule,
		}
	}
	return windows
}

var nodeMaintenanceWindowElemType = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"cron_schedule": {
			Type:     schema.TypeString,
			Required: true,
		},
		"require_idleness": {
			Type:     schema.TypeBool,
			Required: true,
		},
	},
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

func convertStringToPlumbing(porcelain interface{}) string {
	if porcelain == nil {
		return ""
	}
	return porcelain.(string)
}

func convertInt32ToPlumbing(porcelain interface{}) int32 {
	if porcelain == nil {
		return 0
	}
	return int32(porcelain.(int))
}

func convertBoolToPlumbing(porcelain interface{}) bool {
	if porcelain == nil {
		return false
	}
	return porcelain.(bool)
}
