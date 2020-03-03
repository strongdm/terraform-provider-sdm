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
	"context"

	plumbing "github.com/strongdm/strongdm-sdk-go/internal/v1"
)

// AccountAttachments assign an account to a role.
type AccountAttachments struct {
	client plumbing.AccountAttachmentsClient
	parent *Client
}

// Create registers a new AccountAttachment.
func (svc *AccountAttachments) Create(
	ctx context.Context,
	accountAttachment *AccountAttachment,
	options ...*AccountAttachmentCreateOptions) (
	*AccountAttachmentCreateResponse,
	error) {
	req := &plumbing.AccountAttachmentCreateRequest{}

	req.AccountAttachment = accountAttachmentToPlumbing(accountAttachment)
	if len(options) > 1 {
		return nil, &BadRequestError{Message: "use only one AccountAttachmentCreateOptions per Create call"}
	} else if len(options) == 1 {
		req.Options = accountAttachmentCreateOptionsToPlumbing(options[0])
	}
	var plumbingResponse *plumbing.AccountAttachmentCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "AccountAttachments.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountAttachmentCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.AccountAttachment = accountAttachmentToPorcelain(plumbingResponse.AccountAttachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one AccountAttachment by ID.
func (svc *AccountAttachments) Get(
	ctx context.Context,
	id string) (
	*AccountAttachmentGetResponse,
	error) {
	req := &plumbing.AccountAttachmentGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.AccountAttachmentGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "AccountAttachments.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountAttachmentGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.AccountAttachment = accountAttachmentToPorcelain(plumbingResponse.AccountAttachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a AccountAttachment by ID.
func (svc *AccountAttachments) Delete(
	ctx context.Context,
	id string) (
	*AccountAttachmentDeleteResponse,
	error) {
	req := &plumbing.AccountAttachmentDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.AccountAttachmentDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "AccountAttachments.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountAttachmentDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of AccountAttachments matching a given set of criteria.
func (svc *AccountAttachments) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	AccountAttachmentIterator,
	error) {
	req := &plumbing.AccountAttachmentListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newAccountAttachmentIteratorImpl(
		func() (
			[]*AccountAttachment,
			bool, error) {
			var plumbingResponse *plumbing.AccountAttachmentListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "AccountAttachments.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedAccountAttachmentToPorcelain(plumbingResponse.AccountAttachments)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

// AccountGrants connect a resource directly to an account, giving the account the permission to connect to that resource.
type AccountGrants struct {
	client plumbing.AccountGrantsClient
	parent *Client
}

// Create registers a new AccountGrant.
func (svc *AccountGrants) Create(
	ctx context.Context,
	accountGrant *AccountGrant) (
	*AccountGrantCreateResponse,
	error) {
	req := &plumbing.AccountGrantCreateRequest{}

	req.AccountGrant = accountGrantToPlumbing(accountGrant)
	var plumbingResponse *plumbing.AccountGrantCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "AccountGrants.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountGrantCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.AccountGrant = accountGrantToPorcelain(plumbingResponse.AccountGrant)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one AccountGrant by ID.
func (svc *AccountGrants) Get(
	ctx context.Context,
	id string) (
	*AccountGrantGetResponse,
	error) {
	req := &plumbing.AccountGrantGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.AccountGrantGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "AccountGrants.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountGrantGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.AccountGrant = accountGrantToPorcelain(plumbingResponse.AccountGrant)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a AccountGrant by ID.
func (svc *AccountGrants) Delete(
	ctx context.Context,
	id string) (
	*AccountGrantDeleteResponse,
	error) {
	req := &plumbing.AccountGrantDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.AccountGrantDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "AccountGrants.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountGrantDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of AccountGrants matching a given set of criteria.
func (svc *AccountGrants) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	AccountGrantIterator,
	error) {
	req := &plumbing.AccountGrantListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newAccountGrantIteratorImpl(
		func() (
			[]*AccountGrant,
			bool, error) {
			var plumbingResponse *plumbing.AccountGrantListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "AccountGrants.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedAccountGrantToPorcelain(plumbingResponse.AccountGrants)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

// Accounts are users that have access to strongDM.
// There are two types of accounts:
// 1. **Regular users:** humans who are authenticated through username and password or SSO
// 2. **Service users:** machines that are authneticated using a service token
type Accounts struct {
	client plumbing.AccountsClient
	parent *Client
}

// Create registers a new Account.
func (svc *Accounts) Create(
	ctx context.Context,
	account Account) (
	*AccountCreateResponse,
	error) {
	req := &plumbing.AccountCreateRequest{}

	req.Account = accountToPlumbing(account)
	var plumbingResponse *plumbing.AccountCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "Accounts.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Account = accountToPorcelain(plumbingResponse.Account)
	resp.Token = plumbingResponse.Token
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Account by ID.
func (svc *Accounts) Get(
	ctx context.Context,
	id string) (
	*AccountGetResponse,
	error) {
	req := &plumbing.AccountGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.AccountGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "Accounts.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Account = accountToPorcelain(plumbingResponse.Account)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Account by ID.
func (svc *Accounts) Update(
	ctx context.Context,
	account Account) (
	*AccountUpdateResponse,
	error) {
	req := &plumbing.AccountUpdateRequest{}

	req.Account = accountToPlumbing(account)
	var plumbingResponse *plumbing.AccountUpdateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Update(svc.parent.wrapContext(ctx, req, "Accounts.Update"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Account = accountToPorcelain(plumbingResponse.Account)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Account by ID.
func (svc *Accounts) Delete(
	ctx context.Context,
	id string) (
	*AccountDeleteResponse,
	error) {
	req := &plumbing.AccountDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.AccountDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "Accounts.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &AccountDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Accounts matching a given set of criteria.
func (svc *Accounts) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	AccountIterator,
	error) {
	req := &plumbing.AccountListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newAccountIteratorImpl(
		func() (
			[]Account,
			bool, error) {
			var plumbingResponse *plumbing.AccountListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "Accounts.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedAccountToPorcelain(plumbingResponse.Accounts)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

// Nodes make up the strongDM network, and allow your users to connect securely to your resources.
// There are two types of nodes:
// 1. **Relay:** creates connectivity to your datasources, while maintaining the egress-only nature of your firewall
// 1. **Gateways:** a relay that also listens for connections from strongDM clients
type Nodes struct {
	client plumbing.NodesClient
	parent *Client
}

// Create registers a new Node.
func (svc *Nodes) Create(
	ctx context.Context,
	node Node) (
	*NodeCreateResponse,
	error) {
	req := &plumbing.NodeCreateRequest{}

	req.Node = nodeToPlumbing(node)
	var plumbingResponse *plumbing.NodeCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "Nodes.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.Token = plumbingResponse.Token
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Node by ID.
func (svc *Nodes) Get(
	ctx context.Context,
	id string) (
	*NodeGetResponse,
	error) {
	req := &plumbing.NodeGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.NodeGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "Nodes.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Node by ID.
func (svc *Nodes) Update(
	ctx context.Context,
	node Node) (
	*NodeUpdateResponse,
	error) {
	req := &plumbing.NodeUpdateRequest{}

	req.Node = nodeToPlumbing(node)
	var plumbingResponse *plumbing.NodeUpdateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Update(svc.parent.wrapContext(ctx, req, "Nodes.Update"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Node = nodeToPorcelain(plumbingResponse.Node)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Node by ID.
func (svc *Nodes) Delete(
	ctx context.Context,
	id string) (
	*NodeDeleteResponse,
	error) {
	req := &plumbing.NodeDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.NodeDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "Nodes.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &NodeDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Nodes matching a given set of criteria.
func (svc *Nodes) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	NodeIterator,
	error) {
	req := &plumbing.NodeListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newNodeIteratorImpl(
		func() (
			[]Node,
			bool, error) {
			var plumbingResponse *plumbing.NodeListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "Nodes.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedNodeToPorcelain(plumbingResponse.Nodes)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

type Resources struct {
	client plumbing.ResourcesClient
	parent *Client
}

// Create registers a new Resource.
func (svc *Resources) Create(
	ctx context.Context,
	resource Resource) (
	*ResourceCreateResponse,
	error) {
	req := &plumbing.ResourceCreateRequest{}

	req.Resource = resourceToPlumbing(resource)
	var plumbingResponse *plumbing.ResourceCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "Resources.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Resource = resourceToPorcelain(plumbingResponse.Resource)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Resource by ID.
func (svc *Resources) Get(
	ctx context.Context,
	id string) (
	*ResourceGetResponse,
	error) {
	req := &plumbing.ResourceGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.ResourceGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "Resources.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Resource = resourceToPorcelain(plumbingResponse.Resource)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Resource by ID.
func (svc *Resources) Update(
	ctx context.Context,
	resource Resource) (
	*ResourceUpdateResponse,
	error) {
	req := &plumbing.ResourceUpdateRequest{}

	req.Resource = resourceToPlumbing(resource)
	var plumbingResponse *plumbing.ResourceUpdateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Update(svc.parent.wrapContext(ctx, req, "Resources.Update"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Resource = resourceToPorcelain(plumbingResponse.Resource)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Resource by ID.
func (svc *Resources) Delete(
	ctx context.Context,
	id string) (
	*ResourceDeleteResponse,
	error) {
	req := &plumbing.ResourceDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.ResourceDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "Resources.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &ResourceDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Resources matching a given set of criteria.
func (svc *Resources) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	ResourceIterator,
	error) {
	req := &plumbing.ResourceListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newResourceIteratorImpl(
		func() (
			[]Resource,
			bool, error) {
			var plumbingResponse *plumbing.ResourceListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "Resources.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedResourceToPorcelain(plumbingResponse.Resources)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

// RoleAttachments represent relationships between composite roles and the roles
// that make up those composite roles. When a composite role is attached to another
// role, the permissions granted to members of the composite role are augmented to
// include the permissions granted to members of the attached role.
type RoleAttachments struct {
	client plumbing.RoleAttachmentsClient
	parent *Client
}

// Create registers a new RoleAttachment.
func (svc *RoleAttachments) Create(
	ctx context.Context,
	roleAttachment *RoleAttachment) (
	*RoleAttachmentCreateResponse,
	error) {
	req := &plumbing.RoleAttachmentCreateRequest{}

	req.RoleAttachment = roleAttachmentToPlumbing(roleAttachment)
	var plumbingResponse *plumbing.RoleAttachmentCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "RoleAttachments.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleAttachmentCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RoleAttachment = roleAttachmentToPorcelain(plumbingResponse.RoleAttachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one RoleAttachment by ID.
func (svc *RoleAttachments) Get(
	ctx context.Context,
	id string) (
	*RoleAttachmentGetResponse,
	error) {
	req := &plumbing.RoleAttachmentGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleAttachmentGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "RoleAttachments.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleAttachmentGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RoleAttachment = roleAttachmentToPorcelain(plumbingResponse.RoleAttachment)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a RoleAttachment by ID.
func (svc *RoleAttachments) Delete(
	ctx context.Context,
	id string) (
	*RoleAttachmentDeleteResponse,
	error) {
	req := &plumbing.RoleAttachmentDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleAttachmentDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "RoleAttachments.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleAttachmentDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of RoleAttachments matching a given set of criteria.
func (svc *RoleAttachments) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	RoleAttachmentIterator,
	error) {
	req := &plumbing.RoleAttachmentListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newRoleAttachmentIteratorImpl(
		func() (
			[]*RoleAttachment,
			bool, error) {
			var plumbingResponse *plumbing.RoleAttachmentListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "RoleAttachments.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedRoleAttachmentToPorcelain(plumbingResponse.RoleAttachments)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

// RoleGrants represent relationships between composite roles and the roles
// that make up those composite roles. When a composite role is attached to another
// role, the permissions granted to members of the composite role are augmented to
// include the permissions granted to members of the attached role.
type RoleGrants struct {
	client plumbing.RoleGrantsClient
	parent *Client
}

// Create registers a new RoleGrant.
func (svc *RoleGrants) Create(
	ctx context.Context,
	roleGrant *RoleGrant) (
	*RoleGrantCreateResponse,
	error) {
	req := &plumbing.RoleGrantCreateRequest{}

	req.RoleGrant = roleGrantToPlumbing(roleGrant)
	var plumbingResponse *plumbing.RoleGrantCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "RoleGrants.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleGrantCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RoleGrant = roleGrantToPorcelain(plumbingResponse.RoleGrant)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one RoleGrant by ID.
func (svc *RoleGrants) Get(
	ctx context.Context,
	id string) (
	*RoleGrantGetResponse,
	error) {
	req := &plumbing.RoleGrantGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleGrantGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "RoleGrants.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleGrantGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RoleGrant = roleGrantToPorcelain(plumbingResponse.RoleGrant)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a RoleGrant by ID.
func (svc *RoleGrants) Delete(
	ctx context.Context,
	id string) (
	*RoleGrantDeleteResponse,
	error) {
	req := &plumbing.RoleGrantDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleGrantDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "RoleGrants.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleGrantDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of RoleGrants matching a given set of criteria.
func (svc *RoleGrants) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	RoleGrantIterator,
	error) {
	req := &plumbing.RoleGrantListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newRoleGrantIteratorImpl(
		func() (
			[]*RoleGrant,
			bool, error) {
			var plumbingResponse *plumbing.RoleGrantListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "RoleGrants.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedRoleGrantToPorcelain(plumbingResponse.RoleGrants)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}

// Roles are tools for controlling user access to resources. Each Role holds a
// list of resources which they grant access to. Composite roles are a special
// type of Role which have no resource associations of their own, but instead
// grant access to the combined resources associated with a set of child roles.
// Each user can be a member of one Role or composite role.
type Roles struct {
	client plumbing.RolesClient
	parent *Client
}

// Create registers a new Role.
func (svc *Roles) Create(
	ctx context.Context,
	role *Role) (
	*RoleCreateResponse,
	error) {
	req := &plumbing.RoleCreateRequest{}

	req.Role = roleToPlumbing(role)
	var plumbingResponse *plumbing.RoleCreateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Create(svc.parent.wrapContext(ctx, req, "Roles.Create"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleCreateResponse{}
	resp.Meta = createResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Get reads one Role by ID.
func (svc *Roles) Get(
	ctx context.Context,
	id string) (
	*RoleGetResponse,
	error) {
	req := &plumbing.RoleGetRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleGetResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Get(svc.parent.wrapContext(ctx, req, "Roles.Get"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleGetResponse{}
	resp.Meta = getResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Update patches a Role by ID.
func (svc *Roles) Update(
	ctx context.Context,
	role *Role) (
	*RoleUpdateResponse,
	error) {
	req := &plumbing.RoleUpdateRequest{}

	req.Role = roleToPlumbing(role)
	var plumbingResponse *plumbing.RoleUpdateResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Update(svc.parent.wrapContext(ctx, req, "Roles.Update"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleUpdateResponse{}
	resp.Meta = updateResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.Role = roleToPorcelain(plumbingResponse.Role)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// Delete removes a Role by ID.
func (svc *Roles) Delete(
	ctx context.Context,
	id string) (
	*RoleDeleteResponse,
	error) {
	req := &plumbing.RoleDeleteRequest{}

	req.Id = id
	var plumbingResponse *plumbing.RoleDeleteResponse
	var err error
	i := 0
	for {
		plumbingResponse, err = svc.client.Delete(svc.parent.wrapContext(ctx, req, "Roles.Delete"), req)
		if err != nil {
			if !svc.parent.shouldRetry(i, err) {
				return nil, errorToPorcelain(err)
			}
			i++
			svc.parent.jitterSleep(i)
			continue
		}
		break
	}

	resp := &RoleDeleteResponse{}
	resp.Meta = deleteResponseMetadataToPorcelain(plumbingResponse.Meta)
	resp.RateLimit = rateLimitMetadataToPorcelain(plumbingResponse.RateLimit)
	return resp, nil
}

// List gets a list of Roles matching a given set of criteria.
func (svc *Roles) List(
	ctx context.Context,
	filter string,
	args ...interface{}) (
	RoleIterator,
	error) {
	req := &plumbing.RoleListRequest{}

	var filterErr error
	req.Filter, filterErr = quoteFilterArgs(filter, args...)
	if filterErr != nil {
		return nil, filterErr
	}
	req.Meta = &plumbing.ListRequestMetadata{}
	if value := svc.parent.testOption("PageLimit"); value != nil {
		v, ok := value.(int)
		if ok {
			req.Meta.Limit = int32(v)
		}
	}
	return newRoleIteratorImpl(
		func() (
			[]*Role,
			bool, error) {
			var plumbingResponse *plumbing.RoleListResponse
			var err error
			i := 0
			for {
				plumbingResponse, err = svc.client.List(svc.parent.wrapContext(ctx, req, "Roles.List"), req)
				if err != nil {
					if !svc.parent.shouldRetry(i, err) {
						return nil, false, errorToPorcelain(err)
					}
					i++
					svc.parent.jitterSleep(i)
					continue
				}
				break
			}
			result := repeatedRoleToPorcelain(plumbingResponse.Roles)
			req.Meta.Cursor = plumbingResponse.Meta.NextCursor
			return result, req.Meta.Cursor != "", nil
		},
	), nil
}
