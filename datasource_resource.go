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
			"athena": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
	if d.Id() != "" {
		filter += "id:? "
		args = append(args, d.Id())
	}
	// todo
	return filter, args
}

func dataSourceResourceList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := resourceFilterFromResourceData(d)
	resp, err := cc.Resources().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list Resource %s: %w", d.Id(), err)
	}
	vList := []string{}
	for resp.Next() {
		v := resp.Value()
		vList = append(vList, v.GetID())
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", err)
	}

	err = d.Set("resources", vList)
	if err != nil {
		return fmt.Errorf("cannot set vList: %w", err)
	}
	d.SetId("Resource" + filter + fmt.Sprint(args...))
	return nil
}
