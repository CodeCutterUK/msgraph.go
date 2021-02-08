// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// LinkedResourceRequestBuilder is request builder for LinkedResource
type LinkedResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns LinkedResourceRequest
func (b *LinkedResourceRequestBuilder) Request() *LinkedResourceRequest {
	return &LinkedResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LinkedResourceRequest is request for LinkedResource
type LinkedResourceRequest struct{ BaseRequest }

// Get performs GET request for LinkedResource
func (r *LinkedResourceRequest) Get(ctx context.Context) (resObj *LinkedResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LinkedResource
func (r *LinkedResourceRequest) Update(ctx context.Context, reqObj *LinkedResource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LinkedResource
func (r *LinkedResourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}