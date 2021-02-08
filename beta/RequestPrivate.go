// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PrivateLinkResourcePolicyRequestBuilder is request builder for PrivateLinkResourcePolicy
type PrivateLinkResourcePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivateLinkResourcePolicyRequest
func (b *PrivateLinkResourcePolicyRequestBuilder) Request() *PrivateLinkResourcePolicyRequest {
	return &PrivateLinkResourcePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivateLinkResourcePolicyRequest is request for PrivateLinkResourcePolicy
type PrivateLinkResourcePolicyRequest struct{ BaseRequest }

// Get performs GET request for PrivateLinkResourcePolicy
func (r *PrivateLinkResourcePolicyRequest) Get(ctx context.Context) (resObj *PrivateLinkResourcePolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivateLinkResourcePolicy
func (r *PrivateLinkResourcePolicyRequest) Update(ctx context.Context, reqObj *PrivateLinkResourcePolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivateLinkResourcePolicy
func (r *PrivateLinkResourcePolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
