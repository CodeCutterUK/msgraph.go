// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TokenIssuancePolicyRequestBuilder is request builder for TokenIssuancePolicy
type TokenIssuancePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns TokenIssuancePolicyRequest
func (b *TokenIssuancePolicyRequestBuilder) Request() *TokenIssuancePolicyRequest {
	return &TokenIssuancePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TokenIssuancePolicyRequest is request for TokenIssuancePolicy
type TokenIssuancePolicyRequest struct{ BaseRequest }

// Get performs GET request for TokenIssuancePolicy
func (r *TokenIssuancePolicyRequest) Get(ctx context.Context) (resObj *TokenIssuancePolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TokenIssuancePolicy
func (r *TokenIssuancePolicyRequest) Update(ctx context.Context, reqObj *TokenIssuancePolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TokenIssuancePolicy
func (r *TokenIssuancePolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TokenLifetimePolicyRequestBuilder is request builder for TokenLifetimePolicy
type TokenLifetimePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns TokenLifetimePolicyRequest
func (b *TokenLifetimePolicyRequestBuilder) Request() *TokenLifetimePolicyRequest {
	return &TokenLifetimePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TokenLifetimePolicyRequest is request for TokenLifetimePolicy
type TokenLifetimePolicyRequest struct{ BaseRequest }

// Get performs GET request for TokenLifetimePolicy
func (r *TokenLifetimePolicyRequest) Get(ctx context.Context) (resObj *TokenLifetimePolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TokenLifetimePolicy
func (r *TokenLifetimePolicyRequest) Update(ctx context.Context, reqObj *TokenLifetimePolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TokenLifetimePolicy
func (r *TokenLifetimePolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
