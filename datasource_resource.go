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
			"access_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_database": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_header": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_authority": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_authority_filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_certificate_filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_key_filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connect_to_replica": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"database": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"headers_blacklist": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"healthcheck_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"output": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_database": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port_override": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replica_set": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schema": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_access_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_account_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_account_key_filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tls_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"athena": {
							Type:        schema.TypeList,
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
							Optional:    true,
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
	if v, ok := d.GetOk("access_key"); ok {
		filter += "access_key:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("auth_database"); ok {
		filter += "auth_database:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("auth_header"); ok {
		filter += "auth_header:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("certificate_authority"); ok {
		filter += "certificate_authority:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("certificate_authority_filename"); ok {
		filter += "certificate_authority_filename:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("client_certificate"); ok {
		filter += "client_certificate:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("client_certificate_filename"); ok {
		filter += "client_certificate_filename:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("client_key"); ok {
		filter += "client_key:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("client_key_filename"); ok {
		filter += "client_key_filename:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("cluster_name"); ok {
		filter += "cluster_name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("connect_to_replica"); ok {
		filter += "connect_to_replica:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("database"); ok {
		filter += "database:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("default_path"); ok {
		filter += "default_path:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("endpoint"); ok {
		filter += "endpoint:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("headers_blacklist"); ok {
		filter += "headers_blacklist:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("healthcheck_path"); ok {
		filter += "healthcheck_path:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("healthy"); ok {
		filter += "healthy:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("hostname"); ok {
		filter += "hostname:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("output"); ok {
		filter += "output:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("override_database"); ok {
		filter += "override_database:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("password"); ok {
		filter += "password:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("port"); ok {
		filter += "port:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("port_override"); ok {
		filter += "port_override:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("private_key"); ok {
		filter += "private_key:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("project"); ok {
		filter += "project:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("public_key"); ok {
		filter += "public_key:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("region"); ok {
		filter += "region:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("replica_set"); ok {
		filter += "replica_set:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("schema"); ok {
		filter += "schema:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("secret_access_key"); ok {
		filter += "secret_access_key:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("service_account_key"); ok {
		filter += "service_account_key:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("service_account_key_filename"); ok {
		filter += "service_account_key_filename:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("subdomain"); ok {
		filter += "subdomain:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("tls_required"); ok {
		filter += "tls_required:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("token"); ok {
		filter += "token:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("url"); ok {
		filter += "url:? "
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
		return fmt.Errorf("cannot list Resource %s: %w", d.Id(), err)
	}
	vList := []map[string]interface{}{}
	for resp.Next() {
		switch v := resp.Value().(type) {
		case *apiv1.Athena:
			vList = append(vList,
				map[string]interface{}{
					"athena": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"access_key":        v.AccessKey,
							"secret_access_key": v.SecretAccessKey,
							"output":            v.Output,
							"port_override":     v.PortOverride,
							"region":            v.Region,
						},
					},
				},
			)
		case *apiv1.BigQuery:
			vList = append(vList,
				map[string]interface{}{
					"big_query": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"private_key":   v.PrivateKey,
							"project":       v.Project,
							"port_override": v.PortOverride,
							"endpoint":      v.Endpoint,
							"username":      v.Username,
						},
					},
				},
			)
		case *apiv1.Cassandra:
			vList = append(vList,
				map[string]interface{}{
					"cassandra": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"port_override": v.PortOverride,
							"port":          v.Port,
							"tls_required":  v.TlsRequired,
						},
					},
				},
			)
		case *apiv1.Druid:
			vList = append(vList,
				map[string]interface{}{
					"druid": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"port_override": v.PortOverride,
							"username":      v.Username,
							"password":      v.Password,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.DynamoDB:
			vList = append(vList,
				map[string]interface{}{
					"dynamo_db": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"access_key":        v.AccessKey,
							"secret_access_key": v.SecretAccessKey,
							"region":            v.Region,
							"endpoint":          v.Endpoint,
							"port_override":     v.PortOverride,
						},
					},
				},
			)
		case *apiv1.AmazonES:
			vList = append(vList,
				map[string]interface{}{
					"amazon_es": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"region":            v.Region,
							"secret_access_key": v.SecretAccessKey,
							"endpoint":          v.Endpoint,
							"access_key":        v.AccessKey,
							"port_override":     v.PortOverride,
						},
					},
				},
			)
		case *apiv1.Elastic:
			vList = append(vList,
				map[string]interface{}{
					"elastic": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"port_override": v.PortOverride,
							"port":          v.Port,
							"tls_required":  v.TlsRequired,
						},
					},
				},
			)
		case *apiv1.HTTPBasicAuth:
			vList = append(vList,
				map[string]interface{}{
					"http_basic_auth": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"url":               v.Url,
							"healthcheck_path":  v.HealthcheckPath,
							"username":          v.Username,
							"password":          v.Password,
							"headers_blacklist": v.HeadersBlacklist,
							"default_path":      v.DefaultPath,
						},
					},
				},
			)
		case *apiv1.HTTPNoAuth:
			vList = append(vList,
				map[string]interface{}{
					"http_no_auth": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"url":               v.Url,
							"healthcheck_path":  v.HealthcheckPath,
							"headers_blacklist": v.HeadersBlacklist,
							"default_path":      v.DefaultPath,
						},
					},
				},
			)
		case *apiv1.HTTPAuth:
			vList = append(vList,
				map[string]interface{}{
					"http_auth": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"url":               v.Url,
							"healthcheck_path":  v.HealthcheckPath,
							"auth_header":       v.AuthHeader,
							"headers_blacklist": v.HeadersBlacklist,
							"default_path":      v.DefaultPath,
						},
					},
				},
			)
		case *apiv1.Kubernetes:
			vList = append(vList,
				map[string]interface{}{
					"kubernetes": []map[string]interface{}{
						{
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
						},
					},
				},
			)
		case *apiv1.KubernetesBasicAuth:
			vList = append(vList,
				map[string]interface{}{
					"kubernetes_basic_auth": []map[string]interface{}{
						{
							"id":       v.ID,
							"name":     v.Name,
							"hostname": v.Hostname,
							"port":     v.Port,
							"username": v.Username,
							"password": v.Password,
						},
					},
				},
			)
		case *apiv1.AmazonEKS:
			vList = append(vList,
				map[string]interface{}{
					"amazon_eks": []map[string]interface{}{
						{
							"id":                             v.ID,
							"name":                           v.Name,
							"endpoint":                       v.Endpoint,
							"access_key":                     v.AccessKey,
							"secret_access_key":              v.SecretAccessKey,
							"certificate_authority":          v.CertificateAuthority,
							"certificate_authority_filename": v.CertificateAuthorityFilename,
							"region":                         v.Region,
							"cluster_name":                   v.ClusterName,
						},
					},
				},
			)
		case *apiv1.GoogleGKE:
			vList = append(vList,
				map[string]interface{}{
					"google_gke": []map[string]interface{}{
						{
							"id":                             v.ID,
							"name":                           v.Name,
							"endpoint":                       v.Endpoint,
							"certificate_authority":          v.CertificateAuthority,
							"certificate_authority_filename": v.CertificateAuthorityFilename,
							"service_account_key":            v.ServiceAccountKey,
							"service_account_key_filename":   v.ServiceAccountKeyFilename,
						},
					},
				},
			)
		case *apiv1.KubernetesServiceAccount:
			vList = append(vList,
				map[string]interface{}{
					"kubernetes_service_account": []map[string]interface{}{
						{
							"id":       v.ID,
							"name":     v.Name,
							"hostname": v.Hostname,
							"port":     v.Port,
							"token":    v.Token,
						},
					},
				},
			)
		case *apiv1.Memcached:
			vList = append(vList,
				map[string]interface{}{
					"memcached": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.MongoLegacyHost:
			vList = append(vList,
				map[string]interface{}{
					"mongo_legacy_host": []map[string]interface{}{
						{
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
						},
					},
				},
			)
		case *apiv1.MongoLegacyReplicaset:
			vList = append(vList,
				map[string]interface{}{
					"mongo_legacy_replicaset": []map[string]interface{}{
						{
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
						},
					},
				},
			)
		case *apiv1.MongoHost:
			vList = append(vList,
				map[string]interface{}{
					"mongo_host": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"auth_database": v.AuthDatabase,
							"port_override": v.PortOverride,
							"username":      v.Username,
							"password":      v.Password,
							"port":          v.Port,
							"tls_required":  v.TlsRequired,
						},
					},
				},
			)
		case *apiv1.MongoReplicaSet:
			vList = append(vList,
				map[string]interface{}{
					"mongo_replica_set": []map[string]interface{}{
						{
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
						},
					},
				},
			)
		case *apiv1.Mysql:
			vList = append(vList,
				map[string]interface{}{
					"mysql": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"database":      v.Database,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.AuroraMysql:
			vList = append(vList,
				map[string]interface{}{
					"aurora_mysql": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"database":      v.Database,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.Clustrix:
			vList = append(vList,
				map[string]interface{}{
					"clustrix": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"database":      v.Database,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.Maria:
			vList = append(vList,
				map[string]interface{}{
					"maria": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"database":      v.Database,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.Memsql:
			vList = append(vList,
				map[string]interface{}{
					"memsql": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"database":      v.Database,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.Oracle:
			vList = append(vList,
				map[string]interface{}{
					"oracle": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"database":      v.Database,
							"port":          v.Port,
							"port_override": v.PortOverride,
							"tls_required":  v.TlsRequired,
						},
					},
				},
			)
		case *apiv1.Postgres:
			vList = append(vList,
				map[string]interface{}{
					"postgres": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"hostname":          v.Hostname,
							"username":          v.Username,
							"password":          v.Password,
							"database":          v.Database,
							"port_override":     v.PortOverride,
							"port":              v.Port,
							"override_database": v.OverrideDatabase,
						},
					},
				},
			)
		case *apiv1.AuroraPostgres:
			vList = append(vList,
				map[string]interface{}{
					"aurora_postgres": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"hostname":          v.Hostname,
							"username":          v.Username,
							"password":          v.Password,
							"database":          v.Database,
							"port_override":     v.PortOverride,
							"port":              v.Port,
							"override_database": v.OverrideDatabase,
						},
					},
				},
			)
		case *apiv1.Greenplum:
			vList = append(vList,
				map[string]interface{}{
					"greenplum": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"hostname":          v.Hostname,
							"username":          v.Username,
							"password":          v.Password,
							"database":          v.Database,
							"port_override":     v.PortOverride,
							"port":              v.Port,
							"override_database": v.OverrideDatabase,
						},
					},
				},
			)
		case *apiv1.Cockroach:
			vList = append(vList,
				map[string]interface{}{
					"cockroach": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"hostname":          v.Hostname,
							"username":          v.Username,
							"password":          v.Password,
							"database":          v.Database,
							"port_override":     v.PortOverride,
							"port":              v.Port,
							"override_database": v.OverrideDatabase,
						},
					},
				},
			)
		case *apiv1.Redshift:
			vList = append(vList,
				map[string]interface{}{
					"redshift": []map[string]interface{}{
						{
							"id":                v.ID,
							"name":              v.Name,
							"hostname":          v.Hostname,
							"username":          v.Username,
							"password":          v.Password,
							"database":          v.Database,
							"port_override":     v.PortOverride,
							"port":              v.Port,
							"override_database": v.OverrideDatabase,
						},
					},
				},
			)
		case *apiv1.Presto:
			vList = append(vList,
				map[string]interface{}{
					"presto": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"password":      v.Password,
							"database":      v.Database,
							"port_override": v.PortOverride,
							"port":          v.Port,
							"username":      v.Username,
							"tls_required":  v.TlsRequired,
						},
					},
				},
			)
		case *apiv1.RDP:
			vList = append(vList,
				map[string]interface{}{
					"rdp": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.Redis:
			vList = append(vList,
				map[string]interface{}{
					"redis": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"port_override": v.PortOverride,
							"password":      v.Password,
							"port":          v.Port,
						},
					},
				},
			)
		case *apiv1.ElasticacheRedis:
			vList = append(vList,
				map[string]interface{}{
					"elasticache_redis": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"port_override": v.PortOverride,
							"password":      v.Password,
							"port":          v.Port,
							"tls_required":  v.TlsRequired,
						},
					},
				},
			)
		case *apiv1.Snowflake:
			vList = append(vList,
				map[string]interface{}{
					"snowflake": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"database":      v.Database,
							"schema":        v.Schema,
							"port_override": v.PortOverride,
						},
					},
				},
			)
		case *apiv1.SQLServer:
			vList = append(vList,
				map[string]interface{}{
					"sql_server": []map[string]interface{}{
						{
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
						},
					},
				},
			)
		case *apiv1.SSH:
			vList = append(vList,
				map[string]interface{}{
					"ssh": []map[string]interface{}{
						{
							"id":       v.ID,
							"name":     v.Name,
							"hostname": v.Hostname,
							"username": v.Username,
							"port":     v.Port,
						},
					},
				},
			)
		case *apiv1.Sybase:
			vList = append(vList,
				map[string]interface{}{
					"sybase": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"port_override": v.PortOverride,
							"port":          v.Port,
							"password":      v.Password,
						},
					},
				},
			)
		case *apiv1.Teradata:
			vList = append(vList,
				map[string]interface{}{
					"teradata": []map[string]interface{}{
						{
							"id":            v.ID,
							"name":          v.Name,
							"hostname":      v.Hostname,
							"username":      v.Username,
							"password":      v.Password,
							"port_override": v.PortOverride,
							"port":          v.Port,
						},
					},
				},
			)
		}
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
