// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package sdm

import (
	"time"
)

// CreateResponseMetadata is reserved for future use.
type CreateResponseMetadata struct {
}

// GetResponseMetadata is reserved for future use.
type GetResponseMetadata struct {
}

// UpdateResponseMetadata is reserved for future use.
type UpdateResponseMetadata struct {
}

// DeleteResponseMetadata is reserved for future use.
type DeleteResponseMetadata struct {
}

// RateLimitMetadata contains information about remaining requests avaialable
// to the user over some timeframe.
type RateLimitMetadata struct {
	// How many total requests the user/token is authorized to make before being
	// rate limited.
	Limit int64
	// How many remaining requests out of the limit are still avaialable.
	Remaining int64
	// The time when remaining will be reset to limit.
	ResetAt time.Time
	// The bucket this user/token is associated with, which may be shared between
	// multiple users/tokens.
	Bucket string
}

// AccountAttachmentCreateOptions specifies extra options for creating an
// AccountAttachment.
type AccountAttachmentCreateOptions struct {
	// Overwrite clears all account grants before the attachment.
	Overwrite bool
}

// AccountAttachmentCreateResponse reports how the AccountAttachments were created in the system.
type AccountAttachmentCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created AccountAttachment.
	AccountAttachment *AccountAttachment
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountAttachmentGetResponse returns a requested AccountAttachment.
type AccountAttachmentGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested AccountAttachment.
	AccountAttachment *AccountAttachment
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountAttachmentDeleteResponse returns information about a AccountAttachment that was deleted.
type AccountAttachmentDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountAttachments assign an account to a role.
type AccountAttachment struct {
	// Unique identifier of the AccountAttachment.
	ID string
	// The id of the account of this AccountAttachment.
	AccountID string
	// The id of the attached role of this AccountAttachment.
	RoleID string
}

// AccountGrantCreateResponse reports how the AccountGrants were created in the system.
type AccountGrantCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created AccountGrant.
	AccountGrant *AccountGrant
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountGrantGetResponse returns a requested AccountGrant.
type AccountGrantGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested AccountGrant.
	AccountGrant *AccountGrant
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountGrantDeleteResponse returns information about a AccountGrant that was deleted.
type AccountGrantDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
type AccountGrant struct {
	// Unique identifier of the AccountGrant.
	ID string
	// The id of the composite role of this AccountGrant.
	ResourceID string
	// The id of the attached role of this AccountGrant.
	AccountID string
	// The timestamp when the resource will be granted. Optional. Both start_at
	// and end_at must be defined together, or not defined at all.
	StartFrom time.Time
	// The timestamp when the resource grant will expire. Optional. Both
	// start_at and end_at must be defined together, or not defined at all.
	ValidUntil time.Time
}

// AccountCreateResponse reports how the Accounts were created in the system.
type AccountCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Account.
	Account Account
	// The auth token generated for the Account. The Account will use this token to
	// authenticate with the strongDM API.
	Token string
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountGetResponse returns a requested Account.
type AccountGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Account.
	Account Account
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountUpdateResponse returns the fields of a Account after it has been updated by
// a AccountUpdateRequest.
type AccountUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Account.
	Account Account
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// AccountDeleteResponse returns information about a Account that was deleted.
type AccountDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// Accounts are users that have access to strongDM.
// There are two types of accounts:
// 1. **Regular users:** humans who are authenticated through username and password or SSO
// 2. **Service users:** machines that are authneticated using a service token
type Account interface {
	// GetID returns the unique identifier of the Account.
	GetID() string
	isOneOf_Account()
}

func (*User) isOneOf_Account() {}

// GetID returns the unique identifier of the User.
func (m *User) GetID() string     { return m.ID }
func (*Service) isOneOf_Account() {}

// GetID returns the unique identifier of the Service.
func (m *Service) GetID() string { return m.ID }

// A User can connect to resources they are granted directly, or granted
// via roles.
type User struct {
	// Unique identifier of the User.
	ID string
	// The User's email address. Must be unique.
	Email string
	// The User's first name.
	FirstName string
	// The User's last name.
	LastName string
	// The User's suspended state.
	Suspended bool
}

// A Service is a service account that can connect to resources they are granted
// directly, or granted via roles. Services are typically automated jobs.
type Service struct {
	// Unique identifier of the Service.
	ID string
	// Unique human-readable name of the Service.
	Name string
	// The Service's suspended state.
	Suspended bool
}

// A Resource is a server or service which clients connect to through relays.
type Resource interface {
	// GetID returns the unique identifier of the Resource.
	GetID() string
	isOneOf_Resource()
}

func (*Athena) isOneOf_Resource() {}

// GetID returns the unique identifier of the Athena.
func (m *Athena) GetID() string     { return m.ID }
func (*BigQuery) isOneOf_Resource() {}

// GetID returns the unique identifier of the BigQuery.
func (m *BigQuery) GetID() string    { return m.ID }
func (*Cassandra) isOneOf_Resource() {}

// GetID returns the unique identifier of the Cassandra.
func (m *Cassandra) GetID() string { return m.ID }
func (*Druid) isOneOf_Resource()   {}

// GetID returns the unique identifier of the Druid.
func (m *Druid) GetID() string      { return m.ID }
func (*DynamoDB) isOneOf_Resource() {}

// GetID returns the unique identifier of the DynamoDB.
func (m *DynamoDB) GetID() string   { return m.ID }
func (*AmazonES) isOneOf_Resource() {}

// GetID returns the unique identifier of the AmazonES.
func (m *AmazonES) GetID() string  { return m.ID }
func (*Elastic) isOneOf_Resource() {}

// GetID returns the unique identifier of the Elastic.
func (m *Elastic) GetID() string         { return m.ID }
func (*HTTPBasicAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the HTTPBasicAuth.
func (m *HTTPBasicAuth) GetID() string { return m.ID }
func (*HTTPNoAuth) isOneOf_Resource()  {}

// GetID returns the unique identifier of the HTTPNoAuth.
func (m *HTTPNoAuth) GetID() string { return m.ID }
func (*HTTPAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the HTTPAuth.
func (m *HTTPAuth) GetID() string     { return m.ID }
func (*Kubernetes) isOneOf_Resource() {}

// GetID returns the unique identifier of the Kubernetes.
func (m *Kubernetes) GetID() string            { return m.ID }
func (*KubernetesBasicAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the KubernetesBasicAuth.
func (m *KubernetesBasicAuth) GetID() string        { return m.ID }
func (*KubernetesServiceAccount) isOneOf_Resource() {}

// GetID returns the unique identifier of the KubernetesServiceAccount.
func (m *KubernetesServiceAccount) GetID() string { return m.ID }
func (*AmazonEKS) isOneOf_Resource()              {}

// GetID returns the unique identifier of the AmazonEKS.
func (m *AmazonEKS) GetID() string   { return m.ID }
func (*GoogleGKE) isOneOf_Resource() {}

// GetID returns the unique identifier of the GoogleGKE.
func (m *GoogleGKE) GetID() string { return m.ID }
func (*AKS) isOneOf_Resource()     {}

// GetID returns the unique identifier of the AKS.
func (m *AKS) GetID() string            { return m.ID }
func (*AKSBasicAuth) isOneOf_Resource() {}

// GetID returns the unique identifier of the AKSBasicAuth.
func (m *AKSBasicAuth) GetID() string        { return m.ID }
func (*AKSServiceAccount) isOneOf_Resource() {}

// GetID returns the unique identifier of the AKSServiceAccount.
func (m *AKSServiceAccount) GetID() string { return m.ID }
func (*Memcached) isOneOf_Resource()       {}

// GetID returns the unique identifier of the Memcached.
func (m *Memcached) GetID() string         { return m.ID }
func (*MongoLegacyHost) isOneOf_Resource() {}

// GetID returns the unique identifier of the MongoLegacyHost.
func (m *MongoLegacyHost) GetID() string         { return m.ID }
func (*MongoLegacyReplicaset) isOneOf_Resource() {}

// GetID returns the unique identifier of the MongoLegacyReplicaset.
func (m *MongoLegacyReplicaset) GetID() string { return m.ID }
func (*MongoHost) isOneOf_Resource()           {}

// GetID returns the unique identifier of the MongoHost.
func (m *MongoHost) GetID() string         { return m.ID }
func (*MongoReplicaSet) isOneOf_Resource() {}

// GetID returns the unique identifier of the MongoReplicaSet.
func (m *MongoReplicaSet) GetID() string { return m.ID }
func (*Mysql) isOneOf_Resource()         {}

// GetID returns the unique identifier of the Mysql.
func (m *Mysql) GetID() string         { return m.ID }
func (*AuroraMysql) isOneOf_Resource() {}

// GetID returns the unique identifier of the AuroraMysql.
func (m *AuroraMysql) GetID() string { return m.ID }
func (*Clustrix) isOneOf_Resource()  {}

// GetID returns the unique identifier of the Clustrix.
func (m *Clustrix) GetID() string { return m.ID }
func (*Maria) isOneOf_Resource()  {}

// GetID returns the unique identifier of the Maria.
func (m *Maria) GetID() string    { return m.ID }
func (*Memsql) isOneOf_Resource() {}

// GetID returns the unique identifier of the Memsql.
func (m *Memsql) GetID() string   { return m.ID }
func (*Oracle) isOneOf_Resource() {}

// GetID returns the unique identifier of the Oracle.
func (m *Oracle) GetID() string     { return m.ID }
func (*Postgres) isOneOf_Resource() {}

// GetID returns the unique identifier of the Postgres.
func (m *Postgres) GetID() string         { return m.ID }
func (*AuroraPostgres) isOneOf_Resource() {}

// GetID returns the unique identifier of the AuroraPostgres.
func (m *AuroraPostgres) GetID() string { return m.ID }
func (*Greenplum) isOneOf_Resource()    {}

// GetID returns the unique identifier of the Greenplum.
func (m *Greenplum) GetID() string   { return m.ID }
func (*Cockroach) isOneOf_Resource() {}

// GetID returns the unique identifier of the Cockroach.
func (m *Cockroach) GetID() string  { return m.ID }
func (*Redshift) isOneOf_Resource() {}

// GetID returns the unique identifier of the Redshift.
func (m *Redshift) GetID() string { return m.ID }
func (*Presto) isOneOf_Resource() {}

// GetID returns the unique identifier of the Presto.
func (m *Presto) GetID() string { return m.ID }
func (*RDP) isOneOf_Resource()  {}

// GetID returns the unique identifier of the RDP.
func (m *RDP) GetID() string     { return m.ID }
func (*Redis) isOneOf_Resource() {}

// GetID returns the unique identifier of the Redis.
func (m *Redis) GetID() string              { return m.ID }
func (*ElasticacheRedis) isOneOf_Resource() {}

// GetID returns the unique identifier of the ElasticacheRedis.
func (m *ElasticacheRedis) GetID() string { return m.ID }
func (*Snowflake) isOneOf_Resource()      {}

// GetID returns the unique identifier of the Snowflake.
func (m *Snowflake) GetID() string   { return m.ID }
func (*SQLServer) isOneOf_Resource() {}

// GetID returns the unique identifier of the SQLServer.
func (m *SQLServer) GetID() string { return m.ID }
func (*SSH) isOneOf_Resource()     {}

// GetID returns the unique identifier of the SSH.
func (m *SSH) GetID() string      { return m.ID }
func (*Sybase) isOneOf_Resource() {}

// GetID returns the unique identifier of the Sybase.
func (m *Sybase) GetID() string     { return m.ID }
func (*SybaseIQ) isOneOf_Resource() {}

// GetID returns the unique identifier of the SybaseIQ.
func (m *SybaseIQ) GetID() string   { return m.ID }
func (*Teradata) isOneOf_Resource() {}

// GetID returns the unique identifier of the Teradata.
func (m *Teradata) GetID() string { return m.ID }

type Athena struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	AccessKey string

	SecretAccessKey string

	Output string

	PortOverride int32

	Region string
}

type BigQuery struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	PrivateKey string

	Project string

	PortOverride int32

	Endpoint string

	Username string
}

type Cassandra struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	PortOverride int32

	Port int32

	TlsRequired bool
}

type Druid struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	PortOverride int32

	Username string

	Password string

	Port int32
}

type DynamoDB struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	AccessKey string

	SecretAccessKey string

	Region string

	Endpoint string

	PortOverride int32
}

type AmazonES struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Region string

	SecretAccessKey string

	Endpoint string

	AccessKey string

	PortOverride int32
}

type Elastic struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	PortOverride int32

	Port int32

	TlsRequired bool
}

type HTTPBasicAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Url string

	HealthcheckPath string

	Username string

	Password string

	HeadersBlacklist string

	DefaultPath string

	Subdomain string
}

type HTTPNoAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Url string

	HealthcheckPath string

	HeadersBlacklist string

	DefaultPath string

	Subdomain string
}

type HTTPAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Url string

	HealthcheckPath string

	AuthHeader string

	HeadersBlacklist string

	DefaultPath string

	Subdomain string
}

type Kubernetes struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	CertificateAuthority string

	CertificateAuthorityFilename string

	ClientCertificate string

	ClientCertificateFilename string

	ClientKey string

	ClientKeyFilename string
}

type KubernetesBasicAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	Username string

	Password string
}

type KubernetesServiceAccount struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	Token string
}

type AmazonEKS struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Endpoint string

	AccessKey string

	SecretAccessKey string

	CertificateAuthority string

	CertificateAuthorityFilename string

	Region string

	ClusterName string
}

type GoogleGKE struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Endpoint string

	CertificateAuthority string

	CertificateAuthorityFilename string

	ServiceAccountKey string

	ServiceAccountKeyFilename string
}

type AKS struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	CertificateAuthority string

	CertificateAuthorityFilename string

	ClientCertificate string

	ClientCertificateFilename string

	ClientKey string

	ClientKeyFilename string
}

type AKSBasicAuth struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	Username string

	Password string
}

type AKSServiceAccount struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Port int32

	Token string
}

type Memcached struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	PortOverride int32

	Port int32
}

type MongoLegacyHost struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	AuthDatabase string

	PortOverride int32

	Username string

	Password string

	Port int32

	ReplicaSet string

	TlsRequired bool
}

type MongoLegacyReplicaset struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	AuthDatabase string

	PortOverride int32

	Username string

	Password string

	Port int32

	ReplicaSet string

	ConnectToReplica bool

	TlsRequired bool
}

type MongoHost struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	AuthDatabase string

	PortOverride int32

	Username string

	Password string

	Port int32

	TlsRequired bool
}

type MongoReplicaSet struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	AuthDatabase string

	PortOverride int32

	Username string

	Password string

	Port int32

	ReplicaSet string

	ConnectToReplica bool

	TlsRequired bool
}

type Mysql struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type AuroraMysql struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Clustrix struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Maria struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Memsql struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32
}

type Oracle struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	Port int32

	PortOverride int32

	TlsRequired bool
}

type Postgres struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32

	OverrideDatabase bool
}

type AuroraPostgres struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32

	OverrideDatabase bool
}

type Greenplum struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32

	OverrideDatabase bool
}

type Cockroach struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32

	OverrideDatabase bool
}

type Redshift struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Port int32

	OverrideDatabase bool
}

type Presto struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Password string

	Database string

	PortOverride int32

	Port int32

	Username string

	TlsRequired bool
}

type RDP struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	PortOverride int32

	Port int32
}

type Redis struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	PortOverride int32

	Password string

	Port int32
}

type ElasticacheRedis struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	PortOverride int32

	Password string

	Port int32

	TlsRequired bool
}

type Snowflake struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	Schema string

	PortOverride int32
}

type SQLServer struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	Database string

	PortOverride int32

	Schema string

	Port int32

	OverrideDatabase bool
}

type SSH struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Port int32

	PublicKey string
}

type Sybase struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	PortOverride int32

	Port int32

	Password string
}

type SybaseIQ struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	PortOverride int32

	Port int32

	Password string
}

type Teradata struct {
	// Unique identifier of the Resource.
	ID string
	// Unique human-readable name of the Resource.
	Name string
	// True if the datasource is reachable and the credentials are valid.
	Healthy bool

	Hostname string

	Username string

	Password string

	PortOverride int32

	Port int32
}

// NodeCreateResponse reports how the Nodes were created in the system.
type NodeCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Node.
	Node Node
	// The auth token generated for the Node. The Node will use this token to
	// authenticate with the strongDM API.
	Token string
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// NodeGetResponse returns a requested Node.
type NodeGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Node.
	Node Node
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// NodeUpdateResponse returns the fields of a Node after it has been updated by
// a NodeUpdateRequest.
type NodeUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Node.
	Node Node
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// NodeDeleteResponse returns information about a Node that was deleted.
type NodeDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// Nodes make up the strongDM network, and allow your users to connect securely to your resources.
// There are two types of nodes:
// 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
// 1. **Gateways:** a relay that also listens for connections from strongDM clients
type Node interface {
	// GetID returns the unique identifier of the Node.
	GetID() string
	isOneOf_Node()
}

func (*Relay) isOneOf_Node() {}

// GetID returns the unique identifier of the Relay.
func (m *Relay) GetID() string { return m.ID }
func (*Gateway) isOneOf_Node() {}

// GetID returns the unique identifier of the Gateway.
func (m *Gateway) GetID() string { return m.ID }

// Relay represents a StrongDM CLI installation running in relay mode.
type Relay struct {
	// Unique identifier of the Relay.
	ID string
	// Unique human-readable name of the Relay. Generated if not provided on create.
	Name string
	// The current state of the relay. One of: "new", "verifying_restart",
	// "restarting", "started", "stopped", "dead", "unknown",
	State string
}

// Gateway represents a StrongDM CLI installation running in gateway mode.
type Gateway struct {
	// Unique identifier of the Gateway.
	ID string
	// Unique human-readable name of the Gateway. Generated if not provided on create.
	Name string
	// The current state of the gateway. One of: "new", "verifying_restart",
	// "restarting", "started", "stopped", "dead", "unknown"
	State string
	// The public hostname/port tuple at which the gateway will be accessible to clients.
	ListenAddress string
	// The hostname/port tuple which the gateway daemon will bind to.
	// If not provided on create, set to "0.0.0.0:<listen_address_port>".
	BindAddress string
}

// ResourceCreateResponse reports how the Resources were created in the system.
type ResourceCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Resource.
	Resource Resource
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// ResourceGetResponse returns a requested Resource.
type ResourceGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Resource.
	Resource Resource
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// ResourceUpdateResponse returns the fields of a Resource after it has been updated by
// a ResourceUpdateRequest.
type ResourceUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Resource.
	Resource Resource
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// ResourceDeleteResponse returns information about a Resource that was deleted.
type ResourceDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleAttachmentCreateResponse reports how the RoleAttachments were created in the system.
type RoleAttachmentCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created RoleAttachment.
	RoleAttachment *RoleAttachment
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleAttachmentGetResponse returns a requested RoleAttachment.
type RoleAttachmentGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested RoleAttachment.
	RoleAttachment *RoleAttachment
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleAttachmentDeleteResponse returns information about a RoleAttachment that was deleted.
type RoleAttachmentDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// A RoleAttachment assigns a role to a composite role.
type RoleAttachment struct {
	// Unique identifier of the RoleAttachment.
	ID string
	// The id of the composite role of this RoleAttachment.
	CompositeRoleID string
	// The id of the attached role of this RoleAttachment.
	AttachedRoleID string
}

// RoleGrantCreateResponse reports how the RoleGrants were created in the system.
type RoleGrantCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created RoleGrant.
	RoleGrant *RoleGrant
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleGrantGetResponse returns a requested RoleGrant.
type RoleGrantGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested RoleGrant.
	RoleGrant *RoleGrant
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleGrantDeleteResponse returns information about a RoleGrant that was deleted.
type RoleGrantDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// A RoleGrant connects a resource to a role, granting members of the role
// access to that resource.
type RoleGrant struct {
	// Unique identifier of the RoleGrant.
	ID string
	// The id of the resource of this RoleGrant.
	ResourceID string
	// The id of the attached role of this RoleGrant.
	RoleID string
}

// RoleCreateResponse reports how the Roles were created in the system. It can
// communicate partial successes or failures.
type RoleCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata
	// The created Role.
	Role *Role
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleGetResponse returns a requested Role.
type RoleGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata
	// The requested Role.
	Role *Role
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleUpdateResponse returns the fields of a Role after it has been updated by
// a RoleUpdateRequest.
type RoleUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata
	// The updated Role.
	Role *Role
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// RoleDeleteResponse returns information about a Role that was deleted.
type RoleDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata
	// Rate limit information.
	RateLimit *RateLimitMetadata
}

// A Role is a collection of permissions, and typically corresponds to a team, Active Directory OU, or other organizational unit. Users are granted access to resources by assigning them to roles.
type Role struct {
	// Unique identifier of the Role.
	ID string
	// Unique human-readable name of the Role.
	Name string
	// True if the Role is a composite role.
	Composite bool
}

// AccountAttachmentIterator provides read access to a list of AccountAttachment.
// Use it like so:
//     for iterator.Next() {
//         accountAttachment := iterator.Value()
//         // ...
//     }
type AccountAttachmentIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *AccountAttachment
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// AccountGrantIterator provides read access to a list of AccountGrant.
// Use it like so:
//     for iterator.Next() {
//         accountGrant := iterator.Value()
//         // ...
//     }
type AccountGrantIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *AccountGrant
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// AccountIterator provides read access to a list of Account.
// Use it like so:
//     for iterator.Next() {
//         account := iterator.Value()
//         // ...
//     }
type AccountIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() Account
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// NodeIterator provides read access to a list of Node.
// Use it like so:
//     for iterator.Next() {
//         node := iterator.Value()
//         // ...
//     }
type NodeIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() Node
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// ResourceIterator provides read access to a list of Resource.
// Use it like so:
//     for iterator.Next() {
//         resource := iterator.Value()
//         // ...
//     }
type ResourceIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() Resource
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// RoleAttachmentIterator provides read access to a list of RoleAttachment.
// Use it like so:
//     for iterator.Next() {
//         roleAttachment := iterator.Value()
//         // ...
//     }
type RoleAttachmentIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *RoleAttachment
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// RoleGrantIterator provides read access to a list of RoleGrant.
// Use it like so:
//     for iterator.Next() {
//         roleGrant := iterator.Value()
//         // ...
//     }
type RoleGrantIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *RoleGrant
	// Err returns the first error encountered during iteration, if any.
	Err() error
}

// RoleIterator provides read access to a list of Role.
// Use it like so:
//     for iterator.Next() {
//         role := iterator.Value()
//         // ...
//     }
type RoleIterator interface {
	// Next advances the iterator to the next item in the list. It returns
	// true if an item is available to retrieve via the `Value()` function.
	Next() bool
	// Value returns the current item, if one is available.
	Value() *Role
	// Err returns the first error encountered during iteration, if any.
	Err() error
}
