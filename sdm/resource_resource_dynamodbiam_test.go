package sdm

import (
	"testing"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

// --- dynamoDBIAMApplySeValues (resource read path) ---

func TestApplySeValuesDirectArn(t *testing.T) {
	v := &sdm.DynamoDBIAM{RoleArn: "arn:aws:iam::123456789012:role/my-role"}
	se := map[string]string{}
	dynamoDBIAMApplySeValues(v, se)
	if se["role_arn"] != "arn:aws:iam::123456789012:role/my-role" {
		t.Errorf("role_arn: got %q, want direct ARN", se["role_arn"])
	}
	if se["role_assumption_arn"] != "" {
		t.Errorf("role_assumption_arn: got %q, want empty", se["role_assumption_arn"])
	}
}

func TestApplySeValuesSecretStorePath(t *testing.T) {
	v := &sdm.DynamoDBIAM{SecretStoreID: "store-id", RoleArn: "path/to/secret?key=role_arn"}
	se := map[string]string{}
	dynamoDBIAMApplySeValues(v, se)
	if se["role_assumption_arn"] != "path/to/secret?key=role_arn" {
		t.Errorf("role_assumption_arn: got %q, want path", se["role_assumption_arn"])
	}
	if se["role_arn"] != "" {
		t.Errorf("role_arn: got %q, want empty", se["role_arn"])
	}
}

func TestApplySeValuesEmptyRoleArn(t *testing.T) {
	for _, secretStoreID := range []string{"", "store-id"} {
		v := &sdm.DynamoDBIAM{SecretStoreID: secretStoreID, RoleArn: ""}
		se := map[string]string{}
		dynamoDBIAMApplySeValues(v, se)
		if se["role_arn"] != "" || se["role_assumption_arn"] != "" {
			t.Errorf("secretStoreID=%q: unexpected values in seValues: %v", secretStoreID, se)
		}
	}
}

func TestApplySeValuesRoleExternalID(t *testing.T) {
	v := &sdm.DynamoDBIAM{RoleExternalID: "ext-123"}
	se := map[string]string{}
	dynamoDBIAMApplySeValues(v, se)
	if se["role_external_id"] != "ext-123" {
		t.Errorf("role_external_id: got %q, want %q", se["role_external_id"], "ext-123")
	}
}

// --- dynamoDBIAMRoleArnFields (data source read path) ---

func TestRoleArnFieldsDirectArn(t *testing.T) {
	v := &sdm.DynamoDBIAM{RoleArn: "arn:aws:iam::123456789012:role/my-role"}
	roleArn, roleAssumptionArn := dynamoDBIAMRoleArnFields(v)
	if roleArn != "arn:aws:iam::123456789012:role/my-role" {
		t.Errorf("roleArn: got %q, want direct ARN", roleArn)
	}
	if roleAssumptionArn != "" {
		t.Errorf("roleAssumptionArn: got %q, want empty", roleAssumptionArn)
	}
}

func TestRoleArnFieldsSecretStorePath(t *testing.T) {
	v := &sdm.DynamoDBIAM{SecretStoreID: "store-id", RoleArn: "path/to/secret?key=role_arn"}
	roleArn, roleAssumptionArn := dynamoDBIAMRoleArnFields(v)
	if roleAssumptionArn != "path/to/secret?key=role_arn" {
		t.Errorf("roleAssumptionArn: got %q, want path", roleAssumptionArn)
	}
	if roleArn != "" {
		t.Errorf("roleArn: got %q, want empty", roleArn)
	}
}

func TestRoleArnFieldsEmptyRoleArn(t *testing.T) {
	for _, secretStoreID := range []string{"", "store-id"} {
		v := &sdm.DynamoDBIAM{SecretStoreID: secretStoreID, RoleArn: ""}
		roleArn, roleAssumptionArn := dynamoDBIAMRoleArnFields(v)
		if roleArn != "" || roleAssumptionArn != "" {
			t.Errorf("secretStoreID=%q: got (%q, %q), want both empty", secretStoreID, roleArn, roleAssumptionArn)
		}
	}
}

// TestDynamoDBIAMRoleArnDirectValue verifies that role_arn passes through as-is
// when no secret_store_id is set (no secret store path validation).
func TestDynamoDBIAMRoleArnDirectValue(t *testing.T) {
	r := resourceResource()
	d := r.TestResourceData()
	d.Set("dynamo_dbiam", []interface{}{
		map[string]interface{}{
			"name":             "test",
			"endpoint":         "dynamodb.us-east-1.amazonaws.com",
			"region":           "us-east-1",
			"role_arn":         "arn:aws:iam::123456789012:role/my-role",
			"role_assumption_arn":    "",
			"role_external_id": "",
			"secret_store_id":  "",
		},
	})

	seValues, err := secretStoreValuesForResource(d)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if seValues["role_arn"] != "arn:aws:iam::123456789012:role/my-role" {
		t.Errorf("role_arn: got %q, want %q", seValues["role_arn"], "arn:aws:iam::123456789012:role/my-role")
	}
	if seValues["role_assumption_arn"] != "" {
		t.Errorf("role_assumption_arn: got %q, want empty", seValues["role_assumption_arn"])
	}
}

// TestDynamoDBIAMRoleArnPathWithSecretStore verifies that role_assumption_arn is accepted
// as a valid secret store path when secret_store_id is set.
func TestDynamoDBIAMRoleArnPathWithSecretStore(t *testing.T) {
	r := resourceResource()
	d := r.TestResourceData()
	d.Set("dynamo_dbiam", []interface{}{
		map[string]interface{}{
			"name":             "test",
			"endpoint":         "dynamodb.us-east-1.amazonaws.com",
			"region":           "us-east-1",
			"role_arn":         "",
			"role_assumption_arn":    "path/to/secret?key=role_arn",
			"role_external_id": "",
			"secret_store_id":  "store-id-123",
		},
	})

	seValues, err := secretStoreValuesForResource(d)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if seValues["role_assumption_arn"] != "path/to/secret?key=role_arn" {
		t.Errorf("role_assumption_arn: got %q, want %q", seValues["role_assumption_arn"], "path/to/secret?key=role_arn")
	}
}

// TestDynamoDBIAMRoleArnPathInvalidPath verifies that an invalid secret store path
// for role_assumption_arn returns an error when secret_store_id is set.
func TestDynamoDBIAMRoleArnPathInvalidPath(t *testing.T) {
	r := resourceResource()
	d := r.TestResourceData()
	d.Set("dynamo_dbiam", []interface{}{
		map[string]interface{}{
			"name":             "test",
			"endpoint":         "dynamodb.us-east-1.amazonaws.com",
			"region":           "us-east-1",
			"role_arn":         "",
			"role_assumption_arn":    string([]byte{0x7f}), // control character makes it unparseable
			"role_external_id": "",
			"secret_store_id":  "store-id-123",
		},
	})

	_, err := secretStoreValuesForResource(d)
	if err == nil {
		t.Fatal("expected error for invalid role_assumption_arn, got nil")
	}
}

// TestDynamoDBIAMConvertToPlumbingRoleArnDirect verifies that convertResourceToPlumbing
// sets RoleArn from role_arn when no secret store is configured.
func TestDynamoDBIAMConvertToPlumbingRoleArnDirect(t *testing.T) {
	r := resourceResource()
	d := r.TestResourceData()
	d.Set("dynamo_dbiam", []interface{}{
		map[string]interface{}{
			"name":             "test",
			"endpoint":         "dynamodb.us-east-1.amazonaws.com",
			"region":           "us-east-1",
			"role_arn":         "arn:aws:iam::123456789012:role/my-role",
			"role_assumption_arn":    "",
			"role_external_id": "",
			"secret_store_id":  "",
		},
	})

	out := convertResourceToPlumbing(d)
	v, ok := out.(*sdm.DynamoDBIAM)
	if !ok {
		t.Fatalf("expected *sdm.DynamoDBIAM, got %T", out)
	}
	if v.RoleArn != "arn:aws:iam::123456789012:role/my-role" {
		t.Errorf("RoleArn: got %q, want %q", v.RoleArn, "arn:aws:iam::123456789012:role/my-role")
	}
}

// TestDynamoDBIAMConvertToPlumbingRoleArnPath verifies that convertResourceToPlumbing
// sets RoleArn from role_assumption_arn (not role_arn) when a secret store is configured.
func TestDynamoDBIAMConvertToPlumbingRoleArnPath(t *testing.T) {
	r := resourceResource()
	d := r.TestResourceData()
	d.Set("dynamo_dbiam", []interface{}{
		map[string]interface{}{
			"name":             "test",
			"endpoint":         "dynamodb.us-east-1.amazonaws.com",
			"region":           "us-east-1",
			"role_arn":         "",
			"role_assumption_arn":    "path/to/secret?key=role_arn",
			"role_external_id": "",
			"secret_store_id":  "store-id-123",
		},
	})

	out := convertResourceToPlumbing(d)
	v, ok := out.(*sdm.DynamoDBIAM)
	if !ok {
		t.Fatalf("expected *sdm.DynamoDBIAM, got %T", out)
	}
	if v.RoleArn != "path/to/secret?key=role_arn" {
		t.Errorf("RoleArn: got %q, want path from role_assumption_arn %q", v.RoleArn, "path/to/secret?key=role_arn")
	}
	if v.SecretStoreID != "store-id-123" {
		t.Errorf("SecretStoreID: got %q, want %q", v.SecretStoreID, "store-id-123")
	}
}

// TestDynamoDBIAMRoleArnNotTreatedAsSecretPath verifies that a real ARN value
// in role_arn is not rejected when no secret_store_id is set (the old bug: ARNs
// were incorrectly validated as secret store paths when secret_store_id was present).
func TestDynamoDBIAMRoleArnNotTreatedAsSecretPath(t *testing.T) {
	r := resourceResource()
	d := r.TestResourceData()
	// An ARN contains colons which could break the old URL validation when used
	// as a secret store path. Without secret_store_id, it must not be validated.
	d.Set("dynamo_dbiam", []interface{}{
		map[string]interface{}{
			"name":             "test",
			"endpoint":         "dynamodb.us-east-1.amazonaws.com",
			"region":           "us-east-1",
			"role_arn":         "arn:aws:iam::123456789012:role/my-role",
			"role_assumption_arn":    "",
			"role_external_id": "",
			"secret_store_id":  "",
		},
	})

	_, err := secretStoreValuesForResource(d)
	if err != nil {
		t.Errorf("role_arn with ARN value must not be validated as a secret path when no secret_store_id is set: %v", err)
	}
}
