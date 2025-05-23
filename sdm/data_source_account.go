// Code generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: wrapCrudOperation(dataSourceAccountList),
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"manager_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"permission_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"permissions": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"suspended": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     tagsElemType,
			},
			"accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"service": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A Service is a service account that can connect to resources they are granted directly, or granted via roles. Services are typically automated jobs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Service.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Service.",
									},
									"suspended": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "The Service's suspended state.",
									},
									"tags": {
										Type:        schema.TypeMap,
										Elem:        tagsElemType,
										Optional:    true,
										Description: "Tags is a map of key, value pairs.",
									},
								},
							},
						},
						"token": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A Token is an account providing tokenized access for automation or integration use. Tokens include admin tokens, API keys, and SCIM tokens.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Corresponds to the type of token, e.g. api or admin-token.",
									},
									"deadline": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The timestamp when the Token will expire.",
									},
									"duration": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Duration from token creation to expiration.",
									},
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Token.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Token.",
									},
									"permissions": {
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional:    true,
										Description: "Permissions assigned to the token, e.g. role:create.",
									},
									"rekeyed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The timestamp when the Token was last rekeyed.",
									},
									"suspended": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Reserved for future use.  Always false for tokens.",
									},
									"tags": {
										Type:        schema.TypeMap,
										Elem:        tagsElemType,
										Optional:    true,
										Description: "Tags is a map of key, value pairs.",
									},
								},
							},
						},
						"user": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A User can connect to resources they are granted directly, or granted via roles.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scim": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "SCIM contains the raw SCIM metadata for the user. This is a read-only field.",
									},
									"email": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The User's email address. Must be unique.",
									},
									"external_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "External ID is an alternative unique ID this user is represented by within an external service.",
									},
									"first_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The User's first name.",
									},
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the User.",
									},
									"last_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The User's last name.",
									},
									"managed_by": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Managed By is a read only field for what service manages this user, e.g. StrongDM, Okta, Azure.",
									},
									"manager_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Manager ID is the ID of the user's manager. This field is empty when the user has no manager.",
									},
									"permission_level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "PermissionLevel is the user's permission level e.g. admin, DBA, user.",
									},
									"resolved_manager_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resolved Manager ID is the ID of the user's manager derived from the manager_id, if present, or from the SCIM metadata. This is a read-only field that's only populated for get and list.",
									},
									"suspended": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Suspended is a read only field for the User's suspended state.",
									},
									"tags": {
										Type:        schema.TypeMap,
										Elem:        tagsElemType,
										Optional:    true,
										Description: "Tags is a map of key, value pairs.",
									},
								},
							},
						},
					},
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func init() {
	dataSourcesMap["sdm_account"] = dataSourceAccount
}

func convertAccountFilterToPlumbing(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOkExists("type"); ok {
		filter += "type:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("SCIM"); ok {
		filter += "scim:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("account_type"); ok {
		filter += "accounttype:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("deadline"); ok {
		filter += "deadline:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("duration"); ok {
		filter += "duration:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("email"); ok {
		filter += "email:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("external_id"); ok {
		filter += "externalid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("first_name"); ok {
		filter += "firstname:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("last_name"); ok {
		filter += "lastname:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("managed_by"); ok {
		filter += "managedby:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("manager_id"); ok {
		filter += "managerid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("permission_levelRW"); ok {
		filter += "permissionlevelrw:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("permissions"); ok {
		filter += "permissions:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("rekeyed"); ok {
		filter += "rekeyed:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("resolved_manager_id"); ok {
		filter += "resolvedmanagerid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("suspended"); ok {
		filter += "suspended:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("suspendedRO"); ok {
		filter += "suspendedro:? "
		args = append(args, v)
	}
	if v, ok := d.GetOkExists("tags"); ok {
		tags := convertTagsToPlumbing(v)
		for kk, vv := range tags {
			filter += "tag:?=? "
			args = append(args, kk, vv)
		}
	}
	return filter, args
}

func dataSourceAccountList(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	filter, args := convertAccountFilterToPlumbing(d)
	resp, err := cc.Accounts().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Accounts %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]map[string][]entity, 1)
	output[0] = map[string][]entity{
		"service": {},
	}
	for resp.Next() {
		ids = append(ids, resp.Value().GetID())
		switch v := resp.Value().(type) {
		case *sdm.Service:
			output[0]["service"] = append(output[0]["service"], entity{
				"id":        (v.ID),
				"name":      (v.Name),
				"suspended": (v.Suspended),
				"tags":      convertTagsToPorcelain(v.Tags),
			})
		case *sdm.Token:
			output[0]["token"] = append(output[0]["token"], entity{
				"account_type": (v.AccountType),
				"deadline":     convertTimestampToPorcelain(v.Deadline),
				"duration":     convertDurationToPorcelain(v.Duration),
				"id":           (v.ID),
				"name":         (v.Name),
				"permissions":  (v.Permissions),
				"rekeyed":      convertTimestampToPorcelain(v.Rekeyed),
				"suspended":    (v.Suspended),
				"tags":         convertTagsToPorcelain(v.Tags),
			})
		case *sdm.User:
			output[0]["user"] = append(output[0]["user"], entity{
				"scim":                (v.SCIM),
				"email":               (v.Email),
				"external_id":         (v.ExternalID),
				"first_name":          (v.FirstName),
				"id":                  (v.ID),
				"last_name":           (v.LastName),
				"managed_by":          (v.ManagedBy),
				"manager_id":          (v.ManagerID),
				"permission_level":    (v.PermissionLevel),
				"resolved_manager_id": (v.ResolvedManagerID),
				"suspended":           (v.Suspended),
				"tags":                convertTagsToPorcelain(v.Tags),
			})
		}
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("accounts", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("Account" + filter + fmt.Sprint(args...))
	return nil
}
