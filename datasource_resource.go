package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceResource() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceResourceList),
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
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"athena": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"secret_access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"output": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"region": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"big_query": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"private_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"project": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"endpoint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"cassandra": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"druid": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"dynamo_db": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"secret_access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"region": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"endpoint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"amazon_es": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"region": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"secret_access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"endpoint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"elastic": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"http_basic_auth": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"healthcheck_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"headers_blacklist": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"default_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"http_no_auth": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"healthcheck_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"headers_blacklist": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"default_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"http_auth": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"healthcheck_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"auth_header": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"headers_blacklist": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"default_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"kubernetes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"certificate_authority": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"certificate_authority_filename": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"client_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"client_certificate_filename": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"client_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"client_key_filename": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"kubernetes_basic_auth": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"amazon_eks": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"endpoint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"secret_access_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"certificate_authority": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"certificate_authority_filename": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"region": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"cluster_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"google_gke": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"endpoint": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"certificate_authority": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"certificate_authority_filename": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"service_account_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"service_account_key_filename": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"kubernetes_service_account": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"token": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"memcached": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"mongo_legacy_host": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"auth_database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"replica_set": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"mongo_legacy_replicaset": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"auth_database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"replica_set": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"connect_to_replica": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"mongo_host": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"auth_database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"mongo_replica_set": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"auth_database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"replica_set": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"connect_to_replica": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"mysql": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"aurora_mysql": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"clustrix": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"maria": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"memsql": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"oracle": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"postgres": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"override_database": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"aurora_postgres": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"override_database": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"greenplum": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"override_database": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"cockroach": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"override_database": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"redshift": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"override_database": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"presto": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"rdp": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"redis": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"elasticache_redis": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tls_required": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"snowflake": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"schema": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"sql_server": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"database": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"schema": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"override_database": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ssh": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"sybase": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"sybase_iq": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"teradata": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the Resource.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the Resource.",
									},
									"hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"port_override": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
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

func resourceFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOk("type"); ok {
		filter += "type:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("hostname"); ok {
		filter += "hostname:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("port"); ok {
		filter += "port:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("username"); ok {
		filter += "username:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceResourceList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := resourceFilterFromResourceData(d)
	resp, err := cc.Resources().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Resources %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]map[string][]entity, 1)
	output[0] = map[string][]entity{
		"athena": {},
	}
	for resp.Next() {
		ids = append(ids, resp.Value().GetID())
		switch v := resp.Value().(type) {
		case *apiv1.Athena:
			output[0]["athena"] = append(output[0]["athena"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"output":            v.Output,
				"port_override":     v.PortOverride,
				"region":            v.Region,
			})
		case *apiv1.BigQuery:
			output[0]["big_query"] = append(output[0]["big_query"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"private_key":   v.PrivateKey,
				"project":       v.Project,
				"port_override": v.PortOverride,
				"endpoint":      v.Endpoint,
				"username":      v.Username,
			})
		case *apiv1.Cassandra:
			output[0]["cassandra"] = append(output[0]["cassandra"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			})
		case *apiv1.Druid:
			output[0]["druid"] = append(output[0]["druid"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
			})
		case *apiv1.DynamoDB:
			output[0]["dynamo_db"] = append(output[0]["dynamo_db"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"region":            v.Region,
				"endpoint":          v.Endpoint,
				"port_override":     v.PortOverride,
			})
		case *apiv1.AmazonES:
			output[0]["amazon_es"] = append(output[0]["amazon_es"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"region":            v.Region,
				"secret_access_key": v.SecretAccessKey,
				"endpoint":          v.Endpoint,
				"access_key":        v.AccessKey,
				"port_override":     v.PortOverride,
			})
		case *apiv1.Elastic:
			output[0]["elastic"] = append(output[0]["elastic"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			})
		case *apiv1.HTTPBasicAuth:
			output[0]["http_basic_auth"] = append(output[0]["http_basic_auth"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"username":          v.Username,
				"password":          v.Password,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			})
		case *apiv1.HTTPNoAuth:
			output[0]["http_no_auth"] = append(output[0]["http_no_auth"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			})
		case *apiv1.HTTPAuth:
			output[0]["http_auth"] = append(output[0]["http_auth"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"auth_header":       v.AuthHeader,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			})
		case *apiv1.Kubernetes:
			output[0]["kubernetes"] = append(output[0]["kubernetes"], entity{
				"id":                             v.ID,
				"name":                           v.Name,
				"hostname":                       v.Hostname,
				"port":                           v.Port,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"client_certificate":             v.ClientCertificate,
				"client_certificate_filename":    v.ClientCertificateFilename,
				"client_key":                     v.ClientKey,
				"client_key_filename":            v.ClientKeyFilename,
			})
		case *apiv1.KubernetesBasicAuth:
			output[0]["kubernetes_basic_auth"] = append(output[0]["kubernetes_basic_auth"], entity{
				"id":       v.ID,
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"username": v.Username,
				"password": v.Password,
			})
		case *apiv1.AmazonEKS:
			output[0]["amazon_eks"] = append(output[0]["amazon_eks"], entity{
				"id":                             v.ID,
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"access_key":                     v.AccessKey,
				"secret_access_key":              v.SecretAccessKey,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"region":                         v.Region,
				"cluster_name":                   v.ClusterName,
			})
		case *apiv1.GoogleGKE:
			output[0]["google_gke"] = append(output[0]["google_gke"], entity{
				"id":                             v.ID,
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"service_account_key":            v.ServiceAccountKey,
				"service_account_key_filename":   v.ServiceAccountKeyFilename,
			})
		case *apiv1.KubernetesServiceAccount:
			output[0]["kubernetes_service_account"] = append(output[0]["kubernetes_service_account"], entity{
				"id":       v.ID,
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"token":    v.Token,
			})
		case *apiv1.Memcached:
			output[0]["memcached"] = append(output[0]["memcached"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"port":          v.Port,
			})
		case *apiv1.MongoLegacyHost:
			output[0]["mongo_legacy_host"] = append(output[0]["mongo_legacy_host"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"auth_database": v.AuthDatabase,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
				"replica_set":   v.ReplicaSet,
				"tls_required":  v.TlsRequired,
			})
		case *apiv1.MongoLegacyReplicaset:
			output[0]["mongo_legacy_replicaset"] = append(output[0]["mongo_legacy_replicaset"], entity{
				"id":                 v.ID,
				"name":               v.Name,
				"hostname":           v.Hostname,
				"auth_database":      v.AuthDatabase,
				"port_override":      v.PortOverride,
				"username":           v.Username,
				"password":           v.Password,
				"port":               v.Port,
				"replica_set":        v.ReplicaSet,
				"connect_to_replica": v.ConnectToReplica,
				"tls_required":       v.TlsRequired,
			})
		case *apiv1.MongoHost:
			output[0]["mongo_host"] = append(output[0]["mongo_host"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"auth_database": v.AuthDatabase,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			})
		case *apiv1.MongoReplicaSet:
			output[0]["mongo_replica_set"] = append(output[0]["mongo_replica_set"], entity{
				"id":                 v.ID,
				"name":               v.Name,
				"hostname":           v.Hostname,
				"auth_database":      v.AuthDatabase,
				"port_override":      v.PortOverride,
				"username":           v.Username,
				"password":           v.Password,
				"port":               v.Port,
				"replica_set":        v.ReplicaSet,
				"connect_to_replica": v.ConnectToReplica,
				"tls_required":       v.TlsRequired,
			})
		case *apiv1.Mysql:
			output[0]["mysql"] = append(output[0]["mysql"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			})
		case *apiv1.AuroraMysql:
			output[0]["aurora_mysql"] = append(output[0]["aurora_mysql"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			})
		case *apiv1.Clustrix:
			output[0]["clustrix"] = append(output[0]["clustrix"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			})
		case *apiv1.Maria:
			output[0]["maria"] = append(output[0]["maria"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			})
		case *apiv1.Memsql:
			output[0]["memsql"] = append(output[0]["memsql"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			})
		case *apiv1.Oracle:
			output[0]["oracle"] = append(output[0]["oracle"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
				"port_override": v.PortOverride,
				"tls_required":  v.TlsRequired,
			})
		case *apiv1.Postgres:
			output[0]["postgres"] = append(output[0]["postgres"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			})
		case *apiv1.AuroraPostgres:
			output[0]["aurora_postgres"] = append(output[0]["aurora_postgres"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			})
		case *apiv1.Greenplum:
			output[0]["greenplum"] = append(output[0]["greenplum"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			})
		case *apiv1.Cockroach:
			output[0]["cockroach"] = append(output[0]["cockroach"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			})
		case *apiv1.Redshift:
			output[0]["redshift"] = append(output[0]["redshift"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			})
		case *apiv1.Presto:
			output[0]["presto"] = append(output[0]["presto"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"username":      v.Username,
				"tls_required":  v.TlsRequired,
			})
		case *apiv1.RDP:
			output[0]["rdp"] = append(output[0]["rdp"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
			})
		case *apiv1.Redis:
			output[0]["redis"] = append(output[0]["redis"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
			})
		case *apiv1.ElasticacheRedis:
			output[0]["elasticache_redis"] = append(output[0]["elasticache_redis"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			})
		case *apiv1.Snowflake:
			output[0]["snowflake"] = append(output[0]["snowflake"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"schema":        v.Schema,
				"port_override": v.PortOverride,
			})
		case *apiv1.SQLServer:
			output[0]["sql_server"] = append(output[0]["sql_server"], entity{
				"id":                v.ID,
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"schema":            v.Schema,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			})
		case *apiv1.SSH:
			output[0]["ssh"] = append(output[0]["ssh"], entity{
				"id":       v.ID,
				"name":     v.Name,
				"hostname": v.Hostname,
				"username": v.Username,
				"port":     v.Port,
			})
		case *apiv1.Sybase:
			output[0]["sybase"] = append(output[0]["sybase"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"password":      v.Password,
			})
		case *apiv1.SybaseIQ:
			output[0]["sybase_iq"] = append(output[0]["sybase_iq"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"password":      v.Password,
			})
		case *apiv1.Teradata:
			output[0]["teradata"] = append(output[0]["teradata"], entity{
				"id":            v.ID,
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
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
	err = d.Set("resources", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("Resource" + filter + fmt.Sprint(args...))
	return nil
}
