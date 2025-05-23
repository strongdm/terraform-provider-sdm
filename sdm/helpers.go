// Code generated by protogen. DO NOT EDIT.

// Package sdm is a terraform provider for strongDM
package sdm

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

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

func convertBytesToPorcelain(plumbing []byte) string {
	return base64.StdEncoding.EncodeToString(plumbing)
}

func convertAccessRulesToPlumbing(porcelain interface{}) sdm.AccessRules {
	if porcelain == nil {
		return sdm.AccessRules("")
	}
	return sdm.AccessRules(porcelain.(string))
}

func convertAccessRuleToPorcelain(plumbing sdm.AccessRule) string {
	return string(plumbing)
}

func convertAccessRuleToPlumbing(porcelain interface{}) sdm.AccessRule {
	if porcelain == nil {
		return sdm.AccessRule("")
	}
	return sdm.AccessRule(porcelain.(string))
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

func convertBytesToPlumbing(porcelain interface{}) []byte {
	if porcelain == nil {
		return nil
	}
	if v, ok := porcelain.(string); ok {
		result, _ := base64.StdEncoding.DecodeString(v)
		return result
	}
	return nil
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

func accessRuleDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	var oldRule []interface{}
	if err := json.Unmarshal([]byte(old), &oldRule); err != nil {
		return false
	}
	var newRule []interface{}
	if err := json.Unmarshal([]byte(new), &newRule); err != nil {
		return false
	}
	oldNormalized, err := json.Marshal(oldRule)
	if err != nil {
		return false
	}
	newNormalized, err := json.Marshal(newRule)
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

func convertRepeatedStringToPlumbing(porcelain interface{}) []string {
	if porcelain == nil {
		return nil
	}
	interfaces, ok := porcelain.([]interface{})
	if !ok {
		return nil
	}
	strs := make([]string, 0, len(interfaces))
	for _, interf := range interfaces {
		str, ok := interf.(string)
		if !ok {
			return nil
		}
		strs = append(strs, str)
	}
	return strs
}

func convertInt32ToPlumbing(porcelain interface{}) int32 {
	if porcelain == nil {
		return 0
	}
	return int32(porcelain.(int))
}

func convertInt64ToPlumbing(porcelain interface{}) int64 {
	if porcelain == nil {
		return 0
	}
	return int64(porcelain.(int))
}

func convertUint32ToPlumbing(porcelain interface{}) uint32 {
	if porcelain == nil {
		return 0
	}
	return uint32(porcelain.(int))
}

func convertBoolToPlumbing(porcelain interface{}) bool {
	if porcelain == nil {
		return false
	}
	return porcelain.(bool)
}

func convertTimestampToPorcelain(t time.Time) interface{} {
	return t.UTC().Format(time.RFC3339)
}

func convertTimestampToPlumbing(porcelain interface{}) time.Time {
	if porcelain == nil {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, porcelain.(string))
	return t.UTC()
}

func convertDurationToPorcelain(d time.Duration) interface{} {
	return d.String()
}

func convertDurationToPlumbing(porcelain interface{}) time.Duration {
	if porcelain == nil {
		return time.Duration(0)
	}
	d, _ := time.ParseDuration(porcelain.(string))
	return d
}

func convertRepeatedApprovalFlowStepToPlumbing(porcelain interface{}) []*sdm.ApprovalFlowStep {
	if porcelain == nil {
		return nil
	}
	approvalSteps, ok := porcelain.([]interface{})
	if !ok {
		return nil
	}
	ws := make([]*sdm.ApprovalFlowStep, len(approvalSteps))
	for i, w := range approvalSteps {
		approvalStep, ok := w.(map[string]interface{})
		if !ok {
			return nil
		}
		// parse out fields for this step
		quantifier, ok := approvalStep["quantifier"].(string)
		if !ok {
			return nil
		}
		skipAfter := convertDurationToPlumbing(approvalStep["skip_after"])
		approvers := convertRepeatedApprovalFlowApproverToPlumbing(approvalStep["approvers"])
		// parse out list of approvers
		if approvers == nil {
			return nil
		}
		ws[i] = &sdm.ApprovalFlowStep{
			Quantifier: quantifier,
			SkipAfter:  skipAfter,
			Approvers:  approvers,
		}
	}
	return ws
}

func convertRepeatedApprovalFlowStepToPorcelain(ws []*sdm.ApprovalFlowStep) interface{} {
	approvalSteps := make([]map[string]interface{}, len(ws))
	for i, w := range ws {
		approvalSteps[i] = map[string]interface{}{
			"quantifier": w.Quantifier,
			"approvers":  convertRepeatedApprovalFlowApproverToPorcelain(w.Approvers),
			"skip_after": convertDurationToPorcelain(w.SkipAfter),
		}
	}
	return approvalSteps
}

var approvalFlowStepElemType = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"approvers": {
			Type:        schema.TypeList,
			Elem:        approvalFlowApproverElemType,
			Required:    true,
			Description: "The approvers for this approval step",
		},
		"quantifier": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "any",
			Description: "Whether any or all approvers are required to approve for this approval step (optional, defaults to any)",
		},
		"skip_after": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "0s",
			Description: "Duration after which this approval step will be skipped if no approval is given (optional, if not provided this step must be manually approved)",
		},
	},
}

func convertRepeatedApprovalFlowApproverToPlumbing(porcelain interface{}) []*sdm.ApprovalFlowApprover {
	if porcelain == nil {
		return nil
	}
	approvers, ok := porcelain.([]interface{})
	if !ok {
		return nil
	}
	ws := make([]*sdm.ApprovalFlowApprover, len(approvers))
	for i, w := range approvers {
		approver, ok := w.(map[string]interface{})
		if !ok {
			return nil
		}
		ws[i] = &sdm.ApprovalFlowApprover{}
		if approver["account_id"] != nil {
			accountID, ok := approver["account_id"].(string)
			if !ok {
				return nil
			}
			ws[i].AccountID = accountID
		}
		if approver["role_id"] != nil {
			roleID, ok := approver["role_id"].(string)
			if !ok {
				return nil
			}
			ws[i].RoleID = roleID
		}
		if approver["reference"] != nil {
			reference, ok := approver["reference"].(string)
			if !ok {
				return nil
			}
			ws[i].Reference = reference
		}
	}
	return ws
}

func convertRepeatedApprovalFlowApproverToPorcelain(ws []*sdm.ApprovalFlowApprover) interface{} {
	approvers := make([]map[string]interface{}, len(ws))
	for i, w := range ws {
		approvers[i] = map[string]interface{}{}
		if w.AccountID != "" {
			approvers[i]["account_id"] = w.AccountID
		}
		if w.RoleID != "" {
			approvers[i]["role_id"] = w.RoleID
		}
		if w.Reference != "" {
			approvers[i]["reference"] = w.Reference
		}
	}
	return approvers
}

var approvalFlowApproverElemType = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"account_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The account id of the approver (only one of account_id, role_id, or reference may be present for one approver)",
		},
		"role_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The role id of the approver (only one of account_id, role_id, or reference may be present for one approver)",
		},
		"reference": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A reference to an approver: 'manager-of-requester' or 'manager-of-manager-of-requester' (only one of account_id, role_id, or reference may be present for one approver)",
		},
	},
}
