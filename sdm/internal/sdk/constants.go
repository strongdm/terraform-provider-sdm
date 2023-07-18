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
// Code generated by constgen. DO NOT EDIT.

package sdm

// Permission Levels, shared by all entities capable of making authenticated requests against StrongDM.
const (
	PermissionLevelRootAdmin     = "root-admin"
	PermissionLevelAdmin         = "admin"
	PermissionLevelDatabaseAdmin = "database-admin"
	PermissionLevelTeamLeader    = "multi-team-leader"
	PermissionLevelUser          = "user"
	PermissionLevelRelay         = "relay"
	PermissionLevelAdminToken    = "admin-token"
	PermissionLevelSCIMToken     = "scim-token"
	PermissionLevelService       = "service"
	PermissionLevelSuspended     = "suspended"
	PermissionLevelEmpty         = ""
)

// Node Lifecycle States, defining whether a node was last reported online, offline, restarting, etc.
const (
	NodeStateNew              = "new"
	NodeStateVerifyingRestart = "verifying_restart"
	NodeStateAwaitingRestart  = "awaiting_restart"
	NodeStateRestarting       = "restarting"
	NodeStateStarted          = "started"
	NodeStateStopped          = "stopped"
	NodeStateDead             = "dead"
)

// Providers responsible for managing roles and users.
// None, or an empty string, implies the user is managed by strongDM.
// Deprecated: Please use SCIMProvider instead.
const (
	ProviderNone      = ""
	ProviderOkta      = "okta"
	ProviderSailPoint = "sailpoint"
	ProviderAzure     = "azure"
	ProviderGeneric   = "generic"
	ProviderOneLogin  = "onelogin"
	ProviderGoogle    = "google"
)

// Providers responsible for managing roles and users.
// None, or an empty string, implies the user is managed by strongDM.
const (
	SCIMProviderNone      = ""
	SCIMProviderOkta      = "okta"
	SCIMProviderSailPoint = "sailpoint"
	SCIMProviderAzure     = "azure"
	SCIMProviderGeneric   = "generic"
	SCIMProviderOneLogin  = "onelogin"
	SCIMProviderGoogle    = "google"
)

// Providers responsible for SSO authentication.
const (
	AuthProviderAzure           = "azure"
	AuthProviderBitium          = "bitium"
	AuthProviderGoogle          = "google"
	AuthProviderOkta            = "okta"
	AuthProviderStrongDM        = "strongdm"
	AuthProviderActiveDirectory = "active directory"
	AuthProviderGenericOIDC     = "generic oidc"
	AuthProviderOneLoginOIDC    = "oneloginv2"
	AuthProviderKeycloak        = "keycloak"
	AuthProviderShibboleth      = "shibboleth"
	AuthProviderAuth0           = "auth0"
	AuthProviderWorkspaceOne    = "workspace one"
	AuthProviderOneLoginSAML    = "onelogin-saml"
	AuthProviderGenericSAML     = "generic-saml"
	AuthProviderPingIDSAML      = "ping-identity-saml"
)

// Providers responsible for multi-factor authentication
const (
	MFAProviderNone = ""
	MFAProviderDuo  = "duo"
	MFAProviderTOTP = "totp"
)

// Activity Entities, all entity types that can be part of an activity.
const (
	ActivityEntityTypeUser                = "user"
	ActivityEntityTypeRole                = "role"
	ActivityEntityTypeLegacyCompositeRole = "composite_role"
	ActivityEntityTypeDatasource          = "datasource"
	ActivityEntityTypeOrganization        = "organization"
	ActivityEntityTypeInstallation        = "installation"
	ActivityEntityTypeSecretStore         = "secretstore"
	ActivityEntityTypeRemoteIdentityGroup = "remote_identity_group"
	ActivityEntityTypeRemoteIdentity      = "remote_identity"
	ActivityEntityTypeAccessRequest       = "access_request"
	ActivityEntityTypeWorkflow            = "workflow"
	ActivityEntityTypeNode                = "node"
	ActivityEntityTypePeeringGroup        = "peering_group"
)

// Activity Verbs, describe which kind of activity has taken place.
const (
	ActivityVerbUserAdded                                        = "user added"
	ActivityVerbUserDeleted                                      = "user deleted"
	ActivityVerbUserUpdated                                      = "user updated"
	ActivityVerbUserSignup                                       = "user signup"
	ActivityVerbUserTypeChanged                                  = "user type changed"
	ActivityVerbUserTemporaryAccessGranted                       = "user temporary access granted"
	ActivityVerbUserTemporaryAccessRevoked                       = "user temporary access revoked"
	ActivityVerbUserTemporaryAccessExpired                       = "user temporary access expired"
	ActivityVerbUserAddedToRole                                  = "user added to role"
	ActivityVerbUserDeletedFromRole                              = "user deleted from role"
	ActivityVerbUserSuspended                                    = "user suspended"
	ActivityVerbUserReinstated                                   = "user reinstated"
	ActivityVerbUserLoggedIntoTheUI                              = "user logged into the Admin UI"
	ActivityVerbParentAdminLoggedIntoChildOrg                    = "parent admin logged into the child org"
	ActivityVerbUserLoggedIntoTheClient                          = "user logged into the local client"
	ActivityVerbServiceAccountCreated                            = "service account created"
	ActivityVerbServiceAccountExpired                            = "service account expired"
	ActivityVerbAdminTokenAdded                                  = "admin token created"
	ActivityVerbAdminTokenDeleted                                = "admin token deleted"
	ActivityVerbAdminTokenExpired                                = "admin token expired"
	ActivityVerbAdminTokenRekeyed                                = "admin token rekeyed"
	ActivityVerbAdminTokenCloned                                 = "admin token cloned"
	ActivityVerbAdminTokenSuspended                              = "admin token suspended"
	ActivityVerbAdminTokenReinstated                             = "admin token reinstated"
	ActivityVerbSSOUserLoggedIntoTheUI                           = "user logged into the Admin UI using SSO"
	ActivityVerbSSOUserLoggedIntoTheClient                       = "user logged into the local client using SSO"
	ActivityVerbUserLoggedOutFromTheClient                       = "user logged out from the local client"
	ActivityVerbUserLoggedOutFromTheUI                           = "user logged out from the Admin UI"
	ActivityVerbFailedLoginFromTheUI                             = "failed login attempt from the Admin UI"
	ActivityVerbFailedLoginFromTheClient                         = "failed login attempt from the local client"
	ActivityVerbMFADeniedFromTheUI                               = "MFA denied access for the Admin UI"
	ActivityVerbMFADeniedFromTheClient                           = "MFA denied access for the local client"
	ActivityVerbTooManyAttemptsLockout                           = "user account locked due to failed login attempts"
	ActivityVerbAttemptCounterReset                              = "failed login attempt counter reset"
	ActivityVerbSuspendedUserLoginAttemptFromTheClient           = "attempt to login by a suspended user from the local client"
	ActivityVerbSuspendedUserLoginAttemptFromTheUI               = "attempt to login by a suspended user from the Admin UI"
	ActivityVerbServiceAccountLoginAttemptFromTheUI              = "attempted to login by a service account from the Admin UI"
	ActivityVerbSuspendedServiceAccountLoginAttemptFromTheUI     = "attempted to login by a suspended service account from the Admin UI"
	ActivityVerbSuspendedServiceAccountLoginAttemptFromTheClient = "attempt to login by a suspended service account from the local client"
	ActivityVerbUserSetAPassword                                 = "user set a password"
	ActivityVerbUserResetAPassword                               = "user reset their password"
	ActivityVerbUserChangedPassword                              = "user changed their password"
	ActivityVerbUserInvited                                      = "user invited"
	ActivityVerbUserClickedInvitation                            = "user clicked on their invitation"
	ActivityVerbUserClickedPasswordReset                         = "user clicked on their password reset"
	ActivityVerbUserAllowPasswordLogin                           = "user allowed to login via password"
	ActivityVerbUserRequireSSOLogin                              = "user required to login via SSO"
	ActivityVerbUserProvisioningEnabled                          = "user provisioning enabled"
	ActivityVerbUserProvisioningDisabled                         = "user provisioning disabled"
	ActivityVerbAdminInitiatedPasswordReset                      = "admin initiated password reset"
	ActivityVerbRoleAdded                                        = "role added"
	ActivityVerbRoleDeleted                                      = "role deleted"
	ActivityVerbRoleUpdated                                      = "role updated"
	ActivityVerbRoleAccessRulesUpdated                           = "access rules updated"
	ActivityVerbRoleAccessRulesCreated                           = "access rules created"
	ActivityVerbRoleAccessRulesDeleted                           = "access rules deleted"
	ActivityVerbRoleProvisioningEnabled                          = "role provisioning enabled"
	ActivityVerbRoleProvisioningDisabled                         = "role provisioning disabled"
	ActivityVerbDatasourceAdded                                  = "datasource added"
	ActivityVerbDatasourceCloned                                 = "datasource cloned"
	ActivityVerbDatasourceDeleted                                = "datasource deleted"
	ActivityVerbDatasourceUpdated                                = "datasource updated"
	ActivityVerbDatasourcePortOverride                           = "datasource connection port overriden"
	ActivityVerbMultipleDatasourcePortOverride                   = "multiple datasource ports overriden"
	ActivityVerbServerAdded                                      = "server added"
	ActivityVerbServerCloned                                     = "server cloned"
	ActivityVerbServerDeleted                                    = "server deleted"
	ActivityVerbServerUpdated                                    = "server updated"
	ActivityVerbServerPortOverride                               = "server connection port overriden"
	ActivityVerbMultipleServerPortOverride                       = "multiple server ports overriden"
	ActivityVerbClusterAdded                                     = "cluster added"
	ActivityVerbClusterCloned                                    = "cluster cloned"
	ActivityVerbClusterDeleted                                   = "cluster deleted"
	ActivityVerbClusterUpdated                                   = "cluster updated"
	ActivityVerbClusterPortOverride                              = "cluster connection port overriden"
	ActivityVerbMultipleClusterPortOverride                      = "multiple cluster ports overriden"
	ActivityVerbCloudAdded                                       = "cloud added"
	ActivityVerbCloudCloned                                      = "cloud cloned"
	ActivityVerbCloudDeleted                                     = "cloud deleted"
	ActivityVerbCloudUpdated                                     = "cloud updated"
	ActivityVerbWebsiteAdded                                     = "website added"
	ActivityVerbWebsiteCloned                                    = "website cloned"
	ActivityVerbWebsiteDeleted                                   = "website deleted"
	ActivityVerbWebsiteUpdated                                   = "website updated"
	ActivityVerbInstallationCreated                              = "installation created"
	ActivityVerbRelayInstallationCreated                         = "installation created for relay"
	ActivityVerbInstallationApproved                             = "installation approved"
	ActivityVerbInstallationRevoked                              = "installation revoked"
	ActivityVerbRelayCreated                                     = "relay created"
	ActivityVerbRelayUpdatedName                                 = "relay name updated"
	ActivityVerbRelayDeleted                                     = "relay deleted"
	ActivityVerbOrgPublicKeyUpdated                              = "public key updated"
	ActivityVerbOrgEnforcePortOverridesUpdated                   = "port override enforcement updated"
	ActivityVerbOrgServiceAutoConnectUpdated                     = "service account auto-connect updated"
	ActivityVerbOrgSelfRegistrationActivated                     = "self-registration activated"
	ActivityVerbOrgSelfRegistrationDeactivated                   = "self-registration deactivated"
	ActivityVerbOrgNameUpdated                                   = "organization name updated"
	ActivityVerbOrgSettingUpdated                                = "organization setting updated"
	ActivityVerbOrgLogSyncSettingUpdated                         = "organization log stream setting updated"
	ActivityVerbOrgCreated                                       = "organization created"
	ActivityVerbOrgSCIMProvisioningUpdated                       = "SCIM provider set"
	ActivityVerbOrgSCIMProvisioningDeleted                       = "SCIM provider deleted"
	ActivityVerbOrgCustomProvisioningUpdated                     = "Provisioning provider set"
	ActivityVerbOrgCustomProvisioningDeleted                     = "Provisioning provider deleted"
	ActivityVerbChildOrgAdminInvited                             = "child organization admin invited"
	ActivityVerbServiceAccountRekeyed                            = "service account rekeyed"
	ActivityVerbSCIMTokenAdded                                   = "SCIM token created"
	ActivityVerbSCIMTokenDeleted                                 = "SCIM token deleted"
	ActivityVerbSCIMTokenRekeyed                                 = "SCIM token rekeyed"
	ActivityVerbAPIKeyDeleted                                    = "API key deleted"
	ActivityVerbOrgSSHCertificateAuthorityRotated                = "organization SSH certificate authority rotated"
	ActivityVerbOrgSSHAllowPortForwarding                        = "allowed SSH port forwarding"
	ActivityVerbOrgSSHDisallowPortForwarding                     = "disallowed SSH port forwarding"
	ActivityVerbOrgAddChild                                      = "add child organization"
	ActivityVerbOrgRemoveChild                                   = "remove child organization"
	ActivityVerbOrgExtendTrial                                   = "trial extended"
	ActivityVerbSecretStoreAdded                                 = "secret store added"
	ActivityVerbSecretStoreUpdated                               = "secret store updated"
	ActivityVerbSecretStoreDeleted                               = "secret store deleted"
	ActivityVerbRemoteIdentityGroupCreated                       = "remote identity group created"
	ActivityVerbRemoteIdentityGroupUpdated                       = "remote identity group updated"
	ActivityVerbRemoteIdentityGroupDeleted                       = "remote identity group deleted"
	ActivityVerbRemoteIdentityCreated                            = "remote identity created"
	ActivityVerbRemoteIdentityUpdated                            = "remote identity updated"
	ActivityVerbRemoteIdentityDeleted                            = "remote identity deleted"
	ActivityVerbAccessRequestedToResource                        = "access requested to resource"
	ActivityVerbAccessRequestToResourceApprovalAdded             = "access request to resource approval added"
	ActivityVerbAccessRequestToResourceCanceled                  = "access request to resource canceled"
	ActivityVerbAccessRequestToResourceDenied                    = "access request to resource denied"
	ActivityVerbAccessRequestToResourceTimedOut                  = "access request to resource timed out"
	ActivityVerbAccessRequestToResourceGranted                   = "access request to resource granted"
	ActivityVerbAccessRequestToResourceGrantedAutomatically      = "access request to resource granted automatically"
	ActivityVerbOrgVNMSubnetUpdated                              = "organization VNM subnet updated"
	ActivityVerbOrgVNMResourcesAllocated                         = "organization resources allocated within VNM subnet"
	ActivityVerbDeprecatedOrgActivateDeviceApproval              = "activate device approval"
	ActivityVerbDeprecatedOrgDeactivateDeviceApproval            = "deactivate device approval"
	ActivityVerbEmulationMigrationCompleted                      = "emulation migration completed"
	ActivityVerbAccessOverhaulMigrationCompleted                 = "access overhaul migration completed"
	ActivityVerbActivatedSSHPortForwardingAllServer              = "enabled SSH port forwarding on all servers"
	ActivityVerbTOTPEnrollmentAdded                              = "user enrolled a totp device"
	ActivityVerbTOTPEnrollmentDeleted                            = "user reset their totp enrollment"
	ActivityVerbSuspendedUserEnrollAttemptFromTheUI              = "attempt to enroll by a suspended user from the Admin UI"
	ActivityVerbResourceLocked                                   = "user locked a resource"
	ActivityVerbResourceUnlocked                                 = "user unlocked a resource"
	ActivityVerbResourceForceUnlocked                            = "admin force-unlocked a resource"
)

// Permissions, all permissions that may be granted to an account.
const (
	PermissionRelayList                              = "relay:list"
	PermissionRelayCreate                            = "relay:create"
	PermissionDatasourceList                         = "datasource:list"
	PermissionDatasourceCreate                       = "datasource:create"
	PermissionDatasourceHealthcheck                  = "datasource:healthcheck"
	PermissionDeprecatedDatasourceGrant              = "datasource:grant"
	PermissionDatasourceDelete                       = "datasource:delete"
	PermissionDatasourceUpdate                       = "datasource:update"
	PermissionResourceLockDelete                     = "resourcelock:delete"
	PermissionResourceLockList                       = "resourcelock:list"
	PermissionSecretStoreCreate                      = "secretstore:create"
	PermissionSecretStoreList                        = "secretstore:list"
	PermissionSecretStoreDelete                      = "secretstore:delete"
	PermissionSecretStoreUpdate                      = "secretstore:update"
	PermissionSecretStoreStatus                      = "secretstore:status"
	PermissionRemoteIdentityGroupWrite               = "remoteidentitygroup:write"
	PermissionRemoteIdentityGroupRead                = "remoteidentitygroup:read"
	PermissionRemoteIdentityWrite                    = "remoteidentity:write"
	PermissionRemoteIdentityRead                     = "remoteidentity:read"
	PermissionUserCreate                             = "user:create"
	PermissionUserList                               = "user:list"
	PermissionUserUpdateAdmin                        = "user:update_admin"
	PermissionUserCreateAdminToken                   = "user:create_admin_token"
	PermissionUserCreateServiceAccount               = "user:create_service_account"
	PermissionUserSetPermissionLevel                 = "user:set_strong_role"
	PermissionUserUpdate                             = "user:update"
	PermissionUserInitiatePasswordReset              = "user:initiate_password_reset"
	PermissionUserDelete                             = "user:delete"
	PermissionUserAssign                             = "user:assign"
	PermissionUserSuspend                            = "user:suspend"
	PermissionDemoProvisioningRequestCreate          = "demoprovisioningrequest:create"
	PermissionDemoProvisioningRequestList            = "demoprovisioningrequest:list"
	PermissionRoleList                               = "role:list"
	PermissionRoleCreate                             = "role:create"
	PermissionRoleDelete                             = "role:delete"
	PermissionRoleUpdate                             = "role:update"
	PermissionOrgViewSettings                        = "organization:view_settings"
	PermissionOrgEditSettings                        = "organization:edit_settings"
	PermissionOrgDeploymentDoctor                    = "organization:deployment_doctor"
	PermissionOrgListChildren                        = "organization:list_children"
	PermissionOrgCreateChildOrganization             = "organization:create_child_organization"
	PermissionOrgAuditUsers                          = "audit:users"
	PermissionOrgAuditRoles                          = "audit:roles"
	PermissionOrgAuditDatasources                    = "audit:datasources"
	PermissionOrgAuditNodes                          = "audit:nodes"
	PermissionOrgAuditPermissions                    = "audit:permissions"
	PermissionOrgAuditQueries                        = "audit:queries"
	PermissionOrgAuditActivities                     = "audit:activities"
	PermissionOrgAuditSSH                            = "audit:ssh"
	PermissionOrgAuditAccountGrants                  = "audit:accountgrants"
	PermissionOrgAuditOrg                            = "audit:organization"
	PermissionOrgAuditRemoteIdentities               = "audit:remoteidentities"
	PermissionOrgAuditRemoteIdentityGroups           = "audit:remoteidentitygroups"
	PermissionOrgAuditSecretStores                   = "audit:secretstores"
	PermissionWorkflowList                           = "workflow:list"
	PermissionWorkflowEdit                           = "workflow:edit"
	PermissionAccessRequestEdit                      = "accessrequest:edit"
	PermissionAccessRequestList                      = "accessrequest:list"
	PermissionAccessRequestRequester                 = "accessrequest:requester"
	PermissionGlobalRDPRender                        = "rdp:render"
	PermissionGlobalQueryBucketTracker               = "query:bucket_tracker"
	PermissionGlobalAssetsGetLatestVersionCommitHash = "assets:get_latest_version_commit_hash"
	PermissionGlobalSDMOSService                     = "sdmos:service"
	PermissionGlobalSDMOSDeployment                  = "sdmos:deployment"
	PermissionGlobalSDMOSRelease                     = "sdmos:release"
	PermissionGlobalDemoProvisioner                  = "demo:provision"
	PermissionInstallationBless                      = "installation:bless"
	PermissionInstallationCreate                     = "installation:create"
	PermissionInstallationRevoke                     = "installation:revoke"
	PermissionTestingOrgCreate                       = "testing:organization:create"
	PermissionTestingOrgDelete                       = "testing:organization:delete"
	PermissionTestingNoPermissions                   = "testing:noperms"
	PermissionTestingFetchQueries                    = "testing:queries:get"
	PermissionGrantRead                              = "grant:read"
	PermissionGrantWrite                             = "grant:write"
	PermissionReportRead                             = "report:read"
	PermissionBillingRead                            = "billing:read"
)

// Query Categories, all the categories of resource against which queries are logged.
const (
	QueryCategoryKubernetes  = "k8s"
	QueryCategoryDatasources = "queries"
	QueryCategoryRDP         = "rdp"
	QueryCategorySSH         = "ssh"
	QueryCategoryWeb         = "web"
	QueryCategoryCloud       = "cloud"
	QueryCategoryAll         = "all"
)

// LogRemoteEncoder defines the encryption encoder for the queries are stored in the API.
const (
	LogRemoteEncoderStrongDM = "strongdm"
	LogRemoteEncoderPubKey   = "pubkey"
	LogRemoteEncoderHash     = "hash"
)

// LogLocalStorage defines how queries are stored locally.
const (
	LogLocalStorageStdout = "stdout"
	LogLocalStorageFile   = "file"
	LogLocalStorageTCP    = "tcp"
	LogLocalStorageSocket = "socket"
	LogLocalStorageSyslog = "syslog"
	LogLocalStorageNone   = "none"
)

// LogLocalEncoder defines the encryption encoder for queries are stored locally in the relay.
const (
	LogLocalEncoderPlaintext = "plaintext"
	LogLocalEncoderPubKey    = "pubkey"
)

// LogLocalFormat defines the format the queries are stored locally in the relay.
const (
	LogLocalFormatCSV  = "csv"
	LogLocalFormatJSON = "json"
)

// OrgKind defines the types of organizations that may exist.
const (
	OrgKindSolo  = "solo"
	OrgKindRoot  = "root"
	OrgKindChild = "child"
)

// KeyType defines the supported SSH key types
const (
	SSHKeyTypeRSA_2048  = "rsa-2048"
	SSHKeyTypeRSA_4096  = "rsa-4096"
	SSHKeyTypeECDSA_256 = "ecdsa-256"
	SSHKeyTypeECDSA_384 = "ecdsa-384"
	SSHKeyTypeECDSA_521 = "ecdsa-521"
	SSHKeyTypeED25519   = "ed25519"
)

// CaptureType designates what type of SSH/RDP/K8s capture we have.
const (
	CaptureTypeShell          = "shell"
	CaptureTypeScpUpload      = "scp-upload"
	CaptureTypeScpDownload    = "scp-download"
	CaptureTypeCommand        = "command"
	CaptureTypeRDPBasic       = "rdp-basic"
	CaptureTypeRDPEnhanced    = "rdp-enhanced"
	CaptureTypeK8sExec        = "k8s-exec"
	CaptureTypeK8sExecTTY     = "k8s-execTTY"
	CaptureTypeK8sPortForward = "k8s-portForward"
	CaptureTypeK8sCPUpload    = "k8s-cp-upload"
	CaptureTypeK8sCPDownload  = "k8s-cp-download"
	CaptureTypeK8sDescribe    = "k8s-describe"
	CaptureTypeK8sGet         = "k8s-get"
	CaptureTypeK8sDelete      = "k8s-delete"
	CaptureTypeK8sGeneric     = "k8s-generic"
	CaptureTypeK8sApply       = "k8s-apply"
	CaptureTypeSSHPortForward = "ssh-portForward"
)

// Providers responsible for device posture enforcement
const (
	DevicePostureProviderNone        = ""
	DevicePostureProviderSentinelOne = "sentinelone"
	DevicePostureProviderCrowdStrike = "crowdstrike"
)
