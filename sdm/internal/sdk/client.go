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

// Package sdm implements an API client to strongDM restful API.
package sdm

// Code generated by protogen. DO NOT EDIT.

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	plumbing "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	defaultAPIHost   = "api.strongdm.com:443"
	apiVersion       = "2021-08-23"
	defaultUserAgent = "strongdm-sdk-go/4.2.0"
)

var _ = metadata.Pairs

type dialer func(ctx context.Context, addr string) (net.Conn, error)

// Client is the strongDM API client implementation.
type Client struct {
	apiHost               string
	apiToken              string
	apiSecret             []byte
	apiInsecureTransport  bool
	apiTLSConfig          *tls.Config
	exposeRateLimitErrors bool
	userAgent             string
	disableSigning        bool
	pageLimit             int
	snapshotAt            time.Time
	dialer                dialer

	grpcConn *grpc.ClientConn

	maxRetries                  int
	baseRetryDelay              time.Duration
	maxRetryDelay               time.Duration
	accountAttachments          *AccountAttachments
	accountAttachmentsHistory   *AccountAttachmentsHistory
	accountGrants               *AccountGrants
	accountGrantsHistory        *AccountGrantsHistory
	accountPermissions          *AccountPermissions
	accountResources            *AccountResources
	accountResourcesHistory     *AccountResourcesHistory
	accounts                    *Accounts
	accountsHistory             *AccountsHistory
	activities                  *Activities
	controlPanel                *ControlPanel
	nodes                       *Nodes
	nodesHistory                *NodesHistory
	organizationHistory         *OrganizationHistory
	peeringGroupNodes           *PeeringGroupNodes
	peeringGroupPeers           *PeeringGroupPeers
	peeringGroupResources       *PeeringGroupResources
	peeringGroups               *PeeringGroups
	queries                     *Queries
	remoteIdentities            *RemoteIdentities
	remoteIdentitiesHistory     *RemoteIdentitiesHistory
	remoteIdentityGroups        *RemoteIdentityGroups
	remoteIdentityGroupsHistory *RemoteIdentityGroupsHistory
	replays                     *Replays
	resources                   *Resources
	resourcesHistory            *ResourcesHistory
	roleResources               *RoleResources
	roleResourcesHistory        *RoleResourcesHistory
	roles                       *Roles
	rolesHistory                *RolesHistory
	secretStores                *SecretStores
	secretStoresHistory         *SecretStoresHistory
}

// New creates a new strongDM API client.
func New(token, secret string, opts ...ClientOption) (*Client, error) {
	token = strings.TrimSpace(token)
	secret = strings.TrimSpace(secret)
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, convertErrorToPorcelain(fmt.Errorf("invalid secret: %w", err))
	}

	client := &Client{
		apiHost:        defaultAPIHost,
		maxRetries:     defaultMaxRetries,
		baseRetryDelay: defaultBaseRetryDelay,
		maxRetryDelay:  defaultMaxRetryDelay,
		apiToken:       token,
		apiSecret:      decodedSecret,
		userAgent:      defaultUserAgent,
	}

	for _, opt := range opts {
		opt(client)
	}

	var dialOpts []grpc.DialOption
	if client.apiInsecureTransport {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	} else if client.apiTLSConfig != nil {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(client.apiTLSConfig)))
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: false,
		})))
	}
	if client.dialer != nil {
		dialOpts = append(dialOpts, grpc.WithContextDialer(client.dialer))
	}
	cc, err := grpc.Dial(client.apiHost, dialOpts...)
	if err != nil {
		return nil, convertErrorToPorcelain(fmt.Errorf("cannot dial API server: %w", err))
	}
	client.grpcConn = cc
	client.accountAttachments = &AccountAttachments{
		client: plumbing.NewAccountAttachmentsClient(client.grpcConn),
		parent: client,
	}
	client.accountAttachmentsHistory = &AccountAttachmentsHistory{
		client: plumbing.NewAccountAttachmentsHistoryClient(client.grpcConn),
		parent: client,
	}
	client.accountGrants = &AccountGrants{
		client: plumbing.NewAccountGrantsClient(client.grpcConn),
		parent: client,
	}
	client.accountGrantsHistory = &AccountGrantsHistory{
		client: plumbing.NewAccountGrantsHistoryClient(client.grpcConn),
		parent: client,
	}
	client.accountPermissions = &AccountPermissions{
		client: plumbing.NewAccountPermissionsClient(client.grpcConn),
		parent: client,
	}
	client.accountResources = &AccountResources{
		client: plumbing.NewAccountResourcesClient(client.grpcConn),
		parent: client,
	}
	client.accountResourcesHistory = &AccountResourcesHistory{
		client: plumbing.NewAccountResourcesHistoryClient(client.grpcConn),
		parent: client,
	}
	client.accounts = &Accounts{
		client: plumbing.NewAccountsClient(client.grpcConn),
		parent: client,
	}
	client.accountsHistory = &AccountsHistory{
		client: plumbing.NewAccountsHistoryClient(client.grpcConn),
		parent: client,
	}
	client.activities = &Activities{
		client: plumbing.NewActivitiesClient(client.grpcConn),
		parent: client,
	}
	client.controlPanel = &ControlPanel{
		client: plumbing.NewControlPanelClient(client.grpcConn),
		parent: client,
	}
	client.nodes = &Nodes{
		client: plumbing.NewNodesClient(client.grpcConn),
		parent: client,
	}
	client.nodesHistory = &NodesHistory{
		client: plumbing.NewNodesHistoryClient(client.grpcConn),
		parent: client,
	}
	client.organizationHistory = &OrganizationHistory{
		client: plumbing.NewOrganizationHistoryClient(client.grpcConn),
		parent: client,
	}
	client.peeringGroupNodes = &PeeringGroupNodes{
		client: plumbing.NewPeeringGroupNodesClient(client.grpcConn),
		parent: client,
	}
	client.peeringGroupPeers = &PeeringGroupPeers{
		client: plumbing.NewPeeringGroupPeersClient(client.grpcConn),
		parent: client,
	}
	client.peeringGroupResources = &PeeringGroupResources{
		client: plumbing.NewPeeringGroupResourcesClient(client.grpcConn),
		parent: client,
	}
	client.peeringGroups = &PeeringGroups{
		client: plumbing.NewPeeringGroupsClient(client.grpcConn),
		parent: client,
	}
	client.queries = &Queries{
		client: plumbing.NewQueriesClient(client.grpcConn),
		parent: client,
	}
	client.remoteIdentities = &RemoteIdentities{
		client: plumbing.NewRemoteIdentitiesClient(client.grpcConn),
		parent: client,
	}
	client.remoteIdentitiesHistory = &RemoteIdentitiesHistory{
		client: plumbing.NewRemoteIdentitiesHistoryClient(client.grpcConn),
		parent: client,
	}
	client.remoteIdentityGroups = &RemoteIdentityGroups{
		client: plumbing.NewRemoteIdentityGroupsClient(client.grpcConn),
		parent: client,
	}
	client.remoteIdentityGroupsHistory = &RemoteIdentityGroupsHistory{
		client: plumbing.NewRemoteIdentityGroupsHistoryClient(client.grpcConn),
		parent: client,
	}
	client.replays = &Replays{
		client: plumbing.NewReplaysClient(client.grpcConn),
		parent: client,
	}
	client.resources = &Resources{
		client: plumbing.NewResourcesClient(client.grpcConn),
		parent: client,
	}
	client.resourcesHistory = &ResourcesHistory{
		client: plumbing.NewResourcesHistoryClient(client.grpcConn),
		parent: client,
	}
	client.roleResources = &RoleResources{
		client: plumbing.NewRoleResourcesClient(client.grpcConn),
		parent: client,
	}
	client.roleResourcesHistory = &RoleResourcesHistory{
		client: plumbing.NewRoleResourcesHistoryClient(client.grpcConn),
		parent: client,
	}
	client.roles = &Roles{
		client: plumbing.NewRolesClient(client.grpcConn),
		parent: client,
	}
	client.rolesHistory = &RolesHistory{
		client: plumbing.NewRolesHistoryClient(client.grpcConn),
		parent: client,
	}
	client.secretStores = &SecretStores{
		client: plumbing.NewSecretStoresClient(client.grpcConn),
		parent: client,
	}
	client.secretStoresHistory = &SecretStoresHistory{
		client: plumbing.NewSecretStoresHistoryClient(client.grpcConn),
		parent: client,
	}
	return client, nil
}

// Close will close the internal GRPC connection to strongDM. If the client is
// not initialized will return an error. Attempting to use the client after
// Close() may cause panics.
func (c *Client) Close() error {
	if c == nil {
		return &UnknownError{Wrapped: fmt.Errorf("cannot close nil client")}
	}
	if c.grpcConn == nil {
		return &UnknownError{Wrapped: fmt.Errorf("cannot close nil grpc client")}
	}
	return c.grpcConn.Close()
}

// A ClientOption is an optional argument to New that can override the created
// client's default behavior.
type ClientOption func(c *Client)

// WithHost causes a Client to make it's calls against the provided host instead
// of against api.strongdm.com.
func WithHost(host string) ClientOption {
	return func(c *Client) {
		c.apiHost = host
	}
}

// WithInsecure enables a Client to talk to an http server instead of an https
// server. This is potentially useful when communicating through a proxy, but
// should be used with care.
func WithInsecure() ClientOption {
	return func(c *Client) {
		c.apiInsecureTransport = true
	}
}

// WithTLSConfig allows customization of the TLS configuration used to
// communicate with the API server.
func WithTLSConfig(cfg *tls.Config) ClientOption {
	return func(c *Client) {
		c.apiTLSConfig = cfg
	}
}

// WithUserAgentExtra modifies the user agent string to include additional identifying
// information for server-side analytics. The intended use is by extension libraries,
// like a terraform provider wrapping this client.
func WithUserAgentExtra(userAgentExtra string) ClientOption {
	return func(c *Client) {
		c.userAgent += " " + userAgentExtra
	}
}

// WithRateLimitRetries configures whether encountered rate limit errors should
// cause this client to sleep and retry (if enabled), or whether those errors should be
// exposed to the code using this client (if disabled). By default, it is enabled.
func WithRateLimitRetries(enabled bool) ClientOption {
	return func(c *Client) {
		c.exposeRateLimitErrors = !enabled
	}
}

// AccountAttachments assign an account to a role.
func (c *Client) AccountAttachments() *AccountAttachments {
	return c.accountAttachments
}

// AccountAttachmentsHistory records all changes to the state of an AccountAttachment.
func (c *Client) AccountAttachmentsHistory() *AccountAttachmentsHistory {
	return c.accountAttachmentsHistory
}

// AccountGrants assign a resource directly to an account, giving the account the permission to connect to that resource.
func (c *Client) AccountGrants() *AccountGrants {
	return c.accountGrants
}

// AccountGrantsHistory records all changes to the state of an AccountGrant.
func (c *Client) AccountGrantsHistory() *AccountGrantsHistory {
	return c.accountGrantsHistory
}

// AccountPermissions records the granular permissions accounts have, allowing them to execute
// relevant commands via StrongDM's APIs.
func (c *Client) AccountPermissions() *AccountPermissions {
	return c.accountPermissions
}

// AccountResources enumerates the resources to which accounts have access.
// The AccountResources service is read-only.
func (c *Client) AccountResources() *AccountResources {
	return c.accountResources
}

// AccountResourcesHistory records all changes to the state of a AccountResource.
func (c *Client) AccountResourcesHistory() *AccountResourcesHistory {
	return c.accountResourcesHistory
}

// Accounts are users that have access to strongDM. There are two types of accounts:
// 1. **Users:** humans who are authenticated through username and password or SSO.
// 2. **Service Accounts:** machines that are authenticated using a service token.
func (c *Client) Accounts() *Accounts {
	return c.accounts
}

// AccountsHistory records all changes to the state of an Account.
func (c *Client) AccountsHistory() *AccountsHistory {
	return c.accountsHistory
}

// An Activity is a record of an action taken against a strongDM deployment, e.g.
// a user creation, resource deletion, sso configuration change, etc. The Activities
// service is read-only.
func (c *Client) Activities() *Activities {
	return c.activities
}

// ControlPanel contains all administrative controls.
func (c *Client) ControlPanel() *ControlPanel {
	return c.controlPanel
}

// Nodes make up the strongDM network, and allow your users to connect securely to your resources. There are two types of nodes:
// - **Gateways** are the entry points into network. They listen for connection from the strongDM client, and provide access to databases and servers.
// - **Relays** are used to extend the strongDM network into segmented subnets. They provide access to databases and servers but do not listen for incoming connections.
func (c *Client) Nodes() *Nodes {
	return c.nodes
}

// NodesHistory records all changes to the state of a Node.
func (c *Client) NodesHistory() *NodesHistory {
	return c.nodesHistory
}

// OrganizationHistory records all changes to the state of an Organization.
func (c *Client) OrganizationHistory() *OrganizationHistory {
	return c.organizationHistory
}

// PeeringGroupNodes provides the building blocks necessary to obtain attach a node to a peering group.
func (c *Client) PeeringGroupNodes() *PeeringGroupNodes {
	return c.peeringGroupNodes
}

// PeeringGroupPeers provides the building blocks necessary to link two peering groups.
func (c *Client) PeeringGroupPeers() *PeeringGroupPeers {
	return c.peeringGroupPeers
}

// PeeringGroupResources provides the building blocks necessary to obtain attach a resource to a peering group.
func (c *Client) PeeringGroupResources() *PeeringGroupResources {
	return c.peeringGroupResources
}

// PeeringGroups provides the building blocks necessary to obtain explicit network topology and routing.
func (c *Client) PeeringGroups() *PeeringGroups {
	return c.peeringGroups
}

// A Query is a record of a single client request to a resource, such as a SQL query.
// Long-running SSH, RDP, or Kubernetes interactive sessions also count as queries.
// The Queries service is read-only.
func (c *Client) Queries() *Queries {
	return c.queries
}

// RemoteIdentities assign a resource directly to an account, giving the account the permission to connect to that resource.
func (c *Client) RemoteIdentities() *RemoteIdentities {
	return c.remoteIdentities
}

// RemoteIdentitiesHistory records all changes to the state of a RemoteIdentity.
func (c *Client) RemoteIdentitiesHistory() *RemoteIdentitiesHistory {
	return c.remoteIdentitiesHistory
}

// A RemoteIdentityGroup is a named grouping of Remote Identities for Accounts.
// An Account's relationship to a RemoteIdentityGroup is defined via RemoteIdentity objects.
func (c *Client) RemoteIdentityGroups() *RemoteIdentityGroups {
	return c.remoteIdentityGroups
}

// RemoteIdentityGroupsHistory records all changes to the state of a RemoteIdentityGroup.
func (c *Client) RemoteIdentityGroupsHistory() *RemoteIdentityGroupsHistory {
	return c.remoteIdentityGroupsHistory
}

// A Replay captures the data transferred over a long-running SSH, RDP, or Kubernetes interactive session
// (otherwise referred to as a query). The Replays service is read-only.
func (c *Client) Replays() *Replays {
	return c.replays
}

// Resources are databases, servers, clusters, websites, or clouds that strongDM
// delegates access to.
func (c *Client) Resources() *Resources {
	return c.resources
}

// ResourcesHistory records all changes to the state of a Resource.
func (c *Client) ResourcesHistory() *ResourcesHistory {
	return c.resourcesHistory
}

// RoleResources enumerates the resources to which roles have access.
// The RoleResources service is read-only.
func (c *Client) RoleResources() *RoleResources {
	return c.roleResources
}

// RoleResourcesHistory records all changes to the state of a RoleResource.
func (c *Client) RoleResourcesHistory() *RoleResourcesHistory {
	return c.roleResourcesHistory
}

// A Role has a list of access rules which determine which Resources the members
// of the Role have access to. An Account can be a member of multiple Roles via
// AccountAttachments.
func (c *Client) Roles() *Roles {
	return c.roles
}

// RolesHistory records all changes to the state of a Role.
func (c *Client) RolesHistory() *RolesHistory {
	return c.rolesHistory
}

// SecretStores are servers where resource secrets (passwords, keys) are stored.
func (c *Client) SecretStores() *SecretStores {
	return c.secretStores
}

// SecretStoresHistory records all changes to the state of a SecretStore.
func (c *Client) SecretStoresHistory() *SecretStoresHistory {
	return c.secretStoresHistory
}

type SnapshotClient struct {
	client *Client
}

// SnapshotAt constructs a read-only client that will provide historical data
// from the provided timestamp.
func (c *Client) SnapshotAt(t time.Time) *SnapshotClient {
	clientCopy := *c
	snapshotClient := &SnapshotClient{&clientCopy}
	snapshotClient.client.snapshotAt = t
	snapshotClient.client.accountAttachments = &AccountAttachments{
		client: plumbing.NewAccountAttachmentsClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.accountGrants = &AccountGrants{
		client: plumbing.NewAccountGrantsClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.accountPermissions = &AccountPermissions{
		client: plumbing.NewAccountPermissionsClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.accountResources = &AccountResources{
		client: plumbing.NewAccountResourcesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.accounts = &Accounts{
		client: plumbing.NewAccountsClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.nodes = &Nodes{
		client: plumbing.NewNodesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.peeringGroupNodes = &PeeringGroupNodes{
		client: plumbing.NewPeeringGroupNodesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.peeringGroupPeers = &PeeringGroupPeers{
		client: plumbing.NewPeeringGroupPeersClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.peeringGroupResources = &PeeringGroupResources{
		client: plumbing.NewPeeringGroupResourcesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.peeringGroups = &PeeringGroups{
		client: plumbing.NewPeeringGroupsClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.remoteIdentities = &RemoteIdentities{
		client: plumbing.NewRemoteIdentitiesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.remoteIdentityGroups = &RemoteIdentityGroups{
		client: plumbing.NewRemoteIdentityGroupsClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.resources = &Resources{
		client: plumbing.NewResourcesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.roleResources = &RoleResources{
		client: plumbing.NewRoleResourcesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.roles = &Roles{
		client: plumbing.NewRolesClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	snapshotClient.client.secretStores = &SecretStores{
		client: plumbing.NewSecretStoresClient(snapshotClient.client.grpcConn),
		parent: snapshotClient.client,
	}
	return snapshotClient
}

// AccountAttachments assign an account to a role.
func (c *SnapshotClient) AccountAttachments() SnapshotAccountAttachments {
	return c.client.accountAttachments
}

// AccountGrants assign a resource directly to an account, giving the account the permission to connect to that resource.
func (c *SnapshotClient) AccountGrants() SnapshotAccountGrants {
	return c.client.accountGrants
}

// AccountPermissions records the granular permissions accounts have, allowing them to execute
// relevant commands via StrongDM's APIs.
func (c *SnapshotClient) AccountPermissions() SnapshotAccountPermissions {
	return c.client.accountPermissions
}

// AccountResources enumerates the resources to which accounts have access.
// The AccountResources service is read-only.
func (c *SnapshotClient) AccountResources() SnapshotAccountResources {
	return c.client.accountResources
}

// Accounts are users that have access to strongDM. There are two types of accounts:
// 1. **Users:** humans who are authenticated through username and password or SSO.
// 2. **Service Accounts:** machines that are authenticated using a service token.
func (c *SnapshotClient) Accounts() SnapshotAccounts {
	return c.client.accounts
}

// Nodes make up the strongDM network, and allow your users to connect securely to your resources. There are two types of nodes:
// - **Gateways** are the entry points into network. They listen for connection from the strongDM client, and provide access to databases and servers.
// - **Relays** are used to extend the strongDM network into segmented subnets. They provide access to databases and servers but do not listen for incoming connections.
func (c *SnapshotClient) Nodes() SnapshotNodes {
	return c.client.nodes
}

// PeeringGroupNodes provides the building blocks necessary to obtain attach a node to a peering group.
func (c *SnapshotClient) PeeringGroupNodes() SnapshotPeeringGroupNodes {
	return c.client.peeringGroupNodes
}

// PeeringGroupPeers provides the building blocks necessary to link two peering groups.
func (c *SnapshotClient) PeeringGroupPeers() SnapshotPeeringGroupPeers {
	return c.client.peeringGroupPeers
}

// PeeringGroupResources provides the building blocks necessary to obtain attach a resource to a peering group.
func (c *SnapshotClient) PeeringGroupResources() SnapshotPeeringGroupResources {
	return c.client.peeringGroupResources
}

// PeeringGroups provides the building blocks necessary to obtain explicit network topology and routing.
func (c *SnapshotClient) PeeringGroups() SnapshotPeeringGroups {
	return c.client.peeringGroups
}

// RemoteIdentities assign a resource directly to an account, giving the account the permission to connect to that resource.
func (c *SnapshotClient) RemoteIdentities() SnapshotRemoteIdentities {
	return c.client.remoteIdentities
}

// A RemoteIdentityGroup is a named grouping of Remote Identities for Accounts.
// An Account's relationship to a RemoteIdentityGroup is defined via RemoteIdentity objects.
func (c *SnapshotClient) RemoteIdentityGroups() SnapshotRemoteIdentityGroups {
	return c.client.remoteIdentityGroups
}

// Resources are databases, servers, clusters, websites, or clouds that strongDM
// delegates access to.
func (c *SnapshotClient) Resources() SnapshotResources {
	return c.client.resources
}

// RoleResources enumerates the resources to which roles have access.
// The RoleResources service is read-only.
func (c *SnapshotClient) RoleResources() SnapshotRoleResources {
	return c.client.roleResources
}

// A Role has a list of access rules which determine which Resources the members
// of the Role have access to. An Account can be a member of multiple Roles via
// AccountAttachments.
func (c *SnapshotClient) Roles() SnapshotRoles {
	return c.client.roles
}

// SecretStores are servers where resource secrets (passwords, keys) are stored.
func (c *SnapshotClient) SecretStores() SnapshotSecretStores {
	return c.client.secretStores
}

// Sign returns the signature for the given byte array
func (c *Client) Sign(methodName string, message []byte) string {
	if c.disableSigning {
		return ""
	}
	// Current UTC date
	y, m, d := time.Now().UTC().Date()
	currentUTCDate := fmt.Sprintf("%04d-%02d-%02d", y, m, d)

	signingKey := hmacHelper(c.apiSecret, []byte(currentUTCDate))
	signingKey = hmacHelper(signingKey, []byte("sdm_api_v1"))

	hash := sha256.New()
	hash.Write([]byte(methodName))
	hash.Write([]byte{'\n'})
	hash.Write(message)
	hashedMessage := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hmacHelper(signingKey, hashedMessage))
}

func hmacHelper(key, msg []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}

func (c *Client) wrapContext(ctx context.Context, req proto.Message, methodName string) context.Context {
	msg, _ := proto.Marshal(req)
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"x-sdm-authentication": c.apiToken,
		"x-sdm-signature":      c.Sign(methodName, msg),
		"x-sdm-api-version":    apiVersion,
		"x-sdm-user-agent":     c.userAgent,
	}))
}

// These defaults are taken from AWS. Customization of these values
// is a future step in the API.
const (
	defaultMaxRetries     = 3
	defaultBaseRetryDelay = 30 * time.Millisecond
	defaultMaxRetryDelay  = 300 * time.Second
)

func (c *Client) jitterSleep(iter int) {
	durMax := c.baseRetryDelay * time.Duration(2<<iter)
	if durMax > c.maxRetryDelay {
		durMax = c.maxRetryDelay
	}
	// This is a full jitter, ranging from no delay to the maximum
	// this jittering aims to prevent clients that start and conflict
	// at the same time from retrying at the same intervals
	dur := rand.Intn(int(durMax))
	// TODO: use resource.Retry instead of time.Sleep in Terraform provider
	// see https://github.com/bflad/tfproviderlint/tree/main/passes/R018
	time.Sleep(time.Duration(dur)) //lintignore:R018
}

func (c *Client) shouldRetry(iter int, err error) bool {
	if iter >= c.maxRetries-1 {
		return false
	}
	// Internal and unknown errors should be retried
	// Other error types are likely not brief temporary errors
	if s, ok := status.FromError(err); ok {
		if !c.exposeRateLimitErrors && s.Code() == codes.ResourceExhausted {
			var rateLimitError *RateLimitError
			if err != nil && errors.As(convertErrorToPorcelain(err), &rateLimitError) {
				waitFor := time.Until(rateLimitError.RateLimit.ResetAt)
				// If timezones or clock drift causes this calculation to fail,
				// wait at most one minute.
				if waitFor < 0 || waitFor > time.Minute {
					waitFor = time.Minute
				}
				time.Sleep(waitFor)
			}
			return true
		}
		return s.Code() == codes.Internal || s.Code() == codes.Unavailable
	}
	return true
}
