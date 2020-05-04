package sdm

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

var portOverride = NewAtomicCounter(32000)

func init() {
	resource.AddTestSweepers("sdm_resource", &resource.Sweeper{
		Name:         "sdm_resource",
		Dependencies: []string{"sdm_account_grant", "sdm_role_grant"},
		F: func(region string) error {
			client, err := preTestClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			listResp, err := client.Resources().List(ctx, "name:*-test-acc")
			if err != nil {
				return fmt.Errorf("Error listing resources: %s", err)
			}
			for listResp.Next() {
				v := listResp.Value()
				_, err := client.Resources().Delete(ctx, v.GetID())
				if err != nil {
					fmt.Printf("error deleting resource %s %s\n", v.GetID(), err)
				}
			}
			if listResp.Err() != nil {
				return fmt.Errorf("error after listing resource %s", listResp.Err())
			}
			return nil
		},
	})
}

func TestAccSDMResource_Create(t *testing.T) {
	name := randomWithPrefix("test")
	port := portOverride.Count()
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceRedisConfig(name, name, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.name", name),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttrSet("sdm_resource."+name, "redis.0.port_override"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.Redis).Name != name {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Resource.(*sdm.Redis).Name, name)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_Tags(t *testing.T) {
	name := randomWithPrefix("test")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceTagsConfig(name, name, sdm.Tags{"key": "value", "foo": "bar"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.tags.key", "value"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.tags.foo", "bar"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.GetTags()["key"] != "value" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Resource.GetTags()["key"])
						}
						if resp.Resource.GetTags()["foo"] != "bar" {
							return fmt.Errorf("unexpected value '%s' for tag 'key', expected 'value'", resp.Resource.GetTags()["key"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMResourceTagsConfig(name, name, sdm.Tags{"goat": "bananas"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sdm_resource."+name, "redis.0.tags.key"),
					resource.TestCheckNoResourceAttr("sdm_resource."+name, "redis.0.tags.foo"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "redis.0.tags.goat", "bananas"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.GetTags()["goat"] != "bananas" {
							return fmt.Errorf("unexpected value '%s' for tag 'goat', expected 'bananas'", resp.Resource.GetTags()["goat"])
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_CreateSSH(t *testing.T) {
	name := randomWithPrefix("test")
	port := portOverride.Count()
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceSSHConfig(name, name, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.name", name),
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.username", "user"),
					resource.TestCheckResourceAttr("sdm_resource."+name, "ssh.0.port", fmt.Sprint(port)),
					resource.TestCheckResourceAttrSet("sdm_resource."+name, "ssh.0.public_key"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", name)
						if err != nil {
							return err
						}

						// check if it was actually created
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.SSH).Name != name {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Resource.(*sdm.SSH).Name, name)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_Update(t *testing.T) {
	resourceName := randomWithPrefix("test")

	redisName := randomWithPrefix("test")
	updatedRedisName := randomWithPrefix("test2")

	port := portOverride.Count()
	updatedPort := portOverride.Count()

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSDMResourceRedisConfig(resourceName, redisName, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", redisName),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port", fmt.Sprint(port)),
					resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
				),
			},
			{
				ResourceName:      "sdm_resource." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSDMResourceRedisConfig(resourceName, updatedRedisName, updatedPort),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.name", updatedRedisName),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.hostname", "test.com"),
					resource.TestCheckResourceAttr("sdm_resource."+resourceName, "redis.0.port", fmt.Sprint(updatedPort)),
					resource.TestCheckResourceAttrSet("sdm_resource."+resourceName, "redis.0.port_override"),
					func(s *terraform.State) error {
						id, err := testCreatedID(s, "sdm_resource", resourceName)
						if err != nil {
							return err
						}

						// check if it was actually updated
						client := testClient()
						ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
						defer cancel()
						resp, err := client.Resources().Get(ctx, id)
						if err != nil {
							return fmt.Errorf("failed to get created resource: %w", err)
						}

						if resp.Resource.(*sdm.Redis).Name != updatedRedisName {
							return fmt.Errorf("unexpected name '%s', expected '%s'", resp.Resource.(*sdm.Redis).Name, updatedRedisName)
						}

						if resp.Resource.(*sdm.Redis).Port != updatedPort {
							return fmt.Errorf("unexpected port '%d', expected '%d'", resp.Resource.(*sdm.Redis).Port, updatedPort)
						}

						return nil
					},
				),
			},
			{
				ResourceName:      "sdm_resource." + resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSDMResource_UpdateAllTypes(t *testing.T) {
	t.Parallel()

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

	type testCase struct {
		resource string
		pairs    [][2]string
	}
	tcs := []testCase{
		{
			resource: "athena",
			pairs: [][2]string{
				{"name", "\"athena\""},
				{"access_key", "\"AccessKey\""},
				{"secret_access_key", "\"SecretAccessKey\""},
				{"output", "\"Output\""},
			},
		},
		{
			resource: "big_query",
			pairs: [][2]string{
				{"name", `"big_query"`},
				{"endpoint", `"Endpoint"`},
				{"private_key", `"PrivateKey"`},
				{"project", `"Project"`},
			},
		},
		{
			resource: "cassandra",
			pairs: [][2]string{
				{"name", `"cassandra"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
			},
		},
		{
			resource: "druid",
			pairs: [][2]string{
				{"name", `"druid"`},
				{"hostname", `"Hostname"`},
			},
		},
		{
			resource: "dynamo_db",
			pairs: [][2]string{
				{"name", `"dynamo_db"`},
				{"endpoint", `"Endpoint"`},
				{"access_key", `"AccessKey"`},
				{"secret_access_key", `"SecretAccessKey"`},
				{"region", `"Region"`},
			},
		},
		{
			resource: "amazon_es",
			pairs: [][2]string{
				{"name", `"amazon_es"`},
				{"secret_access_key", `"SecretAccessKey"`},
				{"region", `"Region"`},
			},
		},
		{
			resource: "elastic",
			pairs: [][2]string{
				{"name", `"elastic"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
			},
		},
		{
			resource: "http_basic_auth",
			pairs: [][2]string{
				{"name", `"http_basic"`},
				{"url", `"http://example.com"`},
				{"healthcheck_path", `"/"`},
			},
		},
		{
			resource: "http_no_auth",
			pairs: [][2]string{
				{"name", `"http_no_auth"`},
				{"url", `"http://example.com"`},
				{"healthcheck_path", `"/"`},
			},
		},
		{
			resource: "http_auth",
			pairs: [][2]string{
				{"name", `"http_auth"`},
				{"url", `"http://example.com"`},
				{"healthcheck_path", `"/"`},
				{"auth_header", `"AuthHeader"`},
			},
		},
		{
			resource: "kubernetes",
			pairs: [][2]string{
				{"name", `"kubernetes"`},
				{"hostname", `"Hostname"`},
				{"port", "443"},
			},
		},
		{
			resource: "kubernetes_basic_auth",
			pairs: [][2]string{
				{"name", `"kubernetes_basic_auth"`},
				{"hostname", `"Hostname"`},
				{"port", "443"},
				{"username", `"Username"`},
				{"password", `"Password"`},
			},
		},
		{
			resource: "amazon_eks",
			pairs: [][2]string{
				{"name", `"amazon_eks"`},
				{"endpoint", `"Endpoint"`},
				{"access_key", `"AccessKey"`},
				{"secret_access_key", `"SecretAccessKey"`},
				{"certificate_authority", certificateAuthority},
				{"region", `"Region"`},
				{"cluster_name", `"ClusterName"`},
			},
		},
		{
			resource: "google_gke",
			pairs: [][2]string{
				{"name", `"google_gke"`},
				{"endpoint", `"Endpoint"`},
				{"certificate_authority", certificateAuthority},
				{"service_account_key", `"{}"`},
			},
		},
		{
			resource: "memcached",
			pairs: [][2]string{
				{"name", `"memcached"`},
				{"hostname", `"Hostname"`},
			},
		},
		{
			resource: "mongo_legacy_host",
			pairs: [][2]string{
				{"name", `"mongo_legacy_host"`},
				{"hostname", `"Hostname"`},
				{"auth_database", `"AuthDatabase"`},
			},
		},
		{
			resource: "mongo_legacy_replicaset",
			pairs: [][2]string{
				{"name", `"mongo_legacy_replicaset"`},
				{"hostname", `"Hostname"`},
				{"auth_database", `"AuthDatabase"`},
				{"replica_set", `"ReplicaSet"`},
			},
		},
		{
			resource: "mongo_host",
			pairs: [][2]string{
				{"name", `"mongo_host"`},
				{"hostname", `"Hostname"`},
				{"auth_database", `"AuthDatabase"`},
			},
		},
		{
			resource: "mongo_replica_set",
			pairs: [][2]string{
				{"name", `"mongo_replica_set"`},
				{"hostname", `"Hostname"`},
				{"auth_database", `"AuthDatabase"`},
				{"replica_set", `"ReplicaSet"`},
			},
		},
		{
			resource: "mysql",
			pairs: [][2]string{
				{"name", `"mysql"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "aurora_mysql",
			pairs: [][2]string{
				{"name", `"aurora_mysql"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "clustrix",
			pairs: [][2]string{
				{"name", `"clustrix"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "maria",
			pairs: [][2]string{
				{"name", `"maria"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "memsql",
			pairs: [][2]string{
				{"name", `"memsql"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "oracle",
			pairs: [][2]string{
				{"name", `"oracle"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
				{"port", "1521"},
			},
		},
		{
			resource: "postgres",
			pairs: [][2]string{
				{"name", `"postgres"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "aurora_postgres",
			pairs: [][2]string{
				{"name", `"aurora-postgres"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "greenplum",
			pairs: [][2]string{
				{"name", `"greenplum"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "cockroach",
			pairs: [][2]string{
				{"name", `"cockroach"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "redshift",
			pairs: [][2]string{
				{"name", `"redshift"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "presto",
			pairs: [][2]string{
				{"name", `"presto"`},
				{"hostname", `"Hostname"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "rdp",
			pairs: [][2]string{
				{"name", `"rdp"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"port", "3389"},
			},
		},
		{
			resource: "redis",
			pairs: [][2]string{
				{"name", `"redis"`},
				{"hostname", `"Hostname"`},
			},
		},
		{
			resource: "elasticache_redis",
			pairs: [][2]string{
				{"name", `"elasticache_redis"`},
				{"hostname", `"Hostname"`},
			},
		},
		{
			resource: "snowflake",
			pairs: [][2]string{
				{"name", `"snowflake"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
				{"schema", `"Schema"`},
			},
		},
		{
			resource: "sql_server",
			pairs: [][2]string{
				{"name", `"sql_server"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
				{"database", `"Database"`},
			},
		},
		{
			resource: "ssh",
			pairs: [][2]string{
				{"name", `"ssh"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"port", "22"},
				{"port_forwarding", "true"},
			},
		},
		{
			resource: "sybase",
			pairs: [][2]string{
				{"name", `"sybase"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
			},
		},
		{
			resource: "sybase_iq",
			pairs: [][2]string{
				{"name", `"sybase_iq"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
			},
		},
		{
			resource: "teradata",
			pairs: [][2]string{
				{"name", `"teradata"`},
				{"hostname", `"Hostname"`},
				{"username", `"Username"`},
				{"password", `"Password"`},
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

			resource.ParallelTest(t, resource.TestCase{
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

func testAccSDMResourceAnyConfig(resourceName, resourceType string, pairs [][2]string) string {
	s := fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		%s {`, resourceName, resourceType)
	for _, p := range pairs {
		s += "\n"
		s += p[0] + " = " + p[1]
	}
	s += "\n"
	s += `}
	}`
	return s
}

func testAccSDMResourceRedisConfig(resourceName string, sdmResourceName string, port int32) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			port = %d
		}
	}
	`, resourceName, sdmResourceName, port)
}

func testAccSDMResourceTagsConfig(resourceName string, sdmResourceName string, tags sdm.Tags) string {
	tagString := ""
	for key, value := range tags {
		tagString += fmt.Sprintf("\t\t\t\t%s = \"%s\"\n", key, value)
	}
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		redis {
			name = "%s"
			hostname = "test.com"
			tags = {
%s
			}
		}
	}
	`, resourceName, sdmResourceName, tagString)
}

func testAccSDMResourceSSHConfig(resourceName string, sdmResourceName string, port int32) string {
	return fmt.Sprintf(`
	resource "sdm_resource" "%s" {
		ssh {
			name = "%s"
			hostname = "test.com"
			username = "user"
			port = %d
		}
	}
	`, resourceName, sdmResourceName, port)
}

func createRedisesWithPrefix(prefix string, count int) ([]sdm.Resource, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resources := []sdm.Resource{}
	for i := 0; i < count; i++ {
		port := portOverride.Count()
		createResp, err := client.Resources().Create(ctx, &sdm.Redis{
			Name:         randomWithPrefix(prefix),
			Hostname:     randomWithPrefix(prefix),
			Port:         port,
			PortOverride: port,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create redis: %w", err)
		}
		resources = append(resources, createResp.Resource)
	}
	return resources, nil
}

func createSSHesWithPrefix(prefix string, count int) ([]sdm.Resource, error) {
	client, err := preTestClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create test client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resources := []sdm.Resource{}
	for i := 0; i < count; i++ {
		port := portOverride.Count()
		createResp, err := client.Resources().Create(ctx, &sdm.SSH{
			Name:     randomWithPrefix(prefix),
			Hostname: randomWithPrefix(prefix),
			Username: randomWithPrefix(prefix),
			Port:     port,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create ssh: %w", err)
		}
		resources = append(resources, createResp.Resource)
	}
	return resources, nil
}
