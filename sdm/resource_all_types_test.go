package sdm

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func TestAccSDMResource_UpdateAllTypes(t *testing.T) {
	initAcceptanceTest(t)
	certificateAuthorityRaw := []string{
		"\"-----BEGIN CERTIFICATE-----",
		"MIIC5zCCAc+gAwIBAgIBATANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p",
		"a3ViZUNBMB4XDTE5MDkyNDE2MDgwMVoXDTI5MDkyMjE2MDgwMVowFTETMBEGA1UE",
		"AxMKbWluaWt1YmVDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALpg",
		"4/Kjq9CXaU7lSO9v5m7wFHg2+t1U8q7MmO4M9tDmhcgCR3x2lnQZR3cXSuq/BzV+",
		"9VAalARIiA7JYXQRJvucqb9Aj7Q2A/wC9D2CovCCnfslZRGGhcw3zVRM0UwP0zg6",
		"sgBixcls9YqKP2Td2+TwnWYMd43ah9MSO823VLNJi56JXoV0fyrdERpL5+5y6aUM",
		"X3qXXZusVYGwJ4c2ucqWRWcXDDArxYVqNJV7GONeDee2HMBC+k12CUJU7HxIzZlW",
		"QPWSRr9j3nTeJyC8sgbNJDJWLYvIv4j0+0OTUvB/2f8T7vbeWR5VeD7PtmzVN4/M",
		"4U8jI64qUsPZBZGhXg0CAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMB0GA1UdJQQW",
		"MBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3",
		"DQEBCwUAA4IBAQB33WY0z4Aw1RCLYK06V/HtcuZdh5d6TWKHg/b9ncNbroJaEJGk",
		"xPaIuVBzajygn2vdyv8iRX0yDgYM3lZ/P5SLbhD2oZRhi9qgOW6oNuN99EinUnQw",
		"jp1R7F1ug1yxRgZLGgDlBL83jRJV0AgmxgOrbsd8sAVXKI8p70RFYAyBoGd9Pj9D",
		"nohAJ+7Eh2xuPqnU2J7bzddP+ECoucSZ/Ex4qWF+0RFyFUwk/c/2nH9AvTNCUY2H",
		"d/ref47hNcAAjTHG7OKqSUHhAkaOuGUdnGEyNuGHREd11S0x+oTjFDoNaJ6O4PDF",
		"VTzHCUQlVvCxKklV3pArPBB7vJdjBFvZcreB",
		"-----END CERTIFICATE-----\"",
	}
	certificateAuthority := strings.Join(certificateAuthorityRaw, "\\n")
	_ = certificateAuthority

	type testCase struct {
		resource string
		pairs    [][2]string
	}
	tcs := []testCase{
		{
			resource: "aks",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", certificateAuthority},
				{"client_certificate", certificateAuthority},
				{"client_key", `"key"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"port", `443`},
			},
		},
		{
			resource: "aks_basic_auth",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "aks_service_account",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"port", `443`},
				{"token", `"token"`},
			},
		},
		{
			resource: "aks_service_account_user_impersonation",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
				{"token", `"token"`},
			},
		},
		{
			resource: "aks_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", certificateAuthority},
				{"client_certificate", certificateAuthority},
				{"client_key", `"key"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
			},
		},
		{
			resource: "amazon_eks",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", certificateAuthority},
				{"cluster_name", `"cluster_name"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
			},
		},
		{
			resource: "amazon_eks_instance_profile",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", certificateAuthority},
				{"cluster_name", `"cluster_name"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
			},
		},
		{
			resource: "amazon_eks_instance_profile_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", certificateAuthority},
				{"cluster_name", `"cluster_name"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
			},
		},
		{
			resource: "amazon_eks_user_impersonation",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"certificate_authority", certificateAuthority},
				{"cluster_name", `"cluster_name"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
			},
		},
		{
			resource: "amazon_es",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
			},
		},
		{
			resource: "amazon_esiam",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"tls_required", `true`},
			},
		},
		{
			resource: "amazonmq_amqp_091",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "athena",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"egress_filter", `"name:value"`},
				{"name", `"name"`},
				{"output", `"output"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
			},
		},
		{
			resource: "athena_iam",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"name", `"name"`},
				{"output", `"output"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
			},
		},
		{
			resource: "aurora_mysql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "aurora_mysql_iam",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_assumption_arn", `"role_assumption_arn"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "aurora_postgres",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "aurora_postgres_iam",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_assumption_arn", `"role_assumption_arn"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "aws",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_region", `"healthcheck_region"`},
				{"name", `"name"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
			},
		},
		{
			resource: "aws_console",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"enable_env_variables", `true`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"session_expiry", `20000`},
				{"subdomain", `"subdomain"`},
			},
		},
		{
			resource: "aws_console_static_key_pair",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"egress_filter", `"name:value"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
				{"session_expiry", `20000`},
				{"subdomain", `"subdomain"`},
			},
		},
		{
			resource: "aws_instance_profile",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"enable_env_variables", `true`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
			},
		},
		{
			resource: "azure",
			pairs: [][2]string{
				{"app_id", `"app_id"`},
				{"egress_filter", `"name:value"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"tenant_id", `"tenant_id"`},
			},
		},
		{
			resource: "azure_certificate",
			pairs: [][2]string{
				{"app_id", `"app_id"`},
				{"client_certificate", certificateAuthority},
				{"egress_filter", `"name:value"`},
				{"name", `"name"`},
				{"tenant_id", `"tenant_id"`},
			},
		},
		{
			resource: "azure_mysql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "azure_postgres",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "azure_postgres_managed_identity",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "big_query",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"name"`},
				{"private_key", `"private_key"`},
				{"project", `"project"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "cassandra",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "citus",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "click_house_http",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"url", `"https://app.strongdm.com"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "click_house_my_sql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "click_house_tcp",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "clustrix",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "cockroach",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "couchbase_database",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"n_1_ql_port", `443`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "couchbase_web_ui",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"subdomain", `"subdomain"`},
				{"url", `"https://app.strongdm.com"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "db_2_i",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "db_2_luw",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "document_db_host",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "document_db_host_iam",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
				{"region", `"region"`},
			},
		},
		{
			resource: "document_db_replica_set",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"connect_to_replica", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"replica_set", `"replica_set"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "druid",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "dynamo_db",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
			},
		},
		{
			resource: "dynamo_dbiam",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"name"`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
			},
		},
		{
			resource: "elastic",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "elasticache_redis",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "gcp",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"keyfile", `"keyfile"`},
				{"name", `"name"`},
				{"scopes", `"scopes"`},
			},
		},
		{
			resource: "gcp_console",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"session_expiry", `20000`},
				{"subdomain", `"subdomain"`},
				{"workforce_pool_id", `"workforce_pool_id"`},
				{"workforce_provider_id", `"workforce_provider_id"`},
			},
		},
		{
			resource: "gcpwif",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"project_id", `"project_id"`},
				{"scopes", `"scopes"`},
				{"session_expiry", `20000`},
				{"workforce_pool_id", `"workforce_pool_id"`},
				{"workforce_provider_id", `"workforce_provider_id"`},
			},
		},
		{
			resource: "google_gke",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", certificateAuthority},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"service_account_key", `"{}"`},
			},
		},
		{
			resource: "google_gke_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", certificateAuthority},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"name", `"name"`},
				{"service_account_key", `"{}"`},
			},
		},
		{
			resource: "greenplum",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "http_auth",
			pairs: [][2]string{
				{"auth_header", `"auth_header"`},
				{"default_path", `"/"`},
				{"egress_filter", `"name:value"`},
				{"headers_blacklist", `"headers_blacklist"`},
				{"healthcheck_path", `"/"`},
				{"host_override", `"host_override"`},
				{"name", `"name"`},
				{"subdomain", `"subdomain"`},
				{"url", `"https://app.strongdm.com"`},
			},
		},
		{
			resource: "http_basic_auth",
			pairs: [][2]string{
				{"default_path", `"/"`},
				{"egress_filter", `"name:value"`},
				{"headers_blacklist", `"headers_blacklist"`},
				{"healthcheck_path", `"/"`},
				{"host_override", `"host_override"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"subdomain", `"subdomain"`},
				{"url", `"https://app.strongdm.com"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "http_no_auth",
			pairs: [][2]string{
				{"default_path", `"/"`},
				{"egress_filter", `"name:value"`},
				{"headers_blacklist", `"headers_blacklist"`},
				{"healthcheck_path", `"/"`},
				{"host_override", `"host_override"`},
				{"name", `"name"`},
				{"subdomain", `"subdomain"`},
				{"url", `"https://app.strongdm.com"`},
			},
		},
		{
			resource: "kubernetes",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", certificateAuthority},
				{"client_certificate", certificateAuthority},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"port", `443`},
			},
		},
		{
			resource: "kubernetes_basic_auth",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "kubernetes_service_account",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"name"`},
				{"port", `443`},
				{"token", `"token"`},
			},
		},
		{
			resource: "kubernetes_service_account_user_impersonation",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
				{"token", `"token"`},
			},
		},
		{
			resource: "kubernetes_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", certificateAuthority},
				{"client_certificate", certificateAuthority},
				{"client_key", `"key"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
			},
		},
		{
			resource: "maria",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "memcached",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
			},
		},
		{
			resource: "memsql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mongo_host",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mongo_legacy_host",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mongo_legacy_replicaset",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"connect_to_replica", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"replica_set", `"replica_set"`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mongo_replica_set",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"connect_to_replica", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"replica_set", `"replica_set"`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mongo_sharded_cluster",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mtls_mysql",
			pairs: [][2]string{
				{"certificate_authority", certificateAuthority},
				{"client_certificate", certificateAuthority},
				{"client_key", `"key"`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"server_name", `"server_name"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mtls_postgres",
			pairs: [][2]string{
				{"certificate_authority", certificateAuthority},
				{"client_certificate", certificateAuthority},
				{"client_key", `"key"`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"server_name", `"server_name"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "mysql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "neptune",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"name"`},
				{"port", `443`},
			},
		},
		{
			resource: "neptune_iam",
			pairs: [][2]string{
				{"access_key", `"access_key"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"name"`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_arn", `"role_arn"`},
				{"role_external_id", `"role_external_id"`},
				{"secret_access_key", `"secret_access_key"`},
			},
		},
		{
			resource: "oracle",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "postgres",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "presto",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "rabbitmq_amqp_091",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "raw_tcp",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
			},
		},
		{
			resource: "rdp",
			pairs: [][2]string{
				{"downgrade_nla_connections", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"lock_required", `true`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "rdp_cert",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"lock_required", `true`},
				{"name", `"name"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "rds_postgres_iam",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_assumption_arn", `"role_assumption_arn"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "redis",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "redshift",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "single_store",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "snowflake",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"schema", `"schema"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "snowsight",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_username", `"healthcheck_username"`},
				{"name", `"name"`},
				{"saml_metadata", `"saml_metadata"`},
				{"subdomain", `"subdomain"`},
			},
		},
		{
			resource: "sql_server",
			pairs: [][2]string{
				{"allow_deprecated_encryption", `true`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"password", `"password"`},
				{"port", `443`},
				{"schema", `"schema"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "sql_server_azure_ad",
			pairs: [][2]string{
				{"allow_deprecated_encryption", `true`},
				{"client_id", `"client_id"`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"schema", `"schema"`},
				{"secret", `"secret"`},
				{"tenant_id", `"tenant_id"`},
			},
		},
		{
			resource: "sql_server_kerberos_ad",
			pairs: [][2]string{
				{"allow_deprecated_encryption", `true`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"keytab", `"keytab"`},
				{"krb_config", `"krb_config"`},
				{"name", `"name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"realm", `"realm"`},
				{"schema", `"schema"`},
				{"server_spn", `"server_spn"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "ssh",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"key_type", `"rsa-2048"`},
				{"name", `"name"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "ssh_cert",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"key_type", `"rsa-2048"`},
				{"name", `"name"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "ssh_customer_key",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"private_key", `"private_key"`},
				{"username", `"username"`},
			},
		},
		{
			resource: "ssh_password",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"username", `"username"`},
			},
		},
		{
			resource: "sybase",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "sybase_iq",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "teradata",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
		{
			resource: "trino",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"name"`},
				{"password", `"password"`},
				{"port", `443`},
				{"username", `"username"`},
			},
		},
	}

	resourceNameBase := randomWithPrefix("test")

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.resource, func(t *testing.T) {
			name := resourceNameBase + tc.resource
			cfg := testAccSDMResourceAnyConfig(name, tc.resource, tc.pairs)

			checks := make([]resource.TestCheckFunc, len(tc.pairs))
			for i, p := range tc.pairs {
				val := p[1]
				// TF removes quotes around strings
				val = strings.Trim(val, "\"")
				// ... and converts escaped new lines into real newlines
				val = strings.Replace(val, "\\n", "\n", -1)
				checks[i] = resource.TestCheckResourceAttr("sdm_resource."+name, tc.resource+".0."+p[0], val)
			}

			resource.Test(t, resource.TestCase{
				Providers:    testAccProviders,
				CheckDestroy: testCheckDestroy,
				Steps: []resource.TestStep{
					{
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						// there should be no change if the resource is updated from the server
						Taint:  []string{"sdm_resource." + name},
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						ResourceName: "sdm_resource." + name,
						ImportState:  true,
					},
				},
			})
		})
	}
}

func TestAccSDMResource_UpdateAllTypes_SecretStores(t *testing.T) {
	initAcceptanceTest(t)
	client, err := preTestClient()
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// NOTE: resources and secretstores must have matching CredentialType properties.
	resp, err := client.SecretStores().Create(ctx, &sdm.VaultTokenStore{
		Name:          "all-resources-test-store",
		ServerAddress: "allresourcestestaddr",
	})
	if err != nil {
		t.Fatalf("failed to create secret store: %v", err)
	}

	seID := resp.SecretStore.(*sdm.VaultTokenStore).ID

	respSSHCert, err := client.SecretStores().Create(ctx, &sdm.VaultTokenCertSSHStore{
		Name:                 "all-resources-test-cert-ssh-store",
		ServerAddress:        "allresourcestestaddr",
		SigningRole:          "allresourcessigningrole",
		SshMountPoint:        "allresourcessshmountpoint",
		IssuedCertTTLMinutes: 5,
	})
	if err != nil {
		t.Fatalf("failed to create cert ssh secret store: %v", err)
	}

	seSSHID := respSSHCert.SecretStore.(*sdm.VaultTokenCertSSHStore).ID

	respRDPCert, err := client.SecretStores().Create(ctx, &sdm.VaultTokenCertX509Store{
		Name:                 "all-resources-test-cert-rdp-store",
		ServerAddress:        "allresourcestestaddr",
		SigningRole:          "allresourcessigningrole",
		PkiMountPoint:        "allresourcespkimountpoint",
		IssuedCertTTLMinutes: 5,
	})
	if err != nil {
		t.Fatalf("failed to create cert pki secret store: %v", err)
	}

	seRDPID := respRDPCert.SecretStore.(*sdm.VaultTokenCertX509Store).ID

	type testCase struct {
		resource string
		pairs    [][2]string
	}
	tcs := []testCase{

		{
			resource: "aks",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"client_certificate", `"path/to/secret?key=key&encoding=base64"`},
				{"client_key", `"path/to/secret?key=key&encoding=base64"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "aks_basic_auth",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "aks_service_account",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"token", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "aks_service_account_user_impersonation",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"token", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "aks_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"client_certificate", `"path/to/secret?key=key&encoding=base64"`},
				{"client_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "amazon_eks",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"cluster_name", `"cluster_name"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "amazon_eks_instance_profile",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"cluster_name", `"cluster_name"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "amazon_eks_instance_profile_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"cluster_name", `"cluster_name"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "amazon_eks_user_impersonation",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"cluster_name", `"cluster_name"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "amazon_es",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "amazon_esiam",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
			},
		},

		{
			resource: "amazonmq_amqp_091",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "athena",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"name", `"secret_name"`},
				{"output", `"output"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "athena_iam",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"name", `"secret_name"`},
				{"output", `"output"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "aurora_mysql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "aurora_mysql_iam",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_assumption_arn", `"role_assumption_arn"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "aurora_postgres",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "aurora_postgres_iam",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_assumption_arn", `"role_assumption_arn"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "aws",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_region", `"healthcheck_region"`},
				{"name", `"secret_name"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "aws_console",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"enable_env_variables", `true`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"session_expiry", `20000`},
				{"subdomain", `"subdomain"`},
			},
		},

		{
			resource: "aws_console_static_key_pair",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"session_expiry", `20000`},
				{"subdomain", `"subdomain"`},
			},
		},

		{
			resource: "aws_instance_profile",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"enable_env_variables", `true`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "azure",
			pairs: [][2]string{
				{"app_id", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"tenant_id", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "azure_certificate",
			pairs: [][2]string{
				{"app_id", `"path/to/secret?key=key&encoding=base64"`},
				{"client_certificate", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"name", `"secret_name"`},
				{"secret_store_id", `"` + seID + `"`},
				{"tenant_id", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "azure_mysql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "azure_postgres",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "azure_postgres_managed_identity",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "big_query",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"secret_name"`},
				{"private_key", `"path/to/secret?key=key&encoding=base64"`},
				{"project", `"project"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"username"`},
			},
		},

		{
			resource: "cassandra",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "citus",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "click_house_http",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"url", `"https://app.strongdm.com"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "click_house_my_sql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "click_house_tcp",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "clustrix",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "cockroach",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "couchbase_database",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"n_1_ql_port", `443`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "couchbase_web_ui",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"subdomain", `"subdomain"`},
				{"url", `"https://app.strongdm.com"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "db_2_i",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "db_2_luw",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "document_db_host",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "document_db_replica_set",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"connect_to_replica", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"replica_set", `"replica_set"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "druid",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "dynamo_db",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "dynamo_dbiam",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"secret_name"`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "elastic",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "elasticache_redis",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "gcp",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"keyfile", `"path/to/secret?key=key&encoding=base64"`},
				{"name", `"secret_name"`},
				{"scopes", `"scopes"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "google_gke",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"secret_store_id", `"` + seID + `"`},
				{"service_account_key", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "google_gke_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"name", `"secret_name"`},
				{"secret_store_id", `"` + seID + `"`},
				{"service_account_key", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "greenplum",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "http_auth",
			pairs: [][2]string{
				{"auth_header", `"path/to/secret?key=key&encoding=base64"`},
				{"default_path", `"/"`},
				{"egress_filter", `"name:value"`},
				{"headers_blacklist", `"headers_blacklist"`},
				{"healthcheck_path", `"/"`},
				{"host_override", `"host_override"`},
				{"name", `"secret_name"`},
				{"secret_store_id", `"` + seID + `"`},
				{"subdomain", `"subdomain"`},
				{"url", `"https://app.strongdm.com"`},
			},
		},

		{
			resource: "http_basic_auth",
			pairs: [][2]string{
				{"default_path", `"/"`},
				{"egress_filter", `"name:value"`},
				{"headers_blacklist", `"headers_blacklist"`},
				{"healthcheck_path", `"/"`},
				{"host_override", `"host_override"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"subdomain", `"subdomain"`},
				{"url", `"https://app.strongdm.com"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "kubernetes",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"client_certificate", `"path/to/secret?key=key&encoding=base64"`},
				{"client_key", `"path/to/secret?key=key&encoding=base64"`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "kubernetes_basic_auth",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "kubernetes_service_account",
			pairs: [][2]string{
				{"allow_resource_role_bypass", `false`},
				{"discovery_enabled", `true`},
				{"discovery_username", `"discovery_username"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"token", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "kubernetes_service_account_user_impersonation",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"token", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "kubernetes_user_impersonation",
			pairs: [][2]string{
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"client_certificate", `"path/to/secret?key=key&encoding=base64"`},
				{"client_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"healthcheck_namespace", `"healthcheck_namespace"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "maria",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "memsql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mongo_host",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mongo_legacy_host",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mongo_legacy_replicaset",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"connect_to_replica", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"replica_set", `"replica_set"`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mongo_replica_set",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"connect_to_replica", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"replica_set", `"replica_set"`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mongo_sharded_cluster",
			pairs: [][2]string{
				{"auth_database", `"auth_database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mtls_mysql",
			pairs: [][2]string{
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"client_certificate", `"path/to/secret?key=key&encoding=base64"`},
				{"client_key", `"path/to/secret?key=key&encoding=base64"`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"server_name", `"server_name"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mtls_postgres",
			pairs: [][2]string{
				{"certificate_authority", `"path/to/secret?key=key&encoding=base64"`},
				{"client_certificate", `"path/to/secret?key=key&encoding=base64"`},
				{"client_key", `"path/to/secret?key=key&encoding=base64"`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"server_name", `"server_name"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "mysql",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "neptune_iam",
			pairs: [][2]string{
				{"access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"egress_filter", `"name:value"`},
				{"endpoint", `"endpoint"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_arn", `"path/to/secret?key=key&encoding=base64"`},
				{"role_external_id", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_access_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
			},
		},

		{
			resource: "oracle",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "postgres",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "presto",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"username"`},
			},
		},

		{
			resource: "rabbitmq_amqp_091",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "rdp",
			pairs: [][2]string{
				{"downgrade_nla_connections", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"lock_required", `true`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "rdp_cert",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"lock_required", `true`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"secret_store_id", `"` + seRDPID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "rds_postgres_iam",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"region", `"region"`},
				{"role_assumption_arn", `"role_assumption_arn"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "redis",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"tls_required", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "redshift",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "single_store",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"require_native_auth", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"use_azure_single_server_usernames", `true`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "snowflake",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"schema", `"schema"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "snowsight",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"healthcheck_username", `"healthcheck_username"`},
				{"name", `"secret_name"`},
				{"saml_metadata", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"subdomain", `"subdomain"`},
			},
		},

		{
			resource: "sql_server",
			pairs: [][2]string{
				{"allow_deprecated_encryption", `true`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"schema", `"schema"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "sql_server_azure_ad",
			pairs: [][2]string{
				{"allow_deprecated_encryption", `true`},
				{"client_id", `"path/to/secret?key=key&encoding=base64"`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"schema", `"schema"`},
				{"secret", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"tenant_id", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "sql_server_kerberos_ad",
			pairs: [][2]string{
				{"allow_deprecated_encryption", `true`},
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"keytab", `"path/to/secret?key=key&encoding=base64"`},
				{"krb_config", `"path/to/secret?key=key&encoding=base64"`},
				{"name", `"secret_name"`},
				{"override_database", `true`},
				{"port", `443`},
				{"realm", `"path/to/secret?key=key&encoding=base64"`},
				{"schema", `"schema"`},
				{"secret_store_id", `"` + seID + `"`},
				{"server_spn", `"path/to/secret?key=key&encoding=base64"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "ssh",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"key_type", `"rsa-2048"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "ssh_cert",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"identity_alias_healthcheck_username", `"identity_alias_healthcheck_username"`},
				{"key_type", `"rsa-2048"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"secret_store_id", `"` + seSSHID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "ssh_customer_key",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"private_key", `"path/to/secret?key=key&encoding=base64"`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "ssh_password",
			pairs: [][2]string{
				{"allow_deprecated_key_exchanges", `true`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"port_forwarding", `true`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "sybase",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "sybase_iq",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "teradata",
			pairs: [][2]string{
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"path/to/secret?key=key&encoding=base64"`},
			},
		},

		{
			resource: "trino",
			pairs: [][2]string{
				{"database", `"database"`},
				{"egress_filter", `"name:value"`},
				{"hostname", `"hostname"`},
				{"name", `"secret_name"`},
				{"password", `"path/to/secret?key=key&encoding=base64"`},
				{"port", `443`},
				{"secret_store_id", `"` + seID + `"`},
				{"username", `"username"`},
			},
		},
	}
	resourceNameBase := randomWithPrefix("test")

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.resource, func(t *testing.T) {
			name := resourceNameBase + tc.resource
			cfg := testAccSDMResourceAnyConfig(name, tc.resource, tc.pairs)

			checks := make([]resource.TestCheckFunc, len(tc.pairs))
			for i, p := range tc.pairs {
				val := p[1]
				// TF removes quotes around strings
				val = strings.Trim(val, "\"")
				// ... and converts escaped new lines into real newlines
				val = strings.Replace(val, "\\n", "\n", -1)
				checks[i] = resource.TestCheckResourceAttr("sdm_resource."+name, tc.resource+".0."+p[0], val)
			}

			resource.Test(t, resource.TestCase{
				Providers:    testAccProviders,
				CheckDestroy: testCheckDestroy,
				Steps: []resource.TestStep{
					{
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						// there should be no change if the resource is updated from the server
						Taint:  []string{"sdm_resource." + name},
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						ResourceName:      "sdm_resource." + name,
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		})
	}
}

// lintignore:AT012
func TestAccSDMSecretStore_UpdateAllTypes(t *testing.T) {
	initAcceptanceTest(t)
	type testCase struct {
		resource string
		pairs    [][2]string
	}
	tcs := []testCase{
		{
			resource: "active_directory_store",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "aws",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"region", `"region"`},
			},
		},
		{
			resource: "aws_cert_x509",
			pairs: [][2]string{
				{"ca_arn", `"ca_arn"`},
				{"certificate_template_arn", `"certificate_template_arn"`},
				{"issued_cert_ttl_minutes", `"5"`},
				{"name", `"all_secrets_name"`},
				{"region", `"region"`},
				{"signing_algo", `"signing_algo"`},
			},
		},
		{
			resource: "azure_store",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"vault_uri", `"vault_uri"`},
			},
		},
		{
			resource: "cyberark_conjur",
			pairs: [][2]string{
				{"app_url", `"app_url"`},
				{"name", `"all_secrets_name"`},
			},
		},
		{
			resource: "cyberark_pam",
			pairs: [][2]string{
				{"app_url", `"app_url"`},
				{"name", `"all_secrets_name"`},
			},
		},
		{
			resource: "cyberark_pam_experimental",
			pairs: [][2]string{
				{"app_url", `"app_url"`},
				{"name", `"all_secrets_name"`},
			},
		},
		{
			resource: "delinea_store",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"server_url", `"server_url"`},
				{"tenant_name", `"tenant_name"`},
			},
		},
		{
			resource: "gcp_store",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"project_id", `"project_id"`},
			},
		},
		{
			resource: "gcp_cert_x_509_store",
			pairs: [][2]string{
				{"ca_id", `"ca_id"`},
				{"ca_pool_id", `"ca_pool_id"`},
				{"issued_cert_ttl_minutes", `"5"`},
				{"location", `"location"`},
				{"name", `"all_secrets_name"`},
				{"project_id", `"project_id"`},
			},
		},
		{
			resource: "keyfactor_ssh_store",
			pairs: [][2]string{
				{"ca_file_path", `"ca_file_path"`},
				{"certificate_file_path", `"certificate_file_path"`},
				{"default_certificate_authority_name", `"default_certificate_authority_name"`},
				{"default_certificate_profile_name", `"default_certificate_profile_name"`},
				{"default_end_entity_profile_name", `"default_end_entity_profile_name"`},
				{"enrollment_code_env_var", `"enrollment_code_env_var"`},
				{"enrollment_username_env_var", `"enrollment_username_env_var"`},
				{"key_file_path", `"key_file_path"`},
				{"name", `"all_secrets_name"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "keyfactor_x_509_store",
			pairs: [][2]string{
				{"ca_file_path", `"ca_file_path"`},
				{"certificate_file_path", `"certificate_file_path"`},
				{"default_certificate_authority_name", `"default_certificate_authority_name"`},
				{"default_certificate_profile_name", `"default_certificate_profile_name"`},
				{"default_end_entity_profile_name", `"default_end_entity_profile_name"`},
				{"enrollment_code_env_var", `"enrollment_code_env_var"`},
				{"enrollment_username_env_var", `"enrollment_username_env_var"`},
				{"key_file_path", `"key_file_path"`},
				{"name", `"all_secrets_name"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "vault_approle",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "vault_approle_cert_ssh",
			pairs: [][2]string{
				{"issued_cert_ttl_minutes", `"5"`},
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
				{"signing_role", `"signing_role"`},
				{"ssh_mount_point", `"ssh_mount_point"`},
			},
		},
		{
			resource: "vault_approle_cert_x509",
			pairs: [][2]string{
				{"issued_cert_ttl_minutes", `"5"`},
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"pki_mount_point", `"pki_mount_point"`},
				{"server_address", `"server_address"`},
				{"signing_role", `"signing_role"`},
			},
		},
		{
			resource: "vault_aws_ec2",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "vault_aws_iam",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "vault_tls",
			pairs: [][2]string{
				{"ca_cert_path", `"ca_cert_path"`},
				{"client_cert_path", `"client_cert_path"`},
				{"client_key_path", `"client_key_path"`},
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "vault_tls_cert_ssh",
			pairs: [][2]string{
				{"ca_cert_path", `"ca_cert_path"`},
				{"client_cert_path", `"client_cert_path"`},
				{"client_key_path", `"client_key_path"`},
				{"issued_cert_ttl_minutes", `"5"`},
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
				{"signing_role", `"signing_role"`},
				{"ssh_mount_point", `"ssh_mount_point"`},
			},
		},
		{
			resource: "vault_tls_cert_x509",
			pairs: [][2]string{
				{"ca_cert_path", `"ca_cert_path"`},
				{"client_cert_path", `"client_cert_path"`},
				{"client_key_path", `"client_key_path"`},
				{"issued_cert_ttl_minutes", `"5"`},
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"pki_mount_point", `"pki_mount_point"`},
				{"server_address", `"server_address"`},
				{"signing_role", `"signing_role"`},
			},
		},
		{
			resource: "vault_token",
			pairs: [][2]string{
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
			},
		},
		{
			resource: "vault_token_cert_ssh",
			pairs: [][2]string{
				{"issued_cert_ttl_minutes", `"5"`},
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"server_address", `"server_address"`},
				{"signing_role", `"signing_role"`},
				{"ssh_mount_point", `"ssh_mount_point"`},
			},
		},
		{
			resource: "vault_token_cert_x509",
			pairs: [][2]string{
				{"issued_cert_ttl_minutes", `"5"`},
				{"name", `"all_secrets_name"`},
				{"namespace", `"namespace"`},
				{"pki_mount_point", `"pki_mount_point"`},
				{"server_address", `"server_address"`},
				{"signing_role", `"signing_role"`},
			},
		},
	}

	resourceNameBase := randomWithPrefix("test")

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.resource, func(t *testing.T) {
			name := resourceNameBase + tc.resource
			cfg := testAccSDMSecretStoreAnyConfig(name, tc.resource, tc.pairs)

			checks := make([]resource.TestCheckFunc, len(tc.pairs))
			for i, p := range tc.pairs {
				val := p[1]
				// TF removes quotes around strings
				val = strings.Trim(val, "\"")
				// ... and converts escaped new lines into real newlines
				val = strings.Replace(val, "\\n", "\n", -1)
				checks[i] = resource.TestCheckResourceAttr("sdm_secret_store."+name, tc.resource+".0."+p[0], val)
			}

			resource.Test(t, resource.TestCase{
				Providers:    testAccProviders,
				CheckDestroy: testCheckDestroy,
				Steps: []resource.TestStep{
					{
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						// there should be no change if the resource is updated from the server
						Taint:  []string{"sdm_secret_store." + name},
						Config: cfg,
						Check:  resource.ComposeTestCheckFunc(checks...),
					},
					{
						ResourceName: "sdm_secret_store." + name,
						ImportState:  true,
					},
				},
			})
		})
	}
}
