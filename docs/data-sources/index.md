
# Data Sources
* [account_attachment](./account_attachment) - AccountAttachments assign an account to a role or composite role.
* [account_grant](./account_grant) - AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
* [account](./account) - Accounts are users that have access to strongDM. There are two types of accounts:
 1. **Users:** humans who are authenticated through username and password or SSO.
 2. **Service Accounts:** machines that are authenticated using a service token.
* [node](./node) - Nodes make up the strongDM network, and allow your users to connect securely to your resources.
 There are two types of nodes:
 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
 1. **Gateways:** a relay that also listens for connections from strongDM clients
* [resource](./resource) - A Resource is a database or server for which strongDM manages access.
* [role_attachment](./role_attachment) - A RoleAttachment assigns a role to a composite role.
* [role_grant](./role_grant) - A RoleGrant connects a resource to a role, granting members of the role access to that resource.
* [role](./role) - A Role is a collection of access grants, and typically corresponds to a team, Active Directory OU, or other organizational unit. Users are granted access to resources by assigning them to roles.
* [secret_store](./secret_store) - A SecretStore is a beta feature. 
