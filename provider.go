package sdm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	apiV1 "github.com/strongdm/strongdm-sdk-go"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_access_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A GUID identifying the API key used to authenticate with the StrongDM API.",
			},

			"api_secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A base64 encoded secret key used to authenticate with the StrongDM API.",
			},

			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "api.strongdm.com:443",
				Description: "The host and port of the StrongDM API endpoint.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sdm_account_attachment": resourceAccountAttachment(),
			"sdm_account_grant":      resourceAccountGrant(),
			"sdm_account":            resourceAccount(),
			"sdm_node":               resourceNode(),
			"sdm_resource":           resourceResource(),
			"sdm_role_attachment":    resourceRoleAttachment(),
			"sdm_role_grant":         resourceRoleGrant(),
			"sdm_role":               resourceRole(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sdm_account_grant":   dataSourceAccountGrant(),
			"sdm_account":         dataSourceAccount(),
			"sdm_node":            dataSourceNode(),
			"sdm_resource":        dataSourceResource(),
			"sdm_role_attachment": dataSourceRoleAttachment(),
			"sdm_role":            dataSourceRole(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			client, err := apiV1.New(d.Get("host").(string), d.Get("api_access_key").(string), d.Get("api_secret_key").(string))
			if err != nil {
				return nil, fmt.Errorf("cannot dial API server: %w", err)
			}
			return client, nil
		},
	}
}
