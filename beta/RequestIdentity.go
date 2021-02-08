// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// IdentityAPIConnectorRequestBuilder is request builder for IdentityAPIConnector
type IdentityAPIConnectorRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityAPIConnectorRequest
func (b *IdentityAPIConnectorRequestBuilder) Request() *IdentityAPIConnectorRequest {
	return &IdentityAPIConnectorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityAPIConnectorRequest is request for IdentityAPIConnector
type IdentityAPIConnectorRequest struct{ BaseRequest }

// Get performs GET request for IdentityAPIConnector
func (r *IdentityAPIConnectorRequest) Get(ctx context.Context) (resObj *IdentityAPIConnector, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityAPIConnector
func (r *IdentityAPIConnectorRequest) Update(ctx context.Context, reqObj *IdentityAPIConnector) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityAPIConnector
func (r *IdentityAPIConnectorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityContainerRequestBuilder is request builder for IdentityContainer
type IdentityContainerRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityContainerRequest
func (b *IdentityContainerRequestBuilder) Request() *IdentityContainerRequest {
	return &IdentityContainerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityContainerRequest is request for IdentityContainer
type IdentityContainerRequest struct{ BaseRequest }

// Get performs GET request for IdentityContainer
func (r *IdentityContainerRequest) Get(ctx context.Context) (resObj *IdentityContainer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityContainer
func (r *IdentityContainerRequest) Update(ctx context.Context, reqObj *IdentityContainer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityContainer
func (r *IdentityContainerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityGovernanceRequestBuilder is request builder for IdentityGovernance
type IdentityGovernanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityGovernanceRequest
func (b *IdentityGovernanceRequestBuilder) Request() *IdentityGovernanceRequest {
	return &IdentityGovernanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityGovernanceRequest is request for IdentityGovernance
type IdentityGovernanceRequest struct{ BaseRequest }

// Get performs GET request for IdentityGovernance
func (r *IdentityGovernanceRequest) Get(ctx context.Context) (resObj *IdentityGovernance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityGovernance
func (r *IdentityGovernanceRequest) Update(ctx context.Context, reqObj *IdentityGovernance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityGovernance
func (r *IdentityGovernanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityProtectionRootRequestBuilder is request builder for IdentityProtectionRoot
type IdentityProtectionRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityProtectionRootRequest
func (b *IdentityProtectionRootRequestBuilder) Request() *IdentityProtectionRootRequest {
	return &IdentityProtectionRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityProtectionRootRequest is request for IdentityProtectionRoot
type IdentityProtectionRootRequest struct{ BaseRequest }

// Get performs GET request for IdentityProtectionRoot
func (r *IdentityProtectionRootRequest) Get(ctx context.Context) (resObj *IdentityProtectionRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityProtectionRoot
func (r *IdentityProtectionRootRequest) Update(ctx context.Context, reqObj *IdentityProtectionRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityProtectionRoot
func (r *IdentityProtectionRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityProviderRequestBuilder is request builder for IdentityProvider
type IdentityProviderRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityProviderRequest
func (b *IdentityProviderRequestBuilder) Request() *IdentityProviderRequest {
	return &IdentityProviderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityProviderRequest is request for IdentityProvider
type IdentityProviderRequest struct{ BaseRequest }

// Get performs GET request for IdentityProvider
func (r *IdentityProviderRequest) Get(ctx context.Context) (resObj *IdentityProvider, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityProvider
func (r *IdentityProviderRequest) Update(ctx context.Context, reqObj *IdentityProvider) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityProvider
func (r *IdentityProviderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentitySecurityDefaultsEnforcementPolicyRequestBuilder is request builder for IdentitySecurityDefaultsEnforcementPolicy
type IdentitySecurityDefaultsEnforcementPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentitySecurityDefaultsEnforcementPolicyRequest
func (b *IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) Request() *IdentitySecurityDefaultsEnforcementPolicyRequest {
	return &IdentitySecurityDefaultsEnforcementPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentitySecurityDefaultsEnforcementPolicyRequest is request for IdentitySecurityDefaultsEnforcementPolicy
type IdentitySecurityDefaultsEnforcementPolicyRequest struct{ BaseRequest }

// Get performs GET request for IdentitySecurityDefaultsEnforcementPolicy
func (r *IdentitySecurityDefaultsEnforcementPolicyRequest) Get(ctx context.Context) (resObj *IdentitySecurityDefaultsEnforcementPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentitySecurityDefaultsEnforcementPolicy
func (r *IdentitySecurityDefaultsEnforcementPolicyRequest) Update(ctx context.Context, reqObj *IdentitySecurityDefaultsEnforcementPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentitySecurityDefaultsEnforcementPolicy
func (r *IdentitySecurityDefaultsEnforcementPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityUserFlowRequestBuilder is request builder for IdentityUserFlow
type IdentityUserFlowRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityUserFlowRequest
func (b *IdentityUserFlowRequestBuilder) Request() *IdentityUserFlowRequest {
	return &IdentityUserFlowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityUserFlowRequest is request for IdentityUserFlow
type IdentityUserFlowRequest struct{ BaseRequest }

// Get performs GET request for IdentityUserFlow
func (r *IdentityUserFlowRequest) Get(ctx context.Context) (resObj *IdentityUserFlow, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityUserFlow
func (r *IdentityUserFlowRequest) Update(ctx context.Context, reqObj *IdentityUserFlow) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityUserFlow
func (r *IdentityUserFlowRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityUserFlowAttributeRequestBuilder is request builder for IdentityUserFlowAttribute
type IdentityUserFlowAttributeRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityUserFlowAttributeRequest
func (b *IdentityUserFlowAttributeRequestBuilder) Request() *IdentityUserFlowAttributeRequest {
	return &IdentityUserFlowAttributeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityUserFlowAttributeRequest is request for IdentityUserFlowAttribute
type IdentityUserFlowAttributeRequest struct{ BaseRequest }

// Get performs GET request for IdentityUserFlowAttribute
func (r *IdentityUserFlowAttributeRequest) Get(ctx context.Context) (resObj *IdentityUserFlowAttribute, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityUserFlowAttribute
func (r *IdentityUserFlowAttributeRequest) Update(ctx context.Context, reqObj *IdentityUserFlowAttribute) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityUserFlowAttribute
func (r *IdentityUserFlowAttributeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityUserFlowAttributeAssignmentRequestBuilder is request builder for IdentityUserFlowAttributeAssignment
type IdentityUserFlowAttributeAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityUserFlowAttributeAssignmentRequest
func (b *IdentityUserFlowAttributeAssignmentRequestBuilder) Request() *IdentityUserFlowAttributeAssignmentRequest {
	return &IdentityUserFlowAttributeAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityUserFlowAttributeAssignmentRequest is request for IdentityUserFlowAttributeAssignment
type IdentityUserFlowAttributeAssignmentRequest struct{ BaseRequest }

// Get performs GET request for IdentityUserFlowAttributeAssignment
func (r *IdentityUserFlowAttributeAssignmentRequest) Get(ctx context.Context) (resObj *IdentityUserFlowAttributeAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityUserFlowAttributeAssignment
func (r *IdentityUserFlowAttributeAssignmentRequest) Update(ctx context.Context, reqObj *IdentityUserFlowAttributeAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityUserFlowAttributeAssignment
func (r *IdentityUserFlowAttributeAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestBuilder struct{ BaseRequestBuilder }

// SetOrder action undocumented
func (b *B2cIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder) SetOrder(reqObj *IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestParameter) *IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestBuilder {
	bb := &IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setOrder"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// SetOrder action undocumented
func (b *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder) SetOrder(reqObj *IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestParameter) *IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestBuilder {
	bb := &IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setOrder"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IdentityUserFlowAttributeAssignmentCollectionSetOrderRequest struct{ BaseRequest }

//
func (b *IdentityUserFlowAttributeAssignmentCollectionSetOrderRequestBuilder) Request() *IdentityUserFlowAttributeAssignmentCollectionSetOrderRequest {
	return &IdentityUserFlowAttributeAssignmentCollectionSetOrderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IdentityUserFlowAttributeAssignmentCollectionSetOrderRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
