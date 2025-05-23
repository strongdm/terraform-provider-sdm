// Code generated by protogen. DO NOT EDIT.
package sdm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSDMManagedSecretValue_Create_Base64(t *testing.T) {
	name := randomWithPrefix("test")
	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Error(err)
	}
	pubKey := publicKeyPEM(privKey)
	pubKeyBase64 := base64.StdEncoding.EncodeToString([]byte(pubKey))

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		IsUnitTest:   true,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				ResourceName: name,
				Config: fmt.Sprintf(`resource "sdm_managed_secret_value" "%s" {
				value= {
					user_dn = "cn=john a. zoidberg,ou=people,dc=planetexpress,dc=com"
				}
				public_key = trimspace(<<-EOT
					%s
				EOT
				)
			}`, name, pubKeyBase64),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_managed_secret_value."+name, "public_key", strings.TrimSpace(pubKeyBase64)),
					validateEncryptedKeyValue(privKey, "sdm_managed_secret_value."+name, "user_dn", "cn=john a. zoidberg,ou=people,dc=planetexpress,dc=com"),
				),
			},
		},
	})
}

func TestAccSDMManagedSecretValue_Create_PEM(t *testing.T) {
	name := randomWithPrefix("test")
	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Error(err)
	}
	pubKey := publicKeyPEM(privKey)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		IsUnitTest:   true,
		CheckDestroy: testCheckDestroy,
		Steps: []resource.TestStep{
			{
				ResourceName: name,
				Config: fmt.Sprintf(`resource "sdm_managed_secret_value" "%s" {
				value= {
					user_dn = "cn=john a. zoidberg,ou=people,dc=planetexpress,dc=com"
				}
				public_key = trimspace(<<-EOT
					%s
				EOT
				)
			}`, name, pubKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdm_managed_secret_value."+name, "public_key", strings.TrimSpace(pubKey)),
					validateEncryptedKeyValue(privKey, "sdm_managed_secret_value."+name, "user_dn", "cn=john a. zoidberg,ou=people,dc=planetexpress,dc=com"),
				),
			},
		},
	})
}

func validateEncryptedKeyValue(privKey *rsa.PrivateKey, resource string, key, expected string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[resource]
		encrypted := r.Primary.Attributes["encrypted"]
		bytes, err := base64.StdEncoding.DecodeString(encrypted)
		if err != nil {
			return err
		}

		decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, bytes, nil)
		if err != nil {
			return err
		}
		var value map[string]string
		if err := json.Unmarshal(decryptedData, &value); err != nil {
			return err
		}
		if value[key] != expected {
			return fmt.Errorf("%s=%s is different than expected value: %s", key, value[key], expected)
		}
		return nil
	}
}

func publicKeyPEM(privKey *rsa.PrivateKey) string {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(&privKey.PublicKey)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	}))
}
