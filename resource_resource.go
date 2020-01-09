package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func resourceResource() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceResourceCreate),
		Read:   wrapCrudOperation(resourceResourceRead),
		Update: wrapCrudOperation(resourceResourceUpdate),
		Delete: wrapCrudOperation(resourceResourceDelete),
		Schema: map[string]*schema.Schema{
			"athena": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"output": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"private_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"project": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"healthcheck_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"headers_blacklist": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"default_path": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"healthcheck_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"headers_blacklist": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"default_path": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"healthcheck_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_header": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"headers_blacklist": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"default_path": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"certificate_authority": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority_filename": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"client_certificate_filename": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"client_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"client_key_filename": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority_filename": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"cluster_name": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority_filename": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"service_account_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"service_account_key_filename": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"replica_set": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"connect_to_replica": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"replica_set": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"connect_to_replica": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"schema": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"replica_set": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"connect_to_replica": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"schema": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"schema": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
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
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
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

func resourceFromResourceData(d *schema.ResourceData) apiv1.Resource {
	if list := d.Get("athena").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Athena{
			ID:              d.Id(),
			Name:            stringFromMap(raw, "name"),
			AccessKey:       stringFromMap(raw, "access_key"),
			SecretAccessKey: stringFromMap(raw, "secret_access_key"),
			Output:          stringFromMap(raw, "output"),
			PortOverride:    int32FromMap(raw, "port_override"),
			Region:          stringFromMap(raw, "region"),
		}
	}
	if list := d.Get("big_query").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.BigQuery{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PrivateKey:   stringFromMap(raw, "private_key"),
			Project:      stringFromMap(raw, "project"),
			PortOverride: int32FromMap(raw, "port_override"),
			Endpoint:     stringFromMap(raw, "endpoint"),
			Username:     stringFromMap(raw, "username"),
		}
	}
	if list := d.Get("cassandra").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Cassandra{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("druid").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Druid{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("dynamo_db").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.DynamoDB{
			ID:              d.Id(),
			Name:            stringFromMap(raw, "name"),
			AccessKey:       stringFromMap(raw, "access_key"),
			SecretAccessKey: stringFromMap(raw, "secret_access_key"),
			Region:          stringFromMap(raw, "region"),
			Endpoint:        stringFromMap(raw, "endpoint"),
			PortOverride:    int32FromMap(raw, "port_override"),
		}
	}
	if list := d.Get("amazon_es").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.AmazonES{
			ID:              d.Id(),
			Name:            stringFromMap(raw, "name"),
			Region:          stringFromMap(raw, "region"),
			SecretAccessKey: stringFromMap(raw, "secret_access_key"),
			Endpoint:        stringFromMap(raw, "endpoint"),
			AccessKey:       stringFromMap(raw, "access_key"),
			PortOverride:    int32FromMap(raw, "port_override"),
		}
	}
	if list := d.Get("elastic").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Elastic{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("http_basic_auth").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.HTTPBasicAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("http_no_auth").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.HTTPNoAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("http_auth").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.HTTPAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			AuthHeader:       stringFromMap(raw, "auth_header"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("kubernetes").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Kubernetes{
			ID:                           d.Id(),
			Name:                         stringFromMap(raw, "name"),
			Hostname:                     stringFromMap(raw, "hostname"),
			Port:                         int32FromMap(raw, "port"),
			CertificateAuthority:         stringFromMap(raw, "certificate_authority"),
			CertificateAuthorityFilename: stringFromMap(raw, "certificate_authority_filename"),
			ClientCertificate:            stringFromMap(raw, "client_certificate"),
			ClientCertificateFilename:    stringFromMap(raw, "client_certificate_filename"),
			ClientKey:                    stringFromMap(raw, "client_key"),
			ClientKeyFilename:            stringFromMap(raw, "client_key_filename"),
		}
	}
	if list := d.Get("kubernetes_basic_auth").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.KubernetesBasicAuth{
			ID:       d.Id(),
			Name:     stringFromMap(raw, "name"),
			Hostname: stringFromMap(raw, "hostname"),
			Port:     int32FromMap(raw, "port"),
			Username: stringFromMap(raw, "username"),
			Password: stringFromMap(raw, "password"),
		}
	}
	if list := d.Get("amazon_eks").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.AmazonEKS{
			ID:                           d.Id(),
			Name:                         stringFromMap(raw, "name"),
			Endpoint:                     stringFromMap(raw, "endpoint"),
			AccessKey:                    stringFromMap(raw, "access_key"),
			SecretAccessKey:              stringFromMap(raw, "secret_access_key"),
			CertificateAuthority:         stringFromMap(raw, "certificate_authority"),
			CertificateAuthorityFilename: stringFromMap(raw, "certificate_authority_filename"),
			Region:                       stringFromMap(raw, "region"),
			ClusterName:                  stringFromMap(raw, "cluster_name"),
		}
	}
	if list := d.Get("google_gke").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.GoogleGKE{
			ID:                           d.Id(),
			Name:                         stringFromMap(raw, "name"),
			Endpoint:                     stringFromMap(raw, "endpoint"),
			CertificateAuthority:         stringFromMap(raw, "certificate_authority"),
			CertificateAuthorityFilename: stringFromMap(raw, "certificate_authority_filename"),
			ServiceAccountKey:            stringFromMap(raw, "service_account_key"),
			ServiceAccountKeyFilename:    stringFromMap(raw, "service_account_key_filename"),
		}
	}
	if list := d.Get("memcached").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Memcached{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("mongo_legacy_host").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.MongoLegacyHost{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			AuthDatabase:     stringFromMap(raw, "auth_database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Port:             int32FromMap(raw, "port"),
			ReplicaSet:       stringFromMap(raw, "replica_set"),
			ConnectToReplica: boolFromMap(raw, "connect_to_replica"),
			TlsRequired:      boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mongo_legacy_replicaset").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.MongoLegacyReplicaset{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			AuthDatabase:     stringFromMap(raw, "auth_database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Port:             int32FromMap(raw, "port"),
			ReplicaSet:       stringFromMap(raw, "replica_set"),
			ConnectToReplica: boolFromMap(raw, "connect_to_replica"),
			TlsRequired:      boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mongo_host").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.MongoHost{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			AuthDatabase: stringFromMap(raw, "auth_database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
			Schema:       stringFromMap(raw, "schema"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mongo_replica_set").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.MongoReplicaSet{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			AuthDatabase:     stringFromMap(raw, "auth_database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Port:             int32FromMap(raw, "port"),
			ReplicaSet:       stringFromMap(raw, "replica_set"),
			ConnectToReplica: boolFromMap(raw, "connect_to_replica"),
			TlsRequired:      boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mysql").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Mysql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("aurora_mysql").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.AuroraMysql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("clustrix").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Clustrix{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("maria").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Maria{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("memsql").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Memsql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("oracle").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Oracle{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Port:         int32FromMap(raw, "port"),
			PortOverride: int32FromMap(raw, "port_override"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("postgres").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Postgres{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("aurora_postgres").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.AuroraPostgres{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("greenplum").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Greenplum{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("cockroach").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Cockroach{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("redshift").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Redshift{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("presto").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Presto{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			Username:     stringFromMap(raw, "username"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("rdp").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.RDP{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("redis").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Redis{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("elasticache_redis").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.ElasticacheRedis{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("snowflake").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Snowflake{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Schema:       stringFromMap(raw, "schema"),
			PortOverride: int32FromMap(raw, "port_override"),
		}
	}
	if list := d.Get("sql_server").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.SQLServer{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Schema:           stringFromMap(raw, "schema"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("ssh").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.SSH{
			ID:       d.Id(),
			Name:     stringFromMap(raw, "name"),
			Hostname: stringFromMap(raw, "hostname"),
			Username: stringFromMap(raw, "username"),
			Port:     int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("sybase").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Sybase{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			Password:     stringFromMap(raw, "password"),
		}
	}
	if list := d.Get("teradata").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Teradata{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	return nil
}

func resourceResourceCreate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.Resources().Create(ctx, resourceFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create Resource %s: %w", "", err)
	}
	d.SetId(resp.Resource.GetID())
	return resourceResourceRead(d, cc)
}

func resourceResourceRead(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Resources().Get(ctx, d.Id())
	var errNotFound *apiv1.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Resource %s: %w", d.Id(), err)
	}
	switch v := resp.Resource.(type) {
	case *apiv1.Athena:
		d.Set("athena", []map[string]interface{}{
			{
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"output":            v.Output,
				"port_override":     v.PortOverride,
				"region":            v.Region,
			},
		})
	case *apiv1.BigQuery:
		d.Set("big_query", []map[string]interface{}{
			{
				"name":          v.Name,
				"private_key":   v.PrivateKey,
				"project":       v.Project,
				"port_override": v.PortOverride,
				"endpoint":      v.Endpoint,
				"username":      v.Username,
			},
		})
	case *apiv1.Cassandra:
		d.Set("cassandra", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *apiv1.Druid:
		d.Set("druid", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
			},
		})
	case *apiv1.DynamoDB:
		d.Set("dynamo_db", []map[string]interface{}{
			{
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"region":            v.Region,
				"endpoint":          v.Endpoint,
				"port_override":     v.PortOverride,
			},
		})
	case *apiv1.AmazonES:
		d.Set("amazon_es", []map[string]interface{}{
			{
				"name":              v.Name,
				"region":            v.Region,
				"secret_access_key": v.SecretAccessKey,
				"endpoint":          v.Endpoint,
				"access_key":        v.AccessKey,
				"port_override":     v.PortOverride,
			},
		})
	case *apiv1.Elastic:
		d.Set("elastic", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *apiv1.HTTPBasicAuth:
		d.Set("http_basic_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"username":          v.Username,
				"password":          v.Password,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *apiv1.HTTPNoAuth:
		d.Set("http_no_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *apiv1.HTTPAuth:
		d.Set("http_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"auth_header":       v.AuthHeader,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *apiv1.Kubernetes:
		d.Set("kubernetes", []map[string]interface{}{
			{
				"name":                           v.Name,
				"hostname":                       v.Hostname,
				"port":                           v.Port,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"client_certificate":             v.ClientCertificate,
				"client_certificate_filename":    v.ClientCertificateFilename,
				"client_key":                     v.ClientKey,
				"client_key_filename":            v.ClientKeyFilename,
			},
		})
	case *apiv1.KubernetesBasicAuth:
		d.Set("kubernetes_basic_auth", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"username": v.Username,
				"password": v.Password,
			},
		})
	case *apiv1.AmazonEKS:
		d.Set("amazon_eks", []map[string]interface{}{
			{
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"access_key":                     v.AccessKey,
				"secret_access_key":              v.SecretAccessKey,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"region":                         v.Region,
				"cluster_name":                   v.ClusterName,
			},
		})
	case *apiv1.GoogleGKE:
		d.Set("google_gke", []map[string]interface{}{
			{
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"service_account_key":            v.ServiceAccountKey,
				"service_account_key_filename":   v.ServiceAccountKeyFilename,
			},
		})
	case *apiv1.Memcached:
		d.Set("memcached", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *apiv1.MongoLegacyHost:
		d.Set("mongo_legacy_host", []map[string]interface{}{
			{
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
			},
		})
	case *apiv1.MongoLegacyReplicaset:
		d.Set("mongo_legacy_replicaset", []map[string]interface{}{
			{
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
			},
		})
	case *apiv1.MongoHost:
		d.Set("mongo_host", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"auth_database": v.AuthDatabase,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
				"schema":        v.Schema,
				"tls_required":  v.TlsRequired,
			},
		})
	case *apiv1.MongoReplicaSet:
		d.Set("mongo_replica_set", []map[string]interface{}{
			{
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
			},
		})
	case *apiv1.Mysql:
		d.Set("mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *apiv1.AuroraMysql:
		d.Set("aurora_mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *apiv1.Clustrix:
		d.Set("clustrix", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *apiv1.Maria:
		d.Set("maria", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *apiv1.Memsql:
		d.Set("memsql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *apiv1.Oracle:
		d.Set("oracle", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
				"port_override": v.PortOverride,
				"tls_required":  v.TlsRequired,
			},
		})
	case *apiv1.Postgres:
		d.Set("postgres", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *apiv1.AuroraPostgres:
		d.Set("aurora_postgres", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *apiv1.Greenplum:
		d.Set("greenplum", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *apiv1.Cockroach:
		d.Set("cockroach", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *apiv1.Redshift:
		d.Set("redshift", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *apiv1.Presto:
		d.Set("presto", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"username":      v.Username,
				"tls_required":  v.TlsRequired,
			},
		})
	case *apiv1.RDP:
		d.Set("rdp", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *apiv1.Redis:
		d.Set("redis", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
			},
		})
	case *apiv1.ElasticacheRedis:
		d.Set("elasticache_redis", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *apiv1.Snowflake:
		d.Set("snowflake", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"schema":        v.Schema,
				"port_override": v.PortOverride,
			},
		})
	case *apiv1.SQLServer:
		d.Set("sql_server", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"schema":            v.Schema,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *apiv1.SSH:
		d.Set("ssh", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"username": v.Username,
				"port":     v.Port,
			},
		})
	case *apiv1.Sybase:
		d.Set("sybase", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"password":      v.Password,
			},
		})
	case *apiv1.Teradata:
		d.Set("teradata", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	}
	return nil
}

func resourceResourceUpdate(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	resp, err := cc.Resources().Update(ctx, resourceFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Resource %s: %w", d.Id(), err)
	}
	d.SetId(resp.Resource.GetID())
	return resourceResourceRead(d, cc)
}

func resourceResourceDelete(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	_, err := cc.Resources().Delete(ctx, d.Id())
	return err
}
