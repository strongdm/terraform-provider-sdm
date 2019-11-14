package sdm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	apiV1 "github.com/strongdm/strongdm-sdk-go"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"sdm_node": resourceNode(),
			"sdm_role": resourceRole(),
			
		},
		ConfigureFunc: func(*schema.ResourceData) (interface{}, error) {
			client, err := apiV1.New("localhost:8889")
			if err != nil {
				return nil, fmt.Errorf("cannot dial API server: %w", err)
			}
			return client, nil
		},
	}
}
