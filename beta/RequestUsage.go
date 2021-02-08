// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// UsageRightRequestBuilder is request builder for UsageRight
type UsageRightRequestBuilder struct{ BaseRequestBuilder }

// Request returns UsageRightRequest
func (b *UsageRightRequestBuilder) Request() *UsageRightRequest {
	return &UsageRightRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UsageRightRequest is request for UsageRight
type UsageRightRequest struct{ BaseRequest }

// Get performs GET request for UsageRight
func (r *UsageRightRequest) Get(ctx context.Context) (resObj *UsageRight, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UsageRight
func (r *UsageRightRequest) Update(ctx context.Context, reqObj *UsageRight) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UsageRight
func (r *UsageRightRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
