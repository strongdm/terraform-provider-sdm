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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"client_certificate": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"client_key": {
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"service_account_key": {
							Type:        schema.TypeString,
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
					},
				},
			},
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
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Port number override.",
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
						"output": {
							Type:        schema.TypeString,
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
	if list := d.Get("kubernetes").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Kubernetes{
			ID:                   d.Id(),
			Name:                 stringFromMap(raw, "name"),
			PortOverride:         int32FromMap(raw, "port_override"),
			Hostname:             stringFromMap(raw, "hostname"),
			Port:                 int32FromMap(raw, "port"),
			CertificateAuthority: stringFromMap(raw, "certificate_authority"),
			ClientCertificate:    stringFromMap(raw, "client_certificate"),
			ClientKey:            stringFromMap(raw, "client_key"),
		}
	}
	if list := d.Get("amazon_eks").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.AmazonEks{
			ID:                   d.Id(),
			Name:                 stringFromMap(raw, "name"),
			PortOverride:         int32FromMap(raw, "port_override"),
			Endpoint:             stringFromMap(raw, "endpoint"),
			AccessKey:            stringFromMap(raw, "access_key"),
			SecretAccessKey:      stringFromMap(raw, "secret_access_key"),
			CertificateAuthority: stringFromMap(raw, "certificate_authority"),
			Region:               stringFromMap(raw, "region"),
			ClusterName:          stringFromMap(raw, "cluster_name"),
		}
	}
	if list := d.Get("google_gke").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.GoogleGke{
			ID:                   d.Id(),
			Name:                 stringFromMap(raw, "name"),
			PortOverride:         int32FromMap(raw, "port_override"),
			Endpoint:             stringFromMap(raw, "endpoint"),
			CertificateAuthority: stringFromMap(raw, "certificate_authority"),
			ServiceAccountKey:    stringFromMap(raw, "service_account_key"),
		}
	}
	if list := d.Get("ssh").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Ssh{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PortOverride: int32FromMap(raw, "port_override"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("http_basic_auth").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.HttpBasicAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			PortOverride:     int32FromMap(raw, "port_override"),
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
		return &apiv1.HttpNoAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("http_auth").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.HttpAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			AuthHeader:       stringFromMap(raw, "auth_header"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("mysql").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Mysql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PortOverride: int32FromMap(raw, "port_override"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("aurora_mysql").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.AuroraMysql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PortOverride: int32FromMap(raw, "port_override"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("clustrix").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Clustrix{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PortOverride: int32FromMap(raw, "port_override"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("maria").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Maria{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PortOverride: int32FromMap(raw, "port_override"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("memsql").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Memsql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PortOverride: int32FromMap(raw, "port_override"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("athena").([]interface{}); len(list) > 0 {
		raw := list[0].(map[string]interface{})
		return &apiv1.Athena{
			ID:              d.Id(),
			Name:            stringFromMap(raw, "name"),
			PortOverride:    int32FromMap(raw, "port_override"),
			AccessKey:       stringFromMap(raw, "access_key"),
			SecretAccessKey: stringFromMap(raw, "secret_access_key"),
			Region:          stringFromMap(raw, "region"),
			Output:          stringFromMap(raw, "output"),
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
	case *apiv1.Kubernetes:
		d.Set("kubernetes", []map[string]interface{}{
			{
				"name":                  v.Name,
				"port_override":         v.PortOverride,
				"hostname":              v.Hostname,
				"port":                  v.Port,
				"certificate_authority": v.CertificateAuthority,
				"client_certificate":    v.ClientCertificate,
				"client_key":            v.ClientKey,
			},
		})
	case *apiv1.AmazonEKS:
		d.Set("amazon_eks", []map[string]interface{}{
			{
				"name":                  v.Name,
				"port_override":         v.PortOverride,
				"endpoint":              v.Endpoint,
				"access_key":            v.AccessKey,
				"secret_access_key":     v.SecretAccessKey,
				"certificate_authority": v.CertificateAuthority,
				"region":                v.Region,
				"cluster_name":          v.ClusterName,
			},
		})
	case *apiv1.GoogleGKE:
		d.Set("google_gke", []map[string]interface{}{
			{
				"name":                  v.Name,
				"port_override":         v.PortOverride,
				"endpoint":              v.Endpoint,
				"certificate_authority": v.CertificateAuthority,
				"service_account_key":   v.ServiceAccountKey,
			},
		})
	case *apiv1.SSH:
		d.Set("ssh", []map[string]interface{}{
			{
				"name":          v.Name,
				"port_override": v.PortOverride,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port":          v.Port,
			},
		})
	case *apiv1.HTTPBasicAuth:
		d.Set("http_basic_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"port_override":     v.PortOverride,
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
				"port_override":     v.PortOverride,
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
				"port_override":     v.PortOverride,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"auth_header":       v.AuthHeader,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *apiv1.Mysql:
		d.Set("mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"port_override": v.PortOverride,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
			},
		})
	case *apiv1.AuroraMysql:
		d.Set("aurora_mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"port_override": v.PortOverride,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
			},
		})
	case *apiv1.Clustrix:
		d.Set("clustrix", []map[string]interface{}{
			{
				"name":          v.Name,
				"port_override": v.PortOverride,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
			},
		})
	case *apiv1.Maria:
		d.Set("maria", []map[string]interface{}{
			{
				"name":          v.Name,
				"port_override": v.PortOverride,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
			},
		})
	case *apiv1.Memsql:
		d.Set("memsql", []map[string]interface{}{
			{
				"name":          v.Name,
				"port_override": v.PortOverride,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
			},
		})
	case *apiv1.Athena:
		d.Set("athena", []map[string]interface{}{
			{
				"name":              v.Name,
				"port_override":     v.PortOverride,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"region":            v.Region,
				"output":            v.Output,
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
